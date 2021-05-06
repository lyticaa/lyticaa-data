GOBIN?=${GOPATH}/bin

all: lint install

lint-pre:
	@test -z $(gofmt -l .)
	@go mod verify

lint: lint-pre
	@golangci-lint run

lint-verbose: lint-pre
	@golangci-lint run -v

install: go.sum
	GO111MODULE=on go install -v ./cmd/workerd

clean:
	rm -f ${GOBIN}/workerd

tests:
	@go test -v -coverprofile .testCoverage.txt ./...

run-worker-service:
	@workerd

docker-pg:
	@docker-compose -f ./build/docker-compose.yml run --rm -p 5432:5432 --no-deps pg

docker-memcache:
	@docker-compose -f ./build/docker-compose.yml run --rm -p 11211:11211 --no-deps -d memcache

docker-rabbitmq:
	@docker-compose -f ./build/docker-compose.yml run --rm -p 4369:4369 -p 5671:5671 -p 5672:5672 -p 25672:25672 --no-deps -d rabbitmq

create-user:
	PGPASSWORD=password psql -h localhost -U postgres -c "CREATE USER lyticaa WITH SUPERUSER PASSWORD 'password';"

create-database:
	PGPASSWORD=password psql -h localhost -U postgres -c "CREATE DATABASE lyticaa_development OWNER lyticaa;"

drop-database:
	PGPASSWORD=password psql -h localhost -U postgres -c "drop database lyticaa_development;"

migrate:
	@go run cmd/migrate/main.go
