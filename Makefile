current_dir = $(shell pwd)

build-dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml build

run-dev: build-dev
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

build-prod:
	docker-compose build

run-prod: build-prod:
	docker-compose up -d

stop:
	docker-compose stop