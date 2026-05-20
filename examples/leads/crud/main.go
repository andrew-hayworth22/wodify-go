// Command leads/crud demonstrates creating, fetching, updating, and deleting
// a lead using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-crud
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
	lead, err := wc.Leads.Create(ctx, leads.CreateLeadRequest{
		FirstName:   "Go SDK",
		LastName:    "Lead",
		Email:       "api@wodify.com",
		LocationID:  11337,
		GenderID:    1,
		DateOfBirth: models.Date{Time: time.Now().AddDate(-20, 0, 0)},
		Tags:        []string{"api", "test"},
		Notes:       "Created by the Wodify Go SDK example.",
	})
	if err != nil {
		log.Fatalf("creating lead: %v\n", err)
	}
	fmt.Printf("created lead %d: %s %s\n", lead.ID, lead.FirstName, lead.LastName)

	// Fetch the lead.
	lead, err = wc.Leads.Get(ctx, lead.ID)
	if err != nil {
		log.Fatalf("getting lead: %v", err)
	}
	fmt.Printf("fetched lead %d: %s %s\n", lead.ID, lead.FirstName, lead.LastName)

	// Update the lead.
	req := leads.UpdateRequestFrom(lead)
	req.FirstName = "Go SDK (Updated)"
	lead, err = wc.Leads.Update(ctx, lead.ID, req)
	if err != nil {
		log.Fatalf("updating lead: %v\n", err)
	}
	fmt.Printf("updated lead %d: %s %s\n", lead.ID, lead.FirstName, lead.LastName)

	// Delete the lead.
	res, err := wc.Leads.Delete(ctx, lead.ID)
	if err != nil || !res.IsSuccess {
		log.Fatalf("deleting lead: %v\n", err)
	}
	log.Printf("deleted lead %d\n", res.LeadID)
}
