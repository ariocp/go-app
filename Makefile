build:
	docker-compose build go-app

run:
	docker-compose build go-app

migrate:
	migrate -path ./schema -database postgres://postgres:qwerty@localhost:5432?sslmode=disable up