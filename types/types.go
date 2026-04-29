package core

import "time"

// Created is a struct that contains information about an object's creation.
type Created struct {
	CreatedByID int64     `json:"created_by_id"`
	CreatedBy   string    `json:"created_by"`
	CreatedOn   time.Time `json:"created_on"`
}

// Updated is a struct that contains information about an object's last update.
type Updated struct {
	UpdatedByID int64     `json:"updated_by_id"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedOn   time.Time `json:"updated_on"`
}
