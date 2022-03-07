package routes

import (
	"os"

	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/constants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func userRoute(e *echo.Echo, db *gorm.DB) {
	// noAuth := e.Group("/users")
	jwtAuth := e.Group("/users")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv(constants.JWT_SECRET))))
}
