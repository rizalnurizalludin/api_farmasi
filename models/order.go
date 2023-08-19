package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	MedID  uint `json:"med_id"`
	UserID uint `json:"user_id"`
	Qty    uint `json:"qty"`
}
