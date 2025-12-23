// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.CompositionShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "CompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CompositionShapeFormCallback(
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
	case *models.Product:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Product Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ProductShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ProductShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Project:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Project Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProjectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Root:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Root Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RootFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Task:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Task Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TaskShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskShapeFormCallback(
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
