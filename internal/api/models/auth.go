package api_models

import "github.com/golang-jwt/jwt/v5"

type RefreshClaims struct {
	IsAdmin       bool `json:"is_admin"`
	jwt.MapClaims `json:"claims"`
}

type AuthClaims struct {
	UserId        string `json:"user_id"`
	IsAdmin       bool   `json:"is_admin"`
	jwt.MapClaims `json:"claims"`
}

type SignInRepositoryResponse struct {
	UserId       string
	HashPassword string
	IsAdmin      bool
}

type SignInUseCaseResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expiration   int64  `json:"expiration"`
}

type AuthParams struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Description string `json:"description"`
}
