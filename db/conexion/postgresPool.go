package conexion

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon       *gorm.DB
	sqlDB       *sql.DB
	err         error
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_PORT     = os.Getenv("DB_PORT")
)

func InitCon() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}

func CloseCon() {
	sqlDB, err = DBCon.DB()
	if err != nil {
		log.Println(err)
	}
	close := sqlDB.Close()
	if close != nil {
		log.Println(close)
	} else {
		log.Println("conexion close with the DB")
	}
}
