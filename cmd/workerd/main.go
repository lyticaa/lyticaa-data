package main

import (
	"os"

	"github.com/lyticaa/lyticaa-data/internal/worker"
	reports "github.com/lyticaa/lyticaa-data/internal/worker/reports/ingest"

	"github.com/getsentry/sentry-go"
	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent"
	"github.com/rs/zerolog/log"
)

func main() {
	sentryOpts := sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}
	err := sentry.Init(sentryOpts)
	if err != nil {
		panic(err)
	}

	config := newrelic.NewConfig(
		os.Getenv("APP_NAME"),
		os.Getenv("NEW_RELIC_LICENSE_KEY"),
	)

	nr, err := newrelic.NewApplication(config)
	if err != nil {
		panic(err)
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	var workers []worker.Worker
	workers = append(workers, reports.NewIngest(nr, db))

	logger := log.With().Str("module", os.Getenv("APP_NAME")).Logger()
	worker.StartWorkers(&logger, workers)
}
