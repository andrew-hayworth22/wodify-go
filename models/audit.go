package models

// Created is a struct that contains information about an object's creation.
type Created struct {
	// Unique ID of the user that created the record.
	CreatedByID int64 `json:"created_by_id"`
	// Name of the user that created the record.
	CreatedBy string `json:"created_by"`
	// Date and time when the record was created.
	CreatedOn DateTime `json:"created_on"`
}

// Updated is a struct that contains information about an object's last update.
type Updated struct {
	// Unique ID of the user that last updated the record.
	UpdatedByID int64 `json:"updated_by_id"`
	// Name of the user that last updated the record.
	UpdatedBy string `json:"updated_by"`
	// Date and time when the record was last updated.
	UpdatedOn DateTime `json:"updated_on"`
}
