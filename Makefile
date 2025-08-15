vet:
	@go vet --composites=false ./...

tidy:
	@go mod tidy

test:
	@go test -race -coverprofile=coverage.out ./...

lint:
	@golangci-lint run --fix --sort-results --timeout=5m

check:
	make vet
	make lint
	make tidy
	make test

test-brew:
	@./scripts/test-brew.sh
