package config

import (
	"FINALPROJECT/model"
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "Satwikayoga"
    dbname   = "Random_question"

)

func OpenDB()  {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("You connected to your database.")

}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Category{},
	)

	if err != nil {
		return err
	}

	return nil
}
