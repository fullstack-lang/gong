package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
)

// UmlState mirrors joint.shapes.uml.UmlState
// swagger:model UmlState
type UmlState struct {
	Name string

	X float64
	Y float64
}

// Marshall provides the element of state as declaration
func (state *UmlState) Marshall(file *os.File, nbIndentation int) error {
	indent(file, nbIndentation)
	fmt.Fprintf(file, "{\n")
	indent(file, nbIndentation)
	// four spaces are necessary to align on the Name field
	fmt.Fprintf(file, "\tX:    %f,\n", state.X)
	indent(file, nbIndentation)
	fmt.Fprintf(file, "\tY:    %f,\n", state.Y)
	indent(file, nbIndentation)
	fmt.Fprintf(file, "\tName: string(models.%s),\n", state.Name)
	indent(file, nbIndentation)
	fmt.Fprintf(file, "}")

	return nil
}

// MarshallAsVariable ...
func (state *UmlState) MarshallAsVariable(file *os.File) error {

	fmt.Fprintf(file, "var %s uml.UmlState = uml.UmlState", state.Name)
	state.Marshall(file, 0)
	fmt.Fprintf(file, "\n\n")
	return nil
}

// id used when unamrshalling a new shape
// this id is used to create a unique name for each shape
// this unique name is used in different algorithm of metabaron
var stateLastID uint

// StateMap is a Map of all State via their Name
type StateMap map[string]*UmlState

// StateStore is a handy StateMap
var StateStore StateMap = make(map[string]*UmlState, 0)

// Unmarshall updates state values from an ast.Epr
func (state *UmlState) Unmarshall(expr ast.Expr, fset *token.FileSet) {

	// expression should be a composite literal expression
	// models.State{Position: uml.Position{X: 10, Y: 12}, Struct: &(models.Point{})}
	var ok bool
	var cl *ast.CompositeLit
	if cl, ok = expr.(*ast.CompositeLit); !ok {
		log.Panic("Expecting a composite litteral " +
			fset.Position(expr.Pos()).String())
	}
	// check that they are 2 elemnts
	if len(cl.Elts) != 3 {
		log.Panic("Expecting 3 elements in the composite lit " +
			fset.Position(cl.Pos()).String())
	}

	// parse the X value
	UnmarshallNumber(cl.Elts[0], fset, "X", &(state.X))

	// parse the Y value
	UnmarshallNumber(cl.Elts[1], fset, "Y", &(state.Y))

	// parse the Name
	var kve *ast.KeyValueExpr
	if kve, ok = cl.Elts[2].(*ast.KeyValueExpr); !ok {
		log.Panic("Expecting a key value expressin " + fset.Position(cl.Elts[2].Pos()).String())
	}

	var ident *ast.Ident
	if ident, ok = kve.Key.(*ast.Ident); !ok {
		log.Panic("Expecting an ident " + fset.Position(kve.Pos()).String())
	}
	if ident.Name != "Name" {
		log.Panic("Expecting an ident `Name`" + fset.Position(ident.Pos()).String())
	}

	var ce *ast.CallExpr
	if ce, ok = kve.Value.(*ast.CallExpr); !ok {
		log.Panic("Expecting a call expression string( ...) " + fset.Position(ident.Pos()).String())
	}

	if ident, ok = ce.Fun.(*ast.Ident); !ok {
		log.Panic("Expecting an ident " + fset.Position(ce.Fun.Pos()).String())
	}
	if ident.Name != "string" {
		log.Panic("Expecting an ident `string`" + fset.Position(ident.Pos()).String())
	}

	if len(ce.Args) != 1 {
		log.Panic("Expecting only one arg to string(...)" + fset.Position(ce.Pos()).String())
	}

	var se *ast.SelectorExpr
	if se, ok = ce.Args[0].(*ast.SelectorExpr); !ok {
		log.Panic("Expecting a selector expression like models.AVANT_CALCUL" + fset.Position(ce.Pos()).String())
	}

	if ident, ok = se.X.(*ast.Ident); !ok {
		log.Panic("Expecting an ident expression in X" + fset.Position(se.X.Pos()).String())
	}
	if ident.Name != "models" {
		log.Panic("Expecting an ident `models`" + fset.Position(ident.Pos()).String())
	}

	state.Name = se.Sel.Name

	StateStore[state.Name] = state
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (state *UmlState) SerializeToStage() {

	state.Stage()
}
