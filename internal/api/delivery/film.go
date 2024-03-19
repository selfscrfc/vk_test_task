package api_delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"vk_test_task/internal/api/models"
	"vk_test_task/internal/common"
)

// CreateFilm godoc
// @Summary CreateFilm
// @Description creates film instance and returns its uuid. Release date in ISO format (2009-05-27T00:00:00.000Z)
// @Tags Film
// @Param input body api_models.CreateFilmParams true "film info"
// @Accept json
// @Produce json
// @Success 200 {object} api_models.CreateFilmParams
// @Router /film/create [post]
// @Security AccessTokenAuth
func (h Handler) CreateFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api_models.CreateFilmParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			errText := fmt.Sprintf("create film error: %s", err.Error())
			h.logger.Error(errText)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		h.logger.Info(fmt.Sprintf("/film/create request. Params: %v", params))

		params.FilmId, err = h.uc.CreateFilm(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("create film error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
		paramsJson, err := json.Marshal(params)
		w.Write(paramsJson)
	}
}

// GetFilms godoc
// @Summary GetFilms
// @Description return all films with their actors
// @Tags Film
// @Param sort_by query string false "sort column"
// @Param asc query string false "sort asc"
// @Produce json
// @Success 200
// @Router /film/get [get]
// @Security AccessTokenAuth
func (h Handler) GetFilms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var params api_models.GetFilmsParams

		re, _ := regexp.Compile(`/film/get\?sort_by=(\d)&asc=(\d)`)
		matches := re.FindStringSubmatch(r.URL.String())
		if len(matches) == 3 {
			sortByString := matches[1]
			isAscendingString := matches[2]

			params.SortBy, _ = strconv.Atoi(sortByString)
			params.IsAscending, _ = strconv.Atoi(isAscendingString)
		} else {
			params.SortBy = common.SORT_FILM_BY_RATE
			params.IsAscending = common.SORT_FILM_DESC
		}

		h.logger.Info(fmt.Sprintf("/film/get request. Params: %v", params))

		response, err := h.uc.GetFilms(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("get films error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		jsonResponse, err := response.MarshallJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("get films error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

// UpdateFilm godoc
// @Summary UpdateFilm
// @Description updates film info. Release date in ISO format (2009-05-27T00:00:00.000Z)
// @Tags Film
// @Param input body api_models.UpdateFilmParams true "film info"
// @Accept json
// @Success 200
// @Router /film/update [post]
// @Security AccessTokenAuth
func (h Handler) UpdateFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api_models.UpdateFilmParams

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errText := fmt.Sprintf("/film/update error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		h.logger.Info(fmt.Sprintf("/film/update request. Params: %v", params))

		err = h.uc.UpdateFilm(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("/film/update error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// DeleteFilm godoc
// @Summary DeleteFilm
// @Description deletes film by its filmId
// @Tags Film
// @Param input body api_models.DeleteFilmParams true "filmId"
// @Accept json
// @Success 200
// @Router /film/delete [post]
// @Security AccessTokenAuth
func (h Handler) DeleteFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api_models.DeleteFilmParams

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errText := fmt.Sprintf("/film/delete error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		h.logger.Info(fmt.Sprintf("/film/delete request. Params: %v", params))

		err = h.uc.DeleteFilm(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("/film/delete error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// SearchFilm godoc
// @Summary SearchFilm
// @Description accepts path parameters, name prioritized. Defaults: rate, desc
// @Tags Film
// @Param name query string false "film name fragment"
// @Param actor_name query string false "actor name fragment"
// @Produce json
// @Success 200
// @Router /film/search [get]
// @Security AccessTokenAuth
func (h Handler) SearchFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api_models.SearchFilmParams

		flag := true
		re, _ := regexp.Compile(`/film/search\?name=(\w+)`)
		matches := re.FindStringSubmatch(r.URL.String())
		if len(matches) == 2 {
			params.Name = matches[1]
			flag = false
		}

		if flag {
			re, _ = regexp.Compile(`/film/search\?actor_name=(\w+)`)
			matches = re.FindStringSubmatch(r.URL.String())
			if len(matches) == 2 {
				params.ActorName = matches[1]
				flag = false
			}
		}

		if flag {
			w.WriteHeader(http.StatusBadRequest)
			errText := fmt.Sprintf("/film/search error: empty params")
			h.logger.Error(errText)
			return
		}

		h.logger.Info(fmt.Sprintf("/film/search request. Params: %v", params))

		response, err := h.uc.SearchFilm(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("/film/search error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		jsonResponse, err := response.MarshallJSON()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("search films error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
