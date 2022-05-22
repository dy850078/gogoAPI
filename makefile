.PHONY: up down

up: docker-compose-up
down: docker-compose-down

docker-compose-up:
	docker-compose up

docker-compose-down:
	docker-compose down