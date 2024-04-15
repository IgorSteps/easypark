package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	usecases "github.com/IgorSteps/easypark/internal/usecases/notification"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateNotificationTestSuite struct {
	suite.Suite

	testLogger           *logrus.Logger
	loggerHook           *test.Hook
	mockNotificationRepo *mocks.NotificationRepository
	mockSpaceRepo        *mocks.ParkingSpaceRepository
	mockRequestRepo      *mocks.ParkingRequestRepository
	mockAlertCreator     *mocks.AlertCreator
}

func (s *CreateNotificationTestSuite) SetupSuite() {
	s.testLogger, s.loggerHook = test.NewNullLogger()
	s.mockNotificationRepo = &mocks.NotificationRepository{}
	s.mockSpaceRepo = &mocks.ParkingSpaceRepository{}
	s.mockRequestRepo = &mocks.ParkingRequestRepository{}
	s.mockAlertCreator = &mocks.AlertCreator{}
}

func (s *CreateNotificationTestSuite) AfterTest() {
	s.mockAlertCreator.AssertExpectations(s.T())
	s.mockNotificationRepo.AssertExpectations(s.T())
	s.mockRequestRepo.AssertExpectations(s.T())
	s.mockSpaceRepo.AssertExpectations(s.T())
}

func (s *CreateNotificationTestSuite) TestCreateNotification_Execute_HappyPath_Arrival() {
	// --------
	// ASSEMBLE
	// --------
	usecase := usecases.NewCreateNotification(
		s.testLogger,
		s.mockNotificationRepo,
		s.mockSpaceRepo,
		s.mockRequestRepo,
		s.mockAlertCreator,
	)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testParkingRequestID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 0 // arrival

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   testLocation,
		Status: entities.ParkingSpaceStatusAvailable,
	}
	s.mockSpaceRepo.EXPECT().GetSingle(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()

	testParkingRequest := entities.ParkingRequest{
		ID:                      testParkingRequestID,
		UserID:                  testDriverID,
		ParkingSpaceID:          &testSpaceID,
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5 * time.Minute),
		Status:                  entities.RequestStatusApproved,
	}
	s.mockRequestRepo.EXPECT().GetSingle(testCtx, testParkingRequestID).Return(testParkingRequest, nil).Once()

	s.mockNotificationRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()
	testParkingSpace.Status = entities.ParkingSpaceStatusOccupied
	s.mockSpaceRepo.EXPECT().Save(testCtx, &testParkingSpace).Return(nil).Once()

	testParkingRequest.Status = entities.RequestStatusActive
	s.mockRequestRepo.EXPECT().Save(testCtx, &testParkingRequest).Return(nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testParkingRequestID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Error must be nil")
	s.Require().Equal(entities.ArrivalNotification, notification.Type, "Notification type is wrong")
	s.Require().Equal(testDriverID, notification.DriverID)
	s.Require().Equal(testLocation, notification.Location)
	s.Require().Equal(testSpaceID, notification.ParkingSpaceID)
}

func (s *CreateNotificationTestSuite) TestCreateNotification_Execute_HappyPath_Departure() {
	// --------
	// ASSEMBLE
	// --------
	usecase := usecases.NewCreateNotification(
		s.testLogger,
		s.mockNotificationRepo,
		s.mockSpaceRepo,
		s.mockRequestRepo,
		s.mockAlertCreator,
	)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testParkingRequestID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 1 // departure

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   testLocation,
		Status: entities.ParkingSpaceStatusOccupied, // occupied
	}
	s.mockSpaceRepo.EXPECT().GetSingle(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()

	testParkingRequest := entities.ParkingRequest{
		ID:                      testParkingRequestID,
		UserID:                  testDriverID,
		ParkingSpaceID:          &testSpaceID,
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5 * time.Minute),
		Status:                  entities.RequestStatusActive, // active
	}
	s.mockRequestRepo.EXPECT().GetSingle(testCtx, testParkingRequestID).Return(testParkingRequest, nil).Once()

	s.mockNotificationRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

	testParkingSpace.Status = entities.ParkingSpaceStatusAvailable // update space to occupied
	s.mockSpaceRepo.EXPECT().Save(testCtx, &testParkingSpace).Return(nil).Once()

	testParkingRequest.Status = entities.RequestStatusCompleted // update request to completed
	s.mockRequestRepo.EXPECT().Save(testCtx, &testParkingRequest).Return(nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testParkingRequestID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Error must be nil")
	s.Require().Equal(entities.DepartureNotification, notification.Type, "Notification type is wrong")
	s.Require().Equal(testDriverID, notification.DriverID)
	s.Require().Equal(testLocation, notification.Location)
	s.Require().Equal(testSpaceID, notification.ParkingSpaceID)
}

func (s *CreateNotificationTestSuite) TestCreateNotification_Execute_UnhappyPath_DriverIDMismatch() {
	// --------
	// ASSEMBLE
	// --------
	usecase := usecases.NewCreateNotification(
		s.testLogger,
		s.mockNotificationRepo,
		s.mockSpaceRepo,
		s.mockRequestRepo,
		s.mockAlertCreator,
	)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testParkingRequestID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 1 // departure

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   testLocation,
		Status: entities.ParkingSpaceStatusOccupied, // occupied
	}
	s.mockSpaceRepo.EXPECT().GetSingle(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()

	testParkingRequest := entities.ParkingRequest{
		ID:                      testParkingRequestID,
		UserID:                  uuid.New(), // different userID
		ParkingSpaceID:          &testSpaceID,
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5 * time.Minute),
		Status:                  entities.RequestStatusActive, // active
	}
	s.mockRequestRepo.EXPECT().GetSingle(testCtx, testParkingRequestID).Return(testParkingRequest, nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testParkingRequestID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	s.Require().Empty(notification)
	s.Require().IsType(&repositories.InvalidInputError{}, err)
	s.Require().EqualError(err, "userID in the parking request doesn't match userID of whoever created the notification", "Error must be nil")
	s.Require().Equal("userID in the parking request doesn't match userID of whoever created the notification", s.loggerHook.LastEntry().Message)
	s.Require().Equal(logrus.ErrorLevel, s.loggerHook.LastEntry().Level)
}

func (s *CreateNotificationTestSuite) TestCreateNotification_Execute_UnhappyPath_SpaceIDMismatch() {
	// --------
	// ASSEMBLE
	// --------
	usecase := usecases.NewCreateNotification(
		s.testLogger,
		s.mockNotificationRepo,
		s.mockSpaceRepo,
		s.mockRequestRepo,
		s.mockAlertCreator,
	)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testParkingRequestID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 1 // departure

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   testLocation,
		Status: entities.ParkingSpaceStatusOccupied, // occupied
	}
	s.mockSpaceRepo.EXPECT().GetSingle(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()

	differentSpaceID := uuid.New()
	testParkingRequest := entities.ParkingRequest{
		ID:                      testParkingRequestID,
		UserID:                  testDriverID,
		ParkingSpaceID:          &differentSpaceID, // different space id
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5 * time.Minute),
		Status:                  entities.RequestStatusActive, // active
	}
	s.mockRequestRepo.EXPECT().GetSingle(testCtx, testParkingRequestID).Return(testParkingRequest, nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testParkingRequestID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	s.Require().Empty(notification)
	s.Require().IsType(&repositories.InvalidInputError{}, err)
	s.Require().EqualError(err, "parkingSpaceID in the parking request doesn't match parkingSpaceID in the notification", "Error must be nil")
	s.Require().Equal("parkingSpaceID in the parking request doesn't match parkingSpaceID in the notification", s.loggerHook.LastEntry().Message)
	s.Require().Equal(logrus.ErrorLevel, s.loggerHook.LastEntry().Level)
}

func (s *CreateNotificationTestSuite) TestCreateNotification_Execute_UnhappyPath_ArrivalWrongParkingSpaceStatus() {
	// --------
	// ASSEMBLE
	// --------
	usecase := usecases.NewCreateNotification(
		s.testLogger,
		s.mockNotificationRepo,
		s.mockSpaceRepo,
		s.mockRequestRepo,
		s.mockAlertCreator,
	)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testParkingRequestID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 0 // arrival

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   testLocation,
		Status: entities.ParkingSpaceStatusOccupied, // occupied, should be available because of arrival
	}
	s.mockSpaceRepo.EXPECT().GetSingle(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()

	testParkingRequest := entities.ParkingRequest{
		ID:                      testParkingRequestID,
		UserID:                  testDriverID,
		ParkingSpaceID:          &testSpaceID,
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5 * time.Minute),
		Status:                  entities.RequestStatusActive, // active
	}
	s.mockRequestRepo.EXPECT().GetSingle(testCtx, testParkingRequestID).Return(testParkingRequest, nil).Once()
	s.mockNotificationRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testParkingRequestID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	s.Require().Empty(notification)
	s.Require().IsType(&repositories.InvalidInputError{}, err)
	s.Require().EqualError(err, "parking space that isn't available received arrival notification")
	s.Require().Equal("parking space that isn't available received arrival notification", s.loggerHook.LastEntry().Message)
	s.Require().Equal(testSpaceID, s.loggerHook.LastEntry().Data["parking space id"])
	s.Require().Equal(logrus.ErrorLevel, s.loggerHook.LastEntry().Level)
}

func (s *CreateNotificationTestSuite) TestCreateNotification_Execute_UnhappyPath_ArrivalWrongParkingRequestStatus() {
	// --------
	// ASSEMBLE
	// --------
	usecase := usecases.NewCreateNotification(
		s.testLogger,
		s.mockNotificationRepo,
		s.mockSpaceRepo,
		s.mockRequestRepo,
		s.mockAlertCreator,
	)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testParkingRequestID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 0 // arrival

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   testLocation,
		Status: entities.ParkingSpaceStatusAvailable, // available
	}
	s.mockSpaceRepo.EXPECT().GetSingle(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()

	testParkingRequest := entities.ParkingRequest{
		ID:                      testParkingRequestID,
		UserID:                  testDriverID,
		ParkingSpaceID:          &testSpaceID,
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5 * time.Minute),
		Status:                  entities.RequestStatusRejected, // rejected, will error
	}
	s.mockRequestRepo.EXPECT().GetSingle(testCtx, testParkingRequestID).Return(testParkingRequest, nil).Once()
	s.mockNotificationRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testParkingRequestID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	s.Require().Empty(notification)
	s.Require().IsType(&repositories.InvalidInputError{}, err)
	s.Require().EqualError(err, "parking request that isn't approved received arrival notification")
	s.Require().Equal("parking request that isn't approved received arrival notification", s.loggerHook.LastEntry().Message)
	s.Require().Equal(testParkingRequestID, s.loggerHook.LastEntry().Data["parking request id"])
	s.Require().Equal(logrus.ErrorLevel, s.loggerHook.LastEntry().Level)
}

func (s *CreateNotificationTestSuite) TestCreateNotification_Execute_HappyPath_LocationMismatchAlert() {
	// --------
	// ASSEMBLE
	// --------
	usecase := usecases.NewCreateNotification(
		s.testLogger,
		s.mockNotificationRepo,
		s.mockSpaceRepo,
		s.mockRequestRepo,
		s.mockAlertCreator,
	)
	s.testLogger.Level = logrus.DebugLevel
	testCtx := context.Background()
	testDriverID := uuid.New()
	testParkingRequestID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 0 // arrival

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   "different location",
		Status: entities.ParkingSpaceStatusAvailable, // available
	}
	s.mockSpaceRepo.EXPECT().GetSingle(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()

	testParkingRequest := entities.ParkingRequest{
		ID:                      testParkingRequestID,
		UserID:                  testDriverID,
		ParkingSpaceID:          &testSpaceID,
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5 * time.Minute),
		Status:                  entities.RequestStatusApproved, // approved
	}
	s.mockRequestRepo.EXPECT().GetSingle(testCtx, testParkingRequestID).Return(testParkingRequest, nil).Once()
	s.mockNotificationRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

	testAlert := &entities.Alert{}
	testAlertMsg := "driver arrived at wrong parking space"
	testAlert.OnLocationMismatchAlertCreate(testAlertMsg, testDriverID, testSpaceID)
	s.mockAlertCreator.EXPECT().Execute(testCtx, entities.LocationMismatch, testAlertMsg, testDriverID, testSpaceID).Return(testAlert, nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testParkingRequestID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err)
	s.Require().Equal(entities.ArrivalNotification, notification.Type, "Notification type is wrong")
	s.Require().Equal(testDriverID, notification.DriverID)
	s.Require().Equal(testLocation, notification.Location)
	s.Require().Equal(testSpaceID, notification.ParkingSpaceID)

	// assert logger
	s.Require().Equal(2, len(s.loggerHook.Entries), "Logger must log twice")
	s.Require().Equal("created location mismatch alert", s.loggerHook.Entries[0].Message)
	s.Require().Equal(testAlert.ID, s.loggerHook.Entries[0].Data["id"])
	s.Require().Equal(testAlert.Message, s.loggerHook.Entries[0].Data["msg"])
	s.Require().Equal(testAlert.Type, s.loggerHook.Entries[0].Data["type"])
	s.Require().Equal(testAlert.UserID, s.loggerHook.Entries[0].Data["driverID"])
	s.Require().Equal("location mismatch on arrival", s.loggerHook.Entries[1].Message)
}

func (s *CreateNotificationTestSuite) TestCreateNotification_Execute_UnhappyPath_DepartureWrongParkingSpaceStatus() {
	// --------
	// ASSEMBLE
	// --------
	usecase := usecases.NewCreateNotification(
		s.testLogger,
		s.mockNotificationRepo,
		s.mockSpaceRepo,
		s.mockRequestRepo,
		s.mockAlertCreator,
	)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testParkingRequestID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 1 // departure

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   testLocation,
		Status: entities.ParkingSpaceStatusAvailable, // available, will error
	}
	s.mockSpaceRepo.EXPECT().GetSingle(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()

	testParkingRequest := entities.ParkingRequest{
		ID:                      testParkingRequestID,
		UserID:                  testDriverID,
		ParkingSpaceID:          &testSpaceID,
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5 * time.Minute),
		Status:                  entities.RequestStatusActive, // active
	}
	s.mockRequestRepo.EXPECT().GetSingle(testCtx, testParkingRequestID).Return(testParkingRequest, nil).Once()
	s.mockNotificationRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testParkingRequestID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	s.Require().Empty(notification)
	s.Require().IsType(&repositories.InvalidInputError{}, err)
	s.Require().EqualError(err, "parking space that isn't occupied cannot receive departure notification")
	s.Require().Equal("parking space that isn't occupied cannot receive departure notification", s.loggerHook.LastEntry().Message)
	s.Require().Equal(testSpaceID, s.loggerHook.LastEntry().Data["parking space id"])
	s.Require().Equal(logrus.ErrorLevel, s.loggerHook.LastEntry().Level)
}

func (s *CreateNotificationTestSuite) TestCreateNotification_Execute_UnhappyPath_DepartureWrongParkingRequestStatus() {
	// --------
	// ASSEMBLE
	// --------
	usecase := usecases.NewCreateNotification(
		s.testLogger,
		s.mockNotificationRepo,
		s.mockSpaceRepo,
		s.mockRequestRepo,
		s.mockAlertCreator,
	)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testParkingRequestID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 1 // departure

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   testLocation,
		Status: entities.ParkingSpaceStatusOccupied, // occupied
	}
	s.mockSpaceRepo.EXPECT().GetSingle(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()

	testParkingRequest := entities.ParkingRequest{
		ID:                      testParkingRequestID,
		UserID:                  testDriverID,
		ParkingSpaceID:          &testSpaceID,
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5 * time.Minute),
		Status:                  entities.RequestStatusPending, // pending, will error
	}
	s.mockRequestRepo.EXPECT().GetSingle(testCtx, testParkingRequestID).Return(testParkingRequest, nil).Once()
	s.mockNotificationRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testParkingRequestID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	s.Require().Empty(notification)
	s.Require().IsType(&repositories.InvalidInputError{}, err)
	s.Require().EqualError(err, "parking request that isn't active cannot receive departure notification")
	s.Require().Equal("parking request that isn't active cannot receive departure notification", s.loggerHook.LastEntry().Message)
	s.Require().Equal(testParkingRequestID, s.loggerHook.LastEntry().Data["parking request id"])
	s.Require().Equal(logrus.ErrorLevel, s.loggerHook.LastEntry().Level)
}

func (s *CreateNotificationTestSuite) TestCreateNotification_Execute_UnhappyPath_ParsingFailed() {
	// --------
	// ASSEMBLE
	// --------
	usecase := usecases.NewCreateNotification(
		s.testLogger,
		s.mockNotificationRepo,
		s.mockSpaceRepo,
		s.mockRequestRepo,
		s.mockAlertCreator,
	)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testParkingRequestID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 100

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testParkingRequestID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	s.Require().Empty(notification)
	s.Require().EqualError(err, "invalid notification type", "Error message is wrong")
	s.Require().IsType(&repositories.InvalidInputError{}, err, "Error is of wrong type")
}

func TestCreateNotificationTestSuite(t *testing.T) {
	suite.Run(t, new(CreateNotificationTestSuite))
}
