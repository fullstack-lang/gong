package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"strings"
)

// Link represent the UML Link in any diagram
// uni-directional https://en.wikipedia.org/wiki/Association_(object-oriented_programming)
// More specificaly It is a 0..1 ---> 0..1
// swagger:model Link
type Link struct {
	Name string

	// swagger:ignore
	Field         interface{} `gorm:"-"` // field that is diagrammed
	Fieldname     string
	Structname    string
	Fieldtypename string
	Multiplicity  MultiplicityType

	// Vertices at the middle
	// swagger:ignore
	Middlevertice *Vertice
}

type MultiplicityType string

// values for EnumType
const (
	ZERO_ONE MultiplicityType = "0..1"
	ONE      MultiplicityType = "1"
	MANY     MultiplicityType = "*"
)

// Unmarshall
func (link *Link) Unmarshall(expr ast.Expr, fset *token.FileSet) {

	var cl *ast.CompositeLit
	var ok bool
	if cl, ok = expr.(*ast.CompositeLit); !ok {
		log.Panic("Expecting a composite litteral like 	Field: models.Line{}.Start, " +
			fset.Position(expr.Pos()).String())
	}

	// extract all elements
	for _, elt := range cl.Elts {
		var kve *ast.KeyValueExpr
		if kve, ok = elt.(*ast.KeyValueExpr); !ok {
			log.Panic("Expecting 1 key value Field: ... " + fset.Position(cl.Pos()).String())
		}

		// get the key "Field"
		var ident *ast.Ident
		if ident, ok = kve.Key.(*ast.Ident); !ok {
			log.Panic("Expecting 1 ident " + fset.Position(kve.Pos()).String())
		}

		// check Link Field is
		switch ident.Name {
		case "Field":
			var kve *ast.KeyValueExpr
			if kve, ok = elt.(*ast.KeyValueExpr); !ok {
				log.Panic("Expecting 1 key value Field: ... " + fset.Position(cl.Pos()).String())
			}

			// get the key "Field"
			var ident *ast.Ident
			if ident, ok = kve.Key.(*ast.Ident); !ok {
				log.Panic("Expecting 1 ident " + fset.Position(kve.Pos()).String())
			}
			// check Link Field is "Field", ah ah ah !!! happy meta programmer
			if ident.Name != "Field" {
				log.Panic("Expecting 1 Field Field " + fset.Position(ident.Pos()).String())
			}

			var se *ast.SelectorExpr
			if se, ok = kve.Value.(*ast.SelectorExpr); !ok {
				log.Panic("Expecting 1 selector " + fset.Position(kve.Pos()).String())
			}

			var cl2 *ast.CompositeLit
			if cl2, ok = se.X.(*ast.CompositeLit); !ok {
				log.Panic("Expecting 1 composite lit " + fset.Position(se.Pos()).String())
			}

			var se2 *ast.SelectorExpr
			if se2, ok = cl2.Type.(*ast.SelectorExpr); !ok {
				log.Panic("Expecting 1 selector " + fset.Position(cl.Pos()).String())
			}

			var ident2 *ast.Ident
			if ident2, ok = se2.X.(*ast.Ident); !ok {
				log.Panic("Expecting 1 ident " + fset.Position(se2.Pos()).String())
			}

			structnameWithX := ident2.Name + "." + se2.Sel.Name
			link.Structname = se2.Sel.Name
			link.Fieldname = se.Sel.Name
			link.Name = link.Fieldname

			// now, let's find the link target !!!
			fieldname := fmt.Sprintf("%s{}.%s", structnameWithX, link.Fieldname)

			fieldtypename := MapExpToType[fieldname]

			// extract only the selector
			words := strings.Split(fieldtypename, ".")
			link.Fieldtypename = words[len(words)-1]

		case "Middlevertice":
			link.Middlevertice = new(Vertice)
			link.Middlevertice.Unmarshall(kve.Value, fset)
		case "Multiplicity":
			var bl *ast.BasicLit
			if bl, ok = kve.Value.(*ast.BasicLit); !ok {
				// TODO deal with recursive Binary Expressions "titi"+"toto"+"tata"
				log.Panic("Expecting a basic lit " + fset.Position(kve.Value.Pos()).String())
			}
			switch bl.Value {
			case "\"" + string(ZERO_ONE) + "\"":
				link.Multiplicity = ZERO_ONE
			case "\"" + string(MANY) + "\"":
				link.Multiplicity = MANY
			case "\"" + string(ONE) + "\"":
				link.Multiplicity = ONE
			default:
				log.Panic("Unknown multiplicity " + fset.Position(kve.Pos()).String())
			}

		default:
			log.Panic("Expecting 1 Field Middlevertice " + fset.Position(ident.Pos()).String())
		}
	}
}

// Marshall provides the element of link as declaration
func (link *Link) Marshall(file *os.File, nbIndentation int) error {
	indent(file, nbIndentation)
	fmt.Fprintf(file, "{\n")

	indent(file, nbIndentation)
	fmt.Fprintf(file, "\tField: models.%s{}.%s,\n", link.Structname, link.Fieldname)

	if link.Middlevertice != nil {
		indent(file, nbIndentation+1)
		fmt.Fprintf(file, "Middlevertice: ")
		link.Middlevertice.Marshall(file, nbIndentation+1)
	}

	// add multiplicity
	{
		indent(file, nbIndentation+1)
		fmt.Fprintf(file, "Multiplicity: \"%s\",\n", link.Multiplicity)
	}

	indent(file, nbIndentation)
	fmt.Fprintf(file, "}")

	return nil
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (link *Link) SerializeToStage() {

	link.Stage()

	if link.Middlevertice != nil {
		link.Middlevertice.Stage()
	}
}
