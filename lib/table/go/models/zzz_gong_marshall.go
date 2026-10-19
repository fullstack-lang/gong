// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sort"
	"strings"
	"time"
)

var (
	_ = sort.Slice
	_ = time.Hour
)

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`


const GongStringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongMetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

const GongStringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongNumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongPointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongSliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const GongTimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// __gong__toRawStringLiteral formats a string into safe Go source code,
// using backticks to preserve newlines and readability.
func __gong__toRawStringLiteral(s string) string {
	// Step 1: Replace every backtick with a closing backtick,
	// a double-quoted backtick, and an opening backtick.
	escaped := strings.ReplaceAll(s, "`", "` + \"`\" + `")

	// Step 2: Wrap the entire resulting string in backticks.
	result := "`" + escaped + "`"

	// Step 3: Clean up any empty raw strings (``) at the boundaries
	// just in case your original string started or ended with a backtick.
	result = strings.ReplaceAll(result, "`` + ", "")
	result = strings.ReplaceAll(result, " + ``", "")

	return result
}

func __gong__marshallString[T ~string](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongStringInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(string(val)))
}

func __gong__marshallInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", val))
}

func __gong__marshallBool[T ~bool](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", val))
}

func __gong__marshallFloat[T ~float32 | ~float64](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", val))
}

func __gong__marshallTime(ident, fieldName, valStr string) string {
	res := strings.ReplaceAll(GongTimeInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", valStr)
}

func __gong__marshallPointer(ident, fieldName, targetIdent string) string {
	res := strings.ReplaceAll(GongPointerFieldInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", targetIdent)
}

func __gong__marshallSliceOfPointers(ident, fieldName, targetIdent string) string {
	res := strings.ReplaceAll(GongSliceOfPointersFieldInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", targetIdent)
}

func __gong__marshallEnumString(ident, fieldName, codeStr string) string {
	val := "\"\""
	if codeStr != "" {
		val = "models." + codeStr
	}
	res := strings.ReplaceAll(GongStringEnumInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

func __gong__marshallEnumInt(ident, fieldName, codeStr string) string {
	val := "0"
	if codeStr != "" {
		val = "models." + codeStr
	}
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

func __gong__marshallMeta(ident, fieldName, val string) string {
	res := strings.ReplaceAll(GongMetaFieldStructInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
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

// MarshallFile marshall the stage content into a file as an instanciation into a stage
// according to the marshalling policy of the stage.
//
// In GongMarshallingAppendCommit mode, it will append the last commit to the file.
// In other modes, it will rewrite the entire file.
func (stage *Stage) MarshallFile(filename string, args ...string) {

	if stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		contentBytes, err := os.ReadFile(filename)

		// if the file does not exist, marshall the full stage
		if os.IsNotExist(err) {
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, args...)
			return
		}
		if err != nil {
			log.Fatal(err.Error())
		}

		content := string(contentBytes)

		if stage.IsSquashing() {
			// we squash: we want to clear the current function body
			// and let the append logic write the squashed commit
			const fnSig = "func _(stage *models.Stage) {"
			if before, _, ok := strings.Cut(content, fnSig); ok {
				content = before + fnSig + "\n}\n"
			}
		}

		if stage.IsApplyingBackwardCommit() {
			// we are going backward, we need to remove the last forward commit from the file

			// because commitsBehind has been incremented before the call to this function
			// the index of the commit to remove is len(forwardCommits) - commitsBehind
			forwardCommits := stage.GetForwardCommits()
			commitIndexToRemove := len(forwardCommits) - stage.GetCommitsBehind()

			if commitIndexToRemove < 0 || commitIndexToRemove >= len(forwardCommits) {
				return // Should not happen if history is consistent
			}

			commitToRemove := forwardCommits[commitIndexToRemove]

			lastIndex := strings.LastIndex(content, commitToRemove+"\n")
			if lastIndex != -1 {
				newContent := content[:lastIndex] + content[lastIndex+len(commitToRemove)+1:]
				err = os.WriteFile(filename, []byte(newContent), 0644)
				if err != nil {
					log.Fatal(err.Error())
				}
			} else {
				lastIndex = strings.LastIndex(content, commitToRemove)
				if lastIndex != -1 {
					newContent := content[:lastIndex] + content[lastIndex+len(commitToRemove):]
					err = os.WriteFile(filename, []byte(newContent), 0644)
					if err != nil {
						log.Fatal(err.Error())
					}
				} else {
					// The commit block was not found. This typically happens for the initial
					// commit which is formatted differently (the lines after func _(stage *models.Stage) {).
					// We rewrite the entire file with the current (rewound) stage state to safely remove it.
					file, createErr := os.Create(filename)
					if createErr != nil {
						log.Fatal(createErr.Error())
					}
					defer file.Close()
					stage.Marshall(file, args...)
				}
			}
			return // we are done for the backward case
		}

		if stage.IsApplyingForwardCommit() {
			// bypass the modified check
		} else if !stage.IsModified() {
			return
		}

		forwardCommits := stage.GetForwardCommits()
		if len(forwardCommits) == 0 {
			return // nothing to do
		}

		activeCommits := len(forwardCommits) - stage.GetCommitsBehind()
		if activeCommits <= 0 {
			return
		}
		forwardCommit := forwardCommits[activeCommits-1]

		// append before the ending brace of the func
		lastBrace := strings.LastIndex(content, "}")
		if lastBrace == -1 {
			// if no brace, something is wrong with the file, so we rewrite it
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, args...)
			return
		}

		contentBeforeBrace := content[:lastBrace]
		trimmedContentBeforeBrace := strings.TrimSpace(contentBeforeBrace)
		emptyBody := stage.IsSquashing() ||
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *models.Stage) {") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// insertion point for setup of pointers") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)\n\t// ------------------------------------------------------------------------")

		// check if the file ends with stage.Commit() before the brace
		if !emptyBody && !strings.HasSuffix(trimmedContentBeforeBrace, "stage.Commit()") {
			contentBeforeBrace = contentBeforeBrace + "\n\tstage.Commit()\n"
		}

		// insert the commit statements before the last brace
		newContent := contentBeforeBrace + forwardCommit + "\n" + content[lastBrace:]

		err = os.WriteFile(filename, []byte(newContent), 0644)
		if err != nil {
			log.Fatal(err.Error())
		}

	} else {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		stage.Marshall(file, args...)
	}
}

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, args ...string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(args...)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(args ...string) (res string, err error) {
	packageName := "main"
	modelsPackageName := "github.com/fullstack-lang/gong/lib/table/go/models"
	if len(args) == 1 {
		packageName = args[0]
	} else if len(args) >= 2 {
		if args[0] != "" {
			modelsPackageName = args[0]
		}
		packageName = args[1]
	}

	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr
	localAliasPrefix := "__models"
	if stage.IsInDeltaMode() || stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		localAliasPrefix = ""
	}
	_ = localAliasPrefix

	if stage != nil {
		for _, button := range __gong__sortStageSetInstances(stage.Buttons, stage.Button_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			buttonIdent := localAliasPrefix + button.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Button{Name: %s}).Stage(stage)", buttonIdent, __gong__toRawStringLiteral(button.Name)))
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
			if button.ToolTipPosition.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = models.%s", buttonIdent, button.ToolTipPosition.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = \"\"", buttonIdent))
			}
			if button.SVGIcon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + button.SVGIcon.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SVGIcon = %s", buttonIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, cell := range __gong__sortStageSetInstances(stage.Cells, stage.Cell_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cellIdent := localAliasPrefix + cell.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Cell{Name: %s}).Stage(stage)", cellIdent, __gong__toRawStringLiteral(cell.Name)))
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
				targetIdent := localAliasPrefix + cell.CellString.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellString = %s", cellIdent, targetIdent))
			}
			if cell.CellFloat64 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + cell.CellFloat64.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellFloat64 = %s", cellIdent, targetIdent))
			}
			if cell.CellInt != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + cell.CellInt.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellInt = %s", cellIdent, targetIdent))
			}
			if cell.CellBool != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + cell.CellBool.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellBool = %s", cellIdent, targetIdent))
			}
			if cell.CellIcon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + cell.CellIcon.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellIcon = %s", cellIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, cellboolean := range __gong__sortStageSetInstances(stage.CellBooleans, stage.CellBoolean_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cellbooleanIdent := localAliasPrefix + cellboolean.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CellBoolean{Name: %s}).Stage(stage)", cellbooleanIdent, __gong__toRawStringLiteral(cellboolean.Name)))
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
	if stage != nil {
		for _, cellfloat64 := range __gong__sortStageSetInstances(stage.CellFloat64s, stage.CellFloat64_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cellfloat64Ident := localAliasPrefix + cellfloat64.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CellFloat64{Name: %s}).Stage(stage)", cellfloat64Ident, __gong__toRawStringLiteral(cellfloat64.Name)))
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
	if stage != nil {
		for _, cellicon := range __gong__sortStageSetInstances(stage.CellIcons, stage.CellIcon_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			celliconIdent := localAliasPrefix + cellicon.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CellIcon{Name: %s}).Stage(stage)", celliconIdent, __gong__toRawStringLiteral(cellicon.Name)))
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
	if stage != nil {
		for _, cellint := range __gong__sortStageSetInstances(stage.CellInts, stage.CellInt_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cellintIdent := localAliasPrefix + cellint.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CellInt{Name: %s}).Stage(stage)", cellintIdent, __gong__toRawStringLiteral(cellint.Name)))
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
	if stage != nil {
		for _, cellstring := range __gong__sortStageSetInstances(stage.CellStrings, stage.CellString_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cellstringIdent := localAliasPrefix + cellstring.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CellString{Name: %s}).Stage(stage)", cellstringIdent, __gong__toRawStringLiteral(cellstring.Name)))
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
	if stage != nil {
		for _, displayedcolumn := range __gong__sortStageSetInstances(stage.DisplayedColumns, stage.DisplayedColumn_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			displayedcolumnIdent := localAliasPrefix + displayedcolumn.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DisplayedColumn{Name: %s}).Stage(stage)", displayedcolumnIdent, __gong__toRawStringLiteral(displayedcolumn.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", displayedcolumnIdent, __gong__toRawStringLiteral(displayedcolumn.Name)))
		}
	}
	if stage != nil {
		for _, row := range __gong__sortStageSetInstances(stage.Rows, stage.Row_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rowIdent := localAliasPrefix + row.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Row{Name: %s}).Stage(stage)", rowIdent, __gong__toRawStringLiteral(row.Name)))
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
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Cells = append(%s.Cells, %s)", rowIdent, rowIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, svgicon := range __gong__sortStageSetInstances(stage.SVGIcons, stage.SVGIcon_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			svgiconIdent := localAliasPrefix + svgicon.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SVGIcon{Name: %s}).Stage(stage)", svgiconIdent, __gong__toRawStringLiteral(svgicon.Name)))
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
	if stage != nil {
		for _, table := range __gong__sortStageSetInstances(stage.Tables, stage.Table_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tableIdent := localAliasPrefix + table.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Table{Name: %s}).Stage(stage)", tableIdent, __gong__toRawStringLiteral(table.Name)))
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
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DisplayedColumns = append(%s.DisplayedColumns, %s)", tableIdent, tableIdent, targetIdent))
			}
			for _, elem := range table.Rows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rows = append(%s.Rows, %s)", tableIdent, tableIdent, targetIdent))
			}
			for _, elem := range table.RowsSelectedForBulkDelete {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RowsSelectedForBulkDelete = append(%s.RowsSelectedForBulkDelete, %s)", tableIdent, tableIdent, targetIdent))
			}
			for _, elem := range table.Buttons {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Buttons = append(%s.Buttons, %s)", tableIdent, tableIdent, targetIdent))
			}
		}
	}

	var goModuleVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		goModuleVersion = buildInfo.Main.Version
	}

	goModuleVersionWithoutDirty := strings.ReplaceAll(goModuleVersion, "+dirty", "")
	if goModuleVersionWithoutDirty == "" || goModuleVersionWithoutDirty == "(devel)" {
		goModuleVersionWithoutDirty = "latest"
	}

	var metaImports strings.Builder
	var metaDummies strings.Builder
	if len(stage.MetaPackageImports) > 0 {
		for _, imp := range stage.MetaPackageImports {
			metaImports.WriteString(fmt.Sprintf("\n\t%s %s", imp.Alias, imp.Path))
			metaDummies.WriteString(fmt.Sprintf("\nvar _ %s.Stage", imp.Alias))
		}
	} else if stage.MetaPackageImportAlias != "" {
		metaImports.WriteString(fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))
		metaDummies.WriteString(fmt.Sprintf("\nvar _ %s.Stage", stage.MetaPackageImportAlias))
	}

	res = fmt.Sprintf(`// file generated by gong
// go module version: %s
/* to edit this file with the default probe, use the command
go run github.com/fullstack-lang/gong/lib/table/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/table/go/models"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *models.Stage) {

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
`, goModuleVersion, goModuleVersionWithoutDirty, packageName, metaImports.String(), metaDummies.String(), declarations.String(), values.String(), pointers.String())

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/table/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/table/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (button *Button) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := button.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", button.Name)
	case "Icon":
		res = __gong__marshallString(ident, "Icon", button.Icon)
	case "IsDisabled":
		res = __gong__marshallBool(ident, "IsDisabled", button.IsDisabled)
	case "HasToolTip":
		res = __gong__marshallBool(ident, "HasToolTip", button.HasToolTip)
	case "ToolTipText":
		res = __gong__marshallString(ident, "ToolTipText", button.ToolTipText)
	case "ToolTipPosition":
		res = __gong__marshallEnumString(ident, "ToolTipPosition", button.ToolTipPosition.ToCodeString())

	case "SVGIcon":
		if button.SVGIcon != nil {
			res = __gong__marshallPointer(ident, "SVGIcon", button.SVGIcon.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SVGIcon", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Button", fieldName)
	}
	return
}

func (cell *Cell) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := cell.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", cell.Name)

	case "CellString":
		if cell.CellString != nil {
			res = __gong__marshallPointer(ident, "CellString", cell.CellString.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CellString", "nil")
		}
	case "CellFloat64":
		if cell.CellFloat64 != nil {
			res = __gong__marshallPointer(ident, "CellFloat64", cell.CellFloat64.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CellFloat64", "nil")
		}
	case "CellInt":
		if cell.CellInt != nil {
			res = __gong__marshallPointer(ident, "CellInt", cell.CellInt.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CellInt", "nil")
		}
	case "CellBool":
		if cell.CellBool != nil {
			res = __gong__marshallPointer(ident, "CellBool", cell.CellBool.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CellBool", "nil")
		}
	case "CellIcon":
		if cell.CellIcon != nil {
			res = __gong__marshallPointer(ident, "CellIcon", cell.CellIcon.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CellIcon", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Cell", fieldName)
	}
	return
}

func (cellboolean *CellBoolean) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := cellboolean.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", cellboolean.Name)
	case "Value":
		res = __gong__marshallBool(ident, "Value", cellboolean.Value)

	default:
		log.Panicf("Unknown field %s for Gongstruct CellBoolean", fieldName)
	}
	return
}

func (cellfloat64 *CellFloat64) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := cellfloat64.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", cellfloat64.Name)
	case "Value":
		res = __gong__marshallFloat(ident, "Value", cellfloat64.Value)

	default:
		log.Panicf("Unknown field %s for Gongstruct CellFloat64", fieldName)
	}
	return
}

func (cellicon *CellIcon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := cellicon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", cellicon.Name)
	case "Icon":
		res = __gong__marshallString(ident, "Icon", cellicon.Icon)
	case "NeedsConfirmation":
		res = __gong__marshallBool(ident, "NeedsConfirmation", cellicon.NeedsConfirmation)
	case "ConfirmationMessage":
		res = __gong__marshallString(ident, "ConfirmationMessage", cellicon.ConfirmationMessage)

	default:
		log.Panicf("Unknown field %s for Gongstruct CellIcon", fieldName)
	}
	return
}

func (cellint *CellInt) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := cellint.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", cellint.Name)
	case "Value":
		res = __gong__marshallInt(ident, "Value", cellint.Value)

	default:
		log.Panicf("Unknown field %s for Gongstruct CellInt", fieldName)
	}
	return
}

func (cellstring *CellString) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := cellstring.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", cellstring.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", cellstring.Value)

	default:
		log.Panicf("Unknown field %s for Gongstruct CellString", fieldName)
	}
	return
}

func (displayedcolumn *DisplayedColumn) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := displayedcolumn.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", displayedcolumn.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct DisplayedColumn", fieldName)
	}
	return
}

func (row *Row) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := row.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", row.Name)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", row.IsChecked)

	case "Cells":
		var sb strings.Builder
		for _, _cell := range row.Cells {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Cells", _cell.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Row", fieldName)
	}
	return
}

func (svgicon *SVGIcon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := svgicon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", svgicon.Name)
	case "SVG":
		res = __gong__marshallString(ident, "SVG", svgicon.SVG)

	default:
		log.Panicf("Unknown field %s for Gongstruct SVGIcon", fieldName)
	}
	return
}

func (table *Table) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := table.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", table.Name)
	case "HasFiltering":
		res = __gong__marshallBool(ident, "HasFiltering", table.HasFiltering)
	case "HasColumnSorting":
		res = __gong__marshallBool(ident, "HasColumnSorting", table.HasColumnSorting)
	case "HasPaginator":
		res = __gong__marshallBool(ident, "HasPaginator", table.HasPaginator)
	case "HasCheckableRows":
		res = __gong__marshallBool(ident, "HasCheckableRows", table.HasCheckableRows)
	case "HasSaveButton":
		res = __gong__marshallBool(ident, "HasSaveButton", table.HasSaveButton)
	case "SaveButtonLabel":
		res = __gong__marshallString(ident, "SaveButtonLabel", table.SaveButtonLabel)
	case "HasBulkDeleteButton":
		res = __gong__marshallBool(ident, "HasBulkDeleteButton", table.HasBulkDeleteButton)
	case "BulkDeleteButtonTooltip":
		res = __gong__marshallString(ident, "BulkDeleteButtonTooltip", table.BulkDeleteButtonTooltip)
	case "CanDragDropRows":
		res = __gong__marshallBool(ident, "CanDragDropRows", table.CanDragDropRows)
	case "HasCloseButton":
		res = __gong__marshallBool(ident, "HasCloseButton", table.HasCloseButton)
	case "SavingInProgress":
		res = __gong__marshallBool(ident, "SavingInProgress", table.SavingInProgress)
	case "NbOfStickyColumns":
		res = __gong__marshallInt(ident, "NbOfStickyColumns", table.NbOfStickyColumns)

	case "DisplayedColumns":
		var sb strings.Builder
		for _, _displayedcolumn := range table.DisplayedColumns {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DisplayedColumns", _displayedcolumn.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Rows":
		var sb strings.Builder
		for _, _row := range table.Rows {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Rows", _row.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RowsSelectedForBulkDelete":
		var sb strings.Builder
		for _, _row := range table.RowsSelectedForBulkDelete {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RowsSelectedForBulkDelete", _row.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Buttons":
		var sb strings.Builder
		for _, _button := range table.Buttons {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Buttons", _button.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Table", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (button *Button) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(button.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "Icon"))
		pointersInitializesStatements.WriteString(button.GongMarshallField(stage, "SVGIcon"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "IsDisabled"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "HasToolTip"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "ToolTipText"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "ToolTipPosition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cell *Cell) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cell.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellString"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellFloat64"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellInt"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellBool"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellIcon"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cellboolean *CellBoolean) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellboolean.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellboolean.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cellfloat64 *CellFloat64) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellfloat64.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellfloat64.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cellicon *CellIcon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "Icon"))
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "NeedsConfirmation"))
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "ConfirmationMessage"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cellint *CellInt) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellint.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellint.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cellstring *CellString) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellstring.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellstring.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (displayedcolumn *DisplayedColumn) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(displayedcolumn.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (row *Row) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(row.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(row.GongMarshallField(stage, "Cells"))
		initializerStatements.WriteString(row.GongMarshallField(stage, "IsChecked"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (svgicon *SVGIcon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(svgicon.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(svgicon.GongMarshallField(stage, "SVG"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (table *Table) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(table.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "DisplayedColumns"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "Rows"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasFiltering"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasColumnSorting"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasPaginator"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasCheckableRows"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasSaveButton"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "SaveButtonLabel"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasBulkDeleteButton"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "BulkDeleteButtonTooltip"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "RowsSelectedForBulkDelete"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "CanDragDropRows"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasCloseButton"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "SavingInProgress"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "NbOfStickyColumns"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "Buttons"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
