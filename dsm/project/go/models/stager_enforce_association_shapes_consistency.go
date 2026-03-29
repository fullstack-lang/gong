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

	validTaskInputs := make(map[taskProductKey]bool)
	validTaskOutputs := make(map[taskProductKey]bool)
	for _, task := range GetGongstrucsSorted[*Task](stage) {
		for _, prod := range task.Inputs {
			validTaskInputs[taskProductKey{Task: task, Product: prod}] = true
		}
		for _, prod := range task.Outputs {
			validTaskOutputs[taskProductKey{Task: task, Product: prod}] = true
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
			validNoteResources[noteResourceKey{Note: note, Resource: res}] = true
		}
	}

	validResourceTasks := make(map[resourceTaskKey]bool)
	for _, res := range GetGongstrucsSorted[*Resource](stage) {
		for _, task := range res.Tasks {
			validResourceTasks[resourceTaskKey{Resource: res, Task: task}] = true
		}
	}

	for _, shape := range GetGongstrucsSorted[*TaskInputShape](stage) {
		if shape.Task != nil && shape.Product != nil {
			if !validTaskInputs[taskProductKey{Task: shape.Task, Product: shape.Product}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid TaskInputShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*TaskOutputShape](stage) {
		if shape.Task != nil && shape.Product != nil {
			if !validTaskOutputs[taskProductKey{Task: shape.Task, Product: shape.Product}] {
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

	for _, shape := range GetGongstrucsSorted[*NoteResourceShape](stage) {
		if shape.Note != nil && shape.Resource != nil {
			if !validNoteResources[noteResourceKey{Note: shape.Note, Resource: shape.Resource}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid NoteResourceShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*ResourceTaskShape](stage) {
		if shape.Resource != nil && shape.Task != nil {
			if !validResourceTasks[resourceTaskKey{Resource: shape.Resource, Task: shape.Task}] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid ResourceTaskShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	return needCommit
}
