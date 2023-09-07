package interfaces

import (
	"net/http"
	"strconv"
	"strings"
	"watchWebsite/application"
	"watchWebsite/domain/entity"
	"watchWebsite/notifications"

	"github.com/labstack/echo/v4"
)

type Provinces struct {
	p application.ProvinceAppInterface
}

func NewProvinces(p application.ProvinceAppInterface) *Provinces {
	return &Provinces{
		p: p,
	}
}

func (p *Provinces) GetAllProvince(c echo.Context) error {
	provinces, err := p.p.GetAllProvince()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}

	for _, province := range provinces {
		province.ProvinceName = strings.TrimSpace(province.ProvinceName)
	}
	return c.JSON(http.StatusOK, provinces)
}

func (p *Provinces) CreateProvince(c echo.Context) error {
	province := struct {
		ProvinceName string `json:"provinceName"`
	}{}

	if err := c.Bind(&province); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err := p.p.CreateProvince(province.ProvinceName)
	if err != nil {
		if err.Error() == "this province name existed" {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, notifications.SUCCESS)
}

func (p *Provinces) UpdateProvince(c echo.Context) error {
	province := &entity.Province{}

	if err := c.Bind(&province); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err := p.p.UpdateProvince(province)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, notifications.SUCCESS)
}

func (p *Provinces) DeleteProvince(c echo.Context) error {
	provinceIDString := c.Param("provinceID")

	provinceIDInt, err := strconv.Atoi(provinceIDString)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	provinceID := uint32(provinceIDInt)

	err = p.p.DeleteProvince(&provinceID)

	if err != nil {
		if err.Error() == "can not delete province because this province was assigned for another district" {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, notifications.SUCCESS)
}
