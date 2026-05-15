.PHONY: build

build:
	@echo "Building..."
	@go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
	@echo "Done"

lint:
	@golangci-lint run

lint-fix:
	@golangci-lint run --fix
