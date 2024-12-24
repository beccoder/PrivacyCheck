docker-up:
	docker compose --env-file .env -f docker-compose.yaml up --build -d

docker-down:
	docker compose --env-file .env -f docker-compose.yaml down

docker-restart: docker-down docker-up

docker-logs:
	docker compose --env-file .env -f docker-compose.yaml logs -f

# make add-migration filename=create-users
add-migration:
	migrate create -ext sql -dir database/migrations -seq ${filename}
