package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// initalize the err variable here so it can be used in the if statement below.
	var err error

	DB, err = gorm.Open(postgres.Open(AppConfig.DBURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connection established")
}

// // EAsy way to connect to the dataabase.
// var DB *gorm.DB

// func ConnectToDB() {
// 	// Database connection logic here
// 	dsn := os.Getenv("URI")
// 	if dsn == "" {
// 		log.Fatal("URI is missing")
// 	}

// 	var err error

// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	log.Println("Database connection established")
// }
