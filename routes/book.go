package routes

import (
	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/controllers"
	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/lib/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func bookRoute(e *echo.Echo, db *gorm.DB) {
	bdb := database.NewBookDB(db)
	bc := controllers.NewBookController(bdb)
	e.GET("/books", bc.GetBooks)
	e.POST("/books", bc.AddBook)
	e.PUT("/books/:id", bc.EditBooks)
	e.DELETE("/books/:id", bc.DeleteBooks)
}
