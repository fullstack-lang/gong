package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type GongLinkImpl struct {
	gongLink *gong_models.GongLink
	NodeImpl
}

func (gongLinkImpl *GongLinkImpl) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
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
		stage.Checkout()

		noteShapeLink := (&gongdoc_models.NoteShapeLink{Name: stagedNode.GetName()}).Stage()
		noteShapeLink.Identifier =
			gongdoc_models.GongstructAndFieldnameToFieldIdentifier(gongNote.Name, stagedNode.Name)

		noteshape.NoteShapeLinks = append(noteshape.NoteShapeLinks, noteShapeLink)

	}

	// removing a note link
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stage.Checkout()

		// get the relevant gong note link
		var noteLink *gongdoc_models.NoteShapeLink
		for _, _noteLink := range noteshape.NoteShapeLinks {
			if gongdoc_models.IdentifierToFieldName(_noteLink.Identifier) ==
				stagedNode.Name {
				noteLink = _noteLink
			}
		}

		noteLink.Unstage()
		noteshape.NoteShapeLinks = remove(noteshape.NoteShapeLinks, noteLink)

	}

}

func (gongLinkImpl *GongLinkImpl) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {
}
