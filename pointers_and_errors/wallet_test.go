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

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error, but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
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

	t.Run("Withdrawals: Insufficient funds", func(t *testing.T) {
		startingBalance := Digidollar(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Digidollar(100))

		assertError(t, err, "transaction declined; insufficient funds")
		assertBalance(t, wallet, startingBalance)
	})
}
