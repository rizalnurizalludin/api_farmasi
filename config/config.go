package config

import (
	"fmt"
	"gudang-obat/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dbroot := os.Getenv("DB_ROOT")
	dbpass := os.Getenv("DB_PASS")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbroot, dbpass, dbhost, dbport, dbname)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	initMigration()

}

func initMigration() {
	DB.AutoMigrate(&models.Med{}, models.MedIn{}, models.MedOut{}, &models.User{})
}
