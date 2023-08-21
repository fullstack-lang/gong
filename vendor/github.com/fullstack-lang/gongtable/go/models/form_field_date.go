package models

import "time"

type FormFieldDate struct {
	Name string

	// we will only update the day
	Value time.Time
}
