package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"reflect"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// Field represent the UML Field of a Class (a "struct" in go)
// swagger:model Field
type Field struct {
	Name string

	// swagger:ignore
	// actual go field in the models that is diagrammed, for instance "models.Point{}.X"
	Field interface{} `gorm:"-"`

	Fieldname         string
	FieldTypeAsString string
	Structname        string
	Fieldtypename     string
}

// Unmarshall
func (field *Field) Unmarshall(modelPkg *gong_models.ModelPkg, expr ast.Expr, fset *token.FileSet) {

	var cl *ast.CompositeLit
	var ok bool
	if cl, ok = expr.(*ast.CompositeLit); !ok {
		log.Panic("Expecting a composite litteral like 	Field: models.Point{}.X, " +
			fset.Position(expr.Pos()).String())
	}

	{
		var kve *ast.KeyValueExpr
		if kve, ok = cl.Elts[0].(*ast.KeyValueExpr); !ok {
			log.Panic("Expecting 1 key value Field: ... " + fset.Position(cl.Pos()).String())
		}

		// get the key "Field"
		var ident *ast.Ident
		if ident, ok = kve.Key.(*ast.Ident); !ok {
			log.Panic("Expecting 1 ident " + fset.Position(kve.Pos()).String())
		}
		// check Field Field is "Field", ah ah ah !!! happy meta programmer
		if ident.Name != "Field" {
			log.Panic("Expecting 1 Field Field " + fset.Position(ident.Pos()).String())
		}

		// parsing models.Line{}.Start, which is
		// ast.SelectorExpr
		// 	X: ast.CompositeLit
		// 		Type: ast.SelectorExpr
		// 			X: ast.Ident
		// 				Name: "models"
		// 			Sel: ast.Ident
		// 				Name: "Line"
		// 	Sel: ast.Ident
		// 		Name: "Start"
		var se *ast.SelectorExpr
		if se, ok = kve.Value.(*ast.SelectorExpr); !ok {
			log.Panic("Expecting 1 selector " + fset.Position(kve.Pos()).String())
		}

		// in case the classhape is a STRUCT
		var cl2 *ast.CompositeLit
		if cl2, ok = se.X.(*ast.CompositeLit); ok {
			var se2 *ast.SelectorExpr
			if se2, ok = cl2.Type.(*ast.SelectorExpr); !ok {
				log.Panic("Expecting 1 selector " + fset.Position(cl.Pos()).String())
			}

			field.Structname = se2.Sel.Name
			field.Fieldname = se.Sel.Name
			field.Name = field.Fieldname

			// try to find the type of the field
			var typename string
			for _, _struct := range modelPkg.GongStructs {
				if _struct.Name == field.Structname {
					for _, _field := range _struct.GongBasicFields {
						if _field.Name == field.Fieldname {
							typename = _field.DeclaredType
						}
					}
					for _, _field := range _struct.GongTimeFields {
						if _field.Name == field.Fieldname {
							typename = "time.Time"
						}
					}
				}
			}
			fieldtypename := typename

			// extract only the selector
			words := strings.Split(fieldtypename, ".")
			field.Fieldtypename = words[len(words)-1]
		}

		var ident2 *ast.Ident
		if ident2, ok = se.X.(*ast.Ident); ok {
			field.Structname = ident2.Name
			field.Fieldname = se.Sel.Name
		}

	}
}

// Marshall provides the element of field as declaration
func (field *Field) Marshall(file *os.File, nbIndentation int, ClassshapeTargetType ClassshapeTargetType) error {
	indent(file, nbIndentation)
	fmt.Fprintf(file, "{\n")

	indent(file, nbIndentation)
	if ClassshapeTargetType == STRUCT {
		fmt.Fprintf(file, "\tField: models.%s{}.%s,\n", field.Structname, field.Fieldname)
	} else {
		fmt.Fprintf(file, "\tField: models.%s,\n", field.Fieldname)
	}

	indent(file, nbIndentation)
	fmt.Fprintf(file, "}")

	return nil
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (field *Field) SerializeToStage() {

	field.Stage()

	// update name if not done
	if field.Name == "" {
		if field.Field != nil {
			// when a field can be is a struct
			typeofstruct := reflect.TypeOf(field.Field)

			if typeofstruct.Kind() == reflect.Ptr {
				log.Printf(typeofstruct.Elem().String())
				field.Name = strings.Split(typeofstruct.Elem().String(), ".")[1]
				field.Structname = field.Name
			}
		}
	}
}
