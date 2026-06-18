// generated code (do not edit)
package models

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
