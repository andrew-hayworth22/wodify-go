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

	// --- Create Lead ---
	createdLead, err := createLead(ctx, wc)
	if err != nil {
		log.Fatalf("failed to create lead: %v", err)
	}
	log.Printf("created lead %d: %s %s (%s)\n", createdLead.ID, createdLead.FirstName, createdLead.LastName, createdLead.Gender.Name)

	// --- Get Lead ---
	fetchedLead, err := getLead(ctx, wc, createdLead.ID)
	if err != nil {
		log.Fatalf("failed to get lead: %v", err)
	}
	log.Printf("fetched lead %d: %s %s (%s)\n", fetchedLead.ID, fetchedLead.FirstName, fetchedLead.LastName, fetchedLead.Gender.Name)

	// --- List Leads ---
	err = listLeads(ctx, wc)
	if err != nil {
		log.Fatalf("failed to list leads: %v", err)
	}

	// --- Search Leads ---
	err = searchLeads(ctx, wc)
	if err != nil {
		log.Fatalf("failed to search leads: %v", err)
	}

	// --- Delete Lead ---
	err = deleteLead(ctx, wc, createdLead.ID)
	if err != nil {
		log.Fatalf("failed to delete lead: %v", err)
	}
}

func createLead(ctx context.Context, wc *wodify.Client) (*leads.Lead, error) {
	return wc.Leads.Create(ctx, leads.CreateLeadRequest{
		FirstName:             "Go SDK",
		LastName:              "Lead",
		Email:                 "api2@woddifried.com",
		LocationID:            11337,
		Gender:                models.Genders.Female,
		PhoneNumber:           "123-123-1234",
		DateOfBirth:           models.Date{Time: time.Now().AddDate(-20, 0, 0)},
		StreetAddress1:        "123 Example St.",
		StreetAddress2:        "Apt. A",
		City:                  "Testing",
		StateID:               5,
		ZipCode:               "12345",
		CountryID:             57,
		Tags:                  []string{"api", "test"},
		Notes:                 "This is a test lead created by the Wodify Go SDK.",
		EmergencyContactName:  "Jane Doe",
		EmergencyContactPhone: "123-123-1234",
	})
}

func getLead(ctx context.Context, wc *wodify.Client, id int64) (*leads.Lead, error) {
	return wc.Leads.Get(ctx, id)
}

func listLeads(ctx context.Context, wc *wodify.Client) error {
	fmt.Println("listing leads")
	results, err := wc.Leads.List(ctx, leads.ListRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort: sort.NewSort(leads.SortByFirstName, false),
	})
	if err != nil {
		return err
	}

	for _, l := range results.Leads {
		log.Printf("lead %d: %s %s (%s)\n", l.ID, l.FirstName, l.LastName, l.LeadStatus)
	}
	log.Printf("page %d, showing %d results, has more? %t", results.Pagination.Page, results.Pagination.PageSize, results.Pagination.HasMore)
	return nil
}

func searchLeads(ctx context.Context, wc *wodify.Client) error {
	fmt.Println("searching for leads with first name 'Go SDK'")
	results, err := wc.Leads.Search(ctx, leads.SearchRequest{
		Page: models.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
		Sort:  sort.NewSort(leads.SortByFirstName, false),
		Query: leads.NewQuery().Eq(leads.FilterByFirstName, "Go SDK"),
	})
	if err != nil {
		return err
	}

	for _, l := range results.Leads {
		log.Printf("lead %d: %s %s (%s)\n", l.ID, l.FirstName, l.LastName, l.LeadStatus)
	}
	log.Printf("page %d, showing %d results, has more? %t", results.Pagination.Page, results.Pagination.PageSize, results.Pagination.HasMore)
	return nil
}

func deleteLead(ctx context.Context, wc *wodify.Client, id int64) error {
	res, err := wc.Leads.Delete(ctx, id)
	if err != nil {
		return err
	}
	result := "successfully deleted lead"
	if !res.IsSuccess {
		result = "failed to delete lead"
	}

	fmt.Printf("%s %d\n", result, id)
	return nil
}
