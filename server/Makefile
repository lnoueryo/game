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
	docker exec -it  kartenspielen-table go test ./interface/repositories/tests/... ./interface/controllers/tests/...

tests:
	docker exec -it  kartenspielen-table go test ./interface/controllers/tests/... -count 50
	docker exec -it  kartenspielen-table go test ./interface/repositories/tests/... -count 50

delete:
	docker-compose down --rmi all --volumes --remove-orphans
