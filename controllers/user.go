package controllers

import (
	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/lib/database"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userModel  database.UserModel
	loginModel database.LoginModel
	cartModel  database.CartModel
}

func NewUserController(userModel database.UserModel, loginModel database.LoginModel, cartModel database.CartModel) *UserController {
	return &UserController{
		userModel:  userModel,
		loginModel: loginModel,
		cartModel:  cartModel,
	}
}

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
