package main

import (
	"github.com/gin-gonic/gin"

	"server/api"

	dboperations "server/dbOperations"
)

func main() {
	dboperations.CreateDb()
	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.GET("album/:id", api.GetAlbumByID)
	router.POST("/albums", api.PostAlbum)
	router.DELETE("/album/:id", api.DeleteAlbum)
	router.PATCH("/albums/:id", api.UpdateAlbum)

	router.Run("localhost:8080")
}
