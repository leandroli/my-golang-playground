// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	"my-golang-playground/ch9/exercise9.1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	var ok bool
	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	go func() {
		bank.Withdraw(200)
		done <- struct{}{}
	}()

	go func() {
		ok = bank.Withdraw(1000)
		done <- struct{}{}
	}()
	// Wait for both transactions.
	<-done
	<-done
	<-done
	<-done

	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	if OverWithdrawOK := false; OverWithdrawOK != ok {
		t.Error("excessive withdrawal")
	}
}
