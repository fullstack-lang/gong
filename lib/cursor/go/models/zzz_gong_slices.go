// generated code - do not edit
package models

import (
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Cursor
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Cursors {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (cursor *Cursor) GongCopy() GongstructIF {
	newInstance := *cursor
	return &newInstance
}

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	// insertion point per named struct
	var cursors_newInstances []*Cursor
	var cursors_deletedInstances []*Cursor

	// parse all staged instances and check if they have a reference
	for cursor := range stage.Cursors {
		if ref, ok := stage.Cursors_reference[cursor]; !ok {
			cursors_newInstances = append(cursors_newInstances, cursor)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Cursor "+cursor.Name,
				)
			}
		} else {
			diffs := cursor.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Cursor "+cursor.Name + " diffs on fields: "+strings.Join(diffs, ", "),
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cursor := range stage.Cursors_reference {
		if _, ok := stage.Cursors[cursor]; !ok {
			cursors_deletedInstances = append(cursors_deletedInstances, cursor)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Cursor "+cursor.Name,
				)
			}
		}
	}

	lenNewInstances += len(cursors_newInstances)
	lenDeletedInstances += len(cursors_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Cursors_reference = make(map[*Cursor]*Cursor)
	for instance := range stage.Cursors {
		stage.Cursors_reference[instance] = instance.GongCopy().(*Cursor)
	}

}
