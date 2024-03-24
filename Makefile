run:
	@go run ./cmd/api

test:
	@GOFLAGS="-count=1" go test -v -cover -coverprofile=coverage.txt -race ./... 

coverage:
	@go tool cover -html=coverage.txt -o coverage.html

tidy:
	@go mod tidy

vendor:
	@go mod vendor

