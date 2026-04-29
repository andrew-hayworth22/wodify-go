package leads

import (
	"github.com/andrew-hayworth22/wodify-go/types"
)

// Lead represents a Wodify lead.
type Lead struct {
	ID                      int64           `json:"id"`
	FirstName               string          `json:"first_name"`
	LastName                string          `json:"last_name"`
	Email                   string          `json:"email"`
	LeadStatusID            int64           `json:"lead_status_id"`
	LeadStatus              string          `json:"lead_status"`
	LocationID              int64           `json:"location_id"`
	Location                string          `json:"location"`
	GenderID                int64           `json:"gender_id"`
	Gender                  string          `json:"gender"`
	PhoneNumber             string          `json:"phone_number"`
	DateOfBirth             types.Date      `json:"date_of_birth"`
	StreetAddress1          string          `json:"street_address1"`
	StreetAddress2          string          `json:"street_address2"`
	City                    string          `json:"city"`
	StateID                 int             `json:"state_id"`
	State                   string          `json:"state"`
	Province                string          `json:"province"`
	ZipCode                 string          `json:"zip_code"`
	CountryID               int             `json:"country_id"`
	Country                 string          `json:"country"`
	Tags                    []string        `json:"tags"`
	Notes                   string          `json:"notes"`
	LeadGroup               LeadGroup       `json:"lead_group"`
	LastContactDateTime     types.DateTime  `json:"last_contact_date_time"`
	IsConvertedToClient     bool            `json:"is_converted_to_client"`
	StatusHistory           []StatusHistory `json:"status_history"`
	EmergencyContactName    string          `json:"emergency_contact_name"`
	EmergencyContactPhone   string          `json:"emergency_contact_phone"`
	LeadSourceID            int64           `json:"lead_source_id"`
	LeadSource              string          `json:"lead_source"`
	ReferredByFromWeb       string          `json:"referred_by_from_web"`
	ReferredByFromUserId    int64           `json:"referred_by_from_user_id"`
	ReferredByFromUserName  string          `json:"referred_by_from_user_name"`
	IsEmailSubscribed       bool            `json:"is_email_subscribed"`
	IsSMSSubscribed         bool            `json:"is_sms_subscribed"`
	LocationTimezoneID      int64           `json:"location_timezone_id"`
	LocationTimezone        string          `json:"location_timezone"`
	CreatedFromSource       string          `json:"created_from_source"`
	ProfilePhotoURL         string          `json:"profile_photo_url"`
	LeadOwnerID             int64           `json:"lead_owner_id"`
	TotalClassSignIns       int             `json:"total_class_sign_ins"`
	TotalBookingSignIns     int             `json:"total_booking_sign_ins"`
	LastClassSignIn         types.DateTime  `json:"last_class_sign_in"`
	LastBookingSignIn       types.DateTime  `json:"last_booking_sign_in"`
	DaysSinceLastAttendance int             `json:"days_since_last_attendance"`
	IsActive                bool            `json:"is_active"`
	NextClassReservation    types.DateTime  `json:"next_class_reservation"`
	NextAppointmentBooking  types.DateTime  `json:"next_appointment_booking"`
	Created                 types.Updated   `json:"created"`
	Updated                 types.Updated   `json:"updated"`
}

// LeadGroup represents a group of leads.
type LeadGroup struct {
	GroupID                int64              `json:"group_id"`
	GroupRoleID            int64              `json:"group_role_id"`
	GroupRole              string             `json:"group_role"`
	OtherGroupParticipants []GroupParticipant `json:"other_group_participants"`
}

// GroupParticipant represents a participant in a lead group.
type GroupParticipant struct {
	GroupParticipantLeadID int64  `json:"group_participant_lead_id"`
	GroupParticipantName   string `json:"group_participant_name"`
	GroupRoleID            int64  `json:"group_role_id"`
	GroupRole              string `json:"group_role"`
}

// StatusHistory represents the history of a lead's status changes.
type StatusHistory struct {
	FromStatusID         int64          `json:"from_status_id"`
	FromStatus           string         `json:"from_status"`
	ToStatusID           int64          `json:"to_status_id"`
	ToStatus             string         `json:"to_status"`
	StatusChangeDateTime types.DateTime `json:"status_change_date_time"`
}
