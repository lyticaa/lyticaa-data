package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonTransactionType struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func LoadAmazonTransactionTypes(db *sqlx.DB) *[]AmazonTransactionType {
	var transactionTypes []AmazonTransactionType

	query := `SELECT id, name, created_at, updated_at FROM amazon_transaction_types ORDER BY id DESC`
	_ = db.Select(&transactionTypes, query)

	return &transactionTypes
}

func LoadAmazonTransactionType(name string, db *sqlx.DB) *AmazonTransactionType {
	var transactionType AmazonTransactionType

	query := `SELECT * FROM amazon_transaction_types WHERE name = $1`
	_ = db.QueryRow(query, name).Scan(
		&transactionType.ID,
		&transactionType.Name,
		&transactionType.CreatedAt,
		&transactionType.UpdatedAt,
	)

	return &transactionType
}
