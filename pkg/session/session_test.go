package session

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestJsonSession_ToSession(t *testing.T) {
	tests := []struct {
		name          string
		input         JsonSession
		expected      Session
		expectedError error
	}{
		{
			"No Changes",
			JsonSession{
				UserId:          "",
				SessionTitle:    "",
				Routine:         "",
				Timestamp:       "",
				SessionDuration: "",
				Frequency:       "",
				Reminders:       nil,
				Notes:           nil,
			},
			Session{
				userId:          uuid.UUID{},
				sessionTitle:    "",
				routine:         nil,
				timestamp:       time.Time{},
				sessionDuration: 0,
				frequency:       "",
				reminders:       nil,
				notes:           nil,
			},
			nil,
		},
	}

	for index, test := range tests {
		zap.S().Info(fmt.Sprintf("Starting test %d of %d:  %s", index+1, len(tests), test.name))
		res, err := test.input.ToSession()
		if test.expectedError != nil {
			assert.Equal(t, test.expectedError.Error(), err.Error())
		} else {
			assert.Equal(t, test.expected, res)
		}

	}
}