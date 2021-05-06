package amazon

import (
	"github.com/lyticaa/lyticaa-data/internal/worker/reports/ingest/amazon/reports"

	"github.com/jmoiron/sqlx"
)

type Amazon struct {
	Reports *reports.Reports
	db      *sqlx.DB
}

func NewAmazon(db *sqlx.DB) *Amazon {
	return &Amazon{
		Reports: reports.NewReports(db),
		db:      db,
	}
}
