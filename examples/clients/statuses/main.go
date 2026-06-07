// Command clients/statuses demonstrates listing
// client statuses using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make clients-statuses
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/clients"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	ctx := context.Background()

	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("failed to create wodify client: %v", err)
	}

	// List statuses
	statuses, err := wc.Clients.ListStatuses(ctx, clients.NewStatusListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.StatusFieldName),
	))
	if err != nil {
		log.Fatalf("failed to fetch statuses: %v", err)
	}

	fmt.Println("client status list:")
	for _, status := range statuses.Statuses {
		fmt.Printf("status: %+v\n", status)
	}

	// Search statuses that end with an 'e'
	statuses, err = wc.Clients.SearchStatuses(ctx, clients.NewStatusSearchRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.StatusFieldName),
		clients.NewStatusQuery().EndsWith(clients.StatusFieldName, "e"),
	))
	if err != nil {
		log.Fatalf("failed to search statuses: %v", err)
	}

	fmt.Println("client statuses that end with 'e':")
	for _, status := range statuses.Statuses {
		fmt.Printf("status: %+v\n", status)
	}
}
