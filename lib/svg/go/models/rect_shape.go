package models

type RectShape struct {
	X, Y, Width, Height float64
	IsHidden            bool
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

func (s *RectShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *RectShape) GetIsHidden() bool {
	return s.IsHidden
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

	SetIsHidden(isHidden bool)
	GetIsHidden() bool
}
