package models

// PerformanceResult represents a performance result from a Wodify workout.
type PerformanceResult struct {
	// Performance result's ID.
	ID string `json:"id"`
	// ID of the client associated with the result.
	ClientID string `json:"client_id"`
	// Name of the client associated with the result.
	ClientName string `json:"client"`
	// ID of the class sign-in associated with the result.
	ClassSignInID int64 `json:"class_sign_in_id"`
	// Comment on the result.
	Comment string `json:"comment"`
	// Linkified version of the comment.
	LinkifiedComment string `json:"linkified_comment"`
	// Average score for the result.
	AverageScore float64 `json:"average_score"`
	// Calculated one-rep max for the result.
	CalculatedOneRepMax float64 `json:"calculated_one_rep_max"`
	// Calories for the result.
	Calories int64 `json:"calories"`
	// Name of the class associated with the result.
	ClassName string `json:"class"`
	// ID of the class associated with the result.
	ClassID int64 `json:"class_id"`
	// Comparison calculated one-rep max for the result.
	ComparisonCalculatedOneRepMax float64 `json:"comparison_calculated_one_rep_max"`
	// Comparison distance for the result.
	ComparisonDistance float64 `json:"comparison_distance"`
	// Comparison total distance for the result.
	ComparisonTotalDistance float64 `json:"comparison_total_distance"`
	// Comparison total weight for the result.
	ComparisonTotalWeight float64 `json:"comparison_total_weight"`
	// Comparison weight for the result.
	ComparisonWeight float64 `json:"comparison_weight"`
	// ID of the component associated with the result.
	ComponentID int64 `json:"component_id"`
	// Name of the component associated with the result.
	Component string `json:"component"`
	// ID of the component type associated with the result.
	ComponentTypeID int64 `json:"component_type_id"`
	// Name of the component type associated with the result.
	ComponentType string `json:"component_type"`
	// Date of the result.
	Date Date `json:"date"`
	// Distance for the result.
	Distance float64 `json:"distance"`
	// Order of the result within each round (child results only).
	EachRoundOrder int64 `json:"each_round_order"`
	// ID of the each-round type for the result.
	EachRoundTypeID int64 `json:"each_round_type_id"`
	// Name of the each-round type for the result.
	EachRoundType string `json:"each_round_type"`
	// Fully formatted string representation of the result.
	FullyFormattedResult string `json:"fully_formatted_result"`
	// Indicates whether the result has been saved (parent results only).
	HasBeenSaved bool `json:"has_been_saved"`
	// Indicates whether the result has been scored.
	HasBeenScored bool `json:"has_been_scored"`
	// Hours component of the result.
	Hours int64 `json:"hours"`
	// Indicates whether this is a child result.
	IsChildResult bool `json:"is_child_result"`
	// Indicates whether the comment is public.
	IsCommentPublic bool `json:"is_comment_public"`
	// Indicates whether this is the current personal record.
	IsCurrentPersonalRecord bool `json:"is_current_personal_record"`
	// Indicates whether the result is a drop-in.
	IsDropIn bool `json:"is_drop_in"`
	// Indicates whether the result belongs to a lead.
	IsLead bool `json:"is_lead"`
	// Indicates whether this is a max effort result.
	IsMaxEffort bool `json:"is_max_effort"`
	// Indicates whether the media is public.
	IsMediaPublic bool `json:"is_media_public"`
	// Indicates whether the one-rep max is estimated.
	IsOneRepMaxEstimated bool `json:"is_one_rep_max_estimated"`
	// Indicates whether this is a parent result.
	IsParentResult bool `json:"is_parent_result"`
	// Indicates whether this is a personal record.
	IsPersonalRecord bool `json:"is_personal_record"`
	// Indicates whether the result was performed as prescribed (Rx).
	IsRx bool `json:"is_rx"`
	// Indicates whether the result was performed as prescribed plus (Rx+).
	IsRxPlus bool `json:"is_rx_plus"`
	// Indicates whether the result has an associated video.
	IsVideo bool `json:"is_video"`
	// ID of the lead associated with the result.
	LeadID int64 `json:"lead_id"`
	// Name of the lead associated with the result.
	Lead string `json:"lead"`
	// ID of the location associated with the result.
	LocationID int64 `json:"location_id"`
	// Name of the location associated with the result.
	Location string `json:"location"`
	// Maximum score for the result.
	MaximumScore int64 `json:"maximum_score"`
	// URL of the media associated with the result.
	MediaURL string `json:"media_url"`
	// Minimum score for the result.
	MinimumScore int64 `json:"minimum_score"`
	// Minutes component of the result.
	Minutes int64 `json:"minutes"`
	// Number of comments on the result.
	NumberOfComments int64 `json:"number_of_comments"`
	// Number of likes on the result.
	NumberOfLikes int64 `json:"number_of_likes"`
	// ID of the original performance result.
	OriginalPerformanceResultID int64 `json:"original_performance_result_id"`
	// ID of the parent component.
	ParentComponentID int64 `json:"parent_component_id"`
	// Name of the parent component.
	ParentComponent string `json:"parent_component"`
	// Key of the parent result.
	ParentResultKey string `json:"parent_result_key"`
	// GUID of the performance result.
	PerformanceResultGUID string `json:"performance_result_guid"`
	// Custom text for the personal record.
	PersonalRecordCustomText string `json:"personal_record_custom_text"`
	// ID of the personal record type.
	PersonalRecordTypeID int64 `json:"personal_record_type_id"`
	// Name of the personal record type.
	PersonalRecordType string `json:"personal_record_type"`
	// Ranking value for the result.
	RankingValue float64 `json:"ranking_value"`
	// Firebase ranking value for the result.
	RankingValueFirebase float64 `json:"ranking_value_firebase"`
	// Rep scheme for the result.
	RepScheme string `json:"rep_scheme"`
	// Number of reps for the result.
	Reps int64 `json:"reps"`
	// ID of the result type.
	ResultTypeID int64 `json:"result_type_id"`
	// Name of the result type.
	ResultType string `json:"result_type"`
	// Number of rounds for the result.
	Rounds int64 `json:"rounds"`
	// Seconds component of the result.
	Seconds float64 `json:"seconds"`
	// Number of sets for the result.
	Sets int64 `json:"sets"`
	// Total calories for the result.
	TotalCalories float64 `json:"total_calories"`
	// Total distance for the result.
	TotalDistance float64 `json:"total_distance"`
	// Total reps for the result.
	TotalReps float64 `json:"total_reps"`
	// Total seconds for the result.
	TotalSeconds float64 `json:"total_seconds"`
	// Total weight for the result.
	TotalWeight float64 `json:"total_weight"`
	// ID of the distance unit of measure.
	UOMDistanceID int64 `json:"uom_distance_id"`
	// Name of the distance unit of measure.
	UOMDistance string `json:"uom_distance"`
	// ID of the weight unit of measure.
	UOMWeightID int64 `json:"uom_weight_id"`
	// Name of the weight unit of measure.
	UOMWeight string `json:"uom_weight"`
	// Weight for the result.
	Weight float64 `json:"weight"`
	// ID of the workout component associated with the result.
	WorkoutComponentID int64 `json:"workout_component_id"`
	// ID of the workout header associated with the result.
	WorkoutHeaderID int64 `json:"workout_header_id"`
	// Child results nested under this result.
	ChildResults []PerformanceResult `json:"child_results"`
	// Record creation data.
	Created Updated `json:"created"`
	// Record last update data.
	Updated Updated `json:"updated"`
}
