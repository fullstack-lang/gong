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

	validTaskCompositions := make(map[*Task]bool)
	for _, task := range GetGongstrucsSorted[*Task](stage) {
		for _, subTask := range task.SubTasks {
			validTaskCompositions[subTask] = true
		}
	}

	validProductCompositions := make(map[*Product]bool)
	for _, product := range GetGongstrucsSorted[*Product](stage) {
		for _, subProduct := range product.SubProducts {
			validProductCompositions[subProduct] = true
		}
	}

	validResourceCompositions := make(map[*Resource]bool)
	for _, resource := range GetGongstrucsSorted[*Resource](stage) {
		for _, subResource := range resource.SubResources {
			validResourceCompositions[subResource] = true
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

	for _, shape := range GetGongstrucsSorted[*TaskCompositionShape](stage) {
		if shape.Task != nil {
			if !validTaskCompositions[shape.Task] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid TaskCompositionShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*ProductCompositionShape](stage) {
		if shape.Product != nil {
			if !validProductCompositions[shape.Product] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid ProductCompositionShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	for _, shape := range GetGongstrucsSorted[*ResourceCompositionShape](stage) {
		if shape.Resource != nil {
			if !validResourceCompositions[shape.Resource] {
				shape.UnstageVoid(stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid ResourceCompositionShape %s", shape.GetName()))
				needCommit = true
			}
		}
	}

	return needCommit
}
