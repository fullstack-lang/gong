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
		displayselectionOrdered := []*DisplaySelection{}
		for displayselection := range stageSet.Stage.DisplaySelections {
			displayselectionOrdered = append(displayselectionOrdered, displayselection)
		}
		sort.Slice(displayselectionOrdered, func(i, j int) bool {
			return stageSet.Stage.DisplaySelection_stagedOrder[displayselectionOrdered[i]] < stageSet.Stage.DisplaySelection_stagedOrder[displayselectionOrdered[j]]
		})
		for _, displayselection := range displayselectionOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			displayselectionIdent := "__models" + displayselection.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DisplaySelection{Name: %s}).Stage(stageSet.Stage)", displayselectionIdent, __gong__toRawStringLiteral(displayselection.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", displayselectionIdent, __gong__toRawStringLiteral(displayselection.Name)))
			if displayselection.XLFile != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + displayselection.XLFile.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.XLFile = %s", displayselectionIdent, targetIdent))
			}
			if displayselection.XLSheet != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + displayselection.XLSheet.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.XLSheet = %s", displayselectionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		xlcellOrdered := []*XLCell{}
		for xlcell := range stageSet.Stage.XLCells {
			xlcellOrdered = append(xlcellOrdered, xlcell)
		}
		sort.Slice(xlcellOrdered, func(i, j int) bool {
			return stageSet.Stage.XLCell_stagedOrder[xlcellOrdered[i]] < stageSet.Stage.XLCell_stagedOrder[xlcellOrdered[j]]
		})
		for _, xlcell := range xlcellOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			xlcellIdent := "__models" + xlcell.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.XLCell{Name: %s}).Stage(stageSet.Stage)", xlcellIdent, __gong__toRawStringLiteral(xlcell.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xlcellIdent, __gong__toRawStringLiteral(xlcell.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %d", xlcellIdent, xlcell.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %d", xlcellIdent, xlcell.Y))
		}
	}
	if stageSet.Stage != nil {
		xlfileOrdered := []*XLFile{}
		for xlfile := range stageSet.Stage.XLFiles {
			xlfileOrdered = append(xlfileOrdered, xlfile)
		}
		sort.Slice(xlfileOrdered, func(i, j int) bool {
			return stageSet.Stage.XLFile_stagedOrder[xlfileOrdered[i]] < stageSet.Stage.XLFile_stagedOrder[xlfileOrdered[j]]
		})
		for _, xlfile := range xlfileOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			xlfileIdent := "__models" + xlfile.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.XLFile{Name: %s}).Stage(stageSet.Stage)", xlfileIdent, __gong__toRawStringLiteral(xlfile.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xlfileIdent, __gong__toRawStringLiteral(xlfile.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NbSheets = %d", xlfileIdent, xlfile.NbSheets))
			for _, elem := range xlfile.Sheets {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sheets = append(%s.Sheets, %s)", xlfileIdent, xlfileIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		xlrowOrdered := []*XLRow{}
		for xlrow := range stageSet.Stage.XLRows {
			xlrowOrdered = append(xlrowOrdered, xlrow)
		}
		sort.Slice(xlrowOrdered, func(i, j int) bool {
			return stageSet.Stage.XLRow_stagedOrder[xlrowOrdered[i]] < stageSet.Stage.XLRow_stagedOrder[xlrowOrdered[j]]
		})
		for _, xlrow := range xlrowOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			xlrowIdent := "__models" + xlrow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.XLRow{Name: %s}).Stage(stageSet.Stage)", xlrowIdent, __gong__toRawStringLiteral(xlrow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xlrowIdent, __gong__toRawStringLiteral(xlrow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.RowIndex = %d", xlrowIdent, xlrow.RowIndex))
			for _, elem := range xlrow.Cells {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Cells = append(%s.Cells, %s)", xlrowIdent, xlrowIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		xlsheetOrdered := []*XLSheet{}
		for xlsheet := range stageSet.Stage.XLSheets {
			xlsheetOrdered = append(xlsheetOrdered, xlsheet)
		}
		sort.Slice(xlsheetOrdered, func(i, j int) bool {
			return stageSet.Stage.XLSheet_stagedOrder[xlsheetOrdered[i]] < stageSet.Stage.XLSheet_stagedOrder[xlsheetOrdered[j]]
		})
		for _, xlsheet := range xlsheetOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			xlsheetIdent := "__models" + xlsheet.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.XLSheet{Name: %s}).Stage(stageSet.Stage)", xlsheetIdent, __gong__toRawStringLiteral(xlsheet.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xlsheetIdent, __gong__toRawStringLiteral(xlsheet.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxRow = %d", xlsheetIdent, xlsheet.MaxRow))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxCol = %d", xlsheetIdent, xlsheet.MaxCol))
			values.WriteString(fmt.Sprintf("\n\t%s.NbRows = %d", xlsheetIdent, xlsheet.NbRows))
			for _, elem := range xlsheet.Rows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rows = append(%s.Rows, %s)", xlsheetIdent, xlsheetIdent, targetIdent))
			}
			for _, elem := range xlsheet.SheetCells {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SheetCells = append(%s.SheetCells, %s)", xlsheetIdent, xlsheetIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
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
		case "github.com/fullstack-lang/gong/lib/xlsx/go/models":
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
				case "DisplaySelection":
					if !preserveOrder {
						inst := (&DisplaySelection{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DisplaySelection)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "XLCell":
					if !preserveOrder {
						inst := (&XLCell{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(XLCell)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "XLFile":
					if !preserveOrder {
						inst := (&XLFile{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(XLFile)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "XLRow":
					if !preserveOrder {
						inst := (&XLRow{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(XLRow)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "XLSheet":
					if !preserveOrder {
						inst := (&XLSheet{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(XLSheet)
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
				case *DisplaySelection:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "XLFile":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*XLFile); ok {
									inst.XLFile = typedTarget
								}
							}
						}
					case "XLSheet":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*XLSheet); ok {
									inst.XLSheet = typedTarget
								}
							}
						}
					}
				case *XLCell:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractInt(rhs)
					case "Y":
						inst.Y = GongExtractInt(rhs)
					}
				case *XLFile:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NbSheets":
						inst.NbSheets = GongExtractInt(rhs)
					case "Sheets":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*XLSheet); ok {
										inst.Sheets = append(inst.Sheets, typedTarget)
									}
								}
							}
						}
					}
				case *XLRow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RowIndex":
						inst.RowIndex = GongExtractInt(rhs)
					case "Cells":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*XLCell); ok {
										inst.Cells = append(inst.Cells, typedTarget)
									}
								}
							}
						}
					}
				case *XLSheet:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MaxRow":
						inst.MaxRow = GongExtractInt(rhs)
					case "MaxCol":
						inst.MaxCol = GongExtractInt(rhs)
					case "NbRows":
						inst.NbRows = GongExtractInt(rhs)
					case "Rows":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*XLRow); ok {
										inst.Rows = append(inst.Rows, typedTarget)
									}
								}
							}
						}
					case "SheetCells":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*XLCell); ok {
										inst.SheetCells = append(inst.SheetCells, typedTarget)
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
