dev:
	docker compose -f "docker-compose.dev.yaml" up
dev-down:
	docker compose -f "docker-compose.dev.yaml" down
build:
	docker compose -f "docker-compose.dev.yaml" build
proto:
	protoc --go_out=. --go-grpc_out=.  ./pkg/proto/*.proto