package main

import (
	"math/rand"
)

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int
	Balance   int
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        rand.Int(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    getAccounts(),
		Balance:   0,
	}
}
func getAccounts() int {
	return 1
}
