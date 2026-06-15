// Command clients/groups demonstrates multiple operations on
// client groups using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make clients-groups
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/clients"
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

	// List client group roles
	roles, err := wc.Clients.ListGroupRoles(ctx, clients.NewGroupRoleListRequest(
		wodify.NewPaginationRequest(1, 10),
		wodify.SortAscending(clients.GroupRoleFieldName),
	))
	if err != nil {
		log.Fatalf("listing group roles: %v\n", err)
	}

	fmt.Println("list of lead group roles:")
	for _, role := range roles.Roles {
		fmt.Printf("- %s\n", role.Name)
	}

	// Search for 'Dependent' group role
	dependent, err := wc.Clients.SearchGroupRoles(ctx, clients.NewGroupRoleSearchRequest(
		wodify.NewPaginationRequest(1, 1),
		wodify.SortAscending(clients.GroupRoleFieldID),
		clients.NewGroupRoleQuery().Eq(clients.GroupRoleFieldName, "Dependent"),
	))
	if err != nil {
		log.Fatalf("searching group roles: %v\n", err)
	}
	if len(dependent.Roles) == 0 {
		log.Fatal("dependent role not found")
	}

	fmt.Printf("dependent role found: ID=%d Name=%s\n", dependent.Roles[0].ID, dependent.Roles[0].Name)

	// Create a new client group
	group, err := wc.Clients.CreateGroup(ctx, clients.GroupCreateRequest{
		GroupParticipants: []clients.GroupParticipantInput{
			{
				GroupParticipantClientID: 1,
				GroupRoleID:              2,
			},
			{
				GroupParticipantClientID: 2,
				GroupRoleID:              1,
			},
		},
	})
	if err != nil {
		log.Fatalf("creating client group: %v\n", err)
	}

	fmt.Println("created group")

	// Add new clients to group
	group, err = wc.Clients.AddGroupParticipants(ctx, group.Group.ID, clients.GroupParticipantsRequest{
		ClientIDs: []int64{3, 4},
	})
	if err != nil {
		log.Fatalf("adding clients to group: %v\n", err)
	}

	fmt.Println("added clients to group")

	// Remove leads from group
	group, err = wc.Clients.RemoveGroupParticipants(ctx, group.Group.ID, clients.GroupParticipantsRequest{
		ClientIDs: []int64{3},
	})
	if err != nil {
		log.Fatalf("removing clients from group: %v\n", err)
	}

	fmt.Println("removed clients from group")

	// Convert client from dependent
	_, err = wc.Clients.ConvertFromDependent(ctx, 123, clients.ConvertFromDependentRequest{Email: "john.doe@example.com"})
	if err != nil {
		log.Fatalf("converting client from dependent: %v\n", err)
	}

	fmt.Println("converted client from dependent")
}
