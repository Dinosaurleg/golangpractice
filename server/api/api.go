package api

import (
	// "net/http"

	"github.com/gin-gonic/gin"

	"fmt"

	dboperations "server/dbOperations"
	"server/models"
)

func GetAlbums(c *gin.Context) {
	dboperations.GetRecords()
}

// func GetAlbumById(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, album := range albums {
// 		if album.Id == id {
// 			c.IndentedJSON(http.StatusOK, album)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("error with input object")
		return
	}

	dboperations.InsertEntry(newAlbum)
}
