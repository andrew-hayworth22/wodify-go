// Command leads/tags demonstrates adding and deleting tags for
// a lead using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-tags
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/leads"
	"github.com/andrew-hayworth22/wodify-go/models"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	ctx := context.Background()

	// Instantiate the Wodify client.
	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("creating wodify client: %v\n", err)
	}

	// Create a lead.
	initialTags := []string{"api"}
	lead, err := wc.Leads.Create(ctx, leads.CreateLeadRequest{
		FirstName:   "Go SDK",
		LastName:    "Lead Tag Example",
		Email:       "api@wodify.com",
		LocationID:  11337,
		Gender:      3,
		DateOfBirth: models.Date{Time: time.Now().AddDate(-30, 0, 0)},
		Tags:        initialTags,
		Notes:       "Created by a Wodify Go SDK example for lead tags.",
	})
	if err != nil {
		log.Fatalf("creating lead: %v\n", err)
	}
	fmt.Printf("created lead %d: %s %s with tags %v\n", lead.ID, lead.FirstName, lead.LastName, initialTags)

	// Add tags
	addedTags := []string{"added", "deleted"}
	fmt.Printf("adding tags: %v\n", addedTags)
	tags, err := wc.Leads.AddTags(ctx, lead.ID, leads.UpdateTagsRequest{Tags: addedTags})
	if err != nil {
		log.Fatalf("adding tags: %v\n", err)
	}
	fmt.Printf("tags after add: %v\n", tags.Tags)

	// Delete tags
	deletedTags := []string{"deleted"}
	fmt.Printf("deleting tags: %v\n", deletedTags)
	tags, err = wc.Leads.DeleteTags(ctx, lead.ID, leads.UpdateTagsRequest{Tags: deletedTags})
	if err != nil {
		log.Fatalf("deleting tags: %v\n", err)
	}
	fmt.Printf("tags after delete: %v\n", tags.Tags)
}
