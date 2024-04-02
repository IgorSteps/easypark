package models

// UpdateStatusRequest represents the HTTP request to update driver's status.
type UpdateStatusRequest struct {
	Status string `json:"status"`
}

// UpdateStatusResponse represents the response to a HTTP request to update driver's status.
type UpdateStatusResponse struct {
	Message string `json:"message"`
}
