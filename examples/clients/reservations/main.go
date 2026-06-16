// Command clients/reservations demonstrates listing and searching class reservations for
// a client using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make clients-reservations
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/clients"
)

func main() {
	ctx := context.Background()

	// Instantiate the Wodify client.
	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("creating wodify client: %v\n", err)
	}

	// Fetch a client's reservations.
	const clientID = 3297777
	reservations, err := wc.Clients.ListReservations(ctx, clientID, clients.NewReservationListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.ReservationFieldStatusID),
	))
	if err != nil {
		log.Fatalf("listing client reservations: %v\n", err)
	}

	// Print the reservations.
	fmt.Printf("client %d has %d reservation(s)\n", clientID, len(reservations.Reservations))
	for _, reservation := range reservations.Reservations {
		fmt.Printf("reservation: id=%d, class=%s, start-time=%s\n", reservation.ID, reservation.ClassName, reservation.LocalClassStartDateTime)
	}

	// Fetch a client's late cancellations.
	lateCancellations, err := wc.Clients.SearchReservations(ctx, clientID, clients.NewReservationSearchRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.ReservationFieldStatusID),
		clients.NewReservationQuery().Eq(clients.ReservationFieldIsLateCancellation, true),
	))
	if err != nil {
		log.Fatalf("searching client reservations: %v\n", err)
	}

	// Print the client's late cancellations.
	fmt.Printf("client %d has %d late cancellations\n", clientID, len(lateCancellations.Reservations))
	for _, reservation := range lateCancellations.Reservations {
		fmt.Printf("reservation: id=%d, class=%s, start-time=%s\n", reservation.ID, reservation.ClassName, reservation.LocalClassStartDateTime)
	}
}
