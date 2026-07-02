package models

type ControlPoint struct {
	Name string

	// X_Relative is the x coord Relative to the center of the closest Rect. Can vary from §inf to _inf.
	// 0 means ControlPoint is on the Rect center,
	// 1.0 means on the right of the Rect, -1.0 on the left.
	X_Relative,
	Y_Relative float64 // idem for y, but for top or bottom

	ClosestRect *Rect

	Impl ControlPointImplInterface
}

type ControlPointImplInterface interface {

	// ControlPointUpdated function is called each time a ControlPoint is modified
	ControlPointUpdated(updatedControlPoint *ControlPoint)
}

func (controlpoint *ControlPoint) OnAfterUpdate(stage *Stage, _, frontControlPoint *ControlPoint) {

	if controlpoint.Impl != nil {
		controlpoint.Impl.ControlPointUpdated(frontControlPoint)
	}
}
