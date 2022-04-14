module github.com/lyticaa/lyticaa-data

replace github.com/lyticaa/lyticaa-data => ../lyticaa/lyticaa-data

go 1.14

require (
	github.com/aws/aws-sdk-go v1.43.39
	github.com/bufferapp/sqs-worker-go v0.0.0-20181101064454-7e780f286181
	github.com/getsentry/sentry-go v0.7.0
	github.com/golang-migrate/migrate/v4 v4.13.0
	github.com/heroku/x v0.0.26
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.8.0
	github.com/newrelic/go-agent v3.9.0+incompatible
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/rs/zerolog v1.20.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.6.1
	github.com/tealeg/xlsx v1.0.5
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
	syreclabs.com/go/faker v1.2.2
)
