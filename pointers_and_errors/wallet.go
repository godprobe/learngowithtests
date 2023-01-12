package main

import (
	"errors"
	"fmt"
)

type Digidollar int

type Wallet struct {
	balance Digidollar
}

func (w *Wallet) Deposit(amount Digidollar) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("transaction declined, insufficient funds")

func (w *Wallet) Withdraw(amount Digidollar) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Digidollar {
	return w.balance
}

type Stringer interface {
	String() string
}

func (d Digidollar) String() string {
	return fmt.Sprintf("%d DGD", d)
}
