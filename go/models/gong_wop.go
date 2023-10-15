// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type GongBasicField_WOP struct {
	// insertion point
	Name string
	BasicKindName string
	DeclaredType string
	CompositeStructName string
	Index int
	IsDocLink bool
	IsTextArea bool
}

type GongEnum_WOP struct {
	// insertion point
	Name string
	Type GongEnumType
}

type GongEnumValue_WOP struct {
	// insertion point
	Name string
	Value string
}

type GongLink_WOP struct {
	// insertion point
	Name string
	Recv string
	ImportPath string
}

type GongNote_WOP struct {
	// insertion point
	Name string
	Body string
	BodyHTML string
}

type GongStruct_WOP struct {
	// insertion point
	Name string
	HasOnAfterUpdateSignature bool
	IsIgnoredForFront bool
}

type GongTimeField_WOP struct {
	// insertion point
	Name string
	Index int
	CompositeStructName string
}

type Meta_WOP struct {
	// insertion point
	Name string
	Text string
}

type MetaReference_WOP struct {
	// insertion point
	Name string
}

type ModelPkg_WOP struct {
	// insertion point
	Name string
	PkgPath string
}

type PointerToGongStructField_WOP struct {
	// insertion point
	Name string
	Index int
	CompositeStructName string
}

type SliceOfPointerToGongStructField_WOP struct {
	// insertion point
	Name string
	Index int
	CompositeStructName string
}

