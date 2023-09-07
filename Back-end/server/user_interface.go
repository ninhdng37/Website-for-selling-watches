package server

import (
	"net/http"
	"watchWebsite/application"
	"watchWebsite/notifications"

	"github.com/labstack/echo/v4"
)

type Interface struct {
	w application.WatchAppInterface
}

func NewInterface(w application.WatchAppInterface) *Interface {
	return &Interface{
		w: w,
	}
}

type Interfaces interface {
	GetHomeInterface(c echo.Context) error
}

// var _ Interfaces = &Interface{}

func GetLoginInterface(c echo.Context) error {
	return c.File("D:/DoAnThucTap/Front-end/LoginForm.html")
}

func GetHomeInterface(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	// fmt.Println(cookie)
	c.SetCookie(cookie)
	return c.File("D:/DoAnThucTap/Front-end/home.html")
}

func GetRegistryInterface(c echo.Context) error {
	return c.File("D:/DoAnThucTap/Front-end/RegistryForm.html")
}

func GetExpiredTokenInterface(c echo.Context) error {
	return c.File("D:/DoAnThucTap/Front-end/ExpiredTokenForm.html")
}

func GetSuccessVerificationInterface(c echo.Context) error {
	return c.File("D:/DoAnThucTap/Front-end/SuccessVerificationForm.html")
}

func GetForgotingPasswordInterface(c echo.Context) error {
	return c.File("D:/DoAnThucTap/Front-end/ForgottingPasswordForm.html")
}

func GetResettingPasswordInterface(c echo.Context) error {
	cookie, err := c.Cookie("token")
	notices := make(map[string]string)
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	c.SetCookie(cookie)
	c.File("D:/DoAnThucTap/Front-end/ResettingPassword.html")
	notices[notifications.SUCCESS] = notifications.SUCCESS
	return c.JSON(http.StatusOK, notices)
}

func GetLoginInterfaceForManager(c echo.Context) error {
	return c.File("D:/DoAnThucTap/Front-end/LoginFormForManager.html")
}

func GetLoginInterfaceForEmployee(c echo.Context) error {
	return c.File("D:/DoAnThucTap/Front-end/LoginFormForEmployee.html")
}

func GetHomeInterfaceForManager(c echo.Context) error {
	return c.File("D:/DoAnThucTap/Front-end/HomeForManager.html")
}

func GetHomeInterfaceForEmployee(c echo.Context) error {
	return c.File("D:/DoAnThucTap/Front-end/HomeForEmployee.html")
}
