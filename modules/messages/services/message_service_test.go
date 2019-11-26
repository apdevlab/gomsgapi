package services

import (
	"errors"
	"gomsgapi/modules/messages/repositories"
	"testing"

	mockrepos "gomsgapi/gen/mock/modules/messages/repositories"
	"gomsgapi/modules/messages/dto"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMessageRepo := mockrepos.NewMockMessageRepository(ctrl)

	tests := []struct {
		name        string
		mockData    dto.CreateMessageRequest
		mockContext func()
		wantErr     bool
		result      interface{}
	}{
		{
			name:     "failed save",
			mockData: dto.CreateMessageRequest{Message: "test message"},
			mockContext: func() {
				mockMessageRepo.EXPECT().Save(gomock.Any()).Return(errors.New("some error"))
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockContext()
			svc := NewMessageService(mockMessageRepo)
			result, err := svc.Create(test.mockData)

			if test.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, "some error")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.result, result)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name        string
		mockData    uuid.UUID
		mockContext func() repositories.MessageRepository
		wantErr     bool
		errMessage  string
		result      interface{}
	}{
		{
			name:     "get by nil id",
			mockData: uuid.Nil,
			mockContext: func() repositories.MessageRepository {
				mockMessageRepo := mockrepos.NewMockMessageRepository(ctrl)

				return mockMessageRepo
			},
			wantErr:    true,
			errMessage: "invalid id",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := test.mockContext()
			svc := NewMessageService(repo)
			result, err := svc.GetByID(test.mockData)

			if test.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, test.errMessage)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.result, result)
			}
		})
	}
}
