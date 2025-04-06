package router

import (
	"net/http"

	"github.com/GabriellaAmah/go-url-shortner/user"
	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) {
	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/signup", user.UserController.CreateUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/login", user.UserController.LoginUser).Methods(http.MethodPost)
}
