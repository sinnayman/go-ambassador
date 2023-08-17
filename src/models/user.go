package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
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

func (u *User) SetPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	u.Password = hashed
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}
