package interfaces

import (
	"fmt"
	"net/http"
	"strconv"
	"watchWebsite/application"
	"watchWebsite/notifications"

	"github.com/labstack/echo/v4"
)

type Order struct {
	o application.OrderAppInterface
}

func NewOrder(o application.OrderAppInterface) *Order {
	return &Order{
		o: o,
	}
}

func (o *Order) UpdateOrderByCustomer(c echo.Context) error {
	notices := make(map[string]string)

	orderIDString := c.QueryParam("orderID")
	orderIDTemp, err := strconv.Atoi(orderIDString)
	orderID := uint32(orderIDTemp)
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	err = o.o.UpdateOrderByCustomer(&orderID)

	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	notices[notifications.SUCCESS] = notifications.SUCCESS
	return c.JSON(http.StatusOK, notices)
}

func (o *Order) GetAllOrders(c echo.Context) error {
	orders, err := o.o.GetAllOrders()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for _, order := range orders {
		order.TrimSpaces()
	}

	return c.JSON(http.StatusOK, orders)
}

func (o *Order) UpdateOrderStatus(c echo.Context) error {

	order := struct {
		OrderID    uint32 `json:"orderID"`
		Status     uint32 `json:"status"`
		EmployeeID uint32 `json:"employeeID"`
	}{}

	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(order.OrderID)
	fmt.Println(order.EmployeeID)

	err := o.o.UpdateOrderStatus(&order.OrderID, &order.EmployeeID, &order.Status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, notifications.SUCCESS)
}
