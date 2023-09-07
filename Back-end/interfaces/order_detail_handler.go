package interfaces

import (
	"fmt"
	"net/http"
	"strconv"
	"watchWebsite/application"

	"github.com/labstack/echo/v4"
)

type OrderDetail struct {
	od application.OrderDetailAppInterface
}

func NewOrderDetail(od application.OrderDetailAppInterface) *OrderDetail {
	return &OrderDetail{
		od: od,
	}
}

func (od *OrderDetail) GetOrderDetailByOrderID(c echo.Context) error {

	orderIDString := c.Param("orderID")

	orderIDInt, err := strconv.Atoi(orderIDString)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	orderID := uint32(orderIDInt)

	orderDetails, err := od.od.GetOrderDetailByOrderID(&orderID)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orderDetails)
}
