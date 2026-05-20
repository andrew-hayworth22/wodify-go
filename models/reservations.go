package models

// ClassSignIn represents a sign-in for a class.
type ClassSignIn struct {
	// ID of the sign-in.
	ID int64 `json:"id"`
	// Email of the person signed in.
	Email string `json:"email"`
	// Gender of the person signed in.
	Gender Gender `json:"gender_id"`
	// ID of the class signed in to.
	ClassId int64 `json:"class_id"`
	// Name of the class signed in to.
	ClassName string `json:"class"`
	// Date and time of the sign in.
	SignedInAt DateTime `json:"sign_in_date_time"`
	// Source from which the sign in was made.
	SignInSourceID int64 `json:"sign_in_source_id"`
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
