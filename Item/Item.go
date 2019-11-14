package item

import (
	"golang-starter-pack/model"
)

type Store interface {
	GetBySlug(string) (*model.Item, error)
	CreateItem(*model.Item) error
	UpdateItem(*model.Item, []string) error
	DeleteItem(*model.Item) error
	GetPlayerItemBySlug(uint, string) (*model.Item, error)
}
