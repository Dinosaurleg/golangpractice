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

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	album := dboperations.GetRecordByID(uint(u64))
	c.IndentedJSON(http.StatusOK, album)
}

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("error with input object")
		return
	}

	dboperations.InsertEntry(newAlbum)
}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	u64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var updatedAlbum models.Album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		fmt.Println("error with input")
		return
	}
	dboperations.UpdateEntry(uint(u64), updatedAlbum)
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
