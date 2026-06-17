package models

import (
	"fmt"
	"time"
)

// enforceAssociationShapeConsistency check that all associations shapes, like TaskOutputShape
// are matched with a an abstract relation (Task-->Product)
func (stager *Stager) enforceAssociationShapeConsistency() bool {
	var needCommit bool
	stage := stager.stage

	validTaskInputs := make(map[concernProductKey]bool)
	validTaskOutputs := make(map[concernProductKey]bool)
	for _, task := range GetGongstrucsSorted[*Concern](stage) {
		for _, prod := range task.Inputs {
			validTaskInputs[concernProductKey{Concern: task, Product: prod}] = true
		}
		for _, prod := range task.Outputs {
			validTaskOutputs[concernProductKey{Concern: task, Product: prod}] = true
		}
	}

	validNoteProducts := make(map[noteProductKey]bool)
	validNoteTasks := make(map[noteTaskKey]bool)
	validNoteResources := make(map[noteResourceKey]bool)
	for _, note := range GetGongstrucsSorted[*Note](stage) {
		for _, prod := range note.Products {
			validNoteProducts[noteProductKey{Note: note, Product: prod}] = true
		}
		for _, task := range note.Tasks {
			validNoteTasks[noteTaskKey{Note: note, Task: task}] = true
		}
		for _, res := range note.Resources {
			validNoteResources[noteResourceKey{Note: note, Stakeholder: res}] = true
		}
	}

	validResourceTasks := make(map[stakeholderConcernKey]bool)
	for _, res := range GetGongstrucsSorted[*Stakeholder](stage) {
		for _, task := range res.Concerns {
			validResourceTasks[stakeholderConcernKey{Stakeholder: res, Concern: task}] = true
		}
	}

	for _, shape := range GetGongstrucsSorted[*ConcernInputShape](stage) {
		if shape.Concern != nil && shape.Deliverable != nil {
			if !validTaskInputs[concernProductKey{Concern: shape.Concern, Product: shape.Deliverable}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid TaskInputShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*ConcernOutputShape](stage) {
		if shape.Task != nil && shape.Product != nil {
			if !validTaskOutputs[concernProductKey{Concern: shape.Task, Product: shape.Product}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid TaskOutputShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*NoteProductShape](stage) {
		if shape.Note != nil && shape.Product != nil {
			if !validNoteProducts[noteProductKey{Note: shape.Note, Product: shape.Product}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid NoteProductShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*NoteTaskShape](stage) {
		if shape.Note != nil && shape.Task != nil {
			if !validNoteTasks[noteTaskKey{Note: shape.Note, Task: shape.Task}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid NoteTaskShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*NoteStakeholderShape](stage) {
		if shape.Note != nil && shape.Stakeholder != nil {
			if !validNoteResources[noteResourceKey{Note: shape.Note, Stakeholder: shape.Stakeholder}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid NoteResourceShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*StakeholderConcernShape](stage) {
		if shape.Stakeholder != nil && shape.Concern != nil {
			if !validResourceTasks[stakeholderConcernKey{Stakeholder: shape.Stakeholder, Concern: shape.Concern}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid ResourceTaskShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	return needCommit
}
