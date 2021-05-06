package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonFulfillment struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func LoadAmazonFulfillments(db *sqlx.DB) *[]AmazonFulfillment {
	var fulfillments []AmazonFulfillment

	query := `SELECT id, name, created_at, updated_at FROM amazon_fulfillments ORDER BY id DESC`
	_ = db.Select(&fulfillments, query)

	return &fulfillments
}

func LoadAmazonFulfillment(name string, db *sqlx.DB) *AmazonFulfillment {
	var fulfillment AmazonFulfillment

	query := `SELECT * FROM amazon_fulfillments WHERE name = $1`
	_ = db.QueryRow(query, name).Scan(
		&fulfillment.ID,
		&fulfillment.Name,
		&fulfillment.CreatedAt,
		&fulfillment.UpdatedAt,
	)

	return &fulfillment
}
