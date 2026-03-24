package models

import (
	"log"
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) OnUpdate(_ *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {
}

type DiagramTree_Transition_Proxy struct {
	stager          *Stager
	transitionShape *Transition_Shape
	transition      *Transition
	diagram         *Diagram
}

// OnAfterUpdate implements models.NodeImplInterface.
func (proxy *DiagramTree_Transition_Proxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.Name != stagedNode.Name {
		proxy.transition.Name = frontNode.Name
		proxy.transition.IsInRenameMode = false
		proxy.stager.stage.Commit()
		return
	}

	// checking the node of the Transition means we want to add a Transition_Shape to the diagram
	if frontNode.IsChecked && !stagedNode.IsChecked {

		stagedNode.IsChecked = frontNode.IsChecked

		// add the Transition_Shape
		if proxy.transitionShape != nil {
			log.Fatalln("adding a shape to an already shapes transition")
		}

		transitionShape := new(Transition_Shape).Stage(proxy.stager.stage)
		transitionShape.Transition = proxy.transition
		transitionShape.Name = proxy.transition.Name + "-" + proxy.diagram.GetName()
		transitionShape.StartRatio = 0.5
		transitionShape.EndRatio = 0.5
		transitionShape.StartOrientation = ORIENTATION_VERTICAL
		transitionShape.EndOrientation = ORIENTATION_VERTICAL

		if transitionShape.Transition.Start == transitionShape.Transition.End {
			transitionShape.StartRatio = 0.1
			transitionShape.EndRatio = 0.9
		}

		transitionShape.CornerOffsetRatio = 1.3

		proxy.transitionShape = transitionShape

		proxy.diagram.Transition_Shapes = append(proxy.diagram.Transition_Shapes, transitionShape)

		proxy.stager.stage.Commit()
		return
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {

		stagedNode.IsChecked = frontNode.IsChecked

		// one need to remove the Transition_Shape
		if proxy.transitionShape == nil {
			log.Fatalln("remove a non existing shape to Transition")
		}

		transitionShape := proxy.transitionShape
		proxy.transitionShape = nil
		transitionShape.Unstage(proxy.stager.stage)

		idx := slices.Index(proxy.diagram.Transition_Shapes, transitionShape)
		proxy.diagram.Transition_Shapes = slices.Delete(proxy.diagram.Transition_Shapes, idx, idx+1)

		proxy.stager.stage.Commit()
		return
	}

	proxy.stager.probeForm.FillUpFormFromGongstruct(proxy.transition, "Transition")
}
