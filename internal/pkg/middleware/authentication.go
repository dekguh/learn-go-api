package middleware

import (
	"net/http"
	"strings"

	"github.com/dekguh/learn-go-api/internal/pkg/jwt"
	httputils "github.com/dekguh/learn-go-api/internal/pkg/utils/http"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			httputils.NewErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		token := strings.SplitN(authHeader, " ", 2)
		if len(token) != 2 || strings.ToLower(token[0]) != "bearer" {
			httputils.NewErrorResponse(ctx, http.StatusUnauthorized, "Invalid Token")
			ctx.Abort()
			return
		}

		claims, err := jwt.ParseJwt(token[1])
		if err != nil {
			httputils.NewErrorResponse(ctx, http.StatusUnauthorized, "Invalid Token or expired token")
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("user_email", claims.Email)
		ctx.Next()
	}
}
