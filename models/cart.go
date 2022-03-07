package models

import "time"

type Cart struct {
	ID        uint      `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	UserID    uint      `gorm:"primarykey" json:"user_id" form:"user_id"`
	User      User      `gorm:"foreignkey;UserID" json:"-"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
