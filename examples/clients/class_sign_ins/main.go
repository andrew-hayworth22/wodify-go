// Command leads/class_sign_ins demonstrates listing and searching class sign-ins for
// a lead using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-class-sign-ins
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

	// Fetch a client's class sign-ins.
	const clientID = 3297777
	signIns, err := wc.Clients.ListClassSignIns(ctx, clientID, clients.NewClassSignInListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.ClassSignInFieldClassID),
	))
	if err != nil {
		log.Fatalf("listing client sign-ins: %v\n", err)
	}

	// Print the sign-ins.
	fmt.Printf("lead %d has %d sign-in(s)\n", clientID, len(signIns.SignIns))
	for _, signIn := range signIns.SignIns {
		fmt.Printf("booking: id=%d, class=%s, start-time=%s\n", signIn.ID, signIn.ClassName, signIn.LocalClassStartDateTime)
	}

	// Fetch a client's sign-ins where the attended email was sent.
	emailSent, err := wc.Clients.SearchClassSignIns(ctx, clientID, clients.NewClassSignInSearchRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.ClassSignInFieldIsAttendedEmailSent),
		clients.NewClassSignInQuery().Eq(clients.ClassSignInFieldIsAttendedEmailSent, true),
	))
	if err != nil {
		log.Fatalf("searching client sign-ins: %v\n", err)
	}

	// Print the client's sign-ins with sent emails
	fmt.Printf("client %d has %d sign-ins with a sent email\n", clientID, len(emailSent.SignIns))
	for _, signIn := range emailSent.SignIns {
		fmt.Printf("sign-in: id=%d, class=%s, start-time=%s\n", signIn.ID, signIn.ClassName, signIn.LocalClassStartDateTime)
	}
}
