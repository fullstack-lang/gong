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

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

	var newInstancesReverseStmt string
	_ = newInstancesReverseStmt
	var fieldsEditReverseStmt string
	_ = fieldsEditReverseStmt
	var deletedInstancesReverseStmt string
	_ = deletedInstancesReverseStmt

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var cursors_newInstances []*Cursor
	var cursors_deletedInstances []*Cursor

	// parse all staged instances and check if they have a reference
	for cursor := range stage.Cursors {
		if ref, ok := stage.Cursors_reference[cursor]; !ok {
			cursors_newInstances = append(cursors_newInstances, cursor)
			newInstancesStmt += cursor.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += cursor.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := cursor.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cursor.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cursor)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", cursor.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Cursors_reference {
		if _, ok := stage.Cursors[ref]; !ok {
			cursors_deletedInstances = append(cursors_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
		}
	}

	lenNewInstances += len(cursors_newInstances)
	lenDeletedInstances += len(cursors_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		forwardCommit += "\n\tstage.Commit()\n"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		backwardCommit += "\n\tstage.Commit()\n"
		// append to the start of the backward commits slice
		stage.backwardCommits = append([]string{backwardCommit}, stage.backwardCommits...)

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

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Cursors_reference = make(map[*Cursor]*Cursor)
	stage.Cursors_referenceOrder = make(map[*Cursor]uint) // diff Unstage needs the reference order
	for instance := range stage.Cursors {
		stage.Cursors_reference[instance] = instance.GongCopy().(*Cursor)
		stage.Cursors_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (cursor *Cursor) GongGetOrder(stage *Stage) uint {
	return stage.CursorMap_Staged_Order[cursor]
}

func (cursor *Cursor) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Cursors_referenceOrder[cursor]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (cursor *Cursor) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cursor.GongGetGongstructName(), cursor.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cursor *Cursor) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cursor.GongGetGongstructName(), cursor.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (cursor *Cursor) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cursor.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Cursor")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cursor.Name)
	return
}

// insertion point for unstaging
func (cursor *Cursor) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cursor.GongGetReferenceIdentifier(stage))
	return
}
