dev:
	.\scripts\dev.bat

prod:
	.\scripts\prod.bat

build:
	go build -o bin/main.exe main.go

clean:
	if exist bin rmdir /s /q bin

test:
	go test ./...

deps:
	go mod download
	go mod tidy

.PHONY: dev prod build clean test deps
