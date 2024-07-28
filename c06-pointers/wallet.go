package c06

import (
	"errors"
	"fmt"
)

type BitCoin int

// Implement Stringer interface
func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount BitCoin) error {
	if w.balance >= amount {
		w.balance -= amount
		return nil
	}
	return ErrInsufficientFunds
}

func NewWallet() *Wallet {
	return &Wallet{balance: BitCoin(0)}
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

var ErrUhOh = errors.New("oh critical error!!")

func check(num int) error {
	if num == 1 {
		return fmt.Errorf("it's odd")
	} else if num == 2 {
		return ErrUhOh
	}
	return nil
}

func Validate(num int) error {
	err := check(num)
	if err != nil {
		// Warp an error in the new error
		return fmt.Errorf("run error: %w", err)
	}
	return nil
}
