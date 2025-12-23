package models

import (
	"fmt"
	"log"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeWBS(task *Task, parentNode *tree.Node) {
	taskNode := &tree.Node{
		Name:            task.ComputedPrefix + " " + task.Name,
		IsExpanded:      task.IsExpanded,
		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, taskNode)

	taskNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: stager.OnUpdateTask(task),
	}

	addAddItemButton(stager, taskNode, &task.SubTasks)

	for _, task := range task.SubTasks {
		stager.treeWBS(task, taskNode)
	}

	if len(task.Inputs) > 0 {
		inputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(task.Inputs)),
			IsExpanded:           task.IsInputsNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		taskNode.Children = append(taskNode.Children, inputProductsNode)

		inputProductsNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateExpansion(&task.IsInputsNodeExpanded),
		}

		for _, product := range task.Inputs {
			inputProductNode := &tree.Node{
				Name:            product.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			inputProductsNode.Children = append(inputProductsNode.Children, inputProductNode)

			inputProductNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: stager.OnUpdateProduct(product),
			}
		}
	}

	if len(task.Outputs) > 0 {
		outputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(task.Outputs)),
			IsExpanded:           task.IsOutputsNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		taskNode.Children = append(taskNode.Children, outputProductsNode)

		outputProductsNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateExpansion(&task.IsOutputsNodeExpanded),
		}

		for _, product := range task.Outputs {
			outputProductNode := &tree.Node{
				Name:            product.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			outputProductsNode.Children = append(outputProductsNode.Children, outputProductNode)

			outputProductNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: stager.OnUpdateProduct(product),
			}
		}
	}
}

func (stager *Stager) treeWBSinDiagram(diagram *Diagram, task *Task, parentNode *tree.Node) {
	taskNode := &tree.Node{
		Name:       task.ComputedPrefix + " " + task.Name,
		IsExpanded: slices.Index(diagram.TasksWhoseNodeIsExpanded, task) != -1,

		HasCheckboxButton:  true,
		IsCheckboxDisabled: !diagram.IsChecked,

		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		ToolTipText:     "Add task to diagram",

		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, taskNode)

	if _, ok := diagram.map_Task_TaskShape[task]; ok {
		taskNode.IsChecked = true
	}

	// what to do when the task node is clicked
	taskNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: stager.OnUpdateTaskInDiagram(diagram, task),
	}

	// if task has a parent task, add a button to show/hide the link to the parent
	if parentTask := task.parentTask; parentTask != nil {
		if _, ok := diagram.map_Task_TaskShape[parentTask]; ok {
			if _, ok := diagram.map_Task_TaskShape[task]; ok {

				showHideCompositionButton := &tree.Button{
					Name: GetGongstructNameFromPointer(task) + " " + string(buttons.BUTTON_add),

					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if compositionShape, ok := diagram.map_Task_TaskCompositionShape[task]; !ok {
					showHideCompositionButton.Icon = string(buttons.BUTTON_unfold_more)
					showHideCompositionButton.ToolTipText = "Show link from \"" + parentTask.Name +
						"\" to \"" + task.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: stager.OnAddTaskCompositionShape(diagram, parentTask, task),
					}
				} else {
					showHideCompositionButton.Icon = string(buttons.BUTTON_unfold_less)
					showHideCompositionButton.ToolTipText = "Hide link from \"" + parentTask.Name +
						"\" to \"" + task.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: stager.OnRemoveTaskCompositionShape(diagram, compositionShape),
					}
				}
				taskNode.Buttons = append(taskNode.Buttons, showHideCompositionButton)
			}
		}
	}

	for _, task := range task.SubTasks {
		stager.treeWBSinDiagram(diagram, task, taskNode)
	}
}

func (stager *Stager) OnUpdateTaskInDiagram(diagram *Diagram, task *Task) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		// find the shape (if any)
		taskShape := diagram.map_Task_TaskShape[task]

		if frontNode.IsChecked && !stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if taskShape != nil {
				log.Panic("adding a shape to an already task shape")
			}
			taskShape = newTaskShapeToDiagram(task, diagram).Stage(stager.stage)
			stager.stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if taskShape == nil {
				log.Panic("remove a non existing shape to task")
			}
			taskShape.Unstage(stager.stage)
			idx := slices.Index(diagram.Task_Shapes, taskShape)
			diagram.Task_Shapes = slices.Delete(diagram.Task_Shapes, idx, idx+1)
			stager.stage.Commit()
			return
		}

		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			if frontNode.IsExpanded {
				if slices.Index(diagram.TasksWhoseNodeIsExpanded, task) == -1 {
					diagram.TasksWhoseNodeIsExpanded = append(diagram.TasksWhoseNodeIsExpanded, task)
				}
			} else {
				if idx := slices.Index(diagram.TasksWhoseNodeIsExpanded, task); idx != -1 {
					diagram.TasksWhoseNodeIsExpanded = slices.Delete(diagram.TasksWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			return
		}

		stager.probeForm.FillUpFormFromGongstruct(task, GetPointerToGongstructName[*Task]())
		stager.stage.Commit()
	}
}

func (stager *Stager) OnAddTaskCompositionShape(diagram *Diagram, parentTask, task *Task) func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		compositionShape := new(TaskCompositionShape).Stage(stager.stage)
		compositionShape.Name = parentTask.Name + " to " + task.Name
		compositionShape.Task = task
		compositionShape.StartOrientation = ORIENTATION_VERTICAL
		compositionShape.EndOrientation = ORIENTATION_VERTICAL
		compositionShape.CornerOffsetRatio = 1.68 // near the golden ratio
		compositionShape.StartRatio = 0.5
		compositionShape.EndRatio = 0.5

		diagram.TaskComposition_Shapes = append(diagram.TaskComposition_Shapes, compositionShape)
		stager.stage.Commit()
	}
}

func (stager *Stager) OnRemoveTaskCompositionShape(diagram *Diagram, compositionShape *TaskCompositionShape) func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		compositionShape.Unstage(stager.stage)
		idx := slices.Index(diagram.TaskComposition_Shapes, compositionShape)
		diagram.TaskComposition_Shapes = slices.Delete(diagram.TaskComposition_Shapes, idx, idx+1)
		stager.stage.Commit()
	}
}
