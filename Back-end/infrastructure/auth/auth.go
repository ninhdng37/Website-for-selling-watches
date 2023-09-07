package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
	"watchWebsite/application"
	"watchWebsite/infrastructure/security"
	"watchWebsite/notifications"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	cus application.CustomerAppInterface
	em  application.EmployeeAppInterface
}

func NewAuth(cus application.CustomerAppInterface,
	em application.EmployeeAppInterface) *Auth {
	return &Auth{
		cus: cus,
		em:  em,
	}
}

type TokenInterface interface {
	SendEmailToGmail(next echo.HandlerFunc) echo.HandlerFunc
	SendEmailForForgotingPassword(c echo.Context) error
	VerifyEmail(c echo.Context) error
	VerifyEmailForForgottingPassword(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type customerForBinding struct {
	CustomerName string `json:"customerName"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	Password     string `json:"password"`
}

// Auth implements the TokenInterface
var _ TokenInterface = &Auth{}

func sendEmail(tokenString, action string, to []string) error {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("serverEmail"),
		os.Getenv("password_for_verifying_email"),
		os.Getenv("hostEmail"),
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	var verificationLink string
	if action == "create" {
		verificationLink = "http://localhost:8888/customer/verify-email/" + tokenString
	} else {
		verificationLink = "http://localhost:8888/customer/verify-email-forgotting-password/" + tokenString
	}

	msg := "Subject: Your verification code for registrying\n" + headers +
		"\n\n<h2>Link for verification:</h2><h1>" + verificationLink +
		"</h1>\n <p>Please do not share this link for anyone. Link have time to access with 1 minute.</p>"

	err := smtp.SendMail(
		os.Getenv("addrEmail"),
		auth,
		os.Getenv("serverEmail"),
		to,
		[]byte(msg),
	)

	if err != nil {
		return err
	}
	return nil
}

func (a *Auth) SendEmailToGmail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cust := &customerForBinding{}
		notices := make(map[string]string)
		var tokenString string

		if err := c.Bind(&cust); err != nil {
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}

		c.Set("customerName", cust.CustomerName)
		c.Set("email", cust.Email)
		c.Set("phoneNumber", cust.PhoneNumber)
		c.Set("password", cust.Password)

		customer, err := a.cus.GetCustomerByEmail(cust.Email)
		if customer != nil && customer.IsVerified {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.EXISTED_EMAIL_ERROR
			return c.JSON(http.StatusBadRequest, notices)
		}

		if err != nil && err != sql.ErrNoRows {
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}

		customer, err = a.cus.GetCustomerByPhoneNumber(cust.PhoneNumber)
		if customer != nil && customer.IsVerified {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.EXISTED_PHONE_NUMBER_ERROR
			return c.JSON(http.StatusBadRequest, notices)
		}

		if err != nil && err != sql.ErrNoRows {
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}
		// c.Set("isVerified", false)
		tokenString, err = createToken(cust, 1, "")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		var to []string
		to = append(to, cust.Email)
		err = sendEmail(tokenString, "create", to)

		if err != nil {
			return c.JSON(http.StatusOK, map[string]string{"SendingEmail": "Failed"})
		}

		if customer != nil {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.UNVERIFIED_EMAIL_ERROR
			return c.JSON(http.StatusBadRequest, notices)
		}
		return next(c)
	}
}

func (a *Auth) SendEmailForForgotingPassword(c echo.Context) error {
	bindingCust := struct {
		Email string `json:"email"`
	}{}
	notices := make(map[string]string)
	if err := c.Bind(&bindingCust); err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	customer, err := a.cus.GetCustomerByEmail(bindingCust.Email)
	cust := &customerForBinding{}

	if err != nil {
		if err == sql.ErrNoRows {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.EMAIL_NONEXISTENT_ERROR
			return c.JSON(http.StatusBadRequest, notices)
		}
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	cust.Email = strings.TrimSpace(customer.Email)
	cust.CustomerName = strings.TrimSpace(customer.CustomerName)
	cust.PhoneNumber = strings.TrimSpace(customer.PhoneNumber)
	tokenString, err := createToken(cust, 10, "forgot-password")

	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	var to []string
	to = append(to, cust.Email)
	err = sendEmail(tokenString, "forgot", to)

	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	notices[notifications.SUCCESS] = notifications.SUCCESS
	return c.JSON(http.StatusOK, notices)
}

func (a *Auth) VerifyEmail(c echo.Context) error {

	secretKey := os.Getenv("secret_key")
	notices := make(map[string]string)
	tokenString := c.Param("tokenString")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				notices[notifications.BAD_REQUEST_ERROR] = notifications.EXPIRED_TOKEN_ERROR
				return c.Redirect(http.StatusSeeOther, "http://localhost:8888/customer/expired-token")
			}
		}
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)

		email := claims["email"].(string)
		customerName := claims["customerName"].(string)
		phoneNumber := claims["phoneNumber"].(string)
		// action := claims["action"].(string)
		cust := &customerForBinding{
			Email:        email,
			CustomerName: customerName,
			PhoneNumber:  phoneNumber,
		}

		customer, err := a.cus.GetCustomerByEmail(email)

		if err != nil {
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}

		customer.Email = strings.TrimSpace(customer.Email)
		customer.CustomerName = strings.TrimSpace(customer.CustomerName)
		customer.PhoneNumber = strings.TrimSpace(customer.PhoneNumber)

		if customer.Email == cust.Email &&
			customer.CustomerName == cust.CustomerName &&
			customer.PhoneNumber == cust.PhoneNumber &&
			!customer.IsVerified {
			// if action == "forgot-password" {
			// 	return c.Redirect(http.StatusSeeOther, "http://localhost:8888/customer/reset-password")
			// }
			customer.ApartmentNumber = "-1"
			customer.CustomerName = "-1"
			customer.Email = "-1"
			customer.Password = "-1"
			customer.PhoneNumber = "-1"
			var zero uint32 = 0
			customer.ProvinceId = &zero
			customer.WardId = &zero
			customer.DistrictId = &zero
			customer.IsVerified = true
			err := a.cus.UpdateCustomer(customer)
			if err != nil {
				notices[notifications.INTERNAL_SERVER_ERROR] = err.Error()
				return c.JSON(http.StatusInternalServerError, notices)
			}
		}

	}

	notices[notifications.SUCCESS] = notifications.SUCCESS_VERIFICATION
	return c.Redirect(http.StatusSeeOther, "http://localhost:8888/customer/success-verification")
}

func (a *Auth) VerifyEmailForForgottingPassword(c echo.Context) error {
	notices := make(map[string]string)
	tokenString := c.Param("tokenString")
	token, err := IsTokenValid(tokenString)
	if err != nil {
		if err == echo.ErrUnauthorized {
			c.Redirect(http.StatusSeeOther, "http://localhost:8888/customer/expired-token")
			notices[notifications.UNAUTHORIZED_ACCESS_ERROR] = notifications.UNAUTHORIZED_ACCESS_ERROR
			return c.JSON(http.StatusUnauthorized, notices)
		} else if err == echo.ErrBadRequest {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.INVALID_TOKEN_ERROR
			return c.JSON(http.StatusBadRequest, "")
		}
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	claims := token.Claims.(jwt.MapClaims)

	email := claims["email"].(string)
	customerName := claims["customerName"].(string)
	phoneNumber := claims["phoneNumber"].(string)
	cust := &customerForBinding{
		Email:        email,
		CustomerName: customerName,
		PhoneNumber:  phoneNumber,
	}

	customer, err := a.cus.GetCustomerByEmail(email)

	if err != nil {

		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	customer.Email = strings.TrimSpace(customer.Email)
	customer.CustomerName = strings.TrimSpace(customer.CustomerName)
	customer.PhoneNumber = strings.TrimSpace(customer.PhoneNumber)

	if customer.Email != cust.Email ||
		customer.CustomerName != cust.CustomerName ||
		customer.PhoneNumber != cust.PhoneNumber ||
		!customer.IsVerified {
		c.Redirect(http.StatusSeeOther, "http://localhost:8888/customer/expired-token")
		notices[notifications.UNAUTHORIZED_ACCESS_ERROR] = notifications.UNAUTHORIZED_ACCESS_ERROR
		return c.JSON(http.StatusUnauthorized, notices)
	}
	cookie, err := c.Cookie("token")
	if err != nil && err != http.ErrNoCookie {

		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	// cookie := http.Cookie{}
	if cookie == nil {
		cookie = &http.Cookie{}
	}
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(10 * time.Minute)
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Domain = "localhost"

	c.SetCookie(cookie)

	notices[notifications.SUCCESS] = notifications.SUCCESS_VERIFICATION

	return c.Redirect(http.StatusSeeOther,
		"http://localhost:8888/customer/reset-password-page")
}

func (a *Auth) Login(c echo.Context) error {
	account := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	notices := make(map[string]string)

	if err := c.Bind(&account); err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	customer, err := a.cus.GetCustomerByEmailAndPassword(account.Email, account.Password)

	if err != nil {

		if err == sql.ErrNoRows {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.INCORRECT_EMAIL_OR_PASSWORD
			return c.JSON(http.StatusBadRequest, notices)
		}
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	cust := &customerForBinding{
		Email:        customer.Email,
		PhoneNumber:  customer.PhoneNumber,
		CustomerName: customer.CustomerName,
	}
	tokenString, err := createToken(cust, 4320, "")
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour), // Set expiration time for the cookie
		HttpOnly: true,                           // Make the cookie accessible only through HTTP (not JavaScript)
		Path:     "/customer",                    // The cookie is valid for all paths on the domain
		Domain:   "localhost",                    // Set to your domain
	}
	c.SetCookie(&cookie)

	notices[notifications.SUCCESS] = notifications.SUCCESS_LOGIN

	return c.JSON(http.StatusOK, notices)
}

func (a *Auth) LoginForManager(c echo.Context) error {

	notices := make(map[string]string)
	account := struct {
		EmployeeId *uint32 `json:"employeeId"`
		Password   string  `json:"password"`
	}{}

	if err := c.Bind(&account); err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	employee, err := a.em.GetEmployeeByIdAndPassword(account.EmployeeId, account.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.INCORRECT_EMAIL_OR_PASSWORD
			return c.JSON(http.StatusInternalServerError, notices)
		}
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	employee.TrimSpaces()
	notices[notifications.SUCCESS] = notifications.SUCCESS_LOGIN
	return c.JSON(http.StatusOK, employee)
}

func (a *Auth) Logout(c echo.Context) error {
	cookie, err := c.Cookie("token")
	notices := make(map[string]string)
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	// cookie.Expires = time.Now().Add(1 * time.Second)
	// fmt.Println(cookie.Expires)
	cookie.Value = ""
	// cookie.Expires = time.Now().Add(-1 * time.Second)
	c.SetCookie(cookie)

	return nil
}

func createToken(cust *customerForBinding, expiredTime uint, action string) (string, error) {
	// Create a new token with the desired signing method (e.g., HS256)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = cust.Email
	claims["customerName"] = cust.CustomerName
	claims["phoneNumber"] = cust.PhoneNumber
	claims["isCustomer"] = true
	claims["exp"] = time.Now().Add(time.Duration(expiredTime) * time.Minute).Unix() // Token expires in 1 minutes
	if action == "forgot-password" {
		claims["action"] = action
	}

	secretKey := os.Getenv("secret_key")
	// Sign the token with your secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsTokenValid(tokenString string) (*jwt.Token, error) {
	secretKey := os.Getenv("secret_key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, echo.ErrUnauthorized
			}
		}
		fmt.Println(err)
		return nil, echo.ErrInternalServerError
	}

	if !token.Valid {
		return nil, echo.ErrBadRequest
	}

	return token, nil
}

func (a *Auth) ResetPassword(c echo.Context) error {
	data := struct {
		Password string `json:"password"`
	}{}

	notices := make(map[string]string)

	if err := c.Bind(&data); err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	// cookie, err := c.Cookie("token")
	// if err != nil {
	// 	return err
	// }

	cookie := &http.Cookie{}

	cookies := c.Cookies()
	for _, v := range cookies {
		if v.Value != "" && v.Name == "token" {
			cookie = v
		}
	}

	tokenString := cookie.Value
	token, err := IsTokenValid(tokenString)
	if err != nil {
		if err == echo.ErrBadRequest {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.BAD_REQUEST_ERROR
			return c.JSON(http.StatusBadRequest, notices)
		}
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	customer, err := a.cus.GetCustomerByEmail(email)

	if err != nil {
		if err == sql.ErrNoRows {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.EMAIL_NONEXISTENT_ERROR
			return c.JSON(http.StatusBadRequest, notices)
		}
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	customer.Password = security.Hash(data.Password)
	err = a.cus.UpdateCustomer(customer)
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	return c.JSON(http.StatusOK, cookie)
}
