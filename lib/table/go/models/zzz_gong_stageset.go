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
		for _, cell := range __gong__sortStageSetInstances(stageSet.Stage.Cells, stageSet.Stage.Cell_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cellIdent := "__models" + cell.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Cell{Name: %s}).Stage(stageSet.Stage)", cellIdent, __gong__toRawStringLiteral(cell.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cellIdent, __gong__toRawStringLiteral(cell.Name)))
			if cell.CellString != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + cell.CellString.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellString = %s", cellIdent, targetIdent))
			}
			if cell.CellFloat64 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + cell.CellFloat64.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellFloat64 = %s", cellIdent, targetIdent))
			}
			if cell.CellInt != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + cell.CellInt.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellInt = %s", cellIdent, targetIdent))
			}
			if cell.CellBool != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + cell.CellBool.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellBool = %s", cellIdent, targetIdent))
			}
			if cell.CellIcon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + cell.CellIcon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellIcon = %s", cellIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, cellboolean := range __gong__sortStageSetInstances(stageSet.Stage.CellBooleans, stageSet.Stage.CellBoolean_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cellbooleanIdent := "__models" + cellboolean.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CellBoolean{Name: %s}).Stage(stageSet.Stage)", cellbooleanIdent, __gong__toRawStringLiteral(cellboolean.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cellbooleanIdent, __gong__toRawStringLiteral(cellboolean.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %t", cellbooleanIdent, cellboolean.Value))
		}
	}
	if stageSet.Stage != nil {
		for _, cellfloat64 := range __gong__sortStageSetInstances(stageSet.Stage.CellFloat64s, stageSet.Stage.CellFloat64_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cellfloat64Ident := "__models" + cellfloat64.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CellFloat64{Name: %s}).Stage(stageSet.Stage)", cellfloat64Ident, __gong__toRawStringLiteral(cellfloat64.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cellfloat64Ident, __gong__toRawStringLiteral(cellfloat64.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %f", cellfloat64Ident, cellfloat64.Value))
		}
	}
	if stageSet.Stage != nil {
		for _, cellicon := range __gong__sortStageSetInstances(stageSet.Stage.CellIcons, stageSet.Stage.CellIcon_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			celliconIdent := "__models" + cellicon.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CellIcon{Name: %s}).Stage(stageSet.Stage)", celliconIdent, __gong__toRawStringLiteral(cellicon.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", celliconIdent, __gong__toRawStringLiteral(cellicon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Icon = %s", celliconIdent, __gong__toRawStringLiteral(cellicon.Icon)))
			values.WriteString(fmt.Sprintf("\n\t%s.NeedsConfirmation = %t", celliconIdent, cellicon.NeedsConfirmation))
			values.WriteString(fmt.Sprintf("\n\t%s.ConfirmationMessage = %s", celliconIdent, __gong__toRawStringLiteral(cellicon.ConfirmationMessage)))
		}
	}
	if stageSet.Stage != nil {
		for _, cellint := range __gong__sortStageSetInstances(stageSet.Stage.CellInts, stageSet.Stage.CellInt_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cellintIdent := "__models" + cellint.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CellInt{Name: %s}).Stage(stageSet.Stage)", cellintIdent, __gong__toRawStringLiteral(cellint.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cellintIdent, __gong__toRawStringLiteral(cellint.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", cellintIdent, cellint.Value))
		}
	}
	if stageSet.Stage != nil {
		for _, cellstring := range __gong__sortStageSetInstances(stageSet.Stage.CellStrings, stageSet.Stage.CellString_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cellstringIdent := "__models" + cellstring.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CellString{Name: %s}).Stage(stageSet.Stage)", cellstringIdent, __gong__toRawStringLiteral(cellstring.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cellstringIdent, __gong__toRawStringLiteral(cellstring.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", cellstringIdent, __gong__toRawStringLiteral(cellstring.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, displayedcolumn := range __gong__sortStageSetInstances(stageSet.Stage.DisplayedColumns, stageSet.Stage.DisplayedColumn_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			displayedcolumnIdent := "__models" + displayedcolumn.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DisplayedColumn{Name: %s}).Stage(stageSet.Stage)", displayedcolumnIdent, __gong__toRawStringLiteral(displayedcolumn.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", displayedcolumnIdent, __gong__toRawStringLiteral(displayedcolumn.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, row := range __gong__sortStageSetInstances(stageSet.Stage.Rows, stageSet.Stage.Row_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rowIdent := "__models" + row.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Row{Name: %s}).Stage(stageSet.Stage)", rowIdent, __gong__toRawStringLiteral(row.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rowIdent, __gong__toRawStringLiteral(row.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", rowIdent, row.IsChecked))
			for _, elem := range row.Cells {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Cells = append(%s.Cells, %s)", rowIdent, rowIdent, targetIdent))
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
		for _, table := range __gong__sortStageSetInstances(stageSet.Stage.Tables, stageSet.Stage.Table_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tableIdent := "__models" + table.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Table{Name: %s}).Stage(stageSet.Stage)", tableIdent, __gong__toRawStringLiteral(table.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DisplayedColumns = append(%s.DisplayedColumns, %s)", tableIdent, tableIdent, targetIdent))
			}
			for _, elem := range table.Rows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rows = append(%s.Rows, %s)", tableIdent, tableIdent, targetIdent))
			}
			for _, elem := range table.RowsSelectedForBulkDelete {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RowsSelectedForBulkDelete = append(%s.RowsSelectedForBulkDelete, %s)", tableIdent, tableIdent, targetIdent))
			}
			for _, elem := range table.Buttons {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Buttons = append(%s.Buttons, %s)", tableIdent, tableIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/table/go/models"
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
		case "github.com/fullstack-lang/gong/lib/table/go/models":
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
				case "Cell":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Cell), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "CellBoolean":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CellBoolean), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "CellFloat64":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CellFloat64), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "CellIcon":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CellIcon), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "CellInt":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CellInt), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "CellString":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CellString), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DisplayedColumn":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DisplayedColumn), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Row":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Row), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SVGIcon":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SVGIcon), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Table":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Table), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
					}
				case *Cell:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CellString":
						__gong__assignPointer(&inst.CellString, rhs, identifierMap)
					case "CellFloat64":
						__gong__assignPointer(&inst.CellFloat64, rhs, identifierMap)
					case "CellInt":
						__gong__assignPointer(&inst.CellInt, rhs, identifierMap)
					case "CellBool":
						__gong__assignPointer(&inst.CellBool, rhs, identifierMap)
					case "CellIcon":
						__gong__assignPointer(&inst.CellIcon, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.Cells, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.DisplayedColumns, rhs, identifierMap)
					case "Rows":
						__gong__assignSliceOfPointers(&inst.Rows, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.RowsSelectedForBulkDelete, rhs, identifierMap)
					case "CanDragDropRows":
						inst.CanDragDropRows = GongExtractBool(rhs)
					case "HasCloseButton":
						inst.HasCloseButton = GongExtractBool(rhs)
					case "SavingInProgress":
						inst.SavingInProgress = GongExtractBool(rhs)
					case "NbOfStickyColumns":
						inst.NbOfStickyColumns = GongExtractInt(rhs)
					case "Buttons":
						__gong__assignSliceOfPointers(&inst.Buttons, rhs, identifierMap)
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
