package models

import (
	"fmt"
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
			{
				Name: "Class Diagramm Add Button",
				Icon: string(buttons.BUTTON_add),
				Impl: &ButtonNewClassdiagramProxy{
					stager: stager,
				},
				HasToolTip:      true,
				ToolTipText:     "Create a new diagram",
				ToolTipPosition: tree.Above,
			},
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
			IsInEditMode:      classDiagram.IsInRenameMode,
		}
		nodeClassdiagram.Impl = &ClassDiagramNodeProxy{
			node:         nodeClassdiagram,
			stager:       stager,
			classDiagram: classDiagram,
		}

		nodeClassdiagram.Buttons = append(nodeClassdiagram.Buttons,
			&tree.Button{
				Name: classDiagram.GetName() + " " + string(buttons.BUTTON_delete),
				Icon: string(buttons.BUTTON_delete),
				Impl: NewClassDiagramButtonProxy(
					stager,
					classDiagram,
					nodeClassdiagram,
					REMOVE,
				),
				HasToolTip:      true,
				ToolTipText:     "Delete the diagram",
				ToolTipPosition: tree.Above,
			})

		if !classDiagram.IsInRenameMode {
			nodeClassdiagram.Buttons = append(nodeClassdiagram.Buttons,
				&tree.Button{
					Name: classDiagram.GetName() + " " + string(buttons.BUTTON_edit_note),
					Icon: string(buttons.BUTTON_edit_note),
					Impl: NewClassDiagramButtonProxy(
						stager,
						classDiagram,
						nodeClassdiagram,
						RENAME,
					),
					HasToolTip:      true,
					ToolTipText:     "Rename the diagram",
					ToolTipPosition: tree.Above,
				})
		} else {
			nodeClassdiagram.Buttons = append(nodeClassdiagram.Buttons,
				&tree.Button{
					Name: classDiagram.GetName() + " " + string(buttons.BUTTON_edit_off),
					Icon: string(buttons.BUTTON_edit_off),
					Impl: NewClassDiagramButtonProxy(
						stager,
						classDiagram,
						nodeClassdiagram,
						RENAME_CANCEL,
					),
					HasToolTip:      true,
					ToolTipText:     "Cancel renaming",
					ToolTipPosition: tree.Above,
				})
		}

		nodeClassdiagram.Buttons = append(nodeClassdiagram.Buttons,
			&tree.Button{
				Name: classDiagram.GetName() + " " + string(buttons.BUTTON_copy_all),
				Icon: string(buttons.BUTTON_copy_all),
				Impl: NewClassDiagramButtonProxy(
					stager,
					classDiagram,
					nodeClassdiagram,
					DUPLICATE,
				),
				HasToolTip:      true,
				ToolTipText:     "Duplicate diagram",
				ToolTipPosition: tree.Above,
			})

		root.Children = append(root.Children, nodeClassdiagram)

		// if diagramPackage.SelectedClassdiagram != classDiagram {
		// 	continue
		// }

		map_modelElement_shape := stager.compute_map_modelElement_shape(classDiagram, stager.gongStage)

		gongstructs := gong.GetGongstrucsSorted[*gong.GongStruct](stager.gongStage)
		nbGongstructsInDiagram := 0
		for _, gongStruct := range gongstructs {
			_, isInDiagram := map_modelElement_shape[gongStruct]
			if isInDiagram {
				nbGongstructsInDiagram++
			}
		}
		nodeNamedStructs := &tree.Node{
			Name:       fmt.Sprintf("Gongstructs (%d/%d)", nbGongstructsInDiagram, len(gongstructs)),
			IsExpanded: classDiagram.NodeNamedStructsIsExpanded,
			Impl: &ClassDiagramGongstructsNodeProxy{
				stager:       stager,
				classDiagram: classDiagram,
			},
		}
		nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodeNamedStructs)

		gongenums := gong.GetGongstrucsSorted[*gong.GongEnum](stager.gongStage)
		nbGongenumsInDiagram := 0
		for _, gongEnums := range gongenums {
			_, isInDiagram := map_modelElement_shape[gongEnums]
			if isInDiagram {
				nbGongenumsInDiagram++
			}
		}
		nodeGongEnums := &tree.Node{
			Name:       fmt.Sprintf("Gongenums (%d/%d)", nbGongenumsInDiagram, len(gongenums)),
			IsExpanded: classDiagram.NodeGongEnumsIsExpanded,
			Impl: &ClassDiagramGongstructsNodeProxy{
				stager:       stager,
				classDiagram: classDiagram,
			},
		}
		nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodeGongEnums)

		gongnotes := gong.GetGongstrucsSorted[*gong.GongNote](stager.gongStage)
		nbGongnotesInDiagram := 0
		for _, gongNotes := range gongnotes {
			_, isInDiagram := map_modelElement_shape[gongNotes]
			if isInDiagram {
				nbGongnotesInDiagram++
			}
		}
		nodeGongNotes := &tree.Node{
			Name:       fmt.Sprintf("Gongnotes (%d/%d)", nbGongnotesInDiagram, len(gongnotes)),
			IsExpanded: classDiagram.NodeGongNotesIsExpanded,
			Impl: &ClassDiagramGongstructsNodeProxy{
				stager:       stager,
				classDiagram: classDiagram,
			},
		}
		nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodeGongNotes)

		for _, gongStruct := range gongstructs {

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
			nodeNamedStructs.Children = append(nodeNamedStructs.Children, nodeNamedStruct)

			for _, field := range gongStruct.Fields {

				switch field.(type) {
				case *gong.GongBasicField, *gong.GongTimeField:
					shape, isInDiagram := map_modelElement_shape[field]

					attributeShape, isAFieldShape := shape.(*AttributeShape)
					if isInDiagram && !isAFieldShape {
						log.Fatalln("A field should be mapped to a field shape or a link shape")
					}

					nodeField := &tree.Node{
						Name:              field.GetName(),
						HasCheckboxButton: true,
						IsChecked:         isInDiagram,
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
				case *gong.PointerToGongStructField, *gong.SliceOfPointerToGongStructField:
					shape, isInDiagram := map_modelElement_shape[field]

					linkShape, isALinkShape := shape.(*LinkShape)
					if isInDiagram && !isALinkShape {
						log.Fatalln("A pointer or slice of pointer field should be mapped to a link shape")
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

					nodeField.Impl = &LinkFieldNodeProxy{
						node:            nodeNamedStruct,
						stager:          stager,
						classDiagram:    classDiagram,
						gongstruct:      gongStruct,
						gongStructShape: gongStructShape,
						field:           field,
						linkShape:       linkShape,
					}

					nodeNamedStruct.Children = append(nodeNamedStruct.Children, nodeField)
				default:
					log.Fatalln("Unknwon field type")
				}
			}
		}
		for _, gongEnum := range gongenums {
			shape, isInDiagram := map_modelElement_shape[gongEnum]

			gongEnumShape, ok := shape.(*GongEnumShape)
			if isInDiagram && !ok {
				log.Fatalln("a gongenum should be associated to a gongenum shape")
			}
			_ = gongEnumShape

			var isExpanded bool
			if isInDiagram {
				isExpanded = gongEnumShape.IsExpanded
			}
			node := &tree.Node{
				Name:              gongEnum.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
				IsExpanded:        isExpanded,
			}
			nodeGongEnums.Children = append(nodeGongEnums.Children, node)
		}

		for _, gongNote := range gongnotes {
			shape, isInDiagram := map_modelElement_shape[gongNote]

			gongNoteShape, ok := shape.(*NoteShape)
			if isInDiagram && !ok {
				log.Fatalln("a gongnote should be associated to a gongnote shape")
			}
			_ = gongNoteShape

			var isExpanded bool
			if isInDiagram {
				isExpanded = gongNoteShape.IsExpanded
			}
			node := &tree.Node{
				Name:              gongNote.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
				IsExpanded:        isExpanded,
			}
			nodeGongNotes.Children = append(nodeGongNotes.Children, node)
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

	if front.Name != staged.Name {
		proxy.classDiagram.Name = front.Name
		proxy.classDiagram.IsInRenameMode = false

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.stage.Commit()
	}
}

type ClassDiagramGongstructsNodeProxy struct {
	stager       *Stager
	classDiagram *Classdiagram
}

func (proxy *ClassDiagramGongstructsNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	if front.IsExpanded && !staged.IsExpanded {
		proxy.classDiagram.NodeNamedStructsIsExpanded = true
		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		proxy.classDiagram.NodeNamedStructsIsExpanded = false
		front.IsExpanded = true

		proxy.stager.stage.Commit()
	}
}

type ClassDiagramGongEnumsNodeProxy struct {
	stager       *Stager
	classDiagram *Classdiagram
}

func (proxy *ClassDiagramGongEnumsNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	if front.IsExpanded && !staged.IsExpanded {
		proxy.classDiagram.NodeGongEnumsIsExpanded = true
		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		proxy.classDiagram.NodeGongEnumsIsExpanded = false
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

type LinkFieldNodeProxy struct {
	node            *tree.Node
	stager          *Stager
	classDiagram    *Classdiagram
	gongstruct      *gong.GongStruct
	gongStructShape *GongStructShape
	field           gong.FieldInterface
	linkShape       *LinkShape
}

func (proxy *LinkFieldNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		proxy.classDiagram.AddLinkFieldShape(
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
		proxy.classDiagram.RemoveLinkFieldShape(proxy.stager.stage, proxy.linkShape, proxy.gongStructShape)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}
}
