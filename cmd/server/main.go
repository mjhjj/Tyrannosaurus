package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// default router
	router := gin.Default()

	// serve frontend
	router.Static("/home", "./dist")
	router.Static("/css", "./dist/css")
	router.Static("/js", "./dist/js")
	// redirect to main page
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/home")
	})

	// Api for our frontend: v1
	v1 := router.Group("/api/v1")
	{
		// TODO: add handlers
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	// Listen and serve on :8080
	router.Run("127.0.0.1:8080")
}
