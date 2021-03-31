package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjhjj/Tyrannosaurus/internal/domain"
	"github.com/mjhjj/Tyrannosaurus/internal/repository"
)

func main() {
	// Dependencies
	db, err := repository.NewSQLiteDB("psuyhribga8wayvrwatayuog.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepositories(db)

	// default router
	router := gin.Default()
	router.LoadHTMLFiles("add.html")

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
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		// get all places api
		v1.GET("/getAllPlaces", func(c *gin.Context) {
			places, err := repo.Places.SelectAllPlaces()
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "not found",
				})
				return
			}
			c.JSON(http.StatusOK, places)
			fmt.Println(places)
		})
		v1.GET("/addPlacePleace", func(c *gin.Context) {
			c.HTML(http.StatusOK, "add.html", gin.H{
				"title": "Main website",
			})
		})

		v1.GET("/add", func(c *gin.Context) {

			secret := c.Query("secret")
			if secret != "1mr4sist" {
				c.JSON(http.StatusForbidden, gin.H{
					"message": "forbidden",
				})
				return
			}
			x := c.Query("x")
			y := c.Query("y")
			name := c.Query("name")
			address := c.Query("address")
			about := c.Query("about")
			bio := c.Query("bio")
			link := c.Query("link")

			err := repo.Places.Insert(domain.Place{"0", x, y, name, address, about, bio, link})
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "internal error",
				})
				return
			}
			c.Redirect(http.StatusFound, "/home")
		})

	}
	// Listen and serve on :8080
	router.Run()
}
