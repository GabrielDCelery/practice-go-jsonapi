package main

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	FirstName string
	LastName  string
}

type Account struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Balance:   0,
		CreatedAt: time.Now(),
	}
}
