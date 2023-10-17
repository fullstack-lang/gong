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
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.A:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("B", instanceWithInferedType.B, formGroup, playground)
		AssociationSliceToForm("Bs", instanceWithInferedType, &instanceWithInferedType.Bs, formGroup, playground)

	case *models.B:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A"
			rf.Fieldname = "Bs"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A),
					"Bs",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.A, *models.B](
					nil,
					"Bs",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	default:
		_ = instanceWithInferedType
	}
}
