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
	"github.com/andrew-hayworth22/wodify-go/leads"
)

func main() {
	ctx := context.Background()

	// Instantiate the Wodify client.
	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("creating wodify client: %v\n", err)
	}

	// Fetch a lead's class sign-ins.
	const leadID = 2191681
	signIns, err := wc.Leads.ListClassSignIns(ctx, leadID, leads.NewClassSignInListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(leads.ClassSignInFieldClassID),
	))
	if err != nil {
		log.Fatalf("listing lead sign-ins: %v\n", err)
	}

	// Print the sign-ins.
	fmt.Printf("lead %d has %d sign-in(s)\n", leadID, len(signIns.SignIns))
	for _, signIn := range signIns.SignIns {
		fmt.Printf("booking: id=%d, class=%s, start-time=%s\n", signIn.ID, signIn.ClassName, signIn.LocalClassStartDateTime)
	}

	// Fetch a lead's sign-ins where the attended email was sent.
	emailSent, err := wc.Leads.SearchClassSignIns(ctx, leadID, leads.NewClassSignInSearchRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(leads.ClassSignInFieldIsAttendedEmailSent),
		leads.NewClassSignInQuery().Eq(leads.ClassSignInFieldIsAttendedEmailSent, true),
	))
	if err != nil {
		log.Fatalf("searching lead sign-ins: %v\n", err)
	}

	// Print the lead's sign-ins with sent emails
	fmt.Printf("lead %d has %d sign-ins with a sent email\n", leadID, len(emailSent.SignIns))
	for _, signIn := range emailSent.SignIns {
		fmt.Printf("sign-in: id=%d, class=%s, start-time=%s\n", signIn.ID, signIn.ClassName, signIn.LocalClassStartDateTime)
	}
}
