// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/button/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Button:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisabled", instanceWithInferedType.IsDisabled, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MatButtonType", instanceWithInferedType.MatButtonType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MatButtonAppearance", instanceWithInferedType.MatButtonAppearance, instanceWithInferedType, probe.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Buttons"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Group),
					"Buttons",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Group](
					nil,
					"Buttons",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ButtonToggle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisabled", instanceWithInferedType.IsDisabled, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GroupToogle"
			rf.Fieldname = "ButtonToggles"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GroupToogle),
					"ButtonToggles",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GroupToogle](
					nil,
					"ButtonToggles",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Percentage", instanceWithInferedType.Percentage, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Buttons", instanceWithInferedType, &instanceWithInferedType.Buttons, formGroup, probe)
		BasicFieldtoForm("NbColumns", instanceWithInferedType.NbColumns, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layout"
			rf.Fieldname = "Groups"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layout),
					"Groups",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layout](
					nil,
					"Groups",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.GroupToogle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Percentage", instanceWithInferedType.Percentage, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ButtonToggles", instanceWithInferedType, &instanceWithInferedType.ButtonToggles, formGroup, probe)
		BasicFieldtoForm("IsSingleSelector", instanceWithInferedType.IsSingleSelector, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layout"
			rf.Fieldname = "GroupToogles"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layout),
					"GroupToogles",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layout](
					nil,
					"GroupToogles",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		AssociationSliceToForm("GroupToogles", instanceWithInferedType, &instanceWithInferedType.GroupToogles, formGroup, probe)

	default:
		_ = instanceWithInferedType
	}
}
