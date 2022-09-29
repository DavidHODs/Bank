package main_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/DavidHODs/Bank/db/sqlc"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file: ", err)
	}

	return os.Getenv(key)
}

const (
	dbDriver = "postgres"
)

var (
	dbSource = loadEnv("dbSource")
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("could not connect to database:", err)
	}

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
