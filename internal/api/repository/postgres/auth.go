package postgres

import (
	"fmt"
	api_models "vk_test_task/internal/api/models"
	"vk_test_task/internal/common"
)

func (r Repository) SignIn(login string) (api_models.SignInRepositoryResponse, error) {
	if login == "" {
		return api_models.SignInRepositoryResponse{}, fmt.Errorf("repository error: invalid login")
	}

	query := `select user_id, password, is_admin from "user"  where login = $1`

	rows, err := r.db.Query(query, login)
	if err != nil {
		return api_models.SignInRepositoryResponse{}, fmt.Errorf("repository error: %s", err)
	}
	defer rows.Close()

	var response api_models.SignInRepositoryResponse

	rows.Next()

	err = rows.Scan(&response.UserId, &response.HashPassword, &response.IsAdmin)
	if err != nil {
		return api_models.SignInRepositoryResponse{}, fmt.Errorf("wrong login or password")
	}

	return response, nil
}

func (r Repository) SignUp(login, hashPassword, userId string) error {
	if len(login) > common.LOGIN_MAXSIZE || len(login) < common.LOGIN_MINSIZE {
		return fmt.Errorf("invalid login")
	}
	if hashPassword == "" {
		return fmt.Errorf("invalid password")
	}
	if userId == "" {
		return fmt.Errorf("invalid userId")
	}

	checkQuery := `select exists(select 1 from "user" where login = $1)`

	var exists bool
	r.db.QueryRow(checkQuery, login).Scan(&exists)

	if exists {
		return fmt.Errorf("login already exists")
	}

	query := `insert into "user" (user_id, login, password, is_admin) values ($1, $2, $3, $4)`

	_, err := r.db.Exec(query, userId, login, hashPassword, false)

	if err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	return nil
}
