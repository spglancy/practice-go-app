package handler

import (
	"net/http"

	"golang-starter-pack/model"
	"golang-starter-pack/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetItem(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.itemStore.GetBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, newItemResponse(c, a))
}

func (h *Handler) CreateItem(c echo.Context) error {
	var a model.Item
	req := &itemCreateRequest{}
	if err := req.bind(c, &a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	err := h.itemStore.CreateItem(&a)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, newItemResponse(c, &a))
}

func (h *Handler) UpdateItem(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.itemStore.GetPlayerItemBySlug(playerIDFromToken(c), slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := &itemUpdateRequest{}
	req.populate(a)
	if err := req.bind(c, a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newItemResponse(c, a))
}

func (h *Handler) DeleteItem(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.itemStore.GetPlayerItemBySlug(playerIDFromToken(c), slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	err = h.itemStore.DeleteItem(a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}
