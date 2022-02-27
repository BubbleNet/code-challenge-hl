package session

import (
	"github.com/BubbleNet/code-challenge-hl/pkg/reminder"
	"github.com/BubbleNet/code-challenge-hl/pkg/routine"
	"github.com/google/uuid"
	"time"
)

type (
	Session struct {
		userId          uuid.UUID
		sessionTitle    string
		routine         *routine.Routine
		timestamp       time.Time
		sessionDuration time.Duration
		frequency       string
		reminders       []*reminder.Reminder
		notes           []string
	}

	// JsonSession
	//ASSUMPTION: Routine will be passed as a RoutineId
	JsonSession struct {
		UserId          string   `json:"user_id"`
		SessionTitle    string   `json:"session_title"`
		Routine         string   `json:"routine"`
		Timestamp       string   `json:"timestamp"`
		SessionDuration string   `json:"session_duration"`
		Frequency       string   `json:"frequency"`
		Reminders       []string `json:"reminders"`
		Notes           []string `json:"notes"`
	}
)

func (s *Session) GetReminders() []*reminder.Reminder {
	// TODO: Return copies instead of references
	return s.reminders
}

// ToSession converts a json representation of a Session to a Session with go types
func (j *JsonSession) ToSession() (*Session, error) {
	//  TODO: sanity checks as follows
	// session has occurrences in the future, duration is within spec, frequency is within spec
	u, err := uuid.Parse(j.UserId)
	if err != nil {
		return nil, err
	}
	t, err := time.Parse("2006-01-02T15:04:05Z", j.Timestamp)
	if err != nil {
		return nil, err
	}
	d, err := time.ParseDuration(j.SessionDuration)
	if err != nil {
		return nil, err
	}

	// ASSUMPTION: Reminders will be converted from "hour before" to datetime string before api call
	// This isn't ideal and is assumed for simplicity. Actual implementation should anticipate this being
	// a problem if sessions are modified to occur at a different time.

	// ASSUMPTION: using a string for frequency for simplicity. Probably should be an enum with a set number of options.
	// weekly/daily/every-other-day/monthly

	return &Session{
		userId:          u,
		sessionTitle:    j.SessionTitle,
		routine:         &routine.Routine{}, // TODO: Create Routines
		timestamp:       t,
		sessionDuration: d,
		frequency:       j.Frequency,
		reminders:       []*reminder.Reminder{}, // TODO: Create Reminders
		notes:           j.Notes,
	}, nil
}
