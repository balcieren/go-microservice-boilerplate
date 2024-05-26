dev:
	docker compose -f "docker-compose.dev.yaml" -p go-microservice-boilerplate-dev up 
prod:
	docker compose -f "docker-compose.prod.yaml" -p go-microservice-boilerplate-prod up 
swagger:
	swag init -g main.go --output ./docs --quiet --parseDependency --parseInternal
gorm:
	cd pkg/generate && go run gorm_gen.go
proto:
	protoc --go_out=. --go-grpc_out=.  ./pkg/proto/*.proto
