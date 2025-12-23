package models

type RectShape struct {
	X, Y, Width, Height float64
	receiver            GongstructIF // the pointer to the owning object
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
}
