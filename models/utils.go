package models

// Country represents a country
type Country struct {
	// Country's ID.
	ID int `json:"country_id"`
	// Name of the country.
	Name string `json:"country"`
}

// DayOfWeek represents a day of the week
type DayOfWeek struct {
	// ID of the day of the week.
	ID int `json:"day_of_week_id"`
	// Day of the week.
	Day string `json:"day_of_week"`
}

// Gender represents a gender
type Gender struct {
	// ID of the gender.
	ID int `json:"gender_id"`
	// Name of the gender.
	Name string `json:"gender"`
}

// ObjectActionType represents an object action type
type ObjectActionType struct {
	// ID of the object action type.
	ID int `json:"action_type_id"`
	// Name of the object action type.
	Name string `json:"action_type"`
	// ID of the object type.
	ObjectTypeID int `json:"object_type_id"`
	// Name of the object type.
	ObjectTypeName string `json:"object_type"`
	// ID of the role that can perform the action.
	ActionRoleID int `json:"action_role_id"`
	// Name of the role that can perform the action.
	ActionRoleName string `json:"action_role"`
}

// ObjectType represents an object type
type ObjectType struct {
	// ID of the object type.
	ID int `json:"object_type_id"`
	// Name of the object type.
	Name string `json:"object_type"`
}

// State represents a state (location)
type State struct {
	// ID of the state.
	ID int `json:"state_id"`
	// Name of the state.
	Name string `json:"state"`
}

// UnitOfTime represents a unit of time
type UnitOfTime struct {
	// ID of the unit of time.
	ID int `json:"unit_of_time_id"`
	// Name of the unit of time (singular).
	NameSingular string `json:"unit_of_time_singular"`
	// Name of the unit of time (plural).
	NamePlural string `json:"unit_of_time_plural"`
	// Name of the unit of time (singular and plural e.g. "hour(s)").
	NameSingularPlural string `json:"unit_of_time_singular_plural"`
	// Indicates if the unit of time is a time unit (hours/minutes).
	IsTimeUnit bool `json:"is_time_unit"`
	// Number of days in the unit of time.
	NumberOfDays int `json:"number_of_days"`
}
