package adapter

import (
	"log"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type GongNoteNode struct {
	ElementNodeBase
	gongNote *gong_models.GongNote
}

var _ diagrammer.ModelElementNode = &GongNoteNode{}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (gongNoteNode *GongNoteNode) RemoveFromDiagram() {
	gongdocStage := gongNoteNode.portfolioAdapter.gongdocStage
	diagramPackage := gongNoteNode.portfolioAdapter.getDiagramPackage()
	classDiagram := diagramPackage.SelectedClassdiagram

	var noteShape *gongdoc_models.NoteShape
	shape, ok := gongNoteNode.portfolioAdapter.diagrammer.GetElementNodeDisplayed(gongNoteNode)

	if !ok {
		log.Fatalln("unknown note shape")
		return
	}
	if noteShape, ok = shape.(*gongdoc_models.NoteShape); !ok {
		log.Fatalln("not a note shape")
		return
	}

	// remove the links of the note shape
	for _, noteLink := range noteShape.NoteShapeLinks {
		noteLink.Unstage(gongdocStage)
		idx := slices.Index(noteShape.NoteShapeLinks, noteLink)

		noteShape.NoteShapeLinks = slices.Delete(noteShape.NoteShapeLinks, idx, idx+1)
	}
	idx := slices.Index(classDiagram.NoteShapes, noteShape)

	classDiagram.NoteShapes = slices.Delete(classDiagram.NoteShapes, idx, idx+1)
	noteShape.Unstage(gongdocStage)

	gongdocStage.Commit()
}

// AddToDiagram implements diagrammer.ElementNode.
func (gongNoteNode *GongNoteNode) AddToDiagram() {
	diagramPackage := gongNoteNode.portfolioAdapter.getDiagramPackage()
	gongdocStage := gongNoteNode.portfolioAdapter.gongdocStage

	noteShape := (&gongdoc_models.NoteShape{Name: gongNoteNode.GetName()}).Stage(gongdocStage)

	noteShape.Identifier = gongdoc_models.GongStructNameToIdentifier(noteShape.Name)

	noteShape.Body = gongNoteNode.gongNote.Body
	noteShape.X = 30
	noteShape.Y = 30
	noteShape.Width = 240
	noteShape.Height = 63

	diagramPackage.SelectedClassdiagram.NoteShapes =
		append(diagramPackage.SelectedClassdiagram.NoteShapes, noteShape)
	gongdocStage.Commit()
}

var _ diagrammer.ModelElementNode = &GongNoteNode{}

func NewGongNoteNode(
	portfolioAdapter *PortfolioAdapter,
	gongNote *gong_models.GongNote) (gongNoteNode *GongNoteNode) {

	gongNoteNode = &GongNoteNode{ElementNodeBase: ElementNodeBase{portfolioAdapter: portfolioAdapter}}

	gongNoteNode.gongNote = gongNote
	return
}

// GenerateProgeny implements diagrammer.Node.
func (gongNoteNode *GongNoteNode) GenerateProgeny() []diagrammer.ModelNode {

	for _, link := range gongNoteNode.gongNote.Links {
		linkNode := NewLinkNode(gongNoteNode.portfolioAdapter, gongNoteNode, link)
		gongNoteNode.children = append(gongNoteNode.children, linkNode)
	}

	return gongNoteNode.children
}

// GetName implements diagrammer.Node.
func (gongNoteNode *GongNoteNode) GetName() string {
	return gongNoteNode.gongNote.GetName()
}

func (gongNoteNode *GongNoteNode) GetElement() any {
	return gongNoteNode.gongNote
}
