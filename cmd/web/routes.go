package main

import (
	"net/http"

	"github.com/Zdev147/room_booking/pkg/config"
	"github.com/Zdev147/room_booking/pkg/handler"
	"github.com/go-chi/chi/v5"
)

func routes(config config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Get("/", handler.Repo.Home)

	return router
}
