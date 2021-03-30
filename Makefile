.SILENT:

.PHONY: run
run:
	go build -o ./bin/polliter.exe ./cmd/server/main.go
	./bin/polliter.exe

.PHONY: build
build:
	go build -o ./bin/polliter.exe ./cmd/server/main.go

.DEFAULT_GOAL:=run