package models

import "time"

type Staff struct {
	ID          uint      `gorm:"primarykey;AUTO_INCREMENT"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	PhoneNumber string    `gorm:"type:varchar(15);unique;not null" json:"phone_number"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type RegisterStaff struct {
	Name        string `json:"name" form:"name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
}

type ResponseGetStaff struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Username    string `json:"username"`
	Role        string `json:"role"`
}
