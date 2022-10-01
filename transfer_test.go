package main_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	db "github.com/DavidHODs/Bank/db/sqlc"
	"github.com/DavidHODs/Bank/utils"
)

// creates a random detail for transfer creation and testing
func createRandomTransfer(t *testing.T) db.Transfers {
	arg := db.CreateTransferParams{
		FromAccountID: utils.RandomAccountRow(),
		ToAccountID:   utils.RandomAccountRow(),
		Amount:        utils.RandomMoney(),
	}

	transfer, err := utils.TestQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := utils.TestQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	entry1 := createRandomTransfer(t)

	arg := db.UpdateTransferParams{
		ID:            entry1.ID,
		ToAccountID:   entry1.ToAccountID,
		FromAccountID: entry1.FromAccountID,
		Amount:        utils.RandomMoney(),
	}

	entry2, err := utils.TestQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.FromAccountID, entry2.FromAccountID)
	require.Equal(t, entry1.ToAccountID, entry2.ToAccountID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}
