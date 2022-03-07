package models

import "time"

type Book struct {
	ID          uint      `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Name        string    `gorm:"type:varchar(100);unique;not null" json:"name" form:"name"`
	Description string    `gorm:"type:longtext;not null" json:"description" name:"description"`
	Quantity    uint      `gorm:"type:uint;not null" json:"quantity" name:"quantity"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
