// generated code - do not edit
package aixmfeatures

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
	{{Identifier}} := (&aixmfeatures.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *aixmfeatures.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *aixmfeatures.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmfeatures"
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
		for _, a_extension := range __gong__sortStageSetInstances(stage.A_extensions, stage.A_extension_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extensionIdent := a_extension.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension{Name: %s}).Stage(stage)", a_extensionIdent, __gong__toRawStringLiteral(a_extension.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extensionIdent, __gong__toRawStringLiteral(a_extension.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extensionIdent, a_extension.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_1 := range __gong__sortStageSetInstances(stage.A_extension_1s, stage.A_extension_1_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_1Ident := a_extension_1.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_1{Name: %s}).Stage(stage)", a_extension_1Ident, __gong__toRawStringLiteral(a_extension_1.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_1Ident, __gong__toRawStringLiteral(a_extension_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_1Ident, a_extension_1.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_10 := range __gong__sortStageSetInstances(stage.A_extension_10s, stage.A_extension_10_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_10Ident := a_extension_10.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_10{Name: %s}).Stage(stage)", a_extension_10Ident, __gong__toRawStringLiteral(a_extension_10.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_10Ident, __gong__toRawStringLiteral(a_extension_10.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_10Ident, a_extension_10.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_11 := range __gong__sortStageSetInstances(stage.A_extension_11s, stage.A_extension_11_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_11Ident := a_extension_11.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_11{Name: %s}).Stage(stage)", a_extension_11Ident, __gong__toRawStringLiteral(a_extension_11.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_11Ident, __gong__toRawStringLiteral(a_extension_11.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_11Ident, a_extension_11.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_12 := range __gong__sortStageSetInstances(stage.A_extension_12s, stage.A_extension_12_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_12Ident := a_extension_12.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_12{Name: %s}).Stage(stage)", a_extension_12Ident, __gong__toRawStringLiteral(a_extension_12.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_12Ident, __gong__toRawStringLiteral(a_extension_12.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_12Ident, a_extension_12.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_13 := range __gong__sortStageSetInstances(stage.A_extension_13s, stage.A_extension_13_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_13Ident := a_extension_13.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_13{Name: %s}).Stage(stage)", a_extension_13Ident, __gong__toRawStringLiteral(a_extension_13.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_13Ident, __gong__toRawStringLiteral(a_extension_13.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_13Ident, a_extension_13.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_14 := range __gong__sortStageSetInstances(stage.A_extension_14s, stage.A_extension_14_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_14Ident := a_extension_14.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_14{Name: %s}).Stage(stage)", a_extension_14Ident, __gong__toRawStringLiteral(a_extension_14.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_14Ident, __gong__toRawStringLiteral(a_extension_14.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_14Ident, a_extension_14.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_15 := range __gong__sortStageSetInstances(stage.A_extension_15s, stage.A_extension_15_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_15Ident := a_extension_15.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_15{Name: %s}).Stage(stage)", a_extension_15Ident, __gong__toRawStringLiteral(a_extension_15.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_15Ident, __gong__toRawStringLiteral(a_extension_15.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_15Ident, a_extension_15.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_16 := range __gong__sortStageSetInstances(stage.A_extension_16s, stage.A_extension_16_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_16Ident := a_extension_16.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_16{Name: %s}).Stage(stage)", a_extension_16Ident, __gong__toRawStringLiteral(a_extension_16.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_16Ident, __gong__toRawStringLiteral(a_extension_16.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_16Ident, a_extension_16.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_17 := range __gong__sortStageSetInstances(stage.A_extension_17s, stage.A_extension_17_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_17Ident := a_extension_17.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_17{Name: %s}).Stage(stage)", a_extension_17Ident, __gong__toRawStringLiteral(a_extension_17.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_17Ident, __gong__toRawStringLiteral(a_extension_17.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_17Ident, a_extension_17.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_18 := range __gong__sortStageSetInstances(stage.A_extension_18s, stage.A_extension_18_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_18Ident := a_extension_18.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_18{Name: %s}).Stage(stage)", a_extension_18Ident, __gong__toRawStringLiteral(a_extension_18.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_18Ident, __gong__toRawStringLiteral(a_extension_18.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_18Ident, a_extension_18.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_19 := range __gong__sortStageSetInstances(stage.A_extension_19s, stage.A_extension_19_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_19Ident := a_extension_19.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_19{Name: %s}).Stage(stage)", a_extension_19Ident, __gong__toRawStringLiteral(a_extension_19.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_19Ident, __gong__toRawStringLiteral(a_extension_19.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_19Ident, a_extension_19.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_2 := range __gong__sortStageSetInstances(stage.A_extension_2s, stage.A_extension_2_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_2Ident := a_extension_2.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_2{Name: %s}).Stage(stage)", a_extension_2Ident, __gong__toRawStringLiteral(a_extension_2.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_2Ident, __gong__toRawStringLiteral(a_extension_2.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_2Ident, a_extension_2.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_20 := range __gong__sortStageSetInstances(stage.A_extension_20s, stage.A_extension_20_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_20Ident := a_extension_20.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_20{Name: %s}).Stage(stage)", a_extension_20Ident, __gong__toRawStringLiteral(a_extension_20.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_20Ident, __gong__toRawStringLiteral(a_extension_20.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_20Ident, a_extension_20.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_21 := range __gong__sortStageSetInstances(stage.A_extension_21s, stage.A_extension_21_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_21Ident := a_extension_21.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_21{Name: %s}).Stage(stage)", a_extension_21Ident, __gong__toRawStringLiteral(a_extension_21.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_21Ident, __gong__toRawStringLiteral(a_extension_21.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_21Ident, a_extension_21.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_22 := range __gong__sortStageSetInstances(stage.A_extension_22s, stage.A_extension_22_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_22Ident := a_extension_22.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_22{Name: %s}).Stage(stage)", a_extension_22Ident, __gong__toRawStringLiteral(a_extension_22.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_22Ident, __gong__toRawStringLiteral(a_extension_22.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_22Ident, a_extension_22.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_23 := range __gong__sortStageSetInstances(stage.A_extension_23s, stage.A_extension_23_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_23Ident := a_extension_23.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_23{Name: %s}).Stage(stage)", a_extension_23Ident, __gong__toRawStringLiteral(a_extension_23.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_23Ident, __gong__toRawStringLiteral(a_extension_23.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_23Ident, a_extension_23.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_24 := range __gong__sortStageSetInstances(stage.A_extension_24s, stage.A_extension_24_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_24Ident := a_extension_24.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_24{Name: %s}).Stage(stage)", a_extension_24Ident, __gong__toRawStringLiteral(a_extension_24.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_24Ident, __gong__toRawStringLiteral(a_extension_24.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_24Ident, a_extension_24.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_25 := range __gong__sortStageSetInstances(stage.A_extension_25s, stage.A_extension_25_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_25Ident := a_extension_25.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_25{Name: %s}).Stage(stage)", a_extension_25Ident, __gong__toRawStringLiteral(a_extension_25.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_25Ident, __gong__toRawStringLiteral(a_extension_25.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_25Ident, a_extension_25.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_26 := range __gong__sortStageSetInstances(stage.A_extension_26s, stage.A_extension_26_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_26Ident := a_extension_26.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_26{Name: %s}).Stage(stage)", a_extension_26Ident, __gong__toRawStringLiteral(a_extension_26.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_26Ident, __gong__toRawStringLiteral(a_extension_26.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_26Ident, a_extension_26.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_27 := range __gong__sortStageSetInstances(stage.A_extension_27s, stage.A_extension_27_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_27Ident := a_extension_27.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_27{Name: %s}).Stage(stage)", a_extension_27Ident, __gong__toRawStringLiteral(a_extension_27.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_27Ident, __gong__toRawStringLiteral(a_extension_27.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_27Ident, a_extension_27.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_28 := range __gong__sortStageSetInstances(stage.A_extension_28s, stage.A_extension_28_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_28Ident := a_extension_28.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_28{Name: %s}).Stage(stage)", a_extension_28Ident, __gong__toRawStringLiteral(a_extension_28.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_28Ident, __gong__toRawStringLiteral(a_extension_28.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_28Ident, a_extension_28.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_29 := range __gong__sortStageSetInstances(stage.A_extension_29s, stage.A_extension_29_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_29Ident := a_extension_29.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_29{Name: %s}).Stage(stage)", a_extension_29Ident, __gong__toRawStringLiteral(a_extension_29.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_29Ident, __gong__toRawStringLiteral(a_extension_29.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_29Ident, a_extension_29.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_3 := range __gong__sortStageSetInstances(stage.A_extension_3s, stage.A_extension_3_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_3Ident := a_extension_3.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_3{Name: %s}).Stage(stage)", a_extension_3Ident, __gong__toRawStringLiteral(a_extension_3.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_3Ident, __gong__toRawStringLiteral(a_extension_3.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_3Ident, a_extension_3.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_30 := range __gong__sortStageSetInstances(stage.A_extension_30s, stage.A_extension_30_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_30Ident := a_extension_30.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_30{Name: %s}).Stage(stage)", a_extension_30Ident, __gong__toRawStringLiteral(a_extension_30.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_30Ident, __gong__toRawStringLiteral(a_extension_30.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_30Ident, a_extension_30.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_31 := range __gong__sortStageSetInstances(stage.A_extension_31s, stage.A_extension_31_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_31Ident := a_extension_31.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_31{Name: %s}).Stage(stage)", a_extension_31Ident, __gong__toRawStringLiteral(a_extension_31.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_31Ident, __gong__toRawStringLiteral(a_extension_31.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_31Ident, a_extension_31.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_32 := range __gong__sortStageSetInstances(stage.A_extension_32s, stage.A_extension_32_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_32Ident := a_extension_32.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_32{Name: %s}).Stage(stage)", a_extension_32Ident, __gong__toRawStringLiteral(a_extension_32.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_32Ident, __gong__toRawStringLiteral(a_extension_32.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_32Ident, a_extension_32.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_33 := range __gong__sortStageSetInstances(stage.A_extension_33s, stage.A_extension_33_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_33Ident := a_extension_33.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_33{Name: %s}).Stage(stage)", a_extension_33Ident, __gong__toRawStringLiteral(a_extension_33.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_33Ident, __gong__toRawStringLiteral(a_extension_33.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_33Ident, a_extension_33.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_34 := range __gong__sortStageSetInstances(stage.A_extension_34s, stage.A_extension_34_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_34Ident := a_extension_34.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_34{Name: %s}).Stage(stage)", a_extension_34Ident, __gong__toRawStringLiteral(a_extension_34.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_34Ident, __gong__toRawStringLiteral(a_extension_34.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_34Ident, a_extension_34.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_35 := range __gong__sortStageSetInstances(stage.A_extension_35s, stage.A_extension_35_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_35Ident := a_extension_35.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_35{Name: %s}).Stage(stage)", a_extension_35Ident, __gong__toRawStringLiteral(a_extension_35.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_35Ident, __gong__toRawStringLiteral(a_extension_35.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_35Ident, a_extension_35.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_36 := range __gong__sortStageSetInstances(stage.A_extension_36s, stage.A_extension_36_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_36Ident := a_extension_36.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_36{Name: %s}).Stage(stage)", a_extension_36Ident, __gong__toRawStringLiteral(a_extension_36.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_36Ident, __gong__toRawStringLiteral(a_extension_36.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_36Ident, a_extension_36.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_37 := range __gong__sortStageSetInstances(stage.A_extension_37s, stage.A_extension_37_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_37Ident := a_extension_37.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_37{Name: %s}).Stage(stage)", a_extension_37Ident, __gong__toRawStringLiteral(a_extension_37.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_37Ident, __gong__toRawStringLiteral(a_extension_37.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_37Ident, a_extension_37.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_38 := range __gong__sortStageSetInstances(stage.A_extension_38s, stage.A_extension_38_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_38Ident := a_extension_38.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_38{Name: %s}).Stage(stage)", a_extension_38Ident, __gong__toRawStringLiteral(a_extension_38.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_38Ident, __gong__toRawStringLiteral(a_extension_38.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_38Ident, a_extension_38.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_39 := range __gong__sortStageSetInstances(stage.A_extension_39s, stage.A_extension_39_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_39Ident := a_extension_39.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_39{Name: %s}).Stage(stage)", a_extension_39Ident, __gong__toRawStringLiteral(a_extension_39.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_39Ident, __gong__toRawStringLiteral(a_extension_39.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_39Ident, a_extension_39.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_4 := range __gong__sortStageSetInstances(stage.A_extension_4s, stage.A_extension_4_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_4Ident := a_extension_4.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_4{Name: %s}).Stage(stage)", a_extension_4Ident, __gong__toRawStringLiteral(a_extension_4.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_4Ident, __gong__toRawStringLiteral(a_extension_4.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_4Ident, a_extension_4.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_40 := range __gong__sortStageSetInstances(stage.A_extension_40s, stage.A_extension_40_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_40Ident := a_extension_40.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_40{Name: %s}).Stage(stage)", a_extension_40Ident, __gong__toRawStringLiteral(a_extension_40.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_40Ident, __gong__toRawStringLiteral(a_extension_40.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_40Ident, a_extension_40.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_41 := range __gong__sortStageSetInstances(stage.A_extension_41s, stage.A_extension_41_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_41Ident := a_extension_41.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_41{Name: %s}).Stage(stage)", a_extension_41Ident, __gong__toRawStringLiteral(a_extension_41.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_41Ident, __gong__toRawStringLiteral(a_extension_41.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_41Ident, a_extension_41.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_42 := range __gong__sortStageSetInstances(stage.A_extension_42s, stage.A_extension_42_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_42Ident := a_extension_42.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_42{Name: %s}).Stage(stage)", a_extension_42Ident, __gong__toRawStringLiteral(a_extension_42.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_42Ident, __gong__toRawStringLiteral(a_extension_42.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_42Ident, a_extension_42.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_43 := range __gong__sortStageSetInstances(stage.A_extension_43s, stage.A_extension_43_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_43Ident := a_extension_43.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_43{Name: %s}).Stage(stage)", a_extension_43Ident, __gong__toRawStringLiteral(a_extension_43.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_43Ident, __gong__toRawStringLiteral(a_extension_43.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_43Ident, a_extension_43.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_44 := range __gong__sortStageSetInstances(stage.A_extension_44s, stage.A_extension_44_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_44Ident := a_extension_44.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_44{Name: %s}).Stage(stage)", a_extension_44Ident, __gong__toRawStringLiteral(a_extension_44.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_44Ident, __gong__toRawStringLiteral(a_extension_44.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_44Ident, a_extension_44.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_45 := range __gong__sortStageSetInstances(stage.A_extension_45s, stage.A_extension_45_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_45Ident := a_extension_45.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_45{Name: %s}).Stage(stage)", a_extension_45Ident, __gong__toRawStringLiteral(a_extension_45.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_45Ident, __gong__toRawStringLiteral(a_extension_45.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_45Ident, a_extension_45.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_46 := range __gong__sortStageSetInstances(stage.A_extension_46s, stage.A_extension_46_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_46Ident := a_extension_46.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_46{Name: %s}).Stage(stage)", a_extension_46Ident, __gong__toRawStringLiteral(a_extension_46.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_46Ident, __gong__toRawStringLiteral(a_extension_46.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_46Ident, a_extension_46.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_47 := range __gong__sortStageSetInstances(stage.A_extension_47s, stage.A_extension_47_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_47Ident := a_extension_47.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_47{Name: %s}).Stage(stage)", a_extension_47Ident, __gong__toRawStringLiteral(a_extension_47.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_47Ident, __gong__toRawStringLiteral(a_extension_47.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_47Ident, a_extension_47.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_48 := range __gong__sortStageSetInstances(stage.A_extension_48s, stage.A_extension_48_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_48Ident := a_extension_48.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_48{Name: %s}).Stage(stage)", a_extension_48Ident, __gong__toRawStringLiteral(a_extension_48.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_48Ident, __gong__toRawStringLiteral(a_extension_48.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_48Ident, a_extension_48.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_49 := range __gong__sortStageSetInstances(stage.A_extension_49s, stage.A_extension_49_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_49Ident := a_extension_49.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_49{Name: %s}).Stage(stage)", a_extension_49Ident, __gong__toRawStringLiteral(a_extension_49.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_49Ident, __gong__toRawStringLiteral(a_extension_49.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_49Ident, a_extension_49.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_5 := range __gong__sortStageSetInstances(stage.A_extension_5s, stage.A_extension_5_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_5Ident := a_extension_5.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_5{Name: %s}).Stage(stage)", a_extension_5Ident, __gong__toRawStringLiteral(a_extension_5.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_5Ident, __gong__toRawStringLiteral(a_extension_5.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_5Ident, a_extension_5.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_50 := range __gong__sortStageSetInstances(stage.A_extension_50s, stage.A_extension_50_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_50Ident := a_extension_50.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_50{Name: %s}).Stage(stage)", a_extension_50Ident, __gong__toRawStringLiteral(a_extension_50.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_50Ident, __gong__toRawStringLiteral(a_extension_50.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_50Ident, a_extension_50.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_51 := range __gong__sortStageSetInstances(stage.A_extension_51s, stage.A_extension_51_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_51Ident := a_extension_51.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_51{Name: %s}).Stage(stage)", a_extension_51Ident, __gong__toRawStringLiteral(a_extension_51.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_51Ident, __gong__toRawStringLiteral(a_extension_51.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_51Ident, a_extension_51.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_52 := range __gong__sortStageSetInstances(stage.A_extension_52s, stage.A_extension_52_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_52Ident := a_extension_52.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_52{Name: %s}).Stage(stage)", a_extension_52Ident, __gong__toRawStringLiteral(a_extension_52.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_52Ident, __gong__toRawStringLiteral(a_extension_52.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_52Ident, a_extension_52.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_53 := range __gong__sortStageSetInstances(stage.A_extension_53s, stage.A_extension_53_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_53Ident := a_extension_53.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_53{Name: %s}).Stage(stage)", a_extension_53Ident, __gong__toRawStringLiteral(a_extension_53.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_53Ident, __gong__toRawStringLiteral(a_extension_53.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_53Ident, a_extension_53.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_54 := range __gong__sortStageSetInstances(stage.A_extension_54s, stage.A_extension_54_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_54Ident := a_extension_54.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_54{Name: %s}).Stage(stage)", a_extension_54Ident, __gong__toRawStringLiteral(a_extension_54.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_54Ident, __gong__toRawStringLiteral(a_extension_54.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_54Ident, a_extension_54.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_55 := range __gong__sortStageSetInstances(stage.A_extension_55s, stage.A_extension_55_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_55Ident := a_extension_55.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_55{Name: %s}).Stage(stage)", a_extension_55Ident, __gong__toRawStringLiteral(a_extension_55.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_55Ident, __gong__toRawStringLiteral(a_extension_55.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_55Ident, a_extension_55.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_56 := range __gong__sortStageSetInstances(stage.A_extension_56s, stage.A_extension_56_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_56Ident := a_extension_56.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_56{Name: %s}).Stage(stage)", a_extension_56Ident, __gong__toRawStringLiteral(a_extension_56.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_56Ident, __gong__toRawStringLiteral(a_extension_56.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_56Ident, a_extension_56.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_57 := range __gong__sortStageSetInstances(stage.A_extension_57s, stage.A_extension_57_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_57Ident := a_extension_57.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_57{Name: %s}).Stage(stage)", a_extension_57Ident, __gong__toRawStringLiteral(a_extension_57.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_57Ident, __gong__toRawStringLiteral(a_extension_57.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_57Ident, a_extension_57.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_58 := range __gong__sortStageSetInstances(stage.A_extension_58s, stage.A_extension_58_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_58Ident := a_extension_58.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_58{Name: %s}).Stage(stage)", a_extension_58Ident, __gong__toRawStringLiteral(a_extension_58.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_58Ident, __gong__toRawStringLiteral(a_extension_58.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_58Ident, a_extension_58.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_6 := range __gong__sortStageSetInstances(stage.A_extension_6s, stage.A_extension_6_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_6Ident := a_extension_6.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_6{Name: %s}).Stage(stage)", a_extension_6Ident, __gong__toRawStringLiteral(a_extension_6.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_6Ident, __gong__toRawStringLiteral(a_extension_6.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_6Ident, a_extension_6.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_7 := range __gong__sortStageSetInstances(stage.A_extension_7s, stage.A_extension_7_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_7Ident := a_extension_7.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_7{Name: %s}).Stage(stage)", a_extension_7Ident, __gong__toRawStringLiteral(a_extension_7.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_7Ident, __gong__toRawStringLiteral(a_extension_7.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_7Ident, a_extension_7.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_8 := range __gong__sortStageSetInstances(stage.A_extension_8s, stage.A_extension_8_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_8Ident := a_extension_8.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_8{Name: %s}).Stage(stage)", a_extension_8Ident, __gong__toRawStringLiteral(a_extension_8.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_8Ident, __gong__toRawStringLiteral(a_extension_8.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_8Ident, a_extension_8.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_9 := range __gong__sortStageSetInstances(stage.A_extension_9s, stage.A_extension_9_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_9Ident := a_extension_9.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_9{Name: %s}).Stage(stage)", a_extension_9Ident, __gong__toRawStringLiteral(a_extension_9.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_9Ident, __gong__toRawStringLiteral(a_extension_9.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_9Ident, a_extension_9.Owns))
		}
	}
	if stage != nil {
		for _, abstractapproachlegtype := range __gong__sortStageSetInstances(stage.AbstractApproachLegTypes, stage.AbstractApproachLegType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractapproachlegtypeIdent := abstractapproachlegtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractApproachLegType{Name: %s}).Stage(stage)", abstractapproachlegtypeIdent, __gong__toRawStringLiteral(abstractapproachlegtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractapproachlegtypeIdent, __gong__toRawStringLiteral(abstractapproachlegtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractapproachlegtypeIdent, __gong__toRawStringLiteral(abstractapproachlegtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractapproachlegtypeIdent, __gong__toRawStringLiteral(abstractapproachlegtype.Description)))
		}
	}
	if stage != nil {
		for _, abstractproceduretype := range __gong__sortStageSetInstances(stage.AbstractProcedureTypes, stage.AbstractProcedureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractproceduretypeIdent := abstractproceduretype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractProcedureType{Name: %s}).Stage(stage)", abstractproceduretypeIdent, __gong__toRawStringLiteral(abstractproceduretype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractproceduretypeIdent, __gong__toRawStringLiteral(abstractproceduretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractproceduretypeIdent, __gong__toRawStringLiteral(abstractproceduretype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractproceduretypeIdent, __gong__toRawStringLiteral(abstractproceduretype.Description)))
		}
	}
	if stage != nil {
		for _, abstractpropertieswithscheduletype := range __gong__sortStageSetInstances(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractpropertieswithscheduletypeIdent := abstractpropertieswithscheduletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractPropertiesWithScheduleType{Name: %s}).Stage(stage)", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Description)))
		}
	}
	if stage != nil {
		for _, abstractsegmentlegtype := range __gong__sortStageSetInstances(stage.AbstractSegmentLegTypes, stage.AbstractSegmentLegType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractsegmentlegtypeIdent := abstractsegmentlegtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractSegmentLegType{Name: %s}).Stage(stage)", abstractsegmentlegtypeIdent, __gong__toRawStringLiteral(abstractsegmentlegtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractsegmentlegtypeIdent, __gong__toRawStringLiteral(abstractsegmentlegtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractsegmentlegtypeIdent, __gong__toRawStringLiteral(abstractsegmentlegtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractsegmentlegtypeIdent, __gong__toRawStringLiteral(abstractsegmentlegtype.Description)))
		}
	}
	if stage != nil {
		for _, abstractsegmentpointtype := range __gong__sortStageSetInstances(stage.AbstractSegmentPointTypes, stage.AbstractSegmentPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractsegmentpointtypeIdent := abstractsegmentpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractSegmentPointType{Name: %s}).Stage(stage)", abstractsegmentpointtypeIdent, __gong__toRawStringLiteral(abstractsegmentpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractsegmentpointtypeIdent, __gong__toRawStringLiteral(abstractsegmentpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractsegmentpointtypeIdent, __gong__toRawStringLiteral(abstractsegmentpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractsegmentpointtypeIdent, __gong__toRawStringLiteral(abstractsegmentpointtype.Description)))
		}
	}
	if stage != nil {
		for _, aerialrefuellingpointtype := range __gong__sortStageSetInstances(stage.AerialRefuellingPointTypes, stage.AerialRefuellingPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aerialrefuellingpointtypeIdent := aerialrefuellingpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingPointType{Name: %s}).Stage(stage)", aerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointtype.Description)))
			for _, elem := range aerialrefuellingpointtype.FacilityMakeup {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FacilityMakeup = append(%s.FacilityMakeup, %s)", aerialrefuellingpointtypeIdent, aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_fixDesignatedPoint = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_navaidSystem = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_position = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_runwayPoint = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_aimingPoint = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_airportReferencePoint = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.ExtendedServiceVolume != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingpointtype.ExtendedServiceVolume.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtendedServiceVolume = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingpointtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aerialrefuellingpointtypeIdent, aerialrefuellingpointtypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingpointtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aerialrefuellingpointtypeIdent, aerialrefuellingpointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aircraftcharacteristicpropertytype := range __gong__sortStageSetInstances(stage.AircraftCharacteristicPropertyTypes, stage.AircraftCharacteristicPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aircraftcharacteristicpropertytypeIdent := aircraftcharacteristicpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftCharacteristicPropertyType{Name: %s}).Stage(stage)", aircraftcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(aircraftcharacteristicpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(aircraftcharacteristicpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", aircraftcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(aircraftcharacteristicpropertytype.NilReason)))
			if aircraftcharacteristicpropertytype.AircraftCharacteristic != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aircraftcharacteristicpropertytype.AircraftCharacteristic.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCharacteristic = %s", aircraftcharacteristicpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aircraftcharacteristictype := range __gong__sortStageSetInstances(stage.AircraftCharacteristicTypes, stage.AircraftCharacteristicType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aircraftcharacteristictypeIdent := aircraftcharacteristictype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftCharacteristicType{Name: %s}).Stage(stage)", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Description)))
			for _, elem := range aircraftcharacteristictype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aircraftcharacteristictypeIdent, aircraftcharacteristictypeIdent, targetIdent))
			}
			for _, elem := range aircraftcharacteristictype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aircraftcharacteristictypeIdent, aircraftcharacteristictypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airportheliportpropertytype := range __gong__sortStageSetInstances(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airportheliportpropertytypeIdent := airportheliportpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportHeliportPropertyType{Name: %s}).Stage(stage)", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airportheliportpropertytypeIdent, airportheliportpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, altimetersourcepropertytype := range __gong__sortStageSetInstances(stage.AltimeterSourcePropertyTypes, stage.AltimeterSourcePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			altimetersourcepropertytypeIdent := altimetersourcepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AltimeterSourcePropertyType{Name: %s}).Stage(stage)", altimetersourcepropertytypeIdent, __gong__toRawStringLiteral(altimetersourcepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", altimetersourcepropertytypeIdent, __gong__toRawStringLiteral(altimetersourcepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", altimetersourcepropertytypeIdent, altimetersourcepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", altimetersourcepropertytypeIdent, __gong__toRawStringLiteral(altimetersourcepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", altimetersourcepropertytypeIdent, __gong__toRawStringLiteral(altimetersourcepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", altimetersourcepropertytypeIdent, __gong__toRawStringLiteral(altimetersourcepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", altimetersourcepropertytypeIdent, __gong__toRawStringLiteral(altimetersourcepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", altimetersourcepropertytypeIdent, __gong__toRawStringLiteral(altimetersourcepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", altimetersourcepropertytypeIdent, __gong__toRawStringLiteral(altimetersourcepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", altimetersourcepropertytypeIdent, __gong__toRawStringLiteral(altimetersourcepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", altimetersourcepropertytypeIdent, __gong__toRawStringLiteral(altimetersourcepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, altitudeadjustmentpropertytype := range __gong__sortStageSetInstances(stage.AltitudeAdjustmentPropertyTypes, stage.AltitudeAdjustmentPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			altitudeadjustmentpropertytypeIdent := altitudeadjustmentpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AltitudeAdjustmentPropertyType{Name: %s}).Stage(stage)", altitudeadjustmentpropertytypeIdent, __gong__toRawStringLiteral(altitudeadjustmentpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", altitudeadjustmentpropertytypeIdent, __gong__toRawStringLiteral(altitudeadjustmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", altitudeadjustmentpropertytypeIdent, __gong__toRawStringLiteral(altitudeadjustmentpropertytype.NilReason)))
			if altitudeadjustmentpropertytype.AltitudeAdjustment != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := altitudeadjustmentpropertytype.AltitudeAdjustment.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AltitudeAdjustment = %s", altitudeadjustmentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, altitudeadjustmenttype := range __gong__sortStageSetInstances(stage.AltitudeAdjustmentTypes, stage.AltitudeAdjustmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			altitudeadjustmenttypeIdent := altitudeadjustmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AltitudeAdjustmentType{Name: %s}).Stage(stage)", altitudeadjustmenttypeIdent, __gong__toRawStringLiteral(altitudeadjustmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", altitudeadjustmenttypeIdent, __gong__toRawStringLiteral(altitudeadjustmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", altitudeadjustmenttypeIdent, __gong__toRawStringLiteral(altitudeadjustmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", altitudeadjustmenttypeIdent, __gong__toRawStringLiteral(altitudeadjustmenttype.Description)))
			for _, elem := range altitudeadjustmenttype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", altitudeadjustmenttypeIdent, altitudeadjustmenttypeIdent, targetIdent))
			}
			for _, elem := range altitudeadjustmenttype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", altitudeadjustmenttypeIdent, altitudeadjustmenttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, angleindicationpropertytype := range __gong__sortStageSetInstances(stage.AngleIndicationPropertyTypes, stage.AngleIndicationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			angleindicationpropertytypeIdent := angleindicationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleIndicationPropertyType{Name: %s}).Stage(stage)", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", angleindicationpropertytypeIdent, angleindicationpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, angleusepropertytype := range __gong__sortStageSetInstances(stage.AngleUsePropertyTypes, stage.AngleUsePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			angleusepropertytypeIdent := angleusepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleUsePropertyType{Name: %s}).Stage(stage)", angleusepropertytypeIdent, __gong__toRawStringLiteral(angleusepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleusepropertytypeIdent, __gong__toRawStringLiteral(angleusepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", angleusepropertytypeIdent, __gong__toRawStringLiteral(angleusepropertytype.NilReason)))
			if angleusepropertytype.AngleUse != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := angleusepropertytype.AngleUse.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleUse = %s", angleusepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, angleusetype := range __gong__sortStageSetInstances(stage.AngleUseTypes, stage.AngleUseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			angleusetypeIdent := angleusetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleUseType{Name: %s}).Stage(stage)", angleusetypeIdent, __gong__toRawStringLiteral(angleusetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleusetypeIdent, __gong__toRawStringLiteral(angleusetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", angleusetypeIdent, __gong__toRawStringLiteral(angleusetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", angleusetypeIdent, __gong__toRawStringLiteral(angleusetype.Description)))
			for _, elem := range angleusetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", angleusetypeIdent, angleusetypeIdent, targetIdent))
			}
			if angleusetype.TheAngleIndication != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := angleusetype.TheAngleIndication.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheAngleIndication = %s", angleusetypeIdent, targetIdent))
			}
			for _, elem := range angleusetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", angleusetypeIdent, angleusetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, approachaltitudetablepropertytype := range __gong__sortStageSetInstances(stage.ApproachAltitudeTablePropertyTypes, stage.ApproachAltitudeTablePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			approachaltitudetablepropertytypeIdent := approachaltitudetablepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ApproachAltitudeTablePropertyType{Name: %s}).Stage(stage)", approachaltitudetablepropertytypeIdent, __gong__toRawStringLiteral(approachaltitudetablepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", approachaltitudetablepropertytypeIdent, __gong__toRawStringLiteral(approachaltitudetablepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", approachaltitudetablepropertytypeIdent, __gong__toRawStringLiteral(approachaltitudetablepropertytype.NilReason)))
			if approachaltitudetablepropertytype.ApproachAltitudeTable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := approachaltitudetablepropertytype.ApproachAltitudeTable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApproachAltitudeTable = %s", approachaltitudetablepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, approachaltitudetabletype := range __gong__sortStageSetInstances(stage.ApproachAltitudeTableTypes, stage.ApproachAltitudeTableType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			approachaltitudetabletypeIdent := approachaltitudetabletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ApproachAltitudeTableType{Name: %s}).Stage(stage)", approachaltitudetabletypeIdent, __gong__toRawStringLiteral(approachaltitudetabletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", approachaltitudetabletypeIdent, __gong__toRawStringLiteral(approachaltitudetabletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", approachaltitudetabletypeIdent, __gong__toRawStringLiteral(approachaltitudetabletype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", approachaltitudetabletypeIdent, __gong__toRawStringLiteral(approachaltitudetabletype.Description)))
			for _, elem := range approachaltitudetabletype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", approachaltitudetabletypeIdent, approachaltitudetabletypeIdent, targetIdent))
			}
			for _, elem := range approachaltitudetabletype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", approachaltitudetabletypeIdent, approachaltitudetabletypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, approachconditionpropertytype := range __gong__sortStageSetInstances(stage.ApproachConditionPropertyTypes, stage.ApproachConditionPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			approachconditionpropertytypeIdent := approachconditionpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ApproachConditionPropertyType{Name: %s}).Stage(stage)", approachconditionpropertytypeIdent, __gong__toRawStringLiteral(approachconditionpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", approachconditionpropertytypeIdent, __gong__toRawStringLiteral(approachconditionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", approachconditionpropertytypeIdent, __gong__toRawStringLiteral(approachconditionpropertytype.NilReason)))
			if approachconditionpropertytype.ApproachCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := approachconditionpropertytype.ApproachCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApproachCondition = %s", approachconditionpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, approachconditiontype := range __gong__sortStageSetInstances(stage.ApproachConditionTypes, stage.ApproachConditionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			approachconditiontypeIdent := approachconditiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ApproachConditionType{Name: %s}).Stage(stage)", approachconditiontypeIdent, __gong__toRawStringLiteral(approachconditiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", approachconditiontypeIdent, __gong__toRawStringLiteral(approachconditiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", approachconditiontypeIdent, __gong__toRawStringLiteral(approachconditiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", approachconditiontypeIdent, __gong__toRawStringLiteral(approachconditiontype.Description)))
			if approachconditiontype.MinimumSet != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := approachconditiontype.MinimumSet.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinimumSet = %s", approachconditiontypeIdent, targetIdent))
			}
			for _, elem := range approachconditiontype.CirclingRestriction {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CirclingRestriction = append(%s.CirclingRestriction, %s)", approachconditiontypeIdent, approachconditiontypeIdent, targetIdent))
			}
			for _, elem := range approachconditiontype.AircraftCategory {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = append(%s.AircraftCategory, %s)", approachconditiontypeIdent, approachconditiontypeIdent, targetIdent))
			}
			for _, elem := range approachconditiontype.LandingArea {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LandingArea = append(%s.LandingArea, %s)", approachconditiontypeIdent, approachconditiontypeIdent, targetIdent))
			}
			if approachconditiontype.Altimeter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := approachconditiontype.Altimeter.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Altimeter = %s", approachconditiontypeIdent, targetIdent))
			}
			for _, elem := range approachconditiontype.DesignSurface {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignSurface = append(%s.DesignSurface, %s)", approachconditiontypeIdent, approachconditiontypeIdent, targetIdent))
			}
			for _, elem := range approachconditiontype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", approachconditiontypeIdent, approachconditiontypeIdent, targetIdent))
			}
			for _, elem := range approachconditiontype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", approachconditiontypeIdent, approachconditiontypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, approachdistancetablepropertytype := range __gong__sortStageSetInstances(stage.ApproachDistanceTablePropertyTypes, stage.ApproachDistanceTablePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			approachdistancetablepropertytypeIdent := approachdistancetablepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ApproachDistanceTablePropertyType{Name: %s}).Stage(stage)", approachdistancetablepropertytypeIdent, __gong__toRawStringLiteral(approachdistancetablepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", approachdistancetablepropertytypeIdent, __gong__toRawStringLiteral(approachdistancetablepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", approachdistancetablepropertytypeIdent, __gong__toRawStringLiteral(approachdistancetablepropertytype.NilReason)))
			if approachdistancetablepropertytype.ApproachDistanceTable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := approachdistancetablepropertytype.ApproachDistanceTable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApproachDistanceTable = %s", approachdistancetablepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, approachdistancetabletype := range __gong__sortStageSetInstances(stage.ApproachDistanceTableTypes, stage.ApproachDistanceTableType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			approachdistancetabletypeIdent := approachdistancetabletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ApproachDistanceTableType{Name: %s}).Stage(stage)", approachdistancetabletypeIdent, __gong__toRawStringLiteral(approachdistancetabletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", approachdistancetabletypeIdent, __gong__toRawStringLiteral(approachdistancetabletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", approachdistancetabletypeIdent, __gong__toRawStringLiteral(approachdistancetabletype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", approachdistancetabletypeIdent, __gong__toRawStringLiteral(approachdistancetabletype.Description)))
			for _, elem := range approachdistancetabletype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", approachdistancetabletypeIdent, approachdistancetabletypeIdent, targetIdent))
			}
			for _, elem := range approachdistancetabletype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", approachdistancetabletypeIdent, approachdistancetabletypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, approachtimingtablepropertytype := range __gong__sortStageSetInstances(stage.ApproachTimingTablePropertyTypes, stage.ApproachTimingTablePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			approachtimingtablepropertytypeIdent := approachtimingtablepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ApproachTimingTablePropertyType{Name: %s}).Stage(stage)", approachtimingtablepropertytypeIdent, __gong__toRawStringLiteral(approachtimingtablepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", approachtimingtablepropertytypeIdent, __gong__toRawStringLiteral(approachtimingtablepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", approachtimingtablepropertytypeIdent, __gong__toRawStringLiteral(approachtimingtablepropertytype.NilReason)))
			if approachtimingtablepropertytype.ApproachTimingTable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := approachtimingtablepropertytype.ApproachTimingTable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApproachTimingTable = %s", approachtimingtablepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, approachtimingtabletype := range __gong__sortStageSetInstances(stage.ApproachTimingTableTypes, stage.ApproachTimingTableType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			approachtimingtabletypeIdent := approachtimingtabletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ApproachTimingTableType{Name: %s}).Stage(stage)", approachtimingtabletypeIdent, __gong__toRawStringLiteral(approachtimingtabletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", approachtimingtabletypeIdent, __gong__toRawStringLiteral(approachtimingtabletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", approachtimingtabletypeIdent, __gong__toRawStringLiteral(approachtimingtabletype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", approachtimingtabletypeIdent, __gong__toRawStringLiteral(approachtimingtabletype.Description)))
			for _, elem := range approachtimingtabletype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", approachtimingtabletypeIdent, approachtimingtabletypeIdent, targetIdent))
			}
			for _, elem := range approachtimingtabletype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", approachtimingtabletypeIdent, approachtimingtabletypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, arrivalfeederlegtimeslicepropertytype := range __gong__sortStageSetInstances(stage.ArrivalFeederLegTimeSlicePropertyTypes, stage.ArrivalFeederLegTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arrivalfeederlegtimeslicepropertytypeIdent := arrivalfeederlegtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ArrivalFeederLegTimeSlicePropertyType{Name: %s}).Stage(stage)", arrivalfeederlegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrivalfeederlegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", arrivalfeederlegtimeslicepropertytypeIdent, arrivalfeederlegtimeslicepropertytype.Owns))
			if arrivalfeederlegtimeslicepropertytype.ArrivalFeederLegTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivalfeederlegtimeslicepropertytype.ArrivalFeederLegTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArrivalFeederLegTimeSlice = %s", arrivalfeederlegtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, arrivalfeederlegtimeslicetype := range __gong__sortStageSetInstances(stage.ArrivalFeederLegTimeSliceTypes, stage.ArrivalFeederLegTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arrivalfeederlegtimeslicetypeIdent := arrivalfeederlegtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ArrivalFeederLegTimeSliceType{Name: %s}).Stage(stage)", arrivalfeederlegtimeslicetypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrivalfeederlegtimeslicetypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", arrivalfeederlegtimeslicetypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", arrivalfeederlegtimeslicetypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", arrivalfeederlegtimeslicetypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", arrivalfeederlegtimeslicetypeIdent, arrivalfeederlegtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", arrivalfeederlegtimeslicetypeIdent, arrivalfeederlegtimeslicetype.CorrectionNumber))
			if arrivalfeederlegtimeslicetype.StartPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivalfeederlegtimeslicetype.StartPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartPoint = %s", arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			if arrivalfeederlegtimeslicetype.EndPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivalfeederlegtimeslicetype.EndPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPoint = %s", arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			if arrivalfeederlegtimeslicetype.Trajectory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivalfeederlegtimeslicetype.Trajectory.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Trajectory = %s", arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			if arrivalfeederlegtimeslicetype.ArcCentre != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivalfeederlegtimeslicetype.ArcCentre.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcCentre = %s", arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			if arrivalfeederlegtimeslicetype.Angle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivalfeederlegtimeslicetype.Angle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle = %s", arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			if arrivalfeederlegtimeslicetype.Distance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivalfeederlegtimeslicetype.Distance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range arrivalfeederlegtimeslicetype.AircraftCategory {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = append(%s.AircraftCategory, %s)", arrivalfeederlegtimeslicetypeIdent, arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			if arrivalfeederlegtimeslicetype.Holding != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivalfeederlegtimeslicetype.Holding.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Holding = %s", arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range arrivalfeederlegtimeslicetype.DesignSurface {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignSurface = append(%s.DesignSurface, %s)", arrivalfeederlegtimeslicetypeIdent, arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range arrivalfeederlegtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", arrivalfeederlegtimeslicetypeIdent, arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			if arrivalfeederlegtimeslicetype.Approach != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivalfeederlegtimeslicetype.Approach.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Approach = %s", arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range arrivalfeederlegtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", arrivalfeederlegtimeslicetypeIdent, arrivalfeederlegtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, arrivalfeederlegtype := range __gong__sortStageSetInstances(stage.ArrivalFeederLegTypes, stage.ArrivalFeederLegType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arrivalfeederlegtypeIdent := arrivalfeederlegtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ArrivalFeederLegType{Name: %s}).Stage(stage)", arrivalfeederlegtypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrivalfeederlegtypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", arrivalfeederlegtypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", arrivalfeederlegtypeIdent, __gong__toRawStringLiteral(arrivalfeederlegtype.Description)))
			for _, elem := range arrivalfeederlegtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", arrivalfeederlegtypeIdent, arrivalfeederlegtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, arrivallegtimeslicepropertytype := range __gong__sortStageSetInstances(stage.ArrivalLegTimeSlicePropertyTypes, stage.ArrivalLegTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arrivallegtimeslicepropertytypeIdent := arrivallegtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ArrivalLegTimeSlicePropertyType{Name: %s}).Stage(stage)", arrivallegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(arrivallegtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrivallegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(arrivallegtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", arrivallegtimeslicepropertytypeIdent, arrivallegtimeslicepropertytype.Owns))
			if arrivallegtimeslicepropertytype.ArrivalLegTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivallegtimeslicepropertytype.ArrivalLegTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArrivalLegTimeSlice = %s", arrivallegtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, arrivallegtimeslicetype := range __gong__sortStageSetInstances(stage.ArrivalLegTimeSliceTypes, stage.ArrivalLegTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arrivallegtimeslicetypeIdent := arrivallegtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ArrivalLegTimeSliceType{Name: %s}).Stage(stage)", arrivallegtimeslicetypeIdent, __gong__toRawStringLiteral(arrivallegtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrivallegtimeslicetypeIdent, __gong__toRawStringLiteral(arrivallegtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", arrivallegtimeslicetypeIdent, __gong__toRawStringLiteral(arrivallegtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", arrivallegtimeslicetypeIdent, __gong__toRawStringLiteral(arrivallegtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", arrivallegtimeslicetypeIdent, __gong__toRawStringLiteral(arrivallegtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", arrivallegtimeslicetypeIdent, arrivallegtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", arrivallegtimeslicetypeIdent, arrivallegtimeslicetype.CorrectionNumber))
			if arrivallegtimeslicetype.StartPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivallegtimeslicetype.StartPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartPoint = %s", arrivallegtimeslicetypeIdent, targetIdent))
			}
			if arrivallegtimeslicetype.EndPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivallegtimeslicetype.EndPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPoint = %s", arrivallegtimeslicetypeIdent, targetIdent))
			}
			if arrivallegtimeslicetype.Trajectory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivallegtimeslicetype.Trajectory.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Trajectory = %s", arrivallegtimeslicetypeIdent, targetIdent))
			}
			if arrivallegtimeslicetype.ArcCentre != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivallegtimeslicetype.ArcCentre.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcCentre = %s", arrivallegtimeslicetypeIdent, targetIdent))
			}
			if arrivallegtimeslicetype.Angle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivallegtimeslicetype.Angle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle = %s", arrivallegtimeslicetypeIdent, targetIdent))
			}
			if arrivallegtimeslicetype.Distance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivallegtimeslicetype.Distance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", arrivallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range arrivallegtimeslicetype.AircraftCategory {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = append(%s.AircraftCategory, %s)", arrivallegtimeslicetypeIdent, arrivallegtimeslicetypeIdent, targetIdent))
			}
			if arrivallegtimeslicetype.Holding != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivallegtimeslicetype.Holding.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Holding = %s", arrivallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range arrivallegtimeslicetype.DesignSurface {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignSurface = append(%s.DesignSurface, %s)", arrivallegtimeslicetypeIdent, arrivallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range arrivallegtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", arrivallegtimeslicetypeIdent, arrivallegtimeslicetypeIdent, targetIdent))
			}
			if arrivallegtimeslicetype.Arrival != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrivallegtimeslicetype.Arrival.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arrival = %s", arrivallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range arrivallegtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", arrivallegtimeslicetypeIdent, arrivallegtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, arrivallegtype := range __gong__sortStageSetInstances(stage.ArrivalLegTypes, stage.ArrivalLegType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arrivallegtypeIdent := arrivallegtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ArrivalLegType{Name: %s}).Stage(stage)", arrivallegtypeIdent, __gong__toRawStringLiteral(arrivallegtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrivallegtypeIdent, __gong__toRawStringLiteral(arrivallegtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", arrivallegtypeIdent, __gong__toRawStringLiteral(arrivallegtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", arrivallegtypeIdent, __gong__toRawStringLiteral(arrivallegtype.Description)))
			for _, elem := range arrivallegtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", arrivallegtypeIdent, arrivallegtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, circlesectorpropertytype := range __gong__sortStageSetInstances(stage.CircleSectorPropertyTypes, stage.CircleSectorPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circlesectorpropertytypeIdent := circlesectorpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CircleSectorPropertyType{Name: %s}).Stage(stage)", circlesectorpropertytypeIdent, __gong__toRawStringLiteral(circlesectorpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circlesectorpropertytypeIdent, __gong__toRawStringLiteral(circlesectorpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", circlesectorpropertytypeIdent, __gong__toRawStringLiteral(circlesectorpropertytype.NilReason)))
			if circlesectorpropertytype.CircleSector != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlesectorpropertytype.CircleSector.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CircleSector = %s", circlesectorpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, circlesectortype := range __gong__sortStageSetInstances(stage.CircleSectorTypes, stage.CircleSectorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circlesectortypeIdent := circlesectortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CircleSectorType{Name: %s}).Stage(stage)", circlesectortypeIdent, __gong__toRawStringLiteral(circlesectortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circlesectortypeIdent, __gong__toRawStringLiteral(circlesectortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", circlesectortypeIdent, __gong__toRawStringLiteral(circlesectortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", circlesectortypeIdent, __gong__toRawStringLiteral(circlesectortype.Description)))
			for _, elem := range circlesectortype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", circlesectortypeIdent, circlesectortypeIdent, targetIdent))
			}
			for _, elem := range circlesectortype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", circlesectortypeIdent, circlesectortypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, circlingareatimeslicepropertytype := range __gong__sortStageSetInstances(stage.CirclingAreaTimeSlicePropertyTypes, stage.CirclingAreaTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circlingareatimeslicepropertytypeIdent := circlingareatimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CirclingAreaTimeSlicePropertyType{Name: %s}).Stage(stage)", circlingareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(circlingareatimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circlingareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(circlingareatimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", circlingareatimeslicepropertytypeIdent, circlingareatimeslicepropertytype.Owns))
			if circlingareatimeslicepropertytype.CirclingAreaTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlingareatimeslicepropertytype.CirclingAreaTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CirclingAreaTimeSlice = %s", circlingareatimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, circlingareatimeslicetype := range __gong__sortStageSetInstances(stage.CirclingAreaTimeSliceTypes, stage.CirclingAreaTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circlingareatimeslicetypeIdent := circlingareatimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CirclingAreaTimeSliceType{Name: %s}).Stage(stage)", circlingareatimeslicetypeIdent, __gong__toRawStringLiteral(circlingareatimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circlingareatimeslicetypeIdent, __gong__toRawStringLiteral(circlingareatimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", circlingareatimeslicetypeIdent, __gong__toRawStringLiteral(circlingareatimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", circlingareatimeslicetypeIdent, __gong__toRawStringLiteral(circlingareatimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", circlingareatimeslicetypeIdent, __gong__toRawStringLiteral(circlingareatimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", circlingareatimeslicetypeIdent, circlingareatimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", circlingareatimeslicetypeIdent, circlingareatimeslicetype.CorrectionNumber))
			if circlingareatimeslicetype.Extent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlingareatimeslicetype.Extent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extent = %s", circlingareatimeslicetypeIdent, targetIdent))
			}
			if circlingareatimeslicetype.Approach != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlingareatimeslicetype.Approach.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Approach = %s", circlingareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range circlingareatimeslicetype.Condition {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Condition = append(%s.Condition, %s)", circlingareatimeslicetypeIdent, circlingareatimeslicetypeIdent, targetIdent))
			}
			if circlingareatimeslicetype.AircraftCategory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlingareatimeslicetype.AircraftCategory.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = %s", circlingareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range circlingareatimeslicetype.DesignSurface {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignSurface = append(%s.DesignSurface, %s)", circlingareatimeslicetypeIdent, circlingareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range circlingareatimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", circlingareatimeslicetypeIdent, circlingareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range circlingareatimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", circlingareatimeslicetypeIdent, circlingareatimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, circlingareatype := range __gong__sortStageSetInstances(stage.CirclingAreaTypes, stage.CirclingAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circlingareatypeIdent := circlingareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CirclingAreaType{Name: %s}).Stage(stage)", circlingareatypeIdent, __gong__toRawStringLiteral(circlingareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circlingareatypeIdent, __gong__toRawStringLiteral(circlingareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", circlingareatypeIdent, __gong__toRawStringLiteral(circlingareatype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", circlingareatypeIdent, __gong__toRawStringLiteral(circlingareatype.Description)))
			for _, elem := range circlingareatype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", circlingareatypeIdent, circlingareatypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, circlingrestrictionpropertytype := range __gong__sortStageSetInstances(stage.CirclingRestrictionPropertyTypes, stage.CirclingRestrictionPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circlingrestrictionpropertytypeIdent := circlingrestrictionpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CirclingRestrictionPropertyType{Name: %s}).Stage(stage)", circlingrestrictionpropertytypeIdent, __gong__toRawStringLiteral(circlingrestrictionpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circlingrestrictionpropertytypeIdent, __gong__toRawStringLiteral(circlingrestrictionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", circlingrestrictionpropertytypeIdent, __gong__toRawStringLiteral(circlingrestrictionpropertytype.NilReason)))
			if circlingrestrictionpropertytype.CirclingRestriction != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlingrestrictionpropertytype.CirclingRestriction.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CirclingRestriction = %s", circlingrestrictionpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, circlingrestrictiontype := range __gong__sortStageSetInstances(stage.CirclingRestrictionTypes, stage.CirclingRestrictionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circlingrestrictiontypeIdent := circlingrestrictiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CirclingRestrictionType{Name: %s}).Stage(stage)", circlingrestrictiontypeIdent, __gong__toRawStringLiteral(circlingrestrictiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circlingrestrictiontypeIdent, __gong__toRawStringLiteral(circlingrestrictiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", circlingrestrictiontypeIdent, __gong__toRawStringLiteral(circlingrestrictiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", circlingrestrictiontypeIdent, __gong__toRawStringLiteral(circlingrestrictiontype.Description)))
			for _, elem := range circlingrestrictiontype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", circlingrestrictiontypeIdent, circlingrestrictiontypeIdent, targetIdent))
			}
			for _, elem := range circlingrestrictiontype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", circlingrestrictiontypeIdent, circlingrestrictiontypeIdent, targetIdent))
			}
			for _, elem := range circlingrestrictiontype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", circlingrestrictiontypeIdent, circlingrestrictiontypeIdent, targetIdent))
			}
			if circlingrestrictiontype.SectorDescription != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlingrestrictiontype.SectorDescription.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SectorDescription = %s", circlingrestrictiontypeIdent, targetIdent))
			}
			if circlingrestrictiontype.RestrictionArea != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlingrestrictiontype.RestrictionArea.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RestrictionArea = %s", circlingrestrictiontypeIdent, targetIdent))
			}
			for _, elem := range circlingrestrictiontype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", circlingrestrictiontypeIdent, circlingrestrictiontypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, curvepropertytype := range __gong__sortStageSetInstances(stage.CurvePropertyTypes, stage.CurvePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			curvepropertytypeIdent := curvepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CurvePropertyType{Name: %s}).Stage(stage)", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.NilReason)))
			if curvepropertytype.Curve != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := curvepropertytype.Curve.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Curve = %s", curvepropertytypeIdent, targetIdent))
			}
			if curvepropertytype.ElevatedCurve != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := curvepropertytype.ElevatedCurve.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedCurve = %s", curvepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, curvetype := range __gong__sortStageSetInstances(stage.CurveTypes, stage.CurveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			curvetypeIdent := curvetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CurveType{Name: %s}).Stage(stage)", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.SrsName)))
			for _, elem := range curvetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", curvetypeIdent, curvetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, dmepropertytype := range __gong__sortStageSetInstances(stage.DMEPropertyTypes, stage.DMEPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dmepropertytypeIdent := dmepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DMEPropertyType{Name: %s}).Stage(stage)", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", dmepropertytypeIdent, dmepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, departurearrivalconditionpropertytype := range __gong__sortStageSetInstances(stage.DepartureArrivalConditionPropertyTypes, stage.DepartureArrivalConditionPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			departurearrivalconditionpropertytypeIdent := departurearrivalconditionpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DepartureArrivalConditionPropertyType{Name: %s}).Stage(stage)", departurearrivalconditionpropertytypeIdent, __gong__toRawStringLiteral(departurearrivalconditionpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", departurearrivalconditionpropertytypeIdent, __gong__toRawStringLiteral(departurearrivalconditionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", departurearrivalconditionpropertytypeIdent, __gong__toRawStringLiteral(departurearrivalconditionpropertytype.NilReason)))
			if departurearrivalconditionpropertytype.DepartureArrivalCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurearrivalconditionpropertytype.DepartureArrivalCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DepartureArrivalCondition = %s", departurearrivalconditionpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, departurearrivalconditiontype := range __gong__sortStageSetInstances(stage.DepartureArrivalConditionTypes, stage.DepartureArrivalConditionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			departurearrivalconditiontypeIdent := departurearrivalconditiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DepartureArrivalConditionType{Name: %s}).Stage(stage)", departurearrivalconditiontypeIdent, __gong__toRawStringLiteral(departurearrivalconditiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", departurearrivalconditiontypeIdent, __gong__toRawStringLiteral(departurearrivalconditiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", departurearrivalconditiontypeIdent, __gong__toRawStringLiteral(departurearrivalconditiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", departurearrivalconditiontypeIdent, __gong__toRawStringLiteral(departurearrivalconditiontype.Description)))
			if departurearrivalconditiontype.EngineType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurearrivalconditiontype.EngineType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EngineType = %s", departurearrivalconditiontypeIdent, targetIdent))
			}
			for _, elem := range departurearrivalconditiontype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", departurearrivalconditiontypeIdent, departurearrivalconditiontypeIdent, targetIdent))
			}
			for _, elem := range departurearrivalconditiontype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", departurearrivalconditiontypeIdent, departurearrivalconditiontypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, departurelegtimeslicepropertytype := range __gong__sortStageSetInstances(stage.DepartureLegTimeSlicePropertyTypes, stage.DepartureLegTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			departurelegtimeslicepropertytypeIdent := departurelegtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DepartureLegTimeSlicePropertyType{Name: %s}).Stage(stage)", departurelegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(departurelegtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", departurelegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(departurelegtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", departurelegtimeslicepropertytypeIdent, departurelegtimeslicepropertytype.Owns))
			if departurelegtimeslicepropertytype.DepartureLegTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurelegtimeslicepropertytype.DepartureLegTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DepartureLegTimeSlice = %s", departurelegtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, departurelegtimeslicetype := range __gong__sortStageSetInstances(stage.DepartureLegTimeSliceTypes, stage.DepartureLegTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			departurelegtimeslicetypeIdent := departurelegtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DepartureLegTimeSliceType{Name: %s}).Stage(stage)", departurelegtimeslicetypeIdent, __gong__toRawStringLiteral(departurelegtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", departurelegtimeslicetypeIdent, __gong__toRawStringLiteral(departurelegtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", departurelegtimeslicetypeIdent, __gong__toRawStringLiteral(departurelegtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", departurelegtimeslicetypeIdent, __gong__toRawStringLiteral(departurelegtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", departurelegtimeslicetypeIdent, __gong__toRawStringLiteral(departurelegtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", departurelegtimeslicetypeIdent, departurelegtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", departurelegtimeslicetypeIdent, departurelegtimeslicetype.CorrectionNumber))
			if departurelegtimeslicetype.StartPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurelegtimeslicetype.StartPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartPoint = %s", departurelegtimeslicetypeIdent, targetIdent))
			}
			if departurelegtimeslicetype.EndPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurelegtimeslicetype.EndPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPoint = %s", departurelegtimeslicetypeIdent, targetIdent))
			}
			if departurelegtimeslicetype.Trajectory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurelegtimeslicetype.Trajectory.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Trajectory = %s", departurelegtimeslicetypeIdent, targetIdent))
			}
			if departurelegtimeslicetype.ArcCentre != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurelegtimeslicetype.ArcCentre.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcCentre = %s", departurelegtimeslicetypeIdent, targetIdent))
			}
			if departurelegtimeslicetype.Angle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurelegtimeslicetype.Angle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle = %s", departurelegtimeslicetypeIdent, targetIdent))
			}
			if departurelegtimeslicetype.Distance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurelegtimeslicetype.Distance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", departurelegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range departurelegtimeslicetype.AircraftCategory {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = append(%s.AircraftCategory, %s)", departurelegtimeslicetypeIdent, departurelegtimeslicetypeIdent, targetIdent))
			}
			if departurelegtimeslicetype.Holding != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurelegtimeslicetype.Holding.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Holding = %s", departurelegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range departurelegtimeslicetype.DesignSurface {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignSurface = append(%s.DesignSurface, %s)", departurelegtimeslicetypeIdent, departurelegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range departurelegtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", departurelegtimeslicetypeIdent, departurelegtimeslicetypeIdent, targetIdent))
			}
			if departurelegtimeslicetype.Departure != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := departurelegtimeslicetype.Departure.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Departure = %s", departurelegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range departurelegtimeslicetype.Condition {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Condition = append(%s.Condition, %s)", departurelegtimeslicetypeIdent, departurelegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range departurelegtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", departurelegtimeslicetypeIdent, departurelegtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, departurelegtype := range __gong__sortStageSetInstances(stage.DepartureLegTypes, stage.DepartureLegType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			departurelegtypeIdent := departurelegtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DepartureLegType{Name: %s}).Stage(stage)", departurelegtypeIdent, __gong__toRawStringLiteral(departurelegtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", departurelegtypeIdent, __gong__toRawStringLiteral(departurelegtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", departurelegtypeIdent, __gong__toRawStringLiteral(departurelegtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", departurelegtypeIdent, __gong__toRawStringLiteral(departurelegtype.Description)))
			for _, elem := range departurelegtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", departurelegtypeIdent, departurelegtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, designatedpointpropertytype := range __gong__sortStageSetInstances(stage.DesignatedPointPropertyTypes, stage.DesignatedPointPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			designatedpointpropertytypeIdent := designatedpointpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DesignatedPointPropertyType{Name: %s}).Stage(stage)", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", designatedpointpropertytypeIdent, designatedpointpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, distanceindicationpropertytype := range __gong__sortStageSetInstances(stage.DistanceIndicationPropertyTypes, stage.DistanceIndicationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			distanceindicationpropertytypeIdent := distanceindicationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DistanceIndicationPropertyType{Name: %s}).Stage(stage)", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", distanceindicationpropertytypeIdent, distanceindicationpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, elevatedcurvetype := range __gong__sortStageSetInstances(stage.ElevatedCurveTypes, stage.ElevatedCurveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevatedcurvetypeIdent := elevatedcurvetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedCurveType{Name: %s}).Stage(stage)", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.SrsName)))
			for _, elem := range elevatedcurvetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevatedcurvetypeIdent, elevatedcurvetypeIdent, targetIdent))
			}
			for _, elem := range elevatedcurvetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevatedcurvetypeIdent, elevatedcurvetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, elevatedpointtype := range __gong__sortStageSetInstances(stage.ElevatedPointTypes, stage.ElevatedPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevatedpointtypeIdent := elevatedpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedPointType{Name: %s}).Stage(stage)", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.SrsName)))
			for _, elem := range elevatedpointtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevatedpointtypeIdent, elevatedpointtypeIdent, targetIdent))
			}
			for _, elem := range elevatedpointtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevatedpointtypeIdent, elevatedpointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, elevatedsurfacetype := range __gong__sortStageSetInstances(stage.ElevatedSurfaceTypes, stage.ElevatedSurfaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevatedsurfacetypeIdent := elevatedsurfacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedSurfaceType{Name: %s}).Stage(stage)", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.SrsName)))
			for _, elem := range elevatedsurfacetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevatedsurfacetypeIdent, elevatedsurfacetypeIdent, targetIdent))
			}
			for _, elem := range elevatedsurfacetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevatedsurfacetypeIdent, elevatedsurfacetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, enroutesegmentpointtype := range __gong__sortStageSetInstances(stage.EnRouteSegmentPointTypes, stage.EnRouteSegmentPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			enroutesegmentpointtypeIdent := enroutesegmentpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.EnRouteSegmentPointType{Name: %s}).Stage(stage)", enroutesegmentpointtypeIdent, __gong__toRawStringLiteral(enroutesegmentpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", enroutesegmentpointtypeIdent, __gong__toRawStringLiteral(enroutesegmentpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", enroutesegmentpointtypeIdent, __gong__toRawStringLiteral(enroutesegmentpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", enroutesegmentpointtypeIdent, __gong__toRawStringLiteral(enroutesegmentpointtype.Description)))
			for _, elem := range enroutesegmentpointtype.FacilityMakeup {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FacilityMakeup = append(%s.FacilityMakeup, %s)", enroutesegmentpointtypeIdent, enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_fixDesignatedPoint = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_navaidSystem = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_position = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_runwayPoint = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_aimingPoint = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_airportReferencePoint = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.ExtendedServiceVolume != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := enroutesegmentpointtype.ExtendedServiceVolume.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtendedServiceVolume = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			for _, elem := range enroutesegmentpointtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", enroutesegmentpointtypeIdent, enroutesegmentpointtypeIdent, targetIdent))
			}
			for _, elem := range enroutesegmentpointtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", enroutesegmentpointtypeIdent, enroutesegmentpointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, equipmentunavailableadjustmentcolumnpropertytype := range __gong__sortStageSetInstances(stage.EquipmentUnavailableAdjustmentColumnPropertyTypes, stage.EquipmentUnavailableAdjustmentColumnPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			equipmentunavailableadjustmentcolumnpropertytypeIdent := equipmentunavailableadjustmentcolumnpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.EquipmentUnavailableAdjustmentColumnPropertyType{Name: %s}).Stage(stage)", equipmentunavailableadjustmentcolumnpropertytypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmentcolumnpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", equipmentunavailableadjustmentcolumnpropertytypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmentcolumnpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", equipmentunavailableadjustmentcolumnpropertytypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmentcolumnpropertytype.NilReason)))
			if equipmentunavailableadjustmentcolumnpropertytype.EquipmentUnavailableAdjustmentColumn != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := equipmentunavailableadjustmentcolumnpropertytype.EquipmentUnavailableAdjustmentColumn.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EquipmentUnavailableAdjustmentColumn = %s", equipmentunavailableadjustmentcolumnpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, equipmentunavailableadjustmentcolumntype := range __gong__sortStageSetInstances(stage.EquipmentUnavailableAdjustmentColumnTypes, stage.EquipmentUnavailableAdjustmentColumnType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			equipmentunavailableadjustmentcolumntypeIdent := equipmentunavailableadjustmentcolumntype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.EquipmentUnavailableAdjustmentColumnType{Name: %s}).Stage(stage)", equipmentunavailableadjustmentcolumntypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmentcolumntype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", equipmentunavailableadjustmentcolumntypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmentcolumntype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", equipmentunavailableadjustmentcolumntypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmentcolumntype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", equipmentunavailableadjustmentcolumntypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmentcolumntype.Description)))
			for _, elem := range equipmentunavailableadjustmentcolumntype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", equipmentunavailableadjustmentcolumntypeIdent, equipmentunavailableadjustmentcolumntypeIdent, targetIdent))
			}
			for _, elem := range equipmentunavailableadjustmentcolumntype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", equipmentunavailableadjustmentcolumntypeIdent, equipmentunavailableadjustmentcolumntypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, equipmentunavailableadjustmentpropertytype := range __gong__sortStageSetInstances(stage.EquipmentUnavailableAdjustmentPropertyTypes, stage.EquipmentUnavailableAdjustmentPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			equipmentunavailableadjustmentpropertytypeIdent := equipmentunavailableadjustmentpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.EquipmentUnavailableAdjustmentPropertyType{Name: %s}).Stage(stage)", equipmentunavailableadjustmentpropertytypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmentpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", equipmentunavailableadjustmentpropertytypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", equipmentunavailableadjustmentpropertytypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmentpropertytype.NilReason)))
			if equipmentunavailableadjustmentpropertytype.EquipmentUnavailableAdjustment != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := equipmentunavailableadjustmentpropertytype.EquipmentUnavailableAdjustment.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EquipmentUnavailableAdjustment = %s", equipmentunavailableadjustmentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, equipmentunavailableadjustmenttype := range __gong__sortStageSetInstances(stage.EquipmentUnavailableAdjustmentTypes, stage.EquipmentUnavailableAdjustmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			equipmentunavailableadjustmenttypeIdent := equipmentunavailableadjustmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.EquipmentUnavailableAdjustmentType{Name: %s}).Stage(stage)", equipmentunavailableadjustmenttypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", equipmentunavailableadjustmenttypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", equipmentunavailableadjustmenttypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", equipmentunavailableadjustmenttypeIdent, __gong__toRawStringLiteral(equipmentunavailableadjustmenttype.Description)))
			for _, elem := range equipmentunavailableadjustmenttype.AdjustmentINOPCol {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AdjustmentINOPCol = append(%s.AdjustmentINOPCol, %s)", equipmentunavailableadjustmenttypeIdent, equipmentunavailableadjustmenttypeIdent, targetIdent))
			}
			for _, elem := range equipmentunavailableadjustmenttype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", equipmentunavailableadjustmenttypeIdent, equipmentunavailableadjustmenttypeIdent, targetIdent))
			}
			for _, elem := range equipmentunavailableadjustmenttype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", equipmentunavailableadjustmenttypeIdent, equipmentunavailableadjustmenttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, fasdatablockpropertytype := range __gong__sortStageSetInstances(stage.FASDataBlockPropertyTypes, stage.FASDataBlockPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fasdatablockpropertytypeIdent := fasdatablockpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FASDataBlockPropertyType{Name: %s}).Stage(stage)", fasdatablockpropertytypeIdent, __gong__toRawStringLiteral(fasdatablockpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fasdatablockpropertytypeIdent, __gong__toRawStringLiteral(fasdatablockpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", fasdatablockpropertytypeIdent, __gong__toRawStringLiteral(fasdatablockpropertytype.NilReason)))
			if fasdatablockpropertytype.FASDataBlock != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := fasdatablockpropertytype.FASDataBlock.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FASDataBlock = %s", fasdatablockpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, fasdatablocktype := range __gong__sortStageSetInstances(stage.FASDataBlockTypes, stage.FASDataBlockType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fasdatablocktypeIdent := fasdatablocktype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FASDataBlockType{Name: %s}).Stage(stage)", fasdatablocktypeIdent, __gong__toRawStringLiteral(fasdatablocktype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fasdatablocktypeIdent, __gong__toRawStringLiteral(fasdatablocktype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", fasdatablocktypeIdent, __gong__toRawStringLiteral(fasdatablocktype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", fasdatablocktypeIdent, __gong__toRawStringLiteral(fasdatablocktype.Description)))
			for _, elem := range fasdatablocktype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", fasdatablocktypeIdent, fasdatablocktypeIdent, targetIdent))
			}
			for _, elem := range fasdatablocktype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", fasdatablocktypeIdent, fasdatablocktypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, finallegtimeslicepropertytype := range __gong__sortStageSetInstances(stage.FinalLegTimeSlicePropertyTypes, stage.FinalLegTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			finallegtimeslicepropertytypeIdent := finallegtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FinalLegTimeSlicePropertyType{Name: %s}).Stage(stage)", finallegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(finallegtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", finallegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(finallegtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", finallegtimeslicepropertytypeIdent, finallegtimeslicepropertytype.Owns))
			if finallegtimeslicepropertytype.FinalLegTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicepropertytype.FinalLegTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FinalLegTimeSlice = %s", finallegtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, finallegtimeslicetype := range __gong__sortStageSetInstances(stage.FinalLegTimeSliceTypes, stage.FinalLegTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			finallegtimeslicetypeIdent := finallegtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FinalLegTimeSliceType{Name: %s}).Stage(stage)", finallegtimeslicetypeIdent, __gong__toRawStringLiteral(finallegtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", finallegtimeslicetypeIdent, __gong__toRawStringLiteral(finallegtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", finallegtimeslicetypeIdent, __gong__toRawStringLiteral(finallegtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", finallegtimeslicetypeIdent, __gong__toRawStringLiteral(finallegtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", finallegtimeslicetypeIdent, __gong__toRawStringLiteral(finallegtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", finallegtimeslicetypeIdent, finallegtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", finallegtimeslicetypeIdent, finallegtimeslicetype.CorrectionNumber))
			if finallegtimeslicetype.StartPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.StartPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartPoint = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.EndPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.EndPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPoint = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.Trajectory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.Trajectory.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Trajectory = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.ArcCentre != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.ArcCentre.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcCentre = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.Angle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.Angle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.Distance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.Distance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range finallegtimeslicetype.AircraftCategory {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = append(%s.AircraftCategory, %s)", finallegtimeslicetypeIdent, finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.Holding != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.Holding.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Holding = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range finallegtimeslicetype.DesignSurface {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignSurface = append(%s.DesignSurface, %s)", finallegtimeslicetypeIdent, finallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range finallegtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", finallegtimeslicetypeIdent, finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.Approach != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.Approach.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Approach = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range finallegtimeslicetype.Condition {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Condition = append(%s.Condition, %s)", finallegtimeslicetypeIdent, finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.FinalPathAlignmentPoint_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.FinalPathAlignmentPoint_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FinalPathAlignmentPoint_fixDesignatedPoint = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.FinalPathAlignmentPoint_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.FinalPathAlignmentPoint_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FinalPathAlignmentPoint_navaidSystem = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.FinalPathAlignmentPoint_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.FinalPathAlignmentPoint_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FinalPathAlignmentPoint_position = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.FinalPathAlignmentPoint_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.FinalPathAlignmentPoint_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FinalPathAlignmentPoint_runwayPoint = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.FinalPathAlignmentPoint_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.FinalPathAlignmentPoint_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FinalPathAlignmentPoint_aimingPoint = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.FinalPathAlignmentPoint_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.FinalPathAlignmentPoint_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FinalPathAlignmentPoint_airportReferencePoint = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.VisualDescentPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.VisualDescentPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VisualDescentPoint = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			if finallegtimeslicetype.FASData != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finallegtimeslicetype.FASData.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FASData = %s", finallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range finallegtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", finallegtimeslicetypeIdent, finallegtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, finallegtype := range __gong__sortStageSetInstances(stage.FinalLegTypes, stage.FinalLegType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			finallegtypeIdent := finallegtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FinalLegType{Name: %s}).Stage(stage)", finallegtypeIdent, __gong__toRawStringLiteral(finallegtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", finallegtypeIdent, __gong__toRawStringLiteral(finallegtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", finallegtypeIdent, __gong__toRawStringLiteral(finallegtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", finallegtypeIdent, __gong__toRawStringLiteral(finallegtype.Description)))
			for _, elem := range finallegtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", finallegtypeIdent, finallegtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, finalprofilepropertytype := range __gong__sortStageSetInstances(stage.FinalProfilePropertyTypes, stage.FinalProfilePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			finalprofilepropertytypeIdent := finalprofilepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FinalProfilePropertyType{Name: %s}).Stage(stage)", finalprofilepropertytypeIdent, __gong__toRawStringLiteral(finalprofilepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", finalprofilepropertytypeIdent, __gong__toRawStringLiteral(finalprofilepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", finalprofilepropertytypeIdent, __gong__toRawStringLiteral(finalprofilepropertytype.NilReason)))
			if finalprofilepropertytype.FinalProfile != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := finalprofilepropertytype.FinalProfile.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FinalProfile = %s", finalprofilepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, finalprofiletype := range __gong__sortStageSetInstances(stage.FinalProfileTypes, stage.FinalProfileType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			finalprofiletypeIdent := finalprofiletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FinalProfileType{Name: %s}).Stage(stage)", finalprofiletypeIdent, __gong__toRawStringLiteral(finalprofiletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", finalprofiletypeIdent, __gong__toRawStringLiteral(finalprofiletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", finalprofiletypeIdent, __gong__toRawStringLiteral(finalprofiletype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", finalprofiletypeIdent, __gong__toRawStringLiteral(finalprofiletype.Description)))
			for _, elem := range finalprofiletype.Altitude {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Altitude = append(%s.Altitude, %s)", finalprofiletypeIdent, finalprofiletypeIdent, targetIdent))
			}
			for _, elem := range finalprofiletype.Distance {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = append(%s.Distance, %s)", finalprofiletypeIdent, finalprofiletypeIdent, targetIdent))
			}
			for _, elem := range finalprofiletype.Timing {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Timing = append(%s.Timing, %s)", finalprofiletypeIdent, finalprofiletypeIdent, targetIdent))
			}
			for _, elem := range finalprofiletype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", finalprofiletypeIdent, finalprofiletypeIdent, targetIdent))
			}
			for _, elem := range finalprofiletype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", finalprofiletypeIdent, finalprofiletypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingassessmenttimeslicepropertytype := range __gong__sortStageSetInstances(stage.HoldingAssessmentTimeSlicePropertyTypes, stage.HoldingAssessmentTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingassessmenttimeslicepropertytypeIdent := holdingassessmenttimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingAssessmentTimeSlicePropertyType{Name: %s}).Stage(stage)", holdingassessmenttimeslicepropertytypeIdent, __gong__toRawStringLiteral(holdingassessmenttimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingassessmenttimeslicepropertytypeIdent, __gong__toRawStringLiteral(holdingassessmenttimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", holdingassessmenttimeslicepropertytypeIdent, holdingassessmenttimeslicepropertytype.Owns))
			if holdingassessmenttimeslicepropertytype.HoldingAssessmentTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingassessmenttimeslicepropertytype.HoldingAssessmentTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoldingAssessmentTimeSlice = %s", holdingassessmenttimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingassessmenttimeslicetype := range __gong__sortStageSetInstances(stage.HoldingAssessmentTimeSliceTypes, stage.HoldingAssessmentTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingassessmenttimeslicetypeIdent := holdingassessmenttimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingAssessmentTimeSliceType{Name: %s}).Stage(stage)", holdingassessmenttimeslicetypeIdent, __gong__toRawStringLiteral(holdingassessmenttimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingassessmenttimeslicetypeIdent, __gong__toRawStringLiteral(holdingassessmenttimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", holdingassessmenttimeslicetypeIdent, __gong__toRawStringLiteral(holdingassessmenttimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", holdingassessmenttimeslicetypeIdent, __gong__toRawStringLiteral(holdingassessmenttimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", holdingassessmenttimeslicetypeIdent, __gong__toRawStringLiteral(holdingassessmenttimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", holdingassessmenttimeslicetypeIdent, holdingassessmenttimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", holdingassessmenttimeslicetypeIdent, holdingassessmenttimeslicetype.CorrectionNumber))
			if holdingassessmenttimeslicetype.HoldingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingassessmenttimeslicetype.HoldingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoldingPoint = %s", holdingassessmenttimeslicetypeIdent, targetIdent))
			}
			if holdingassessmenttimeslicetype.UnplannedHolding != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingassessmenttimeslicetype.UnplannedHolding.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UnplannedHolding = %s", holdingassessmenttimeslicetypeIdent, targetIdent))
			}
			if holdingassessmenttimeslicetype.AssessedHoldingPattern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingassessmenttimeslicetype.AssessedHoldingPattern.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AssessedHoldingPattern = %s", holdingassessmenttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range holdingassessmenttimeslicetype.ObstacleAssessment {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstacleAssessment = append(%s.ObstacleAssessment, %s)", holdingassessmenttimeslicetypeIdent, holdingassessmenttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range holdingassessmenttimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", holdingassessmenttimeslicetypeIdent, holdingassessmenttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range holdingassessmenttimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", holdingassessmenttimeslicetypeIdent, holdingassessmenttimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingassessmenttype := range __gong__sortStageSetInstances(stage.HoldingAssessmentTypes, stage.HoldingAssessmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingassessmenttypeIdent := holdingassessmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingAssessmentType{Name: %s}).Stage(stage)", holdingassessmenttypeIdent, __gong__toRawStringLiteral(holdingassessmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingassessmenttypeIdent, __gong__toRawStringLiteral(holdingassessmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", holdingassessmenttypeIdent, __gong__toRawStringLiteral(holdingassessmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", holdingassessmenttypeIdent, __gong__toRawStringLiteral(holdingassessmenttype.Description)))
			for _, elem := range holdingassessmenttype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", holdingassessmenttypeIdent, holdingassessmenttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingpatterndistancepropertytype := range __gong__sortStageSetInstances(stage.HoldingPatternDistancePropertyTypes, stage.HoldingPatternDistancePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingpatterndistancepropertytypeIdent := holdingpatterndistancepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingPatternDistancePropertyType{Name: %s}).Stage(stage)", holdingpatterndistancepropertytypeIdent, __gong__toRawStringLiteral(holdingpatterndistancepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingpatterndistancepropertytypeIdent, __gong__toRawStringLiteral(holdingpatterndistancepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", holdingpatterndistancepropertytypeIdent, __gong__toRawStringLiteral(holdingpatterndistancepropertytype.NilReason)))
			if holdingpatterndistancepropertytype.HoldingPatternDistance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingpatterndistancepropertytype.HoldingPatternDistance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoldingPatternDistance = %s", holdingpatterndistancepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingpatterndistancetype := range __gong__sortStageSetInstances(stage.HoldingPatternDistanceTypes, stage.HoldingPatternDistanceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingpatterndistancetypeIdent := holdingpatterndistancetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingPatternDistanceType{Name: %s}).Stage(stage)", holdingpatterndistancetypeIdent, __gong__toRawStringLiteral(holdingpatterndistancetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingpatterndistancetypeIdent, __gong__toRawStringLiteral(holdingpatterndistancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", holdingpatterndistancetypeIdent, __gong__toRawStringLiteral(holdingpatterndistancetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", holdingpatterndistancetypeIdent, __gong__toRawStringLiteral(holdingpatterndistancetype.Description)))
			for _, elem := range holdingpatterndistancetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", holdingpatterndistancetypeIdent, holdingpatterndistancetypeIdent, targetIdent))
			}
			for _, elem := range holdingpatterndistancetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", holdingpatterndistancetypeIdent, holdingpatterndistancetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingpatterndurationpropertytype := range __gong__sortStageSetInstances(stage.HoldingPatternDurationPropertyTypes, stage.HoldingPatternDurationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingpatterndurationpropertytypeIdent := holdingpatterndurationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingPatternDurationPropertyType{Name: %s}).Stage(stage)", holdingpatterndurationpropertytypeIdent, __gong__toRawStringLiteral(holdingpatterndurationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingpatterndurationpropertytypeIdent, __gong__toRawStringLiteral(holdingpatterndurationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", holdingpatterndurationpropertytypeIdent, __gong__toRawStringLiteral(holdingpatterndurationpropertytype.NilReason)))
			if holdingpatterndurationpropertytype.HoldingPatternDuration != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingpatterndurationpropertytype.HoldingPatternDuration.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoldingPatternDuration = %s", holdingpatterndurationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingpatterndurationtype := range __gong__sortStageSetInstances(stage.HoldingPatternDurationTypes, stage.HoldingPatternDurationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingpatterndurationtypeIdent := holdingpatterndurationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingPatternDurationType{Name: %s}).Stage(stage)", holdingpatterndurationtypeIdent, __gong__toRawStringLiteral(holdingpatterndurationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingpatterndurationtypeIdent, __gong__toRawStringLiteral(holdingpatterndurationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", holdingpatterndurationtypeIdent, __gong__toRawStringLiteral(holdingpatterndurationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", holdingpatterndurationtypeIdent, __gong__toRawStringLiteral(holdingpatterndurationtype.Description)))
			for _, elem := range holdingpatterndurationtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", holdingpatterndurationtypeIdent, holdingpatterndurationtypeIdent, targetIdent))
			}
			for _, elem := range holdingpatterndurationtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", holdingpatterndurationtypeIdent, holdingpatterndurationtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingpatternpropertytype := range __gong__sortStageSetInstances(stage.HoldingPatternPropertyTypes, stage.HoldingPatternPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingpatternpropertytypeIdent := holdingpatternpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingPatternPropertyType{Name: %s}).Stage(stage)", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", holdingpatternpropertytypeIdent, holdingpatternpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, holdingpatterntimeslicepropertytype := range __gong__sortStageSetInstances(stage.HoldingPatternTimeSlicePropertyTypes, stage.HoldingPatternTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingpatterntimeslicepropertytypeIdent := holdingpatterntimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingPatternTimeSlicePropertyType{Name: %s}).Stage(stage)", holdingpatterntimeslicepropertytypeIdent, __gong__toRawStringLiteral(holdingpatterntimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingpatterntimeslicepropertytypeIdent, __gong__toRawStringLiteral(holdingpatterntimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", holdingpatterntimeslicepropertytypeIdent, holdingpatterntimeslicepropertytype.Owns))
			if holdingpatterntimeslicepropertytype.HoldingPatternTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingpatterntimeslicepropertytype.HoldingPatternTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoldingPatternTimeSlice = %s", holdingpatterntimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingpatterntimeslicetype := range __gong__sortStageSetInstances(stage.HoldingPatternTimeSliceTypes, stage.HoldingPatternTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingpatterntimeslicetypeIdent := holdingpatterntimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingPatternTimeSliceType{Name: %s}).Stage(stage)", holdingpatterntimeslicetypeIdent, __gong__toRawStringLiteral(holdingpatterntimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingpatterntimeslicetypeIdent, __gong__toRawStringLiteral(holdingpatterntimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", holdingpatterntimeslicetypeIdent, __gong__toRawStringLiteral(holdingpatterntimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", holdingpatterntimeslicetypeIdent, __gong__toRawStringLiteral(holdingpatterntimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", holdingpatterntimeslicetypeIdent, __gong__toRawStringLiteral(holdingpatterntimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", holdingpatterntimeslicetypeIdent, holdingpatterntimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", holdingpatterntimeslicetypeIdent, holdingpatterntimeslicetype.CorrectionNumber))
			if holdingpatterntimeslicetype.OutboundLegSpan_endPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingpatterntimeslicetype.OutboundLegSpan_endPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OutboundLegSpan_endPoint = %s", holdingpatterntimeslicetypeIdent, targetIdent))
			}
			if holdingpatterntimeslicetype.OutboundLegSpan_endDistance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingpatterntimeslicetype.OutboundLegSpan_endDistance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OutboundLegSpan_endDistance = %s", holdingpatterntimeslicetypeIdent, targetIdent))
			}
			if holdingpatterntimeslicetype.OutboundLegSpan_endTime != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingpatterntimeslicetype.OutboundLegSpan_endTime.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OutboundLegSpan_endTime = %s", holdingpatterntimeslicetypeIdent, targetIdent))
			}
			if holdingpatterntimeslicetype.HoldingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingpatterntimeslicetype.HoldingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoldingPoint = %s", holdingpatterntimeslicetypeIdent, targetIdent))
			}
			if holdingpatterntimeslicetype.Extent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingpatterntimeslicetype.Extent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extent = %s", holdingpatterntimeslicetypeIdent, targetIdent))
			}
			for _, elem := range holdingpatterntimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", holdingpatterntimeslicetypeIdent, holdingpatterntimeslicetypeIdent, targetIdent))
			}
			for _, elem := range holdingpatterntimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", holdingpatterntimeslicetypeIdent, holdingpatterntimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingpatterntype := range __gong__sortStageSetInstances(stage.HoldingPatternTypes, stage.HoldingPatternType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingpatterntypeIdent := holdingpatterntype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingPatternType{Name: %s}).Stage(stage)", holdingpatterntypeIdent, __gong__toRawStringLiteral(holdingpatterntype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingpatterntypeIdent, __gong__toRawStringLiteral(holdingpatterntype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", holdingpatterntypeIdent, __gong__toRawStringLiteral(holdingpatterntype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", holdingpatterntypeIdent, __gong__toRawStringLiteral(holdingpatterntype.Description)))
			for _, elem := range holdingpatterntype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", holdingpatterntypeIdent, holdingpatterntypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingusepropertytype := range __gong__sortStageSetInstances(stage.HoldingUsePropertyTypes, stage.HoldingUsePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingusepropertytypeIdent := holdingusepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingUsePropertyType{Name: %s}).Stage(stage)", holdingusepropertytypeIdent, __gong__toRawStringLiteral(holdingusepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingusepropertytypeIdent, __gong__toRawStringLiteral(holdingusepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", holdingusepropertytypeIdent, __gong__toRawStringLiteral(holdingusepropertytype.NilReason)))
			if holdingusepropertytype.HoldingUse != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingusepropertytype.HoldingUse.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoldingUse = %s", holdingusepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, holdingusetype := range __gong__sortStageSetInstances(stage.HoldingUseTypes, stage.HoldingUseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holdingusetypeIdent := holdingusetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingUseType{Name: %s}).Stage(stage)", holdingusetypeIdent, __gong__toRawStringLiteral(holdingusetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingusetypeIdent, __gong__toRawStringLiteral(holdingusetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", holdingusetypeIdent, __gong__toRawStringLiteral(holdingusetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", holdingusetypeIdent, __gong__toRawStringLiteral(holdingusetype.Description)))
			for _, elem := range holdingusetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", holdingusetypeIdent, holdingusetypeIdent, targetIdent))
			}
			if holdingusetype.TheHoldingPattern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := holdingusetype.TheHoldingPattern.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheHoldingPattern = %s", holdingusetypeIdent, targetIdent))
			}
			for _, elem := range holdingusetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", holdingusetypeIdent, holdingusetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, initiallegtimeslicepropertytype := range __gong__sortStageSetInstances(stage.InitialLegTimeSlicePropertyTypes, stage.InitialLegTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			initiallegtimeslicepropertytypeIdent := initiallegtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InitialLegTimeSlicePropertyType{Name: %s}).Stage(stage)", initiallegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(initiallegtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", initiallegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(initiallegtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", initiallegtimeslicepropertytypeIdent, initiallegtimeslicepropertytype.Owns))
			if initiallegtimeslicepropertytype.InitialLegTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := initiallegtimeslicepropertytype.InitialLegTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InitialLegTimeSlice = %s", initiallegtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, initiallegtimeslicetype := range __gong__sortStageSetInstances(stage.InitialLegTimeSliceTypes, stage.InitialLegTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			initiallegtimeslicetypeIdent := initiallegtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InitialLegTimeSliceType{Name: %s}).Stage(stage)", initiallegtimeslicetypeIdent, __gong__toRawStringLiteral(initiallegtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", initiallegtimeslicetypeIdent, __gong__toRawStringLiteral(initiallegtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", initiallegtimeslicetypeIdent, __gong__toRawStringLiteral(initiallegtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", initiallegtimeslicetypeIdent, __gong__toRawStringLiteral(initiallegtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", initiallegtimeslicetypeIdent, __gong__toRawStringLiteral(initiallegtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", initiallegtimeslicetypeIdent, initiallegtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", initiallegtimeslicetypeIdent, initiallegtimeslicetype.CorrectionNumber))
			if initiallegtimeslicetype.StartPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := initiallegtimeslicetype.StartPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartPoint = %s", initiallegtimeslicetypeIdent, targetIdent))
			}
			if initiallegtimeslicetype.EndPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := initiallegtimeslicetype.EndPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPoint = %s", initiallegtimeslicetypeIdent, targetIdent))
			}
			if initiallegtimeslicetype.Trajectory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := initiallegtimeslicetype.Trajectory.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Trajectory = %s", initiallegtimeslicetypeIdent, targetIdent))
			}
			if initiallegtimeslicetype.ArcCentre != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := initiallegtimeslicetype.ArcCentre.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcCentre = %s", initiallegtimeslicetypeIdent, targetIdent))
			}
			if initiallegtimeslicetype.Angle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := initiallegtimeslicetype.Angle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle = %s", initiallegtimeslicetypeIdent, targetIdent))
			}
			if initiallegtimeslicetype.Distance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := initiallegtimeslicetype.Distance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", initiallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range initiallegtimeslicetype.AircraftCategory {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = append(%s.AircraftCategory, %s)", initiallegtimeslicetypeIdent, initiallegtimeslicetypeIdent, targetIdent))
			}
			if initiallegtimeslicetype.Holding != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := initiallegtimeslicetype.Holding.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Holding = %s", initiallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range initiallegtimeslicetype.DesignSurface {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignSurface = append(%s.DesignSurface, %s)", initiallegtimeslicetypeIdent, initiallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range initiallegtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", initiallegtimeslicetypeIdent, initiallegtimeslicetypeIdent, targetIdent))
			}
			if initiallegtimeslicetype.Approach != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := initiallegtimeslicetype.Approach.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Approach = %s", initiallegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range initiallegtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", initiallegtimeslicetypeIdent, initiallegtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, initiallegtype := range __gong__sortStageSetInstances(stage.InitialLegTypes, stage.InitialLegType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			initiallegtypeIdent := initiallegtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InitialLegType{Name: %s}).Stage(stage)", initiallegtypeIdent, __gong__toRawStringLiteral(initiallegtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", initiallegtypeIdent, __gong__toRawStringLiteral(initiallegtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", initiallegtypeIdent, __gong__toRawStringLiteral(initiallegtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", initiallegtypeIdent, __gong__toRawStringLiteral(initiallegtype.Description)))
			for _, elem := range initiallegtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", initiallegtypeIdent, initiallegtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, instrumentapproachprocedurepropertytype := range __gong__sortStageSetInstances(stage.InstrumentApproachProcedurePropertyTypes, stage.InstrumentApproachProcedurePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			instrumentapproachprocedurepropertytypeIdent := instrumentapproachprocedurepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InstrumentApproachProcedurePropertyType{Name: %s}).Stage(stage)", instrumentapproachprocedurepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachprocedurepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentapproachprocedurepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachprocedurepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", instrumentapproachprocedurepropertytypeIdent, instrumentapproachprocedurepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", instrumentapproachprocedurepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachprocedurepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", instrumentapproachprocedurepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachprocedurepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", instrumentapproachprocedurepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachprocedurepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", instrumentapproachprocedurepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachprocedurepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", instrumentapproachprocedurepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachprocedurepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", instrumentapproachprocedurepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachprocedurepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", instrumentapproachprocedurepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachprocedurepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", instrumentapproachprocedurepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachprocedurepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, instrumentapproachproceduretimeslicepropertytype := range __gong__sortStageSetInstances(stage.InstrumentApproachProcedureTimeSlicePropertyTypes, stage.InstrumentApproachProcedureTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			instrumentapproachproceduretimeslicepropertytypeIdent := instrumentapproachproceduretimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InstrumentApproachProcedureTimeSlicePropertyType{Name: %s}).Stage(stage)", instrumentapproachproceduretimeslicepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentapproachproceduretimeslicepropertytypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", instrumentapproachproceduretimeslicepropertytypeIdent, instrumentapproachproceduretimeslicepropertytype.Owns))
			if instrumentapproachproceduretimeslicepropertytype.InstrumentApproachProcedureTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := instrumentapproachproceduretimeslicepropertytype.InstrumentApproachProcedureTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InstrumentApproachProcedureTimeSlice = %s", instrumentapproachproceduretimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, instrumentapproachproceduretimeslicetype := range __gong__sortStageSetInstances(stage.InstrumentApproachProcedureTimeSliceTypes, stage.InstrumentApproachProcedureTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			instrumentapproachproceduretimeslicetypeIdent := instrumentapproachproceduretimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InstrumentApproachProcedureTimeSliceType{Name: %s}).Stage(stage)", instrumentapproachproceduretimeslicetypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentapproachproceduretimeslicetypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrumentapproachproceduretimeslicetypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", instrumentapproachproceduretimeslicetypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", instrumentapproachproceduretimeslicetypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", instrumentapproachproceduretimeslicetypeIdent, instrumentapproachproceduretimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", instrumentapproachproceduretimeslicetypeIdent, instrumentapproachproceduretimeslicetype.CorrectionNumber))
			for _, elem := range instrumentapproachproceduretimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", instrumentapproachproceduretimeslicetypeIdent, instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range instrumentapproachproceduretimeslicetype.AirportHeliport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", instrumentapproachproceduretimeslicetypeIdent, instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range instrumentapproachproceduretimeslicetype.AircraftCharacteristic {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCharacteristic = append(%s.AircraftCharacteristic, %s)", instrumentapproachproceduretimeslicetypeIdent, instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range instrumentapproachproceduretimeslicetype.FlightTransition {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightTransition = append(%s.FlightTransition, %s)", instrumentapproachproceduretimeslicetypeIdent, instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			if instrumentapproachproceduretimeslicetype.GuidanceFacility_specialNavigationSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := instrumentapproachproceduretimeslicetype.GuidanceFacility_specialNavigationSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GuidanceFacility_specialNavigationSystem = %s", instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			if instrumentapproachproceduretimeslicetype.GuidanceFacility_radar != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := instrumentapproachproceduretimeslicetype.GuidanceFacility_radar.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GuidanceFacility_radar = %s", instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			if instrumentapproachproceduretimeslicetype.GuidanceFacility_navaid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := instrumentapproachproceduretimeslicetype.GuidanceFacility_navaid.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GuidanceFacility_navaid = %s", instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range instrumentapproachproceduretimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", instrumentapproachproceduretimeslicetypeIdent, instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			if instrumentapproachproceduretimeslicetype.SafeAltitude != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := instrumentapproachproceduretimeslicetype.SafeAltitude.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SafeAltitude = %s", instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			if instrumentapproachproceduretimeslicetype.Landing != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := instrumentapproachproceduretimeslicetype.Landing.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Landing = %s", instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range instrumentapproachproceduretimeslicetype.MissedInstruction {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MissedInstruction = append(%s.MissedInstruction, %s)", instrumentapproachproceduretimeslicetypeIdent, instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			if instrumentapproachproceduretimeslicetype.FinalProfile != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := instrumentapproachproceduretimeslicetype.FinalProfile.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FinalProfile = %s", instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range instrumentapproachproceduretimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", instrumentapproachproceduretimeslicetypeIdent, instrumentapproachproceduretimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, instrumentapproachproceduretype := range __gong__sortStageSetInstances(stage.InstrumentApproachProcedureTypes, stage.InstrumentApproachProcedureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			instrumentapproachproceduretypeIdent := instrumentapproachproceduretype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InstrumentApproachProcedureType{Name: %s}).Stage(stage)", instrumentapproachproceduretypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentapproachproceduretypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrumentapproachproceduretypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", instrumentapproachproceduretypeIdent, __gong__toRawStringLiteral(instrumentapproachproceduretype.Description)))
			for _, elem := range instrumentapproachproceduretype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", instrumentapproachproceduretypeIdent, instrumentapproachproceduretypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, intermediatelegtimeslicepropertytype := range __gong__sortStageSetInstances(stage.IntermediateLegTimeSlicePropertyTypes, stage.IntermediateLegTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			intermediatelegtimeslicepropertytypeIdent := intermediatelegtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.IntermediateLegTimeSlicePropertyType{Name: %s}).Stage(stage)", intermediatelegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(intermediatelegtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", intermediatelegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(intermediatelegtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", intermediatelegtimeslicepropertytypeIdent, intermediatelegtimeslicepropertytype.Owns))
			if intermediatelegtimeslicepropertytype.IntermediateLegTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := intermediatelegtimeslicepropertytype.IntermediateLegTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediateLegTimeSlice = %s", intermediatelegtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, intermediatelegtimeslicetype := range __gong__sortStageSetInstances(stage.IntermediateLegTimeSliceTypes, stage.IntermediateLegTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			intermediatelegtimeslicetypeIdent := intermediatelegtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.IntermediateLegTimeSliceType{Name: %s}).Stage(stage)", intermediatelegtimeslicetypeIdent, __gong__toRawStringLiteral(intermediatelegtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", intermediatelegtimeslicetypeIdent, __gong__toRawStringLiteral(intermediatelegtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", intermediatelegtimeslicetypeIdent, __gong__toRawStringLiteral(intermediatelegtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", intermediatelegtimeslicetypeIdent, __gong__toRawStringLiteral(intermediatelegtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", intermediatelegtimeslicetypeIdent, __gong__toRawStringLiteral(intermediatelegtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", intermediatelegtimeslicetypeIdent, intermediatelegtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", intermediatelegtimeslicetypeIdent, intermediatelegtimeslicetype.CorrectionNumber))
			if intermediatelegtimeslicetype.StartPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := intermediatelegtimeslicetype.StartPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartPoint = %s", intermediatelegtimeslicetypeIdent, targetIdent))
			}
			if intermediatelegtimeslicetype.EndPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := intermediatelegtimeslicetype.EndPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPoint = %s", intermediatelegtimeslicetypeIdent, targetIdent))
			}
			if intermediatelegtimeslicetype.Trajectory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := intermediatelegtimeslicetype.Trajectory.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Trajectory = %s", intermediatelegtimeslicetypeIdent, targetIdent))
			}
			if intermediatelegtimeslicetype.ArcCentre != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := intermediatelegtimeslicetype.ArcCentre.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcCentre = %s", intermediatelegtimeslicetypeIdent, targetIdent))
			}
			if intermediatelegtimeslicetype.Angle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := intermediatelegtimeslicetype.Angle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle = %s", intermediatelegtimeslicetypeIdent, targetIdent))
			}
			if intermediatelegtimeslicetype.Distance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := intermediatelegtimeslicetype.Distance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", intermediatelegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range intermediatelegtimeslicetype.AircraftCategory {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = append(%s.AircraftCategory, %s)", intermediatelegtimeslicetypeIdent, intermediatelegtimeslicetypeIdent, targetIdent))
			}
			if intermediatelegtimeslicetype.Holding != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := intermediatelegtimeslicetype.Holding.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Holding = %s", intermediatelegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range intermediatelegtimeslicetype.DesignSurface {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignSurface = append(%s.DesignSurface, %s)", intermediatelegtimeslicetypeIdent, intermediatelegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range intermediatelegtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", intermediatelegtimeslicetypeIdent, intermediatelegtimeslicetypeIdent, targetIdent))
			}
			if intermediatelegtimeslicetype.Approach != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := intermediatelegtimeslicetype.Approach.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Approach = %s", intermediatelegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range intermediatelegtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", intermediatelegtimeslicetypeIdent, intermediatelegtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, intermediatelegtype := range __gong__sortStageSetInstances(stage.IntermediateLegTypes, stage.IntermediateLegType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			intermediatelegtypeIdent := intermediatelegtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.IntermediateLegType{Name: %s}).Stage(stage)", intermediatelegtypeIdent, __gong__toRawStringLiteral(intermediatelegtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", intermediatelegtypeIdent, __gong__toRawStringLiteral(intermediatelegtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", intermediatelegtypeIdent, __gong__toRawStringLiteral(intermediatelegtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", intermediatelegtypeIdent, __gong__toRawStringLiteral(intermediatelegtype.Description)))
			for _, elem := range intermediatelegtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", intermediatelegtypeIdent, intermediatelegtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, landingtakeoffareacollectionpropertytype := range __gong__sortStageSetInstances(stage.LandingTakeoffAreaCollectionPropertyTypes, stage.LandingTakeoffAreaCollectionPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			landingtakeoffareacollectionpropertytypeIdent := landingtakeoffareacollectionpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LandingTakeoffAreaCollectionPropertyType{Name: %s}).Stage(stage)", landingtakeoffareacollectionpropertytypeIdent, __gong__toRawStringLiteral(landingtakeoffareacollectionpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", landingtakeoffareacollectionpropertytypeIdent, __gong__toRawStringLiteral(landingtakeoffareacollectionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", landingtakeoffareacollectionpropertytypeIdent, __gong__toRawStringLiteral(landingtakeoffareacollectionpropertytype.NilReason)))
			if landingtakeoffareacollectionpropertytype.LandingTakeoffAreaCollection != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := landingtakeoffareacollectionpropertytype.LandingTakeoffAreaCollection.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LandingTakeoffAreaCollection = %s", landingtakeoffareacollectionpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, landingtakeoffareacollectiontype := range __gong__sortStageSetInstances(stage.LandingTakeoffAreaCollectionTypes, stage.LandingTakeoffAreaCollectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			landingtakeoffareacollectiontypeIdent := landingtakeoffareacollectiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LandingTakeoffAreaCollectionType{Name: %s}).Stage(stage)", landingtakeoffareacollectiontypeIdent, __gong__toRawStringLiteral(landingtakeoffareacollectiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", landingtakeoffareacollectiontypeIdent, __gong__toRawStringLiteral(landingtakeoffareacollectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", landingtakeoffareacollectiontypeIdent, __gong__toRawStringLiteral(landingtakeoffareacollectiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", landingtakeoffareacollectiontypeIdent, __gong__toRawStringLiteral(landingtakeoffareacollectiontype.Description)))
			for _, elem := range landingtakeoffareacollectiontype.Runway {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Runway = append(%s.Runway, %s)", landingtakeoffareacollectiontypeIdent, landingtakeoffareacollectiontypeIdent, targetIdent))
			}
			for _, elem := range landingtakeoffareacollectiontype.TLOF {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TLOF = append(%s.TLOF, %s)", landingtakeoffareacollectiontypeIdent, landingtakeoffareacollectiontypeIdent, targetIdent))
			}
			for _, elem := range landingtakeoffareacollectiontype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", landingtakeoffareacollectiontypeIdent, landingtakeoffareacollectiontypeIdent, targetIdent))
			}
			for _, elem := range landingtakeoffareacollectiontype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", landingtakeoffareacollectiontypeIdent, landingtakeoffareacollectiontypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, linguisticnotepropertytype := range __gong__sortStageSetInstances(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linguisticnotepropertytypeIdent := linguisticnotepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LinguisticNotePropertyType{Name: %s}).Stage(stage)", linguisticnotepropertytypeIdent, __gong__toRawStringLiteral(linguisticnotepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linguisticnotepropertytypeIdent, __gong__toRawStringLiteral(linguisticnotepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", linguisticnotepropertytypeIdent, __gong__toRawStringLiteral(linguisticnotepropertytype.NilReason)))
			if linguisticnotepropertytype.LinguisticNote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := linguisticnotepropertytype.LinguisticNote.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LinguisticNote = %s", linguisticnotepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, linguisticnotetype := range __gong__sortStageSetInstances(stage.LinguisticNoteTypes, stage.LinguisticNoteType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linguisticnotetypeIdent := linguisticnotetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LinguisticNoteType{Name: %s}).Stage(stage)", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Description)))
			for _, elem := range linguisticnotetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", linguisticnotetypeIdent, linguisticnotetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, minimapropertytype := range __gong__sortStageSetInstances(stage.MinimaPropertyTypes, stage.MinimaPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			minimapropertytypeIdent := minimapropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MinimaPropertyType{Name: %s}).Stage(stage)", minimapropertytypeIdent, __gong__toRawStringLiteral(minimapropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", minimapropertytypeIdent, __gong__toRawStringLiteral(minimapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", minimapropertytypeIdent, __gong__toRawStringLiteral(minimapropertytype.NilReason)))
			if minimapropertytype.Minima != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := minimapropertytype.Minima.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Minima = %s", minimapropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, minimatype := range __gong__sortStageSetInstances(stage.MinimaTypes, stage.MinimaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			minimatypeIdent := minimatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MinimaType{Name: %s}).Stage(stage)", minimatypeIdent, __gong__toRawStringLiteral(minimatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", minimatypeIdent, __gong__toRawStringLiteral(minimatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", minimatypeIdent, __gong__toRawStringLiteral(minimatype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", minimatypeIdent, __gong__toRawStringLiteral(minimatype.Description)))
			for _, elem := range minimatype.AdjustmentINOP {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AdjustmentINOP = append(%s.AdjustmentINOP, %s)", minimatypeIdent, minimatypeIdent, targetIdent))
			}
			for _, elem := range minimatype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", minimatypeIdent, minimatypeIdent, targetIdent))
			}
			for _, elem := range minimatype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", minimatypeIdent, minimatypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, missedapproachgrouppropertytype := range __gong__sortStageSetInstances(stage.MissedApproachGroupPropertyTypes, stage.MissedApproachGroupPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			missedapproachgrouppropertytypeIdent := missedapproachgrouppropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MissedApproachGroupPropertyType{Name: %s}).Stage(stage)", missedapproachgrouppropertytypeIdent, __gong__toRawStringLiteral(missedapproachgrouppropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", missedapproachgrouppropertytypeIdent, __gong__toRawStringLiteral(missedapproachgrouppropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", missedapproachgrouppropertytypeIdent, __gong__toRawStringLiteral(missedapproachgrouppropertytype.NilReason)))
			if missedapproachgrouppropertytype.MissedApproachGroup != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := missedapproachgrouppropertytype.MissedApproachGroup.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MissedApproachGroup = %s", missedapproachgrouppropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, missedapproachgrouptype := range __gong__sortStageSetInstances(stage.MissedApproachGroupTypes, stage.MissedApproachGroupType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			missedapproachgrouptypeIdent := missedapproachgrouptype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MissedApproachGroupType{Name: %s}).Stage(stage)", missedapproachgrouptypeIdent, __gong__toRawStringLiteral(missedapproachgrouptype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", missedapproachgrouptypeIdent, __gong__toRawStringLiteral(missedapproachgrouptype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", missedapproachgrouptypeIdent, __gong__toRawStringLiteral(missedapproachgrouptype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", missedapproachgrouptypeIdent, __gong__toRawStringLiteral(missedapproachgrouptype.Description)))
			for _, elem := range missedapproachgrouptype.Altimeter {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Altimeter = append(%s.Altimeter, %s)", missedapproachgrouptypeIdent, missedapproachgrouptypeIdent, targetIdent))
			}
			for _, elem := range missedapproachgrouptype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", missedapproachgrouptypeIdent, missedapproachgrouptypeIdent, targetIdent))
			}
			for _, elem := range missedapproachgrouptype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", missedapproachgrouptypeIdent, missedapproachgrouptypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, missedapproachlegtimeslicepropertytype := range __gong__sortStageSetInstances(stage.MissedApproachLegTimeSlicePropertyTypes, stage.MissedApproachLegTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			missedapproachlegtimeslicepropertytypeIdent := missedapproachlegtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MissedApproachLegTimeSlicePropertyType{Name: %s}).Stage(stage)", missedapproachlegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(missedapproachlegtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", missedapproachlegtimeslicepropertytypeIdent, __gong__toRawStringLiteral(missedapproachlegtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", missedapproachlegtimeslicepropertytypeIdent, missedapproachlegtimeslicepropertytype.Owns))
			if missedapproachlegtimeslicepropertytype.MissedApproachLegTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := missedapproachlegtimeslicepropertytype.MissedApproachLegTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MissedApproachLegTimeSlice = %s", missedapproachlegtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, missedapproachlegtimeslicetype := range __gong__sortStageSetInstances(stage.MissedApproachLegTimeSliceTypes, stage.MissedApproachLegTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			missedapproachlegtimeslicetypeIdent := missedapproachlegtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MissedApproachLegTimeSliceType{Name: %s}).Stage(stage)", missedapproachlegtimeslicetypeIdent, __gong__toRawStringLiteral(missedapproachlegtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", missedapproachlegtimeslicetypeIdent, __gong__toRawStringLiteral(missedapproachlegtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", missedapproachlegtimeslicetypeIdent, __gong__toRawStringLiteral(missedapproachlegtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", missedapproachlegtimeslicetypeIdent, __gong__toRawStringLiteral(missedapproachlegtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", missedapproachlegtimeslicetypeIdent, __gong__toRawStringLiteral(missedapproachlegtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", missedapproachlegtimeslicetypeIdent, missedapproachlegtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", missedapproachlegtimeslicetypeIdent, missedapproachlegtimeslicetype.CorrectionNumber))
			if missedapproachlegtimeslicetype.StartPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := missedapproachlegtimeslicetype.StartPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartPoint = %s", missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			if missedapproachlegtimeslicetype.EndPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := missedapproachlegtimeslicetype.EndPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPoint = %s", missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			if missedapproachlegtimeslicetype.Trajectory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := missedapproachlegtimeslicetype.Trajectory.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Trajectory = %s", missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			if missedapproachlegtimeslicetype.ArcCentre != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := missedapproachlegtimeslicetype.ArcCentre.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcCentre = %s", missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			if missedapproachlegtimeslicetype.Angle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := missedapproachlegtimeslicetype.Angle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle = %s", missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			if missedapproachlegtimeslicetype.Distance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := missedapproachlegtimeslicetype.Distance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range missedapproachlegtimeslicetype.AircraftCategory {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = append(%s.AircraftCategory, %s)", missedapproachlegtimeslicetypeIdent, missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			if missedapproachlegtimeslicetype.Holding != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := missedapproachlegtimeslicetype.Holding.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Holding = %s", missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range missedapproachlegtimeslicetype.DesignSurface {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignSurface = append(%s.DesignSurface, %s)", missedapproachlegtimeslicetypeIdent, missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range missedapproachlegtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", missedapproachlegtimeslicetypeIdent, missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			if missedapproachlegtimeslicetype.Approach != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := missedapproachlegtimeslicetype.Approach.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Approach = %s", missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range missedapproachlegtimeslicetype.Condition {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Condition = append(%s.Condition, %s)", missedapproachlegtimeslicetypeIdent, missedapproachlegtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range missedapproachlegtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", missedapproachlegtimeslicetypeIdent, missedapproachlegtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, missedapproachlegtype := range __gong__sortStageSetInstances(stage.MissedApproachLegTypes, stage.MissedApproachLegType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			missedapproachlegtypeIdent := missedapproachlegtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MissedApproachLegType{Name: %s}).Stage(stage)", missedapproachlegtypeIdent, __gong__toRawStringLiteral(missedapproachlegtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", missedapproachlegtypeIdent, __gong__toRawStringLiteral(missedapproachlegtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", missedapproachlegtypeIdent, __gong__toRawStringLiteral(missedapproachlegtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", missedapproachlegtypeIdent, __gong__toRawStringLiteral(missedapproachlegtype.Description)))
			for _, elem := range missedapproachlegtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", missedapproachlegtypeIdent, missedapproachlegtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navaidpropertytype := range __gong__sortStageSetInstances(stage.NavaidPropertyTypes, stage.NavaidPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidpropertytypeIdent := navaidpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidPropertyType{Name: %s}).Stage(stage)", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", navaidpropertytypeIdent, navaidpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, navigationarearestrictiontimeslicepropertytype := range __gong__sortStageSetInstances(stage.NavigationAreaRestrictionTimeSlicePropertyTypes, stage.NavigationAreaRestrictionTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navigationarearestrictiontimeslicepropertytypeIdent := navigationarearestrictiontimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavigationAreaRestrictionTimeSlicePropertyType{Name: %s}).Stage(stage)", navigationarearestrictiontimeslicepropertytypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navigationarearestrictiontimeslicepropertytypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", navigationarearestrictiontimeslicepropertytypeIdent, navigationarearestrictiontimeslicepropertytype.Owns))
			if navigationarearestrictiontimeslicepropertytype.NavigationAreaRestrictionTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationarearestrictiontimeslicepropertytype.NavigationAreaRestrictionTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavigationAreaRestrictionTimeSlice = %s", navigationarearestrictiontimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navigationarearestrictiontimeslicetype := range __gong__sortStageSetInstances(stage.NavigationAreaRestrictionTimeSliceTypes, stage.NavigationAreaRestrictionTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navigationarearestrictiontimeslicetypeIdent := navigationarearestrictiontimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavigationAreaRestrictionTimeSliceType{Name: %s}).Stage(stage)", navigationarearestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navigationarearestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navigationarearestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navigationarearestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", navigationarearestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", navigationarearestrictiontimeslicetypeIdent, navigationarearestrictiontimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", navigationarearestrictiontimeslicetypeIdent, navigationarearestrictiontimeslicetype.CorrectionNumber))
			for _, elem := range navigationarearestrictiontimeslicetype.Procedure {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Procedure = append(%s.Procedure, %s)", navigationarearestrictiontimeslicetypeIdent, navigationarearestrictiontimeslicetypeIdent, targetIdent))
			}
			if navigationarearestrictiontimeslicetype.DesignSurface != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationarearestrictiontimeslicetype.DesignSurface.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignSurface = %s", navigationarearestrictiontimeslicetypeIdent, targetIdent))
			}
			if navigationarearestrictiontimeslicetype.SectorDefinition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationarearestrictiontimeslicetype.SectorDefinition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SectorDefinition = %s", navigationarearestrictiontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navigationarearestrictiontimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navigationarearestrictiontimeslicetypeIdent, navigationarearestrictiontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navigationarearestrictiontimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navigationarearestrictiontimeslicetypeIdent, navigationarearestrictiontimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navigationarearestrictiontype := range __gong__sortStageSetInstances(stage.NavigationAreaRestrictionTypes, stage.NavigationAreaRestrictionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navigationarearestrictiontypeIdent := navigationarearestrictiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavigationAreaRestrictionType{Name: %s}).Stage(stage)", navigationarearestrictiontypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navigationarearestrictiontypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navigationarearestrictiontypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navigationarearestrictiontypeIdent, __gong__toRawStringLiteral(navigationarearestrictiontype.Description)))
			for _, elem := range navigationarearestrictiontype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", navigationarearestrictiontypeIdent, navigationarearestrictiontypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navigationareasectorpropertytype := range __gong__sortStageSetInstances(stage.NavigationAreaSectorPropertyTypes, stage.NavigationAreaSectorPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navigationareasectorpropertytypeIdent := navigationareasectorpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavigationAreaSectorPropertyType{Name: %s}).Stage(stage)", navigationareasectorpropertytypeIdent, __gong__toRawStringLiteral(navigationareasectorpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navigationareasectorpropertytypeIdent, __gong__toRawStringLiteral(navigationareasectorpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", navigationareasectorpropertytypeIdent, __gong__toRawStringLiteral(navigationareasectorpropertytype.NilReason)))
			if navigationareasectorpropertytype.NavigationAreaSector != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareasectorpropertytype.NavigationAreaSector.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavigationAreaSector = %s", navigationareasectorpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navigationareasectortype := range __gong__sortStageSetInstances(stage.NavigationAreaSectorTypes, stage.NavigationAreaSectorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navigationareasectortypeIdent := navigationareasectortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavigationAreaSectorType{Name: %s}).Stage(stage)", navigationareasectortypeIdent, __gong__toRawStringLiteral(navigationareasectortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navigationareasectortypeIdent, __gong__toRawStringLiteral(navigationareasectortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navigationareasectortypeIdent, __gong__toRawStringLiteral(navigationareasectortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navigationareasectortypeIdent, __gong__toRawStringLiteral(navigationareasectortype.Description)))
			if navigationareasectortype.SectorDefinition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareasectortype.SectorDefinition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SectorDefinition = %s", navigationareasectortypeIdent, targetIdent))
			}
			for _, elem := range navigationareasectortype.SignificantObstacle {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantObstacle = append(%s.SignificantObstacle, %s)", navigationareasectortypeIdent, navigationareasectortypeIdent, targetIdent))
			}
			if navigationareasectortype.Extent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareasectortype.Extent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extent = %s", navigationareasectortypeIdent, targetIdent))
			}
			for _, elem := range navigationareasectortype.SectorCriteria {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SectorCriteria = append(%s.SectorCriteria, %s)", navigationareasectortypeIdent, navigationareasectortypeIdent, targetIdent))
			}
			for _, elem := range navigationareasectortype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navigationareasectortypeIdent, navigationareasectortypeIdent, targetIdent))
			}
			for _, elem := range navigationareasectortype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navigationareasectortypeIdent, navigationareasectortypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navigationareatimeslicepropertytype := range __gong__sortStageSetInstances(stage.NavigationAreaTimeSlicePropertyTypes, stage.NavigationAreaTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navigationareatimeslicepropertytypeIdent := navigationareatimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavigationAreaTimeSlicePropertyType{Name: %s}).Stage(stage)", navigationareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(navigationareatimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navigationareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(navigationareatimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", navigationareatimeslicepropertytypeIdent, navigationareatimeslicepropertytype.Owns))
			if navigationareatimeslicepropertytype.NavigationAreaTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareatimeslicepropertytype.NavigationAreaTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavigationAreaTimeSlice = %s", navigationareatimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navigationareatimeslicetype := range __gong__sortStageSetInstances(stage.NavigationAreaTimeSliceTypes, stage.NavigationAreaTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navigationareatimeslicetypeIdent := navigationareatimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavigationAreaTimeSliceType{Name: %s}).Stage(stage)", navigationareatimeslicetypeIdent, __gong__toRawStringLiteral(navigationareatimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navigationareatimeslicetypeIdent, __gong__toRawStringLiteral(navigationareatimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navigationareatimeslicetypeIdent, __gong__toRawStringLiteral(navigationareatimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navigationareatimeslicetypeIdent, __gong__toRawStringLiteral(navigationareatimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", navigationareatimeslicetypeIdent, __gong__toRawStringLiteral(navigationareatimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", navigationareatimeslicetypeIdent, navigationareatimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", navigationareatimeslicetypeIdent, navigationareatimeslicetype.CorrectionNumber))
			if navigationareatimeslicetype.Departure != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareatimeslicetype.Departure.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Departure = %s", navigationareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navigationareatimeslicetype.Sector {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sector = append(%s.Sector, %s)", navigationareatimeslicetypeIdent, navigationareatimeslicetypeIdent, targetIdent))
			}
			if navigationareatimeslicetype.CentrePoint_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareatimeslicetype.CentrePoint_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_fixDesignatedPoint = %s", navigationareatimeslicetypeIdent, targetIdent))
			}
			if navigationareatimeslicetype.CentrePoint_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareatimeslicetype.CentrePoint_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_navaidSystem = %s", navigationareatimeslicetypeIdent, targetIdent))
			}
			if navigationareatimeslicetype.CentrePoint_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareatimeslicetype.CentrePoint_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_position = %s", navigationareatimeslicetypeIdent, targetIdent))
			}
			if navigationareatimeslicetype.CentrePoint_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareatimeslicetype.CentrePoint_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_runwayPoint = %s", navigationareatimeslicetypeIdent, targetIdent))
			}
			if navigationareatimeslicetype.CentrePoint_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareatimeslicetype.CentrePoint_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_aimingPoint = %s", navigationareatimeslicetypeIdent, targetIdent))
			}
			if navigationareatimeslicetype.CentrePoint_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navigationareatimeslicetype.CentrePoint_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_airportReferencePoint = %s", navigationareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navigationareatimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navigationareatimeslicetypeIdent, navigationareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navigationareatimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navigationareatimeslicetypeIdent, navigationareatimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navigationareatype := range __gong__sortStageSetInstances(stage.NavigationAreaTypes, stage.NavigationAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navigationareatypeIdent := navigationareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavigationAreaType{Name: %s}).Stage(stage)", navigationareatypeIdent, __gong__toRawStringLiteral(navigationareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navigationareatypeIdent, __gong__toRawStringLiteral(navigationareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navigationareatypeIdent, __gong__toRawStringLiteral(navigationareatype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navigationareatypeIdent, __gong__toRawStringLiteral(navigationareatype.Description)))
			for _, elem := range navigationareatype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", navigationareatypeIdent, navigationareatypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, notepropertytype := range __gong__sortStageSetInstances(stage.NotePropertyTypes, stage.NotePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notepropertytypeIdent := notepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NotePropertyType{Name: %s}).Stage(stage)", notepropertytypeIdent, __gong__toRawStringLiteral(notepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notepropertytypeIdent, __gong__toRawStringLiteral(notepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", notepropertytypeIdent, __gong__toRawStringLiteral(notepropertytype.NilReason)))
			if notepropertytype.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notepropertytype.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, notetype := range __gong__sortStageSetInstances(stage.NoteTypes, stage.NoteType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notetypeIdent := notetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NoteType{Name: %s}).Stage(stage)", notetypeIdent, __gong__toRawStringLiteral(notetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Description)))
			for _, elem := range notetype.TranslatedNote {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TranslatedNote = append(%s.TranslatedNote, %s)", notetypeIdent, notetypeIdent, targetIdent))
			}
			for _, elem := range notetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", notetypeIdent, notetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, obstacleassessmentareapropertytype := range __gong__sortStageSetInstances(stage.ObstacleAssessmentAreaPropertyTypes, stage.ObstacleAssessmentAreaPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			obstacleassessmentareapropertytypeIdent := obstacleassessmentareapropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstacleAssessmentAreaPropertyType{Name: %s}).Stage(stage)", obstacleassessmentareapropertytypeIdent, __gong__toRawStringLiteral(obstacleassessmentareapropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleassessmentareapropertytypeIdent, __gong__toRawStringLiteral(obstacleassessmentareapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", obstacleassessmentareapropertytypeIdent, __gong__toRawStringLiteral(obstacleassessmentareapropertytype.NilReason)))
			if obstacleassessmentareapropertytype.ObstacleAssessmentArea != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstacleassessmentareapropertytype.ObstacleAssessmentArea.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstacleAssessmentArea = %s", obstacleassessmentareapropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, obstacleassessmentareatype := range __gong__sortStageSetInstances(stage.ObstacleAssessmentAreaTypes, stage.ObstacleAssessmentAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			obstacleassessmentareatypeIdent := obstacleassessmentareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstacleAssessmentAreaType{Name: %s}).Stage(stage)", obstacleassessmentareatypeIdent, __gong__toRawStringLiteral(obstacleassessmentareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleassessmentareatypeIdent, __gong__toRawStringLiteral(obstacleassessmentareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", obstacleassessmentareatypeIdent, __gong__toRawStringLiteral(obstacleassessmentareatype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", obstacleassessmentareatypeIdent, __gong__toRawStringLiteral(obstacleassessmentareatype.Description)))
			for _, elem := range obstacleassessmentareatype.AircraftCategory {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = append(%s.AircraftCategory, %s)", obstacleassessmentareatypeIdent, obstacleassessmentareatypeIdent, targetIdent))
			}
			for _, elem := range obstacleassessmentareatype.SignificantObstacle {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantObstacle = append(%s.SignificantObstacle, %s)", obstacleassessmentareatypeIdent, obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.Surface != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstacleassessmentareatype.Surface.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Surface = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.StartingCurve != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstacleassessmentareatype.StartingCurve.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartingCurve = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			for _, elem := range obstacleassessmentareatype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", obstacleassessmentareatypeIdent, obstacleassessmentareatypeIdent, targetIdent))
			}
			for _, elem := range obstacleassessmentareatype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", obstacleassessmentareatypeIdent, obstacleassessmentareatypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, obstacleplacementpropertytype := range __gong__sortStageSetInstances(stage.ObstaclePlacementPropertyTypes, stage.ObstaclePlacementPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			obstacleplacementpropertytypeIdent := obstacleplacementpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstaclePlacementPropertyType{Name: %s}).Stage(stage)", obstacleplacementpropertytypeIdent, __gong__toRawStringLiteral(obstacleplacementpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleplacementpropertytypeIdent, __gong__toRawStringLiteral(obstacleplacementpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", obstacleplacementpropertytypeIdent, __gong__toRawStringLiteral(obstacleplacementpropertytype.NilReason)))
			if obstacleplacementpropertytype.ObstaclePlacement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstacleplacementpropertytype.ObstaclePlacement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstaclePlacement = %s", obstacleplacementpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, obstacleplacementtype := range __gong__sortStageSetInstances(stage.ObstaclePlacementTypes, stage.ObstaclePlacementType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			obstacleplacementtypeIdent := obstacleplacementtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstaclePlacementType{Name: %s}).Stage(stage)", obstacleplacementtypeIdent, __gong__toRawStringLiteral(obstacleplacementtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleplacementtypeIdent, __gong__toRawStringLiteral(obstacleplacementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", obstacleplacementtypeIdent, __gong__toRawStringLiteral(obstacleplacementtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", obstacleplacementtypeIdent, __gong__toRawStringLiteral(obstacleplacementtype.Description)))
			for _, elem := range obstacleplacementtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", obstacleplacementtypeIdent, obstacleplacementtypeIdent, targetIdent))
			}
			for _, elem := range obstacleplacementtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", obstacleplacementtypeIdent, obstacleplacementtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, obstructionpropertytype := range __gong__sortStageSetInstances(stage.ObstructionPropertyTypes, stage.ObstructionPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			obstructionpropertytypeIdent := obstructionpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstructionPropertyType{Name: %s}).Stage(stage)", obstructionpropertytypeIdent, __gong__toRawStringLiteral(obstructionpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstructionpropertytypeIdent, __gong__toRawStringLiteral(obstructionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", obstructionpropertytypeIdent, __gong__toRawStringLiteral(obstructionpropertytype.NilReason)))
			if obstructionpropertytype.Obstruction != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstructionpropertytype.Obstruction.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Obstruction = %s", obstructionpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, obstructiontype := range __gong__sortStageSetInstances(stage.ObstructionTypes, stage.ObstructionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			obstructiontypeIdent := obstructiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstructionType{Name: %s}).Stage(stage)", obstructiontypeIdent, __gong__toRawStringLiteral(obstructiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstructiontypeIdent, __gong__toRawStringLiteral(obstructiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", obstructiontypeIdent, __gong__toRawStringLiteral(obstructiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", obstructiontypeIdent, __gong__toRawStringLiteral(obstructiontype.Description)))
			if obstructiontype.TheVerticalStructure != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstructiontype.TheVerticalStructure.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheVerticalStructure = %s", obstructiontypeIdent, targetIdent))
			}
			for _, elem := range obstructiontype.Adjustment {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Adjustment = append(%s.Adjustment, %s)", obstructiontypeIdent, obstructiontypeIdent, targetIdent))
			}
			for _, elem := range obstructiontype.ObstaclePlacement {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstaclePlacement = append(%s.ObstaclePlacement, %s)", obstructiontypeIdent, obstructiontypeIdent, targetIdent))
			}
			for _, elem := range obstructiontype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", obstructiontypeIdent, obstructiontypeIdent, targetIdent))
			}
			for _, elem := range obstructiontype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", obstructiontypeIdent, obstructiontypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, organisationauthoritypropertytype := range __gong__sortStageSetInstances(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			organisationauthoritypropertytypeIdent := organisationauthoritypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityPropertyType{Name: %s}).Stage(stage)", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", organisationauthoritypropertytypeIdent, organisationauthoritypropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, pointpropertytype := range __gong__sortStageSetInstances(stage.PointPropertyTypes, stage.PointPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointpropertytypeIdent := pointpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PointPropertyType{Name: %s}).Stage(stage)", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.NilReason)))
			if pointpropertytype.Point != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := pointpropertytype.Point.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Point = %s", pointpropertytypeIdent, targetIdent))
			}
			if pointpropertytype.ElevatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := pointpropertytype.ElevatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedPoint = %s", pointpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pointreferencepropertytype := range __gong__sortStageSetInstances(stage.PointReferencePropertyTypes, stage.PointReferencePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointreferencepropertytypeIdent := pointreferencepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PointReferencePropertyType{Name: %s}).Stage(stage)", pointreferencepropertytypeIdent, __gong__toRawStringLiteral(pointreferencepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointreferencepropertytypeIdent, __gong__toRawStringLiteral(pointreferencepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", pointreferencepropertytypeIdent, __gong__toRawStringLiteral(pointreferencepropertytype.NilReason)))
			if pointreferencepropertytype.PointReference != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := pointreferencepropertytype.PointReference.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointReference = %s", pointreferencepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pointreferencetype := range __gong__sortStageSetInstances(stage.PointReferenceTypes, stage.PointReferenceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointreferencetypeIdent := pointreferencetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PointReferenceType{Name: %s}).Stage(stage)", pointreferencetypeIdent, __gong__toRawStringLiteral(pointreferencetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointreferencetypeIdent, __gong__toRawStringLiteral(pointreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pointreferencetypeIdent, __gong__toRawStringLiteral(pointreferencetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pointreferencetypeIdent, __gong__toRawStringLiteral(pointreferencetype.Description)))
			if pointreferencetype.Point != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := pointreferencetype.Point.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Point = %s", pointreferencetypeIdent, targetIdent))
			}
			for _, elem := range pointreferencetype.FacilityAngle {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FacilityAngle = append(%s.FacilityAngle, %s)", pointreferencetypeIdent, pointreferencetypeIdent, targetIdent))
			}
			for _, elem := range pointreferencetype.FacilityDistance {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FacilityDistance = append(%s.FacilityDistance, %s)", pointreferencetypeIdent, pointreferencetypeIdent, targetIdent))
			}
			if pointreferencetype.FixToleranceArea != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := pointreferencetype.FixToleranceArea.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FixToleranceArea = %s", pointreferencetypeIdent, targetIdent))
			}
			for _, elem := range pointreferencetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", pointreferencetypeIdent, pointreferencetypeIdent, targetIdent))
			}
			for _, elem := range pointreferencetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", pointreferencetypeIdent, pointreferencetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pointtype := range __gong__sortStageSetInstances(stage.PointTypes, stage.PointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointtypeIdent := pointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PointType{Name: %s}).Stage(stage)", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.SrsName)))
			for _, elem := range pointtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", pointtypeIdent, pointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, procedureavailabilitypropertytype := range __gong__sortStageSetInstances(stage.ProcedureAvailabilityPropertyTypes, stage.ProcedureAvailabilityPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			procedureavailabilitypropertytypeIdent := procedureavailabilitypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedureAvailabilityPropertyType{Name: %s}).Stage(stage)", procedureavailabilitypropertytypeIdent, __gong__toRawStringLiteral(procedureavailabilitypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", procedureavailabilitypropertytypeIdent, __gong__toRawStringLiteral(procedureavailabilitypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", procedureavailabilitypropertytypeIdent, __gong__toRawStringLiteral(procedureavailabilitypropertytype.NilReason)))
			if procedureavailabilitypropertytype.ProcedureAvailability != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := procedureavailabilitypropertytype.ProcedureAvailability.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProcedureAvailability = %s", procedureavailabilitypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, procedureavailabilitytype := range __gong__sortStageSetInstances(stage.ProcedureAvailabilityTypes, stage.ProcedureAvailabilityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			procedureavailabilitytypeIdent := procedureavailabilitytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedureAvailabilityType{Name: %s}).Stage(stage)", procedureavailabilitytypeIdent, __gong__toRawStringLiteral(procedureavailabilitytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", procedureavailabilitytypeIdent, __gong__toRawStringLiteral(procedureavailabilitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", procedureavailabilitytypeIdent, __gong__toRawStringLiteral(procedureavailabilitytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", procedureavailabilitytypeIdent, __gong__toRawStringLiteral(procedureavailabilitytype.Description)))
			for _, elem := range procedureavailabilitytype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", procedureavailabilitytypeIdent, procedureavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range procedureavailabilitytype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", procedureavailabilitytypeIdent, procedureavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range procedureavailabilitytype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", procedureavailabilitytypeIdent, procedureavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range procedureavailabilitytype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", procedureavailabilitytypeIdent, procedureavailabilitytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, proceduredmetimeslicepropertytype := range __gong__sortStageSetInstances(stage.ProcedureDMETimeSlicePropertyTypes, stage.ProcedureDMETimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			proceduredmetimeslicepropertytypeIdent := proceduredmetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedureDMETimeSlicePropertyType{Name: %s}).Stage(stage)", proceduredmetimeslicepropertytypeIdent, __gong__toRawStringLiteral(proceduredmetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", proceduredmetimeslicepropertytypeIdent, __gong__toRawStringLiteral(proceduredmetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", proceduredmetimeslicepropertytypeIdent, proceduredmetimeslicepropertytype.Owns))
			if proceduredmetimeslicepropertytype.ProcedureDMETimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := proceduredmetimeslicepropertytype.ProcedureDMETimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProcedureDMETimeSlice = %s", proceduredmetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, proceduredmetimeslicetype := range __gong__sortStageSetInstances(stage.ProcedureDMETimeSliceTypes, stage.ProcedureDMETimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			proceduredmetimeslicetypeIdent := proceduredmetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedureDMETimeSliceType{Name: %s}).Stage(stage)", proceduredmetimeslicetypeIdent, __gong__toRawStringLiteral(proceduredmetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", proceduredmetimeslicetypeIdent, __gong__toRawStringLiteral(proceduredmetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", proceduredmetimeslicetypeIdent, __gong__toRawStringLiteral(proceduredmetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", proceduredmetimeslicetypeIdent, __gong__toRawStringLiteral(proceduredmetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", proceduredmetimeslicetypeIdent, __gong__toRawStringLiteral(proceduredmetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", proceduredmetimeslicetypeIdent, proceduredmetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", proceduredmetimeslicetypeIdent, proceduredmetimeslicetype.CorrectionNumber))
			if proceduredmetimeslicetype.DME != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := proceduredmetimeslicetype.DME.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DME = %s", proceduredmetimeslicetypeIdent, targetIdent))
			}
			if proceduredmetimeslicetype.SegmentLeg != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := proceduredmetimeslicetype.SegmentLeg.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SegmentLeg = %s", proceduredmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range proceduredmetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", proceduredmetimeslicetypeIdent, proceduredmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range proceduredmetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", proceduredmetimeslicetypeIdent, proceduredmetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, proceduredmetype := range __gong__sortStageSetInstances(stage.ProcedureDMETypes, stage.ProcedureDMEType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			proceduredmetypeIdent := proceduredmetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedureDMEType{Name: %s}).Stage(stage)", proceduredmetypeIdent, __gong__toRawStringLiteral(proceduredmetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", proceduredmetypeIdent, __gong__toRawStringLiteral(proceduredmetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", proceduredmetypeIdent, __gong__toRawStringLiteral(proceduredmetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", proceduredmetypeIdent, __gong__toRawStringLiteral(proceduredmetype.Description)))
			for _, elem := range proceduredmetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", proceduredmetypeIdent, proceduredmetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, procedurepropertytype := range __gong__sortStageSetInstances(stage.ProcedurePropertyTypes, stage.ProcedurePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			procedurepropertytypeIdent := procedurepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedurePropertyType{Name: %s}).Stage(stage)", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", procedurepropertytypeIdent, procedurepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, proceduretransitionlegpropertytype := range __gong__sortStageSetInstances(stage.ProcedureTransitionLegPropertyTypes, stage.ProcedureTransitionLegPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			proceduretransitionlegpropertytypeIdent := proceduretransitionlegpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedureTransitionLegPropertyType{Name: %s}).Stage(stage)", proceduretransitionlegpropertytypeIdent, __gong__toRawStringLiteral(proceduretransitionlegpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", proceduretransitionlegpropertytypeIdent, __gong__toRawStringLiteral(proceduretransitionlegpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", proceduretransitionlegpropertytypeIdent, __gong__toRawStringLiteral(proceduretransitionlegpropertytype.NilReason)))
			if proceduretransitionlegpropertytype.ProcedureTransitionLeg != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := proceduretransitionlegpropertytype.ProcedureTransitionLeg.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProcedureTransitionLeg = %s", proceduretransitionlegpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, proceduretransitionlegtype := range __gong__sortStageSetInstances(stage.ProcedureTransitionLegTypes, stage.ProcedureTransitionLegType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			proceduretransitionlegtypeIdent := proceduretransitionlegtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedureTransitionLegType{Name: %s}).Stage(stage)", proceduretransitionlegtypeIdent, __gong__toRawStringLiteral(proceduretransitionlegtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", proceduretransitionlegtypeIdent, __gong__toRawStringLiteral(proceduretransitionlegtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", proceduretransitionlegtypeIdent, __gong__toRawStringLiteral(proceduretransitionlegtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", proceduretransitionlegtypeIdent, __gong__toRawStringLiteral(proceduretransitionlegtype.Description)))
			for _, elem := range proceduretransitionlegtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", proceduretransitionlegtypeIdent, proceduretransitionlegtypeIdent, targetIdent))
			}
			if proceduretransitionlegtype.TheSegmentLeg != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := proceduretransitionlegtype.TheSegmentLeg.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheSegmentLeg = %s", proceduretransitionlegtypeIdent, targetIdent))
			}
			for _, elem := range proceduretransitionlegtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", proceduretransitionlegtypeIdent, proceduretransitionlegtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, proceduretransitionpropertytype := range __gong__sortStageSetInstances(stage.ProcedureTransitionPropertyTypes, stage.ProcedureTransitionPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			proceduretransitionpropertytypeIdent := proceduretransitionpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedureTransitionPropertyType{Name: %s}).Stage(stage)", proceduretransitionpropertytypeIdent, __gong__toRawStringLiteral(proceduretransitionpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", proceduretransitionpropertytypeIdent, __gong__toRawStringLiteral(proceduretransitionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", proceduretransitionpropertytypeIdent, __gong__toRawStringLiteral(proceduretransitionpropertytype.NilReason)))
			if proceduretransitionpropertytype.ProcedureTransition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := proceduretransitionpropertytype.ProcedureTransition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProcedureTransition = %s", proceduretransitionpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, proceduretransitiontype := range __gong__sortStageSetInstances(stage.ProcedureTransitionTypes, stage.ProcedureTransitionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			proceduretransitiontypeIdent := proceduretransitiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedureTransitionType{Name: %s}).Stage(stage)", proceduretransitiontypeIdent, __gong__toRawStringLiteral(proceduretransitiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", proceduretransitiontypeIdent, __gong__toRawStringLiteral(proceduretransitiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", proceduretransitiontypeIdent, __gong__toRawStringLiteral(proceduretransitiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", proceduretransitiontypeIdent, __gong__toRawStringLiteral(proceduretransitiontype.Description)))
			if proceduretransitiontype.DepartureRunwayTransition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := proceduretransitiontype.DepartureRunwayTransition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DepartureRunwayTransition = %s", proceduretransitiontypeIdent, targetIdent))
			}
			if proceduretransitiontype.Trajectory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := proceduretransitiontype.Trajectory.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Trajectory = %s", proceduretransitiontypeIdent, targetIdent))
			}
			for _, elem := range proceduretransitiontype.TransitionLeg {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TransitionLeg = append(%s.TransitionLeg, %s)", proceduretransitiontypeIdent, proceduretransitiontypeIdent, targetIdent))
			}
			for _, elem := range proceduretransitiontype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", proceduretransitiontypeIdent, proceduretransitiontypeIdent, targetIdent))
			}
			for _, elem := range proceduretransitiontype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", proceduretransitiontypeIdent, proceduretransitiontypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, radarsystempropertytype := range __gong__sortStageSetInstances(stage.RadarSystemPropertyTypes, stage.RadarSystemPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radarsystempropertytypeIdent := radarsystempropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarSystemPropertyType{Name: %s}).Stage(stage)", radarsystempropertytypeIdent, __gong__toRawStringLiteral(radarsystempropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarsystempropertytypeIdent, __gong__toRawStringLiteral(radarsystempropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", radarsystempropertytypeIdent, radarsystempropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", radarsystempropertytypeIdent, __gong__toRawStringLiteral(radarsystempropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", radarsystempropertytypeIdent, __gong__toRawStringLiteral(radarsystempropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", radarsystempropertytypeIdent, __gong__toRawStringLiteral(radarsystempropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", radarsystempropertytypeIdent, __gong__toRawStringLiteral(radarsystempropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", radarsystempropertytypeIdent, __gong__toRawStringLiteral(radarsystempropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", radarsystempropertytypeIdent, __gong__toRawStringLiteral(radarsystempropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", radarsystempropertytypeIdent, __gong__toRawStringLiteral(radarsystempropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", radarsystempropertytypeIdent, __gong__toRawStringLiteral(radarsystempropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, radiofrequencyareapropertytype := range __gong__sortStageSetInstances(stage.RadioFrequencyAreaPropertyTypes, stage.RadioFrequencyAreaPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radiofrequencyareapropertytypeIdent := radiofrequencyareapropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioFrequencyAreaPropertyType{Name: %s}).Stage(stage)", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", radiofrequencyareapropertytypeIdent, radiofrequencyareapropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, runwaycentrelinepointpropertytype := range __gong__sortStageSetInstances(stage.RunwayCentrelinePointPropertyTypes, stage.RunwayCentrelinePointPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			runwaycentrelinepointpropertytypeIdent := runwaycentrelinepointpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RunwayCentrelinePointPropertyType{Name: %s}).Stage(stage)", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", runwaycentrelinepointpropertytypeIdent, runwaycentrelinepointpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, runwaydirectionpropertytype := range __gong__sortStageSetInstances(stage.RunwayDirectionPropertyTypes, stage.RunwayDirectionPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			runwaydirectionpropertytypeIdent := runwaydirectionpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RunwayDirectionPropertyType{Name: %s}).Stage(stage)", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", runwaydirectionpropertytypeIdent, runwaydirectionpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, safealtitudeareapropertytype := range __gong__sortStageSetInstances(stage.SafeAltitudeAreaPropertyTypes, stage.SafeAltitudeAreaPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			safealtitudeareapropertytypeIdent := safealtitudeareapropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SafeAltitudeAreaPropertyType{Name: %s}).Stage(stage)", safealtitudeareapropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareapropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", safealtitudeareapropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", safealtitudeareapropertytypeIdent, safealtitudeareapropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", safealtitudeareapropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareapropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", safealtitudeareapropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareapropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", safealtitudeareapropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareapropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", safealtitudeareapropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareapropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", safealtitudeareapropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareapropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", safealtitudeareapropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareapropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", safealtitudeareapropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareapropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", safealtitudeareapropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareapropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, safealtitudeareasectorpropertytype := range __gong__sortStageSetInstances(stage.SafeAltitudeAreaSectorPropertyTypes, stage.SafeAltitudeAreaSectorPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			safealtitudeareasectorpropertytypeIdent := safealtitudeareasectorpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SafeAltitudeAreaSectorPropertyType{Name: %s}).Stage(stage)", safealtitudeareasectorpropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareasectorpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", safealtitudeareasectorpropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareasectorpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", safealtitudeareasectorpropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareasectorpropertytype.NilReason)))
			if safealtitudeareasectorpropertytype.SafeAltitudeAreaSector != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := safealtitudeareasectorpropertytype.SafeAltitudeAreaSector.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SafeAltitudeAreaSector = %s", safealtitudeareasectorpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, safealtitudeareasectortype := range __gong__sortStageSetInstances(stage.SafeAltitudeAreaSectorTypes, stage.SafeAltitudeAreaSectorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			safealtitudeareasectortypeIdent := safealtitudeareasectortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SafeAltitudeAreaSectorType{Name: %s}).Stage(stage)", safealtitudeareasectortypeIdent, __gong__toRawStringLiteral(safealtitudeareasectortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", safealtitudeareasectortypeIdent, __gong__toRawStringLiteral(safealtitudeareasectortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", safealtitudeareasectortypeIdent, __gong__toRawStringLiteral(safealtitudeareasectortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", safealtitudeareasectortypeIdent, __gong__toRawStringLiteral(safealtitudeareasectortype.Description)))
			if safealtitudeareasectortype.Extent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := safealtitudeareasectortype.Extent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extent = %s", safealtitudeareasectortypeIdent, targetIdent))
			}
			for _, elem := range safealtitudeareasectortype.SignificantObstacle {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantObstacle = append(%s.SignificantObstacle, %s)", safealtitudeareasectortypeIdent, safealtitudeareasectortypeIdent, targetIdent))
			}
			if safealtitudeareasectortype.SectorDefinition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := safealtitudeareasectortype.SectorDefinition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SectorDefinition = %s", safealtitudeareasectortypeIdent, targetIdent))
			}
			for _, elem := range safealtitudeareasectortype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", safealtitudeareasectortypeIdent, safealtitudeareasectortypeIdent, targetIdent))
			}
			for _, elem := range safealtitudeareasectortype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", safealtitudeareasectortypeIdent, safealtitudeareasectortypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, safealtitudeareatimeslicepropertytype := range __gong__sortStageSetInstances(stage.SafeAltitudeAreaTimeSlicePropertyTypes, stage.SafeAltitudeAreaTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			safealtitudeareatimeslicepropertytypeIdent := safealtitudeareatimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SafeAltitudeAreaTimeSlicePropertyType{Name: %s}).Stage(stage)", safealtitudeareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareatimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", safealtitudeareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(safealtitudeareatimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", safealtitudeareatimeslicepropertytypeIdent, safealtitudeareatimeslicepropertytype.Owns))
			if safealtitudeareatimeslicepropertytype.SafeAltitudeAreaTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := safealtitudeareatimeslicepropertytype.SafeAltitudeAreaTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SafeAltitudeAreaTimeSlice = %s", safealtitudeareatimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, safealtitudeareatimeslicetype := range __gong__sortStageSetInstances(stage.SafeAltitudeAreaTimeSliceTypes, stage.SafeAltitudeAreaTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			safealtitudeareatimeslicetypeIdent := safealtitudeareatimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SafeAltitudeAreaTimeSliceType{Name: %s}).Stage(stage)", safealtitudeareatimeslicetypeIdent, __gong__toRawStringLiteral(safealtitudeareatimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", safealtitudeareatimeslicetypeIdent, __gong__toRawStringLiteral(safealtitudeareatimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", safealtitudeareatimeslicetypeIdent, __gong__toRawStringLiteral(safealtitudeareatimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", safealtitudeareatimeslicetypeIdent, __gong__toRawStringLiteral(safealtitudeareatimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", safealtitudeareatimeslicetypeIdent, __gong__toRawStringLiteral(safealtitudeareatimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", safealtitudeareatimeslicetypeIdent, safealtitudeareatimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", safealtitudeareatimeslicetypeIdent, safealtitudeareatimeslicetype.CorrectionNumber))
			if safealtitudeareatimeslicetype.CentrePoint_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := safealtitudeareatimeslicetype.CentrePoint_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_fixDesignatedPoint = %s", safealtitudeareatimeslicetypeIdent, targetIdent))
			}
			if safealtitudeareatimeslicetype.CentrePoint_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := safealtitudeareatimeslicetype.CentrePoint_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_navaidSystem = %s", safealtitudeareatimeslicetypeIdent, targetIdent))
			}
			if safealtitudeareatimeslicetype.CentrePoint_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := safealtitudeareatimeslicetype.CentrePoint_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_position = %s", safealtitudeareatimeslicetypeIdent, targetIdent))
			}
			if safealtitudeareatimeslicetype.CentrePoint_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := safealtitudeareatimeslicetype.CentrePoint_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_runwayPoint = %s", safealtitudeareatimeslicetypeIdent, targetIdent))
			}
			if safealtitudeareatimeslicetype.CentrePoint_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := safealtitudeareatimeslicetype.CentrePoint_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_aimingPoint = %s", safealtitudeareatimeslicetypeIdent, targetIdent))
			}
			if safealtitudeareatimeslicetype.CentrePoint_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := safealtitudeareatimeslicetype.CentrePoint_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CentrePoint_airportReferencePoint = %s", safealtitudeareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range safealtitudeareatimeslicetype.Sector {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sector = append(%s.Sector, %s)", safealtitudeareatimeslicetypeIdent, safealtitudeareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range safealtitudeareatimeslicetype.Location {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = append(%s.Location, %s)", safealtitudeareatimeslicetypeIdent, safealtitudeareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range safealtitudeareatimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", safealtitudeareatimeslicetypeIdent, safealtitudeareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range safealtitudeareatimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", safealtitudeareatimeslicetypeIdent, safealtitudeareatimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, safealtitudeareatype := range __gong__sortStageSetInstances(stage.SafeAltitudeAreaTypes, stage.SafeAltitudeAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			safealtitudeareatypeIdent := safealtitudeareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SafeAltitudeAreaType{Name: %s}).Stage(stage)", safealtitudeareatypeIdent, __gong__toRawStringLiteral(safealtitudeareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", safealtitudeareatypeIdent, __gong__toRawStringLiteral(safealtitudeareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", safealtitudeareatypeIdent, __gong__toRawStringLiteral(safealtitudeareatype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", safealtitudeareatypeIdent, __gong__toRawStringLiteral(safealtitudeareatype.Description)))
			for _, elem := range safealtitudeareatype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", safealtitudeareatypeIdent, safealtitudeareatypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, sectordesignpropertytype := range __gong__sortStageSetInstances(stage.SectorDesignPropertyTypes, stage.SectorDesignPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sectordesignpropertytypeIdent := sectordesignpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SectorDesignPropertyType{Name: %s}).Stage(stage)", sectordesignpropertytypeIdent, __gong__toRawStringLiteral(sectordesignpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sectordesignpropertytypeIdent, __gong__toRawStringLiteral(sectordesignpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", sectordesignpropertytypeIdent, __gong__toRawStringLiteral(sectordesignpropertytype.NilReason)))
			if sectordesignpropertytype.SectorDesign != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := sectordesignpropertytype.SectorDesign.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SectorDesign = %s", sectordesignpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, sectordesigntype := range __gong__sortStageSetInstances(stage.SectorDesignTypes, stage.SectorDesignType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sectordesigntypeIdent := sectordesigntype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SectorDesignType{Name: %s}).Stage(stage)", sectordesigntypeIdent, __gong__toRawStringLiteral(sectordesigntype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sectordesigntypeIdent, __gong__toRawStringLiteral(sectordesigntype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", sectordesigntypeIdent, __gong__toRawStringLiteral(sectordesigntype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", sectordesigntypeIdent, __gong__toRawStringLiteral(sectordesigntype.Description)))
			for _, elem := range sectordesigntype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", sectordesigntypeIdent, sectordesigntypeIdent, targetIdent))
			}
			for _, elem := range sectordesigntype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", sectordesigntypeIdent, sectordesigntypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, segmentlegpropertytype := range __gong__sortStageSetInstances(stage.SegmentLegPropertyTypes, stage.SegmentLegPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			segmentlegpropertytypeIdent := segmentlegpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SegmentLegPropertyType{Name: %s}).Stage(stage)", segmentlegpropertytypeIdent, __gong__toRawStringLiteral(segmentlegpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", segmentlegpropertytypeIdent, __gong__toRawStringLiteral(segmentlegpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", segmentlegpropertytypeIdent, segmentlegpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", segmentlegpropertytypeIdent, __gong__toRawStringLiteral(segmentlegpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", segmentlegpropertytypeIdent, __gong__toRawStringLiteral(segmentlegpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", segmentlegpropertytypeIdent, __gong__toRawStringLiteral(segmentlegpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", segmentlegpropertytypeIdent, __gong__toRawStringLiteral(segmentlegpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", segmentlegpropertytypeIdent, __gong__toRawStringLiteral(segmentlegpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", segmentlegpropertytypeIdent, __gong__toRawStringLiteral(segmentlegpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", segmentlegpropertytypeIdent, __gong__toRawStringLiteral(segmentlegpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", segmentlegpropertytypeIdent, __gong__toRawStringLiteral(segmentlegpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, segmentpointpropertytype := range __gong__sortStageSetInstances(stage.SegmentPointPropertyTypes, stage.SegmentPointPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			segmentpointpropertytypeIdent := segmentpointpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SegmentPointPropertyType{Name: %s}).Stage(stage)", segmentpointpropertytypeIdent, __gong__toRawStringLiteral(segmentpointpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", segmentpointpropertytypeIdent, __gong__toRawStringLiteral(segmentpointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", segmentpointpropertytypeIdent, __gong__toRawStringLiteral(segmentpointpropertytype.NilReason)))
			if segmentpointpropertytype.AerialRefuellingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := segmentpointpropertytype.AerialRefuellingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerialRefuellingPoint = %s", segmentpointpropertytypeIdent, targetIdent))
			}
			if segmentpointpropertytype.EnRouteSegmentPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := segmentpointpropertytype.EnRouteSegmentPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnRouteSegmentPoint = %s", segmentpointpropertytypeIdent, targetIdent))
			}
			if segmentpointpropertytype.TerminalSegmentPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := segmentpointpropertytype.TerminalSegmentPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TerminalSegmentPoint = %s", segmentpointpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialnavigationsystempropertytype := range __gong__sortStageSetInstances(stage.SpecialNavigationSystemPropertyTypes, stage.SpecialNavigationSystemPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialnavigationsystempropertytypeIdent := specialnavigationsystempropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationSystemPropertyType{Name: %s}).Stage(stage)", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", specialnavigationsystempropertytypeIdent, specialnavigationsystempropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, standardinstrumentarrivalpropertytype := range __gong__sortStageSetInstances(stage.StandardInstrumentArrivalPropertyTypes, stage.StandardInstrumentArrivalPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardinstrumentarrivalpropertytypeIdent := standardinstrumentarrivalpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardInstrumentArrivalPropertyType{Name: %s}).Stage(stage)", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardinstrumentarrivalpropertytypeIdent, standardinstrumentarrivalpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, standardinstrumentarrivaltimeslicepropertytype := range __gong__sortStageSetInstances(stage.StandardInstrumentArrivalTimeSlicePropertyTypes, stage.StandardInstrumentArrivalTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardinstrumentarrivaltimeslicepropertytypeIdent := standardinstrumentarrivaltimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardInstrumentArrivalTimeSlicePropertyType{Name: %s}).Stage(stage)", standardinstrumentarrivaltimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardinstrumentarrivaltimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardinstrumentarrivaltimeslicepropertytypeIdent, standardinstrumentarrivaltimeslicepropertytype.Owns))
			if standardinstrumentarrivaltimeslicepropertytype.StandardInstrumentArrivalTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentarrivaltimeslicepropertytype.StandardInstrumentArrivalTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardInstrumentArrivalTimeSlice = %s", standardinstrumentarrivaltimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardinstrumentarrivaltimeslicetype := range __gong__sortStageSetInstances(stage.StandardInstrumentArrivalTimeSliceTypes, stage.StandardInstrumentArrivalTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardinstrumentarrivaltimeslicetypeIdent := standardinstrumentarrivaltimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardInstrumentArrivalTimeSliceType{Name: %s}).Stage(stage)", standardinstrumentarrivaltimeslicetypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardinstrumentarrivaltimeslicetypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardinstrumentarrivaltimeslicetypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardinstrumentarrivaltimeslicetypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", standardinstrumentarrivaltimeslicetypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", standardinstrumentarrivaltimeslicetypeIdent, standardinstrumentarrivaltimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", standardinstrumentarrivaltimeslicetypeIdent, standardinstrumentarrivaltimeslicetype.CorrectionNumber))
			for _, elem := range standardinstrumentarrivaltimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", standardinstrumentarrivaltimeslicetypeIdent, standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardinstrumentarrivaltimeslicetype.AirportHeliport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", standardinstrumentarrivaltimeslicetypeIdent, standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardinstrumentarrivaltimeslicetype.AircraftCharacteristic {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCharacteristic = append(%s.AircraftCharacteristic, %s)", standardinstrumentarrivaltimeslicetypeIdent, standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardinstrumentarrivaltimeslicetype.FlightTransition {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightTransition = append(%s.FlightTransition, %s)", standardinstrumentarrivaltimeslicetypeIdent, standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
			if standardinstrumentarrivaltimeslicetype.GuidanceFacility_specialNavigationSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentarrivaltimeslicetype.GuidanceFacility_specialNavigationSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GuidanceFacility_specialNavigationSystem = %s", standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
			if standardinstrumentarrivaltimeslicetype.GuidanceFacility_radar != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentarrivaltimeslicetype.GuidanceFacility_radar.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GuidanceFacility_radar = %s", standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
			if standardinstrumentarrivaltimeslicetype.GuidanceFacility_navaid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentarrivaltimeslicetype.GuidanceFacility_navaid.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GuidanceFacility_navaid = %s", standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardinstrumentarrivaltimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", standardinstrumentarrivaltimeslicetypeIdent, standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
			if standardinstrumentarrivaltimeslicetype.SafeAltitude != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentarrivaltimeslicetype.SafeAltitude.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SafeAltitude = %s", standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
			if standardinstrumentarrivaltimeslicetype.Arrival != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentarrivaltimeslicetype.Arrival.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arrival = %s", standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardinstrumentarrivaltimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", standardinstrumentarrivaltimeslicetypeIdent, standardinstrumentarrivaltimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardinstrumentarrivaltype := range __gong__sortStageSetInstances(stage.StandardInstrumentArrivalTypes, stage.StandardInstrumentArrivalType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardinstrumentarrivaltypeIdent := standardinstrumentarrivaltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardInstrumentArrivalType{Name: %s}).Stage(stage)", standardinstrumentarrivaltypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardinstrumentarrivaltypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardinstrumentarrivaltypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardinstrumentarrivaltypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivaltype.Description)))
			for _, elem := range standardinstrumentarrivaltype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", standardinstrumentarrivaltypeIdent, standardinstrumentarrivaltypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardinstrumentdeparturepropertytype := range __gong__sortStageSetInstances(stage.StandardInstrumentDeparturePropertyTypes, stage.StandardInstrumentDeparturePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardinstrumentdeparturepropertytypeIdent := standardinstrumentdeparturepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardInstrumentDeparturePropertyType{Name: %s}).Stage(stage)", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardinstrumentdeparturepropertytypeIdent, standardinstrumentdeparturepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, standardinstrumentdeparturetimeslicepropertytype := range __gong__sortStageSetInstances(stage.StandardInstrumentDepartureTimeSlicePropertyTypes, stage.StandardInstrumentDepartureTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardinstrumentdeparturetimeslicepropertytypeIdent := standardinstrumentdeparturetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardInstrumentDepartureTimeSlicePropertyType{Name: %s}).Stage(stage)", standardinstrumentdeparturetimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardinstrumentdeparturetimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardinstrumentdeparturetimeslicepropertytypeIdent, standardinstrumentdeparturetimeslicepropertytype.Owns))
			if standardinstrumentdeparturetimeslicepropertytype.StandardInstrumentDepartureTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentdeparturetimeslicepropertytype.StandardInstrumentDepartureTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardInstrumentDepartureTimeSlice = %s", standardinstrumentdeparturetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardinstrumentdeparturetimeslicetype := range __gong__sortStageSetInstances(stage.StandardInstrumentDepartureTimeSliceTypes, stage.StandardInstrumentDepartureTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardinstrumentdeparturetimeslicetypeIdent := standardinstrumentdeparturetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardInstrumentDepartureTimeSliceType{Name: %s}).Stage(stage)", standardinstrumentdeparturetimeslicetypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardinstrumentdeparturetimeslicetypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardinstrumentdeparturetimeslicetypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardinstrumentdeparturetimeslicetypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", standardinstrumentdeparturetimeslicetypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", standardinstrumentdeparturetimeslicetypeIdent, standardinstrumentdeparturetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", standardinstrumentdeparturetimeslicetypeIdent, standardinstrumentdeparturetimeslicetype.CorrectionNumber))
			for _, elem := range standardinstrumentdeparturetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", standardinstrumentdeparturetimeslicetypeIdent, standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardinstrumentdeparturetimeslicetype.AirportHeliport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", standardinstrumentdeparturetimeslicetypeIdent, standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardinstrumentdeparturetimeslicetype.AircraftCharacteristic {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCharacteristic = append(%s.AircraftCharacteristic, %s)", standardinstrumentdeparturetimeslicetypeIdent, standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardinstrumentdeparturetimeslicetype.FlightTransition {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightTransition = append(%s.FlightTransition, %s)", standardinstrumentdeparturetimeslicetypeIdent, standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
			if standardinstrumentdeparturetimeslicetype.GuidanceFacility_specialNavigationSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentdeparturetimeslicetype.GuidanceFacility_specialNavigationSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GuidanceFacility_specialNavigationSystem = %s", standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
			if standardinstrumentdeparturetimeslicetype.GuidanceFacility_radar != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentdeparturetimeslicetype.GuidanceFacility_radar.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GuidanceFacility_radar = %s", standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
			if standardinstrumentdeparturetimeslicetype.GuidanceFacility_navaid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentdeparturetimeslicetype.GuidanceFacility_navaid.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GuidanceFacility_navaid = %s", standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardinstrumentdeparturetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", standardinstrumentdeparturetimeslicetypeIdent, standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
			if standardinstrumentdeparturetimeslicetype.SafeAltitude != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentdeparturetimeslicetype.SafeAltitude.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SafeAltitude = %s", standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
			if standardinstrumentdeparturetimeslicetype.Takeoff != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardinstrumentdeparturetimeslicetype.Takeoff.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Takeoff = %s", standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardinstrumentdeparturetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", standardinstrumentdeparturetimeslicetypeIdent, standardinstrumentdeparturetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardinstrumentdeparturetype := range __gong__sortStageSetInstances(stage.StandardInstrumentDepartureTypes, stage.StandardInstrumentDepartureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardinstrumentdeparturetypeIdent := standardinstrumentdeparturetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardInstrumentDepartureType{Name: %s}).Stage(stage)", standardinstrumentdeparturetypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardinstrumentdeparturetypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardinstrumentdeparturetypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardinstrumentdeparturetypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturetype.Description)))
			for _, elem := range standardinstrumentdeparturetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", standardinstrumentdeparturetypeIdent, standardinstrumentdeparturetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, surfacepropertytype := range __gong__sortStageSetInstances(stage.SurfacePropertyTypes, stage.SurfacePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			surfacepropertytypeIdent := surfacepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SurfacePropertyType{Name: %s}).Stage(stage)", surfacepropertytypeIdent, __gong__toRawStringLiteral(surfacepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surfacepropertytypeIdent, __gong__toRawStringLiteral(surfacepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", surfacepropertytypeIdent, __gong__toRawStringLiteral(surfacepropertytype.NilReason)))
			if surfacepropertytype.Surface != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := surfacepropertytype.Surface.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Surface = %s", surfacepropertytypeIdent, targetIdent))
			}
			if surfacepropertytype.ElevatedSurface != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := surfacepropertytype.ElevatedSurface.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedSurface = %s", surfacepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, surfacetype := range __gong__sortStageSetInstances(stage.SurfaceTypes, stage.SurfaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			surfacetypeIdent := surfacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SurfaceType{Name: %s}).Stage(stage)", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.SrsName)))
			for _, elem := range surfacetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", surfacetypeIdent, surfacetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, terminalarrivalareasectorpropertytype := range __gong__sortStageSetInstances(stage.TerminalArrivalAreaSectorPropertyTypes, stage.TerminalArrivalAreaSectorPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			terminalarrivalareasectorpropertytypeIdent := terminalarrivalareasectorpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TerminalArrivalAreaSectorPropertyType{Name: %s}).Stage(stage)", terminalarrivalareasectorpropertytypeIdent, __gong__toRawStringLiteral(terminalarrivalareasectorpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", terminalarrivalareasectorpropertytypeIdent, __gong__toRawStringLiteral(terminalarrivalareasectorpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", terminalarrivalareasectorpropertytypeIdent, __gong__toRawStringLiteral(terminalarrivalareasectorpropertytype.NilReason)))
			if terminalarrivalareasectorpropertytype.TerminalArrivalAreaSector != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareasectorpropertytype.TerminalArrivalAreaSector.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TerminalArrivalAreaSector = %s", terminalarrivalareasectorpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, terminalarrivalareasectortype := range __gong__sortStageSetInstances(stage.TerminalArrivalAreaSectorTypes, stage.TerminalArrivalAreaSectorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			terminalarrivalareasectortypeIdent := terminalarrivalareasectortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TerminalArrivalAreaSectorType{Name: %s}).Stage(stage)", terminalarrivalareasectortypeIdent, __gong__toRawStringLiteral(terminalarrivalareasectortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", terminalarrivalareasectortypeIdent, __gong__toRawStringLiteral(terminalarrivalareasectortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", terminalarrivalareasectortypeIdent, __gong__toRawStringLiteral(terminalarrivalareasectortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", terminalarrivalareasectortypeIdent, __gong__toRawStringLiteral(terminalarrivalareasectortype.Description)))
			if terminalarrivalareasectortype.SectorDefinition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareasectortype.SectorDefinition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SectorDefinition = %s", terminalarrivalareasectortypeIdent, targetIdent))
			}
			if terminalarrivalareasectortype.Extent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareasectortype.Extent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extent = %s", terminalarrivalareasectortypeIdent, targetIdent))
			}
			for _, elem := range terminalarrivalareasectortype.SignificantObstacle {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantObstacle = append(%s.SignificantObstacle, %s)", terminalarrivalareasectortypeIdent, terminalarrivalareasectortypeIdent, targetIdent))
			}
			for _, elem := range terminalarrivalareasectortype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", terminalarrivalareasectortypeIdent, terminalarrivalareasectortypeIdent, targetIdent))
			}
			for _, elem := range terminalarrivalareasectortype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", terminalarrivalareasectortypeIdent, terminalarrivalareasectortypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, terminalarrivalareatimeslicepropertytype := range __gong__sortStageSetInstances(stage.TerminalArrivalAreaTimeSlicePropertyTypes, stage.TerminalArrivalAreaTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			terminalarrivalareatimeslicepropertytypeIdent := terminalarrivalareatimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TerminalArrivalAreaTimeSlicePropertyType{Name: %s}).Stage(stage)", terminalarrivalareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(terminalarrivalareatimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", terminalarrivalareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(terminalarrivalareatimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", terminalarrivalareatimeslicepropertytypeIdent, terminalarrivalareatimeslicepropertytype.Owns))
			if terminalarrivalareatimeslicepropertytype.TerminalArrivalAreaTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicepropertytype.TerminalArrivalAreaTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TerminalArrivalAreaTimeSlice = %s", terminalarrivalareatimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, terminalarrivalareatimeslicetype := range __gong__sortStageSetInstances(stage.TerminalArrivalAreaTimeSliceTypes, stage.TerminalArrivalAreaTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			terminalarrivalareatimeslicetypeIdent := terminalarrivalareatimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TerminalArrivalAreaTimeSliceType{Name: %s}).Stage(stage)", terminalarrivalareatimeslicetypeIdent, __gong__toRawStringLiteral(terminalarrivalareatimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", terminalarrivalareatimeslicetypeIdent, __gong__toRawStringLiteral(terminalarrivalareatimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", terminalarrivalareatimeslicetypeIdent, __gong__toRawStringLiteral(terminalarrivalareatimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", terminalarrivalareatimeslicetypeIdent, __gong__toRawStringLiteral(terminalarrivalareatimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", terminalarrivalareatimeslicetypeIdent, __gong__toRawStringLiteral(terminalarrivalareatimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", terminalarrivalareatimeslicetypeIdent, terminalarrivalareatimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", terminalarrivalareatimeslicetypeIdent, terminalarrivalareatimeslicetype.CorrectionNumber))
			if terminalarrivalareatimeslicetype.IF_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IF_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IF_fixDesignatedPoint = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IF_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IF_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IF_navaidSystem = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IF_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IF_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IF_position = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IF_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IF_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IF_runwayPoint = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IF_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IF_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IF_aimingPoint = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IF_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IF_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IF_airportReferencePoint = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IAF_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IAF_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IAF_fixDesignatedPoint = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IAF_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IAF_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IAF_navaidSystem = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IAF_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IAF_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IAF_position = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IAF_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IAF_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IAF_runwayPoint = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IAF_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IAF_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IAF_aimingPoint = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.IAF_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.IAF_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IAF_airportReferencePoint = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.Buffer != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.Buffer.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Buffer = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range terminalarrivalareatimeslicetype.Sector {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sector = append(%s.Sector, %s)", terminalarrivalareatimeslicetypeIdent, terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			if terminalarrivalareatimeslicetype.ApproachRNAV != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalarrivalareatimeslicetype.ApproachRNAV.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApproachRNAV = %s", terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range terminalarrivalareatimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", terminalarrivalareatimeslicetypeIdent, terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range terminalarrivalareatimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", terminalarrivalareatimeslicetypeIdent, terminalarrivalareatimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, terminalarrivalareatype := range __gong__sortStageSetInstances(stage.TerminalArrivalAreaTypes, stage.TerminalArrivalAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			terminalarrivalareatypeIdent := terminalarrivalareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TerminalArrivalAreaType{Name: %s}).Stage(stage)", terminalarrivalareatypeIdent, __gong__toRawStringLiteral(terminalarrivalareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", terminalarrivalareatypeIdent, __gong__toRawStringLiteral(terminalarrivalareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", terminalarrivalareatypeIdent, __gong__toRawStringLiteral(terminalarrivalareatype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", terminalarrivalareatypeIdent, __gong__toRawStringLiteral(terminalarrivalareatype.Description)))
			for _, elem := range terminalarrivalareatype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", terminalarrivalareatypeIdent, terminalarrivalareatypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, terminalsegmentpointpropertytype := range __gong__sortStageSetInstances(stage.TerminalSegmentPointPropertyTypes, stage.TerminalSegmentPointPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			terminalsegmentpointpropertytypeIdent := terminalsegmentpointpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TerminalSegmentPointPropertyType{Name: %s}).Stage(stage)", terminalsegmentpointpropertytypeIdent, __gong__toRawStringLiteral(terminalsegmentpointpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", terminalsegmentpointpropertytypeIdent, __gong__toRawStringLiteral(terminalsegmentpointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", terminalsegmentpointpropertytypeIdent, __gong__toRawStringLiteral(terminalsegmentpointpropertytype.NilReason)))
			if terminalsegmentpointpropertytype.TerminalSegmentPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalsegmentpointpropertytype.TerminalSegmentPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TerminalSegmentPoint = %s", terminalsegmentpointpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, terminalsegmentpointtype := range __gong__sortStageSetInstances(stage.TerminalSegmentPointTypes, stage.TerminalSegmentPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			terminalsegmentpointtypeIdent := terminalsegmentpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TerminalSegmentPointType{Name: %s}).Stage(stage)", terminalsegmentpointtypeIdent, __gong__toRawStringLiteral(terminalsegmentpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", terminalsegmentpointtypeIdent, __gong__toRawStringLiteral(terminalsegmentpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", terminalsegmentpointtypeIdent, __gong__toRawStringLiteral(terminalsegmentpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", terminalsegmentpointtypeIdent, __gong__toRawStringLiteral(terminalsegmentpointtype.Description)))
			for _, elem := range terminalsegmentpointtype.FacilityMakeup {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FacilityMakeup = append(%s.FacilityMakeup, %s)", terminalsegmentpointtypeIdent, terminalsegmentpointtypeIdent, targetIdent))
			}
			if terminalsegmentpointtype.PointChoice_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalsegmentpointtype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_fixDesignatedPoint = %s", terminalsegmentpointtypeIdent, targetIdent))
			}
			if terminalsegmentpointtype.PointChoice_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalsegmentpointtype.PointChoice_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_navaidSystem = %s", terminalsegmentpointtypeIdent, targetIdent))
			}
			if terminalsegmentpointtype.PointChoice_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalsegmentpointtype.PointChoice_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_position = %s", terminalsegmentpointtypeIdent, targetIdent))
			}
			if terminalsegmentpointtype.PointChoice_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalsegmentpointtype.PointChoice_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_runwayPoint = %s", terminalsegmentpointtypeIdent, targetIdent))
			}
			if terminalsegmentpointtype.PointChoice_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalsegmentpointtype.PointChoice_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_aimingPoint = %s", terminalsegmentpointtypeIdent, targetIdent))
			}
			if terminalsegmentpointtype.PointChoice_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalsegmentpointtype.PointChoice_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_airportReferencePoint = %s", terminalsegmentpointtypeIdent, targetIdent))
			}
			if terminalsegmentpointtype.ExtendedServiceVolume != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := terminalsegmentpointtype.ExtendedServiceVolume.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtendedServiceVolume = %s", terminalsegmentpointtypeIdent, targetIdent))
			}
			for _, elem := range terminalsegmentpointtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", terminalsegmentpointtypeIdent, terminalsegmentpointtypeIdent, targetIdent))
			}
			for _, elem := range terminalsegmentpointtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", terminalsegmentpointtypeIdent, terminalsegmentpointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, timesheetpropertytype := range __gong__sortStageSetInstances(stage.TimesheetPropertyTypes, stage.TimesheetPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timesheetpropertytypeIdent := timesheetpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TimesheetPropertyType{Name: %s}).Stage(stage)", timesheetpropertytypeIdent, __gong__toRawStringLiteral(timesheetpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timesheetpropertytypeIdent, __gong__toRawStringLiteral(timesheetpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", timesheetpropertytypeIdent, __gong__toRawStringLiteral(timesheetpropertytype.NilReason)))
			if timesheetpropertytype.Timesheet != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := timesheetpropertytype.Timesheet.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Timesheet = %s", timesheetpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, timesheettype := range __gong__sortStageSetInstances(stage.TimesheetTypes, stage.TimesheetType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timesheettypeIdent := timesheettype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TimesheetType{Name: %s}).Stage(stage)", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Description)))
			for _, elem := range timesheettype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", timesheettypeIdent, timesheettypeIdent, targetIdent))
			}
			for _, elem := range timesheettype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", timesheettypeIdent, timesheettypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, touchdownliftoffpropertytype := range __gong__sortStageSetInstances(stage.TouchDownLiftOffPropertyTypes, stage.TouchDownLiftOffPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			touchdownliftoffpropertytypeIdent := touchdownliftoffpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TouchDownLiftOffPropertyType{Name: %s}).Stage(stage)", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", touchdownliftoffpropertytypeIdent, touchdownliftoffpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, unplannedholdingpropertytype := range __gong__sortStageSetInstances(stage.UnplannedHoldingPropertyTypes, stage.UnplannedHoldingPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unplannedholdingpropertytypeIdent := unplannedholdingpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnplannedHoldingPropertyType{Name: %s}).Stage(stage)", unplannedholdingpropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unplannedholdingpropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", unplannedholdingpropertytypeIdent, unplannedholdingpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", unplannedholdingpropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", unplannedholdingpropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", unplannedholdingpropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", unplannedholdingpropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", unplannedholdingpropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", unplannedholdingpropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", unplannedholdingpropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", unplannedholdingpropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, unplannedholdingtimeslicepropertytype := range __gong__sortStageSetInstances(stage.UnplannedHoldingTimeSlicePropertyTypes, stage.UnplannedHoldingTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unplannedholdingtimeslicepropertytypeIdent := unplannedholdingtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnplannedHoldingTimeSlicePropertyType{Name: %s}).Stage(stage)", unplannedholdingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unplannedholdingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(unplannedholdingtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", unplannedholdingtimeslicepropertytypeIdent, unplannedholdingtimeslicepropertytype.Owns))
			if unplannedholdingtimeslicepropertytype.UnplannedHoldingTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := unplannedholdingtimeslicepropertytype.UnplannedHoldingTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UnplannedHoldingTimeSlice = %s", unplannedholdingtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, unplannedholdingtimeslicetype := range __gong__sortStageSetInstances(stage.UnplannedHoldingTimeSliceTypes, stage.UnplannedHoldingTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unplannedholdingtimeslicetypeIdent := unplannedholdingtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnplannedHoldingTimeSliceType{Name: %s}).Stage(stage)", unplannedholdingtimeslicetypeIdent, __gong__toRawStringLiteral(unplannedholdingtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unplannedholdingtimeslicetypeIdent, __gong__toRawStringLiteral(unplannedholdingtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", unplannedholdingtimeslicetypeIdent, __gong__toRawStringLiteral(unplannedholdingtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", unplannedholdingtimeslicetypeIdent, __gong__toRawStringLiteral(unplannedholdingtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", unplannedholdingtimeslicetypeIdent, __gong__toRawStringLiteral(unplannedholdingtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", unplannedholdingtimeslicetypeIdent, unplannedholdingtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", unplannedholdingtimeslicetypeIdent, unplannedholdingtimeslicetype.CorrectionNumber))
			if unplannedholdingtimeslicetype.HoldingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := unplannedholdingtimeslicetype.HoldingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoldingPoint = %s", unplannedholdingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unplannedholdingtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", unplannedholdingtimeslicetypeIdent, unplannedholdingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unplannedholdingtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", unplannedholdingtimeslicetypeIdent, unplannedholdingtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, unplannedholdingtype := range __gong__sortStageSetInstances(stage.UnplannedHoldingTypes, stage.UnplannedHoldingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unplannedholdingtypeIdent := unplannedholdingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnplannedHoldingType{Name: %s}).Stage(stage)", unplannedholdingtypeIdent, __gong__toRawStringLiteral(unplannedholdingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unplannedholdingtypeIdent, __gong__toRawStringLiteral(unplannedholdingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", unplannedholdingtypeIdent, __gong__toRawStringLiteral(unplannedholdingtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", unplannedholdingtypeIdent, __gong__toRawStringLiteral(unplannedholdingtype.Description)))
			for _, elem := range unplannedholdingtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", unplannedholdingtypeIdent, unplannedholdingtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, verticalstructurepropertytype := range __gong__sortStageSetInstances(stage.VerticalStructurePropertyTypes, stage.VerticalStructurePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			verticalstructurepropertytypeIdent := verticalstructurepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructurePropertyType{Name: %s}).Stage(stage)", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", verticalstructurepropertytypeIdent, verticalstructurepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Actuate)))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmfeatures/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmfeatures"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *aixmfeatures.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *aixmfeatures.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmfeatures" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmfeatures\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (a_extension *A_extension) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension", fieldName)
	}
	return
}

func (a_extension_1 *A_extension_1) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_1.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_1.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_1.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_1", fieldName)
	}
	return
}

func (a_extension_10 *A_extension_10) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_10.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_10.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_10.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_10", fieldName)
	}
	return
}

func (a_extension_11 *A_extension_11) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_11.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_11.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_11.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_11", fieldName)
	}
	return
}

func (a_extension_12 *A_extension_12) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_12.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_12.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_12.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_12", fieldName)
	}
	return
}

func (a_extension_13 *A_extension_13) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_13.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_13.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_13.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_13", fieldName)
	}
	return
}

func (a_extension_14 *A_extension_14) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_14.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_14.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_14.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_14", fieldName)
	}
	return
}

func (a_extension_15 *A_extension_15) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_15.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_15.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_15.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_15", fieldName)
	}
	return
}

func (a_extension_16 *A_extension_16) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_16.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_16.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_16.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_16", fieldName)
	}
	return
}

func (a_extension_17 *A_extension_17) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_17.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_17.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_17.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_17", fieldName)
	}
	return
}

func (a_extension_18 *A_extension_18) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_18.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_18.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_18.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_18", fieldName)
	}
	return
}

func (a_extension_19 *A_extension_19) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_19.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_19.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_19.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_19", fieldName)
	}
	return
}

func (a_extension_2 *A_extension_2) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_2.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_2.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_2.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_2", fieldName)
	}
	return
}

func (a_extension_20 *A_extension_20) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_20.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_20.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_20.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_20", fieldName)
	}
	return
}

func (a_extension_21 *A_extension_21) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_21.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_21.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_21.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_21", fieldName)
	}
	return
}

func (a_extension_22 *A_extension_22) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_22.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_22.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_22.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_22", fieldName)
	}
	return
}

func (a_extension_23 *A_extension_23) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_23.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_23.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_23.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_23", fieldName)
	}
	return
}

func (a_extension_24 *A_extension_24) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_24.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_24.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_24.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_24", fieldName)
	}
	return
}

func (a_extension_25 *A_extension_25) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_25.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_25.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_25.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_25", fieldName)
	}
	return
}

func (a_extension_26 *A_extension_26) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_26.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_26.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_26.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_26", fieldName)
	}
	return
}

func (a_extension_27 *A_extension_27) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_27.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_27.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_27.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_27", fieldName)
	}
	return
}

func (a_extension_28 *A_extension_28) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_28.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_28.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_28.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_28", fieldName)
	}
	return
}

func (a_extension_29 *A_extension_29) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_29.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_29.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_29.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_29", fieldName)
	}
	return
}

func (a_extension_3 *A_extension_3) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_3.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_3.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_3.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_3", fieldName)
	}
	return
}

func (a_extension_30 *A_extension_30) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_30.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_30.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_30.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_30", fieldName)
	}
	return
}

func (a_extension_31 *A_extension_31) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_31.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_31.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_31.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_31", fieldName)
	}
	return
}

func (a_extension_32 *A_extension_32) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_32.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_32.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_32.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_32", fieldName)
	}
	return
}

func (a_extension_33 *A_extension_33) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_33.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_33.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_33.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_33", fieldName)
	}
	return
}

func (a_extension_34 *A_extension_34) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_34.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_34.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_34.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_34", fieldName)
	}
	return
}

func (a_extension_35 *A_extension_35) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_35.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_35.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_35.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_35", fieldName)
	}
	return
}

func (a_extension_36 *A_extension_36) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_36.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_36.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_36.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_36", fieldName)
	}
	return
}

func (a_extension_37 *A_extension_37) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_37.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_37.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_37.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_37", fieldName)
	}
	return
}

func (a_extension_38 *A_extension_38) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_38.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_38.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_38.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_38", fieldName)
	}
	return
}

func (a_extension_39 *A_extension_39) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_39.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_39.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_39.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_39", fieldName)
	}
	return
}

func (a_extension_4 *A_extension_4) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_4.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_4.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_4.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_4", fieldName)
	}
	return
}

func (a_extension_40 *A_extension_40) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_40.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_40.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_40.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_40", fieldName)
	}
	return
}

func (a_extension_41 *A_extension_41) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_41.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_41.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_41.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_41", fieldName)
	}
	return
}

func (a_extension_42 *A_extension_42) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_42.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_42.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_42.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_42", fieldName)
	}
	return
}

func (a_extension_43 *A_extension_43) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_43.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_43.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_43.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_43", fieldName)
	}
	return
}

func (a_extension_44 *A_extension_44) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_44.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_44.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_44.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_44", fieldName)
	}
	return
}

func (a_extension_45 *A_extension_45) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_45.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_45.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_45.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_45", fieldName)
	}
	return
}

func (a_extension_46 *A_extension_46) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_46.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_46.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_46.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_46", fieldName)
	}
	return
}

func (a_extension_47 *A_extension_47) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_47.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_47.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_47.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_47", fieldName)
	}
	return
}

func (a_extension_48 *A_extension_48) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_48.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_48.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_48.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_48", fieldName)
	}
	return
}

func (a_extension_49 *A_extension_49) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_49.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_49.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_49.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_49", fieldName)
	}
	return
}

func (a_extension_5 *A_extension_5) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_5.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_5.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_5.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_5", fieldName)
	}
	return
}

func (a_extension_50 *A_extension_50) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_50.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_50.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_50.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_50", fieldName)
	}
	return
}

func (a_extension_51 *A_extension_51) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_51.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_51.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_51.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_51", fieldName)
	}
	return
}

func (a_extension_52 *A_extension_52) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_52.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_52.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_52.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_52", fieldName)
	}
	return
}

func (a_extension_53 *A_extension_53) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_53.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_53.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_53.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_53", fieldName)
	}
	return
}

func (a_extension_54 *A_extension_54) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_54.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_54.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_54.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_54", fieldName)
	}
	return
}

func (a_extension_55 *A_extension_55) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_55.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_55.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_55.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_55", fieldName)
	}
	return
}

func (a_extension_56 *A_extension_56) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_56.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_56.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_56.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_56", fieldName)
	}
	return
}

func (a_extension_57 *A_extension_57) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_57.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_57.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_57.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_57", fieldName)
	}
	return
}

func (a_extension_58 *A_extension_58) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_58.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_58.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_58.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_58", fieldName)
	}
	return
}

func (a_extension_6 *A_extension_6) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_6.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_6.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_6.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_6", fieldName)
	}
	return
}

func (a_extension_7 *A_extension_7) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_7.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_7.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_7.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_7", fieldName)
	}
	return
}

func (a_extension_8 *A_extension_8) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_8.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_8.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_8.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_8", fieldName)
	}
	return
}

func (a_extension_9 *A_extension_9) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_9.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_9.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_9.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_9", fieldName)
	}
	return
}

func (abstractapproachlegtype *AbstractApproachLegType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractapproachlegtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractapproachlegtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractapproachlegtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractapproachlegtype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractApproachLegType", fieldName)
	}
	return
}

func (abstractproceduretype *AbstractProcedureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractproceduretype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractproceduretype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractproceduretype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractproceduretype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractProcedureType", fieldName)
	}
	return
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractpropertieswithscheduletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractpropertieswithscheduletype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractpropertieswithscheduletype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractpropertieswithscheduletype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractPropertiesWithScheduleType", fieldName)
	}
	return
}

func (abstractsegmentlegtype *AbstractSegmentLegType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractsegmentlegtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractsegmentlegtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractsegmentlegtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractsegmentlegtype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractSegmentLegType", fieldName)
	}
	return
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractsegmentpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractsegmentpointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractsegmentpointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractsegmentpointtype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractSegmentPointType", fieldName)
	}
	return
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aerialrefuellingpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aerialrefuellingpointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aerialrefuellingpointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aerialrefuellingpointtype.Description)

	case "FacilityMakeup":
		var sb strings.Builder
		for _, _pointreferencepropertytype := range aerialrefuellingpointtype.FacilityMakeup {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FacilityMakeup", _pointreferencepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PointChoice_fixDesignatedPoint":
		if aerialrefuellingpointtype.PointChoice_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_fixDesignatedPoint", aerialrefuellingpointtype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_fixDesignatedPoint", "nil")
		}
	case "PointChoice_navaidSystem":
		if aerialrefuellingpointtype.PointChoice_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "PointChoice_navaidSystem", aerialrefuellingpointtype.PointChoice_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_navaidSystem", "nil")
		}
	case "PointChoice_position":
		if aerialrefuellingpointtype.PointChoice_position != nil {
			res = __gong__marshallPointer(ident, "PointChoice_position", aerialrefuellingpointtype.PointChoice_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_position", "nil")
		}
	case "PointChoice_runwayPoint":
		if aerialrefuellingpointtype.PointChoice_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_runwayPoint", aerialrefuellingpointtype.PointChoice_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_runwayPoint", "nil")
		}
	case "PointChoice_aimingPoint":
		if aerialrefuellingpointtype.PointChoice_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_aimingPoint", aerialrefuellingpointtype.PointChoice_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_aimingPoint", "nil")
		}
	case "PointChoice_airportReferencePoint":
		if aerialrefuellingpointtype.PointChoice_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_airportReferencePoint", aerialrefuellingpointtype.PointChoice_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_airportReferencePoint", "nil")
		}
	case "ExtendedServiceVolume":
		if aerialrefuellingpointtype.ExtendedServiceVolume != nil {
			res = __gong__marshallPointer(ident, "ExtendedServiceVolume", aerialrefuellingpointtype.ExtendedServiceVolume.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ExtendedServiceVolume", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range aerialrefuellingpointtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension := range aerialrefuellingpointtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingPointType", fieldName)
	}
	return
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aircraftcharacteristicpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aircraftcharacteristicpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", aircraftcharacteristicpropertytype.NilReason)

	case "AircraftCharacteristic":
		if aircraftcharacteristicpropertytype.AircraftCharacteristic != nil {
			res = __gong__marshallPointer(ident, "AircraftCharacteristic", aircraftcharacteristicpropertytype.AircraftCharacteristic.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AircraftCharacteristic", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AircraftCharacteristicPropertyType", fieldName)
	}
	return
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aircraftcharacteristictype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aircraftcharacteristictype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aircraftcharacteristictype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aircraftcharacteristictype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range aircraftcharacteristictype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_51 := range aircraftcharacteristictype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_51.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AircraftCharacteristicType", fieldName)
	}
	return
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airportheliportpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airportheliportpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airportheliportpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airportheliportpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", airportheliportpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", airportheliportpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", airportheliportpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", airportheliportpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", airportheliportpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", airportheliportpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", airportheliportpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct AirportHeliportPropertyType", fieldName)
	}
	return
}

func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := altimetersourcepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", altimetersourcepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", altimetersourcepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", altimetersourcepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", altimetersourcepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", altimetersourcepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", altimetersourcepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", altimetersourcepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", altimetersourcepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", altimetersourcepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", altimetersourcepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct AltimeterSourcePropertyType", fieldName)
	}
	return
}

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := altitudeadjustmentpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", altitudeadjustmentpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", altitudeadjustmentpropertytype.NilReason)

	case "AltitudeAdjustment":
		if altitudeadjustmentpropertytype.AltitudeAdjustment != nil {
			res = __gong__marshallPointer(ident, "AltitudeAdjustment", altitudeadjustmentpropertytype.AltitudeAdjustment.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AltitudeAdjustment", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AltitudeAdjustmentPropertyType", fieldName)
	}
	return
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := altitudeadjustmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", altitudeadjustmenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", altitudeadjustmenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", altitudeadjustmenttype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range altitudeadjustmenttype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_54 := range altitudeadjustmenttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_54.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AltitudeAdjustmentType", fieldName)
	}
	return
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := angleindicationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", angleindicationpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", angleindicationpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", angleindicationpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", angleindicationpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", angleindicationpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", angleindicationpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", angleindicationpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", angleindicationpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", angleindicationpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", angleindicationpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct AngleIndicationPropertyType", fieldName)
	}
	return
}

func (angleusepropertytype *AngleUsePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := angleusepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", angleusepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", angleusepropertytype.NilReason)

	case "AngleUse":
		if angleusepropertytype.AngleUse != nil {
			res = __gong__marshallPointer(ident, "AngleUse", angleusepropertytype.AngleUse.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AngleUse", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AngleUsePropertyType", fieldName)
	}
	return
}

func (angleusetype *AngleUseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := angleusetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", angleusetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", angleusetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", angleusetype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range angleusetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheAngleIndication":
		if angleusetype.TheAngleIndication != nil {
			res = __gong__marshallPointer(ident, "TheAngleIndication", angleusetype.TheAngleIndication.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheAngleIndication", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_8 := range angleusetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_8.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AngleUseType", fieldName)
	}
	return
}

func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := approachaltitudetablepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", approachaltitudetablepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", approachaltitudetablepropertytype.NilReason)

	case "ApproachAltitudeTable":
		if approachaltitudetablepropertytype.ApproachAltitudeTable != nil {
			res = __gong__marshallPointer(ident, "ApproachAltitudeTable", approachaltitudetablepropertytype.ApproachAltitudeTable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ApproachAltitudeTable", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ApproachAltitudeTablePropertyType", fieldName)
	}
	return
}

func (approachaltitudetabletype *ApproachAltitudeTableType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := approachaltitudetabletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", approachaltitudetabletype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", approachaltitudetabletype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", approachaltitudetabletype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range approachaltitudetabletype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_14 := range approachaltitudetabletype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_14.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ApproachAltitudeTableType", fieldName)
	}
	return
}

func (approachconditionpropertytype *ApproachConditionPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := approachconditionpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", approachconditionpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", approachconditionpropertytype.NilReason)

	case "ApproachCondition":
		if approachconditionpropertytype.ApproachCondition != nil {
			res = __gong__marshallPointer(ident, "ApproachCondition", approachconditionpropertytype.ApproachCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ApproachCondition", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ApproachConditionPropertyType", fieldName)
	}
	return
}

func (approachconditiontype *ApproachConditionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := approachconditiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", approachconditiontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", approachconditiontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", approachconditiontype.Description)

	case "MinimumSet":
		if approachconditiontype.MinimumSet != nil {
			res = __gong__marshallPointer(ident, "MinimumSet", approachconditiontype.MinimumSet.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MinimumSet", "nil")
		}
	case "CirclingRestriction":
		var sb strings.Builder
		for _, _circlingrestrictionpropertytype := range approachconditiontype.CirclingRestriction {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "CirclingRestriction", _circlingrestrictionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AircraftCategory":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range approachconditiontype.AircraftCategory {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCategory", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "LandingArea":
		var sb strings.Builder
		for _, _landingtakeoffareacollectionpropertytype := range approachconditiontype.LandingArea {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "LandingArea", _landingtakeoffareacollectionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Altimeter":
		if approachconditiontype.Altimeter != nil {
			res = __gong__marshallPointer(ident, "Altimeter", approachconditiontype.Altimeter.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Altimeter", "nil")
		}
	case "DesignSurface":
		var sb strings.Builder
		for _, _obstacleassessmentareapropertytype := range approachconditiontype.DesignSurface {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DesignSurface", _obstacleassessmentareapropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range approachconditiontype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_15 := range approachconditiontype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_15.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ApproachConditionType", fieldName)
	}
	return
}

func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := approachdistancetablepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", approachdistancetablepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", approachdistancetablepropertytype.NilReason)

	case "ApproachDistanceTable":
		if approachdistancetablepropertytype.ApproachDistanceTable != nil {
			res = __gong__marshallPointer(ident, "ApproachDistanceTable", approachdistancetablepropertytype.ApproachDistanceTable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ApproachDistanceTable", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ApproachDistanceTablePropertyType", fieldName)
	}
	return
}

func (approachdistancetabletype *ApproachDistanceTableType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := approachdistancetabletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", approachdistancetabletype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", approachdistancetabletype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", approachdistancetabletype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range approachdistancetabletype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_16 := range approachdistancetabletype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_16.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ApproachDistanceTableType", fieldName)
	}
	return
}

func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := approachtimingtablepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", approachtimingtablepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", approachtimingtablepropertytype.NilReason)

	case "ApproachTimingTable":
		if approachtimingtablepropertytype.ApproachTimingTable != nil {
			res = __gong__marshallPointer(ident, "ApproachTimingTable", approachtimingtablepropertytype.ApproachTimingTable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ApproachTimingTable", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ApproachTimingTablePropertyType", fieldName)
	}
	return
}

func (approachtimingtabletype *ApproachTimingTableType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := approachtimingtabletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", approachtimingtabletype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", approachtimingtabletype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", approachtimingtabletype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range approachtimingtabletype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_17 := range approachtimingtabletype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_17.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ApproachTimingTableType", fieldName)
	}
	return
}

func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arrivalfeederlegtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arrivalfeederlegtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", arrivalfeederlegtimeslicepropertytype.Owns)

	case "ArrivalFeederLegTimeSlice":
		if arrivalfeederlegtimeslicepropertytype.ArrivalFeederLegTimeSlice != nil {
			res = __gong__marshallPointer(ident, "ArrivalFeederLegTimeSlice", arrivalfeederlegtimeslicepropertytype.ArrivalFeederLegTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ArrivalFeederLegTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ArrivalFeederLegTimeSlicePropertyType", fieldName)
	}
	return
}

func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arrivalfeederlegtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arrivalfeederlegtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", arrivalfeederlegtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", arrivalfeederlegtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", arrivalfeederlegtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", arrivalfeederlegtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", arrivalfeederlegtimeslicetype.CorrectionNumber)

	case "StartPoint":
		if arrivalfeederlegtimeslicetype.StartPoint != nil {
			res = __gong__marshallPointer(ident, "StartPoint", arrivalfeederlegtimeslicetype.StartPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartPoint", "nil")
		}
	case "EndPoint":
		if arrivalfeederlegtimeslicetype.EndPoint != nil {
			res = __gong__marshallPointer(ident, "EndPoint", arrivalfeederlegtimeslicetype.EndPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndPoint", "nil")
		}
	case "Trajectory":
		if arrivalfeederlegtimeslicetype.Trajectory != nil {
			res = __gong__marshallPointer(ident, "Trajectory", arrivalfeederlegtimeslicetype.Trajectory.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Trajectory", "nil")
		}
	case "ArcCentre":
		if arrivalfeederlegtimeslicetype.ArcCentre != nil {
			res = __gong__marshallPointer(ident, "ArcCentre", arrivalfeederlegtimeslicetype.ArcCentre.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ArcCentre", "nil")
		}
	case "Angle":
		if arrivalfeederlegtimeslicetype.Angle != nil {
			res = __gong__marshallPointer(ident, "Angle", arrivalfeederlegtimeslicetype.Angle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Angle", "nil")
		}
	case "Distance":
		if arrivalfeederlegtimeslicetype.Distance != nil {
			res = __gong__marshallPointer(ident, "Distance", arrivalfeederlegtimeslicetype.Distance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Distance", "nil")
		}
	case "AircraftCategory":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range arrivalfeederlegtimeslicetype.AircraftCategory {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCategory", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Holding":
		if arrivalfeederlegtimeslicetype.Holding != nil {
			res = __gong__marshallPointer(ident, "Holding", arrivalfeederlegtimeslicetype.Holding.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Holding", "nil")
		}
	case "DesignSurface":
		var sb strings.Builder
		for _, _obstacleassessmentareapropertytype := range arrivalfeederlegtimeslicetype.DesignSurface {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DesignSurface", _obstacleassessmentareapropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range arrivalfeederlegtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Approach":
		if arrivalfeederlegtimeslicetype.Approach != nil {
			res = __gong__marshallPointer(ident, "Approach", arrivalfeederlegtimeslicetype.Approach.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Approach", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_37 := range arrivalfeederlegtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_37.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ArrivalFeederLegTimeSliceType", fieldName)
	}
	return
}

func (arrivalfeederlegtype *ArrivalFeederLegType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arrivalfeederlegtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arrivalfeederlegtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", arrivalfeederlegtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", arrivalfeederlegtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _arrivalfeederlegtimeslicepropertytype := range arrivalfeederlegtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _arrivalfeederlegtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ArrivalFeederLegType", fieldName)
	}
	return
}

func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arrivallegtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arrivallegtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", arrivallegtimeslicepropertytype.Owns)

	case "ArrivalLegTimeSlice":
		if arrivallegtimeslicepropertytype.ArrivalLegTimeSlice != nil {
			res = __gong__marshallPointer(ident, "ArrivalLegTimeSlice", arrivallegtimeslicepropertytype.ArrivalLegTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ArrivalLegTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ArrivalLegTimeSlicePropertyType", fieldName)
	}
	return
}

func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arrivallegtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arrivallegtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", arrivallegtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", arrivallegtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", arrivallegtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", arrivallegtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", arrivallegtimeslicetype.CorrectionNumber)

	case "StartPoint":
		if arrivallegtimeslicetype.StartPoint != nil {
			res = __gong__marshallPointer(ident, "StartPoint", arrivallegtimeslicetype.StartPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartPoint", "nil")
		}
	case "EndPoint":
		if arrivallegtimeslicetype.EndPoint != nil {
			res = __gong__marshallPointer(ident, "EndPoint", arrivallegtimeslicetype.EndPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndPoint", "nil")
		}
	case "Trajectory":
		if arrivallegtimeslicetype.Trajectory != nil {
			res = __gong__marshallPointer(ident, "Trajectory", arrivallegtimeslicetype.Trajectory.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Trajectory", "nil")
		}
	case "ArcCentre":
		if arrivallegtimeslicetype.ArcCentre != nil {
			res = __gong__marshallPointer(ident, "ArcCentre", arrivallegtimeslicetype.ArcCentre.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ArcCentre", "nil")
		}
	case "Angle":
		if arrivallegtimeslicetype.Angle != nil {
			res = __gong__marshallPointer(ident, "Angle", arrivallegtimeslicetype.Angle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Angle", "nil")
		}
	case "Distance":
		if arrivallegtimeslicetype.Distance != nil {
			res = __gong__marshallPointer(ident, "Distance", arrivallegtimeslicetype.Distance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Distance", "nil")
		}
	case "AircraftCategory":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range arrivallegtimeslicetype.AircraftCategory {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCategory", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Holding":
		if arrivallegtimeslicetype.Holding != nil {
			res = __gong__marshallPointer(ident, "Holding", arrivallegtimeslicetype.Holding.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Holding", "nil")
		}
	case "DesignSurface":
		var sb strings.Builder
		for _, _obstacleassessmentareapropertytype := range arrivallegtimeslicetype.DesignSurface {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DesignSurface", _obstacleassessmentareapropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range arrivallegtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Arrival":
		if arrivallegtimeslicetype.Arrival != nil {
			res = __gong__marshallPointer(ident, "Arrival", arrivallegtimeslicetype.Arrival.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Arrival", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_38 := range arrivallegtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_38.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ArrivalLegTimeSliceType", fieldName)
	}
	return
}

func (arrivallegtype *ArrivalLegType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arrivallegtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arrivallegtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", arrivallegtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", arrivallegtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _arrivallegtimeslicepropertytype := range arrivallegtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _arrivallegtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ArrivalLegType", fieldName)
	}
	return
}

func (circlesectorpropertytype *CircleSectorPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circlesectorpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circlesectorpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", circlesectorpropertytype.NilReason)

	case "CircleSector":
		if circlesectorpropertytype.CircleSector != nil {
			res = __gong__marshallPointer(ident, "CircleSector", circlesectorpropertytype.CircleSector.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CircleSector", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CircleSectorPropertyType", fieldName)
	}
	return
}

func (circlesectortype *CircleSectorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circlesectortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circlesectortype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", circlesectortype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", circlesectortype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range circlesectortype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_52 := range circlesectortype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_52.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CircleSectorType", fieldName)
	}
	return
}

func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circlingareatimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circlingareatimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", circlingareatimeslicepropertytype.Owns)

	case "CirclingAreaTimeSlice":
		if circlingareatimeslicepropertytype.CirclingAreaTimeSlice != nil {
			res = __gong__marshallPointer(ident, "CirclingAreaTimeSlice", circlingareatimeslicepropertytype.CirclingAreaTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CirclingAreaTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CirclingAreaTimeSlicePropertyType", fieldName)
	}
	return
}

func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circlingareatimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circlingareatimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", circlingareatimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", circlingareatimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", circlingareatimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", circlingareatimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", circlingareatimeslicetype.CorrectionNumber)

	case "Extent":
		if circlingareatimeslicetype.Extent != nil {
			res = __gong__marshallPointer(ident, "Extent", circlingareatimeslicetype.Extent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Extent", "nil")
		}
	case "Approach":
		if circlingareatimeslicetype.Approach != nil {
			res = __gong__marshallPointer(ident, "Approach", circlingareatimeslicetype.Approach.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Approach", "nil")
		}
	case "Condition":
		var sb strings.Builder
		for _, _approachconditionpropertytype := range circlingareatimeslicetype.Condition {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Condition", _approachconditionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AircraftCategory":
		if circlingareatimeslicetype.AircraftCategory != nil {
			res = __gong__marshallPointer(ident, "AircraftCategory", circlingareatimeslicetype.AircraftCategory.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AircraftCategory", "nil")
		}
	case "DesignSurface":
		var sb strings.Builder
		for _, _obstacleassessmentareapropertytype := range circlingareatimeslicetype.DesignSurface {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DesignSurface", _obstacleassessmentareapropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range circlingareatimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_24 := range circlingareatimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_24.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CirclingAreaTimeSliceType", fieldName)
	}
	return
}

func (circlingareatype *CirclingAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circlingareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circlingareatype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", circlingareatype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", circlingareatype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _circlingareatimeslicepropertytype := range circlingareatype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _circlingareatimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CirclingAreaType", fieldName)
	}
	return
}

func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circlingrestrictionpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circlingrestrictionpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", circlingrestrictionpropertytype.NilReason)

	case "CirclingRestriction":
		if circlingrestrictionpropertytype.CirclingRestriction != nil {
			res = __gong__marshallPointer(ident, "CirclingRestriction", circlingrestrictionpropertytype.CirclingRestriction.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CirclingRestriction", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CirclingRestrictionPropertyType", fieldName)
	}
	return
}

func (circlingrestrictiontype *CirclingRestrictionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circlingrestrictiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circlingrestrictiontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", circlingrestrictiontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", circlingrestrictiontype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range circlingrestrictiontype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range circlingrestrictiontype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range circlingrestrictiontype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SectorDescription":
		if circlingrestrictiontype.SectorDescription != nil {
			res = __gong__marshallPointer(ident, "SectorDescription", circlingrestrictiontype.SectorDescription.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SectorDescription", "nil")
		}
	case "RestrictionArea":
		if circlingrestrictiontype.RestrictionArea != nil {
			res = __gong__marshallPointer(ident, "RestrictionArea", circlingrestrictiontype.RestrictionArea.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RestrictionArea", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_25 := range circlingrestrictiontype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_25.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CirclingRestrictionType", fieldName)
	}
	return
}

func (curvepropertytype *CurvePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := curvepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", curvepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", curvepropertytype.NilReason)

	case "Curve":
		if curvepropertytype.Curve != nil {
			res = __gong__marshallPointer(ident, "Curve", curvepropertytype.Curve.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Curve", "nil")
		}
	case "ElevatedCurve":
		if curvepropertytype.ElevatedCurve != nil {
			res = __gong__marshallPointer(ident, "ElevatedCurve", curvepropertytype.ElevatedCurve.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ElevatedCurve", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CurvePropertyType", fieldName)
	}
	return
}

func (curvetype *CurveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := curvetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", curvetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", curvetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", curvetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", curvetype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range curvetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CurveType", fieldName)
	}
	return
}

func (dmepropertytype *DMEPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := dmepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", dmepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", dmepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", dmepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", dmepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", dmepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", dmepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", dmepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", dmepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", dmepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", dmepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct DMEPropertyType", fieldName)
	}
	return
}

func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := departurearrivalconditionpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", departurearrivalconditionpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", departurearrivalconditionpropertytype.NilReason)

	case "DepartureArrivalCondition":
		if departurearrivalconditionpropertytype.DepartureArrivalCondition != nil {
			res = __gong__marshallPointer(ident, "DepartureArrivalCondition", departurearrivalconditionpropertytype.DepartureArrivalCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DepartureArrivalCondition", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DepartureArrivalConditionPropertyType", fieldName)
	}
	return
}

func (departurearrivalconditiontype *DepartureArrivalConditionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := departurearrivalconditiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", departurearrivalconditiontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", departurearrivalconditiontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", departurearrivalconditiontype.Description)

	case "EngineType":
		if departurearrivalconditiontype.EngineType != nil {
			res = __gong__marshallPointer(ident, "EngineType", departurearrivalconditiontype.EngineType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EngineType", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range departurearrivalconditiontype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_30 := range departurearrivalconditiontype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_30.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DepartureArrivalConditionType", fieldName)
	}
	return
}

func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := departurelegtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", departurelegtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", departurelegtimeslicepropertytype.Owns)

	case "DepartureLegTimeSlice":
		if departurelegtimeslicepropertytype.DepartureLegTimeSlice != nil {
			res = __gong__marshallPointer(ident, "DepartureLegTimeSlice", departurelegtimeslicepropertytype.DepartureLegTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DepartureLegTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DepartureLegTimeSlicePropertyType", fieldName)
	}
	return
}

func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := departurelegtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", departurelegtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", departurelegtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", departurelegtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", departurelegtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", departurelegtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", departurelegtimeslicetype.CorrectionNumber)

	case "StartPoint":
		if departurelegtimeslicetype.StartPoint != nil {
			res = __gong__marshallPointer(ident, "StartPoint", departurelegtimeslicetype.StartPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartPoint", "nil")
		}
	case "EndPoint":
		if departurelegtimeslicetype.EndPoint != nil {
			res = __gong__marshallPointer(ident, "EndPoint", departurelegtimeslicetype.EndPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndPoint", "nil")
		}
	case "Trajectory":
		if departurelegtimeslicetype.Trajectory != nil {
			res = __gong__marshallPointer(ident, "Trajectory", departurelegtimeslicetype.Trajectory.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Trajectory", "nil")
		}
	case "ArcCentre":
		if departurelegtimeslicetype.ArcCentre != nil {
			res = __gong__marshallPointer(ident, "ArcCentre", departurelegtimeslicetype.ArcCentre.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ArcCentre", "nil")
		}
	case "Angle":
		if departurelegtimeslicetype.Angle != nil {
			res = __gong__marshallPointer(ident, "Angle", departurelegtimeslicetype.Angle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Angle", "nil")
		}
	case "Distance":
		if departurelegtimeslicetype.Distance != nil {
			res = __gong__marshallPointer(ident, "Distance", departurelegtimeslicetype.Distance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Distance", "nil")
		}
	case "AircraftCategory":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range departurelegtimeslicetype.AircraftCategory {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCategory", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Holding":
		if departurelegtimeslicetype.Holding != nil {
			res = __gong__marshallPointer(ident, "Holding", departurelegtimeslicetype.Holding.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Holding", "nil")
		}
	case "DesignSurface":
		var sb strings.Builder
		for _, _obstacleassessmentareapropertytype := range departurelegtimeslicetype.DesignSurface {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DesignSurface", _obstacleassessmentareapropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range departurelegtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Departure":
		if departurelegtimeslicetype.Departure != nil {
			res = __gong__marshallPointer(ident, "Departure", departurelegtimeslicetype.Departure.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Departure", "nil")
		}
	case "Condition":
		var sb strings.Builder
		for _, _departurearrivalconditionpropertytype := range departurelegtimeslicetype.Condition {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Condition", _departurearrivalconditionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_39 := range departurelegtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_39.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DepartureLegTimeSliceType", fieldName)
	}
	return
}

func (departurelegtype *DepartureLegType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := departurelegtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", departurelegtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", departurelegtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", departurelegtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _departurelegtimeslicepropertytype := range departurelegtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _departurelegtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DepartureLegType", fieldName)
	}
	return
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := designatedpointpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", designatedpointpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", designatedpointpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", designatedpointpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", designatedpointpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", designatedpointpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", designatedpointpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", designatedpointpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", designatedpointpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", designatedpointpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", designatedpointpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct DesignatedPointPropertyType", fieldName)
	}
	return
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := distanceindicationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", distanceindicationpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", distanceindicationpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", distanceindicationpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", distanceindicationpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", distanceindicationpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", distanceindicationpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", distanceindicationpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", distanceindicationpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", distanceindicationpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", distanceindicationpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct DistanceIndicationPropertyType", fieldName)
	}
	return
}

func (elevatedcurvetype *ElevatedCurveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevatedcurvetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevatedcurvetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", elevatedcurvetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", elevatedcurvetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", elevatedcurvetype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range elevatedcurvetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_1 := range elevatedcurvetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_1.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedCurveType", fieldName)
	}
	return
}

func (elevatedpointtype *ElevatedPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevatedpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevatedpointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", elevatedpointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", elevatedpointtype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", elevatedpointtype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range elevatedpointtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_2 := range elevatedpointtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_2.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedPointType", fieldName)
	}
	return
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevatedsurfacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevatedsurfacetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", elevatedsurfacetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", elevatedsurfacetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", elevatedsurfacetype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range elevatedsurfacetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_3 := range elevatedsurfacetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_3.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedSurfaceType", fieldName)
	}
	return
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := enroutesegmentpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", enroutesegmentpointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", enroutesegmentpointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", enroutesegmentpointtype.Description)

	case "FacilityMakeup":
		var sb strings.Builder
		for _, _pointreferencepropertytype := range enroutesegmentpointtype.FacilityMakeup {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FacilityMakeup", _pointreferencepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PointChoice_fixDesignatedPoint":
		if enroutesegmentpointtype.PointChoice_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_fixDesignatedPoint", enroutesegmentpointtype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_fixDesignatedPoint", "nil")
		}
	case "PointChoice_navaidSystem":
		if enroutesegmentpointtype.PointChoice_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "PointChoice_navaidSystem", enroutesegmentpointtype.PointChoice_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_navaidSystem", "nil")
		}
	case "PointChoice_position":
		if enroutesegmentpointtype.PointChoice_position != nil {
			res = __gong__marshallPointer(ident, "PointChoice_position", enroutesegmentpointtype.PointChoice_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_position", "nil")
		}
	case "PointChoice_runwayPoint":
		if enroutesegmentpointtype.PointChoice_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_runwayPoint", enroutesegmentpointtype.PointChoice_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_runwayPoint", "nil")
		}
	case "PointChoice_aimingPoint":
		if enroutesegmentpointtype.PointChoice_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_aimingPoint", enroutesegmentpointtype.PointChoice_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_aimingPoint", "nil")
		}
	case "PointChoice_airportReferencePoint":
		if enroutesegmentpointtype.PointChoice_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_airportReferencePoint", enroutesegmentpointtype.PointChoice_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_airportReferencePoint", "nil")
		}
	case "ExtendedServiceVolume":
		if enroutesegmentpointtype.ExtendedServiceVolume != nil {
			res = __gong__marshallPointer(ident, "ExtendedServiceVolume", enroutesegmentpointtype.ExtendedServiceVolume.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ExtendedServiceVolume", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range enroutesegmentpointtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_9 := range enroutesegmentpointtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_9.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct EnRouteSegmentPointType", fieldName)
	}
	return
}

func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := equipmentunavailableadjustmentcolumnpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", equipmentunavailableadjustmentcolumnpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", equipmentunavailableadjustmentcolumnpropertytype.NilReason)

	case "EquipmentUnavailableAdjustmentColumn":
		if equipmentunavailableadjustmentcolumnpropertytype.EquipmentUnavailableAdjustmentColumn != nil {
			res = __gong__marshallPointer(ident, "EquipmentUnavailableAdjustmentColumn", equipmentunavailableadjustmentcolumnpropertytype.EquipmentUnavailableAdjustmentColumn.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EquipmentUnavailableAdjustmentColumn", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct EquipmentUnavailableAdjustmentColumnPropertyType", fieldName)
	}
	return
}

func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := equipmentunavailableadjustmentcolumntype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", equipmentunavailableadjustmentcolumntype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", equipmentunavailableadjustmentcolumntype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", equipmentunavailableadjustmentcolumntype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range equipmentunavailableadjustmentcolumntype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_27 := range equipmentunavailableadjustmentcolumntype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_27.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct EquipmentUnavailableAdjustmentColumnType", fieldName)
	}
	return
}

func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := equipmentunavailableadjustmentpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", equipmentunavailableadjustmentpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", equipmentunavailableadjustmentpropertytype.NilReason)

	case "EquipmentUnavailableAdjustment":
		if equipmentunavailableadjustmentpropertytype.EquipmentUnavailableAdjustment != nil {
			res = __gong__marshallPointer(ident, "EquipmentUnavailableAdjustment", equipmentunavailableadjustmentpropertytype.EquipmentUnavailableAdjustment.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EquipmentUnavailableAdjustment", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct EquipmentUnavailableAdjustmentPropertyType", fieldName)
	}
	return
}

func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := equipmentunavailableadjustmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", equipmentunavailableadjustmenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", equipmentunavailableadjustmenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", equipmentunavailableadjustmenttype.Description)

	case "AdjustmentINOPCol":
		var sb strings.Builder
		for _, _equipmentunavailableadjustmentcolumnpropertytype := range equipmentunavailableadjustmenttype.AdjustmentINOPCol {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AdjustmentINOPCol", _equipmentunavailableadjustmentcolumnpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range equipmentunavailableadjustmenttype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_26 := range equipmentunavailableadjustmenttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_26.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct EquipmentUnavailableAdjustmentType", fieldName)
	}
	return
}

func (fasdatablockpropertytype *FASDataBlockPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := fasdatablockpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", fasdatablockpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", fasdatablockpropertytype.NilReason)

	case "FASDataBlock":
		if fasdatablockpropertytype.FASDataBlock != nil {
			res = __gong__marshallPointer(ident, "FASDataBlock", fasdatablockpropertytype.FASDataBlock.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FASDataBlock", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FASDataBlockPropertyType", fieldName)
	}
	return
}

func (fasdatablocktype *FASDataBlockType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := fasdatablocktype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", fasdatablocktype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", fasdatablocktype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", fasdatablocktype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range fasdatablocktype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_23 := range fasdatablocktype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_23.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FASDataBlockType", fieldName)
	}
	return
}

func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := finallegtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", finallegtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", finallegtimeslicepropertytype.Owns)

	case "FinalLegTimeSlice":
		if finallegtimeslicepropertytype.FinalLegTimeSlice != nil {
			res = __gong__marshallPointer(ident, "FinalLegTimeSlice", finallegtimeslicepropertytype.FinalLegTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FinalLegTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FinalLegTimeSlicePropertyType", fieldName)
	}
	return
}

func (finallegtimeslicetype *FinalLegTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := finallegtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", finallegtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", finallegtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", finallegtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", finallegtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", finallegtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", finallegtimeslicetype.CorrectionNumber)

	case "StartPoint":
		if finallegtimeslicetype.StartPoint != nil {
			res = __gong__marshallPointer(ident, "StartPoint", finallegtimeslicetype.StartPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartPoint", "nil")
		}
	case "EndPoint":
		if finallegtimeslicetype.EndPoint != nil {
			res = __gong__marshallPointer(ident, "EndPoint", finallegtimeslicetype.EndPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndPoint", "nil")
		}
	case "Trajectory":
		if finallegtimeslicetype.Trajectory != nil {
			res = __gong__marshallPointer(ident, "Trajectory", finallegtimeslicetype.Trajectory.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Trajectory", "nil")
		}
	case "ArcCentre":
		if finallegtimeslicetype.ArcCentre != nil {
			res = __gong__marshallPointer(ident, "ArcCentre", finallegtimeslicetype.ArcCentre.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ArcCentre", "nil")
		}
	case "Angle":
		if finallegtimeslicetype.Angle != nil {
			res = __gong__marshallPointer(ident, "Angle", finallegtimeslicetype.Angle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Angle", "nil")
		}
	case "Distance":
		if finallegtimeslicetype.Distance != nil {
			res = __gong__marshallPointer(ident, "Distance", finallegtimeslicetype.Distance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Distance", "nil")
		}
	case "AircraftCategory":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range finallegtimeslicetype.AircraftCategory {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCategory", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Holding":
		if finallegtimeslicetype.Holding != nil {
			res = __gong__marshallPointer(ident, "Holding", finallegtimeslicetype.Holding.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Holding", "nil")
		}
	case "DesignSurface":
		var sb strings.Builder
		for _, _obstacleassessmentareapropertytype := range finallegtimeslicetype.DesignSurface {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DesignSurface", _obstacleassessmentareapropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range finallegtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Approach":
		if finallegtimeslicetype.Approach != nil {
			res = __gong__marshallPointer(ident, "Approach", finallegtimeslicetype.Approach.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Approach", "nil")
		}
	case "Condition":
		var sb strings.Builder
		for _, _approachconditionpropertytype := range finallegtimeslicetype.Condition {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Condition", _approachconditionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "FinalPathAlignmentPoint_fixDesignatedPoint":
		if finallegtimeslicetype.FinalPathAlignmentPoint_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_fixDesignatedPoint", finallegtimeslicetype.FinalPathAlignmentPoint_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_fixDesignatedPoint", "nil")
		}
	case "FinalPathAlignmentPoint_navaidSystem":
		if finallegtimeslicetype.FinalPathAlignmentPoint_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_navaidSystem", finallegtimeslicetype.FinalPathAlignmentPoint_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_navaidSystem", "nil")
		}
	case "FinalPathAlignmentPoint_position":
		if finallegtimeslicetype.FinalPathAlignmentPoint_position != nil {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_position", finallegtimeslicetype.FinalPathAlignmentPoint_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_position", "nil")
		}
	case "FinalPathAlignmentPoint_runwayPoint":
		if finallegtimeslicetype.FinalPathAlignmentPoint_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_runwayPoint", finallegtimeslicetype.FinalPathAlignmentPoint_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_runwayPoint", "nil")
		}
	case "FinalPathAlignmentPoint_aimingPoint":
		if finallegtimeslicetype.FinalPathAlignmentPoint_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_aimingPoint", finallegtimeslicetype.FinalPathAlignmentPoint_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_aimingPoint", "nil")
		}
	case "FinalPathAlignmentPoint_airportReferencePoint":
		if finallegtimeslicetype.FinalPathAlignmentPoint_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_airportReferencePoint", finallegtimeslicetype.FinalPathAlignmentPoint_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FinalPathAlignmentPoint_airportReferencePoint", "nil")
		}
	case "VisualDescentPoint":
		if finallegtimeslicetype.VisualDescentPoint != nil {
			res = __gong__marshallPointer(ident, "VisualDescentPoint", finallegtimeslicetype.VisualDescentPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VisualDescentPoint", "nil")
		}
	case "FASData":
		if finallegtimeslicetype.FASData != nil {
			res = __gong__marshallPointer(ident, "FASData", finallegtimeslicetype.FASData.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FASData", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_40 := range finallegtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_40.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FinalLegTimeSliceType", fieldName)
	}
	return
}

func (finallegtype *FinalLegType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := finallegtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", finallegtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", finallegtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", finallegtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _finallegtimeslicepropertytype := range finallegtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _finallegtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FinalLegType", fieldName)
	}
	return
}

func (finalprofilepropertytype *FinalProfilePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := finalprofilepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", finalprofilepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", finalprofilepropertytype.NilReason)

	case "FinalProfile":
		if finalprofilepropertytype.FinalProfile != nil {
			res = __gong__marshallPointer(ident, "FinalProfile", finalprofilepropertytype.FinalProfile.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FinalProfile", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FinalProfilePropertyType", fieldName)
	}
	return
}

func (finalprofiletype *FinalProfileType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := finalprofiletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", finalprofiletype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", finalprofiletype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", finalprofiletype.Description)

	case "Altitude":
		var sb strings.Builder
		for _, _approachaltitudetablepropertytype := range finalprofiletype.Altitude {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Altitude", _approachaltitudetablepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Distance":
		var sb strings.Builder
		for _, _approachdistancetablepropertytype := range finalprofiletype.Distance {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Distance", _approachdistancetablepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Timing":
		var sb strings.Builder
		for _, _approachtimingtablepropertytype := range finalprofiletype.Timing {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Timing", _approachtimingtablepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range finalprofiletype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_18 := range finalprofiletype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_18.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FinalProfileType", fieldName)
	}
	return
}

func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingassessmenttimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingassessmenttimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", holdingassessmenttimeslicepropertytype.Owns)

	case "HoldingAssessmentTimeSlice":
		if holdingassessmenttimeslicepropertytype.HoldingAssessmentTimeSlice != nil {
			res = __gong__marshallPointer(ident, "HoldingAssessmentTimeSlice", holdingassessmenttimeslicepropertytype.HoldingAssessmentTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HoldingAssessmentTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingAssessmentTimeSlicePropertyType", fieldName)
	}
	return
}

func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingassessmenttimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingassessmenttimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", holdingassessmenttimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", holdingassessmenttimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", holdingassessmenttimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", holdingassessmenttimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", holdingassessmenttimeslicetype.CorrectionNumber)

	case "HoldingPoint":
		if holdingassessmenttimeslicetype.HoldingPoint != nil {
			res = __gong__marshallPointer(ident, "HoldingPoint", holdingassessmenttimeslicetype.HoldingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HoldingPoint", "nil")
		}
	case "UnplannedHolding":
		if holdingassessmenttimeslicetype.UnplannedHolding != nil {
			res = __gong__marshallPointer(ident, "UnplannedHolding", holdingassessmenttimeslicetype.UnplannedHolding.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "UnplannedHolding", "nil")
		}
	case "AssessedHoldingPattern":
		if holdingassessmenttimeslicetype.AssessedHoldingPattern != nil {
			res = __gong__marshallPointer(ident, "AssessedHoldingPattern", holdingassessmenttimeslicetype.AssessedHoldingPattern.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AssessedHoldingPattern", "nil")
		}
	case "ObstacleAssessment":
		var sb strings.Builder
		for _, _obstacleassessmentareapropertytype := range holdingassessmenttimeslicetype.ObstacleAssessment {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ObstacleAssessment", _obstacleassessmentareapropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range holdingassessmenttimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_55 := range holdingassessmenttimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_55.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingAssessmentTimeSliceType", fieldName)
	}
	return
}

func (holdingassessmenttype *HoldingAssessmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingassessmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingassessmenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", holdingassessmenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", holdingassessmenttype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _holdingassessmenttimeslicepropertytype := range holdingassessmenttype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _holdingassessmenttimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingAssessmentType", fieldName)
	}
	return
}

func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingpatterndistancepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingpatterndistancepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", holdingpatterndistancepropertytype.NilReason)

	case "HoldingPatternDistance":
		if holdingpatterndistancepropertytype.HoldingPatternDistance != nil {
			res = __gong__marshallPointer(ident, "HoldingPatternDistance", holdingpatterndistancepropertytype.HoldingPatternDistance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HoldingPatternDistance", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingPatternDistancePropertyType", fieldName)
	}
	return
}

func (holdingpatterndistancetype *HoldingPatternDistanceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingpatterndistancetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingpatterndistancetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", holdingpatterndistancetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", holdingpatterndistancetype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range holdingpatterndistancetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_5 := range holdingpatterndistancetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_5.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingPatternDistanceType", fieldName)
	}
	return
}

func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingpatterndurationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingpatterndurationpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", holdingpatterndurationpropertytype.NilReason)

	case "HoldingPatternDuration":
		if holdingpatterndurationpropertytype.HoldingPatternDuration != nil {
			res = __gong__marshallPointer(ident, "HoldingPatternDuration", holdingpatterndurationpropertytype.HoldingPatternDuration.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HoldingPatternDuration", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingPatternDurationPropertyType", fieldName)
	}
	return
}

func (holdingpatterndurationtype *HoldingPatternDurationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingpatterndurationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingpatterndurationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", holdingpatterndurationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", holdingpatterndurationtype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range holdingpatterndurationtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_6 := range holdingpatterndurationtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_6.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingPatternDurationType", fieldName)
	}
	return
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingpatternpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingpatternpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", holdingpatternpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", holdingpatternpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", holdingpatternpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", holdingpatternpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", holdingpatternpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", holdingpatternpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", holdingpatternpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", holdingpatternpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", holdingpatternpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingPatternPropertyType", fieldName)
	}
	return
}

func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingpatterntimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingpatterntimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", holdingpatterntimeslicepropertytype.Owns)

	case "HoldingPatternTimeSlice":
		if holdingpatterntimeslicepropertytype.HoldingPatternTimeSlice != nil {
			res = __gong__marshallPointer(ident, "HoldingPatternTimeSlice", holdingpatterntimeslicepropertytype.HoldingPatternTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HoldingPatternTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingPatternTimeSlicePropertyType", fieldName)
	}
	return
}

func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingpatterntimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingpatterntimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", holdingpatterntimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", holdingpatterntimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", holdingpatterntimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", holdingpatterntimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", holdingpatterntimeslicetype.CorrectionNumber)

	case "OutboundLegSpan_endPoint":
		if holdingpatterntimeslicetype.OutboundLegSpan_endPoint != nil {
			res = __gong__marshallPointer(ident, "OutboundLegSpan_endPoint", holdingpatterntimeslicetype.OutboundLegSpan_endPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OutboundLegSpan_endPoint", "nil")
		}
	case "OutboundLegSpan_endDistance":
		if holdingpatterntimeslicetype.OutboundLegSpan_endDistance != nil {
			res = __gong__marshallPointer(ident, "OutboundLegSpan_endDistance", holdingpatterntimeslicetype.OutboundLegSpan_endDistance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OutboundLegSpan_endDistance", "nil")
		}
	case "OutboundLegSpan_endTime":
		if holdingpatterntimeslicetype.OutboundLegSpan_endTime != nil {
			res = __gong__marshallPointer(ident, "OutboundLegSpan_endTime", holdingpatterntimeslicetype.OutboundLegSpan_endTime.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OutboundLegSpan_endTime", "nil")
		}
	case "HoldingPoint":
		if holdingpatterntimeslicetype.HoldingPoint != nil {
			res = __gong__marshallPointer(ident, "HoldingPoint", holdingpatterntimeslicetype.HoldingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HoldingPoint", "nil")
		}
	case "Extent":
		if holdingpatterntimeslicetype.Extent != nil {
			res = __gong__marshallPointer(ident, "Extent", holdingpatterntimeslicetype.Extent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Extent", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range holdingpatterntimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_4 := range holdingpatterntimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_4.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingPatternTimeSliceType", fieldName)
	}
	return
}

func (holdingpatterntype *HoldingPatternType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingpatterntype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingpatterntype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", holdingpatterntype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", holdingpatterntype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _holdingpatterntimeslicepropertytype := range holdingpatterntype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _holdingpatterntimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingPatternType", fieldName)
	}
	return
}

func (holdingusepropertytype *HoldingUsePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingusepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingusepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", holdingusepropertytype.NilReason)

	case "HoldingUse":
		if holdingusepropertytype.HoldingUse != nil {
			res = __gong__marshallPointer(ident, "HoldingUse", holdingusepropertytype.HoldingUse.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HoldingUse", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingUsePropertyType", fieldName)
	}
	return
}

func (holdingusetype *HoldingUseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := holdingusetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", holdingusetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", holdingusetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", holdingusetype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range holdingusetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheHoldingPattern":
		if holdingusetype.TheHoldingPattern != nil {
			res = __gong__marshallPointer(ident, "TheHoldingPattern", holdingusetype.TheHoldingPattern.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheHoldingPattern", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_41 := range holdingusetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_41.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct HoldingUseType", fieldName)
	}
	return
}

func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := initiallegtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", initiallegtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", initiallegtimeslicepropertytype.Owns)

	case "InitialLegTimeSlice":
		if initiallegtimeslicepropertytype.InitialLegTimeSlice != nil {
			res = __gong__marshallPointer(ident, "InitialLegTimeSlice", initiallegtimeslicepropertytype.InitialLegTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "InitialLegTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct InitialLegTimeSlicePropertyType", fieldName)
	}
	return
}

func (initiallegtimeslicetype *InitialLegTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := initiallegtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", initiallegtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", initiallegtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", initiallegtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", initiallegtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", initiallegtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", initiallegtimeslicetype.CorrectionNumber)

	case "StartPoint":
		if initiallegtimeslicetype.StartPoint != nil {
			res = __gong__marshallPointer(ident, "StartPoint", initiallegtimeslicetype.StartPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartPoint", "nil")
		}
	case "EndPoint":
		if initiallegtimeslicetype.EndPoint != nil {
			res = __gong__marshallPointer(ident, "EndPoint", initiallegtimeslicetype.EndPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndPoint", "nil")
		}
	case "Trajectory":
		if initiallegtimeslicetype.Trajectory != nil {
			res = __gong__marshallPointer(ident, "Trajectory", initiallegtimeslicetype.Trajectory.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Trajectory", "nil")
		}
	case "ArcCentre":
		if initiallegtimeslicetype.ArcCentre != nil {
			res = __gong__marshallPointer(ident, "ArcCentre", initiallegtimeslicetype.ArcCentre.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ArcCentre", "nil")
		}
	case "Angle":
		if initiallegtimeslicetype.Angle != nil {
			res = __gong__marshallPointer(ident, "Angle", initiallegtimeslicetype.Angle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Angle", "nil")
		}
	case "Distance":
		if initiallegtimeslicetype.Distance != nil {
			res = __gong__marshallPointer(ident, "Distance", initiallegtimeslicetype.Distance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Distance", "nil")
		}
	case "AircraftCategory":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range initiallegtimeslicetype.AircraftCategory {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCategory", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Holding":
		if initiallegtimeslicetype.Holding != nil {
			res = __gong__marshallPointer(ident, "Holding", initiallegtimeslicetype.Holding.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Holding", "nil")
		}
	case "DesignSurface":
		var sb strings.Builder
		for _, _obstacleassessmentareapropertytype := range initiallegtimeslicetype.DesignSurface {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DesignSurface", _obstacleassessmentareapropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range initiallegtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Approach":
		if initiallegtimeslicetype.Approach != nil {
			res = __gong__marshallPointer(ident, "Approach", initiallegtimeslicetype.Approach.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Approach", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_42 := range initiallegtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_42.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct InitialLegTimeSliceType", fieldName)
	}
	return
}

func (initiallegtype *InitialLegType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := initiallegtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", initiallegtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", initiallegtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", initiallegtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _initiallegtimeslicepropertytype := range initiallegtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _initiallegtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct InitialLegType", fieldName)
	}
	return
}

func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := instrumentapproachprocedurepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", instrumentapproachprocedurepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", instrumentapproachprocedurepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", instrumentapproachprocedurepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", instrumentapproachprocedurepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", instrumentapproachprocedurepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", instrumentapproachprocedurepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", instrumentapproachprocedurepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", instrumentapproachprocedurepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", instrumentapproachprocedurepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", instrumentapproachprocedurepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct InstrumentApproachProcedurePropertyType", fieldName)
	}
	return
}

func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := instrumentapproachproceduretimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", instrumentapproachproceduretimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", instrumentapproachproceduretimeslicepropertytype.Owns)

	case "InstrumentApproachProcedureTimeSlice":
		if instrumentapproachproceduretimeslicepropertytype.InstrumentApproachProcedureTimeSlice != nil {
			res = __gong__marshallPointer(ident, "InstrumentApproachProcedureTimeSlice", instrumentapproachproceduretimeslicepropertytype.InstrumentApproachProcedureTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "InstrumentApproachProcedureTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct InstrumentApproachProcedureTimeSlicePropertyType", fieldName)
	}
	return
}

func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := instrumentapproachproceduretimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", instrumentapproachproceduretimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", instrumentapproachproceduretimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", instrumentapproachproceduretimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", instrumentapproachproceduretimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", instrumentapproachproceduretimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", instrumentapproachproceduretimeslicetype.CorrectionNumber)

	case "Availability":
		var sb strings.Builder
		for _, _procedureavailabilitypropertytype := range instrumentapproachproceduretimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _procedureavailabilitypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AirportHeliport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range instrumentapproachproceduretimeslicetype.AirportHeliport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AirportHeliport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AircraftCharacteristic":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range instrumentapproachproceduretimeslicetype.AircraftCharacteristic {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCharacteristic", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "FlightTransition":
		var sb strings.Builder
		for _, _proceduretransitionpropertytype := range instrumentapproachproceduretimeslicetype.FlightTransition {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FlightTransition", _proceduretransitionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GuidanceFacility_specialNavigationSystem":
		if instrumentapproachproceduretimeslicetype.GuidanceFacility_specialNavigationSystem != nil {
			res = __gong__marshallPointer(ident, "GuidanceFacility_specialNavigationSystem", instrumentapproachproceduretimeslicetype.GuidanceFacility_specialNavigationSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GuidanceFacility_specialNavigationSystem", "nil")
		}
	case "GuidanceFacility_radar":
		if instrumentapproachproceduretimeslicetype.GuidanceFacility_radar != nil {
			res = __gong__marshallPointer(ident, "GuidanceFacility_radar", instrumentapproachproceduretimeslicetype.GuidanceFacility_radar.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GuidanceFacility_radar", "nil")
		}
	case "GuidanceFacility_navaid":
		if instrumentapproachproceduretimeslicetype.GuidanceFacility_navaid != nil {
			res = __gong__marshallPointer(ident, "GuidanceFacility_navaid", instrumentapproachproceduretimeslicetype.GuidanceFacility_navaid.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GuidanceFacility_navaid", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range instrumentapproachproceduretimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SafeAltitude":
		if instrumentapproachproceduretimeslicetype.SafeAltitude != nil {
			res = __gong__marshallPointer(ident, "SafeAltitude", instrumentapproachproceduretimeslicetype.SafeAltitude.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SafeAltitude", "nil")
		}
	case "Landing":
		if instrumentapproachproceduretimeslicetype.Landing != nil {
			res = __gong__marshallPointer(ident, "Landing", instrumentapproachproceduretimeslicetype.Landing.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Landing", "nil")
		}
	case "MissedInstruction":
		var sb strings.Builder
		for _, _missedapproachgrouppropertytype := range instrumentapproachproceduretimeslicetype.MissedInstruction {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "MissedInstruction", _missedapproachgrouppropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "FinalProfile":
		if instrumentapproachproceduretimeslicetype.FinalProfile != nil {
			res = __gong__marshallPointer(ident, "FinalProfile", instrumentapproachproceduretimeslicetype.FinalProfile.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FinalProfile", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_19 := range instrumentapproachproceduretimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_19.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct InstrumentApproachProcedureTimeSliceType", fieldName)
	}
	return
}

func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := instrumentapproachproceduretype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", instrumentapproachproceduretype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", instrumentapproachproceduretype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", instrumentapproachproceduretype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _instrumentapproachproceduretimeslicepropertytype := range instrumentapproachproceduretype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _instrumentapproachproceduretimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct InstrumentApproachProcedureType", fieldName)
	}
	return
}

func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := intermediatelegtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", intermediatelegtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", intermediatelegtimeslicepropertytype.Owns)

	case "IntermediateLegTimeSlice":
		if intermediatelegtimeslicepropertytype.IntermediateLegTimeSlice != nil {
			res = __gong__marshallPointer(ident, "IntermediateLegTimeSlice", intermediatelegtimeslicepropertytype.IntermediateLegTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IntermediateLegTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct IntermediateLegTimeSlicePropertyType", fieldName)
	}
	return
}

func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := intermediatelegtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", intermediatelegtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", intermediatelegtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", intermediatelegtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", intermediatelegtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", intermediatelegtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", intermediatelegtimeslicetype.CorrectionNumber)

	case "StartPoint":
		if intermediatelegtimeslicetype.StartPoint != nil {
			res = __gong__marshallPointer(ident, "StartPoint", intermediatelegtimeslicetype.StartPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartPoint", "nil")
		}
	case "EndPoint":
		if intermediatelegtimeslicetype.EndPoint != nil {
			res = __gong__marshallPointer(ident, "EndPoint", intermediatelegtimeslicetype.EndPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndPoint", "nil")
		}
	case "Trajectory":
		if intermediatelegtimeslicetype.Trajectory != nil {
			res = __gong__marshallPointer(ident, "Trajectory", intermediatelegtimeslicetype.Trajectory.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Trajectory", "nil")
		}
	case "ArcCentre":
		if intermediatelegtimeslicetype.ArcCentre != nil {
			res = __gong__marshallPointer(ident, "ArcCentre", intermediatelegtimeslicetype.ArcCentre.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ArcCentre", "nil")
		}
	case "Angle":
		if intermediatelegtimeslicetype.Angle != nil {
			res = __gong__marshallPointer(ident, "Angle", intermediatelegtimeslicetype.Angle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Angle", "nil")
		}
	case "Distance":
		if intermediatelegtimeslicetype.Distance != nil {
			res = __gong__marshallPointer(ident, "Distance", intermediatelegtimeslicetype.Distance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Distance", "nil")
		}
	case "AircraftCategory":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range intermediatelegtimeslicetype.AircraftCategory {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCategory", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Holding":
		if intermediatelegtimeslicetype.Holding != nil {
			res = __gong__marshallPointer(ident, "Holding", intermediatelegtimeslicetype.Holding.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Holding", "nil")
		}
	case "DesignSurface":
		var sb strings.Builder
		for _, _obstacleassessmentareapropertytype := range intermediatelegtimeslicetype.DesignSurface {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DesignSurface", _obstacleassessmentareapropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range intermediatelegtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Approach":
		if intermediatelegtimeslicetype.Approach != nil {
			res = __gong__marshallPointer(ident, "Approach", intermediatelegtimeslicetype.Approach.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Approach", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_43 := range intermediatelegtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_43.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct IntermediateLegTimeSliceType", fieldName)
	}
	return
}

func (intermediatelegtype *IntermediateLegType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := intermediatelegtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", intermediatelegtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", intermediatelegtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", intermediatelegtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _intermediatelegtimeslicepropertytype := range intermediatelegtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _intermediatelegtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct IntermediateLegType", fieldName)
	}
	return
}

func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := landingtakeoffareacollectionpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", landingtakeoffareacollectionpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", landingtakeoffareacollectionpropertytype.NilReason)

	case "LandingTakeoffAreaCollection":
		if landingtakeoffareacollectionpropertytype.LandingTakeoffAreaCollection != nil {
			res = __gong__marshallPointer(ident, "LandingTakeoffAreaCollection", landingtakeoffareacollectionpropertytype.LandingTakeoffAreaCollection.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LandingTakeoffAreaCollection", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LandingTakeoffAreaCollectionPropertyType", fieldName)
	}
	return
}

func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := landingtakeoffareacollectiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", landingtakeoffareacollectiontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", landingtakeoffareacollectiontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", landingtakeoffareacollectiontype.Description)

	case "Runway":
		var sb strings.Builder
		for _, _runwaydirectionpropertytype := range landingtakeoffareacollectiontype.Runway {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Runway", _runwaydirectionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TLOF":
		var sb strings.Builder
		for _, _touchdownliftoffpropertytype := range landingtakeoffareacollectiontype.TLOF {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TLOF", _touchdownliftoffpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range landingtakeoffareacollectiontype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_44 := range landingtakeoffareacollectiontype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_44.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LandingTakeoffAreaCollectionType", fieldName)
	}
	return
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := linguisticnotepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", linguisticnotepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", linguisticnotepropertytype.NilReason)

	case "LinguisticNote":
		if linguisticnotepropertytype.LinguisticNote != nil {
			res = __gong__marshallPointer(ident, "LinguisticNote", linguisticnotepropertytype.LinguisticNote.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LinguisticNote", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LinguisticNotePropertyType", fieldName)
	}
	return
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := linguisticnotetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", linguisticnotetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", linguisticnotetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", linguisticnotetype.Description)

	case "Extension":
		var sb strings.Builder
		for _, _a_extension_12 := range linguisticnotetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_12.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LinguisticNoteType", fieldName)
	}
	return
}

func (minimapropertytype *MinimaPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := minimapropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", minimapropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", minimapropertytype.NilReason)

	case "Minima":
		if minimapropertytype.Minima != nil {
			res = __gong__marshallPointer(ident, "Minima", minimapropertytype.Minima.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Minima", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MinimaPropertyType", fieldName)
	}
	return
}

func (minimatype *MinimaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := minimatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", minimatype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", minimatype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", minimatype.Description)

	case "AdjustmentINOP":
		var sb strings.Builder
		for _, _equipmentunavailableadjustmentpropertytype := range minimatype.AdjustmentINOP {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AdjustmentINOP", _equipmentunavailableadjustmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range minimatype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_28 := range minimatype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_28.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MinimaType", fieldName)
	}
	return
}

func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := missedapproachgrouppropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", missedapproachgrouppropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", missedapproachgrouppropertytype.NilReason)

	case "MissedApproachGroup":
		if missedapproachgrouppropertytype.MissedApproachGroup != nil {
			res = __gong__marshallPointer(ident, "MissedApproachGroup", missedapproachgrouppropertytype.MissedApproachGroup.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MissedApproachGroup", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MissedApproachGroupPropertyType", fieldName)
	}
	return
}

func (missedapproachgrouptype *MissedApproachGroupType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := missedapproachgrouptype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", missedapproachgrouptype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", missedapproachgrouptype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", missedapproachgrouptype.Description)

	case "Altimeter":
		var sb strings.Builder
		for _, _altimetersourcepropertytype := range missedapproachgrouptype.Altimeter {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Altimeter", _altimetersourcepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range missedapproachgrouptype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_20 := range missedapproachgrouptype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_20.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MissedApproachGroupType", fieldName)
	}
	return
}

func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := missedapproachlegtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", missedapproachlegtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", missedapproachlegtimeslicepropertytype.Owns)

	case "MissedApproachLegTimeSlice":
		if missedapproachlegtimeslicepropertytype.MissedApproachLegTimeSlice != nil {
			res = __gong__marshallPointer(ident, "MissedApproachLegTimeSlice", missedapproachlegtimeslicepropertytype.MissedApproachLegTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MissedApproachLegTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MissedApproachLegTimeSlicePropertyType", fieldName)
	}
	return
}

func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := missedapproachlegtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", missedapproachlegtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", missedapproachlegtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", missedapproachlegtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", missedapproachlegtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", missedapproachlegtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", missedapproachlegtimeslicetype.CorrectionNumber)

	case "StartPoint":
		if missedapproachlegtimeslicetype.StartPoint != nil {
			res = __gong__marshallPointer(ident, "StartPoint", missedapproachlegtimeslicetype.StartPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartPoint", "nil")
		}
	case "EndPoint":
		if missedapproachlegtimeslicetype.EndPoint != nil {
			res = __gong__marshallPointer(ident, "EndPoint", missedapproachlegtimeslicetype.EndPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndPoint", "nil")
		}
	case "Trajectory":
		if missedapproachlegtimeslicetype.Trajectory != nil {
			res = __gong__marshallPointer(ident, "Trajectory", missedapproachlegtimeslicetype.Trajectory.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Trajectory", "nil")
		}
	case "ArcCentre":
		if missedapproachlegtimeslicetype.ArcCentre != nil {
			res = __gong__marshallPointer(ident, "ArcCentre", missedapproachlegtimeslicetype.ArcCentre.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ArcCentre", "nil")
		}
	case "Angle":
		if missedapproachlegtimeslicetype.Angle != nil {
			res = __gong__marshallPointer(ident, "Angle", missedapproachlegtimeslicetype.Angle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Angle", "nil")
		}
	case "Distance":
		if missedapproachlegtimeslicetype.Distance != nil {
			res = __gong__marshallPointer(ident, "Distance", missedapproachlegtimeslicetype.Distance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Distance", "nil")
		}
	case "AircraftCategory":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range missedapproachlegtimeslicetype.AircraftCategory {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCategory", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Holding":
		if missedapproachlegtimeslicetype.Holding != nil {
			res = __gong__marshallPointer(ident, "Holding", missedapproachlegtimeslicetype.Holding.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Holding", "nil")
		}
	case "DesignSurface":
		var sb strings.Builder
		for _, _obstacleassessmentareapropertytype := range missedapproachlegtimeslicetype.DesignSurface {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DesignSurface", _obstacleassessmentareapropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range missedapproachlegtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Approach":
		if missedapproachlegtimeslicetype.Approach != nil {
			res = __gong__marshallPointer(ident, "Approach", missedapproachlegtimeslicetype.Approach.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Approach", "nil")
		}
	case "Condition":
		var sb strings.Builder
		for _, _approachconditionpropertytype := range missedapproachlegtimeslicetype.Condition {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Condition", _approachconditionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_45 := range missedapproachlegtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_45.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MissedApproachLegTimeSliceType", fieldName)
	}
	return
}

func (missedapproachlegtype *MissedApproachLegType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := missedapproachlegtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", missedapproachlegtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", missedapproachlegtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", missedapproachlegtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _missedapproachlegtimeslicepropertytype := range missedapproachlegtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _missedapproachlegtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MissedApproachLegType", fieldName)
	}
	return
}

func (navaidpropertytype *NavaidPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", navaidpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", navaidpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", navaidpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", navaidpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", navaidpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", navaidpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", navaidpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", navaidpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", navaidpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidPropertyType", fieldName)
	}
	return
}

func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navigationarearestrictiontimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navigationarearestrictiontimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", navigationarearestrictiontimeslicepropertytype.Owns)

	case "NavigationAreaRestrictionTimeSlice":
		if navigationarearestrictiontimeslicepropertytype.NavigationAreaRestrictionTimeSlice != nil {
			res = __gong__marshallPointer(ident, "NavigationAreaRestrictionTimeSlice", navigationarearestrictiontimeslicepropertytype.NavigationAreaRestrictionTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "NavigationAreaRestrictionTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NavigationAreaRestrictionTimeSlicePropertyType", fieldName)
	}
	return
}

func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navigationarearestrictiontimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navigationarearestrictiontimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", navigationarearestrictiontimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", navigationarearestrictiontimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", navigationarearestrictiontimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", navigationarearestrictiontimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", navigationarearestrictiontimeslicetype.CorrectionNumber)

	case "Procedure":
		var sb strings.Builder
		for _, _procedurepropertytype := range navigationarearestrictiontimeslicetype.Procedure {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Procedure", _procedurepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DesignSurface":
		if navigationarearestrictiontimeslicetype.DesignSurface != nil {
			res = __gong__marshallPointer(ident, "DesignSurface", navigationarearestrictiontimeslicetype.DesignSurface.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DesignSurface", "nil")
		}
	case "SectorDefinition":
		if navigationarearestrictiontimeslicetype.SectorDefinition != nil {
			res = __gong__marshallPointer(ident, "SectorDefinition", navigationarearestrictiontimeslicetype.SectorDefinition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SectorDefinition", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range navigationarearestrictiontimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_46 := range navigationarearestrictiontimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_46.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NavigationAreaRestrictionTimeSliceType", fieldName)
	}
	return
}

func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navigationarearestrictiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navigationarearestrictiontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", navigationarearestrictiontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", navigationarearestrictiontype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _navigationarearestrictiontimeslicepropertytype := range navigationarearestrictiontype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _navigationarearestrictiontimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NavigationAreaRestrictionType", fieldName)
	}
	return
}

func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navigationareasectorpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navigationareasectorpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", navigationareasectorpropertytype.NilReason)

	case "NavigationAreaSector":
		if navigationareasectorpropertytype.NavigationAreaSector != nil {
			res = __gong__marshallPointer(ident, "NavigationAreaSector", navigationareasectorpropertytype.NavigationAreaSector.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "NavigationAreaSector", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NavigationAreaSectorPropertyType", fieldName)
	}
	return
}

func (navigationareasectortype *NavigationAreaSectorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navigationareasectortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navigationareasectortype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", navigationareasectortype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", navigationareasectortype.Description)

	case "SectorDefinition":
		if navigationareasectortype.SectorDefinition != nil {
			res = __gong__marshallPointer(ident, "SectorDefinition", navigationareasectortype.SectorDefinition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SectorDefinition", "nil")
		}
	case "SignificantObstacle":
		var sb strings.Builder
		for _, _obstructionpropertytype := range navigationareasectortype.SignificantObstacle {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SignificantObstacle", _obstructionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extent":
		if navigationareasectortype.Extent != nil {
			res = __gong__marshallPointer(ident, "Extent", navigationareasectortype.Extent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Extent", "nil")
		}
	case "SectorCriteria":
		var sb strings.Builder
		for _, _sectordesignpropertytype := range navigationareasectortype.SectorCriteria {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SectorCriteria", _sectordesignpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range navigationareasectortype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_32 := range navigationareasectortype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_32.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NavigationAreaSectorType", fieldName)
	}
	return
}

func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navigationareatimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navigationareatimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", navigationareatimeslicepropertytype.Owns)

	case "NavigationAreaTimeSlice":
		if navigationareatimeslicepropertytype.NavigationAreaTimeSlice != nil {
			res = __gong__marshallPointer(ident, "NavigationAreaTimeSlice", navigationareatimeslicepropertytype.NavigationAreaTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "NavigationAreaTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NavigationAreaTimeSlicePropertyType", fieldName)
	}
	return
}

func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navigationareatimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navigationareatimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", navigationareatimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", navigationareatimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", navigationareatimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", navigationareatimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", navigationareatimeslicetype.CorrectionNumber)

	case "Departure":
		if navigationareatimeslicetype.Departure != nil {
			res = __gong__marshallPointer(ident, "Departure", navigationareatimeslicetype.Departure.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Departure", "nil")
		}
	case "Sector":
		var sb strings.Builder
		for _, _navigationareasectorpropertytype := range navigationareatimeslicetype.Sector {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sector", _navigationareasectorpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "CentrePoint_fixDesignatedPoint":
		if navigationareatimeslicetype.CentrePoint_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_fixDesignatedPoint", navigationareatimeslicetype.CentrePoint_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_fixDesignatedPoint", "nil")
		}
	case "CentrePoint_navaidSystem":
		if navigationareatimeslicetype.CentrePoint_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_navaidSystem", navigationareatimeslicetype.CentrePoint_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_navaidSystem", "nil")
		}
	case "CentrePoint_position":
		if navigationareatimeslicetype.CentrePoint_position != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_position", navigationareatimeslicetype.CentrePoint_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_position", "nil")
		}
	case "CentrePoint_runwayPoint":
		if navigationareatimeslicetype.CentrePoint_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_runwayPoint", navigationareatimeslicetype.CentrePoint_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_runwayPoint", "nil")
		}
	case "CentrePoint_aimingPoint":
		if navigationareatimeslicetype.CentrePoint_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_aimingPoint", navigationareatimeslicetype.CentrePoint_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_aimingPoint", "nil")
		}
	case "CentrePoint_airportReferencePoint":
		if navigationareatimeslicetype.CentrePoint_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_airportReferencePoint", navigationareatimeslicetype.CentrePoint_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_airportReferencePoint", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range navigationareatimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_31 := range navigationareatimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_31.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NavigationAreaTimeSliceType", fieldName)
	}
	return
}

func (navigationareatype *NavigationAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navigationareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navigationareatype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", navigationareatype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", navigationareatype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _navigationareatimeslicepropertytype := range navigationareatype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _navigationareatimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NavigationAreaType", fieldName)
	}
	return
}

func (notepropertytype *NotePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := notepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", notepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", notepropertytype.NilReason)

	case "Note":
		if notepropertytype.Note != nil {
			res = __gong__marshallPointer(ident, "Note", notepropertytype.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NotePropertyType", fieldName)
	}
	return
}

func (notetype *NoteType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := notetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", notetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", notetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", notetype.Description)

	case "TranslatedNote":
		var sb strings.Builder
		for _, _linguisticnotepropertytype := range notetype.TranslatedNote {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TranslatedNote", _linguisticnotepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_13 := range notetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_13.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteType", fieldName)
	}
	return
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := obstacleassessmentareapropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", obstacleassessmentareapropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", obstacleassessmentareapropertytype.NilReason)

	case "ObstacleAssessmentArea":
		if obstacleassessmentareapropertytype.ObstacleAssessmentArea != nil {
			res = __gong__marshallPointer(ident, "ObstacleAssessmentArea", obstacleassessmentareapropertytype.ObstacleAssessmentArea.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ObstacleAssessmentArea", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ObstacleAssessmentAreaPropertyType", fieldName)
	}
	return
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := obstacleassessmentareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", obstacleassessmentareatype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", obstacleassessmentareatype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", obstacleassessmentareatype.Description)

	case "AircraftCategory":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range obstacleassessmentareatype.AircraftCategory {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCategory", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SignificantObstacle":
		var sb strings.Builder
		for _, _obstructionpropertytype := range obstacleassessmentareatype.SignificantObstacle {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SignificantObstacle", _obstructionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Surface":
		if obstacleassessmentareatype.Surface != nil {
			res = __gong__marshallPointer(ident, "Surface", obstacleassessmentareatype.Surface.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Surface", "nil")
		}
	case "StartingCurve":
		if obstacleassessmentareatype.StartingCurve != nil {
			res = __gong__marshallPointer(ident, "StartingCurve", obstacleassessmentareatype.StartingCurve.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartingCurve", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range obstacleassessmentareatype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_56 := range obstacleassessmentareatype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_56.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ObstacleAssessmentAreaType", fieldName)
	}
	return
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := obstacleplacementpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", obstacleplacementpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", obstacleplacementpropertytype.NilReason)

	case "ObstaclePlacement":
		if obstacleplacementpropertytype.ObstaclePlacement != nil {
			res = __gong__marshallPointer(ident, "ObstaclePlacement", obstacleplacementpropertytype.ObstaclePlacement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ObstaclePlacement", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ObstaclePlacementPropertyType", fieldName)
	}
	return
}

func (obstacleplacementtype *ObstaclePlacementType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := obstacleplacementtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", obstacleplacementtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", obstacleplacementtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", obstacleplacementtype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range obstacleplacementtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_57 := range obstacleplacementtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_57.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ObstaclePlacementType", fieldName)
	}
	return
}

func (obstructionpropertytype *ObstructionPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := obstructionpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", obstructionpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", obstructionpropertytype.NilReason)

	case "Obstruction":
		if obstructionpropertytype.Obstruction != nil {
			res = __gong__marshallPointer(ident, "Obstruction", obstructionpropertytype.Obstruction.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Obstruction", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ObstructionPropertyType", fieldName)
	}
	return
}

func (obstructiontype *ObstructionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := obstructiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", obstructiontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", obstructiontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", obstructiontype.Description)

	case "TheVerticalStructure":
		if obstructiontype.TheVerticalStructure != nil {
			res = __gong__marshallPointer(ident, "TheVerticalStructure", obstructiontype.TheVerticalStructure.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheVerticalStructure", "nil")
		}
	case "Adjustment":
		var sb strings.Builder
		for _, _altitudeadjustmentpropertytype := range obstructiontype.Adjustment {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Adjustment", _altitudeadjustmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ObstaclePlacement":
		var sb strings.Builder
		for _, _obstacleplacementpropertytype := range obstructiontype.ObstaclePlacement {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ObstaclePlacement", _obstacleplacementpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range obstructiontype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_58 := range obstructiontype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_58.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ObstructionType", fieldName)
	}
	return
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := organisationauthoritypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", organisationauthoritypropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", organisationauthoritypropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", organisationauthoritypropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", organisationauthoritypropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", organisationauthoritypropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", organisationauthoritypropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", organisationauthoritypropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", organisationauthoritypropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", organisationauthoritypropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", organisationauthoritypropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct OrganisationAuthorityPropertyType", fieldName)
	}
	return
}

func (pointpropertytype *PointPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pointpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pointpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", pointpropertytype.NilReason)

	case "Point":
		if pointpropertytype.Point != nil {
			res = __gong__marshallPointer(ident, "Point", pointpropertytype.Point.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Point", "nil")
		}
	case "ElevatedPoint":
		if pointpropertytype.ElevatedPoint != nil {
			res = __gong__marshallPointer(ident, "ElevatedPoint", pointpropertytype.ElevatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ElevatedPoint", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PointPropertyType", fieldName)
	}
	return
}

func (pointreferencepropertytype *PointReferencePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pointreferencepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pointreferencepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", pointreferencepropertytype.NilReason)

	case "PointReference":
		if pointreferencepropertytype.PointReference != nil {
			res = __gong__marshallPointer(ident, "PointReference", pointreferencepropertytype.PointReference.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointReference", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PointReferencePropertyType", fieldName)
	}
	return
}

func (pointreferencetype *PointReferenceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pointreferencetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pointreferencetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", pointreferencetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", pointreferencetype.Description)

	case "Point":
		if pointreferencetype.Point != nil {
			res = __gong__marshallPointer(ident, "Point", pointreferencetype.Point.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Point", "nil")
		}
	case "FacilityAngle":
		var sb strings.Builder
		for _, _angleusepropertytype := range pointreferencetype.FacilityAngle {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FacilityAngle", _angleusepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "FacilityDistance":
		var sb strings.Builder
		for _, _distanceindicationpropertytype := range pointreferencetype.FacilityDistance {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FacilityDistance", _distanceindicationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "FixToleranceArea":
		if pointreferencetype.FixToleranceArea != nil {
			res = __gong__marshallPointer(ident, "FixToleranceArea", pointreferencetype.FixToleranceArea.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FixToleranceArea", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range pointreferencetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_10 := range pointreferencetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_10.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PointReferenceType", fieldName)
	}
	return
}

func (pointtype *PointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", pointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", pointtype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", pointtype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range pointtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PointType", fieldName)
	}
	return
}

func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := procedureavailabilitypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", procedureavailabilitypropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", procedureavailabilitypropertytype.NilReason)

	case "ProcedureAvailability":
		if procedureavailabilitypropertytype.ProcedureAvailability != nil {
			res = __gong__marshallPointer(ident, "ProcedureAvailability", procedureavailabilitypropertytype.ProcedureAvailability.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ProcedureAvailability", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcedureAvailabilityPropertyType", fieldName)
	}
	return
}

func (procedureavailabilitytype *ProcedureAvailabilityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := procedureavailabilitytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", procedureavailabilitytype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", procedureavailabilitytype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", procedureavailabilitytype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range procedureavailabilitytype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range procedureavailabilitytype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range procedureavailabilitytype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_50 := range procedureavailabilitytype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_50.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcedureAvailabilityType", fieldName)
	}
	return
}

func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := proceduredmetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", proceduredmetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", proceduredmetimeslicepropertytype.Owns)

	case "ProcedureDMETimeSlice":
		if proceduredmetimeslicepropertytype.ProcedureDMETimeSlice != nil {
			res = __gong__marshallPointer(ident, "ProcedureDMETimeSlice", proceduredmetimeslicepropertytype.ProcedureDMETimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ProcedureDMETimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcedureDMETimeSlicePropertyType", fieldName)
	}
	return
}

func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := proceduredmetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", proceduredmetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", proceduredmetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", proceduredmetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", proceduredmetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", proceduredmetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", proceduredmetimeslicetype.CorrectionNumber)

	case "DME":
		if proceduredmetimeslicetype.DME != nil {
			res = __gong__marshallPointer(ident, "DME", proceduredmetimeslicetype.DME.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DME", "nil")
		}
	case "SegmentLeg":
		if proceduredmetimeslicetype.SegmentLeg != nil {
			res = __gong__marshallPointer(ident, "SegmentLeg", proceduredmetimeslicetype.SegmentLeg.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SegmentLeg", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range proceduredmetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_47 := range proceduredmetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_47.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcedureDMETimeSliceType", fieldName)
	}
	return
}

func (proceduredmetype *ProcedureDMEType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := proceduredmetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", proceduredmetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", proceduredmetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", proceduredmetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _proceduredmetimeslicepropertytype := range proceduredmetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _proceduredmetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcedureDMEType", fieldName)
	}
	return
}

func (procedurepropertytype *ProcedurePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := procedurepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", procedurepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", procedurepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", procedurepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", procedurepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", procedurepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", procedurepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", procedurepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", procedurepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", procedurepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", procedurepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct ProcedurePropertyType", fieldName)
	}
	return
}

func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := proceduretransitionlegpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", proceduretransitionlegpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", proceduretransitionlegpropertytype.NilReason)

	case "ProcedureTransitionLeg":
		if proceduretransitionlegpropertytype.ProcedureTransitionLeg != nil {
			res = __gong__marshallPointer(ident, "ProcedureTransitionLeg", proceduretransitionlegpropertytype.ProcedureTransitionLeg.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ProcedureTransitionLeg", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcedureTransitionLegPropertyType", fieldName)
	}
	return
}

func (proceduretransitionlegtype *ProcedureTransitionLegType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := proceduretransitionlegtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", proceduretransitionlegtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", proceduretransitionlegtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", proceduretransitionlegtype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range proceduretransitionlegtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheSegmentLeg":
		if proceduretransitionlegtype.TheSegmentLeg != nil {
			res = __gong__marshallPointer(ident, "TheSegmentLeg", proceduretransitionlegtype.TheSegmentLeg.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheSegmentLeg", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_49 := range proceduretransitionlegtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_49.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcedureTransitionLegType", fieldName)
	}
	return
}

func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := proceduretransitionpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", proceduretransitionpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", proceduretransitionpropertytype.NilReason)

	case "ProcedureTransition":
		if proceduretransitionpropertytype.ProcedureTransition != nil {
			res = __gong__marshallPointer(ident, "ProcedureTransition", proceduretransitionpropertytype.ProcedureTransition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ProcedureTransition", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcedureTransitionPropertyType", fieldName)
	}
	return
}

func (proceduretransitiontype *ProcedureTransitionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := proceduretransitiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", proceduretransitiontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", proceduretransitiontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", proceduretransitiontype.Description)

	case "DepartureRunwayTransition":
		if proceduretransitiontype.DepartureRunwayTransition != nil {
			res = __gong__marshallPointer(ident, "DepartureRunwayTransition", proceduretransitiontype.DepartureRunwayTransition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DepartureRunwayTransition", "nil")
		}
	case "Trajectory":
		if proceduretransitiontype.Trajectory != nil {
			res = __gong__marshallPointer(ident, "Trajectory", proceduretransitiontype.Trajectory.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Trajectory", "nil")
		}
	case "TransitionLeg":
		var sb strings.Builder
		for _, _proceduretransitionlegpropertytype := range proceduretransitiontype.TransitionLeg {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TransitionLeg", _proceduretransitionlegpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range proceduretransitiontype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_48 := range proceduretransitiontype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_48.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcedureTransitionType", fieldName)
	}
	return
}

func (radarsystempropertytype *RadarSystemPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radarsystempropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radarsystempropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", radarsystempropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", radarsystempropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", radarsystempropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", radarsystempropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", radarsystempropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", radarsystempropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", radarsystempropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", radarsystempropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", radarsystempropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct RadarSystemPropertyType", fieldName)
	}
	return
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radiofrequencyareapropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radiofrequencyareapropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", radiofrequencyareapropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", radiofrequencyareapropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", radiofrequencyareapropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", radiofrequencyareapropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", radiofrequencyareapropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", radiofrequencyareapropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", radiofrequencyareapropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", radiofrequencyareapropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", radiofrequencyareapropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct RadioFrequencyAreaPropertyType", fieldName)
	}
	return
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := runwaycentrelinepointpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", runwaycentrelinepointpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", runwaycentrelinepointpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", runwaycentrelinepointpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", runwaycentrelinepointpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", runwaycentrelinepointpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", runwaycentrelinepointpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", runwaycentrelinepointpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", runwaycentrelinepointpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", runwaycentrelinepointpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", runwaycentrelinepointpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct RunwayCentrelinePointPropertyType", fieldName)
	}
	return
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := runwaydirectionpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", runwaydirectionpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", runwaydirectionpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", runwaydirectionpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", runwaydirectionpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", runwaydirectionpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", runwaydirectionpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", runwaydirectionpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", runwaydirectionpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", runwaydirectionpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", runwaydirectionpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct RunwayDirectionPropertyType", fieldName)
	}
	return
}

func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := safealtitudeareapropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", safealtitudeareapropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", safealtitudeareapropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", safealtitudeareapropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", safealtitudeareapropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", safealtitudeareapropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", safealtitudeareapropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", safealtitudeareapropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", safealtitudeareapropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", safealtitudeareapropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", safealtitudeareapropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct SafeAltitudeAreaPropertyType", fieldName)
	}
	return
}

func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := safealtitudeareasectorpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", safealtitudeareasectorpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", safealtitudeareasectorpropertytype.NilReason)

	case "SafeAltitudeAreaSector":
		if safealtitudeareasectorpropertytype.SafeAltitudeAreaSector != nil {
			res = __gong__marshallPointer(ident, "SafeAltitudeAreaSector", safealtitudeareasectorpropertytype.SafeAltitudeAreaSector.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SafeAltitudeAreaSector", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SafeAltitudeAreaSectorPropertyType", fieldName)
	}
	return
}

func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := safealtitudeareasectortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", safealtitudeareasectortype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", safealtitudeareasectortype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", safealtitudeareasectortype.Description)

	case "Extent":
		if safealtitudeareasectortype.Extent != nil {
			res = __gong__marshallPointer(ident, "Extent", safealtitudeareasectortype.Extent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Extent", "nil")
		}
	case "SignificantObstacle":
		var sb strings.Builder
		for _, _obstructionpropertytype := range safealtitudeareasectortype.SignificantObstacle {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SignificantObstacle", _obstructionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SectorDefinition":
		if safealtitudeareasectortype.SectorDefinition != nil {
			res = __gong__marshallPointer(ident, "SectorDefinition", safealtitudeareasectortype.SectorDefinition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SectorDefinition", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range safealtitudeareasectortype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_36 := range safealtitudeareasectortype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_36.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SafeAltitudeAreaSectorType", fieldName)
	}
	return
}

func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := safealtitudeareatimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", safealtitudeareatimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", safealtitudeareatimeslicepropertytype.Owns)

	case "SafeAltitudeAreaTimeSlice":
		if safealtitudeareatimeslicepropertytype.SafeAltitudeAreaTimeSlice != nil {
			res = __gong__marshallPointer(ident, "SafeAltitudeAreaTimeSlice", safealtitudeareatimeslicepropertytype.SafeAltitudeAreaTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SafeAltitudeAreaTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SafeAltitudeAreaTimeSlicePropertyType", fieldName)
	}
	return
}

func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := safealtitudeareatimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", safealtitudeareatimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", safealtitudeareatimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", safealtitudeareatimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", safealtitudeareatimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", safealtitudeareatimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", safealtitudeareatimeslicetype.CorrectionNumber)

	case "CentrePoint_fixDesignatedPoint":
		if safealtitudeareatimeslicetype.CentrePoint_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_fixDesignatedPoint", safealtitudeareatimeslicetype.CentrePoint_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_fixDesignatedPoint", "nil")
		}
	case "CentrePoint_navaidSystem":
		if safealtitudeareatimeslicetype.CentrePoint_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_navaidSystem", safealtitudeareatimeslicetype.CentrePoint_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_navaidSystem", "nil")
		}
	case "CentrePoint_position":
		if safealtitudeareatimeslicetype.CentrePoint_position != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_position", safealtitudeareatimeslicetype.CentrePoint_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_position", "nil")
		}
	case "CentrePoint_runwayPoint":
		if safealtitudeareatimeslicetype.CentrePoint_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_runwayPoint", safealtitudeareatimeslicetype.CentrePoint_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_runwayPoint", "nil")
		}
	case "CentrePoint_aimingPoint":
		if safealtitudeareatimeslicetype.CentrePoint_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_aimingPoint", safealtitudeareatimeslicetype.CentrePoint_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_aimingPoint", "nil")
		}
	case "CentrePoint_airportReferencePoint":
		if safealtitudeareatimeslicetype.CentrePoint_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "CentrePoint_airportReferencePoint", safealtitudeareatimeslicetype.CentrePoint_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CentrePoint_airportReferencePoint", "nil")
		}
	case "Sector":
		var sb strings.Builder
		for _, _safealtitudeareasectorpropertytype := range safealtitudeareatimeslicetype.Sector {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sector", _safealtitudeareasectorpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Location":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range safealtitudeareatimeslicetype.Location {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Location", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range safealtitudeareatimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_35 := range safealtitudeareatimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_35.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SafeAltitudeAreaTimeSliceType", fieldName)
	}
	return
}

func (safealtitudeareatype *SafeAltitudeAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := safealtitudeareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", safealtitudeareatype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", safealtitudeareatype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", safealtitudeareatype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _safealtitudeareatimeslicepropertytype := range safealtitudeareatype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _safealtitudeareatimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SafeAltitudeAreaType", fieldName)
	}
	return
}

func (sectordesignpropertytype *SectorDesignPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := sectordesignpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", sectordesignpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", sectordesignpropertytype.NilReason)

	case "SectorDesign":
		if sectordesignpropertytype.SectorDesign != nil {
			res = __gong__marshallPointer(ident, "SectorDesign", sectordesignpropertytype.SectorDesign.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SectorDesign", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SectorDesignPropertyType", fieldName)
	}
	return
}

func (sectordesigntype *SectorDesignType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := sectordesigntype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", sectordesigntype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", sectordesigntype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", sectordesigntype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range sectordesigntype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_33 := range sectordesigntype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_33.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SectorDesignType", fieldName)
	}
	return
}

func (segmentlegpropertytype *SegmentLegPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := segmentlegpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", segmentlegpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", segmentlegpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", segmentlegpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", segmentlegpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", segmentlegpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", segmentlegpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", segmentlegpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", segmentlegpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", segmentlegpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", segmentlegpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct SegmentLegPropertyType", fieldName)
	}
	return
}

func (segmentpointpropertytype *SegmentPointPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := segmentpointpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", segmentpointpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", segmentpointpropertytype.NilReason)

	case "AerialRefuellingPoint":
		if segmentpointpropertytype.AerialRefuellingPoint != nil {
			res = __gong__marshallPointer(ident, "AerialRefuellingPoint", segmentpointpropertytype.AerialRefuellingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AerialRefuellingPoint", "nil")
		}
	case "EnRouteSegmentPoint":
		if segmentpointpropertytype.EnRouteSegmentPoint != nil {
			res = __gong__marshallPointer(ident, "EnRouteSegmentPoint", segmentpointpropertytype.EnRouteSegmentPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EnRouteSegmentPoint", "nil")
		}
	case "TerminalSegmentPoint":
		if segmentpointpropertytype.TerminalSegmentPoint != nil {
			res = __gong__marshallPointer(ident, "TerminalSegmentPoint", segmentpointpropertytype.TerminalSegmentPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TerminalSegmentPoint", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SegmentPointPropertyType", fieldName)
	}
	return
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialnavigationsystempropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialnavigationsystempropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", specialnavigationsystempropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", specialnavigationsystempropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", specialnavigationsystempropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", specialnavigationsystempropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", specialnavigationsystempropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", specialnavigationsystempropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", specialnavigationsystempropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", specialnavigationsystempropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", specialnavigationsystempropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialNavigationSystemPropertyType", fieldName)
	}
	return
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardinstrumentarrivalpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardinstrumentarrivalpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", standardinstrumentarrivalpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", standardinstrumentarrivalpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", standardinstrumentarrivalpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", standardinstrumentarrivalpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", standardinstrumentarrivalpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", standardinstrumentarrivalpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", standardinstrumentarrivalpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", standardinstrumentarrivalpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", standardinstrumentarrivalpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct StandardInstrumentArrivalPropertyType", fieldName)
	}
	return
}

func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardinstrumentarrivaltimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardinstrumentarrivaltimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", standardinstrumentarrivaltimeslicepropertytype.Owns)

	case "StandardInstrumentArrivalTimeSlice":
		if standardinstrumentarrivaltimeslicepropertytype.StandardInstrumentArrivalTimeSlice != nil {
			res = __gong__marshallPointer(ident, "StandardInstrumentArrivalTimeSlice", standardinstrumentarrivaltimeslicepropertytype.StandardInstrumentArrivalTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StandardInstrumentArrivalTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardInstrumentArrivalTimeSlicePropertyType", fieldName)
	}
	return
}

func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardinstrumentarrivaltimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardinstrumentarrivaltimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardinstrumentarrivaltimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardinstrumentarrivaltimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", standardinstrumentarrivaltimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", standardinstrumentarrivaltimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", standardinstrumentarrivaltimeslicetype.CorrectionNumber)

	case "Availability":
		var sb strings.Builder
		for _, _procedureavailabilitypropertytype := range standardinstrumentarrivaltimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _procedureavailabilitypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AirportHeliport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range standardinstrumentarrivaltimeslicetype.AirportHeliport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AirportHeliport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AircraftCharacteristic":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range standardinstrumentarrivaltimeslicetype.AircraftCharacteristic {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCharacteristic", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "FlightTransition":
		var sb strings.Builder
		for _, _proceduretransitionpropertytype := range standardinstrumentarrivaltimeslicetype.FlightTransition {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FlightTransition", _proceduretransitionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GuidanceFacility_specialNavigationSystem":
		if standardinstrumentarrivaltimeslicetype.GuidanceFacility_specialNavigationSystem != nil {
			res = __gong__marshallPointer(ident, "GuidanceFacility_specialNavigationSystem", standardinstrumentarrivaltimeslicetype.GuidanceFacility_specialNavigationSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GuidanceFacility_specialNavigationSystem", "nil")
		}
	case "GuidanceFacility_radar":
		if standardinstrumentarrivaltimeslicetype.GuidanceFacility_radar != nil {
			res = __gong__marshallPointer(ident, "GuidanceFacility_radar", standardinstrumentarrivaltimeslicetype.GuidanceFacility_radar.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GuidanceFacility_radar", "nil")
		}
	case "GuidanceFacility_navaid":
		if standardinstrumentarrivaltimeslicetype.GuidanceFacility_navaid != nil {
			res = __gong__marshallPointer(ident, "GuidanceFacility_navaid", standardinstrumentarrivaltimeslicetype.GuidanceFacility_navaid.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GuidanceFacility_navaid", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range standardinstrumentarrivaltimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SafeAltitude":
		if standardinstrumentarrivaltimeslicetype.SafeAltitude != nil {
			res = __gong__marshallPointer(ident, "SafeAltitude", standardinstrumentarrivaltimeslicetype.SafeAltitude.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SafeAltitude", "nil")
		}
	case "Arrival":
		if standardinstrumentarrivaltimeslicetype.Arrival != nil {
			res = __gong__marshallPointer(ident, "Arrival", standardinstrumentarrivaltimeslicetype.Arrival.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Arrival", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_29 := range standardinstrumentarrivaltimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_29.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardInstrumentArrivalTimeSliceType", fieldName)
	}
	return
}

func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardinstrumentarrivaltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardinstrumentarrivaltype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardinstrumentarrivaltype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardinstrumentarrivaltype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _standardinstrumentarrivaltimeslicepropertytype := range standardinstrumentarrivaltype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _standardinstrumentarrivaltimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardInstrumentArrivalType", fieldName)
	}
	return
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardinstrumentdeparturepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardinstrumentdeparturepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", standardinstrumentdeparturepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", standardinstrumentdeparturepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", standardinstrumentdeparturepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", standardinstrumentdeparturepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", standardinstrumentdeparturepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", standardinstrumentdeparturepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", standardinstrumentdeparturepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", standardinstrumentdeparturepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", standardinstrumentdeparturepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct StandardInstrumentDeparturePropertyType", fieldName)
	}
	return
}

func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardinstrumentdeparturetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardinstrumentdeparturetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", standardinstrumentdeparturetimeslicepropertytype.Owns)

	case "StandardInstrumentDepartureTimeSlice":
		if standardinstrumentdeparturetimeslicepropertytype.StandardInstrumentDepartureTimeSlice != nil {
			res = __gong__marshallPointer(ident, "StandardInstrumentDepartureTimeSlice", standardinstrumentdeparturetimeslicepropertytype.StandardInstrumentDepartureTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StandardInstrumentDepartureTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardInstrumentDepartureTimeSlicePropertyType", fieldName)
	}
	return
}

func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardinstrumentdeparturetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardinstrumentdeparturetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardinstrumentdeparturetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardinstrumentdeparturetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", standardinstrumentdeparturetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", standardinstrumentdeparturetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", standardinstrumentdeparturetimeslicetype.CorrectionNumber)

	case "Availability":
		var sb strings.Builder
		for _, _procedureavailabilitypropertytype := range standardinstrumentdeparturetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _procedureavailabilitypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AirportHeliport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range standardinstrumentdeparturetimeslicetype.AirportHeliport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AirportHeliport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AircraftCharacteristic":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range standardinstrumentdeparturetimeslicetype.AircraftCharacteristic {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AircraftCharacteristic", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "FlightTransition":
		var sb strings.Builder
		for _, _proceduretransitionpropertytype := range standardinstrumentdeparturetimeslicetype.FlightTransition {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FlightTransition", _proceduretransitionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GuidanceFacility_specialNavigationSystem":
		if standardinstrumentdeparturetimeslicetype.GuidanceFacility_specialNavigationSystem != nil {
			res = __gong__marshallPointer(ident, "GuidanceFacility_specialNavigationSystem", standardinstrumentdeparturetimeslicetype.GuidanceFacility_specialNavigationSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GuidanceFacility_specialNavigationSystem", "nil")
		}
	case "GuidanceFacility_radar":
		if standardinstrumentdeparturetimeslicetype.GuidanceFacility_radar != nil {
			res = __gong__marshallPointer(ident, "GuidanceFacility_radar", standardinstrumentdeparturetimeslicetype.GuidanceFacility_radar.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GuidanceFacility_radar", "nil")
		}
	case "GuidanceFacility_navaid":
		if standardinstrumentdeparturetimeslicetype.GuidanceFacility_navaid != nil {
			res = __gong__marshallPointer(ident, "GuidanceFacility_navaid", standardinstrumentdeparturetimeslicetype.GuidanceFacility_navaid.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GuidanceFacility_navaid", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range standardinstrumentdeparturetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SafeAltitude":
		if standardinstrumentdeparturetimeslicetype.SafeAltitude != nil {
			res = __gong__marshallPointer(ident, "SafeAltitude", standardinstrumentdeparturetimeslicetype.SafeAltitude.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SafeAltitude", "nil")
		}
	case "Takeoff":
		if standardinstrumentdeparturetimeslicetype.Takeoff != nil {
			res = __gong__marshallPointer(ident, "Takeoff", standardinstrumentdeparturetimeslicetype.Takeoff.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Takeoff", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_34 := range standardinstrumentdeparturetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_34.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardInstrumentDepartureTimeSliceType", fieldName)
	}
	return
}

func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardinstrumentdeparturetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardinstrumentdeparturetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardinstrumentdeparturetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardinstrumentdeparturetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _standardinstrumentdeparturetimeslicepropertytype := range standardinstrumentdeparturetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _standardinstrumentdeparturetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardInstrumentDepartureType", fieldName)
	}
	return
}

func (surfacepropertytype *SurfacePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := surfacepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", surfacepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", surfacepropertytype.NilReason)

	case "Surface":
		if surfacepropertytype.Surface != nil {
			res = __gong__marshallPointer(ident, "Surface", surfacepropertytype.Surface.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Surface", "nil")
		}
	case "ElevatedSurface":
		if surfacepropertytype.ElevatedSurface != nil {
			res = __gong__marshallPointer(ident, "ElevatedSurface", surfacepropertytype.ElevatedSurface.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ElevatedSurface", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SurfacePropertyType", fieldName)
	}
	return
}

func (surfacetype *SurfaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := surfacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", surfacetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", surfacetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", surfacetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", surfacetype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range surfacetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SurfaceType", fieldName)
	}
	return
}

func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := terminalarrivalareasectorpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", terminalarrivalareasectorpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", terminalarrivalareasectorpropertytype.NilReason)

	case "TerminalArrivalAreaSector":
		if terminalarrivalareasectorpropertytype.TerminalArrivalAreaSector != nil {
			res = __gong__marshallPointer(ident, "TerminalArrivalAreaSector", terminalarrivalareasectorpropertytype.TerminalArrivalAreaSector.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TerminalArrivalAreaSector", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TerminalArrivalAreaSectorPropertyType", fieldName)
	}
	return
}

func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := terminalarrivalareasectortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", terminalarrivalareasectortype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", terminalarrivalareasectortype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", terminalarrivalareasectortype.Description)

	case "SectorDefinition":
		if terminalarrivalareasectortype.SectorDefinition != nil {
			res = __gong__marshallPointer(ident, "SectorDefinition", terminalarrivalareasectortype.SectorDefinition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SectorDefinition", "nil")
		}
	case "Extent":
		if terminalarrivalareasectortype.Extent != nil {
			res = __gong__marshallPointer(ident, "Extent", terminalarrivalareasectortype.Extent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Extent", "nil")
		}
	case "SignificantObstacle":
		var sb strings.Builder
		for _, _obstructionpropertytype := range terminalarrivalareasectortype.SignificantObstacle {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SignificantObstacle", _obstructionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range terminalarrivalareasectortype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_22 := range terminalarrivalareasectortype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_22.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TerminalArrivalAreaSectorType", fieldName)
	}
	return
}

func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := terminalarrivalareatimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", terminalarrivalareatimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", terminalarrivalareatimeslicepropertytype.Owns)

	case "TerminalArrivalAreaTimeSlice":
		if terminalarrivalareatimeslicepropertytype.TerminalArrivalAreaTimeSlice != nil {
			res = __gong__marshallPointer(ident, "TerminalArrivalAreaTimeSlice", terminalarrivalareatimeslicepropertytype.TerminalArrivalAreaTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TerminalArrivalAreaTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TerminalArrivalAreaTimeSlicePropertyType", fieldName)
	}
	return
}

func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := terminalarrivalareatimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", terminalarrivalareatimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", terminalarrivalareatimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", terminalarrivalareatimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", terminalarrivalareatimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", terminalarrivalareatimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", terminalarrivalareatimeslicetype.CorrectionNumber)

	case "IF_fixDesignatedPoint":
		if terminalarrivalareatimeslicetype.IF_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "IF_fixDesignatedPoint", terminalarrivalareatimeslicetype.IF_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IF_fixDesignatedPoint", "nil")
		}
	case "IF_navaidSystem":
		if terminalarrivalareatimeslicetype.IF_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "IF_navaidSystem", terminalarrivalareatimeslicetype.IF_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IF_navaidSystem", "nil")
		}
	case "IF_position":
		if terminalarrivalareatimeslicetype.IF_position != nil {
			res = __gong__marshallPointer(ident, "IF_position", terminalarrivalareatimeslicetype.IF_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IF_position", "nil")
		}
	case "IF_runwayPoint":
		if terminalarrivalareatimeslicetype.IF_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "IF_runwayPoint", terminalarrivalareatimeslicetype.IF_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IF_runwayPoint", "nil")
		}
	case "IF_aimingPoint":
		if terminalarrivalareatimeslicetype.IF_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "IF_aimingPoint", terminalarrivalareatimeslicetype.IF_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IF_aimingPoint", "nil")
		}
	case "IF_airportReferencePoint":
		if terminalarrivalareatimeslicetype.IF_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "IF_airportReferencePoint", terminalarrivalareatimeslicetype.IF_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IF_airportReferencePoint", "nil")
		}
	case "IAF_fixDesignatedPoint":
		if terminalarrivalareatimeslicetype.IAF_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "IAF_fixDesignatedPoint", terminalarrivalareatimeslicetype.IAF_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IAF_fixDesignatedPoint", "nil")
		}
	case "IAF_navaidSystem":
		if terminalarrivalareatimeslicetype.IAF_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "IAF_navaidSystem", terminalarrivalareatimeslicetype.IAF_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IAF_navaidSystem", "nil")
		}
	case "IAF_position":
		if terminalarrivalareatimeslicetype.IAF_position != nil {
			res = __gong__marshallPointer(ident, "IAF_position", terminalarrivalareatimeslicetype.IAF_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IAF_position", "nil")
		}
	case "IAF_runwayPoint":
		if terminalarrivalareatimeslicetype.IAF_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "IAF_runwayPoint", terminalarrivalareatimeslicetype.IAF_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IAF_runwayPoint", "nil")
		}
	case "IAF_aimingPoint":
		if terminalarrivalareatimeslicetype.IAF_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "IAF_aimingPoint", terminalarrivalareatimeslicetype.IAF_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IAF_aimingPoint", "nil")
		}
	case "IAF_airportReferencePoint":
		if terminalarrivalareatimeslicetype.IAF_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "IAF_airportReferencePoint", terminalarrivalareatimeslicetype.IAF_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IAF_airportReferencePoint", "nil")
		}
	case "Buffer":
		if terminalarrivalareatimeslicetype.Buffer != nil {
			res = __gong__marshallPointer(ident, "Buffer", terminalarrivalareatimeslicetype.Buffer.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Buffer", "nil")
		}
	case "Sector":
		var sb strings.Builder
		for _, _terminalarrivalareasectorpropertytype := range terminalarrivalareatimeslicetype.Sector {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sector", _terminalarrivalareasectorpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ApproachRNAV":
		if terminalarrivalareatimeslicetype.ApproachRNAV != nil {
			res = __gong__marshallPointer(ident, "ApproachRNAV", terminalarrivalareatimeslicetype.ApproachRNAV.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ApproachRNAV", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range terminalarrivalareatimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_21 := range terminalarrivalareatimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_21.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TerminalArrivalAreaTimeSliceType", fieldName)
	}
	return
}

func (terminalarrivalareatype *TerminalArrivalAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := terminalarrivalareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", terminalarrivalareatype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", terminalarrivalareatype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", terminalarrivalareatype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _terminalarrivalareatimeslicepropertytype := range terminalarrivalareatype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _terminalarrivalareatimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TerminalArrivalAreaType", fieldName)
	}
	return
}

func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := terminalsegmentpointpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", terminalsegmentpointpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", terminalsegmentpointpropertytype.NilReason)

	case "TerminalSegmentPoint":
		if terminalsegmentpointpropertytype.TerminalSegmentPoint != nil {
			res = __gong__marshallPointer(ident, "TerminalSegmentPoint", terminalsegmentpointpropertytype.TerminalSegmentPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TerminalSegmentPoint", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TerminalSegmentPointPropertyType", fieldName)
	}
	return
}

func (terminalsegmentpointtype *TerminalSegmentPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := terminalsegmentpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", terminalsegmentpointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", terminalsegmentpointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", terminalsegmentpointtype.Description)

	case "FacilityMakeup":
		var sb strings.Builder
		for _, _pointreferencepropertytype := range terminalsegmentpointtype.FacilityMakeup {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FacilityMakeup", _pointreferencepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PointChoice_fixDesignatedPoint":
		if terminalsegmentpointtype.PointChoice_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_fixDesignatedPoint", terminalsegmentpointtype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_fixDesignatedPoint", "nil")
		}
	case "PointChoice_navaidSystem":
		if terminalsegmentpointtype.PointChoice_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "PointChoice_navaidSystem", terminalsegmentpointtype.PointChoice_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_navaidSystem", "nil")
		}
	case "PointChoice_position":
		if terminalsegmentpointtype.PointChoice_position != nil {
			res = __gong__marshallPointer(ident, "PointChoice_position", terminalsegmentpointtype.PointChoice_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_position", "nil")
		}
	case "PointChoice_runwayPoint":
		if terminalsegmentpointtype.PointChoice_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_runwayPoint", terminalsegmentpointtype.PointChoice_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_runwayPoint", "nil")
		}
	case "PointChoice_aimingPoint":
		if terminalsegmentpointtype.PointChoice_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_aimingPoint", terminalsegmentpointtype.PointChoice_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_aimingPoint", "nil")
		}
	case "PointChoice_airportReferencePoint":
		if terminalsegmentpointtype.PointChoice_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_airportReferencePoint", terminalsegmentpointtype.PointChoice_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_airportReferencePoint", "nil")
		}
	case "ExtendedServiceVolume":
		if terminalsegmentpointtype.ExtendedServiceVolume != nil {
			res = __gong__marshallPointer(ident, "ExtendedServiceVolume", terminalsegmentpointtype.ExtendedServiceVolume.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ExtendedServiceVolume", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range terminalsegmentpointtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_11 := range terminalsegmentpointtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_11.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TerminalSegmentPointType", fieldName)
	}
	return
}

func (timesheetpropertytype *TimesheetPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := timesheetpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", timesheetpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", timesheetpropertytype.NilReason)

	case "Timesheet":
		if timesheetpropertytype.Timesheet != nil {
			res = __gong__marshallPointer(ident, "Timesheet", timesheetpropertytype.Timesheet.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Timesheet", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TimesheetPropertyType", fieldName)
	}
	return
}

func (timesheettype *TimesheetType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := timesheettype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", timesheettype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", timesheettype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", timesheettype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range timesheettype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_53 := range timesheettype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_53.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TimesheetType", fieldName)
	}
	return
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := touchdownliftoffpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", touchdownliftoffpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", touchdownliftoffpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", touchdownliftoffpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", touchdownliftoffpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", touchdownliftoffpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", touchdownliftoffpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", touchdownliftoffpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", touchdownliftoffpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", touchdownliftoffpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", touchdownliftoffpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct TouchDownLiftOffPropertyType", fieldName)
	}
	return
}

func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unplannedholdingpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unplannedholdingpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", unplannedholdingpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", unplannedholdingpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", unplannedholdingpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", unplannedholdingpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", unplannedholdingpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", unplannedholdingpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", unplannedholdingpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", unplannedholdingpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", unplannedholdingpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct UnplannedHoldingPropertyType", fieldName)
	}
	return
}

func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unplannedholdingtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unplannedholdingtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", unplannedholdingtimeslicepropertytype.Owns)

	case "UnplannedHoldingTimeSlice":
		if unplannedholdingtimeslicepropertytype.UnplannedHoldingTimeSlice != nil {
			res = __gong__marshallPointer(ident, "UnplannedHoldingTimeSlice", unplannedholdingtimeslicepropertytype.UnplannedHoldingTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "UnplannedHoldingTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct UnplannedHoldingTimeSlicePropertyType", fieldName)
	}
	return
}

func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unplannedholdingtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unplannedholdingtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", unplannedholdingtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", unplannedholdingtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", unplannedholdingtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", unplannedholdingtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", unplannedholdingtimeslicetype.CorrectionNumber)

	case "HoldingPoint":
		if unplannedholdingtimeslicetype.HoldingPoint != nil {
			res = __gong__marshallPointer(ident, "HoldingPoint", unplannedholdingtimeslicetype.HoldingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HoldingPoint", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range unplannedholdingtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_7 := range unplannedholdingtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_7.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct UnplannedHoldingTimeSliceType", fieldName)
	}
	return
}

func (unplannedholdingtype *UnplannedHoldingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unplannedholdingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unplannedholdingtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", unplannedholdingtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", unplannedholdingtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _unplannedholdingtimeslicepropertytype := range unplannedholdingtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _unplannedholdingtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct UnplannedHoldingType", fieldName)
	}
	return
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := verticalstructurepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", verticalstructurepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", verticalstructurepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", verticalstructurepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", verticalstructurepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", verticalstructurepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", verticalstructurepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", verticalstructurepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", verticalstructurepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", verticalstructurepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", verticalstructurepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalStructurePropertyType", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (a_extension *A_extension) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_1 *A_extension_1) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_1.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_1.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_10 *A_extension_10) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_10.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_10.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_11 *A_extension_11) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_11.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_11.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_12 *A_extension_12) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_12.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_12.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_13 *A_extension_13) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_13.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_13.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_14 *A_extension_14) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_14.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_14.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_15 *A_extension_15) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_15.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_15.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_16 *A_extension_16) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_16.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_16.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_17 *A_extension_17) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_17.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_17.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_18 *A_extension_18) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_18.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_18.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_19 *A_extension_19) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_19.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_19.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_2 *A_extension_2) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_2.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_2.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_20 *A_extension_20) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_20.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_20.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_21 *A_extension_21) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_21.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_21.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_22 *A_extension_22) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_22.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_22.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_23 *A_extension_23) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_23.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_23.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_24 *A_extension_24) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_24.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_24.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_25 *A_extension_25) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_25.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_25.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_26 *A_extension_26) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_26.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_26.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_27 *A_extension_27) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_27.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_27.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_28 *A_extension_28) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_28.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_28.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_29 *A_extension_29) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_29.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_29.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_3 *A_extension_3) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_3.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_3.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_30 *A_extension_30) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_30.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_30.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_31 *A_extension_31) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_31.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_31.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_32 *A_extension_32) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_32.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_32.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_33 *A_extension_33) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_33.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_33.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_34 *A_extension_34) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_34.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_34.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_35 *A_extension_35) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_35.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_35.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_36 *A_extension_36) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_36.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_36.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_37 *A_extension_37) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_37.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_37.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_38 *A_extension_38) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_38.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_38.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_39 *A_extension_39) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_39.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_39.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_4 *A_extension_4) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_4.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_4.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_40 *A_extension_40) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_40.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_40.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_41 *A_extension_41) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_41.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_41.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_42 *A_extension_42) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_42.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_42.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_43 *A_extension_43) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_43.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_43.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_44 *A_extension_44) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_44.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_44.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_45 *A_extension_45) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_45.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_45.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_46 *A_extension_46) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_46.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_46.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_47 *A_extension_47) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_47.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_47.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_48 *A_extension_48) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_48.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_48.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_49 *A_extension_49) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_49.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_49.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_5 *A_extension_5) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_5.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_5.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_50 *A_extension_50) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_50.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_50.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_51 *A_extension_51) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_51.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_51.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_52 *A_extension_52) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_52.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_52.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_53 *A_extension_53) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_53.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_53.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_54 *A_extension_54) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_54.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_54.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_55 *A_extension_55) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_55.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_55.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_56 *A_extension_56) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_56.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_56.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_57 *A_extension_57) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_57.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_57.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_58 *A_extension_58) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_58.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_58.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_6 *A_extension_6) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_6.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_6.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_7 *A_extension_7) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_7.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_7.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_8 *A_extension_8) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_8.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_8.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_9 *A_extension_9) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_9.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_9.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractapproachlegtype *AbstractApproachLegType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractapproachlegtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractapproachlegtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractapproachlegtype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractproceduretype *AbstractProcedureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractproceduretype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractproceduretype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractproceduretype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractpropertieswithscheduletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractpropertieswithscheduletype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractpropertieswithscheduletype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractsegmentlegtype *AbstractSegmentLegType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractsegmentlegtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractsegmentlegtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractsegmentlegtype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractsegmentpointtype *AbstractSegmentPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractsegmentpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractsegmentpointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractsegmentpointtype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aerialrefuellingpointtype *AerialRefuellingPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "FacilityMakeup"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "PointChoice_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "PointChoice_navaidSystem"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "PointChoice_position"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "PointChoice_runwayPoint"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "PointChoice_aimingPoint"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "PointChoice_airportReferencePoint"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "ExtendedServiceVolume"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aircraftcharacteristicpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aircraftcharacteristicpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(aircraftcharacteristicpropertytype.GongMarshallField(stage, "AircraftCharacteristic"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aircraftcharacteristictype *AircraftCharacteristicType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aircraftcharacteristictype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aircraftcharacteristictype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aircraftcharacteristictype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(aircraftcharacteristictype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(aircraftcharacteristictype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(altimetersourcepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(altimetersourcepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(altimetersourcepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(altimetersourcepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(altimetersourcepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(altimetersourcepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(altimetersourcepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(altimetersourcepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(altimetersourcepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(altimetersourcepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(altitudeadjustmentpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(altitudeadjustmentpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(altitudeadjustmentpropertytype.GongMarshallField(stage, "AltitudeAdjustment"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (altitudeadjustmenttype *AltitudeAdjustmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(altitudeadjustmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(altitudeadjustmenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(altitudeadjustmenttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(altitudeadjustmenttype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(altitudeadjustmenttype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (angleindicationpropertytype *AngleIndicationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(angleindicationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(angleindicationpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(angleindicationpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(angleindicationpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(angleindicationpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(angleindicationpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(angleindicationpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(angleindicationpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(angleindicationpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(angleindicationpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (angleusepropertytype *AngleUsePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(angleusepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(angleusepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(angleusepropertytype.GongMarshallField(stage, "AngleUse"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (angleusetype *AngleUseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(angleusetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(angleusetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(angleusetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(angleusetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(angleusetype.GongMarshallField(stage, "TheAngleIndication"))
		pointersInitializesStatements.WriteString(angleusetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(approachaltitudetablepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(approachaltitudetablepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(approachaltitudetablepropertytype.GongMarshallField(stage, "ApproachAltitudeTable"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (approachaltitudetabletype *ApproachAltitudeTableType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(approachaltitudetabletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(approachaltitudetabletype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(approachaltitudetabletype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(approachaltitudetabletype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(approachaltitudetabletype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (approachconditionpropertytype *ApproachConditionPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(approachconditionpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(approachconditionpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(approachconditionpropertytype.GongMarshallField(stage, "ApproachCondition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (approachconditiontype *ApproachConditionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(approachconditiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(approachconditiontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(approachconditiontype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(approachconditiontype.GongMarshallField(stage, "MinimumSet"))
		pointersInitializesStatements.WriteString(approachconditiontype.GongMarshallField(stage, "CirclingRestriction"))
		pointersInitializesStatements.WriteString(approachconditiontype.GongMarshallField(stage, "AircraftCategory"))
		pointersInitializesStatements.WriteString(approachconditiontype.GongMarshallField(stage, "LandingArea"))
		pointersInitializesStatements.WriteString(approachconditiontype.GongMarshallField(stage, "Altimeter"))
		pointersInitializesStatements.WriteString(approachconditiontype.GongMarshallField(stage, "DesignSurface"))
		pointersInitializesStatements.WriteString(approachconditiontype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(approachconditiontype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(approachdistancetablepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(approachdistancetablepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(approachdistancetablepropertytype.GongMarshallField(stage, "ApproachDistanceTable"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (approachdistancetabletype *ApproachDistanceTableType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(approachdistancetabletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(approachdistancetabletype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(approachdistancetabletype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(approachdistancetabletype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(approachdistancetabletype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(approachtimingtablepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(approachtimingtablepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(approachtimingtablepropertytype.GongMarshallField(stage, "ApproachTimingTable"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (approachtimingtabletype *ApproachTimingTableType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(approachtimingtabletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(approachtimingtabletype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(approachtimingtabletype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(approachtimingtabletype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(approachtimingtabletype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arrivalfeederlegtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arrivalfeederlegtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicepropertytype.GongMarshallField(stage, "ArrivalFeederLegTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "StartPoint"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "EndPoint"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Trajectory"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "ArcCentre"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Angle"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Distance"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "AircraftCategory"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Holding"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "DesignSurface"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Approach"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arrivalfeederlegtype *ArrivalFeederLegType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arrivalfeederlegtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arrivalfeederlegtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(arrivalfeederlegtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(arrivalfeederlegtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arrivallegtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arrivallegtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicepropertytype.GongMarshallField(stage, "ArrivalLegTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "StartPoint"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "EndPoint"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Trajectory"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "ArcCentre"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Angle"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Distance"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "AircraftCategory"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Holding"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "DesignSurface"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Arrival"))
		pointersInitializesStatements.WriteString(arrivallegtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arrivallegtype *ArrivalLegType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arrivallegtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arrivallegtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(arrivallegtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(arrivallegtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circlesectorpropertytype *CircleSectorPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlesectorpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circlesectorpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(circlesectorpropertytype.GongMarshallField(stage, "CircleSector"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circlesectortype *CircleSectorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlesectortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circlesectortype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(circlesectortype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(circlesectortype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(circlesectortype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlingareatimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circlingareatimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(circlingareatimeslicepropertytype.GongMarshallField(stage, "CirclingAreaTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "Extent"))
		pointersInitializesStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "Approach"))
		pointersInitializesStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "Condition"))
		pointersInitializesStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "AircraftCategory"))
		pointersInitializesStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "DesignSurface"))
		pointersInitializesStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(circlingareatimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circlingareatype *CirclingAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlingareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circlingareatype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(circlingareatype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(circlingareatype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlingrestrictionpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circlingrestrictionpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(circlingrestrictionpropertytype.GongMarshallField(stage, "CirclingRestriction"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circlingrestrictiontype *CirclingRestrictionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlingrestrictiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circlingrestrictiontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(circlingrestrictiontype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(circlingrestrictiontype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(circlingrestrictiontype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(circlingrestrictiontype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(circlingrestrictiontype.GongMarshallField(stage, "SectorDescription"))
		pointersInitializesStatements.WriteString(circlingrestrictiontype.GongMarshallField(stage, "RestrictionArea"))
		pointersInitializesStatements.WriteString(circlingrestrictiontype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (curvepropertytype *CurvePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Curve"))
		pointersInitializesStatements.WriteString(curvepropertytype.GongMarshallField(stage, "ElevatedCurve"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (curvetype *CurveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(curvetype.GongMarshallField(stage, "Annotation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dmepropertytype *DMEPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dmepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(dmepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(dmepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(dmepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(dmepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(dmepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(dmepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(dmepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(dmepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(dmepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(departurearrivalconditionpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(departurearrivalconditionpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(departurearrivalconditionpropertytype.GongMarshallField(stage, "DepartureArrivalCondition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (departurearrivalconditiontype *DepartureArrivalConditionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(departurearrivalconditiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(departurearrivalconditiontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(departurearrivalconditiontype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(departurearrivalconditiontype.GongMarshallField(stage, "EngineType"))
		pointersInitializesStatements.WriteString(departurearrivalconditiontype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(departurearrivalconditiontype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(departurelegtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(departurelegtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(departurelegtimeslicepropertytype.GongMarshallField(stage, "DepartureLegTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "StartPoint"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "EndPoint"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Trajectory"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "ArcCentre"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Angle"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Distance"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "AircraftCategory"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Holding"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "DesignSurface"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Departure"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Condition"))
		pointersInitializesStatements.WriteString(departurelegtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (departurelegtype *DepartureLegType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(departurelegtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(departurelegtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(departurelegtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(departurelegtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (designatedpointpropertytype *DesignatedPointPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(designatedpointpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(designatedpointpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(designatedpointpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(designatedpointpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(designatedpointpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(designatedpointpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(designatedpointpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(designatedpointpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(designatedpointpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(designatedpointpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(distanceindicationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(distanceindicationpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(distanceindicationpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(distanceindicationpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(distanceindicationpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(distanceindicationpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(distanceindicationpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(distanceindicationpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(distanceindicationpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(distanceindicationpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevatedcurvetype *ElevatedCurveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevatedpointtype *ElevatedPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (enroutesegmentpointtype *EnRouteSegmentPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "FacilityMakeup"))
		pointersInitializesStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "PointChoice_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "PointChoice_navaidSystem"))
		pointersInitializesStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "PointChoice_position"))
		pointersInitializesStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "PointChoice_runwayPoint"))
		pointersInitializesStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "PointChoice_aimingPoint"))
		pointersInitializesStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "PointChoice_airportReferencePoint"))
		pointersInitializesStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "ExtendedServiceVolume"))
		pointersInitializesStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(enroutesegmentpointtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(equipmentunavailableadjustmentcolumnpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(equipmentunavailableadjustmentcolumnpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(equipmentunavailableadjustmentcolumnpropertytype.GongMarshallField(stage, "EquipmentUnavailableAdjustmentColumn"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(equipmentunavailableadjustmentcolumntype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(equipmentunavailableadjustmentcolumntype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(equipmentunavailableadjustmentcolumntype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(equipmentunavailableadjustmentcolumntype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(equipmentunavailableadjustmentcolumntype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(equipmentunavailableadjustmentpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(equipmentunavailableadjustmentpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(equipmentunavailableadjustmentpropertytype.GongMarshallField(stage, "EquipmentUnavailableAdjustment"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(equipmentunavailableadjustmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(equipmentunavailableadjustmenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(equipmentunavailableadjustmenttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(equipmentunavailableadjustmenttype.GongMarshallField(stage, "AdjustmentINOPCol"))
		pointersInitializesStatements.WriteString(equipmentunavailableadjustmenttype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(equipmentunavailableadjustmenttype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (fasdatablockpropertytype *FASDataBlockPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(fasdatablockpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(fasdatablockpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(fasdatablockpropertytype.GongMarshallField(stage, "FASDataBlock"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (fasdatablocktype *FASDataBlockType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(fasdatablocktype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(fasdatablocktype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(fasdatablocktype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(fasdatablocktype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(fasdatablocktype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(finallegtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(finallegtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(finallegtimeslicepropertytype.GongMarshallField(stage, "FinalLegTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (finallegtimeslicetype *FinalLegTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "StartPoint"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "EndPoint"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Trajectory"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "ArcCentre"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Angle"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Distance"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "AircraftCategory"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Holding"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "DesignSurface"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Approach"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Condition"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "FinalPathAlignmentPoint_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "FinalPathAlignmentPoint_navaidSystem"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "FinalPathAlignmentPoint_position"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "FinalPathAlignmentPoint_runwayPoint"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "FinalPathAlignmentPoint_aimingPoint"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "FinalPathAlignmentPoint_airportReferencePoint"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "VisualDescentPoint"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "FASData"))
		pointersInitializesStatements.WriteString(finallegtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (finallegtype *FinalLegType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(finallegtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(finallegtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(finallegtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(finallegtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (finalprofilepropertytype *FinalProfilePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(finalprofilepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(finalprofilepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(finalprofilepropertytype.GongMarshallField(stage, "FinalProfile"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (finalprofiletype *FinalProfileType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(finalprofiletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(finalprofiletype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(finalprofiletype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(finalprofiletype.GongMarshallField(stage, "Altitude"))
		pointersInitializesStatements.WriteString(finalprofiletype.GongMarshallField(stage, "Distance"))
		pointersInitializesStatements.WriteString(finalprofiletype.GongMarshallField(stage, "Timing"))
		pointersInitializesStatements.WriteString(finalprofiletype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(finalprofiletype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingassessmenttimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingassessmenttimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(holdingassessmenttimeslicepropertytype.GongMarshallField(stage, "HoldingAssessmentTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "HoldingPoint"))
		pointersInitializesStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "UnplannedHolding"))
		pointersInitializesStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "AssessedHoldingPattern"))
		pointersInitializesStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "ObstacleAssessment"))
		pointersInitializesStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(holdingassessmenttimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingassessmenttype *HoldingAssessmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingassessmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingassessmenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(holdingassessmenttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(holdingassessmenttype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingpatterndistancepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingpatterndistancepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(holdingpatterndistancepropertytype.GongMarshallField(stage, "HoldingPatternDistance"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingpatterndistancetype *HoldingPatternDistanceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingpatterndistancetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingpatterndistancetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(holdingpatterndistancetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(holdingpatterndistancetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(holdingpatterndistancetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingpatterndurationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingpatterndurationpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(holdingpatterndurationpropertytype.GongMarshallField(stage, "HoldingPatternDuration"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingpatterndurationtype *HoldingPatternDurationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingpatterndurationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingpatterndurationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(holdingpatterndurationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(holdingpatterndurationtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(holdingpatterndurationtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingpatternpropertytype *HoldingPatternPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingpatternpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingpatternpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(holdingpatternpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(holdingpatternpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(holdingpatternpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(holdingpatternpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(holdingpatternpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(holdingpatternpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(holdingpatternpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(holdingpatternpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingpatterntimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingpatterntimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(holdingpatterntimeslicepropertytype.GongMarshallField(stage, "HoldingPatternTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "OutboundLegSpan_endPoint"))
		pointersInitializesStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "OutboundLegSpan_endDistance"))
		pointersInitializesStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "OutboundLegSpan_endTime"))
		pointersInitializesStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "HoldingPoint"))
		pointersInitializesStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "Extent"))
		pointersInitializesStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(holdingpatterntimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingpatterntype *HoldingPatternType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingpatterntype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingpatterntype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(holdingpatterntype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(holdingpatterntype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingusepropertytype *HoldingUsePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingusepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingusepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(holdingusepropertytype.GongMarshallField(stage, "HoldingUse"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (holdingusetype *HoldingUseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(holdingusetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(holdingusetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(holdingusetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(holdingusetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(holdingusetype.GongMarshallField(stage, "TheHoldingPattern"))
		pointersInitializesStatements.WriteString(holdingusetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(initiallegtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(initiallegtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(initiallegtimeslicepropertytype.GongMarshallField(stage, "InitialLegTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (initiallegtimeslicetype *InitialLegTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "StartPoint"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "EndPoint"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Trajectory"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "ArcCentre"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Angle"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Distance"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "AircraftCategory"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Holding"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "DesignSurface"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Approach"))
		pointersInitializesStatements.WriteString(initiallegtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (initiallegtype *InitialLegType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(initiallegtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(initiallegtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(initiallegtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(initiallegtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(instrumentapproachprocedurepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(instrumentapproachprocedurepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(instrumentapproachprocedurepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(instrumentapproachprocedurepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(instrumentapproachprocedurepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(instrumentapproachprocedurepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(instrumentapproachprocedurepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(instrumentapproachprocedurepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(instrumentapproachprocedurepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(instrumentapproachprocedurepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(instrumentapproachproceduretimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(instrumentapproachproceduretimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicepropertytype.GongMarshallField(stage, "InstrumentApproachProcedureTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "AircraftCharacteristic"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "FlightTransition"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "GuidanceFacility_specialNavigationSystem"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "GuidanceFacility_radar"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "GuidanceFacility_navaid"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "SafeAltitude"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "Landing"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "MissedInstruction"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "FinalProfile"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(instrumentapproachproceduretype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(instrumentapproachproceduretype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(instrumentapproachproceduretype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(instrumentapproachproceduretype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(intermediatelegtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(intermediatelegtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicepropertytype.GongMarshallField(stage, "IntermediateLegTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "StartPoint"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "EndPoint"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Trajectory"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "ArcCentre"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Angle"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Distance"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "AircraftCategory"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Holding"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "DesignSurface"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Approach"))
		pointersInitializesStatements.WriteString(intermediatelegtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (intermediatelegtype *IntermediateLegType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(intermediatelegtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(intermediatelegtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(intermediatelegtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(intermediatelegtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(landingtakeoffareacollectionpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(landingtakeoffareacollectionpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(landingtakeoffareacollectionpropertytype.GongMarshallField(stage, "LandingTakeoffAreaCollection"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(landingtakeoffareacollectiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(landingtakeoffareacollectiontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(landingtakeoffareacollectiontype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(landingtakeoffareacollectiontype.GongMarshallField(stage, "Runway"))
		pointersInitializesStatements.WriteString(landingtakeoffareacollectiontype.GongMarshallField(stage, "TLOF"))
		pointersInitializesStatements.WriteString(landingtakeoffareacollectiontype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(landingtakeoffareacollectiontype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(linguisticnotepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(linguisticnotepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(linguisticnotepropertytype.GongMarshallField(stage, "LinguisticNote"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (linguisticnotetype *LinguisticNoteType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(linguisticnotetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(linguisticnotetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(linguisticnotetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(linguisticnotetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (minimapropertytype *MinimaPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(minimapropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(minimapropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(minimapropertytype.GongMarshallField(stage, "Minima"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (minimatype *MinimaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(minimatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(minimatype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(minimatype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(minimatype.GongMarshallField(stage, "AdjustmentINOP"))
		pointersInitializesStatements.WriteString(minimatype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(minimatype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(missedapproachgrouppropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(missedapproachgrouppropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(missedapproachgrouppropertytype.GongMarshallField(stage, "MissedApproachGroup"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (missedapproachgrouptype *MissedApproachGroupType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(missedapproachgrouptype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(missedapproachgrouptype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(missedapproachgrouptype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(missedapproachgrouptype.GongMarshallField(stage, "Altimeter"))
		pointersInitializesStatements.WriteString(missedapproachgrouptype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(missedapproachgrouptype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(missedapproachlegtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(missedapproachlegtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicepropertytype.GongMarshallField(stage, "MissedApproachLegTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "StartPoint"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "EndPoint"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Trajectory"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "ArcCentre"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Angle"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Distance"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "AircraftCategory"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Holding"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "DesignSurface"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Approach"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Condition"))
		pointersInitializesStatements.WriteString(missedapproachlegtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (missedapproachlegtype *MissedApproachLegType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(missedapproachlegtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(missedapproachlegtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(missedapproachlegtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(missedapproachlegtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidpropertytype *NavaidPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(navaidpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(navaidpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(navaidpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(navaidpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(navaidpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(navaidpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(navaidpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(navaidpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navigationarearestrictiontimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navigationarearestrictiontimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(navigationarearestrictiontimeslicepropertytype.GongMarshallField(stage, "NavigationAreaRestrictionTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "Procedure"))
		pointersInitializesStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "DesignSurface"))
		pointersInitializesStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "SectorDefinition"))
		pointersInitializesStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(navigationarearestrictiontimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navigationarearestrictiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navigationarearestrictiontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(navigationarearestrictiontype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(navigationarearestrictiontype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navigationareasectorpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navigationareasectorpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(navigationareasectorpropertytype.GongMarshallField(stage, "NavigationAreaSector"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navigationareasectortype *NavigationAreaSectorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navigationareasectortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navigationareasectortype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(navigationareasectortype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(navigationareasectortype.GongMarshallField(stage, "SectorDefinition"))
		pointersInitializesStatements.WriteString(navigationareasectortype.GongMarshallField(stage, "SignificantObstacle"))
		pointersInitializesStatements.WriteString(navigationareasectortype.GongMarshallField(stage, "Extent"))
		pointersInitializesStatements.WriteString(navigationareasectortype.GongMarshallField(stage, "SectorCriteria"))
		pointersInitializesStatements.WriteString(navigationareasectortype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(navigationareasectortype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navigationareatimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navigationareatimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(navigationareatimeslicepropertytype.GongMarshallField(stage, "NavigationAreaTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "Departure"))
		pointersInitializesStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "Sector"))
		pointersInitializesStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "CentrePoint_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "CentrePoint_navaidSystem"))
		pointersInitializesStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "CentrePoint_position"))
		pointersInitializesStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "CentrePoint_runwayPoint"))
		pointersInitializesStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "CentrePoint_aimingPoint"))
		pointersInitializesStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "CentrePoint_airportReferencePoint"))
		pointersInitializesStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(navigationareatimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navigationareatype *NavigationAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navigationareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navigationareatype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(navigationareatype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(navigationareatype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (notepropertytype *NotePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(notepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(notepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(notepropertytype.GongMarshallField(stage, "Note"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (notetype *NoteType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(notetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(notetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(notetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(notetype.GongMarshallField(stage, "TranslatedNote"))
		pointersInitializesStatements.WriteString(notetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(obstacleassessmentareapropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(obstacleassessmentareapropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(obstacleassessmentareapropertytype.GongMarshallField(stage, "ObstacleAssessmentArea"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(obstacleassessmentareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(obstacleassessmentareatype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(obstacleassessmentareatype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(obstacleassessmentareatype.GongMarshallField(stage, "AircraftCategory"))
		pointersInitializesStatements.WriteString(obstacleassessmentareatype.GongMarshallField(stage, "SignificantObstacle"))
		pointersInitializesStatements.WriteString(obstacleassessmentareatype.GongMarshallField(stage, "Surface"))
		pointersInitializesStatements.WriteString(obstacleassessmentareatype.GongMarshallField(stage, "StartingCurve"))
		pointersInitializesStatements.WriteString(obstacleassessmentareatype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(obstacleassessmentareatype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(obstacleplacementpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(obstacleplacementpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(obstacleplacementpropertytype.GongMarshallField(stage, "ObstaclePlacement"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (obstacleplacementtype *ObstaclePlacementType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(obstacleplacementtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(obstacleplacementtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(obstacleplacementtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(obstacleplacementtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(obstacleplacementtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (obstructionpropertytype *ObstructionPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(obstructionpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(obstructionpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(obstructionpropertytype.GongMarshallField(stage, "Obstruction"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (obstructiontype *ObstructionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(obstructiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(obstructiontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(obstructiontype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(obstructiontype.GongMarshallField(stage, "TheVerticalStructure"))
		pointersInitializesStatements.WriteString(obstructiontype.GongMarshallField(stage, "Adjustment"))
		pointersInitializesStatements.WriteString(obstructiontype.GongMarshallField(stage, "ObstaclePlacement"))
		pointersInitializesStatements.WriteString(obstructiontype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(obstructiontype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pointpropertytype *PointPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Point"))
		pointersInitializesStatements.WriteString(pointpropertytype.GongMarshallField(stage, "ElevatedPoint"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pointreferencepropertytype *PointReferencePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pointreferencepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pointreferencepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(pointreferencepropertytype.GongMarshallField(stage, "PointReference"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pointreferencetype *PointReferenceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pointreferencetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pointreferencetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(pointreferencetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(pointreferencetype.GongMarshallField(stage, "Point"))
		pointersInitializesStatements.WriteString(pointreferencetype.GongMarshallField(stage, "FacilityAngle"))
		pointersInitializesStatements.WriteString(pointreferencetype.GongMarshallField(stage, "FacilityDistance"))
		pointersInitializesStatements.WriteString(pointreferencetype.GongMarshallField(stage, "FixToleranceArea"))
		pointersInitializesStatements.WriteString(pointreferencetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(pointreferencetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pointtype *PointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(pointtype.GongMarshallField(stage, "Annotation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(procedureavailabilitypropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(procedureavailabilitypropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(procedureavailabilitypropertytype.GongMarshallField(stage, "ProcedureAvailability"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (procedureavailabilitytype *ProcedureAvailabilityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(procedureavailabilitytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(procedureavailabilitytype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(procedureavailabilitytype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(procedureavailabilitytype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(procedureavailabilitytype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(procedureavailabilitytype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(procedureavailabilitytype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(proceduredmetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(proceduredmetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(proceduredmetimeslicepropertytype.GongMarshallField(stage, "ProcedureDMETimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(proceduredmetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(proceduredmetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(proceduredmetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(proceduredmetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(proceduredmetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(proceduredmetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(proceduredmetimeslicetype.GongMarshallField(stage, "DME"))
		pointersInitializesStatements.WriteString(proceduredmetimeslicetype.GongMarshallField(stage, "SegmentLeg"))
		pointersInitializesStatements.WriteString(proceduredmetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(proceduredmetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (proceduredmetype *ProcedureDMEType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(proceduredmetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(proceduredmetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(proceduredmetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(proceduredmetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (procedurepropertytype *ProcedurePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(procedurepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(procedurepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(procedurepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(procedurepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(procedurepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(procedurepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(procedurepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(procedurepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(procedurepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(procedurepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(proceduretransitionlegpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(proceduretransitionlegpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(proceduretransitionlegpropertytype.GongMarshallField(stage, "ProcedureTransitionLeg"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (proceduretransitionlegtype *ProcedureTransitionLegType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(proceduretransitionlegtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(proceduretransitionlegtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(proceduretransitionlegtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(proceduretransitionlegtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(proceduretransitionlegtype.GongMarshallField(stage, "TheSegmentLeg"))
		pointersInitializesStatements.WriteString(proceduretransitionlegtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(proceduretransitionpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(proceduretransitionpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(proceduretransitionpropertytype.GongMarshallField(stage, "ProcedureTransition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (proceduretransitiontype *ProcedureTransitionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(proceduretransitiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(proceduretransitiontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(proceduretransitiontype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(proceduretransitiontype.GongMarshallField(stage, "DepartureRunwayTransition"))
		pointersInitializesStatements.WriteString(proceduretransitiontype.GongMarshallField(stage, "Trajectory"))
		pointersInitializesStatements.WriteString(proceduretransitiontype.GongMarshallField(stage, "TransitionLeg"))
		pointersInitializesStatements.WriteString(proceduretransitiontype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(proceduretransitiontype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radarsystempropertytype *RadarSystemPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radarsystempropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radarsystempropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(radarsystempropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(radarsystempropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(radarsystempropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(radarsystempropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(radarsystempropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(radarsystempropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(radarsystempropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(radarsystempropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radiofrequencyareapropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radiofrequencyareapropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(radiofrequencyareapropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(radiofrequencyareapropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(radiofrequencyareapropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(radiofrequencyareapropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(radiofrequencyareapropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(radiofrequencyareapropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(radiofrequencyareapropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(radiofrequencyareapropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(runwaycentrelinepointpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(runwaycentrelinepointpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(runwaycentrelinepointpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(runwaycentrelinepointpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(runwaycentrelinepointpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(runwaycentrelinepointpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(runwaycentrelinepointpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(runwaycentrelinepointpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(runwaycentrelinepointpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(runwaycentrelinepointpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(safealtitudeareapropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(safealtitudeareapropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(safealtitudeareapropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(safealtitudeareapropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(safealtitudeareapropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(safealtitudeareapropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(safealtitudeareapropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(safealtitudeareapropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(safealtitudeareapropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(safealtitudeareapropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(safealtitudeareasectorpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(safealtitudeareasectorpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(safealtitudeareasectorpropertytype.GongMarshallField(stage, "SafeAltitudeAreaSector"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(safealtitudeareasectortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(safealtitudeareasectortype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(safealtitudeareasectortype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(safealtitudeareasectortype.GongMarshallField(stage, "Extent"))
		pointersInitializesStatements.WriteString(safealtitudeareasectortype.GongMarshallField(stage, "SignificantObstacle"))
		pointersInitializesStatements.WriteString(safealtitudeareasectortype.GongMarshallField(stage, "SectorDefinition"))
		pointersInitializesStatements.WriteString(safealtitudeareasectortype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(safealtitudeareasectortype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(safealtitudeareatimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(safealtitudeareatimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicepropertytype.GongMarshallField(stage, "SafeAltitudeAreaTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "CentrePoint_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "CentrePoint_navaidSystem"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "CentrePoint_position"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "CentrePoint_runwayPoint"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "CentrePoint_aimingPoint"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "CentrePoint_airportReferencePoint"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "Sector"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(safealtitudeareatimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (safealtitudeareatype *SafeAltitudeAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(safealtitudeareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(safealtitudeareatype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(safealtitudeareatype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(safealtitudeareatype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (sectordesignpropertytype *SectorDesignPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(sectordesignpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(sectordesignpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(sectordesignpropertytype.GongMarshallField(stage, "SectorDesign"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (sectordesigntype *SectorDesignType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(sectordesigntype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(sectordesigntype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(sectordesigntype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(sectordesigntype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(sectordesigntype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (segmentlegpropertytype *SegmentLegPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(segmentlegpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(segmentlegpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(segmentlegpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(segmentlegpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(segmentlegpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(segmentlegpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(segmentlegpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(segmentlegpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(segmentlegpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(segmentlegpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (segmentpointpropertytype *SegmentPointPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(segmentpointpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(segmentpointpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(segmentpointpropertytype.GongMarshallField(stage, "AerialRefuellingPoint"))
		pointersInitializesStatements.WriteString(segmentpointpropertytype.GongMarshallField(stage, "EnRouteSegmentPoint"))
		pointersInitializesStatements.WriteString(segmentpointpropertytype.GongMarshallField(stage, "TerminalSegmentPoint"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialnavigationsystempropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialnavigationsystempropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(specialnavigationsystempropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(specialnavigationsystempropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(specialnavigationsystempropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(specialnavigationsystempropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(specialnavigationsystempropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(specialnavigationsystempropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(specialnavigationsystempropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(specialnavigationsystempropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardinstrumentarrivalpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardinstrumentarrivalpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(standardinstrumentarrivalpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(standardinstrumentarrivalpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(standardinstrumentarrivalpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(standardinstrumentarrivalpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(standardinstrumentarrivalpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(standardinstrumentarrivalpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(standardinstrumentarrivalpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(standardinstrumentarrivalpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardinstrumentarrivaltimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardinstrumentarrivaltimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicepropertytype.GongMarshallField(stage, "StandardInstrumentArrivalTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "AircraftCharacteristic"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "FlightTransition"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "GuidanceFacility_specialNavigationSystem"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "GuidanceFacility_radar"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "GuidanceFacility_navaid"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "SafeAltitude"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "Arrival"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardinstrumentarrivaltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardinstrumentarrivaltype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardinstrumentarrivaltype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(standardinstrumentarrivaltype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardinstrumentdeparturepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardinstrumentdeparturepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(standardinstrumentdeparturepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(standardinstrumentdeparturepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(standardinstrumentdeparturepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(standardinstrumentdeparturepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(standardinstrumentdeparturepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(standardinstrumentdeparturepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(standardinstrumentdeparturepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(standardinstrumentdeparturepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardinstrumentdeparturetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardinstrumentdeparturetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicepropertytype.GongMarshallField(stage, "StandardInstrumentDepartureTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "AircraftCharacteristic"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "FlightTransition"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "GuidanceFacility_specialNavigationSystem"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "GuidanceFacility_radar"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "GuidanceFacility_navaid"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "SafeAltitude"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "Takeoff"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardinstrumentdeparturetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardinstrumentdeparturetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardinstrumentdeparturetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(standardinstrumentdeparturetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (surfacepropertytype *SurfacePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(surfacepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(surfacepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(surfacepropertytype.GongMarshallField(stage, "Surface"))
		pointersInitializesStatements.WriteString(surfacepropertytype.GongMarshallField(stage, "ElevatedSurface"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (surfacetype *SurfaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(surfacetype.GongMarshallField(stage, "Annotation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(terminalarrivalareasectorpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(terminalarrivalareasectorpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(terminalarrivalareasectorpropertytype.GongMarshallField(stage, "TerminalArrivalAreaSector"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(terminalarrivalareasectortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(terminalarrivalareasectortype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(terminalarrivalareasectortype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(terminalarrivalareasectortype.GongMarshallField(stage, "SectorDefinition"))
		pointersInitializesStatements.WriteString(terminalarrivalareasectortype.GongMarshallField(stage, "Extent"))
		pointersInitializesStatements.WriteString(terminalarrivalareasectortype.GongMarshallField(stage, "SignificantObstacle"))
		pointersInitializesStatements.WriteString(terminalarrivalareasectortype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(terminalarrivalareasectortype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(terminalarrivalareatimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(terminalarrivalareatimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicepropertytype.GongMarshallField(stage, "TerminalArrivalAreaTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IF_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IF_navaidSystem"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IF_position"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IF_runwayPoint"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IF_aimingPoint"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IF_airportReferencePoint"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IAF_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IAF_navaidSystem"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IAF_position"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IAF_runwayPoint"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IAF_aimingPoint"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "IAF_airportReferencePoint"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "Buffer"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "Sector"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "ApproachRNAV"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(terminalarrivalareatimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (terminalarrivalareatype *TerminalArrivalAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(terminalarrivalareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(terminalarrivalareatype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(terminalarrivalareatype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(terminalarrivalareatype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(terminalsegmentpointpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(terminalsegmentpointpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(terminalsegmentpointpropertytype.GongMarshallField(stage, "TerminalSegmentPoint"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (terminalsegmentpointtype *TerminalSegmentPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "FacilityMakeup"))
		pointersInitializesStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "PointChoice_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "PointChoice_navaidSystem"))
		pointersInitializesStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "PointChoice_position"))
		pointersInitializesStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "PointChoice_runwayPoint"))
		pointersInitializesStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "PointChoice_aimingPoint"))
		pointersInitializesStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "PointChoice_airportReferencePoint"))
		pointersInitializesStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "ExtendedServiceVolume"))
		pointersInitializesStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(terminalsegmentpointtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (timesheetpropertytype *TimesheetPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(timesheetpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(timesheetpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(timesheetpropertytype.GongMarshallField(stage, "Timesheet"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (timesheettype *TimesheetType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(timesheettype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(timesheettype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(timesheettype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(timesheettype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(timesheettype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(touchdownliftoffpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(touchdownliftoffpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(touchdownliftoffpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(touchdownliftoffpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(touchdownliftoffpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(touchdownliftoffpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(touchdownliftoffpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(touchdownliftoffpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(touchdownliftoffpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(touchdownliftoffpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unplannedholdingpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unplannedholdingpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(unplannedholdingpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(unplannedholdingpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(unplannedholdingpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(unplannedholdingpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(unplannedholdingpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(unplannedholdingpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(unplannedholdingpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(unplannedholdingpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unplannedholdingtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unplannedholdingtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(unplannedholdingtimeslicepropertytype.GongMarshallField(stage, "UnplannedHoldingTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unplannedholdingtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unplannedholdingtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(unplannedholdingtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(unplannedholdingtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(unplannedholdingtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(unplannedholdingtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(unplannedholdingtimeslicetype.GongMarshallField(stage, "HoldingPoint"))
		pointersInitializesStatements.WriteString(unplannedholdingtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(unplannedholdingtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unplannedholdingtype *UnplannedHoldingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unplannedholdingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unplannedholdingtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(unplannedholdingtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(unplannedholdingtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticalstructurepropertytype *VerticalStructurePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
