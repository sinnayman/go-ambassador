package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email" gorm:"unique"`
	Password         []byte `json:"-"`
	IsAmbassador     bool   `json:"is_ambassador"`
	PasswordValidate string `json:"-" gorm:"-"`
	PasswordConfirm  string `json:"-" gorm:"-"`
}

type UserWrite struct {
	gorm.Model
	User
}

func (UserWrite) TableName() string {
	return "users"
}

type UserRead struct {
	ID int `json:"id"`
	User
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (UserRead) TableName() string {
	return "users"
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
