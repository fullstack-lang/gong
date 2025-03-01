// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
	"github.com/fullstack-lang/gong/lib/slider/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Checkbox:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ValueBool", instanceWithInferedType.ValueBool, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LabelForTrue", instanceWithInferedType.LabelForTrue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LabelForFalse", instanceWithInferedType.LabelForFalse, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Checkboxes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Group),
					"Checkboxes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Group, *models.Checkbox](
					nil,
					"Checkboxes",
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
		AssociationSliceToForm("Sliders", instanceWithInferedType, &instanceWithInferedType.Sliders, formGroup, probe)
		AssociationSliceToForm("Checkboxes", instanceWithInferedType, &instanceWithInferedType.Checkboxes, formGroup, probe)
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

	case *models.Slider:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsFloat64", instanceWithInferedType.IsFloat64, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInt", instanceWithInferedType.IsInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinInt", instanceWithInferedType.MinInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxInt", instanceWithInferedType.MaxInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StepInt", instanceWithInferedType.StepInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ValueInt", instanceWithInferedType.ValueInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinFloat64", instanceWithInferedType.MinFloat64, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxFloat64", instanceWithInferedType.MaxFloat64, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StepFloat64", instanceWithInferedType.StepFloat64, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ValueFloat64", instanceWithInferedType.ValueFloat64, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Sliders"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Group),
					"Sliders",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Group, *models.Slider](
					nil,
					"Sliders",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
