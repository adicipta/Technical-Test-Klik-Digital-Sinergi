package database

import (
	"errors"

	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/models"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

type UserModel interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUser([]models.ResponseGetUser, error)
	UpdateUser(id uint, newUser models.User) (models.User, error)
	GetUserProfile(id uint) (models.ResponseGetUser, error)
	GetUserByID(userID uint) (models.User, error)
}

func (m *UserDB) CreateUser(user models.User) (models.User, error) {
	if err := m.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (m *UserDB) GetAllUser() ([]models.ResponseGetUser, error) {
	var users []models.ResponseGetUser

	if err := m.db.Raw("SELECT u.id, u.name, u.address, u.gender, u.phone_number, l.username, l.role FROM users JOIN logins l ON u.id = l.user_id").Scan(&users).Error; err != nil {
		return nil, err
	} else if len(users) == 0 {
		err := errors.New("is empty")
		return nil, err
	}

	return users, nil
}

func (m *UserDB) UpdateUser(id uint, newUser models.User) (models.User, error) {
	var user models.User

	if err := m.db.First(&user, id).Error; err != nil {
		return user, err
	}
	user.Name = newUser.Name
	user.Address = newUser.Address
	user.Gender = newUser.Gender
	user.PhoneNumber = newUser.PhoneNumber
	if err := m.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (m *UserDB) GetUserProfile(id uint) (models.ResponseGetUser, error) {
	var profile models.ResponseGetUser

	if err := m.db.Raw("SELECT u.id, u.name, u.address, u.gender, u.phone_number, l.username, l.role FROM users JOIN logins l ON u.id = l.user_id WHERE user_id = ?", id).Scan(&profile).Error; err != nil {
		return profile, err
	}

	return profile, nil
}

func (m *UserDB) GetUserByID(userID uint) (models.User, error) {
	var user models.User
	if err := m.db.First(&user, userID).Error; err != nil {
		return user, err
	}

	return user, nil
}
