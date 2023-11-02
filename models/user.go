package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Hash      string
	NeviseHa  []Nevise `gorm:"foreignKey:UserRefer"`
	NazarHa   []Nazar  `gorm:"foreignKey:UserRefer"`
}
