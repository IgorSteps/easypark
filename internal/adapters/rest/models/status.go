package models

// UpdateStatusRequest represents the HTTP request to update driver's status.
type UpdateStatusRequest struct {
	Status string `json:"status"`
}

type UpdateStatusResponse struct {
	Message string `json:"message"`
}
