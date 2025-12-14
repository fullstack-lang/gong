// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Action:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Criticality", instanceWithInferedType.Criticality, instanceWithInferedType, probe.formStage, formGroup)

	case *models.Activities:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Criticality", instanceWithInferedType.Criticality, instanceWithInferedType, probe.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "State"
			rf.Fieldname = "Activities"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.State),
					"Activities",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.State](
					nil,
					"Activities",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Architecture:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("StateMachines", instanceWithInferedType, &instanceWithInferedType.StateMachines, formGroup, probe)
		AssociationSliceToForm("Roles", instanceWithInferedType, &instanceWithInferedType.Roles, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Diagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsEditable_", instanceWithInferedType.IsEditable_, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInRenameMode", instanceWithInferedType.IsInRenameMode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("State_Shapes", instanceWithInferedType, &instanceWithInferedType.State_Shapes, formGroup, probe)
		AssociationSliceToForm("Transition_Shapes", instanceWithInferedType, &instanceWithInferedType.Transition_Shapes, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "State"
			rf.Fieldname = "Diagrams"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.State),
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.State](
					nil,
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "StateMachine"
			rf.Fieldname = "Diagrams"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.StateMachine),
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.StateMachine](
					nil,
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Transition"
			rf.Fieldname = "Diagrams"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Transition),
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Transition](
					nil,
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Kill:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Message:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("MessageType", instanceWithInferedType.MessageType, formGroup, probe)
		AssociationFieldToForm("OriginTransition", instanceWithInferedType.OriginTransition, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Object"
			rf.Fieldname = "Messages"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Object),
					"Messages",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Object](
					nil,
					"Messages",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.MessageType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Transition"
			rf.Fieldname = "GeneratedMessages"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Transition),
					"GeneratedMessages",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Transition](
					nil,
					"GeneratedMessages",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Object:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("State", instanceWithInferedType.State, formGroup, probe)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Rank", instanceWithInferedType.Rank, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DOF", instanceWithInferedType.DOF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Messages", instanceWithInferedType, &instanceWithInferedType.Messages, formGroup, probe)

	case *models.Role:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Acronym", instanceWithInferedType.Acronym, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RolesWithSamePermissions", instanceWithInferedType, &instanceWithInferedType.RolesWithSamePermissions, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Architecture"
			rf.Fieldname = "Roles"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Architecture),
					"Roles",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Architecture](
					nil,
					"Roles",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Role"
			rf.Fieldname = "RolesWithSamePermissions"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Role),
					"RolesWithSamePermissions",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Role](
					nil,
					"RolesWithSamePermissions",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Transition"
			rf.Fieldname = "RolesWithPermissions"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Transition),
					"RolesWithPermissions",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Transition](
					nil,
					"RolesWithPermissions",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.State:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, false, 0)
		AssociationFieldToForm("Parent", instanceWithInferedType.Parent, formGroup, probe)
		BasicFieldtoForm("IsDecisionNode", instanceWithInferedType.IsDecisionNode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsFictif", instanceWithInferedType.IsFictif, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsEndState", instanceWithInferedType.IsEndState, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubStates", instanceWithInferedType, &instanceWithInferedType.SubStates, formGroup, probe)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		AssociationFieldToForm("Entry", instanceWithInferedType.Entry, formGroup, probe)
		AssociationSliceToForm("Activities", instanceWithInferedType, &instanceWithInferedType.Activities, formGroup, probe)
		AssociationFieldToForm("Exit", instanceWithInferedType.Exit, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "State"
			rf.Fieldname = "SubStates"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.State),
					"SubStates",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.State](
					nil,
					"SubStates",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "StateMachine"
			rf.Fieldname = "States"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.StateMachine),
					"States",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.StateMachine](
					nil,
					"States",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.StateMachine:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsNodeExpanded", instanceWithInferedType.IsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("States", instanceWithInferedType, &instanceWithInferedType.States, formGroup, probe)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		AssociationFieldToForm("InitialState", instanceWithInferedType.InitialState, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Architecture"
			rf.Fieldname = "StateMachines"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Architecture),
					"StateMachines",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Architecture](
					nil,
					"StateMachines",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.StateShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("State", instanceWithInferedType.State, formGroup, probe)
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
			rf.Fieldname = "State_Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"State_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"State_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Transition:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, false, 0)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, probe)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, probe)
		AssociationSliceToForm("RolesWithPermissions", instanceWithInferedType, &instanceWithInferedType.RolesWithPermissions, formGroup, probe)
		AssociationSliceToForm("GeneratedMessages", instanceWithInferedType, &instanceWithInferedType.GeneratedMessages, formGroup, probe)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)

	case *models.Transition_Shape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Transition", instanceWithInferedType.Transition, formGroup, probe)
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
			rf.Fieldname = "Transition_Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"Transition_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"Transition_Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
