package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// Link represent the UML Link in any diagram
// uni-directional https://en.wikipedia.org/wiki/Association_(object-oriented_programming)
// More specificaly It is a 0..1 ---> 0..1
// swagger:model Link
type Link struct {
	Name string

	// swagger:ignore
	Field interface{} `gorm:"-"` // field that is diagrammed

	Fieldname          string
	Structname         string
	Fieldtypename      string
	TargetMultiplicity MultiplicityType
	SourceMultiplicity MultiplicityType

	// Vertices at the middle
	Middlevertice *Vertice
}

// Unmarshall
func (link *Link) Unmarshall(modelPkg *gong_models.ModelPkg, expr ast.Expr, fset *token.FileSet) {

	var cl *ast.CompositeLit
	var ok bool
	if cl, ok = expr.(*ast.CompositeLit); !ok {
		log.Panic("Expecting a composite litteral like 	Field: models.Line{}.Start, " +
			fset.Position(expr.Pos()).String())
	}

	// fieldType is the type of the field.
	//
	//  It is stored here in order to adjust it later if the type in the case of an N_M association field.
	// in a N_M assciation field, the field is an slice of pointer to a type with a suffix "Use"
	// and whose only pointer field is to the actual type
	var fieldType *gong_models.GongStruct

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

		// parse elements
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

			link.Structname = se2.Sel.Name
			link.Fieldname = se.Sel.Name
			link.Name = link.Fieldname

			// try to find the type of the field
			var typename string
			for _, _struct := range modelPkg.GongStructs {
				if _struct.Name == link.Structname {
					for _, _field := range _struct.PointerToGongStructFields {
						if _field.Name == link.Fieldname {
							typename = _field.GongStruct.Name
							fieldType = _field.GongStruct
						}
					}
					for _, _field := range _struct.SliceOfPointerToGongStructFields {
						if _field.Name == link.Fieldname {
							typename = _field.GongStruct.Name
							fieldType = _field.GongStruct
						}
					}
				}
			}
			_ = typename

			fieldtypename := typename

			// extract only the selector
			words := strings.Split(fieldtypename, ".")
			link.Fieldtypename = words[len(words)-1]

		case "Middlevertice":
			link.Middlevertice = new(Vertice)
			link.Middlevertice.Unmarshall(kve.Value, fset)
		case "TargetMultiplicity":
			var bl *ast.BasicLit
			if bl, ok = kve.Value.(*ast.BasicLit); !ok {
				// TODO deal with recursive Binary Expressions "titi"+"toto"+"tata"
				log.Panic("Expecting a basic lit " + fset.Position(kve.Value.Pos()).String())
			}
			switch bl.Value {
			case "\"" + string(ZERO_ONE) + "\"":
				link.TargetMultiplicity = ZERO_ONE
			case "\"" + string(MANY) + "\"":
				link.TargetMultiplicity = MANY
			case "\"" + string(ONE) + "\"":
				link.TargetMultiplicity = ONE
			default:
				log.Panic("Unknown multiplicity " + fset.Position(kve.Pos()).String())
			}
		case "SourceMultiplicity":
			var bl *ast.BasicLit
			if bl, ok = kve.Value.(*ast.BasicLit); !ok {
				// TODO deal with recursive Binary Expressions "titi"+"toto"+"tata"
				log.Panic("Expecting a basic lit " + fset.Position(kve.Value.Pos()).String())
			}
			switch bl.Value {
			case "\"" + string(ZERO_ONE) + "\"":
				link.SourceMultiplicity = ZERO_ONE
			case "\"" + string(MANY) + "\"":
				link.SourceMultiplicity = MANY
			case "\"" + string(ONE) + "\"":
				link.SourceMultiplicity = ONE
			default:
				log.Panic("Unknown multiplicity " + fset.Position(kve.Pos()).String())
			}

		default:
			log.Panic("Expecting 1 Field Middlevertice " + fset.Position(ident.Pos()).String())
		}
	}

	// process the special case of N_M association field
	// if both source and end multiplicity are MANY, then the field is a N_M association field
	if link.SourceMultiplicity == MANY && link.TargetMultiplicity == MANY {
		// get the field of the the field type
		if fieldType == nil {
			log.Panic(" N_M association field, expecting a field type " + fset.Position(cl.Pos()).String())
		}
		if !strings.HasSuffix(fieldType.Name, "Use") {
			log.Panic(" N_M association field, expecting a field type with a Use suffix" + fset.Position(cl.Pos()).String())
		}
		if len(fieldType.PointerToGongStructFields) != 1 {
			log.Panic(" N_M association field, expecting a field type with a single field" + fset.Position(cl.Pos()).String())
		}
		link.Fieldtypename = fieldType.PointerToGongStructFields[0].GongStruct.Name
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
		fmt.Fprintf(file, "TargetMultiplicity: \"%s\",\n", link.TargetMultiplicity)
		fmt.Fprintf(file, "\t\t\t\t\tSourceMultiplicity: \"%s\",\n", link.SourceMultiplicity)
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
