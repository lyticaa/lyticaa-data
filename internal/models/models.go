package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Filter struct {
	Start     int64
	Length    int64
	StartDate time.Time
	EndDate   time.Time
}

type User struct {
	UserID string `db:"user_id"`
}

type Product struct {
	SKU         string `db:"sku"`
	UserID      string `db:"user_id"`
	Description string `db:"description"`
	Marketplace string `db:"marketplace"`
}

var (
	ranges = []string{
		"today",
		"yesterday",
		"last_thirty_days",
		"previous_thirty_days",
		"this_month",
		"last_month",
		"month_before_last",
		"last_three_months",
		"previous_three_months",
		"last_six_months",
		"previous_six_months",
		"this_year",
		"last_year",
		"all_time",
	}
)

func NewFilter() *Filter {
	return &Filter{}
}

func RefreshViews(db *sqlx.DB) {
	for _, r := range ranges {
		_, _ = db.Exec(fmt.Sprintf(`REFRESH MATERIALIZED VIEW amazon_custom_transactions_%v`, r))
		_, _ = db.Exec(fmt.Sprintf(`REFRESH MATERIALIZED VIEW amazon_sponsored_products_%v`, r))
	}

	_, _ = db.Exec(`REFRESH MATERIALIZED VIEW amazon_users`)
	_, _ = db.Exec(`REFRESH MATERIALIZED VIEW amazon_products`)
}
