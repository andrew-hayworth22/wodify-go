package utils

import (
	"context"
	"net/http"
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/internal/search"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

///////////////////////////////////////////////////////////////////////
// Client methods
///////////////////////////////////////////////////////////////////////

// ListObjectTypes fetches a list of object types.
func (c *Client) ListObjectTypes(ctx context.Context, req ListObjectTypesRequest) (*ListObjectTypesResponse, error) {
	var out ListObjectTypesResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/object-types", req.ToQuery(), nil, &out)
	return &out, err
}

// SearchObjectTypes fetches a list of object types matching a search criteria.
func (c *Client) SearchObjectTypes(ctx context.Context, req SearchObjectTypesRequest) (*ListObjectTypesResponse, error) {
	var out ListObjectTypesResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/object-types/search", req.ToQuery(), nil, &out)
	return &out, err
}

// ListObjectActionTypes fetches a list of object action types.
func (c *Client) ListObjectActionTypes(ctx context.Context, req ListObjectActionTypesRequest) (*ListObjectActionTypesResponse, error) {
	var out ListObjectActionTypesResponse
	err := c.hc.Do(ctx, http.MethodGet, "/utilities/object-type-action-types", req.ToQuery(), nil, &out)
	return &out, err
}

// SearchObjectActionTypes fetches a list of object action types matching a search criteria.
func (c *Client) SearchObjectActionTypes(ctx context.Context, req SearchObjectActionTypesRequest) (*ListObjectActionTypesResponse, error) {
	var out ListObjectActionTypesResponse
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

// ObjectTypeSort represents an object type sort order.
type ObjectTypeSort = sort.Sort[ObjectTypeField]

// NewObjectTypeSort creates a new object type sort.
func NewObjectTypeSort(field ObjectTypeField, isDescending bool) ObjectTypeSort {
	return sort.NewSort(field, isDescending)
}

// ObjectTypeQuery represents an object type search query.
type ObjectTypeQuery = search.Builder[ObjectTypeField]

// NewObjectTypeQuery creates a new object type search query builder.
func NewObjectTypeQuery() *ObjectTypeQuery {
	return search.New[ObjectTypeField]()
}

// ListObjectTypesRequest represents a request to list object types.
type ListObjectTypesRequest struct {
	Page models.PaginationRequest
	Sort ObjectTypeSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListObjectTypesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

// SearchObjectTypesRequest represents a request to search object types.
type SearchObjectTypesRequest struct {
	Page  models.PaginationRequest
	Sort  ObjectTypeSort
	Query *ObjectTypeQuery
}

// ToQuery converts the request to URL query string parameters.
func (r SearchObjectTypesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	if r.Query != nil {
		q.Set("q", r.Query.String())
	}
	return q
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

// ObjectActionTypeSort represents an object action type sort order.
type ObjectActionTypeSort = sort.Sort[ObjectActionTypeField]

// NewObjectActionTypeSort creates a new object action type sort.
func NewObjectActionTypeSort(field ObjectActionTypeField, isDescending bool) ObjectActionTypeSort {
	return sort.NewSort(field, isDescending)
}

// ObjectActionTypeQuery represents an object action type search query.
type ObjectActionTypeQuery = search.Builder[ObjectActionTypeField]

// NewObjectActionTypeQuery creates a new object action type search query builder.
func NewObjectActionTypeQuery() *ObjectActionTypeQuery {
	return search.New[ObjectActionTypeField]()
}

// ListObjectActionTypesRequest represents a request to list object action types.
type ListObjectActionTypesRequest struct {
	Page models.PaginationRequest
	Sort ObjectActionTypeSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListObjectActionTypesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

// SearchObjectActionTypesRequest represents a request to search object action types.
type SearchObjectActionTypesRequest struct {
	Page  models.PaginationRequest
	Sort  ObjectActionTypeSort
	Query *ObjectActionTypeQuery
}

// ToQuery converts the request to URL query string parameters.
func (r SearchObjectActionTypesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	if r.Query != nil {
		q.Set("q", r.Query.String())
	}
	return q
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// ListObjectTypesResponse represents a response to a list object types request.
type ListObjectTypesResponse struct {
	ObjectTypes []models.ObjectType `json:"object_types"`
	Pagination  models.PaginationResponse
}

// ListObjectActionTypesResponse represents a response to a list object action types request.
type ListObjectActionTypesResponse struct {
	ObjectActionTypes []models.ObjectActionType `json:"object_type_actions"`
	Pagination        models.PaginationResponse
}
