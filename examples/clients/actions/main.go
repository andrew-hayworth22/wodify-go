// Commanclientss/crud demonstrates deactivating, reactivating, suspending, and reinstating
// clients using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make client-actions
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
		LastName:    "Actions Client",
		Email:       "api-actions@wodify.com",
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

	// Deactivate client
	resp, err := wc.Clients.Deactivate(ctx, client.ID)
	if err != nil {
		log.Fatalf("deactivating client: %v\n", err)
	}
	fmt.Printf("successfully deactivated %s %s\n", resp.Client.FirstName, resp.Client.LastName)

	// Reactivate client
	resp, err = wc.Clients.Reactivate(ctx, client.ID)
	if err != nil {
		log.Fatalf("deactivating client: %v\n", err)
	}
	fmt.Printf("successfully reactivated %s %s\n", resp.Client.FirstName, resp.Client.LastName)

	// Suspend client
	resp, err = wc.Clients.Suspend(ctx, client.ID)
	if err != nil {
		log.Fatalf("suspending client: %v\n", err)
	}
	fmt.Printf("successfully suspended %s %s\n", resp.Client.FirstName, resp.Client.LastName)

	// Reinstate client
	resp, err = wc.Clients.Reinstate(ctx, client.ID)
	if err != nil {
		log.Fatalf("reinstating client: %v\n", err)
	}
	fmt.Printf("successfully reinstated %s %s\n", resp.Client.FirstName, resp.Client.LastName)
}
