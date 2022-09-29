package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	db "github.com/DavidHODs/Bank/db/sqlc"
	"github.com/DavidHODs/Bank/utils"
)

func TestCreateAccount(t *testing.T) {
	arg := db.CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
