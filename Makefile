db-up:
	sudo docker-compose up --remove-orphans --build

db-down:
	sudo docker-compose down

app-run:
	go run ./cmd/main.go 10,11,14,15

.PHONY: service-up service-down app-run