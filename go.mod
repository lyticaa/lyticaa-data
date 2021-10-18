module github.com/lyticaa/lyticaa-data

replace github.com/lyticaa/lyticaa-data => ../lyticaa/lyticaa-data

go 1.14

require (
	github.com/aws/aws-sdk-go v1.35.22
	github.com/bufferapp/sqs-worker-go v0.0.0-20181101064454-7e780f286181
	github.com/cockroachdb/cockroach-go v0.0.0-20190925194419-606b3d062051 // indirect
	github.com/getsentry/sentry-go v0.7.0
	github.com/golang-migrate/migrate/v4 v4.15.1
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/heroku/x v0.0.26
	github.com/jmoiron/sqlx v1.3.1
	github.com/lib/pq v1.10.0
	github.com/newrelic/go-agent v3.9.0+incompatible
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/rs/zerolog v1.20.0
	github.com/snowflakedb/glog v0.0.0-20180824191149-f5055e6f21ce // indirect
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/tealeg/xlsx v1.0.5
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	syreclabs.com/go/faker v1.2.2
)
