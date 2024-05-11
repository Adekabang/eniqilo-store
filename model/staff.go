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

type StaffData struct {
	UserId      string
	PhoneNumber string
	Name        string
}
type AuthenticationStaffResponse struct {
	Status  string    `json:"status"`
	Data    StaffData `json:"data"`
	Message string    `json:"message"`
}
