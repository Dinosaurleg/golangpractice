package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"fmt"

	dboperations "server/dbOperations"
	"server/models"
)

func GetAlbums(c *gin.Context) {
	var albums []models.Album
	albums = dboperations.GetRecords()
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("error with input object")
		return
	}

	dboperations.InsertEntry(newAlbum)
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	err = dboperations.DeleteEntry(uint(u64))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error deleting album"})
		return
	}
}
