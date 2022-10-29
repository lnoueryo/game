start:
	docker-compose down
	docker-compose up -d
restart:
	docker-compose down
	docker-compose build --no-cache
	docker-compose up -d
delete:
	docker-compose down --rmi all --volumes --remove-orphans