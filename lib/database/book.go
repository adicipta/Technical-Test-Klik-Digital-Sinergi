package database

import (
	"errors"

	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/models"
	"gorm.io/gorm"
)

type BookModel interface {
	GetBooks() ([]models.Book, error)
	GetBookById(id uint) (models.Book, error)
	AddBook(book models.Book) (models.Book, error)
	EditBookById(id uint, newBook models.Book) (models.Book, error)
	DeleteBookById(id uint) error
}

type BookDB struct {
	db *gorm.DB
}

func NewBookDB(db *gorm.DB) *BookDB {
	return &BookDB{db: db}
}

func (bdb *BookDB) GetBooks() ([]models.Book, error) {
	var books []models.Book
	if err := bdb.db.Find(&books).Error; err != nil {
		return nil, err
	} else if len(books) == 0 {
		err := errors.New("is empty")
		return nil, err
	}
	return books, nil
}

func (bdb *BookDB) GetBookById(id uint) (models.Book, error) {
	var book models.Book
	err := bdb.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (bdb *BookDB) AddBook(books models.Book) (models.Book, error) {
	if err := bdb.db.Save(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func (bdb *BookDB) EditBookById(id uint, newBook models.Book) (models.Book, error) {
	var book models.Book
	if err := bdb.db.First(&book, id).Error; err != nil {
		return book, err
	}
	book.Name = newBook.Name
	book.Description = newBook.Description
	if err := bdb.db.Save(&book).Error; err != nil {
		return book, err
	} else if book.ID == 0 {
		err := errors.New("not found")
		return book, err
	}
	return book, nil
}

func (bdb *BookDB) DeleteBookById(id uint) error {
	rows := bdb.db.Delete(&models.Book{}, id).RowsAffected
	if rows == 0 {
		err := errors.New("not found")
		return err
	}
	return nil
}
