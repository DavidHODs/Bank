package main_test

import (
	"context"
	"testing"

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
