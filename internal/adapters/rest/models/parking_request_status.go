package models

// UpdateParkingRequestStatusRequest represents the HTTP request to update parking request status.
type UpdateParkingRequestStatusRequest struct {
	Status string `json:"status"`
}

// UpdateParkingRequestStatusResponse represents the response to a HTTP request to update parking request status.
type UpdateParkingRequestStatusResponse struct {
	Message string `json:"message"`
}
