package leads

import (
	"github.com/andrew-hayworth22/wodify-go/models"
)

// LeadListItem represents a lead in a list.
type LeadListItem struct {
	// Lead's ID.
	ID int64 `json:"id"`
	// Lead's first name.
	FirstName string `json:"first_name"`
	// Lead's last name.
	LastName string `json:"last_name"`
	// Lead's email address.
	Email string `json:"email"`
	// Lead's status ID.
	LeadStatusID int64 `json:"lead_status_id"`
	// Lead's status label.
	LeadStatus string `json:"lead_status"`
	// Lead's default location ID.
	LocationID int64 `json:"location_id"`
	// Lead's default location name.
	Location string `json:"location"`
	// Lead's gender.
	Gender models.Gender `json:"gender_id"`
	// Lead's phone number.
	PhoneNumber string `json:"phone_number"`
	// Lead's date of birth.
	DateOfBirth models.Date `json:"date_of_birth"`
	// Lead's street address (line 1).
	StreetAddress1 string `json:"street_address1"`
	// Lead's street address (line 2).
	StreetAddress2 string `json:"street_address2"`
	// Lead's city.
	City string `json:"city"`
	// Lead's state.
	StateID int `json:"state_id"`
	// Lead's state name.
	State string `json:"state"`
	// Lead's province, if applicable.
	Province string `json:"province"`
	// Lead's ZIP code.
	ZipCode string `json:"zipcode"`
	// Lead's country ID.
	CountryID int `json:"country_id"`
	// Lead's country name.
	Country string `json:"country"`
	// Lead's tags as a text list.
	Tags []string `json:"tags"`
	// Last time the lead was contacted.
	LastContactDateTime models.DateTime `json:"last_contact_date_time"`
	// Indicates whether the lead has been converted to a client.
	IsConvertedToClient bool `json:"is_converted_to_client"`
	// Name of the Lead's emergency contact.
	EmergencyContactName string `json:"emergency_contact_name"`
	// Phone number of the Lead's emergency contact.
	EmergencyContactPhone string `json:"emergency_contact_phone"`
	// Lead's source ID.
	LeadSourceID int64 `json:"lead_source_id"`
	// Lead's source name.
	LeadSource string `json:"lead_source"`
	// Who referred the lead on the Web (free text).
	ReferredByFromWeb string `json:"referred_by_from_web"`
	// Unique ID of the user that referred the lead.
	ReferredByUserId int64 `json:"referred_by_user_id"`
	// Name of the user that referred the lead.
	ReferredByFromUserName string `json:"referred_by_from_user_name"`
	// Indicates whether the lead has subscribed to email notifications.
	IsEmailSubscribed bool `json:"is_email_subscribed"`
	// Indicates whether the lead has subscribed to SMS notifications.
	IsSMSSubscribed bool `json:"is_sms_subscribed"`
	// ID of the lead's default location timezone.'
	LocationTimezoneID int64 `json:"location_timezone_id"`
	// Timezone of the lead's default location.
	LocationTimezone string `json:"location_timezone"`
	// The ID of the source from which the lead was created.
	CreatedFromSource string `json:"created_from_source"`
	// Lead's profile photo URL.'
	ProfilePhotoURL string `json:"profile_photo_url"`
	// Unique ID of the lead's owner.
	LeadOwnerID int64 `json:"lead_owner_id"`
	// Total number of classes that the Lead has signed in to.
	TotalClassSignIns int `json:"total_class_sign_ins"`
	// Total number of appointment bookings that the Lead has signed in to.
	TotalBookingSignIns int `json:"total_booking_sign_ins"`
	// Last time the Lead signed in to a class.
	LastClassSignIn models.DateTime `json:"last_class_sign_in"`
	// Last time the Lead signed in to an appointment booking.
	LastBookingSignIn models.DateTime `json:"last_booking_sign_in"`
	// Number of days since the Lead last attended a class or appointment booking.
	DaysSinceLastAttendance int `json:"days_since_last_attendance"`
	// Next reservation date and time for the Lead.
	NextClassReservation models.DateTime `json:"next_class_reservation"`
	// Next appointment booking date and time for the Lead.
	NextAppointmentBooking models.DateTime `json:"next_appointment_booking"`
	// Record creation data.
	Created models.Updated `json:"created"`
	// Record last update data.
	Updated models.Updated `json:"updated"`
}

// ListResponse represents a response to a list request.
type ListResponse struct {
	// List of fetched leads.
	Leads []LeadListItem `json:"leads"`
	// Pagination information.
	Pagination models.PaginationResponse `json:"pagination"`
}

// DeleteLeadResponse represents a response to a delete request.
type DeleteLeadResponse struct {
	// ID of the deleted lead.
	LeadID int64 `json:"lead_id"`
	// Indicates whether the deletion was successful.
	IsSuccess bool `json:"is_success"`
}

// ConvertLeadResponse represents a response to a lead conversion to a client.
type ConvertLeadResponse struct {
	// ID of the lead that was converted to a client.
	ConvertedLeadID int64 `json:"converted_lead_id"`
	// Indicates whether the conversion was successful.
	IsSuccess bool `json:"is_success"`
	// Client created from the lead conversion.
	ClientData models.Client `json:"client_data"`
}

// ListStatusesResponse represents a response to a lead status fetch
type ListStatusesResponse struct {
	Statuses   []models.LeadStatus       `json:"statuses"`
	Pagination models.PaginationResponse `json:"pagination"`
}

// ListSourcesResponse represents a response to a lead source fetch
type ListSourcesResponse struct {
	Sources    []models.LeadSource       `json:"sources"`
	Pagination models.PaginationResponse `json:"pagination"`
}

// UpdateTagsResponse represents a response to an add lead tags request
type UpdateTagsResponse struct {
	Tags      []string `json:"active_tags"`
	IsSuccess bool     `json:"is_success"`
}

// ListBookingsResponse represents a response to a lead appointment booking fetch
type ListBookingsResponse struct {
	Bookings                  []models.LeadBooking `json:"lead_appointment_bookings"`
	models.PaginationResponse `json:"pagination"`
}
