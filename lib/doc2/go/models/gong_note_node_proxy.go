package models

import (
	"log" // Added for logging errors

	gong "github.com/fullstack-lang/gong/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	// Updated import path for the nodestates package
	"github.com/fullstack-lang/gong/lib/doc2/go/models/nodestates"
)

type GongNoteNodeProxy struct {
	node          *tree.Node
	stager        *Stager
	classDiagram  *Classdiagram
	gongNote      *gong.GongNote
	gongNoteShape *GongNoteShape
	rank          int // This corresponds to the 'index' in nodestates functions
}

func (proxy *GongNoteNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		proxy.classDiagram.AddGongNoteShape(proxy.stager.stage, proxy.gongNote, diagramPackage, proxy.gongNote.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitFormStage()
		proxy.stager.UpdateAndCommitSVGStage()
		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		proxy.classDiagram.RemoveGongNoteShape(proxy.stager.stage, proxy.gongNote.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitFormStage()
		proxy.stager.UpdateAndCommitSVGStage()
		proxy.stager.stage.Commit()
	}

	// expansionToggled tracks if the user action was an expansion or collapse
	expansionToggled := false
	currentExpansionStateInDiagram := proxy.classDiagram.NodeGongNoteNodeExpansion // For readability

	// User expanded the node
	if front.IsExpanded && !staged.IsExpanded {
		err := nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, proxy.rank)
		if err != nil {
			log.Printf("Error toggling node expansion for GongNote %s (rank %d) to expanded: %v", proxy.gongNote.Name, proxy.rank, err)
		} else {
			proxy.classDiagram.NodeGongNoteNodeExpansion = currentExpansionStateInDiagram
			expansionToggled = true
		}
		// front.IsExpanded is already true, reflecting user's action.
		// The backend model (classDiagram.NodeGongNoteNodeExpansion) is now updated.
	}

	// User collapsed the node
	if !front.IsExpanded && staged.IsExpanded {
		err := nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, proxy.rank)
		if err != nil {
			log.Printf("Error toggling node expansion for GongNote %s (rank %d) to collapsed: %v", proxy.gongNote.Name, proxy.rank, err)
		} else {
			proxy.classDiagram.NodeGongNoteNodeExpansion = currentExpansionStateInDiagram
			expansionToggled = true
		}
		// front.IsExpanded is already false, reflecting user's action.
	}

	if expansionToggled {
		proxy.stager.stage.Commit() // Commit changes to the classDiagram model
		// Optionally, if immediate tree refresh based on the new string state is needed:
		// proxy.stager.UpdateAndCommitTreeStage()
		// However, this might be redundant if other actions already trigger it.
		// The critical part is that classDiagram.NodeGongNoteNodeExpansion is updated.
	}
}
