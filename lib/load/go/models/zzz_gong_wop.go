// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type FileToDownload_WOP struct {
	// insertion point

	Name string

	Base64EncodedContent string
}

func (from *FileToDownload) GongCopyBasicFields(to *FileToDownload) {
	// insertion point
	to.Name = from.Name
	to.Base64EncodedContent = from.Base64EncodedContent
}

type FileToUpload_WOP struct {
	// insertion point

	Name string

	Base64EncodedContent string
}

func (from *FileToUpload) GongCopyBasicFields(to *FileToUpload) {
	// insertion point
	to.Name = from.Name
	to.Base64EncodedContent = from.Base64EncodedContent
}

type Message_WOP struct {
	// insertion point

	Name string
}

func (from *Message) GongCopyBasicFields(to *Message) {
	// insertion point
	to.Name = from.Name
}

// end of insertion point
