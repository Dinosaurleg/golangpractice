package main

import (
	"github.com/gin-gonic/gin"

	"server/api"
)

func main() {
	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.GET("albums/:id", api.GetAlbumById)
	router.POST("/albums", api.PostAlbum)

	router.Run("localhost:8080")
}
