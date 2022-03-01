package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			t.Errorf("received error %q", err)
		}
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingbilance := Bitcoin(20)
		wallet := Wallet{startingbilance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingbilance)
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
