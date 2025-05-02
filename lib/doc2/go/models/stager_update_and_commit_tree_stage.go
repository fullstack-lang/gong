package models

import (
	"log"

	gong "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) UpdateAndCommitTreeStage() {

	stager.treeStage.Reset()

	// put a "class diagram button at the root"
	root := &tree.Node{
		Name:       "Class Diagrams",
		IsExpanded: true,
		Buttons: []*tree.Button{
			tree.NewButton(
				&ButtonNewClassdiagramProxy{
					stager: stager,
				},
				"Class Diagramm Add Button",
				string(buttons.BUTTON_add)),
		},
	}

	// append a node below for each diagram
	diagramPackage := getTheDiagramPackage(stager.stage)
	for _, classDiagram := range diagramPackage.Classdiagrams {
		var selected bool
		if diagramPackage.SelectedClassdiagram == classDiagram {
			selected = true
		}
		nodeClassdiagram := &tree.Node{
			Name:              classDiagram.Name,
			HasCheckboxButton: true,
			IsChecked:         selected,
			IsExpanded:        classDiagram.IsExpanded,
		}
		nodeClassdiagram.Impl = &ClassDiagramNodeProxy{
			node:         nodeClassdiagram,
			stager:       stager,
			classDiagram: classDiagram,
		}

		root.Children = append(root.Children, nodeClassdiagram)

		if diagramPackage.SelectedClassdiagram != classDiagram {
			continue
		}

		map_modelElement_shape := stager.compute_map_modelElement_shape(classDiagram, stager.gongStage)
		for _, gongStruct := range gong.GetGongstrucsSorted[*gong.GongStruct](stager.gongStage) {

			shape, isInDiagram := map_modelElement_shape[gongStruct]

			gongStructShape, ok := shape.(*GongStructShape)
			if isInDiagram && !ok {
				log.Fatalln("A gongstruct shape should be mapped to a gongstruct")
			}
			var isExpanded bool
			if isInDiagram {
				isExpanded = gongStructShape.IsExpanded
			}

			nodeNamedStruct := &tree.Node{
				Name:              gongStruct.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
				IsExpanded:        isExpanded,
			}
			nodeNamedStruct.Impl = &GongstructNodeProxy{
				node:            nodeNamedStruct,
				stager:          stager,
				classDiagram:    classDiagram,
				gongstruct:      gongStruct,
				gongStructShape: gongStructShape,
			}
			nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodeNamedStruct)

			for _, field := range gongStruct.Fields {
				shape, isInDiagram := map_modelElement_shape[field]

				attributeShape, isAFieldShape := shape.(*AttributeShape)
				if isInDiagram && !isAFieldShape {
					log.Fatalln("A field should be mapped to a field shape or a link shape")
				}

				// if field is a link to another, disable the checkbox
				// if the target gong struct shape is not in the diagram
				isCheckboxDisabled := false
				if link, ok := field.(*gong.PointerToGongStructField); ok {
					if _, ok = map_modelElement_shape[link.GongStruct]; !ok {
						isCheckboxDisabled = true
					}
				}
				if link, ok := field.(*gong.SliceOfPointerToGongStructField); ok {
					if _, ok = map_modelElement_shape[link.GongStruct]; !ok {
						isCheckboxDisabled = true
					}
				}

				nodeField := &tree.Node{
					Name:               field.GetName(),
					HasCheckboxButton:  true,
					IsChecked:          isInDiagram,
					IsCheckboxDisabled: isCheckboxDisabled,
				}

				nodeField.Impl = &AttributeFieldNodeProxy{
					node:            nodeNamedStruct,
					stager:          stager,
					classDiagram:    classDiagram,
					gongstruct:      gongStruct,
					gongStructShape: gongStructShape,
					field:           field,
					attributeShape:  attributeShape,
				}

				nodeNamedStruct.Children = append(nodeNamedStruct.Children, nodeField)
			}
		}

	}
	tree.StageBranch(stager.treeStage,
		&tree.Tree{
			Name:      stager.stage.GetProbeTreeSidebarStageName(),
			RootNodes: []*tree.Node{root},
		},
	)

	stager.treeStage.Commit()
}

type ClassDiagramNodeProxy struct {
	node         *tree.Node
	stager       *Stager
	classDiagram *Classdiagram
}

// OnAfterUpdate is called when a node is checked/unchecked by the user
func (proxy *ClassDiagramNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		diagramPackage.SelectedClassdiagram = proxy.classDiagram

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		diagramPackage.SelectedClassdiagram = nil

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	if front.IsExpanded && !staged.IsExpanded {
		proxy.classDiagram.IsExpanded = true
		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		proxy.classDiagram.IsExpanded = false
		front.IsExpanded = true

		proxy.stager.stage.Commit()
	}
}

type GongstructNodeProxy struct {
	node            *tree.Node
	stager          *Stager
	classDiagram    *Classdiagram
	gongstruct      *gong.GongStruct
	gongStructShape *GongStructShape
}

// OnAfterUpdate is called when a node is checked/unchecked by the user
func (proxy *GongstructNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		proxy.classDiagram.AddGongStructShape(proxy.stager.stage, diagramPackage, proxy.gongstruct.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		proxy.classDiagram.RemoveGongStructShape(proxy.stager.stage, proxy.gongstruct.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	if front.IsExpanded && !staged.IsExpanded {
		proxy.gongStructShape.IsExpanded = true
		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		proxy.gongStructShape.IsExpanded = false
		front.IsExpanded = true

		proxy.stager.stage.Commit()
	}
}

type AttributeFieldNodeProxy struct {
	node            *tree.Node
	stager          *Stager
	classDiagram    *Classdiagram
	gongstruct      *gong.GongStruct
	gongStructShape *GongStructShape
	field           gong.FieldInterface
	attributeShape  *AttributeShape
}

func (proxy *AttributeFieldNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		proxy.classDiagram.AddAttributeFieldShape(
			proxy.stager.stage,
			proxy.stager.gongStage,
			proxy.gongstruct,
			proxy.field,
			proxy.gongStructShape)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		proxy.classDiagram.RemoveAttributeFieldShape(proxy.stager.stage, proxy.attributeShape, proxy.gongStructShape)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}
}
