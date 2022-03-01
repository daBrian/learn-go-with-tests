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

	t.Run("Withdraww with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingbilance := Bitcoin(20)
		wallet := Wallet{startingbilance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFund)
		assertBalance(t, wallet, startingbilance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
