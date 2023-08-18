package models

import (
	"time"
)

type MedOut struct {
	ID      int
	OutDate time.Time
	Jumlah  int
	MedID   int
}
