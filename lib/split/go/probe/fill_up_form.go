// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.AsSplit:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("AsSplitAreas", instanceWithInferedType, &instanceWithInferedType.AsSplitAreas, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AsSplitArea"
			rf.Fieldname = "AsSplits"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.AsSplitArea),
					"AsSplits",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.AsSplitArea, *models.AsSplit](
					nil,
					"AsSplits",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.AsSplitArea:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsAny", instanceWithInferedType.IsAny, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("AsSplits", instanceWithInferedType, &instanceWithInferedType.AsSplits, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AsSplit"
			rf.Fieldname = "AsSplitAreas"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.AsSplit),
					"AsSplitAreas",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.AsSplit, *models.AsSplitArea](
					nil,
					"AsSplitAreas",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "View"
			rf.Fieldname = "RootAsSplitAreas"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.View),
					"RootAsSplitAreas",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.View, *models.AsSplitArea](
					nil,
					"RootAsSplitAreas",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.View:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RootAsSplitAreas", instanceWithInferedType, &instanceWithInferedType.RootAsSplitAreas, formGroup, probe)

	default:
		_ = instanceWithInferedType
	}
}
