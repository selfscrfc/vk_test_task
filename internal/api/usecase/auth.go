package api_usecase

import (
	"fmt"
	"github.com/google/uuid"
	api_models "vk_test_task/internal/api/models"
	"vk_test_task/internal/common"
	"vk_test_task/internal/utils/encryption"
)

func (u UseCase) SignIn(params api_models.AuthParams) (api_models.SignInUseCaseResponse, error) {
	if len(params.Login) < common.LOGIN_MINSIZE || len(params.Password) < common.PASSWORD_MINSIZE {
		return api_models.SignInUseCaseResponse{}, fmt.Errorf("usecase error: wrong params")
	}
	repoResponse, err := u.db.SignIn(params.Login)
	if err != nil {
		return api_models.SignInUseCaseResponse{}, fmt.Errorf("usecase error: %s", err.Error())
	}

	if ok := encryption.CheckPasswordHash(params.Password, repoResponse.HashPassword); ok != true {
		return api_models.SignInUseCaseResponse{}, fmt.Errorf("wrong login or password")
	}

	accessToken, refreshToken, exp, err := u.rdb.CreateTokensPair(repoResponse.UserId, repoResponse.IsAdmin)
	if err != nil {
		return api_models.SignInUseCaseResponse{}, fmt.Errorf("usecase error: %s", err.Error())
	}

	return api_models.SignInUseCaseResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiration:   exp,
	}, nil
}

func (u UseCase) SignUp(params api_models.AuthParams) error {
	if len(params.Login) < common.LOGIN_MINSIZE {
		return fmt.Errorf("usecase error: login is too short")
	}
	if len(params.Password) < common.PASSWORD_MINSIZE {
		return fmt.Errorf("usecase error: password is too short")
	}
	if len(params.Password) > common.PASSWORD_MAXSIZE {
		return fmt.Errorf("usecase error: password is too long")
	}

	userId, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("usecase error: %s", err.Error())
	}

	hashPassword, err := encryption.HashPassword(params.Password)
	if err != nil {
		return fmt.Errorf("usecase error: %s", err.Error())
	}

	err = u.db.SignUp(params.Login, hashPassword, userId.String())
	if err != nil {
		return fmt.Errorf("usecase error: %s", err.Error())
	}

	return nil
}
