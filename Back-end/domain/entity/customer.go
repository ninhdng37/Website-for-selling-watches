package entity

import (
	"regexp"
	"strconv"
	"watchWebsite/notifications"
)

type Customer struct {
	CustomerId      *uint32 `json:"customerId"`
	CustomerName    string  `json:"customerName"`
	Email           string  `json:"email"`
	PhoneNumber     string  `json:"phoneNumber"`
	Password        string  `json:"password"`
	ProvinceId      *uint32 `json:"provinceId"`
	DistrictId      *uint32 `json:"districtId"`
	WardId          *uint32 `json:"wardId"`
	ApartmentNumber string  `json:"apartmentNumber"`
	IsVerified      bool    `json:"isVerified"`
	// VerificationToken string  `json:"verificationToken"`
}

func isValidEmail(email string) bool {
	// Regular expression pattern for email validation
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression pattern
	regex := regexp.MustCompile(emailPattern)

	// Use the MatchString method to check if the email matches the pattern
	return regex.MatchString(email)
}

func (cust *Customer) Validate(action string) map[string]string {
	var errMessages = make(map[string]string)

	if len(cust.CustomerName) > notifications.CUSTOMER_NAME_MAX_LENGTH {
		errMessages["customer_length_error"] = "Customer name have not to exceed " +
			strconv.Itoa(notifications.CUSTOMER_NAME_MAX_LENGTH) + "  characters."
	}

	if len(cust.Email) > notifications.EMAIL_MAX_LENGTH {
		errMessages["email_length_error"] = "Email have not to exceed " +
			strconv.Itoa(notifications.EMAIL_MAX_LENGTH) + "  characters."
	}

	if len(cust.PhoneNumber) > notifications.PHONE_NUMBER_MAX_LENGTH {
		errMessages["phone_number_length_error"] = "Phone number have not to exceed " +
			strconv.Itoa(notifications.PHONE_NUMBER_MAX_LENGTH) + "  characters."
	}

	if len(cust.Password) < notifications.PASSWORD_MIN_LENGTH &&
		len(cust.Password) > notifications.PASSWORD_MAX_LENGTH {
		errMessages["password_length_error"] = "Password only contains from " +
			strconv.Itoa(notifications.PASSWORD_MIN_LENGTH) + "to " +
			strconv.Itoa(notifications.PASSWORD_MAX_LENGTH) + "  characters."
	}

	if !isValidEmail(cust.Email) {
		// fmt.Println(cust.Email)
		errMessages["invalid_email"] = "Email is not valid"
	}

	return errMessages
}
