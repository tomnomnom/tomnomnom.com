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
		"templates/tools.tmpl",
		"templates/videos.tmpl",
		"templates/sheep.tmpl",
		"templates/writing.tmpl",
		"templates/article.tmpl",
	)
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicons/favicon.ico")
	r.StaticFile("/site.webmanifest", "./static/favicons/site.webmanifest")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			// TODO: page <title>
			"title": "TomNomNom.com",
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

	r.GET("/tools", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tools.tmpl", gin.H{})
	})

	posts, err := getArticles()
	if err != nil {
		return err
	}

	r.GET("/writing", func(c *gin.Context) {
		c.HTML(http.StatusOK, "writing.tmpl", gin.H{
			"posts": posts,
		})
	})

	r.GET("/writing/:post_id", func(c *gin.Context) {

		post := posts.find(c.Param("post_id"))
		if post == nil {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}

		c.HTML(http.StatusOK, "article.tmpl", gin.H{
			"content": post.Content,
		})
	})

	return r.Run(":8080")
}
