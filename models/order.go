package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	MedID int
	Qty   int
}
