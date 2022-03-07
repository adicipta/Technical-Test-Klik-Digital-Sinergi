package models

type CartItem struct {
	ID       uint `gorm:"primarykey;not null;AUTO_INCREMENT" json:"id" form:"id"`
	CartID   uint `json:"cart_id" form:"cart_id"`
	Cart     Cart `gorm:"foreignkey:CardID" json:"-"`
	BookID   uint `json:"book_id" form:"book_id"`
	Book     Book `gorm:"foreingkey:BookID" json:"-"`
	Quantity uint
}
