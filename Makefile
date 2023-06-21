dev:
	docker compose -f "docker-compose.dev.yaml" up
dev-down:
	docker compose -f "docker-compose.dev.yaml" down
build:
	docker compose -f "docker-compose.yaml" build
proto:
	protoc --go_out=. --go-grpc_out=.  ./pkg/proto/*.proto
swagger:
	swag init -g /app/proxy/main.go --parseDependency true --parseInternal true