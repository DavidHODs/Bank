package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Randomly generates a random integer between min and max value
func RandomInteger(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Randomly generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates random amount of money
func RandomMoney() int64 {
	return RandomInteger(0, 1000)
}

// RandomCurrency generates a random currency unit
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "NGN", "CAD"}
	return currencies[rand.Intn(len(currencies))]
}

// RandomAccountRow returns a random account id from the created accounts list
func RandomAccountRow() int64 {
	return RandomInteger(1, TotalRows())
}
