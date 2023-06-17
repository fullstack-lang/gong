package node2gongdoc

import (
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type NodeImplLink struct {
	NodeImplGongObjectAbstract

	gongNote *gong_models.GongNote
	gongLink *gong_models.GongLink

	nodeOfGongNote *gongtree_models.Node
	nodeOfLink     *gongtree_models.Node
}

func NewNodeImplLink(
	gongNote *gong_models.GongNote,
	gongLink *gong_models.GongLink,
	nodeOfGongNote *gongtree_models.Node,
	nodeOfLink *gongtree_models.Node,
	NodeImplGongObjectAbstract NodeImplGongObjectAbstract,
) (nodeImplLink *NodeImplLink) {

	nodeImplLink = new(NodeImplLink)
	nodeImplLink.gongNote = gongNote
	nodeImplLink.gongLink = gongLink
	nodeImplLink.nodeOfGongNote = nodeOfGongNote
	nodeImplLink.nodeOfLink = nodeOfLink

	nodeImplLink.NodeImplGongObjectAbstract = NodeImplGongObjectAbstract

	return
}

func (nodeImplLink *NodeImplLink) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	classdiagram := nodeImplLink.diagramPackage.SelectedClassdiagram
	gongdocStage := nodeImplLink.diagramPackage.Stage_

	// find the classhape in the classdiagram
	foundNoteshape := false
	var noteshape *gongdoc_models.NoteShape
	_ = noteshape

	for _, _noteshape := range classdiagram.NoteShapes {
		// strange behavior when the gongstructshape is remove within the loop
		if gongdoc_models.IdentifierToGongObjectName(_noteshape.Identifier) ==
			nodeImplLink.gongNote.Name && !foundNoteshape {
			noteshape = _noteshape
		}
	}

	// adding a note link
	if !stagedNode.IsChecked && frontNode.IsChecked {
		gongdocStage.Checkout()

		noteShapeLink := (&gongdoc_models.NoteShapeLink{Name: stagedNode.GetName()}).Stage(gongdocStage)

		if strings.ContainsAny(stagedNode.Name, ".") {

			subStrings := strings.Split(stagedNode.Name, ".")

			noteShapeLink.Type = gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_FIELD
			noteShapeLink.Identifier =
				gongdoc_models.GongstructAndFieldnameToFieldIdentifier(subStrings[0], subStrings[1])

		} else {
			noteShapeLink.Type = gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE
			noteShapeLink.Identifier =
				gongdoc_models.GongStructNameToIdentifier(stagedNode.Name)

		}

		noteshape.NoteShapeLinks = append(noteshape.NoteShapeLinks, noteShapeLink)
	}

	// removing a note link
	if stagedNode.IsChecked && !frontNode.IsChecked {
		gongdocStage.Checkout()

		// get the relevant gong note link
		var noteShapeLink *gongdoc_models.NoteShapeLink
		for _, _noteShapeLink := range noteshape.NoteShapeLinks {

			switch _noteShapeLink.Type {
			case gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE:
				if gongdoc_models.IdentifierToGongObjectName(_noteShapeLink.Identifier) == stagedNode.Name {
					noteShapeLink = _noteShapeLink
				}
			case gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_FIELD:
				receiver, fieldName := gongdoc_models.IdentifierToReceiverAndFieldName(_noteShapeLink.Identifier)

				if receiver+"."+fieldName == stagedNode.Name {
					noteShapeLink = _noteShapeLink
				}
			}
		}

		noteShapeLink.Unstage(gongdocStage)
		noteshape.NoteShapeLinks = remove(noteshape.NoteShapeLinks, noteShapeLink)
	}

	computeGongNodesConfigurations(
		gongdocStage,
		nodeImplLink.diagramPackage.SelectedClassdiagram,
		nodeImplLink.treeOfGongObjects)
	gongdocStage.Commit()
}
