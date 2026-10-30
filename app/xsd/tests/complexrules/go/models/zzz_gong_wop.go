// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Catalog_WOP struct {
	// insertion point

	Name string
}

func (from *Catalog) GongCopyBasicFields(to *Catalog) {
	// insertion point
	to.Name = from.Name
}

type DerivedItemType_WOP struct {
	// insertion point

	Name string

	Code string

	Score string

	TargetRef string

	OptionA string

	OptionB int
}

func (from *DerivedItemType) GongCopyBasicFields(to *DerivedItemType) {
	// insertion point
	*to = *from
}

// end of insertion point
