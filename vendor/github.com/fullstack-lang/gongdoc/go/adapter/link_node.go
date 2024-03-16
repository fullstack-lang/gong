package adapter

import (
	"log"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type LinkNode struct {
	ElementNodeBase
	gongNoteNode *GongNoteNode
	stage        *gong_models.StageStruct
	link         *gong_models.GongLink
}

func NewLinkNode(
	portfolioAdapter *PortfolioAdapter,
	gongNoteNode *GongNoteNode,
	link *gong_models.GongLink) (linkNode *LinkNode) {
	linkNode = &LinkNode{ElementNodeBase: ElementNodeBase{portfolioAdapter: portfolioAdapter}}

	linkNode.gongNoteNode = gongNoteNode
	linkNode.link = link
	return
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (linkNode *LinkNode) RemoveFromDiagram() {

	gongdocStage := linkNode.portfolioAdapter.gongdocStage

	// get the holding note shape
	var noteShape *gongdoc_models.NoteShape
	shape, ok := linkNode.portfolioAdapter.diagrammer.GetElementNodeDisplayed(linkNode.gongNoteNode)

	if !ok {
		log.Fatalln("unknown note shape")
		return
	}
	if noteShape, ok = shape.(*gongdoc_models.NoteShape); !ok {
		log.Fatalln("not a note shape")
		return
	}

	// get the relevant gong note link
	var noteShapeLink *gongdoc_models.NoteShapeLink
	for _, _noteShapeLink := range noteShape.NoteShapeLinks {
		if _noteShapeLink.Name == linkNode.link.Name {
			noteShapeLink = _noteShapeLink
		}
	}

	noteShapeLink.Unstage(gongdocStage)
	idx := slices.Index(noteShape.NoteShapeLinks, noteShapeLink)
	noteShape.NoteShapeLinks = slices.Delete(noteShape.NoteShapeLinks, idx, idx+1)

	gongdocStage.Commit()
}

// AddToDiagram implements diagrammer.ElementNode.
func (linkNode *LinkNode) AddToDiagram() {
	gongdocStage := linkNode.portfolioAdapter.gongdocStage

	var noteShape *gongdoc_models.NoteShape
	shape, ok := linkNode.portfolioAdapter.diagrammer.GetElementNodeDisplayed(linkNode.gongNoteNode)

	if !ok {
		log.Fatalln("unknown gong enum shape")
		return
	}
	if noteShape, ok = shape.(*gongdoc_models.NoteShape); !ok {
		log.Fatalln("not a gong enum shape")
		return
	}
	noteShapeLink := (&gongdoc_models.NoteShapeLink{Name: linkNode.GetName()}).Stage(gongdocStage)

	if linkNode.link.Recv != "" {

		noteShapeLink.Type = gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_FIELD
		noteShapeLink.Identifier =
			gongdoc_models.GongstructAndFieldnameToFieldIdentifier(
				linkNode.link.Recv, linkNode.link.Name)

	} else {
		noteShapeLink.Type = gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE
		noteShapeLink.Identifier =
			gongdoc_models.GongStructNameToIdentifier(linkNode.GetName())

	}

	noteShape.NoteShapeLinks = append(noteShape.NoteShapeLinks, noteShapeLink)

	gongdocStage.Commit()
}

// GetChildren implements diagrammer.ModelElementNode.
func (*LinkNode) GetChildren() []diagrammer.ModelNode {
	return nil
}

var _ diagrammer.ModelElementNode = &LinkNode{}

func (linkNode *LinkNode) IsExpanded() bool {
	return true
}

// GenerateProgeny implements diagrammer.Node.
func (linkNode *LinkNode) GenerateProgeny() (children []diagrammer.ModelNode) {
	return
}

// GetName implements diagrammer.Node.
func (linkNode *LinkNode) GetName() string {
	return linkNode.link.GetName()
}

// IsNameEditable implements diagrammer.Node.
func (linkNode *LinkNode) IsNameEditable() bool {
	return false
}

func (linkNode *LinkNode) GetElement() any {
	return linkNode.link
}

func (linkNode *LinkNode) CanBeAddedToDiagram() (result bool) {

	diagrammer := linkNode.portfolioAdapter.diagrammer
	gongStage := linkNode.portfolioAdapter.gongStage
	link := linkNode.link

	// the parent node must already be displayed in order to be able to display the node
	result = diagrammer.IsElementNodeDisplayed(linkNode.gongNoteNode)

	// first case is when the gong link points to a shape
	if result && link.Recv == "" {

		// is it a gongstruct ?
		map_ := *gong_models.GetGongstructInstancesMap[gong_models.GongStruct](gongStage)
		gongStruct, ok := map_[link.Name]

		if ok {
			result = result && diagrammer.IsElementDisplayed(gongStruct)
		}

		// is it a gongsenum ?
		map2_ := *gong_models.GetGongstructInstancesMap[gong_models.GongEnum](gongStage)
		gongEnum, ok := map2_[link.Name]

		if ok {
			result = result && diagrammer.IsElementDisplayed(gongEnum)
		}
	} else // the other case (Recv != "") is when the gonglink points to a link
	{
		fieldName := link.Recv + "." + link.Name
		_ = fieldName

		// if the receiver gongstruct present ?
		map_ := *gong_models.GetGongstructInstancesMap[gong_models.GongStruct](gongStage)
		gongStruct, ok := map_[link.Recv]

		if ok {
			result = result && diagrammer.IsElementDisplayed(gongStruct)
		}

		// if yes, is the field displayed ?
		if result && ok {
			for _, field := range gongStruct.Fields {
				if field.GetName() == link.Name {
					result = result && diagrammer.IsElementDisplayed(field)
				}
			}
		}
	}

	return
}
