package repository

import (
	"database/sql"
	"log"

	"github.com/Adekabang/eniqilo-store/model"
	"github.com/Adekabang/eniqilo-store/utils"
	"github.com/google/uuid"
)

type StaffRepository struct {
	Db *sql.DB
}

func NewStaffRepository(db *sql.DB) StaffRepositoryInterface {
	return &StaffRepository{Db: db}
}

// RegisterStaff implements UserRepositoryInterface
func (m *StaffRepository) RegisterStaff(payload model.RegisterStaff) bool {

	hashedPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		log.Println(err.Error())
		return false
	}

	uuidUser := uuid.New()

	stmt, err := m.Db.Prepare("INSERT INTO staff (id, phone_number, name, password_hash) VALUES ($1,$2,$3,$4)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(uuidUser, payload.PhoneNumber, payload.Name, hashedPassword)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}
