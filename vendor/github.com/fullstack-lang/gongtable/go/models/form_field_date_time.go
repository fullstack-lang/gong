package models

import "time"

type FormFieldDateTime struct {
	Name string

	// we will only update the day
	Value time.Time
}
