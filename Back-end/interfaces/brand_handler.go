package interfaces

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"watchWebsite/application"
	"watchWebsite/domain/entity"
	"watchWebsite/notifications"

	"github.com/labstack/echo/v4"
)

type Brand struct {
	brand application.BrandAppInterface
}

func NewBrands(brand application.BrandAppInterface) *Brand {
	return &Brand{
		brand: brand,
	}
}

func (b *Brand) GetAllBrands(c echo.Context) error {
	brands, err := b.brand.GetAllBrands()

	for _, brand := range brands {
		brand.BrandName = strings.TrimSpace(brand.BrandName)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}
	return c.JSON(http.StatusOK, brands)
}

func (b *Brand) CreateBrand(c echo.Context) error {
	brand := struct {
		BrandName string `json:"brandName"`
	}{}

	if err := c.Bind(&brand); err != nil {
		fmt.Println(1)
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err := b.brand.CreateBrand(brand.BrandName)

	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, notifications.SUCCESS)
}

func (b *Brand) UpdateBrand(c echo.Context) error {

	brand := &entity.Brand{}

	if err := c.Bind(&brand); err != nil {
		fmt.Println(1)
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}

	err := b.brand.UpdateBrand(brand)

	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}

	return c.JSON(http.StatusOK, notifications.SUCCESS)
}

func (b *Brand) DeleteBrand(c echo.Context) error {

	brandIDString := c.Param("brandID")
	brandIDInt, err := strconv.Atoi(brandIDString)

	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}

	brandID := uint32(brandIDInt)
	err = b.brand.DeleteBrand(&brandID)

	if err != nil {
		if err.Error() == "can not delete brand because this brand was assigned for watches" {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		fmt.Println(2)
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, notifications.SUCCESS)
}
