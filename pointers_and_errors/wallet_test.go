package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Digidollar) {
		got := wallet.balance
		if got != want {
			t.Errorf("got %s, want %s", got, &want)
		}
	}

	t.Run("Deposits", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Digidollar(10))
		assertBalance(t, wallet, Digidollar(10))
	})

	t.Run("Withdrawals", func(t *testing.T) {
		wallet := Wallet{balance: Digidollar(50)}
		wallet.Withdraw(Digidollar(35))
		assertBalance(t, wallet, Digidollar(15))
	})
}
