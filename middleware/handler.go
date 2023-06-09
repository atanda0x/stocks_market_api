package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	MESSAGE string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := gotdotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")
	return db

}
