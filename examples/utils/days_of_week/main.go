// Command utils/days_of_week demonstrates listing days of the week
// using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make utils-days-of-week
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

	// Fetch a list of days of the week.
	days, err := wc.Utils.ListDaysOfWeek(ctx, utils.NewDayOfWeekListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(utils.DayOfWeekFieldID),
	))
	if err != nil {
		log.Fatalf("listing days of the week: %v\n", err)
	}

	// Print days of the week.
	fmt.Println("days of the week:")
	for _, day := range days.DaysOfWeek {
		fmt.Printf("%+v\n", day)
	}
}
