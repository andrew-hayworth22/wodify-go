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

// ListReservations fetches a list of a leads' class reservations
func (c *Client) ListReservations(ctx context.Context, id int64, req ReservationListRequest) (*ReservationListResponse, error) {
	var out ReservationListResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/classes/reservations", id), req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// SearchReservations fetches a list of a leads' class reservations matching a query criteria
func (c *Client) SearchReservations(ctx context.Context, id int64, req ReservationSearchRequest) (*ReservationListResponse, error) {
	var out ReservationListResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/classes/reservations/search", id), req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// ReservationField represents a field that class reservation lists can be sorted/filtered by.
type ReservationField string

const (
	ReservationFieldID                      ReservationField = "reservation_id"
	ReservationFieldClassID                 ReservationField = "class_id"
	ReservationFieldClassName               ReservationField = "class"
	ReservationFieldUTCClassStartDateTime   ReservationField = "utc_class_start_datetime"
	ReservationFieldUTCClassEndDateTime     ReservationField = "utc_class_end_datetime"
	ReservationFieldLocalClassStartDateTime ReservationField = "local_class_start_datetime"
	ReservationFieldLocalClassEndDateTime   ReservationField = "local_class_end_datetime"
	ReservationFieldProgramID               ReservationField = "program_id"
	ReservationFieldProgramName             ReservationField = "program"
	ReservationFieldLocationID              ReservationField = "location_id"
	ReservationFieldLocationName            ReservationField = "location"
	ReservationFieldStatusID                ReservationField = "reservation_status_id"
	ReservationFieldStatusName              ReservationField = "reservation_status"
	ReservationFieldMembershipID            ReservationField = "membership_id"
	ReservationFieldMembershipName          ReservationField = "membership"
	ReservationFieldDropInName              ReservationField = "drop_in_name"
	ReservationFieldDropInEmail             ReservationField = "drop_in_email"
	ReservationFieldIsCanceledFromWaitlist  ReservationField = "is_cancelled_from_waitlist"
	ReservationFieldIsLateCancellation      ReservationField = "is_late_cancellation"
)

// ReservationListRequest represents a request to list a lead's class reservations.
type ReservationListRequest = request.ListRequest[ReservationField]

// NewReservationListRequest creates a new ReservationListRequest with the given pagination and sort.
func NewReservationListRequest(pagination request.PaginationRequest, sort sort.Sort[ReservationField]) ReservationListRequest {
	return ReservationListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// ReservationSearchRequest represents a request to query a lead's class reservations.
type ReservationSearchRequest = request.SearchRequest[ReservationField]

// NewReservationSearchRequest creates a new ReservationSearchRequest with the given pagination, sort, and query.
func NewReservationSearchRequest(pagination request.PaginationRequest, sort sort.Sort[ReservationField], query *query.Builder[ReservationField]) ReservationSearchRequest {
	return ReservationSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewReservationQuery creates a new lead booking query builder.
func NewReservationQuery() *query.Builder[ReservationField] {
	return query.New[ReservationField]()
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// ReservationListResponse represents a response to a lead reservations fetch
type ReservationListResponse struct {
	Reservations              []models.LeadReservation `json:"lead_class_reservations"`
	models.PaginationResponse `json:"pagination"`
}
