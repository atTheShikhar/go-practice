package main

import (
	"testing"
)

func TestWaller(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("insufficient funds error", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "insufficient balance")
	})
}

func assertBalance(tb testing.TB, wallet Wallet, want Bitcoin) {
	tb.Helper()
	got := wallet.Balance()

	if got != want {
		tb.Errorf("got %s, want %s", got, want)
	}
}

func assertError(tb testing.TB, got error, want string) {
	tb.Helper()

	if got == nil {
		tb.Fatal("wanted an error but didn't get one")
	}

	if got.Error() != want {
		tb.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(tb testing.TB, got error) {
	tb.Helper()

	if got != nil {
		tb.Fatal("got an error but didn't expected one")
	}
}
