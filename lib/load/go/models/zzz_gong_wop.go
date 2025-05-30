// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type FileToDownload_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *FileToDownload) CopyBasicFields(to *FileToDownload) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type FileToUpload_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *FileToUpload) CopyBasicFields(to *FileToUpload) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type Message_WOP struct {
	// insertion point

	Name string
}

func (from *Message) CopyBasicFields(to *Message) {
	// insertion point
	to.Name = from.Name
}

