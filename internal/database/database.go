package database

import (
	"github.com/BubbleNet/code-challenge-hl/pkg/reminder"
	"github.com/BubbleNet/code-challenge-hl/pkg/session"
	"github.com/google/uuid"
)

type (
	//Database - Mock a relational database by storing Sessions and Reminders in maps
	// Use maps to mock indexed queries
	Database struct {
		sessions  map[uuid.UUID]*session.Session
		reminders map[uuid.UUID]*reminder.Reminder
	}
)

func New() *Database {
	return &Database{
		sessions:  make(map[uuid.UUID]*session.Session),
		reminders: make(map[uuid.UUID]*reminder.Reminder),
	}
}

// SetSession creates an id for a new session and then inserts the session into the db.
// Returns the generated sessionId
func (d *Database) SetSession(s *session.Session) (uuid.UUID, error) {
	sessionId, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}
	d.sessions[sessionId] = s
	return sessionId, nil
}

// GetReminders returns a copy of all reminders in a session
// TODO: Support pagination
func (d *Database) GetReminders(sessionId uuid.UUID) ([]*reminder.Reminder, error) {
	if s, ok := d.sessions[sessionId]; ok {
		return s.GetReminders(), nil
	}
	return []*reminder.Reminder{}, nil
}
