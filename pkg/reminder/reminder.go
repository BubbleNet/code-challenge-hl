package reminder

import "time"

type (
	// Reminder
	// TODO: Reminder should have properties allowing dynamic alerting based on session time
	Reminder struct {
		AlertTime time.Time `json:"alert_time"`
	}
)
