// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour
var _ = __GONG_time_The_fool_doth_think_he_is_wise__

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Chapter_WOP struct {
	// insertion point

	Name string

	MardownContent string
}

func (from *Chapter) CopyBasicFields(to *Chapter) {
	// insertion point
	to.Name = from.Name
	to.MardownContent = from.MardownContent
}

type Content_WOP struct {
	// insertion point

	Name string

	MardownContent string

	ContentPath string

	OutputPath string

	LayoutPath string

	StaticPath string

	Target Target
}

func (from *Content) CopyBasicFields(to *Content) {
	// insertion point
	to.Name = from.Name
	to.MardownContent = from.MardownContent
	to.ContentPath = from.ContentPath
	to.OutputPath = from.OutputPath
	to.LayoutPath = from.LayoutPath
	to.StaticPath = from.StaticPath
	to.Target = from.Target
}

type Page_WOP struct {
	// insertion point

	Name string

	MardownContent string
}

func (from *Page) CopyBasicFields(to *Page) {
	// insertion point
	to.Name = from.Name
	to.MardownContent = from.MardownContent
}

