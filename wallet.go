package main

import (
	"errors"
	"fmt"
)

// InsufficientFundsError - Error when withdraw value greater than balance
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

// Bitcoin - A Bitcoin type
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet - A Wallet sruct
type Wallet struct {
	balance Bitcoin
}

// Deposit - Deposit coin to account
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw - Withdraw coin from account
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}

// Balance - Check account balance
func (w Wallet) Balance() Bitcoin {
	return w.balance
}
