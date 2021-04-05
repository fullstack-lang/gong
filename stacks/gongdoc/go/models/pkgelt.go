package models

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/stacks/gongdoc/go/walk"
	"golang.org/x/tools/go/packages"
)

// Pkgelt
// swagger:model Pkgelt
type Pkgelt struct {
	Name string

	Path string

	Classdiagrams []*Classdiagram

	Umlscs []*Umlsc
}

const pkgLoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo

func closeFile(f *os.File) {
	fmt.Println("Closing file ", f.Name())

	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// map of types for expressions
var MapExpToType map[string]string

const preludeRef string = `package diagrams

import (
	uml "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"

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
				strings.ReplaceAll(pkgelt.Name, "diagrams", "models")+"\"")
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
// it is "../diagrams" relative to the "models"
func (pkgelt *Pkgelt) Unmarshall(DiagramPackagePath string) {

	pkgelt.Path = DiagramPackagePath

	directory, err := filepath.Abs(DiagramPackagePath)
	if err != nil {
		log.Panic("Diagram package path does not exist %s ;" + directory)
	}
	log.Println("Loading package " + directory)

	pkgelt.FillUpMapExprComments(DiagramPackagePath)

	// we fill up the map of fields of variable with their
	// corresponding type, it is usefull to find the type of a field
	// for instance
	// Links: []*uml.Link{
	// 	{
	// 		Field: models.Line{}.Start,
	// 	},
	// },
	MapExpToType = make(map[string]string, 0)

	// get all files from the directory
	// files, err := ioutil.ReadDir(directory)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	if !file.IsDir() {
	// 		fileExtension := filepath.Ext(file.Name())

	// 		if fileExtension == ".go" {
	// 			extractTypeFromVariables(directory, file.Name(), &MapExpToType)
	// 		}
	// 	}
	// }

	var fset token.FileSet
	cfg := &packages.Config{
		Dir:   directory,
		Mode:  pkgLoadMode,
		Tests: false,

		Fset: &fset,
	}
	pkgs, err2 := packages.Load(cfg, "./...")
	if err2 != nil {
		s := fmt.Sprintf("cannot process package at path %s, err %s", DiagramPackagePath, err.Error())
		log.Panic(s)
	}

	if len(pkgs) != 1 {
		// empty package
		return
	}
	pkg := pkgs[0]
	pkgelt.Name = pkg.ID

	for expr, tv := range pkg.TypesInfo.Types {
		var buf bytes.Buffer
		posn := fset.Position(expr.Pos())
		typeValueString := tv.Type.String()
		if tv.Value != nil {
			typeValueString += " = " + tv.Value.String()
		}
		str := exprString(&fset, expr)
		// line:col | expr | mode : type = value
		if str == "models.Line{}.Start" {
			// log.Printf("trouvé %s", tvstr)
		}
		if str == "target_models.Line{}" {
			// log.Printf("trouvé %s", tvstr)
			// log.Printf("%2d:%2d | %-19s | %-7s : %s",
			// 	posn.Line, posn.Column, exprString(fset, expr),
			// 	mode(tv), tvstr)
		}
		if str == "models.Line{}" {
			// log.Printf("trouvé %s", tvstr)
		}
		MapExpToType[str] = typeValueString

		fmt.Fprintf(&buf, "%2d:%2d | %-19s | %-7s : %s",
			posn.Line, posn.Column, str,
			mode(tv), typeValueString)
	}

	// fetch all gd (generic declaration node)
	for _, f := range pkg.Syntax {
		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok {
				continue
			}

			// we are interested in gd with "var" lexical token
			if gd.Tok != token.VAR {
				continue
			}

			// we should find only one ValueSpec (constant or variable declaration)
			if len(gd.Specs) != 1 {
				log.Panicf("One variable declaration should be present instead of %d, %s",
					len(gd.Specs), fset.Position(gd.Pos()))
			}

			// value spec is the top node for
			// "var position Position = position{ X : 10, Y : 20 }"
			vs, ok := gd.Specs[0].(*ast.ValueSpec)
			if !ok {
				log.Panicf("Expected a variable declaration at %s", fset.Position(vs.Pos()))
			}

			// analyse name
			// produce error is vs.Names is no a single element array
			// of type ast.Ident. We expect the name of the diagram
			if len(vs.Names) != 1 {
				log.Panicf("bad variable declaration: %s", fset.Position(vs.Pos()))
			}
			variableName := vs.Names[0]

			// analyse the type of the variable declaration
			// X   Expr   // expression
			// Sel *Ident // field selector
			var se *ast.SelectorExpr
			if se, ok = vs.Type.(*ast.SelectorExpr); !ok {
				log.Panicf("bad type variable declaration, if should something like uml.Classshape or uml.<...>: %s",
					fset.Position(vs.Pos()))
			}

			//  fetch the X field which should an Ident node
			// var X *ast.Ident
			// if X, ok = se.X.(*ast.Ident); !ok {
			// 	log.Panicf("bad type variable declaration, selector should something like uml: %s",
			// 		fset.Position(se.Pos()))
			// }
			// log.Printf("expression is %s, selector is %s", X.Name, se.Sel.Name)

			// produce error if vs.Values is not a single element array
			// of type ast.Ident. We expect the value of the diagram
			if len(vs.Values) != 1 {
				log.Panicf("variable declaration with more than one variable at %s", fset.Position(vs.Pos()))
			}

			switch se.Sel.Name {
			case "Classdiagram":
				var classdiagram Classdiagram
				classdiagram.Name = variableName.Name
				classdiagram.Unmarshall(vs.Values[0], &fset)

				pkgelt.Classdiagrams = append(pkgelt.Classdiagrams, &classdiagram)

			case "Umlsc":
				var umlsc Umlsc
				umlsc.Name = variableName.Name
				umlsc.Unmarshall(vs.Values[0], &fset)

				pkgelt.Umlscs = append(pkgelt.Umlscs, &umlsc)

			case "Document":
			default:
				log.Panicf("Unexpected type of variable: %s at pos %s", se.Sel.Name, fset.Position(se.Pos()))
			}
		}
	}
}

func mode(tv types.TypeAndValue) string {
	switch {
	case tv.IsVoid():
		return "void"
	case tv.IsType():
		return "type"
	case tv.IsBuiltin():
		return "builtin"
	case tv.IsNil():
		return "nil"
	case tv.Assignable():
		if tv.Addressable() {
			return "var"
		}
		return "mapindex"
	case tv.IsValue():
		return "value"
	default:
		return "unknown"
	}
}

func exprString(fset *token.FileSet, expr ast.Expr) string {
	var buf bytes.Buffer
	format.Node(&buf, fset, expr)
	return buf.String()
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

// MapExprComments provides a map of expression, comments
var MapExprComments = make(map[string]string)

// FillUpMapExprComments parse the models package, and for each expression and fill up MapExprComments
// with comments
func (pkgelt *Pkgelt) FillUpMapExprComments(DiagramPackagePath string) {

	modelsPackagePath := filepath.Join(DiagramPackagePath, "../models")

	var err error
	var directory string
	if directory, err = filepath.Abs(modelsPackagePath); err != nil {
		log.Panic("Path does not exist %s ;" + directory)
	}
	log.Println("Loading package " + directory)

	cfg := &packages.Config{
		Dir:   directory,
		Mode:  pkgLoadMode,
		Tests: false,
	}

	var pkgs []*packages.Package
	if pkgs, err = packages.Load(cfg, "./..."); err != nil {
		s := fmt.Sprintf("cannot process package at path %s, err %s", modelsPackagePath, err.Error())
		log.Panic(s)
	}

	if len(pkgs) != 1 {
		log.Panicf("Expected 1 package to scope, found %d", len(pkgs))
	}
	pkg := pkgs[0]
	for _, f := range pkg.Syntax {
		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, s := range gd.Specs {
				if ts, ok := s.(*ast.TypeSpec); ok {
					cg := gd.Doc

					MapExprComments[ts.Name.Name] = cg.Text()
				}
			}
		}
	}

}
