package controllers

import (
	"net/http"

	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/lib/database"
	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	userModel  database.UserModel
	loginModel database.LoginModel
	staffModel database.StaffModel
}

func NewLoginController(userModel database.UserModel, loginModel database.LoginModel, staffModel database.StaffModel) *LoginController {
	return &LoginController{
		userModel:  userModel,
		loginModel: loginModel,
		staffModel: staffModel,
	}
}

func (controllers *LoginController) Login(c echo.Context) error {
	var requestLogin models.RequestLogin

	if err := c.Bind(&requestLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	account, err := controllers.loginModel.GetAccountByUsername(requestLogin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect username")
	}

	check := CheckPasswordHash(requestLogin.Password, account.Password)
	if !check {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect Password")
	}

	var id int
	if account.Role == "user" {
		id = account.UserID
	}else if account.Role = "staff" {
		id = account.StaffID
	}

	loginID := account.ID
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
