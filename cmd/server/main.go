package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mjhjj/Tyrannosaurus/internal/domain"
	"github.com/mjhjj/Tyrannosaurus/internal/repository"
)

func main() {
	// Dependencies
	db, err := repository.NewSQLiteDB(os.Getenv("PLACES_BD_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepositories(db)

	// default router
	router := gin.Default()
	router.LoadHTMLFiles("add.html", "docs.html")

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
		})
		v1.GET("/addPlacePleace", func(c *gin.Context) {
			c.HTML(http.StatusOK, "add.html", gin.H{
				"title": "add",
			})
		})

		v1.Static("/images", "./images")
		v1.GET("/docs", func(c *gin.Context) {
			c.HTML(http.StatusOK, "docs.html", gin.H{
				"title": "docs",
			})
		})

		v1.GET("/add", func(c *gin.Context) {

			secret := c.Query("secret")
			if secret != os.Getenv("ADD_PLACE_SECRET") {
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
