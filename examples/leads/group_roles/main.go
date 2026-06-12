// Command leads/group_roles demonstrates fetching and searching for
// lead group roles using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-group-roles
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/leads"
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

	// List lead group roles
	roles, err := wc.Leads.ListGroupRoles(ctx, leads.NewGroupRoleListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(leads.GroupRoleFieldName),
	))

	fmt.Println("list of lead group roles:")
	for _, role := range roles.Roles {
		fmt.Printf("- %s\n", role.Name)
	}

	// Search for 'Guardian' group role
	guardian, err := wc.Leads.SearchGroupRoles(ctx, leads.NewGroupRoleSearchRequest(
		wodify.NewPaginationRequest(1, 1),
		wodify.SortAscending(leads.GroupRoleFieldID),
		leads.NewGroupRoleQuery().Eq(leads.GroupRoleFieldName, "Guardian"),
	))
	if err != nil {
		log.Fatalf("searching group roles: %v\n", err)
	}
	if len(guardian.Roles) == 0 {
		log.Fatal("guardian role not found")
	}

	fmt.Printf("guardian role found: ID=%d Name=%s\n", guardian.Roles[0].ID, guardian.Roles[0].Name)
}
