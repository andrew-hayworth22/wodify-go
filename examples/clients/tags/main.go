// Command clients/tags demonstrates adding and deleting tags for
// a client using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make clients-tags
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

	// Create a lead.
	initialTags := []string{"api"}
	client, err := wc.Clients.Create(ctx, clients.ClientCreateRequest{
		FirstName:   "Go SDK",
		LastName:    "Client Tag Example",
		Email:       "api+tags@wodify.com",
		LocationID:  11337,
		GenderID:    3,
		DateOfBirth: models.Date{Time: time.Now().AddDate(-30, 0, 0)},
		Tags:        initialTags,
		Notes:       "Created by a Wodify Go SDK example for lead tags.",
	})
	if err != nil {
		log.Fatalf("creating client: %v\n", err)
	}
	fmt.Printf("created client %d: %s %s with tags %v\n", client.ID, client.FirstName, client.LastName, initialTags)

	// Add tags
	addedTags := []string{"added", "deleted"}
	fmt.Printf("adding tags: %v\n", addedTags)
	tags, err := wc.Clients.AddTags(ctx, client.ID, clients.TagsUpdateRequest{Tags: addedTags})
	if err != nil {
		log.Fatalf("adding tags: %v\n", err)
	}
	fmt.Printf("tags after add: %v\n", tags.Tags)

	// Delete tags
	deletedTags := []string{"deleted"}
	fmt.Printf("deleting tags: %v\n", deletedTags)
	tags, err = wc.Clients.DeleteTags(ctx, client.ID, clients.TagsUpdateRequest{Tags: deletedTags})
	if err != nil {
		log.Fatalf("deleting tags: %v\n", err)
	}
	fmt.Printf("tags after delete: %v\n", tags.Tags)
}
