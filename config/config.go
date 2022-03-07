package config

import (
	"fmt"
	"log"
	"os"

	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	errEnv := godotenv.Load("./.env")
	if errEnv != nil {
		log.Fatal("Failed to load .env file")
	}
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	var err error
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigration(db)
	return db
}

func InitMigration(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Staff{})
	db.AutoMigrate(&models.Login{})
}
