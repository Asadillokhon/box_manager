include .env
export

service-run:
	go run cmd/main.go

migrate-up:
	migrate -path cmd/migrations -database ${CONN_STRING} up

migrate-down:
	migrate -path cmd/migrations -database ${CONN_STRING} down
	
make-up:
	sudo docker compose up -d --build

make-down:
	sudo docker compose down