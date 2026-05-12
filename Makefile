include .env
export

postgres-up:
	docker compose up -d kinotower-postgres

postgres-down:
	docker compose down kinotower-postgres

migrate-create:
	docker compose run --rm kinotower-postgres-migrate create -ext sql -dir /migrations -seq "$(name)"


migrate-up:
	docker compose run --rm kinotower-postgres-migrate -path /migrations -database postgresql://postgres:123456@kinotower-postgres:5432/kinotower?sslmode=disable up

migrate-down:
	docker compose run --rm kinotower-postgres-migrate -path /migrations -database postgresql://postgres:123456@kinotower-postgres:5432/kinotower?sslmode=disable down