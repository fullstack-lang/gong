package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
)

// Vertice mirrors joint.dia.Point
// swagger:model Vertice
type Vertice struct {
	X    float64
	Y    float64
	Name string // temporary
}

// Marshall put vertice into a go variable
func (vertice *Vertice) Marshall(file *os.File, nbIndentation int) error {
	fmt.Fprintf(file, "&uml.Vertice{\n")
	indent(file, nbIndentation)
	fmt.Fprintf(file, "\tX: %f,\n", vertice.X)
	indent(file, nbIndentation)
	fmt.Fprintf(file, "\tY: %f,\n", vertice.Y)
	indent(file, nbIndentation)
	fmt.Fprintf(file, "},\n")

	return nil
}

// id used when unamrshalling a new shape
// this id is used to create a unique name for each shape
// this unique name is used in different algorithm of metabaron
var verticeLastID uint

// VerticeMap is a Map of all Vertice via their Name
type VerticeMap map[string]*Vertice

// VerticeStore is a handy VerticeMap
var VerticeStore VerticeMap = make(map[string]*Vertice, 0)

// Unmarshall updates vertice value from an ast.Expr
func (vertice *Vertice) Unmarshall(expr ast.Expr, fset *token.FileSet) {

	if ue, ok := expr.(*ast.UnaryExpr); !ok {
		log.Panic("Expecting an unary expression " +
			fset.Position(expr.Pos()).String())
	} else {

		vertice.Name = fmt.Sprintf("Vertice-%04d", verticeLastID)
		verticeLastID++

		// expression should be a composite literal expression {X : 10, Y: 12}
		if cl, ok := ue.X.(*ast.CompositeLit); ok {

			// check that they are 2 elemnts
			if len(cl.Elts) != 2 {
				log.Panic("Expecting 2 elements in the composite lit " +
					fset.Position(cl.Pos()).String())
			}

			// parse the X value
			UnmarshallNumber(cl.Elts[0], fset, "X", &(vertice.X))

			// parse the Y value
			UnmarshallNumber(cl.Elts[1], fset, "Y", &(vertice.Y))

		} else {
			log.Panic("not a composite lit " + fset.Position(cl.Pos()).String())
		}
	}

	VerticeStore[vertice.Name] = vertice
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (vertice *Vertice) SerializeToStage() {

	vertice.Stage()
}
