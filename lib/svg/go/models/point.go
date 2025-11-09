package models

type Point struct {
	Name string
	X, Y float64

	Impl PointImplInterface
}

type PointImplInterface interface {

	// PointUpdated function is called each time a Point is modified
	PointUpdated(updatedPoint *Point)
}

func (point *Point) OnAfterUpdate(stage *Stage, _, frontPoint *Point) {

	if point.Impl != nil {
		point.Impl.PointUpdated(frontPoint)
	}
}
