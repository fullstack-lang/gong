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
	validTaskPredecessors := make(map[taskPredecessorKey]bool)
	for _, task := range stage.GetInstancesSorted[*Task]() {
		for _, prod := range task.Inputs {
			validTaskInputs[taskProductKey{Task: task, Product: prod}] = true
		}
		for _, prod := range task.Outputs {
			validTaskOutputs[taskProductKey{Task: task, Product: prod}] = true
		}
		for _, pred := range task.Predecessors {
			validTaskPredecessors[taskPredecessorKey{Task: task, Predecessor: pred}] = true
		}
	}

	validNoteProducts := make(map[noteProductKey]bool)
	validNoteTasks := make(map[noteTaskKey]bool)
	validNoteResources := make(map[noteResourceKey]bool)
	for _, note := range stage.GetInstancesSorted[*Note]() {
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
	for _, res := range stage.GetInstancesSorted[*Resource]() {
		for _, task := range res.Tasks {
			validResourceTasks[resourceTaskKey{Resource: res, Task: task}] = true
		}
	}

	validTaskCompositions := make(map[*Task]bool)
	for _, task := range stage.GetInstancesSorted[*Task]() {
		for _, subTask := range task.SubTasks {
			validTaskCompositions[subTask] = true
		}
	}

	validProductCompositions := make(map[*Product]bool)
	validProductReferences := make(map[productReferenceKey]bool)
	for _, product := range stage.GetInstancesSorted[*Product]() {
		for _, subProduct := range product.SubProducts {
			validProductCompositions[subProduct] = true
		}
		if product.ReferencedProduct != nil {
			validProductReferences[productReferenceKey{Product: product, ReferencedProduct: product.ReferencedProduct}] = true
		}
	}

	validResourceCompositions := make(map[*Resource]bool)
	for _, resource := range stage.GetInstancesSorted[*Resource]() {
		for _, subResource := range resource.SubResources {
			validResourceCompositions[subResource] = true
		}
	}

	for _, shape := range stage.GetInstancesSorted[*TaskInputShape]() {
		if shape.Task != nil && shape.Product != nil {
			if !validTaskInputs[taskProductKey{Task: shape.Task, Product: shape.Product}] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid TaskInputShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	for _, shape := range stage.GetInstancesSorted[*TaskOutputShape]() {
		if shape.Task != nil && shape.Product != nil {
			if !validTaskOutputs[taskProductKey{Task: shape.Task, Product: shape.Product}] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid TaskOutputShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	for _, shape := range stage.GetInstancesSorted[*TaskPredecessorShape]() {
		if shape.Task != nil && shape.Predecessor != nil {
			if !validTaskPredecessors[taskPredecessorKey{Task: shape.Task, Predecessor: shape.Predecessor}] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid TaskPredecessorShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	for _, shape := range stage.GetInstancesSorted[*NoteProductShape]() {
		if shape.Note != nil && shape.Product != nil {
			if !validNoteProducts[noteProductKey{Note: shape.Note, Product: shape.Product}] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid NoteProductShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	for _, shape := range stage.GetInstancesSorted[*NoteTaskShape]() {
		if shape.Note != nil && shape.Task != nil {
			if !validNoteTasks[noteTaskKey{Note: shape.Note, Task: shape.Task}] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid NoteTaskShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	for _, shape := range stage.GetInstancesSorted[*NoteResourceShape]() {
		if shape.Note != nil && shape.Resource != nil {
			if !validNoteResources[noteResourceKey{Note: shape.Note, Resource: shape.Resource}] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid NoteResourceShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	for _, shape := range stage.GetInstancesSorted[*ResourceTaskShape]() {
		if shape.Resource != nil && shape.Task != nil {
			if !validResourceTasks[resourceTaskKey{Resource: shape.Resource, Task: shape.Task}] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid ResourceTaskShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	for _, shape := range stage.GetInstancesSorted[*TaskCompositionShape]() {
		if shape.Task != nil {
			if !validTaskCompositions[shape.Task] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid TaskCompositionShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	for _, shape := range stage.GetInstancesSorted[*ProductCompositionShape]() {
		if shape.Product != nil {
			if !validProductCompositions[shape.Product] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid ProductCompositionShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	for _, shape := range stage.GetInstancesSorted[*ProductReferenceShape]() {
		if shape.Product != nil && shape.ReferencedProduct != nil {
			if !validProductReferences[productReferenceKey{Product: shape.Product, ReferencedProduct: shape.ReferencedProduct}] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid ProductReferenceShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	for _, shape := range stage.GetInstancesSorted[*ResourceCompositionShape]() {
		if shape.Resource != nil {
			if !validResourceCompositions[shape.Resource] {
				shape.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged invalid ResourceCompositionShape %s", shape.GetName()))
				}
				needCommit = true
			}
		}
	}

	return needCommit
}
