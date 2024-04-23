package deductions

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_Service(t *testing.T) {
	mockRepo := NewMockRepository(t)
	mockRepo.On("UpsertMaximumDeduction", mock.Anything, mock.Anything).Return(nil)
	svc := NewService(mockRepo)
	var req = Request{
		Amount: decimal.NewFromFloat(70000),
	}

	resp, err := svc.Process(context.Background(), req, "personal")

	assert.NoError(t, err)
	assert.True(t, resp.PersonalDeduction.Equal(decimal.NewFromFloat(70000)), "%s should be equal to 70000", resp.PersonalDeduction)
}

func Test_ServiceFailUpsert(t *testing.T) {
	mockRepo := NewMockRepository(t)
	mockRepo.On("UpsertMaximumDeduction", mock.Anything, mock.Anything).Return(gorm.ErrInvalidTransaction)
	svc := NewService(mockRepo)
	var req = Request{
		Amount: decimal.NewFromFloat(70000),
	}

	resp, err := svc.Process(context.Background(), req, "personal")

	assert.Error(t, err)
	assert.True(t, resp.PersonalDeduction.Equal(decimal.Zero), "%s should be equal to 0", resp.PersonalDeduction)
}
