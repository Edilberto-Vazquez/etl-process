package migrations

import (
	"log"

	"github.com/Edilberto-Vazquez/weathercloud-etl-process/db/conexion"
	"github.com/Edilberto-Vazquez/weathercloud-etl-process/db/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate() {
	m := gormigrate.New(conexion.DBCon, gormigrate.DefaultOptions, []*gormigrate.Migration{})

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&models.ElectricField{},
			&models.WeatherCloud{},
		)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")
}
