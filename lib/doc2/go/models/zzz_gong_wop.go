// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type AttributeShape_WOP struct {
	// insertion point
	Name string
	Identifier string
	FieldTypeAsString string
	Structname string
	Fieldtypename string
}

func (from *AttributeShape) CopyBasicFields(to *AttributeShape) {
	// insertion point
	to.Name = from.Name
	to.Identifier = from.Identifier
	to.FieldTypeAsString = from.FieldTypeAsString
	to.Structname = from.Structname
	to.Fieldtypename = from.Fieldtypename
}

type Classdiagram_WOP struct {
	// insertion point
	Name string
	IsInRenameMode bool
	IsExpanded bool
	NodeGongStructsIsExpanded bool
	NodeGongStructNodeExpansionBinaryEncoding int
	NodeGongEnumsIsExpanded bool
	NodeGongEnumNodeExpansionBinaryEncoding int
	NodeGongNotesIsExpanded bool
	NodeGongNoteNodeExpansionBinaryEncoding int
}

func (from *Classdiagram) CopyBasicFields(to *Classdiagram) {
	// insertion point
	to.Name = from.Name
	to.IsInRenameMode = from.IsInRenameMode
	to.IsExpanded = from.IsExpanded
	to.NodeGongStructsIsExpanded = from.NodeGongStructsIsExpanded
	to.NodeGongStructNodeExpansionBinaryEncoding = from.NodeGongStructNodeExpansionBinaryEncoding
	to.NodeGongEnumsIsExpanded = from.NodeGongEnumsIsExpanded
	to.NodeGongEnumNodeExpansionBinaryEncoding = from.NodeGongEnumNodeExpansionBinaryEncoding
	to.NodeGongNotesIsExpanded = from.NodeGongNotesIsExpanded
	to.NodeGongNoteNodeExpansionBinaryEncoding = from.NodeGongNoteNodeExpansionBinaryEncoding
}

type DiagramPackage_WOP struct {
	// insertion point
	Name string
	Path string
	GongModelPath string
	AbsolutePathToDiagramPackage string
}

func (from *DiagramPackage) CopyBasicFields(to *DiagramPackage) {
	// insertion point
	to.Name = from.Name
	to.Path = from.Path
	to.GongModelPath = from.GongModelPath
	to.AbsolutePathToDiagramPackage = from.AbsolutePathToDiagramPackage
}

type GongEnumShape_WOP struct {
	// insertion point
	Name string
	X float64
	Y float64
	Identifier string
	Width float64
	Height float64
	IsExpanded bool
}

func (from *GongEnumShape) CopyBasicFields(to *GongEnumShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Identifier = from.Identifier
	to.Width = from.Width
	to.Height = from.Height
	to.IsExpanded = from.IsExpanded
}

type GongEnumValueShape_WOP struct {
	// insertion point
	Name string
	Identifier string
}

func (from *GongEnumValueShape) CopyBasicFields(to *GongEnumValueShape) {
	// insertion point
	to.Name = from.Name
	to.Identifier = from.Identifier
}

type GongNoteLinkShape_WOP struct {
	// insertion point
	Name string
	Identifier string
	Type NoteShapeLinkType
}

func (from *GongNoteLinkShape) CopyBasicFields(to *GongNoteLinkShape) {
	// insertion point
	to.Name = from.Name
	to.Identifier = from.Identifier
	to.Type = from.Type
}

type GongNoteShape_WOP struct {
	// insertion point
	Name string
	Identifier string
	Body string
	BodyHTML string
	X float64
	Y float64
	Width float64
	Height float64
	Matched bool
	IsExpanded bool
}

func (from *GongNoteShape) CopyBasicFields(to *GongNoteShape) {
	// insertion point
	to.Name = from.Name
	to.Identifier = from.Identifier
	to.Body = from.Body
	to.BodyHTML = from.BodyHTML
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.Matched = from.Matched
	to.IsExpanded = from.IsExpanded
}

type GongStructShape_WOP struct {
	// insertion point
	Name string
	X float64
	Y float64
	Identifier string
	ShowNbInstances bool
	NbInstances int
	Width float64
	Height float64
	IsSelected bool
}

func (from *GongStructShape) CopyBasicFields(to *GongStructShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Identifier = from.Identifier
	to.ShowNbInstances = from.ShowNbInstances
	to.NbInstances = from.NbInstances
	to.Width = from.Width
	to.Height = from.Height
	to.IsSelected = from.IsSelected
}

type LinkShape_WOP struct {
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
	X float64
	Y float64
	StartOrientation OrientationType
	StartRatio float64
	EndOrientation OrientationType
	EndRatio float64
	CornerOffsetRatio float64
}

func (from *LinkShape) CopyBasicFields(to *LinkShape) {
	// insertion point
	to.Name = from.Name
	to.Identifier = from.Identifier
	to.Fieldtypename = from.Fieldtypename
	to.FieldOffsetX = from.FieldOffsetX
	to.FieldOffsetY = from.FieldOffsetY
	to.TargetMultiplicity = from.TargetMultiplicity
	to.TargetMultiplicityOffsetX = from.TargetMultiplicityOffsetX
	to.TargetMultiplicityOffsetY = from.TargetMultiplicityOffsetY
	to.SourceMultiplicity = from.SourceMultiplicity
	to.SourceMultiplicityOffsetX = from.SourceMultiplicityOffsetX
	to.SourceMultiplicityOffsetY = from.SourceMultiplicityOffsetY
	to.X = from.X
	to.Y = from.Y
	to.StartOrientation = from.StartOrientation
	to.StartRatio = from.StartRatio
	to.EndOrientation = from.EndOrientation
	to.EndRatio = from.EndRatio
	to.CornerOffsetRatio = from.CornerOffsetRatio
}

