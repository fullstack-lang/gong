// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
)

var (
	_ = time.Hour
	_ = slices.Index[[]int, int]
	_ = sort.Slice
	_ = strconv.Itoa
)

// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.YStage != nil {
		stageSet.YStage.Commit()
	}
	if stageSet.XStage != nil {
		stageSet.XStage.Commit()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.YStage != nil {
		stageSet.YStage.Checkout()
	}
	if stageSet.XStage != nil {
		stageSet.XStage.Checkout()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.YStage != nil {
		stageSet.YStage.Reset()
	}
	if stageSet.XStage != nil {
		stageSet.XStage.Reset()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.YStage != nil {
		stageSet.YStage.Clean()
	}
	if stageSet.XStage != nil {
		stageSet.XStage.Clean()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.YStage != nil {
		stageSet.YStage.ComputeReverseMaps()
	}
	if stageSet.XStage != nil {
		stageSet.XStage.ComputeReverseMaps()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.YStage != nil {
		stageSet.YStage.ComputeInstancesNb()
	}
	if stageSet.XStage != nil {
		stageSet.XStage.ComputeInstancesNb()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.YStage != nil {
		stageSet.YStage.ComputeReferenceAndOrders()
	}
	if stageSet.XStage != nil {
		stageSet.XStage.ComputeReferenceAndOrders()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	subPath_XStage := "x"
	if path != "" {
	subPath_XStage = path + "_x"
	}
	stageSet.XStage = x.NewStage(subPath_XStage)
	subPath_YStage := "y"
	if path != "" {
	subPath_YStage = path + "_y"
	}
	stageSet.YStage = y.NewStage(subPath_YStage)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
	subPath_XStage := "x"
	if stage != nil && stage.GetName() != "" {
	subPath_XStage = stage.GetName() + "_x"
	}
	stageSet.XStage = x.NewStage(subPath_XStage)
	subPath_YStage := "y"
	if stage != nil && stage.GetName() != "" {
	subPath_YStage = stage.GetName() + "_y"
	}
	stageSet.YStage = y.NewStage(subPath_YStage)
	return stageSet
}

// GetProbeSplitStageName returns the split stage name for the StageSet probe
func (stageSet *StageSet) GetProbeSplitStageName() string {
	if stageSet.Stage != nil {
		return stageSet.Stage.GetProbeSplitStageName() + "_stageset"
	}
	return "stageset_probe_split"
}

// MarshallFile marshalls all stages into a file
func (stageSet *StageSet) MarshallFile(filename, packageName string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stageSet.Marshall(file, packageName)
}

// Marshall marshalls all stages into an open file
func (stageSet *StageSet) Marshall(file *os.File, packageName string) {
	res, err := stageSet.MarshallToString(packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}
	fmt.Fprintln(file, res)
}

// MarshallToString marshalls all stages into a Go code string
func (stageSet *StageSet) MarshallToString(packageName string) (res string, err error) {
	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder

	if stageSet.YStage != nil {
		yOrdered := []*y.Y{}
		for y := range stageSet.YStage.Ys {
			yOrdered = append(yOrdered, y)
		}
		sort.Slice(yOrdered, func(i, j int) bool {
			return stageSet.YStage.Y_stagedOrder[yOrdered[i]] < stageSet.YStage.Y_stagedOrder[yOrdered[j]]
		})
		for _, y := range yOrdered {
			yIdent := "__stage_2" + y.GongGetIdentifier(stageSet.YStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_2__.Y{Name: %s}).Stage(stageSet.YStage)", yIdent, __gong__toRawStringLiteral(y.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", yIdent, __gong__toRawStringLiteral(y.Name)))
		}
	}
	if stageSet.XStage != nil {
		xOrdered := []*x.X{}
		for x := range stageSet.XStage.Xs {
			xOrdered = append(xOrdered, x)
		}
		sort.Slice(xOrdered, func(i, j int) bool {
			return stageSet.XStage.X_stagedOrder[xOrdered[i]] < stageSet.XStage.X_stagedOrder[xOrdered[j]]
		})
		for _, x := range xOrdered {
			xIdent := "__stage_1" + x.GongGetIdentifier(stageSet.XStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_1__.X{Name: %s}).Stage(stageSet.XStage)", xIdent, __gong__toRawStringLiteral(x.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xIdent, __gong__toRawStringLiteral(x.Name)))
			if x.Y != nil {
				targetIdent := "__stage_2" + x.Y.GongGetIdentifier(stageSet.YStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Y = %s", xIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		aOrdered := []*A{}
		for a := range stageSet.Stage.As {
			aOrdered = append(aOrdered, a)
		}
		sort.Slice(aOrdered, func(i, j int) bool {
			return stageSet.Stage.A_stagedOrder[aOrdered[i]] < stageSet.Stage.A_stagedOrder[aOrdered[j]]
		})
		for _, a := range aOrdered {
			aIdent := "__stage_0" + a.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.A{Name: %s}).Stage(stageSet.Stage)", aIdent, __gong__toRawStringLiteral(a.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aIdent, __gong__toRawStringLiteral(a.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumberField = %d", aIdent, a.NumberField))
			values.WriteString(fmt.Sprintf("\n\t%s.Foo = %d", aIdent, a.Foo))
			values.WriteString(fmt.Sprintf("\n\t%s.Bar = %f", aIdent, a.Bar))
			values.WriteString(fmt.Sprintf("\n\t%s.Zorgh = %s", aIdent, __gong__toRawStringLiteral(a.Zorgh)))
			if a.B != nil {
				targetIdent := "__stage_0" + a.B.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.B = %s", aIdent, targetIdent))
			}
			for _, elem := range a.Bs {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bs = append(%s.Bs, %s)", aIdent, aIdent, targetIdent))
			}
			if a.X != nil {
				targetIdent := "__stage_1" + a.X.GongGetIdentifier(stageSet.XStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.X = %s", aIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		bOrdered := []*B{}
		for b := range stageSet.Stage.Bs {
			bOrdered = append(bOrdered, b)
		}
		sort.Slice(bOrdered, func(i, j int) bool {
			return stageSet.Stage.B_stagedOrder[bOrdered[i]] < stageSet.Stage.B_stagedOrder[bOrdered[j]]
		})
		for _, b := range bOrdered {
			bIdent := "__stage_0" + b.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.B{Name: %s}).Stage(stageSet.Stage)", bIdent, __gong__toRawStringLiteral(b.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bIdent, __gong__toRawStringLiteral(b.Name)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/test/test2/go/models"
	__stage_1__ "github.com/fullstack-lang/gong/test/test2/go/models/x"
	__stage_2__ "github.com/fullstack-lang/gong/test/test2/go/models/y"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *__stage_0__.Stage
	_ *__stage_1__.Stage
	_ *__stage_2__.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *__stage_0__.StageSet) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, packageName, declarations.String(), values.String(), pointers.String())

	return res, nil
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file into stageSet
func (stageSet *StageSet) ParseAstFile(pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist " + pathToFile + " ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file into stageSet
func (stageSet *StageSet) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New("Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string into stageSet
func (stageSet *StageSet) ParseAstString(blob string, preserveOrder bool) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", blob, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances into stageSet
func (stageSet *StageSet) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	identifierMap := make(map[string]any)

	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var pkgAlias string
					var typeName string
					var instanceName string

					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								if pkgId, ok := selExpr.X.(*ast.Ident); ok {
									pkgAlias = pkgId.Name
								}
								typeName = selExpr.Sel.Name
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					switch pkgAlias {
			case "__stage_2__":
				switch typeName {
				case "Y":
					if !preserveOrder {
						inst := (&y.Y{Name: instanceName}).Stage(stageSet.YStage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(y.Y)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.YStage, uint(order))
						identifierMap[ident.Name] = inst
					}
				}
			case "__stage_1__":
				switch typeName {
				case "X":
					if !preserveOrder {
						inst := (&x.X{Name: instanceName}).Stage(stageSet.XStage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(x.X)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.XStage, uint(order))
						identifierMap[ident.Name] = inst
					}
				}
			case "__stage_0__":
				switch typeName {
				case "A":
					if !preserveOrder {
						inst := (&A{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "B":
					if !preserveOrder {
						inst := (&B{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(B)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *y.Y:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *x.X:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Y":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*y.Y); ok {
									inst.Y = typedTarget
								}
							}
						}
					}
				case *A:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumberField":
						inst.NumberField = GongExtractInt(rhs)
					case "B":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*B); ok {
									inst.B = typedTarget
								}
							}
						}
					case "Bs":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*B); ok {
										inst.Bs = append(inst.Bs, typedTarget)
									}
								}
							}
						}
					case "X":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*x.X); ok {
									inst.X = typedTarget
								}
							}
						}
					case "Foo":
						inst.Foo = GongExtractInt(rhs)
					case "Bar":
						inst.Bar = GongExtractFloat(rhs)
					case "Zorgh":
						inst.Zorgh = GongExtractString(rhs)
					}
				case *B:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}
