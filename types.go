package main

import "github.com/google/uuid"

type Account struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Balance   int    `json:"balance"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Balance:   0,
	}
}
