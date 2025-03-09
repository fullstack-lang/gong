// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/button/go/models"
	"github.com/fullstack-lang/gong/lib/button/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Buttons"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Group),
					"Buttons",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Group, *models.Button](
					nil,
					"Buttons",
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layout"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layout),
					"Groups",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layout, *models.Group](
					nil,
					"Groups",
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

	default:
		_ = instanceWithInferedType
	}
}
