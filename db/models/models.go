package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type ElectricField struct {
	gorm.Model
	TimeStamp      time.Time `gorm:"<-:create;not null"`
	TimeStampRound time.Time `gorm:"<-:create;not null"`
	Lightning      bool      `gorm:"<-:create;not null"`
	Distance       uint8     `gorm:"<-:create;not null"`
	ElectricField  float32   `gorm:"<-:create;not null"`
	RotorStatus    bool      `gorm:"<-:create;not null"`
}

type WeatherCloud struct {
	gorm.Model
	TimeStamp time.Time       `gorm:"<-:create;not null"`
	TempIn    sql.NullFloat64 `gorm:"<-:create"`
	Temp      sql.NullFloat64 `gorm:"<-:create"`
	Chill     sql.NullFloat64 `gorm:"<-:create"`
	DewIn     sql.NullFloat64 `gorm:"<-:create"`
	Dew       sql.NullFloat64 `gorm:"<-:create"`
	HeatIn    sql.NullFloat64 `gorm:"<-:create"`
	Heat      sql.NullFloat64 `gorm:"<-:create"`
	Humin     sql.NullFloat64 `gorm:"<-:create"`
	Hum       sql.NullFloat64 `gorm:"<-:create"`
	Wspdhi    sql.NullFloat64 `gorm:"<-:create"`
	Wspdavg   sql.NullFloat64 `gorm:"<-:create"`
	Wdiravg   sql.NullFloat64 `gorm:"<-:create"`
	Bar       sql.NullFloat64 `gorm:"<-:create"`
	Rain      sql.NullFloat64 `gorm:"<-:create"`
	RainRate  sql.NullFloat64 `gorm:"<-:create"`
}
