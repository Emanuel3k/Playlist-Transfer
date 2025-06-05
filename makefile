include .env

migration-up:
	@migrate -path ./database/migration -database ${DATABASE_CONNECTION_PATH} -verbose up
	
test:
	@go test ./...

build :
	@go build ./cmd/http/main.go