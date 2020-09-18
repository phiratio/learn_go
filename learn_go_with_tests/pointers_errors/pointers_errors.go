package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

// create method String on bitcoin type
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//  this interface is defined in the fmt package and
// lets you define how your type is printed when used with the %s format string in prints.
// https://golang.org/pkg/fmt/#Stringer
//type Stringer interface {
//	String() string
//}

type Wallet struct {
	balance Bitcoin
}

// In Go, when you call a function or a method the arguments are copied.
// wallet  =>  0xc0000 => "balance is here"
// operates with pointer of Wallet struct type
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in wallet is %v \n", &w.balance) // get the data this pointer has
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
