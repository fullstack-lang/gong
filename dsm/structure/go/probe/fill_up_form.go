// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.AllocatedProcessShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Participant", instanceWithInferedType.Participant, formGroup, probe)
		AssociationFieldToForm("Process", instanceWithInferedType.Process, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.AllocatedProcessShape](
				"DiagramProcess",
				"AllocatedProcessShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.AllocatedProcessShape {
					return owner.AllocatedProcessShapes
				})
		}

	case *models.AllocatedResourceShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Participant", instanceWithInferedType.Participant, formGroup, probe)
		AssociationFieldToForm("Resource", instanceWithInferedType.Resource, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.AllocatedResourceShape](
				"DiagramProcess",
				"AllocatedResourceShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.AllocatedResourceShape {
					return owner.AllocatedResourceShapes
				})
		}

	case *models.ControlFlow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, probe)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.ControlFlow](
				"DiagramProcess",
				"ControlFlowsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.ControlFlow {
					return owner.ControlFlowsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Participant, *models.ControlFlow](
				"Participant",
				"ControlFlows",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Participant) []*models.ControlFlow {
					return owner.ControlFlows
				})
		}

	case *models.ControlFlowShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ControlFlow", instanceWithInferedType.ControlFlow, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.ControlFlowShape](
				"DiagramProcess",
				"ControlFlow_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.ControlFlowShape {
					return owner.ControlFlow_Shapes
				})
		}

	case *models.Data:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Acronym", instanceWithInferedType.Acronym, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("SVG_Path", instanceWithInferedType.SVG_Path, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)
		BasicFieldtoForm("InverseAppliedScaling", instanceWithInferedType.InverseAppliedScaling, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DataFlow, *models.Data](
				"DataFlow",
				"Datas",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DataFlow) []*models.Data {
					return owner.Datas
				})
		}
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.Data](
				"DiagramProcess",
				"DatasWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.Data {
					return owner.DatasWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Data](
				"Library",
				"RootDatas",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Data {
					return owner.RootDatas
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Data](
				"Library",
				"DatasWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Data {
					return owner.DatasWhoseNodeIsExpanded
				})
		}

	case *models.DataFlow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Datas", instanceWithInferedType, &instanceWithInferedType.Datas, formGroup, probe)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("StartTask", instanceWithInferedType.StartTask, formGroup, probe)
		AssociationFieldToForm("EndTask", instanceWithInferedType.EndTask, formGroup, probe)
		AssociationFieldToForm("StartExternalParticipant", instanceWithInferedType.StartExternalParticipant, formGroup, probe)
		AssociationFieldToForm("EndExternalParticipant", instanceWithInferedType.EndExternalParticipant, formGroup, probe)
		BasicFieldtoForm("IsDatasNodeExpanded", instanceWithInferedType.IsDatasNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.DataFlow](
				"DiagramProcess",
				"DataFlowsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.DataFlow {
					return owner.DataFlowsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.DataFlow](
				"DiagramProcess",
				"DataFlowsWhoseDataNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.DataFlow {
					return owner.DataFlowsWhoseDataNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.DataFlow](
				"Library",
				"RootDataFlows",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.DataFlow {
					return owner.RootDataFlows
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.DataFlow](
				"Library",
				"DataFlowsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.DataFlow {
					return owner.DataFlowsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Process, *models.DataFlow](
				"Process",
				"DataFlows",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Process) []*models.DataFlow {
					return owner.DataFlows
				})
		}

	case *models.DataFlowShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DataFlow", instanceWithInferedType.DataFlow, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.DataFlowShape](
				"DiagramProcess",
				"DataFlow_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.DataFlowShape {
					return owner.DataFlow_Shapes
				})
		}

	case *models.DataShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Data", instanceWithInferedType.Data, formGroup, probe)
		AssociationFieldToForm("DataFlow", instanceWithInferedType.DataFlow, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.DataShape](
				"DiagramProcess",
				"Data_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.DataShape {
					return owner.Data_Shapes
				})
		}

	case *models.DiagramProcess:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
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
		BasicFieldtoForm("IsProcesssNodeExpanded", instanceWithInferedType.IsProcesssNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ProcesssWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ProcesssWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("Participant_Shapes", instanceWithInferedType, &instanceWithInferedType.Participant_Shapes, formGroup, probe)
		BasicFieldtoForm("IsParticipantsNodeExpanded", instanceWithInferedType.IsParticipantsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ParticipantWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ParticipantWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ExternalParticipant_Shapes", instanceWithInferedType, &instanceWithInferedType.ExternalParticipant_Shapes, formGroup, probe)
		BasicFieldtoForm("IsExternalParticipantsNodeExpanded", instanceWithInferedType.IsExternalParticipantsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ExternalParticipantWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ExternalParticipantWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ExternalParticipantsWhoseInDataFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("TasksWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TasksWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("Task_Shapes", instanceWithInferedType, &instanceWithInferedType.Task_Shapes, formGroup, probe)
		AssociationSliceToForm("ControlFlowsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ControlFlowsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ControlFlow_Shapes", instanceWithInferedType, &instanceWithInferedType.ControlFlow_Shapes, formGroup, probe)
		AssociationSliceToForm("DataFlowsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DataFlowsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("DataFlow_Shapes", instanceWithInferedType, &instanceWithInferedType.DataFlow_Shapes, formGroup, probe)
		AssociationSliceToForm("DatasWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DatasWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("Data_Shapes", instanceWithInferedType, &instanceWithInferedType.Data_Shapes, formGroup, probe)
		AssociationSliceToForm("DataFlowsWhoseDataNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DataFlowsWhoseDataNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("AllocatedResourcesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.AllocatedResourcesWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("AllocatedResourceShapes", instanceWithInferedType, &instanceWithInferedType.AllocatedResourceShapes, formGroup, probe)
		AssociationSliceToForm("AllocatedProcessesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.AllocatedProcessesWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("AllocatedProcessShapes", instanceWithInferedType, &instanceWithInferedType.AllocatedProcessShapes, formGroup, probe)
		AssociationSliceToForm("Note_Shapes", instanceWithInferedType, &instanceWithInferedType.Note_Shapes, formGroup, probe)
		AssociationSliceToForm("NotesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.NotesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsNotesNodeExpanded", instanceWithInferedType.IsNotesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("NoteTaskShapes", instanceWithInferedType, &instanceWithInferedType.NoteTaskShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Process, *models.DiagramProcess](
				"Process",
				"DiagramProcesss",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Process) []*models.DiagramProcess {
					return owner.DiagramProcesss
				})
		}
		{
			AssociationReverseSliceToForm[*models.Process, *models.DiagramProcess](
				"Process",
				"DiagramProcessWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Process) []*models.DiagramProcess {
					return owner.DiagramProcessWhoseNodeIsExpanded
				})
		}

	case *models.ExternalParticipantShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Participant", instanceWithInferedType.Participant, formGroup, probe)
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
		BasicFieldtoForm("TailHeigth", instanceWithInferedType.TailHeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.ExternalParticipantShape](
				"DiagramProcess",
				"ExternalParticipant_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.ExternalParticipantShape {
					return owner.ExternalParticipant_Shapes
				})
		}

	case *models.Library:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsRootLibrary", instanceWithInferedType.IsRootLibrary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		BasicFieldtoForm("IsSubLibrariesNodeExpanded", instanceWithInferedType.IsSubLibrariesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubLibrariesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SubLibrariesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LogoSVGFile", instanceWithInferedType.LogoSVGFile, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)
		AssociationSliceToForm("RootProcesses", instanceWithInferedType, &instanceWithInferedType.RootProcesses, formGroup, probe)
		BasicFieldtoForm("IsProcessesNodeExpanded", instanceWithInferedType.IsProcessesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ProcesssWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ProcesssWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("RootDataFlows", instanceWithInferedType, &instanceWithInferedType.RootDataFlows, formGroup, probe)
		BasicFieldtoForm("IsDataFlowsNodeExpanded", instanceWithInferedType.IsDataFlowsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DataFlowsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DataFlowsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("RootDatas", instanceWithInferedType, &instanceWithInferedType.RootDatas, formGroup, probe)
		BasicFieldtoForm("IsDatasNodeExpanded", instanceWithInferedType.IsDatasNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DatasWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DatasWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("RootResources", instanceWithInferedType, &instanceWithInferedType.RootResources, formGroup, probe)
		BasicFieldtoForm("IsResourcesNodeExpanded", instanceWithInferedType.IsResourcesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ResourcesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ResourcesWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ParticipantsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ParticipantsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("RootNotes", instanceWithInferedType, &instanceWithInferedType.RootNotes, formGroup, probe)
		BasicFieldtoForm("IsNotesNodeExpanded", instanceWithInferedType.IsNotesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("NotesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.NotesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsExpandedTmp", instanceWithInferedType.IsExpandedTmp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.Library](
				"Library",
				"SubLibraries",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Library {
					return owner.SubLibraries
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Library](
				"Library",
				"SubLibrariesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Library {
					return owner.SubLibrariesWhoseNodeIsExpanded
				})
		}

	case *models.Note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsTasksNodeExpanded", instanceWithInferedType.IsTasksNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Tasks", instanceWithInferedType, &instanceWithInferedType.Tasks, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.Note](
				"DiagramProcess",
				"NotesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.Note {
					return owner.NotesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Note](
				"Library",
				"RootNotes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Note {
					return owner.RootNotes
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Note](
				"Library",
				"NotesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Note {
					return owner.NotesWhoseNodeIsExpanded
				})
		}

	case *models.NoteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.NoteShape](
				"DiagramProcess",
				"Note_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.NoteShape {
					return owner.Note_Shapes
				})
		}

	case *models.NoteTaskShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Task", instanceWithInferedType.Task, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.NoteTaskShape](
				"DiagramProcess",
				"NoteTaskShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.NoteTaskShape {
					return owner.NoteTaskShapes
				})
		}

	case *models.Participant:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsProcessResource", instanceWithInferedType.IsProcessResource, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("Resources", instanceWithInferedType, &instanceWithInferedType.Resources, formGroup, probe)
		BasicFieldtoForm("IsResourcesNodeExpanded", instanceWithInferedType.IsResourcesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Processes", instanceWithInferedType, &instanceWithInferedType.Processes, formGroup, probe)
		BasicFieldtoForm("IsProcessesNodeExpanded", instanceWithInferedType.IsProcessesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsTasksNodeExpanded", instanceWithInferedType.IsTasksNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Tasks", instanceWithInferedType, &instanceWithInferedType.Tasks, formGroup, probe)
		BasicFieldtoForm("IsControlFlowsNodeExpanded", instanceWithInferedType.IsControlFlowsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlFlows", instanceWithInferedType, &instanceWithInferedType.ControlFlows, formGroup, probe)
		AssociationSliceToForm("TaskWhoseOutControlFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TaskWhoseOutControlFlowsNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("TaskWhoseInControlFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TaskWhoseInControlFlowsNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsDataFlowsNodeExpanded", instanceWithInferedType.IsDataFlowsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("TaskWhoseOutDataFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TaskWhoseOutDataFlowsNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("TaskWhoseInDataFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TaskWhoseInDataFlowsNodeIsExpanded, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.Participant](
				"DiagramProcess",
				"ParticipantWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.Participant {
					return owner.ParticipantWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.Participant](
				"DiagramProcess",
				"ExternalParticipantWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.Participant {
					return owner.ExternalParticipantWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.Participant](
				"DiagramProcess",
				"ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.Participant {
					return owner.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.Participant](
				"DiagramProcess",
				"ExternalParticipantsWhoseInDataFlowsNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.Participant {
					return owner.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Participant](
				"Library",
				"ParticipantsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Participant {
					return owner.ParticipantsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Process, *models.Participant](
				"Process",
				"Participants",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Process) []*models.Participant {
					return owner.Participants
				})
		}
		{
			AssociationReverseSliceToForm[*models.Process, *models.Participant](
				"Process",
				"ParticipantWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Process) []*models.Participant {
					return owner.ParticipantWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Process, *models.Participant](
				"Process",
				"ExternalParticipants",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Process) []*models.Participant {
					return owner.ExternalParticipants
				})
		}
		{
			AssociationReverseSliceToForm[*models.Process, *models.Participant](
				"Process",
				"ExternalParticipantWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Process) []*models.Participant {
					return owner.ExternalParticipantWhoseNodeIsExpanded
				})
		}

	case *models.ParticipantShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Participant", instanceWithInferedType.Participant, formGroup, probe)
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
		BasicFieldtoForm("WidthWeight", instanceWithInferedType.WidthWeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.ParticipantShape](
				"DiagramProcess",
				"Participant_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.ParticipantShape {
					return owner.Participant_Shapes
				})
		}

	case *models.Process:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("SVG_Path", instanceWithInferedType.SVG_Path, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)
		BasicFieldtoForm("InverseAppliedScaling", instanceWithInferedType.InverseAppliedScaling, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DiagramProcesss", instanceWithInferedType, &instanceWithInferedType.DiagramProcesss, formGroup, probe)
		AssociationSliceToForm("DiagramProcessWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DiagramProcessWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsSubProcessNodeExpanded", instanceWithInferedType.IsSubProcessNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubProcesses", instanceWithInferedType, &instanceWithInferedType.SubProcesses, formGroup, probe)
		AssociationSliceToForm("Participants", instanceWithInferedType, &instanceWithInferedType.Participants, formGroup, probe)
		AssociationSliceToForm("ParticipantWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ParticipantWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("DataFlows", instanceWithInferedType, &instanceWithInferedType.DataFlows, formGroup, probe)
		BasicFieldtoForm("IsDataFlowsNodeExpanded", instanceWithInferedType.IsDataFlowsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ExternalParticipants", instanceWithInferedType, &instanceWithInferedType.ExternalParticipants, formGroup, probe)
		AssociationSliceToForm("ExternalParticipantWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ExternalParticipantWhoseNodeIsExpanded, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.Process](
				"DiagramProcess",
				"ProcesssWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.Process {
					return owner.ProcesssWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.Process](
				"DiagramProcess",
				"AllocatedProcessesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.Process {
					return owner.AllocatedProcessesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Process](
				"Library",
				"RootProcesses",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Process {
					return owner.RootProcesses
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Process](
				"Library",
				"ProcesssWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Process {
					return owner.ProcesssWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Participant, *models.Process](
				"Participant",
				"Processes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Participant) []*models.Process {
					return owner.Processes
				})
		}
		{
			AssociationReverseSliceToForm[*models.Process, *models.Process](
				"Process",
				"SubProcesses",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Process) []*models.Process {
					return owner.SubProcesses
				})
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.ProcessShape](
				"DiagramProcess",
				"Process_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.ProcessShape {
					return owner.Process_Shapes
				})
		}

	case *models.Resource:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Acronym", instanceWithInferedType.Acronym, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("SVG_Path", instanceWithInferedType.SVG_Path, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)
		BasicFieldtoForm("InverseAppliedScaling", instanceWithInferedType.InverseAppliedScaling, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.Resource](
				"DiagramProcess",
				"AllocatedResourcesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.Resource {
					return owner.AllocatedResourcesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Resource](
				"Library",
				"RootResources",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Resource {
					return owner.RootResources
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Resource](
				"Library",
				"ResourcesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Resource {
					return owner.ResourcesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Participant, *models.Resource](
				"Participant",
				"Resources",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Participant) []*models.Resource {
					return owner.Resources
				})
		}

	case *models.Task:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsStartTask", instanceWithInferedType.IsStartTask, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsEndTask", instanceWithInferedType.IsEndTask, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Type", instanceWithInferedType.Type, formGroup, probe)
		BasicFieldtoForm("IsTaskNameNotProcessName", instanceWithInferedType.IsTaskNameNotProcessName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.Task](
				"DiagramProcess",
				"TasksWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.Task {
					return owner.TasksWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Task](
				"Note",
				"Tasks",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Task {
					return owner.Tasks
				})
		}
		{
			AssociationReverseSliceToForm[*models.Participant, *models.Task](
				"Participant",
				"Tasks",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Participant) []*models.Task {
					return owner.Tasks
				})
		}
		{
			AssociationReverseSliceToForm[*models.Participant, *models.Task](
				"Participant",
				"TaskWhoseOutControlFlowsNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Participant) []*models.Task {
					return owner.TaskWhoseOutControlFlowsNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Participant, *models.Task](
				"Participant",
				"TaskWhoseInControlFlowsNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Participant) []*models.Task {
					return owner.TaskWhoseInControlFlowsNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Participant, *models.Task](
				"Participant",
				"TaskWhoseOutDataFlowsNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Participant) []*models.Task {
					return owner.TaskWhoseOutDataFlowsNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Participant, *models.Task](
				"Participant",
				"TaskWhoseInDataFlowsNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Participant) []*models.Task {
					return owner.TaskWhoseInDataFlowsNodeIsExpanded
				})
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
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramProcess, *models.TaskShape](
				"DiagramProcess",
				"Task_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramProcess) []*models.TaskShape {
					return owner.Task_Shapes
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
