package administration

import (
	"errors"
)

var (
	FailGet         = errors.New("sensor : Data Not Found")
	FailInsert      = errors.New("sensor : Fail Input Data")
	FailInsertExist = errors.New("sensor : Data Already Exist")
	FailUpdate      = errors.New("sensor : Fail Update Data")
	FailDelete      = errors.New("sensor : Fail Delete Data")
)

type Users struct {
	Id    int64  `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
	City  string `json:"city"`
}

type UsersInsert struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
	City  string `json:"city"`
}

type UsersService interface {
	Create(r *UsersInsert) error
	//Get(int) (*Users, error)
	//Update(int, Users) error
	//Remove(int) error
}
