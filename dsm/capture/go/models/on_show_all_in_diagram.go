package models

import (
)

// onShowAllInDiagram resets the diagram, creates and add all all possible concrete instances to the diagram,
// and organize them in a pseudo pert diagram
func onShowAllInDiagram(stager *Stager, diagram *Diagram) func() {
	return func() {
		// 1. Reset the diagram: remove all shapes from the stage to start fresh
		diagram.Deliverable_Shapes = []*DeliverableShape{}
		diagram.Concern_Shapes = []*ConcernShape{}
		diagram.Note_Shapes = []*NoteShape{}

		// 2. Compute the rank of each node (Deliverable or Task)
		// The rank is determined by the longest path from a root node.
		// Rank 0 = Root. Rank N = Dependencies have Max(Rank) = N-1.
		map_Node_Rank := make(map[any]int)

		project := stager.stage.Library_Diagrams_reverseMap[diagram]

		// collect all Deliverables and Tasks that are reachable from the Project
		deliverables := collectProjectElements(project.RootDeliverables, func(p *Deliverable) []*Deliverable {
			return p.SubDeliverables
		})

		tasks := collectProjectElements(project.RootConcerns, func(t *Concern) []*Concern {
			return t.SubConcerns
		})

		// Relax edges to propagate ranks.
		// We loop enough times to cover the maximum possible path length (number of nodes).
		// Edges considered:
		// - Deliverable -> SubDeliverable
		// - Task -> SubTask
		// - Task -> Output (Deliverable)
		// - Deliverable -> Input (Task)
		nbNodes := len(deliverables) + len(tasks)
		for i := 0; i < nbNodes; i++ {

			// Deliverable Composition: Parent Rank < Child Rank
			for _, deliverable := range deliverables {
				for _, subDeliverable := range deliverable.SubDeliverables {
					if map_Node_Rank[subDeliverable] <= map_Node_Rank[deliverable] {
						map_Node_Rank[subDeliverable] = map_Node_Rank[deliverable] + 1
					}
				}
			}

			// Task Composition: Parent Rank < Child Rank
			for _, task := range tasks {
				for _, subTask := range task.SubConcerns {
					if map_Node_Rank[subTask] <= map_Node_Rank[task] {
						map_Node_Rank[subTask] = map_Node_Rank[task] + 1
					}
				}
			}

			// Task Output: Task Rank < Output Deliverable Rank
			for _, task := range tasks {
				for _, output := range task.Outputs {
					if map_Node_Rank[output] <= map_Node_Rank[task] {
						map_Node_Rank[output] = map_Node_Rank[task] + 1
					}
				}
			}

			// Task Input: Input Deliverable Rank < Task Rank
			for _, task := range tasks {
				for _, input := range task.Inputs {
					if map_Node_Rank[task] <= map_Node_Rank[input] {
						map_Node_Rank[task] = map_Node_Rank[input] + 1
					}
				}
			}
		}

		// 3. Create and Position Shapes
		const (
			X_STEP   = 300.0
			Y_STEP   = 100.0
			OFFSET_X = 50.0
			OFFSET_Y = 50.0
		)

		// Helper to keep track of vertical stacking within a rank column
		rankCounters := make(map[int]int)

		// Create Deliverable Shapes
		for _, deliverable := range deliverables {
			rank := map_Node_Rank[deliverable]
			index := rankCounters[rank]
			rankCounters[rank]++

			x := OFFSET_X + float64(rank)*X_STEP
			y := OFFSET_Y + float64(index)*Y_STEP

			// Note: We do not Stage the shape here, StageBranch will handle it later
			shape := &DeliverableShape{Name: deliverable.Name}
			shape.Deliverable = deliverable
			shape.X = x
			shape.Y = y
			shape.Width = 200
			shape.Height = 60
			diagram.Deliverable_Shapes = append(diagram.Deliverable_Shapes, shape)
		}

		// Create Task Shapes
		for _, task := range tasks {
			rank := map_Node_Rank[task]
			index := rankCounters[rank]
			rankCounters[rank]++

			x := OFFSET_X + float64(rank)*X_STEP
			y := OFFSET_Y + float64(index)*Y_STEP

			// Note: We do not Stage the shape here, StageBranch will handle it later
			shape := &ConcernShape{Name: task.Name}
			shape.Concern = task
			shape.X = x
			shape.Y = y
			shape.Width = 200
			shape.Height = 60
			diagram.Concern_Shapes = append(diagram.Concern_Shapes, shape)
		}

		// 4. Create Link Shapes using newConcreteAssociation
		// Deliverable Composition
		for _, deliverable := range deliverables {
			for _, subDeliverable := range deliverable.SubDeliverables {
				newConcreteAssociation(
					deliverable, subDeliverable, &diagram.DeliverableComposition_Shapes)
			}
		}

		// Task Composition
		for _, task := range tasks {
			for _, subTask := range task.SubConcerns {
				newConcreteAssociation(
					task, subTask, &diagram.ConcernComposition_Shapes)
			}
		}

		// Task Inputs (Deliverable -> Task) and Outputs (Task -> Deliverable)
		for _, task := range tasks {
			// Task Inputs: Task is the concrete start, Deliverable is the concrete end
			for _, input := range task.Inputs {
				newConcreteAssociation(
					task, input, &diagram.ConcernInputShapes)
			}
			// Task Outputs: Task is the concrete start, Deliverable is the concrete end
			for _, output := range task.Outputs {
				newConcreteAssociation(
					task, output, &diagram.ConcernOutputShapes)
			}
		}

		// 5. Unstage the diagram and re-stage it with StageBranch to include all new shapes recursively
		diagram.Unstage(stager.stage)
		StageBranch(stager.stage, diagram)

		// Save changes
		stager.stage.Commit()
	}
}
