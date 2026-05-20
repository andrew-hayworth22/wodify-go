// Command leads/statuses demonstrates listing
// lead sources using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-sources
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/leads"
	"github.com/andrew-hayworth22/wodify-go/models"
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

	sources, err := wc.Leads.ListSources(ctx, leads.ListSourcesRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: leads.NewSourceSort(leads.SourceFieldName, false),
	})
	if err != nil {
		log.Fatalf("failed to fetch statuses: %v", err)
	}

	for _, source := range sources.Sources {
		fmt.Printf("source: %+v\n", source)
	}
}
