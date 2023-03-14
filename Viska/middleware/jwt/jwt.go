package middleware

import (
	"net/http"
	"strings"
	"viska/storage/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/thedevsaddam/renderer"
)

type TokenGenerator struct {
	render *renderer.Render
}

func NewTokenGenerator(render *renderer.Render) *TokenGenerator {
	return &TokenGenerator{
		render: render,
	}
}

var JWT_KEY = []byte("12345asdqwerty")

type JWTClaims struct {
	Username string
	jwt.RegisteredClaims
}

func (t *TokenGenerator) MiddlewareJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			t.render.JSON(w, http.StatusBadRequest, &models.Template3{
				ResponseCode:    http.StatusBadRequest,
				ResponseMessage: "Status Bad Request",
			})
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				t.render.JSON(w, http.StatusUnauthorized, &models.Template1{
					Message: "Token is Invalid",
					Data:    nil,
				})
				return
			case jwt.ValidationErrorExpired:
				t.render.JSON(w, http.StatusUnauthorized, &models.Template1{
					Message: "Token is Unauthorized",
					Data:    nil,
				})
				return
			default:
				t.render.JSON(w, http.StatusUnauthorized, &models.Template1{
					Message: "Token is Unauthorized",
					Data:    nil,
				})
				return
			}
		}

		if !token.Valid {
			t.render.JSON(w, http.StatusUnauthorized, &models.Template1{
				Message: "Token is Invalid",
				Data:    nil,
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
