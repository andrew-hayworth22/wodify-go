package leads

import "github.com/andrew-hayworth22/wodify-go/types"

// CreateLeadRequest represents a request to create a new lead.
type CreateLeadRequest struct {
	FirstName             string       `json:"first_name"`
	LastName              string       `json:"last_name"`
	Email                 string       `json:"email"`
	LeadStatusID          int64        `json:"lead_status_id"`
	LocationID            int64        `json:"location_id"`
	Gender                types.Gender `json:"gender_id"`
	PhoneNumber           string       `json:"phone_number"`
	DateOfBirth           types.Date   `json:"date_of_birth"`
	StreetAddress1        string       `json:"street_address1"`
	StreetAddress2        string       `json:"street_address2"`
	City                  string       `json:"city"`
	StateID               int          `json:"state_id"`
	Province              string       `json:"province"`
	ZipCode               string       `json:"zipcode"`
	CountryID             int          `json:"country_id"`
	Tags                  []string     `json:"tags"`
	Notes                 string       `json:"notes"`
	EmergencyContactName  string       `json:"emergency_contact_name"`
	EmergencyContactPhone string       `json:"emergency_contact_phone"`
	LeadSourceID          int64        `json:"lead_source_id"`
	ReferredByFromWeb     string       `json:"referred_by_from_web"`
	ReferredByUserId      int64        `json:"referred_by_user_id"`
	IsEmailSubscribed     bool         `json:"is_email_subscribed"`
	IsSMSSubscribed       bool         `json:"is_sms_subscribed"`
	LeadOwnerID           int64        `json:"lead_owner_id"`
}
