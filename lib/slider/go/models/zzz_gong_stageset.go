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
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
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

	if stageSet.Stage != nil {
		checkboxOrdered := []*Checkbox{}
		for checkbox := range stageSet.Stage.Checkboxs {
			checkboxOrdered = append(checkboxOrdered, checkbox)
		}
		sort.Slice(checkboxOrdered, func(i, j int) bool {
			return stageSet.Stage.Checkbox_stagedOrder[checkboxOrdered[i]] < stageSet.Stage.Checkbox_stagedOrder[checkboxOrdered[j]]
		})
		for _, checkbox := range checkboxOrdered {
			checkboxIdent := "__stage_0" + checkbox.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Checkbox{Name: %s}).Stage(stageSet.Stage)", checkboxIdent, __gong__toRawStringLiteral(checkbox.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkboxIdent, __gong__toRawStringLiteral(checkbox.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ValueBool = %t", checkboxIdent, checkbox.ValueBool))
			values.WriteString(fmt.Sprintf("\n\t%s.LabelForTrue = %s", checkboxIdent, __gong__toRawStringLiteral(checkbox.LabelForTrue)))
			values.WriteString(fmt.Sprintf("\n\t%s.LabelForFalse = %s", checkboxIdent, __gong__toRawStringLiteral(checkbox.LabelForFalse)))
		}
	}
	if stageSet.Stage != nil {
		groupOrdered := []*Group{}
		for group := range stageSet.Stage.Groups {
			groupOrdered = append(groupOrdered, group)
		}
		sort.Slice(groupOrdered, func(i, j int) bool {
			return stageSet.Stage.Group_stagedOrder[groupOrdered[i]] < stageSet.Stage.Group_stagedOrder[groupOrdered[j]]
		})
		for _, group := range groupOrdered {
			groupIdent := "__stage_0" + group.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Group{Name: %s}).Stage(stageSet.Stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupIdent, __gong__toRawStringLiteral(group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Percentage = %f", groupIdent, group.Percentage))
			for _, elem := range group.Sliders {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sliders = append(%s.Sliders, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Checkboxes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Checkboxes = append(%s.Checkboxes, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		layoutOrdered := []*Layout{}
		for layout := range stageSet.Stage.Layouts {
			layoutOrdered = append(layoutOrdered, layout)
		}
		sort.Slice(layoutOrdered, func(i, j int) bool {
			return stageSet.Stage.Layout_stagedOrder[layoutOrdered[i]] < stageSet.Stage.Layout_stagedOrder[layoutOrdered[j]]
		})
		for _, layout := range layoutOrdered {
			layoutIdent := "__stage_0" + layout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Layout{Name: %s}).Stage(stageSet.Stage)", layoutIdent, __gong__toRawStringLiteral(layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", layoutIdent, __gong__toRawStringLiteral(layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCustomGutterSize = %t", layoutIdent, layout.IsWithCustomGutterSize))
			values.WriteString(fmt.Sprintf("\n\t%s.GutterSize = %f", layoutIdent, layout.GutterSize))
			for _, elem := range layout.Groups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", layoutIdent, layoutIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		sliderOrdered := []*Slider{}
		for slider := range stageSet.Stage.Sliders {
			sliderOrdered = append(sliderOrdered, slider)
		}
		sort.Slice(sliderOrdered, func(i, j int) bool {
			return stageSet.Stage.Slider_stagedOrder[sliderOrdered[i]] < stageSet.Stage.Slider_stagedOrder[sliderOrdered[j]]
		})
		for _, slider := range sliderOrdered {
			sliderIdent := "__stage_0" + slider.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Slider{Name: %s}).Stage(stageSet.Stage)", sliderIdent, __gong__toRawStringLiteral(slider.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sliderIdent, __gong__toRawStringLiteral(slider.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsFloat64 = %t", sliderIdent, slider.IsFloat64))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInt = %t", sliderIdent, slider.IsInt))
			values.WriteString(fmt.Sprintf("\n\t%s.MinInt = %d", sliderIdent, slider.MinInt))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxInt = %d", sliderIdent, slider.MaxInt))
			values.WriteString(fmt.Sprintf("\n\t%s.StepInt = %d", sliderIdent, slider.StepInt))
			values.WriteString(fmt.Sprintf("\n\t%s.ValueInt = %d", sliderIdent, slider.ValueInt))
			values.WriteString(fmt.Sprintf("\n\t%s.MinFloat64 = %f", sliderIdent, slider.MinFloat64))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxFloat64 = %f", sliderIdent, slider.MaxFloat64))
			values.WriteString(fmt.Sprintf("\n\t%s.StepFloat64 = %f", sliderIdent, slider.StepFloat64))
			values.WriteString(fmt.Sprintf("\n\t%s.ValueFloat64 = %f", sliderIdent, slider.ValueFloat64))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDisabled = %t", sliderIdent, slider.IsDisabled))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/lib/slider/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *__stage_0__.Stage
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
			case "__stage_0__":
				switch typeName {
				case "Checkbox":
					if !preserveOrder {
						inst := (&Checkbox{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Checkbox)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Group":
					if !preserveOrder {
						inst := (&Group{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Group)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Layout":
					if !preserveOrder {
						inst := (&Layout{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Layout)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Slider":
					if !preserveOrder {
						inst := (&Slider{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Slider)
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
				case *Checkbox:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ValueBool":
						inst.ValueBool = GongExtractBool(rhs)
					case "LabelForTrue":
						inst.LabelForTrue = GongExtractString(rhs)
					case "LabelForFalse":
						inst.LabelForFalse = GongExtractString(rhs)
					}
				case *Group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Percentage":
						inst.Percentage = GongExtractFloat(rhs)
					case "Sliders":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Slider); ok {
										inst.Sliders = append(inst.Sliders, typedTarget)
									}
								}
							}
						}
					case "Checkboxes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Checkbox); ok {
										inst.Checkboxes = append(inst.Checkboxes, typedTarget)
									}
								}
							}
						}
					}
				case *Layout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Groups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
										inst.Groups = append(inst.Groups, typedTarget)
									}
								}
							}
						}
					case "IsWithCustomGutterSize":
						inst.IsWithCustomGutterSize = GongExtractBool(rhs)
					case "GutterSize":
						inst.GutterSize = GongExtractFloat(rhs)
					}
				case *Slider:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsFloat64":
						inst.IsFloat64 = GongExtractBool(rhs)
					case "IsInt":
						inst.IsInt = GongExtractBool(rhs)
					case "MinInt":
						inst.MinInt = GongExtractInt(rhs)
					case "MaxInt":
						inst.MaxInt = GongExtractInt(rhs)
					case "StepInt":
						inst.StepInt = GongExtractInt(rhs)
					case "ValueInt":
						inst.ValueInt = GongExtractInt(rhs)
					case "MinFloat64":
						inst.MinFloat64 = GongExtractFloat(rhs)
					case "MaxFloat64":
						inst.MaxFloat64 = GongExtractFloat(rhs)
					case "StepFloat64":
						inst.StepFloat64 = GongExtractFloat(rhs)
					case "ValueFloat64":
						inst.ValueFloat64 = GongExtractFloat(rhs)
					case "IsDisabled":
						inst.IsDisabled = GongExtractBool(rhs)
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
