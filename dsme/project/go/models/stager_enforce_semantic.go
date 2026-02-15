package models

import (
	"slices"
	"time"
)

func (stager *Stager) enforceSemantic() (needCommit bool) {
	stage := stager.stage

	// VERY important because the probe only unstages objects
	// this is the Clean that delete them from slices and pointers that reference
	// them. If the checkout is not performed, the stage might be dirty
	// with slices of pointer or pointer to unstaged instance
	needCommit = stage.Clean() || needCommit

	// Ensures that there is one and only one root
	// prune the other
	// check that there is at least one root
	// and that one can safely access stager.root
	roots := GetGongstrucsSorted[*Root](stager.stage)
	if len(roots) == 0 {
		stager.root = (&Root{Name: "Root"}).Stage(stager.stage)
		needCommit = true
	} else {
		stager.root = roots[0]

		if len(roots) > 1 {
			for _, root := range roots[1:] {
				root.Unstage(stage)
				needCommit = true
			}
		}
	}

	root := stager.root

	// Enforce that all projects are appended to the [root]
	// if one project is not appended, append it
	for _, project := range GetGongstrucsSorted[*Project](stage) {
		if slices.Index(root.Projects, project) == -1 {
			root.Projects = append(root.Projects, project)
			needCommit = true
		}
	}

	needCommit = stager.enforceObjectValues() || needCommit
	needCommit = stager.enforceDAG() || needCommit
	needCommit = stager.enforceHierarchy() || needCommit
	needCommit = stager.enforceUniquenessInProjects() || needCommit
	needCommit = stager.unstageAllOrphans() || needCommit
	needCommit = stager.enforceComputedPrefix() || needCommit

	// Semantic for shapes
	{
		// enforce visibility will unstage shapes that are not visible
		// and indirectly shapes whose abstract element is not staged (because they are not visible)
		needCommit = stager.enforceVisibility() || needCommit
		needCommit = stager.enforceRelationDuplicates() || needCommit
		needCommit = stager.enforceNodeShapeDuplicates() || needCommit
		needCommit = stager.enforceShapeOrphans() || needCommit
		needCommit = stager.enforceShapesAbstractConsistency(stage, needCommit) || needCommit
	}

	needCommit = stager.enforceTaskInputOutputProjectConsistency() || needCommit
	needCommit = stager.enforceDuplicateRemove() || needCommit

	// computes fields that are not persisted
	stager.enforceProducersConsumers()
	stager.enforceDiagramMaps()

	if needCommit {
		stager.stage.Clean()

		stager.enforceProducersConsumers()
		stager.enforceDiagramMaps()

		stager.stage.CommitWithSuspendedCallbacks()
		stager.probeForm.AddNotification(time.Now(), "Stage was modified to enforce semantic, please check the diagram for details")
		stager.probeForm.CommitNotificationTable()

		{
			// for debug purpose, to check that the stage is not dirty after the commit
			var selectedDiagram *Diagram
			for _, diagram := range GetGongstrucsSorted[*Diagram](stage) {
				if diagram.IsEditable() {
					selectedDiagram = diagram
					break
				}
			}
			if selectedDiagram != nil {
				TaskOutputShapes := selectedDiagram.TaskOutputShapes
				_ = TaskOutputShapes
			}
		}
	}

	return
}
