package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/GabriellaAmah/go-url-shortner/user"
	"github.com/GabriellaAmah/go-url-shortner/util"
)

type IAuthenticationMiddleware struct {
	userService user.IUserService
}

func (amw *IAuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		splitedAuth := strings.Split(authorization, " ")
		if len(splitedAuth) < 2 {
			 util.HttpErrorResponse(res, util.MakeError{
				IsError: true,
				Message: "Bearer auth token is required",
				Code:    401,
			})

			return
		}

		token := splitedAuth[1]

		decodedJwt, err := util.DecodeAuthJwt(token)
		if err.IsError {
			util.HttpErrorResponse(res, err)
			return 
		}

		validUser, err := amw.userService.GetUserDetailsById(decodedJwt.UserId)
		if err.IsError || validUser.Email == "" {
			 util.HttpErrorResponse(res, util.MakeError{
				IsError: true,
				Message: "Invalid User",
				Code:    401,
			})
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), "user", validUser))

		next.ServeHTTP(res, req)
	})
}

var AuthenticationMiddleware = IAuthenticationMiddleware{userService: user.UserService}
