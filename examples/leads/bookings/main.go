// Command leads/bookings demonstrates listing and searching appointment bookings for
// a lead using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-bookings
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/leads"
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

	// Fetch a lead's bookings.
	const leadID = 2191681
	bookings, err := wc.Leads.ListBookings(ctx, leadID, leads.NewBookingListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(leads.BookingFieldStatusID),
	))
	if err != nil {
		log.Fatalf("listing lead bookings: %v\n", err)
	}

	// Print the bookings.
	fmt.Printf("lead %d has %d booking(s)\n", leadID, len(bookings.Bookings))
	for _, booking := range bookings.Bookings {
		fmt.Printf("booking: id=%d, service=%s, start-time=%s\n", booking.BookingID, booking.ServiceName, booking.LocalAppointmentStartDateTime)
	}

	// Fetch a lead's free trial bookings
	freeTrials, err := wc.Leads.SearchBookings(ctx, leadID, leads.NewBookingSearchRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(leads.BookingFieldStatusID),
		leads.NewBookingQuery().Eq(leads.BookingFieldIsFreeTrial, true),
	))
	if err != nil {
		log.Fatalf("searching lead bookings: %v\n", err)
	}

	// Print the lead's free trial bookings.
	fmt.Printf("lead %d has %d free trial booking(s)\n", leadID, len(freeTrials.Bookings))
	for _, booking := range freeTrials.Bookings {
		fmt.Printf("booking: id=%d, service=%s, start-time=%s\n", booking.BookingID, booking.ServiceName, booking.LocalAppointmentStartDateTime)
	}
}
