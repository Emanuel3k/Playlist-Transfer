include .env

migration-up:
	@migrate -path ./database/migration -database ${DATABASE_CONNECTION_PATH} -verbose up