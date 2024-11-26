package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func serve() error {

	r := gin.Default()

	r.LoadHTMLFiles(
		"templates/header.tmpl",
		"templates/footer.tmpl",
		"templates/index.tmpl",
		"templates/videos.tmpl",
		"templates/sheep.tmpl",
	)
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicons/favicon.ico")
	r.StaticFile("/site.webmanifest", "./static/favicons/site.webmanifest")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/videos.json", func(c *gin.Context) {
		videos, err := getVideos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, videos)
	})

	r.GET("/videos", func(c *gin.Context) {
		videos, err := getVideos()
		if err != nil {
			// TODO: HTML error pages
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.HTML(http.StatusOK, "videos.tmpl", gin.H{
			"videos": videos,
		})
	})

	r.GET("/sheep", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sheep.tmpl", gin.H{})
	})

	return r.Run(":8080")
}
