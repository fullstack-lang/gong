package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeProcessesWithinDiagramProcessWithinParticipant(
	diagramProcess *DiagramProcess,
	process *Process,
	participant *Participant,
	parentNode *tree.Node) {

	stage := stager.stage
	key := allocatedProcessShapeKey{
		participant: participant,
		process:     process,
	}
	allocatedProcessShape, ok := diagramProcess.map_AllocatedProcessShapeKey_AllocatedProcessShape[key]

	node := &tree.Node{
		Name:              process.Name,
		IsNodeClickable:   true,
		IsInEditMode:      process.GetIsInRenameMode(),
		IsChecked:         ok,
		HasCheckboxButton: true,
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(process, node, stager)

	node.OnNameChange = stager.onNameChange(process)
	node.OnClick = onNodeClicked(stager, process)

	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !ok {
			allocatedProcessShape = (&AllocatedProcessShape{
				Name:        diagramProcess.GetName() + "-" + participant.GetName() + "-" + process.GetName(),
				Participant: participant,
				Process:     process,
			}).Stage(stage)
			diagramProcess.AllocatedProcessShapes = append(diagramProcess.AllocatedProcessShapes, allocatedProcessShape)
			stage.Commit()
		} else if !isChecked && ok {
			allocatedProcessShape.Unstage(stage)
			stage.Commit()
		}
	}
}
