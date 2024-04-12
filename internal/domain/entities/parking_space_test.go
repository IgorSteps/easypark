package entities_test

import (
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestParkingSpace_OnCreate(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	space := entities.ParkingSpace{}
	testName := "cool name"
	testParkingLotID := uuid.New()

	// --------
	// ACT
	// --------
	space.OnCreate(testName, testParkingLotID)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, testName, space.Name, "Parking space names must match")
	assert.Equal(t, testParkingLotID, space.ParkingLotID, "Parking space's parking lot ids must match")
	assert.Equal(t, entities.StatusAvailable, space.Status, "Parking space statuses mast match")
	assert.NotNil(t, space.ID, "Parking space must have ID set")
}

func TestParkingSpace_CheckForOverlap(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	now := time.Now()
	oneHourLater := now.Add(time.Hour)
	twoHoursLater := now.Add(2 * time.Hour)
	threeHoursLater := now.Add(3 * time.Hour)

	// Test setup: create a list of parking requests that simulate different scenarios
	parkingRequests := []entities.ParkingRequest{
		// A request that starts now and ends in an hour
		{StartTime: now, EndTime: oneHourLater},
		// A request that starts in two hours and ends in three hours
		{StartTime: twoHoursLater, EndTime: threeHoursLater},
	}

	type fields struct {
		ID              uuid.UUID
		ParkingLotID    uuid.UUID
		Name            string
		Status          entities.ParkingSpaceStatus
		ParkingRequests []entities.ParkingRequest
	}
	type args struct {
		requestStartTime time.Time
		requestEndTime   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Overlap with the first existing request",
			fields: fields{
				ParkingRequests: parkingRequests,
			},
			args: args{
				requestStartTime: now.Add(30 * time.Minute), // Starts 30 minutes after the first request starts
				requestEndTime:   twoHoursLater,             // Ends after the first request ends but before the second starts
			},
			want: true,
		},
		{
			name: "No overlap, starts and ends before any existing request",
			fields: fields{
				ParkingRequests: parkingRequests,
			},
			args: args{
				requestStartTime: now.Add(-2 * time.Hour), // Starts 2 hours before now
				requestEndTime:   now.Add(-1 * time.Hour), // Ends 1 hour before now
			},
			want: false,
		},
		{
			name: "No overlap, starts and ends after all existing requests",
			fields: fields{
				ParkingRequests: parkingRequests,
			},
			args: args{
				requestStartTime: threeHoursLater.Add(30 * time.Minute), // Starts after the last request ends
				requestEndTime:   threeHoursLater.Add(2 * time.Hour),    // Ends later
			},
			want: false,
		},
		{
			name: "Overlap with the second existing request",
			fields: fields{
				ParkingRequests: parkingRequests,
			},
			args: args{
				requestStartTime: twoHoursLater.Add(-30 * time.Minute),   // Starts before the second request starts
				requestEndTime:   threeHoursLater.Add(-30 * time.Minute), // Ends before the second request ends
			},
			want: true,
		},
		{
			name: "Overlap with both existing requests",
			fields: fields{
				ParkingRequests: parkingRequests,
			},
			args: args{
				requestStartTime: now.Add(-30 * time.Minute),            // Starts 30 minutes before the first request starts
				requestEndTime:   threeHoursLater.Add(30 * time.Minute), // Ends 30 minutes after the second request ends
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &entities.ParkingSpace{
				ID:              tt.fields.ID,
				ParkingLotID:    tt.fields.ParkingLotID,
				Name:            tt.fields.Name,
				Status:          tt.fields.Status,
				ParkingRequests: tt.fields.ParkingRequests,
			}

			// --------
			// ACT
			// --------
			if got := s.CheckForOverlap(tt.args.requestStartTime, tt.args.requestEndTime); got != tt.want {

				// --------
				// ASSERT
				// --------
				t.Errorf("Overlap check returned %v, want %v", got, tt.want)
			}
		})
	}
}
