// Command leads/convert demonstrates creating and converting
// a lead to a client using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-convert
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
	l, err := wc.Leads.Create(ctx, leads.LeadCreateRequest{
		FirstName:   "Go SDK",
		LastName:    "Converted Lead",
		Email:       "converted-lead@wodify.com",
		LocationID:  11337,
		GenderID:    2,
		DateOfBirth: models.Date{Time: time.Now().AddDate(-20, 0, 0)},
		Tags:        []string{"api", "test"},
		Notes:       "Created and converted by the Wodify Go SDK example.",
	})
	if err != nil {
		log.Fatalf("creating lead: %v\n", err)
	}
	fmt.Printf("created lead %d: %s %s\n", l.ID, l.FirstName, l.LastName)

	// Convert the lead to a client
	req := leads.LeadConvertRequestFrom(l)
	converted, err := wc.Leads.Convert(ctx, l.ID, req)
	if err != nil {
		log.Fatalf("converting lead: %v\n", err)
	}
	fmt.Printf("converted lead ID=%d to client ID=%d", l.ID, converted.ClientData.ID)
}
