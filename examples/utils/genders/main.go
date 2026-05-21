// Command utils/genders demonstrates listing genders
// using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make utils-genders
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/models"
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

	// Fetch a list of genders.
	genders, err := wc.Utils.ListGenders(ctx, utils.ListGendersRequest{
		Page: models.PaginationRequest{
			PageSize: 7,
			Page:     1,
		},
		Sort: utils.NewGenderSort(utils.GenderFieldID, false),
	})
	if err != nil {
		log.Fatalf("listing genders: %v\n", err)
	}

	// Print genders.
	fmt.Println("genders:")
	for _, gender := range genders.Genders {
		fmt.Printf("%+v\n", gender)
	}
}
