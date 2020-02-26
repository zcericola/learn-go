package main

import (
	"errors"
	"fmt"
)

//Bitcoin is the type of currency in the wallet
type Bitcoin int

//Stringer stringifies Bitcoin
type Stringer interface {
	String() string
}

//String is a method on the Stringer interface that returns a formatted string
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Wallet is a type that holds a balance of currency
type Wallet struct {
	balance Bitcoin
}

//Deposit is a method on Wallet that adds money
func (w *Wallet) Deposit(amount Bitcoin) {
	/*Deposit takes a pointer to Wallet as the input because
	otherwise, the amount we add to balance will be to a copy of balance
	and not to the one that the test is checking against
	*/
	(*w).balance += amount
	/*using (*w) dereferences the pointer so we can read it
	however it is not necessary to write it like this because golang will
	do it implicitly for you
	*/

}

//ErrInsufficientFunds is an error message
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Withdraw removes Bitcoin from the Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	(*w).balance -= amount
	return nil

}

//Balance is a method on Wallet that reads the balance
func (w *Wallet) Balance() Bitcoin {
	/*Same here, we use a pointer because we want to see
	the balance of the original wallet that we have in memory and not a copy
	*/
	return w.balance //this still returns correctly without having to type (*w)

}

func main() {

}
