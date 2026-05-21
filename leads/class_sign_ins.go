package leads

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/internal/search"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

///////////////////////////////////////////////////////////////////////
// Client methods
///////////////////////////////////////////////////////////////////////

// ListClassSignIns fetches a list of a leads' class sign-ins
func (c *Client) ListClassSignIns(ctx context.Context, id int64, req ListClassSignInsRequest) (*ListClassSignInsResponse, error) {
	var out ListClassSignInsResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/classes/sign-ins", id), req.ToQuery(), nil, &out)
	return &out, err
}

// SearchClassSignIns fetches a list of a leads' class sign-ins matching a search criteria
func (c *Client) SearchClassSignIns(ctx context.Context, id int64, req SearchClassSignInsRequest) (*ListClassSignInsResponse, error) {
	var out ListClassSignInsResponse
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

// ClassSignInSort represents a booking sort order.
type ClassSignInSort = sort.Sort[ClassSignInField]

// NewClassSignInSort creates a new booking sort.
func NewClassSignInSort(field ClassSignInField, isDescending bool) ClassSignInSort {
	return sort.NewSort(field, isDescending)
}

// ClassSignInQuery represents a lead booking query.
type ClassSignInQuery = search.Builder[ClassSignInField]

// NewClassSignInQuery creates a new lead booking query builder.
func NewClassSignInQuery() *ClassSignInQuery {
	return search.New[ClassSignInField]()
}

// ListClassSignInsRequest represents a request to list a lead's bookings.
type ListClassSignInsRequest struct {
	Page models.PaginationRequest
	Sort ClassSignInSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListClassSignInsRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

// SearchClassSignInsRequest represents a request to search a lead's bookings.
type SearchClassSignInsRequest struct {
	Page  models.PaginationRequest
	Sort  ClassSignInSort
	Query *ClassSignInQuery
}

// ToQuery converts the request to URL query string parameters.
func (r SearchClassSignInsRequest) ToQuery() url.Values {
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

// ListClassSignInsResponse represents a response to a lead class sign-ins fetch
type ListClassSignInsResponse struct {
	SignIns                   []models.LeadClassSignIn `json:"class_sign_ins"`
	models.PaginationResponse `json:"pagination"`
}
