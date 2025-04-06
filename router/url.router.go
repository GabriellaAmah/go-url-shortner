package router

import (
	"net/http"

	"github.com/GabriellaAmah/go-url-shortner/middleware"
	"github.com/GabriellaAmah/go-url-shortner/url"
	"github.com/gorilla/mux"
)

func UrlRouter(r *mux.Router) {
	urlRouter := r.PathPrefix("/url").Subrouter()

	urlRouter.Use(middleware.AuthenticationMiddleware.Middleware)
	urlRouter.HandleFunc("/", url.UrlController.CreateUrl).Methods(http.MethodPost)
}
