package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/andrew-hayworth22/wodify-go"
)

func main() {
	ctx := context.Background()

	wc := wodify.New(
		wodify.WithAPIKey("aB6cZQIa8p5qwbzCuhF1j4DI7kUO4BM549XoCwqE"),
		wodify.WithBaseURL("https://devapi.wodify.com/v1"),
		wodify.WithHTTPClient(&http.Client{}),
	)

	lead, err := wc.Leads.Get(ctx, 2191681)
	if err != nil {
		log.Fatalf("failed to get lead: %v", err)
	}
	fmt.Printf("%+v\n", lead)
}
