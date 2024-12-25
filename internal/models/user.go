package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          int       `json:"id" gorm:"primary_key;auto_increment"`
	Username    string    `json:"username" gorm:"column:username;type:varchar(25);unique;not null" validate:"required"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(100);unique;not null" validate:"required,email"`
	Password    string    `json:"password" gorm:"column:password;type:varchar(255);not null" validate:"required"`
	Fullname    string    `json:"fullname" gorm:"column:fullname;type:varchar(100);not null" validate:"required"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(15);not null" validate:"required"`
	Address     string    `json:"address" gorm:"column:address;type:text"`
	Dob         string    `json:"dob" gorm:"column:dob;type:date"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (*User) TableName() string {
	return "users"
}

func (l *User) validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  uint      `json:"id" gorm:"primarykey"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	UserID              int       `json:"user_id" gorm:"type:int;" validate:"required"`
	AccessToken         string    `json:"access_token" gorm:"type:varchar(255);" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:varchar(255);" validate:"required"`
	AccessTokenExpired  time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l *UserSession) validate() error {
	v := validator.New()
	return v.Struct(l)
}
