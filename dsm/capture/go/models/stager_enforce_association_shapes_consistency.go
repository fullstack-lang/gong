package models

import (
	"fmt"
	"time"
)

// enforceAssociationShapeConsistency check that all associations shapes, like TaskOutputShape
// are matched with a an abstract relation (Task-->Deliverable)
func (stager *Stager) enforceAssociationShapeConsistency() bool {
	var needCommit bool
	stage := stager.stage

	validTaskInputs := make(map[concernDeliverableKey]bool)
	validTaskOutputs := make(map[concernDeliverableKey]bool)
	for _, task := range GetGongstrucsSorted[*Concern](stage) {
		for _, prod := range task.Inputs {
			validTaskInputs[concernDeliverableKey{Concern: task, Deliverable: prod}] = true
		}
		for _, prod := range task.Outputs {
			validTaskOutputs[concernDeliverableKey{Concern: task, Deliverable: prod}] = true
		}
	}

	validDeliverableConcepts := make(map[deliverableConceptKey]bool)
	for _, deliverable := range GetGongstrucsSorted[*Deliverable](stage) {
		for _, concept := range deliverable.Concepts {
			validDeliverableConcepts[deliverableConceptKey{Deliverable: deliverable, Concept: concept}] = true
		}
	}

	validNoteDeliverables := make(map[noteDeliverableKey]bool)
	validNoteTasks := make(map[noteTaskKey]bool)
	validNoteResources := make(map[noteResourceKey]bool)
	for _, note := range GetGongstrucsSorted[*Note](stage) {
		for _, prod := range note.Deliverables {
			validNoteDeliverables[noteDeliverableKey{Note: note, Deliverable: prod}] = true
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
			if !validTaskInputs[concernDeliverableKey{Concern: shape.Concern, Deliverable: shape.Deliverable}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid TaskInputShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*ConcernOutputShape](stage) {
		if shape.Concern != nil && shape.Deliverable != nil {
			if !validTaskOutputs[concernDeliverableKey{Concern: shape.Concern, Deliverable: shape.Deliverable}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid TaskOutputShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*NoteDeliverableShape](stage) {
		if shape.Note != nil && shape.Deliverable != nil {
			if !validNoteDeliverables[noteDeliverableKey{Note: shape.Note, Deliverable: shape.Deliverable}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid NoteDeliverableShape %s", shape.GetName()))
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

	for _, shape := range GetGongstrucsSorted[*DeliverableConceptShape](stage) {
		if shape.Deliverable != nil && shape.Concept != nil {
			if !validDeliverableConcepts[deliverableConceptKey{Deliverable: shape.Deliverable, Concept: shape.Concept}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid DeliverableConceptShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	return needCommit
}
