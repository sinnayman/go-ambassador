package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Password         []byte `json:"-"`
	PasswordValidate string `json:"-" gorm:"-"`
	PasswordConfirm  string `json:"-" gorm:"-"`
	IsAmbassador     bool   `json:"is_ambassador"`
}

func (u *User) Validate() error {
	if u.PasswordValidate != u.PasswordConfirm {
		return errors.New("passwords don't match")
	}
	return nil
}
