package models

// Booking represents an appointment booking.
type Booking struct {
	// ID of the booking.
	BookingID int64 `json:"booking_id"`
	// ID of the appointment booked
	AppointmentID int64 `json:"appointment_id"`
	// UTC Timestamp of the appointment start
	UTCAppointmentStartDateTime DateTime `json:"utc_appointment_start_datetime"`
	// UTC Timestamp of the appointment end
	UTCAppointmentEndDateTime DateTime `json:"utc_appointment_end_datetime"`
	// Local Timestamp of the appointment start
	LocalAppointmentStartDateTime DateTime `json:"local_appointment_start_datetime"`
	// Local Timestamp of the appointment end
	LocalAppointmentEndDateTime DateTime `json:"local_appointment_end_datetime"`
	// ID of service provided by the appointment
	ServiceID string `json:"service_id"`
	// Name of service provided by the appointment
	ServiceName string `json:"service"`
	// ID of the location where the appointment is held
	LocationID int64 `json:"location_id"`
	// Name of the location where the appointment is held
	LocationName string `json:"location"`
	// ID of the provider who is providing the appointment
	ProviderID string `json:"provider_id"`
	// Name of the provider who is providing the appointment
	ProviderName string `json:"provider"`
	// ID of the status of the booking
	BookingStatusID int64 `json:"booking_status_id"`
	// Name of the status of the booking
	BookingStatus string `json:"booking_status"`
	// ID of the membership associated with the booking
	MembershipID int64 `json:"membership_id"`
	// Name of the membership associated with the booking
	MembershipName string `json:"membership"`
	// Indicates whether the booking is a free trial
	IsFreeTrial bool `json:"is_free_trial"`
	// Indicates whether the booking is a late cancellation
	IsLateCancellation bool `json:"is_late_cancellation"`
	// Information surrounding the cancellation of the booking
	Canceled Canceled `json:"cancelled"`
	// Record creation data
	Created Created `json:"created"`
	// Record last update data
	Updated Updated `json:"updated"`
}

// Canceled represents information surrounding the cancellation of a booking.
type Canceled struct {
	// ID of the user that canceled the booking.
	CanceledById int64 `json:"cancelled_by_id"`
	// Name of the user that canceled the booking.
	CanceledByName string   `json:"cancelled_by"`
	CanceledOn     DateTime `json:"cancelled_on_datetime"`
}
