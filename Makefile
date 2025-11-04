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

local-run:
	go run ./cmd/order-app/main.go

# https://github.com/githubnemo/CompileDaemon
local-run-hot:
	CompileDaemon -exclude-dir=.git -exclude-dir=bin -exclude-dir=openapi -build='make build' -command='./bin/user' -color=true

docker-build:
	docker build -t order-service:latest -f ./Dockerfile .

docker-run:
	docker run --name order-service -d -p 8080:8080 -e CRM_PORT=8080 order-service:latest

compose-up:
	docker compose -f docker-compose.yaml --project-directory=./ up -d --build

compose-up-local:
	docker compose -f docker-compose.yaml up -d --build \
		postgres kafka kafka-ui migrations-up

local-dev: compose-up-local local-run

compose-down:
	docker compose -v -f docker-compose.yaml --project-directory=./ down

test:
	go test ./... -v -race -count 100

.PHONY: build build-producer lint clean-bin tidy local-run local-run-hot docker-build docker-run compose-up compose-up-local compose-down local-dev test
