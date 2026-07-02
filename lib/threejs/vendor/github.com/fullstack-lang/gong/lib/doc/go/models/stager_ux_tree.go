package models

import (
	"fmt"
	"go/types"
	"log"

	gong "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	// Import for the nodestates package
	"github.com/fullstack-lang/gong/lib/doc/go/models/nodestates"
)

func (stager *Stager) tree() {
	stager.treeStage.Reset()

	classdiagramsTree := &tree.Tree{
		Name: stager.stage.GetProbeTreeSidebarStageName(),
	}

	if !stager.embeddedDiagrams {
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
				OnClick: func() {
					proxy := &ButtonNewClassdiagramProxy{
						stager: stager,
					}
					proxy.ButtonUpdated(nil, nil, nil)
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

		nodeClassdiagram.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked {
				// uncheck all other diagram
				diagramPackage := getTheDiagramPackage(stager.stage)
				diagramPackage.SelectedClassdiagram = classDiagram

				stager.stage.Commit()
			} else {
				diagramPackage := getTheDiagramPackage(stager.stage)
				diagramPackage.SelectedClassdiagram = nil

				stager.stage.Commit()
			}
		}

		nodeClassdiagram.OnIsExpandedChange = func(isExpanded bool) {
			classDiagram.IsExpanded = isExpanded
			stager.stage.Commit()
		}

		nodeClassdiagram.OnNameChange = func(newName string) {
			classDiagram.Name = newName
			classDiagram.IsInRenameMode = false

			stager.stage.Commit()
		}

		nodeClassdiagram.OnIsSecondCheckboxCheckedChanged = func(isChecked bool) {
			classDiagram.IsIncludedInStaticWebSite = isChecked
			stager.stage.Commit()
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
			OnIsExpandedChange: func(isExpanded bool) {
				classDiagram.NodeGongStructsIsExpanded = isExpanded
				stager.stage.Commit()
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
			OnIsExpandedChange: func(isExpanded bool) {
				classDiagram.NodeGongEnumsIsExpanded = isExpanded
				stager.stage.Commit()
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
			OnIsExpandedChange: func(isExpanded bool) {
				classDiagram.NodeGongNotesIsExpanded = isExpanded
				stager.stage.Commit()
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
				CheckboxHasToolTip: true,
				CheckboxToolTipText: func() string {
					if isGongStructShapeInDiagram {
						return "Remove from diagram"
					} else {
						return "Add to diagram"
					}
				}(),
				CheckboxToolTipPosition: tree.Above,
			}
			nodeNamedStruct.OnIsCheckedChanged = func(isChecked bool) {
				if isChecked {
					diagramPackage := getTheDiagramPackage(stager.stage)
					classDiagram.AddGongStructShape(stager.stage, diagramPackage, gongStruct.Name)
				} else {
					classDiagram.RemoveGongStructShape(stager.stage, gongStruct.Name)
				}
			}
			nodeNamedStruct.OnIsExpandedChange = func(isExpanded bool) {
				currentExpansionStateInDiagram := classDiagram.NodeGongStructNodeExpansion
				if nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, idx) == nil {
					classDiagram.NodeGongStructNodeExpansion = currentExpansionStateInDiagram
					stager.stage.Commit()
				}
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
						CheckboxHasToolTip: true,
						CheckboxToolTipText: func() string {
							if isInDiagram {
								return "Remove from diagram"
							} else {
								return "Add to diagram"
							}
						}(),
						CheckboxToolTipPosition: tree.Above,
						HasToolTip:              true,
						ToolTipText: func() string {
							switch field := field.(type) {
							case *gong.GongBasicField:
								switch field.GetBasicKind() {
								case types.Int, types.Int64, types.Int32, types.Float64:
									if field.GongEnum != nil {
										return fmt.Sprintf("%s (enum)", field.DeclaredType)
									} else {
										return fmt.Sprintf("%s", field.DeclaredType)
									}
								case types.String:
									if field.GongEnum != nil {
										return field.GongEnum.Name + " (enum)"
									} else {
										return "string"
									}
								case types.Bool:
									return "boolean"
								case types.UntypedNil:
									return "any type (used for referencing identifiers)"
								default:
									return ""
								}
							case *gong.GongTimeField:
								return "Type: time.Time"
							default:
								return ""
							}
						}(),
						ToolTipPosition: tree.Right,
					}

					nodeField.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							classDiagram.AddAttributeFieldShape(
								stager.stage,
								stager.gongStage,
								gongStruct,
								field,
								gongStructShape)
							stager.stage.Commit()
						} else {
							classDiagram.RemoveAttributeFieldShape(stager.stage, attributeShape, gongStructShape)
							stager.stage.Commit()
						}
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
						HasToolTip:         true,
						ToolTipText: func() string {
							switch field := field.(type) {
							case *gong.PointerToGongStructField:
								return fmt.Sprintf("*%s", field.GongStruct.Name)
							case *gong.SliceOfPointerToGongStructField:
								return fmt.Sprintf("[]*%s", field.GongStruct.Name)
							default:
								return ""
							}
						}(),
						ToolTipPosition: tree.Right,
					}

					nodeField.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							classDiagram.AddLinkShape(
								stager.stage,
								stager.gongStage,
								gongStruct,
								field,
								gongStructShape)
							stager.stage.Commit()
						} else {
							classDiagram.RemoveLinkFieldShape(stager.stage, linkShape, gongStructShape)
							stager.stage.Commit()
						}
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
			node.OnIsCheckedChanged = func(isChecked bool) {
				if isChecked {
					diagramPackage := getTheDiagramPackage(stager.stage)
					classDiagram.AddGongEnumShape(stager.stage, diagramPackage, gongEnum.Name)
					stager.stage.Commit()
				} else {
					classDiagram.RemoveGongEnumShape(stager.stage, gongEnum.Name)
					stager.stage.Commit()
				}
			}
			node.OnIsExpandedChange = func(isExpanded bool) {
				currentExpansionStateInDiagram := classDiagram.NodeGongEnumNodeExpansion
				if nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, idx) == nil {
					classDiagram.NodeGongEnumNodeExpansion = currentExpansionStateInDiagram
					stager.stage.Commit()
				}
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
				nodeEnumValue.OnIsCheckedChanged = func(isChecked bool) {
					if isChecked {
						classDiagram.AddGongEnumValueShapeToDiagram(
							stager.stage,
							gongEnumShape,
							gongEnum,
							gongEnumValue)
						stager.stage.Commit()
					} else {
						classDiagram.RemoveGongEnumValueShapeFromDiagram(
							stager.stage,
							gongEnumShape,
							gongEnumValue,
						)
						stager.stage.Commit()
					}
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
			gongNoteNode.OnIsCheckedChanged = func(isChecked bool) {
				if isChecked {
					diagramPackage := getTheDiagramPackage(stager.stage)
					classDiagram.AddGongNoteShape(stager.stage, gongNote, diagramPackage, gongNote.Name)
					stager.stage.Commit()
				} else {
					classDiagram.RemoveGongNoteShape(stager.stage, gongNote.Name)
					stager.stage.Commit()
				}
			}
			gongNoteNode.OnIsExpandedChange = func(isExpanded bool) {
				currentExpansionStateInDiagram := classDiagram.NodeGongNoteNodeExpansion
				if nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, idx) == nil {
					classDiagram.NodeGongNoteNodeExpansion = currentExpansionStateInDiagram
					stager.stage.Commit()
				}
			}
			nodeGongNotes.Children = append(nodeGongNotes.Children, gongNoteNode)

			for _, gongLink := range gongNote.Links {

				_, isGongLinkShapeInDiagram := map_modelElement_shape[gongLink] // Corrected variable name from original
				// gongNoteLinkShape, ok := shape.(*GongNoteLinkShape) // 'shape' not defined in this scope for link
				// if isGongLinkShapeInDiagram && !ok {
				// 	log.Fatalln("a gongnote link should be associated to a gongnote link shape")
				// }
				// _ = gongNoteLinkShape

				name := gongLink.Name
				if gongLink.Recv != "" {
					name = gongLink.Recv + "." + gongLink.Name
				}

				isTargetAbsent := true
				if gongLink.Recv == "" {
					// check if a GongStructShape, GongEnumShape, or GongNoteShape exists with that name
					for _, shape := range classDiagram.GongStructShapes {
						if IdentifierMetaToGongStructName(shape.IdentifierMeta) == gongLink.Name {
							isTargetAbsent = false
						}
					}
					for _, shape := range classDiagram.GongEnumShapes {
						if GongEnumIdentifierMetaToGongEnumName(shape.IdentifierMeta) == gongLink.Name {
							isTargetAbsent = false
						}
					}
					for _, shape := range classDiagram.GongNoteShapes {
						if IdentifierToGongStructName(shape.Identifier) == gongLink.Name {
							isTargetAbsent = false
						}
					}
				} else {
					// check if a LinkShape exists for the receiver
					for _, shape := range classDiagram.GongStructShapes {
						if IdentifierMetaToGongStructName(shape.IdentifierMeta) == gongLink.Recv {
							for _, linkShape := range shape.LinkShapes {
								if IdentifierMetaToFieldName(linkShape.IdentifierMeta) == gongLink.Name {
									isTargetAbsent = false
								}
							}
						}
					}
				}

				docLinkNode := &tree.Node{
					Name:               name,
					HasCheckboxButton:  true,
					IsChecked:          isGongLinkShapeInDiagram,
					IsExpanded:         noteIsExpanded, // Links inherit expansion from parent note
					IsCheckboxDisabled: !isGongNoteShapeInDiagram || isTargetAbsent || stager.embeddedDiagrams || !selected,
				}
				docLinkNode.OnIsCheckedChanged = func(isChecked bool) {
					if isChecked {
						classDiagram.AddGongNoteLinkShapeToDiagram(stager.stage, gongNoteShape, gongLink)
						stager.stage.Commit()
					} else {
						classDiagram.RemoveGongNoteLinkShapeFromDiagram(stager.stage, gongNoteShape, gongLink)
						stager.stage.Commit()
					}
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
			OnClick: func() {
				classDiagram.ShowNbInstances = !classDiagram.ShowNbInstances
				stager.stage.Commit()
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
			OnClick: func() {
				classDiagram.ShowMultiplicity = !classDiagram.ShowMultiplicity
				stager.stage.Commit()
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
			OnClick: func() {
				classDiagram.ShowLinkNames = !classDiagram.ShowLinkNames
				stager.stage.Commit()
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
			OnClick: func() {
				NewClassDiagramButtonProxy(
					stager,
					classDiagram,
					nodeClassdiagram,
					REMOVE,
				).ButtonUpdated(nil, nil, nil)
			},
			HasToolTip:      true,
			ToolTipText:     "Delete the diagram",
			ToolTipPosition: tree.Above,
		})

	if !classDiagram.IsInRenameMode {
		nodeClassdiagram.Buttons = append(nodeClassdiagram.Buttons,
			&tree.Button{
				Name: classDiagram.GetName() + " " + string(buttons.BUTTON_edit_note),
				Icon: string(buttons.BUTTON_edit_note),
				OnClick: func() {
					NewClassDiagramButtonProxy(
						stager,
						classDiagram,
						nodeClassdiagram,
						RENAME,
					).ButtonUpdated(nil, nil, nil)
				},
				HasToolTip:      true,
				ToolTipText:     "Rename the diagram",
				ToolTipPosition: tree.Above,
			})
	} else {
		nodeClassdiagram.Buttons = append(nodeClassdiagram.Buttons,
			&tree.Button{
				Name: classDiagram.GetName() + " " + string(buttons.BUTTON_edit_off),
				Icon: string(buttons.BUTTON_edit_off),
				OnClick: func() {
					NewClassDiagramButtonProxy(
						stager,
						classDiagram,
						nodeClassdiagram,
						RENAME_CANCEL,
					).ButtonUpdated(nil, nil, nil)
				},
				HasToolTip:      true,
				ToolTipText:     "Cancel renaming",
				ToolTipPosition: tree.Above,
			})
	}

	nodeClassdiagram.Buttons = append(nodeClassdiagram.Buttons,
		&tree.Button{
			Name: classDiagram.GetName() + " " + string(buttons.BUTTON_copy_all),
			Icon: string(buttons.BUTTON_copy_all),
			OnClick: func() {
				NewClassDiagramButtonProxy(
					stager,
					classDiagram,
					nodeClassdiagram,
					DUPLICATE,
				).ButtonUpdated(nil, nil, nil)
			},
			HasToolTip:      true,
			ToolTipText:     "Duplicate diagram",
			ToolTipPosition: tree.Right,
		})

	// add a second checkbox for including the diagram into
	nodeClassdiagram.HasSecondCheckboxButton = true
}
