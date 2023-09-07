package interfaces

import (
	"net/http"
	"watchWebsite/application"
	"watchWebsite/notifications"

	"github.com/labstack/echo/v4"
)

type Wards struct {
	w application.WardAppInterface
}

func NewWards(w application.WardAppInterface) *Wards {
	return &Wards{
		w: w,
	}
}

func (w *Wards) GetAllWard(c echo.Context) error {
	wards, err := w.w.GetAllWard()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}
	return c.JSON(http.StatusOK, wards)
}
