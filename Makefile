run:
	docker-compose up --build

down:
	docker-compose down --build

migrate:
	docker-compose exec db psql -U postgres -d mini_ecommerce -f /app/migrations/init.sql

lint:
	golangci-lint run

test:
	go test ./...
