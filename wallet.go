package main

// Wallet - A wallet sruct
type Wallet struct {
	balance int
}

// Deposit - Deposit coin to account
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance - Check account balance
func (w Wallet) Balance() int {
	return w.balance
}
