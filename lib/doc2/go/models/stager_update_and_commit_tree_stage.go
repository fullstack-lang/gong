package models

import (
	"fmt"
	"log"

	gong "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	// Import for the nodestates package
	"github.com/fullstack-lang/gong/lib/doc2/go/models/nodestates"
)

func (stager *Stager) UpdateAndCommitTreeStage() {

	stager.treeStage.Reset()

	classdiagramsTree := &tree.Tree{
		Name: stager.stage.GetProbeTreeSidebarStageName(),
	}

	root := &tree.Node{
		Name:       "Class Diagrams",
		IsExpanded: true,
	}
	classdiagramsTree.RootNodes = append(classdiagramsTree.RootNodes, root)

	// if diagram can be edited
	// 1/ put a "add a class diagram" button
	// 2/ put a "generate sss" button
	if !stager.embeddedDiagrams {
		root.Buttons = append(root.Buttons,
			&tree.Button{
				Name: "Class Diagramm Add Button",
				Icon: string(buttons.BUTTON_add),
				Impl: &ButtonNewClassdiagramProxy{
					stager: stager,
				},
				HasToolTip:      true,
				ToolTipText:     "Create a new diagram",
				ToolTipPosition: tree.Above,
			},
			&tree.Button{
				Name: "Generates the documentation static web site",
				Icon: string(buttons.BUTTON_web_asset),
				Impl: &ButtonGeneratesStaticWebSiteProxy{
					stager: stager,
				},
				HasToolTip:      true,
				ToolTipText:     "Generates the documentation static web site",
				ToolTipPosition: tree.Above,
			},
		)
	}

	// append a node below for each diagram
	diagramPackage := getTheDiagramPackage(stager.stage)
	for _, classDiagram := range diagramPackage.Classdiagrams {
		var selected bool
		if diagramPackage.SelectedClassdiagram == classDiagram {
			selected = true
		}
		nodeClassdiagram := &tree.Node{
			Name: classDiagram.Name,

			IsChecked:           selected,
			CheckboxHasToolTip:  true,
			CheckboxToolTipText: "Select this diagram for display",

			IsExpanded: classDiagram.IsExpanded,

			IsInEditMode: classDiagram.IsInRenameMode,

			HasCheckboxButton:       true,
			IsSecondCheckboxChecked: classDiagram.IsIncludedInStaticWebSite,
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
		nodeGongStructs := &tree.Node{
			Name:       fmt.Sprintf("Gongstructs (%d/%d)", nbGongstructsInDiagram, len(gongstructs)),
			IsExpanded: classDiagram.NodeGongStructsIsExpanded, // This remains a boolean for the category node
			Impl: &ClassDiagramGongStructsNodeProxy{
				stager:       stager,
				classDiagram: classDiagram,
			},
		}
		nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodeGongStructs)

		gongenums := gong.GetGongstrucsSorted[*gong.GongEnum](stager.gongStage)
		nbGongenumsInDiagram := 0
		for _, gongEnumItem := range gongenums { // Renamed variable to avoid conflict
			_, isInDiagram := map_modelElement_shape[gongEnumItem]
			if isInDiagram {
				nbGongenumsInDiagram++
			}
		}
		nodeGongEnums := &tree.Node{
			Name:       fmt.Sprintf("Gongenums (%d/%d)", nbGongenumsInDiagram, len(gongenums)),
			IsExpanded: classDiagram.NodeGongEnumsIsExpanded, // This remains a boolean for the category node
			Impl: &ClassDiagramGongEnumsNodeProxy{
				stager:       stager,
				classDiagram: classDiagram,
			},
		}
		nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodeGongEnums)

		gongnotes := gong.GetGongstrucsSorted[*gong.GongNote](stager.gongStage)
		nbGongnotesInDiagram := 0
		for _, gongNoteItem := range gongnotes { // Renamed variable to avoid conflict
			_, isInDiagram := map_modelElement_shape[gongNoteItem]
			if isInDiagram {
				nbGongnotesInDiagram++
			}
		}
		nodeGongNotes := &tree.Node{
			Name:       fmt.Sprintf("Gongnotes (%d/%d)", nbGongnotesInDiagram, len(gongnotes)),
			IsExpanded: classDiagram.NodeGongNotesIsExpanded, // This remains a boolean for the category node
			Impl: &ClassDiagramGongNotesNodeProxy{
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
			// Refactored to use nodestates package and the new string field
			isExpanded, err := nodestates.IsNodeExpanded(classDiagram.NodeGongStructNodeExpansion, idx)
			if err != nil {
				log.Printf("Error checking expansion state for GongStruct %s (index %d): %v. Defaulting to not expanded.", gongStruct.Name, idx, err)
				isExpanded = false
			}

			nodeNamedStruct := &tree.Node{
				Name:               gongStruct.Name,
				HasCheckboxButton:  true,
				IsCheckboxDisabled: stager.embeddedDiagrams || !selected,
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
			nodeGongStructs.Children = append(nodeGongStructs.Children, nodeNamedStruct)

			for _, field := range gongStruct.Fields {

				switch field := field.(type) { // Shadowing field variable is idiomatic Go in type switches
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
						IsCheckboxDisabled: !isGongStructShapeInDiagram || stager.embeddedDiagrams || !selected,
					}

					nodeField.Impl = &AttributeFieldNodeProxy{
						node:            nodeNamedStruct, // Original code had nodeNamedStruct, implies parent
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
					if ptrField, ok := field.(*gong.PointerToGongStructField); ok {
						if _, ok2 := map_modelElement_shape[ptrField.GongStruct]; !ok2 {
							isTargetGongstructAbsent = true
						}
					}
					if sliceField, ok := field.(*gong.SliceOfPointerToGongStructField); ok {
						if _, ok2 := map_modelElement_shape[sliceField.GongStruct]; !ok2 {
							isTargetGongstructAbsent = true
						}
					}

					nodeField := &tree.Node{
						Name:               field.GetName(),
						HasCheckboxButton:  true,
						IsChecked:          isInDiagram,
						IsCheckboxDisabled: !isGongStructShapeInDiagram || isTargetGongstructAbsent || stager.embeddedDiagrams || !selected,
					}

					nodeField.Impl = &LinkFieldNodeProxy{
						node:            nodeNamedStruct, // Original code had nodeNamedStruct, implies parent
						stager:          stager,
						classDiagram:    classDiagram,
						gongstruct:      gongStruct,
						gongStructShape: gongStructShape,
						field:           field,
						linkShape:       linkShape,
					}

					nodeNamedStruct.Children = append(nodeNamedStruct.Children, nodeField)
				default:
					log.Printf("Unknown field type encountered: %T for field %s", field, field.GetName())
				}
			}
		}
		for idx, gongEnum := range gongenums {
			shape, isEnumInDiagram := map_modelElement_shape[gongEnum]

			gongEnumShape, ok := shape.(*GongEnumShape)
			if isEnumInDiagram && !ok {
				log.Fatalln("a gongenum should be associated to a gongenum shape")
			}
			_ = gongEnumShape // To avoid unused variable error if not used otherwise later

			// Refactored to use nodestates package and the new string field
			isExpanded, err := nodestates.IsNodeExpanded(classDiagram.NodeGongEnumNodeExpansion, idx)
			if err != nil {
				log.Printf("Error checking expansion state for GongEnum %s (index %d): %v. Defaulting to not expanded.", gongEnum.Name, idx, err)
				isExpanded = false
			}

			node := &tree.Node{
				Name:               gongEnum.Name,
				HasCheckboxButton:  true,
				IsChecked:          isEnumInDiagram,
				IsExpanded:         isExpanded,
				IsCheckboxDisabled: stager.embeddedDiagrams || !selected,
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
					IsCheckboxDisabled: !isEnumInDiagram || stager.embeddedDiagrams || !selected,
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

		var noteIsExpanded bool // Variable to store current note's expansion for its links
		for idx, gongNote := range gongnotes {
			shape, isGongNoteShapeInDiagram := map_modelElement_shape[gongNote]

			gongNoteShape, ok := shape.(*GongNoteShape)
			if isGongNoteShapeInDiagram && !ok {
				log.Fatalln("a gongnote should be associated to a gongnote shape")
			}
			_ = gongNoteShape // To avoid unused variable error if not used otherwise later

			// Refactored to use nodestates package and the new string field
			var err error
			noteIsExpanded, err = nodestates.IsNodeExpanded(classDiagram.NodeGongNoteNodeExpansion, idx)
			if err != nil {
				log.Printf("Error checking expansion state for GongNote %s (index %d): %v. Defaulting to not expanded.", gongNote.Name, idx, err)
				noteIsExpanded = false
			}

			gongNoteNode := &tree.Node{
				Name:               gongNote.Name,
				HasCheckboxButton:  true,
				IsChecked:          isGongNoteShapeInDiagram,
				IsExpanded:         noteIsExpanded,
				IsCheckboxDisabled: stager.embeddedDiagrams || !selected,
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

			for _, gongLink := range gongNote.Links {

				_, isGongLinkShapeInDiagram := map_modelElement_shape[gongLink] // Corrected variable name from original
				// gongNoteLinkShape, ok := shape.(*GongNoteLinkShape) // 'shape' not defined in this scope for link
				// if isGongLinkShapeInDiagram && !ok {
				// 	log.Fatalln("a gongnote link should be associated to a gongnote link shape")
				// }
				// _ = gongNoteLinkShape

				docLinkNode := &tree.Node{
					Name:              gongLink.Name,
					HasCheckboxButton: true,
					IsChecked:         isGongLinkShapeInDiagram,
					IsExpanded:        noteIsExpanded, // Links inherit expansion from parent note
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

	// add a second checkbox for including the diagram into
	nodeClassdiagram.HasSecondCheckboxButton = true

}
