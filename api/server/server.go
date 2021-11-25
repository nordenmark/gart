package server

import (
	"gart/config"
	"gart/images"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getImageNames(engine *gin.Engine) []string {
	images := []string{}

	for _, route := range engine.Routes() {
		if route.Path[0:3] == "/g/" {
			images = append(images, route.Path[3:])
		}
	}

	return images
}

var configurationMap = map[string]config.Configuration{
	"janus":    images.JanusConfiguration{},
	"shapes":   images.ShapesConfiguration{},
	"contours": images.ContoursConfiguration{},
}

func CreateServer() *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.Use(cors.Default())

	generate := r.Group("/g")
	{
		generate.GET("/janus", func(c *gin.Context) {
			var config images.JanusConfiguration

			if err := c.BindQuery(&config); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			path := images.Janus(config)

			c.File(path)
		})

		generate.GET("/shapes", func(c *gin.Context) {
			var config images.ShapesConfiguration

			if err := c.BindQuery(&config); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			path := images.Shapes(config)

			c.File(path)
		})

		generate.GET("/contours", func(c *gin.Context) {
			var config images.ContoursConfiguration

			if err := c.BindQuery(&config); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			path := images.Contours(config)

			c.File(path)
		})
	}

	api := r.Group("/api")
	{
		api.GET("/images", func(c *gin.Context) {
			imageNames := getImageNames(r)

			imageObjects := []config.ImageObject{}

			for _, imageName := range imageNames {
				imageObjects = append(imageObjects, config.ImageObject{
					Name:       imageName,
					Parameters: configurationMap[imageName].Parameters(),
				})
			}

			c.JSON(http.StatusOK, imageObjects)
		})
	}

	return r
}
