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

	// token belom
	return true
}

func (m *StaffRepository) LoginStaff(payload model.LoginStaff) bool {
	// response := model.ResponseMessage{Status: "failed", Msg: "user not found"}
	query, err := m.Db.Query("SELECT * FROM staff WHERE phone_number = $1", payload.PhoneNumber)
	if err != nil {
		log.Println(err)
		return false
		// response = model.ResponseMessage{Status: "failed", Msg: "server error"}
	}
	defer query.Close()
	if query != nil {
		for query.Next() {
			var (
				id            string
				created_at    string
				email         string
				name          string
				password_hash string
			)
			err := query.Scan(&id, &created_at, &email, &name, &password_hash)
			if err != nil {
				log.Println(err)
			}
			err2 := utils.VerifyPassword(payload.Password, password_hash)
			if err2 != nil {
				log.Println(err2)
				return false
				// response = model.ResponseMessage{Status: "failed", Msg: "wrong password"}
			} else {
				// token, err := utils.GenerateToken(id)
				// if err != nil {
				// 	log.Println(err)
				// 	response = model.ResponseMessage{Status: "failed", Msg: "server error"}
				// }
				return true
				// response = model.ResponseMessage{Status: "success", Msg: token, Data: model.UserData{Email: email, Name: name}}
			}

		}
	}
	// return response
	return true
}
