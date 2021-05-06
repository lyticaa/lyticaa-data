package sponsored

import (
	"github.com/jmoiron/sqlx"
)

type Sponsored struct {
	db *sqlx.DB
}

func NewSponsored(db *sqlx.DB) *Sponsored {
	return &Sponsored{
		db: db,
	}
}
