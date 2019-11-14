package handler

import (
	"golang-starter-pack/model"
	player "golang-starter-pack/player"
	"golang-starter-pack/utils"

	"github.com/labstack/echo/v4"
)

type playerResponse struct {
	Player struct {
		Email    string `gorm:"unique_index;not null"`
		Password string `gorm:"not null"`
		Name     string `gorm:"unique_index;not null"`
		Health   int64
		Token    string `json:"token"`
	} `json:"player"`
}

func newPlayerResponse(u *model.Player) *playerResponse {
	r := new(playerResponse)
	r.Player.Email = u.Email
	r.Player.Password = u.Password
	r.Player.Name = u.Name
	r.Player.Health = u.Health
	r.Player.Token = utils.GenerateJWT(u.ID)
	return r
}

type itemResponse struct {
	Slug    string `gorm:"unique_index;not null"`
	Title   string `gorm:"not null"`
	Ammo    int64
	Damage  int64
	Healing int64
	Player  struct {
		Email    string `gorm:"unique_index;not null"`
		Password string `gorm:"not null"`
		Name     string `gorm:"unique_index;not null"`
		Health   int64
	} `json:"player"`
}

type singleItemResponse struct {
	Item *itemResponse `json:"item"`
}

type itemListResponse struct {
	Items      []*itemResponse `json:"items"`
	ItemsCount int             `json:"itemsCount"`
}

func newItemResponse(c echo.Context, a *model.Item) *singleItemResponse {
	ar := new(itemResponse)
	ar.Slug = a.Slug
	ar.Title = a.Title
	ar.Ammo = a.Ammo
	ar.Damage = a.Damage
	ar.Player.Name = a.Player.Name
	ar.Player.Health = a.Player.Health
	return &singleItemResponse{ar}
}

func newItemListResponse(us player.Store, playerID uint, items []model.Item, count int) *itemListResponse {
	r := new(itemListResponse)
	r.Items = make([]*itemResponse, 0)
	for _, a := range items {
		ar := new(itemResponse)
		ar.Slug = a.Slug
		ar.Title = a.Title
		ar.Ammo = a.Ammo
		ar.Damage = a.Damage
		ar.Player.Name = a.Player.Name
		ar.Player.Health = a.Player.Health
		r.Items = append(r.Items, ar)
	}
	r.ItemsCount = count
	return r
}
