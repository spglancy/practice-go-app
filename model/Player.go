package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Player to store player data and Items
type Player struct {
	gorm.Model
	Email     string `gorm:"unique_index;not null"`
	Password  string `gorm:"not null"`
	Name      string `gorm:"unique_index;not null"`
	Health    int64
	Inventory []Item `gorm:"many2many:inventory;"`
}

func (p *Player) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (p *Player) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(plain))
	return err == nil
}
