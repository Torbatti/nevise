package models

import "time"

type Nazar struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
}
