// generated code - do not edit
package models

import (
	"fmt"
	"sort"
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

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var freqencys_newInstances []*Freqency
	var freqencys_deletedInstances []*Freqency

	// parse all staged instances and check if they have a reference
	for freqency := range stage.Freqencys {
		if ref, ok := stage.Freqencys_reference[freqency]; !ok {
			freqencys_newInstances = append(freqencys_newInstances, freqency)
			newInstancesSlice = append(newInstancesSlice, freqency.GongMarshallIdentifier(stage))
			if stage.Freqencys_referenceOrder == nil {
				stage.Freqencys_referenceOrder = make(map[*Freqency]uint)
			}
			stage.Freqencys_referenceOrder[freqency] = stage.FreqencyMap_Staged_Order[freqency]
			newInstancesReverseSlice = append(newInstancesReverseSlice, freqency.GongMarshallUnstaging(stage))
			delete(stage.Freqencys_referenceOrder, freqency)
			fieldInitializers, pointersInitializations := freqency.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FreqencyMap_Staged_Order[ref] = stage.FreqencyMap_Staged_Order[freqency]
			diffs := freqency.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, freqency)
			delete(stage.FreqencyMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", freqency.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Freqencys_reference {
		if _, ok := stage.Freqencys[ref]; !ok {
			freqencys_deletedInstances = append(freqencys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, note.GongMarshallIdentifier(stage))
			if stage.Notes_referenceOrder == nil {
				stage.Notes_referenceOrder = make(map[*Note]uint)
			}
			stage.Notes_referenceOrder[note] = stage.NoteMap_Staged_Order[note]
			newInstancesReverseSlice = append(newInstancesReverseSlice, note.GongMarshallUnstaging(stage))
			delete(stage.Notes_referenceOrder, note)
			fieldInitializers, pointersInitializations := note.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteMap_Staged_Order[ref] = stage.NoteMap_Staged_Order[note]
			diffs := note.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, note)
			delete(stage.NoteMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", note.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Notes_reference {
		if _, ok := stage.Notes[ref]; !ok {
			notes_deletedInstances = append(notes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, player.GongMarshallIdentifier(stage))
			if stage.Players_referenceOrder == nil {
				stage.Players_referenceOrder = make(map[*Player]uint)
			}
			stage.Players_referenceOrder[player] = stage.PlayerMap_Staged_Order[player]
			newInstancesReverseSlice = append(newInstancesReverseSlice, player.GongMarshallUnstaging(stage))
			delete(stage.Players_referenceOrder, player)
			fieldInitializers, pointersInitializations := player.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PlayerMap_Staged_Order[ref] = stage.PlayerMap_Staged_Order[player]
			diffs := player.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, player)
			delete(stage.PlayerMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", player.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Players_reference {
		if _, ok := stage.Players[ref]; !ok {
			players_deletedInstances = append(players_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(players_newInstances)
	lenDeletedInstances += len(players_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {

	// insertion point per named struct
	stage.Freqencys_reference = make(map[*Freqency]*Freqency)
	stage.Freqencys_referenceOrder = make(map[*Freqency]uint) // diff Unstage needs the reference order
	for instance := range stage.Freqencys {
		stage.Freqencys_reference[instance] = instance.GongCopy().(*Freqency)
		stage.Freqencys_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint) // diff Unstage needs the reference order
	for instance := range stage.Notes {
		stage.Notes_reference[instance] = instance.GongCopy().(*Note)
		stage.Notes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Players_reference = make(map[*Player]*Player)
	stage.Players_referenceOrder = make(map[*Player]uint) // diff Unstage needs the reference order
	for instance := range stage.Players {
		stage.Players_reference[instance] = instance.GongCopy().(*Player)
		stage.Players_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (freqency *Freqency) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FreqencyMap_Staged_Order[freqency]; ok {
		return order
	}
	return stage.Freqencys_referenceOrder[freqency]
}

func (freqency *Freqency) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Freqencys_referenceOrder[freqency]
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteMap_Staged_Order[note]; ok {
		return order
	}
	return stage.Notes_referenceOrder[note]
}

func (note *Note) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Notes_referenceOrder[note]
}

func (player *Player) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PlayerMap_Staged_Order[player]; ok {
		return order
	}
	return stage.Players_referenceOrder[player]
}

func (player *Player) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Players_referenceOrder[player]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (freqency *Freqency) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", freqency.GongGetGongstructName(), freqency.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (freqency *Freqency) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", freqency.GongGetGongstructName(), freqency.GongGetReferenceOrder(stage))
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetReferenceOrder(stage))
}

func (player *Player) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", player.GongGetGongstructName(), player.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (player *Player) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", player.GongGetGongstructName(), player.GongGetReferenceOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{Identifier}}", freqency.GongGetReferenceIdentifier(stage))
	return
}
func (note *Note) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetReferenceIdentifier(stage))
	return
}
func (player *Player) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", player.GongGetReferenceIdentifier(stage))
	return
}
