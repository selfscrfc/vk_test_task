package api_delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vk_test_task/internal/api/models"
)

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
