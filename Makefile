include .env

up:
	docker compose up -d

down:
	docker compose down

bash:
	docker compose exec -it app /bin/bash

run:
	docker compose exec -it app go run .

tidy:
	docker compose exec -it app go mod tidy

test:
	docker compose exec -it app go test -count=1 ./...

swag-init:
	docker compose exec -it app swag init

swag-fmt:
	docker compose exec -it app swag fmt

wire:
	docker compose exec -it app wire gen

# ex.) make migrate-create name=foo
migrate-create:
	docker compose exec -it app migrate create -dir ./migrations -ext sql -seq -digits 6 ${name}

migrate-up:
	docker compose exec -it app migrate -database "postgres://${PG_USERNAME}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_DATABASE}?sslmode=${PG_SSLMODE}" -path ./migrations up

migrate-down:
	docker compose exec -it app migrate -database "postgres://${PG_USERNAME}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_DATABASE}?sslmode=${PG_SSLMODE}" -path ./migrations down -all

# ex.) make migrate-force version=1
migrate-force:
	docker compose exec -it app migrate -database "postgres://${PG_USERNAME}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_DATABASE}?sslmode=${PG_SSLMODE}" -path ./migrations force ${version}

redoc-gen:
	npx @redocly/cli build-docs docs/swagger.yaml -o docs/redoc-static.html

api-docs:
	docker compose exec -it app swag fmt
	docker compose exec -it app swag init
	npx @redocly/cli build-docs docs/swagger.yaml -o docs/redoc-static.html

redoc-open:
	open docs/redoc-static.html
