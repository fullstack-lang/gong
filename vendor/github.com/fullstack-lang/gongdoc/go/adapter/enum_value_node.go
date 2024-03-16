package adapter

import (
	"log"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type EnumValueNode struct {
	ElementNodeBase
	gongEnumNode *GongEnumNode
	enumValue    *gong_models.GongEnumValue
}

var _ diagrammer.ModelElementNode = &EnumValueNode{}

func NewEnumValueNode(
	portfolioAdapter *PortfolioAdapter,
	gongEnumNode *GongEnumNode,
	enumvalue *gong_models.GongEnumValue) (enumvalueNode *EnumValueNode) {
	enumvalueNode = &EnumValueNode{ElementNodeBase: ElementNodeBase{portfolioAdapter: portfolioAdapter}}

	enumvalueNode.gongEnumNode = gongEnumNode
	enumvalueNode.enumValue = enumvalue
	return
}

// GetChildren implements diagrammer.ModelElementNode.
func (*EnumValueNode) GetChildren() []diagrammer.ModelNode {
	return nil
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (enumValueNode *EnumValueNode) RemoveFromDiagram() {
	gongdocStage := enumValueNode.portfolioAdapter.gongdocStage

	var gongEnumShape *gongdoc_models.GongEnumShape
	shape, ok := enumValueNode.portfolioAdapter.diagrammer.GetElementNodeDisplayed(enumValueNode.gongEnumNode)

	if !ok {
		log.Fatalln("unknown gong enum shape")
		return
	}
	if gongEnumShape, ok = shape.(*gongdoc_models.GongEnumShape); !ok {
		log.Fatalln("not a gong enum shape")
		return
	}

	var gongEnumValueEntry *gongdoc_models.GongEnumValueEntry

	for _, _field := range gongEnumShape.GongEnumValueEntrys {
		if gongdoc_models.IdentifierToFieldName(_field.Identifier) == enumValueNode.GetName() {
			gongEnumValueEntry = _field
		}
	}
	if gongEnumValueEntry != nil {
		idx := slices.Index(gongEnumShape.GongEnumValueEntrys, gongEnumValueEntry)
		gongEnumShape.GongEnumValueEntrys = slices.Delete(gongEnumShape.GongEnumValueEntrys, idx, idx+1)
		gongEnumShape.Height = gongEnumShape.Height - 15
		gongEnumValueEntry.Unstage(gongdocStage)
	}
	gongdocStage.Commit()
}

// AddToDiagram implements diagrammer.ElementNode.
func (enumValueNode *EnumValueNode) AddToDiagram() {

	gongdocStage := enumValueNode.portfolioAdapter.gongdocStage

	var gongEnumShape *gongdoc_models.GongEnumShape
	shape, ok := enumValueNode.portfolioAdapter.diagrammer.GetElementNodeDisplayed(enumValueNode.gongEnumNode)

	if !ok {
		log.Fatalln("unknown gong enum shape")
		return
	}
	if gongEnumShape, ok = shape.(*gongdoc_models.GongEnumShape); !ok {
		log.Fatalln("not a gong enum shape")
		return
	}

	// insert the value at the correct spot in the classhape
	map_Value_rankInEnum := make(map[gong_models.GongEnumValue]int, 0)
	map_ValueName_Value := make(map[string]gong_models.GongEnumValue, 0)

	// what is the index of the field to insert in the gong struct ?
	rankkInEnum := 0

	gongEnumShape.Height = gongEnumShape.Height + 15

	var gongEnumValueEntry gongdoc_models.GongEnumValueEntry
	gongEnumValueEntry.Name = enumValueNode.GetName()
	gongEnumValueEntry.Identifier = gongdoc_models.GongstructAndFieldnameToFieldIdentifier(
		enumValueNode.gongEnumNode.gongEnum.Name, enumValueNode.GetName())

	for idx, gongEnum := range enumValueNode.gongEnumNode.gongEnum.GongEnumValues {

		map_Value_rankInEnum[*gongEnum] = idx
		map_ValueName_Value[gongEnum.GetName()] = *gongEnum

		if gongEnum.GetName() == gongEnumValueEntry.Name {
			rankkInEnum = idx
		}
	}

	// compute insertionIndex (index where to insert the field to display)
	insertionIndex := 0
	for idx, field := range gongEnumShape.GongEnumValueEntrys {
		value := map_ValueName_Value[gongdoc_models.IdentifierToFieldName(field.Identifier)]
		_rankInEnum := map_Value_rankInEnum[value]
		if rankkInEnum > _rankInEnum {
			insertionIndex = idx + 1
		}
	}

	// append the filed to display in the right index
	if insertionIndex == len(gongEnumShape.GongEnumValueEntrys) {
		gongEnumShape.GongEnumValueEntrys = append(gongEnumShape.GongEnumValueEntrys, &gongEnumValueEntry)
	} else {
		gongEnumShape.GongEnumValueEntrys = append(gongEnumShape.GongEnumValueEntrys[:insertionIndex+1],
			gongEnumShape.GongEnumValueEntrys[insertionIndex:]...)
		gongEnumShape.GongEnumValueEntrys[insertionIndex] = &gongEnumValueEntry
	}
	gongEnumValueEntry.Stage(gongdocStage)
	gongdocStage.Commit()
}

// GenerateProgeny implements diagrammer.Node.
func (enumvalueNode *EnumValueNode) GenerateProgeny() (children []diagrammer.ModelNode) {
	return
}

// GetName implements diagrammer.Node.
func (enumvalueNode *EnumValueNode) GetName() string {
	return enumvalueNode.enumValue.GetName()
}

// GetElement implements diagrammer.ModelNode.
func (enumvalueNode *EnumValueNode) GetElement() any {
	return enumvalueNode.enumValue
}

func (enumvalueNode *EnumValueNode) CanBeAddedToDiagram() (result bool) {

	diagrammer := enumvalueNode.portfolioAdapter.diagrammer

	// the parent node must already be displayed in order to be able to display the node
	result = diagrammer.IsElementNodeDisplayed(enumvalueNode.gongEnumNode)

	return
}
