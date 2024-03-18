package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"strings"
	api_models "vk_test_task/internal/api/models"
	"vk_test_task/internal/common"
)

func (r Repository) CreateFilm(params api_models.CreateFilmParams) error {
	if params.FilmId == "" {
		return fmt.Errorf("repository error: invalid film id")
	}

	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("repository error: transaction error: %s", err.Error())
	}
	defer tx.Rollback()

	query := `insert into film(name, description, date_released, rate, id) values ($1, $2, $3, $4, $5)`

	_, err = tx.Exec(query, params.Name, params.Description, params.ReleaseDate, params.Rate, params.FilmId)

	if err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	if len(params.Actors) == 0 {
		if err = tx.Commit(); err != nil {
			return fmt.Errorf("repository error: transaction error: %s", err.Error())
		}
		return nil
	}

	relationQuery := `insert into film_actor(film_id, actor_id) values`

	for i := 0; i < len(params.Actors); i++ {
		relationQuery += fmt.Sprintf(` ($1, $%d),`, i+2)
	}
	if len(params.Actors) > 0 {
		relationQuery = strings.TrimSuffix(relationQuery, ",")
	}
	args := []interface{}{params.FilmId}
	for _, v := range params.Actors {
		args = append(args, v)
	}
	_, err = tx.Exec(relationQuery, args...)

	if err != nil {
		fmt.Println(relationQuery, append([]string{params.FilmId}, params.Actors...))
		return fmt.Errorf("repository error: %s", err.Error())
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repository error: transaction error: %s", err.Error())
	}

	return nil
}

func (r Repository) GetFilms(params api_models.GetFilmsParams) (api_models.GetFilmsResponse, error) {
	var queryAscending, querySortBy = "desc", "rate"
	if params.IsAscending == common.SORT_FILM_ASC {
		queryAscending = ""
	}

	switch params.SortBy {
	case common.SORT_FILM_BY_NAME:
		querySortBy = "name"
	case common.SORT_FILM_BY_RELEASE_DATE:
		querySortBy = "date_released"
	}

	query := fmt.Sprintf(`select film.*, array_agg(actor.name) as actors
	from film
	left join film_actor on film.id = film_actor.film_id
	left join actor on actor.id = film_actor.actor_id
	group by film.id, film.%s
    order by film.%s %s`, querySortBy, querySortBy, queryAscending)

	var response api_models.GetFilmsResponse

	rows, err := r.db.Query(query)
	if err != nil {
		return api_models.GetFilmsResponse{}, fmt.Errorf("repository error: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var filmAndActors api_models.FilmAndActors

		err = rows.Scan(&filmAndActors.Name, &filmAndActors.Description,
			&filmAndActors.ReleaseDate, &filmAndActors.Rate, &filmAndActors.FilmId,
			&filmAndActors.Actors)

		if err != nil {
			return api_models.GetFilmsResponse{}, fmt.Errorf("repository error: %s", err.Error())
		}

		response.Response = append(response.Response, filmAndActors)
	}

	return response, nil
}

func (r Repository) UpdateFilm(params api_models.UpdateFilmParams) error {
	if params.FilmId == "" {
		return fmt.Errorf("repository error: invalid filmId")
	}

	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("repository error: transaction error: %s", err.Error())
	}
	defer tx.Rollback()

	var name, description, releaseDate, rate = "name", "description", "date_released", "rate"
	if params.Name != "" {
		name = "@name"
	}
	if params.Description != "" {
		description = "@description"
	}
	if !params.ReleaseDate.IsZero() {
		releaseDate = "@date_released"
	}
	if params.Rate == 0 {
		rate = "@rate"
	}

	query := fmt.Sprintf(`update film set 
                name = %s, description = %s, date_released = %s, rate = %s where id = @id`,
		name, description, releaseDate, rate)

	args := pgx.NamedArgs{
		"name":          params.Name,
		"description":   params.Description,
		"date_released": params.ReleaseDate,
		"id":            params.FilmId,
	}

	_, err = tx.Exec(query, args)
	if err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	if len(params.Actors) == 0 {
		if err = tx.Commit(); err != nil {
			return fmt.Errorf("repository error: transaction error: %s", err.Error())
		}
		return nil
	}

	relationDeleteQuery := `delete from film_actor where film_id = $1`
	_, err = tx.Exec(relationDeleteQuery, params.FilmId)
	if err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repository error: transaction error: %s", err.Error())
	}

	return nil
}

func (r Repository) DeleteFilm(filmId string) error {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("repository error: transaction error: %s", err.Error())
	}
	defer tx.Rollback()

	relationQuery := `delete from film_actor where film_id = $1`

	_, err = tx.Exec(relationQuery, filmId)
	if err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	query := `delete from film where id = $1`

	_, err = tx.Exec(query, filmId)
	if err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	return nil
}

func (r Repository) SearchFilmByName(name string) (api_models.SearchFilmResponse, error) {
	if name == "" {
		return api_models.SearchFilmResponse{}, fmt.Errorf("repository error: invalid name")
	}

	query := fmt.Sprintf(`select film.*, array_agg(actor.name) as actors
	from film
	left join film_actor on film.id = film_actor.film_id
	left join actor on actor.id = film_actor.actor_id
	where film.name ilike $1
	group by film.id
	`)

	var response api_models.SearchFilmResponse

	regex := fmt.Sprintf("%%%s%%", name)

	rows, err := r.db.Query(query, regex)
	if err != nil {
		return api_models.SearchFilmResponse{}, fmt.Errorf("repository error: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var filmAndActors api_models.FilmAndActors

		err = rows.Scan(&filmAndActors.Name, &filmAndActors.Description,
			&filmAndActors.ReleaseDate, &filmAndActors.Rate,
			&filmAndActors.FilmId, &filmAndActors.Actors)

		if err != nil {
			return api_models.SearchFilmResponse{}, fmt.Errorf("repository error: %s", err.Error())
		}

		response.Response = append(response.Response, filmAndActors)
	}

	return response, nil
}

func (r Repository) SearchFilmByActorName(actorName string) (api_models.SearchFilmResponse, error) {
	if actorName == "" {
		return api_models.SearchFilmResponse{}, fmt.Errorf("repository error: invalid name")
	}

	query := fmt.Sprintf(`with films as (select film.*, array_agg(actor.name) as actors
	from film
	left join film_actor on film.id = film_actor.film_id
	left join actor on actor.id = film_actor.actor_id
	group by film.id)

	select films.*
	from films
	where array_to_string(actors, ' ') ilike $1`)

	var response api_models.SearchFilmResponse

	regex := fmt.Sprintf("%%%s%%", actorName)

	rows, err := r.db.Query(query, regex)
	if err != nil {
		return api_models.SearchFilmResponse{}, fmt.Errorf("repository error: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var filmAndActors api_models.FilmAndActors

		err = rows.Scan(&filmAndActors.Name, &filmAndActors.Description,
			&filmAndActors.ReleaseDate, &filmAndActors.Rate,
			&filmAndActors.FilmId, &filmAndActors.Actors)

		if err != nil {
			return api_models.SearchFilmResponse{}, fmt.Errorf("repository error: %s", err.Error())
		}

		response.Response = append(response.Response, filmAndActors)
	}

	return response, nil
}
