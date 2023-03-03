package node2gongdoc

import (
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type GongLinkImpl struct {
	gongLink *gong_models.GongLink
	NodeImpl
}

func (gongLinkImpl *GongLinkImpl) OnAfterUpdate(
	gongdocStage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	classdiagram := gongLinkImpl.nodeCb.GetSelectedClassdiagram()

	// find the parent node to find the gongstruct to find the gongstructshape
	// the node is field, one needs to find the gongstruct that contains it
	// get the parent node
	parentNode := gongLinkImpl.nodeCb.map_Children_Parent[stagedNode]

	gongNoteImpl := parentNode.Impl.(*GongNoteImpl)
	gongNote := gongNoteImpl.gongNote

	// find the classhape in the classdiagram
	foundNoteshape := false
	var noteshape *gongdoc_models.NoteShape
	_ = noteshape

	for _, _noteshape := range classdiagram.NoteShapes {
		// strange behavior when the gongstructshape is remove within the loop
		if gongdoc_models.IdentifierToGongObjectName(_noteshape.Identifier) ==
			gongNote.Name && !foundNoteshape {
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

}

func (gongLinkImpl *GongLinkImpl) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {
}
