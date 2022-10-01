package utils

import (
	"database/sql"
	"log"
	"os"

	db "github.com/DavidHODs/Bank/db/sqlc"
	"github.com/joho/godotenv"
)

func loadEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file: ", err)
	}

	return os.Getenv(key)
}

const (
	DbDriver = "postgres"
)

var (
	DbSource = loadEnv("dbSource")
)

var TestQueries *db.Queries
var TestDB *sql.DB
