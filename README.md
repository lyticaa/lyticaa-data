# lyticaa-data

Lyticaa Data (process uploads).

## Setup

### Golang

If you are new to Golang, please follow the setup instructions [here](https://golang.org/doc/install).

### Environment

Before running this project, please ensure that you have the following environment variables set:

```bash
ENV=
APP_NAME=
SENTRY_DSN=
NEW_RELIC_LICENSE_KEY=
DATABASE_URL=
AWS_REGION=
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_SQS_QUEUE=
AWS_S3_UPLOAD_BUCKET=
GMT=
```

If you are unsure as to what these values ought to be, then please check with a colleague.

### Linter

To run the linter:

```bash
make lint
```

### Tests

To run the tests and see test coverage:

```bash
make tests
```

### Install

To compile and install the binary:

```bash
make install
```

### Run the Worker

```bash
make run-worker-service
```

The worker will then connect to SQS and listen for any incoming messages to process.

## Database

This project makes use of Postgres.

### Setup

To start a local Postgres instance, run:

```bash
make docker-pg
```

Then, to create the database and apply the correct role:

```bash
make create-database
make create-user
```

### Migrations

Add your migrations to the `db/migrations` folder. To apply the migrations:

```bash
make migrate
```
