package leads

import (
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/internal/search"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

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

// ReservationSort represents a booking sort order.
type ReservationSort = sort.Sort[ReservationField]

// NewReservationSort creates a new booking sort.
func NewReservationSort(field ReservationField, isDescending bool) ReservationSort {
	return sort.NewSort(field, isDescending)
}

// ReservationQuery represents a lead booking query.
type ReservationQuery = search.Builder[ReservationField]

// NewReservationQuery creates a new lead booking query builder.
func NewReservationQuery() *ReservationQuery {
	return search.New[ReservationField]()
}

// ListReservationsRequest represents a request to list a lead's bookings.
type ListReservationsRequest struct {
	Page models.PaginationRequest
	Sort ReservationSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListReservationsRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

// SearchReservationsRequest represents a request to search a lead's bookings.
type SearchReservationsRequest struct {
	Page  models.PaginationRequest
	Sort  ReservationSort
	Query *ReservationQuery
}

// ToQuery converts the request to URL query string parameters.
func (r SearchReservationsRequest) ToQuery() url.Values {
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

// ListReservationsResponse represents a response to a lead reservations fetch
type ListReservationsResponse struct {
	Reservations              []models.LeadReservation `json:"lead_class_reservations"`
	models.PaginationResponse `json:"pagination"`
}
