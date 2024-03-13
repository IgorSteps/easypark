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
	s.BaseURL = "http://localhost:8081"
	s.Client = &http.Client{}
}

// TearDownSuite cleans up resources after suite execution.
func (s *RestClientSuite) TearDownSuite() {
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
	requestBody, err := json.Marshal(req)
	if err != nil {
		s.T().Logf("failed to marshal request: %v", err)
		return nil, 0, err
	}

	url := fmt.Sprintf("%s/register", s.BaseURL)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		s.T().Logf("failed to create request: %v", err)
		return nil, 0, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(httpReq)
	if err != nil {
		s.T().Logf("failed to send request: %v", err)
		return nil, 0, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		s.T().Logf("failed to read response body: %v", err)
		return nil, 0, err
	}

	return responseBody, resp.StatusCode, nil
}

func (s *RestClientSuite) UnmarshalCreateUserResponse(responseBody []byte, targetModel interface{}) error {
	if err := json.Unmarshal(responseBody, &targetModel); err != nil {
		s.T().Logf("failed to unmarshal response body: %v", err)
		return err
	}

	return nil
}
