package models

import (
	"log"
	"path/filepath"
	"time"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// GongdocCommand is the struct of the instance that is updated by the front for issuing commands
// swagger:model GongdocCommand
type GongdocCommand struct {
	Name        string
	Command     GongdocCommandType
	DiagramName string // diagram involved in the command
	Date        string // important in the back for computing wether if the command has been completed or not

	GongdocNodeType GongdocNodeType // node in the tree that is involved in the command (for delete or create commands)

	StructName string // name of the gong struct (if command is about a struct or a field of a struct)
	FieldName  string // name of the field (if command is about a field)

	FieldTypeName string // type of the field (if command is about a field)
	PositionX     int
	PositionY     int

	NoteName string // name of the note to create/delete

	GongdocCommandCallback GongdocCommandCallback
}

type GongdocCommandCallback interface {
	HasSelected(gongstructName string)
}

var GongdocCommandSingloton = (&GongdocCommand{
	Name:    "Gongdoc Command Singloton",
	Command: MARSHALL_DIAGRAM,
	Date:    "",
}).Stage()

// init enables GongdocCommand to periodicaly watch the GongdocCommand
// if a more recent GongdocCommand arrives, it marshall diagrams
func init() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	nbCheckCommand := 0

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:

				// The idea of this command pooler is to checkout the GongdocCommandSingloton
				// in order to see if a command has arrived
				// In order to not perform the same command twice, the CommandCompletionDate
				// of the GongdocStatusSingloton is updated and cross checked against
				// the GongdocCommandSingloton date
				GongdocStatusSingloton.CommandCompletionDate = GongdocCommandSingloton.Date
				GongdocCommandSingloton.Checkout()

				nbCheckCommand = nbCheckCommand + 1
				// log.Println("nbCheckCommand ", nbCheckCommand, " date ", GongdocCommandSingloton.Date)
				if GongdocCommandSingloton.Date == GongdocStatusSingloton.CommandCompletionDate {
					continue
				}

				log.Println("Receive command of date ", GongdocCommandSingloton.Date, " at ", t)
				log.Println("Command status at date ", GongdocStatusSingloton.CommandCompletionDate, " at ", t)
				GongdocStatusSingloton.CommandCompletionDate = GongdocCommandSingloton.Date

				// put to stage the gongdoc stack elements in the back repo
				Stage.Checkout()

				if len(Stage.Pkgelts) != 1 {
					log.Panicf("There should be 1 package. Instead, there are %d", len(Stage.Pkgelts))
				}

				// fetch the only package
				var pkgelt *Pkgelt
				for _pkgelt := range *GetGongstructInstancesSet[Pkgelt]() {
					pkgelt = _pkgelt
				}

				var gongPackage *gong_models.ModelPkg
				for _modelPkg := range *gong_models.GetGongstructInstancesSet[gong_models.ModelPkg]() {
					gongPackage = _modelPkg
				}
				log.Println("Fetched gong package: ", gongPackage.Name)

				// fetch the classdiagram of interest
				var classDiagram *Classdiagram
				for _, _classDiagram := range pkgelt.Classdiagrams {
					if _classDiagram.Name == GongdocCommandSingloton.DiagramName {
						classDiagram = _classDiagram
					}
				}
				if classDiagram == nil && GongdocCommandSingloton.Command != DIAGRAM_GONGSTRUCT_SELECT {
					log.Panicf("Unknown class diagram %s", GongdocCommandSingloton.DiagramName)
				}
				switch GongdocCommandSingloton.Command {
				case MARSHALL_DIAGRAM:
					classDiagram.Marshall(pkgelt, filepath.Join(pkgelt.Path, "../diagrams"))
				case DIAGRAM_ELEMENT_DELETE:
					log.Printf("Delete command node type :%s", GongdocCommandSingloton.GongdocNodeType)

					foundClassshape := false
					var classshape *Classshape
					var idx int
					for _idx, _classshape := range classDiagram.Classshapes {

						// strange behavior when the classshape is remove within the loop
						if _classshape.Structname == GongdocCommandSingloton.StructName && !foundClassshape {
							classshape = _classshape
							idx = _idx
						}
					}
					if classshape == nil && GongdocCommandSingloton.GongdocNodeType != GONG_NOTE {
						log.Panicf("Unknown classshape  %s", GongdocCommandSingloton.StructName)
					}

					switch GongdocCommandSingloton.GongdocNodeType {
					case GONG_STRUCT:
						classDiagram.Classshapes = removeClassshapeFromSlice(classDiagram.Classshapes, idx)
						classshape.Position.Unstage()
						classshape.Unstage()

						// remove links that go from this classshape
						for _, link := range classshape.Links {
							link.Middlevertice.Unstage()
							link.Unstage()
						}
						classshape.Links = []*Link{}

						// remove links that go to this classshape
						for _, fromClassshape := range classDiagram.Classshapes {

							newSliceOfLinks := make([]*Link, 0)
							for _, link := range fromClassshape.Links {
								if link.Fieldtypename == classshape.Structname {
									link.Middlevertice.Unstage()
									link.Unstage()
								} else {
									newSliceOfLinks = append(newSliceOfLinks, link)
								}
							}
							fromClassshape.Links = newSliceOfLinks
						}

						// remove fields of the classshape
						for _, field := range classshape.Fields {
							field.Unstage()
						}

						Stage.Commit()
					case BASIC_FIELD, TIME_FIELD:
						var basicField *Field
						var idx int
						for _idx, _basicField := range classshape.Fields {
							if _basicField.Fieldname == GongdocCommandSingloton.FieldName {
								idx = _idx
								basicField = _basicField
							}
						}
						classshape.Fields = removeFieldFromSlice(classshape.Fields, idx)
						classshape.Heigth = classshape.Heigth - 15

						basicField.Unstage()
						Stage.Commit()
					case POINTER_TO_STRUCT, SLICE_OF_POINTER_TO_STRUCT, M_N_ASSOCIATION_FIELD:
						// check wether the classshape of the basic field is present
						foundSourceClassshape := false
						var fromClassshape *Classshape
						for _, _classshape := range classDiagram.Classshapes {

							// strange behavior when the classshape is remove within the loop
							if _classshape.Structname == GongdocCommandSingloton.StructName && !foundSourceClassshape {
								foundSourceClassshape = true
								fromClassshape = _classshape
							}
						}
						if !foundSourceClassshape {
							log.Panicf("Classshape %s of field not present ", GongdocCommandSingloton.StructName)
						}
						_ = fromClassshape

						toClassshapeFound := false
						var toClassshape *Classshape
						for _, _classshape := range classDiagram.Classshapes {

							// strange behavior when the classshape is remove within the loop
							if _classshape.Structname == GongdocCommandSingloton.FieldTypeName && !toClassshapeFound {
								toClassshapeFound = true
								toClassshape = _classshape
							}
						}
						if !toClassshapeFound {
							log.Panicf("Classshape %s of field not present ", GongdocCommandSingloton.FieldTypeName)
						}
						_ = toClassshape

						newSliceOfLinks := make([]*Link, 0)
						for _, link := range fromClassshape.Links {
							if link.Fieldname == GongdocCommandSingloton.FieldName &&
								link.Fieldtypename == GongdocCommandSingloton.FieldTypeName {
								link.Middlevertice.Unstage()
								link.Unstage()
							} else {
								newSliceOfLinks = append(newSliceOfLinks, link)
							}
						}
						fromClassshape.Links = newSliceOfLinks

						Stage.Commit()
					case GONG_NOTE:
						foundNote := false
						var note *Note
						var idx int
						var _note *Note
						for idx, _note = range classDiagram.Notes {

							// strange behavior when the note is remove within the loop
							if _note.Name == GongdocCommandSingloton.NoteName && !foundNote {
								foundNote = true
								note = _note
							}
						}
						if !foundNote {
							log.Panicf("Note %s of field not present ", GongdocCommandSingloton.StructName)
						}
						classDiagram.Notes = removeNoteFromSlice(classDiagram.Notes, idx)
						note.Unstage()
						Stage.Commit()

					}

				case DIAGRAM_ELEMENT_CREATE:
					log.Printf("Create command node type :%s", GongdocCommandSingloton.GongdocNodeType)

					switch GongdocCommandSingloton.GongdocNodeType {
					case GONG_STRUCT:
						var classshape Classshape
						classshape.Name = GongdocCommandSingloton.DiagramName + "-" + GongdocCommandSingloton.StructName
						classshape.Structname = GongdocCommandSingloton.StructName
						classshape.ClassshapeTargetType = STRUCT
						classshape.Width = 240
						classshape.Heigth = 63

						// attach GongStruct to classshape
						gongStruct, ok := Stage.GongStructs_mapString[classshape.Structname]
						if ok {
							classshape.GongStruct = gongStruct
							classshape.ShowNbInstances = true
							classshape.NbInstances = gongStruct.NbInstances
						}
						classshape.Stage()

						var position Position
						position.Name = "Pos-" + classshape.Name
						position.X = float64(GongdocCommandSingloton.PositionX)
						position.Y = float64(GongdocCommandSingloton.PositionY)
						classshape.Position = &position
						position.Stage()

						classDiagram.Classshapes = append(classDiagram.Classshapes, &classshape)
						Stage.Commit()
					case BASIC_FIELD, TIME_FIELD:
						// check wether the classshape of the basic field is present
						foundClassshape := false
						var classshape *Classshape
						for _, _classshape := range classDiagram.Classshapes {

							// strange behavior when the classshape is remove within the loop
							if _classshape.Structname == GongdocCommandSingloton.StructName && !foundClassshape {
								foundClassshape = true
								classshape = _classshape
							}
						}
						if !foundClassshape {
							log.Panicf("Classshape %s of field not present ", GongdocCommandSingloton.StructName)
						}
						_ = classshape

						var basicOrTimeField Field
						basicOrTimeField.Name = GongdocCommandSingloton.FieldName
						basicOrTimeField.Fieldname = GongdocCommandSingloton.FieldName
						if GongdocCommandSingloton.GongdocNodeType != TIME_FIELD {
							basicOrTimeField.Fieldtypename = GongdocCommandSingloton.FieldTypeName
						} else {
							basicOrTimeField.Fieldtypename = "Time"
						}
						basicOrTimeField.Structname = classshape.Structname
						basicOrTimeField.Stage()

						classshape.Heigth = classshape.Heigth + 15

						// construct ordered slice of fields
						rankOfFieldsInTheOriginalGongStruct := make(map[gong_models.FieldInterface]int, 0)
						nameOfFields := make(map[string]gong_models.FieldInterface, 0)

						// what is the index of the field to insert in the gong struct ?
						indexOfFieldToInsertInTheOriginalGongStruct := 0

						// let's compute it by parsing the field of the gongstruct
						gongStruct_ := gong_models.Stage.GongStructs_mapString[GongdocCommandSingloton.StructName]
						for idx, gongField := range gongStruct_.Fields {

							rankOfFieldsInTheOriginalGongStruct[gongField] = idx
							nameOfFields[gongField.GetName()] = gongField

							if gongField.GetName() == basicOrTimeField.Name {
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
							classshape.Fields = append(classshape.Fields, &basicOrTimeField)
						} else {
							classshape.Fields = append(classshape.Fields[:indexOfFieldToInsertInTheGongStructToDisplay+1],
								classshape.Fields[indexOfFieldToInsertInTheGongStructToDisplay:]...)
							classshape.Fields[indexOfFieldToInsertInTheGongStructToDisplay] = &basicOrTimeField
						}

						Stage.Commit()

					case POINTER_TO_STRUCT, SLICE_OF_POINTER_TO_STRUCT, M_N_ASSOCIATION_FIELD:
						// check wether the classshape of the basic field is present
						foundSourceClassshape := false
						var sourceClassshape *Classshape
						for _, _classshape := range classDiagram.Classshapes {

							// strange behavior when the classshape is remove within the loop
							if _classshape.Structname == GongdocCommandSingloton.StructName && !foundSourceClassshape {
								foundSourceClassshape = true
								sourceClassshape = _classshape
							}
						}
						if !foundSourceClassshape {
							log.Panicf("Classshape %s of field not present ", GongdocCommandSingloton.StructName)
						}
						_ = sourceClassshape

						targetSourceClassshape := false
						var targetClassshape *Classshape
						for _, _classshape := range classDiagram.Classshapes {

							// strange behavior when the classshape is remove within the loop
							if _classshape.Structname == GongdocCommandSingloton.FieldTypeName && !targetSourceClassshape {
								targetSourceClassshape = true
								targetClassshape = _classshape
							}
						}
						if !targetSourceClassshape {
							log.Panicf("Classshape %s of field not present ", GongdocCommandSingloton.StructName)
						}
						_ = targetClassshape

						link := new(Link).Stage()
						link.Name = GongdocCommandSingloton.FieldName
						link.Fieldname = GongdocCommandSingloton.FieldName
						link.Structname = GongdocCommandSingloton.StructName
						link.Fieldtypename = GongdocCommandSingloton.FieldTypeName
						switch GongdocCommandSingloton.GongdocNodeType {
						case POINTER_TO_STRUCT:
							link.SourceMultiplicity = MANY
							link.TargetMultiplicity = ZERO_ONE
						case SLICE_OF_POINTER_TO_STRUCT:
							link.SourceMultiplicity = ZERO_ONE
							link.TargetMultiplicity = MANY
						case M_N_ASSOCIATION_FIELD:
							link.SourceMultiplicity = MANY
							link.TargetMultiplicity = MANY
						}
						sourceClassshape.Links = append(sourceClassshape.Links, link)
						link.Middlevertice = new(Vertice).Stage()
						link.Middlevertice.Name = "Verticle in class diagram " + classDiagram.Name +
							" in middle between " + sourceClassshape.Name + " and " + targetClassshape.Name
						link.Middlevertice.X = (sourceClassshape.Position.X+targetClassshape.Position.X)/2.0 +
							sourceClassshape.Width*1.5
						link.Middlevertice.Y = (sourceClassshape.Position.Y+targetClassshape.Position.Y)/2.0 +
							sourceClassshape.Heigth/2.0
						Stage.Commit()
					case GONG_NOTE:
						log.Println("Note selected ", GongdocCommandSingloton.NoteName)
						note := (&Note{Name: GongdocCommandSingloton.NoteName}).Stage()

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

						classDiagram.Notes = append(classDiagram.Notes, note)
						Stage.Commit()
					}
				case DIAGRAM_GONGSTRUCT_SELECT:
					log.Println("UML Shape selected ", GongdocCommandSingloton.StructName)
					gongStruct, ok := Stage.GongStructs_mapString[GongdocCommandSingloton.StructName]
					if ok {
						if GongdocCommandSingloton.GongdocCommandCallback != nil {
							GongdocCommandSingloton.GongdocCommandCallback.HasSelected(gongStruct.Name)
						}
					}

				}
			} // end of polling function
		}
	}()
}

func removeClassshapeFromSlice(s []*Classshape, i int) []*Classshape {
	return append(s[:i], s[i+1:]...)
}

func removeNoteFromSlice(s []*Note, i int) []*Note {
	return append(s[:i], s[i+1:]...)
}

func removeLinkFromSlice(s []*Link, i int) []*Link {
	return append(s[:i], s[i+1:]...)
}

func removeFieldFromSlice(s []*Field, i int) []*Field {
	return append(s[:i], s[i+1:]...)
}
