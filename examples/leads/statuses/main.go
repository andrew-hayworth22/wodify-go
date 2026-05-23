// Command leads/statuses demonstrates listing
// lead statuses using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-statuses
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/leads"
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

	statuses, err := wc.Leads.ListStatuses(ctx, leads.NewStatusListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(leads.StatusFieldName),
	))
	if err != nil {
		log.Fatalf("failed to fetch statuses: %v", err)
	}

	for _, status := range statuses.Statuses {
		fmt.Printf("status: %+v\n", status)
	}
}
