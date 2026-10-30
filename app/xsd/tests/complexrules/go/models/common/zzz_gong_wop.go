// generated code - do not edit
package common

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type BaseItemType_WOP struct {
	// insertion point

	Name string

	Code string

	Score string
}

func (from *BaseItemType) GongCopyBasicFields(to *BaseItemType) {
	// insertion point
	*to = *from
}

// end of insertion point
