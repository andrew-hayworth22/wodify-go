// Command utils/object_types demonstrates listing object types and object action types
// using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make utils-object-types
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

	// Fetch a list of object types.
	objectTypes, err := wc.Utils.ListObjectTypes(ctx, utils.ListObjectTypesRequest{
		Page: models.PaginationRequest{
			PageSize: 10,
			Page:     1,
		},
		Sort: utils.NewObjectTypeSort(utils.ObjectTypeFieldName, false),
	})
	if err != nil {
		log.Fatalf("listing object types: %v\n", err)
	}

	// Print object types.
	fmt.Println("object types:")
	for _, objectType := range objectTypes.ObjectTypes {
		fmt.Printf("%+v\n", objectType)
	}

	// Search for the Appointment object type.
	appointment, err := wc.Utils.SearchObjectTypes(ctx, utils.SearchObjectTypesRequest{
		Page: models.PaginationRequest{
			PageSize: 1,
			Page:     1,
		},
		Query: utils.NewObjectTypeQuery().Eq(utils.ObjectTypeFieldName, "Appointment"),
	})
	if err != nil {
		log.Fatalf("searching object types: %v\n", err)
	}
	if len(appointment.ObjectTypes) == 0 {
		log.Fatalf("no object types found with name 'Appointment'")
	}

	// Print the Appointment object type.
	fmt.Printf("Appointment object type: %+v\n", appointment.ObjectTypes[0])

	// Fetch a list of object action types.
	objectActionTypes, err := wc.Utils.ListObjectActionTypes(ctx, utils.ListObjectActionTypesRequest{
		Page: models.PaginationRequest{
			PageSize: 10,
			Page:     1,
		},
		Sort: utils.NewObjectActionTypeSort(utils.ObjectActionTypeFieldName, false),
	})
	if err != nil {
		log.Fatalf("listing object action types: %v\n", err)
	}

	// Print object action types.
	fmt.Println("object action types:")
	for _, objectActionType := range objectActionTypes.ObjectActionTypes {
		fmt.Printf("%+v\n", objectActionType)
	}

	// Search for Client object action types.
	clientActionTypes, err := wc.Utils.SearchObjectActionTypes(ctx, utils.SearchObjectActionTypesRequest{
		Sort:  utils.NewObjectActionTypeSort(utils.ObjectActionTypeFieldName, false),
		Query: utils.NewObjectActionTypeQuery().Eq(utils.ObjectActionTypeFieldObjectTypeName, "Client"),
	})
	if err != nil {
		log.Fatalf("searching client types: %v\n", err)
	}

	// Print Client object action types.
	fmt.Println("client action types:")
	for _, objectActionType := range clientActionTypes.ObjectActionTypes {
		fmt.Printf("%+v\n", objectActionType)
	}
}
