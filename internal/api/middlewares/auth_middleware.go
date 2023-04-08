package middlewares

import "github.com/gin-gonic/gin"

type AuthMiddleware struct {
}

func (am AuthMiddleware) Authenticated(c *gin.Context) {
	// some logic
	c.Next()
}
