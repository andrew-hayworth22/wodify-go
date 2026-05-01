package leads

import (
	"net/url"

	"github.com/andrew-hayworth22/wodify-go/models"
	"github.com/andrew-hayworth22/wodify-go/search"
)

// sortableField represents a field that lead lists can be sorted by.
type sortableField string

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
	Sort models.Sort[sortableField]
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
	Sort  models.Sort[sortableField]
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
