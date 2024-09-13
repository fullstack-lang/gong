package adapter

import (
	"log"
	"slices"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type FieldNode struct {
	ElementNodeBase
	gongStructNode *GongStructNode // parent node
	Field          gong_models.FieldInterface
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (fieldNode *FieldNode) RemoveFromDiagram() {
	gongdocStage := fieldNode.portfolioAdapter.gongdocStage

	var gongStructShape *gongdoc_models.GongStructShape
	shape, ok := fieldNode.portfolioAdapter.diagrammer.GetElementNodeDisplayed(fieldNode.gongStructNode)

	if !ok {
		log.Fatalln("unknown gongstruct shape")
		return
	}
	if gongStructShape, ok = shape.(*gongdoc_models.GongStructShape); !ok {
		log.Fatalln("not a gongstruct shape")
		return
	}

	{
		var field *gongdoc_models.Field

		for _, _field := range gongStructShape.Fields {
			if gongdoc_models.IdentifierToFieldName(_field.Identifier) == fieldNode.GetName() {
				field = _field
			}
		}
		if field != nil {
			idx := slices.Index(gongStructShape.Fields, field)
			gongStructShape.Fields = slices.Delete(gongStructShape.Fields, idx, idx+1)
			gongStructShape.Height = gongStructShape.Height - 15
			field.Unstage(gongdocStage)
		}
	}
	{
		var link *gongdoc_models.Link

		for _, _field := range gongStructShape.Links {
			if gongdoc_models.IdentifierToFieldName(_field.Identifier) == fieldNode.GetName() {
				link = _field
			}
		}
		if link != nil {
			idx := slices.Index(gongStructShape.Links, link)
			gongStructShape.Links = slices.Delete(gongStructShape.Links, idx, idx+1)
			link.Unstage(gongdocStage)
		}
	}

	gongdocStage.Commit()
}

// AddToDiagram implements diagrammer.ElementNode.
func (fieldNode *FieldNode) AddToDiagram() {
	diagramPackage := fieldNode.portfolioAdapter.getDiagramPackage()
	gongdocStage := fieldNode.portfolioAdapter.gongdocStage

	// get the gongstruct shape
	// diagramNode := fieldNode.portfolioAdapter.GetSelectedPortfolioDiagramNode()

	var gongStructShape *gongdoc_models.GongStructShape
	shape, ok := fieldNode.portfolioAdapter.diagrammer.GetElementNodeDisplayed(fieldNode.gongStructNode)

	if !ok {
		log.Fatalln("unknown gongstruct shape")
		return
	}
	if gongStructShape, ok = shape.(*gongdoc_models.GongStructShape); !ok {
		log.Fatalln("not a gongstruct shape")
		return
	}

	switch fieldNode.Field.(type) {
	case *gong_models.GongBasicField, *gong_models.GongTimeField:

		// concrete in the sense of UML concrete syntax
		var concreteField gongdoc_models.Field
		concreteField.Name = fieldNode.GetName()
		concreteField.Identifier = gongdoc_models.GongstructAndFieldnameToFieldIdentifier(
			fieldNode.gongStructNode.gongStruct.Name, fieldNode.GetName())

		switch realField := fieldNode.Field.(type) {
		case *gong_models.GongBasicField:

			// get the type after the "."
			names := strings.Split(realField.DeclaredType, ".")
			fieldTypeName := names[len(names)-1]

			concreteField.Fieldtypename = fieldTypeName
		case *gong_models.GongTimeField:
			concreteField.Fieldtypename = "Time"
		case *gong_models.PointerToGongStructField:
		case *gong_models.SliceOfPointerToGongStructField:
		}

		concreteField.Structname = gongdoc_models.IdentifierToGongObjectName(gongStructShape.Identifier)
		concreteField.Stage(gongdocStage)

		gongStructShape.Height = gongStructShape.Height + 15

		// construct ordered slice of fields
		map_Field_Rank := make(map[gong_models.FieldInterface]int, 0)
		map_Name_Field := make(map[string]gong_models.FieldInterface, 0)

		// what is the index of the field to insert in the gong struct ?
		fieldRank := 0

		// let's compute it by parsing the field of the gongstruct
		gongStruct_ := (*gong_models.GetGongstructInstancesMap[gong_models.GongStruct](
			diagramPackage.ModelPkg.GetStage()))[fieldNode.gongStructNode.gongStruct.Name]
		for idx, gongField := range gongStruct_.Fields {

			map_Field_Rank[gongField] = idx
			map_Name_Field[gongField.GetName()] = gongField

			if gongField.GetName() == concreteField.Name {
				fieldRank = idx
			}
		}

		// compute insertionIndex (index where to insert the field to display)
		insertionIndex := 0
		for idx, field := range gongStructShape.Fields {
			gongField := map_Name_Field[gongdoc_models.IdentifierToFieldName(field.Identifier)]
			_fieldRank := map_Field_Rank[gongField]
			if fieldRank > _fieldRank {
				insertionIndex = idx + 1
			}
		}

		// append the filed to display in the right index
		if insertionIndex == len(gongStructShape.Fields) {
			gongStructShape.Fields = append(gongStructShape.Fields, &concreteField)
		} else {
			gongStructShape.Fields = append(gongStructShape.Fields[:insertionIndex+1],
				gongStructShape.Fields[insertionIndex:]...)
			gongStructShape.Fields[insertionIndex] = &concreteField
		}

	case *gong_models.PointerToGongStructField, *gong_models.SliceOfPointerToGongStructField:

		var targetStructName string
		var sourceMultiplicity gongdoc_models.MultiplicityType
		var targetMultiplicity gongdoc_models.MultiplicityType

		switch realField := fieldNode.Field.(type) {
		case *gong_models.PointerToGongStructField:
			targetStructName = realField.GongStruct.Name
			sourceMultiplicity = gongdoc_models.MANY
			targetMultiplicity = gongdoc_models.ZERO_ONE
		case *gong_models.SliceOfPointerToGongStructField:
			targetStructName = realField.GongStruct.Name
			sourceMultiplicity = gongdoc_models.MANY
			targetMultiplicity = gongdoc_models.MANY
		}
		targetSourceGongStructShape := false
		var targetGongStructShape *gongdoc_models.GongStructShape
		for _, _gongstructshape := range diagramPackage.SelectedClassdiagram.GongStructShapes {

			// strange behavior when the gongstructshape is remove within the loop
			if gongdoc_models.IdentifierToGongObjectName(_gongstructshape.Identifier) == targetStructName && !targetSourceGongStructShape {
				targetSourceGongStructShape = true
				targetGongStructShape = _gongstructshape
			}
		}
		if !targetSourceGongStructShape {
			log.Panicf("GongStructShape %s of field not present ", targetStructName)
		}
		_ = targetGongStructShape

		link := new(gongdoc_models.Link).Stage(gongdocStage)
		link.Name = fieldNode.GetName()
		link.SourceMultiplicity = sourceMultiplicity
		link.SourceMultiplicityOffsetX = 10
		link.SourceMultiplicityOffsetY = -50

		link.TargetMultiplicity = targetMultiplicity
		link.TargetMultiplicityOffsetX = -50
		link.TargetMultiplicityOffsetY = 16

		link.FieldOffsetX = -50
		link.FieldOffsetY = -16

		link.Identifier =
			gongdoc_models.GongstructAndFieldnameToFieldIdentifier(fieldNode.gongStructNode.gongStruct.Name, fieldNode.GetName())
		link.Fieldtypename = gongdoc_models.GongStructNameToIdentifier(targetStructName)

		gongStructShape.Links = append(gongStructShape.Links, link)
		link.Middlevertice = new(gongdoc_models.Vertice).Stage(gongdocStage)
		link.Middlevertice.Name = "Verticle in class diagram " + diagramPackage.SelectedClassdiagram.Name +
			" in middle between " + gongStructShape.Name + " and " + targetGongStructShape.Name
		link.Middlevertice.X = (gongStructShape.Position.X+targetGongStructShape.Position.X)/2.0 +
			gongStructShape.Width*1.5
		link.Middlevertice.Y = (gongStructShape.Position.Y+targetGongStructShape.Position.Y)/2.0 +
			gongStructShape.Height/2.0

		link.StartOrientation = gongdoc_models.ORIENTATION_HORIZONTAL
		link.StartRatio = 0.5
		link.EndOrientation = gongdoc_models.ORIENTATION_HORIZONTAL
		link.EndRatio = 0.5
		link.CornerOffsetRatio = 1.38
	}
	gongdocStage.Commit()
}

// GetChildren implements diagrammer.ModelElementNode.
func (*FieldNode) GetChildren() []diagrammer.ModelNode {
	return nil
}

var _ diagrammer.ModelElementNode = &FieldNode{}

func NewFieldNode(
	portfolioAdapter *PortfolioAdapter,
	gongStructNode *GongStructNode,
	field gong_models.FieldInterface) (fieldNode *FieldNode) {
	fieldNode = &FieldNode{ElementNodeBase: ElementNodeBase{portfolioAdapter: portfolioAdapter}}

	fieldNode.gongStructNode = gongStructNode
	fieldNode.Field = field
	return
}

// GenerateProgeny implements diagrammer.Node.
func (fieldNode *FieldNode) GenerateProgeny() (children []diagrammer.ModelNode) {
	return
}

// GetName implements diagrammer.Node.
func (fieldNode *FieldNode) GetName() string {
	return fieldNode.Field.GetName()
}

func (fieldNode *FieldNode) GetElement() any {
	return fieldNode.Field
}

func (fieldNode *FieldNode) CanBeAddedToDiagram() (result bool) {

	diagrammer := fieldNode.portfolioAdapter.diagrammer
	// the parent node must already be displayed in order to be able to display the node
	result = diagrammer.IsElementNodeDisplayed(fieldNode.gongStructNode)

	if pointerToGongStructField, ok := fieldNode.Field.(*gong_models.PointerToGongStructField); ok {
		result = result && diagrammer.IsElementDisplayed(pointerToGongStructField.GongStruct)
	}

	if sliceOfPointerToGongStructField, ok := fieldNode.Field.(*gong_models.SliceOfPointerToGongStructField); ok {
		result = result && diagrammer.IsElementDisplayed(sliceOfPointerToGongStructField.GongStruct)
	}

	return
}

func (fieldNode *FieldNode) HasSecondCheckbox() (result bool) {

	// if pointerToGongStructField, ok := fieldNode.Field.(*gong_models.PointerToGongStructField); ok {
	// 	_ = pointerToGongStructField
	// 	result = !fieldNode.CanBeAddedToDiagram()
	// }

	// if sliceOfPointerToGongStructField, ok := fieldNode.Field.(*gong_models.SliceOfPointerToGongStructField); ok {
	// 	_ = sliceOfPointerToGongStructField
	// 	result = !fieldNode.CanBeAddedToDiagram()
	// }
	return
}

func (fieldNode *FieldNode) HasSecondName() (result bool) {

	if pointerToGongStructField, ok := fieldNode.Field.(*gong_models.PointerToGongStructField); ok {
		_ = pointerToGongStructField
		result = true
	}

	if sliceOfPointerToGongStructField, ok := fieldNode.Field.(*gong_models.SliceOfPointerToGongStructField); ok {
		_ = sliceOfPointerToGongStructField
		result = true
	}
	return
}

func (fieldNode *FieldNode) GetSecondName() (result string) {

	if pointerToGongStructField, ok := fieldNode.Field.(*gong_models.PointerToGongStructField); ok {
		_ = pointerToGongStructField
		result = pointerToGongStructField.GongStruct.GetName()
	}

	if sliceOfPointerToGongStructField, ok := fieldNode.Field.(*gong_models.SliceOfPointerToGongStructField); ok {
		_ = sliceOfPointerToGongStructField
		result = sliceOfPointerToGongStructField.GongStruct.GetName()
	}

	return
}
