package model

import (
	"github.com/jinzhu/gorm"
)

// Item to go in inventory of the Player
type Item struct {
	gorm.Model
	Slug     string `gorm:"unique_index;not null"`
	Title    string `gorm:"not null"`
	Ammo     int64
	Damage   int64
	Healing  int64
	PlayerID uint
	Player   Player
}
