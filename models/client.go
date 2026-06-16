package models

// Client represents a business client
type Client struct {
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
	// Client's billing credit card email.
	BillingCCEmail string `json:"billing_cc_email"`
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
	DateOfBirth Date `json:"date_of_birth"`
	// Client's gender ID.
	GenderID int `json:"gender_id"`
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
	CountryID int `json:"country_id"`
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
	// Notes associated with the client.
	Notes string `json:"notes"`
	// Client's group data.
	ClientGroup ClientGroupMemberInfo `json:"group"`
	// Status history of the client.
	StatusHistory []ClientStatusHistory `json:"status_history"`
	// Emergency contact name.
	EmergencyContactName string `json:"emergency_contact_name"`
	// Emergency contact phone number.
	EmergencyContactPhone string `json:"emergency_contact_phone"`
	// Indicates whether the client has overdue invoices.
	HasOverdueInvoices bool `json:"has_overdue_invoices"`
	// Indicates whether the client has membership.
	HasMembership bool `json:"has_membership"`
	// Indicates whether the client has a payment method.
	HasPaymentMethod bool `json:"has_payment_method"`
	// Indicates whether the client has store credit.
	HasStoreCredit bool `json:"has_store_credit"`
	// Last time the client attended a class or appointment booking.
	LastAttendance DateTime `json:"last_attendance"`
	// Number of days since the client last attended a class or appointment booking.
	DaysSinceLastAttendance int `json:"days_since_last_attended"`
	// Date the client joined the business.
	MemberSince Date `json:"member_since"`
	// Last time the client was contacted.
	LastContacted Date `json:"last_contacted"`
	// Indicates whether the client is at risk.
	IsAtRisk bool `json:"is_at_risk"`
	// Date that the client will be snoozed until
	RetainSnoozeUntilDate Date `json:"retain_snooze_until_date"`
	// Indicates whether the client has an initial email sent.
	IsInitialEmailSent bool `json:"is_initial_email_sent"`
	// Identifier of the lead that the client was converted from.
	LeadId int64 `json:"lead_id"`
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
	// Reason for deactivation of the client.
	MemberDeactivationReason string `json:"member_deactivation_reason"`
	// Client's roles.
	Roles []ClientRole `json:"roles"`
	// Client's programs.
	Programs []ClientProgram `json:"available_programs"`
	// Client's owner identifier.
	ClientOwnerID int64 `json:"client_owner_id"`
	// Client's owner name.
	ClientOwnerName string `json:"client_owner"`
	// Number of classes the client has signed in to.
	TotalClassSignIns int `json:"total_class_sign_ins"`
	// Number of bookings the client has signed in to.
	TotalBookingSignIns int `json:"total_booking_sign_ins"`
	// Last time the client signed in to a class.
	LastClassSignIn DateTime `json:"last_class_sign_in"`
	// Last time the client signed in to an appointment booking.
	LastBookingSignIn DateTime `json:"last_booking_sign_in"`
	// Current streak of weeks the client has attended a class or appointment booking.
	CurrentWeekstreak int `json:"current_weekstreak"`
	// Highest streak of weeks the client has attended a class or appointment booking.
	HighestWeekstreak int `json:"highest_weekstreak"`
	// Last time the client's streak was updated.
	CurrentWeekstreakUpdatedOn DateTime `json:"current_weekstreak_updatedon"`
	// Coach's title.
	CoachTitle string `json:"coach_title"`
	// Coach's bio HTML.
	CoachBio string `json:"coach_bio"`
	// Indicates whether the coach should be displayed on the mobile app.
	IsDisplayCoachOnMobileApp bool `json:"is_display_coach_on_mobile_app"`
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
	// Time of the client's next class reservation.
	NextClassReservation DateTime `json:"next_class_reservation"`
	// Time of the client's next appointment booking.
	NextAppointmentBooking DateTime `json:"next_appointment_booking"`
	// Services that the client can coach.
	CoachServices []ClientService `json:"coach_services"`
	// Programs that the client can coach.
	CoachPrograms []ClientProgram `json:"coach_programs"`
	// Locations that the client can coach at.
	CoachLocations []ClientLocation `json:"coach_locations"`
	// Record creation data.
	Created Created `json:"created"`
	// Record last update data.
	Updated Updated `json:"updated"`
}

// ClientGroup represents a group of clients.
type ClientGroup struct {
	ID           int64                    `json:"group_id"`
	Participants []ClientGroupParticipant `json:"group_participants"`
}

// ClientGroupMemberInfo represents a group of clients.
type ClientGroupMemberInfo struct {
	// ID of the Client group.
	GroupID int64 `json:"group_id"`
	// ID of the Client's role in the group.
	GroupRoleID int64 `json:"group_role_id"`
	// Name of the Client group's role.
	GroupRole string `json:"group_role"`
	// List of the other members of the Client's group.
	OtherGroupParticipants []ClientGroupParticipant `json:"other_group_participants"`
}

// ClientGroupParticipant represents a participant in a lead group.
type ClientGroupParticipant struct {
	// ID of the group member.
	ClientID int64 `json:"group_participant_client_id"`
	// Name of the group member.
	Name string `json:"group_participant_name"`
	// ID of the group member's role.'
	GroupRoleID int64 `json:"group_role_id"`
	// Name of the group member's role.'
	GroupRole string `json:"group_role"`
}

// ClientGroupRole represents a role associated with a client in a group
type ClientGroupRole struct {
	// ID of role
	ID int64 `json:"id"`
	// Name of role
	Name string `json:"group_role"`
}

// ClientStatusHistory represents the history of a lead's status changes.
type ClientStatusHistory struct {
	// ID of the client status before change.
	FromStatusID int64 `json:"from_status_id"`
	// Name of the client status before change.
	FromStatus string `json:"from_status"`
	// ID of the client status after change.
	ToStatusID int64 `json:"to_status_id"`
	// Name of the client status after change.
	ToStatus string `json:"to_status"`
	// Date and time of the client status change.
	StatusChangeDateTime DateTime `json:"status_change_datetime"`
}

// ClientRole represents a role assigned to a client.
type ClientRole struct {
	// ID of the role.
	ID int64 `json:"id"`
	// Name of the role.
	Name string `json:"role"`
}

// ClientProgram represents a program available to a client.
type ClientProgram struct {
	// ID of the program.
	ID int64 `json:"program_id"`
	// Name of the program.
	Name string `json:"program"`
}

// ClientService represents a service available to a client.
type ClientService struct {
	// ID of the service.
	ID int64 `json:"service_id"`
	// Name of the service.
	Name string `json:"service"`
}

// ClientLocation represents a location associated with a client.
type ClientLocation struct {
	// ID of the location.
	ID int64 `json:"location_id"`
	// Name of the location.
	Name string `json:"location"`
}

// ClientStatus represents a status of a client.
type ClientStatus struct {
	// ID of the status
	ID int64 `json:"id"`
	// Name of the status
	Name string `json:"status"`
}

// ClientBooking represents a booking for a client
type ClientBooking struct {
	Booking
	// ID of the client booked
	ClientID int64 `json:"client_id"`
	// Name of the client booked
	ClientName string `json:"client"`
}

// ClientClassSignIn represents a sign-in for a client
type ClientClassSignIn struct {
	ClassSignIn
	// ID of the client signed in
	ClientID int64 `json:"client_id"`
	// Name of the client signed in
	ClientName string `json:"client"`
}

// ClientReservation represents a reservation for a client
type ClientReservation struct {
	Reservation
	// ID of the client signed in
	ClientID int64 `json:"client_id"`
	// Name of the client signed in
	ClientName string `json:"client"`
}
