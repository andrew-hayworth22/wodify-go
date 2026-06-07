package models

// Reservation represents a reservation for a class.
type Reservation struct {
	// ID of the reservation
	ID int64 `json:"reservation_id"`
	// ID of the class reserved for
	ClassId int64 `json:"class_id"`
	// Name of the class reserved for
	ClassName string `json:"class"`
	// UTC Timestamp of the class start
	UTCClassStartDateTime DateTime `json:"utc_class_start_datetime"`
	// UTC Timestamp of the class end
	UTCClassEndDateTime DateTime `json:"utc_class_end_datetime"`
	// Local Timestamp of the class start
	LocalClassStartDateTime DateTime `json:"local_class_start_datetime"`
	// Local Timestamp of the class end
	LocalClassEndDateTime DateTime `json:"local_class_end_datetime"`
	// ID of the program the class is part of
	ProgramID string `json:"program_id"`
	// Name of the program the class is part of
	ProgramName string `json:"program"`
	// ID of the location the class is held at
	LocationID int64 `json:"location_id"`
	// Name of the location the class is held at
	LocationName string `json:"location_name"`
	// Coaches for the class
	Coaches []Coach `json:"coaches"`
	// ID of the reservation status
	ReservationStatusId int64 `json:"reservation_status_id"`
	// Name of the reservation status
	ReservationStatusName string `json:"reservation_status"`
	// ID of the membership associated with the reservation
	MembershipID int64 `json:"membership_id"`
	// Name of the membership associated with the reservation
	MembershipName string `json:"membership"`
	// Name of drop in
	DropInName string `json:"drop_in_name"`
	// Email of drop in
	DropInEmail string `json:"drop_in_email"`
	// Whether the reservation was canceled from the waitlist
	IsCanceledFromWaitlist bool `json:"is_cancelled_from_waitlist"`
	// Whether the reservation was a late cancellation
	IsLateCancellation bool `json:"is_late_cancellation"`
	// Record creation data
	Created Created `json:"created"`
	// Record last update data
	Updated Updated `json:"updated"`
}

// Coach represents a coach for a class.
type Coach struct {
	// User ID of the coach
	ID int64 `json:"coach_id"`
	// Name of the coach
	Name string `json:"coach"`
	// ID of the payroll position of the coach
	PayrollPositionId int64 `json:"payroll_position_id"`
	// Name of the payroll position of the coach
	PayrollPositionName string `json:"payroll_position"`
}

// ClassSignIn represents a sign-in for a class.
type ClassSignIn struct {
	// ID of the sign-in.
	ID int64 `json:"id"`
	// Email of the person signed in.
	Email string `json:"email"`
	// Gender ID of the person signed in.
	GenderID int `json:"gender_id"`
	// Name of the gender of signed in person.
	GenderName string `json:"gender"`
	// ID of the class signed in to.
	ClassId int64 `json:"class_id"`
	// Name of the class signed in to.
	ClassName string `json:"class"`
	// Date and time of the sign in.
	SignedInAt DateTime `json:"sign_in_date_time"`
	// ID of source from which the sign in was made.
	SignInSourceID int64 `json:"sign_in_source_id"`
	// Name of source from which the sign in was made.
	SignInSourceName string `json:"sign_in_source"`
	// UTC Timestamp of the class start
	UTCClassStartDateTime DateTime `json:"utc_class_start_datetime"`
	// UTC Timestamp of the class end
	UTCClassEndDateTime DateTime `json:"utc_class_end_datetime"`
	// Local Timestamp of the class start
	LocalClassStartDateTime DateTime `json:"local_class_start_datetime"`
	// Local Timestamp of the class end
	LocalClassEndDateTime DateTime `json:"local_class_end_datetime"`
	// ID of the program the class is part of
	ProgramID int64 `json:"program_id"`
	// Name of the program the class is part of
	ProgramName string `json:"program"`
	// ID of the location the class is held at
	LocationID int64 `json:"location_id"`
	// Name of the location the class is held at
	LocationName string `json:"location"`
	// ID of the membership the class sign-in is associated with
	MembershipID int64 `json:"membership_id"`
	// Name of the membership the class sign-in is associated with
	MembershipName string `json:"membership"`
	// ID of online membership sale class sign-in is associated with
	OnlineMembershipSaleID int64 `json:"online_membership_sale_id"`
	// Whether the sign-in is a drop-in
	IsDropIn bool `json:"is_drop_in"`
	// Whether the sign-in happened automatically
	IsAutoSignIn bool `json:"is_auto_sign_in"`
	// Whether the sign-in counts against attendance limits
	CountsTowardsAttendanceLimits bool `json:"counts_towards_attendance_limits"`
	// Whether membership enforcement was enabled at the time of sign-in
	IsMembershipEnforcementEnabled bool `json:"is_membership_enforcement_enabled"`
	// ID of the enforcement type for limited class plans at the time of sign-in
	LimitedPlanEnforcementTypeID int64 `json:"limited_plan_enforcement_type_id"`
	// Name of the enforcement type for limited class plans at the time of sign-in
	LimitedPlanEnforcementTypeName string `json:"limited_plan_enforcement_type"`
	// ID of the enforcement type for class packs at the time of sign-in
	ClassPackEnforcementTypeID int64 `json:"class_pack_enforcement_type_id"`
	// Name of the enforcement type for class packs at the time of sign-in
	ClassPackEnforcementTypeName string `json:"class_pack_enforcement_type"`
	// Whether the attended email was sent
	IsAttendedEmailSent bool `json:"is_attended_email_sent"`
	// Record creation data
	Created Created `json:"created"`
	// Record last update data
	Updated Updated `json:"updated"`
}
