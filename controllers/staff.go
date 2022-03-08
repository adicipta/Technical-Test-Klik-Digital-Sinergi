package controllers

import (
	"net/http"

	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/lib/database"
	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/models"
	"github.com/labstack/echo/v4"
)

type StaffController struct {
	staffModel database.StaffModel
	loginModel database.LoginModel
}

func NewStaffController(staffModel database.StaffModel, loginModel database.LoginModel) *StaffController {
	return &StaffController{
		staffModel: staffModel,
		loginModel: loginModel,
	}
}

func (controllers *StaffController) AddStaff(c echo.Context) error {
	var register models.RegisterStaff

	if err := c.Bind(&register); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	row := controllers.loginModel.GetUsername(register.Username)
	if row != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Username is already registered")
	}

	hashPassword, err := GenerateHashPassword(register.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error password in hash")
	}

	var staff models.Staff
	staff.Name = register.Name
	staff.PhoneNumber = register.PhoneNumber

	staff, err = controllers.staffModel.CreateStaff(staff)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Cannot create staff")
	}

	var login models.Login
	login.Username = register.Username
	login.Password = hashPassword
	login.Role = "staff"
	login.StaffID = staff.ID

	login, err = controllers.loginModel.CreateLogin(login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Cannot Create Staff")
	}

	var response models.ResponseGetStaff
	response.ID = staff.ID
	response.Name = staff.Name
	response.PhoneNumber = staff.PhoneNumber
	response.Username = login.Username
	response.Role = login.Role

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"data":   response,
	})
}

// func (controllers *StaffController) GetAllStaff(c echo.Context) error {
// 	staff, err := controllers.staffModel.GetAllStaff()

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status": "success",
// 		"data":   staff,
// 	})
// }
