init: ## setup docker build, network, and databases
	docker network create --subnet=192.168.0.1/24 game || true
	docker compose up -d

server:
	docker stop kartenspielen-table
	docker-compose up -d kartenspielen-table

start:
	docker-compose down
	docker-compose up -d

restart:
	docker-compose down
	docker-compose build --no-cache
	docker-compose up -d

seed:
	docker-compose up -d migrate

test:
	go test -v ./interface/repositories/tests/... -count 30

delete:
	docker-compose down --rmi all --volumes --remove-orphans
