package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) updateTaskTreeStage() {

	stager.treeTasksStage.Reset()

	root := stager.root

	treeInstance := &tree.Tree{Name: "PBS"}

	allProjectsNode := &tree.Node{
		Name:       "** Tree of Projects **",
		IsExpanded: true,
	}
	treeInstance.RootNodes = append(treeInstance.RootNodes, allProjectsNode)

	for _, project := range root.Projects {
		projectNode := &tree.Node{
			Name:            project.Name,
			IsExpanded:      project.IsExpanded,
			IsNodeClickable: true,
		}
		treeInstance.RootNodes = append(treeInstance.RootNodes, projectNode)
		projectNode.Impl = &NodeProxy[*Project]{
			stager:   stager,
			node:     projectNode,
			instance: project,
		}

		addAddItemButton(stager, projectNode, &project.RootTasks,
			func(items *[]*Task, item *Task) {
				*items = append(*items, item)
			})

		for _, task := range project.RootTasks {
			stager.generateTreeOfTask(task, projectNode)
		}
	}

	if len(root.OrphanedTasks) > 0 {
		orphansTaskNode := &tree.Node{Name: "Orphans Tasks", IsExpanded: true}
		treeInstance.RootNodes = append(treeInstance.RootNodes, orphansTaskNode)
		for _, task := range root.OrphanedTasks {
			stager.generateTreeOfTask(task, orphansTaskNode)
		}
	}

	tree.StageBranch(stager.treeTasksStage, treeInstance)

	stager.treeTasksStage.Commit()
}

func (stager *Stager) generateTreeOfTask(task *Task, parentNode *tree.Node) {

	taskNode := &tree.Node{
		Name:            task.ComputedPrefix + " " + task.Name,
		IsExpanded:      true,
		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, taskNode)

	taskNode.Impl = &NodeProxy[*Task]{
		stager:   stager,
		node:     taskNode,
		instance: task,
	}

	addAddItemButton(stager, taskNode, &task.SubTasks,
		func(items *[]*Task, item *Task) {
			*items = append(*items, item)
		})

	for _, task := range task.SubTasks {
		stager.generateTreeOfTask(task, taskNode)
	}
}
