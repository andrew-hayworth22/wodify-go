package clients

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

// ListBookings fetches a list of appointment bookings for a client
func (c *Client) ListBookings(ctx context.Context, id int64, req BookingListRequest) (*BookingListResponse, error) {
	var out BookingListResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/clients/%d/appointments/bookings", id), req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// SearchBookings fetches a list of a client's bookings matching a query criteria
func (c *Client) SearchBookings(ctx context.Context, id int64, req BookingSearchRequest) (*BookingListResponse, error) {
	if req.Query != nil {
		if err := req.Query.Err(); err != nil {
			return nil, err
		}
	}
	var out BookingListResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/clients/%d/appointments/bookings/search", id), req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// BookingField represents a field that booking lists can be sorted/filtered by.
type BookingField string

const (
	BookingFieldID                            BookingField = "booking_id"
	BookingFieldAppointmentID                 BookingField = "appointment_id"
	BookingFieldUTCAppointmentStartDateTime   BookingField = "utc_appointment_start_datetime"
	BookingFieldUTCAppointmentEndDateTime     BookingField = "utc_appointment_end_datetime"
	BookingFieldLocalAppointmentStartDateTime BookingField = "local_appointment_start_datetime"
	BookingFieldLocalAppointmentEndDateTime   BookingField = "local_appointment_end_datetime"
	BookingFieldServiceID                     BookingField = "service_id"
	BookingFieldServiceName                   BookingField = "service"
	BookingFieldLocationID                    BookingField = "location_id"
	BookingFieldLocationName                  BookingField = "location"
	BookingFieldProviderID                    BookingField = "provider_id"
	BookingFieldProviderName                  BookingField = "provider"
	BookingFieldStatusID                      BookingField = "booking_status_id"
	BookingFieldStatus                        BookingField = "booking_status"
	BookingFieldMembershipID                  BookingField = "membership_id"
	BookingFieldMembershipName                BookingField = "membership"
	BookingFieldIsFreeTrial                   BookingField = "is_free_trial"
	BookingFieldIsLateCancellation            BookingField = "is_late_cancellation"
)

// BookingListRequest represents a request to list a client's bookings.
type BookingListRequest = request.ListRequest[BookingField]

// NewBookingListRequest creates a new BookingListRequest with the given pagination and sort.
func NewBookingListRequest(pagination request.PaginationRequest, sort sort.Sort[BookingField]) BookingListRequest {
	return BookingListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// BookingSearchRequest represents a request to query a client's bookings.
type BookingSearchRequest = request.SearchRequest[BookingField]

// NewBookingSearchRequest creates a new BookingSearchRequest with the given pagination, sort, and query.
func NewBookingSearchRequest(pagination request.PaginationRequest, sort sort.Sort[BookingField], query *query.Builder[BookingField]) BookingSearchRequest {
	return BookingSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewBookingQuery creates a new client booking query builder.
func NewBookingQuery() *query.Builder[BookingField] {
	return query.New[BookingField]()
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// BookingListResponse represents a response to a client appointment booking fetch
type BookingListResponse struct {
	Bookings                  []models.ClientBooking `json:"client_appointment_bookings"`
	models.PaginationResponse `json:"pagination"`
}
