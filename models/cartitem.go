package models

type CartItem struct {
	ID         int  `gorm:"primarykey;not null;AUTO_INCREMENT" json:"id" form:"id"`
	CartUserID int  `json:"cart_id" form:"cart_user_id"`
	Cart       Cart `gorm:"foreignkey:CardID" json:"-"`
	BookID     int  `json:"book_id" form:"book_id"`
	Book       Book `gorm:"foreingkey:BookID" json:"-"`
	Quantity   int  `gorm:"type:int;not null" json:"quantity" form:"quantity"`
}

type CartItem_Input struct {
	BookID   int `json:"book_id" form:"book_id"`
	Quantity int `json:"quantity" form:"quantity"`
}
