package database

import (
	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/models"
	"gorm.io/gorm"
)

type LoginDB struct {
	db *gorm.DB
}

type LoginModel interface {
	GetUsername(string) int
	CreateLogin(login models.Login) (models.Login, error)
	CreateLoginStaff(login models.Login) (models.Login, error)
	GetAccountByUsername(requestlogin models.RequestLogin) (models.Login, error)
	GetLoginUserID(userID int) (models.Login, error)
	UpdateToken(id int, token string) (models.Login, error)
	UpdateLogin(id int, login models.Login) (models.Login, error)
}

func (m *LoginDB) GetUsername(username string) int {
	var login models.Login

	row := m.db.Where("username =?", username).Find(&login).RowsAffected
	return int(row)
}

func (m *LoginDB) CreateLogin(login models.Login) (models.Login, error) {
	if err := m.db.Select("username", "password", "role", "user_id").Create(&login).Error; err != nil {
		return login, err
	}

	return login, nil
}

func (m *LoginDB) CreateLoginStaff(login models.Login) (models.Login, error) {
	if err := m.db.Select("username", "password", "role", "staff_id").Create(&login).Error; err != nil {
		return login, err
	}

	return login, nil
}

func (m *LoginDB) GetAccountByUsername(requestlogin models.RequestLogin) (models.Login, error) {
	var login models.Login

	if err := m.db.Where("username = ?", requestlogin.Username).First(&login).Error; err != nil {
		return login, err
	}

	return login, nil
}

func (m *LoginDB) GetLoginUserID(userID int) (models.Login, error) {
	var login models.Login

	if err := m.db.Where("user_id = ?", userID).First(&login).Error; err != nil {
		return login, err
	}

	return login, nil
}

func (m *LoginDB) UpdateToken(id int, newToken string) (models.Login, error) {
	var login models.Login

	if err := m.db.First(&login, id).Error; err != nil {
		return login, err
	}
	login.Token = newToken

	if err := m.db.Model(&login).Update("token", newToken).Error; err != nil {
		return login, err
	}

	return login, nil
}

func (m *LoginDB) UpdateLogin(id int, newLogin models.Login) (models.Login, error) {
	var login models.Login

	if err := m.db.First(&login, id).Error; err != nil {
		return login, err
	}

	login.Username = newLogin.Username
	login.Password = newLogin.Password

	if err := m.db.Model(&login).Updates(models.Login{
		Username: login.Username,
		Password: login.Password,
	}).Error; err != nil {
		return login, err
	}

	return login, nil
}
