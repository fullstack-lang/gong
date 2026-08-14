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

	var checkedDiagram *models.Clock3DDiagram
	for _, d := range plant.Clock3DDiagrams {
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

	if checkedDiagram != nil && checkedDiagram != nil {
		params.Rendered3DShape = checkedDiagram.Rendered3DShape
		params.IsHiddenTorus3DShape = checkedDiagram.IsHiddenTorus3DShape
		params.IsHiddenTopCurveShape = checkedDiagram.IsHiddenClockTopCurveShape
		params.IsHiddenSampledPoints3DShape = checkedDiagram.IsHiddenSampledPoints3DShape
	}

	base := cylinderstage3d.RenderCylinder3DBase(clock3dStage, stager, plant, params)
	if base != nil && base.Canvas != nil {
		if checkedDiagram == nil || checkedDiagram == nil || !checkedDiagram.IsHiddenTiledFloor3DShape {
			cylinderstage3d.AddFloorTiles(clock3dStage, base.Canvas, base.GlobalR, base.FloorMinY)
		}
	}

	clock3dStage.Commit()
}
