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

// Get fetches a single lead by ID.
func (c *Client) Get(ctx context.Context, id int64) (*models.Lead, error) {
	var out models.Lead
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/leads/%d", id), nil, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// List fetches a list of leads.
func (c *Client) List(ctx context.Context, req LeadListRequest) (*LeadListResponse, error) {
	var out LeadListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/leads", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Search fetches a list of leads matching the query criteria.
func (c *Client) Search(ctx context.Context, req LeadSearchRequest) (*LeadListResponse, error) {
	var out LeadListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/leads/search", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Create creates a new lead.
func (c *Client) Create(ctx context.Context, req LeadCreateRequest) (*models.Lead, error) {
	var out models.Lead
	err := c.hc.Do(ctx, http.MethodPost, "/leads", nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Delete deletes a lead by ID.
func (c *Client) Delete(ctx context.Context, id int64) (*LeadDeleteResponse, error) {
	var out LeadDeleteResponse
	err := c.hc.Do(ctx, http.MethodDelete, fmt.Sprintf("/leads/%d", id), nil, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Update updates a lead by ID.
func (c *Client) Update(ctx context.Context, id int64, req LeadUpdateRequest) (*models.Lead, error) {
	var out models.Lead
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/leads/%d", id), nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Convert converts a lead to a Client.
func (c *Client) Convert(ctx context.Context, id int64, req LeadConvertRequest) (*LeadConvertResponse, error) {
	var out LeadConvertResponse
	err := c.hc.Do(ctx, http.MethodPost, fmt.Sprintf("/leads/%d/convert", id), nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// LeadField represents a field that lead lists can be sorted/filtered by.
type LeadField string

const (
	LeadFieldID                      LeadField = "id"
	LeadFieldFirstName               LeadField = "first_name"
	LeadFieldLastName                LeadField = "last_name"
	LeadFieldEmail                   LeadField = "email"
	LeadFieldStatusID                LeadField = "lead_status_id"
	LeadFieldStatusName              LeadField = "lead_status"
	LeadFieldLocationID              LeadField = "location_id"
	LeadFieldLocationName            LeadField = "location"
	LeadFieldGenderID                LeadField = "gender_id"
	LeadFieldGenderName              LeadField = "gender"
	LeadFieldPhoneNumber             LeadField = "phone_number"
	LeadFieldDateOfBirth             LeadField = "date_of_birth"
	LeadFieldStreetAddress1          LeadField = "street_address1"
	LeadFieldStreetAddress2          LeadField = "street_address2"
	LeadFieldCity                    LeadField = "city"
	LeadFieldStateID                 LeadField = "state_id"
	LeadFieldStateName               LeadField = "state"
	LeadFieldProvince                LeadField = "province"
	LeadFieldZipCode                 LeadField = "zipcode"
	LeadFieldCountryID               LeadField = "country_id"
	LeadFieldCountryName             LeadField = "country"
	LeadFieldTags                    LeadField = "tags"
	LeadFieldNotes                   LeadField = "notes"
	LeadFieldIsConvertedToClient     LeadField = "is_converted_to_client"
	LeadFieldEmergencyContactName    LeadField = "emergency_contact_name"
	LeadFieldEmergencyContactPhone   LeadField = "emergency_contact_phone"
	LeadFieldLeadSourceID            LeadField = "lead_source_id"
	LeadFieldLeadSourceName          LeadField = "lead_source"
	LeadFieldReferredByFromWeb       LeadField = "referred_by_from_web"
	LeadFieldReferredByUserID        LeadField = "referred_by_user_id"
	LeadFieldReferredByFromUserName  LeadField = "referred_by_from_user_name"
	LeadFieldIsEmailSubscribed       LeadField = "is_email_subscribed"
	LeadFieldIsSMSSubscribed         LeadField = "is_sms_subscribed"
	LeadFieldLocationTimezoneID      LeadField = "location_timezone_id"
	LeadFieldLocationTimezoneName    LeadField = "location_timezone"
	LeadFieldCreatedFromSource       LeadField = "created_from_source"
	LeadFieldProfilePhotoURL         LeadField = "profile_photo_url"
	LeadFieldLeadOwnerID             LeadField = "lead_owner_id"
	LeadFieldTotalClassSignIns       LeadField = "total_class_sign_ins"
	LeadFieldTotalBookingSignIns     LeadField = "total_booking_sign_ins"
	LeadFieldLastClassSignIn         LeadField = "last_class_sign_in"
	LeadFieldLastBookingSignIn       LeadField = "last_booking_sign_in"
	LeadFieldDaysSinceLastAttendance LeadField = "days_since_last_attendance"
	LeadFieldIsActive                LeadField = "is_active"
	LeadFieldNextClassReservation    LeadField = "next_class_reservation"
	LeadFieldNextAppointmentBooking  LeadField = "next_appointment_booking"
)

// LeadListRequest represents a request to list leads.
type LeadListRequest = request.ListRequest[LeadField]

// NewLeadListRequest creates a new LeadListRequest with the given pagination and sort.
func NewLeadListRequest(pagination request.PaginationRequest, sort sort.Sort[LeadField]) LeadListRequest {
	return LeadListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// LeadSearchRequest represents a request to query leads.
type LeadSearchRequest = request.SearchRequest[LeadField]

// NewLeadSearchRequest creates a new LeadSearchRequest with the given pagination, sort, and query.
func NewLeadSearchRequest(pagination request.PaginationRequest, sort sort.Sort[LeadField], query *query.Builder[LeadField]) LeadSearchRequest {
	return LeadSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewLeadQuery creates a new lead query builder.
func NewLeadQuery() *query.Builder[LeadField] {
	return query.New[LeadField]()
}

// LeadCreateRequest represents a request to create a new lead.
type LeadCreateRequest struct {
	// Lead's first name.
	FirstName string `json:"first_name"`
	// Lead's last name.
	LastName string `json:"last_name"`
	// Lead's email address.
	Email string `json:"email"`
	// Lead's status ID.
	LeadStatusID int64 `json:"lead_status_id"`
	// Lead's default location ID.
	LocationID int64 `json:"location_id"`
	// Lead's gender ID.
	GenderID int `json:"gender_id"`
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
	// Lead's province, if applicable.
	Province string `json:"province"`
	// Lead's ZIP code.
	ZipCode string `json:"zipcode"`
	// Lead's country ID.
	CountryID int `json:"country_id"`
	// Lead's tags as a text list.
	Tags []string `json:"tags"`
	// Lead's notes.
	Notes string `json:"notes"`
	// Lead's Emergency contact name.
	EmergencyContactName string `json:"emergency_contact_name"`
	// Lead's Emergency contact phone number.
	EmergencyContactPhone string `json:"emergency_contact_phone"`
	// Lead's source ID.
	LeadSourceID int64 `json:"lead_source_id"`
	// Who referred the Lead on the Web (free text).
	ReferredByFromWeb string `json:"referred_by_from_web"`
	// Unique ID of the user that referred the Lead.
	ReferredByUserID int64 `json:"referred_by_user_id"`
	// Indicates whether the Lead has subscribed to email notifications.
	IsEmailSubscribed bool `json:"is_email_subscribed"`
	// Indicates whether the Lead has subscribed to SMS notifications.
	IsSMSSubscribed bool `json:"is_sms_subscribed"`
	// Unique ID of the Lead's owner.
	LeadOwnerID int64 `json:"lead_owner_id"`
}

// LeadUpdateRequest represents a request to update a lead.
type LeadUpdateRequest struct {
	// Lead's first name.
	FirstName string `json:"first_name"`
	// Lead's last name.
	LastName string `json:"last_name"`
	// Lead's email address.
	Email string `json:"email"`
	// Lead's status ID.
	LeadStatusID int64 `json:"lead_status_id"`
	// Lead's default location ID.
	LocationID int64 `json:"location_id"`
	// Lead's gender ID.
	GenderID int `json:"gender_id"`
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
	// Lead's province, if applicable.
	Province string `json:"province"`
	// Lead's ZIP code.
	ZipCode string `json:"zipcode"`
	// Lead's country ID.
	CountryID int `json:"country_id"`
	// Lead's notes.
	Notes string `json:"notes"`
	// Lead's Emergency contact name.
	EmergencyContactName string `json:"emergency_contact_name"`
	// Lead's Emergency contact phone number.
	EmergencyContactPhone string `json:"emergency_contact_phone"`
	// Lead's source ID.
	LeadSourceID int64 `json:"lead_source_id"`
	// Who referred the Lead on the Web (free text).
	ReferredByFromWeb string `json:"referred_by_from_web"`
	// Unique ID of the user that referred the Lead.
	ReferredByUserID int64 `json:"referred_by_user_id"`
	// Indicates whether the Lead has subscribed to email notifications.
	IsEmailSubscribed bool `json:"is_email_subscribed"`
	// Indicates whether the Lead has subscribed to SMS notifications.
	IsSMSSubscribed bool `json:"is_sms_subscribed"`
	// Unique ID of the Lead's owner.
	LeadOwnerID int64 `json:"lead_owner_id"`
}

// LeadUpdateRequestFrom converts a Lead to a LeadUpdateRequest.
func LeadUpdateRequestFrom(l *models.Lead) LeadUpdateRequest {
	return LeadUpdateRequest{
		FirstName:             l.FirstName,
		LastName:              l.LastName,
		Email:                 l.Email,
		LeadStatusID:          l.LeadStatusID,
		LocationID:            l.LocationID,
		GenderID:              l.GenderID,
		PhoneNumber:           l.PhoneNumber,
		DateOfBirth:           l.DateOfBirth,
		StreetAddress1:        l.StreetAddress1,
		StreetAddress2:        l.StreetAddress2,
		City:                  l.City,
		StateID:               l.StateID,
		Province:              l.Province,
		ZipCode:               l.ZipCode,
		CountryID:             l.CountryID,
		Notes:                 l.Notes,
		EmergencyContactName:  l.EmergencyContactName,
		EmergencyContactPhone: l.EmergencyContactPhone,
		LeadSourceID:          l.LeadSourceID,
		ReferredByFromWeb:     l.ReferredByFromWeb,
		ReferredByUserID:      l.ReferredByUserID,
		IsEmailSubscribed:     l.IsEmailSubscribed,
		IsSMSSubscribed:       l.IsSMSSubscribed,
		LeadOwnerID:           l.LeadOwnerID,
	}
}

// LeadConvertRequest represents a request to convert a lead to a client.
type LeadConvertRequest struct {
	// ID of the converted lead's default location.
	LocationID int64 `json:"location_id"`
	// Email the client will have after conversion.
	Email string `json:"email"`
	// First name the client will have after conversion.
	FirstName string `json:"first_name"`
	// Last name the client will have after conversion.
	LastName string `json:"last_name"`
	// Status of the status that the client will have after conversion.
	ClientStatusID int64 `json:"client_status_id"`
	// Gender ID the client will have after conversion.
	GenderID int `json:"gender_id"`
	// Billing credit card email the client will have after conversion.
	BillingCCEmail string `json:"billing_cc_email"`
	// Mobile number the client will have after conversion.
	MobileNumber string `json:"mobile_number"`
	// Date of birth the client will have after conversion.
	DateOfBirth models.Date `json:"date_of_birth"`
	// Street address (line 1) the client will have after conversion.
	StreetAddress1 string `json:"street_address1"`
	// Street address (line 2) the client will have after conversion.
	StreetAddress2 string `json:"street_address2"`
	// City the client will have after conversion.
	City string `json:"city"`
	// State the client will have after conversion.
	StateID int `json:"state_id"`
	// Province the client will have after conversion.
	Province string `json:"province"`
	// Country the client will have after conversion.
	CountryID int `json:"country_id"`
	// ZIP code the client will have after conversion.
	ZipCode string `json:"zipcode"`
	// Client owner the client will have after conversion.
	ClientOwnerID int64 `json:"client_owner_id"`
}

// LeadConvertRequestFrom converts a Lead to a LeadConvertRequest.
func LeadConvertRequestFrom(l *models.Lead) LeadConvertRequest {
	return LeadConvertRequest{
		LocationID:     l.LocationID,
		Email:          l.Email,
		FirstName:      l.FirstName,
		LastName:       l.LastName,
		ClientStatusID: 0,
		GenderID:       l.GenderID,
		BillingCCEmail: "",
		MobileNumber:   l.PhoneNumber,
		DateOfBirth:    l.DateOfBirth,
		StreetAddress1: l.StreetAddress1,
		StreetAddress2: l.StreetAddress2,
		City:           l.City,
		StateID:        l.StateID,
		Province:       l.Province,
		CountryID:      l.CountryID,
		ZipCode:        l.ZipCode,
		ClientOwnerID:  l.LeadOwnerID,
	}
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

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
	// Lead's gender ID.
	GenderID int `json:"gender_id"`
	// Lead's gender name.
	GenderName string `json:"gender"`
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
	StateName string `json:"state"`
	// Lead's province, if applicable.
	Province string `json:"province"`
	// Lead's ZIP code.
	ZipCode string `json:"zipcode"`
	// Lead's country ID.
	CountryID int `json:"country_id"`
	// Lead's country name.
	CountryName string `json:"country"`
	// Lead's tags as a text list.
	Tags []string `json:"tags"`
	// Last time the lead was contacted.
	LastContactDateTime models.DateTime `json:"last_contact_datetime"`
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
	ReferredByUserID int64 `json:"referred_by_user_id"`
	// Name of the user that referred the lead.
	ReferredByUserName string `json:"referred_by_user_name"`
	// ID of the lead's default location timezone.'
	LocationTimezoneID int64 `json:"location_timezone_id"`
	// Timezone of the lead's default location.
	LocationTimezoneName string `json:"location_time_zone"`
	// The ID of the source from which the lead was created.
	CreatedFromSourceID int64 `json:"created_from_source_id"`
	// The Name of the source from which the lead was created.
	CreatedFromSourceName string `json:"created_from_source"`
	// Lead's profile photo URL.'
	ProfilePhotoURL string `json:"profile_photo_url"`
	// Unique ID of the lead's owner.
	LeadOwnerID int64 `json:"lead_owner_id"`
	// Name of the lead's owner.
	LeadOwnerName string `json:"lead_owner"`
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
	Created models.Created `json:"created"`
	// Record last update data.
	Updated models.Updated `json:"updated"`
}

// LeadListResponse represents a response to a list request.
type LeadListResponse struct {
	// List of fetched leads.
	Leads []LeadListItem `json:"leads"`
	// Pagination information.
	Pagination models.PaginationResponse `json:"pagination"`
}

// LeadDeleteResponse represents a response to a delete request.
type LeadDeleteResponse struct {
	// ID of the deleted lead.
	LeadID int64 `json:"lead_id"`
	// Indicates whether the deletion was successful.
	IsSuccess bool `json:"is_success"`
}

// LeadConvertResponse represents a response to a lead conversion to a client.
type LeadConvertResponse struct {
	// ID of the lead that was converted to a client.
	ConvertedLeadID int64 `json:"converted_lead_id"`
	// Indicates whether the conversion was successful.
	IsSuccess bool `json:"is_success"`
	// Client created from the lead conversion.
	ClientData models.Client `json:"client_data"`
}
