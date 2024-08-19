package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store interface {
	CreateAccount(*Account) error
	DeleteAccount(id string) error
	GetAccountByID(id string) (error, *Account)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (error, *PostgresStore) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err, nil
	}
	if err := db.Ping(); err != nil {
		return err, nil
	}
	return nil, &PostgresStore{
		db: db,
	}
}

func (store *PostgresStore) CreateAccount(account *Account) error {
	return nil
}
func (store *PostgresStore) DeleteAccount(id string) error {
	return nil
}
func (store *PostgresStore) GetAccountByID(id string) (error, *Account) {
	return nil, &Account{}
}
