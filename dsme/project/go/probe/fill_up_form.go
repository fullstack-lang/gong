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
	case *models.Diagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsEditable_", instanceWithInferedType.IsEditable_, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInRenameMode", instanceWithInferedType.IsInRenameMode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Product_Shapes", instanceWithInferedType, &instanceWithInferedType.Product_Shapes, formGroup, probe)
		AssociationSliceToForm("ProductsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ProductsWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsPBSNodeExpanded", instanceWithInferedType.IsPBSNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ProductComposition_Shapes", instanceWithInferedType, &instanceWithInferedType.ProductComposition_Shapes, formGroup, probe)
		BasicFieldtoForm("IsWBSNodeExpanded", instanceWithInferedType.IsWBSNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Task_Shapes", instanceWithInferedType, &instanceWithInferedType.Task_Shapes, formGroup, probe)
		AssociationSliceToForm("TasksWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TasksWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("TasksWhoseInputNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TasksWhoseInputNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("TaskComposition_Shapes", instanceWithInferedType, &instanceWithInferedType.TaskComposition_Shapes, formGroup, probe)
		AssociationSliceToForm("TaskInputShapes", instanceWithInferedType, &instanceWithInferedType.TaskInputShapes, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Project"
			rf.Fieldname = "Diagrams"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Project),
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Project](
					nil,
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Product:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubProducts", instanceWithInferedType, &instanceWithInferedType.SubProducts, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsProducersNodeExpanded", instanceWithInferedType.IsProducersNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsConsumersNodeExpanded", instanceWithInferedType.IsConsumersNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "ProductsWhoseNodeIsExpanded"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"ProductsWhoseNodeIsExpanded",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"ProductsWhoseNodeIsExpanded",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
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
			rf.Fieldname = "Inputs"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Task),
					"Inputs",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Task](
					nil,
					"Inputs",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Task"
			rf.Fieldname = "Outputs"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Task),
					"Outputs",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Task](
					nil,
					"Outputs",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ProductCompositionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Product", instanceWithInferedType.Product, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "ProductComposition_Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"ProductComposition_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"ProductComposition_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ProductShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Product", instanceWithInferedType.Product, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "Product_Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"Product_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"Product_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Project:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RootProducts", instanceWithInferedType, &instanceWithInferedType.RootProducts, formGroup, probe)
		BasicFieldtoForm("IsPBSNodeExpanded", instanceWithInferedType.IsPBSNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RootTasks", instanceWithInferedType, &instanceWithInferedType.RootTasks, formGroup, probe)
		BasicFieldtoForm("IsWBSNodeExpanded", instanceWithInferedType.IsWBSNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		BasicFieldtoForm("IsDiagramsNodeExpanded", instanceWithInferedType.IsDiagramsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
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
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Task:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubTasks", instanceWithInferedType, &instanceWithInferedType.SubTasks, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Inputs", instanceWithInferedType, &instanceWithInferedType.Inputs, formGroup, probe)
		BasicFieldtoForm("IsInputsNodeExpanded", instanceWithInferedType.IsInputsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Outputs", instanceWithInferedType, &instanceWithInferedType.Outputs, formGroup, probe)
		BasicFieldtoForm("IsOutputsNodeExpanded", instanceWithInferedType.IsOutputsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "TasksWhoseNodeIsExpanded"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"TasksWhoseNodeIsExpanded",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"TasksWhoseNodeIsExpanded",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "TasksWhoseInputNodeIsExpanded"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"TasksWhoseInputNodeIsExpanded",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"TasksWhoseInputNodeIsExpanded",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
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

	case *models.TaskCompositionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Task", instanceWithInferedType.Task, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "TaskComposition_Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"TaskComposition_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"TaskComposition_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.TaskInputShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Task", instanceWithInferedType.Task, formGroup, probe)
		AssociationFieldToForm("Product", instanceWithInferedType.Product, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "TaskInputShapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"TaskInputShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"TaskInputShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.TaskShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Task", instanceWithInferedType.Task, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "Task_Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"Task_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"Task_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
