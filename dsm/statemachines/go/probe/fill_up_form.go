// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/statemachines/go/models"
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
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Criticality", instanceWithInferedType.Criticality, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Activities:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Criticality", instanceWithInferedType.Criticality, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.State, *models.Activities](
				"State",
				"Activities",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.State) []*models.Activities {
					return owner.Activities
				})
		}

	case *models.Architecture:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("StateMachines", instanceWithInferedType, &instanceWithInferedType.StateMachines, formGroup, probe)
		AssociationSliceToForm("Roles", instanceWithInferedType, &instanceWithInferedType.Roles, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Diagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsEditable_", instanceWithInferedType.IsEditable_, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsStatesNodeExpanded", instanceWithInferedType.IsStatesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsNotesNodeExpanded", instanceWithInferedType.IsNotesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("NotesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.NotesWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("State_Shapes", instanceWithInferedType, &instanceWithInferedType.State_Shapes, formGroup, probe)
		AssociationSliceToForm("StatesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.StatesWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("Transition_Shapes", instanceWithInferedType, &instanceWithInferedType.Transition_Shapes, formGroup, probe)
		AssociationSliceToForm("Note_Shapes", instanceWithInferedType, &instanceWithInferedType.Note_Shapes, formGroup, probe)
		AssociationSliceToForm("NoteState_Shapes", instanceWithInferedType, &instanceWithInferedType.NoteState_Shapes, formGroup, probe)
		AssociationSliceToForm("NoteTransition_Shapes", instanceWithInferedType, &instanceWithInferedType.NoteTransition_Shapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.Diagram](
				"Library",
				"Diagrams",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Diagram {
					return owner.Diagrams
				})
		}
		{
			AssociationReverseSliceToForm[*models.State, *models.Diagram](
				"State",
				"Diagrams",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.State) []*models.Diagram {
					return owner.Diagrams
				})
		}
		{
			AssociationReverseSliceToForm[*models.StateMachine, *models.Diagram](
				"StateMachine",
				"Diagrams",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StateMachine) []*models.Diagram {
					return owner.Diagrams
				})
		}
		{
			AssociationReverseSliceToForm[*models.Transition, *models.Diagram](
				"Transition",
				"Diagrams",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Transition) []*models.Diagram {
					return owner.Diagrams
				})
		}

	case *models.Guard:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Kill:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Library:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LogoSVGFile", instanceWithInferedType.LogoSVGFile, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsRootLibrary", instanceWithInferedType.IsRootLibrary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		AssociationSliceToForm("RootStateMachines", instanceWithInferedType, &instanceWithInferedType.RootStateMachines, formGroup, probe)
		BasicFieldtoForm("IsStateMachinesNodeExpanded", instanceWithInferedType.IsStateMachinesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("StateMachinesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.StateMachinesWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("RootNotes", instanceWithInferedType, &instanceWithInferedType.RootNotes, formGroup, probe)
		BasicFieldtoForm("IsNotesNodeExpanded", instanceWithInferedType.IsNotesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("NotesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.NotesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsSubLibrariesNodeExpanded", instanceWithInferedType.IsSubLibrariesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubLibrariesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SubLibrariesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsExpandedTmp", instanceWithInferedType.IsExpandedTmp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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

	case *models.Message:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("MessageType", instanceWithInferedType.MessageType, formGroup, probe)
		AssociationFieldToForm("OriginTransition", instanceWithInferedType.OriginTransition, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Object, *models.Message](
				"Object",
				"Messages",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Object) []*models.Message {
					return owner.Messages
				})
		}

	case *models.MessageType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Transition, *models.MessageType](
				"Transition",
				"GeneratedMessages",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Transition) []*models.MessageType {
					return owner.GeneratedMessages
				})
		}

	case *models.Note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("States", instanceWithInferedType, &instanceWithInferedType.States, formGroup, probe)
		AssociationSliceToForm("Transitions", instanceWithInferedType, &instanceWithInferedType.Transitions, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Note](
				"Diagram",
				"NotesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Note {
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
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		BasicFieldtoForm("OverideLayoutDirection", instanceWithInferedType.OverideLayoutDirection, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.NoteShape](
				"Diagram",
				"Note_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.NoteShape {
					return owner.Note_Shapes
				})
		}

	case *models.NoteStateShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("State", instanceWithInferedType.State, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.NoteStateShape](
				"Diagram",
				"NoteState_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.NoteStateShape {
					return owner.NoteState_Shapes
				})
		}

	case *models.NoteTransitionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Transition", instanceWithInferedType.Transition, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.NoteTransitionShape](
				"Diagram",
				"NoteTransition_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.NoteTransitionShape {
					return owner.NoteTransition_Shapes
				})
		}

	case *models.Object:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("State", instanceWithInferedType.State, formGroup, probe)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Rank", instanceWithInferedType.Rank, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DOF", instanceWithInferedType.DOF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Messages", instanceWithInferedType, &instanceWithInferedType.Messages, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Role:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Acronym", instanceWithInferedType.Acronym, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("RolesWithSamePermissions", instanceWithInferedType, &instanceWithInferedType.RolesWithSamePermissions, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Architecture, *models.Role](
				"Architecture",
				"Roles",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Architecture) []*models.Role {
					return owner.Roles
				})
		}
		{
			AssociationReverseSliceToForm[*models.Role, *models.Role](
				"Role",
				"RolesWithSamePermissions",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Role) []*models.Role {
					return owner.RolesWithSamePermissions
				})
		}
		{
			AssociationReverseSliceToForm[*models.Transition, *models.Role](
				"Transition",
				"RolesWithPermissions",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Transition) []*models.Role {
					return owner.RolesWithPermissions
				})
		}

	case *models.State:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		AssociationFieldToForm("Parent", instanceWithInferedType.Parent, formGroup, probe)
		BasicFieldtoForm("IsDecisionNode", instanceWithInferedType.IsDecisionNode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsFictious", instanceWithInferedType.IsFictious, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsEndState", instanceWithInferedType.IsEndState, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubStates", instanceWithInferedType, &instanceWithInferedType.SubStates, formGroup, probe)
		AssociationFieldToForm("Entry", instanceWithInferedType.Entry, formGroup, probe)
		AssociationSliceToForm("Activities", instanceWithInferedType, &instanceWithInferedType.Activities, formGroup, probe)
		AssociationFieldToForm("Exit", instanceWithInferedType.Exit, formGroup, probe)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.State](
				"Diagram",
				"StatesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.State {
					return owner.StatesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.State](
				"Note",
				"States",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.State {
					return owner.States
				})
		}
		{
			AssociationReverseSliceToForm[*models.State, *models.State](
				"State",
				"SubStates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.State) []*models.State {
					return owner.SubStates
				})
		}
		{
			AssociationReverseSliceToForm[*models.StateMachine, *models.State](
				"StateMachine",
				"States",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StateMachine) []*models.State {
					return owner.States
				})
		}

	case *models.StateMachine:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		AssociationFieldToForm("InitialState", instanceWithInferedType.InitialState, formGroup, probe)
		AssociationSliceToForm("States", instanceWithInferedType, &instanceWithInferedType.States, formGroup, probe)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		BasicFieldtoForm("IsWithTransitionNameAutonamticalyGenerated", instanceWithInferedType.IsWithTransitionNameAutonamticalyGenerated, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Architecture, *models.StateMachine](
				"Architecture",
				"StateMachines",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Architecture) []*models.StateMachine {
					return owner.StateMachines
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.StateMachine](
				"Library",
				"RootStateMachines",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.StateMachine {
					return owner.RootStateMachines
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.StateMachine](
				"Library",
				"StateMachinesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.StateMachine {
					return owner.StateMachinesWhoseNodeIsExpanded
				})
		}

	case *models.StateShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("State", instanceWithInferedType.State, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.StateShape](
				"Diagram",
				"State_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.StateShape {
					return owner.State_Shapes
				})
		}

	case *models.Transition:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, false, 0, false)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, probe)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, probe)
		AssociationSliceToForm("RolesWithPermissions", instanceWithInferedType, &instanceWithInferedType.RolesWithPermissions, formGroup, probe)
		AssociationSliceToForm("GeneratedMessages", instanceWithInferedType, &instanceWithInferedType.GeneratedMessages, formGroup, probe)
		AssociationFieldToForm("Guard", instanceWithInferedType.Guard, formGroup, probe)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Note, *models.Transition](
				"Note",
				"Transitions",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Transition {
					return owner.Transitions
				})
		}

	case *models.Transition_Shape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Transition", instanceWithInferedType.Transition, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Transition_Shape](
				"Diagram",
				"Transition_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Transition_Shape {
					return owner.Transition_Shapes
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
