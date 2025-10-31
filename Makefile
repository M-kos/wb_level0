build:
	go build -o ./bin/order ./cmd/order-app/main.go

build-producer:
	go build -o ./bin/order-producer ./cmd/order-producer/main.go

lint:
	golangci-lint run ./... -v

clean-bin:
	rm -rf ./bin

tidy:
	go mod tidy -v

local-api-up:
	go run ./cmd/order-app/main.go

# https://github.com/githubnemo/CompileDaemon
local-api-up-hot:
	CompileDaemon -exclude-dir=.git -exclude-dir=bin -exclude-dir=openapi -build='make build' -command='./bin/user' -color=true

docker-build:
	docker build -t order-service:latest -f ./Dockerfile .

docker-run:
	docker run --name order-service -d -p 8080:8080 -e CRM_PORT=8080 order-service:latest

compose-up:
	docker compose -f docker-compose.yaml --project-directory=./ up -d --build

compose-down:
	docker compose -v -f docker-compose.yaml --project-directory=./ down

migrate-down:
	docker compose --profile migration-down -f docker-compose.yaml up

.PHONY: generate-openapi build build-with-generate lint clean-bin tidy dev-up dev-down dev-api-test dev-db local-api-up local-api-up-hot docker-build docker-run compose-up compose-down migrate-down
