DB_URL=postgres://postgres:postgres@localhost:5432/crypto_flow?sslmode=disable

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

migration:
	migrate create -ext sql -dir migrations -seq $(name)