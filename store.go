package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	CreateAccount(*Account) error
	DeleteAccountByID(id string) error
	GetAccountByID(id string) (error, *Account)
	GetAccounts() (error, []*Account)
}

type PostgresStore struct {
	conn *pgxpool.Pool
}

func NewPostgresStore() (error, *PostgresStore) {
	const (
		host     = "localhost"
		port     = 5432
		username = "postgres"
		password = "gobank"
		database = "postgres"
		sslmode  = "disable"
	)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", username, password, host, port, database, sslmode)
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return err, nil
	}
	return nil, &PostgresStore{
		conn: pool,
	}
}

func (store *PostgresStore) Init() error {
	err := store.createAccountTable()
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresStore) createAccountTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS accounts(
			id VARCHAR(255),
			first_name VARCHAR(255),
			last_name VARCHAR(255),
			balance INT,
			created_at timestamp default current_timestamp
		);
	`
	_, err := store.conn.Query(context.Background(), sql)
	return err
}

func (store *PostgresStore) CreateAccount(account *Account) error {
	sql := `
		INSERT INTO accounts VALUES(
			$1,
			$2,
			$3,
			$4
		);
	`
	_, err := store.conn.Query(context.Background(), sql, account.ID, account.FirstName, account.LastName, account.Balance)
	return err
}

func (store *PostgresStore) DeleteAccountByID(id string) error {
	sql := `
		DELETE FROM accounts WHERE id = $1
	`
	_, err := store.conn.Query(context.Background(), sql, id)
	return err
}

func (store *PostgresStore) GetAccountByID(id string) (error, *Account) {
	sql := `
		SELECT id, first_name, last_name, balance, created_at FROM accounts WHERE id = $1;
	`
	account := &Account{}
	err := store.conn.QueryRow(context.Background(), sql, id).Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		return err, &Account{}
	}
	return nil, account
}

func (store *PostgresStore) GetAccounts() (error, []*Account) {
	return nil, []*Account{}
}
