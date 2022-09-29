package main_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/DavidHODs/Bank/db/sqlc"
	"github.com/DavidHODs/Bank/utils"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(utils.DbDriver, utils.DbSource)
	if err != nil {
		log.Fatal("could not connect to database:", err)
	}

	utils.TestQueries = db.New(conn)

	os.Exit(m.Run())
}
