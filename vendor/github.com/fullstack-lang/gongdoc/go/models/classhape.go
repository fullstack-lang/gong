package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"path"
	"sort"
	"strconv"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

const ClassshapeDefaultWidth = 240.0
const ClassshapeDefaultHeigth = 48.0

// Classshape mirrors joint.shapes.uml.Class
// swagger:model Classshape
type Classshape struct {
	Name string

	Position *Position

	// swagger:ignore, an "interface" field cannot be used by gong, therefore, one specifies
	// both swagger and gorm ignore magic code
	ReferencedGong interface{} `gorm:"-"` // pointer to the struct/enum of the model that it is diagramming
	ReferenceName  string

	// the related gong struct or gong enum
	Reference *Reference

	// Identifier is the identifier of the struct referenced by the shape in the modeled package
	//gong:ident
	Identifier string

	// gongdoc can be integrated in a runtime application
	// the application can then set up the number of instances of Struct
	ShowNbInstances bool
	NbInstances     int

	// models of the composition of Field
	Fields []*Field

	// models of the composition of Link
	Links []*Link

	// with and height of the shape when they are rendered on SVG or with jointjs
	// They are optional fields. they can be computed when empty
	Width, Heigth float64

	// this is always false in the backend, but it can be set to true by the front end
	// this means it is selected by the user
	IsSelected bool
}

// Marshall provides the element of classshape as declaration
func (classshape *Classshape) Marshall(file *os.File, nbIndentation int) error {
	indent(file, nbIndentation)
	fmt.Fprintf(file, "{\n")

	indent(file, nbIndentation)
	if classshape.Reference.Type == REFERENCE_GONG_STRUCT {
		fmt.Fprintf(file, "\tReferencedGong: &(models.%s{}),\n", classshape.ReferenceName)
	}
	if classshape.Reference.Type == REFERENCE_GONG_ENUM {
		fmt.Fprintf(file, "\tReferencedGong: new(models.%s),\n", classshape.ReferenceName)
	}

	if err := classshape.Position.Marshall(file, nbIndentation+1); err != nil {
		return err
	}

	if classshape.Width != 0.0 {
		indent(file, nbIndentation)
		fmt.Fprintf(file, "\tWidth:  %f,\n", classshape.Width)
	}
	if classshape.Heigth != 0.0 {
		indent(file, nbIndentation)
		fmt.Fprintf(file, "\tHeigth: %f,\n", classshape.Heigth)
	}

	if len(classshape.Links) > 0 {
		// sort Links
		sort.Slice(classshape.Links[:], func(i, j int) bool {
			return classshape.Links[i].Fieldname < classshape.Links[j].Fieldname
		})
		indent(file, nbIndentation+1)
		fmt.Fprintf(file, "Links: []*uml.Link{\n")
		for _, link := range classshape.Links {
			link.Marshall(file, nbIndentation+2)
			fmt.Fprintf(file, ",\n")
		}
		indent(file, nbIndentation+1)
		fmt.Fprintf(file, "},\n")
	}

	if len(classshape.Fields) > 0 {
		// sort Fields
		sort.Slice(classshape.Fields[:], func(i, j int) bool {
			return (classshape.Fields[i].Structname + "." + classshape.Fields[i].Fieldname) <
				(classshape.Fields[j].Structname + "." + classshape.Fields[j].Fieldname)
		})
		indent(file, nbIndentation+1)
		fmt.Fprintf(file, "Fields: []*uml.Field{\n")
		for _, field := range classshape.Fields {
			field.Marshall(file, nbIndentation+2, classshape.Reference.Type)
			fmt.Fprintf(file, ",\n")
		}
		indent(file, nbIndentation+1)
		fmt.Fprintf(file, "},\n")
	}

	indent(file, nbIndentation)
	fmt.Fprintf(file, "}")

	return nil
}

// MarshallAsVariable
func (classshape *Classshape) MarshallAsVariable(file *os.File) error {

	fmt.Fprintf(file, "var %s uml.Classshape = uml.Classshape", classshape.Name)
	classshape.Marshall(file, 0)
	fmt.Fprintf(file, "\n\n")
	return nil
}

// id used when unamrshalling a new shape
// this id is used to create a unique name for each shape
// this unique name is used in different algorithm of metabaron
var classshapeLastID uint

// ClassshapeMap is a Map of all Classshape via their Name
type ClassshapeMap map[string]*Classshape

// ClassshapeStore is a handy ClassshapeMap
var ClassshapeStore ClassshapeMap = make(map[string]*Classshape, 0)

// Unmarshall updates classshape values from an ast.Epr
func (classshape *Classshape) Unmarshall(modelPkg *gong_models.ModelPkg, expr ast.Expr, fset *token.FileSet) {

	// expression should be a composite literal expression
	// models.Classshape{Position: uml.Position{X: 10, Y: 12}, ReferencedGong: &(models.Point{})}
	if cl, ok := expr.(*ast.CompositeLit); !ok {
		log.Fatal("Expecting a composite litteral " +
			fset.Position(expr.Pos()).String())
	} else {

		for _, elt := range cl.Elts {
			var kve *ast.KeyValueExpr
			if kve, ok = elt.(*ast.KeyValueExpr); !ok {
				log.Fatal("Element should be a KeyValueExpr" + fset.Position(elt.Pos()).String())
			}

			var ident *ast.Ident
			if ident, ok = kve.Key.(*ast.Ident); !ok {
				log.Fatal("Element Key should be an Ident" + fset.Position(kve.Pos()).String())
			}

			switch ident.Name {

			case "Position":
				var position Position
				classshape.Position = &position
				classshape.Position.Unmarshall(kve.Value, fset)
			case "ReferencedGong":
				// expect an Unary Expression if it is a struct
				var ok bool
				var valuekey *ast.UnaryExpr
				if valuekey, ok = kve.Value.(*ast.UnaryExpr); ok {

					var pe *ast.ParenExpr
					if pe, ok = valuekey.X.(*ast.ParenExpr); !ok {
						log.Fatal("Expression should be parenthese expression" +
							fset.Position(valuekey.Pos()).String())
					}

					// expect a Composite Litteral with no Element <type>{}
					var cl *ast.CompositeLit
					if cl, ok = pe.X.(*ast.CompositeLit); !ok {
						log.Fatal("Expression should be a composite lit" +
							fset.Position(cl.Pos()).String())
					}

					// get type models.xxxx
					var se *ast.SelectorExpr
					if se, ok = cl.Type.(*ast.SelectorExpr); !ok {
						log.Fatal("Expression should be a selector" +
							fset.Position(cl.Pos()).String())
					}

					// assign Strut
					classshape.ReferencedGong = se.Sel.Name
					classshape.ReferenceName = se.Sel.Name
					classshape.Identifier = RefPrefixReferencedPackage + path.Base(modelPkg.PkgPath) + "." + se.Sel.Name

					// attach reference to classshape
					var ok bool
					var reference *Reference
					reference, ok = (*GetGongstructInstancesMap[Reference]())[classshape.ReferenceName]
					if !ok {
						reference = (&Reference{Name: classshape.ReferenceName}).Stage()
						reference.Type = REFERENCE_GONG_STRUCT
					}
					classshape.Reference = reference
					classshape.ShowNbInstances = true
					classshape.NbInstances = reference.NbInstances

					if classshape.Name == "" {
						classshape.Name = fmt.Sprintf("Classshape%04d", classshapeLastID)
					}
				}

				// if this is an enum, we expect a se
				// get type new(models.xxxx)
				var ce *ast.CallExpr
				if ce, ok = kve.Value.(*ast.CallExpr); ok {

					if len(ce.Args) != 1 {
						log.Fatal("Expression should be new(...)" +
							fset.Position(cl.Pos()).String())
					}

					var se *ast.SelectorExpr
					if se, ok = ce.Args[0].(*ast.SelectorExpr); !ok {
						log.Fatal("Expression should be a selector" +
							fset.Position(cl.Pos()).String())
					}

					classshape.ReferencedGong = se.Sel.Name
					classshape.ReferenceName = se.Sel.Name
					classshape.Identifier = RefPrefixReferencedPackage + path.Base(modelPkg.PkgPath) + "." + se.Sel.Name

					var ok bool
					var reference *Reference
					reference, ok = (*GetGongstructInstancesMap[Reference]())[se.Sel.Name]
					if !ok {
						reference = (&Reference{Name: se.Sel.Name}).Stage()
						reference.Type = REFERENCE_GONG_ENUM
					}
					classshape.Reference = reference

					if classshape.Name == "" {
						classshape.Name = fmt.Sprintf("Classshape%04d", classshapeLastID)
					}
				}

				classshapeLastID++
			case "Links":
				var cl *ast.CompositeLit
				var ok bool
				if cl, ok = kve.Value.(*ast.CompositeLit); !ok {
					log.Fatal("Value shoud be a composite lit" +
						fset.Position(cl.Pos()).String())
				}
				// get the array of Link (as definition only)
				for _, expr := range cl.Elts {

					switch exp := expr.(type) {
					case *ast.CompositeLit: // this is a definition
						var link Link
						link.Unmarshall(modelPkg, exp, fset)

						classshape.Links = append(classshape.Links, &link)
					default:
						log.Fatal("Value shoud be a composite lit" +
							fset.Position(kve.Pos()).String())
					}
				}
			case "Fields":
				var cl *ast.CompositeLit
				var ok bool
				if cl, ok = kve.Value.(*ast.CompositeLit); !ok {
					log.Fatal("Value shoud be a composite lit" +
						fset.Position(cl.Pos()).String())
				}
				// get the array of Field (as definition only)
				for _, expr := range cl.Elts {

					switch exp := expr.(type) {
					case *ast.CompositeLit: // this is a definition
						var field Field
						field.Unmarshall(modelPkg, exp, fset)

						classshape.Fields = append(classshape.Fields, &field)
					default:
						log.Fatal("Value shoud be a composite lit" +
							fset.Position(kve.Pos()).String())
					}
				}
			case "Width":
				switch val := kve.Value.(type) {
				case *ast.BasicLit:
					if val.Kind != token.INT && val.Kind != token.FLOAT {
						log.Panic("Width is not an INT or a FLOAT" +
							fset.Position(val.Pos()).String())
					}
					X, err := strconv.ParseFloat(val.Value, 64)
					if err != nil {
						log.Panic("Problem parsing float" + fset.Position(val.Pos()).String())
					}
					classshape.Width = X
				}
			case "Heigth":
				switch val := kve.Value.(type) {
				case *ast.BasicLit:
					if val.Kind != token.INT && val.Kind != token.FLOAT {
						log.Panic("Width is not an INT or a FLOAT" +
							fset.Position(val.Pos()).String())
					}
					X, err := strconv.ParseFloat(val.Value, 64)
					if err != nil {
						log.Panic("Problem parsing float" + fset.Position(val.Pos()).String())
					}
					classshape.Heigth = X
				}
			default:
				log.Fatal("Key should be Position, Struct, Field or Link" +
					fset.Position(kve.Pos()).String())
			}
		}
	}

	// init width with default value
	if classshape.Width == 0.0 {
		classshape.Width = ClassshapeDefaultWidth
	}
	// init Heigth with default value
	if classshape.Heigth == 0.0 {
		classshape.Heigth = ClassshapeDefaultHeigth + 15.0*float64(len(classshape.Fields))
	}

	ClassshapeStore[classshape.Name] = classshape
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (classshape *Classshape) SerializeToStage() {

	classshape.Stage()

	classshape.Position.Stage()
	classshape.Reference.Stage()

	for _, link := range classshape.Links {
		link.SerializeToStage()
	}

	for _, field := range classshape.Fields {
		field.SerializeToStage()
	}
}

func (classshape *Classshape) SerializeToUnstage() {

	classshape.Unstage()

	classshape.Position.Unstage()
	classshape.Reference.Unstage()

	for _, link := range classshape.Links {
		link.SerializeToUnstage()
	}

	for _, field := range classshape.Fields {
		field.Unstage()
	}
}
