package notifications

const (
	CUSTOMER_NAME_MAX_LENGTH    = 50
	EMAIL_MAX_LENGTH            = 50
	PHONE_NUMBER_MAX_LENGTH     = 50
	APARTMENT_NUMBER_MAX_LENGTH = 50
	PASSWORD_MAX_LENGTH         = 20
	PASSWORD_MIN_LENGTH         = 8
)

// type Errors struct {
// 	Message []string
// }

// Error keys
const (
	INTERNAL_SERVER_ERROR     = "InternalServerError"
	BAD_REQUEST_ERROR         = "BadRequest"
	UNAUTHORIZED_ACCESS_ERROR = "UnauthorizedError"
)

// Error value
const (
	SYSTEM_ERROR               = "System error!"
	LINK_ERROR                 = "Link is incorrect!"
	UNEXIST_EMAIL_ERROR        = "Email is not registered!"
	EXISTED_EMAIL_ERROR        = "Email existed! Please use a different email!"
	EXISTED_PHONE_NUMBER_ERROR = "Phone number existed! Please use a different phone number!"
	EMAIL_NONEXISTENT_ERROR    = "Email is invalid!"
	EXPIRED_TOKEN_ERROR        = "Token is expired!"
	INVALID_TOKEN_ERROR        = "Token is invalid!"
	EMPTY_NAME_ERROR           = "Name is not empty"
	UNVERIFIED_EMAIL_ERROR     = "You registered but your email is unverified." +
		" Please access your email to verify!"
	INCORRECT_EMAIL_OR_PASSWORD = "Incorrect email or password!"
)
