package models

import "time"

// CheckForLateArrivalsRequest represents a body of incoming HTTP request to check for late arrivals.
type CheckForLateArrivalsRequest struct {
	Threshold time.Duration `json:"threshold"`
}
