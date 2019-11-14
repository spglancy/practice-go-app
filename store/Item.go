package store

import (
	"golang-starter-pack/model"

	"github.com/jinzhu/gorm"
)

type ItemStore struct {
	db *gorm.DB
}

func NewItemStore(db *gorm.DB) *ItemStore {
	return &ItemStore{
		db: db,
	}
}

func (ps *PlayerStore) GetByEmail(e string) (*model.Player, error) {
	var m model.Player
	if err := ps.db.Where(&model.Player{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (is *ItemStore) GetBySlug(s string) (*model.Item, error) {
	var m model.Item
	err := is.db.Where(&model.Item{Slug: s}).Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, err
}

func (as *ItemStore) GetPlayerItemBySlug(playerID uint, slug string) (*model.Item, error) {
	var m model.Item
	err := as.db.Where(&model.Item{Slug: slug, PlayerID: playerID}).Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, err
}

func (is *ItemStore) CreateItem(a *model.Item) error {
	return is.db.Create(a).Error
}

func (is *ItemStore) UpdateItem(a *model.Item, tagList []string) error {
	return is.db.Model(a).Update(a).Error
}

func (is *ItemStore) DeleteItem(a *model.Item) error {
	return is.db.Delete(a).Error
}
