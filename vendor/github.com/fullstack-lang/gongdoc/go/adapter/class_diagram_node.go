package adapter

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/doc2svg"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type ClassDiagramNode struct {
	portfolioAdapter *PortfolioAdapter
	classdiagram     *gongdoc_models.Classdiagram
	isInRenameMode   bool

	// mementos is the implementation of the mementos pattern
	mementos []*gongdoc_models.Classdiagram

	parentNode diagrammer.PortfolioNode
}

// static check that it meets the intended interface
var _ diagrammer.PortfolioDiagramNode = &(ClassDiagramNode{})

func NewClassDiagramNode(
	portfolioAdapter *PortfolioAdapter,
	parentNode diagrammer.PortfolioNode,
	classDiagram *gongdoc_models.Classdiagram,
) (classDiagramNode *ClassDiagramNode) {

	classDiagramNode = &ClassDiagramNode{
		portfolioAdapter: portfolioAdapter,
		parentNode:       parentNode,
	}

	classDiagramNode.classdiagram = classDiagram

	return
}

// GetParent implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) GetParent() diagrammer.PortfolioNode {
	return classDiagramNode.parentNode
}

// GetChildren implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) GetChildren() []diagrammer.PortfolioNode {
	return nil
}

// AppendChildren implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) AppendChildren(diagrammer.PortfolioNode) {
	panic("unimplemented")
}

// IsInRenameMode implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) IsInRenameMode() bool {
	return classDiagramNode.isInRenameMode
}

func (classDiagramNode *ClassDiagramNode) SetIsInRenameMode(isInRenameMode bool) {
	classDiagramNode.isInRenameMode = isInRenameMode
}

// HasDiagramRenameButton implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) HasDiagramRenameButton() bool {
	return classDiagramNode.portfolioAdapter.getDiagramPackage().IsEditable
}

// RenameDiagram implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) RenameDiagram(newName string) (err error) {

	// check that the name is a correct identifer in go
	for _, b := range newName {
		if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_' || '0' <= b && b <= '9' {
			// Avoid assigning a rune for the common case of an ascii character.
			continue
		} else {
			err = errors.New("The name of the diagram is not a correct identifier in go: " + newName)
			return
		}
	}

	classdiagram := classDiagramNode.classdiagram
	gongdocStage := classDiagramNode.portfolioAdapter.gongdocStage
	diagramPackage := classDiagramNode.portfolioAdapter.getDiagramPackage()

	// rename the diagram file if it exists
	// remove the actual classdiagram file if it exsits
	oldClassdiagramFilePath := filepath.Join(diagramPackage.Path,
		"../diagrams", classdiagram.Name) + ".go"
	newClassdiagramFilePath := filepath.Join(diagramPackage.Path,
		"../diagrams", newName) + ".go"

	if _, err := os.Stat(oldClassdiagramFilePath); err != nil {
		return err
	}
	if err := os.Remove(oldClassdiagramFilePath); err != nil {
		return err
	}

	file, err := os.Create(newClassdiagramFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// checkout in order to get the latest version of the diagram before
	// modifying it updated by the front
	gongdocStage.Checkout()
	gongdocStage.Unstage()
	gongdoc_models.StageBranch(gongdocStage, classdiagram)
	classdiagram.Name = newName

	gongdoc_models.SetupMapDocLinkRenaming(diagramPackage.ModelPkg.Stage_, gongdocStage)

	// save the diagram
	gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

	// restore the original stage
	gongdocStage.Unstage()
	gongdocStage.Checkout()

	classdiagram.Name = newName
	gongdocStage.Commit()

	classDiagramNode.isInRenameMode = false

	return
}

// GeneratesProgeny implements diagrammer.Node.
func (classDiagramNode *ClassDiagramNode) GeneratesProgeny() (children []diagrammer.PortfolioNode) {
	return
}

// GetName implements diagrammer.Node.
func (classDiagramNode *ClassDiagramNode) GetName() string {
	return classDiagramNode.classdiagram.GetName()
}

// IsExpanded implements bridge.PortfolioNode.
func (*ClassDiagramNode) IsExpanded() bool {
	return true
}

// IsNameEditable implements bridge.PortfolioNode.
func (*ClassDiagramNode) IsNameEditable() bool {
	return false
}

// RemoveChildren implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) RemoveChildren(diagrammer.PortfolioNode) {
	panic("unimplemented")
}

// DisplayDiagram implements diagrammer.Portfolio.
//
// It generated the SVG
func (classDiagramNode *ClassDiagramNode) DisplayDiagram() (
	setOfModelElementNode map[diagrammer.ModelElementNode]diagrammer.Shape) {

	gongdocStage := classDiagramNode.portfolioAdapter.gongdocStage
	gongsvgStage := classDiagramNode.portfolioAdapter.gongsvgStage
	gongStage := classDiagramNode.portfolioAdapter.gongStage

	diagramPackage := classDiagramNode.portfolioAdapter.getDiagramPackage()

	diagramPackage.SelectedClassdiagram = classDiagramNode.classdiagram
	selectedClassdiagram := diagramPackage.SelectedClassdiagram

	docSVGMapper := doc2svg.NewDocSVGMapper(gongsvgStage)
	docSVGMapper.GenerateSvg(gongdocStage)

	setOfModelElementNode = classDiagramNode.getSetOfModelElementNodesInDiagram(gongStage, selectedClassdiagram)

	classDiagramNode.portfolioAdapter.setSelectedClassdiagramNode(classDiagramNode)

	return
}

// compute the set of Model Node
// 1. Create the map of model element to model node
// 2. Parse the selected diagram, for every shape, get the element model of the shape
// use the aforementioned map to populate the set
// parse fields of gongstruct to match the field shape
// parse values of gongstruct to match the value shape
func (classDiagramNode *ClassDiagramNode) getSetOfModelElementNodesInDiagram(
	gongStage *gong_models.StageStruct, selectedClassdiagram *gongdoc_models.Classdiagram) (
	setOfModelElementNode map[diagrammer.ModelElementNode]diagrammer.Shape) {
	setOfModelElementNode = make(map[diagrammer.ModelElementNode]diagrammer.Shape)

	map_ModelElement_ModelNode := make(map[any]diagrammer.ModelElementNode)
	for modelNode := range classDiagramNode.portfolioAdapter.diagrammer.GetMap_elementNode_treeNode() {

		if elementNode, ok := modelNode.(diagrammer.ModelElementNode); ok {
			if modelElement := elementNode.GetElement(); modelElement != nil {
				map_ModelElement_ModelNode[modelElement] = modelNode
			}
		}
	}

	gongStructSet := *gong_models.GetGongstructInstancesMap[gong_models.GongStruct](gongStage)
	for _, gongStructShape := range selectedClassdiagram.GongStructShapes {

		gongStructName := gongdoc_models.IdentifierToGongObjectName(gongStructShape.Identifier)
		gongStruct, ok := gongStructSet[gongStructName]

		gongStructNode, ok := map_ModelElement_ModelNode[gongStruct]
		if !ok {
			log.Fatalln("unkown element", gongStructShape.Identifier)
		}

		setOfModelElementNode[gongStructNode] = gongStructShape

		for _, fieldShape := range gongStructShape.Fields {
			fieldShapeName := gongdoc_models.IdentifierToFieldName(fieldShape.Identifier)

			for _, field := range gongStruct.Fields {
				if field.GetName() == fieldShapeName {

					gongStructFieldNode, ok := map_ModelElement_ModelNode[field]
					if !ok {
						log.Fatalln("unkown element", fieldShape.Identifier)
					}

					setOfModelElementNode[gongStructFieldNode] = fieldShape
				}
			}
		}

		for _, linkShape := range gongStructShape.Links {
			linkShapeName := gongdoc_models.IdentifierToFieldName(linkShape.Identifier)

			for _, link := range gongStruct.SliceOfPointerToGongStructFields {
				if link.GetName() == linkShapeName {
					linkNode, ok := map_ModelElement_ModelNode[link]
					if !ok {
						log.Fatalln("unkown element", linkShape.Identifier)
					}

					setOfModelElementNode[linkNode] = linkShape
				}
			}

			for _, link := range gongStruct.PointerToGongStructFields {
				if link.GetName() == linkShapeName {
					linkNode, ok := map_ModelElement_ModelNode[link]
					if !ok {
						log.Fatalln("unkown element", linkShape.Identifier)
					}

					setOfModelElementNode[linkNode] = linkShape
				}
			}
		}
	}

	gongEnumSet := *gong_models.GetGongstructInstancesMap[gong_models.GongEnum](gongStage)
	for _, gongEnumShape := range selectedClassdiagram.GongEnumShapes {

		gongEnumName := gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier)
		gongEnum, ok := gongEnumSet[gongEnumName]

		gongEnumNode, ok := map_ModelElement_ModelNode[gongEnum]
		if !ok {
			log.Fatalln("unkown element", gongEnumShape.Identifier)
		}

		setOfModelElementNode[gongEnumNode] = gongEnumShape

		for _, valueShape := range gongEnumShape.GongEnumValueEntrys {
			valueShapeName := gongdoc_models.IdentifierToFieldName(valueShape.Identifier)

			for _, value := range gongEnum.GongEnumValues {
				if value.GetName() == valueShapeName {

					gongEnumValueNode, ok := map_ModelElement_ModelNode[value]
					if !ok {
						log.Fatalln("unkown element", valueShape.Identifier)
					}

					setOfModelElementNode[gongEnumValueNode] = valueShape
				}
			}
		}
	}

	gongNoteSet := *gong_models.GetGongstructInstancesMap[gong_models.GongNote](gongStage)
	for _, gongNoteShape := range selectedClassdiagram.NoteShapes {

		gongNoteName := gongdoc_models.IdentifierToGongObjectName(gongNoteShape.Identifier)
		gongNote, ok := gongNoteSet[gongNoteName]

		gongNoteNode, ok := map_ModelElement_ModelNode[gongNote]
		if !ok {
			log.Fatalln("unkown element", gongNoteShape.Identifier)
		}

		setOfModelElementNode[gongNoteNode] = gongNoteShape

		for _, noteLinkShape := range gongNoteShape.NoteShapeLinks {

			switch noteLinkShape.Type {
			case gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE:
				nodeLinkShapeTarget := gongdoc_models.IdentifierToGongObjectName(noteLinkShape.Identifier)
				for _, link := range gongNote.Links {
					if link.GetName() == nodeLinkShapeTarget {
						linkNode, ok := map_ModelElement_ModelNode[link]
						if !ok {
							log.Fatalln("unkown element", noteLinkShape.Identifier)
						}

						setOfModelElementNode[linkNode] = noteLinkShape
					}
				}
			case gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_FIELD:
				receiver, fieldName := gongdoc_models.IdentifierToReceiverAndFieldName(noteLinkShape.Identifier)
				for _, link := range gongNote.Links {
					if link.GetName() == fieldName && link.Recv == receiver {
						linkNode, ok := map_ModelElement_ModelNode[link]
						if !ok {
							log.Fatalln("unkown element", noteLinkShape.Identifier)
						}

						setOfModelElementNode[linkNode] = noteLinkShape
					}
				}

			}
		}
	}
	return
}

// HasDuplicateButton implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) HasDuplicateButton() bool {
	return classDiagramNode.portfolioAdapter.getDiagramPackage().IsEditable
}

// DuplicateDiagram implements diagrammer.PortfolioDiagramNode
//
// DuplicateDiagram performs the duplication with a deep copy
func (classDiagramNode *ClassDiagramNode) DuplicateDiagram() diagrammer.PortfolioDiagramNode {

	gongdocStage := classDiagramNode.portfolioAdapter.gongdocStage

	diagramPackage := classDiagramNode.portfolioAdapter.getDiagramPackage()
	diagramPackage.SelectedClassdiagram = classDiagramNode.classdiagram
	selectedClassdiagram := diagramPackage.SelectedClassdiagram

	var hasNameCollision bool
	initialName := selectedClassdiagram.Name
	newClassdiagramName := initialName
	index := 0
	// loop until the name of the new diagram is not in collision with an existing
	// diagram name
	for index == 0 || hasNameCollision {
		index++
		hasNameCollision = false
		for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {
			if classdiagram.Name == newClassdiagramName {
				hasNameCollision = true
			}
		}
		if hasNameCollision {
			newClassdiagramName = initialName + fmt.Sprintf("_%d", index)
		}
	}

	log.Println("New name is", newClassdiagramName)

	newClassdiagram := selectedClassdiagram.DuplicateDiagram()
	newClassdiagram.Name = newClassdiagramName
	diagramPackage.Classdiagrams = append(diagramPackage.Classdiagrams, newClassdiagram)
	gongdoc_models.StageBranch(gongdocStage, newClassdiagram)

	// the gongstage has now the new class diagram within
	gongdocStage.Commit()

	classDiagramNode.marshallDiagram(gongdocStage, newClassdiagram, diagramPackage)

	parent := classDiagramNode.GetParent()
	newClassDiagramNode := NewClassDiagramNode(classDiagramNode.portfolioAdapter, parent, newClassdiagram)
	parent.AppendChildren(newClassDiagramNode)

	return newClassDiagramNode
}

// to generate the diagram file, one uses the function that marshall the whole gongdoc stage
// into a file.
// in order to do that:
//
// 1. one empties the stage
// 2. fill up the stage with only this diagram
// 3. marshall the stage into the dedicated file
// 4. restore the stage
// 5. empties the stage
// 6. stage the branch with the duplicated diagram
// 7. marshall the stage
// gongdoc_models.SetupMapDocLinkRenaming(diagramPackage.ModelPkg.Stage_, gongdocStage)
// 8. restore the stage
func (*ClassDiagramNode) marshallDiagram(
	gongdocStage *gongdoc_models.StageStruct,
	newClassdiagram *gongdoc_models.Classdiagram,
	diagramPackage *gongdoc_models.DiagramPackage) {
	gongdocStage.Checkout()

	gongdocStage.Unstage()

	gongdoc_models.StageBranch(gongdocStage, newClassdiagram)

	newClassdiagramFilePath := filepath.Join(diagramPackage.Path, "../diagrams", newClassdiagram.GetName()) + ".go"

	file, err := os.Create(newClassdiagramFilePath)
	if err != nil {
		log.Fatal("Cannot open diagram file" + err.Error())
	}
	gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")
	file.Close()

	gongdocStage.Checkout()
}

// HasDeleteButton implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) HasDeleteButton() bool {
	return classDiagramNode.portfolioAdapter.getDiagramPackage().IsEditable
}

// DeleteDiagram implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) DeleteDiagram() {

	gongdocStage := classDiagramNode.portfolioAdapter.gongdocStage
	gongsvgStage := classDiagramNode.portfolioAdapter.gongsvgStage

	// checkout the stage, it shall remove the link between
	// the parent node and the staged node because 0..1->0..N association
	// is stored in the staged node as a reverse pointer
	gongdocStage.Checkout()
	diagramPackage := classDiagramNode.portfolioAdapter.getDiagramPackage()
	diagramPackage.SelectedClassdiagram = classDiagramNode.classdiagram
	selectedClassdiagram := diagramPackage.SelectedClassdiagram

	// remove the classdiagram node from the pkg element node
	idx := slices.Index(diagramPackage.Classdiagrams, selectedClassdiagram)
	diagramPackage.Classdiagrams = slices.Delete(diagramPackage.Classdiagrams, idx, idx+1)
	gongdoc_models.UnstageBranch(gongdocStage, selectedClassdiagram)

	// remove the actual classdiagram file if it exsits
	classdiagramFilePath := filepath.Join(diagramPackage.Path, "../diagrams", selectedClassdiagram.Name) + ".go"
	if _, err := os.Stat(classdiagramFilePath); err == nil {
		if err := os.Remove(classdiagramFilePath); err != nil {
			log.Println("Error while deleting file " + classdiagramFilePath + " : " + err.Error())
		}
	}

	diagramPackage.SelectedClassdiagram = nil
	gongdocStage.Commit()

	docSVGMapper := doc2svg.NewDocSVGMapper(gongsvgStage)
	docSVGMapper.GenerateSvg(gongdocStage)

}

// HasDrawButton implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) HasDrawButton() bool {
	return classDiagramNode.portfolioAdapter.getDiagramPackage().IsEditable
}

// IsInDrawingMode implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) IsInDrawingMode() bool {
	return classDiagramNode.classdiagram.IsInDrawMode
}

// SetIsInEditMode implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) DrawDiagram() (
	setOfModelElementNode map[diagrammer.ModelElementNode]diagrammer.Shape) {

	gongsvgStage := classDiagramNode.portfolioAdapter.gongsvgStage
	gongdocStage := classDiagramNode.portfolioAdapter.gongdocStage
	gongStage := classDiagramNode.portfolioAdapter.gongStage

	// initiate the mementos
	classDiagramNode.mementos = classDiagramNode.mementos[:0]
	classDiagramNode.mementos = append(classDiagramNode.mementos,
		classDiagramNode.classdiagram.DuplicateDiagram())

	// set the diagram node in draw mode
	classDiagramNode.classdiagram.IsInDrawMode = true
	gongdocStage.Commit()

	docSVGMapper := doc2svg.NewDocSVGMapper(gongsvgStage)
	docSVGMapper.GenerateSvg(gongdocStage)

	diagramPackage := classDiagramNode.portfolioAdapter.getDiagramPackage()
	selectedClassdiagram := diagramPackage.SelectedClassdiagram
	setOfModelElementNode = classDiagramNode.getSetOfModelElementNodesInDiagram(gongStage, selectedClassdiagram)

	return
}

// CancelEdit implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) CancelEdit() (
	setOfModelElementNode map[diagrammer.ModelElementNode]diagrammer.Shape) {

	gongsvgStage := classDiagramNode.portfolioAdapter.gongsvgStage
	gongdocStage := classDiagramNode.portfolioAdapter.gongdocStage
	gongStage := classDiagramNode.portfolioAdapter.gongStage

	diagramPackage := classDiagramNode.portfolioAdapter.getDiagramPackage()
	selectedClassdiagram := diagramPackage.SelectedClassdiagram

	// remove the classdiagram node from the pkg element node
	idx := slices.Index(diagramPackage.Classdiagrams, selectedClassdiagram)
	diagramPackage.Classdiagrams = slices.Delete(diagramPackage.Classdiagrams, idx, idx+1)
	gongdoc_models.UnstageBranch(gongdocStage, selectedClassdiagram)

	revertedClassdiagram := classDiagramNode.mementos[0]
	gongdoc_models.StageBranch(gongdocStage, revertedClassdiagram)

	diagramPackage.Classdiagrams = append(diagramPackage.Classdiagrams, revertedClassdiagram)
	diagramPackage.SelectedClassdiagram = revertedClassdiagram
	classDiagramNode.classdiagram = revertedClassdiagram

	classDiagramNode.classdiagram.IsInDrawMode = false
	gongdocStage.Commit()

	docSVGMapper := doc2svg.NewDocSVGMapper(gongsvgStage)
	docSVGMapper.GenerateSvg(gongdocStage)

	setOfModelElementNode = classDiagramNode.getSetOfModelElementNodesInDiagram(gongStage, revertedClassdiagram)

	return
}

// SaveDiagram implements diagrammer.PortfolioDiagramNode.
func (classDiagramNode *ClassDiagramNode) SaveDiagram() (
	setOfModelElementNode map[diagrammer.ModelElementNode]diagrammer.Shape) {
	gongdocStage := classDiagramNode.portfolioAdapter.gongdocStage
	gongStage := classDiagramNode.portfolioAdapter.gongStage

	diagramPackage := classDiagramNode.portfolioAdapter.getDiagramPackage()
	diagramPackage.SelectedClassdiagram = classDiagramNode.classdiagram
	selectedClassdiagram := diagramPackage.SelectedClassdiagram
	selectedClassdiagram.IsInDrawMode = false
	gongdocStage.Commit()

	classDiagramNode.marshallDiagram(gongdocStage, selectedClassdiagram, diagramPackage)

	setOfModelElementNode = classDiagramNode.getSetOfModelElementNodesInDiagram(gongStage, selectedClassdiagram)

	return
}
