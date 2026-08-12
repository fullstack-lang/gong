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

	IsHiddenEyeSeatBottomCurveShape bool
	EyeSeatBottomCurveShape         *EyeSeatBottomCurveShape

	IsHiddenEyeStoolBottomCurveShape bool
	EyeStoolBottomCurveShape         *EyeStoolBottomCurveShape

	IsHiddenSeat3DShape bool
	Seat3DShape         *Seat3DShape

	IsHiddenEyeVolume3DShape bool
	EyeVolume3DShape         *EyeVolume3DShape

	IsHiddenSeatAndLegs3DShape bool
	SeatAndLegs3DShape         *SeatAndLegs3DShape

	IsHiddenRotatedSeatAndLegs3DShape bool
	RotatedSeatAndLegs3DShape         *RotatedSeatAndLegs3DShape

	IsHiddenTiledFloor3DShape bool
	TiledFloor3DShape         *TiledFloor3DShape

	Rendered3DShape *Rendered3DShape
}
