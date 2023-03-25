package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"strings"
	"time"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

const LegacyDiagramUmarshalling = false

// DiagramPackage stores all diagrams related to a gong package
// swagger:model DiagramPackage
type DiagramPackage struct {
	Name string

	// Stage_ where the DiagamPackage lives
	Stage_ *StageStruct

	// Path to the "diagrams" directory
	Path string

	// GongModelPath is the package path, e.g. "fullstack-lang/gongxlsx/go/models"
	GongModelPath string

	// Classdiagrams store UML Classdiagrams
	Classdiagrams []*Classdiagram

	// SelectedClassdiagram is the diagram of interest
	SelectedClassdiagram *Classdiagram

	// list of files in the "diagrams" directory
	Files []string
	Ast   *ast.Package
	Fset  *token.FileSet

	// Umlscs stores UML State charts diagrams
	Umlscs []*Umlsc

	// IsEditable indicates wether the end user can edit the diagram
	// When a diagram is used in production for navigation, the
	// model is not IsEditable.
	IsEditable bool

	// IsReload indicates if a reload of the go definition of the diagram is rrequested
	// from the end user.
	// on the back stage, this value is allways false
	// on the front stage, this value is set to true when the end user requests a reload
	// after the checkout from the front, the value is set back to false
	IsReloaded bool

	// pointer to the model package
	ModelPkg                     *gong_models.ModelPkg
	AbsolutePathToDiagramPackage string

	// swagger:ignore
	Map_Identifier_NbInstances map[string]int
}

const preludeRef string = `package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model{{Imports}}
)

`

func (diagramPackage *DiagramPackage) UnmarshallOneDiagram(stage *StageStruct, diagramName string, inFile *ast.File, fset *token.FileSet) (classdiagram *Classdiagram) {

	var err error
	startParser := time.Now()
	err = ParseAstFileFromAst(stage, inFile, fset)
	log.Printf("Parsing of %s took %s", diagramName, time.Since(startParser))

	if err != nil {
		log.Fatalln("Unable to parse", diagramName, err.Error())
	} else {
		// there should be one diagram on the stage and it has to be
		// appended to the diagram package
		var ok bool
		classdiagram, ok = (*GetGongstructInstancesMap[Classdiagram](stage))[diagramName]

		if !ok {
			// log.Println("Unable to find", diagramName, ". It might be a docs.go file")
			return
		}

		// the parsed diagram might have been in draw mode
		// let's not keep that
		classdiagram.IsInDrawMode = false

		diagramPackage.Classdiagrams = append(diagramPackage.Classdiagrams,
			classdiagram)

		for gongStructShape := range *GetGongstructInstancesSet[GongStructShape](stage) {

			_, ok := (*gong_models.GetGongstructInstancesMap[gong_models.GongStruct](diagramPackage.ModelPkg.GetStage()))[IdentifierToGongObjectName(gongStructShape.Identifier)]

			if !ok {
				log.Println("UnmarshallOneDiagram: In diagram", classdiagram.Name, "unknown note related to note shape", gongStructShape.Identifier)
				gongStructShape.Unstage(diagramPackage.Stage_)

				if contains(classdiagram.GongStructShapes, gongStructShape) {
					classdiagram.GongStructShapes = remove(classdiagram.GongStructShapes, gongStructShape)
				}
				continue
			}
		}

		// refresh all notes body from the original gong note in the package models
		// because, note are not synchronized via the gopls renaming request
		//
		// if a can be traced, this is probably for a lack of diagram maintenance
		for noteShape := range *GetGongstructInstancesSet[NoteShape](stage) {

			note, ok := (*gong_models.GetGongstructInstancesMap[gong_models.GongNote](diagramPackage.ModelPkg.GetStage()))[IdentifierToGongObjectName(noteShape.Identifier)]

			if !ok {
				log.Println("UnmarshallOneDiagram: In diagram", classdiagram.Name, "unknown note related to note shape", noteShape.Identifier)
				noteShape.Unstage(diagramPackage.Stage_)

				if contains(classdiagram.NoteShapes, noteShape) {
					classdiagram.NoteShapes = remove(classdiagram.NoteShapes, noteShape)
				}
				continue
			}

			noteShape.Body = note.Body
			noteShape.BodyHTML = note.BodyHTML
		}

		// legacy diagram file may have Fieldtypename without the ident `Point`
		// the following will turn it into `ref_models.Point`
		for link := range *GetGongstructInstancesSet[Link](stage) {

			if !strings.ContainsAny(link.Fieldtypename, ".") {
				link.Fieldtypename = GongStructNameToIdentifier(link.Fieldtypename)
			}
		}

	}
	return
}

func contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func closeFile(f *os.File) {
	fmt.Println("Closing file (legacy)", f.Name())

	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
