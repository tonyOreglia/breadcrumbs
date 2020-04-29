current_dir = $(shell pwd)

run-dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

run-prod: build-prod
	docker-compose up -d

build-prod:
	docker-compose build --no-cache

stop:
	docker-compose stop