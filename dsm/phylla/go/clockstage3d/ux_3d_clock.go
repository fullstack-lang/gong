package clockstage3d

import (
	"github.com/fullstack-lang/gong/dsm/phylla/go/cylinderstage3d"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
)

func (u *Clock3DStageUpdater) ux_3d_clock(stager *models.Stager) {
	clock3dStage := stager.GetClock3dStage()
	if clock3dStage == nil {
		return
	}

	plant := stager.GetCurrentPlant()
	if plant == nil || plant.PlantType != models.Clock || plant.ClockAbstract == nil {
		clock3dStage.Reset()
		clock3dStage.Commit()
		return
	}

	var checkedDiagram *models.PlantDiagram
	for _, d := range plant.PlantDiagrams {
		if d.IsChecked {
			checkedDiagram = d
			break
		}
	}

	params := cylinderstage3d.Cylinder3DParams{
		NamePrefix:            "Clock",
		CanvasName:            "Clock 3D Canvas",
		RadialRepetitions:     plant.ClockAbstract.RadialRepetitions,
		Transparency:          plant.ClockAbstract.Transparency,
		RelativeTubeDiameter:  plant.ClockAbstract.RelativeTubeDiameter,
		RelativeHeight3DTorus: plant.ClockAbstract.RelativeHeight3DTorus,
		VerticalScale:         plant.ClockAbstract.ClockTorusVerticalScale,
		RelativeHeight:        plant.ClockAbstract.RelativeHeight,
		ProjectionAngle:       plant.ClockAbstract.ProjectionAngle,
		HasRotatedShapes:      false,
	}

	if checkedDiagram != nil && checkedDiagram.ClockDiagram != nil {
		params.Rendered3DShape = checkedDiagram.ClockDiagram.Rendered3DShape
		params.IsHiddenTorus3DShape = checkedDiagram.ClockDiagram.IsHiddenTorus3DShape
		params.IsHiddenTopCurveShape = checkedDiagram.ClockDiagram.IsHiddenClockTopCurveShape
		params.IsHiddenSampledPoints3DShape = checkedDiagram.ClockDiagram.IsHiddenSampledPoints3DShape
	}

	base := cylinderstage3d.RenderCylinder3DBase(clock3dStage, stager, plant, params)
	if base != nil && base.Canvas != nil {
		cylinderstage3d.AddFloorTiles(clock3dStage, base.Canvas, base.GlobalR, base.FloorMinY)
	}

	clock3dStage.Commit()
}
