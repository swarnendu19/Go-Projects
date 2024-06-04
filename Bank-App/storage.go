package main

import "database/sql"

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostGresStore struct {
	db *sql.DB
}

func NewPostGresStore() (*PostGresStore, error) {

}
