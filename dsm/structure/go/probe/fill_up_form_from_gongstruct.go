// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.DiagramStructure:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DiagramStructure Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramStructureFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Library:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Library Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Link:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LinkAssociationShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "LinkAssociationShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkAssociationShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Part Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "PartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.System:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "System Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SystemFormCallback(
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
