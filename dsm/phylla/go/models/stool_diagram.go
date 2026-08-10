package models

type StoolDiagram struct {
	Name string

	IsHiddenSeatTopCurveShape bool
	SeatTopCurveShape         *SeatTopCurveShape

	IsHiddenRotatedSeatTopCurveShape bool
	RotatedSeatTopCurveShape         *PartiallyRotatedSeatTopCurveShape

	IsHiddenSeatBottomCurveShape bool
	SeatBottomCurveShape         *SeatBottomCurveShape

	IsHiddenRotatedSeatBottomCurveShape bool
	RotatedSeatBottomCurveShape         *PartiallyRotatedSeatBottomCurveShape

	IsHiddenTorus3DShape bool
	Torus3DShape         *Torus3DShape

	IsHiddenRotatedTorusShape bool
	RotatedTorusShape         *PartiallyRotatedTorusShape

	IsHiddenSampledPoints3DShape bool
	SampledPoints3DShape         *SampledPoints3DShape

	IsHiddenRotatedSampledPoints3DShape bool
	RotatedSampledPoints3DShape         *RotatedSampledPoints3DShape

	IsHiddenEyeSampledPoints3DShape bool
	EyeSampledPoints3DShape         *EyeSampledPoints3DShape

	IsHiddenEyeCornersSampledPoints3DShape bool
	EyeCornersSampledPoints3DShape         *EyeCornersSampledPoints3DShape

	IsHiddenEye3DShape bool
	Eye3DShape         *Eye3DShape

	Rendered3DShape *Rendered3DShape
}
