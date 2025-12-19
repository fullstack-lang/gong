// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Product_WOP struct {
	// insertion point

	Name string

	IsExpanded bool
}

func (from *Product) CopyBasicFields(to *Product) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
}

type Project_WOP struct {
	// insertion point

	Name string

	IsExpanded bool
}

func (from *Project) CopyBasicFields(to *Project) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
}

type Root_WOP struct {
	// insertion point

	Name string
}

func (from *Root) CopyBasicFields(to *Root) {
	// insertion point
	to.Name = from.Name
}

type Task_WOP struct {
	// insertion point

	Name string
}

func (from *Task) CopyBasicFields(to *Task) {
	// insertion point
	to.Name = from.Name
}

