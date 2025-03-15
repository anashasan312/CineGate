package server

import "github.com/gin-gonic/gin"

func setupUserRoutes(r *gin.RouterGroup, s *HttpServer) {
	r.POST("",func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
}
