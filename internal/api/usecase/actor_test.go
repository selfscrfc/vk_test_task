package api_usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	mock_api "vk_test_task/internal/api/mocks"
	api_models "vk_test_task/internal/api/models"
)

func TestUseCase_CreateActor(t *testing.T) {
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
				repo.EXPECT().CreateActor(gomock.Any()).Return(nil)
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
			},
			wantErr: true,
		},
		{
			name: "invalid sex",
			args: api_models.CreateActorParams{
				Name:  "Name",
				Sex:   999,
				Birth: time.Now(),
			},
			mockBehaviour: func(params api_models.CreateActorParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid birth",
			args: api_models.CreateActorParams{
				Name:  "Name",
				Sex:   1,
				Birth: time.Now().Add(10 * time.Hour),
			},
			mockBehaviour: func(params api_models.CreateActorParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid name case",
			args: api_models.CreateActorParams{
				Name:  "name",
				Sex:   1,
				Birth: time.Now(),
			},
			mockBehaviour: func(params api_models.CreateActorParams) {
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			_, err := uc.CreateActor(test.args)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestUseCase_UpdateActor(t *testing.T) {
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
				ActorId: "id1",
				Name:    "Name",
				Sex:     1,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
				repo.EXPECT().UpdateActor(params).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "invalid actor id",
			args: api_models.UpdateActorParams{
				ActorId: "",
				Name:    "Name",
				Sex:     1,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid actor name",
			args: api_models.UpdateActorParams{
				ActorId: "asdfafa",
				Name:    "name",
				Sex:     1,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid actor name",
			args: api_models.UpdateActorParams{
				ActorId: "asdasd",
				Name:    "",
				Sex:     1,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid actor sex",
			args: api_models.UpdateActorParams{
				ActorId: "asdasd",
				Name:    "Aaa",
				Sex:     -1,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid actor name",
			args: api_models.UpdateActorParams{
				ActorId: "asdasd",
				Name:    "Asdf",
				Sex:     1,
				Birth:   time.Now().Add(10 * time.Hour),
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			err := uc.UpdateActor(test.args)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestUseCase_GetActors(t *testing.T) {
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

	type mockBehaviour func()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "default",
			mockBehaviour: func() {
				repo.EXPECT().GetActors().Return(api_models.GetActorsResponse{}, nil)
			},
			wantErr: false,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour()

			_, err := uc.GetActors()

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestUseCase_DeleteActors(t *testing.T) {
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

	type mockBehaviour func(actorId string)

	testTable := []struct {
		name          string
		actorId       string
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name:    "default",
			actorId: "id",
			mockBehaviour: func(actorId string) {
				repo.EXPECT().DeleteActor(actorId).Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "invalid actorId",
			actorId: "",
			mockBehaviour: func(actorId string) {
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.actorId)

			err := uc.DeleteActor(api_models.DeleteActorParams{ActorId: test.actorId})

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
