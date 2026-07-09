// generated code (do not edit)
package models

type AbstractType interface {
	GongstructIF
	GetIsExpanded() bool
	SetIsExpanded(bool)
	GetComputedPrefix() string
	SetComputedPrefix(string)
	GetIsInRenameMode() bool
	SetIsInRenameMode(bool)
	SetComputedPrefixInt([]int)
}

type TreeAbstractType interface {
	AbstractType

	GetComputedWidth() int
	SetComputedWidth(int)

	GetLayoutDirection() LayoutDirection
	SetLayoutDirection(LayoutDirection)
}

type LayoutDirection int

const (
	Vertical LayoutDirection = iota
	Horizontal
)

type AbstractTypeFields struct {
	// ComputedPrefix is automaticaly computed by the semantic enforcing mechanism
	ComputedPrefix string
	computedPrefix []int

	// nodes can be edited
	isInRenameMode bool
	IsExpanded     bool // to be made private once in production (no need to persist)
}

type TreeAbstractTypeFields struct {
	// When the full PBS is displayed, the computedWidth is the number of node
	// aligned below. A leaf node has a computedWidth of 1
	computedWidth int

	// Directive for display in the concrete diagram
	LayoutDirection LayoutDirection
}

func (r *TreeAbstractTypeFields) GetLayoutDirection() LayoutDirection {
	return r.LayoutDirection
}

func (r *TreeAbstractTypeFields) SetLayoutDirection(d LayoutDirection) {
	r.LayoutDirection = d
}

func (r *TreeAbstractTypeFields) GetComputedWidth() int {
	return r.computedWidth
}

func (r *TreeAbstractTypeFields) SetComputedWidth(w int) {
	r.computedWidth = w
}



func (r *AbstractTypeFields) SetComputedPrefixInt(p []int) {
	r.computedPrefix = p
}

func (r *AbstractTypeFields) GetIsExpanded() bool {
	return r.IsExpanded
}

func (r *AbstractTypeFields) SetIsExpanded(isExpanded bool) {
	r.IsExpanded = isExpanded
}

func (r *AbstractTypeFields) GetComputedPrefix() string {
	return r.ComputedPrefix
}

func (r *AbstractTypeFields) SetComputedPrefix(ComputedPrefix string) {
	r.ComputedPrefix = ComputedPrefix
}

func (r *AbstractTypeFields) GetIsInRenameMode() bool {
	return r.isInRenameMode
}

func (r *AbstractTypeFields) SetIsInRenameMode(isInRenameMode bool) {
	r.isInRenameMode = isInRenameMode
}

type RectShape struct {
	X, Y, Width, Height float64
	receiver            GongstructIF // the pointer to the owning object
	IsHidden            bool
}

func (s *RectShape) Stage(stage *Stage) {
	s.receiver.StageVoid(stage)
}

func (s *RectShape) StageVoid(stage *Stage) {
	s.receiver.StageVoid(stage)
}

func (s *RectShape) SetHeight(height float64) {
	s.Height = height
}

func (s *RectShape) SetWidth(width float64) {
	s.Width = width
}

func (s *RectShape) SetX(x float64) {
	s.X = x
}

func (s *RectShape) SetY(y float64) {
	s.Y = y
}

func (s *RectShape) GetX() float64 {
	return s.X
}

func (s *RectShape) GetY() float64 {
	return s.Y
}

func (s *RectShape) GetWidth() float64 {
	return s.Width
}

func (s *RectShape) GetHeight() float64 {
	return s.Height
}

func (s *RectShape) SetReceiver(receiver GongstructIF) {
	s.receiver = receiver
}

func (s *RectShape) GetReceiver() GongstructIF {
	return s.receiver
}

func (s *RectShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *RectShape) GetIsHidden() bool {
	return s.IsHidden
}

type RectShapeInterface interface {
	SetX(x float64)
	SetY(x float64)
	SetWidth(width float64)
	SetHeight(height float64)

	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64

	StageVoid(stage *Stage)

	SetReceiver(receiver GongstructIF)
	GetReceiver() GongstructIF

	SetIsHidden(isHidden bool)
	GetIsHidden() bool
}

type OrientationType string

// values for EnumType
const (
	ORIENTATION_HORIZONTAL OrientationType = "ORIENTATION_HORIZONTAL"
	ORIENTATION_VERTICAL   OrientationType = "ORIENTATION_VERTICAL"
)

type LinkShape struct {
	StartRatio float64
	EndRatio   float64

	StartOrientation OrientationType
	EndOrientation   OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

// Setter functions for StateTransitionShape
func (s *LinkShape) SetStartRatio(ratio float64) {
	s.StartRatio = ratio
}

func (s *LinkShape) SetEndRatio(ratio float64) {
	s.EndRatio = ratio
}

func (s *LinkShape) SetStartOrientation(orientation OrientationType) {
	s.StartOrientation = orientation
}

func (s *LinkShape) SetEndOrientation(orientation OrientationType) {
	s.EndOrientation = orientation
}

func (s *LinkShape) SetCornerOffsetRatio(ratio float64) {
	s.CornerOffsetRatio = ratio
}

func (s *LinkShape) GetStartRatio() float64 {
	return s.StartRatio
}

func (s *LinkShape) GetEndRatio() float64 {
	return s.EndRatio
}

func (s *LinkShape) GetStartOrientation() OrientationType {
	return s.StartOrientation
}

func (s *LinkShape) GetEndOrientation() OrientationType {
	return s.EndOrientation
}

func (s *LinkShape) GetCornerOffsetRatio() float64 {
	return s.CornerOffsetRatio
}

func (s *LinkShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *LinkShape) GetIsHidden() bool {
	return s.IsHidden
}

// Interface for StateTransitionShape
type LinkShapeInterface interface {
	SetStartRatio(ratio float64)
	SetEndRatio(ratio float64)
	SetStartOrientation(orientation OrientationType)
	SetEndOrientation(orientation OrientationType)
	SetCornerOffsetRatio(ratio float64)

	GetStartRatio() float64
	GetEndRatio() float64
	GetStartOrientation() OrientationType
	GetEndOrientation() OrientationType
	GetCornerOffsetRatio() float64

	SetIsHidden(isHidden bool)
	GetIsHidden() bool
}

type ConcreteType interface {
	GongstructIF
	GetAbstractElement() AbstractType
	SetAbstractElement(AbstractType)
}

type LayoutConcreteType interface {
	RectShapeInterface
	ConcreteType
	GetConcreteLayoutDirection() LayoutDirection
	SetConcreteLayoutDirection(LayoutDirection)
	GetOverideLayoutDirection() bool
	SetOverideLayoutDirection(bool)
}

type ConcreteTypeFields struct {
	OverideLayoutDirection bool
	LayoutDirection        LayoutDirection
}

func (c *ConcreteTypeFields) GetConcreteLayoutDirection() LayoutDirection {
	return c.LayoutDirection
}

func (c *ConcreteTypeFields) SetConcreteLayoutDirection(d LayoutDirection) {
	c.LayoutDirection = d
}

func (c *ConcreteTypeFields) GetOverideLayoutDirection() bool {
	return c.OverideLayoutDirection
}

func (c *ConcreteTypeFields) SetOverideLayoutDirection(b bool) {
	c.OverideLayoutDirection = b
}

type AssociationConcreteType interface {
	GongstructIF
	GetAbstractStartElement() AbstractType
	SetAbstractStartElement(AbstractType)
	GetAbstractEndElement() AbstractType
	SetAbstractEndElement(AbstractType)
}

type LayoutAssociationType interface {
	LinkShapeInterface
	AssociationConcreteType
}

type AssociationConcreteType2[SourceAT AbstractType, TargetAT AbstractType] interface {
	GongstructIF
	GetAbstractStartElement() SourceAT
	SetAbstractStartElement(SourceAT)
	GetAbstractEndElement() TargetAT
	SetAbstractEndElement(TargetAT)
}

type DiagramIF interface {
	GetName() string
	GetIsChecked() bool
	SetIsChecked(bool)
	SetIsShowPrefix(bool)
	GetIsShowPrefix() bool

	GetDefaultBoxWidth() float64
	GetDefaultBoxHeigth() float64
	IsEditable() bool

	// within the tree branch of one diagram, when an element is present in more than one diagram,
	// it is possible to access it via a list. Only one element have a list that is available per diagram.
	GetDiagramListElement() AbstractType
	SetDiagramListElement(AbstractType)
}
