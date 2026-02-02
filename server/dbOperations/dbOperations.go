package dboperations

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"server/models"
)

var sqlLiteDb *gorm.DB

func openDb() *gorm.DB {
	sqlLiteDb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	return sqlLiteDb
}

func closeDb(sqliteDb *gorm.DB) {
	defer sqliteDb.DB()
}

func CreateDb() {
	sqlLiteDb := openDb()
	err := sqlLiteDb.AutoMigrate(&models.Album{})

	if err != nil {
		panic("failed to automigrate database schema")
	}
	closeDb(sqlLiteDb)
}

func GetRecords() {
	sqlLiteDb := openDb()

	var albums []models.Album
	sqlLiteDb.Find(&albums)
	fmt.Println(albums)
}

func InsertEntry(album models.Album) {
	sqlLiteDb := openDb()

	result := sqlLiteDb.Create(&album)
	fmt.Printf("Inserted user with ID: %d, Rows Affected: %d\n", album.Id, result.RowsAffected)

	closeDb(sqlLiteDb)
}
