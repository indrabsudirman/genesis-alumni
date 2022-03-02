package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	var err error

	//Using PostgreSQL
	const POSTGRES = "postgresql://postgres:Indra19@localhost:5432/genesisalumni?sslmode=disable&TimeZone=Asia/Jakarta"
	dsn := POSTGRES
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	log.Println("Connected to database")

}
