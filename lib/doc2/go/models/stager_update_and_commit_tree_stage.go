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

	classdiagramsTree := &tree.Tree{
		Name: stager.stage.GetProbeTreeSidebarStageName(),
	}

	// put a "class diagram button at the root"
	addDiagramButton := &tree.Button{
		Name: "Class Diagramm Add Button",
		Icon: string(buttons.BUTTON_add),
		Impl: &ButtonNewClassdiagramProxy{
			stager: stager,
		},
		HasToolTip:      true,
		ToolTipText:     "Create a new diagram",
		ToolTipPosition: tree.Above,
	}
	root := &tree.Node{
		Name:       "Class Diagrams",
		IsExpanded: true,
	}
	if !stager.embeddedDiagrams {
		root.Buttons = append(root.Buttons, addDiagramButton)
	}
	classdiagramsTree.RootNodes = append(classdiagramsTree.RootNodes, root)

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

		if !stager.embeddedDiagrams {
			stager.addButtonsToClassdiagramNode(nodeClassdiagram, classDiagram)
		}

		// if the classdiagram appear as sub node of the classdiagram node
		// uses the following line instead
		// 	root.Children = append(root.Children, nodeClassdiagram)
		classdiagramsTree.RootNodes = append(classdiagramsTree.RootNodes, nodeClassdiagram)

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
			IsExpanded: classDiagram.NodeGongStructsIsExpanded,
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
			Impl: &ClassDiagramGongEnumsNodeProxy{
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

		for idx, gongStruct := range gongstructs {

			shape, isGongStructShapeInDiagram := map_modelElement_shape[gongStruct]

			gongStructShape, ok := shape.(*GongStructShape)
			if isGongStructShapeInDiagram && !ok {
				log.Fatalln("A gongstruct shape should be mapped to a gongstruct")
			}
			isExpanded := IsNodeExpanded(classDiagram.NodeGongStructNodeExpansionBinaryEncoding, idx)

			nodeNamedStruct := &tree.Node{
				Name:               gongStruct.Name,
				HasCheckboxButton:  true,
				IsCheckboxDisabled: stager.embeddedDiagrams,
				IsChecked:          isGongStructShapeInDiagram,
				IsExpanded:         isExpanded,
			}
			nodeNamedStruct.Impl = &GongStructNodeProxy{
				node:            nodeNamedStruct,
				stager:          stager,
				classDiagram:    classDiagram,
				gongstruct:      gongStruct,
				gongStructShape: gongStructShape,
				rank:            idx,
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
						Name:               field.GetName(),
						HasCheckboxButton:  true,
						IsChecked:          isInDiagram,
						IsCheckboxDisabled: !isGongStructShapeInDiagram || stager.embeddedDiagrams,
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
					isTargetGongstructAbsent := false
					if link, ok := field.(*gong.PointerToGongStructField); ok {
						if _, ok = map_modelElement_shape[link.GongStruct]; !ok {
							isTargetGongstructAbsent = true
						}
					}
					if link, ok := field.(*gong.SliceOfPointerToGongStructField); ok {
						if _, ok = map_modelElement_shape[link.GongStruct]; !ok {
							isTargetGongstructAbsent = true
						}
					}

					nodeField := &tree.Node{
						Name:               field.GetName(),
						HasCheckboxButton:  true,
						IsChecked:          isInDiagram,
						IsCheckboxDisabled: !isGongStructShapeInDiagram || isTargetGongstructAbsent || stager.embeddedDiagrams,
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
		for idx, gongEnum := range gongenums {
			shape, isEnumInDiagram := map_modelElement_shape[gongEnum]

			gongEnumShape, ok := shape.(*GongEnumShape)
			if isEnumInDiagram && !ok {
				log.Fatalln("a gongenum should be associated to a gongenum shape")
			}
			_ = gongEnumShape

			isExpanded := IsNodeExpanded(classDiagram.NodeGongEnumNodeExpansionBinaryEncoding, idx)

			node := &tree.Node{
				Name:               gongEnum.Name,
				HasCheckboxButton:  true,
				IsChecked:          isEnumInDiagram,
				IsExpanded:         isExpanded,
				IsCheckboxDisabled: stager.embeddedDiagrams,
			}
			node.Impl = &GongEnumNodeProxy{
				node:          node,
				stager:        stager,
				classDiagram:  classDiagram,
				gongEnum:      gongEnum,
				gongEnumShape: gongEnumShape,
				rank:          idx,
			}
			nodeGongEnums.Children = append(nodeGongEnums.Children, node)

			for _, gongEnumValue := range gongEnum.GongEnumValues {

				_, isEnumValueInDiagram := map_modelElement_shape[gongEnumValue]
				nodeEnumValue := &tree.Node{
					Name:               gongEnumValue.Name,
					HasCheckboxButton:  true,
					IsChecked:          isEnumValueInDiagram,
					IsCheckboxDisabled: !isEnumInDiagram || stager.embeddedDiagrams,
				}
				nodeEnumValue.Impl = &GongEnumNodeValueProxy{
					stager:        stager,
					classDiagram:  classDiagram,
					gongEnumShape: gongEnumShape,
					gongEnum:      gongEnum,
					gongEnumValue: gongEnumValue,
				}
				node.Children = append(node.Children, nodeEnumValue)
			}
		}

		for idx, gongNote := range gongnotes {
			shape, isInDiagram := map_modelElement_shape[gongNote]

			gongNoteShape, ok := shape.(*GongNoteShape)
			if isInDiagram && !ok {
				log.Fatalln("a gongnote should be associated to a gongnote shape")
			}
			_ = gongNoteShape

			isExpanded := IsNodeExpanded(classDiagram.NodeGongNoteNodeExpansionBinaryEncoding, idx)

			gongNoteNode := &tree.Node{
				Name:               gongNote.Name,
				HasCheckboxButton:  true,
				IsChecked:          isInDiagram,
				IsExpanded:         isExpanded,
				IsCheckboxDisabled: stager.embeddedDiagrams,
			}
			gongNoteNode.Impl = &GongNoteNodeProxy{
				node:          gongNoteNode,
				stager:        stager,
				classDiagram:  classDiagram,
				gongNote:      gongNote,
				gongNoteShape: gongNoteShape,
				rank:          idx,
			}
			nodeGongNotes.Children = append(nodeGongNotes.Children, gongNoteNode)

			for _, docLink := range gongNote.Links {

				docLinkNode := &tree.Node{
					Name:              docLink.Name,
					HasCheckboxButton: true,
					IsChecked:         isInDiagram,
					IsExpanded:        isExpanded,
				}
				gongNoteNode.Children = append(gongNoteNode.Children, docLinkNode)
			}
		}
	}
	tree.StageBranch(stager.treeStage,
		classdiagramsTree,
	)

	stager.treeStage.Commit()
}

func (stager *Stager) addButtonsToClassdiagramNode(nodeClassdiagram *tree.Node, classDiagram *Classdiagram) {
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
}
