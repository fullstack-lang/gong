// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/cld/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Category1:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Category1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category1FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Category1Shape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Category1Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category1ShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Category2:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Category2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Category2Shape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Category2Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category2ShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Category3:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Category3 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category3FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Category3Shape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Category3Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category3ShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ControlPointShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ControlPointShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Desk:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Desk Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeskFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Diagram:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Diagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Influence:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Influence Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InfluenceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.InfluenceShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "InfluenceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InfluenceShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
