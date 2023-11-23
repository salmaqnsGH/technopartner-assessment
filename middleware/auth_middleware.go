package middleware

import (
	"net/http"
	"strings"
	"technopartner/test/helper"
	"technopartner/test/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/api/v1/users/register" || req.URL.Path == "/api/v1/users/login" {
		middleware.Handler.ServeHTTP(writer, req)
		return
	}

	authHeader := req.Header.Get("Authorization")

	if !strings.Contains(authHeader, "Bearer") {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	_, err := helper.ValidateToken(tokenString)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	middleware.Handler.ServeHTTP(writer, req)
}
