package middlewares

import (
	"Rent-And-Buy-App/pkg/Response"
	"Rent-And-Buy-App/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//type AuthMiddleware struct {
//	jwtManager *auth.JWTManager
//}

func JWTAuth(jwtMgr *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			Response.JSON(c, http.StatusUnauthorized, "missing authorization header", nil)
			c.Abort()
			return
		}
		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			Response.JSON(c, http.StatusUnauthorized, "invalid authorization header", nil)
			c.Abort()
			return
		}
		token := parts[1]

		userId, err := jwtMgr.VerifyToken(token)
		if err != nil {
			Response.JSON(c, http.StatusUnauthorized, "invalid token", nil)
			c.Abort()
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
