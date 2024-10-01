run:
	@go run ./cmd/app/main.go

migrateup:
	migrate -path migrations -database "postgres://postgres:postgres@192.168.56.103:5432/music?sslmode=disable" up

migratedown:
	migrate -path migrations -database "postgres://postgres:postgres@192.168.56.103:5432/music?sslmode=disable" down

