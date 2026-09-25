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
	*to = *from
}

type FileToUpload_WOP struct {
	// insertion point

	Name string

	Base64EncodedContent string
}

func (from *FileToUpload) GongCopyBasicFields(to *FileToUpload) {
	// insertion point
	*to = *from
}

type Message_WOP struct {
	// insertion point

	Name string
}

func (from *Message) GongCopyBasicFields(to *Message) {
	// insertion point
	*to = *from
}

// end of insertion point
