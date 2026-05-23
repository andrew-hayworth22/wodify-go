// Command utils/units_of_time demonstrates listing and searching for units of time
// using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make utils-units-of-time
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/utils"
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

	// Fetch a list of units of time.
	units, err := wc.Utils.ListUnitsOfTime(ctx, utils.NewUnitOfTimeListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(utils.UnitOfTimeFieldNameSingular),
	))
	if err != nil {
		log.Fatalf("listing units of time: %v\n", err)
	}

	// Print units of time.
	fmt.Println("units of time:")
	for _, unit := range units.UnitsOfTime {
		fmt.Printf("%+v\n", unit)
	}

	// Search for Minute unit of time.
	minute, err := wc.Utils.SearchUnitsOfTime(ctx, utils.NewUnitOfTimeSearchRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(utils.UnitOfTimeFieldNameSingular),
		utils.NewUnitOfTimeQuery().Eq(utils.UnitOfTimeFieldNameSingular, "Minute"),
	))
	if err != nil {
		log.Fatalf("searching units of time: %v\n", err)
	}

	if len(minute.UnitsOfTime) == 0 {
		log.Fatalf("no units of time found with singular 'Minute'")
	}

	fmt.Printf("minute unit of time: %+v\n", minute.UnitsOfTime[0])
}
