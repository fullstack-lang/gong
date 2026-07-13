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
	case *models.BottomEndArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "BottomEndArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BottomEndArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BottomEndArcShapeV2Grid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "BottomEndArcShapeV2Grid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BottomEndArcShapeV2GridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BottomStackGrowthCurveEndArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "BottomStackGrowthCurveEndArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BottomStackGrowthCurveEndArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BottomStackGrowthCurveStartArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "BottomStackGrowthCurveStartArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BottomStackGrowthCurveStartArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BottomStackOfGrowthCurveV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "BottomStackOfGrowthCurveV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BottomStackOfGrowthCurveV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BottomStartArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "BottomStartArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BottomStartArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BottomStartArcShapeV2Grid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "BottomStartArcShapeV2Grid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BottomStartArcShapeV2GridFormCallback(
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
	case *models.EndArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "EndArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EndArcShapeV2Grid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "EndArcShapeV2Grid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndArcShapeV2GridFormCallback(
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
	case *models.GrowthCurveBezierShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "GrowthCurveBezierShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurveBezierShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GrowthCurveBezierShapeGrid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "GrowthCurveBezierShapeGrid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurveBezierShapeGridFormCallback(
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
	case *models.NextCircleShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "NextCircleShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NextCircleShapeFormCallback(
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
	case *models.StackGrowthCurveBezierShape:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackGrowthCurveBezierShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurveBezierShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackGrowthCurveEndArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackGrowthCurveEndArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurveEndArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackGrowthCurveStartArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackGrowthCurveStartArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurveStartArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackOfGrowthCurve:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackOfGrowthCurve",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfGrowthCurveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StackOfGrowthCurveV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StackOfGrowthCurveV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfGrowthCurveV2FormCallback(
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
	case *models.StartArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StartArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StartArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StartArcShapeV2Grid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "StartArcShapeV2Grid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StartArcShapeV2GridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopEndArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopEndArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopEndArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopEndArcShapeV2Grid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopEndArcShapeV2Grid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopEndArcShapeV2GridFormCallback(
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
	case *models.TopStackGrowthCurveEndArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStackGrowthCurveEndArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackGrowthCurveEndArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStackGrowthCurveStartArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStackGrowthCurveStartArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackGrowthCurveStartArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStackOfGrowthCurveV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStackOfGrowthCurveV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackOfGrowthCurveV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStartArcShapeV2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStartArcShapeV2",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStartArcShapeV2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TopStartArcShapeV2Grid:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "TopStartArcShapeV2Grid",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStartArcShapeV2GridFormCallback(
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
