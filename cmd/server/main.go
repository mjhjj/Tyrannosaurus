package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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

		v1.POST("/add", func(c *gin.Context) {

			secret := c.PostForm("secret")
			if secret != os.Getenv("ADD_PLACE_SECRET") {
				c.JSON(http.StatusForbidden, gin.H{
					"message": "forbidden",
				})
				return
			}
			x := c.PostForm("x")
			y := c.PostForm("y")
			name := c.PostForm("name")
			address := c.PostForm("address")
			about := c.PostForm("about")
			bio := c.PostForm("bio")
			link := c.PostForm("link")
			linkName := c.PostForm("linkName")
			sity := c.PostForm("sity")
			nameForImage := time.Now().Unix()
			image, header, err := c.Request.FormFile("image")
			if err != nil {
				log.Println(err.Error())
			}
			// save image to images directory
			if header != nil {
				out, err := os.Create("images/" + fmt.Sprintf("%d", nameForImage))
				if err != nil {
					log.Println(err)
				}
				defer out.Close()
				_, err = io.Copy(out, image)
				if err != nil {
					log.Println(err)
				}

				err = repo.Places.Insert(domain.Place{Id: "0",
					PositionX:   x,
					PositionY:   y,
					Name:        name,
					Address:     address,
					About:       about,
					Bio:         bio,
					PanoramLink: link,
					LinkName:    linkName,
					Sity:        sity,
					Image:       "/api/v1/images/" + fmt.Sprintf("%d", nameForImage),
				})
				c.Redirect(http.StatusFound, "/home")
				return
			}
			// if image is not uploaded
			err = repo.Places.Insert(domain.Place{Id: "0",
				PositionX:   x,
				PositionY:   y,
				Name:        name,
				Address:     address,
				About:       about,
				Bio:         bio,
				PanoramLink: link,
				LinkName:    linkName,
				Sity:        sity,
				Image:       "",
			})
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
