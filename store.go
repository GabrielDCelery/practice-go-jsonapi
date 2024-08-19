package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	CreateAccount(*Account) error
	DeleteAccount(id string) error
	GetAccountByID(id string) (error, *Account)
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
	// defer store.conn.Close(context.Background())
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
	// defer store.conn.Close(context.Background())
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
func (store *PostgresStore) DeleteAccount(id string) error {
	return nil
}
func (store *PostgresStore) GetAccountByID(id string) (error, *Account) {
	return nil, &Account{}
}
