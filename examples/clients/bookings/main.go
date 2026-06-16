// Command clients/bookings demonstrates listing and searching appointment bookings for
// a client using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make clients-bookings
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/clients"
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

	// Fetch a client's bookings.
	const clientID = 3297777
	bookings, err := wc.Clients.ListBookings(ctx, clientID, clients.NewBookingListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.BookingFieldStatusID),
	))
	if err != nil {
		log.Fatalf("listing client bookings: %v\n", err)
	}

	// Print the bookings.
	fmt.Printf("client %d has %d booking(s)\n", clientID, len(bookings.Bookings))
	for _, booking := range bookings.Bookings {
		fmt.Printf("booking: id=%d, service=%s, start-time=%s\n", booking.BookingID, booking.ServiceName, booking.LocalAppointmentStartDateTime)
	}

	// Fetch a client's free trial bookings
	freeTrials, err := wc.Clients.SearchBookings(ctx, clientID, clients.NewBookingSearchRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.BookingFieldStatusID),
		clients.NewBookingQuery().Eq(clients.BookingFieldIsFreeTrial, true),
	))
	if err != nil {
		log.Fatalf("searching client bookings: %v\n", err)
	}

	// Print the client's free trial bookings.
	fmt.Printf("client %d has %d free trial booking(s)\n", clientID, len(freeTrials.Bookings))
	for _, booking := range freeTrials.Bookings {
		fmt.Printf("booking: id=%d, service=%s, start-time=%s\n", booking.BookingID, booking.ServiceName, booking.LocalAppointmentStartDateTime)
	}
}
