package db

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	
	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable",
						os.Getenv("PGUSER"),
						os.Getenv("PGPASSWORD"),
						os.Getenv("PGDATABASE"),
  )

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	err = DB.AutoMigrate(&User{}, &Video{}, &Like{}, &WatchHistory{})
	if err != nil {
			log.Fatal("Failed to migrate DB: ", err)
	}



}
