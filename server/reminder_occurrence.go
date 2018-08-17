package main

import (
	"time"
)

type ReminderOccurrence struct {

	ReminderOccurrenceId string

	ReminderId string

	Occurrence time.Time

	Snoozed time.Time

	Repeat string
}