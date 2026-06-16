// Command clients/crud demonstrates creating, fetching, and updating
// a client using the Wodify Go SDK.
// For actions like deactivating, reactivating, suspending, and reinstating clients,
// see the client/actions example.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make client-crud
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/clients"
	"github.com/andrew-hayworth22/wodify-go/models"
)

func main() {
	ctx := context.Background()

	// Instantiate the Wodify client.
	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("creating wodify client: %v\n", err)
	}

	// Create a client.
	client, err := wc.Clients.Create(ctx, clients.ClientCreateRequest{
		FirstName:   "Go SDK",
		LastName:    "Client",
		Email:       "api@wodify.com",
		LocationID:  11337,
		GenderID:    1,
		DateOfBirth: models.Date{Time: time.Now().AddDate(-20, 0, 0)},
		Tags:        []string{"api", "test"},
		Notes:       "Created by the Wodify Go SDK example.",
	})
	if err != nil {
		log.Fatalf("creating client: %v\n", err)
	}
	fmt.Printf("created client %d: %s %s\n", client.ID, client.FirstName, client.LastName)

	// Fetch the client.
	client, err = wc.Clients.Get(ctx, client.ID)
	if err != nil {
		log.Fatalf("getting client: %v", err)
	}
	fmt.Printf("fetched client %d: %s %s\n", client.ID, client.FirstName, client.LastName)

	// Update the client.
	req := clients.ClientUpdateRequestFrom(client)
	req.FirstName = "Go SDK (Updated)"
	client, err = wc.Clients.Update(ctx, client.ID, req)
	if err != nil {
		log.Fatalf("updating client: %v\n", err)
	}
	fmt.Printf("updated client %d: %s %s\n", client.ID, client.FirstName, client.LastName)
}
