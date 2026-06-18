package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/maticons/maticons"
)

func (stager *Stager) ux_tree() {
	stager.treeStage.Reset()

	treeInstance := &tree.Tree{Name: "Sidebar"}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	list := GetGongstrucsSorted[*Analysis](stager.stage)

	for _, analysis := range list {
		stager.treeAnalysis(treeInstance, analysis, &treeInstance.RootNodes)
	}

	// add a node for the Add Analysis
	newAnalysisNode := new(tree.Node)
	newAnalysisNode.Name = "New Analysis"
	newAnalysisNode.FontStyle = tree.ITALIC
	newAnalysisNode.IsWithPreceedingIcon = true
	newAnalysisNode.PreceedingIcon = string(maticons.BUTTON_library_books)
	treeInstance.RootNodes = append(treeInstance.RootNodes, newAnalysisNode)

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}
