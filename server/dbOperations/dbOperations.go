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

func GetRecords() []models.Album {
	sqlLiteDb := openDb()

	var albums []models.Album
	sqlLiteDb.Find(&albums)
	fmt.Println(albums)
	closeDb(sqlLiteDb)
	return albums
}

func GetRecordByID(ID uint) models.Album {
	sqlLiteDb := openDb()

	var album models.Album
	sqlLiteDb.First(&album, ID)
	fmt.Println(album)
	closeDb(sqlLiteDb)
	return album
}

func InsertEntry(album models.Album) {
	sqlLiteDb := openDb()

	result := sqlLiteDb.Create(&album)
	fmt.Printf("Inserted user with ID: %d, Rows Affected: %d\n", album.ID, result.RowsAffected)

	closeDb(sqlLiteDb)
}

func UpdateEntry(ID uint, album models.Album) error {
	sqlLiteDb := openDb()

	fmt.Printf("WE ARE BEFOR RESULT")
	result := sqlLiteDb.Model(&models.Album{}).Where("id = ?", ID).Updates(&album)
	fmt.Printf("WE ARE AFER RESULT")

	if result.Error != nil {
		log.Fatal("failed to update recordL ", result.Error)
		return result.Error
	}
	fmt.Printf("Updated user with ID: %d", album.ID)
	closeDb(sqlLiteDb)
	return nil
}

func DeleteEntry(ID uint) error {
	sqlLiteDb := openDb()
	album := models.Album{ID: ID}
	result := sqlLiteDb.Delete(&album)

	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("Rows deleted: %d\n", result.RowsAffected)
	closeDb(sqlLiteDb)
	return nil
}
