package models

import "gorm.io/gorm"

type Med struct {
	gorm.Model
	Name  string `json:"name"`
	Type  string `json:"type"`
	Price uint   `json:"price"`
	Stock uint   `json:"stock"`
}
