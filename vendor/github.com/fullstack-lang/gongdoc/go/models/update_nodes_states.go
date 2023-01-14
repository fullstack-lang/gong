package models

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// updateNodesStates updates the tree of symbols
// according to the selected diagram
func updateNodesStates(stage *StageStruct, nodesCb *NodeCallbacksSingloton) {

	// get the diagram package of interest
	var pkglet *DiagramPackage
	for _pkgelt := range *GetGongstructInstancesSet[DiagramPackage]() {
		pkglet = _pkgelt
	}
	nodesCb.idTree.UncheckAndDisable()

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range nodesCb.ClassdiagramsRootNode.Children {

		classdiagramNode.HasEditButton = false
		classdiagramNode.HasDeleteButton = false
		classdiagramNode.HasDrawButton = false
		classdiagramNode.HasDrawOffButton = false

		if !classdiagramNode.IsChecked {
			classdiagramNode.IsInEditMode = false
			classdiagramNode.IsInDrawMode = false
			continue
		}

		editable := pkglet.IsEditable && !classdiagramNode.IsInEditMode && !classdiagramNode.IsInDrawMode

		classdiagramNode.HasEditButton = editable
		classdiagramNode.HasDeleteButton = editable
		classdiagramNode.HasDrawButton = editable

		// get the diagram
		classDiagram := classdiagramNode.Classdiagram

		// 1. allow all gongstructs node to be checked/unckecked
		for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
			nodesCb.idTree.nodeMap[gongStruct.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode
		}

		//
		// Enum
		//
		// 1. for each enum of the model, enable the check button if the class diagram is in draw mode
		for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {
			nodesCb.idTree.nodeMap[gongEnum.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode
		}

		// 2. get the all classshape and check the node of the
		// referenced gongstructs if it is referenced by the classshape
		for _, classshape := range classDiagram.Classshapes {
			ref_GongStruct := classshape.Reference
			node, ok := nodesCb.idTree.nodeMap[ref_GongStruct.Name]

			if !ok {
				log.Println("Unknown node, diagram not synchro with model ", ref_GongStruct.Name)
				continue
			}
			node.IsChecked = true

			// disable checkbox of all children of the gongstruct
			for _, gongfieldNode := range node.Children {
				gongfieldNode.IsCheckboxDisabled = !classDiagram.IsInDrawMode
			}

			for _, field := range classshape.Fields {
				nodeId := ref_GongStruct.Name + "." + field.Fieldname
				// node, ok := map_FieldId_Node[nodeId]
				node, _ := nodesCb.idTree.nodeMap[nodeId]

				node.IsChecked = true
				node.IsCheckboxDisabled = !classDiagram.IsInDrawMode

			}
			for _, link := range classshape.Links {
				nodeId := ref_GongStruct.Name + "." + link.Fieldname
				node, _ := nodesCb.idTree.nodeMap[nodeId]

				node.IsChecked = true
				node.IsCheckboxDisabled = !classDiagram.IsInDrawMode
			}
		}

		// disable checkbox field node that reference a gongstruct whose classshape is
		// not present in the diagram
		// first, construct map of all gongstructs present in the diagram
		map_OfGongstruct := make(map[string]bool)
		for _, classshape := range classDiagram.Classshapes {
			map_OfGongstruct[classshape.ReferenceName] = true
		}
		// then iterate over all fields of all gongstructs node
		for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
			gognstructNode := nodesCb.idTree.nodeMap[gongStruct.Name]

			for _, fieldNode := range gognstructNode.Children {
				// then disable the checkbox
				if fieldNode.Type == GONG_STRUCT_FIELD {
					switch fieldWithRef := fieldNode.Gongfield.(type) {
					case *gong_models.PointerToGongStructField:
						if _, ok := map_OfGongstruct[fieldWithRef.GongStruct.Name]; !ok {
							fieldNode.IsCheckboxDisabled = true
						}
					case *gong_models.SliceOfPointerToGongStructField:
						if _, ok := map_OfGongstruct[fieldWithRef.GongStruct.Name]; !ok {
							fieldNode.IsCheckboxDisabled = true
						}
					}
				}
			}
		}

		//
		// For notes
		//
		// 1. for each note of the model, enable the check button if the class diagram is in draw mode
		for note := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {
			nodesCb.idTree.nodeMap[note.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode

			for _, link := range note.Links {
				nodesCb.idTree.nodeMap[note.Name+"."+link.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode
			}
		}
		// 2. for each note of the diagram, set the check button to true
		for _, note := range classDiagram.Notes {
			nodesCb.idTree.nodeMap[note.Name].IsChecked = true
		}

	}

	for _, stateDiagramNode := range nodesCb.StateDiagramsRootNode.Children {
		stateDiagramNode.HasEditButton = pkglet.IsEditable && stateDiagramNode.IsChecked && !stateDiagramNode.IsInEditMode
		stateDiagramNode.HasDeleteButton = pkglet.IsEditable && stateDiagramNode.IsChecked && !stateDiagramNode.IsInEditMode
		stateDiagramNode.HasDrawButton = pkglet.IsEditable && stateDiagramNode.IsChecked && !stateDiagramNode.IsInEditMode
	}

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("UpdateNodeStates, after  commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())

}
