package pointersanderrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) WithDraw(value Bitcoin) error {
	if value > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= value
	return nil
}

func (w Wallet) Balance() (balance Bitcoin) {
	balance = w.balance
	return
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
