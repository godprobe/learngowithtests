package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposits", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Digidollar(10))

		got := wallet.Balance()

		want := Digidollar(10)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Withdrawals", func(t *testing.T) {
		wallet := Wallet{balance: Digidollar(50)}

		wallet.Withdraw(Digidollar(35))

		got := wallet.Balance()

		want := Digidollar(15)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
