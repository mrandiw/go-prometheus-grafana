run:
	go run main.go

install-prometheus:
	docker-compose -f docker-compose.yml up -d

run-k6:
	cd k6 && docker compose up -d

down-prometheus:
	docker-compose -f docker-compose.yml down

down-k6:
	cd k6 && docker-compose -f docker-compose.yml down