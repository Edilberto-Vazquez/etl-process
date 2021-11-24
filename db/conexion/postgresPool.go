package conexion

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
	sqlDB *sql.DB
	err   error
)

func InitCon() {
	dsn := "host= user= password= dbname=inaoe port=5432 sslmode=disable"
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
