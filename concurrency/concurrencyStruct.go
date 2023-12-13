package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account struct {
	balance int
	sync.RWMutex
}

func (a *Account) getBalance() int {
	a.RLock()         // read lock can call two routine at a time
	defer a.RUnlock() // to run after return statement
	return a.balance
}

func (a *Account) deposit(amt int) {
	a.Lock()
	a.balance = a.balance + amt
	a.Unlock()
}

func (a *Account) withdraw(amt int) error {
	a.Lock()
	defer a.Unlock()
	if amt < a.balance {
		a.balance = a.balance - amt
		return nil
	}
	return errors.New("Insufficient Balance")
}

func main() {
	acct := Account{balance: 500}
	completed := make(chan struct{})
	go func() {
		defer func() { completed <- struct{}{} }()
		err := acct.withdraw(200)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Successfully withdrawn 200 \n")
		}
	}()
	go func() {
		defer func() { completed <- struct{}{} }()
		err := acct.withdraw(400)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Successfully withdrawn 200 \n")
		}
	}()
	<-completed
	<-completed
}
