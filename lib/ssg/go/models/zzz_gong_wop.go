// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Chapter_WOP struct {
	// insertion point
	Name string
	Weigth float64
	Description string
}

func (from *Chapter) CopyBasicFields(to *Chapter) {
	// insertion point
	to.Name = from.Name
	to.Weigth = from.Weigth
	to.Description = from.Description
}

type Content_WOP struct {
	// insertion point
	Name string
	Text string
	ContentPath string
}

func (from *Content) CopyBasicFields(to *Content) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
	to.ContentPath = from.ContentPath
}

