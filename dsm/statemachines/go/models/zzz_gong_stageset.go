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
		actionOrdered := []*Action{}
		for action := range stageSet.Stage.Actions {
			actionOrdered = append(actionOrdered, action)
		}
		sort.Slice(actionOrdered, func(i, j int) bool {
			return stageSet.Stage.Action_stagedOrder[actionOrdered[i]] < stageSet.Stage.Action_stagedOrder[actionOrdered[j]]
		})
		for _, action := range actionOrdered {
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
		activitiesOrdered := []*Activities{}
		for activities := range stageSet.Stage.Activitiess {
			activitiesOrdered = append(activitiesOrdered, activities)
		}
		sort.Slice(activitiesOrdered, func(i, j int) bool {
			return stageSet.Stage.Activities_stagedOrder[activitiesOrdered[i]] < stageSet.Stage.Activities_stagedOrder[activitiesOrdered[j]]
		})
		for _, activities := range activitiesOrdered {
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
		diagramOrdered := []*Diagram{}
		for diagram := range stageSet.Stage.Diagrams {
			diagramOrdered = append(diagramOrdered, diagram)
		}
		sort.Slice(diagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Diagram_stagedOrder[diagramOrdered[i]] < stageSet.Stage.Diagram_stagedOrder[diagramOrdered[j]]
		})
		for _, diagram := range diagramOrdered {
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
		guardOrdered := []*Guard{}
		for guard := range stageSet.Stage.Guards {
			guardOrdered = append(guardOrdered, guard)
		}
		sort.Slice(guardOrdered, func(i, j int) bool {
			return stageSet.Stage.Guard_stagedOrder[guardOrdered[i]] < stageSet.Stage.Guard_stagedOrder[guardOrdered[j]]
		})
		for _, guard := range guardOrdered {
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
		killOrdered := []*Kill{}
		for kill := range stageSet.Stage.Kills {
			killOrdered = append(killOrdered, kill)
		}
		sort.Slice(killOrdered, func(i, j int) bool {
			return stageSet.Stage.Kill_stagedOrder[killOrdered[i]] < stageSet.Stage.Kill_stagedOrder[killOrdered[j]]
		})
		for _, kill := range killOrdered {
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
		libraryOrdered := []*Library{}
		for library := range stageSet.Stage.Librarys {
			libraryOrdered = append(libraryOrdered, library)
		}
		sort.Slice(libraryOrdered, func(i, j int) bool {
			return stageSet.Stage.Library_stagedOrder[libraryOrdered[i]] < stageSet.Stage.Library_stagedOrder[libraryOrdered[j]]
		})
		for _, library := range libraryOrdered {
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
		}
	}
	if stageSet.Stage != nil {
		messageOrdered := []*Message{}
		for message := range stageSet.Stage.Messages {
			messageOrdered = append(messageOrdered, message)
		}
		sort.Slice(messageOrdered, func(i, j int) bool {
			return stageSet.Stage.Message_stagedOrder[messageOrdered[i]] < stageSet.Stage.Message_stagedOrder[messageOrdered[j]]
		})
		for _, message := range messageOrdered {
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
		messagetypeOrdered := []*MessageType{}
		for messagetype := range stageSet.Stage.MessageTypes {
			messagetypeOrdered = append(messagetypeOrdered, messagetype)
		}
		sort.Slice(messagetypeOrdered, func(i, j int) bool {
			return stageSet.Stage.MessageType_stagedOrder[messagetypeOrdered[i]] < stageSet.Stage.MessageType_stagedOrder[messagetypeOrdered[j]]
		})
		for _, messagetype := range messagetypeOrdered {
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
		noteOrdered := []*Note{}
		for note := range stageSet.Stage.Notes {
			noteOrdered = append(noteOrdered, note)
		}
		sort.Slice(noteOrdered, func(i, j int) bool {
			return stageSet.Stage.Note_stagedOrder[noteOrdered[i]] < stageSet.Stage.Note_stagedOrder[noteOrdered[j]]
		})
		for _, note := range noteOrdered {
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
		noteshapeOrdered := []*NoteShape{}
		for noteshape := range stageSet.Stage.NoteShapes {
			noteshapeOrdered = append(noteshapeOrdered, noteshape)
		}
		sort.Slice(noteshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteShape_stagedOrder[noteshapeOrdered[i]] < stageSet.Stage.NoteShape_stagedOrder[noteshapeOrdered[j]]
		})
		for _, noteshape := range noteshapeOrdered {
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
		notestateshapeOrdered := []*NoteStateShape{}
		for notestateshape := range stageSet.Stage.NoteStateShapes {
			notestateshapeOrdered = append(notestateshapeOrdered, notestateshape)
		}
		sort.Slice(notestateshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteStateShape_stagedOrder[notestateshapeOrdered[i]] < stageSet.Stage.NoteStateShape_stagedOrder[notestateshapeOrdered[j]]
		})
		for _, notestateshape := range notestateshapeOrdered {
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
		objectOrdered := []*Object{}
		for object := range stageSet.Stage.Objects {
			objectOrdered = append(objectOrdered, object)
		}
		sort.Slice(objectOrdered, func(i, j int) bool {
			return stageSet.Stage.Object_stagedOrder[objectOrdered[i]] < stageSet.Stage.Object_stagedOrder[objectOrdered[j]]
		})
		for _, object := range objectOrdered {
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
		roleOrdered := []*Role{}
		for role := range stageSet.Stage.Roles {
			roleOrdered = append(roleOrdered, role)
		}
		sort.Slice(roleOrdered, func(i, j int) bool {
			return stageSet.Stage.Role_stagedOrder[roleOrdered[i]] < stageSet.Stage.Role_stagedOrder[roleOrdered[j]]
		})
		for _, role := range roleOrdered {
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
		stateOrdered := []*State{}
		for state := range stageSet.Stage.States {
			stateOrdered = append(stateOrdered, state)
		}
		sort.Slice(stateOrdered, func(i, j int) bool {
			return stageSet.Stage.State_stagedOrder[stateOrdered[i]] < stageSet.Stage.State_stagedOrder[stateOrdered[j]]
		})
		for _, state := range stateOrdered {
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
		statemachineOrdered := []*StateMachine{}
		for statemachine := range stageSet.Stage.StateMachines {
			statemachineOrdered = append(statemachineOrdered, statemachine)
		}
		sort.Slice(statemachineOrdered, func(i, j int) bool {
			return stageSet.Stage.StateMachine_stagedOrder[statemachineOrdered[i]] < stageSet.Stage.StateMachine_stagedOrder[statemachineOrdered[j]]
		})
		for _, statemachine := range statemachineOrdered {
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
		stateshapeOrdered := []*StateShape{}
		for stateshape := range stageSet.Stage.StateShapes {
			stateshapeOrdered = append(stateshapeOrdered, stateshape)
		}
		sort.Slice(stateshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.StateShape_stagedOrder[stateshapeOrdered[i]] < stageSet.Stage.StateShape_stagedOrder[stateshapeOrdered[j]]
		})
		for _, stateshape := range stateshapeOrdered {
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
		transitionOrdered := []*Transition{}
		for transition := range stageSet.Stage.Transitions {
			transitionOrdered = append(transitionOrdered, transition)
		}
		sort.Slice(transitionOrdered, func(i, j int) bool {
			return stageSet.Stage.Transition_stagedOrder[transitionOrdered[i]] < stageSet.Stage.Transition_stagedOrder[transitionOrdered[j]]
		})
		for _, transition := range transitionOrdered {
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
		transition_shapeOrdered := []*Transition_Shape{}
		for transition_shape := range stageSet.Stage.Transition_Shapes {
			transition_shapeOrdered = append(transition_shapeOrdered, transition_shape)
		}
		sort.Slice(transition_shapeOrdered, func(i, j int) bool {
			return stageSet.Stage.Transition_Shape_stagedOrder[transition_shapeOrdered[i]] < stageSet.Stage.Transition_Shape_stagedOrder[transition_shapeOrdered[j]]
		})
		for _, transition_shape := range transition_shapeOrdered {
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
					if !preserveOrder {
						inst := (&Action{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Action)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Activities":
					if !preserveOrder {
						inst := (&Activities{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Activities)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Diagram":
					if !preserveOrder {
						inst := (&Diagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Diagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Guard":
					if !preserveOrder {
						inst := (&Guard{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Guard)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Kill":
					if !preserveOrder {
						inst := (&Kill{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Kill)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Library":
					if !preserveOrder {
						inst := (&Library{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Library)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Message":
					if !preserveOrder {
						inst := (&Message{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Message)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MessageType":
					if !preserveOrder {
						inst := (&MessageType{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MessageType)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Note":
					if !preserveOrder {
						inst := (&Note{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Note)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "NoteShape":
					if !preserveOrder {
						inst := (&NoteShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NoteShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "NoteStateShape":
					if !preserveOrder {
						inst := (&NoteStateShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NoteStateShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Object":
					if !preserveOrder {
						inst := (&Object{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Object)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Role":
					if !preserveOrder {
						inst := (&Role{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Role)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "State":
					if !preserveOrder {
						inst := (&State{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(State)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StateMachine":
					if !preserveOrder {
						inst := (&StateMachine{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StateMachine)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StateShape":
					if !preserveOrder {
						inst := (&StateShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StateShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Transition":
					if !preserveOrder {
						inst := (&Transition{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Transition)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Transition_Shape":
					if !preserveOrder {
						inst := (&Transition_Shape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Transition_Shape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*StateShape); ok {
										inst.State_Shapes = append(inst.State_Shapes, typedTarget)
									}
								}
							}
						}
					case "StatesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*State); ok {
										inst.StatesWhoseNodeIsExpanded = append(inst.StatesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "Transition_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Transition_Shape); ok {
										inst.Transition_Shapes = append(inst.Transition_Shapes, typedTarget)
									}
								}
							}
						}
					case "Note_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*NoteShape); ok {
										inst.Note_Shapes = append(inst.Note_Shapes, typedTarget)
									}
								}
							}
						}
					case "NoteState_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*NoteStateShape); ok {
										inst.NoteState_Shapes = append(inst.NoteState_Shapes, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Library); ok {
										inst.SubLibraries = append(inst.SubLibraries, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Diagram); ok {
										inst.Diagrams = append(inst.Diagrams, typedTarget)
									}
								}
							}
						}
					case "RootStateMachines":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*StateMachine); ok {
										inst.RootStateMachines = append(inst.RootStateMachines, typedTarget)
									}
								}
							}
						}
					case "IsStateMachinesNodeExpanded":
						inst.IsStateMachinesNodeExpanded = GongExtractBool(rhs)
					case "StateMachinesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*StateMachine); ok {
										inst.StateMachinesWhoseNodeIsExpanded = append(inst.StateMachinesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsSubLibrariesNodeExpanded":
						inst.IsSubLibrariesNodeExpanded = GongExtractBool(rhs)
					case "SubLibrariesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Library); ok {
										inst.SubLibrariesWhoseNodeIsExpanded = append(inst.SubLibrariesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsExpandedTmp":
						inst.IsExpandedTmp = GongExtractBool(rhs)
					case "Roles":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Role); ok {
										inst.Roles = append(inst.Roles, typedTarget)
									}
								}
							}
						}
					}
				case *Message:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					case "MessageType":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MessageType); ok {
									inst.MessageType = typedTarget
								}
							}
						}
					case "OriginTransition":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Transition); ok {
									inst.OriginTransition = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*State); ok {
									inst.State = typedTarget
								}
							}
						}
					}
				case *NoteShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
									inst.Note = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
									inst.Note = typedTarget
								}
							}
						}
					case "State":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*State); ok {
									inst.State = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*State); ok {
									inst.State = typedTarget
								}
							}
						}
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					case "DOF":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.DOF, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "Messages":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Message); ok {
										inst.Messages = append(inst.Messages, typedTarget)
									}
								}
							}
						}
					}
				case *Role:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Acronym":
						inst.Acronym = GongExtractString(rhs)
					case "RolesWithSamePermissions":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Role); ok {
										inst.RolesWithSamePermissions = append(inst.RolesWithSamePermissions, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*State); ok {
										inst.SubStates = append(inst.SubStates, typedTarget)
									}
								}
							}
						}
					case "Entry":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Action); ok {
									inst.Entry = typedTarget
								}
							}
						}
					case "Activities":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Activities); ok {
										inst.Activities = append(inst.Activities, typedTarget)
									}
								}
							}
						}
					case "Exit":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Action); ok {
									inst.Exit = typedTarget
								}
							}
						}
					case "Parent":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*State); ok {
									inst.Parent = typedTarget
								}
							}
						}
					case "IsFictious":
						inst.IsFictious = GongExtractBool(rhs)
					case "Diagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Diagram); ok {
										inst.Diagrams = append(inst.Diagrams, typedTarget)
									}
								}
							}
						}
					case "Notes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Note); ok {
										inst.Notes = append(inst.Notes, typedTarget)
									}
								}
							}
						}
					}
				case *StateMachine:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "InitialState":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*State); ok {
									inst.InitialState = typedTarget
								}
							}
						}
					case "States":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*State); ok {
										inst.States = append(inst.States, typedTarget)
									}
								}
							}
						}
					case "Diagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Diagram); ok {
										inst.Diagrams = append(inst.Diagrams, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*State); ok {
									inst.State = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*State); ok {
									inst.Start = typedTarget
								}
							}
						}
					case "End":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*State); ok {
									inst.End = typedTarget
								}
							}
						}
					case "RolesWithPermissions":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Role); ok {
										inst.RolesWithPermissions = append(inst.RolesWithPermissions, typedTarget)
									}
								}
							}
						}
					case "GeneratedMessages":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*MessageType); ok {
										inst.GeneratedMessages = append(inst.GeneratedMessages, typedTarget)
									}
								}
							}
						}
					case "Guard":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Guard); ok {
									inst.Guard = typedTarget
								}
							}
						}
					case "Diagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Diagram); ok {
										inst.Diagrams = append(inst.Diagrams, typedTarget)
									}
								}
							}
						}
					}
				case *Transition_Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Transition":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Transition); ok {
									inst.Transition = typedTarget
								}
							}
						}
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
