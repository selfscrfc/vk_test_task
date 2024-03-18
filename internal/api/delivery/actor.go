package api_delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vk_test_task/internal/api/models"
)

// CreateActor godoc
// @Summary CreateActor
// @Description creates actor instance and returns its uuid. Birth in ISO format (2009-05-27T00:00:00.000Z)
// @Tags Actor
// @Param input body api_models.CreateActorParams true "actor info"
// @Accept json
// @Produce json
// @Success 200 {object} api_models.CreateActorParams
// @Router /actor/create [post]
// @Security AccessTokenAuth
func (h Handler) CreateActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api_models.CreateActorParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			errText := fmt.Sprintf("create actor error: %s", err.Error())
			h.logger.Error(errText)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		h.logger.Info(fmt.Sprintf("/actore/create request. Params: %v", params))

		params.ActorId, err = h.uc.CreateActor(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("create actor error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
		paramsJson, err := json.Marshal(params)
		w.Write(paramsJson)
	}
}

// GetActors godoc
// @Summary GetActors
// @Description return all actors with their films
// @Tags Actor
// @Produce json
// @Success 200
// @Router /actor/get [get]
// @Security AccessTokenAuth
func (h Handler) GetActors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.logger.Info(fmt.Sprintf("/actor/get request."))
		response, err := h.uc.GetActors()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("get actors error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		jsonResponse, err := response.MarshallJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("get actors error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

// UpdateActor godoc
// @Summary UpdateActor
// @Description updates actor info. Birth in ISO format (2009-05-27T00:00:00.000Z)
// @Tags Actor
// @Param input body api_models.UpdateActorParams true "actor info"
// @Accept json
// @Success 200
// @Router /actor/update [post]
// @Security AccessTokenAuth
func (h Handler) UpdateActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api_models.UpdateActorParams

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errText := fmt.Sprintf("/actor/update error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		h.logger.Info(fmt.Sprintf("/actor/update request. Params: %v", params))

		err = h.uc.UpdateActor(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("/actor/update error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// DeleteActor godoc
// @Summary DeleteActor
// @Description deletes actor by its actorId
// @Tags Actor
// @Param input body api_models.DeleteActorParams true "actorId"
// @Accept json
// @Success 200
// @Router /actor/delete [post]
// @Security AccessTokenAuth
func (h Handler) DeleteActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api_models.DeleteActorParams

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errText := fmt.Sprintf("/actor/delete error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		h.logger.Info(fmt.Sprintf("/actor/delete request. Params: %v", params))

		err = h.uc.DeleteActor(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errText := fmt.Sprintf("/actor/delete error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
