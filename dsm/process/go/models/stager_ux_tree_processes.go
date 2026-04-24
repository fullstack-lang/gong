package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeProcessesWithinDiagram(diagramProcess *DiagramProcess, process *Process, parentNode *tree.Node) {
	processNode := addNodeToTree(
		stager,
		diagramProcess,
		parentNode,
		process,
		process.parentProcess,
		&diagramProcess.ProcesssWhoseNodeIsExpanded,
		&diagramProcess.Process_Shapes,
		diagramProcess.map_Process_ProcessShape,
		diagramProcess.map_Process_ProcessCompositionShape,
		&diagramProcess.ProcessComposition_Shapes,
	)

	addAddItemButton(stager, &diagramProcess.ProcesssWhoseNodeIsExpanded, process, nil, processNode, &process.SubProcesses, diagramProcess, &diagramProcess.Process_Shapes, &diagramProcess.ProcessComposition_Shapes)

}
