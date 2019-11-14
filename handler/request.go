package handler

import (
	"golang-starter-pack/model"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
)

type playerUpdateRequest struct {
	Player struct {
		Email    string `gorm:"unique_index;not null"`
		Password string `gorm:"not null"`
		Name     string `gorm:"unique_index;not null"`
		Health   int64
	} `json:"player"`
}

func newPlayerUpdateRequest() *playerUpdateRequest {
	return new(playerUpdateRequest)
}

func (r *playerUpdateRequest) populate(u *model.Player) {
	r.Player.Name = u.Name
	r.Player.Email = u.Email
	r.Player.Password = u.Password
}

func (r *playerUpdateRequest) bind(c echo.Context, u *model.Player) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Name = r.Player.Name
	u.Email = r.Player.Email
	if r.Player.Password != u.Password {
		h, err := u.HashPassword(r.Player.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}
	return nil
}

type playerRegisterRequest struct {
	Player struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"player"`
}

func (r *playerRegisterRequest) bind(c echo.Context, u *model.Player) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Name = r.Player.Name
	u.Email = r.Player.Email
	h, err := u.HashPassword(r.Player.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

type playerLoginRequest struct {
	Player struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"player"`
}

func (r *playerLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type itemCreateRequest struct {
	Item struct {
		Slug    string `gorm:"unique_index;not null"`
		Title   string `gorm:"not null"`
		Ammo    int64
		Damage  int64
		Healing int64
	} `json:"item"`
}

func (r *itemCreateRequest) bind(c echo.Context, a *model.Item) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Item.Title
	a.Ammo = r.Item.Ammo
	a.Damage = r.Item.Damage
	a.Slug = slug.Make(r.Item.Title)
	return nil
}

type itemUpdateRequest struct {
	Item struct {
		Slug   string `gorm:"unique_index;not null"`
		Title  string `gorm:"not null"`
		Ammo   int64
		Damage int64
	} `json:"item"`
}

func (r *itemUpdateRequest) populate(a *model.Item) {
	r.Item.Title = a.Title
	r.Item.Ammo = a.Ammo
	r.Item.Damage = a.Damage
}

func (r *itemUpdateRequest) bind(c echo.Context, a *model.Item) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Item.Title
	a.Slug = slug.Make(a.Title)
	a.Ammo = r.Item.Ammo
	a.Damage = r.Item.Damage
	return nil
}
