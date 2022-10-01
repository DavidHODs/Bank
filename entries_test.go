package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	db "github.com/DavidHODs/Bank/db/sqlc"
	"github.com/DavidHODs/Bank/utils"
)

// creates a random detail for entry creation and testing
func createRandomEntry(t *testing.T) db.Entries {
	arg := db.CreateEntryParams{
		AccountID: utils.RandomAccountRow(),
		Amount:    utils.RandomMoney(),
	}

	entry, err := utils.TestQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}
