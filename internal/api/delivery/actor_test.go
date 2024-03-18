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
	"time"
	mock_api "vk_test_task/internal/api/mocks"
	api_models "vk_test_task/internal/api/models"
)

func TestHandler_CreateActor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func(params api_models.CreateActorParams)

	testTable := []struct {
		name          string
		args          api_models.CreateActorParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.CreateActorParams{
				Name:  "Name",
				Sex:   1,
				Birth: time.Now(),
			},
			mockBehaviour: func(params api_models.CreateActorParams) {
				uc.EXPECT().CreateActor(gomock.Any()).Return("userid", nil)
			},
			wantErr: false,
		},
		{
			name: "invalid name",
			args: api_models.CreateActorParams{
				Name:  "",
				Sex:   1,
				Birth: time.Now(),
			},
			mockBehaviour: func(params api_models.CreateActorParams) {
				uc.EXPECT().CreateActor(gomock.Any()).Return("", fmt.Errorf(""))
			},
			wantErr: true,
		},
		{
			name: "invalid params",
			args: api_models.CreateActorParams{},
			mockBehaviour: func(params api_models.CreateActorParams) {
				uc.EXPECT().CreateActor(gomock.Any()).Return("", fmt.Errorf(""))
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			ts := httptest.NewServer(h.CreateActor())
			defer ts.Close()
			r, _ := json.Marshal(test.args)
			res, _ := http.Post(ts.URL, "application/json", bytes.NewReader(r))
			var response api_models.CreateActorParams
			json.NewDecoder(res.Body).Decode(&response)

			if test.wantErr {
				assert.NotEqual(t, "200 OK", res.Status)
			} else {
				assert.Equal(t, "200 OK", res.Status)
			}
		})
	}
}

func TestHandler_UpdateActor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func(params api_models.UpdateActorParams)

	testTable := []struct {
		name          string
		args          api_models.UpdateActorParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.UpdateActorParams{
				ActorId: "id",
				Name:    "Name",
				Sex:     1,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
				uc.EXPECT().UpdateActor(gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "bad request",
			args: api_models.UpdateActorParams{
				ActorId: "",
				Name:    "Name",
				Sex:     1,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
				uc.EXPECT().UpdateActor(gomock.Any()).Return(fmt.Errorf(""))
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			ts := httptest.NewServer(h.UpdateActor())
			defer ts.Close()
			r, _ := json.Marshal(test.args)
			res, _ := http.Post(ts.URL, "application/json", bytes.NewReader(r))
			var response api_models.UpdateActorParams
			json.NewDecoder(res.Body).Decode(&response)

			if test.wantErr {
				assert.NotEqual(t, "200 OK", res.Status)
			} else {
				assert.Equal(t, "200 OK", res.Status)
			}
		})
	}

}

func TestHandler_GetActors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			mockBehaviour: func() {
				uc.EXPECT().GetActors().Return(api_models.GetActorsResponse{}, nil)
			},
			wantErr: false,
		},
		{
			name: "internal server error",
			mockBehaviour: func() {
				uc.EXPECT().GetActors().Return(api_models.GetActorsResponse{}, fmt.Errorf(""))
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour()

			ts := httptest.NewServer(h.GetActors())
			defer ts.Close()
			res, _ := http.Post(ts.URL, "application/json", nil)
			var response api_models.GetActorsResponse
			json.NewDecoder(res.Body).Decode(&response)

			if test.wantErr {
				assert.NotEqual(t, "200 OK", res.Status)
			} else {
				assert.Equal(t, "200 OK", res.Status)
			}
		})
	}

}

func TestHandler_DeleteActor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func(params api_models.DeleteActorParams)

	testTable := []struct {
		name          string
		args          api_models.DeleteActorParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.DeleteActorParams{
				ActorId: "id",
			},
			mockBehaviour: func(params api_models.DeleteActorParams) {
				uc.EXPECT().DeleteActor(params).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "internal server error",
			args: api_models.DeleteActorParams{
				ActorId: "",
			},
			mockBehaviour: func(params api_models.DeleteActorParams) {
				uc.EXPECT().DeleteActor(params).Return(fmt.Errorf(""))
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			ts := httptest.NewServer(h.DeleteActor())
			defer ts.Close()
			r, _ := json.Marshal(test.args)
			res, _ := http.Post(ts.URL, "application/json", bytes.NewReader(r))
			var response api_models.DeleteActorParams
			json.NewDecoder(res.Body).Decode(&response)

			if test.wantErr {
				assert.NotEqual(t, "200 OK", res.Status)
			} else {
				assert.Equal(t, "200 OK", res.Status)
			}
		})
	}

}
