gen:
	mockgen -source=internal/repo/repo.go -destination=internal/repo/mocks/mock_repo.go

up-http-postgres:
	TRANSPORT_MODE=http STORAGE_MODE=postgres docker compose up --build

up-http-cache:
	TRANSPORT_MODE=http STORAGE_MODE=cache docker compose up --build

up-grpc-postgres:
	TRANSPORT_MODE=grpc STORAGE_MODE=postgres docker compose up --build

up-grpc-cache:
	TRANSPORT_MODE=grpc STORAGE_MODE=cache docker compose up --build

test:
	go test ./internal/service

cover:
	go tool cover -html=./internal/service/coverage.out