package models

import "time"

type Milestone struct {
	Name string
	Date time.Time

	// DisplayVerticalBar indicates wether the milestone
	// has a vertical vertical on the whole gantt
	DisplayVerticalBar bool

	// a red diamond a text anchor will be displayed
	LanesToDisplay []*Lane
}
