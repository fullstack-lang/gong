// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.ArcNormalVectorShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ArcNormalVectorShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArcNormalVectorShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ArcNormalVectorShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ArcNormalVectorShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArcNormalVectorShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AxesShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "AxesShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AxesShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BaseVectorShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "BaseVectorShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BaseVectorShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BaseVectorShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "BaseVectorShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BaseVectorShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ChosenP1P2PairShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ChosenP1P2PairShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ChosenP1P2PairShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.CircleGridShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "CircleGridShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleGridShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EndArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "EndArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EndArcShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "EndArcShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndArcShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EndHalfwayArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "EndHalfwayArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndHalfwayArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EndHalfwayArcShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "EndHalfwayArcShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndHalfwayArcShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ExplanationTextShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ExplanationTextShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExplanationTextShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GridPathShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "GridPathShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GridPathShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GrowthCurve2D:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "GrowthCurve2D",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurve2DFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GrowthCurve2DRibbon:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "GrowthCurve2DRibbon",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurve2DRibbonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GrowthCurve2DRibbonEndShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "GrowthCurve2DRibbonEndShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurve2DRibbonEndShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GrowthCurve2DRibbonStartShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "GrowthCurve2DRibbonStartShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurve2DRibbonStartShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GrowthCurveRhombusGridShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "GrowthCurveRhombusGridShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurveRhombusGridShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GrowthCurveRhombusShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "GrowthCurveRhombusShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurveRhombusShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GrowthVectorShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "GrowthVectorShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthVectorShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.InitialRhombusGridShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "InitialRhombusGridShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InitialRhombusGridShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.InitialRhombusShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "InitialRhombusShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InitialRhombusShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Library:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "Library",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MidArcVectorShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "MidArcVectorShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MidArcVectorShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MidArcVectorShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "MidArcVectorShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MidArcVectorShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DRibbon:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DRibbon",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DRibbonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DRibbonEndShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DRibbonEndShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DRibbonEndShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DRibbonStartShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DRibbonStartShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DRibbonStartShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DTrajectory:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DTrajectory",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DTrajectoryP1CurveShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DTrajectoryP1P2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DTrajectoryP1P2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1P2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DTrajectoryP1PointShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DTrajectoryP1PointShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DTrajectoryP2CurveShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DTrajectoryP2PointShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DTrajectoryP2PointShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyGrowthCurve2DTrajectoryShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyGrowthCurve2DTrajectoryShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartiallyRotatedTorusShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PartiallyRotatedTorusShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyRotatedTorusShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PerpendicularVector:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PerpendicularVector",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PerpendicularVectorFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PerpendicularVectorGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PerpendicularVectorGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PerpendicularVectorGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PerpendicularVectorGridHalfway:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PerpendicularVectorGridHalfway",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PerpendicularVectorGridHalfwayFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PerpendicularVectorHalfway:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PerpendicularVectorHalfway",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PerpendicularVectorHalfwayFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Plant:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "Plant",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlantFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PlantCircumferenceShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PlantCircumferenceShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlantCircumferenceShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PlantDiagram:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PlantDiagram",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlantDiagramFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PxShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "PxShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PxShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Rendered3DShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "Rendered3DShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Rendered3DShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RhombusShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "RhombusShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RhombusShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RhombusStuff:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "RhombusStuff",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RhombusStuffFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RotatedRhombusGridShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "RotatedRhombusGridShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RotatedRhombusGridShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RotatedRhombusShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "RotatedRhombusShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RotatedRhombusShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ShiftedBottomTopStartArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ShiftedBottomTopStartArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedBottomTopStartArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ShiftedBottomTopStartArcShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ShiftedBottomTopStartArcShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedBottomTopStartArcShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ShiftedLeftStackGrowthCurveEndArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ShiftedLeftStackGrowthCurveEndArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftStackGrowthCurveEndArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ShiftedLeftStackGrowthCurveStartArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ShiftedLeftStackGrowthCurveStartArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftStackGrowthCurveStartArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ShiftedLeftStackNormalVector:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ShiftedLeftStackNormalVector",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftStackNormalVectorFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ShiftedLeftStackOfGrowthCurve:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ShiftedLeftStackOfGrowthCurve",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftStackOfGrowthCurveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ShiftedLeftStackOfNormalVector:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ShiftedLeftStackOfNormalVector",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftStackOfNormalVectorFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ShiftedRightGrowthCurve2DRibbon:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ShiftedRightGrowthCurve2DRibbon",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedRightGrowthCurve2DRibbonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ShiftedRightGrowthCurve2DRibbonEndShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ShiftedRightGrowthCurve2DRibbonEndShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ShiftedRightGrowthCurve2DRibbonStartShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "ShiftedRightGrowthCurve2DRibbonStartShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackGrowthCurve2DEndHalfwayArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackGrowthCurve2DEndHalfwayArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurve2DEndHalfwayArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackGrowthCurve2DRibbonEndShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackGrowthCurve2DRibbonEndShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurve2DRibbonEndShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackGrowthCurve2DRibbonStartShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackGrowthCurve2DRibbonStartShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurve2DRibbonStartShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackGrowthCurve2DStartHalfwayArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackGrowthCurve2DStartHalfwayArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurve2DStartHalfwayArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackOfGrowthCurve2D:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackOfGrowthCurve2D",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfGrowthCurve2DFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackOfGrowthCurve2DRibbon:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackOfGrowthCurve2DRibbon",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfGrowthCurve2DRibbonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackOfPartiallyRotatedTorusShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackOfPartiallyRotatedTorusShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfPartiallyRotatedTorusShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackOfRotatedGrowthCurve2D:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackOfRotatedGrowthCurve2D",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfRotatedGrowthCurve2DFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackOfRotatedGrowthCurve2DRibbon:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackOfRotatedGrowthCurve2DRibbon",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfRotatedGrowthCurve2DRibbonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackRotatedGrowthCurve2DEndArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackRotatedGrowthCurve2DEndArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DEndArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackRotatedGrowthCurve2DRibbonEndShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackRotatedGrowthCurve2DRibbonEndShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DRibbonEndShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackRotatedGrowthCurve2DRibbonStartShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackRotatedGrowthCurve2DRibbonStartShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DRibbonStartShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackRotatedGrowthCurve2DStartArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackRotatedGrowthCurve2DStartArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DStartArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StartArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StartArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StartArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StartArcShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StartArcShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StartArcShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StartHalfwayArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StartHalfwayArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StartHalfwayArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StartHalfwayArcShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StartHalfwayArcShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StartHalfwayArcShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopEndArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopEndArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopEndArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopEndArcShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopEndArcShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopEndArcShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopEndHalfwayArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopEndHalfwayArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopEndHalfwayArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopEndHalfwayArcShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopEndHalfwayArcShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopEndHalfwayArcShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopGrowthCurve2D:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopGrowthCurve2D",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopGrowthCurve2DFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopMidArcVectorShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopMidArcVectorShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopMidArcVectorShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopMidArcVectorShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopMidArcVectorShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopMidArcVectorShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStackGrowthCurve2DEndHalfwayArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStackGrowthCurve2DEndHalfwayArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStackGrowthCurve2DStartHalfwayArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStackGrowthCurve2DStartHalfwayArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStackOfGrowthCurve2D:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStackOfGrowthCurve2D",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackOfGrowthCurve2DFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStackOfRotatedGrowthCurve2D:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStackOfRotatedGrowthCurve2D",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackOfRotatedGrowthCurve2DFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStackOfRotatedGrowthCurve2DEndArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStackOfRotatedGrowthCurve2DEndArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStackOfRotatedGrowthCurve2DStartArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStackOfRotatedGrowthCurve2DStartArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStartArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStartArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStartArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStartArcShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStartArcShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStartArcShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStartHalfwayArcShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStartHalfwayArcShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStartHalfwayArcShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStartHalfwayArcShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStartHalfwayArcShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStartHalfwayArcShapeGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TorusStackShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TorusStackShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TorusStackShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.VerticalTorusStackShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "VerticalTorusStackShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VerticalTorusStackShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
