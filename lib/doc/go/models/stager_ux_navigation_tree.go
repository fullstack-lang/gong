package models

import (
	"fmt"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) navigationTree() {

	stager.treeNavigationStage.Reset()

	tree := &tree_models.Tree{
		Name: stager.stage.GetProbeNavigationTreeSidebarStageName(),
	}

	stageOfInterest := stager.stage

	deltaNode := &tree_models.Node{}
	tree.RootNodes = append(tree.RootNodes, deltaNode)

	backwardButton := &tree_models.Button{
		Name:       "BackwardButton",
		Icon:       string(tree_buttons.BUTTON_undo),
		HasToolTip: true,
		ToolTipText: fmt.Sprintf("Go to previous commit (%d/%d)",
			len(stageOfInterest.GetBackwardCommits()), stageOfInterest.GetCommitsBehind()),
		ToolTipPosition: tree_models.Below,
		OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
			err := stageOfInterest.ApplyBackwardCommit()
			if err != nil {
				panic(err)
			}
			// probe.Refresh()
		},
	}
	deltaNode.Buttons = append(deltaNode.Buttons, backwardButton)

	if stageOfInterest.GetCommitsBehind() == len(stageOfInterest.GetBackwardCommits()) {
		backwardButton.IsDisabled = true
		backwardButton.ToolTipText = "No more previous commits"
	}

	forwardButton := &tree_models.Button{
		Name:       "ForwardButton",
		Icon:       string(tree_buttons.BUTTON_redo),
		HasToolTip: true,
		ToolTipText: fmt.Sprintf("Go to next commit (%d/%d)",
			len(stageOfInterest.GetBackwardCommits()), stageOfInterest.GetCommitsBehind()),
		ToolTipPosition: tree_models.Below,
		OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
			err := stageOfInterest.ApplyForwardCommit()
			if err != nil {
				panic(err)
			}
			// probe.Refresh()
		},
	}
	deltaNode.Buttons = append(deltaNode.Buttons, forwardButton)

	if stageOfInterest.GetCommitsBehind() == 0 {
		forwardButton.IsDisabled = true
		forwardButton.ToolTipText = "No more next commits"
	}

	if stageOfInterest.GetCommitsBehind() > 0 {
		discardButton := &tree_models.Button{
			Name:            "DiscardButton",
			Icon:            string(tree_buttons.BUTTON_cancel),
			HasToolTip:      true,
			ToolTipText:     "Discard commits ahead (git reset --hard HEAD)",
			ToolTipPosition: tree_models.Below,
			OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
				stageOfInterest.ResetHard()
				// probe.Refresh()
			},
		}
		deltaNode.Buttons = append(deltaNode.Buttons, discardButton)
	}

	squashButton := &tree_models.Button{
		Name:            "SqueashButton",
		Icon:            string(tree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Discard all commits history (git orphan)",
		ToolTipPosition: tree_models.Below,
		OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
			stageOfInterest.Squash()
			// probe.Refresh()
		},
	}
	if len(stageOfInterest.GetBackwardCommits()) == 0 {
		squashButton.IsDisabled = true
		squashButton.ToolTipText = "No commits to orphan"
	}
	deltaNode.Buttons = append(deltaNode.Buttons, squashButton)

	tree_models.StageBranch(stager.treeNavigationStage, tree)

	stager.treeNavigationStage.Commit()
}
