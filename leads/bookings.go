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

// ListBookings fetches a list of appointment bookings for a Lead
func (c *Client) ListBookings(ctx context.Context, id int64, req ListBookingsRequest) (*ListBookingsResponse, error) {
	var out ListBookingsResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/appointments/bookings", id), req.ToQuery(), nil, &out)
	return &out, err
}

// SearchBookings fetches a list of a leads' bookings matching a search criteria
func (c *Client) SearchBookings(ctx context.Context, id int64, req SearchBookingsRequest) (*ListBookingsResponse, error) {
	var out ListBookingsResponse
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d/appointments/bookings/search", id), req.ToQuery(), nil, &out)
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

// BookingSort represents a booking sort order.
type BookingSort = sort.Sort[BookingField]

// NewBookingSort creates a new booking sort.
func NewBookingSort(field BookingField, isDescending bool) BookingSort {
	return sort.NewSort(field, isDescending)
}

// BookingQuery represents a lead booking query.
type BookingQuery = search.Builder[BookingField]

// NewBookingQuery creates a new lead booking query builder.
func NewBookingQuery() *BookingQuery {
	return search.New[BookingField]()
}

// ListBookingsRequest represents a request to list a lead's bookings.
type ListBookingsRequest struct {
	Page models.PaginationRequest
	Sort BookingSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListBookingsRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

// SearchBookingsRequest represents a request to search a lead's bookings.
type SearchBookingsRequest struct {
	Page  models.PaginationRequest
	Sort  BookingSort
	Query *BookingQuery
}

// ToQuery converts the request to URL query string parameters.
func (r SearchBookingsRequest) ToQuery() url.Values {
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

// ListBookingsResponse represents a response to a lead appointment booking fetch
type ListBookingsResponse struct {
	Bookings                  []models.LeadBooking `json:"lead_appointment_bookings"`
	models.PaginationResponse `json:"pagination"`
}
