package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonMarketplace struct {
	ID             int64     `db:"id"`
	Name           string    `db:"name"`
	ExchangeRateID int64     `db:"exchange_rate_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func LoadAmazonMarketplaces(db *sqlx.DB) *[]AmazonMarketplace {
	var marketplaces []AmazonMarketplace

	query := `SELECT id, name, exchange_rate_id, created_at, updated_at FROM amazon_marketplaces ORDER BY id DESC`
	_ = db.Select(&marketplaces, query)

	return &marketplaces
}

func LoadAmazonMarketplace(name string, db *sqlx.DB) *AmazonMarketplace {
	var marketplace AmazonMarketplace

	query := `SELECT * FROM amazon_marketplaces WHERE name = $1`
	_ = db.QueryRow(query, name).Scan(
		&marketplace.ID,
		&marketplace.Name,
		&marketplace.ExchangeRateID,
		&marketplace.CreatedAt,
		&marketplace.UpdatedAt,
	)

	return &marketplace
}
