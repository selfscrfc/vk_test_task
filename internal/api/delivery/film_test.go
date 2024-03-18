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

func TestHandler_CreateFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func(params api_models.CreateFilmParams)

	testTable := []struct {
		name          string
		args          api_models.CreateFilmParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.CreateFilmParams{
				Name:        "Name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        2,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.CreateFilmParams) {
				uc.EXPECT().CreateFilm(gomock.Any()).Return("", nil)
			},
			wantErr: false,
		},
		{
			name: "invalid name",
			args: api_models.CreateFilmParams{
				Name:        "",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        2,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.CreateFilmParams) {
				uc.EXPECT().CreateFilm(gomock.Any()).Return("", fmt.Errorf(""))
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			ts := httptest.NewServer(h.CreateFilm())
			defer ts.Close()
			r, _ := json.Marshal(test.args)
			res, _ := http.Post(ts.URL, "application/json", bytes.NewReader(r))
			var response api_models.CreateFilmParams
			json.NewDecoder(res.Body).Decode(&response)

			if test.wantErr {
				assert.NotEqual(t, "200 OK", res.Status)
			} else {
				assert.Equal(t, "200 OK", res.Status)
			}
		})
	}
}

func TestHandler_UpdateFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func(params api_models.UpdateFilmParams)

	testTable := []struct {
		name          string
		args          api_models.UpdateFilmParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.UpdateFilmParams{
				FilmId:      "id",
				Name:        "",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        2,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.UpdateFilmParams) {
				uc.EXPECT().UpdateFilm(gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "bad request",
			args: api_models.UpdateFilmParams{
				FilmId:      "",
				Name:        "",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        2,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.UpdateFilmParams) {
				uc.EXPECT().UpdateFilm(gomock.Any()).Return(fmt.Errorf(""))
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			ts := httptest.NewServer(h.UpdateFilm())
			defer ts.Close()
			r, _ := json.Marshal(test.args)
			res, _ := http.Post(ts.URL, "application/json", bytes.NewReader(r))
			var response api_models.UpdateFilmParams
			json.NewDecoder(res.Body).Decode(&response)

			if test.wantErr {
				assert.NotEqual(t, "200 OK", res.Status)
			} else {
				assert.Equal(t, "200 OK", res.Status)
			}
		})
	}

}

func TestHandler_GetFilms(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func(params api_models.GetFilmsParams)

	testTable := []struct {
		name          string
		args          api_models.GetFilmsParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.GetFilmsParams{
				SortBy:      1,
				IsAscending: 1,
			},
			mockBehaviour: func(params api_models.GetFilmsParams) {
				uc.EXPECT().GetFilms(params).Return(api_models.GetFilmsResponse{}, nil)
			},
			wantErr: false,
		},
		{
			name: "internal server error",
			args: api_models.GetFilmsParams{
				SortBy:      1,
				IsAscending: 1,
			},
			mockBehaviour: func(params api_models.GetFilmsParams) {
				uc.EXPECT().GetFilms(params).Return(api_models.GetFilmsResponse{}, fmt.Errorf(""))
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			ts := httptest.NewServer(h.GetFilms())
			defer ts.Close()
			r, _ := json.Marshal(test.args)
			res, _ := http.Post(fmt.Sprintf(ts.URL+"/film/get?sort_by=%d&asc=%d",
				test.args.SortBy, test.args.IsAscending),
				"application/json", bytes.NewReader(r))
			var response api_models.GetFilmsResponse
			json.NewDecoder(res.Body).Decode(&response)

			if test.wantErr {
				assert.NotEqual(t, "200 OK", res.Status)
			} else {
				assert.Equal(t, "200 OK", res.Status)
			}
		})
	}

}

func TestHandler_DeleteFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func(params api_models.DeleteFilmParams)

	testTable := []struct {
		name          string
		args          api_models.DeleteFilmParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.DeleteFilmParams{
				FilmId: "id",
			},
			mockBehaviour: func(params api_models.DeleteFilmParams) {
				uc.EXPECT().DeleteFilm(params).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "internal server error",
			args: api_models.DeleteFilmParams{
				FilmId: "",
			},
			mockBehaviour: func(params api_models.DeleteFilmParams) {
				uc.EXPECT().DeleteFilm(params).Return(fmt.Errorf(""))
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			ts := httptest.NewServer(h.DeleteFilm())
			defer ts.Close()
			r, _ := json.Marshal(test.args)
			res, _ := http.Post(ts.URL, "application/json", bytes.NewReader(r))
			var response api_models.DeleteFilmParams
			json.NewDecoder(res.Body).Decode(&response)

			if test.wantErr {
				assert.NotEqual(t, "200 OK", res.Status)
			} else {
				assert.Equal(t, "200 OK", res.Status)
			}
		})
	}

}

func TestHandler_SearchFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_api.NewMockUseCaseInterface(ctrl)
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{}))
	h := New(nil, l, uc)

	type mockBehaviour func(params api_models.SearchFilmParams)

	type test struct {
		name          string
		args          api_models.SearchFilmParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}

	test1 := test{
		name: "default",
		args: api_models.SearchFilmParams{
			Name:      "name",
			ActorName: "",
		},
		mockBehaviour: func(params api_models.SearchFilmParams) {
			uc.EXPECT().SearchFilm(params).Return(api_models.SearchFilmResponse{}, nil)
		},
		wantErr: false,
	}
	t.Run(test1.name, func(t *testing.T) {
		test1.mockBehaviour(test1.args)

		ts := httptest.NewServer(h.SearchFilm())
		defer ts.Close()
		r, _ := json.Marshal(test1.args)
		res, _ := http.Post(fmt.Sprintf(ts.URL+"/film/search?name=%s",
			test1.args.Name),
			"application/json", bytes.NewReader(r))
		var response api_models.GetFilmsResponse
		json.NewDecoder(res.Body).Decode(&response)

		if test1.wantErr {
			assert.NotEqual(t, "200 OK", res.Status)
		} else {
			assert.Equal(t, "200 OK", res.Status)
		}
	})
	test2 := test{
		name: "internal server error",
		args: api_models.SearchFilmParams{
			Name:      "",
			ActorName: "",
		},
		mockBehaviour: func(params api_models.SearchFilmParams) {
		},
		wantErr: true,
	}
	t.Run(test2.name, func(t *testing.T) {
		test2.mockBehaviour(test2.args)

		ts := httptest.NewServer(h.SearchFilm())
		defer ts.Close()
		r, _ := json.Marshal(test2.args)
		res, _ := http.Post(ts.URL+"/film/search",
			"application/json", bytes.NewReader(r))

		if test2.wantErr {
			assert.NotEqual(t, "200 OK", res.Status)
		} else {
			assert.Equal(t, "200 OK", res.Status)
		}
	})
	test3 := test{
		name: "by actor name",
		args: api_models.SearchFilmParams{
			Name:      "",
			ActorName: "actor",
		},
		mockBehaviour: func(params api_models.SearchFilmParams) {
			uc.EXPECT().SearchFilm(params).Return(api_models.SearchFilmResponse{}, nil)
		},
		wantErr: false,
	}
	t.Run(test3.name, func(t *testing.T) {
		test3.mockBehaviour(test3.args)

		ts := httptest.NewServer(h.SearchFilm())
		defer ts.Close()
		r, _ := json.Marshal(test3.args)
		res, _ := http.Post(ts.URL+fmt.Sprintf("/film/search?actor_name=%s", test3.args.ActorName),
			"application/json", bytes.NewReader(r))
		var response api_models.GetFilmsResponse
		json.NewDecoder(res.Body).Decode(&response)

		if test3.wantErr {
			assert.NotEqual(t, "200 OK", res.Status)
		} else {
			assert.Equal(t, "200 OK", res.Status)
		}
	})
}
