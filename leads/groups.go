package leads

import (
	"context"
	"net/http"

	"github.com/andrew-hayworth22/wodify-go/internal/query"
	"github.com/andrew-hayworth22/wodify-go/internal/request"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

///////////////////////////////////////////////////////////////////////
// Client methods
///////////////////////////////////////////////////////////////////////

// ListGroupRoles fetches a list of lead group roles
func (c *Client) ListGroupRoles(ctx context.Context, req GroupRoleListRequest) (*GroupRoleListResponse, error) {
	var out GroupRoleListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/leads/group/roles", req.ToQuery(), nil, &out)
	return &out, err
}

// SearchGroupRoles fetches a list of lead group roles matching the search criteria
func (c *Client) SearchGroupRoles(ctx context.Context, req GroupRoleSearchRequest) (*GroupRoleListResponse, error) {
	var out GroupRoleListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/leads/group/roles/search", req.ToQuery(), nil, &out)
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// GroupRoleField represents a field that group lists can be sorted/filtered by.
type GroupRoleField string

const (
	GroupRoleFieldID   GroupRoleField = "id"
	GroupRoleFieldName GroupRoleField = "lead_group_role"
)

// GroupRoleListRequest represents a request to list group roles.
type GroupRoleListRequest = request.ListRequest[GroupRoleField]

// NewGroupRoleListRequest creates a new GroupRoleListRequest with the given pagination and sort.
func NewGroupRoleListRequest(pagination request.PaginationRequest, sort sort.Sort[GroupRoleField]) GroupRoleListRequest {
	return GroupRoleListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// GroupRoleSearchRequest represents a request to query group roles.
type GroupRoleSearchRequest = request.SearchRequest[GroupRoleField]

func NewGroupRoleSearchRequest(pagination request.PaginationRequest, sort sort.Sort[GroupRoleField], query *query.Builder[GroupRoleField]) GroupRoleSearchRequest {
	return GroupRoleSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewGroupRoleQuery creates a new group role query builder.
func NewGroupRoleQuery() *query.Builder[GroupRoleField] {
	return query.New[GroupRoleField]()
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// GroupRoleListResponse represents a response to a list group roles request.
type GroupRoleListResponse struct {
	Roles      []models.LeadGroupRole    `json:"lead_group_roles"`
	Pagination models.PaginationResponse `json:"pagination"`
}
