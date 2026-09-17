// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type All_WOP struct {
	// insertion point

	Name string

	OuterElementName string

	Order int

	Depth int

	MinOccurs string

	MaxOccurs string
}

func (from *All) GongCopyBasicFields(to *All) {
	// insertion point
	to.Name = from.Name
	to.OuterElementName = from.OuterElementName
	to.Order = from.Order
	to.Depth = from.Depth
	to.MinOccurs = from.MinOccurs
	to.MaxOccurs = from.MaxOccurs
}

type Annotation_WOP struct {
	// insertion point

	Name string
}

func (from *Annotation) GongCopyBasicFields(to *Annotation) {
	// insertion point
	to.Name = from.Name
}

type Attribute_WOP struct {
	// insertion point

	Name string

	NameXSD string

	Type string

	HasNameConflict bool

	GoIdentifier string

	Default string

	Use string

	Form string

	Fixed string

	Ref string

	TargetNamespace string

	SimpleType string

	IDXSD string
}

func (from *Attribute) GongCopyBasicFields(to *Attribute) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
	to.Type = from.Type
	to.HasNameConflict = from.HasNameConflict
	to.GoIdentifier = from.GoIdentifier
	to.Default = from.Default
	to.Use = from.Use
	to.Form = from.Form
	to.Fixed = from.Fixed
	to.Ref = from.Ref
	to.TargetNamespace = from.TargetNamespace
	to.SimpleType = from.SimpleType
	to.IDXSD = from.IDXSD
}

type AttributeGroup_WOP struct {
	// insertion point

	Name string

	NameXSD string

	HasNameConflict bool

	GoIdentifier string

	Ref string

	Order int

	Depth int
}

func (from *AttributeGroup) GongCopyBasicFields(to *AttributeGroup) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
	to.HasNameConflict = from.HasNameConflict
	to.GoIdentifier = from.GoIdentifier
	to.Ref = from.Ref
	to.Order = from.Order
	to.Depth = from.Depth
}

type Choice_WOP struct {
	// insertion point

	Name string

	OuterElementName string

	Order int

	Depth int

	MinOccurs string

	MaxOccurs string

	IsDuplicatedInXSD bool
}

func (from *Choice) GongCopyBasicFields(to *Choice) {
	// insertion point
	to.Name = from.Name
	to.OuterElementName = from.OuterElementName
	to.Order = from.Order
	to.Depth = from.Depth
	to.MinOccurs = from.MinOccurs
	to.MaxOccurs = from.MaxOccurs
	to.IsDuplicatedInXSD = from.IsDuplicatedInXSD
}

type ComplexContent_WOP struct {
	// insertion point

	Name string
}

func (from *ComplexContent) GongCopyBasicFields(to *ComplexContent) {
	// insertion point
	to.Name = from.Name
}

type ComplexType_WOP struct {
	// insertion point

	Name string

	HasNameConflict bool

	GoIdentifier string

	IsAnonymous bool

	NameXSD string

	OuterElementName string

	Order int

	Depth int

	MinOccurs string

	MaxOccurs string

	IsDuplicatedInXSD bool
}

func (from *ComplexType) GongCopyBasicFields(to *ComplexType) {
	// insertion point
	to.Name = from.Name
	to.HasNameConflict = from.HasNameConflict
	to.GoIdentifier = from.GoIdentifier
	to.IsAnonymous = from.IsAnonymous
	to.NameXSD = from.NameXSD
	to.OuterElementName = from.OuterElementName
	to.Order = from.Order
	to.Depth = from.Depth
	to.MinOccurs = from.MinOccurs
	to.MaxOccurs = from.MaxOccurs
	to.IsDuplicatedInXSD = from.IsDuplicatedInXSD
}

type Documentation_WOP struct {
	// insertion point

	Name string

	Text string

	Source string

	Lang string
}

func (from *Documentation) GongCopyBasicFields(to *Documentation) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
	to.Source = from.Source
	to.Lang = from.Lang
}

type Element_WOP struct {
	// insertion point

	Name string

	Order int

	Depth int

	HasNameConflict bool

	GoIdentifier string

	NameXSD string

	Type string

	MinOccurs string

	MaxOccurs string

	Default string

	Fixed string

	Nillable string

	Ref string

	Abstract string

	Form string

	Block string

	Final string

	IsDuplicatedInXSD bool
}

func (from *Element) GongCopyBasicFields(to *Element) {
	// insertion point
	to.Name = from.Name
	to.Order = from.Order
	to.Depth = from.Depth
	to.HasNameConflict = from.HasNameConflict
	to.GoIdentifier = from.GoIdentifier
	to.NameXSD = from.NameXSD
	to.Type = from.Type
	to.MinOccurs = from.MinOccurs
	to.MaxOccurs = from.MaxOccurs
	to.Default = from.Default
	to.Fixed = from.Fixed
	to.Nillable = from.Nillable
	to.Ref = from.Ref
	to.Abstract = from.Abstract
	to.Form = from.Form
	to.Block = from.Block
	to.Final = from.Final
	to.IsDuplicatedInXSD = from.IsDuplicatedInXSD
}

type Enumeration_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *Enumeration) GongCopyBasicFields(to *Enumeration) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type Extension_WOP struct {
	// insertion point

	Name string

	OuterElementName string

	Order int

	Depth int

	MinOccurs string

	MaxOccurs string

	Base string

	Ref string
}

func (from *Extension) GongCopyBasicFields(to *Extension) {
	// insertion point
	to.Name = from.Name
	to.OuterElementName = from.OuterElementName
	to.Order = from.Order
	to.Depth = from.Depth
	to.MinOccurs = from.MinOccurs
	to.MaxOccurs = from.MaxOccurs
	to.Base = from.Base
	to.Ref = from.Ref
}

type Group_WOP struct {
	// insertion point

	Name string

	NameXSD string

	Ref string

	IsAnonymous bool

	HasNameConflict bool

	GoIdentifier string

	OuterElementName string

	Order int

	Depth int

	MinOccurs string

	MaxOccurs string
}

func (from *Group) GongCopyBasicFields(to *Group) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
	to.Ref = from.Ref
	to.IsAnonymous = from.IsAnonymous
	to.HasNameConflict = from.HasNameConflict
	to.GoIdentifier = from.GoIdentifier
	to.OuterElementName = from.OuterElementName
	to.Order = from.Order
	to.Depth = from.Depth
	to.MinOccurs = from.MinOccurs
	to.MaxOccurs = from.MaxOccurs
}

type Length_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *Length) GongCopyBasicFields(to *Length) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type MaxInclusive_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *MaxInclusive) GongCopyBasicFields(to *MaxInclusive) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type MaxLength_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *MaxLength) GongCopyBasicFields(to *MaxLength) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type MinInclusive_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *MinInclusive) GongCopyBasicFields(to *MinInclusive) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type MinLength_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *MinLength) GongCopyBasicFields(to *MinLength) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type Pattern_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *Pattern) GongCopyBasicFields(to *Pattern) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type Restriction_WOP struct {
	// insertion point

	Name string

	Base string
}

func (from *Restriction) GongCopyBasicFields(to *Restriction) {
	// insertion point
	to.Name = from.Name
	to.Base = from.Base
}

type Schema_WOP struct {
	// insertion point

	Name string

	Xs string

	Order int

	Depth int
}

func (from *Schema) GongCopyBasicFields(to *Schema) {
	// insertion point
	to.Name = from.Name
	to.Xs = from.Xs
	to.Order = from.Order
	to.Depth = from.Depth
}

type Sequence_WOP struct {
	// insertion point

	Name string

	OuterElementName string

	Order int

	Depth int

	MinOccurs string

	MaxOccurs string
}

func (from *Sequence) GongCopyBasicFields(to *Sequence) {
	// insertion point
	to.Name = from.Name
	to.OuterElementName = from.OuterElementName
	to.Order = from.Order
	to.Depth = from.Depth
	to.MinOccurs = from.MinOccurs
	to.MaxOccurs = from.MaxOccurs
}

type SimpleContent_WOP struct {
	// insertion point

	Name string
}

func (from *SimpleContent) GongCopyBasicFields(to *SimpleContent) {
	// insertion point
	to.Name = from.Name
}

type SimpleType_WOP struct {
	// insertion point

	Name string

	NameXSD string

	Order int

	Depth int
}

func (from *SimpleType) GongCopyBasicFields(to *SimpleType) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
	to.Order = from.Order
	to.Depth = from.Depth
}

type TotalDigit_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *TotalDigit) GongCopyBasicFields(to *TotalDigit) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type Union_WOP struct {
	// insertion point

	Name string

	MemberTypes string
}

func (from *Union) GongCopyBasicFields(to *Union) {
	// insertion point
	to.Name = from.Name
	to.MemberTypes = from.MemberTypes
}

type WhiteSpace_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *WhiteSpace) GongCopyBasicFields(to *WhiteSpace) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

// end of insertion point
