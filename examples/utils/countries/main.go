// Command utils/countries demonstrates listing and searching for countries
// using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make utils-countries
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

	// Fetch a list of countries.
	countries, err := wc.Utils.ListCountries(ctx, utils.ListCountriesRequest{
		Page: models.PaginationRequest{
			PageSize: 10,
			Page:     1,
		},
		Sort: utils.NewCountrySort(utils.CountryFieldName, false),
	})
	if err != nil {
		log.Fatalf("listing countries: %v\n", err)
	}

	// Print countries.
	fmt.Printf("top 10 countries:")
	for _, country := range countries.Countries {
		fmt.Printf("%+v\n", country)
	}

	// Search for countries containing "America" in the name.
	countries, err = wc.Utils.SearchCountries(ctx, utils.SearchCountriesRequest{
		Page: models.PaginationRequest{
			PageSize: 15,
			Page:     1,
		},
		Sort:  utils.NewCountrySort(utils.CountryFieldName, false),
		Query: utils.NewCountryQuery().Contains(utils.CountryFieldName, "America"),
	})
	if err != nil {
		log.Fatalf("searching countries: %v\n", err)
	}

	// Print countries.
	fmt.Println("countries containing 'America':")
	for _, country := range countries.Countries {
		fmt.Printf("%+v\n", country)
	}
}
