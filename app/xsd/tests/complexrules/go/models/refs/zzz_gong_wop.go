// generated code - do not edit
package refs

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type RefTargetType_WOP struct {
	// insertion point

	Name string

	Id string

	Status Enum_StatusEnum

	CreatedBy string

	Label string
}

func (from *RefTargetType) GongCopyBasicFields(to *RefTargetType) {
	// insertion point
	*to = *from
}

// end of insertion point
