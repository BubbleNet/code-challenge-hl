package scheduler

import "time"

func Start() {
	stop := false
	// upcomingReminders = []reminder.Reminders
	for {
		time.Sleep(1 * time.Minute)
		if stop {
			break
		}
		// Keep a list of upcoming reminders to check for every x minutes
		// This is done so we aren't checking the db too frequently
		// Update this list whenever reminders are modified
		// upon a server restart, get all reminders from db occurring within the next x hours, x being memory dependant

		// Order reminder list with head of list = newest reminder
		// remove first element from list and execute notification when time.now() > reminderTime
	}
}
