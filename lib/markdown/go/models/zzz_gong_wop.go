// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Content_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *Content) CopyBasicFields(to *Content) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

