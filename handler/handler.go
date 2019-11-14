package handler

import (
	item "golang-starter-pack/item"
	player "golang-starter-pack/player"
)

type Handler struct {
	playerStore player.Store
	itemStore   item.Store
}

func NewHandler(ps player.Store, is item.Store) *Handler {
	return &Handler{
		playerStore: ps,
		itemStore:   is,
	}
}
