package model

type RegisterStaff struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type LoginStaff struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}
