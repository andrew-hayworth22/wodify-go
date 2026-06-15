# Tests
test:
	go test ./... -cover

test-wodify:
	go test . -cover

test-models:
	go test ./models -cover

test-internal:
	go test ./internal/... -cover

test-utils:
	go test ./utils -cover

test-leads:
	go test ./leads -cover

test-clients:
	go test ./clients -cover

# Examples
utils-countries:
	go run ./examples/utils/countries

utils-states:
	go run ./examples/utils/states

utils-days-of-week:
	go run ./examples/utils/days_of_week

utils-genders:
	go run ./examples/utils/genders

utils-object-types:
	go run ./examples/utils/object_types

utils-units-of-time:
	go run ./examples/utils/units_of_time

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

leads-groups:
	go run ./examples/leads/groups

leads-bookings:
	go run ./examples/leads/bookings

leads-class-sign-ins:
	go run ./examples/leads/class_sign_ins

leads-reservations:
	go run ./examples/leads/reservations

leads-performance-results:
	go run ./examples/leads/performance_results


clients-crud:
	go run ./examples/clients/crud

clients-search:
	go run ./examples/clients/search

clients-statuses:
	go run ./examples/clients/statuses

clients-actions:
	go run ./examples/clients/actions

clients-groups:
	go run ./examples/clients/groups

clients-register-links:
	go run ./examples/clients/register_links

clients-tags:
	go run ./examples/clients/tags

# Profiling
profile-test:
	go test $$(go list ./... | grep -v testutil | grep -v /examples) -coverprofile=coverage.out && go tool cover -html=coverage.out
