// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/statemachines/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.Action:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Criticality", inst.Criticality, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.States {
				if src.Entry == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.State", "Entry", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.States {
				if src.Exit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.State", "Exit", refNames, formGroup, probe.formStage)
		}
	case *models.Activities:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Criticality", inst.Criticality, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.States {
				for _, target := range src.Activities {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.State", "Activities", refNames, formGroup, probe.formStage)
		}
	case *models.Diagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsEditable_", inst.IsEditable_, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsStatesNodeExpanded", inst.IsStatesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: State_Shapes
			div := (&form.FormDiv{Name: "State_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.State_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "State_Shapes",
				Label: "State_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: StatesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "StatesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.StatesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "StatesWhoseNodeIsExpanded",
				Label: "StatesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Transition_Shapes
			div := (&form.FormDiv{Name: "Transition_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Transition_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Transition_Shapes",
				Label: "Transition_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Note_Shapes
			div := (&form.FormDiv{Name: "Note_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Note_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Note_Shapes",
				Label: "Note_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: NoteState_Shapes
			div := (&form.FormDiv{Name: "NoteState_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NoteState_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NoteState_Shapes",
				Label: "NoteState_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ShowRoles", inst.ShowRoles, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowMessages", inst.ShowMessages, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.Diagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "Diagrams", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.States {
				for _, target := range src.Diagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.State", "Diagrams", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StateMachines {
				for _, target := range src.Diagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StateMachine", "Diagrams", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Transitions {
				for _, target := range src.Diagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Transition", "Diagrams", refNames, formGroup, probe.formStage)
		}
	case *models.Guard:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Transitions {
				if src.Guard == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Transition", "Guard", refNames, formGroup, probe.formStage)
		}
	case *models.Kill:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	case *models.Library:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: SubLibraries
			div := (&form.FormDiv{Name: "SubLibraries"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubLibraries {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubLibraries",
				Label: "SubLibraries",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("NbPixPerCharacter", inst.NbPixPerCharacter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LogoSVGFile", inst.LogoSVGFile, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsRootLibrary", inst.IsRootLibrary, probe.formStage, formGroup)

		{
			// Slice of pointers: Diagrams
			div := (&form.FormDiv{Name: "Diagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Diagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Diagrams",
				Label: "Diagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootStateMachines
			div := (&form.FormDiv{Name: "RootStateMachines"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootStateMachines {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootStateMachines",
				Label: "RootStateMachines",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsStateMachinesNodeExpanded", inst.IsStateMachinesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: StateMachinesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "StateMachinesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.StateMachinesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "StateMachinesWhoseNodeIsExpanded",
				Label: "StateMachinesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsSubLibrariesNodeExpanded", inst.IsSubLibrariesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: SubLibrariesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "SubLibrariesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubLibrariesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubLibrariesWhoseNodeIsExpanded",
				Label: "SubLibrariesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsExpandedTmp", inst.IsExpandedTmp, probe.formStage, formGroup)

		{
			// Slice of pointers: Roles
			div := (&form.FormDiv{Name: "Roles"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Roles {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Roles",
				Label: "Roles",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsRolesNodeExpanded", inst.IsRolesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: MessageTypes
			div := (&form.FormDiv{Name: "MessageTypes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MessageTypes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MessageTypes",
				Label: "MessageTypes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsMessageTypesNodeExpanded", inst.IsMessageTypesNodeExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.SubLibraries {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "SubLibraries", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.SubLibrariesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "SubLibrariesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}
	case *models.Message:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsSelected", inst.IsSelected, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("MessageType", inst.MessageType, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.MessageType](), probe.formStage)
		StageSetAssociationFieldToForm("OriginTransition", inst.OriginTransition, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Transition](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Objects {
				for _, target := range src.Messages {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Object", "Messages", refNames, formGroup, probe.formStage)
		}
	case *models.MessageType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.MessageTypes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "MessageTypes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Messages {
				if src.MessageType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Message", "MessageType", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Transitions {
				for _, target := range src.GeneratedMessages {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Transition", "GeneratedMessages", refNames, formGroup, probe.formStage)
		}
	case *models.Note:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("State", inst.State, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.State](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteShapes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteShape", "Note", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteStateShapes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteStateShape", "Note", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.States {
				for _, target := range src.Notes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.State", "Notes", refNames, formGroup, probe.formStage)
		}
	case *models.NoteShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetBasicFieldtoForm("IsLayoutDirectionDifferent", inst.IsLayoutDirectionDifferent, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Note_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Note_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.NoteStateShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetAssociationFieldToForm("State", inst.State, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.State](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.NoteState_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "NoteState_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Object:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("State", inst.State, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.State](), probe.formStage)
		StageSetBasicFieldtoForm("IsSelected", inst.IsSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rank", inst.Rank, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DOF", inst.DOF, probe.formStage, formGroup)

		{
			// Slice of pointers: Messages
			div := (&form.FormDiv{Name: "Messages"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Messages {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Messages",
				Label: "Messages",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *models.Role:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Acronym", inst.Acronym, probe.formStage, formGroup)

		{
			// Slice of pointers: RolesWithSamePermissions
			div := (&form.FormDiv{Name: "RolesWithSamePermissions"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RolesWithSamePermissions {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RolesWithSamePermissions",
				Label: "RolesWithSamePermissions",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.Roles {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "Roles", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Roles {
				for _, target := range src.RolesWithSamePermissions {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Role", "RolesWithSamePermissions", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Transitions {
				for _, target := range src.RolesWithPermissions {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Transition", "RolesWithPermissions", refNames, formGroup, probe.formStage)
		}
	case *models.State:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsEndState", inst.IsEndState, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsDecisionNode", inst.IsDecisionNode, probe.formStage, formGroup)

		{
			// Slice of pointers: SubStates
			div := (&form.FormDiv{Name: "SubStates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubStates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubStates",
				Label: "SubStates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Entry", inst.Entry, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Action](), probe.formStage)

		{
			// Slice of pointers: Activities
			div := (&form.FormDiv{Name: "Activities"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Activities {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Activities",
				Label: "Activities",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Exit", inst.Exit, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Action](), probe.formStage)
		StageSetAssociationFieldToForm("Parent", inst.Parent, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.State](), probe.formStage)
		StageSetBasicFieldtoForm("IsFictious", inst.IsFictious, probe.formStage, formGroup)

		{
			// Slice of pointers: Diagrams
			div := (&form.FormDiv{Name: "Diagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Diagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Diagrams",
				Label: "Diagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Notes
			div := (&form.FormDiv{Name: "Notes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Notes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Notes",
				Label: "Notes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.StatesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "StatesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				if src.State == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "State", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteStateShapes {
				if src.State == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteStateShape", "State", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Objects {
				if src.State == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Object", "State", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.States {
				for _, target := range src.SubStates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.State", "SubStates", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.States {
				if src.Parent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.State", "Parent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StateMachines {
				if src.InitialState == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StateMachine", "InitialState", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StateMachines {
				for _, target := range src.States {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StateMachine", "States", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StateShapes {
				if src.State == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StateShape", "State", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Transitions {
				if src.Start == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Transition", "Start", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Transitions {
				if src.End == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Transition", "End", refNames, formGroup, probe.formStage)
		}
	case *models.StateMachine:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("InitialState", inst.InitialState, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.State](), probe.formStage)

		{
			// Slice of pointers: States
			div := (&form.FormDiv{Name: "States"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.States {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "States",
				Label: "States",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Diagrams
			div := (&form.FormDiv{Name: "Diagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Diagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Diagrams",
				Label: "Diagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsWithTransitionNameAutonamticalyGenerated", inst.IsWithTransitionNameAutonamticalyGenerated, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootStateMachines {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootStateMachines", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.StateMachinesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "StateMachinesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}
	case *models.StateShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("State", inst.State, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.State](), probe.formStage)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.State_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "State_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Transition:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Start", inst.Start, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.State](), probe.formStage)
		StageSetAssociationFieldToForm("End", inst.End, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.State](), probe.formStage)

		{
			// Slice of pointers: RolesWithPermissions
			div := (&form.FormDiv{Name: "RolesWithPermissions"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RolesWithPermissions {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RolesWithPermissions",
				Label: "RolesWithPermissions",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GeneratedMessages
			div := (&form.FormDiv{Name: "GeneratedMessages"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GeneratedMessages {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GeneratedMessages",
				Label: "GeneratedMessages",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Guard", inst.Guard, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Guard](), probe.formStage)

		{
			// Slice of pointers: Diagrams
			div := (&form.FormDiv{Name: "Diagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Diagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Diagrams",
				Label: "Diagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsRolesNodeExpanded", inst.IsRolesNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsMessagesNodeExpanded", inst.IsMessagesNodeExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Messages {
				if src.OriginTransition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Message", "OriginTransition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Transition_Shapes {
				if src.Transition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Transition_Shape", "Transition", refNames, formGroup, probe.formStage)
		}
	case *models.Transition_Shape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Transition", inst.Transition, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Transition](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Transition_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Transition_Shapes", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
