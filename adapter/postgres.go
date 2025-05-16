package adapter

import (
	"fmt"
	"point-of-sales-app/config"
	"point-of-sales-app/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {

	// dsn := "host=localhosttttt user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.ENV.DB_HOST,
		config.ENV.DB_USER,
		config.ENV.DB_PASSWORD,
		config.ENV.DB_NAME,
		config.ENV.DB_PORT,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&entities.User{}); err != nil {
		return err
	}

	DB = db

	return nil
}
