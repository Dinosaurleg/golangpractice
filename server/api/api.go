package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"

	"server/models"
)

var albums = []models.Album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	fmt.Println("doing a get")
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.Id == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("error with input object")
		return
	}

	fmt.Println("object inserted")
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
