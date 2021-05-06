package transaction

import (
	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	db *sqlx.DB
}

func NewTransaction(db *sqlx.DB) *Transaction {
	return &Transaction{
		db: db,
	}
}
