build:
	docker-compose build go-app

run:
	docker-compose build go-app

migrate:
	migrate -path ./migrations -database postgres://zescnmxq:pMfdF8MvPAq2nVjDjb1ugMhzmnCHFsu1@balarama.db.elephantsql.com:5432/zescnmxq?sslmode=disable up

