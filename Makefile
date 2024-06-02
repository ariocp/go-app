build:
	docker-compose build go-app

run:
	docker-compose build go-app

migrate:
	migrate -path ./migrations -database postgres://postgres:qwerty@localhost:5432?sslmode=disable up

swag:
	swag init -g cmd/app/main.go

dev:
	go run ./cmd/app/main.go
