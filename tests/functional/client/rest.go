package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type RestClientSuite struct {
	suite.Suite

	BaseURL string
	Client  *http.Client
}

// SetupSuite initializes the REST client.
func (s *RestClientSuite) SetupSuite() {
	s.BaseURL = "http://localhost:8080"
	s.Client = &http.Client{}
}

// TearDownTest cleans up resources after every test.
func (s *RestClientSuite) TearDownTest() {
	// TODO: Clean up resources properly
	cmd := exec.Command("../../build/cleandb.sh")
	err := cmd.Run()
	if err != nil {
		s.T().Logf("Failed to truncate tables: %v", err)
	} else {
		s.T().Log("Tables truncated successfully.")
	}

	s.T().Log("Tearing down suite...")
}

// CreateUser interacts with the REST API to create a new user.
func (s *RestClientSuite) CreateUser(ctx context.Context, req *models.UserCreationRequest) ([]byte, int, error) {
	return s.sendRequest(ctx, "POST", "/register", req)
}

// LoginUser interacts with the REST API to login a user.
func (s *RestClientSuite) LoginUser(ctx context.Context, req *models.LoginUserRequest) ([]byte, int, error) {
	return s.sendRequest(ctx, http.MethodPost, "/login", req)
}

// PlaceholderDriverRoute interacts with the REST API to a placeholder for driver routrs.
func (s *RestClientSuite) PlaceholderDriverRoute(ctx context.Context, token string) ([]byte, int, error) {
	return s.sendRequestWithToken(ctx, http.MethodGet, "/driver", nil, token)
}

// GetAllDrivers interacts with the REST API to get get all drivers.
func (s *RestClientSuite) GetAllDrivers(ctx context.Context, token string) ([]byte, int, error) {
	return s.sendRequestWithToken(ctx, http.MethodGet, "/drivers", nil, token)
}

// BanDriver interacts with the REST API to ban a driver with the given ID.
func (s *RestClientSuite) BanDriver(ctx context.Context, token, id string, req *models.UpdateStatusRequest) ([]byte, int, error) {
	return s.sendRequestWithToken(ctx, http.MethodPatch, "/drivers/"+id+"/status", req, token)
}

// CreateParkingRequest interacts with the REST API to create a parkig request for the given userID.
func (s *RestClientSuite) CreateParkingRequest(ctx context.Context, token, userID string, req *models.CreateParkingRequestRequest) ([]byte, int, error) {
	return s.sendRequestWithToken(ctx, http.MethodPost, "/drivers/"+userID+"/parking-requests", req, token)
}

// UpdateParkingRequestStatus interacts with the REST API to update a parking request with the given requestID.
func (s *RestClientSuite) UpdateParkingRequestStatus(ctx context.Context, token, requestID string, req *models.UpdateParkingRequestStatusRequest) ([]byte, int, error) {
	return s.sendRequestWithToken(ctx, http.MethodPatch, "/parking-requests/"+requestID+"/status", req, token)
}

// CreateParkingLot interacts with the REST API to create a parking lot.
func (s *RestClientSuite) CreateParkingLot(ctx context.Context, token string, req *models.CreateParkingLotRequest) ([]byte, int, error) {
	return s.sendRequestWithToken(ctx, http.MethodPost, "/parking-lots", req, token)
}

// UpdateParkingRequestSpace interacts with the REST API to update a parking request with a parking space.
func (s *RestClientSuite) UpdateParkingRequestSpace(ctx context.Context, token string, parkingReqID uuid.UUID, req *models.ParkingRequestSpaceUpdateRequest) ([]byte, int, error) {
	return s.sendRequestWithToken(ctx, http.MethodPatch, "/parking-requests/"+parkingReqID.String()+"/space", req, token)
}

// GetAllParkingRequests interacts with the REST API to get all parking requests.
func (s *RestClientSuite) GetAllParkingRequests(ctx context.Context, token string) ([]byte, int, error) {
	return s.sendRequestWithToken(ctx, http.MethodGet, "/parking-requests", nil, token)
}

// sendRequest sends a HTTP request via provided method and path.
func (s *RestClientSuite) sendRequest(ctx context.Context, method, path string, body interface{}) ([]byte, int, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, 0, err
	}

	url := fmt.Sprintf("%s%s", s.BaseURL, path)
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return responseBody, resp.StatusCode, nil
}

// sendRequestWithToken sends a HTTP request via provided method and path with an auth token.
func (s *RestClientSuite) sendRequestWithToken(ctx context.Context, method, path string, body interface{}, token string) ([]byte, int, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, 0, err
	}

	url := fmt.Sprintf("%s%s", s.BaseURL, path)
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return responseBody, resp.StatusCode, nil
}

func (s *RestClientSuite) UnmarshalHTTPResponse(responseBody []byte, targetModel interface{}) error {
	if err := json.Unmarshal(responseBody, &targetModel); err != nil {
		s.T().Logf("failed to unmarshal response body: %v", err)
		return err
	}

	return nil
}
