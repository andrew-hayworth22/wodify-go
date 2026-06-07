package utils

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

// ListObjectTypes fetches a list of object types.
func (c *Client) ListObjectTypes(ctx context.Context, req ObjectTypeListRequest) (*ObjectTypeListResponse, error) {
	var out ObjectTypeListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/object-types", req.ToQuery(), nil, &out)
	return &out, err
}

// SearchObjectTypes fetches a list of object types matching a query criteria.
func (c *Client) SearchObjectTypes(ctx context.Context, req ObjectTypeSearchRequest) (*ObjectTypeListResponse, error) {
	var out ObjectTypeListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/object-types/search", req.ToQuery(), nil, &out)
	return &out, err
}

// ListObjectActionTypes fetches a list of object action types.
func (c *Client) ListObjectActionTypes(ctx context.Context, req ObjectActionTypeListRequest) (*ObjectActionTypeListResponse, error) {
	var out ObjectActionTypeListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/object-type-action-types", req.ToQuery(), nil, &out)
	return &out, err
}

// SearchObjectActionTypes fetches a list of object action types matching a query criteria.
func (c *Client) SearchObjectActionTypes(ctx context.Context, req ObjectActionTypeSearchRequest) (*ObjectActionTypeListResponse, error) {
	var out ObjectActionTypeListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/object-type-action-types/search", req.ToQuery(), nil, &out)
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// ObjectTypeField represents a field that object types can be sorted/filtered by.
type ObjectTypeField string

const (
	ObjectTypeFieldID   ObjectTypeField = "object_type_id"
	ObjectTypeFieldName ObjectTypeField = "object_type"
)

// ObjectTypeListRequest represents a request to list object types.
type ObjectTypeListRequest = request.ListRequest[ObjectTypeField]

// NewObjectTypeListRequest creates a new ObjectTypeListRequest with the given pagination and sort.
func NewObjectTypeListRequest(pagination request.PaginationRequest, sort sort.Sort[ObjectTypeField]) ObjectTypeListRequest {
	return ObjectTypeListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// ObjectTypeSearchRequest represents a request to query object types.
type ObjectTypeSearchRequest = request.SearchRequest[ObjectTypeField]

// NewObjectTypeSearchRequest creates a new ObjectTypeSearchRequest with the given pagination, sort, and query.
func NewObjectTypeSearchRequest(pagination request.PaginationRequest, sort sort.Sort[ObjectTypeField], query *query.Builder[ObjectTypeField]) ObjectTypeSearchRequest {
	return ObjectTypeSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewObjectTypeQuery creates a new object type query query builder.
func NewObjectTypeQuery() *query.Builder[ObjectTypeField] {
	return query.New[ObjectTypeField]()
}

// ObjectActionTypeField represents a field that object action types can be sorted/filtered by.
type ObjectActionTypeField string

const (
	ObjectActionTypeFieldID             ObjectActionTypeField = "action_type_id"
	ObjectActionTypeFieldName           ObjectActionTypeField = "action_type"
	ObjectActionTypeFieldObjectTypeID   ObjectActionTypeField = "object_type_id"
	ObjectActionTypeFieldObjectTypeName ObjectActionTypeField = "object_type"
	ObjectActionTypeFieldRoleID         ObjectActionTypeField = "action_role_id"
	ObjectActionTypeFieldRoleName       ObjectActionTypeField = "action_role"
)

// ObjectActionTypeListRequest represents a request to list object action types.
type ObjectActionTypeListRequest = request.ListRequest[ObjectActionTypeField]

// NewObjectActionTypeListRequest creates a new ObjectActionTypeListRequest with the given pagination and sort.
func NewObjectActionTypeListRequest(pagination request.PaginationRequest, sort sort.Sort[ObjectActionTypeField]) ObjectActionTypeListRequest {
	return ObjectActionTypeListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// ObjectActionTypeSearchRequest represents a request to query object action types.
type ObjectActionTypeSearchRequest = request.SearchRequest[ObjectActionTypeField]

// NewObjectActionTypeSearchRequest creates a new ObjectActionTypeSearchRequest with the given pagination, sort, and query.
func NewObjectActionTypeSearchRequest(pagination request.PaginationRequest, sort sort.Sort[ObjectActionTypeField], query *query.Builder[ObjectActionTypeField]) ObjectActionTypeSearchRequest {
	return ObjectActionTypeSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewObjectActionTypeQuery creates a new object action type query query builder.
func NewObjectActionTypeQuery() *query.Builder[ObjectActionTypeField] {
	return query.New[ObjectActionTypeField]()
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// ObjectTypeListResponse represents a response to a list object types request.
type ObjectTypeListResponse struct {
	ObjectTypes []models.ObjectType       `json:"object_types"`
	Pagination  models.PaginationResponse `json:"pagination"`
}

// ObjectActionTypeListResponse represents a response to a list object action types request.
type ObjectActionTypeListResponse struct {
	ObjectActionTypes []models.ObjectActionType `json:"object_type_actions"`
	Pagination        models.PaginationResponse `json:"pagination"`
}
