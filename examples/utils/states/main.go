// Command utils/states demonstrates listing and searching for US states
// using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make utils-states
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

func main() {
	ctx := context.Background()

	// Instantiate the Wodify client.
	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("creating wodify client: %v\n", err)
	}

	// Fetch a list of states.
	states, err := wc.Utils.ListStates(ctx, utils.NewStateListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(utils.StateFieldName),
	))
	if err != nil {
		log.Fatalf("listing states: %v\n", err)
	}

	// Print states.
	fmt.Printf("top 10 US states:")
	for _, state := range states.States {
		fmt.Printf("%+v\n", state)
	}

	// Search for States starting with "O"
	oStates, err := wc.Utils.SearchStates(ctx, utils.NewStateSearchRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(utils.StateFieldName),
		utils.NewStateQuery().StartsWith(utils.StateFieldName, "O"),
	))
	if err != nil {
		log.Fatalf("searching states: %v\n", err)
	}

	// Print states that start with "O".
	fmt.Println("states starting with 'O':")
	for _, state := range oStates.States {
		fmt.Printf("%+v\n", state)
	}
}
