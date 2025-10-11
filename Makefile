.PHONY: migrate-up
migrate-up:
	goose -dir ./db/migrations postgres "postgres://user:pass@localhost:5432/library?sslmode=disable" up
.PHONY: migrate-down
migrate-down:
	goose -dir ./db/migrations postgres "postgres://user:pass@localhost:5432/library?sslmode=disable" down
