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
	model "github.com/fullstack-lang/gong/test/test2/go/models/x/models"
)

var (
	_ = time.Hour
	_ = slices.Index[[]int, int]
	_ = sort.Slice
	_ = strconv.Itoa
)

// StageSet coordinates multiple stages across packages
type StageSet struct {
	Stage *Stage
	XStage *x.Stage
	YStage *y.Stage
	ModelStage *model.Stage
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.ModelStage != nil {
		stageSet.ModelStage.Commit()
	}
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
	if stageSet.ModelStage != nil {
		stageSet.ModelStage.Checkout()
	}
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
	if stageSet.ModelStage != nil {
		stageSet.ModelStage.Reset()
	}
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
	if stageSet.ModelStage != nil {
		stageSet.ModelStage.Clean()
	}
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
	if stageSet.ModelStage != nil {
		stageSet.ModelStage.ComputeReverseMaps()
	}
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
	if stageSet.ModelStage != nil {
		stageSet.ModelStage.ComputeInstancesNb()
	}
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
	if stageSet.ModelStage != nil {
		stageSet.ModelStage.ComputeReferenceAndOrders()
	}
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
	subPath_ModelStage := "model"
	if path != "" {
	subPath_ModelStage = path + "_model"
	}
	stageSet.ModelStage = model.NewStage(subPath_ModelStage)
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
	subPath_ModelStage := "model"
	if stage != nil && stage.GetName() != "" {
	subPath_ModelStage = stage.GetName() + "_model"
	}
	stageSet.ModelStage = model.NewStage(subPath_ModelStage)
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
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stageSet.ModelStage != nil {
		for _, submodel := range __gong__sortStageSetInstances(stageSet.ModelStage.SubModels, stageSet.ModelStage.SubModel_stagedOrder) {
			if lastStageDecl != "ModelStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ModelStage"
			}
			submodelIdent := "__model" + submodel.GongGetIdentifier(stageSet.ModelStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&model.SubModel{Name: %s}).Stage(stageSet.ModelStage)", submodelIdent, __gong__toRawStringLiteral(submodel.Name)))
			if lastStageVal != "ModelStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ModelStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", submodelIdent, __gong__toRawStringLiteral(submodel.Name)))
		}
	}
	if stageSet.YStage != nil {
		for _, y := range __gong__sortStageSetInstances(stageSet.YStage.Ys, stageSet.YStage.Y_stagedOrder) {
			if lastStageDecl != "YStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "YStage"
			}
			yIdent := "__y" + y.GongGetIdentifier(stageSet.YStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&y.Y{Name: %s}).Stage(stageSet.YStage)", yIdent, __gong__toRawStringLiteral(y.Name)))
			if lastStageVal != "YStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "YStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", yIdent, __gong__toRawStringLiteral(y.Name)))
			if y.SubModel != nil {
				if lastStagePtr != "YStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "YStage"
				}
				targetIdent := "__model" + y.SubModel.GongGetIdentifier(stageSet.ModelStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubModel = %s", yIdent, targetIdent))
			}
		}
	}
	if stageSet.XStage != nil {
		for _, x := range __gong__sortStageSetInstances(stageSet.XStage.Xs, stageSet.XStage.X_stagedOrder) {
			if lastStageDecl != "XStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "XStage"
			}
			xIdent := "__x" + x.GongGetIdentifier(stageSet.XStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&x.X{Name: %s}).Stage(stageSet.XStage)", xIdent, __gong__toRawStringLiteral(x.Name)))
			if lastStageVal != "XStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "XStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xIdent, __gong__toRawStringLiteral(x.Name)))
			if x.Y != nil {
				if lastStagePtr != "XStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "XStage"
				}
				targetIdent := "__y" + x.Y.GongGetIdentifier(stageSet.YStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Y = %s", xIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a := range __gong__sortStageSetInstances(stageSet.Stage.As, stageSet.Stage.A_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aIdent := "__models" + a.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A{Name: %s}).Stage(stageSet.Stage)", aIdent, __gong__toRawStringLiteral(a.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aIdent, __gong__toRawStringLiteral(a.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumberField = %d", aIdent, a.NumberField))
			values.WriteString(fmt.Sprintf("\n\t%s.Foo = %d", aIdent, a.Foo))
			values.WriteString(fmt.Sprintf("\n\t%s.Bar = %f", aIdent, a.Bar))
			values.WriteString(fmt.Sprintf("\n\t%s.Zorgh = %s", aIdent, __gong__toRawStringLiteral(a.Zorgh)))
			if a.B != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + a.B.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.B = %s", aIdent, targetIdent))
			}
			for _, elem := range a.Bs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bs = append(%s.Bs, %s)", aIdent, aIdent, targetIdent))
			}
			if a.X != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__x" + a.X.GongGetIdentifier(stageSet.XStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.X = %s", aIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, b := range __gong__sortStageSetInstances(stageSet.Stage.Bs, stageSet.Stage.B_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bIdent := "__models" + b.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.B{Name: %s}).Stage(stageSet.Stage)", bIdent, __gong__toRawStringLiteral(b.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bIdent, __gong__toRawStringLiteral(b.Name)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
	model "github.com/fullstack-lang/gong/test/test2/go/models/x/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
	_ *x.Stage
	_ *y.Stage
	_ *model.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *models.StageSet) {

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

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/test/test2/go/models/x/models":
			aliasToCanonical[alias] = "model"
		case "github.com/fullstack-lang/gong/test/test2/go/models/y":
			aliasToCanonical[alias] = "y"
		case "github.com/fullstack-lang/gong/test/test2/go/models/x":
			aliasToCanonical[alias] = "x"
		case "github.com/fullstack-lang/gong/test/test2/go/models":
			aliasToCanonical[alias] = "models"
		}
	}

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

					if canonical, ok := aliasToCanonical[pkgAlias]; ok {
						pkgAlias = canonical
					}

					switch pkgAlias {
			case "model":
				switch typeName {
				case "SubModel":
					identifierMap[ident.Name] = __gong__stageSetInit(new(model.SubModel), stageSet.ModelStage, ident.Name, instanceName, preserveOrder)
				}
			case "y":
				switch typeName {
				case "Y":
					identifierMap[ident.Name] = __gong__stageSetInit(new(y.Y), stageSet.YStage, ident.Name, instanceName, preserveOrder)
				}
			case "x":
				switch typeName {
				case "X":
					identifierMap[ident.Name] = __gong__stageSetInit(new(x.X), stageSet.XStage, ident.Name, instanceName, preserveOrder)
				}
			case "models":
				switch typeName {
				case "A":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "B":
					identifierMap[ident.Name] = __gong__stageSetInit(new(B), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *model.SubModel:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *y.Y:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SubModel":
						__gong__assignPointer(&inst.SubModel, rhs, identifierMap)
					}
				case *x.X:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Y":
						__gong__assignPointer(&inst.Y, rhs, identifierMap)
					}
				case *A:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumberField":
						inst.NumberField = GongExtractInt(rhs)
					case "B":
						__gong__assignPointer(&inst.B, rhs, identifierMap)
					case "Bs":
						__gong__assignSliceOfPointers(&inst.Bs, rhs, identifierMap)
					case "X":
						__gong__assignPointer(&inst.X, rhs, identifierMap)
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

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
}

func __gong__stageSetInit[P interface {
	SetName(string)
	StageVoid(S)
	StagePreserveOrder(S, uint)
}, S any](instance P, stage S, identifier string, instanceName string, preserveOrder bool) any {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if order, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifier", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, order)
		}
	}
	return instance
}

func __gong__assignPointer[T any](targetPtr **T, rhs ast.Expr, identifierMap map[string]any) {
	if rIdent, ok := rhs.(*ast.Ident); ok {
		if rIdent.Name == "nil" {
			*targetPtr = nil
			return
		}
		if target, ok := identifierMap[rIdent.Name]; ok {
			if typedTarget, ok := target.(*T); ok {
				*targetPtr = typedTarget
			}
		}
	}
}

func __gong__assignSliceOfPointers[T any](slice *[]*T, rhs ast.Expr, identifierMap map[string]any) {
	if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
		if rIdent, ok := call.Args[1].(*ast.Ident); ok {
			if target, ok := identifierMap[rIdent.Name]; ok {
				if typedTarget, ok := target.(*T); ok {
					*slice = append(*slice, typedTarget)
				}
			}
		}
	}
}
