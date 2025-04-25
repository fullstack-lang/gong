// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Chapter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MardownContent", instanceWithInferedType.MardownContent, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)
		AssociationSliceToForm("Pages", instanceWithInferedType, &instanceWithInferedType.Pages, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Content"
			rf.Fieldname = "Chapters"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Content),
					"Chapters",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Content, *models.Chapter](
					nil,
					"Chapters",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Content:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MardownContent", instanceWithInferedType.MardownContent, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)
		BasicFieldtoForm("ContentPath", instanceWithInferedType.ContentPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OutputPath", instanceWithInferedType.OutputPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LayoutPath", instanceWithInferedType.LayoutPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StaticPath", instanceWithInferedType.StaticPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Target", instanceWithInferedType.Target, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("Chapters", instanceWithInferedType, &instanceWithInferedType.Chapters, formGroup, probe)

	case *models.Page:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MardownContent", instanceWithInferedType.MardownContent, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Chapter"
			rf.Fieldname = "Pages"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Chapter),
					"Pages",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Chapter, *models.Page](
					nil,
					"Pages",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
