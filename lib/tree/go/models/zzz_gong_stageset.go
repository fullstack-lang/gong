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
		for _, button := range __gong__sortStageSetInstances(stageSet.Stage.Buttons, stageSet.Stage.Button_stagedOrder) {
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
			values.WriteString(fmt.Sprintf("\n\t%s.Icon = %s", buttonIdent, __gong__toRawStringLiteral(button.Icon)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDisabled = %t", buttonIdent, button.IsDisabled))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", buttonIdent, button.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", buttonIdent, __gong__toRawStringLiteral(button.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = %s", buttonIdent, __gong__toRawStringLiteral(string(button.ToolTipPosition))))
			values.WriteString(fmt.Sprintf("\n\t%s.ClientOnX = %f", buttonIdent, button.ClientOnX))
			values.WriteString(fmt.Sprintf("\n\t%s.ClientOnY = %f", buttonIdent, button.ClientOnY))
			if button.SVGIcon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + button.SVGIcon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SVGIcon = %s", buttonIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, menu := range __gong__sortStageSetInstances(stageSet.Stage.Menus, stageSet.Stage.Menu_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			menuIdent := "__models" + menu.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Menu{Name: %s}).Stage(stageSet.Stage)", menuIdent, __gong__toRawStringLiteral(menu.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", menuIdent, __gong__toRawStringLiteral(menu.Name)))
			for _, elem := range menu.Buttons {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Buttons = append(%s.Buttons, %s)", menuIdent, menuIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, node := range __gong__sortStageSetInstances(stageSet.Stage.Nodes, stageSet.Stage.Node_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			nodeIdent := "__models" + node.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Node{Name: %s}).Stage(stageSet.Stage)", nodeIdent, __gong__toRawStringLiteral(node.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nodeIdent, __gong__toRawStringLiteral(node.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithPrefix = %t", nodeIdent, node.IsWithPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.Prefix = %s", nodeIdent, __gong__toRawStringLiteral(node.Prefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", nodeIdent, __gong__toRawStringLiteral(string(node.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.BackgroundColor = %s", nodeIdent, __gong__toRawStringLiteral(node.BackgroundColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", nodeIdent, node.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.HasCheckboxButton = %t", nodeIdent, node.HasCheckboxButton))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", nodeIdent, node.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsCheckboxDisabled = %t", nodeIdent, node.IsCheckboxDisabled))
			values.WriteString(fmt.Sprintf("\n\t%s.CheckboxHasToolTip = %t", nodeIdent, node.CheckboxHasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.CheckboxToolTipText = %s", nodeIdent, __gong__toRawStringLiteral(node.CheckboxToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.CheckboxToolTipPosition = %s", nodeIdent, __gong__toRawStringLiteral(string(node.CheckboxToolTipPosition))))
			values.WriteString(fmt.Sprintf("\n\t%s.HasSecondCheckboxButton = %t", nodeIdent, node.HasSecondCheckboxButton))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSecondCheckboxChecked = %t", nodeIdent, node.IsSecondCheckboxChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSecondCheckboxDisabled = %t", nodeIdent, node.IsSecondCheckboxDisabled))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondCheckboxHasToolTip = %t", nodeIdent, node.SecondCheckboxHasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondCheckboxToolTipText = %s", nodeIdent, __gong__toRawStringLiteral(node.SecondCheckboxToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondCheckboxToolTipPosition = %s", nodeIdent, __gong__toRawStringLiteral(string(node.SecondCheckboxToolTipPosition))))
			values.WriteString(fmt.Sprintf("\n\t%s.TextAfterSecondCheckbox = %s", nodeIdent, __gong__toRawStringLiteral(node.TextAfterSecondCheckbox)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", nodeIdent, node.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", nodeIdent, __gong__toRawStringLiteral(node.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = %s", nodeIdent, __gong__toRawStringLiteral(string(node.ToolTipPosition))))
			values.WriteString(fmt.Sprintf("\n\t%s.ClientOnY = %f", nodeIdent, node.ClientOnY))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInEditMode = %t", nodeIdent, node.IsInEditMode))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNodeClickable = %t", nodeIdent, node.IsNodeClickable))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithPreceedingIcon = %t", nodeIdent, node.IsWithPreceedingIcon))
			values.WriteString(fmt.Sprintf("\n\t%s.PreceedingIcon = %s", nodeIdent, __gong__toRawStringLiteral(node.PreceedingIcon)))
			if node.PreceedingSVGIcon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + node.PreceedingSVGIcon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PreceedingSVGIcon = %s", nodeIdent, targetIdent))
			}
			for _, elem := range node.Children {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Children = append(%s.Children, %s)", nodeIdent, nodeIdent, targetIdent))
			}
			for _, elem := range node.Buttons {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Buttons = append(%s.Buttons, %s)", nodeIdent, nodeIdent, targetIdent))
			}
			if node.Menu != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + node.Menu.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Menu = %s", nodeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, svgicon := range __gong__sortStageSetInstances(stageSet.Stage.SVGIcons, stageSet.Stage.SVGIcon_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			svgiconIdent := "__models" + svgicon.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SVGIcon{Name: %s}).Stage(stageSet.Stage)", svgiconIdent, __gong__toRawStringLiteral(svgicon.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgiconIdent, __gong__toRawStringLiteral(svgicon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", svgiconIdent, __gong__toRawStringLiteral(svgicon.SVG)))
		}
	}
	if stageSet.Stage != nil {
		for _, tree := range __gong__sortStageSetInstances(stageSet.Stage.Trees, stageSet.Stage.Tree_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			treeIdent := "__models" + tree.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tree{Name: %s}).Stage(stageSet.Stage)", treeIdent, __gong__toRawStringLiteral(tree.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", treeIdent, __gong__toRawStringLiteral(tree.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.HaveSearch = %t", treeIdent, tree.HaveSearch))
			for _, elem := range tree.RootNodes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootNodes = append(%s.RootNodes, %s)", treeIdent, treeIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
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
		case "github.com/fullstack-lang/gong/lib/tree/go/models":
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
					identifierMap[ident.Name] = __gong__stageSetInit(new(Button), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Menu":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Menu), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Node":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Node), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SVGIcon":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SVGIcon), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tree":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tree), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
					case "Icon":
						inst.Icon = GongExtractString(rhs)
					case "SVGIcon":
						__gong__assignPointer(&inst.SVGIcon, rhs, identifierMap)
					case "IsDisabled":
						inst.IsDisabled = GongExtractBool(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "ToolTipPosition":
						inst.ToolTipPosition = ToolTipPositionEnum(GongExtractString(rhs))
					case "ClientOnX":
						inst.ClientOnX = GongExtractFloat(rhs)
					case "ClientOnY":
						inst.ClientOnY = GongExtractFloat(rhs)
					}
				case *Menu:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Buttons":
						__gong__assignSliceOfPointers(&inst.Buttons, rhs, identifierMap)
					}
				case *Node:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsWithPrefix":
						inst.IsWithPrefix = GongExtractBool(rhs)
					case "Prefix":
						inst.Prefix = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = FontStyleEnum(GongExtractString(rhs))
					case "BackgroundColor":
						inst.BackgroundColor = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "HasCheckboxButton":
						inst.HasCheckboxButton = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "IsCheckboxDisabled":
						inst.IsCheckboxDisabled = GongExtractBool(rhs)
					case "CheckboxHasToolTip":
						inst.CheckboxHasToolTip = GongExtractBool(rhs)
					case "CheckboxToolTipText":
						inst.CheckboxToolTipText = GongExtractString(rhs)
					case "CheckboxToolTipPosition":
						inst.CheckboxToolTipPosition = ToolTipPositionEnum(GongExtractString(rhs))
					case "HasSecondCheckboxButton":
						inst.HasSecondCheckboxButton = GongExtractBool(rhs)
					case "IsSecondCheckboxChecked":
						inst.IsSecondCheckboxChecked = GongExtractBool(rhs)
					case "IsSecondCheckboxDisabled":
						inst.IsSecondCheckboxDisabled = GongExtractBool(rhs)
					case "SecondCheckboxHasToolTip":
						inst.SecondCheckboxHasToolTip = GongExtractBool(rhs)
					case "SecondCheckboxToolTipText":
						inst.SecondCheckboxToolTipText = GongExtractString(rhs)
					case "SecondCheckboxToolTipPosition":
						inst.SecondCheckboxToolTipPosition = ToolTipPositionEnum(GongExtractString(rhs))
					case "TextAfterSecondCheckbox":
						inst.TextAfterSecondCheckbox = GongExtractString(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "ToolTipPosition":
						inst.ToolTipPosition = ToolTipPositionEnum(GongExtractString(rhs))
					case "ClientOnY":
						inst.ClientOnY = GongExtractFloat(rhs)
					case "IsInEditMode":
						inst.IsInEditMode = GongExtractBool(rhs)
					case "IsNodeClickable":
						inst.IsNodeClickable = GongExtractBool(rhs)
					case "IsWithPreceedingIcon":
						inst.IsWithPreceedingIcon = GongExtractBool(rhs)
					case "PreceedingIcon":
						inst.PreceedingIcon = GongExtractString(rhs)
					case "PreceedingSVGIcon":
						__gong__assignPointer(&inst.PreceedingSVGIcon, rhs, identifierMap)
					case "Children":
						__gong__assignSliceOfPointers(&inst.Children, rhs, identifierMap)
					case "Buttons":
						__gong__assignSliceOfPointers(&inst.Buttons, rhs, identifierMap)
					case "Menu":
						__gong__assignPointer(&inst.Menu, rhs, identifierMap)
					}
				case *SVGIcon:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *Tree:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RootNodes":
						__gong__assignSliceOfPointers(&inst.RootNodes, rhs, identifierMap)
					case "HaveSearch":
						inst.HaveSearch = GongExtractBool(rhs)
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
