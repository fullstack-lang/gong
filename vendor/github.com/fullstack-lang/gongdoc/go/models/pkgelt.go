package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/walk"
)

// Pkgelt stores all diagrams related to a gong package
// swagger:model Pkgelt
type Pkgelt struct {
	Name string

	// Path to the "diagrams" directory
	Path string

	// GongModelPath is the package path, e.g. "fullstack-lang/gongxlsx/go/models"
	GongModelPath string

	// Classdiagrams store UML Classdiagrams
	Classdiagrams []*Classdiagram

	// Umlscs stores UML State charts diagrams
	Umlscs []*Umlsc

	// Editable indicates wether the end user can edit the diagram
	// When a diagram is used in production for navigation, the
	// model is not Editable.
	Editable bool
}

func closeFile(f *os.File) {
	fmt.Println("Closing file ", f.Name())

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
func (pkgelt *Pkgelt) Marshall(pkgPath string) error {

	// sort Classdiagrams
	sort.Slice(pkgelt.Classdiagrams[:], func(i, j int) bool {
		return pkgelt.Classdiagrams[i].Name < pkgelt.Classdiagrams[j].Name
	})
	for _, classdiagram := range pkgelt.Classdiagrams {
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
				pkgelt.GongModelPath+"\"")
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

	for _, umlsc := range pkgelt.Umlscs {
		// open file
		file, err := os.Create(filepath.Join(pkgPath, umlsc.Name) + ".go")
		defer closeFile(file)

		prelude := strings.ReplaceAll(preludeRef, "{{Imports}}", strings.ReplaceAll(pkgelt.Name, "diagrams", "models"))
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
type PkgeltMap map[string]*Pkgelt

// PkgeltStore is a handy ClassdiagramsMap
var PkgeltStore PkgeltMap = make(map[string]*Pkgelt, 0)

// Unmarshall parse the diagram package to get diagrams
// diagramPackagePath is "../diagrams" relative to the "models"
// gongModelPackagePath is the model package path, e.g. "github.com/fullstack-lang/gongxlsx/go/models"
func (pkgelt *Pkgelt) Unmarshall(modelPkg *gong_models.ModelPkg, astPkg *ast.Package, fset2 *token.FileSet, diagramPackagePath string) {

	pkgelt.Path = diagramPackagePath
	pkgelt.GongModelPath = modelPkg.PkgPath

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
							var classdiagram Classdiagram
							classdiagram.Name = vs.Names[0].Name
							_ = astPkg
							// log.Println("nb files ", len(astPkg.Files))
							astNode := vs.Values[0]
							classdiagram.Unmarshall(modelPkg, astNode, fset2)

							pkgelt.Classdiagrams = append(pkgelt.Classdiagrams, &classdiagram)
						case "Umlsc":
							var umlsc Umlsc
							umlsc.Name = vs.Names[0].Name
							astNode := vs.Values[0]
							umlsc.Unmarshall(modelPkg, astNode, fset2)

							pkgelt.Umlscs = append(pkgelt.Umlscs, &umlsc)
						}

					}
				}
			}
		}
		return true
	})
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (pkgelt *Pkgelt) SerializeToStage() {

	pkgelt.Stage()

	for _, classdiagram := range pkgelt.Classdiagrams {
		classdiagram.SerializeToStage()
	}

	for _, umlsc := range pkgelt.Umlscs {
		umlsc.SerializeToStage()
	}

}
