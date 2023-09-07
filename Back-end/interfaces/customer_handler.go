package interfaces

import (
	"fmt"
	"net/http"
	"watchWebsite/application"
	"watchWebsite/domain/entity"

	"github.com/labstack/echo/v4"
)

type Customers struct {
	cus application.CustomerAppInterface
}

func NewCustomers(cus application.CustomerAppInterface) *Customers {
	return &Customers{
		cus: cus,
	}
}

func (cust *Customers) CreateCustomer(c echo.Context) error {
	var customer entity.Customer
	customer.CustomerName = c.Get("customerName").(string)
	customer.Email = c.Get("email").(string)
	customer.PhoneNumber = c.Get("phoneNumber").(string)
	customer.Password = c.Get("password").(string)
	//validate request
	validateErr := customer.Validate("create")
	if len(validateErr) > 0 {
		return c.JSON(http.StatusUnprocessableEntity, validateErr)
	}

	err := cust.cus.CreateCustomer(&customer)
	if err != nil {
		fmt.Println("3")
		return c.JSON(http.StatusInternalServerError, err)
	}

	notice := "Created an account successfully!"
	return c.JSON(http.StatusCreated, notice)
}
