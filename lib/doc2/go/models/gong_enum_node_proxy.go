package models

import (
	"log" // Added for logging errors

	gong "github.com/fullstack-lang/gong/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	// Updated import path for the nodestates package
	"github.com/fullstack-lang/gong/lib/doc2/go/models/nodestates"
)

type GongEnumNodeProxy struct {
	node          *tree.Node
	stager        *Stager
	classDiagram  *Classdiagram
	gongEnum      *gong.GongEnum
	gongEnumShape *GongEnumShape
	rank          int // This corresponds to the 'index' in nodestates functions
}

func (proxy *GongEnumNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		proxy.classDiagram.AddGongEnumShape(proxy.stager.stage, diagramPackage, proxy.gongEnum.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitFormStage()
		proxy.stager.UpdateAndCommitSVGStage()
		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		proxy.classDiagram.RemoveGongEnumShape(proxy.stager.stage, proxy.gongEnum.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitFormStage()
		proxy.stager.UpdateAndCommitSVGStage()
		proxy.stager.stage.Commit()
	}

	expansionToggled := false
	currentExpansionStateInDiagram := proxy.classDiagram.NodeGongEnumNodeExpansion

	if front.IsExpanded && !staged.IsExpanded {
		err := nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, proxy.rank)
		if err != nil {
			log.Printf("Error toggling node expansion for GongEnum %s (rank %d) to expanded: %v", proxy.gongEnum.Name, proxy.rank, err)
		} else {
			proxy.classDiagram.NodeGongEnumNodeExpansion = currentExpansionStateInDiagram
			expansionToggled = true
		}
	}

	if !front.IsExpanded && staged.IsExpanded {
		err := nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, proxy.rank)
		if err != nil {
			log.Printf("Error toggling node expansion for GongEnum %s (rank %d) to collapsed: %v", proxy.gongEnum.Name, proxy.rank, err)
		} else {
			proxy.classDiagram.NodeGongEnumNodeExpansion = currentExpansionStateInDiagram
			expansionToggled = true
		}
	}

	if expansionToggled {
		proxy.stager.stage.Commit()
		// proxy.stager.UpdateAndCommitTreeStage() // Consider if this is needed or handled elsewhere
	}
}
