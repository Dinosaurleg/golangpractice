package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Id     uint    `gorm:"primaryKey"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
