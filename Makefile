.PHONY: run build clean test fmt check

BINARY = bin/mizni-godis.exe

run:
	go run ./cmd/server

build:
	if not exist bin mkdir bin
	go build -o $(BINARY) ./cmd/server

test:
	go test ./...

fmt:
	go fmt ./...

check: fmt
	go vet ./...

clean:
	if exist bin rmdir /s /q bin