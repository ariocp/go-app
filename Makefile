build:
	docker-compose build go-app

run:
	docker-compose build go-app

migrate:
<<<<<<< HEAD
	migrate -path ./migrations -database postgres://postgres:qwerty@localhost:5432?sslmode=disable up

swag:
	swag init -g cmd/app/main.go

dev:
	go run ./cmd/app/main.go
=======
	migrate -path ./schema -database postgres://postgres:qwerty@localhost:5432?sslmode=disable up
>>>>>>> 39260269a65d547ef035ec84c6d4c737c0756251
