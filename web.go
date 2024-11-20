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
		"templates/talks.tmpl",
	)
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/talks.json", func(c *gin.Context) {
		talks, err := getTalks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, talks)
	})

	r.GET("/talks", func(c *gin.Context) {
		talks, err := getTalks()
		if err != nil {
			// TODO: HTML error pages
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.HTML(http.StatusOK, "talks.tmpl", gin.H{
			"talks": talks,
		})
	})

	return r.Run(":8080")
}
