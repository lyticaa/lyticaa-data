package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type ExchangeRate struct {
	ID        int64     `db:"id"`
	Code      string    `db:"code"`
	Rate      float64   `db:"rate"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func LoadExchangeRates(db *sqlx.DB) *[]ExchangeRate {
	var exchangeRates []ExchangeRate

	query := `SELECT id, code, rate, created_at, updated_at FROM exchange_rates ORDER BY id DESC`
	_ = db.Select(&exchangeRates, query)

	return &exchangeRates
}
