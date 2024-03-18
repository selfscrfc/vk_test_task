package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
	"vk_test_task/config"
	"vk_test_task/internal/api/models"
)

type Repository struct {
	Cfg    *config.Config
	Logger *slog.Logger
	DB     *redis.Client
}

func New(cfg *config.Config, logger *slog.Logger) Repository {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		DB:   cfg.Redis.Database,
	})

	logger.Debug("redis database connected")

	return Repository{
		Cfg:    cfg,
		Logger: logger,
		DB:     client,
	}
}

func (r Repository) CreateAccessToken(userId string, isAdmin bool) (string, int64, error) {
	if userId == "" {
		return "", 0, fmt.Errorf("redis error: invalid userId")
	}

	exp := time.Now().Add(time.Second * time.Duration(r.Cfg.Server.AccessLifetime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, api_models.AuthClaims{
		UserId:  userId,
		IsAdmin: isAdmin,
		MapClaims: jwt.MapClaims{
			"exp": exp,
		},
	})

	tokenString, err := token.SignedString([]byte(r.Cfg.Server.AccessSecret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, exp, nil
}

func (r Repository) CreateRefreshToken(isAdmin bool) (string, int64, error) {
	exp := time.Now().Add(time.Second * time.Duration(r.Cfg.Server.RefreshLifetime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, api_models.RefreshClaims{
		IsAdmin: isAdmin,
		MapClaims: jwt.MapClaims{
			"exp": exp,
		},
	})

	tokenString, err := token.SignedString([]byte(r.Cfg.Server.RefreshSecret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, exp, nil
}

func (r Repository) VerifyRefreshToken(userId string, tokenString string) (bool, error) {
	if userId == "" {
		return false, errors.New("invalid userId")
	}
	if tokenString == "" {
		return false, errors.New("invalid token")
	}

	var claims api_models.RefreshClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(token *jwt.Token) (interface{}, error) { return []byte(r.Cfg.Server.RefreshSecret), nil })
	if err != nil {
		return false, err
	}

	if token.Valid == false {
		return false, errors.New("invalid jwt token")
	}

	cmd := r.DB.Get(context.Background(), fmt.Sprintf("refresh-token-%s", userId))

	if err = cmd.Err(); err != nil {
		return false, fmt.Errorf("redis error: %s", err.Error())
	}

	if cmd.Val() != tokenString {
		return false, errors.New("out of date token")
	}

	return claims.IsAdmin, nil
}

func (r Repository) UpdateAccessToken(userId string, refreshToken string) (string, int64, error) {
	if userId == "" {
		return "", 0, fmt.Errorf("redis error: invalid userId")
	}
	if refreshToken == "" {
		return "", 0, fmt.Errorf("redis error: invalid refreshToken")
	}

	isAdmin, err := r.VerifyRefreshToken(userId, refreshToken)
	if err != nil {
		return "", 0, err
	}

	token, exp, err := r.CreateAccessToken(userId, isAdmin)
	if err != nil {
		return "", 0, err
	}

	return token, exp, nil
}

func (r Repository) CreateTokensPair(userId string, isAdmin bool) (string, string, int64, error) {
	accessToken, exp, err := r.CreateAccessToken(userId, isAdmin)
	if err != nil {
		return "", "", 0, err
	}

	refreshToken, _, err := r.CreateRefreshToken(isAdmin)
	if err != nil {
		return "", "", 0, err
	}

	cmd := r.DB.Set(
		context.Background(),
		fmt.Sprintf("refresh-token-%s", userId),
		refreshToken,
		0,
	)

	if err = cmd.Err(); err != nil {
		return "", "", 0, fmt.Errorf("redis error: %s", err.Error())
	}

	return accessToken, refreshToken, exp, nil
}
