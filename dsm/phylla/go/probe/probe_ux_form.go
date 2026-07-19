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
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *ArcNormalVectorShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ArcNormalVectorShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.arcnormalvectorshape, probe)
			}
		case *ArcNormalVectorShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ArcNormalVectorShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.arcnormalvectorshapegrid, probe)
			}
		case *AxesShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "AxesShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.axesshape, probe)
			}
		case *BaseVectorShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "BaseVectorShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.basevectorshape, probe)
			}
		case *BaseVectorShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "BaseVectorShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.basevectorshapegrid, probe)
			}
		case *CircleGridShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "CircleGridShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.circlegridshape, probe)
			}
		case *DiscreteTorusShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DiscreteTorusShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.discretetorusshape, probe)
			}
		case *DiscreteTorusStackShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DiscreteTorusStackShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.discretetorusstackshape, probe)
			}
		case *EndArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "EndArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.endarcshape, probe)
			}
		case *EndArcShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "EndArcShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.endarcshapegrid, probe)
			}
		case *EndHalfwayArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "EndHalfwayArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.endhalfwayarcshape, probe)
			}
		case *EndHalfwayArcShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "EndHalfwayArcShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.endhalfwayarcshapegrid, probe)
			}
		case *ExplanationTextShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ExplanationTextShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.explanationtextshape, probe)
			}
		case *GridPathShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GridPathShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.gridpathshape, probe)
			}
		case *GrowthCurve2DFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GrowthCurve2D", true)
			} else {
				FillUpFormFromGongstruct(onSave.growthcurve2d, probe)
			}
		case *GrowthCurveRhombusGridShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GrowthCurveRhombusGridShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.growthcurverhombusgridshape, probe)
			}
		case *GrowthCurveRhombusShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GrowthCurveRhombusShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.growthcurverhombusshape, probe)
			}
		case *GrowthVectorShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GrowthVectorShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.growthvectorshape, probe)
			}
		case *InitialRhombusGridShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "InitialRhombusGridShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.initialrhombusgridshape, probe)
			}
		case *InitialRhombusShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "InitialRhombusShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.initialrhombusshape, probe)
			}
		case *LibraryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Library", true)
			} else {
				FillUpFormFromGongstruct(onSave.library, probe)
			}
		case *MidArcVectorShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MidArcVectorShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.midarcvectorshape, probe)
			}
		case *MidArcVectorShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MidArcVectorShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.midarcvectorshapegrid, probe)
			}
		case *PerpendicularVectorFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PerpendicularVector", true)
			} else {
				FillUpFormFromGongstruct(onSave.perpendicularvector, probe)
			}
		case *PerpendicularVectorGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PerpendicularVectorGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.perpendicularvectorgrid, probe)
			}
		case *PerpendicularVectorGridHalfwayFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PerpendicularVectorGridHalfway", true)
			} else {
				FillUpFormFromGongstruct(onSave.perpendicularvectorgridhalfway, probe)
			}
		case *PerpendicularVectorHalfwayFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PerpendicularVectorHalfway", true)
			} else {
				FillUpFormFromGongstruct(onSave.perpendicularvectorhalfway, probe)
			}
		case *PlantFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Plant", true)
			} else {
				FillUpFormFromGongstruct(onSave.plant, probe)
			}
		case *PlantCircumferenceShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PlantCircumferenceShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.plantcircumferenceshape, probe)
			}
		case *PlantDiagramFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PlantDiagram", true)
			} else {
				FillUpFormFromGongstruct(onSave.plantdiagram, probe)
			}
		case *Rendered3DShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Rendered3DShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.rendered3dshape, probe)
			}
		case *RhombusShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RhombusShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.rhombusshape, probe)
			}
		case *RhombusStuffFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RhombusStuff", true)
			} else {
				FillUpFormFromGongstruct(onSave.rhombusstuff, probe)
			}
		case *RotatedRhombusGridShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RotatedRhombusGridShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.rotatedrhombusgridshape, probe)
			}
		case *RotatedRhombusShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RotatedRhombusShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.rotatedrhombusshape, probe)
			}
		case *ShiftedBottomTopStartArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ShiftedBottomTopStartArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.shiftedbottomtopstartarcshape, probe)
			}
		case *ShiftedBottomTopStartArcShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ShiftedBottomTopStartArcShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.shiftedbottomtopstartarcshapegrid, probe)
			}
		case *ShiftedLeftStackGrowthCurveEndArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ShiftedLeftStackGrowthCurveEndArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.shiftedleftstackgrowthcurveendarcshape, probe)
			}
		case *ShiftedLeftStackGrowthCurveStartArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ShiftedLeftStackGrowthCurveStartArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.shiftedleftstackgrowthcurvestartarcshape, probe)
			}
		case *ShiftedLeftStackNormalVectorFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ShiftedLeftStackNormalVector", true)
			} else {
				FillUpFormFromGongstruct(onSave.shiftedleftstacknormalvector, probe)
			}
		case *ShiftedLeftStackOfGrowthCurveFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ShiftedLeftStackOfGrowthCurve", true)
			} else {
				FillUpFormFromGongstruct(onSave.shiftedleftstackofgrowthcurve, probe)
			}
		case *ShiftedLeftStackOfNormalVectorFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ShiftedLeftStackOfNormalVector", true)
			} else {
				FillUpFormFromGongstruct(onSave.shiftedleftstackofnormalvector, probe)
			}
		case *StackGrowthCurve2DEndHalfwayArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackGrowthCurve2DEndHalfwayArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackgrowthcurve2dendhalfwayarcshape, probe)
			}
		case *StackGrowthCurve2DRibbonEndShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackGrowthCurve2DRibbonEndShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackgrowthcurve2dribbonendshape, probe)
			}
		case *StackGrowthCurve2DRibbonStartShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackGrowthCurve2DRibbonStartShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackgrowthcurve2dribbonstartshape, probe)
			}
		case *StackGrowthCurve2DStartHalfwayArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackGrowthCurve2DStartHalfwayArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackgrowthcurve2dstarthalfwayarcshape, probe)
			}
		case *StackGrowthCurveEndArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackGrowthCurveEndArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackgrowthcurveendarcshape, probe)
			}
		case *StackGrowthCurveStartArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackGrowthCurveStartArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackgrowthcurvestartarcshape, probe)
			}
		case *StackOfGrowthCurveFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackOfGrowthCurve", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackofgrowthcurve, probe)
			}
		case *StackOfGrowthCurve2DFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackOfGrowthCurve2D", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackofgrowthcurve2d, probe)
			}
		case *StackOfGrowthCurve2DRibbonFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackOfGrowthCurve2DRibbon", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackofgrowthcurve2dribbon, probe)
			}
		case *StartArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StartArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.startarcshape, probe)
			}
		case *StartArcShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StartArcShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.startarcshapegrid, probe)
			}
		case *StartHalfwayArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StartHalfwayArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.starthalfwayarcshape, probe)
			}
		case *StartHalfwayArcShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StartHalfwayArcShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.starthalfwayarcshapegrid, probe)
			}
		case *TopEndArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopEndArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.topendarcshape, probe)
			}
		case *TopEndArcShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopEndArcShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.topendarcshapegrid, probe)
			}
		case *TopEndHalfwayArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopEndHalfwayArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.topendhalfwayarcshape, probe)
			}
		case *TopEndHalfwayArcShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopEndHalfwayArcShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.topendhalfwayarcshapegrid, probe)
			}
		case *TopGrowthCurve2DFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopGrowthCurve2D", true)
			} else {
				FillUpFormFromGongstruct(onSave.topgrowthcurve2d, probe)
			}
		case *TopMidArcVectorShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopMidArcVectorShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.topmidarcvectorshape, probe)
			}
		case *TopMidArcVectorShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopMidArcVectorShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.topmidarcvectorshapegrid, probe)
			}
		case *TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopStackGrowthCurve2DEndHalfwayArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.topstackgrowthcurve2dendhalfwayarcshape, probe)
			}
		case *TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopStackGrowthCurve2DStartHalfwayArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.topstackgrowthcurve2dstarthalfwayarcshape, probe)
			}
		case *TopStackGrowthCurveEndArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopStackGrowthCurveEndArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.topstackgrowthcurveendarcshape, probe)
			}
		case *TopStackGrowthCurveStartArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopStackGrowthCurveStartArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.topstackgrowthcurvestartarcshape, probe)
			}
		case *TopStackOfGrowthCurveFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopStackOfGrowthCurve", true)
			} else {
				FillUpFormFromGongstruct(onSave.topstackofgrowthcurve, probe)
			}
		case *TopStackOfGrowthCurve2DFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopStackOfGrowthCurve2D", true)
			} else {
				FillUpFormFromGongstruct(onSave.topstackofgrowthcurve2d, probe)
			}
		case *TopStartArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopStartArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.topstartarcshape, probe)
			}
		case *TopStartArcShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopStartArcShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.topstartarcshapegrid, probe)
			}
		case *TopStartHalfwayArcShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopStartHalfwayArcShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.topstarthalfwayarcshape, probe)
			}
		case *TopStartHalfwayArcShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TopStartHalfwayArcShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.topstarthalfwayarcshapegrid, probe)
			}
		case *TorusStackShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TorusStackShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.torusstackshape, probe)
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
	case "DiscreteTorusShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DiscreteTorusShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiscreteTorusShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		discretetorusshape := new(models.DiscreteTorusShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(discretetorusshape, formGroup, probe)
	case "DiscreteTorusStackShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DiscreteTorusStackShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiscreteTorusStackShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		discretetorusstackshape := new(models.DiscreteTorusStackShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(discretetorusstackshape, formGroup, probe)
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
	case "Plant":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Plant Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlantFormCallback(
			nil,
			probe,
			formGroup,
		)
		plant := new(models.Plant)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(plant, formGroup, probe)
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
	case "PlantDiagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PlantDiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlantDiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		plantdiagram := new(models.PlantDiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(plantdiagram, formGroup, probe)
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
	case "StackGrowthCurveEndArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackGrowthCurveEndArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurveEndArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackgrowthcurveendarcshape := new(models.StackGrowthCurveEndArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackgrowthcurveendarcshape, formGroup, probe)
	case "StackGrowthCurveStartArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackGrowthCurveStartArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurveStartArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackgrowthcurvestartarcshape := new(models.StackGrowthCurveStartArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackgrowthcurvestartarcshape, formGroup, probe)
	case "StackOfGrowthCurve":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackOfGrowthCurve Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackOfGrowthCurveFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackofgrowthcurve := new(models.StackOfGrowthCurve)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackofgrowthcurve, formGroup, probe)
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
	case "TopStackGrowthCurveEndArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStackGrowthCurveEndArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackGrowthCurveEndArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstackgrowthcurveendarcshape := new(models.TopStackGrowthCurveEndArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstackgrowthcurveendarcshape, formGroup, probe)
	case "TopStackGrowthCurveStartArcShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStackGrowthCurveStartArcShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackGrowthCurveStartArcShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstackgrowthcurvestartarcshape := new(models.TopStackGrowthCurveStartArcShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstackgrowthcurvestartarcshape, formGroup, probe)
	case "TopStackOfGrowthCurve":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TopStackOfGrowthCurve Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TopStackOfGrowthCurveFormCallback(
			nil,
			probe,
			formGroup,
		)
		topstackofgrowthcurve := new(models.TopStackOfGrowthCurve)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(topstackofgrowthcurve, formGroup, probe)
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
	}
	formStage.Commit()
}
