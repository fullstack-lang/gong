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
		case *RhombusGridShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RhombusGridShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.rhombusgridshape, probe)
			}
		case *RhombusShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RhombusShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.rhombusshape, probe)
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
	case "RhombusGridShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RhombusGridShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RhombusGridShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		rhombusgridshape := new(models.RhombusGridShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rhombusgridshape, formGroup, probe)
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
	}
	formStage.Commit()
}
