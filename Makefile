test:
	go test ./... -cover

test-search:
	go test ./internal/search -cover

test-leads:
	go test ./leads -cover

leads-crud:
	go run ./examples/leads/crud

leads-search:
	go run ./examples/leads/search