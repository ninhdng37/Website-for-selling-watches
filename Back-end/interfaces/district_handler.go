package interfaces

import (
	"net/http"
	"watchWebsite/application"
	"watchWebsite/notifications"

	"github.com/labstack/echo/v4"
)

type Districts struct {
	d application.DistrictAppInterface
}

func NewDistricts(d application.DistrictAppInterface) *Districts {
	return &Districts{
		d: d,
	}
}

func (d *Districts) GetAllDistrict(c echo.Context) error {
	districts, err := d.d.GetAllDistrict()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}
	return c.JSON(http.StatusOK, districts)
}
