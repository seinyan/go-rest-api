package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/seinyan/go-rest-api/service"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(jwtServ service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {

		const BearerSchema = "Bearer"

		authHeader := c.GetHeader("Authorization")
		parts := strings.Fields(authHeader)

		// Token base validation
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API token required",
			})
			return
		} else if parts[0] != BearerSchema { // Token base validation
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header must start with Bearer",
			})
			return
		} else if len(parts) == 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token not found",
			})
			return
		} else if len(parts) > 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header must be Bearer and token",
			})
			return
		}

		token, err := jwtServ.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}


		if token.Valid {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "errMessage",
			})
		}
	}
}
