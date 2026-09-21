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
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stageSet.Stage != nil {
		buttonOrdered := []*Button{}
		for button := range stageSet.Stage.Buttons {
			buttonOrdered = append(buttonOrdered, button)
		}
		sort.Slice(buttonOrdered, func(i, j int) bool {
			return stageSet.Stage.Button_stagedOrder[buttonOrdered[i]] < stageSet.Stage.Button_stagedOrder[buttonOrdered[j]]
		})
		for _, button := range buttonOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			buttonIdent := "__models" + button.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Button{Name: %s}).Stage(stageSet.Stage)", buttonIdent, __gong__toRawStringLiteral(button.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", buttonIdent, __gong__toRawStringLiteral(button.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", buttonIdent, __gong__toRawStringLiteral(button.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.Icon = %s", buttonIdent, __gong__toRawStringLiteral(button.Icon)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDisabled = %t", buttonIdent, button.IsDisabled))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", buttonIdent, __gong__toRawStringLiteral(string(button.Color))))
			values.WriteString(fmt.Sprintf("\n\t%s.MatButtonType = %s", buttonIdent, __gong__toRawStringLiteral(string(button.MatButtonType))))
			values.WriteString(fmt.Sprintf("\n\t%s.MatButtonAppearance = %s", buttonIdent, __gong__toRawStringLiteral(string(button.MatButtonAppearance))))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", buttonIdent, button.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", buttonIdent, __gong__toRawStringLiteral(button.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = %s", buttonIdent, __gong__toRawStringLiteral(string(button.ToolTipPosition))))
		}
	}
	if stageSet.Stage != nil {
		buttontoggleOrdered := []*ButtonToggle{}
		for buttontoggle := range stageSet.Stage.ButtonToggles {
			buttontoggleOrdered = append(buttontoggleOrdered, buttontoggle)
		}
		sort.Slice(buttontoggleOrdered, func(i, j int) bool {
			return stageSet.Stage.ButtonToggle_stagedOrder[buttontoggleOrdered[i]] < stageSet.Stage.ButtonToggle_stagedOrder[buttontoggleOrdered[j]]
		})
		for _, buttontoggle := range buttontoggleOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			buttontoggleIdent := "__models" + buttontoggle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ButtonToggle{Name: %s}).Stage(stageSet.Stage)", buttontoggleIdent, __gong__toRawStringLiteral(buttontoggle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", buttontoggleIdent, __gong__toRawStringLiteral(buttontoggle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", buttontoggleIdent, __gong__toRawStringLiteral(buttontoggle.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.Icon = %s", buttontoggleIdent, __gong__toRawStringLiteral(buttontoggle.Icon)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDisabled = %t", buttontoggleIdent, buttontoggle.IsDisabled))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", buttontoggleIdent, buttontoggle.IsChecked))
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
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupIdent := "__models" + group.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group{Name: %s}).Stage(stageSet.Stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupIdent, __gong__toRawStringLiteral(group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Percentage = %f", groupIdent, group.Percentage))
			values.WriteString(fmt.Sprintf("\n\t%s.NbColumns = %d", groupIdent, group.NbColumns))
			for _, elem := range group.Buttons {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Buttons = append(%s.Buttons, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		grouptoogleOrdered := []*GroupToogle{}
		for grouptoogle := range stageSet.Stage.GroupToogles {
			grouptoogleOrdered = append(grouptoogleOrdered, grouptoogle)
		}
		sort.Slice(grouptoogleOrdered, func(i, j int) bool {
			return stageSet.Stage.GroupToogle_stagedOrder[grouptoogleOrdered[i]] < stageSet.Stage.GroupToogle_stagedOrder[grouptoogleOrdered[j]]
		})
		for _, grouptoogle := range grouptoogleOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			grouptoogleIdent := "__models" + grouptoogle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GroupToogle{Name: %s}).Stage(stageSet.Stage)", grouptoogleIdent, __gong__toRawStringLiteral(grouptoogle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", grouptoogleIdent, __gong__toRawStringLiteral(grouptoogle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Percentage = %f", grouptoogleIdent, grouptoogle.Percentage))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSingleSelector = %t", grouptoogleIdent, grouptoogle.IsSingleSelector))
			for _, elem := range grouptoogle.ButtonToggles {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ButtonToggles = append(%s.ButtonToggles, %s)", grouptoogleIdent, grouptoogleIdent, targetIdent))
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
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			layoutIdent := "__models" + layout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Layout{Name: %s}).Stage(stageSet.Stage)", layoutIdent, __gong__toRawStringLiteral(layout.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", layoutIdent, __gong__toRawStringLiteral(layout.Name)))
			for _, elem := range layout.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", layoutIdent, layoutIdent, targetIdent))
			}
			for _, elem := range layout.GroupToogles {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupToogles = append(%s.GroupToogles, %s)", layoutIdent, layoutIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/button/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
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
		case "github.com/fullstack-lang/gong/lib/button/go/models":
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
			case "models":
				switch typeName {
				case "Button":
					if !preserveOrder {
						inst := (&Button{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Button)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ButtonToggle":
					if !preserveOrder {
						inst := (&ButtonToggle{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ButtonToggle)
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
				case "GroupToogle":
					if !preserveOrder {
						inst := (&GroupToogle{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GroupToogle)
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
				case *Button:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "Icon":
						inst.Icon = GongExtractString(rhs)
					case "IsDisabled":
						inst.IsDisabled = GongExtractBool(rhs)
					case "Color":
						inst.Color = MatButtonPaletteType(GongExtractString(rhs))
					case "MatButtonType":
						inst.MatButtonType = MatButtonType(GongExtractString(rhs))
					case "MatButtonAppearance":
						inst.MatButtonAppearance = MatButtonAppearance(GongExtractString(rhs))
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "ToolTipPosition":
						inst.ToolTipPosition = ToolTipPositionEnum(GongExtractString(rhs))
					}
				case *ButtonToggle:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "Icon":
						inst.Icon = GongExtractString(rhs)
					case "IsDisabled":
						inst.IsDisabled = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					}
				case *Group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Percentage":
						inst.Percentage = GongExtractFloat(rhs)
					case "Buttons":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Button); ok {
										inst.Buttons = append(inst.Buttons, typedTarget)
									}
								}
							}
						}
					case "NbColumns":
						inst.NbColumns = GongExtractInt(rhs)
					}
				case *GroupToogle:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Percentage":
						inst.Percentage = GongExtractFloat(rhs)
					case "ButtonToggles":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ButtonToggle); ok {
										inst.ButtonToggles = append(inst.ButtonToggles, typedTarget)
									}
								}
							}
						}
					case "IsSingleSelector":
						inst.IsSingleSelector = GongExtractBool(rhs)
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
					case "GroupToogles":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GroupToogle); ok {
										inst.GroupToogles = append(inst.GroupToogles, typedTarget)
									}
								}
							}
						}
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
