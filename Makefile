current_dir = $(shell pwd)

run-dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

run-prod:
	docker-compose up -d

stop:
	docker-compose stop