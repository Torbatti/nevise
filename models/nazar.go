package models

import "time"

type Nazar struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	Text        string `json:"nazar_text"`
	UserRefer   uint
	NeviseRefer uint
}
