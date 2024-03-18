package api_usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	mock_api "vk_test_task/internal/api/mocks"
	api_models "vk_test_task/internal/api/models"
	"vk_test_task/internal/utils/encryption"
)

func TestUseCase_SignIn(t *testing.T) {
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
				pass, _ := encryption.HashPassword(params.Password)
				repo.EXPECT().SignIn(params.Login).Return(api_models.SignInRepositoryResponse{
					UserId:       "id",
					HashPassword: pass,
					IsAdmin:      false,
				}, nil)
				tokenRepo.EXPECT().CreateTokensPair("id", false).Return("", "", int64(1), nil)
			},
			wantErr: false,
		},
		{
			name: "wrong password",
			args: api_models.AuthParams{
				Login:    "login",
				Password: "password",
			},
			mockBehaviour: func(params api_models.AuthParams) {
				pass := "wrong pass"
				repo.EXPECT().SignIn(params.Login).Return(api_models.SignInRepositoryResponse{
					UserId:       "id",
					HashPassword: pass,
					IsAdmin:      false,
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "invalid login",
			args: api_models.AuthParams{
				Login:    "",
				Password: "pass",
			},
			mockBehaviour: func(params api_models.AuthParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid pass",
			args: api_models.AuthParams{
				Login:    "asdfasdfad",
				Password: "",
			},
			mockBehaviour: func(params api_models.AuthParams) {
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			_, err := uc.SignIn(test.args)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestUseCase_SignUp(t *testing.T) {
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
				repo.EXPECT().SignUp(params.Login, gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "invalid login",
			args: api_models.AuthParams{
				Login:    "",
				Password: "password",
			},
			mockBehaviour: func(params api_models.AuthParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid pass",
			args: api_models.AuthParams{
				Login:    "login",
				Password: "",
			},
			mockBehaviour: func(params api_models.AuthParams) {
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehaviour(test.args)

			err := uc.SignUp(test.args)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
