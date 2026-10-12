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

func (d LayoutDirection) String() string {
	if d == Vertical {
		return "Vertical"
	}
	return "Horizontal"
}

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

	IsPresentationOverridden bool
	Color                    string
	FillOpacity              float64
	Stroke                   string
	StrokeWidth              float64
	StrokeOpacity            float64
	StrokeDashArray          string
	RX                       float64

	FontSize   string
	FontWeight string
	FontColor  string
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

func (s *RectShape) GetIsPresentationOverridden() bool {
	return s.IsPresentationOverridden
}

func (s *RectShape) SetIsPresentationOverridden(b bool) {
	s.IsPresentationOverridden = b
}

func (s *RectShape) GetColor() string {
	return s.Color
}

func (s *RectShape) SetColor(c string) {
	s.Color = c
}

func (s *RectShape) GetFillOpacity() float64 {
	return s.FillOpacity
}

func (s *RectShape) SetFillOpacity(fo float64) {
	s.FillOpacity = fo
}

func (s *RectShape) GetStroke() string {
	return s.Stroke
}

func (s *RectShape) SetStroke(st string) {
	s.Stroke = st
}

func (s *RectShape) GetStrokeWidth() float64 {
	return s.StrokeWidth
}

func (s *RectShape) SetStrokeWidth(w float64) {
	s.StrokeWidth = w
}

func (s *RectShape) GetStrokeOpacity() float64 {
	return s.StrokeOpacity
}

func (s *RectShape) SetStrokeOpacity(so float64) {
	s.StrokeOpacity = so
}

func (s *RectShape) GetStrokeDashArray() string {
	return s.StrokeDashArray
}

func (s *RectShape) SetStrokeDashArray(sda string) {
	s.StrokeDashArray = sda
}

func (s *RectShape) GetRX() float64 {
	return s.RX
}

func (s *RectShape) SetRX(rx float64) {
	s.RX = rx
}

func (s *RectShape) GetFontSize() string {
	return s.FontSize
}

func (s *RectShape) SetFontSize(fs string) {
	s.FontSize = fs
}

func (s *RectShape) GetFontWeight() string {
	return s.FontWeight
}

func (s *RectShape) SetFontWeight(fw string) {
	s.FontWeight = fw
}

func (s *RectShape) GetFontColor() string {
	return s.FontColor
}

func (s *RectShape) SetFontColor(fc string) {
	s.FontColor = fc
}

type RectShapeInterface interface {
	SetX(x float64)
	SetY(y float64)
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

	GetIsPresentationOverridden() bool
	SetIsPresentationOverridden(bool)

	GetColor() string
	SetColor(string)

	GetFillOpacity() float64
	SetFillOpacity(float64)

	GetStroke() string
	SetStroke(string)

	GetStrokeWidth() float64
	SetStrokeWidth(float64)

	GetStrokeOpacity() float64
	SetStrokeOpacity(float64)

	GetStrokeDashArray() string
	SetStrokeDashArray(string)

	GetRX() float64
	SetRX(float64)

	GetFontSize() string
	SetFontSize(string)

	GetFontWeight() string
	SetFontWeight(string)

	GetFontColor() string
	SetFontColor(string)
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

	IsPresentationOverridden bool
	Stroke                   string
	StrokeWidth              float64
	StrokeOpacity            float64
	StrokeDashArray          string

	FontSize  string
	FontColor string
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

func (s *LinkShape) GetIsPresentationOverridden() bool {
	return s.IsPresentationOverridden
}

func (s *LinkShape) SetIsPresentationOverridden(b bool) {
	s.IsPresentationOverridden = b
}

func (s *LinkShape) GetStroke() string {
	return s.Stroke
}

func (s *LinkShape) SetStroke(st string) {
	s.Stroke = st
}

func (s *LinkShape) GetStrokeWidth() float64 {
	return s.StrokeWidth
}

func (s *LinkShape) SetStrokeWidth(w float64) {
	s.StrokeWidth = w
}

func (s *LinkShape) GetStrokeOpacity() float64 {
	return s.StrokeOpacity
}

func (s *LinkShape) SetStrokeOpacity(so float64) {
	s.StrokeOpacity = so
}

func (s *LinkShape) GetStrokeDashArray() string {
	return s.StrokeDashArray
}

func (s *LinkShape) SetStrokeDashArray(sda string) {
	s.StrokeDashArray = sda
}

func (s *LinkShape) GetFontSize() string {
	return s.FontSize
}

func (s *LinkShape) SetFontSize(fs string) {
	s.FontSize = fs
}

func (s *LinkShape) GetFontColor() string {
	return s.FontColor
}

func (s *LinkShape) SetFontColor(fc string) {
	s.FontColor = fc
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

	GetIsPresentationOverridden() bool
	SetIsPresentationOverridden(bool)

	GetStroke() string
	SetStroke(string)

	GetStrokeWidth() float64
	SetStrokeWidth(float64)

	GetStrokeOpacity() float64
	SetStrokeOpacity(float64)

	GetStrokeDashArray() string
	SetStrokeDashArray(string)

	GetFontSize() string
	SetFontSize(string)

	GetFontColor() string
	SetFontColor(string)
}

type ConcreteType interface {
	GongstructIF
	GetAbstractElement() AbstractType
	SetAbstractElement(AbstractType)
}

type LayoutConcreteType interface {
	RectShapeInterface
	ConcreteType
	GetIsLayoutDirectionDifferent() bool
	SetIsLayoutDirectionDifferent(bool)
}

type ConcreteTypeFields struct {
	IsLayoutDirectionDifferent bool
}

func (c *ConcreteTypeFields) GetIsLayoutDirectionDifferent() bool {
	return c.IsLayoutDirectionDifferent
}

func (c *ConcreteTypeFields) SetIsLayoutDirectionDifferent(b bool) {
	c.IsLayoutDirectionDifferent = b
}

func GetConcreteLayoutDirection(shape LayoutConcreteType) LayoutDirection {
	abstractLayoutDirection := Vertical
	if shape.GetAbstractElement() != nil {
		if treeNode, ok := shape.GetAbstractElement().(TreeAbstractType); ok {
			abstractLayoutDirection = treeNode.GetLayoutDirection()
		}
	}
	if shape.GetIsLayoutDirectionDifferent() {
		if abstractLayoutDirection == Vertical {
			return Horizontal
		}
		return Vertical
	}
	return abstractLayoutDirection
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

	GetIsInAutoLayoutMode() bool
	SetIsInAutoLayoutMode(bool)

	GetDefaultBoxWidth() float64
	GetDefaultBoxHeigth() float64
	IsEditable() bool

	// within the tree branch of one diagram, when an element is present in more than one diagram,
	// it is possible to access it via a list. Only one element have a list that is available per diagram.
	GetDiagramListElement() AbstractType
	SetDiagramListElement(AbstractType)
}
