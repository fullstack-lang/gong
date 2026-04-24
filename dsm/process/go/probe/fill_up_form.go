// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
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
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInRenameMode", instanceWithInferedType.IsInRenameMode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsEditable_", instanceWithInferedType.IsEditable_, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsShowPrefix", instanceWithInferedType.IsShowPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DefaultBoxWidth", instanceWithInferedType.DefaultBoxWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DefaultBoxHeigth", instanceWithInferedType.DefaultBoxHeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Process_Shapes", instanceWithInferedType, &instanceWithInferedType.Process_Shapes, formGroup, probe)
		AssociationSliceToForm("ProcesssWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ProcesssWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsProcesssNodeExpanded", instanceWithInferedType.IsProcesssNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ProcessComposition_Shapes", instanceWithInferedType, &instanceWithInferedType.ProcessComposition_Shapes, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Library"
			rf.Fieldname = "Diagrams"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Library),
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Library](
					nil,
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Library:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInRenameMode", instanceWithInferedType.IsInRenameMode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LogoSVGFile", instanceWithInferedType.LogoSVGFile, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Library"
			rf.Fieldname = "SubLibraries"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Library),
					"SubLibraries",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Library](
					nil,
					"SubLibraries",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Process:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInRenameMode", instanceWithInferedType.IsInRenameMode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubProcesses", instanceWithInferedType, &instanceWithInferedType.SubProcesses, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"ProcesssWhoseNodeIsExpanded",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"ProcesssWhoseNodeIsExpanded",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Process"
			rf.Fieldname = "SubProcesses"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Process),
					"SubProcesses",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Process](
					nil,
					"SubProcesses",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ProcessCompositionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Process", instanceWithInferedType.Process, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "ProcessComposition_Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"ProcessComposition_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"ProcessComposition_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ProcessShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Process", instanceWithInferedType.Process, formGroup, probe)
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
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "Process_Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"Process_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"Process_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
