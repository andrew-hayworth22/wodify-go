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

// Get fetches a single client by ID.
func (c *Client) Get(ctx context.Context, id int64) (*models.Client, error) {
	var out models.Client
	err := c.hc.Do(ctx, http.MethodGet, fmt.Sprintf("/clients/%d", id), nil, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// List fetches a list of clients.
func (c *Client) List(ctx context.Context, req ClientListRequest) (*ClientListResponse, error) {
	var out ClientListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/clients", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Search fetches a list of clients matching the query criteria
func (c *Client) Search(ctx context.Context, req ClientSearchRequest) (*ClientListResponse, error) {
	if req.Query != nil {
		if err := req.Query.Err(); err != nil {
			return nil, err
		}
	}
	var out ClientListResponse
	err := c.hc.Do(ctx, http.MethodGet, "/clients/search", req.ToQuery(), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Create creates a new client
func (c *Client) Create(ctx context.Context, req ClientCreateRequest) (*models.Client, error) {
	var out models.Client
	err := c.hc.Do(ctx, http.MethodPost, "/clients", nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Deactivate deactivates a client by their ID.
func (c *Client) Deactivate(ctx context.Context, id int64) (*ClientActionResponse, error) {
	var out ClientActionResponse
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/clients/%d/deactivate", id), nil, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Reactivate reactivates a deactivated client by their ID.
func (c *Client) Reactivate(ctx context.Context, id int64) (*ClientActionResponse, error) {
	var out ClientActionResponse
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/clients/%d/reactivate", id), nil, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Suspend suspends a client by their ID.
func (c *Client) Suspend(ctx context.Context, id int64) (*ClientActionResponse, error) {
	var out ClientActionResponse
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/clients/%d/suspend", id), nil, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Reinstate reinstates a suspended client by their ID.
func (c *Client) Reinstate(ctx context.Context, id int64) (*ClientActionResponse, error) {
	var out ClientActionResponse
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/clients/%d/reinstate", id), nil, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Update updates a client by ID.
func (c *Client) Update(ctx context.Context, id int64, req ClientUpdateRequest) (*models.Client, error) {
	var out models.Client
	err := c.hc.Do(ctx, http.MethodPut, fmt.Sprintf("/clients/%d", id), nil, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *Client) GenerateRegisterLink(ctx context.Context, id int64) (*GenerateLinkResponse, error) {
	var out GenerateLinkResponse
	err := c.hc.Do(ctx, http.MethodPost, fmt.Sprintf("/clients/%d/register-link", id), nil, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

///////////////////////////////////////////////////////////////////////
// Request Types
///////////////////////////////////////////////////////////////////////

// ClientField represents a field that client lists can be sorted/filtered by.
type ClientField string

const (
	ClientFieldID                         ClientField = "id"
	ClientFieldFirstName                  ClientField = "first_name"
	ClientFieldLastName                   ClientField = "last_name"
	ClientFieldEmail                      ClientField = "email"
	ClientFieldPhoneNumber                ClientField = "phone_number"
	ClientFieldStatusID                   ClientField = "client_status_id"
	ClientFieldStatusName                 ClientField = "client_status"
	ClientFieldLocationID                 ClientField = "location_id"
	ClientFieldLocationName               ClientField = "location"
	ClientFieldDefaultProgramID           ClientField = "default_program_id"
	ClientFieldDefaultProgramName         ClientField = "default_program"
	ClientFieldDateOfBirth                ClientField = "date_of_birth"
	ClientFieldGenderID                   ClientField = "gender_id"
	ClientFieldGenderName                 ClientField = "gender"
	ClientFieldStreetAddress1             ClientField = "street_address_1"
	ClientFieldStreetAddress2             ClientField = "street_address_2"
	ClientFieldCity                       ClientField = "city"
	ClientFieldStateID                    ClientField = "state_id"
	ClientFieldStateName                  ClientField = "state"
	ClientFieldProvince                   ClientField = "province"
	ClientFieldZipCode                    ClientField = "zipcode"
	ClientFieldCountryID                  ClientField = "country_id"
	ClientFieldCountryName                ClientField = "country"
	ClientFieldTimezoneID                 ClientField = "timezone_id"
	ClientFieldTimezoneName               ClientField = "timezone"
	ClientFieldHeight1Measurement         ClientField = "height_measurement_1"
	ClientFieldHeight1SystemOfMeasure     ClientField = "height1_system_of_measure"
	ClientFieldHeight2Measurement         ClientField = "height_measurement_2"
	ClientFieldHeight2SystemOfMeasure     ClientField = "height2_system_of_measure"
	ClientFieldWeight                     ClientField = "weight"
	ClientFieldWeightSystemOfMeasure      ClientField = "weight_system_of_measure"
	ClientFieldIsEmailSubscribed          ClientField = "is_email_subscribed"
	ClientFieldIsSMSSubscribed            ClientField = "is_sms_subscribed"
	ClientFieldTags                       ClientField = "tags"
	ClientFieldEmergencyContactName       ClientField = "emergency_contact_name"
	ClientFieldEmergencyContactPhone      ClientField = "emergency_contact_phone"
	ClientFieldLastAttendance             ClientField = "last_attendance"
	ClientFieldDaysSinceLastAttendance    ClientField = "days_since_last_attendance"
	ClientFieldMemberSince                ClientField = "member_since"
	ClientFieldLastContacted              ClientField = "last_contacted"
	ClientFieldIsAtRisk                   ClientField = "is_at_risk"
	ClientFieldRetainSnoozeUntilDate      ClientField = "retain_snooze_until_date"
	ClientFieldLeadSourceID               ClientField = "lead_source_id"
	ClientFieldLeadSourceName             ClientField = "lead_source"
	ClientFieldReferringUserID            ClientField = "referring_user_id"
	ClientFieldReferringUserName          ClientField = "referring_user"
	ClientFieldIsConvertedFromLead        ClientField = "is_converted_from_lead"
	ClientFieldOwnerID                    ClientField = "client_owner_id"
	ClientFieldOwnerName                  ClientField = "client_owner"
	ClientFieldTotalClassSignIns          ClientField = "total_class_sign_ins"
	ClientFieldTotalBookingSignIns        ClientField = "total_booking_sign_ins"
	ClientFieldLastClassSignIn            ClientField = "last_class_sign_in"
	ClientFieldLastBookingSignIn          ClientField = "last_booking_sign_in"
	ClientFieldCurrentWeekstreak          ClientField = "current_weekstreak"
	ClientFieldHighestWeekstreak          ClientField = "highest_weekstreak"
	ClientFieldCurrentWeekstreakUpdatedOn ClientField = "current_weekstreak_updatedon"
	ClientFieldCoachTitle                 ClientField = "coach_title"
	ClientFieldCoachBio                   ClientField = "coach_bio"
	ClientFieldIsDisplayCoachOnMobileApp  ClientField = "is_display_coach_on_mobile_app"
	ClientFieldCoachLink1Icon             ClientField = "coach_link_1_icon"
	ClientFieldCoachLink1URL              ClientField = "coach_link_1_url"
	ClientFieldCoachLink2Icon             ClientField = "coach_link_2_icon"
	ClientFieldCoachLink2URL              ClientField = "coach_link_2_url"
	ClientFieldCoachLink3Icon             ClientField = "coach_link_3_icon"
	ClientFieldCoachLink3URL              ClientField = "coach_link_3_url"
	ClientFieldCoachLink4Icon             ClientField = "coach_link_4_icon"
	ClientFieldCoachLink4URL              ClientField = "coach_link_4_url"
	ClientFieldCoachLink5Icon             ClientField = "coach_link_5_icon"
	ClientFieldCoachLink5URL              ClientField = "coach_link_5_url"
	ClientFieldNextClassReservation       ClientField = "next_class_reservation"
	ClientFieldNextAppointmentBooking     ClientField = "next_appointment_booking"
)

// ClientListRequest represents a request to list clients.
type ClientListRequest = request.ListRequest[ClientField]

// NewClientListRequest creates a new ClientListRequest with the given pagination and sort.
func NewClientListRequest(pagination request.PaginationRequest, sort sort.Sort[ClientField]) ClientListRequest {
	return ClientListRequest{
		Page: pagination,
		Sort: sort,
	}
}

// ClientSearchRequest represents a request to search clients.
type ClientSearchRequest = request.SearchRequest[ClientField]

// NewClientSearchRequest creates a new ClientSearchRequest with the given pagination, sort, and query.
func NewClientSearchRequest(pagination request.PaginationRequest, sort sort.Sort[ClientField], query *query.Builder[ClientField]) ClientSearchRequest {
	return ClientSearchRequest{
		Page:  pagination,
		Sort:  sort,
		Query: query,
	}
}

// NewClientQuery creates a new client query query builder.
func NewClientQuery() *query.Builder[ClientField] {
	return query.New[ClientField]()
}

// ClientCreateRequest represents a request to create a new client.
type ClientCreateRequest struct {
	// Client's first name.
	FirstName string `json:"first_name"`
	// Client's last name.
	LastName string `json:"last_name"`
	// Client's email address.
	Email string `json:"email"`
	// Client's phone number.
	PhoneNumber string `json:"phone_number"`
	// Client's billing credit card email.
	BillingCCEmail string `json:"billing_cc_email"`
	// Client's status identifier.
	ClientStatusID int64 `json:"client_status_id"`
	// Client's default location identifier.
	LocationID int64 `json:"location_id"`
	// Client's date of birth.
	DateOfBirth models.Date `json:"date_of_birth"`
	// Client's gender ID.
	GenderID int64 `json:"gender_id"`
	// Client's street address (line 1).
	StreetAddress1 string `json:"street_address_1"`
	// Client's street address (line 2).
	StreetAddress2 string `json:"street_address_2"`
	// Client's city.
	City string `json:"city"`
	// Client's state.
	StateID int64 `json:"state_id"`
	// Client's province, if applicable.
	Province string `json:"province"`
	// Client's ZIP code.
	ZipCode string `json:"zipcode"`
	// Client's country ID.
	CountryID int64 `json:"country_id"`
	// Client's timezone identifier.
	TimezoneID int64 `json:"timezone_id"`
	// Client's first height measurement (feet/meters based on UOM).
	Height1Measurement int `json:"height_measurement_1"`
	// Client's second height measurement (inches/centimeters based on UOM).
	Height2Measurement int `json:"height_measurement_2"`
	// Client's weight measurement (pounds/kilograms based on UOM).
	Weight float64 `json:"weight"`
	// IsEmailSubscribed indicates whether the client has subscribed to email notifications.
	IsEmailSubscribed bool `json:"is_email_subscribed"`
	// Tags associated with the client.
	Tags []string `json:"tags"`
	// Notes associated with the client.
	Notes string `json:"notes"`
	// Emergency contact name.
	EmergencyContactName string `json:"emergency_contact_name"`
	// Emergency contact phone number.
	EmergencyContactPhone string `json:"emergency_contact_phone"`
	// Identifier of the lead source that the client was converted from.
	LeadSourceID int64 `json:"lead_source_id"`
	// Identifier of the user that referred the client.
	ReferringUserID int64 `json:"referring_user_id"`
	// IsSMSSubscribed indicates whether the client has subscribed to SMS notifications.
	IsSMSSubscribed bool `json:"is_sms_subscribed"`
	// Client's owner identifier.
	ClientOwnerID int64 `json:"client_owner_id"`
}

// ClientUpdateRequest represents a request to update a client.
type ClientUpdateRequest struct {
	// Client's first name.
	FirstName string `json:"first_name"`
	// Client's last name.
	LastName string `json:"last_name"`
	// Client's email address.
	Email string `json:"email"`
	// Client's phone number.
	PhoneNumber string `json:"phone_number"`
	// Client's billing credit card email.
	BillingCCEmail string `json:"billing_cc_email"`
	// Client's status identifier.
	ClientStatusID int64 `json:"client_status_id"`
	// Client's default location identifier.
	LocationID int64 `json:"location_id"`
	// Client's default program identifier.
	DefaultProgramID int64 `json:"default_program_id"`
	// Client's date of birth.
	DateOfBirth models.Date `json:"date_of_birth"`
	// Client's gender ID.
	GenderID int `json:"gender_id"`
	// Client's street address (line 1).
	StreetAddress1 string `json:"street_address_1"`
	// Client's street address (line 2).
	StreetAddress2 string `json:"street_address_2"`
	// Client's city.
	City string `json:"city"`
	// Client's state.
	StateID int64 `json:"state_id"`
	// Client's province, if applicable.
	Province string `json:"province"`
	// Client's ZIP code.
	ZipCode string `json:"zipcode"`
	// Client's country ID.
	CountryID int `json:"country_id"`
	// Client's timezone identifier.
	TimezoneID int64 `json:"timezone_id"`
	// Client's first height measurement (feet/meters based on UOM).
	Height1Measurement int `json:"height_measurement_1"`
	// Client's second height measurement (inches/centimeters based on UOM).
	Height2Measurement int `json:"height_measurement_2"`
	// Client's weight measurement (pounds/kilograms based on UOM).
	Weight float64 `json:"weight"`
	// IsEmailSubscribed indicates whether the client has subscribed to email notifications.
	IsEmailSubscribed bool `json:"is_email_subscribed"`
	// IsSMSSubscribed indicates whether the client has subscribed to SMS notifications.
	IsSMSSubscribed bool `json:"is_sms_subscribed"`
	// Notes associated with the client.
	Notes string `json:"notes"`
	// IsOverwriteNotes indicates whether to overwrite existing notes.
	IsOverwriteNotes bool `json:"is_overwrite_notes"`
	// Emergency contact name.
	EmergencyContactName string `json:"emergency_contact_name"`
	// Emergency contact phone number.
	EmergencyContactPhone string `json:"emergency_contact_phone"`
	// Identifier of the lead source that the client was converted from.
	LeadSourceID int64 `json:"lead_source_id"`
	// Identifier of the user that referred the client.
	ReferringUserID int64 `json:"referring_user_id"`
	// Client's owner identifier.
	ClientOwnerID int64 `json:"client_owner_id"`
	// Coach's title.
	CoachTitle string `json:"coach_title"`
	// Coach's bio HTML.
	CoachBio string `json:"coach_bio"`
	// Coach's first link icon.
	CoachLink1Icon string `json:"coach_link_1_icon"`
	// Coach's first link URL.
	CoachLink1URL string `json:"coach_link_1_url"`
	// Coach's second link icon.
	CoachLink2Icon string `json:"coach_link_2_icon"`
	// Coach's second link URL.
	CoachLink2URL string `json:"coach_link_2_url"`
	// Coach's third link icon.
	CoachLink3Icon string `json:"coach_link_3_icon"`
	// Coach's third link URL.
	CoachLink3URL string `json:"coach_link_3_url"`
	// Coach's fourth link icon.
	CoachLink4Icon string `json:"coach_link_4_icon"`
	// Coach's fourth link URL.
	CoachLink4URL string `json:"coach_link_4_url"`
	// Coach's fifth link icon.
	CoachLink5Icon string `json:"coach_link_5_icon"`
	// Coach's fifth link URL.
	CoachLink5URL string `json:"coach_link_5_url"`
}

// ClientUpdateRequestFrom converts a Client to a ClientUpdateRequest
func ClientUpdateRequestFrom(c *models.Client) ClientUpdateRequest {
	return ClientUpdateRequest{
		FirstName:             c.FirstName,
		LastName:              c.LastName,
		Email:                 c.Email,
		PhoneNumber:           c.PhoneNumber,
		BillingCCEmail:        c.BillingCCEmail,
		ClientStatusID:        c.ClientStatusID,
		LocationID:            c.LocationID,
		DefaultProgramID:      c.DefaultProgramID,
		DateOfBirth:           c.DateOfBirth,
		GenderID:              c.GenderID,
		StreetAddress1:        c.StreetAddress1,
		StreetAddress2:        c.StreetAddress2,
		City:                  c.City,
		StateID:               c.StateID,
		Province:              c.Province,
		ZipCode:               c.ZipCode,
		CountryID:             c.CountryID,
		TimezoneID:            c.TimezoneID,
		Height1Measurement:    c.Height1Measurement,
		Height2Measurement:    c.Height2Measurement,
		Weight:                c.Weight,
		IsEmailSubscribed:     c.IsEmailSubscribed,
		IsSMSSubscribed:       c.IsSMSSubscribed,
		Notes:                 c.Notes,
		IsOverwriteNotes:      true,
		EmergencyContactName:  c.EmergencyContactName,
		EmergencyContactPhone: c.EmergencyContactPhone,
		LeadSourceID:          c.LeadSourceID,
		ReferringUserID:       c.ReferringUserID,
		ClientOwnerID:         c.ClientOwnerID,
		CoachTitle:            c.CoachTitle,
		CoachBio:              c.CoachBio,
		CoachLink1Icon:        c.CoachLink1Icon,
		CoachLink1URL:         c.CoachLink1URL,
		CoachLink2Icon:        c.CoachLink2Icon,
		CoachLink2URL:         c.CoachLink2URL,
		CoachLink3Icon:        c.CoachLink3Icon,
		CoachLink3URL:         c.CoachLink3URL,
		CoachLink4Icon:        c.CoachLink4Icon,
		CoachLink4URL:         c.CoachLink4URL,
		CoachLink5Icon:        c.CoachLink5Icon,
		CoachLink5URL:         c.CoachLink5URL,
	}
}

///////////////////////////////////////////////////////////////////////
// Response Types
///////////////////////////////////////////////////////////////////////

// ClientListItem represents a client in a list.
type ClientListItem struct {
	// Client's ID.
	ID int64 `json:"id"`
	// Client's first name.
	FirstName string `json:"first_name"`
	// Client's last name.
	LastName string `json:"last_name"`
	// Client's email address.
	Email string `json:"email"`
	// Client's phone number.
	PhoneNumber string `json:"phone_number"`
	// Client's status identifier.
	ClientStatusID int64 `json:"client_status_id"`
	// Client's status label.
	ClientStatus string `json:"client_status"`
	// Client's default location identifier.
	LocationID int64 `json:"location_id"`
	// Client's default location name.
	Location string `json:"location"`
	// Client's default program identifier.
	DefaultProgramID int64 `json:"default_program_id"`
	// Default program name.
	DefaultProgram string `json:"default_program"`
	// Client's date of birth.
	DateOfBirth string `json:"date_of_birth"`
	// Client's gender ID.
	GenderID int64 `json:"gender_id"`
	// Client's gender name.
	GenderName string `json:"gender"`
	// Client's street address (line 1).
	StreetAddress1 string `json:"street_address_1"`
	// Client's street address (line 2).
	StreetAddress2 string `json:"street_address_2"`
	// Client's city.
	City string `json:"city"`
	// Client's state.
	StateID int64 `json:"state_id"`
	// Client's state name.
	StateName string `json:"state"`
	// Client's province, if applicable.
	Province string `json:"province"`
	// Client's ZIP code.
	ZipCode string `json:"zipcode"`
	// Client's country ID.
	CountryID int64 `json:"country_id"`
	// Client's country name.
	CountryName string `json:"country"`
	// Client's timezone identifier.
	TimezoneID int64 `json:"timezone_id"`
	// Client's timezone name.
	TimezoneName string `json:"timezone"`
	// Client's first height measurement (feet/meters based on UOM).
	Height1Measurement int `json:"height_measurement_1"`
	// Client's first height measurement unit.
	Height1SystemOfMeasure string `json:"height1_system_of_measure"`
	// Client's second height measurement (inches/centimeters based on UOM).
	Height2Measurement int `json:"height_measurement_2"`
	// Client's second height measurement unit.
	Height2SystemOfMeasure string `json:"height2_system_of_measure"`
	// Client's weight measurement (pounds/kilograms based on UOM).
	Weight float64 `json:"weight"`
	// Client's weight measurement unit.
	WeightSystemOfMeasure string `json:"weight_system_of_measure"`
	// IsEmailSubscribed indicates whether the client has subscribed to email notifications.
	IsEmailSubscribed bool `json:"is_email_subscribed"`
	// IsSMSSubscribed indicates whether the client has subscribed to SMS notifications.
	IsSMSSubscribed bool `json:"is_sms_subscribed"`
	// Tags associated with the client.
	Tags []string `json:"tags"`
	// Emergency contact name.
	EmergencyContactName string `json:"emergency_contact_name"`
	// Emergency contact phone number.
	EmergencyContactPhone string `json:"emergency_contact_phone"`
	// Last time the client attended a class or appointment booking.
	LastAttendance models.DateTime `json:"last_attendance"`
	// Number of days since the client last attended a class or appointment booking.
	DaysSinceLastAttendance int `json:"days_since_last_attendance"`
	// Date the client joined the business.
	MemberSince string `json:"member_since"`
	// Last time the client was contacted.
	LastContacted string `json:"last_contacted"`
	// Indicates whether the client is at risk.
	IsAtRisk bool `json:"is_at_risk"`
	// Date that the client will be snoozed until
	RetainSnoozeUntilDate string `json:"retain_snooze_until_date"`
	// Identifier of the lead source that the client was converted from.
	LeadSourceID int64 `json:"lead_source_id"`
	// Name of the lead source that the client was converted from.
	LeadSourceName string `json:"lead_source"`
	// Identifier of the user that referred the client.
	ReferringUserID int64 `json:"referring_user_id"`
	// Name of the user that referred the client.
	ReferringUser string `json:"referring_user"`
	// Indicates whether the client was converted from a lead.
	IsConvertedFromLead bool `json:"is_converted_from_lead"`
	// Client's owner identifier.
	ClientOwnerID int64 `json:"client_owner_id"`
	// Client's owner name.
	ClientOwnerName string `json:"client_owner"`
	// Number of classes the client has signed in to.
	TotalClassSignIns int `json:"total_class_sign_ins"`
	// Number of bookings the client has signed in to.
	TotalBookingSignIns int `json:"total_booking_sign_ins"`
	// Last time the client signed in to a class.
	LastClassSignIn models.DateTime `json:"last_class_sign_in"`
	// Last time the client signed in to an appointment booking.
	LastBookingSignIn models.DateTime `json:"last_booking_sign_in"`
	// Current streak of weeks the client has attended a class or appointment booking.
	CurrentWeekstreak int `json:"current_weekstreak"`
	// Highest streak of weeks the client has attended a class or appointment booking.
	HighestWeekstreak int `json:"highest_weekstreak"`
	// Last time the client's streak was updated.
	CurrentWeekstreakUpdatedon models.DateTime `json:"current_weekstreak_updatedon"`
	// Coach's title.
	CoachTitle string `json:"coach_title"`
	// Coach's bio HTML.
	CoachBio string `json:"coach_bio"`
	// Indicates whether the coach should be displayed on the mobile app.
	IsDisplayCoachOnMobileApp bool `json:"is_display_coach_on_mobile_app"`
	// Coach's first link icon.
	CoachLink1Icon string `json:"coach_link_1_icon"`
	// Coach's first link URL.
	CoachLink1Url string `json:"coach_link_1_url"`
	// Coach's second link icon.
	CoachLink2Icon string `json:"coach_link_2_icon"`
	// Coach's second link URL.
	CoachLink2URL string `json:"coach_link_2_url"`
	// Coach's third link icon.
	CoachLink3Icon string `json:"coach_link_3_icon"`
	// Coach's third link URL.
	CoachLink3URL string `json:"coach_link_3_url"`
	// Coach's fourth link icon.
	CoachLink4Icon string `json:"coach_link_4_icon"`
	// Coach's fourth link URL.
	CoachLink4URL string `json:"coach_link_4_url"`
	// Coach's fifth link icon.
	CoachLink5Icon string `json:"coach_link_5_icon"`
	// Coach's fifth link URL.
	CoachLink5URL string `json:"coach_link_5_url"`
	// Time of the client's next class reservation.
	NextClassReservation models.DateTime `json:"next_class_reservation"`
	// Time of the client's next appointment booking.
	NextAppointmentBooking models.DateTime `json:"next_appointment_booking"`
	// Record creation data.
	Created models.Created `json:"created_on"`
	// Record last update data.
	Updated models.Updated `json:"updated"`
}

// ClientListResponse represents a response to a client list request.
type ClientListResponse struct {
	Clients    []ClientListItem          `json:"clients"`
	Pagination models.PaginationResponse `json:"pagination"`
}

// ClientActionResponse represents a response to the following actions:
//   - Client.Deactivate
//   - Client.Reactivate
//   - Client.Suspend
//   - Client.Reinstate
type ClientActionResponse struct {
	// Indicates whether the client was successfully deactivated or reactivated.
	IsSuccess bool `json:"is_success"`
	// Client data after the deactivation or reactivation.
	Client models.Client `json:"ClientResponse"`
}

// GenerateLinkResponse represents a client registration link registration response.
type GenerateLinkResponse struct {
	// Registration link.
	Link string `json:"register_link"`
}
