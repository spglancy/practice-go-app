package player

import (
	"golang-starter-pack/model"
)

type Store interface {
	GetByName(string) (*model.Player, error)
	Create(*model.Player) error
	Update(*model.Player) error
	GetByID(uint) (*model.Player, error)
	GetByEmail(string) (*model.Player, error)
}
