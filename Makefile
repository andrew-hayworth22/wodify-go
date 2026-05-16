# Tests
test:
	go test ./... -cover

test-search:
	go test ./internal/search -cover

test-leads:
	go test ./leads -cover

# Examples
leads-crud:
	go run ./examples/leads/crud

leads-search:
	go run ./examples/leads/search

leads-statuses:
	go run ./examples/leads/statuses

leads-sources:
	go run ./examples/leads/sources

# Profiling
profile-test:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

# Utilities
lint:
	golangci-lint run