package models

import "time"

type Nevise struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
}
