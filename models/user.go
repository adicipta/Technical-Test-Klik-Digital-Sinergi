package models

import (
	"time"
)

type User struct {
	ID          int       `gorm:"primarykey;AUTO_INCREMENT"`
	Name        string    `gorm:"type:varchar(100)" json:"name"`
	Address     string    `gorm:"type:longtext;not null" json:"address"`
	Gender      string    `gorm:"type:enum('male','female');not null" json:"gender"`
	PhoneNumber string    `gorm:"type:varchar(15);unique;not null" json:"phone_number"`
	Username    string    `gorm:"type:varchar(55);unique" json:"username"`
	Password    string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type RegisterUser struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	Gender      string `json:"gender" form:"gender"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
}

type ResponseGetUser struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Username    string `json:"username"`
	Role        string `json:"role"`
}
