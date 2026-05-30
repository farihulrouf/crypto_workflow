migrate \
-path migrations \
-database "postgres://postgres:postgres@localhost:5432/crypto_flow?sslmode=disable" \
up