// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Content_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *Content) GongCopyBasicFields(to *Content) {
	// insertion point
	*to = *from
}

type JpgImage_WOP struct {
	// insertion point

	Name string

	Base64Content string
}

func (from *JpgImage) GongCopyBasicFields(to *JpgImage) {
	// insertion point
	*to = *from
}

type PngImage_WOP struct {
	// insertion point

	Name string

	Base64Content string
}

func (from *PngImage) GongCopyBasicFields(to *PngImage) {
	// insertion point
	*to = *from
}

type SvgImage_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *SvgImage) GongCopyBasicFields(to *SvgImage) {
	// insertion point
	*to = *from
}

// end of insertion point
