package models

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

type NodeCallbacksSingloton struct {
	ClassdiagramsRootNode *Node
	StateDiagramsRootNode *Node

	GongstructsRootNode *Node
	GongenumsRootNode   *Node
	GongnotesRootNode   *Node
}

func (nodeCallbacksSingloton *NodeCallbacksSingloton) OnAfterUpdate(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:

		// if node is selected
		if !stagedNode.IsChecked && frontNode.IsChecked {

			// setting the value of the staged node	to the new value
			stagedNode.IsChecked = true
			stagedNode.Commit()

			// parse all nodes and uncheck them if necessary
			diagramNodes := append(
				nodeCallbacksSingloton.ClassdiagramsRootNode.Children,
				nodeCallbacksSingloton.StateDiagramsRootNode.Children...)

			for _, otherDiagramNode := range diagramNodes {
				if otherDiagramNode == stagedNode {
					continue
				}

				// uncheck the other node
				if otherDiagramNode.IsChecked {
					// log.Println("Node " + node.Name + " is checked and should be unchecked")
					otherDiagramNode.IsChecked = false
					otherDiagramNode.Commit()
				}
			}

			// update the tree of identifiers in order to show
			// which identifiers are presents in the selected diagram
			updateNodesStates(stage, nodeCallbacksSingloton)
		}

		// node was checked and user wants to uncheck it. This is not possible
		// from a application logic point of view
		// on need to commit the staged node for the front to reconstruct
		// the node as checked and overides the unchecking action
		if stagedNode.IsChecked && !frontNode.IsChecked {
			stagedNode.Commit()
		}

		// in case the front change the name of the diagram
		// one need to indicate the change as an increase in the commit
		// from the back, otherwise, the front wont be able to detect
		// the change
		// change the name of the diagram
		if stagedNode.Name != frontNode.Name {

			switch stagedNode.Type {
			case CLASS_DIAGRAM:

				// check that the name is a correct identifer in go
				for _, b := range frontNode.Name {
					if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_' || '0' <= b && b <= '9' {
						// Avoid assigning a rune for the common case of an ascii character.
						continue
					} else {
						log.Println("The name of the diagram is not a correct identifier in go: " + frontNode.Name)
						stagedNode.Commit()
						return
					}
				}

				// rename the diagram file if it exists
				// remove the actual classdiagram file if it exsits
				fieldName := GetAssociationName[DiagramPackage]().Classdiagrams[0].Name
				mapReverse := GetSliceOfPointersReverseMap[DiagramPackage, Classdiagram](fieldName)
				pkgelt := mapReverse[stagedNode.Classdiagram]

				oldClassdiagramFilePath := filepath.Join(pkgelt.Path, "../diagrams", stagedNode.Classdiagram.Name) + ".go"
				newClassdiagramFilePath := filepath.Join(pkgelt.Path, "../diagrams", frontNode.Name) + ".go"

				if _, err := os.Stat(oldClassdiagramFilePath); err == nil {
					if err := os.Rename(oldClassdiagramFilePath, newClassdiagramFilePath); err != nil {
						log.Println("Error while renaming file " + oldClassdiagramFilePath + " : " + err.Error())
					} else {
						// change the name of the go variable that describes the diagram
						// in the package
						input, err := ioutil.ReadFile(newClassdiagramFilePath)
						if err != nil {
							log.Fatalln(err)
						}

						lines := strings.Split(string(input), "\n")

						for i, line := range lines {
							if strings.Contains(line, "var "+stagedNode.Classdiagram.Name) {
								lines[i] = strings.Replace(line, stagedNode.Classdiagram.Name, frontNode.Name, 1)
								continue
							}
						}
						output := strings.Join(lines, "\n")
						err = ioutil.WriteFile(newClassdiagramFilePath, []byte(output), 0644)
						if err != nil {
							log.Fatalln(err)
						}

						stagedNode.Classdiagram.Name = frontNode.Name
						stagedNode.Classdiagram.Commit()
					}
				}

			case STATE_DIAGRAM:
				stagedNode.Umlsc.Name = frontNode.Name
				stagedNode.Umlsc.Commit()
			}
			switch stagedNode.Type {
			case CLASS_DIAGRAM, STATE_DIAGRAM:
				stagedNode.Name = frontNode.Name
				stagedNode.IsInEditMode = false
				updateNodesStates(stage, nodeCallbacksSingloton)
			}
		}

		if stagedNode.IsInEditMode != frontNode.IsInEditMode {
			stagedNode.IsInEditMode = frontNode.IsInEditMode
			updateNodesStates(stage, nodeCallbacksSingloton)
		}

		if stagedNode.IsInDrawMode != frontNode.IsInDrawMode {
			stagedNode.IsInDrawMode = frontNode.IsInDrawMode
			switch stagedNode.Type {
			case CLASS_DIAGRAM:
				stagedNode.Classdiagram.IsInDrawMode = frontNode.IsInDrawMode
			case STATE_DIAGRAM:
				stagedNode.Umlsc.IsInDrawMode = frontNode.IsInDrawMode
			}
			updateNodesStates(stage, nodeCallbacksSingloton)
		}

		// marshall diagram to switch to saved state
		if !stagedNode.IsSaved && frontNode.IsSaved {
			// fetch the only package
			var pkgelt *DiagramPackage
			for _pkgelt := range *GetGongstructInstancesSet[DiagramPackage]() {
				pkgelt = _pkgelt
			}
			switch stagedNode.Type {
			case CLASS_DIAGRAM:

				// checkout in order to get the latest version of the diagram before modifying it updated
				// by the front
				Stage.Checkout()
				stagedNode.IsSaved = false
				stagedNode.Classdiagram.Marshall(pkgelt, filepath.Join(pkgelt.Path, "../diagrams"))
				stage.Commit()
			case STATE_DIAGRAM:
				// stagedNode.Umlsc = stagedNode.Umlsc.Saved
			}

			updateNodesStates(stage, nodeCallbacksSingloton)
		}
	case GONG_STRUCT:

		// if node is unchecked
		if stagedNode.IsChecked && !frontNode.IsChecked {

			// get the latest version of the diagram before modifying it
			stage.Checkout()

			// remove the classshape from the selected diagram
			for _, classdiagramNode := range nodeCallbacksSingloton.ClassdiagramsRootNode.Children {
				if classdiagramNode.IsChecked {
					// get the diagram
					classDiagram := classdiagramNode.Classdiagram

					// get the referenced gongstructs
					for _, classshape := range classDiagram.Classshapes {
						gongstruct := classshape.Reference
						if gongstruct.Name == stagedNode.Name {
							classDiagram.RemoveClassshape(gongstruct.Name)
						}
					}
				}

				updateNodesStates(stage, nodeCallbacksSingloton)
			}
		}

		// if node is checked, add classshape
		if !stagedNode.IsChecked && frontNode.IsChecked {

			// get the latest version of the diagram before modifying it
			stage.Checkout()

			for _, classdiagramNode := range nodeCallbacksSingloton.ClassdiagramsRootNode.Children {
				if classdiagramNode.IsChecked {
					// get the diagram
					classDiagram := classdiagramNode.Classdiagram
					classDiagram.AddClassshape(frontNode.Name, REFERENCE_GONG_STRUCT)
				}

				updateNodesStates(stage, nodeCallbacksSingloton)
			}
		}

	case GONG_STRUCT_FIELD:

		// find classdiagram
		var classdiagram *Classdiagram
		for _, classdiagramNode := range nodeCallbacksSingloton.ClassdiagramsRootNode.Children {
			if classdiagramNode.IsChecked {
				// get the diagram
				classdiagram = classdiagramNode.Classdiagram
			}
		}

		// find the parent node to find the gongstruct to find the classshape
		// the node is field, one needs to find the gongstruct that contains it
		// get the parent node
		fieldName := GetAssociationName[Node]().Children[0].Name
		map_ReverseNodeParentNodeChildrend := GetSliceOfPointersReverseMap[Node, Node](fieldName)
		parentNode := map_ReverseNodeParentNodeChildrend[stagedNode]
		gongStruct := parentNode.Gongstruct

		// find the classhape in the classdiagram
		foundClassshape := false
		var classshape *Classshape
		for _, _classshape := range classdiagram.Classshapes {
			// strange behavior when the classshape is remove within the loop
			if _classshape.ReferenceName == gongStruct.Name && !foundClassshape {
				classshape = _classshape
			}
		}

		if stagedNode.IsChecked && !frontNode.IsChecked {

			// get the latest version of the diagram before modifying it
			stage.Checkout()

			{
				var field *Field

				for _, _field := range classshape.Fields {
					if _field.Fieldname == stagedNode.Name {
						field = _field
					}
				}
				if field != nil {
					classshape.Fields = remove(classshape.Fields, field)
					classshape.Heigth = classshape.Heigth - 15
					field.Unstage()
				}
			}
			{
				var link *Link

				for _, _field := range classshape.Links {
					if _field.Fieldname == stagedNode.Name {
						link = _field
					}
				}
				if link != nil {
					classshape.Links = remove(classshape.Links, link)
					link.Unstage()
				}
			}

			Stage.Commit()
		}
		if !stagedNode.IsChecked && frontNode.IsChecked {

			// get the latest version of the diagram before modifying it
			stage.Checkout()

			switch stagedNode.Gongfield.(type) {
			case *gong_models.GongBasicField, *gong_models.GongTimeField:

				var field Field
				field.Name = stagedNode.Name
				field.Fieldname = stagedNode.Name

				switch realField := stagedNode.Gongfield.(type) {
				case *gong_models.GongBasicField:

					// get the type after the "."
					names := strings.Split(realField.DeclaredType, ".")
					fieldTypeName := names[len(names)-1]

					field.Fieldtypename = fieldTypeName
				case *gong_models.GongTimeField:
					field.Fieldtypename = "Time"
				case *gong_models.PointerToGongStructField:
				case *gong_models.SliceOfPointerToGongStructField:
				}

				field.Structname = classshape.ReferenceName
				field.Stage()

				classshape.Heigth = classshape.Heigth + 15

				// construct ordered slice of fields
				rankOfFieldsInTheOriginalGongStruct := make(map[gong_models.FieldInterface]int, 0)
				nameOfFields := make(map[string]gong_models.FieldInterface, 0)

				// what is the index of the field to insert in the gong struct ?
				indexOfFieldToInsertInTheOriginalGongStruct := 0

				// let's compute it by parsing the field of the gongstruct
				gongStruct_ := (*gong_models.GetGongstructInstancesMap[gong_models.GongStruct]())[gongStruct.Name]
				for idx, gongField := range gongStruct_.Fields {

					rankOfFieldsInTheOriginalGongStruct[gongField] = idx
					nameOfFields[gongField.GetName()] = gongField

					if gongField.GetName() == field.Name {
						indexOfFieldToInsertInTheOriginalGongStruct = idx
					}
				}

				// compute indexOfFieldToInsertInTheGongStructToDisplay (index where to insert the field to display)
				indexOfFieldToInsertInTheGongStructToDisplay := 0
				for idx, field := range classshape.Fields {
					gongField := nameOfFields[field.Fieldname]
					rankInTheOriginalGoncStructOfField := rankOfFieldsInTheOriginalGongStruct[gongField]
					if indexOfFieldToInsertInTheOriginalGongStruct > rankInTheOriginalGoncStructOfField {
						indexOfFieldToInsertInTheGongStructToDisplay = idx + 1
					}
				}

				// append the filed to display in the right index
				if indexOfFieldToInsertInTheGongStructToDisplay == len(classshape.Fields) {
					classshape.Fields = append(classshape.Fields, &field)
				} else {
					classshape.Fields = append(classshape.Fields[:indexOfFieldToInsertInTheGongStructToDisplay+1],
						classshape.Fields[indexOfFieldToInsertInTheGongStructToDisplay:]...)
					classshape.Fields[indexOfFieldToInsertInTheGongStructToDisplay] = &field
				}

			case *gong_models.PointerToGongStructField, *gong_models.SliceOfPointerToGongStructField:

				var targetStructName string
				var sourceMultiplicity MultiplicityType
				var targetMultiplicity MultiplicityType

				switch realField := stagedNode.Gongfield.(type) {
				case *gong_models.PointerToGongStructField:
					targetStructName = realField.GongStruct.Name
					sourceMultiplicity = MANY
					targetMultiplicity = ZERO_ONE
				case *gong_models.SliceOfPointerToGongStructField:
					targetStructName = realField.GongStruct.Name
					sourceMultiplicity = ZERO_ONE
					targetMultiplicity = MANY
				}
				targetSourceClassshape := false
				var targetClassshape *Classshape
				for _, _classshape := range classdiagram.Classshapes {

					// strange behavior when the classshape is remove within the loop
					if _classshape.ReferenceName == targetStructName && !targetSourceClassshape {
						targetSourceClassshape = true
						targetClassshape = _classshape
					}
				}
				if !targetSourceClassshape {
					log.Panicf("Classshape %s of field not present ", targetStructName)
				}
				_ = targetClassshape

				link := new(Link).Stage()
				link.Name = stagedNode.Name
				link.SourceMultiplicity = sourceMultiplicity
				link.TargetMultiplicity = targetMultiplicity
				link.Fieldname = stagedNode.Name
				link.Structname = gongStruct.Name
				link.Fieldtypename = targetStructName

				classshape.Links = append(classshape.Links, link)
				link.Middlevertice = new(Vertice).Stage()
				link.Middlevertice.Name = "Verticle in class diagram " + classdiagram.Name +
					" in middle between " + classshape.Name + " and " + targetClassshape.Name
				link.Middlevertice.X = (classshape.Position.X+targetClassshape.Position.X)/2.0 +
					classshape.Width*1.5
				link.Middlevertice.Y = (classshape.Position.Y+targetClassshape.Position.Y)/2.0 +
					classshape.Heigth/2.0
				Stage.Commit()
			}

			Stage.Commit()
		}

	case GONG_NOTE:
		var classdiagram *Classdiagram
		for _, classdiagramNode := range nodeCallbacksSingloton.ClassdiagramsRootNode.Children {
			if classdiagramNode.IsChecked {
				// get the diagram
				classdiagram = classdiagramNode.Classdiagram
			}
		}

		if !stagedNode.IsChecked && frontNode.IsChecked {
			stage.Checkout()

			note := (&NoteShape{Name: stagedNode.Name}).Stage()

			mapOfGongNotes := *gong_models.GetGongstructInstancesMap[gong_models.GongNote]()

			gongNote, ok := mapOfGongNotes[note.Name]
			if !ok {
				log.Fatal("Unkown note ", note.Name)
			}
			note.Body = gongNote.Body
			note.X = 30
			note.Y = 30
			note.Width = 240
			note.Heigth = 63

			classdiagram.Notes = append(classdiagram.Notes, note)
			Stage.Commit()

		}
		if stagedNode.IsChecked && !frontNode.IsChecked {
			stage.Checkout()
			foundNote := false
			var note *NoteShape
			var _note *NoteShape
			for _, _note = range classdiagram.Notes {

				// strange behavior when the note is removed within the loop
				if _note.Name == stagedNode.Name && !foundNote {
					foundNote = true
					note = _note
				}
			}
			if !foundNote {
				log.Panicf("Note %s of field not present ", stagedNode.Name)
			}
			classdiagram.Notes = remove(classdiagram.Notes, note)
			note.Unstage()
			Stage.Commit()
		}
	case GONG_ENUM:
		// if node is unchecked
		if stagedNode.IsChecked && !frontNode.IsChecked {

			// get the latest version of the diagram before modifying it
			stage.Checkout()

			// remove the classshape from the selected diagram
			for _, classdiagramNode := range nodeCallbacksSingloton.ClassdiagramsRootNode.Children {
				if classdiagramNode.IsChecked {
					// get the diagram
					classDiagram := classdiagramNode.Classdiagram

					// get the referenced gongstructs
					for _, classshape := range classDiagram.Classshapes {
						reference := classshape.Reference
						if reference.Name == stagedNode.Name {
							classDiagram.RemoveClassshape(reference.Name)
						}
					}
				}

				updateNodesStates(stage, nodeCallbacksSingloton)
			}
		}

		// if node is checked, add classshape
		if !stagedNode.IsChecked && frontNode.IsChecked {

			// get the latest version of the diagram before modifying it
			stage.Checkout()

			for _, classdiagramNode := range nodeCallbacksSingloton.ClassdiagramsRootNode.Children {
				if classdiagramNode.IsChecked {
					// get the diagram
					classDiagram := classdiagramNode.Classdiagram
					classDiagram.AddClassshape(frontNode.Name, REFERENCE_GONG_ENUM)
				}

				updateNodesStates(stage, nodeCallbacksSingloton)
			}
		}
	case GONG_ENUM_VALUE:
		// find classdiagram
		var classdiagram *Classdiagram
		for _, classdiagramNode := range nodeCallbacksSingloton.ClassdiagramsRootNode.Children {
			if classdiagramNode.IsChecked {
				// get the diagram
				classdiagram = classdiagramNode.Classdiagram
			}
		}

		// find the parent node to find the gongstruct to find the classshape
		// the node is field, one needs to find the gongstruct that contains it
		// get the parent node
		fieldName := GetAssociationName[Node]().Children[0].Name
		map_ReverseNodeParentNodeChildrend := GetSliceOfPointersReverseMap[Node, Node](fieldName)
		parentNode := map_ReverseNodeParentNodeChildrend[stagedNode]
		gongEnum := parentNode.GongEnum

		// find the classhape in the classdiagram
		foundClassshape := false
		var classshape *Classshape
		for _, _classshape := range classdiagram.Classshapes {
			// strange behavior when the classshape is remove within the loop
			if _classshape.ReferenceName == gongEnum.Name && !foundClassshape {
				classshape = _classshape
			}
		}
		_ = classshape

		rankOfValuesInTheOriginalGongEnum := make(map[gong_models.GongEnumValue]int, 0)
		nameOfFields := make(map[string]gong_models.GongEnumValue, 0)

		// what is the index of the field to insert in the gong struct ?
		indexOfValueToInsertInTheOriginalGongEnum := 0

		if !stagedNode.IsChecked && frontNode.IsChecked {

			// get the latest version of the diagram before modifying it
			stage.Checkout()

			classshape.Heigth = classshape.Heigth + 15

			var field Field
			field.Name = stagedNode.Name
			field.Fieldname = stagedNode.Name

			for idx, gongEnum := range gongEnum.GongEnumValues {

				rankOfValuesInTheOriginalGongEnum[*gongEnum] = idx
				nameOfFields[gongEnum.GetName()] = *gongEnum

				if gongEnum.GetName() == field.Name {
					indexOfValueToInsertInTheOriginalGongEnum = idx
				}
			}

			// compute indexOfValueToInsertInTheGongEnumToDisplay (index where to insert the field to display)
			indexOfValueToInsertInTheGongEnumToDisplay := 0
			for idx, field := range classshape.Fields {
				gongField := nameOfFields[field.Fieldname]
				rankInTheOriginalGoncStructOfField := rankOfValuesInTheOriginalGongEnum[gongField]
				if indexOfValueToInsertInTheOriginalGongEnum > rankInTheOriginalGoncStructOfField {
					indexOfValueToInsertInTheGongEnumToDisplay = idx + 1
				}
			}

			// append the filed to display in the right index
			if indexOfValueToInsertInTheGongEnumToDisplay == len(classshape.Fields) {
				classshape.Fields = append(classshape.Fields, &field)
			} else {
				classshape.Fields = append(classshape.Fields[:indexOfValueToInsertInTheGongEnumToDisplay+1],
					classshape.Fields[indexOfValueToInsertInTheGongEnumToDisplay:]...)
				classshape.Fields[indexOfValueToInsertInTheGongEnumToDisplay] = &field
			}
			field.Stage()
			Stage.Commit()

		}
		if stagedNode.IsChecked && !frontNode.IsChecked {

			// get the latest version of the diagram before modifying it
			stage.Checkout()

			{
				var field *Field

				for _, _field := range classshape.Fields {
					if _field.Fieldname == stagedNode.Name {
						field = _field
					}
				}
				if field != nil {
					classshape.Fields = remove(classshape.Fields, field)
					classshape.Heigth = classshape.Heigth - 15
					field.Unstage()
				}
			}

			Stage.Commit()
		}
	}

	if stagedNode.IsExpanded != frontNode.IsExpanded {
		// editablelog.Println("Node " + stagedNode.Name + " is updated with value IsExpanded to " + strconv.FormatBool(frontNode.IsExpanded))

		// setting the value of the staged node	to the new value
		stagedNode.IsExpanded = frontNode.IsExpanded
	}
}

func (callbacksSingloton *NodeCallbacksSingloton) OnAfterCreate(
	stage *StageStruct,
	newDiagramNode *Node) {

	log.Println("Node " + newDiagramNode.Name + " is created")

	switch newDiagramNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:
		newDiagramNode.HasCheckboxButton = true

		// fetch the only package
		var pkgelt *DiagramPackage
		for _pkgelt := range *GetGongstructInstancesSet[DiagramPackage]() {
			pkgelt = _pkgelt
		}

		classdiagram := (&Classdiagram{Name: newDiagramNode.Name}).Stage()
		pkgelt.Classdiagrams = append(pkgelt.Classdiagrams, classdiagram)
		newDiagramNode.Classdiagram = classdiagram
		newDiagramNode.IsInEditMode = true
		newDiagramNode.IsInDrawMode = false
		newDiagramNode.HasEditButton = false

		callbacksSingloton.ClassdiagramsRootNode.Children =
			append(callbacksSingloton.ClassdiagramsRootNode.Children, newDiagramNode)

		updateNodesStates(stage, callbacksSingloton)
		stage.Commit()
	}
}

// OnAfterDelete is called after a node is deleted
// notice that the fontNode only have its basic fields updated
// its pointers are not ok
func (callbacksSingloton *NodeCallbacksSingloton) OnAfterDelete(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:
		// checkout the stage, it shall remove the link between
		// the parent node and the staged node because 0..1->0..N association
		// is stored in the staged node as a reverse pointer
		stage.Checkout()
	}
	switch stagedNode.Type {
	case CLASS_DIAGRAM:
		// remove the classdiagram node from the pkg element node
		fieldName := GetAssociationName[DiagramPackage]().Classdiagrams[0].Name
		mapReverse := GetSliceOfPointersReverseMap[DiagramPackage, Classdiagram](fieldName)
		pkgelt := mapReverse[stagedNode.Classdiagram]
		pkgelt.Classdiagrams = remove(pkgelt.Classdiagrams, stagedNode.Classdiagram)
		stagedNode.Classdiagram.Unstage()

		// remove the actual classdiagram file if it exsits
		classdiagramFilePath := filepath.Join(pkgelt.Path, "../diagrams", stagedNode.Classdiagram.Name) + ".go"
		if _, err := os.Stat(classdiagramFilePath); err == nil {
			if err := os.Remove(classdiagramFilePath); err != nil {
				log.Println("Error while deleting file " + classdiagramFilePath + " : " + err.Error())
			}
		}

	case STATE_DIAGRAM:
		// remove the umlsc node from the pkg element node
		fieldName := GetAssociationName[DiagramPackage]().Umlscs[0].Name
		mapReverse := GetSliceOfPointersReverseMap[DiagramPackage, Umlsc](fieldName)
		pkgelt := mapReverse[stagedNode.Umlsc]
		pkgelt.Umlscs = remove(pkgelt.Umlscs, stagedNode.Umlsc)
		stagedNode.Umlsc.Unstage()

		// remove the actual classdiagram file if it exsits
		statediagramFilePath := filepath.Join(pkgelt.Path, "../diagrams", stagedNode.Umlsc.Name) + ".go"
		if _, err := os.Stat(statediagramFilePath); err == nil {
			if err := os.Remove(statediagramFilePath); err != nil {
				log.Println("Error while deleting file " + statediagramFilePath + " : " + err.Error())
			}
		}
	}
	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:

		// commit will clean up the stage associations
		stage.Commit()
	}
}
