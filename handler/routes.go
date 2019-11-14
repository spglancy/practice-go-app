package handler

import (
	"golang-starter-pack/router/middleware"
	"golang-starter-pack/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestPlayers := v1.Group("/players")
	guestPlayers.POST("", h.SignUp)
	guestPlayers.POST("/login", h.Login)

	player := v1.Group("/player", jwtMiddleware)
	player.GET("", h.CurrentPlayer)
	player.PUT("", h.UpdatePlayer)

	items := v1.Group("/items", middleware.JWTWithConfig(
		middleware.JWTConfig{
			Skipper: func(c echo.Context) bool {
				if c.Request().Method == "GET" && c.Path() != "/api/items/feed" {
					return true
				}
				return false
			},
			SigningKey: utils.JWTSecret,
		},
	))
	items.POST("", h.CreateItem)
	items.PUT("/:slug", h.UpdateItem)
	items.DELETE("/:slug", h.DeleteItem)
	items.GET("/:slug", h.GetItem)
}
