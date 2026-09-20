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
		buttonOrdered := []*Button{}
		for button := range stageSet.Stage.Buttons {
			buttonOrdered = append(buttonOrdered, button)
		}
		sort.Slice(buttonOrdered, func(i, j int) bool {
			return stageSet.Stage.Button_stagedOrder[buttonOrdered[i]] < stageSet.Stage.Button_stagedOrder[buttonOrdered[j]]
		})
		for _, button := range buttonOrdered {
			buttonIdent := "__stage_0" + button.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Button{Name: %s}).Stage(stageSet.Stage)", buttonIdent, __gong__toRawStringLiteral(button.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", buttonIdent, __gong__toRawStringLiteral(button.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Icon = %s", buttonIdent, __gong__toRawStringLiteral(button.Icon)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDisabled = %t", buttonIdent, button.IsDisabled))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", buttonIdent, button.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", buttonIdent, __gong__toRawStringLiteral(button.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = %s", buttonIdent, __gong__toRawStringLiteral(string(button.ToolTipPosition))))
			if button.SVGIcon != nil {
				targetIdent := "__stage_0" + button.SVGIcon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SVGIcon = %s", buttonIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		cellOrdered := []*Cell{}
		for cell := range stageSet.Stage.Cells {
			cellOrdered = append(cellOrdered, cell)
		}
		sort.Slice(cellOrdered, func(i, j int) bool {
			return stageSet.Stage.Cell_stagedOrder[cellOrdered[i]] < stageSet.Stage.Cell_stagedOrder[cellOrdered[j]]
		})
		for _, cell := range cellOrdered {
			cellIdent := "__stage_0" + cell.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Cell{Name: %s}).Stage(stageSet.Stage)", cellIdent, __gong__toRawStringLiteral(cell.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cellIdent, __gong__toRawStringLiteral(cell.Name)))
			if cell.CellString != nil {
				targetIdent := "__stage_0" + cell.CellString.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellString = %s", cellIdent, targetIdent))
			}
			if cell.CellFloat64 != nil {
				targetIdent := "__stage_0" + cell.CellFloat64.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellFloat64 = %s", cellIdent, targetIdent))
			}
			if cell.CellInt != nil {
				targetIdent := "__stage_0" + cell.CellInt.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellInt = %s", cellIdent, targetIdent))
			}
			if cell.CellBool != nil {
				targetIdent := "__stage_0" + cell.CellBool.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellBool = %s", cellIdent, targetIdent))
			}
			if cell.CellIcon != nil {
				targetIdent := "__stage_0" + cell.CellIcon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellIcon = %s", cellIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		cellbooleanOrdered := []*CellBoolean{}
		for cellboolean := range stageSet.Stage.CellBooleans {
			cellbooleanOrdered = append(cellbooleanOrdered, cellboolean)
		}
		sort.Slice(cellbooleanOrdered, func(i, j int) bool {
			return stageSet.Stage.CellBoolean_stagedOrder[cellbooleanOrdered[i]] < stageSet.Stage.CellBoolean_stagedOrder[cellbooleanOrdered[j]]
		})
		for _, cellboolean := range cellbooleanOrdered {
			cellbooleanIdent := "__stage_0" + cellboolean.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.CellBoolean{Name: %s}).Stage(stageSet.Stage)", cellbooleanIdent, __gong__toRawStringLiteral(cellboolean.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cellbooleanIdent, __gong__toRawStringLiteral(cellboolean.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %t", cellbooleanIdent, cellboolean.Value))
		}
	}
	if stageSet.Stage != nil {
		cellfloat64Ordered := []*CellFloat64{}
		for cellfloat64 := range stageSet.Stage.CellFloat64s {
			cellfloat64Ordered = append(cellfloat64Ordered, cellfloat64)
		}
		sort.Slice(cellfloat64Ordered, func(i, j int) bool {
			return stageSet.Stage.CellFloat64_stagedOrder[cellfloat64Ordered[i]] < stageSet.Stage.CellFloat64_stagedOrder[cellfloat64Ordered[j]]
		})
		for _, cellfloat64 := range cellfloat64Ordered {
			cellfloat64Ident := "__stage_0" + cellfloat64.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.CellFloat64{Name: %s}).Stage(stageSet.Stage)", cellfloat64Ident, __gong__toRawStringLiteral(cellfloat64.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cellfloat64Ident, __gong__toRawStringLiteral(cellfloat64.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %f", cellfloat64Ident, cellfloat64.Value))
		}
	}
	if stageSet.Stage != nil {
		celliconOrdered := []*CellIcon{}
		for cellicon := range stageSet.Stage.CellIcons {
			celliconOrdered = append(celliconOrdered, cellicon)
		}
		sort.Slice(celliconOrdered, func(i, j int) bool {
			return stageSet.Stage.CellIcon_stagedOrder[celliconOrdered[i]] < stageSet.Stage.CellIcon_stagedOrder[celliconOrdered[j]]
		})
		for _, cellicon := range celliconOrdered {
			celliconIdent := "__stage_0" + cellicon.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.CellIcon{Name: %s}).Stage(stageSet.Stage)", celliconIdent, __gong__toRawStringLiteral(cellicon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", celliconIdent, __gong__toRawStringLiteral(cellicon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Icon = %s", celliconIdent, __gong__toRawStringLiteral(cellicon.Icon)))
			values.WriteString(fmt.Sprintf("\n\t%s.NeedsConfirmation = %t", celliconIdent, cellicon.NeedsConfirmation))
			values.WriteString(fmt.Sprintf("\n\t%s.ConfirmationMessage = %s", celliconIdent, __gong__toRawStringLiteral(cellicon.ConfirmationMessage)))
		}
	}
	if stageSet.Stage != nil {
		cellintOrdered := []*CellInt{}
		for cellint := range stageSet.Stage.CellInts {
			cellintOrdered = append(cellintOrdered, cellint)
		}
		sort.Slice(cellintOrdered, func(i, j int) bool {
			return stageSet.Stage.CellInt_stagedOrder[cellintOrdered[i]] < stageSet.Stage.CellInt_stagedOrder[cellintOrdered[j]]
		})
		for _, cellint := range cellintOrdered {
			cellintIdent := "__stage_0" + cellint.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.CellInt{Name: %s}).Stage(stageSet.Stage)", cellintIdent, __gong__toRawStringLiteral(cellint.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cellintIdent, __gong__toRawStringLiteral(cellint.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", cellintIdent, cellint.Value))
		}
	}
	if stageSet.Stage != nil {
		cellstringOrdered := []*CellString{}
		for cellstring := range stageSet.Stage.CellStrings {
			cellstringOrdered = append(cellstringOrdered, cellstring)
		}
		sort.Slice(cellstringOrdered, func(i, j int) bool {
			return stageSet.Stage.CellString_stagedOrder[cellstringOrdered[i]] < stageSet.Stage.CellString_stagedOrder[cellstringOrdered[j]]
		})
		for _, cellstring := range cellstringOrdered {
			cellstringIdent := "__stage_0" + cellstring.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.CellString{Name: %s}).Stage(stageSet.Stage)", cellstringIdent, __gong__toRawStringLiteral(cellstring.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cellstringIdent, __gong__toRawStringLiteral(cellstring.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", cellstringIdent, __gong__toRawStringLiteral(cellstring.Value)))
		}
	}
	if stageSet.Stage != nil {
		displayedcolumnOrdered := []*DisplayedColumn{}
		for displayedcolumn := range stageSet.Stage.DisplayedColumns {
			displayedcolumnOrdered = append(displayedcolumnOrdered, displayedcolumn)
		}
		sort.Slice(displayedcolumnOrdered, func(i, j int) bool {
			return stageSet.Stage.DisplayedColumn_stagedOrder[displayedcolumnOrdered[i]] < stageSet.Stage.DisplayedColumn_stagedOrder[displayedcolumnOrdered[j]]
		})
		for _, displayedcolumn := range displayedcolumnOrdered {
			displayedcolumnIdent := "__stage_0" + displayedcolumn.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DisplayedColumn{Name: %s}).Stage(stageSet.Stage)", displayedcolumnIdent, __gong__toRawStringLiteral(displayedcolumn.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", displayedcolumnIdent, __gong__toRawStringLiteral(displayedcolumn.Name)))
		}
	}
	if stageSet.Stage != nil {
		rowOrdered := []*Row{}
		for row := range stageSet.Stage.Rows {
			rowOrdered = append(rowOrdered, row)
		}
		sort.Slice(rowOrdered, func(i, j int) bool {
			return stageSet.Stage.Row_stagedOrder[rowOrdered[i]] < stageSet.Stage.Row_stagedOrder[rowOrdered[j]]
		})
		for _, row := range rowOrdered {
			rowIdent := "__stage_0" + row.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Row{Name: %s}).Stage(stageSet.Stage)", rowIdent, __gong__toRawStringLiteral(row.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rowIdent, __gong__toRawStringLiteral(row.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", rowIdent, row.IsChecked))
			for _, elem := range row.Cells {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Cells = append(%s.Cells, %s)", rowIdent, rowIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		svgiconOrdered := []*SVGIcon{}
		for svgicon := range stageSet.Stage.SVGIcons {
			svgiconOrdered = append(svgiconOrdered, svgicon)
		}
		sort.Slice(svgiconOrdered, func(i, j int) bool {
			return stageSet.Stage.SVGIcon_stagedOrder[svgiconOrdered[i]] < stageSet.Stage.SVGIcon_stagedOrder[svgiconOrdered[j]]
		})
		for _, svgicon := range svgiconOrdered {
			svgiconIdent := "__stage_0" + svgicon.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.SVGIcon{Name: %s}).Stage(stageSet.Stage)", svgiconIdent, __gong__toRawStringLiteral(svgicon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgiconIdent, __gong__toRawStringLiteral(svgicon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", svgiconIdent, __gong__toRawStringLiteral(svgicon.SVG)))
		}
	}
	if stageSet.Stage != nil {
		tableOrdered := []*Table{}
		for table := range stageSet.Stage.Tables {
			tableOrdered = append(tableOrdered, table)
		}
		sort.Slice(tableOrdered, func(i, j int) bool {
			return stageSet.Stage.Table_stagedOrder[tableOrdered[i]] < stageSet.Stage.Table_stagedOrder[tableOrdered[j]]
		})
		for _, table := range tableOrdered {
			tableIdent := "__stage_0" + table.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Table{Name: %s}).Stage(stageSet.Stage)", tableIdent, __gong__toRawStringLiteral(table.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tableIdent, __gong__toRawStringLiteral(table.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasFiltering = %t", tableIdent, table.HasFiltering))
			values.WriteString(fmt.Sprintf("\n\t%s.HasColumnSorting = %t", tableIdent, table.HasColumnSorting))
			values.WriteString(fmt.Sprintf("\n\t%s.HasPaginator = %t", tableIdent, table.HasPaginator))
			values.WriteString(fmt.Sprintf("\n\t%s.HasCheckableRows = %t", tableIdent, table.HasCheckableRows))
			values.WriteString(fmt.Sprintf("\n\t%s.HasSaveButton = %t", tableIdent, table.HasSaveButton))
			values.WriteString(fmt.Sprintf("\n\t%s.SaveButtonLabel = %s", tableIdent, __gong__toRawStringLiteral(table.SaveButtonLabel)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasBulkDeleteButton = %t", tableIdent, table.HasBulkDeleteButton))
			values.WriteString(fmt.Sprintf("\n\t%s.BulkDeleteButtonTooltip = %s", tableIdent, __gong__toRawStringLiteral(table.BulkDeleteButtonTooltip)))
			values.WriteString(fmt.Sprintf("\n\t%s.CanDragDropRows = %t", tableIdent, table.CanDragDropRows))
			values.WriteString(fmt.Sprintf("\n\t%s.HasCloseButton = %t", tableIdent, table.HasCloseButton))
			values.WriteString(fmt.Sprintf("\n\t%s.SavingInProgress = %t", tableIdent, table.SavingInProgress))
			values.WriteString(fmt.Sprintf("\n\t%s.NbOfStickyColumns = %d", tableIdent, table.NbOfStickyColumns))
			for _, elem := range table.DisplayedColumns {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DisplayedColumns = append(%s.DisplayedColumns, %s)", tableIdent, tableIdent, targetIdent))
			}
			for _, elem := range table.Rows {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rows = append(%s.Rows, %s)", tableIdent, tableIdent, targetIdent))
			}
			for _, elem := range table.RowsSelectedForBulkDelete {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RowsSelectedForBulkDelete = append(%s.RowsSelectedForBulkDelete, %s)", tableIdent, tableIdent, targetIdent))
			}
			for _, elem := range table.Buttons {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Buttons = append(%s.Buttons, %s)", tableIdent, tableIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/lib/table/go/models"
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
				case "Cell":
					if !preserveOrder {
						inst := (&Cell{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Cell)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "CellBoolean":
					if !preserveOrder {
						inst := (&CellBoolean{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(CellBoolean)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "CellFloat64":
					if !preserveOrder {
						inst := (&CellFloat64{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(CellFloat64)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "CellIcon":
					if !preserveOrder {
						inst := (&CellIcon{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(CellIcon)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "CellInt":
					if !preserveOrder {
						inst := (&CellInt{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(CellInt)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "CellString":
					if !preserveOrder {
						inst := (&CellString{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(CellString)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DisplayedColumn":
					if !preserveOrder {
						inst := (&DisplayedColumn{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DisplayedColumn)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Row":
					if !preserveOrder {
						inst := (&Row{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Row)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SVGIcon":
					if !preserveOrder {
						inst := (&SVGIcon{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SVGIcon)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Table":
					if !preserveOrder {
						inst := (&Table{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Table)
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
					case "Icon":
						inst.Icon = GongExtractString(rhs)
					case "SVGIcon":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SVGIcon); ok {
									inst.SVGIcon = typedTarget
								}
							}
						}
					case "IsDisabled":
						inst.IsDisabled = GongExtractBool(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "ToolTipPosition":
						inst.ToolTipPosition = ToolTipPositionEnum(GongExtractString(rhs))
					}
				case *Cell:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CellString":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*CellString); ok {
									inst.CellString = typedTarget
								}
							}
						}
					case "CellFloat64":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*CellFloat64); ok {
									inst.CellFloat64 = typedTarget
								}
							}
						}
					case "CellInt":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*CellInt); ok {
									inst.CellInt = typedTarget
								}
							}
						}
					case "CellBool":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*CellBoolean); ok {
									inst.CellBool = typedTarget
								}
							}
						}
					case "CellIcon":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*CellIcon); ok {
									inst.CellIcon = typedTarget
								}
							}
						}
					}
				case *CellBoolean:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractBool(rhs)
					}
				case *CellFloat64:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractFloat(rhs)
					}
				case *CellIcon:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Icon":
						inst.Icon = GongExtractString(rhs)
					case "NeedsConfirmation":
						inst.NeedsConfirmation = GongExtractBool(rhs)
					case "ConfirmationMessage":
						inst.ConfirmationMessage = GongExtractString(rhs)
					}
				case *CellInt:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *CellString:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *DisplayedColumn:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Row:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Cells":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Cell); ok {
										inst.Cells = append(inst.Cells, typedTarget)
									}
								}
							}
						}
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					}
				case *SVGIcon:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *Table:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DisplayedColumns":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DisplayedColumn); ok {
										inst.DisplayedColumns = append(inst.DisplayedColumns, typedTarget)
									}
								}
							}
						}
					case "Rows":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Row); ok {
										inst.Rows = append(inst.Rows, typedTarget)
									}
								}
							}
						}
					case "HasFiltering":
						inst.HasFiltering = GongExtractBool(rhs)
					case "HasColumnSorting":
						inst.HasColumnSorting = GongExtractBool(rhs)
					case "HasPaginator":
						inst.HasPaginator = GongExtractBool(rhs)
					case "HasCheckableRows":
						inst.HasCheckableRows = GongExtractBool(rhs)
					case "HasSaveButton":
						inst.HasSaveButton = GongExtractBool(rhs)
					case "SaveButtonLabel":
						inst.SaveButtonLabel = GongExtractString(rhs)
					case "HasBulkDeleteButton":
						inst.HasBulkDeleteButton = GongExtractBool(rhs)
					case "BulkDeleteButtonTooltip":
						inst.BulkDeleteButtonTooltip = GongExtractString(rhs)
					case "RowsSelectedForBulkDelete":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Row); ok {
										inst.RowsSelectedForBulkDelete = append(inst.RowsSelectedForBulkDelete, typedTarget)
									}
								}
							}
						}
					case "CanDragDropRows":
						inst.CanDragDropRows = GongExtractBool(rhs)
					case "HasCloseButton":
						inst.HasCloseButton = GongExtractBool(rhs)
					case "SavingInProgress":
						inst.SavingInProgress = GongExtractBool(rhs)
					case "NbOfStickyColumns":
						inst.NbOfStickyColumns = GongExtractInt(rhs)
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
