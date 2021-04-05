package models

import (
	"log"
	"path/filepath"
	"time"
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
}

var GongdocCommandSingloton = (&GongdocCommand{
	Name:    "Gongdoc Command Singloton",
	Command: MARSHALL_DIAGRAM,
	Date:    "",
}).Stage()

//
// init enables GongdocCommand to periodicaly watch the GongdocCommand
// if a more recent GongdocCommand arrives, it marshall diagrams
//
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
				// log.Println("nbCheckCommand ", nbCheckCommand)
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
				for _pkgelt := range Stage.Pkgelts {
					pkgelt = _pkgelt
				}

				// fetch the classdiagram of interest
				var classDiagram *Classdiagram
				for _, _classDiagram := range pkgelt.Classdiagrams {
					if _classDiagram.Name == GongdocCommandSingloton.DiagramName {
						classDiagram = _classDiagram
					}
				}
				if classDiagram == nil {
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
					if classshape == nil {
						log.Panicf("Unknown classshape  %s", GongdocCommandSingloton.StructName)
					}

					switch GongdocCommandSingloton.GongdocNodeType {
					case GONG_STRUCT:
						classDiagram.Classshapes = removeClassshapeFromSlice(classDiagram.Classshapes, idx)
						classshape.Position.Unstage()
						classshape.Unstage()
						Stage.Commit()
					case BASIC_FIELD:
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
						classshape.Stage()

						var position Position
						position.Name = "Pos-" + classshape.Name
						position.X = float64(GongdocCommandSingloton.PositionX)
						position.Y = float64(GongdocCommandSingloton.PositionY)
						classshape.Position = &position
						position.Stage()

						classDiagram.Classshapes = append(classDiagram.Classshapes, &classshape)
						Stage.Commit()
					case BASIC_FIELD:
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

						var basicField Field
						basicField.Name = GongdocCommandSingloton.FieldName
						basicField.Fieldname = GongdocCommandSingloton.FieldName
						basicField.Fieldtypename = GongdocCommandSingloton.FieldTypeName
						basicField.Structname = classshape.Structname
						basicField.Stage()

						classshape.Heigth = classshape.Heigth + 15
						classshape.Fields = append(classshape.Fields, &basicField)
						Stage.Commit()

					case POINTER_TO_STRUCT:
					case SLICE_OF_POINTER_TO_STRUCT:
					}
				}
			} // end of polling function
		}
	}()
}

func removeClassshapeFromSlice(s []*Classshape, i int) []*Classshape {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func removeFieldFromSlice(s []*Field, i int) []*Field {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
