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
)

func main() {
	ctx := context.Background()

	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("failed to create wodify client: %v", err)
	}

	sources, err := wc.Leads.ListSources(ctx, leads.NewSourceListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(leads.SourceFieldName),
	))
	if err != nil {
		log.Fatalf("failed to fetch statuses: %v", err)
	}

	for _, source := range sources.Sources {
		fmt.Printf("source: %+v\n", source)
	}
}
