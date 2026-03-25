package initializers

import (
	"log"

	"IkonKutz.API/models"
)

func SyncDatabase() {
	// AutoMigrate creates the table if it does not exist,
	// and adds missing columns if the model changes.
	err := DB.AutoMigrate(
		&models.Service{},
		&models.Appointment{},
	)
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("database synced successfully")
}
