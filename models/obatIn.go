package models

import "time"

type MedIn struct {
	ID     int
	InDate time.Time
	Jumlah int
	MedID  int
}
