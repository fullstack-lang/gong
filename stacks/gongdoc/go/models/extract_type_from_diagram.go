package models

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"path/filepath"
)

// extractTypeFromVariables takes a directory and a filename in this directory and generates
// a map linking Expressions to Type
//
// in details, it populates a
//
// 	info := types.Info{
// 		Types: make(map[ast.Expr]types.TypeAndValue),
// 		Defs:  make(map[*ast.Ident]types.Object),
// 		Uses:  make(map[*ast.Ident]types.Object),
// 	}
//
// with the call to
//
// 	type.Config.Check()
//
// Then, it ranges all expression in the map[ast.Expr]types.TypeAndValue and for each expression
// it extract the associated string and if the string match a pattern that is of interest for
// the diagrams,  it stores the match into
//
//	  MapExpToType *map[string]string)
//
// Which will used later
func extractTypeFromVariables(directory string, filename string, MapExpToType *map[string]string) {

	diagramFilepath := filepath.Join(directory, filename)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, diagramFilepath, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Type-check the package.
	// We create an empty map for each kind of input
	// we're interested in, and Check populates them.
	//
	// Type-check the package.
	// We create an empty map for each kind of input
	// we're interested in, and Check populates them.
	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	var config types.Config
	config.Importer = importer.Default()
	_, err = config.Check(directory, fset, []*ast.File{f}, &info)
	if err != nil {
		log.Fatal(err)
	}

	// Print package-level variables in initialization order.
	fmt.Printf("InitOrder: %v\n\n", info.InitOrder)

	fmt.Println("Types and Values of each expression:")
	for expr, tv := range info.Types {
		var buf bytes.Buffer
		posn := fset.Position(expr.Pos())
		typeValueString := tv.Type.String()
		if tv.Value != nil {
			typeValueString += " = " + tv.Value.String()
		}
		// line:col | expr | mode : type = value
		if exprString(fset, expr) == "models.Line{}.Start" {
			// log.Printf("trouvé %s", tvstr)
		}
		if exprString(fset, expr) == "target_models.Line{}" {
			// log.Printf("trouvé %s", tvstr)
			// log.Printf("%2d:%2d | %-19s | %-7s : %s",
			// 	posn.Line, posn.Column, exprString(fset, expr),
			// 	mode(tv), tvstr)
		}
		if exprString(fset, expr) == "models.Line{}" {
			// log.Printf("trouvé %s", tvstr)
		}
		(*MapExpToType)[exprString(fset, expr)] = typeValueString

		fmt.Fprintf(&buf, "%2d:%2d | %-19s | %-7s : %s",
			posn.Line, posn.Column, exprString(fset, expr),
			mode(tv), typeValueString)
	}

}
