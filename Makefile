include .env
export

MIGRATION_DIR = ./db/migrations

.PHONY: migrate-new migrate-up migrate-down sqlc-gen

migrate-new:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database "sqlite3://$(DB_PATH)" up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database "sqlite3://$(DB_PATH)" down

sqlc-gen:
	sqlc generate