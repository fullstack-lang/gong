package models

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// GenGoDefaultDiagram generates the Ng Enum
func GenGoDefaultDiagram(modelPkg *gong_models.ModelPkg, pkgPath string) {

	// generate diagrams for documentation
	var pkgelt Pkgelt
	// parse the diagram package
	diagramPkgPath := filepath.Join(pkgPath, "../diagrams")

	// check existance of path
	_, err := os.Stat(diagramPkgPath)

	// if directory does not exist, creates it
	if os.IsNotExist(err) {
		errd := os.Mkdir(diagramPkgPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + diagramPkgPath)
		}
	}

	// generates default diagram
	{
		var pkgelt_default Pkgelt
		pkgelt_default.Name = modelPkg.PkgPath

		defaultClassDiagramm := new(Classdiagram)
		defaultClassDiagramm.Name = "defaultDiagram"
		pkgelt_default.Classdiagrams = append(pkgelt_default.Classdiagrams, defaultClassDiagramm)

		idx := 0.0
		for _, _enum := range modelPkg.GongEnums {
			classshape := new(Classshape)
			classshape.ClassshapeTargetType = ENUM
			classshape.Name = _enum.Name

			classshape.Position = new(Position)
			classshape.Structname = _enum.Name

			classshape.Position.X = 40.0 + 300.0*idx
			idx = idx + 1
			classshape.Position.Y = 500.0

			defaultClassDiagramm.Classshapes = append(defaultClassDiagramm.Classshapes, classshape)

			for _, value := range _enum.GongEnumValues {
				field := new(Field)
				field.Fieldname = value.Name
				field.Structname = _enum.Name
				field.Field = value
				classshape.Fields = append(classshape.Fields, field)
			}
		}

		idx = 0.0
		for _, _struct := range modelPkg.GongStructs {
			classshape := new(Classshape)
			classshape.ClassshapeTargetType = STRUCT

			classshape.Name = _struct.Name

			classshape.Position = new(Position)
			classshape.Structname = _struct.Name

			classshape.Position.X = 40.0 + 300.0*idx
			idx = idx + 1.0
			classshape.Position.Y = 40.0

			defaultClassDiagramm.Classshapes = append(defaultClassDiagramm.Classshapes, classshape)

			// get fields

			linkIndex := 0
			for _, _field := range _struct.Fields {

				// only put exported fields
				if _field.GetName() == strings.Title(_field.GetName()) {

					switch _field.(type) {
					case *gong_models.PointerToGongStructField:
						pointerToGongStructField := _field.(*gong_models.PointerToGongStructField)

						link := new(Link)
						link.Fieldname = pointerToGongStructField.Name
						link.Structname = _struct.Name
						link.Field = _field
						link.TargetMultiplicity = ZERO_ONE

						link.Middlevertice = new(Vertice)
						link.Middlevertice.X = 40 + 300.0*float64(idx) + 250
						link.Middlevertice.Y = 200.0 + float64(linkIndex)*50
						classshape.Links = append(classshape.Links, link)

						linkIndex = linkIndex + 1
					case *gong_models.SliceOfPointerToGongStructField:
						sliceOfPointerToGongStructField := _field.(*gong_models.SliceOfPointerToGongStructField)

						link := new(Link)
						link.Fieldname = sliceOfPointerToGongStructField.Name
						link.Structname = _struct.Name
						link.Field = _field
						link.TargetMultiplicity = MANY

						link.Middlevertice = new(Vertice)
						link.Middlevertice.X = 40 + 300.0*float64(idx) + 250
						link.Middlevertice.Y = 200.0 + float64(linkIndex)*50
						classshape.Links = append(classshape.Links, link)

						linkIndex = linkIndex + 1
					case *gong_models.GongBasicField:
						basicField := _field.(*gong_models.GongBasicField)

						field := new(Field)
						field.Fieldname = basicField.Name
						field.Structname = _struct.Name
						field.Field = _field
						classshape.Fields = append(classshape.Fields, field)

					default:
					}

				}

			}
		}
		pkgelt_default.Marshall(diagramPkgPath)
	}

	// generates all diagrams
	if !os.IsNotExist(err) {
		pkgelt.Unmarshall(modelPkg, nil, nil, diagramPkgPath)

		for _, classDiagram := range pkgelt.Classdiagrams {
			classDiagram.OutputSVG(diagramPkgPath)
		}
	}

}
