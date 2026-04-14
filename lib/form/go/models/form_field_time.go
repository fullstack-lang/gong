package models

import "time"

type FormFieldTime struct {
	Name string

	// we will only update the hours, minutes, seconds, ...
	Value time.Time

	// For an <input> element with type="time",
	// the step attribute determines the step size between time values. Its value is specified in seconds.
	// When you set step="1", it means the smallest interval of time the user can select or input is one second.
	//  This allows users to pick any time down to the second. If you omit the step attribute or set it to
	// a value like 60 (which corresponds to one minute), users can only pick times in one-minute intervals.
	//
	// For type="time", it determines the step between times.
	// If you want granularity down to milliseconds, you can set step="0.001", which is 1 millisecond
	Step float64
}
