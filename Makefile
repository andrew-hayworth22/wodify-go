# Tests
test:
	go test ./...

test-wodify:
	go test .

test-internal:
	go test $$(go list ./internal/... | grep -v testutil)

test-leads:
	go test ./leads

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

leads-reservations:
	go run ./examples/leads/reservations

# Profiling
profile-test:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
