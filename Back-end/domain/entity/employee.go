package entity

import "strings"

type Employee struct {
	EmployeeID     *uint32 `json:"employeeId"`
	Fullname       string  `json:"fullname"`
	IdentityNumber string  `json:"identityNumber"`
	Position       string  `json:"position"`
	Email          string  `json:"email"`
	PhoneNumber    string  `json:"phoneNumber"`
	Password       string  `json:"password"`
}

func (e *Employee) TrimSpaces() {
	e.Fullname = strings.TrimSpace(e.Fullname)
	e.IdentityNumber = strings.TrimSpace(e.IdentityNumber)
	e.Position = strings.TrimSpace(e.Position)
	e.Email = strings.TrimSpace(e.Email)
	e.PhoneNumber = strings.TrimSpace(e.PhoneNumber)
	e.Password = strings.TrimSpace(e.Password)
}
