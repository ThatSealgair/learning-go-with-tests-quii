package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	fmt.Printf("Address of balance in test is '%p'\n", &wallet.balance)

	want := Bitcoin(10)

	if got != want {
		t.Errorf("Got '%s' want '%s'.", got, want)
	}
}
