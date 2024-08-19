package main

import (
	"github.com/jackc/pgx"
)

type Store interface {
	CreateAccount(*Account) error
	DeleteAccount(id string) error
	GetAccountByID(id string) (error, *Account)
}

type PostgresStore struct {
	conn *pgx.Conn
}

func NewPostgresStore() (error, *PostgresStore) {
	connConfig := pgx.ConnConfig{
		Host:     "localhost",
		Database: "postgres",
		Port:     5432,
		User:     "postgres",
		Password: "gobank",
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		return err, nil
	}
	return nil, &PostgresStore{
		conn: conn,
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
