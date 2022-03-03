package migration

import (
	"genesis-alumni/database"
	"genesis-alumni/model/entity"
	"log"
)

func Migration() {
	errMigration := database.DB.AutoMigrate(&entity.AlumnusName{})
	if errMigration != nil {
		log.Println("Error while migrate: ", errMigration)
	}
	log.Println("Database success migrated")
}
