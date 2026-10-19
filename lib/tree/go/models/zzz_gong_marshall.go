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
	modelsPackageName := "github.com/fullstack-lang/gong/lib/tree/go/models"
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

	if stage != nil {
		for _, button := range __gong__sortStageSetInstances(stage.Buttons, stage.Button_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			buttonIdent := button.GongGetIdentifier(stage)
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
			values.WriteString(fmt.Sprintf("\n\t%s.ClientOnX = %f", buttonIdent, button.ClientOnX))
			values.WriteString(fmt.Sprintf("\n\t%s.ClientOnY = %f", buttonIdent, button.ClientOnY))
			if button.SVGIcon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := button.SVGIcon.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SVGIcon = %s", buttonIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, menu := range __gong__sortStageSetInstances(stage.Menus, stage.Menu_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			menuIdent := menu.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Menu{Name: %s}).Stage(stage)", menuIdent, __gong__toRawStringLiteral(menu.Name)))
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
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Buttons = append(%s.Buttons, %s)", menuIdent, menuIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, node := range __gong__sortStageSetInstances(stage.Nodes, stage.Node_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			nodeIdent := node.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Node{Name: %s}).Stage(stage)", nodeIdent, __gong__toRawStringLiteral(node.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nodeIdent, __gong__toRawStringLiteral(node.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithPrefix = %t", nodeIdent, node.IsWithPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.Prefix = %s", nodeIdent, __gong__toRawStringLiteral(node.Prefix)))
			if node.FontStyle.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = models.%s", nodeIdent, node.FontStyle.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = \"\"", nodeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.BackgroundColor = %s", nodeIdent, __gong__toRawStringLiteral(node.BackgroundColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", nodeIdent, node.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.HasCheckboxButton = %t", nodeIdent, node.HasCheckboxButton))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", nodeIdent, node.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsCheckboxDisabled = %t", nodeIdent, node.IsCheckboxDisabled))
			values.WriteString(fmt.Sprintf("\n\t%s.CheckboxHasToolTip = %t", nodeIdent, node.CheckboxHasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.CheckboxToolTipText = %s", nodeIdent, __gong__toRawStringLiteral(node.CheckboxToolTipText)))
			if node.CheckboxToolTipPosition.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.CheckboxToolTipPosition = models.%s", nodeIdent, node.CheckboxToolTipPosition.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.CheckboxToolTipPosition = \"\"", nodeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.HasSecondCheckboxButton = %t", nodeIdent, node.HasSecondCheckboxButton))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSecondCheckboxChecked = %t", nodeIdent, node.IsSecondCheckboxChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSecondCheckboxDisabled = %t", nodeIdent, node.IsSecondCheckboxDisabled))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondCheckboxHasToolTip = %t", nodeIdent, node.SecondCheckboxHasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondCheckboxToolTipText = %s", nodeIdent, __gong__toRawStringLiteral(node.SecondCheckboxToolTipText)))
			if node.SecondCheckboxToolTipPosition.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.SecondCheckboxToolTipPosition = models.%s", nodeIdent, node.SecondCheckboxToolTipPosition.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.SecondCheckboxToolTipPosition = \"\"", nodeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.TextAfterSecondCheckbox = %s", nodeIdent, __gong__toRawStringLiteral(node.TextAfterSecondCheckbox)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", nodeIdent, node.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", nodeIdent, __gong__toRawStringLiteral(node.ToolTipText)))
			if node.ToolTipPosition.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = models.%s", nodeIdent, node.ToolTipPosition.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = \"\"", nodeIdent))
			}
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
				targetIdent := node.PreceedingSVGIcon.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PreceedingSVGIcon = %s", nodeIdent, targetIdent))
			}
			for _, elem := range node.Children {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Children = append(%s.Children, %s)", nodeIdent, nodeIdent, targetIdent))
			}
			for _, elem := range node.Buttons {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Buttons = append(%s.Buttons, %s)", nodeIdent, nodeIdent, targetIdent))
			}
			if node.Menu != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := node.Menu.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Menu = %s", nodeIdent, targetIdent))
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
			svgiconIdent := svgicon.GongGetIdentifier(stage)
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
		for _, tree := range __gong__sortStageSetInstances(stage.Trees, stage.Tree_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			treeIdent := tree.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tree{Name: %s}).Stage(stage)", treeIdent, __gong__toRawStringLiteral(tree.Name)))
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
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootNodes = append(%s.RootNodes, %s)", treeIdent, treeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/lib/tree/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/tree/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/tree/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/tree/go/models\"", fmt.Sprintf("%q", modelsPackageName))
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
	case "ClientOnX":
		res = __gong__marshallFloat(ident, "ClientOnX", button.ClientOnX)
	case "ClientOnY":
		res = __gong__marshallFloat(ident, "ClientOnY", button.ClientOnY)

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

func (menu *Menu) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := menu.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", menu.Name)

	case "Buttons":
		var sb strings.Builder
		for _, _button := range menu.Buttons {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Buttons", _button.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Menu", fieldName)
	}
	return
}

func (node *Node) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := node.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", node.Name)
	case "IsWithPrefix":
		res = __gong__marshallBool(ident, "IsWithPrefix", node.IsWithPrefix)
	case "Prefix":
		res = __gong__marshallString(ident, "Prefix", node.Prefix)
	case "FontStyle":
		res = __gong__marshallEnumString(ident, "FontStyle", node.FontStyle.ToCodeString())
	case "BackgroundColor":
		res = __gong__marshallString(ident, "BackgroundColor", node.BackgroundColor)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", node.IsExpanded)
	case "HasCheckboxButton":
		res = __gong__marshallBool(ident, "HasCheckboxButton", node.HasCheckboxButton)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", node.IsChecked)
	case "IsCheckboxDisabled":
		res = __gong__marshallBool(ident, "IsCheckboxDisabled", node.IsCheckboxDisabled)
	case "CheckboxHasToolTip":
		res = __gong__marshallBool(ident, "CheckboxHasToolTip", node.CheckboxHasToolTip)
	case "CheckboxToolTipText":
		res = __gong__marshallString(ident, "CheckboxToolTipText", node.CheckboxToolTipText)
	case "CheckboxToolTipPosition":
		res = __gong__marshallEnumString(ident, "CheckboxToolTipPosition", node.CheckboxToolTipPosition.ToCodeString())
	case "HasSecondCheckboxButton":
		res = __gong__marshallBool(ident, "HasSecondCheckboxButton", node.HasSecondCheckboxButton)
	case "IsSecondCheckboxChecked":
		res = __gong__marshallBool(ident, "IsSecondCheckboxChecked", node.IsSecondCheckboxChecked)
	case "IsSecondCheckboxDisabled":
		res = __gong__marshallBool(ident, "IsSecondCheckboxDisabled", node.IsSecondCheckboxDisabled)
	case "SecondCheckboxHasToolTip":
		res = __gong__marshallBool(ident, "SecondCheckboxHasToolTip", node.SecondCheckboxHasToolTip)
	case "SecondCheckboxToolTipText":
		res = __gong__marshallString(ident, "SecondCheckboxToolTipText", node.SecondCheckboxToolTipText)
	case "SecondCheckboxToolTipPosition":
		res = __gong__marshallEnumString(ident, "SecondCheckboxToolTipPosition", node.SecondCheckboxToolTipPosition.ToCodeString())
	case "TextAfterSecondCheckbox":
		res = __gong__marshallString(ident, "TextAfterSecondCheckbox", node.TextAfterSecondCheckbox)
	case "HasToolTip":
		res = __gong__marshallBool(ident, "HasToolTip", node.HasToolTip)
	case "ToolTipText":
		res = __gong__marshallString(ident, "ToolTipText", node.ToolTipText)
	case "ToolTipPosition":
		res = __gong__marshallEnumString(ident, "ToolTipPosition", node.ToolTipPosition.ToCodeString())
	case "ClientOnY":
		res = __gong__marshallFloat(ident, "ClientOnY", node.ClientOnY)
	case "IsInEditMode":
		res = __gong__marshallBool(ident, "IsInEditMode", node.IsInEditMode)
	case "IsNodeClickable":
		res = __gong__marshallBool(ident, "IsNodeClickable", node.IsNodeClickable)
	case "IsWithPreceedingIcon":
		res = __gong__marshallBool(ident, "IsWithPreceedingIcon", node.IsWithPreceedingIcon)
	case "PreceedingIcon":
		res = __gong__marshallString(ident, "PreceedingIcon", node.PreceedingIcon)

	case "PreceedingSVGIcon":
		if node.PreceedingSVGIcon != nil {
			res = __gong__marshallPointer(ident, "PreceedingSVGIcon", node.PreceedingSVGIcon.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PreceedingSVGIcon", "nil")
		}
	case "Children":
		var sb strings.Builder
		for _, _node := range node.Children {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Children", _node.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Buttons":
		var sb strings.Builder
		for _, _button := range node.Buttons {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Buttons", _button.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Menu":
		if node.Menu != nil {
			res = __gong__marshallPointer(ident, "Menu", node.Menu.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Menu", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Node", fieldName)
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

func (tree *Tree) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tree.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tree.Name)
	case "HaveSearch":
		res = __gong__marshallBool(ident, "HaveSearch", tree.HaveSearch)

	case "RootNodes":
		var sb strings.Builder
		for _, _node := range tree.RootNodes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootNodes", _node.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Tree", fieldName)
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
		initializerStatements.WriteString(button.GongMarshallField(stage, "ClientOnX"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "ClientOnY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (menu *Menu) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(menu.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(menu.GongMarshallField(stage, "Buttons"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (node *Node) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(node.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "IsWithPrefix"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "Prefix"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "FontStyle"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "BackgroundColor"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "HasCheckboxButton"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "IsCheckboxDisabled"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "CheckboxHasToolTip"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "CheckboxToolTipText"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "CheckboxToolTipPosition"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "HasSecondCheckboxButton"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "IsSecondCheckboxChecked"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "IsSecondCheckboxDisabled"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "SecondCheckboxHasToolTip"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "SecondCheckboxToolTipText"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "SecondCheckboxToolTipPosition"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "TextAfterSecondCheckbox"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "HasToolTip"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "ToolTipText"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "ToolTipPosition"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "ClientOnY"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "IsInEditMode"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "IsNodeClickable"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "IsWithPreceedingIcon"))
		initializerStatements.WriteString(node.GongMarshallField(stage, "PreceedingIcon"))
		pointersInitializesStatements.WriteString(node.GongMarshallField(stage, "PreceedingSVGIcon"))
		pointersInitializesStatements.WriteString(node.GongMarshallField(stage, "Children"))
		pointersInitializesStatements.WriteString(node.GongMarshallField(stage, "Buttons"))
		pointersInitializesStatements.WriteString(node.GongMarshallField(stage, "Menu"))
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
func (tree *Tree) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tree.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(tree.GongMarshallField(stage, "RootNodes"))
		initializerStatements.WriteString(tree.GongMarshallField(stage, "HaveSearch"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
