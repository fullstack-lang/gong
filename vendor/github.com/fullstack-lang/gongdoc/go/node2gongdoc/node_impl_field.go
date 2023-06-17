package node2gongdoc

import (
	"log"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

// NodeImplField handles the passage fo information
// from the field node in the tree of identifiers to
// other nodes in the tree of identifiers and to the
// the diagram
type NodeImplField struct {
	NodeImplGongObjectAbstract

	gongStruct *gong_models.GongStruct
	field      gong_models.FieldInterface

	nodeOfGongstruct *gongtree_models.Node
	nodeOfField      *gongtree_models.Node
}

func NewNodeImplField(
	gongStruct *gong_models.GongStruct,
	field gong_models.FieldInterface,
	nodeOfGongstruct *gongtree_models.Node,
	nodeOfField *gongtree_models.Node,
	NodeImplGongObjectAbstract NodeImplGongObjectAbstract,
) (nodeImplField *NodeImplField) {

	nodeImplField = new(NodeImplField)
	nodeImplField.gongStruct = gongStruct
	nodeImplField.field = field
	nodeImplField.nodeOfGongstruct = nodeOfGongstruct
	nodeImplField.nodeOfField = nodeOfField

	nodeImplField.NodeImplGongObjectAbstract = NodeImplGongObjectAbstract

	return
}

func (nodeImplField *NodeImplField) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	gongdocStage := nodeImplField.diagramPackage.Stage_

	// find the classhape in the classdiagram
	foundGongStructShape := false
	var gongStructShape *gongdoc_models.GongStructShape
	for _, _gongstructshape := range nodeImplField.diagramPackage.SelectedClassdiagram.GongStructShapes {
		// strange behavior when the gongstructshape is remove within the loop
		if gongdoc_models.IdentifierToGongObjectName(_gongstructshape.Identifier) ==
			nodeImplField.gongStruct.Name && !foundGongStructShape {
			gongStructShape = _gongstructshape
		}
	}

	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		gongdocStage.Checkout()

		{
			var field *gongdoc_models.Field

			for _, _field := range gongStructShape.Fields {
				if gongdoc_models.IdentifierToFieldName(_field.Identifier) == stagedNode.Name {
					field = _field
				}
			}
			if field != nil {
				gongStructShape.Fields = remove(gongStructShape.Fields, field)
				gongStructShape.Heigth = gongStructShape.Heigth - 15
				field.Unstage(gongdocStage)
			}
		}
		{
			var link *gongdoc_models.Link

			for _, _field := range gongStructShape.Links {
				if gongdoc_models.IdentifierToFieldName(_field.Identifier) == stagedNode.Name {
					link = _field
				}
			}
			if link != nil {
				gongStructShape.Links = remove(gongStructShape.Links, link)
				link.Unstage(gongdocStage)
			}
		}

		gongdocStage.Commit()
	}
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		gongdocStage.Checkout()

		switch nodeImplField.field.(type) {
		case *gong_models.GongBasicField, *gong_models.GongTimeField:

			// concrete in the sense of UML concrete syntax
			var concreteField gongdoc_models.Field
			concreteField.Name = stagedNode.Name
			concreteField.Identifier = gongdoc_models.GongstructAndFieldnameToFieldIdentifier(
				nodeImplField.gongStruct.Name, stagedNode.Name)

			switch realField := nodeImplField.field.(type) {
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

			gongStructShape.Heigth = gongStructShape.Heigth + 15

			// construct ordered slice of fields
			map_Field_Rank := make(map[gong_models.FieldInterface]int, 0)
			map_Name_Field := make(map[string]gong_models.FieldInterface, 0)

			// what is the index of the field to insert in the gong struct ?
			fieldRank := 0

			// let's compute it by parsing the field of the gongstruct
			gongStruct_ := (*gong_models.GetGongstructInstancesMap[gong_models.GongStruct](
				nodeImplField.diagramPackage.ModelPkg.GetStage()))[nodeImplField.gongStruct.Name]
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

			switch realField := nodeImplField.field.(type) {
			case *gong_models.PointerToGongStructField:
				targetStructName = realField.GongStruct.Name
				sourceMultiplicity = gongdoc_models.MANY
				targetMultiplicity = gongdoc_models.ZERO_ONE
			case *gong_models.SliceOfPointerToGongStructField:
				targetStructName = realField.GongStruct.Name
				sourceMultiplicity = gongdoc_models.ZERO_ONE
				targetMultiplicity = gongdoc_models.MANY
			}
			targetSourceGongStructShape := false
			var targetGongStructShape *gongdoc_models.GongStructShape
			for _, _gongstructshape := range nodeImplField.diagramPackage.SelectedClassdiagram.GongStructShapes {

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
			link.Name = stagedNode.Name
			link.SourceMultiplicity = sourceMultiplicity
			link.SourceMultiplicityOffsetX = 10
			link.SourceMultiplicityOffsetY = -50

			link.TargetMultiplicity = targetMultiplicity
			link.TargetMultiplicityOffsetX = -50
			link.TargetMultiplicityOffsetY = 16

			link.FieldOffsetX = -50
			link.FieldOffsetY = -16

			link.Identifier =
				gongdoc_models.GongstructAndFieldnameToFieldIdentifier(nodeImplField.gongStruct.Name, stagedNode.Name)
			link.Fieldtypename = gongdoc_models.GongStructNameToIdentifier(targetStructName)

			gongStructShape.Links = append(gongStructShape.Links, link)
			link.Middlevertice = new(gongdoc_models.Vertice).Stage(gongdocStage)
			link.Middlevertice.Name = "Verticle in class diagram " + nodeImplField.diagramPackage.SelectedClassdiagram.Name +
				" in middle between " + gongStructShape.Name + " and " + targetGongStructShape.Name
			link.Middlevertice.X = (gongStructShape.Position.X+targetGongStructShape.Position.X)/2.0 +
				gongStructShape.Width*1.5
			link.Middlevertice.Y = (gongStructShape.Position.Y+targetGongStructShape.Position.Y)/2.0 +
				gongStructShape.Heigth/2.0

			link.StartOrientation = gongdoc_models.ORIENTATION_HORIZONTAL
			link.StartRatio = 0.5
			link.EndOrientation = gongdoc_models.ORIENTATION_HORIZONTAL
			link.EndRatio = 0.5
			link.CornerOffsetRatio = 0.8
		}
	}
	computeGongNodesConfigurations(
		gongdocStage,
		nodeImplField.diagramPackage.SelectedClassdiagram,
		nodeImplField.treeOfGongObjects)
	gongdocStage.Commit()
}
