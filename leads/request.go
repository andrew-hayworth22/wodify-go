package leads

import (
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/internal/search"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

// LeadField represents a field that lead lists can be sorted/filtered by.
type LeadField string

const (
	LeadFieldFirstName   LeadField = "first_name"
	LeadFieldLastName    LeadField = "last_name"
	LeadFieldEmail       LeadField = "email"
	LeadFieldStatus      LeadField = "status"
	LeadFieldLocation    LeadField = "location"
	LeadFieldGender      LeadField = "gender"
	LeadFieldPhone       LeadField = "phone"
	LeadFieldDateOfBirth LeadField = "date_of_birth"
	LeadFieldCreatedAt   LeadField = "created_at"
)

// LeadSort represents a lead sort order.
type LeadSort = sort.Sort[LeadField]

// NewLeadSort creates a new lead sort.
func NewLeadSort(field LeadField, isDescending bool) LeadSort {
	return sort.NewSort(field, isDescending)
}

// LeadQuery represents a lead search query.
type LeadQuery = search.Builder[LeadField]

// NewLeadQuery creates a new lead search query builder.
func NewLeadQuery() *LeadQuery {
	return search.New[LeadField]()
}

// CreateLeadRequest represents a request to create a new lead.
type CreateLeadRequest struct {
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

// ListRequest represents a request to list leads.
type ListRequest struct {
	Page models.PaginationRequest
	Sort LeadSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

// SearchRequest represents a request to search leads.
type SearchRequest struct {
	Page  models.PaginationRequest
	Sort  LeadSort
	Query *LeadQuery
}

// ToQuery converts the request to URL query string parameters.
func (r SearchRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	if r.Query != nil {
		q.Set("q", r.Query.String())
	}
	return q
}

// UpdateLeadRequest represents a request to update a lead.
type UpdateLeadRequest struct {
	// Lead's first name.
	FirstName string `json:"first_name,omitempty"`
	// Lead's last name.
	LastName string `json:"last_name,omitempty"`
	// Lead's email address.
	Email string `json:"email,omitempty"`
	// Lead's status ID.
	LeadStatusID int64 `json:"lead_status_id,omitempty"`
	// Lead's default location ID.
	LocationID int64 `json:"location_id,omitempty"`
	// Lead's gender.
	Gender models.Gender `json:"gender_id,omitempty"`
	// Lead's phone number.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Lead's date of birth.
	DateOfBirth models.Date `json:"date_of_birth,omitempty"`
	// Lead's street address (line 1).
	StreetAddress1 string `json:"street_address1,omitempty"`
	// Lead's street address (line 2).
	StreetAddress2 string `json:"street_address2,omitempty"`
	// Lead's city.
	City string `json:"city,omitempty"`
	// Lead's state.
	StateID int `json:"state_id,omitempty"`
	// Lead's province, if applicable.
	Province string `json:"province,omitempty"`
	// Lead's ZIP code.
	ZipCode string `json:"zipcode,omitempty"`
	// Lead's country ID.
	CountryID int `json:"country_id,omitempty"`
	// Lead's notes.
	Notes string `json:"notes,omitempty"`
	// Lead's Emergency contact name.
	EmergencyContactName string `json:"emergency_contact_name,omitempty"`
	// Lead's Emergency contact phone number.
	EmergencyContactPhone string `json:"emergency_contact_phone,omitempty"`
	// Lead's source ID.
	LeadSourceID int64 `json:"lead_source_id,omitempty"`
	// Who referred the Lead on the Web (free text).
	ReferredByFromWeb string `json:"referred_by_from_web,omitempty"`
	// Unique ID of the user that referred the Lead.
	ReferredByUserId int64 `json:"referred_by_user_id,omitempty"`
	// Indicates whether the Lead has subscribed to email notifications.
	IsEmailSubscribed bool `json:"is_email_subscribed,omitempty"`
	// Indicates whether the Lead has subscribed to SMS notifications.
	IsSMSSubscribed bool `json:"is_sms_subscribed,omitempty"`
	// Unique ID of the Lead's owner.
	LeadOwnerID int64 `json:"lead_owner_id,omitempty"`
}

// UpdateRequestFrom converts a Lead to an UpdateLeadRequest.
func UpdateRequestFrom(l *models.Lead) UpdateLeadRequest {
	return UpdateLeadRequest{
		FirstName:             l.FirstName,
		LastName:              l.LastName,
		Email:                 l.Email,
		LeadStatusID:          l.LeadStatusID,
		LocationID:            l.LocationID,
		Gender:                l.Gender,
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
		ReferredByUserId:      l.ReferredByUserId,
		IsEmailSubscribed:     l.IsEmailSubscribed,
		IsSMSSubscribed:       l.IsSMSSubscribed,
		LeadOwnerID:           l.LeadOwnerID,
	}
}

// ConvertLeadRequest represents a request to convert a lead to a client.
type ConvertLeadRequest struct {
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
	// Gender the client will have after conversion.
	Gender models.Gender `json:"gender_id"`
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

// ConversionRequestFrom converts a Lead to a ConvertLeadRequest.
func ConversionRequestFrom(l *models.Lead) ConvertLeadRequest {
	return ConvertLeadRequest{
		LocationID:     l.LocationID,
		Email:          l.Email,
		FirstName:      l.FirstName,
		LastName:       l.LastName,
		ClientStatusID: 0,
		Gender:         l.Gender,
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

// StatusField represents a field that lead statuses can be sorted/filtered on
type StatusField string

const (
	StatusFieldID   StatusField = "id"
	StatusFieldName StatusField = "status"
)

// StatusSort represents a lead status sort order
type StatusSort = sort.Sort[StatusField]

// NewStatusSort creates a new lead status sort
func NewStatusSort(field StatusField, isDescending bool) StatusSort {
	return sort.NewSort(field, isDescending)
}

// ListStatusesRequest represents a request to list lead statuses
type ListStatusesRequest struct {
	Page models.PaginationRequest
	Sort StatusSort
}

// ToQuery converts the request to URL query string parameters.
func (r ListStatusesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}

// SourceField represents a field that lead sources can be sorted/filtered on
type SourceField string

const (
	SourceFieldID   SourceField = "id"
	SourceFieldName SourceField = "source"
)

// SourceSort represents a lead source sort order
type SourceSort = sort.Sort[SourceField]

// NewSourceSort creates a new lead source sort
func NewSourceSort(field SourceField, isDescending bool) SourceSort {
	return sort.NewSort(field, isDescending)
}

// ListSourcesRequest represents a request to list lead sources
type ListSourcesRequest struct {
	Page models.PaginationRequest
	Sort SourceSort
}

func (r ListSourcesRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	return q
}
