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
		case *AxesShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "AxesShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.axesshape, probe)
			}
		case *CircleGridShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "CircleGridShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.circlegridshape, probe)
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
		case *GrowthCurveBezierShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GrowthCurveBezierShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.growthcurvebeziershape, probe)
			}
		case *GrowthCurveBezierShapeGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GrowthCurveBezierShapeGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.growthcurvebeziershapegrid, probe)
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
		case *NextCircleShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NextCircleShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.nextcircleshape, probe)
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
		case *StackGrowthCurveBezierShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackGrowthCurveBezierShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackgrowthcurvebeziershape, probe)
			}
		case *StackOfGrowthCurveFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StackOfGrowthCurve", true)
			} else {
				FillUpFormFromGongstruct(onSave.stackofgrowthcurve, probe)
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
	case "GrowthCurveBezierShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GrowthCurveBezierShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurveBezierShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		growthcurvebeziershape := new(models.GrowthCurveBezierShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(growthcurvebeziershape, formGroup, probe)
	case "GrowthCurveBezierShapeGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GrowthCurveBezierShapeGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GrowthCurveBezierShapeGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		growthcurvebeziershapegrid := new(models.GrowthCurveBezierShapeGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(growthcurvebeziershapegrid, formGroup, probe)
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
	case "NextCircleShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NextCircleShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NextCircleShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		nextcircleshape := new(models.NextCircleShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(nextcircleshape, formGroup, probe)
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
	case "StackGrowthCurveBezierShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StackGrowthCurveBezierShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StackGrowthCurveBezierShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stackgrowthcurvebeziershape := new(models.StackGrowthCurveBezierShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stackgrowthcurvebeziershape, formGroup, probe)
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
	}
	formStage.Commit()
}
