package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id"`
	Token     string         `json:"token"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy uint           `json:"created_by"`
	UpdatedAt time.Time      `json:"updated_at"`
	UpdatedBy uint           `json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"-"`
	ExpiredAt time.Time      `json:"expired_at"`
}
