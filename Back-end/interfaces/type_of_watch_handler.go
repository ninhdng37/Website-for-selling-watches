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

type TypeOfWatch struct {
	typeOfWatch application.TypeOfWatchAppInterface
}

func NewTypeOfWatch(typeOfWatch application.TypeOfWatchAppInterface) *TypeOfWatch {
	return &TypeOfWatch{
		typeOfWatch: typeOfWatch,
	}
}

func (b *TypeOfWatch) GetAllTypesOfWatch(c echo.Context) error {
	typesOfWatch, err := b.typeOfWatch.GetAllTypesOfWatch()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}

	for _, typeOfWatch := range typesOfWatch {
		typeOfWatch.TypeOfWatchName = strings.TrimSpace(typeOfWatch.TypeOfWatchName)
	}

	return c.JSON(http.StatusOK, typesOfWatch)
}
func (b *TypeOfWatch) CreateTypeOfWatch(c echo.Context) error {
	typeOfWatch := struct {
		TypeOfWatchName string `json:"typeOfWatchName"`
	}{}

	if err := c.Bind(&typeOfWatch); err != nil {
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}

	err := b.typeOfWatch.CreateTypeOfWatch(typeOfWatch.TypeOfWatchName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, notifications.SUCCESS)
}

func (b *TypeOfWatch) UpdateTypeOfWatch(c echo.Context) error {
	typeOfWatch := &entity.TypeOfWatch{}

	if err := c.Bind(&typeOfWatch); err != nil {
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}

	err := b.typeOfWatch.UpdateTypeOfWatch(typeOfWatch)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}
	return c.JSON(http.StatusOK, notifications.SUCCESS)
}

func (b *TypeOfWatch) DeleteTypeOfWatch(c echo.Context) error {
	typeOfWatchIDString := c.Param("typeOfWatchID")

	typeOfWatchIDInt, err := strconv.Atoi(typeOfWatchIDString)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	typeOfWatchID := uint32(typeOfWatchIDInt)
	fmt.Println(typeOfWatchID)
	err = b.typeOfWatch.DeleteTypeOfWatch(&typeOfWatchID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, notifications.SUCCESS)
}
