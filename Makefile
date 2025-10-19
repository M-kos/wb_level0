build:
	go build -o ./bin/orderRepository ./cmd/main.go

lint:
	golangci-lint run ./... -v

clean-bin:
	rm -rf ./bin

tidy:
	go mod tidy -v

# запускает локально api
local-api-up:
	CRM_ENV_FILE=".env" go run ./cmd/main.go

# https://github.com/githubnemo/CompileDaemon
local-api-up-hot:
	CompileDaemon -exclude-dir=.git -exclude-dir=bin -exclude-dir=openapi -build='make build' -command='./bin/user' -color=true

docker-build:
	docker build -t orderRepository-service:latest -f ./Dockerfile .

docker-run:
	docker run --name orderRepository-service -d -p 8080:8080 -e CRM_PORT=8080 orderRepository-service:latest

compose-up:
	docker compose -f docker-compose.yaml --project-directory=./ up -d --build

compose-down:
	docker compose -f docker-compose.yaml --project-directory=./ down

migrate-down:
	docker compose --profile migration-down -f docker-compose.yaml up

.PHONY: generate-openapi build build-with-generate lint clean-bin tidy dev-up dev-down dev-api-test dev-db local-api-up local-api-up-hot docker-build docker-run compose-up compose-down migrate-down
