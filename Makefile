all: start_service

start_service:
	docker-compose up --build

test:
	go test ./...