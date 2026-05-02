test:
	go test ./... -cover

test-search:
	go test ./search -cover

example-lead:
	go run ./examples/leads
