package main

import "://github.com"

func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/manifest.webmanifest", func(c *gin.Context) {
		c.Header("Content-Type", "application/manifest+json")
		c.File("./static/manifest.webmanifest")
	})

	r.Run(":8080")
}
