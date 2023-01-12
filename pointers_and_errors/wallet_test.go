package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
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

	t.Run("Withdrawals: Insufficient funds", func(t *testing.T) {
		startingBalance := Digidollar(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Digidollar(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Digidollar) {
	t.Helper()
	got := wallet.balance

	if got != want {
		t.Errorf("got %s, want %s", got, &want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error, but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
