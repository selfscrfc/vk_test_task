package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	api_models "vk_test_task/internal/api/models"
	"vk_test_task/internal/common"
)

func (r Repository) CreateActor(params api_models.CreateActorParams) error {
	if params.ActorId == "" {
		return errors.New("invalid actorId")
	}
	if len(params.Name) > common.ACTOR_NAME_MAXSIZE {
		return errors.New("name is too long")
	}
	if params.Sex != common.ACTOR_SEX_MALE &&
		params.Sex != common.ACTOR_SEX_FEMALE {
		return errors.New("invalid sex")
	}

	query := `insert into actor(id, name, sex, birth) values($1, $2, $3, $4)`

	_, err := r.db.Exec(query, params.ActorId, params.Name, params.Sex, params.Birth)

	if err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}
	return nil
}

func (r Repository) GetActors() (api_models.GetActorsResponse, error) {
	query := `select actor.*, array_agg(film.name) as films
	from actor
	left join film_actor on actor.id = film_actor.actor_id
	left join film on film_actor.film_id = film.id
	group by actor.id`

	var response api_models.GetActorsResponse

	rows, err := r.db.Query(query)
	if err != nil {
		return api_models.GetActorsResponse{}, fmt.Errorf("repository error: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var actorAndFilms api_models.ActorAndFilms

		err = rows.Scan(&actorAndFilms.Name, &actorAndFilms.Sex,
			&actorAndFilms.Birth, &actorAndFilms.ActorId,
			&actorAndFilms.Films)

		if err != nil {
			return api_models.GetActorsResponse{}, fmt.Errorf("repository error: %s", err.Error())
		}

		response.Response = append(response.Response, actorAndFilms)
	}

	return response, nil
}

func (r Repository) UpdateActor(params api_models.UpdateActorParams) error {
	if params.ActorId == "" {
		return fmt.Errorf("repository error: invalid actor id")
	}
	var queryName, querySex, queryBirth = "name", "sex", "birth"
	if params.Name != "" {
		queryName = "@name"
	}
	if params.Sex != 0 {
		querySex = "@sex"
	}
	if !params.Birth.IsZero() {
		queryBirth = "@birth"
	}

	query := fmt.Sprintf(`update actor set name = %s, sex = %s, birth = %s where id = @id`,
		queryName, querySex, queryBirth)

	args := pgx.NamedArgs{
		"name":  params.Name,
		"sex":   params.Sex,
		"birth": params.Birth,
		"id":    params.ActorId,
	}

	_, err := r.db.Exec(query, args)

	if err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	return nil
}

func (r Repository) DeleteActor(actorId string) error {
	if actorId == "" {
		return fmt.Errorf("repository err: invalid actor id")
	}
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("repository error: transaction error: %s", err.Error())
	}
	defer tx.Rollback()

	relationQuery := `delete from film_actor where actor_id = $1`

	_, err = tx.Exec(relationQuery, actorId)
	if err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	query := `delete from actor where id = $1`

	_, err = tx.Exec(query, actorId)
	if err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	return nil
}
