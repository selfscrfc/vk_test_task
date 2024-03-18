package api_delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/lmittmann/tint"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	mock_api "vk_test_task/internal/api/mocks"
	api_models "vk_test_task/internal/api/models"
)

func TestHandler_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func(params api_models.AuthParams)

	testTable := []struct {
		name          string
		args          api_models.AuthParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.AuthParams{
				Login:    "login",
				Password: "password",
			},
			mockBehaviour: func(params api_models.AuthParams) {
				uc.EXPECT().SignIn(params).Return(api_models.SignInUseCaseResponse{}, nil)
			},
			wantErr: false,
		},
		{
			name: "invalid pass",
			args: api_models.AuthParams{
				Login:    "login",
				Password: "",
			},
			mockBehaviour: func(params api_models.AuthParams) {
				uc.EXPECT().SignIn(params).Return(api_models.SignInUseCaseResponse{}, fmt.Errorf(""))
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			ts := httptest.NewServer(h.SignIn())
			defer ts.Close()
			r, _ := json.Marshal(test.args)
			res, _ := http.Post(ts.URL, "application/json", bytes.NewReader(r))
			var response api_models.SignInUseCaseResponse
			json.NewDecoder(res.Body).Decode(&response)

			if test.wantErr {
				assert.NotEqual(t, "200 OK", res.Status)
			} else {
				assert.Equal(t, "200 OK", res.Status)
			}
		})
	}

}

func TestHandler_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func(params api_models.AuthParams)

	testTable := []struct {
		name          string
		args          api_models.AuthParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.AuthParams{
				Login:    "login",
				Password: "password",
			},
			mockBehaviour: func(params api_models.AuthParams) {
				uc.EXPECT().SignUp(params).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "invalid pass",
			args: api_models.AuthParams{
				Login:    "login",
				Password: "",
			},
			mockBehaviour: func(params api_models.AuthParams) {
				uc.EXPECT().SignUp(params).Return(fmt.Errorf(""))
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			ts := httptest.NewServer(h.SignUp())
			defer ts.Close()
			r, _ := json.Marshal(test.args)
			res, _ := http.Post(ts.URL, "application/json", bytes.NewReader(r))

			if test.wantErr {
				assert.NotEqual(t, "200 OK", res.Status)
			} else {
				assert.Equal(t, "200 OK", res.Status)
			}
		})
	}

}
