// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

)

var (
	_ = time.Hour
	_ = slices.Index[[]int, int]
	_ = sort.Slice
	_ = strconv.Itoa
)

// StageSet coordinates multiple stages across packages
type StageSet struct {
	Stage *Stage
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
	return stageSet
}

// GetProbeSplitStageName returns the split stage name for the StageSet probe
func (stageSet *StageSet) GetProbeSplitStageName() string {
	if stageSet.Stage != nil {
		return stageSet.Stage.GetProbeSplitStageName() + "_stageset"
	}
	return "stageset_probe_split"
}

// MarshallFile marshalls all stages into a file
func (stageSet *StageSet) MarshallFile(filename, packageName string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stageSet.Marshall(file, packageName)
}

// Marshall marshalls all stages into an open file
func (stageSet *StageSet) Marshall(file *os.File, packageName string) {
	res, err := stageSet.MarshallToString(packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}
	fmt.Fprintln(file, res)
}

// MarshallToString marshalls all stages into a Go code string
func (stageSet *StageSet) MarshallToString(packageName string) (res string, err error) {
	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stageSet.Stage != nil {
		for _, action := range __gong__sortStageSetInstances(stageSet.Stage.Actions, stageSet.Stage.Action_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			actionIdent := "__models" + action.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Action{Name: %s}).Stage(stageSet.Stage)", actionIdent, __gong__toRawStringLiteral(action.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", actionIdent, __gong__toRawStringLiteral(action.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Criticality = %s", actionIdent, __gong__toRawStringLiteral(string(action.Criticality))))
		}
	}
	if stageSet.Stage != nil {
		for _, activities := range __gong__sortStageSetInstances(stageSet.Stage.Activitiess, stageSet.Stage.Activities_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			activitiesIdent := "__models" + activities.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Activities{Name: %s}).Stage(stageSet.Stage)", activitiesIdent, __gong__toRawStringLiteral(activities.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", activitiesIdent, __gong__toRawStringLiteral(activities.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Criticality = %s", activitiesIdent, __gong__toRawStringLiteral(string(activities.Criticality))))
		}
	}
	if stageSet.Stage != nil {
		for _, diagram := range __gong__sortStageSetInstances(stageSet.Stage.Diagrams, stageSet.Stage.Diagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramIdent := "__models" + diagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Diagram{Name: %s}).Stage(stageSet.Stage)", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramIdent, diagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramIdent, diagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramIdent, diagram.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStatesNodeExpanded = %t", diagramIdent, diagram.IsStatesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowRoles = %t", diagramIdent, diagram.ShowRoles))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowMessages = %t", diagramIdent, diagram.ShowMessages))
			for _, elem := range diagram.State_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.State_Shapes = append(%s.State_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.StatesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StatesWhoseNodeIsExpanded = append(%s.StatesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Transition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Transition_Shapes = append(%s.Transition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Note_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteState_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteState_Shapes = append(%s.NoteState_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, guard := range __gong__sortStageSetInstances(stageSet.Stage.Guards, stageSet.Stage.Guard_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			guardIdent := "__models" + guard.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Guard{Name: %s}).Stage(stageSet.Stage)", guardIdent, __gong__toRawStringLiteral(guard.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", guardIdent, __gong__toRawStringLiteral(guard.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, kill := range __gong__sortStageSetInstances(stageSet.Stage.Kills, stageSet.Stage.Kill_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			killIdent := "__models" + kill.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Kill{Name: %s}).Stage(stageSet.Stage)", killIdent, __gong__toRawStringLiteral(kill.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", killIdent, __gong__toRawStringLiteral(kill.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, library := range __gong__sortStageSetInstances(stageSet.Stage.Librarys, stageSet.Stage.Library_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			libraryIdent := "__models" + library.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Library{Name: %s}).Stage(stageSet.Stage)", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStateMachinesNodeExpanded = %t", libraryIdent, library.IsStateMachinesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubLibrariesNodeExpanded = %t", libraryIdent, library.IsSubLibrariesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpandedTmp = %t", libraryIdent, library.IsExpandedTmp))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRolesNodeExpanded = %t", libraryIdent, library.IsRolesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMessageTypesNodeExpanded = %t", libraryIdent, library.IsMessageTypesNodeExpanded))
			for _, elem := range library.SubLibraries {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootStateMachines {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootStateMachines = append(%s.RootStateMachines, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.StateMachinesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StateMachinesWhoseNodeIsExpanded = append(%s.StateMachinesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibrariesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibrariesWhoseNodeIsExpanded = append(%s.SubLibrariesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.Roles {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Roles = append(%s.Roles, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.MessageTypes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MessageTypes = append(%s.MessageTypes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, message := range __gong__sortStageSetInstances(stageSet.Stage.Messages, stageSet.Stage.Message_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			messageIdent := "__models" + message.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Message{Name: %s}).Stage(stageSet.Stage)", messageIdent, __gong__toRawStringLiteral(message.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", messageIdent, __gong__toRawStringLiteral(message.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", messageIdent, message.IsSelected))
			if message.MessageType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + message.MessageType.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MessageType = %s", messageIdent, targetIdent))
			}
			if message.OriginTransition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + message.OriginTransition.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OriginTransition = %s", messageIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, messagetype := range __gong__sortStageSetInstances(stageSet.Stage.MessageTypes, stageSet.Stage.MessageType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			messagetypeIdent := "__models" + messagetype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MessageType{Name: %s}).Stage(stageSet.Stage)", messagetypeIdent, __gong__toRawStringLiteral(messagetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", messagetypeIdent, __gong__toRawStringLiteral(messagetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", messagetypeIdent, __gong__toRawStringLiteral(messagetype.Description)))
		}
	}
	if stageSet.Stage != nil {
		for _, note := range __gong__sortStageSetInstances(stageSet.Stage.Notes, stageSet.Stage.Note_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteIdent := "__models" + note.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Note{Name: %s}).Stage(stageSet.Stage)", noteIdent, __gong__toRawStringLiteral(note.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteIdent, __gong__toRawStringLiteral(note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", noteIdent, __gong__toRawStringLiteral(note.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteIdent, note.IsExpanded))
			if note.State != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.State.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.State = %s", noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, noteshape := range __gong__sortStageSetInstances(stageSet.Stage.NoteShapes, stageSet.Stage.NoteShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteshapeIdent := "__models" + noteshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteShape{Name: %s}).Stage(stageSet.Stage)", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OverideLayoutDirection = %t", noteshapeIdent, noteshape.OverideLayoutDirection))
			values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = %d", noteshapeIdent, int(noteshape.LayoutDirection)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", noteshapeIdent, noteshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", noteshapeIdent, noteshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", noteshapeIdent, noteshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", noteshapeIdent, noteshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteshapeIdent, noteshape.IsHidden))
			if noteshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, notestateshape := range __gong__sortStageSetInstances(stageSet.Stage.NoteStateShapes, stageSet.Stage.NoteStateShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notestateshapeIdent := "__models" + notestateshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteStateShape{Name: %s}).Stage(stageSet.Stage)", notestateshapeIdent, __gong__toRawStringLiteral(notestateshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notestateshapeIdent, __gong__toRawStringLiteral(notestateshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notestateshapeIdent, notestateshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notestateshapeIdent, notestateshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notestateshapeIdent, __gong__toRawStringLiteral(string(notestateshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notestateshapeIdent, __gong__toRawStringLiteral(string(notestateshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notestateshapeIdent, notestateshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notestateshapeIdent, notestateshape.IsHidden))
			if notestateshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notestateshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notestateshapeIdent, targetIdent))
			}
			if notestateshape.State != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notestateshape.State.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.State = %s", notestateshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, object := range __gong__sortStageSetInstances(stageSet.Stage.Objects, stageSet.Stage.Object_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			objectIdent := "__models" + object.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Object{Name: %s}).Stage(stageSet.Stage)", objectIdent, __gong__toRawStringLiteral(object.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", objectIdent, __gong__toRawStringLiteral(object.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", objectIdent, object.IsSelected))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", objectIdent, object.Rank))
			values.WriteString(fmt.Sprintf("\n\t%s.DOF, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", objectIdent, object.DOF.String()))
			if object.State != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + object.State.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.State = %s", objectIdent, targetIdent))
			}
			for _, elem := range object.Messages {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Messages = append(%s.Messages, %s)", objectIdent, objectIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, role := range __gong__sortStageSetInstances(stageSet.Stage.Roles, stageSet.Stage.Role_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			roleIdent := "__models" + role.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Role{Name: %s}).Stage(stageSet.Stage)", roleIdent, __gong__toRawStringLiteral(role.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", roleIdent, __gong__toRawStringLiteral(role.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Acronym = %s", roleIdent, __gong__toRawStringLiteral(role.Acronym)))
			for _, elem := range role.RolesWithSamePermissions {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RolesWithSamePermissions = append(%s.RolesWithSamePermissions, %s)", roleIdent, roleIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, state := range __gong__sortStageSetInstances(stageSet.Stage.States, stageSet.Stage.State_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stateIdent := "__models" + state.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.State{Name: %s}).Stage(stageSet.Stage)", stateIdent, __gong__toRawStringLiteral(state.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stateIdent, __gong__toRawStringLiteral(state.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEndState = %t", stateIdent, state.IsEndState))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDecisionNode = %t", stateIdent, state.IsDecisionNode))
			values.WriteString(fmt.Sprintf("\n\t%s.IsFictious = %t", stateIdent, state.IsFictious))
			for _, elem := range state.SubStates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubStates = append(%s.SubStates, %s)", stateIdent, stateIdent, targetIdent))
			}
			if state.Entry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + state.Entry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Entry = %s", stateIdent, targetIdent))
			}
			for _, elem := range state.Activities {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Activities = append(%s.Activities, %s)", stateIdent, stateIdent, targetIdent))
			}
			if state.Exit != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + state.Exit.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Exit = %s", stateIdent, targetIdent))
			}
			if state.Parent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + state.Parent.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parent = %s", stateIdent, targetIdent))
			}
			for _, elem := range state.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", stateIdent, stateIdent, targetIdent))
			}
			for _, elem := range state.Notes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notes = append(%s.Notes, %s)", stateIdent, stateIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, statemachine := range __gong__sortStageSetInstances(stageSet.Stage.StateMachines, stageSet.Stage.StateMachine_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			statemachineIdent := "__models" + statemachine.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StateMachine{Name: %s}).Stage(stageSet.Stage)", statemachineIdent, __gong__toRawStringLiteral(statemachine.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", statemachineIdent, __gong__toRawStringLiteral(statemachine.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithTransitionNameAutonamticalyGenerated = %t", statemachineIdent, statemachine.IsWithTransitionNameAutonamticalyGenerated))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", statemachineIdent, __gong__toRawStringLiteral(statemachine.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", statemachineIdent, statemachine.IsExpanded))
			if statemachine.InitialState != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + statemachine.InitialState.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InitialState = %s", statemachineIdent, targetIdent))
			}
			for _, elem := range statemachine.States {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.States = append(%s.States, %s)", statemachineIdent, statemachineIdent, targetIdent))
			}
			for _, elem := range statemachine.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", statemachineIdent, statemachineIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, stateshape := range __gong__sortStageSetInstances(stageSet.Stage.StateShapes, stageSet.Stage.StateShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stateshapeIdent := "__models" + stateshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StateShape{Name: %s}).Stage(stageSet.Stage)", stateshapeIdent, __gong__toRawStringLiteral(stateshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stateshapeIdent, __gong__toRawStringLiteral(stateshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", stateshapeIdent, stateshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", stateshapeIdent, stateshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", stateshapeIdent, stateshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", stateshapeIdent, stateshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stateshapeIdent, stateshape.IsHidden))
			if stateshape.State != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stateshape.State.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.State = %s", stateshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, transition := range __gong__sortStageSetInstances(stageSet.Stage.Transitions, stageSet.Stage.Transition_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			transitionIdent := "__models" + transition.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Transition{Name: %s}).Stage(stageSet.Stage)", transitionIdent, __gong__toRawStringLiteral(transition.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", transitionIdent, __gong__toRawStringLiteral(transition.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", transitionIdent, transition.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRolesNodeExpanded = %t", transitionIdent, transition.IsRolesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMessagesNodeExpanded = %t", transitionIdent, transition.IsMessagesNodeExpanded))
			if transition.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + transition.Start.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", transitionIdent, targetIdent))
			}
			if transition.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + transition.End.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", transitionIdent, targetIdent))
			}
			for _, elem := range transition.RolesWithPermissions {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RolesWithPermissions = append(%s.RolesWithPermissions, %s)", transitionIdent, transitionIdent, targetIdent))
			}
			for _, elem := range transition.GeneratedMessages {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeneratedMessages = append(%s.GeneratedMessages, %s)", transitionIdent, transitionIdent, targetIdent))
			}
			if transition.Guard != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + transition.Guard.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Guard = %s", transitionIdent, targetIdent))
			}
			for _, elem := range transition.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", transitionIdent, transitionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, transition_shape := range __gong__sortStageSetInstances(stageSet.Stage.Transition_Shapes, stageSet.Stage.Transition_Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			transition_shapeIdent := "__models" + transition_shape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Transition_Shape{Name: %s}).Stage(stageSet.Stage)", transition_shapeIdent, __gong__toRawStringLiteral(transition_shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", transition_shapeIdent, __gong__toRawStringLiteral(transition_shape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", transition_shapeIdent, transition_shape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", transition_shapeIdent, transition_shape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", transition_shapeIdent, __gong__toRawStringLiteral(string(transition_shape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", transition_shapeIdent, __gong__toRawStringLiteral(string(transition_shape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", transition_shapeIdent, transition_shape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", transition_shapeIdent, transition_shape.IsHidden))
			if transition_shape.Transition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + transition_shape.Transition.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Transition = %s", transition_shapeIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/statemachines/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *models.StageSet) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, packageName, declarations.String(), values.String(), pointers.String())

	return res, nil
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file into stageSet
func (stageSet *StageSet) ParseAstFile(pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist " + pathToFile + " ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file into stageSet
func (stageSet *StageSet) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New("Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string into stageSet
func (stageSet *StageSet) ParseAstString(blob string, preserveOrder bool) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", blob, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances into stageSet
func (stageSet *StageSet) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	identifierMap := make(map[string]any)

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/dsm/statemachines/go/models":
			aliasToCanonical[alias] = "models"
		}
	}

	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var pkgAlias string
					var typeName string
					var instanceName string

					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								if pkgId, ok := selExpr.X.(*ast.Ident); ok {
									pkgAlias = pkgId.Name
								}
								typeName = selExpr.Sel.Name
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					if canonical, ok := aliasToCanonical[pkgAlias]; ok {
						pkgAlias = canonical
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "Action":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Action), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Activities":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Activities), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Diagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Diagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Guard":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Guard), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Kill":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Kill), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Message":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Message), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MessageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MessageType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteStateShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteStateShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Object":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Object), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Role":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Role), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "State":
					identifierMap[ident.Name] = __gong__stageSetInit(new(State), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StateMachine":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StateMachine), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StateShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StateShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Transition":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Transition), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Transition_Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Transition_Shape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *Action:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Criticality":
						inst.Criticality = Criticality(GongExtractString(rhs))
					}
				case *Activities:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Criticality":
						inst.Criticality = Criticality(GongExtractString(rhs))
					}
				case *Diagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsEditable_":
						inst.IsEditable_ = GongExtractBool(rhs)
					case "IsStatesNodeExpanded":
						inst.IsStatesNodeExpanded = GongExtractBool(rhs)
					case "State_Shapes":
						__gong__assignSliceOfPointers(&inst.State_Shapes, rhs, identifierMap)
					case "StatesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.StatesWhoseNodeIsExpanded, rhs, identifierMap)
					case "Transition_Shapes":
						__gong__assignSliceOfPointers(&inst.Transition_Shapes, rhs, identifierMap)
					case "Note_Shapes":
						__gong__assignSliceOfPointers(&inst.Note_Shapes, rhs, identifierMap)
					case "NoteState_Shapes":
						__gong__assignSliceOfPointers(&inst.NoteState_Shapes, rhs, identifierMap)
					case "ShowRoles":
						inst.ShowRoles = GongExtractBool(rhs)
					case "ShowMessages":
						inst.ShowMessages = GongExtractBool(rhs)
					}
				case *Guard:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Kill:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Library:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SubLibraries":
						__gong__assignSliceOfPointers(&inst.SubLibraries, rhs, identifierMap)
					case "NbPixPerCharacter":
						inst.NbPixPerCharacter = GongExtractFloat(rhs)
					case "LogoSVGFile":
						inst.LogoSVGFile = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsRootLibrary":
						inst.IsRootLibrary = GongExtractBool(rhs)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					case "RootStateMachines":
						__gong__assignSliceOfPointers(&inst.RootStateMachines, rhs, identifierMap)
					case "IsStateMachinesNodeExpanded":
						inst.IsStateMachinesNodeExpanded = GongExtractBool(rhs)
					case "StateMachinesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.StateMachinesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsSubLibrariesNodeExpanded":
						inst.IsSubLibrariesNodeExpanded = GongExtractBool(rhs)
					case "SubLibrariesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.SubLibrariesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsExpandedTmp":
						inst.IsExpandedTmp = GongExtractBool(rhs)
					case "Roles":
						__gong__assignSliceOfPointers(&inst.Roles, rhs, identifierMap)
					case "IsRolesNodeExpanded":
						inst.IsRolesNodeExpanded = GongExtractBool(rhs)
					case "MessageTypes":
						__gong__assignSliceOfPointers(&inst.MessageTypes, rhs, identifierMap)
					case "IsMessageTypesNodeExpanded":
						inst.IsMessageTypesNodeExpanded = GongExtractBool(rhs)
					}
				case *Message:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					case "MessageType":
						__gong__assignPointer(&inst.MessageType, rhs, identifierMap)
					case "OriginTransition":
						__gong__assignPointer(&inst.OriginTransition, rhs, identifierMap)
					}
				case *MessageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *Note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "State":
						__gong__assignPointer(&inst.State, rhs, identifierMap)
					}
				case *NoteShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "OverideLayoutDirection":
						inst.OverideLayoutDirection = GongExtractBool(rhs)
					case "LayoutDirection":
						inst.LayoutDirection = LayoutDirection(GongExtractInt(rhs))
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					}
				case *NoteStateShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "State":
						__gong__assignPointer(&inst.State, rhs, identifierMap)
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					}
				case *Object:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "State":
						__gong__assignPointer(&inst.State, rhs, identifierMap)
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					case "DOF":
						inst.DOF = GongExtractDate(rhs)
					case "Messages":
						__gong__assignSliceOfPointers(&inst.Messages, rhs, identifierMap)
					}
				case *Role:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Acronym":
						inst.Acronym = GongExtractString(rhs)
					case "RolesWithSamePermissions":
						__gong__assignSliceOfPointers(&inst.RolesWithSamePermissions, rhs, identifierMap)
					}
				case *State:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsEndState":
						inst.IsEndState = GongExtractBool(rhs)
					case "IsDecisionNode":
						inst.IsDecisionNode = GongExtractBool(rhs)
					case "SubStates":
						__gong__assignSliceOfPointers(&inst.SubStates, rhs, identifierMap)
					case "Entry":
						__gong__assignPointer(&inst.Entry, rhs, identifierMap)
					case "Activities":
						__gong__assignSliceOfPointers(&inst.Activities, rhs, identifierMap)
					case "Exit":
						__gong__assignPointer(&inst.Exit, rhs, identifierMap)
					case "Parent":
						__gong__assignPointer(&inst.Parent, rhs, identifierMap)
					case "IsFictious":
						inst.IsFictious = GongExtractBool(rhs)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					case "Notes":
						__gong__assignSliceOfPointers(&inst.Notes, rhs, identifierMap)
					}
				case *StateMachine:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "InitialState":
						__gong__assignPointer(&inst.InitialState, rhs, identifierMap)
					case "States":
						__gong__assignSliceOfPointers(&inst.States, rhs, identifierMap)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					case "IsWithTransitionNameAutonamticalyGenerated":
						inst.IsWithTransitionNameAutonamticalyGenerated = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *StateShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "State":
						__gong__assignPointer(&inst.State, rhs, identifierMap)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					}
				case *Transition:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Start":
						__gong__assignPointer(&inst.Start, rhs, identifierMap)
					case "End":
						__gong__assignPointer(&inst.End, rhs, identifierMap)
					case "RolesWithPermissions":
						__gong__assignSliceOfPointers(&inst.RolesWithPermissions, rhs, identifierMap)
					case "GeneratedMessages":
						__gong__assignSliceOfPointers(&inst.GeneratedMessages, rhs, identifierMap)
					case "Guard":
						__gong__assignPointer(&inst.Guard, rhs, identifierMap)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsRolesNodeExpanded":
						inst.IsRolesNodeExpanded = GongExtractBool(rhs)
					case "IsMessagesNodeExpanded":
						inst.IsMessagesNodeExpanded = GongExtractBool(rhs)
					}
				case *Transition_Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Transition":
						__gong__assignPointer(&inst.Transition, rhs, identifierMap)
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					}
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
}

func __gong__stageSetInit[P interface {
	SetName(string)
	StageVoid(S)
	StagePreserveOrder(S, uint)
}, S any](instance P, stage S, identifier string, instanceName string, preserveOrder bool) any {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if order, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifier", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, order)
		}
	}
	return instance
}

func __gong__assignPointer[T any](targetPtr **T, rhs ast.Expr, identifierMap map[string]any) {
	if rIdent, ok := rhs.(*ast.Ident); ok {
		if rIdent.Name == "nil" {
			*targetPtr = nil
			return
		}
		if target, ok := identifierMap[rIdent.Name]; ok {
			if typedTarget, ok := target.(*T); ok {
				*targetPtr = typedTarget
			}
		}
	}
}

func __gong__assignSliceOfPointers[T any](slice *[]*T, rhs ast.Expr, identifierMap map[string]any) {
	if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
		if rIdent, ok := call.Args[1].(*ast.Ident); ok {
			if target, ok := identifierMap[rIdent.Name]; ok {
				if typedTarget, ok := target.(*T); ok {
					*slice = append(*slice, typedTarget)
				}
			}
		}
	}
}
