package main

import (
	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/config"
	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/middlewares"
	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/routes"
)

func main() {
	db := config.SetupDB()
	e := routes.New(db)

	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
