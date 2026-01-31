package probe

import (
	"fmt"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/gong/test/test3/go/models"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (probe *Probe) AddCommitNavigationNode(appendChildrenNodeFunc func(models.GongNodeIF)) {

	stageOfInterest := probe.stageOfInterest

	deltaNode := &tree.Node{}
	deltaNode.Name += fmt.Sprintf(".... <-->   %d   <-   %d ",
		len(stageOfInterest.GetBackwardCommits()),
		stageOfInterest.GetCommitsBehind())

	backwardButton := &tree.Button{
		Name:            "BackwardButton",
		Icon:            string(gongtree_buttons.BUTTON_arrow_back),
		HasToolTip:      true,
		ToolTipText:     "Go to previous commit",
		ToolTipPosition: tree.Below,
	}
	deltaNode.Buttons = append(deltaNode.Buttons, backwardButton)
	backwardButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage,
			stagedButton, frontButton *tree.Button) {
			err := stageOfInterest.ApplyBackwardCommit()
			if err != nil {
				panic(err)
			}
			probe.Refresh()
		},
	}

	if stageOfInterest.GetCommitsBehind() == len(stageOfInterest.GetBackwardCommits()) {
		backwardButton.IsDisabled = true
		backwardButton.Icon = string(gongtree_buttons.BUTTON_do_not_disturb)
		backwardButton.ToolTipText = "No more previous commits"
	}

	forwardButton := &tree.Button{
		Name:            "ForwardButton",
		Icon:            string(gongtree_buttons.BUTTON_arrow_forward),
		HasToolTip:      true,
		ToolTipText:     "Go to next commit",
		ToolTipPosition: tree.Below,
	}
	deltaNode.Buttons = append(deltaNode.Buttons, forwardButton)
	forwardButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage,
			stagedButton, frontButton *tree.Button) {
			err := stageOfInterest.ApplyForwardCommit()
			if err != nil {
				panic(err)
			}
			probe.Refresh()
		},
	}

	if stageOfInterest.GetCommitsBehind() == 0 {
		forwardButton.IsDisabled = true
		forwardButton.Icon = string(gongtree_buttons.BUTTON_do_not_disturb)
		forwardButton.ToolTipText = "No more next commits"
	}

	if stageOfInterest.GetCommitsBehind() > 0 {
		discardButton := &tree.Button{
			Name:            "DiscardButton",
			Icon:            string(gongtree_buttons.BUTTON_cancel),
			HasToolTip:      true,
			ToolTipText:     "Discard commits ahead (git reset --hard HEAD)",
			ToolTipPosition: tree.Below,
		}
		deltaNode.Buttons = append(deltaNode.Buttons, discardButton)
		discardButton.Impl = &tree.FunctionalButtonProxy{
			OnUpdated: func(stage *tree.Stage,
				stagedButton, frontButton *tree.Button) {
				stageOfInterest.ResetHard()
				probe.Refresh()
			},
		}
	}
	appendChildrenNodeFunc(deltaNode)
}
