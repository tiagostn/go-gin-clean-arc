# Main
run:
	go run ./cmd/server/main.go

build:
	go build ./cmd/server/main.go

dev:
	CompileDaemon --build="go build cmd/server/main.go" --command=./main

test:
	go test -cover ./...


# Modules
tidy:
	go mod tidy
	go mod vendor
