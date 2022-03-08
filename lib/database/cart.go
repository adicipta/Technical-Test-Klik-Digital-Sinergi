package database

import (
	"errors"

	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/models"
	"gorm.io/gorm"
)

type CartDB struct {
	db *gorm.DB
}

func NewCartDB(db *gorm.DB) *CartDB {
	return &CartDB{db: db}
}

type CartModel interface {
	CreateCart(cart models.Cart) error
	AddToCart(userID int, input models.CartItem_Input) (models.CartItem, error)
	GetCartItem(userID int) ([]models.CartItem, error)
	EditCartItem(cartItemID int, input models.CartItem_Input) (models.CartItem, error)
	DeleteCartItem(cartItemID int) error
	CheckCartByBookID(userID int, bookID int) bool
	GetItemInCart(userID, bookID int) (models.CartItem, error)
}

func (cdb *CartDB) CreateCart(cart models.Cart) error {
	if err := cdb.db.Save(&cart).Error; err != nil {
		return err
	}

	return nil
}

func (cdb *CartDB) AddToCart(userID int, input models.CartItem_Input) (models.CartItem, error) {
	var cartItem models.CartItem

	cartItem.CartUserID = userID
	cartItem.BookID = input.BookID
	cartItem.Quantity = input.Quantity
	if err := cdb.db.Select("cart_user_id", "book_id", "quantity").Create(&cartItem).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (cdb *CartDB) GetCartItem(userID int) ([]models.CartItem, error) {
	var cartItems []models.CartItem

	if err := cdb.db.Find(&cartItems, "cart_user_id = ? IS NULL", userID).Error; err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (cdb *CartDB) EditCartItem(cartItemID int, input models.CartItem_Input) (models.CartItem, error) {
	var cartItems models.CartItem

	if err := cdb.db.First(&cartItems, cartItemID).Update("quantity", input.Quantity).Error; err != nil {
		return cartItems, err
	} else if cartItems.ID == 0 {
		err := errors.New("not found")
		return cartItems, err
	}

	return cartItems, nil
}

func (cdb *CartDB) DeleteCartItem(cartItemID int) error {
	rows := cdb.db.Delete(&models.CartItem{}, cartItemID).RowsAffected
	if rows == 0 {
		err := errors.New("cart item is not found")
		return err
	}

	return nil
}

func (cdb *CartDB) CheckCartByBookID(userID int, bookID int) bool {
	var cartItem models.CartItem

	row := cdb.db.Where("cart_user_id = ? and book_id = ? IS NULL", userID, bookID).Find(&cartItem).RowsAffected
	if row == 1 {
		return true
	}

	return false
}

func (cdb *CartDB) GetItemInCart(userID, bookID int) (models.CartItem, error) {
	var cartItem models.CartItem

	if err := cdb.db.Where("cart_user_id = ? and book_id = ? IS NULL", userID, bookID).First(&cartItem).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}
