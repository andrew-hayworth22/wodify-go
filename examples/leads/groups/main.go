// Command leads/groups demonstrates multiple operations on
// lead groups using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-groups
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
	if err != nil {
		log.Fatalf("listing group roles: %v\n", err)
	}

	fmt.Println("list of lead group roles:")
	for _, role := range roles.Roles {
		fmt.Printf("- %s\n", role.Name)
	}

	// Search for 'Dependent' group role
	dependent, err := wc.Leads.SearchGroupRoles(ctx, leads.NewGroupRoleSearchRequest(
		wodify.NewPaginationRequest(1, 1),
		wodify.SortAscending(leads.GroupRoleFieldID),
		leads.NewGroupRoleQuery().Eq(leads.GroupRoleFieldName, "Dependent"),
	))
	if err != nil {
		log.Fatalf("searching group roles: %v\n", err)
	}
	if len(dependent.Roles) == 0 {
		log.Fatal("dependent role not found")
	}

	fmt.Printf("dependent role found: ID=%d Name=%s\n", dependent.Roles[0].ID, dependent.Roles[0].Name)

	// Create a new lead group
	group, err := wc.Leads.CreateGroup(ctx, leads.GroupCreateRequest{
		GroupParticipants: []leads.GroupParticipantInput{
			{
				GroupParticipantLeadID: 1,
				GroupRoleID:            2,
			},
			{
				GroupParticipantLeadID: 2,
				GroupRoleID:            1,
			},
		},
	})
	if err != nil {
		log.Fatalf("creating lead group: %v\n", err)
	}

	fmt.Println("created group")

	// Add new leads to group
	group, err = wc.Leads.AddGroupParticipants(ctx, group.Group.ID, leads.GroupParticipantsRequest{
		LeadIDs: []int64{3, 4},
	})
	if err != nil {
		log.Fatalf("adding leads to group: %v\n", err)
	}

	fmt.Println("added leads to group")

	// Remove leads from group
	group, err = wc.Leads.RemoveGroupParticipants(ctx, group.Group.ID, leads.GroupParticipantsRequest{
		LeadIDs: []int64{3},
	})
	if err != nil {
		log.Fatalf("removing leads from group: %v\n", err)
	}

	fmt.Println("removed leads from group")

	// Convert lead from dependent
	_, err = wc.Leads.ConvertFromDependent(ctx, 123, leads.ConvertFromDependentRequest{Email: "john.doe@example.com"})
	if err != nil {
		log.Fatalf("converting lead from dependent: %v\n", err)
	}

	fmt.Println("converted lead from dependent")
}
