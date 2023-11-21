// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Classdiagram_WOP struct {
	// insertion point
	Name string
	IsInDrawMode bool
}

type DiagramPackage_WOP struct {
	// insertion point
	Name string
	Path string
	GongModelPath string
	IsEditable bool
	IsReloaded bool
	AbsolutePathToDiagramPackage string
}

type Field_WOP struct {
	// insertion point
	Name string
	Identifier string
	FieldTypeAsString string
	Structname string
	Fieldtypename string
}

type GongEnumShape_WOP struct {
	// insertion point
	Name string
	Identifier string
	Width float64
	Heigth float64
}

type GongEnumValueEntry_WOP struct {
	// insertion point
	Name string
	Identifier string
}

type GongStructShape_WOP struct {
	// insertion point
	Name string
	Identifier string
	ShowNbInstances bool
	NbInstances int
	Width float64
	Height float64
	IsSelected bool
}

type Link_WOP struct {
	// insertion point
	Name string
	Identifier string
	Fieldtypename string
	FieldOffsetX float64
	FieldOffsetY float64
	TargetMultiplicity MultiplicityType
	TargetMultiplicityOffsetX float64
	TargetMultiplicityOffsetY float64
	SourceMultiplicity MultiplicityType
	SourceMultiplicityOffsetX float64
	SourceMultiplicityOffsetY float64
	StartOrientation OrientationType
	StartRatio float64
	EndOrientation OrientationType
	EndRatio float64
	CornerOffsetRatio float64
}

type NoteShape_WOP struct {
	// insertion point
	Name string
	Identifier string
	Body string
	BodyHTML string
	X float64
	Y float64
	Width float64
	Heigth float64
	Matched bool
}

type NoteShapeLink_WOP struct {
	// insertion point
	Name string
	Identifier string
	Type NoteShapeLinkType
}

type Position_WOP struct {
	// insertion point
	X float64
	Y float64
	Name string
}

type UmlState_WOP struct {
	// insertion point
	Name string
	X float64
	Y float64
}

type Umlsc_WOP struct {
	// insertion point
	Name string
	Activestate string
	IsInDrawMode bool
}

type Vertice_WOP struct {
	// insertion point
	X float64
	Y float64
	Name string
}

