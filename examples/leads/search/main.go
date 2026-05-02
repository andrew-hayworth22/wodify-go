// Command leads/search demonstrates listing and searching for
// leads using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-search
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
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

	listResults, err := wc.Leads.List(ctx, leads.ListRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: sort.NewSort(leads.SortByFirstName, false),
	})
	if err != nil {
		log.Fatalf("failed to fetch leads: %v\n", err)
	}

	fmt.Println("listing leads:")
	for _, l := range listResults.Leads {
		fmt.Printf("lead %d: %s %s (%s)\n", l.ID, l.FirstName, l.LastName, l.LeadStatus)
	}
	fmt.Printf("page %d, showing %d results, has more? %t\n", listResults.Pagination.Page, listResults.Pagination.PageSize, listResults.Pagination.HasMore)
	fmt.Println("--------------------------------------------------------")

	fmt.Println("searching for leads with first name 'Go SDK':")
	searchResults, err := wc.Leads.Search(ctx, leads.SearchRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort:  leads.NewSort(leads.SortByFirstName, false),
		Query: leads.NewQuery().Eq(leads.FilterByFirstName, "Go SDK"),
	})
	if err != nil {
		log.Fatalf("failed to search leads: %v\n", err)
	}

	for _, l := range searchResults.Leads {
		fmt.Printf("lead %d: %s %s (%s)\n", l.ID, l.FirstName, l.LastName, l.LeadStatus)
	}
	fmt.Printf("page %d, showing %d results, has more? %t\n", searchResults.Pagination.Page, searchResults.Pagination.PageSize, searchResults.Pagination.HasMore)
}
