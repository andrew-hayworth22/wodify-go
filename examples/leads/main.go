// Command leads demonstrates basic usage of the Wodify Go SDK leads service.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	go run ./example/leads
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/leads"
	"github.com/andrew-hayworth22/wodify-go/types"
)

func main() {
	ctx := context.Background()

	// Create a new Wodify client.
	wc := wodify.New(
		wodify.WithAPIKey("[enter your API key here]"),
	)

	// Create a new lead.
	req := leads.CreateLeadRequest{
		FirstName:             "API",
		LastName:              "Lead 2",
		Email:                 "api2@woddifried.com",
		LocationID:            11337,
		Gender:                types.Genders.Female,
		PhoneNumber:           "123-123-1234",
		DateOfBirth:           types.Date{Time: time.Now().AddDate(-20, 0, 0)},
		StreetAddress1:        "123 Example St.",
		StreetAddress2:        "Apt. A",
		City:                  "Testing",
		StateID:               5,
		Province:              "",
		ZipCode:               "12345",
		CountryID:             57,
		Tags:                  []string{"api", "test"},
		Notes:                 "This is a test lead created by the Wodify Go SDK.",
		EmergencyContactName:  "Jane Doe",
		EmergencyContactPhone: "123-123-1234",
	}
	createdLead, err := wc.Leads.Create(ctx, req)
	if err != nil {
		log.Fatalf("failed to create lead: %v", err)
	}

	fmt.Printf("created lead %d: %s %s %s\n", createdLead.ID, createdLead.FirstName, createdLead.LastName, createdLead.Gender.Name)

	// Fetch the newly created lead.
	lead, err := wc.Leads.Get(ctx, createdLead.ID)
	if err != nil {
		log.Fatalf("failed to get lead: %v", err)
	}

	fmt.Printf("fetched lead %d: %s %s %s\n", lead.ID, lead.FirstName, lead.LastName, lead.Gender.Name)
}
