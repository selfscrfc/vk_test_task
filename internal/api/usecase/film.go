package api_usecase

import (
	"fmt"
	"github.com/google/uuid"
	api_models "vk_test_task/internal/api/models"
	"vk_test_task/internal/common"
)

func (u UseCase) CreateFilm(params api_models.CreateFilmParams) (string, error) {
	if len(params.Name) > common.FILM_NAME_MAXSIZE ||
		len(params.Name) < common.FILM_NAME_MINSIZE {
		return "", fmt.Errorf("usecase error: invalid film name")
	}

	if len(params.Description) > common.FILM_DESCRIPTION_MAXSIZE {
		return "", fmt.Errorf("usecase error: description is too long")
	}

	if params.Rate < 0 || params.Rate > 10 {
		return "", fmt.Errorf("usecase error: invalid film rate")
	}

	filmId, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("usecase error: %s", err.Error())
	}
	params.FilmId = filmId.String()

	err = u.db.CreateFilm(params)
	if err != nil {
		return "", fmt.Errorf("usecase error: %s", err.Error())
	}

	return params.FilmId, nil
}

func (u UseCase) GetFilms(params api_models.GetFilmsParams) (api_models.GetFilmsResponse, error) {
	if params.IsAscending != common.SORT_FILM_ASC &&
		params.IsAscending != common.SORT_FILM_DESC {
		return api_models.GetFilmsResponse{}, fmt.Errorf("usecase error: invalid sort asc parameter")
	}
	if params.SortBy != common.SORT_FILM_BY_RATE &&
		params.SortBy != common.SORT_FILM_BY_NAME &&
		params.SortBy != common.SORT_FILM_BY_RELEASE_DATE {
		return api_models.GetFilmsResponse{}, fmt.Errorf("usecase error: invalid sort by parameter")
	}

	response, err := u.db.GetFilms(params)
	if err != nil {
		return api_models.GetFilmsResponse{}, fmt.Errorf("usecase error: %s", err.Error())
	}
	return response, nil
}

func (u UseCase) UpdateFilm(params api_models.UpdateFilmParams) error {
	if params.FilmId == "" {
		return fmt.Errorf("usecase error: invalid id")
	}
	if len(params.Name) > common.FILM_NAME_MAXSIZE {
		return fmt.Errorf("usecase error: invalid film name")
	}
	if len(params.Description) > common.FILM_DESCRIPTION_MAXSIZE {
		return fmt.Errorf("usecase error: film description is too long")
	}
	if params.Rate < 0 || params.Rate > 10 {
		return fmt.Errorf("usecase error: invalid film rate")
	}

	err := u.db.UpdateFilm(params)
	if err != nil {
		return fmt.Errorf("usecase error: %s", err.Error())
	}

	return nil
}

func (u UseCase) DeleteFilm(params api_models.DeleteFilmParams) error {
	if params.FilmId == "" {
		return fmt.Errorf("usecase error: invalid film id")
	}

	err := u.db.DeleteFilm(params.FilmId)
	if err != nil {
		return fmt.Errorf("usecase error: %s", err.Error())
	}

	return nil
}

func (u UseCase) SearchFilm(params api_models.SearchFilmParams) (api_models.SearchFilmResponse, error) {
	if params.ActorName == "" && params.Name == "" {
		return api_models.SearchFilmResponse{}, fmt.Errorf("usecase error: invalid params")
	}

	var response api_models.SearchFilmResponse
	var err error

	if name := params.Name; name != "" {
		response, err = u.db.SearchFilmByName(name)
	} else {
		response, err = u.db.SearchFilmByActorName(params.ActorName)
	}

	if err != nil {
		return api_models.SearchFilmResponse{}, fmt.Errorf("usecase error: %s", err.Error())
	}

	return response, nil
}
