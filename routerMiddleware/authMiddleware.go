package routerMiddleware

import (
	"net/http"
	"strings"

	"interviewProject/auth"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		prefix := "Bearer "
		authHeader := c.GetHeader("authorization")
		clientToken := strings.TrimPrefix(authHeader, prefix)
		
		claims, err1 := auth.ValidateToken(clientToken)
		if err1 != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "You are unauthorized to access this resource"})
			c.Abort()
			return
		}
		c.Set("id",claims.ID)
	}
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://commit-interview.herokuapp.com")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
