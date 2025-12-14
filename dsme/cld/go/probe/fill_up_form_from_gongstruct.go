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
	case *models.ArtefactType:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ArtefactType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArtefactTypeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ArtefactTypeShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ArtefactTypeShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArtefactTypeShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Artist:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Artist Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArtistFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ArtistShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ArtistShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArtistShapeFormCallback(
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
	case *models.Movement:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Movement Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MovementFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MovementShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MovementShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MovementShapeFormCallback(
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
