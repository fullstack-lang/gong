// generated code - do not edit
package models

import (
	"fmt"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Freqency
	// insertion point per field

	// Compute reverse map for named struct Note
	// insertion point per field
	stage.Note_Frequencies_reverseMap = make(map[*Freqency]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _freqency := range note.Frequencies {
			stage.Note_Frequencies_reverseMap[_freqency] = note
		}
	}

	// Compute reverse map for named struct Player
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Freqencys {
		res = append(res, instance)
	}

	for instance := range stage.Notes {
		res = append(res, instance)
	}

	for instance := range stage.Players {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (freqency *Freqency) GongCopy() GongstructIF {
	newInstance := *freqency
	return &newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := *note
	return &newInstance
}

func (player *Player) GongCopy() GongstructIF {
	newInstance := *player
	return &newInstance
}

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int
	
	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

	// insertion point per named struct
	var freqencys_newInstances []*Freqency
	var freqencys_deletedInstances []*Freqency

	// parse all staged instances and check if they have a reference
	for freqency := range stage.Freqencys {
		if ref, ok := stage.Freqencys_reference[freqency]; !ok {
			freqencys_newInstances = append(freqencys_newInstances, freqency)
			newInstancesStmt += freqency.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := freqency.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := freqency.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\t// modifications for instance %s \n", freqency.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for freqency := range stage.Freqencys_reference {
		if _, ok := stage.Freqencys[freqency]; !ok {
			freqencys_deletedInstances = append(freqencys_deletedInstances, freqency)
			deletedInstancesStmt += freqency.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(freqencys_newInstances)
	lenDeletedInstances += len(freqencys_deletedInstances)
	var notes_newInstances []*Note
	var notes_deletedInstances []*Note

	// parse all staged instances and check if they have a reference
	for note := range stage.Notes {
		if ref, ok := stage.Notes_reference[note]; !ok {
			notes_newInstances = append(notes_newInstances, note)
			newInstancesStmt += note.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := note.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := note.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\t// modifications for instance %s \n", note.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for note := range stage.Notes_reference {
		if _, ok := stage.Notes[note]; !ok {
			notes_deletedInstances = append(notes_deletedInstances, note)
			deletedInstancesStmt += note.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(notes_newInstances)
	lenDeletedInstances += len(notes_deletedInstances)
	var players_newInstances []*Player
	var players_deletedInstances []*Player

	// parse all staged instances and check if they have a reference
	for player := range stage.Players {
		if ref, ok := stage.Players_reference[player]; !ok {
			players_newInstances = append(players_newInstances, player)
			newInstancesStmt += player.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := player.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := player.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\t// modifications for instance %s \n", player.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for player := range stage.Players_reference {
		if _, ok := stage.Players[player]; !ok {
			players_deletedInstances = append(players_deletedInstances, player)
			deletedInstancesStmt += player.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(players_newInstances)
	lenDeletedInstances += len(players_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				newInstancesStmt+fieldsEditStmt+deletedInstancesStmt,
			)
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Freqencys_reference = make(map[*Freqency]*Freqency)
	for instance := range stage.Freqencys {
		stage.Freqencys_reference[instance] = instance.GongCopy().(*Freqency)
	}

	stage.Notes_reference = make(map[*Note]*Note)
	for instance := range stage.Notes {
		stage.Notes_reference[instance] = instance.GongCopy().(*Note)
	}

	stage.Players_reference = make(map[*Player]*Player)
	for instance := range stage.Players {
		stage.Players_reference[instance] = instance.GongCopy().(*Player)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (freqency *Freqency) GongGetOrder(stage *Stage) uint {
	return stage.FreqencyMap_Staged_Order[freqency]
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	return stage.NoteMap_Staged_Order[note]
}

func (player *Player) GongGetOrder(stage *Stage) uint {
	return stage.PlayerMap_Staged_Order[player]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (freqency *Freqency) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", freqency.GongGetGongstructName(), freqency.GongGetOrder(stage))
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

func (player *Player) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", player.GongGetGongstructName(), player.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (freqency *Freqency) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", freqency.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Freqency")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", freqency.Name)
	return
}
func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", note.Name)
	return
}
func (player *Player) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", player.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Player")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", player.Name)
	return
}

// insertion point for unstaging
func (freqency *Freqency) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", freqency.GongGetIdentifier(stage))
	return
}
func (note *Note) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	return
}
func (player *Player) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", player.GongGetIdentifier(stage))
	return
}
