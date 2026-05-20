package models

// Lead represents a Wodify lead.
type Lead struct {
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
	Gender Gender `json:"gender_id"`
	// Lead's phone number.
	PhoneNumber string `json:"phone_number"`
	// Lead's date of birth.
	DateOfBirth Date `json:"date_of_birth"`
	// Lead's street address (line 1).
	StreetAddress1 string `json:"street_address_1"`
	// Lead's street address (line 2).
	StreetAddress2 string `json:"street_address_2"`
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
	// Lead's notes.
	Notes string `json:"notes"`
	// Lead's group data.
	LeadGroup LeadGroup `json:"lead_group"`
	// Last time the lead was contacted.
	LastContactDateTime DateTime `json:"last_contact_date_time"`
	// Indicates whether the lead has been converted to a client.
	IsConvertedToClient bool `json:"is_converted_to_client"`
	// Lead's status history
	StatusHistory []LeadStatusHistory `json:"status_history"`
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
	LastClassSignIn DateTime `json:"last_class_sign_in"`
	// Last time the Lead signed in to an appointment booking.
	LastBookingSignIn DateTime `json:"last_booking_sign_in"`
	// Number of days since the Lead last attended a class or appointment booking.
	DaysSinceLastAttendance int `json:"days_since_last_attendance"`
	// Indicates whether the Lead is active.
	IsActive bool `json:"is_active"`
	// Next reservation date and time for the Lead.
	NextClassReservation DateTime `json:"next_class_reservation"`
	// Next appointment booking date and time for the Lead.
	NextAppointmentBooking DateTime `json:"next_appointment_booking"`
	// Record creation data.
	Created Updated `json:"created"`
	// Record last update data.
	Updated Updated `json:"updated"`
}

// LeadGroup represents a group of leads.
type LeadGroup struct {
	// ID of the Lead group.
	ID int64 `json:"group_id"`
	// ID of the Lead group's role.
	RoleID int64 `json:"group_role_id"`
	// Name of the Lead group's role.
	Role string `json:"group_role"`
	// List of the other members of the Lead's group.
	OtherGroupParticipants []LeadGroupParticipant `json:"other_group_participants"`
}

// LeadGroupParticipant represents a participant in a lead group.
type LeadGroupParticipant struct {
	// ID of the group member.
	GroupParticipantLeadID int64 `json:"group_participant_lead_id"`
	// Name of the group member.
	GroupParticipantName string `json:"group_participant_name"`
	// ID of the group member's role.'
	GroupRoleID int64 `json:"group_role_id"`
	// Name of the group member's role.'
	GroupRole string `json:"group_role"`
}

// LeadStatusHistory represents the history of a lead's status changes.
type LeadStatusHistory struct {
	// ID of the lead status before change.
	FromStatusID int64 `json:"from_status_id"`
	// Name of the lead status before change.
	FromStatus string `json:"from_status"`
	// ID of the lead status after change.
	ToStatusID int64 `json:"to_status_id"`
	// Name of the lead status after change.
	ToStatus string `json:"to_status"`
	// Date and time of the lead status change.
	StatusChangeDateTime DateTime `json:"status_change_datetime"`
}

// LeadStatus represents a status for a Lead
type LeadStatus struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

// LeadSource represents a source of a Lead
type LeadSource struct {
	ID     int64  `json:"id"`
	Source string `json:"source"`
}

// LeadBooking represents a booking for a Lead
type LeadBooking struct {
	Booking
	// ID of the lead booked
	LeadID int64 `json:"lead_id"`
	// Name of the lead booked
	LeadName string `json:"lead"`
}
