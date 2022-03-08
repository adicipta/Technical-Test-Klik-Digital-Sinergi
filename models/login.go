package models

import "time"

type Login struct {
	ID        int       `gorm:"primarykey;AUTO_INCREMENT"`
	Username  string    `gorm:"type:varchar(55);unique" json:"username"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	Role      string    `gorm:"type:enum('staff','user')" json:"role"`
	UserID    int       `gorm:"type:uint"`
	User      User      `gorm:"foreignkey:UserID;"`
	Token     string    `gorm:"type:longtext;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type RequestLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type ResponseLogin struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
