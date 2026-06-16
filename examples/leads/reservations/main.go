// Command leads/reservations demonstrates listing and searching class reservations for
// a lead using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-reservations
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/leads"
)

func main() {
	ctx := context.Background()

	// Instantiate the Wodify client.
	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("creating wodify client: %v\n", err)
	}

	// Fetch a lead's reservations.
	const leadID = 2191681
	reservations, err := wc.Leads.ListReservations(ctx, leadID, leads.NewReservationListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(leads.ReservationFieldStatusID),
	))
	if err != nil {
		log.Fatalf("listing lead reservations: %v\n", err)
	}

	// Print the reservations.
	fmt.Printf("lead %d has %d reservation(s)\n", leadID, len(reservations.Reservations))
	for _, reservation := range reservations.Reservations {
		fmt.Printf("reservation: id=%d, class=%s, start-time=%s\n", reservation.ID, reservation.ClassName, reservation.LocalClassStartDateTime)
	}

	// Fetch a lead's late cancellations.
	lateCancellations, err := wc.Leads.SearchReservations(ctx, leadID, leads.NewReservationSearchRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(leads.ReservationFieldStatusID),
		leads.NewReservationQuery().Eq(leads.ReservationFieldIsLateCancellation, true),
	))
	if err != nil {
		log.Fatalf("searching lead reservations: %v\n", err)
	}

	// Print the lead's late cancellations.
	fmt.Printf("lead %d has %d late cancellations\n", leadID, len(lateCancellations.Reservations))
	for _, reservation := range lateCancellations.Reservations {
		fmt.Printf("reservation: id=%d, class=%s, start-time=%s\n", reservation.ID, reservation.ClassName, reservation.LocalClassStartDateTime)
	}
}
