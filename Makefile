dev:
	docker compose -f "docker-compose.dev.yaml" up
dev-down:
	docker compose -f "docker-compose.dev.yaml" down
build:
	docker compose -f "docker-compose.yaml" up
proto:
	protoc --go_out=. --go-grpc_out=.  ./pkg/proto/*.proto
ent:
	go generate ./pkg/ent
swagger:
	swag init -g /app/proxy/main.go --output ./app/proxy/docs --parseDependency true --parseInternal true 
go-global:
	go install github.com/swaggo/swag/cmd/swag@latest 
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 
	export PATH="$PATH:$(go env GOPATH)/bin"