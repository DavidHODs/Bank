package utils

import (
	"context"
	"log"
)

func TotalRows() int64 {
	num, err := TestQueries.TotalRows(context.Background())
	if err != nil {
		log.Fatal("could not get total number of account rows: err")
	}

	return num
}
