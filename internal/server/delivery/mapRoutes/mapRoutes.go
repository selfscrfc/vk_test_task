package mapRoutes

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
	"vk_test_task/config"
	_ "vk_test_task/docs"
	"vk_test_task/internal/api"
	"vk_test_task/internal/middleware"
)

func MapApiRoutes(cfg *config.Config, logger *slog.Logger, h api.HandlerInterface) {
	secret := cfg.Server.AccessSecret

	http.HandleFunc("/actor/create", middleware.JWTAdminAuth(secret, logger, h.CreateActor()))
	http.HandleFunc("/actor/get", middleware.JWTUserAuth(secret, logger, h.GetActors()))
	http.HandleFunc("/actor/update", middleware.JWTAdminAuth(secret, logger, h.UpdateActor()))
	http.HandleFunc("/actor/delete", middleware.JWTAdminAuth(secret, logger, h.DeleteActor()))

	http.HandleFunc("/film/create", middleware.JWTAdminAuth(secret, logger, h.CreateFilm()))
	http.HandleFunc("/film/get", middleware.JWTUserAuth(secret, logger, h.GetFilms()))
	http.HandleFunc("/film/update", middleware.JWTAdminAuth(secret, logger, h.UpdateFilm()))
	http.HandleFunc("/film/delete", middleware.JWTAdminAuth(secret, logger, h.DeleteFilm()))
	http.HandleFunc("/film/search", middleware.JWTUserAuth(secret, logger, h.SearchFilm()))

	http.HandleFunc("/sign_in", h.SignIn())
	http.HandleFunc("/sign_up", h.SignUp())

	swagUrl := fmt.Sprintf("http://localhost:%s/swagger/doc.json", cfg.Server.Port)
	http.Handle("/swagger/*", httpSwagger.Handler(httpSwagger.URL(swagUrl)))
}
