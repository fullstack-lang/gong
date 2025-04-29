package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) UpdateAndCommitTreeStage() {

	stager.treeStage.Reset()

	tree.StageBranch(stager.treeStage,
		&tree.Tree{
			Name: stager.stage.GetProbeTreeSidebarStageName(),
			RootNodes: []*tree.Node{
				{
					Name:       "Class Diagrams",
					IsExpanded: true,
				},
			},
		},
	)

	stager.treeStage.Commit()
}

type NewClassdiagramButtonProxy struct {
	stager *Stager
}
