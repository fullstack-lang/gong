package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"sort"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// Umlsc diagram struct store a class diagram
// temporary here
// swagger:model Umlsc
type Umlsc struct {
	Name string

	// this is the memory model (and not the "memory motel" of the Rolling Stones)
	// it is not ignored by swagger because it is used by the angular model
	States []*UmlState

	// in this version, only one state is active mong the states
	// (with the embedded states version, that might change)
	Activestate string
}

// MarshallAsVariable transform the class diagram into a var
func (umlsc *Umlsc) MarshallAsVariable(file *os.File) error {

	fmt.Fprintf(file, "var %s uml.Umlsc = uml.Umlsc{\n", umlsc.Name)

	if umlsc.Activestate != "" {
		fmt.Fprintf(file, "\tActivestate: string(models.%s),\n", umlsc.Activestate)
	}

	// sort States
	sort.Slice(umlsc.States[:], func(i, j int) bool {
		return umlsc.States[i].Name < umlsc.States[j].Name
	})
	fmt.Fprintf(file, "\tStates: []*uml.UmlState{\n")
	for _, state := range umlsc.States {
		state.Marshall(file, 2)
		fmt.Fprintf(file, ",\n")
	}
	fmt.Fprintf(file, "\t},\n")

	fmt.Fprintf(file, "}\n\n")
	return nil
}

// UmlscMap is a Map of all Umlsc via their Name
type UmlscMap map[string]*Umlsc

// UmlscStore is a handy UmlscMap
var UmlscStore UmlscMap = make(map[string]*Umlsc, 0)

// Unmarshall updates a umlsc values from an ast.Epr
// and appends it to the UmlscStore
func (umlsc *Umlsc) Unmarshall(modelPkg *gong_models.ModelPkg, expr ast.Expr, fset *token.FileSet) {

	// expression should be a composite literal expression
	// "uml.Umlsc{
	//   	States: []*uml.UmlState{ ...

	if cl, ok := expr.(*ast.CompositeLit); ok {

		// parse all KeyValues of the Umlsc
		for _, elt := range cl.Elts {
			if kve, ok := elt.(*ast.KeyValueExpr); !ok {
				log.Panic("Expression should be a struct key" +
					fset.Position(kve.Pos()).String())
			} else {

				if key, ok := kve.Key.(*ast.Ident); !ok {
					log.Panic("Key shoud be an ident" +
						fset.Position(key.Pos()).String())
				} else {
					switch key.Name {
					case "Name":
						var bl *ast.BasicLit
						if bl, ok = kve.Value.(*ast.BasicLit); !ok {
							// TODO deal with recursive Binary Expressions "titi"+"toto"+"tata"
							log.Panic("Expecting a basic lit " + fset.Position(kve.Value.Pos()).String())
						}
						name := bl.Value
						name = strings.TrimPrefix(name, "\"")
						name = strings.TrimSuffix(name, "\"")
						name = strings.TrimRight(name, "\"")
						// need to remove prefix and trailling "
						umlsc.Name = name
					case "States":
						var cl *ast.CompositeLit
						var ok bool
						if cl, ok = kve.Value.(*ast.CompositeLit); !ok {
							log.Panic("Value shoud be a composite lit" +
								fset.Position(kve.Pos()).String())
						}
						// get the array of shapes (either as definition or as reference)
						for _, expr := range cl.Elts {

							var state *UmlState
							switch exp := expr.(type) {
							case *ast.CompositeLit: // this is a definition
								var _state UmlState
								state = &_state
								state.Unmarshall(exp, fset)
							default:
								log.Panic("Value shoud be a composite lit or a unary" +
									fset.Position(kve.Pos()).String())
							}

							umlsc.States = append(umlsc.States, state)
						}
					case "Activestate":
						var ce *ast.CallExpr
						var ident *ast.Ident
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
						umlsc.Activestate = se.Sel.Name

						if ident, ok = se.X.(*ast.Ident); !ok {
							log.Panic("Expecting an ident expression in X" + fset.Position(se.X.Pos()).String())
						}
						if ident.Name != "models" {
							log.Panic("Expecting an ident `models`" + fset.Position(ident.Pos()).String())
						}

					default:
						log.Panic("Unknow state " + key.Name + " " +
							fset.Position(key.Pos()).String())
					}
				}
			}
		}
	}
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (umlsc *Umlsc) SerializeToStage() {

	umlsc.Stage()

	for _, state := range umlsc.States {
		state.SerializeToStage()
	}
}
