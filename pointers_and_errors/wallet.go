package main

import "fmt"

type Digidollar int

type Wallet struct {
	balance Digidollar
}

func (w *Wallet) Deposit(amount Digidollar) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Digidollar) {
	w.balance -= amount
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
