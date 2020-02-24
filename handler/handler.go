package handler

import (
	"back_core/user"
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct {
	userStore    user.Store
}

func (h *Handler) test(c echo.Context) error {
	return c.JSON(http.StatusCreated, "GOLA KURCINA")

}

func NewHandler(us user.Store) *Handler {
	return &Handler{
		userStore:    us,
	}
}