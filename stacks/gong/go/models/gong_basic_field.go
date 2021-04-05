package models

import "go/types"

type GongBasicField struct {
	Name string

	// swagger:ignore
	Type types.Type `gorm:"-"`

	// swagger:ignore
	basicKind     types.BasicKind
	BasicKindName string

	GongEnum *GongEnum // not null if it is an enum variable
}

func (basicField *GongBasicField) GetName() string {
	return basicField.Name
}
