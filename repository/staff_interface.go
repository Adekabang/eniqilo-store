package repository

import "github.com/Adekabang/eniqilo-store/model"

type StaffRepositoryInterface interface {
	RegisterStaff(model.RegisterStaff) model.AuthenticationStaffResponse
	LoginStaff(model.LoginStaff) bool
}
