module github.com/lyticaa/lyticaa-data

replace github.com/lyticaa/lyticaa-data => ../lyticaa/lyticaa-data

go 1.14

require (
	github.com/aws/aws-sdk-go v1.35.22
	github.com/bufferapp/sqs-worker-go v0.0.0-20181101064454-7e780f286181
	github.com/getsentry/sentry-go v0.7.0
	github.com/golang-migrate/migrate/v4 v4.13.0
	github.com/heroku/x v0.0.44
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.8.0
	github.com/newrelic/go-agent v3.9.0+incompatible
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/rs/zerolog v1.20.0
	github.com/shirou/gopsutil v0.0.0-20180427012116-c95755e4bcd7 // indirect
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/tealeg/xlsx v1.0.5
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
	syreclabs.com/go/faker v1.2.2
)
