package models

import (
	"fmt"
	"go/types"
	"log"
	"sort"
	"strings"

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

	classDiagrams := stager.stage.GetInstancesSorted[*Classdiagram]()

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

		gongstructs := stager.gongStage.GetInstancesSorted[*gong.GongStruct]()
		structToGlobalIdx := make(map[*gong.GongStruct]int, len(gongstructs))
		for i, gs := range gongstructs {
			structToGlobalIdx[gs] = i
		}

		gongenums := stager.gongStage.GetInstancesSorted[*gong.GongEnum]()
		enumToGlobalIdx := make(map[*gong.GongEnum]int, len(gongenums))
		for i, ge := range gongenums {
			enumToGlobalIdx[ge] = i
		}

		gongnotes := stager.gongStage.GetInstancesSorted[*gong.GongNote]()
		noteToGlobalIdx := make(map[*gong.GongNote]int, len(gongnotes))
		for i, gn := range gongnotes {
			noteToGlobalIdx[gn] = i
		}

		if len(stager.stage.MetaPackageImports) == 0 {
			nbGongstructsInDiagram := 0
			for _, gongStruct := range gongstructs {
				if _, isInDiagram := map_modelElement_shape[gongStruct]; isInDiagram {
					nbGongstructsInDiagram++
				}
			}
			nodeGongStructs := &tree.Node{
				Name:       fmt.Sprintf("Gongstructs (%d/%d)", nbGongstructsInDiagram, len(gongstructs)),
				IsExpanded: classDiagram.NodeGongStructsIsExpanded,
				OnIsExpandedChange: func(isExpanded bool) {
					classDiagram.NodeGongStructsIsExpanded = isExpanded
					stager.stage.Commit()
				},
			}
			nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodeGongStructs)

			nbGongenumsInDiagram := 0
			for _, gongEnumItem := range gongenums {
				if _, isInDiagram := map_modelElement_shape[gongEnumItem]; isInDiagram {
					nbGongenumsInDiagram++
				}
			}
			nodeGongEnums := &tree.Node{
				Name:       fmt.Sprintf("Gongenums (%d/%d)", nbGongenumsInDiagram, len(gongenums)),
				IsExpanded: classDiagram.NodeGongEnumsIsExpanded,
				OnIsExpandedChange: func(isExpanded bool) {
					classDiagram.NodeGongEnumsIsExpanded = isExpanded
					stager.stage.Commit()
				},
			}
			nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodeGongEnums)

			nbGongnotesInDiagram := 0
			for _, gongNoteItem := range gongnotes {
				if _, isInDiagram := map_modelElement_shape[gongNoteItem]; isInDiagram {
					nbGongnotesInDiagram++
				}
			}
			nodeGongNotes := &tree.Node{
				Name:       fmt.Sprintf("Gongnotes (%d/%d)", nbGongnotesInDiagram, len(gongnotes)),
				IsExpanded: classDiagram.NodeGongNotesIsExpanded,
				OnIsExpandedChange: func(isExpanded bool) {
					classDiagram.NodeGongNotesIsExpanded = isExpanded
					stager.stage.Commit()
				},
			}
			nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodeGongNotes)

			for idx, gongStruct := range gongstructs {
				pkgName := "models"
				if gongStruct.ModelPkg != nil && gongStruct.ModelPkg.PkgGoName != "" {
					pkgName = gongStruct.ModelPkg.PkgGoName
				}
				nodeGongStructs.Children = append(nodeGongStructs.Children, stager.createStructNode(classDiagram, gongStruct, idx, pkgName, selected, map_modelElement_shape))
			}
			for idx, gongEnum := range gongenums {
				nodeGongEnums.Children = append(nodeGongEnums.Children, stager.createEnumNode(classDiagram, gongEnum, idx, "models", selected, map_modelElement_shape))
			}
			for idx, gongNote := range gongnotes {
				nodeGongNotes.Children = append(nodeGongNotes.Children, stager.createNoteNode(classDiagram, gongNote, idx, selected, map_modelElement_shape))
			}
		} else {
			// Separate root package (ref_models) and subpackages
			var rootImports []*MetaPackageImport
			var subPkgImports []*MetaPackageImport
			for _, imp := range stager.stage.MetaPackageImports {
				if imp.Alias == "ref_models" {
					rootImports = append(rootImports, imp)
				} else {
					subPkgImports = append(subPkgImports, imp)
				}
			}

			// Sort subPkgImports so ancestor packages appear before descendant packages
			sort.SliceStable(subPkgImports, func(i, j int) bool {
				pathI := strings.Trim(subPkgImports[i].Path, "\"")
				pathJ := strings.Trim(subPkgImports[j].Path, "\"")
				if strings.HasPrefix(pathJ, pathI+"/") {
					return true
				}
				if strings.HasPrefix(pathI, pathJ+"/") {
					return false
				}
				return false
			})

			allImports := append(rootImports, subPkgImports...)
			pkgNodeMap := make(map[*MetaPackageImport]*tree.Node)

			for _, imp := range allImports {
				cleanPath := strings.Trim(imp.Path, "\"")
				pkgName := strings.TrimPrefix(imp.Alias, "ref_")

				var matchedModelPkg *gong.ModelPkg
				for mPkg := range *stager.gongStage.GetInstancesSet[*gong.ModelPkg]() {
					if mPkg.PkgPath == cleanPath {
						matchedModelPkg = mPkg
						break
					}
				}
				if matchedModelPkg == nil {
					for mPkg := range *stager.gongStage.GetInstancesSet[*gong.ModelPkg]() {
						if mPkg.PkgGoName == pkgName {
							matchedModelPkg = mPkg
							break
						}
					}
				}

				var pkgStructs []*gong.GongStruct
				for _, gs := range gongstructs {
					if (gs.ModelPkg != nil && gs.ModelPkg.PkgPath != "" && gs.ModelPkg.PkgPath == cleanPath) ||
						(gs.ModelPkg != nil && gs.ModelPkg == matchedModelPkg) {
						pkgStructs = append(pkgStructs, gs)
					}
				}
				if len(pkgStructs) == 0 && matchedModelPkg != nil {
					for _, gs := range gongstructs {
						for _, mgs := range matchedModelPkg.GongStructs {
							if mgs == gs {
								pkgStructs = append(pkgStructs, gs)
								break
							}
						}
					}
				}

				var pkgEnums []*gong.GongEnum
				if matchedModelPkg != nil {
					for _, ge := range gongenums {
						for _, mge := range matchedModelPkg.GongEnums {
							if mge == ge {
								pkgEnums = append(pkgEnums, ge)
								break
							}
						}
					}
				}

				var pkgNotes []*gong.GongNote
				if matchedModelPkg != nil {
					for _, gn := range gongnotes {
						for _, mgn := range matchedModelPkg.GongNotes {
							if mgn == gn {
								pkgNotes = append(pkgNotes, gn)
								break
							}
						}
					}
				}

				isRootPackage := (imp.Alias == "ref_models")

				var parentNode *tree.Node
				if isRootPackage {
					// Root package elements are direct children of the diagram node (no "models" wrapper node)
					parentNode = nodeClassdiagram
				} else {
					// Find closest ancestor package among subPkgImports
					var parentImp *MetaPackageImport
					longestPrefix := ""
					for _, other := range subPkgImports {
						if other == imp {
							continue
						}
						otherCleanPath := strings.Trim(other.Path, "\"")
						prefix := otherCleanPath + "/"
						if strings.HasPrefix(cleanPath, prefix) && len(prefix) > len(longestPrefix) {
							parentImp = other
							longestPrefix = prefix
						}
					}

					nodePkg := &tree.Node{
						Name:       pkgName,
						IsExpanded: true,
					}
					nodePkg.OnIsExpandedChange = func(isExpanded bool) {
						nodePkg.IsExpanded = isExpanded
					}
					pkgNodeMap[imp] = nodePkg

					if parentImp != nil && pkgNodeMap[parentImp] != nil {
						pkgNodeMap[parentImp].Children = append(pkgNodeMap[parentImp].Children, nodePkg)
					} else {
						nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodePkg)
					}
					parentNode = nodePkg
				}

				// 1. Gongstructs
				nbPkgGongstructsInDiagram := 0
				for _, gs := range pkgStructs {
					if _, isInDiagram := map_modelElement_shape[gs]; isInDiagram {
						nbPkgGongstructsInDiagram++
					}
				}
				nodePkgGongStructs := &tree.Node{
					Name:       fmt.Sprintf("Gongstructs (%d/%d)", nbPkgGongstructsInDiagram, len(pkgStructs)),
					IsExpanded: classDiagram.NodeGongStructsIsExpanded,
					OnIsExpandedChange: func(isExpanded bool) {
						classDiagram.NodeGongStructsIsExpanded = isExpanded
						stager.stage.Commit()
					},
				}
				parentNode.Children = append(parentNode.Children, nodePkgGongStructs)
				for _, gs := range pkgStructs {
					nodePkgGongStructs.Children = append(nodePkgGongStructs.Children, stager.createStructNode(classDiagram, gs, structToGlobalIdx[gs], pkgName, selected, map_modelElement_shape))
				}

				// 2. Gongenums (if any)
				if len(pkgEnums) > 0 {
					nbPkgGongenumsInDiagram := 0
					for _, ge := range pkgEnums {
						if _, isInDiagram := map_modelElement_shape[ge]; isInDiagram {
							nbPkgGongenumsInDiagram++
						}
					}
					nodePkgGongEnums := &tree.Node{
						Name:       fmt.Sprintf("Gongenums (%d/%d)", nbPkgGongenumsInDiagram, len(pkgEnums)),
						IsExpanded: classDiagram.NodeGongEnumsIsExpanded,
						OnIsExpandedChange: func(isExpanded bool) {
							classDiagram.NodeGongEnumsIsExpanded = isExpanded
							stager.stage.Commit()
						},
					}
					parentNode.Children = append(parentNode.Children, nodePkgGongEnums)
					for _, ge := range pkgEnums {
						nodePkgGongEnums.Children = append(nodePkgGongEnums.Children, stager.createEnumNode(classDiagram, ge, enumToGlobalIdx[ge], pkgName, selected, map_modelElement_shape))
					}
				}

				// 3. Gongnotes (if any)
				if len(pkgNotes) > 0 {
					nbPkgGongnotesInDiagram := 0
					for _, gn := range pkgNotes {
						if _, isInDiagram := map_modelElement_shape[gn]; isInDiagram {
							nbPkgGongnotesInDiagram++
						}
					}
					nodePkgGongNotes := &tree.Node{
						Name:       fmt.Sprintf("Gongnotes (%d/%d)", nbPkgGongnotesInDiagram, len(pkgNotes)),
						IsExpanded: classDiagram.NodeGongNotesIsExpanded,
						OnIsExpandedChange: func(isExpanded bool) {
							classDiagram.NodeGongNotesIsExpanded = isExpanded
							stager.stage.Commit()
						},
					}
					parentNode.Children = append(parentNode.Children, nodePkgGongNotes)
					for _, gn := range pkgNotes {
						nodePkgGongNotes.Children = append(nodePkgGongNotes.Children, stager.createNoteNode(classDiagram, gn, noteToGlobalIdx[gn], selected, map_modelElement_shape))
					}
				}
			}
		}
	}
	stager.treeStage.StageBranch(
		classdiagramsTree,
	)

	stager.treeStage.Commit()
}

func (stager *Stager) createStructNode(
	classDiagram *Classdiagram,
	gongStruct *gong.GongStruct,
	idx int,
	pkgName string,
	selected bool,
	map_modelElement_shape map[ModelElement]Shape,
) *tree.Node {
	shape, isGongStructShapeInDiagram := map_modelElement_shape[gongStruct]

	gongStructShape, ok := shape.(*GongStructShape)
	if isGongStructShapeInDiagram && !ok {
		log.Fatalln("A gongstruct shape should be mapped to a gongstruct")
	}
	isExpanded, err := nodestates.IsNodeExpanded(classDiagram.NodeGongStructNodeExpansion, idx)
	if err != nil {
		log.Printf("Error checking expansion state for GongStruct %s (index %d): %v. Defaulting to not expanded.", gongStruct.Name, idx, err)
		isExpanded = false
	}

	nodeName := gongStruct.Name
	if len(stager.stage.MetaPackageImports) == 0 && pkgName != "" && pkgName != "models" {
		nodeName = gongStruct.Name + " (" + pkgName + ")"
	}

	nodeNamedStruct := &tree.Node{
		Name:               nodeName,
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
			classDiagram.AddGongStructShapeWithPackage(stager.stage, diagramPackage, pkgName, gongStruct.Name)
		} else {
			classDiagram.RemoveGongStructShapeWithPackage(stager.stage, pkgName, gongStruct.Name)
		}
	}
	nodeNamedStruct.OnIsExpandedChange = func(isExpanded bool) {
		currentExpansionStateInDiagram := classDiagram.NodeGongStructNodeExpansion
		if nodestates.ToggleNodeExpanded(&currentExpansionStateInDiagram, idx) == nil {
			classDiagram.NodeGongStructNodeExpansion = currentExpansionStateInDiagram
			stager.stage.Commit()
		}
	}

	for _, field := range gongStruct.Fields {
		switch field := field.(type) {
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

	return nodeNamedStruct
}

func (stager *Stager) createEnumNode(
	classDiagram *Classdiagram,
	gongEnum *gong.GongEnum,
	idx int,
	pkgName string,
	selected bool,
	map_modelElement_shape map[ModelElement]Shape,
) *tree.Node {
	shape, isEnumInDiagram := map_modelElement_shape[gongEnum]

	gongEnumShape, ok := shape.(*GongEnumShape)
	if isEnumInDiagram && !ok {
		log.Fatalln("a gongenum should be associated to a gongenum shape")
	}
	_ = gongEnumShape

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
			classDiagram.AddGongEnumShapeWithPackage(stager.stage, diagramPackage, pkgName, gongEnum.Name)
			stager.stage.Commit()
		} else {
			classDiagram.RemoveGongEnumShapeWithPackage(stager.stage, pkgName, gongEnum.Name)
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

	return node
}

func (stager *Stager) createNoteNode(
	classDiagram *Classdiagram,
	gongNote *gong.GongNote,
	idx int,
	selected bool,
	map_modelElement_shape map[ModelElement]Shape,
) *tree.Node {
	shape, isGongNoteShapeInDiagram := map_modelElement_shape[gongNote]

	gongNoteShape, ok := shape.(*GongNoteShape)
	if isGongNoteShapeInDiagram && !ok {
		log.Fatalln("a gongnote should be associated to a gongnote shape")
	}
	_ = gongNoteShape

	noteIsExpanded, err := nodestates.IsNodeExpanded(classDiagram.NodeGongNoteNodeExpansion, idx)
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

	for _, gongLink := range gongNote.Links {
		_, isGongLinkShapeInDiagram := map_modelElement_shape[gongLink]

		name := gongLink.Name
		if gongLink.Recv != "" {
			name = gongLink.Recv + "." + gongLink.Name
		}

		isTargetAbsent := true
		if gongLink.Recv == "" {
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
			IsExpanded:         noteIsExpanded,
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

	return gongNoteNode
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
