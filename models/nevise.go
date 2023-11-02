package models

import "time"

type Nevise struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Title     string `json:"nevise_title"`
	Text      string `json:"nevise_text"`
	IsSave    bool   `json:"is_draft"`
	UserRefer uint
	NazarHa   []Nazar `gorm:"foreignKey:NeviseRefer"`
}
