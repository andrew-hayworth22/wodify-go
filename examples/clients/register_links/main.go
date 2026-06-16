// Command clients/register_links demonstrates generating a registration link
// for a client using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make clients-register-links
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
)

func main() {
	ctx := context.Background()

	// Instantiate the Wodify client.
	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("creating wodify client: %v\n", err)
	}

	// Generate a registration link for a client.
	const clientID = 3297777
	resp, err := wc.Clients.GenerateRegisterLink(ctx, clientID)
	if err != nil {
		log.Fatalf("generating registration link: %v\n", err)
	}

	// Print the registration link.
	fmt.Printf("registration link: %s\n", resp.Link)
}
