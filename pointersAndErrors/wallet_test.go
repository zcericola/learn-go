package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit money into the wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw money from the wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)} //declares a wallet with 20 BTC
		err := wallet.Withdraw(Bitcoin(10))    //removes 10
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

//helpers declared below so the assertions are the focus

//helper function that abstracts reused code into a separate fn
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

//helper function to declare errors
func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("error was expected but did not occur") //t.Fatal throws an Error immediately
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want) //t.Errorf will just log?
	}

}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but did not expect one")
	}
}
