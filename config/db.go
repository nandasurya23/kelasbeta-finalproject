package config

import (
	"fmt"
	"kelasbeta/finalproject/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

var Postgres PostgresDB

func OpenDB() {
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PASS"))

	postgresConn, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Postgres = PostgresDB{
		DB: postgresConn,
	}

	err = autoMigrate(postgresConn)
	if err != nil {
		log.Fatal(err)
	}
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Question{},
		&models.Answer{},
	)

	if err != nil {
		return err
	}

	return nil
}
