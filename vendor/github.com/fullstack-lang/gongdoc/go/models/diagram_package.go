package models

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/walk"
)

const LegacyDiagramUmarshalling = false

// DiagramPackage stores all diagrams related to a gong package
// swagger:model DiagramPackage
type DiagramPackage struct {
	Name string

	// Path to the "diagrams" directory
	Path string

	// GongModelPath is the package path, e.g. "fullstack-lang/gongxlsx/go/models"
	GongModelPath string

	// Classdiagrams store UML Classdiagrams
	// only one diagram is present at a time
	Classdiagrams []*Classdiagram

	// list of files in the "diagrams" directory
	Files []string
	ast   *ast.Package
	fset  *token.FileSet

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
}

func (diagramPackage *DiagramPackage) Reload() {

	gong_models.Stage.Checkout()
	gong_models.Stage.Reset()
	modelPkg, _ := gong_models.LoadSource(
		filepath.Join(diagramPackage.AbsolutePathToDiagramPackage, "../models"))
	gong_models.Stage.Commit()

	Stage.Checkout()
	Stage.Reset()
	Stage.Commit()

	diagramPackage.Classdiagrams = nil
	diagramPackage.Umlscs = nil
	diagramPackage.ModelPkg = modelPkg

	fset := token.NewFileSet()
	startParser := time.Now()
	pkgsParser, errParser := parser.ParseDir(fset,
		diagramPackage.AbsolutePathToDiagramPackage,
		nil,
		parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		log.Panic("Unable to parser ")
	}
	if len(pkgsParser) != 1 {
		log.Println("Unable to parser, wrong number of parsers ", len(pkgsParser))
	} else {
		diagramPackage.Unmarshall(
			diagramPackage.ModelPkg,
			pkgsParser["diagrams"],
			fset, filepath.Join(diagramPackage.GongModelPath, "../diagrams"), "")
	}
	diagramPackage.SerializeToStage()
	FillUpNodeTree(diagramPackage)
	Stage.Commit()
}

func closeFile(f *os.File) {
	fmt.Println("Closing file (legacy)", f.Name())

	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

const preludeRef string = `package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model{{Imports}}
)

`

// Marshall translates all elements of a Pkgelt into a go file
// it recusively call Marshall function into the elements
func (diagramPackage *DiagramPackage) Marshall(pkgPath string) error {

	// sort Classdiagrams
	sort.Slice(diagramPackage.Classdiagrams[:], func(i, j int) bool {
		return diagramPackage.Classdiagrams[i].Name < diagramPackage.Classdiagrams[j].Name
	})
	for _, classdiagram := range diagramPackage.Classdiagrams {
		// open file
		file, err := os.Create(filepath.Join(pkgPath, classdiagram.Name) + ".go")
		defer closeFile(file)

		log.SetFlags(log.Lshortfile)
		filename := walk.CaptureOutput(func() { log.Printf("") })
		log.SetFlags(log.LstdFlags)
		prelude := strings.ReplaceAll(preludeRef, "{{filename}}", filename)
		prelude = strings.ReplaceAll(prelude, "{{ClassdiagramName}}", classdiagram.Name)
		if len(classdiagram.Classshapes) > 0 {
			prelude = strings.ReplaceAll(prelude, "{{Imports}}", "\n\t\""+
				diagramPackage.GongModelPath+"\"")
		} else {
			prelude = strings.ReplaceAll(prelude, "{{Imports}}", "")
		}
		prelude = strings.ReplaceAll(prelude, "docs", "models")

		if err == nil {
			fmt.Fprintf(file, prelude)
		} else {
			log.Fatal(err)
		}
		if err2 := classdiagram.MarshallAsVariable(file); err != nil {
			return err2
		}
	}

	for _, umlsc := range diagramPackage.Umlscs {
		// open file
		file, err := os.Create(filepath.Join(pkgPath, umlsc.Name) + ".go")
		defer closeFile(file)

		prelude := strings.ReplaceAll(preludeRef, "{{Imports}}", strings.ReplaceAll(diagramPackage.Name, "diagrams", "models"))
		prelude = strings.ReplaceAll(prelude, "docs", "models")

		if err == nil {
			fmt.Fprintf(file, prelude)
		} else {
			log.Fatal(err)
		}

		if err2 := umlsc.MarshallAsVariable(file); err != nil {
			return err2
		}
	}

	return nil
}

// PkgeltMap is a Map of all Classdiagrams via their Name
type PkgeltMap map[string]*DiagramPackage

// PkgeltStore is a handy ClassdiagramsMap
var PkgeltStore PkgeltMap = make(map[string]*DiagramPackage, 0)

// Unmarshall parse the diagram package to get diagrams
// diagramPackagePath is "../diagrams" relative to the "models"
// gongModelPackagePath is the model package path, e.g. "github.com/fullstack-lang/gongxlsx/go/models"
func (diagramPackage *DiagramPackage) Unmarshall(
	modelPkg *gong_models.ModelPkg,
	astPkg *ast.Package,
	fset2 *token.FileSet,
	diagramPackagePath string,
	diagramName string) (classdiagram *Classdiagram) {

	diagramPackage.Path = diagramPackagePath
	diagramPackage.GongModelPath = modelPkg.PkgPath

	ast.Inspect(astPkg, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			if len(x.Specs) > 0 {
				// log.Println("Found declaration ")
				switch vs := x.Specs[0].(type) {
				case *ast.ValueSpec:
					// log.Println("Found value spec ", vs.Names[0])

					switch se := vs.Type.(type) {
					case *ast.SelectorExpr:
						switch se.Sel.Name {
						case "Classdiagram":
							if vs.Names[0].Name != diagramName {
								return false
							}

							classdiagram = new(Classdiagram)
							classdiagram.Name = vs.Names[0].Name
							_ = astPkg
							// log.Println("nb files ", len(astPkg.Files))
							astNode := vs.Values[0]
							classdiagram.Unmarshall(modelPkg, astNode, fset2)

							diagramPackage.Classdiagrams = append(diagramPackage.Classdiagrams, classdiagram)
						case "Umlsc":
							var umlsc Umlsc
							umlsc.Name = vs.Names[0].Name
							astNode := vs.Values[0]
							umlsc.Unmarshall(modelPkg, astNode, fset2)

							diagramPackage.Umlscs = append(diagramPackage.Umlscs, &umlsc)
						}

					}
				}
			}
		}
		return true
	})
	return
}

func (diagramPackage *DiagramPackage) UnmarshallOneDiagram(diagramName string) (classdiagram *Classdiagram) {

	// for debug purposes
	gongdocStage := Stage
	_ = gongdocStage

	diagramFileName :=
		filepath.Join(diagramPackage.AbsolutePathToDiagramPackage, "../diagrams", diagramName) + ".go"
	if err := ParseAstFile(diagramFileName); err != nil {
		log.Fatalln("Unable to parse", diagramFileName)
	} else {
		log.Println("Parsed", diagramFileName)

		// there should be one diagram on the stage and it has to be
		// appended to the diagram package
		var ok bool
		classdiagram, ok = (*GetGongstructInstancesMap[Classdiagram]())[diagramName]

		if !ok {
			log.Println("Unable to find", diagramName, "in", diagramFileName, ". It might be a docs.go file")
			return
		}

		diagramPackage.Classdiagrams = append(diagramPackage.Classdiagrams,
			classdiagram)
	}
	return
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (diagramPackage *DiagramPackage) SerializeToStage() {

	diagramPackage.Stage()

	for _, classdiagram := range diagramPackage.Classdiagrams {
		classdiagram.SerializeToStage()
	}

	for _, umlsc := range diagramPackage.Umlscs {
		umlsc.SerializeToStage()
	}

}
