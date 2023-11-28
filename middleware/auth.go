package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func Authorization() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/users/register"{
			ctx.Next()
			return
		}
		var resp dto.Response

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			resp.Message = "Invalid token"
			ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		err := godotenv.Load()
		if err != nil {
			log.Printf("unable to load env: %v\n", err)
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, apperror.ErrSigningMethodInvalid
			} else if method != jwt.SigningMethodHS256 {
				return nil, apperror.ErrSigningMethodInvalid
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			resp.Message = err.Error()
			ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			resp.Message = "Invalid token"
			ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			return
		}

		ctx.Set("user_id", claims["user_id"])

		ctx.Next()

	}
}