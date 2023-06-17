package node2gongdoc

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type NodeImplGongnote struct {
	NodeImplGongObjectAbstract
	gongNote *gong_models.GongNote
}

func NewNodeImplGongnote(
	gongNote *gong_models.GongNote,
	nodeImplGongObjectAbstract NodeImplGongObjectAbstract,
) (nodeImplGongnote *NodeImplGongnote) {

	nodeImplGongnote = new(NodeImplGongnote)
	nodeImplGongnote.NodeImplGongObjectAbstract = nodeImplGongObjectAbstract
	nodeImplGongnote.gongNote = gongNote

	return
}

func (nodeImplGongnote *NodeImplGongnote) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	classDiagram := nodeImplGongnote.diagramPackage.SelectedClassdiagram
	gongdocStage := nodeImplGongnote.diagramPackage.Stage_

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {
		gongdocStage.Checkout()
		foundNote := false
		var noteShape *gongdoc_models.NoteShape
		var _noteShape *gongdoc_models.NoteShape
		for _, _noteShape = range classDiagram.NoteShapes {

			// strange behavior when the note is removed within the loop
			if gongdoc_models.IdentifierToGongObjectName(_noteShape.Identifier) == stagedNode.Name && !foundNote {
				foundNote = true
				noteShape = _noteShape
			}
		}
		if !foundNote {
			log.Panicf("Note %s of field not present ", stagedNode.Name)
		}

		// remove the links of the note shape
		for _, noteLink := range noteShape.NoteShapeLinks {
			noteLink.Unstage(gongdocStage)
			noteShape.NoteShapeLinks = remove(noteShape.NoteShapeLinks, noteLink)
		}
		classDiagram.NoteShapes = remove(classDiagram.NoteShapes, noteShape)
		noteShape.Unstage(gongdocStage)
	}

	// if node is checked, add gongnoteshape
	if !stagedNode.IsChecked && frontNode.IsChecked {
		gongdocStage.Checkout()

		noteShape := (&gongdoc_models.NoteShape{Name: stagedNode.Name}).Stage(gongdocStage)

		mapOfGongNotes := *gong_models.GetGongstructInstancesMap[gong_models.GongNote](nodeImplGongnote.diagramPackage.ModelPkg.GetStage())

		gongNote, ok := mapOfGongNotes[noteShape.Name]
		if !ok {
			log.Fatal("Unkown note ", noteShape.Name)
		}

		noteShape.Identifier = gongdoc_models.GongStructNameToIdentifier(noteShape.Name)

		noteShape.Body = gongNote.Body
		noteShape.X = 30
		noteShape.Y = 30
		noteShape.Width = 240
		noteShape.Heigth = 63

		classDiagram.NoteShapes = append(classDiagram.NoteShapes, noteShape)
	}

	computeGongNodesConfigurations(
		gongdocStage,
		classDiagram,
		nodeImplGongnote.treeOfGongObjects)

	gongdocStage.Commit()
}
