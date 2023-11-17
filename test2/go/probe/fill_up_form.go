// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test2/go/models"
	"github.com/fullstack-lang/gong/test2/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.A:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("NumberField", instanceWithInferedType.NumberField, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationFieldToForm("B", instanceWithInferedType.B, formGroup, probe)
		AssociationSliceToForm("Bs", instanceWithInferedType, &instanceWithInferedType.Bs, formGroup, probe)

	case *models.B:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A"
			rf.Fieldname = "Bs"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A),
					"Bs",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A, *models.B](
					nil,
					"Bs",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	default:
		_ = instanceWithInferedType
	}
}
