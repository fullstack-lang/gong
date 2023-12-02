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

	// IsDocLink is true if the field is a string field
	// that is set at compile time by a DocLink directive //gong:ident
	IsDocLink bool

	// if there is a gong:text gong directive
	IsTextArea bool

	// IsBespokeWidth indicates wether the field shall be displayed with a
	// beskope width
	// this is set with a gong:width <width in px> gong directive
	IsBespokeWidth bool
	BespokeWidth   int

	//
	IsBespokeHeight bool
	BespokeHeight   int
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
