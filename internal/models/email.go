package models

import "time"

type Email struct {
	ID        uint
	Email     string
	CreatedAt time.Time
}
