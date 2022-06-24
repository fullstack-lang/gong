package models

import "go/types"

type GongBasicField struct {
	Name string

	// swagger:ignore
	basicKind     types.BasicKind
	BasicKindName string

	GongEnum     *GongEnum // not null if it is an enum variable
	DeclaredType string    // "time.Duration" for instance (the underlying type being int64)

	CompositeStructName string

	Index int
}

func (gongBasicField *GongBasicField) GetIndex() int {
	return gongBasicField.Index
}

func (gongBasicField *GongBasicField) GetBasicKind() types.BasicKind {
	return gongBasicField.basicKind
}

func (gongBasicField *GongBasicField) GetCompositeStructName() string {
	return gongBasicField.CompositeStructName
}
