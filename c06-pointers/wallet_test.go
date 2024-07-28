package c06_test

import (
	"errors"
	"fmt"
	"testing"

	"example.com/c06-pointers"
	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	t.Run("wallet deposits coin into empty wallet correctly", func(t *testing.T) {
		// Arrange
		wallet := c06.NewWallet()

		// Act
		wallet.Deposit(c06.BitCoin(10))

		// Assert
		actual := wallet.Balance()
		expected := c06.BitCoin(10)
		assert.Equal(t, expected, actual)
	})

	t.Run("wallet withdraws coin correctly", func(t *testing.T) {
		// Arrange
		wallet := c06.NewWallet()
		wallet.Deposit(c06.BitCoin(10))

		// Act
		_ = wallet.Withdraw(c06.BitCoin(5))

		// Assert
		actual := wallet.Balance()
		expected := c06.BitCoin(5)
		assert.Equal(t, expected, actual, "expected %s but got %s", expected, actual)
	})

	t.Run("wallet throws an error when withdrawing from insufficient funds", func(t *testing.T) {
		// Arrange
		wallet := c06.NewWallet()
		wallet.Deposit(c06.BitCoin(10))

		// Act
		actual := wallet.Withdraw(c06.BitCoin(20))

		// Assert
		assert.ErrorIs(t, actual, c06.ErrInsufficientFunds)
	})
}

// 1: some error is present .. run error: it's odd
// 2: some error is present .. run error: oh critical error!!
// 3: valid number only ..!
// 4: valid number only ..!
// 5: valid number only ..!
func TestErrorWrapping(t *testing.T) {
	for n := 1; n <= 5; n++ {
		err := c06.Validate(n)
		// Find ErrUhOh in the error, even though it is wrapped
		if errors.Is(err, c06.ErrUhOh) {
			fmt.Printf("%v: oh no something has happened! .. %+v\n", n, err)
		} else if err != nil {
			fmt.Printf("%v: some error is present .. %+v\n", n, err)
		} else {
			fmt.Printf("%v: valid number only ..!\n", n)
		}
	}
}
