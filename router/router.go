package router

import (
	"net/http"

	"github.com/GabriellaAmah/go-url-shortner/health"
	"github.com/gorilla/mux"
)

func RegisterRouters(r *mux.Router) {
	r.HandleFunc("/health", health.HealthController.GetAppHealth).Methods(http.MethodGet)
	UserRouter(r)
	UrlRouter(r)
}
