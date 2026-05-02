package leads

import (
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/internal/search"
	"github.com/andrew-hayworth22/wodify-go/internal/sort"
	"github.com/andrew-hayworth22/wodify-go/models"
)

// sortableField represents a field that lead lists can be sorted by.
type sortableField string

// Sort represents a lead sort order.
type Sort = sort.Sort[sortableField]

const (
	SortByFirstName   sortableField = "first_name"
	SortByLastName    sortableField = "last_name"
	SortByEmail       sortableField = "email"
	SortByStatus      sortableField = "status"
	SortByLocation    sortableField = "location"
	SortByGender      sortableField = "gender"
	SortByPhone       sortableField = "phone"
	SortByDateOfBirth sortableField = "date_of_birth"
	SortByCreatedAt   sortableField = "created_at"
)

// NewSort creates a new lead sort.
func NewSort(field sortableField, isDescending bool) Sort {
	return sort.NewSort(field, isDescending)
}

// filterableField represents a field that lead lists can be filtered by.
type filterableField string

// Query represents a lead search query.
type Query = search.Builder[filterableField]

const (
	FilterByFirstName filterableField = "first_name"
	FilterByLastName  filterableField = "last_name"
	FilterByStatus    filterableField = "status"
)

// NewQuery creates a new lead search query builder.
func NewQuery() *Query {
	return search.New[filterableField]()
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
	ReferredByUserId int64 `json:"referred_by_user_id"`
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
	Sort sort.Sort[sortableField]
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
	Sort  sort.Sort[sortableField]
	Query *Query
}

// ToQuery converts the request to URL query string parameters.
func (r SearchRequest) ToQuery() url.Values {
	q := r.Page.ToQuery()
	if r.Sort.Field != "" {
		q.Set("sort", r.Sort.String())
	}
	q.Set("q", r.Query.Encode())
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
	GenderID models.Gender `json:"gender_id"`
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
