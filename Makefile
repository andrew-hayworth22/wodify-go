# Tests
test:
	go test ./... -cover

test-wodify:
	go test . -cover

test-internal:
	go test $$(go list ./internal/... | grep -v testutil) -cover

test-utils:
	go test ./utils -cover

test-leads:
	go test ./leads -cover

# Examples
utils-countries:
	go run ./examples/utils/countries

utils-days-of-week:
	go run ./examples/utils/days_of_week

utils-genders:
	go run ./examples/utils/genders

utils-object-types:
	go run ./examples/utils/object_types

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
	go run ./examples/leads/class_sign_ins

leads-reservations:
	go run ./examples/leads/reservations

leads-performance-results:
	go run ./examples/leads/performance_results

# Profiling
profile-test:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
