// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		if onSave, ok := formGroup.OnSave.(FormCallbackIF); ok {
			if onSave.GetCreationMode() {
				FillUpFormFromGongstructName(probe, onSave.GetGongstructName(), true)
			} else {
				FillUpFormFromGongstruct(onSave.GetInstance(), probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Angle0Shape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Angle0Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Angle0ShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		angle0shape := new(models.Angle0Shape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(angle0shape, formGroup, probe)
	case "ArcNormalVectorShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ArcNormalVectorShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArcNormalVectorShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		arcnormalvectorshape := new(models.ArcNormalVectorShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(arcnormalvectorshape, formGroup, probe)
	case "ArcNormalVectorShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ArcNormalVectorShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArcNormalVectorShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		arcnormalvectorshapegrid := new(models.ArcNormalVectorShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(arcnormalvectorshapegrid, formGroup, probe)
	case "AxesShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AxesShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AxesShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		axesshape := new(models.AxesShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(axesshape, formGroup, probe)
	case "BaseVectorShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "BaseVectorShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BaseVectorShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		basevectorshape := new(models.BaseVectorShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(basevectorshape, formGroup, probe)
	case "BaseVectorShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "BaseVectorShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BaseVectorShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		basevectorshapegrid := new(models.BaseVectorShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(basevectorshapegrid, formGroup, probe)
	case "BottomCurvePlane1Shape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "BottomCurvePlane1Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BottomCurvePlane1ShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		bottomcurveplane1shape := new(models.BottomCurvePlane1Shape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bottomcurveplane1shape, formGroup, probe)
	case "BottomCurvePlane2Shape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "BottomCurvePlane2Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BottomCurvePlane2ShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		bottomcurveplane2shape := new(models.BottomCurvePlane2Shape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bottomcurveplane2shape, formGroup, probe)
	case "ChosenP1P2PairShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ChosenP1P2PairShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ChosenP1P2PairShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		chosenp1p2pairshape := new(models.ChosenP1P2PairShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(chosenp1p2pairshape, formGroup, probe)
	case "CircleGridShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CircleGridShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleGridShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		circlegridshape := new(models.CircleGridShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(circlegridshape, formGroup, probe)
	case "Circumference3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Circumference3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Circumference3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		circumference3dshape := new(models.Circumference3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(circumference3dshape, formGroup, probe)
	case "Clock2DDiagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Clock2DDiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Clock2DDiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		clock2ddiagram := new(models.Clock2DDiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(clock2ddiagram, formGroup, probe)
	case "Clock3DDiagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Clock3DDiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Clock3DDiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		clock3ddiagram := new(models.Clock3DDiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(clock3ddiagram, formGroup, probe)
	case "ClockTopCurveShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ClockTopCurveShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ClockTopCurveShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		clocktopcurveshape := new(models.ClockTopCurveShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(clocktopcurveshape, formGroup, probe)
	case "CutLine3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CutLine3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CutLine3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		cutline3dshape := new(models.CutLine3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cutline3dshape, formGroup, probe)
	case "EndArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EndArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		endarcshape := new(models.EndArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(endarcshape, formGroup, probe)
	case "EndArcShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EndArcShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndArcShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		endarcshapegrid := new(models.EndArcShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(endarcshapegrid, formGroup, probe)
	case "EndHalfwayArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EndHalfwayArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndHalfwayArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		endhalfwayarcshape := new(models.EndHalfwayArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(endhalfwayarcshape, formGroup, probe)
	case "EndHalfwayArcShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EndHalfwayArcShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndHalfwayArcShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		endhalfwayarcshapegrid := new(models.EndHalfwayArcShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(endhalfwayarcshapegrid, formGroup, probe)
	case "ExplanationTextShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ExplanationTextShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExplanationTextShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		explanationtextshape := new(models.ExplanationTextShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(explanationtextshape, formGroup, probe)
	case "Eye3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Eye3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Eye3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		eye3dshape := new(models.Eye3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(eye3dshape, formGroup, probe)
	case "EyeCornersSampledPoints3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EyeCornersSampledPoints3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EyeCornersSampledPoints3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		eyecornerssampledpoints3dshape := new(models.EyeCornersSampledPoints3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(eyecornerssampledpoints3dshape, formGroup, probe)
	case "EyeSampledPoints3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EyeSampledPoints3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EyeSampledPoints3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		eyesampledpoints3dshape := new(models.EyeSampledPoints3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(eyesampledpoints3dshape, formGroup, probe)
	case "EyeSeatBottomCurveShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EyeSeatBottomCurveShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EyeSeatBottomCurveShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		eyeseatbottomcurveshape := new(models.EyeSeatBottomCurveShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(eyeseatbottomcurveshape, formGroup, probe)
	case "EyeStoolBottomCurveShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EyeStoolBottomCurveShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EyeStoolBottomCurveShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		eyestoolbottomcurveshape := new(models.EyeStoolBottomCurveShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(eyestoolbottomcurveshape, formGroup, probe)
	case "EyeVolume3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EyeVolume3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EyeVolume3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		eyevolume3dshape := new(models.EyeVolume3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(eyevolume3dshape, formGroup, probe)
	case "GridPathShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GridPathShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GridPathShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		gridpathshape := new(models.GridPathShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gridpathshape, formGroup, probe)
	case "GrowthCurve2D":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GrowthCurve2D Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurve2DFormCallback(
			nil,
			probe,
			formGroup,
		)
		growthcurve2d := new(models.GrowthCurve2D)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(growthcurve2d, formGroup, probe)
	case "GrowthCurve2DRibbon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GrowthCurve2DRibbon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurve2DRibbonFormCallback(
			nil,
			probe,
			formGroup,
		)
		growthcurve2dribbon := new(models.GrowthCurve2DRibbon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(growthcurve2dribbon, formGroup, probe)
	case "GrowthCurve2DRibbonEndShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GrowthCurve2DRibbonEndShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		growthcurve2dribbonendshape := new(models.GrowthCurve2DRibbonEndShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(growthcurve2dribbonendshape, formGroup, probe)
	case "GrowthCurve2DRibbonStartShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GrowthCurve2DRibbonStartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		growthcurve2dribbonstartshape := new(models.GrowthCurve2DRibbonStartShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(growthcurve2dribbonstartshape, formGroup, probe)
	case "GrowthCurveRhombusGridShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GrowthCurveRhombusGridShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurveRhombusGridShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		growthcurverhombusgridshape := new(models.GrowthCurveRhombusGridShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(growthcurverhombusgridshape, formGroup, probe)
	case "GrowthCurveRhombusShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GrowthCurveRhombusShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurveRhombusShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		growthcurverhombusshape := new(models.GrowthCurveRhombusShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(growthcurverhombusshape, formGroup, probe)
	case "GrowthVectorShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GrowthVectorShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthVectorShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		growthvectorshape := new(models.GrowthVectorShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(growthvectorshape, formGroup, probe)
	case "InitialRhombusGridShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "InitialRhombusGridShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InitialRhombusGridShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		initialrhombusgridshape := new(models.InitialRhombusGridShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(initialrhombusgridshape, formGroup, probe)
	case "InitialRhombusShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "InitialRhombusShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InitialRhombusShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		initialrhombusshape := new(models.InitialRhombusShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(initialrhombusshape, formGroup, probe)
	case "Key3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Key3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Key3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		key3dshape := new(models.Key3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(key3dshape, formGroup, probe)
	case "KeyHole3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "KeyHole3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KeyHole3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		keyhole3dshape := new(models.KeyHole3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(keyhole3dshape, formGroup, probe)
	case "KeyHoleShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "KeyHoleShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KeyHoleShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		keyholeshape := new(models.KeyHoleShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(keyholeshape, formGroup, probe)
	case "Leaves3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Leaves3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Leaves3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		leaves3dshape := new(models.Leaves3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(leaves3dshape, formGroup, probe)
	case "Library":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Library Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			nil,
			probe,
			formGroup,
		)
		library := new(models.Library)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(library, formGroup, probe)
	case "MidArcVectorShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MidArcVectorShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MidArcVectorShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		midarcvectorshape := new(models.MidArcVectorShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(midarcvectorshape, formGroup, probe)
	case "MidArcVectorShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MidArcVectorShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MidArcVectorShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		midarcvectorshapegrid := new(models.MidArcVectorShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(midarcvectorshapegrid, formGroup, probe)
	case "OriginalPoints3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "OriginalPoints3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OriginalPoints3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		originalpoints3dshape := new(models.OriginalPoints3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(originalpoints3dshape, formGroup, probe)
	case "ParastichyMCurves3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ParastichyMCurves3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParastichyMCurves3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		parastichymcurves3dshape := new(models.ParastichyMCurves3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(parastichymcurves3dshape, formGroup, probe)
	case "ParastichyNCurves3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ParastichyNCurves3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParastichyNCurves3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		parastichyncurves3dshape := new(models.ParastichyNCurves3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(parastichyncurves3dshape, formGroup, probe)
	case "PartiallyGrowthCurve2DRibbon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DRibbon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DRibbonFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dribbon := new(models.PartiallyGrowthCurve2DRibbon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dribbon, formGroup, probe)
	case "PartiallyGrowthCurve2DRibbonEndShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DRibbonEndShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dribbonendshape := new(models.PartiallyGrowthCurve2DRibbonEndShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dribbonendshape, formGroup, probe)
	case "PartiallyGrowthCurve2DRibbonStartShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DRibbonStartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dribbonstartshape := new(models.PartiallyGrowthCurve2DRibbonStartShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dribbonstartshape, formGroup, probe)
	case "PartiallyGrowthCurve2DTrajectory":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DTrajectory Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dtrajectory := new(models.PartiallyGrowthCurve2DTrajectory)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dtrajectory, formGroup, probe)
	case "PartiallyGrowthCurve2DTrajectoryP1CurveShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DTrajectoryP1CurveShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dtrajectoryp1curveshape := new(models.PartiallyGrowthCurve2DTrajectoryP1CurveShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dtrajectoryp1curveshape, formGroup, probe)
	case "PartiallyGrowthCurve2DTrajectoryP1P2":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DTrajectoryP1P2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1P2FormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dtrajectoryp1p2 := new(models.PartiallyGrowthCurve2DTrajectoryP1P2)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dtrajectoryp1p2, formGroup, probe)
	case "PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dtrajectoryp1p2pairlineshape := new(models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dtrajectoryp1p2pairlineshape, formGroup, probe)
	case "PartiallyGrowthCurve2DTrajectoryP1PointShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DTrajectoryP1PointShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dtrajectoryp1pointshape := new(models.PartiallyGrowthCurve2DTrajectoryP1PointShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dtrajectoryp1pointshape, formGroup, probe)
	case "PartiallyGrowthCurve2DTrajectoryP2CurveShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DTrajectoryP2CurveShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dtrajectoryp2curveshape := new(models.PartiallyGrowthCurve2DTrajectoryP2CurveShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dtrajectoryp2curveshape, formGroup, probe)
	case "PartiallyGrowthCurve2DTrajectoryP2PointShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DTrajectoryP2PointShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dtrajectoryp2pointshape := new(models.PartiallyGrowthCurve2DTrajectoryP2PointShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dtrajectoryp2pointshape, formGroup, probe)
	case "PartiallyGrowthCurve2DTrajectoryShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyGrowthCurve2DTrajectoryShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallygrowthcurve2dtrajectoryshape := new(models.PartiallyGrowthCurve2DTrajectoryShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallygrowthcurve2dtrajectoryshape, formGroup, probe)
	case "PartiallyRotatedSeatBottomCurveShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyRotatedSeatBottomCurveShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyRotatedSeatBottomCurveShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallyrotatedseatbottomcurveshape := new(models.PartiallyRotatedSeatBottomCurveShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallyrotatedseatbottomcurveshape, formGroup, probe)
	case "PartiallyRotatedSeatTopCurveShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyRotatedSeatTopCurveShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyRotatedSeatTopCurveShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallyrotatedseattopcurveshape := new(models.PartiallyRotatedSeatTopCurveShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallyrotatedseattopcurveshape, formGroup, probe)
	case "PartiallyRotatedTorusShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartiallyRotatedTorusShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartiallyRotatedTorusShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partiallyrotatedtorusshape := new(models.PartiallyRotatedTorusShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partiallyrotatedtorusshape, formGroup, probe)
	case "PerpendicularVector":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PerpendicularVector Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PerpendicularVectorFormCallback(
			nil,
			probe,
			formGroup,
		)
		perpendicularvector := new(models.PerpendicularVector)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(perpendicularvector, formGroup, probe)
	case "PerpendicularVectorGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PerpendicularVectorGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PerpendicularVectorGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		perpendicularvectorgrid := new(models.PerpendicularVectorGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(perpendicularvectorgrid, formGroup, probe)
	case "PerpendicularVectorGridHalfway":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PerpendicularVectorGridHalfway Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PerpendicularVectorGridHalfwayFormCallback(
			nil,
			probe,
			formGroup,
		)
		perpendicularvectorgridhalfway := new(models.PerpendicularVectorGridHalfway)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(perpendicularvectorgridhalfway, formGroup, probe)
	case "PerpendicularVectorHalfway":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PerpendicularVectorHalfway Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PerpendicularVectorHalfwayFormCallback(
			nil,
			probe,
			formGroup,
		)
		perpendicularvectorhalfway := new(models.PerpendicularVectorHalfway)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(perpendicularvectorhalfway, formGroup, probe)
	case "Plant2DDiagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Plant2DDiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Plant2DDiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		plant2ddiagram := new(models.Plant2DDiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(plant2ddiagram, formGroup, probe)
	case "Plant3DDiagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Plant3DDiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Plant3DDiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		plant3ddiagram := new(models.Plant3DDiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(plant3ddiagram, formGroup, probe)
	case "PlantAbstract":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PlantAbstract Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlantAbstractFormCallback(
			nil,
			probe,
			formGroup,
		)
		plantabstract := new(models.PlantAbstract)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(plantabstract, formGroup, probe)
	case "PlantCircumferenceShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PlantCircumferenceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlantCircumferenceShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		plantcircumferenceshape := new(models.PlantCircumferenceShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(plantcircumferenceshape, formGroup, probe)
	case "PointsAndLines3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PointsAndLines3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PointsAndLines3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		pointsandlines3dshape := new(models.PointsAndLines3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pointsandlines3dshape, formGroup, probe)
	case "PxShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PxShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PxShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		pxshape := new(models.PxShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pxshape, formGroup, probe)
	case "Rendered3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Rendered3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Rendered3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		rendered3dshape := new(models.Rendered3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rendered3dshape, formGroup, probe)
	case "RhombusShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RhombusShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RhombusShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		rhombusshape := new(models.RhombusShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rhombusshape, formGroup, probe)
	case "RhombusStuff":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RhombusStuff Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RhombusStuffFormCallback(
			nil,
			probe,
			formGroup,
		)
		rhombusstuff := new(models.RhombusStuff)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rhombusstuff, formGroup, probe)
	case "RotatedRhombusGridShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RotatedRhombusGridShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RotatedRhombusGridShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		rotatedrhombusgridshape := new(models.RotatedRhombusGridShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rotatedrhombusgridshape, formGroup, probe)
	case "RotatedRhombusShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RotatedRhombusShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RotatedRhombusShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		rotatedrhombusshape := new(models.RotatedRhombusShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rotatedrhombusshape, formGroup, probe)
	case "RotatedSampledPoints3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RotatedSampledPoints3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RotatedSampledPoints3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		rotatedsampledpoints3dshape := new(models.RotatedSampledPoints3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rotatedsampledpoints3dshape, formGroup, probe)
	case "RotatedSeatAndLegs3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RotatedSeatAndLegs3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RotatedSeatAndLegs3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		rotatedseatandlegs3dshape := new(models.RotatedSeatAndLegs3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rotatedseatandlegs3dshape, formGroup, probe)
	case "SampledPoints3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SampledPoints3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SampledPoints3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		sampledpoints3dshape := new(models.SampledPoints3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(sampledpoints3dshape, formGroup, probe)
	case "Seat3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Seat3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Seat3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		seat3dshape := new(models.Seat3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(seat3dshape, formGroup, probe)
	case "SeatAndLegs3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SeatAndLegs3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SeatAndLegs3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		seatandlegs3dshape := new(models.SeatAndLegs3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(seatandlegs3dshape, formGroup, probe)
	case "SeatBottomCurveShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SeatBottomCurveShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SeatBottomCurveShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		seatbottomcurveshape := new(models.SeatBottomCurveShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(seatbottomcurveshape, formGroup, probe)
	case "SeatTopCurveShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SeatTopCurveShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SeatTopCurveShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		seattopcurveshape := new(models.SeatTopCurveShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(seattopcurveshape, formGroup, probe)
	case "ShiftedBottomTopStartArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedBottomTopStartArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedBottomTopStartArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedbottomtopstartarcshape := new(models.ShiftedBottomTopStartArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedbottomtopstartarcshape, formGroup, probe)
	case "ShiftedBottomTopStartArcShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedBottomTopStartArcShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedBottomTopStartArcShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedbottomtopstartarcshapegrid := new(models.ShiftedBottomTopStartArcShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedbottomtopstartarcshapegrid, formGroup, probe)
	case "ShiftedLeftGrowthCurve2DRibbon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftGrowthCurve2DRibbon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftGrowthCurve2DRibbonFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftgrowthcurve2dribbon := new(models.ShiftedLeftGrowthCurve2DRibbon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftgrowthcurve2dribbon, formGroup, probe)
	case "ShiftedLeftGrowthCurve2DRibbonEndShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftGrowthCurve2DRibbonEndShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftGrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftgrowthcurve2dribbonendshape := new(models.ShiftedLeftGrowthCurve2DRibbonEndShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftgrowthcurve2dribbonendshape, formGroup, probe)
	case "ShiftedLeftGrowthCurve2DRibbonStartShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftGrowthCurve2DRibbonStartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftGrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftgrowthcurve2dribbonstartshape := new(models.ShiftedLeftGrowthCurve2DRibbonStartShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftgrowthcurve2dribbonstartshape, formGroup, probe)
	case "ShiftedLeftPartiallyGrowthCurve2DRibbon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftPartiallyGrowthCurve2DRibbon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftPartiallyGrowthCurve2DRibbonFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftpartiallygrowthcurve2dribbon := new(models.ShiftedLeftPartiallyGrowthCurve2DRibbon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftpartiallygrowthcurve2dribbon, formGroup, probe)
	case "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftpartiallygrowthcurve2dribbonendshape := new(models.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftpartiallygrowthcurve2dribbonendshape, formGroup, probe)
	case "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftpartiallygrowthcurve2dribbonstartshape := new(models.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftpartiallygrowthcurve2dribbonstartshape, formGroup, probe)
	case "ShiftedLeftStackGrowthCurveEndArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftStackGrowthCurveEndArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftStackGrowthCurveEndArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftstackgrowthcurveendarcshape := new(models.ShiftedLeftStackGrowthCurveEndArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftstackgrowthcurveendarcshape, formGroup, probe)
	case "ShiftedLeftStackGrowthCurveStartArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftStackGrowthCurveStartArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftStackGrowthCurveStartArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftstackgrowthcurvestartarcshape := new(models.ShiftedLeftStackGrowthCurveStartArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftstackgrowthcurvestartarcshape, formGroup, probe)
	case "ShiftedLeftStackNormalVector":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftStackNormalVector Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftStackNormalVectorFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftstacknormalvector := new(models.ShiftedLeftStackNormalVector)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftstacknormalvector, formGroup, probe)
	case "ShiftedLeftStackOfGrowthCurve":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftStackOfGrowthCurve Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftStackOfGrowthCurveFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftstackofgrowthcurve := new(models.ShiftedLeftStackOfGrowthCurve)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftstackofgrowthcurve, formGroup, probe)
	case "ShiftedLeftStackOfNormalVector":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedLeftStackOfNormalVector Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedLeftStackOfNormalVectorFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedleftstackofnormalvector := new(models.ShiftedLeftStackOfNormalVector)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedleftstackofnormalvector, formGroup, probe)
	case "ShiftedRightGrowthCurve2DRibbon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedRightGrowthCurve2DRibbon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedRightGrowthCurve2DRibbonFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedrightgrowthcurve2dribbon := new(models.ShiftedRightGrowthCurve2DRibbon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedrightgrowthcurve2dribbon, formGroup, probe)
	case "ShiftedRightGrowthCurve2DRibbonEndShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedRightGrowthCurve2DRibbonEndShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedrightgrowthcurve2dribbonendshape := new(models.ShiftedRightGrowthCurve2DRibbonEndShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedrightgrowthcurve2dribbonendshape, formGroup, probe)
	case "ShiftedRightGrowthCurve2DRibbonStartShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShiftedRightGrowthCurve2DRibbonStartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		shiftedrightgrowthcurve2dribbonstartshape := new(models.ShiftedRightGrowthCurve2DRibbonStartShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shiftedrightgrowthcurve2dribbonstartshape, formGroup, probe)
	case "StackGrowthCurve2DEndHalfwayArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackGrowthCurve2DEndHalfwayArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurve2DEndHalfwayArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackgrowthcurve2dendhalfwayarcshape := new(models.StackGrowthCurve2DEndHalfwayArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackgrowthcurve2dendhalfwayarcshape, formGroup, probe)
	case "StackGrowthCurve2DRibbonEndShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackGrowthCurve2DRibbonEndShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackgrowthcurve2dribbonendshape := new(models.StackGrowthCurve2DRibbonEndShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackgrowthcurve2dribbonendshape, formGroup, probe)
	case "StackGrowthCurve2DRibbonStartShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackGrowthCurve2DRibbonStartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackgrowthcurve2dribbonstartshape := new(models.StackGrowthCurve2DRibbonStartShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackgrowthcurve2dribbonstartshape, formGroup, probe)
	case "StackGrowthCurve2DStartHalfwayArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackGrowthCurve2DStartHalfwayArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurve2DStartHalfwayArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackgrowthcurve2dstarthalfwayarcshape := new(models.StackGrowthCurve2DStartHalfwayArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackgrowthcurve2dstarthalfwayarcshape, formGroup, probe)
	case "StackOfGrowthCurve2D":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackOfGrowthCurve2D Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfGrowthCurve2DFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackofgrowthcurve2d := new(models.StackOfGrowthCurve2D)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackofgrowthcurve2d, formGroup, probe)
	case "StackOfGrowthCurve2DByGrowthVector":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackOfGrowthCurve2DByGrowthVector Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfGrowthCurve2DByGrowthVectorFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackofgrowthcurve2dbygrowthvector := new(models.StackOfGrowthCurve2DByGrowthVector)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackofgrowthcurve2dbygrowthvector, formGroup, probe)
	case "StackOfGrowthCurve2DRibbon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackOfGrowthCurve2DRibbon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfGrowthCurve2DRibbonFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackofgrowthcurve2dribbon := new(models.StackOfGrowthCurve2DRibbon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackofgrowthcurve2dribbon, formGroup, probe)
	case "StackOfPartiallyRotatedTorusShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackOfPartiallyRotatedTorusShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfPartiallyRotatedTorusShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackofpartiallyrotatedtorusshape := new(models.StackOfPartiallyRotatedTorusShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackofpartiallyrotatedtorusshape, formGroup, probe)
	case "StackOfRotatedGrowthCurve2D":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackOfRotatedGrowthCurve2D Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfRotatedGrowthCurve2DFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackofrotatedgrowthcurve2d := new(models.StackOfRotatedGrowthCurve2D)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackofrotatedgrowthcurve2d, formGroup, probe)
	case "StackOfRotatedGrowthCurve2DRibbon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackOfRotatedGrowthCurve2DRibbon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfRotatedGrowthCurve2DRibbonFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackofrotatedgrowthcurve2dribbon := new(models.StackOfRotatedGrowthCurve2DRibbon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackofrotatedgrowthcurve2dribbon, formGroup, probe)
	case "StackOfRotatedVaseTrapezeRingsShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackOfRotatedVaseTrapezeRingsShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfRotatedVaseTrapezeRingsShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackofrotatedvasetrapezeringsshape := new(models.StackOfRotatedVaseTrapezeRingsShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackofrotatedvasetrapezeringsshape, formGroup, probe)
	case "StackOfVaseTrapezeRingsShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackOfVaseTrapezeRingsShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfVaseTrapezeRingsShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackofvasetrapezeringsshape := new(models.StackOfVaseTrapezeRingsShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackofvasetrapezeringsshape, formGroup, probe)
	case "StackRotatedGrowthCurve2DEndArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackRotatedGrowthCurve2DEndArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DEndArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackrotatedgrowthcurve2dendarcshape := new(models.StackRotatedGrowthCurve2DEndArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackrotatedgrowthcurve2dendarcshape, formGroup, probe)
	case "StackRotatedGrowthCurve2DRibbonEndShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackRotatedGrowthCurve2DRibbonEndShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackrotatedgrowthcurve2dribbonendshape := new(models.StackRotatedGrowthCurve2DRibbonEndShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackrotatedgrowthcurve2dribbonendshape, formGroup, probe)
	case "StackRotatedGrowthCurve2DRibbonStartShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackRotatedGrowthCurve2DRibbonStartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackrotatedgrowthcurve2dribbonstartshape := new(models.StackRotatedGrowthCurve2DRibbonStartShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackrotatedgrowthcurve2dribbonstartshape, formGroup, probe)
	case "StackRotatedGrowthCurve2DStartArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackRotatedGrowthCurve2DStartArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DStartArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackrotatedgrowthcurve2dstartarcshape := new(models.StackRotatedGrowthCurve2DStartArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackrotatedgrowthcurve2dstartarcshape, formGroup, probe)
	case "StartArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StartArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StartArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		startarcshape := new(models.StartArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(startarcshape, formGroup, probe)
	case "StartArcShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StartArcShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StartArcShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		startarcshapegrid := new(models.StartArcShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(startarcshapegrid, formGroup, probe)
	case "StartHalfwayArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StartHalfwayArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StartHalfwayArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		starthalfwayarcshape := new(models.StartHalfwayArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(starthalfwayarcshape, formGroup, probe)
	case "StartHalfwayArcShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StartHalfwayArcShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StartHalfwayArcShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		starthalfwayarcshapegrid := new(models.StartHalfwayArcShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(starthalfwayarcshapegrid, formGroup, probe)
	case "StemCylinder3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StemCylinder3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StemCylinder3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stemcylinder3dshape := new(models.StemCylinder3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stemcylinder3dshape, formGroup, probe)
	case "Stool2DDiagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Stool2DDiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Stool2DDiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		stool2ddiagram := new(models.Stool2DDiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stool2ddiagram, formGroup, probe)
	case "Stool3DDiagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Stool3DDiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Stool3DDiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		stool3ddiagram := new(models.Stool3DDiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stool3ddiagram, formGroup, probe)
	case "TiledFloor3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TiledFloor3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TiledFloor3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		tiledfloor3dshape := new(models.TiledFloor3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tiledfloor3dshape, formGroup, probe)
	case "TopCurvePlane1Shape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopCurvePlane1Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopCurvePlane1ShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topcurveplane1shape := new(models.TopCurvePlane1Shape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topcurveplane1shape, formGroup, probe)
	case "TopCurvePlane2Shape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopCurvePlane2Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopCurvePlane2ShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topcurveplane2shape := new(models.TopCurvePlane2Shape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topcurveplane2shape, formGroup, probe)
	case "TopEndArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopEndArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopEndArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topendarcshape := new(models.TopEndArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topendarcshape, formGroup, probe)
	case "TopEndArcShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopEndArcShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopEndArcShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		topendarcshapegrid := new(models.TopEndArcShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topendarcshapegrid, formGroup, probe)
	case "TopEndHalfwayArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopEndHalfwayArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopEndHalfwayArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topendhalfwayarcshape := new(models.TopEndHalfwayArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topendhalfwayarcshape, formGroup, probe)
	case "TopEndHalfwayArcShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopEndHalfwayArcShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopEndHalfwayArcShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		topendhalfwayarcshapegrid := new(models.TopEndHalfwayArcShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topendhalfwayarcshapegrid, formGroup, probe)
	case "TopGrowthCurve2D":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopGrowthCurve2D Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopGrowthCurve2DFormCallback(
			nil,
			probe,
			formGroup,
		)
		topgrowthcurve2d := new(models.TopGrowthCurve2D)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topgrowthcurve2d, formGroup, probe)
	case "TopMidArcVectorShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopMidArcVectorShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopMidArcVectorShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topmidarcvectorshape := new(models.TopMidArcVectorShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topmidarcvectorshape, formGroup, probe)
	case "TopMidArcVectorShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopMidArcVectorShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopMidArcVectorShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		topmidarcvectorshapegrid := new(models.TopMidArcVectorShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topmidarcvectorshapegrid, formGroup, probe)
	case "TopStackGrowthCurve2DEndHalfwayArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStackGrowthCurve2DEndHalfwayArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstackgrowthcurve2dendhalfwayarcshape := new(models.TopStackGrowthCurve2DEndHalfwayArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstackgrowthcurve2dendhalfwayarcshape, formGroup, probe)
	case "TopStackGrowthCurve2DStartHalfwayArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStackGrowthCurve2DStartHalfwayArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstackgrowthcurve2dstarthalfwayarcshape := new(models.TopStackGrowthCurve2DStartHalfwayArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstackgrowthcurve2dstarthalfwayarcshape, formGroup, probe)
	case "TopStackOfGrowthCurve2D":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStackOfGrowthCurve2D Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackOfGrowthCurve2DFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstackofgrowthcurve2d := new(models.TopStackOfGrowthCurve2D)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstackofgrowthcurve2d, formGroup, probe)
	case "TopStackOfRotatedGrowthCurve2D":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStackOfRotatedGrowthCurve2D Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackOfRotatedGrowthCurve2DFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstackofrotatedgrowthcurve2d := new(models.TopStackOfRotatedGrowthCurve2D)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstackofrotatedgrowthcurve2d, formGroup, probe)
	case "TopStackOfRotatedGrowthCurve2DEndArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStackOfRotatedGrowthCurve2DEndArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstackofrotatedgrowthcurve2dendarcshape := new(models.TopStackOfRotatedGrowthCurve2DEndArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstackofrotatedgrowthcurve2dendarcshape, formGroup, probe)
	case "TopStackOfRotatedGrowthCurve2DStartArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStackOfRotatedGrowthCurve2DStartArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstackofrotatedgrowthcurve2dstartarcshape := new(models.TopStackOfRotatedGrowthCurve2DStartArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstackofrotatedgrowthcurve2dstartarcshape, formGroup, probe)
	case "TopStartArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStartArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStartArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstartarcshape := new(models.TopStartArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstartarcshape, formGroup, probe)
	case "TopStartArcShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStartArcShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStartArcShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstartarcshapegrid := new(models.TopStartArcShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstartarcshapegrid, formGroup, probe)
	case "TopStartHalfwayArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStartHalfwayArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStartHalfwayArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstarthalfwayarcshape := new(models.TopStartHalfwayArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstarthalfwayarcshape, formGroup, probe)
	case "TopStartHalfwayArcShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStartHalfwayArcShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStartHalfwayArcShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstarthalfwayarcshapegrid := new(models.TopStartHalfwayArcShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstarthalfwayarcshapegrid, formGroup, probe)
	case "Torus3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Torus3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Torus3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		torus3dshape := new(models.Torus3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(torus3dshape, formGroup, probe)
	case "TorusEdge3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TorusEdge3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TorusEdge3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		torusedge3dshape := new(models.TorusEdge3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(torusedge3dshape, formGroup, probe)
	case "TorusStackShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TorusStackShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TorusStackShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		torusstackshape := new(models.TorusStackShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(torusstackshape, formGroup, probe)
	case "TubeVase3DDiagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TubeVase3DDiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TubeVase3DDiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		tubevase3ddiagram := new(models.TubeVase3DDiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tubevase3ddiagram, formGroup, probe)
	case "TubeVaseAbstract":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TubeVaseAbstract Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TubeVaseAbstractFormCallback(
			nil,
			probe,
			formGroup,
		)
		tubevaseabstract := new(models.TubeVaseAbstract)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tubevaseabstract, formGroup, probe)
	case "Vase2DDiagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Vase2DDiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Vase2DDiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		vase2ddiagram := new(models.Vase2DDiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(vase2ddiagram, formGroup, probe)
	case "VaseTrapezeRingShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "VaseTrapezeRingShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VaseTrapezeRingShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		vasetrapezeringshape := new(models.VaseTrapezeRingShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(vasetrapezeringshape, formGroup, probe)
	case "VerticalTorusStackShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "VerticalTorusStackShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VerticalTorusStackShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		verticaltorusstackshape := new(models.VerticalTorusStackShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(verticaltorusstackshape, formGroup, probe)
	case "VolumeKey3DShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "VolumeKey3DShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VolumeKey3DShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		volumekey3dshape := new(models.VolumeKey3DShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(volumekey3dshape, formGroup, probe)
	}
	formStage.Commit()
}
