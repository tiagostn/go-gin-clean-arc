# Main
run:
	go run ./cmd/server/main.go

build:
	go build ./cmd/server/main.go

dev:
	CompileDaemon --build="go build cmd/server/main.go" --command=./main

test:
	go test -cover ./...

# Docker
dockerdev:
	docker-compose -f docker-compose.dev.yml up --build

# Modules
tidy:
	go mod tidy
	go mod vendor

# Tools
run-linter:
	golangci-lint run ./...

swaggo:
	swag init -g cmd/server/main.go
