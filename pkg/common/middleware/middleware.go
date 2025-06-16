package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func Auth(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := ctx.GetHeader("Authorization")

		tokenStr := strings.TrimPrefix(h, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(key), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, genError("invalid claims"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, genError("invalid claims"))
			return
		}

		uId, ok := claims["user_id"].(float64)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, genError("invalid claims"))
			return
		}

		ctx.Set("userId", int32(uId))
		ctx.Next()
	}
}

func genError(msg string) gin.H {
	return gin.H{"error": msg}
}
