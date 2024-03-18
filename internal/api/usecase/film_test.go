package api_usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	mock_api "vk_test_task/internal/api/mocks"
	api_models "vk_test_task/internal/api/models"
)

func TestUseCase_CreateFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_api.NewMockRepositoryInterface(ctrl)
	tokenRepo := mock_api.NewMockTokenRepositoryInterface(ctrl)

	uc := New(
		nil,
		nil,
		repo,
		tokenRepo,
	)

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
				Name:        "name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        10,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.CreateFilmParams) {
				repo.EXPECT().CreateFilm(gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "invalid name",
			args: api_models.CreateFilmParams{
				Name:        "",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        10,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.CreateFilmParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid rate",
			args: api_models.CreateFilmParams{
				Name:        "name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        31233,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.CreateFilmParams) {
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			_, err := uc.CreateFilm(test.args)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestUseCase_UpdateFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_api.NewMockRepositoryInterface(ctrl)
	tokenRepo := mock_api.NewMockTokenRepositoryInterface(ctrl)

	uc := New(
		nil,
		nil,
		repo,
		tokenRepo,
	)

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
				Name:        "Name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        10,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.UpdateFilmParams) {
				repo.EXPECT().UpdateFilm(params).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "invalid filmId",
			args: api_models.UpdateFilmParams{
				FilmId:      "",
				Name:        "Name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        10,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.UpdateFilmParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid rate",
			args: api_models.UpdateFilmParams{
				FilmId:      "id",
				Name:        "Name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        19999999,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.UpdateFilmParams) {
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			err := uc.UpdateFilm(test.args)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestUseCase_GetFilms(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_api.NewMockRepositoryInterface(ctrl)
	tokenRepo := mock_api.NewMockTokenRepositoryInterface(ctrl)

	uc := New(
		nil,
		nil,
		repo,
		tokenRepo,
	)

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
				repo.EXPECT().GetFilms(params).Return(api_models.GetFilmsResponse{}, nil)
			},
			wantErr: false,
		},
		{
			name: "invalid sort by",
			args: api_models.GetFilmsParams{
				SortBy:      9999,
				IsAscending: 1,
			},
			mockBehaviour: func(params api_models.GetFilmsParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid sort asc",
			args: api_models.GetFilmsParams{
				SortBy:      1,
				IsAscending: -1111,
			},
			mockBehaviour: func(params api_models.GetFilmsParams) {
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			_, err := uc.GetFilms(test.args)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestUseCase_DeleteFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_api.NewMockRepositoryInterface(ctrl)
	tokenRepo := mock_api.NewMockTokenRepositoryInterface(ctrl)

	uc := New(
		nil,
		nil,
		repo,
		tokenRepo,
	)

	type mockBehaviour func(filmId string)

	testTable := []struct {
		name          string
		filmId        string
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name:   "default",
			filmId: "id",
			mockBehaviour: func(filmId string) {
				repo.EXPECT().DeleteFilm(filmId).Return(nil)
			},
			wantErr: false,
		},
		{
			name:   "invalid filmId",
			filmId: "",
			mockBehaviour: func(filmId string) {
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.filmId)

			err := uc.DeleteFilm(api_models.DeleteFilmParams{FilmId: test.filmId})

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestUseCase_SearchFilm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_api.NewMockRepositoryInterface(ctrl)
	tokenRepo := mock_api.NewMockTokenRepositoryInterface(ctrl)

	uc := New(
		nil,
		nil,
		repo,
		tokenRepo,
	)

	type mockBehaviour func(params api_models.SearchFilmParams)

	testTable := []struct {
		name          string
		args          api_models.SearchFilmParams
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.SearchFilmParams{
				Name:      "name",
				ActorName: "",
			},
			mockBehaviour: func(params api_models.SearchFilmParams) {
				repo.EXPECT().SearchFilmByName(params.Name).Return(api_models.SearchFilmResponse{}, nil)
			},
			wantErr: false,
		},
		{
			name: "by actor",
			args: api_models.SearchFilmParams{
				Name:      "",
				ActorName: "actorname",
			},
			mockBehaviour: func(params api_models.SearchFilmParams) {
				repo.EXPECT().SearchFilmByActorName(params.ActorName).Return(api_models.SearchFilmResponse{}, nil)
			},
			wantErr: false,
		},
		{
			name: "invalid params",
			args: api_models.SearchFilmParams{
				Name:      "",
				ActorName: "",
			},
			mockBehaviour: func(params api_models.SearchFilmParams) {
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			_, err := uc.SearchFilm(test.args)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
