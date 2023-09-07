package interfaces

import (
	"fmt"
	"net/http"
	"strconv"
	"watchWebsite/application"
	"watchWebsite/domain/entity"
	"watchWebsite/notifications"

	"github.com/labstack/echo/v4"
)

type Employee struct {
	e application.EmployeeAppInterface
}

func NewEmployee(e application.EmployeeAppInterface) *Employee {
	return &Employee{
		e: e,
	}
}

func (e *Employee) CreateEmployee(c echo.Context) error {
	notices := make(map[string]string)
	em := struct {
		Fullname       string `json:"fullname"`
		IdentityNumber string `json:"identityNumber"`
		Position       string `json:"position"`
		Email          string `json:"email"`
		PhoneNumber    string `json:"phoneNumber"`
	}{}

	if err := c.Bind(&em); err != nil {
		fmt.Println(1)
		fmt.Println(err)
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	employee := &entity.Employee{}
	employee.Fullname = em.Fullname
	employee.IdentityNumber = em.IdentityNumber
	employee.Position = em.Position
	employee.Email = em.Email
	employee.PhoneNumber = em.PhoneNumber
	// Define the characters that can be used in the password
	// characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// // Initialize a random seed using the current time
	// rand.NewSource(time.Now().UnixNano())

	// // Generate the random password
	// password := make([]byte, 8)
	// for i := 0; i < 8; i++ {
	// 	password[i] = characters[rand.Intn(len(characters))]
	// }
	employee.Password = "12345678"

	employeeId, err := e.e.CreateEmployee(employee)
	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	employee.EmployeeID = employeeId
	notices[notifications.SUCCESS] = notifications.SUCCESS_CREATING_EMPLOYEE
	return c.JSON(http.StatusOK, notices)
}

func (e *Employee) GetAllEmployee(c echo.Context) error {
	employees, err := e.e.GetAllEmployee()
	notices := make(map[string]string)

	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	for _, emp := range employees {
		emp.TrimSpaces()
	}

	return c.JSON(http.StatusOK, employees)
}

func (e *Employee) UpdateEmployee(c echo.Context) error {
	emp := struct {
		EmployeeID     *uint32 `json:"employeeId"`
		Fullname       string  `json:"fullname"`
		IdentityNumber string  `json:"identityNumber"`
		Position       string  `json:"position"`
		Email          string  `json:"email"`
		PhoneNumber    string  `json:"phoneNumber"`
	}{}
	notices := make(map[string]string)

	if err := c.Bind(&emp); err != nil {
		fmt.Println(1)
		fmt.Println(err)
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	employee := &entity.Employee{}

	employee.EmployeeID = emp.EmployeeID
	employee.Fullname = emp.Fullname
	employee.IdentityNumber = emp.IdentityNumber
	employee.Position = emp.Position
	employee.Email = emp.Email
	employee.PhoneNumber = emp.PhoneNumber

	err := e.e.UpdateEmployee(employee)

	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		notices[notifications.INTERNAL_SERVER_ERROR] = err.Error()
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	notices[notifications.SUCCESS] = notifications.SUCCESS
	return c.JSON(http.StatusOK, notices)
}

func (e *Employee) DeleteEmployee(c echo.Context) error {
	employeeID, err := strconv.Atoi(c.Param("employeeID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error")
	}
	ID := uint32(employeeID)
	err = e.e.DeleteEmployee(&ID)

	if err != nil {
		if err.Error() == "Can not delete employee because this employee approved orders" {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, notifications.SYSTEM_ERROR)
	}

	return c.JSON(http.StatusOK, "Success")
}
