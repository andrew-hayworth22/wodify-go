# Tests
test:
	go test ./... -cover

test-wodify:
	go test . -cover

test-internal:
	go test ./internal/... -cover

test-leads:
	go test ./leads -cover

# Examples
leads-crud:
	go run ./examples/leads/crud

leads-search:
	go run ./examples/leads/search

leads-convert:
	go run ./examples/leads/convert

leads-statuses:
	go run ./examples/leads/statuses

leads-sources:
	go run ./examples/leads/sources

leads-tags:
	go run ./examples/leads/tags

leads-bookings:
	go run ./examples/leads/bookings

leads-class-sign-ins:
	go run ./examples/leads/classsignins

# Profiling
profile-test:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
