package api_delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vk_test_task/internal/api/models"
)

func (h Handler) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api_models.AuthParams

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errText := fmt.Sprintf("sign in error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		h.logger.Info(fmt.Sprintf("/sign_in request."))

		resp, err := h.uc.SignIn(params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorResponse, _ := json.Marshal(api_models.ErrorResponse{Description: err.Error()})
			w.Write(errorResponse)
			errText := fmt.Sprintf("sign in error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResponse, err := json.Marshal(resp)
		_, err = w.Write(jsonResponse)
		if err != nil {
			errText := fmt.Sprintf("sign in error: %s", err.Error())
			h.logger.Error(errText)
			return
		}
	}
}

func (h Handler) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api_models.AuthParams

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errText := fmt.Sprintf("sign in error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		h.logger.Info(fmt.Sprintf("/sign_up request."))

		err = h.uc.SignUp(params)
		if err != nil {
			errorResponse, _ := json.Marshal(api_models.ErrorResponse{Description: err.Error()})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorResponse)
			errText := fmt.Sprintf("sign in error: %s", err.Error())
			h.logger.Error(errText)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}
}
