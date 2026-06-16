// Command clients/search demonstrates listing and searching for
// clients using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make clients-search
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/clients"
)

func main() {
	ctx := context.Background()

	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("failed to create wodify client: %v", err)
	}

	listResults, err := wc.Clients.List(ctx, clients.NewClientListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.ClientFieldFirstName),
	))
	if err != nil {
		log.Fatalf("failed to fetch clients: %v\n", err)
	}

	fmt.Println("listing clients:")
	for _, c := range listResults.Clients {
		fmt.Printf("client %d: %s %s (%s)\n", c.ID, c.FirstName, c.LastName, c.ClientStatus)
	}
	fmt.Printf("page %d, showing %d results, has more? %t\n", listResults.Pagination.Page, listResults.Pagination.PageSize, listResults.Pagination.HasMore)
	fmt.Println("--------------------------------------------------------")

	fmt.Println("searching for clients that were converted from leads:")
	searchResults, err := wc.Clients.Search(ctx, clients.NewClientSearchRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.ClientFieldLastName),
		clients.NewClientQuery().Eq(clients.ClientFieldIsConvertedFromLead, true),
	))
	if err != nil {
		log.Fatalf("failed to query clients: %v\n", err)
	}

	for _, c := range searchResults.Clients {
		fmt.Printf("client %d: %s %s (%s)\n", c.ID, c.FirstName, c.LastName, c.ClientStatus)
	}
	fmt.Printf("page %d, showing %d results, has more? %t\n", searchResults.Pagination.Page, searchResults.Pagination.PageSize, searchResults.Pagination.HasMore)
}
