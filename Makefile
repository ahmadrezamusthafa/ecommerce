start-infra:
	@echo "Running infrastructure" && \
	docker compose -f docker-compose.yaml up -d

stop-infra:
	@echo "Stopping infrastructure" && \
	docker compose -f docker-compose.yaml down

run-migration:
	@echo "Running all database migration" && \
    go build -o ecommerce && \
    ./ecommerce migrate/up

run-service:
	@echo "Running service" && \
	go run ./main.go
