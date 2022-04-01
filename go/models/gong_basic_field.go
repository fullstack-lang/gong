package models

import "go/types"

type GongBasicField struct {
	Name string

	// swagger:ignore
	Type types.Type `gorm:"-"`

	// swagger:ignore
	basicKind     types.BasicKind
	BasicKindName string

	GongEnum     *GongEnum // not null if it is an enum variable
	DeclaredType string    // "time.Duration" for instance (the underlying type being int64)

	Index int
}

func (gongBasicField *GongBasicField) GetIndex() int {
	return gongBasicField.Index
}

func (gongBasicField *GongBasicField) GetBasicKind() types.BasicKind {
	return gongBasicField.basicKind
}
