package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Digidollar(10))

	got := wallet.Balance()

	want := Digidollar(9000)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
