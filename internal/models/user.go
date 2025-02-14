package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Username    string    `json:"username" gorm:"column:username;type:varchar(20);not null;unique" validate:"required"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(100);not null;unique" validate:"required,email"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(15);not null;unique" validate:"required"`
	FullName    string    `json:"full_name" gorm:"column:full_name;type:varchar(100);not null" validate:"required"`
	Address     string    `json:"address" gorm:"column:address;type:text;not null" validate:"required"`
	Dob         time.Time `json:"dob" gorm:"column:dob;not null" validate:"required"`
	Password    string    `json:"password" gorm:"column:password;type:varchar(255);not null" validate:"required,min=8"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;autoUpdateTime"`
}

func (*User) TableName() string {
	return "users"
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserID              int       `json:"user_id" gorm:"column:user_id;not null;index;constraint:OnDelete:CASCADE" validate:"required"`
	Token               string    `json:"token" gorm:"column:token;type:varchar(255);not null" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"column:refresh_token;type:varchar(255);not null" validate:"required"`
	TokenExpired        time.Time `json:"token_expired" gorm:"column:token_expired;not null"`
	RefreshTokenExpired time.Time `json:"refresh_token_expired" gorm:"column:refresh_token_expired;not null"`
	CreatedAt           time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt           time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;autoUpdateTime"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
