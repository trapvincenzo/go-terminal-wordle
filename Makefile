.PHONY: run
run:
	@go run cmd/main.go

.PHONY: build
build:
	@go build -o build/wordle cmd/main.go