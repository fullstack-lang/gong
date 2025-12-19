// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Product:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubProducts", instanceWithInferedType, &instanceWithInferedType.SubProducts, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Product"
			rf.Fieldname = "SubProducts"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Product),
					"SubProducts",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Product](
					nil,
					"SubProducts",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Project"
			rf.Fieldname = "RootProducts"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Project),
					"RootProducts",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Project](
					nil,
					"RootProducts",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Root"
			rf.Fieldname = "OrphanedProducts"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Root),
					"OrphanedProducts",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Root](
					nil,
					"OrphanedProducts",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Task"
			rf.Fieldname = "InputProducts"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Task),
					"InputProducts",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Task](
					nil,
					"InputProducts",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Task"
			rf.Fieldname = "OutputProducts"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Task),
					"OutputProducts",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Task](
					nil,
					"OutputProducts",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Project:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RootTasks", instanceWithInferedType, &instanceWithInferedType.RootTasks, formGroup, probe)
		AssociationSliceToForm("RootProducts", instanceWithInferedType, &instanceWithInferedType.RootProducts, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Root"
			rf.Fieldname = "Projects"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Root),
					"Projects",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Root](
					nil,
					"Projects",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Root:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Projects", instanceWithInferedType, &instanceWithInferedType.Projects, formGroup, probe)
		AssociationSliceToForm("OrphanedProducts", instanceWithInferedType, &instanceWithInferedType.OrphanedProducts, formGroup, probe)
		AssociationSliceToForm("OrphanedTasks", instanceWithInferedType, &instanceWithInferedType.OrphanedTasks, formGroup, probe)

	case *models.Task:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubTasks", instanceWithInferedType, &instanceWithInferedType.SubTasks, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("InputProducts", instanceWithInferedType, &instanceWithInferedType.InputProducts, formGroup, probe)
		BasicFieldtoForm("IsInputProducsNodeExpanded", instanceWithInferedType.IsInputProducsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("OutputProducts", instanceWithInferedType, &instanceWithInferedType.OutputProducts, formGroup, probe)
		BasicFieldtoForm("IsOutputProducsNodeExpanded", instanceWithInferedType.IsOutputProducsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Project"
			rf.Fieldname = "RootTasks"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Project),
					"RootTasks",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Project](
					nil,
					"RootTasks",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Root"
			rf.Fieldname = "OrphanedTasks"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Root),
					"OrphanedTasks",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Root](
					nil,
					"OrphanedTasks",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Task"
			rf.Fieldname = "SubTasks"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Task),
					"SubTasks",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Task](
					nil,
					"SubTasks",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
