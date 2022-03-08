package models

type CartItem struct {
	ID       int  `gorm:"primarykey;not null;AUTO_INCREMENT" json:"id" form:"id"`
	CartID   int  `json:"cart_id" form:"cart_id"`
	Cart     Cart `gorm:"foreignkey:CardID" json:"-"`
	BookID   int  `json:"book_id" form:"book_id"`
	Book     Book `gorm:"foreingkey:BookID" json:"-"`
	Quantity int
}
