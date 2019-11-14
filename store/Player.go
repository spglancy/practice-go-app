package store

import (
	"golang-starter-pack/model"

	"github.com/jinzhu/gorm"
)

type PlayerStore struct {
	db *gorm.DB
}

func NewPlayerStore(db *gorm.DB) *PlayerStore {
	return &PlayerStore{
		db: db,
	}
}

func (ps *PlayerStore) GetByID(id uint) (*model.Player, error) {
	var m model.Player
	if err := ps.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ps *PlayerStore) GetByName(Name string) (*model.Player, error) {
	var m model.Player
	if err := ps.db.Where(&model.Player{Name: Name}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ps *PlayerStore) Create(p *model.Player) (err error) {
	return ps.db.Create(p).Error
}

func (ps *PlayerStore) Update(p *model.Player) error {
	return ps.db.Model(p).Update(p).Error
}

func (ps *PlayerStore) Delete(p *model.Player) error {
	return ps.db.Model(p).Delete(p).Error
}
