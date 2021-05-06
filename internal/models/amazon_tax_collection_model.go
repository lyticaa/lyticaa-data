package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonTaxCollectionModel struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func LoadAmazonTaxCollectionModels(db *sqlx.DB) *[]AmazonTaxCollectionModel {
	var taxCollectionModels []AmazonTaxCollectionModel

	query := `SELECT id, name, created_at, updated_at FROM amazon_tax_collection_models ORDER BY id DESC`
	_ = db.Select(&taxCollectionModels, query)

	return &taxCollectionModels
}

func LoadAmazonTaxCollectionModel(name string, db *sqlx.DB) *AmazonTaxCollectionModel {
	var taxCollectionModel AmazonTaxCollectionModel

	query := `SELECT * FROM amazon_tax_collection_models WHERE name = $1`
	_ = db.QueryRow(query, name).Scan(
		&taxCollectionModel.ID,
		&taxCollectionModel.Name,
		&taxCollectionModel.CreatedAt,
		&taxCollectionModel.UpdatedAt,
	)

	return &taxCollectionModel
}
