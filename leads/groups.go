package leads

import (
	"context"
	"fmt"
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
	if err != nil {
		return nil, err
	}
	return &out, err
}

// SearchGroupRoles fetches a list of lead group roles matching the search criteria
func (c *Client) SearchGroupRoles(ctx context.Context, req GroupRoleSearchRequest) (*GroupRoleListResponse, error) {
	if req.Query != nil {
		if err := req.Query.Err(); err != nil {
			return nil, err
		}
	}
	var out GroupRoleListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/leads/group/roles/search", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// CreateGroup creates a new lead group
func (c *Client) CreateGroup(ctx context.Context, req GroupCreateRequest) (*GroupResponse, error) {
	var out GroupResponse
	err := c.hc.Do(ctx, http.MethodPost, "/leads/group/participants", nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// AddGroupParticipants adds lead participants to a group
func (c *Client) AddGroupParticipants(ctx context.Context, groupID int64, req GroupParticipantsRequest) (*GroupResponse, error) {
	var out GroupResponse
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/leads/group/%d/participants", groupID), nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// RemoveGroupParticipants removes lead participants from a group
func (c *Client) RemoveGroupParticipants(ctx context.Context, groupID int64, req GroupParticipantsRequest) (*GroupResponse, error) {
	var out GroupResponse
	err := c.hc.Do(ctx, http.MethodDelete, fmt.Sprintf("/leads/group/%d/participants", groupID), nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// ConvertFromDependent converts a lead to a full group member from a dependent
func (c *Client) ConvertFromDependent(ctx context.Context, leadID int64, req ConvertFromDependentRequest) (*ConvertFromDependentResponse, error) {
	var out ConvertFromDependentResponse
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/leads/%d/convert-from-dependent", leadID), nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
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

// GroupCreateRequest represents a request to create a new LeadGroup.
type GroupCreateRequest struct {
	// Participants of the new group.
	GroupParticipants []GroupParticipantInput `json:"group_participants"`
}

// GroupParticipantInput represents a group participant used when creating a LeadGroup.
type GroupParticipantInput struct {
	// ID of the group member.
	GroupParticipantLeadID int64 `json:"group_participant_lead_id"`
	// ID of the group member's role.'
	GroupRoleID int64 `json:"group_role_id"`
}

// GroupParticipantsRequest represents a request to add or remove participants from a group.
type GroupParticipantsRequest struct {
	LeadIDs []int64 `json:"group_participants"`
}

// ConvertFromDependentRequest represents a request to convert a lead to a full group member from a dependent.
type ConvertFromDependentRequest struct {
	// New email of the full group member.
	Email string `json:"email"`
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// GroupRoleListResponse represents a response to a list group roles request.
type GroupRoleListResponse struct {
	Roles      []models.LeadGroupRole    `json:"lead_group_roles"`
	Pagination models.PaginationResponse `json:"pagination"`
}

// GroupResponse represents a response to a group request.
type GroupResponse struct {
	Group models.LeadGroup `json:"lead_group"`
}

// ConvertFromDependentResponse represents a response to a lead dependent conversion request.
type ConvertFromDependentResponse struct {
	IsSuccess bool        `json:"is_success"`
	Lead      models.Lead `json:"lead_data"`
}
