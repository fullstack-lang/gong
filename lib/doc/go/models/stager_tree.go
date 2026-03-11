package models

import (
	"fmt"
	"log"

	gong "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	// Import for the nodestates package
	"github.com/fullstack-lang/gong/lib/doc/go/models/nodestates"
)

func (stager *Stager) tree() {
	log.Println("dod.stager.tree()")

	stager.treeStage.Reset()

	classdiagramsTree := &tree.Tree{
		Name: stager.stage.GetProbeTreeSidebarStageName(),
	}

	root := &tree.Node{
		Name:       "Class Diagrams",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}
	classdiagramsTree.RootNodes = append(classdiagramsTree.RootNodes, root)

	if stager.embeddedDiagrams {
		root.Name += " (embbeded in executable, for consultation only)"
	}

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
		)
	}

	// append a node below for each diagram
	diagramPackage := getTheDiagramPackage(stager.stage)

	classDiagrams := GetGongstrucsSorted[*Classdiagram](stager.stage)

	for _, classDiagram := range classDiagrams {
		var selected bool
		if diagramPackage.SelectedClassdiagram == classDiagram {
			selected = true
		}
		nodeClassdiagram := &tree.Node{
			Name: classDiagram.Name,

			IsChecked:               selected,
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Above,

			IsExpanded: classDiagram.IsExpanded,

			IsInEditMode: classDiagram.IsInRenameMode,

			HasCheckboxButton:             true,
			IsSecondCheckboxChecked:       classDiagram.IsIncludedInStaticWebSite,
			SecondCheckboxHasToolTip:      true,
			SecondCheckboxToolTipPosition: tree.Above,
		}

		if selected {
			nodeClassdiagram.CheckboxToolTipText = "Hide diagram"
		} else {
			nodeClassdiagram.CheckboxToolTipText = "Show diagram"
		}

		if classDiagram.IsIncludedInStaticWebSite {
			nodeClassdiagram.SecondCheckboxToolTipText = "Remove from documentation"
		} else {
			nodeClassdiagram.SecondCheckboxToolTipText = "Add to documentation"
		}

		nodeClassdiagram.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: func(stage *tree.Stage, staged, front *tree.Node) {
				// intercept update to the node that are when the node is checked
				if front.IsChecked && !staged.IsChecked {
					// uncheck all other diagram
					diagramPackage := getTheDiagramPackage(stager.stage)
					diagramPackage.SelectedClassdiagram = classDiagram

					stager.stage.Commit()
				}

				// the checked node is unchecked
				if !front.IsChecked && staged.IsChecked {
					diagramPackage := getTheDiagramPackage(stager.stage)
					diagramPackage.SelectedClassdiagram = nil

					stager.stage.Commit()
				}

				if front.IsExpanded && !staged.IsExpanded {
					classDiagram.IsExpanded = true
					stager.stage.Commit()
				}
				if !front.IsExpanded && staged.IsExpanded {
					classDiagram.IsExpanded = false
					stager.stage.Commit()
				}

				if front.Name != staged.Name {
					classDiagram.Name = front.Name
					classDiagram.IsInRenameMode = false

					stager.stage.Commit()
				}

				// second checkbox
				if front.IsSecondCheckboxChecked != staged.IsSecondCheckboxChecked {

					staged.IsSecondCheckboxChecked = front.IsSecondCheckboxChecked
					classDiagram.IsIncludedInStaticWebSite = front.IsSecondCheckboxChecked
					stager.stage.Commit()
				}
			},
		}

		if !stager.embeddedDiagrams {
			stager.addDeleteRenameCopyButtons(nodeClassdiagram, classDiagram)
			stager.addDisplayModeButtons(classDiagram, nodeClassdiagram)
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
			Impl: &tree.FunctionalNodeProxy{
				OnUpdate: func(stage *tree.Stage, staged, front *tree.Node) {
					if front.IsExpanded && !staged.IsExpanded {
						classDiagram.NodeGongStructsIsExpanded = true
						stager.stage.Commit()
					}
					if !front.IsExpanded && staged.IsExpanded {
						classDiagram.NodeGongStructsIsExpanded = false
						stager.stage.Commit()
					}
				},
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
			Impl: &tree.FunctionalNodeProxy{
				OnUpdate: func(stage *tree.Stage, staged, front *tree.Node) {
					if front.IsExpanded && !staged.IsExpanded {
						classDiagram.NodeGongEnumsIsExpanded = true
						stager.stage.Commit()
					}
					if !front.IsExpanded && staged.IsExpanded {
						classDiagram.NodeGongEnumsIsExpanded = false
						stager.stage.Commit()
					}
				},
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
			Impl: &tree.FunctionalNodeProxy{
				OnUpdate: func(stage *tree.Stage, staged, front *tree.Node) {
					if front.IsExpanded && !staged.IsExpanded {
						classDiagram.NodeGongNotesIsExpanded = true
						stager.stage.Commit()
					}
					if !front.IsExpanded && staged.IsExpanded {
						classDiagram.NodeGongNotesIsExpanded = false
						stager.stage.Commit()
					}
				},
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
			nodeNamedStruct.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: func(stage *tree.Stage, staged, front *tree.Node) {

					// intercept update to the node that are when the node is checked
					if front.IsChecked && !staged.IsChecked {
						diagramPackage := getTheDiagramPackage(stager.stage)
						classDiagram.AddGongStructShape(stager.stage, diagramPackage, gongStruct.Name)
						return
					}

					// the checked node is unchecked
					if !front.IsChecked && staged.IsChecked {
						classDiagram.RemoveGongStructShape(stager.stage, gongStruct.Name)
						return
					}

					expansionToggled := false
					currentExpansionStateInDiagram := classDiagram.NodeGongStructNodeExpansion

					if front.IsExpanded && !staged.IsExpanded {
						if nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, idx) == nil {
							classDiagram.NodeGongStructNodeExpansion = currentExpansionStateInDiagram
							expansionToggled = true
						}
					}

					if !front.IsExpanded && staged.IsExpanded {
						if nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, idx) == nil {
							classDiagram.NodeGongStructNodeExpansion = currentExpansionStateInDiagram
							expansionToggled = true
						}
					}

					if expansionToggled {
						stager.stage.Commit()
					}
				},
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

					nodeField.Impl = &tree.FunctionalNodeProxy{
						OnUpdate: func(stage *tree.Stage, staged, front *tree.Node) {
							// intercept update to the node that are when the node is checked
							if front.IsChecked && !staged.IsChecked {
								// uncheck all other diagram
								classDiagram.AddAttributeFieldShape(
									stager.stage,
									stager.gongStage,
									gongStruct,
									field,
									gongStructShape)
								stager.stage.Commit()
							}
							// the checked node is unchecked
							if !front.IsChecked && staged.IsChecked {
								classDiagram.RemoveAttributeFieldShape(stager.stage, attributeShape, gongStructShape)
								stager.stage.Commit()
							}
						},
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

					nodeField.Impl = &tree.FunctionalNodeProxy{
						OnUpdate: func(stage *tree.Stage, staged, front *tree.Node) {
							// intercept update to the node that are when the node is checked
							if front.IsChecked && !staged.IsChecked {
								// uncheck all other diagram
								classDiagram.AddLinkShape(
									stager.stage,
									stager.gongStage,
									gongStruct,
									field,
									gongStructShape)
								stager.stage.Commit()
							}
							// the checked node is unchecked
							if !front.IsChecked && staged.IsChecked {
								classDiagram.RemoveLinkFieldShape(stager.stage, linkShape, gongStructShape)
								stager.stage.Commit()
							}
						},
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
			node.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: func(stage *tree.Stage, staged, front *tree.Node) {

					// intercept update to the node that are when the node is checked
					if front.IsChecked && !staged.IsChecked {
						diagramPackage := getTheDiagramPackage(stager.stage)
						classDiagram.AddGongEnumShape(stager.stage, diagramPackage, gongEnum.Name)

						stager.stage.Commit()
					}

					// the checked node is unchecked
					if !front.IsChecked && staged.IsChecked {
						classDiagram.RemoveGongEnumShape(stager.stage, gongEnum.Name)
						stager.stage.Commit()
					}

					expansionToggled := false
					currentExpansionStateInDiagram := classDiagram.NodeGongEnumNodeExpansion

					if front.IsExpanded && !staged.IsExpanded {
						if nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, idx) == nil {
							classDiagram.NodeGongEnumNodeExpansion = currentExpansionStateInDiagram
							expansionToggled = true
						}
					}

					if !front.IsExpanded && staged.IsExpanded {
						if nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, idx) == nil {
							classDiagram.NodeGongEnumNodeExpansion = currentExpansionStateInDiagram
							expansionToggled = true
						}
					}

					if expansionToggled {
						stager.stage.Commit()
					}
				},
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
				nodeEnumValue.Impl = &tree.FunctionalNodeProxy{
					OnUpdate: func(stage *tree.Stage, staged, front *tree.Node) {
						// intercept update to the node that are when the node is checked
						if front.IsChecked && !staged.IsChecked {
							classDiagram.AddGongEnumValueShapeToDiagram(
								stager.stage,
								gongEnumShape,
								gongEnum,
								gongEnumValue)
							stager.stage.Commit()
						}
						// the checked node is unchecked
						if !front.IsChecked && staged.IsChecked {
							classDiagram.RemoveGongEnumValueShapeFromDiagram(
								stager.stage,
								gongEnumShape,
								gongEnumValue,
							)
							stager.stage.Commit()
						}
					},
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
			gongNoteNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: func(stage *tree.Stage, staged, front *tree.Node) {

					// intercept update to the node that are when the node is checked
					if front.IsChecked && !staged.IsChecked {
						diagramPackage := getTheDiagramPackage(stager.stage)
						classDiagram.AddGongNoteShape(stager.stage, gongNote, diagramPackage, gongNote.Name)
						stager.stage.Commit()
					}

					// the checked node is unchecked
					if !front.IsChecked && staged.IsChecked {
						classDiagram.RemoveGongNoteShape(stager.stage, gongNote.Name)
						stager.stage.Commit()
					}

					// expansionToggled tracks if the user action was an expansion or collapse
					expansionToggled := false
					currentExpansionStateInDiagram := classDiagram.NodeGongNoteNodeExpansion // For readability

					// User expanded the node
					if front.IsExpanded && !staged.IsExpanded {
						if nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, idx) == nil {
							classDiagram.NodeGongNoteNodeExpansion = currentExpansionStateInDiagram
							expansionToggled = true
						}
					}

					// User collapsed the node
					if !front.IsExpanded && staged.IsExpanded {
						if nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, idx) == nil {
							classDiagram.NodeGongNoteNodeExpansion = currentExpansionStateInDiagram
							expansionToggled = true
						}
					}

					if expansionToggled {
						stager.stage.Commit() // Commit changes to the classDiagram model
					}
				},
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

func (stager *Stager) addDisplayModeButtons(classDiagram *Classdiagram, nodeClassdiagram *tree.Node) {
	{
		button := &tree.Button{
			Name: "Show/Unshow number of instances",
			Impl: &toggleButtonProxy{
				stager:      stager,
				toggleValue: &classDiagram.ShowNbInstances,
			},
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
		}

		if !classDiagram.ShowNbInstances {
			button.ToolTipText = "Show nb of instances"
			button.Icon = string(buttons.BUTTON_visibility)
		} else {
			button.ToolTipText = "Hide nb of instances"
			button.Icon = string(buttons.BUTTON_visibility_off)
		}

		nodeClassdiagram.Buttons = append(nodeClassdiagram.Buttons, button)
	}

	{
		button := &tree.Button{
			Name: "Show/Unshow multiplicity",
			Impl: &toggleButtonProxy{
				stager:      stager,
				toggleValue: &classDiagram.ShowMultiplicity,
			},
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
		}

		if !classDiagram.ShowMultiplicity {
			button.ToolTipText = "Show multiplicity"
			button.Icon = string(buttons.BUTTON_visibility)
		} else {
			button.ToolTipText = "Hide multiplicity"
			button.Icon = string(buttons.BUTTON_visibility_off)
		}

		nodeClassdiagram.Buttons = append(nodeClassdiagram.Buttons, button)
	}

	{
		button := &tree.Button{
			Name: "Show/Unshow Link Names",
			Impl: &toggleButtonProxy{
				stager:      stager,
				toggleValue: &classDiagram.ShowLinkNames,
			},
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
		}

		if !classDiagram.ShowLinkNames {
			button.ToolTipText = "Show Link Names"
			button.Icon = string(buttons.BUTTON_visibility)
		} else {
			button.ToolTipText = "Hide Link Names"
			button.Icon = string(buttons.BUTTON_visibility_off)
		}

		nodeClassdiagram.Buttons = append(nodeClassdiagram.Buttons, button)
	}
}

func (stager *Stager) addDeleteRenameCopyButtons(nodeClassdiagram *tree.Node, classDiagram *Classdiagram) {

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
			ToolTipPosition: tree.Right,
		})

	// add a second checkbox for including the diagram into
	nodeClassdiagram.HasSecondCheckboxButton = true

}
