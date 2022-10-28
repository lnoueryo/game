init: ## setup docker build, network, and databases
	docker network create game || true
	docker compose up -d

server:
	docker-compose down server_api
	docker-compose up -d api

start:
	docker-compose down
	docker-compose up -d

stop:
	docker-compose down

seed:
	docker-compose up -d migrate

test:
	go test ./interface/repositories/... -count 30

settest:
	docker-compose down
	docker-compose up -d
	go test ./interface/repositories/... -count 30
	docker-compose down