package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"strconv"
)

// Position mirrors joint.dia.Point
// swagger:model Position
type Position struct {
	X    float64
	Y    float64
	Name string // temporary
}

// Marshall put position into a go variable
func (position *Position) Marshall(file *os.File, nbIndentation int) error {
	indent(file, nbIndentation)
	fmt.Fprintf(file, "Position: &uml.Position{\n")
	indent(file, nbIndentation)
	fmt.Fprintf(file, "\tX: %f,\n", position.X)
	indent(file, nbIndentation)
	fmt.Fprintf(file, "\tY: %f,\n", position.Y)
	indent(file, nbIndentation)
	fmt.Fprintf(file, "},\n")

	return nil
}

// id used when unamrshalling a new shape
// this id is used to create a unique name for each shape
// this unique name is used in different algorithm of metabaron
var positionLastID uint

// PositionMap is a Map of all Position via their Name
type PositionMap map[string]*Position

// PositionStore is a handy PositionMap
var PositionStore PositionMap = make(map[string]*Position, 0)

// Unmarshall updates position value from an ast.Expr
func (position *Position) Unmarshall(expr ast.Expr, fset *token.FileSet) {

	if ue, ok := expr.(*ast.UnaryExpr); !ok {
		log.Panic("Expecting an unary expression " +
			fset.Position(expr.Pos()).String())
	} else {

		position.Name = fmt.Sprintf("Position-%04d", positionLastID)
		positionLastID++

		// expression should be a composite literal expression {X : 10, Y: 12}
		if cl, ok := ue.X.(*ast.CompositeLit); ok {

			// check that they are 2 elemnts
			if len(cl.Elts) != 2 {
				log.Panic("Expecting 2 elements in the composite lit " +
					fset.Position(cl.Pos()).String())
			}

			// parse the X value
			UnmarshallNumber(cl.Elts[0], fset, "X", &(position.X))

			// parse the Y value
			UnmarshallNumber(cl.Elts[1], fset, "Y", &(position.Y))

		} else {
			log.Panic("not a composite lit " + fset.Position(cl.Pos()).String())
		}
	}

	PositionStore[position.Name] = position
}

// UnmarshallNumber a float value
func UnmarshallNumber(expr ast.Expr, fset *token.FileSet, name string, x *float64) {
	// parse the value
	if xvaluekeyexpr, ok := expr.(*ast.KeyValueExpr); ok {

		if xkey, ok := xvaluekeyexpr.Key.(*ast.Ident); ok {
			if xkey.Name != name {
				log.Panic("First element of position is not " + name +
					" " + fset.Position(xkey.Pos()).String())
			}
		}

		// if value is positive, the xvale is a basic litteral
		// if value is negative, it is a unary expression with SUB and the BasicLit
		switch val := xvaluekeyexpr.Value.(type) {
		case *ast.BasicLit:
			if val.Kind != token.INT && val.Kind != token.FLOAT {
				log.Panic(name + " is not an INT or a FLOAT" +
					fset.Position(val.Pos()).String())
			}
			X, err := strconv.ParseFloat(val.Value, 64)
			if err != nil {
				log.Panic("Problem parsing float" + fset.Position(val.Pos()).String())
			}
			*x = X
		// negative value
		case *ast.UnaryExpr:
			if val.Op != token.SUB {
				log.Panic("Unkown sign" + fset.Position(val.Pos()).String())
			}

			if bl, ok2 := val.X.(*ast.BasicLit); !ok2 {
				log.Panic("Problem parsing float" + fset.Position(bl.Pos()).String())
			} else {
				if bl.Kind != token.INT && bl.Kind != token.FLOAT {
					log.Panic(name + " is not an INT or a FLOAT" +
						fset.Position(val.Pos()).String())
				}
				X, err := strconv.ParseFloat(bl.Value, 64)
				if err != nil {
					log.Panic("Problem parsing float" + fset.Position(val.Pos()).String())
				}
				*x = -X
			}
		default:
			log.Panic("Problem parsing float" +
				fset.Position(val.Pos()).String())
		}
	}
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (position *Position) SerializeToStage() {

	position.Stage()

}
