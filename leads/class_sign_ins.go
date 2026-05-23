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

// ListClassSignIns fetches a list of a leads' class sign-ins
func (c *Client) ListClassSignIns(ctx context.Context, id int64, req ClassSignInListRequest) (*ClassSignInListResponse, error) {
	var out ClassSignInListResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/classes/sign-ins", id), req.ToQuery(), nil, &out)
	return &out, err
}

// SearchClassSignIns fetches a list of a leads' class sign-ins matching a query criteria
func (c *Client) SearchClassSignIns(ctx context.Context, id int64, req ClassSignInSearchRequest) (*ClassSignInListResponse, error) {
	var out ClassSignInListResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/classes/sign-ins/search", id), req.ToQuery(), nil, &out)
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// ClassSignInField represents a field that class sign-in lists can be sorted/filtered by.
type ClassSignInField string

const (
	ClassSignInFieldID                             ClassSignInField = "id"
	ClassSignInFieldClassID                        ClassSignInField = "class_id"
	ClassSignInFieldClassName                      ClassSignInField = "class"
	ClassSignInFieldUTCClassStartDateTime          ClassSignInField = "utc_class_start_datetime"
	ClassSignInFieldUTCClassEndDateTime            ClassSignInField = "utc_class_end_datetime"
	ClassSignInFieldLocalClassStartDateTime        ClassSignInField = "local_class_start_datetime"
	ClassSignInFieldLocalClassEndDateTime          ClassSignInField = "local_class_end_datetime"
	ClassSignInFieldProgramID                      ClassSignInField = "program_id"
	ClassSignInFieldProgramName                    ClassSignInField = "program"
	ClassSignInFieldLocationID                     ClassSignInField = "location_id"
	ClassSignInFieldLocationName                   ClassSignInField = "location"
	ClassSignInFieldMembershipID                   ClassSignInField = "membership_id"
	ClassSignInFieldMembershipName                 ClassSignInField = "membership"
	ClassSignInFieldIsDropIn                       ClassSignInField = "is_drop_in"
	ClassSignInFieldIsAutoSignIn                   ClassSignInField = "is_auto_sign_in"
	ClassSignInFieldCountsTowardsAttendanceLimits  ClassSignInField = "counts_towards_attendance_limits"
	ClassSignInFieldIsMembershipEnforcementEnabled ClassSignInField = "is_membership_enforcement_enabled"
	ClassSignInFieldLimitedPlanEnforcementTypeID   ClassSignInField = "limited_plan_enforcement_type_id"
	ClassSignInFieldLimitedPlanEnforcementTypeName ClassSignInField = "limited_plan_enforcement_type"
	ClassSignInFieldClassPackEnforcementTypeID     ClassSignInField = "class_pack_enforcement_type_id"
	ClassSignInFieldClassPackEnforcementTypeName   ClassSignInField = "class_pack_enforcement_type"
	ClassSignInFieldIsAttendedEmailSent            ClassSignInField = "is_attended_email_sent"
)

// ClassSignInListRequest represents a request to list a lead's class sign-ins.
type ClassSignInListRequest = request.ListRequest[ClassSignInField]

// NewClassSignInListRequest creates a new ClassSignInListRequest with the given pagination and sort.
func NewClassSignInListRequest(pagination request.PaginationRequest, sort sort.Sort[ClassSignInField]) ClassSignInListRequest {
	return ClassSignInListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// ClassSignInSort represents a booking sort order.
type ClassSignInSort = sort.Sort[ClassSignInField]

// ClassSignInSearchRequest represents a request to query a lead's bookings.
type ClassSignInSearchRequest = request.SearchRequest[ClassSignInField]

// NewClassSignInSearchRequest creates a new ClassSignInSearchRequest with the given pagination, sort, and query.
func NewClassSignInSearchRequest(pagination request.PaginationRequest, sort sort.Sort[ClassSignInField], query *query.Builder[ClassSignInField]) ClassSignInSearchRequest {
	return ClassSignInSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewClassSignInQuery creates a new lead class sign-in query builder.
func NewClassSignInQuery() *query.Builder[ClassSignInField] {
	return query.New[ClassSignInField]()
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// ClassSignInListResponse represents a response to a lead class sign-ins fetch
type ClassSignInListResponse struct {
	SignIns                   []models.LeadClassSignIn `json:"class_sign_ins"`
	models.PaginationResponse `json:"pagination"`
}
