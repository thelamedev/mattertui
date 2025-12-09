.PHONY: migrate-up migrate-down sqlc db-reset
	
build-all: build-server build-tui
	
build-server:
	go build -o ./bin/server ./cmd/server
	
build-tui:
	go build -o ./bin/tui ./cmd/tui
	
run-server:
	go run ./cmd/server
	
run-tui:
	go run ./cmd/tui

migrate-up:
	go run cmd/migrate/migrate.go -direction up

migrate-down:
	go run cmd/migrate/migrate.go -direction down

sqlc:
	sqlc generate

db-reset: migrate-down migrate-up
