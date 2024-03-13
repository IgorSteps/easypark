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
	return s.sendRequest(ctx, "POST", "/login", req)
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

func (s *RestClientSuite) UnmarshalHTTPResponse(responseBody []byte, targetModel interface{}) error {
	if err := json.Unmarshal(responseBody, &targetModel); err != nil {
		s.T().Logf("failed to unmarshal response body: %v", err)
		return err
	}

	return nil
}
