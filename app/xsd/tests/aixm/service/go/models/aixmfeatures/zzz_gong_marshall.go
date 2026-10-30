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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures"
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
		for _, abstractairportgroundservicetype := range __gong__sortStageSetInstances(stage.AbstractAirportGroundServiceTypes, stage.AbstractAirportGroundServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractairportgroundservicetypeIdent := abstractairportgroundservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractAirportGroundServiceType{Name: %s}).Stage(stage)", abstractairportgroundservicetypeIdent, __gong__toRawStringLiteral(abstractairportgroundservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractairportgroundservicetypeIdent, __gong__toRawStringLiteral(abstractairportgroundservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractairportgroundservicetypeIdent, __gong__toRawStringLiteral(abstractairportgroundservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractairportgroundservicetypeIdent, __gong__toRawStringLiteral(abstractairportgroundservicetype.Description)))
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
		for _, abstractservicetype := range __gong__sortStageSetInstances(stage.AbstractServiceTypes, stage.AbstractServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractservicetypeIdent := abstractservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractServiceType{Name: %s}).Stage(stage)", abstractservicetypeIdent, __gong__toRawStringLiteral(abstractservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractservicetypeIdent, __gong__toRawStringLiteral(abstractservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractservicetypeIdent, __gong__toRawStringLiteral(abstractservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractservicetypeIdent, __gong__toRawStringLiteral(abstractservicetype.Description)))
		}
	}
	if stage != nil {
		for _, abstracttrafficseparationservicetype := range __gong__sortStageSetInstances(stage.AbstractTrafficSeparationServiceTypes, stage.AbstractTrafficSeparationServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstracttrafficseparationservicetypeIdent := abstracttrafficseparationservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractTrafficSeparationServiceType{Name: %s}).Stage(stage)", abstracttrafficseparationservicetypeIdent, __gong__toRawStringLiteral(abstracttrafficseparationservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstracttrafficseparationservicetypeIdent, __gong__toRawStringLiteral(abstracttrafficseparationservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstracttrafficseparationservicetypeIdent, __gong__toRawStringLiteral(abstracttrafficseparationservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstracttrafficseparationservicetypeIdent, __gong__toRawStringLiteral(abstracttrafficseparationservicetype.Description)))
		}
	}
	if stage != nil {
		for _, aerialrefuellingpropertytype := range __gong__sortStageSetInstances(stage.AerialRefuellingPropertyTypes, stage.AerialRefuellingPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aerialrefuellingpropertytypeIdent := aerialrefuellingpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingPropertyType{Name: %s}).Stage(stage)", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", aerialrefuellingpropertytypeIdent, aerialrefuellingpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, airtrafficcontrolservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AirTrafficControlServiceTimeSlicePropertyTypes, stage.AirTrafficControlServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airtrafficcontrolservicetimeslicepropertytypeIdent := airtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficControlServiceTimeSlicePropertyType{Name: %s}).Stage(stage)", airtrafficcontrolservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficcontrolservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airtrafficcontrolservicetimeslicepropertytypeIdent, airtrafficcontrolservicetimeslicepropertytype.Owns))
			if airtrafficcontrolservicetimeslicepropertytype.AirTrafficControlServiceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airtrafficcontrolservicetimeslicepropertytype.AirTrafficControlServiceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirTrafficControlServiceTimeSlice = %s", airtrafficcontrolservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airtrafficcontrolservicetimeslicetype := range __gong__sortStageSetInstances(stage.AirTrafficControlServiceTimeSliceTypes, stage.AirTrafficControlServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airtrafficcontrolservicetimeslicetypeIdent := airtrafficcontrolservicetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficControlServiceTimeSliceType{Name: %s}).Stage(stage)", airtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetype.CorrectionNumber))
			if airtrafficcontrolservicetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.Call_sign {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientAirport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirport = append(%s.ClientAirport, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientAirspace {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirspace = append(%s.ClientAirspace, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientRoute {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientRoute = append(%s.ClientRoute, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientProcedure {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientProcedure = append(%s.ClientProcedure, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientHolding {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientHolding = append(%s.ClientHolding, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientAerialRefuelling {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAerialRefuelling = append(%s.ClientAerialRefuelling, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.AircraftLocator != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.AircraftLocator.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftLocator = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airtrafficcontrolservicetype := range __gong__sortStageSetInstances(stage.AirTrafficControlServiceTypes, stage.AirTrafficControlServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airtrafficcontrolservicetypeIdent := airtrafficcontrolservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficControlServiceType{Name: %s}).Stage(stage)", airtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetype.Description)))
			for _, elem := range airtrafficcontrolservicetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airtrafficcontrolservicetypeIdent, airtrafficcontrolservicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airtrafficmanagementservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AirTrafficManagementServiceTimeSlicePropertyTypes, stage.AirTrafficManagementServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airtrafficmanagementservicetimeslicepropertytypeIdent := airtrafficmanagementservicetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficManagementServiceTimeSlicePropertyType{Name: %s}).Stage(stage)", airtrafficmanagementservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficmanagementservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airtrafficmanagementservicetimeslicepropertytypeIdent, airtrafficmanagementservicetimeslicepropertytype.Owns))
			if airtrafficmanagementservicetimeslicepropertytype.AirTrafficManagementServiceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airtrafficmanagementservicetimeslicepropertytype.AirTrafficManagementServiceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirTrafficManagementServiceTimeSlice = %s", airtrafficmanagementservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airtrafficmanagementservicetimeslicetype := range __gong__sortStageSetInstances(stage.AirTrafficManagementServiceTimeSliceTypes, stage.AirTrafficManagementServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airtrafficmanagementservicetimeslicetypeIdent := airtrafficmanagementservicetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficManagementServiceTimeSliceType{Name: %s}).Stage(stage)", airtrafficmanagementservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficmanagementservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airtrafficmanagementservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airtrafficmanagementservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airtrafficmanagementservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetype.CorrectionNumber))
			if airtrafficmanagementservicetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.Call_sign {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.ClientAirspace {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirspace = append(%s.ClientAirspace, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.ClientAerialRefuelling {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAerialRefuelling = append(%s.ClientAerialRefuelling, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.ClientRoute {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientRoute = append(%s.ClientRoute, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airtrafficmanagementservicetype := range __gong__sortStageSetInstances(stage.AirTrafficManagementServiceTypes, stage.AirTrafficManagementServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airtrafficmanagementservicetypeIdent := airtrafficmanagementservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficManagementServiceType{Name: %s}).Stage(stage)", airtrafficmanagementservicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficmanagementservicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airtrafficmanagementservicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airtrafficmanagementservicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetype.Description)))
			for _, elem := range airtrafficmanagementservicetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airtrafficmanagementservicetypeIdent, airtrafficmanagementservicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aircraftgroundservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AircraftGroundServiceTimeSlicePropertyTypes, stage.AircraftGroundServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aircraftgroundservicetimeslicepropertytypeIdent := aircraftgroundservicetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftGroundServiceTimeSlicePropertyType{Name: %s}).Stage(stage)", aircraftgroundservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftgroundservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", aircraftgroundservicetimeslicepropertytypeIdent, aircraftgroundservicetimeslicepropertytype.Owns))
			if aircraftgroundservicetimeslicepropertytype.AircraftGroundServiceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aircraftgroundservicetimeslicepropertytype.AircraftGroundServiceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftGroundServiceTimeSlice = %s", aircraftgroundservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aircraftgroundservicetimeslicetype := range __gong__sortStageSetInstances(stage.AircraftGroundServiceTimeSliceTypes, stage.AircraftGroundServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aircraftgroundservicetimeslicetypeIdent := aircraftgroundservicetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftGroundServiceTimeSliceType{Name: %s}).Stage(stage)", aircraftgroundservicetimeslicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftgroundservicetimeslicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aircraftgroundservicetimeslicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aircraftgroundservicetimeslicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", aircraftgroundservicetimeslicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetype.CorrectionNumber))
			if aircraftgroundservicetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.Call_sign {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.AirportHeliport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aircraftgroundservicetype := range __gong__sortStageSetInstances(stage.AircraftGroundServiceTypes, stage.AircraftGroundServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aircraftgroundservicetypeIdent := aircraftgroundservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftGroundServiceType{Name: %s}).Stage(stage)", aircraftgroundservicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftgroundservicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aircraftgroundservicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aircraftgroundservicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetype.Description)))
			for _, elem := range aircraftgroundservicetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", aircraftgroundservicetypeIdent, aircraftgroundservicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airportclearanceservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AirportClearanceServiceTimeSlicePropertyTypes, stage.AirportClearanceServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airportclearanceservicetimeslicepropertytypeIdent := airportclearanceservicetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportClearanceServiceTimeSlicePropertyType{Name: %s}).Stage(stage)", airportclearanceservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportclearanceservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airportclearanceservicetimeslicepropertytypeIdent, airportclearanceservicetimeslicepropertytype.Owns))
			if airportclearanceservicetimeslicepropertytype.AirportClearanceServiceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airportclearanceservicetimeslicepropertytype.AirportClearanceServiceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportClearanceServiceTimeSlice = %s", airportclearanceservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airportclearanceservicetimeslicetype := range __gong__sortStageSetInstances(stage.AirportClearanceServiceTimeSliceTypes, stage.AirportClearanceServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airportclearanceservicetimeslicetypeIdent := airportclearanceservicetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportClearanceServiceTimeSliceType{Name: %s}).Stage(stage)", airportclearanceservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportclearanceservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airportclearanceservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airportclearanceservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airportclearanceservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetype.CorrectionNumber))
			if airportclearanceservicetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airportclearanceservicetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			if airportclearanceservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airportclearanceservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.Call_sign {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.AirportHeliport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airportclearanceservicetype := range __gong__sortStageSetInstances(stage.AirportClearanceServiceTypes, stage.AirportClearanceServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airportclearanceservicetypeIdent := airportclearanceservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportClearanceServiceType{Name: %s}).Stage(stage)", airportclearanceservicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportclearanceservicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airportclearanceservicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airportclearanceservicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetype.Description)))
			for _, elem := range airportclearanceservicetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airportclearanceservicetypeIdent, airportclearanceservicetypeIdent, targetIdent))
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
		for _, airportsuppliesservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AirportSuppliesServiceTimeSlicePropertyTypes, stage.AirportSuppliesServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airportsuppliesservicetimeslicepropertytypeIdent := airportsuppliesservicetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportSuppliesServiceTimeSlicePropertyType{Name: %s}).Stage(stage)", airportsuppliesservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportsuppliesservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airportsuppliesservicetimeslicepropertytypeIdent, airportsuppliesservicetimeslicepropertytype.Owns))
			if airportsuppliesservicetimeslicepropertytype.AirportSuppliesServiceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airportsuppliesservicetimeslicepropertytype.AirportSuppliesServiceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportSuppliesServiceTimeSlice = %s", airportsuppliesservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airportsuppliesservicetimeslicetype := range __gong__sortStageSetInstances(stage.AirportSuppliesServiceTimeSliceTypes, stage.AirportSuppliesServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airportsuppliesservicetimeslicetypeIdent := airportsuppliesservicetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportSuppliesServiceTimeSliceType{Name: %s}).Stage(stage)", airportsuppliesservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportsuppliesservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airportsuppliesservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airportsuppliesservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airportsuppliesservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetype.CorrectionNumber))
			if airportsuppliesservicetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.Call_sign {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.AirportHeliport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.FuelSupply {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FuelSupply = append(%s.FuelSupply, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.OilSupply {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OilSupply = append(%s.OilSupply, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.NitrogenSupply {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NitrogenSupply = append(%s.NitrogenSupply, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.OxygenSupply {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OxygenSupply = append(%s.OxygenSupply, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airportsuppliesservicetype := range __gong__sortStageSetInstances(stage.AirportSuppliesServiceTypes, stage.AirportSuppliesServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airportsuppliesservicetypeIdent := airportsuppliesservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportSuppliesServiceType{Name: %s}).Stage(stage)", airportsuppliesservicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportsuppliesservicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airportsuppliesservicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airportsuppliesservicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetype.Description)))
			for _, elem := range airportsuppliesservicetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airportsuppliesservicetypeIdent, airportsuppliesservicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacepropertytype := range __gong__sortStageSetInstances(stage.AirspacePropertyTypes, stage.AirspacePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacepropertytypeIdent := airspacepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspacePropertyType{Name: %s}).Stage(stage)", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airspacepropertytypeIdent, airspacepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, callsigndetailpropertytype := range __gong__sortStageSetInstances(stage.CallsignDetailPropertyTypes, stage.CallsignDetailPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			callsigndetailpropertytypeIdent := callsigndetailpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CallsignDetailPropertyType{Name: %s}).Stage(stage)", callsigndetailpropertytypeIdent, __gong__toRawStringLiteral(callsigndetailpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", callsigndetailpropertytypeIdent, __gong__toRawStringLiteral(callsigndetailpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", callsigndetailpropertytypeIdent, __gong__toRawStringLiteral(callsigndetailpropertytype.NilReason)))
			if callsigndetailpropertytype.CallsignDetail != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := callsigndetailpropertytype.CallsignDetail.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CallsignDetail = %s", callsigndetailpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, callsigndetailtype := range __gong__sortStageSetInstances(stage.CallsignDetailTypes, stage.CallsignDetailType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			callsigndetailtypeIdent := callsigndetailtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CallsignDetailType{Name: %s}).Stage(stage)", callsigndetailtypeIdent, __gong__toRawStringLiteral(callsigndetailtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", callsigndetailtypeIdent, __gong__toRawStringLiteral(callsigndetailtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", callsigndetailtypeIdent, __gong__toRawStringLiteral(callsigndetailtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", callsigndetailtypeIdent, __gong__toRawStringLiteral(callsigndetailtype.Description)))
			for _, elem := range callsigndetailtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", callsigndetailtypeIdent, callsigndetailtypeIdent, targetIdent))
			}
			for _, elem := range callsigndetailtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", callsigndetailtypeIdent, callsigndetailtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, contactinformationpropertytype := range __gong__sortStageSetInstances(stage.ContactInformationPropertyTypes, stage.ContactInformationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			contactinformationpropertytypeIdent := contactinformationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ContactInformationPropertyType{Name: %s}).Stage(stage)", contactinformationpropertytypeIdent, __gong__toRawStringLiteral(contactinformationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", contactinformationpropertytypeIdent, __gong__toRawStringLiteral(contactinformationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", contactinformationpropertytypeIdent, __gong__toRawStringLiteral(contactinformationpropertytype.NilReason)))
			if contactinformationpropertytype.ContactInformation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := contactinformationpropertytype.ContactInformation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ContactInformation = %s", contactinformationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, contactinformationtype := range __gong__sortStageSetInstances(stage.ContactInformationTypes, stage.ContactInformationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			contactinformationtypeIdent := contactinformationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ContactInformationType{Name: %s}).Stage(stage)", contactinformationtypeIdent, __gong__toRawStringLiteral(contactinformationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", contactinformationtypeIdent, __gong__toRawStringLiteral(contactinformationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", contactinformationtypeIdent, __gong__toRawStringLiteral(contactinformationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", contactinformationtypeIdent, __gong__toRawStringLiteral(contactinformationtype.Description)))
			for _, elem := range contactinformationtype.Address {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Address = append(%s.Address, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
			}
			for _, elem := range contactinformationtype.NetworkNode {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NetworkNode = append(%s.NetworkNode, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
			}
			for _, elem := range contactinformationtype.PhoneFax {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PhoneFax = append(%s.PhoneFax, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
			}
			for _, elem := range contactinformationtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
			}
			for _, elem := range contactinformationtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
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
		for _, directionfinderpropertytype := range __gong__sortStageSetInstances(stage.DirectionFinderPropertyTypes, stage.DirectionFinderPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directionfinderpropertytypeIdent := directionfinderpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectionFinderPropertyType{Name: %s}).Stage(stage)", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", directionfinderpropertytypeIdent, directionfinderpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, elevatedpointpropertytype := range __gong__sortStageSetInstances(stage.ElevatedPointPropertyTypes, stage.ElevatedPointPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevatedpointpropertytypeIdent := elevatedpointpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedPointPropertyType{Name: %s}).Stage(stage)", elevatedpointpropertytypeIdent, __gong__toRawStringLiteral(elevatedpointpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedpointpropertytypeIdent, __gong__toRawStringLiteral(elevatedpointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", elevatedpointpropertytypeIdent, __gong__toRawStringLiteral(elevatedpointpropertytype.NilReason)))
			if elevatedpointpropertytype.ElevatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elevatedpointpropertytype.ElevatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedPoint = %s", elevatedpointpropertytypeIdent, targetIdent))
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
		for _, firefightingservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.FireFightingServiceTimeSlicePropertyTypes, stage.FireFightingServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			firefightingservicetimeslicepropertytypeIdent := firefightingservicetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FireFightingServiceTimeSlicePropertyType{Name: %s}).Stage(stage)", firefightingservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", firefightingservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", firefightingservicetimeslicepropertytypeIdent, firefightingservicetimeslicepropertytype.Owns))
			if firefightingservicetimeslicepropertytype.FireFightingServiceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := firefightingservicetimeslicepropertytype.FireFightingServiceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FireFightingServiceTimeSlice = %s", firefightingservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, firefightingservicetimeslicetype := range __gong__sortStageSetInstances(stage.FireFightingServiceTimeSliceTypes, stage.FireFightingServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			firefightingservicetimeslicetypeIdent := firefightingservicetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FireFightingServiceTimeSliceType{Name: %s}).Stage(stage)", firefightingservicetimeslicetypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", firefightingservicetimeslicetypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", firefightingservicetimeslicetypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", firefightingservicetimeslicetypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", firefightingservicetimeslicetypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetype.CorrectionNumber))
			if firefightingservicetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := firefightingservicetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := firefightingservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.Call_sign {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.AirportHeliport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, firefightingservicetype := range __gong__sortStageSetInstances(stage.FireFightingServiceTypes, stage.FireFightingServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			firefightingservicetypeIdent := firefightingservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FireFightingServiceType{Name: %s}).Stage(stage)", firefightingservicetypeIdent, __gong__toRawStringLiteral(firefightingservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", firefightingservicetypeIdent, __gong__toRawStringLiteral(firefightingservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", firefightingservicetypeIdent, __gong__toRawStringLiteral(firefightingservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", firefightingservicetypeIdent, __gong__toRawStringLiteral(firefightingservicetype.Description)))
			for _, elem := range firefightingservicetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", firefightingservicetypeIdent, firefightingservicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, fuelpropertytype := range __gong__sortStageSetInstances(stage.FuelPropertyTypes, stage.FuelPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fuelpropertytypeIdent := fuelpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FuelPropertyType{Name: %s}).Stage(stage)", fuelpropertytypeIdent, __gong__toRawStringLiteral(fuelpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fuelpropertytypeIdent, __gong__toRawStringLiteral(fuelpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", fuelpropertytypeIdent, __gong__toRawStringLiteral(fuelpropertytype.NilReason)))
			if fuelpropertytype.Fuel != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := fuelpropertytype.Fuel.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fuel = %s", fuelpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, fueltype := range __gong__sortStageSetInstances(stage.FuelTypes, stage.FuelType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fueltypeIdent := fueltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FuelType{Name: %s}).Stage(stage)", fueltypeIdent, __gong__toRawStringLiteral(fueltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fueltypeIdent, __gong__toRawStringLiteral(fueltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", fueltypeIdent, __gong__toRawStringLiteral(fueltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", fueltypeIdent, __gong__toRawStringLiteral(fueltype.Description)))
			for _, elem := range fueltype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", fueltypeIdent, fueltypeIdent, targetIdent))
			}
			for _, elem := range fueltype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", fueltypeIdent, fueltypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, groundlightsystempropertytype := range __gong__sortStageSetInstances(stage.GroundLightSystemPropertyTypes, stage.GroundLightSystemPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groundlightsystempropertytypeIdent := groundlightsystempropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GroundLightSystemPropertyType{Name: %s}).Stage(stage)", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", groundlightsystempropertytypeIdent, groundlightsystempropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, groundtrafficcontrolservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.GroundTrafficControlServiceTimeSlicePropertyTypes, stage.GroundTrafficControlServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groundtrafficcontrolservicetimeslicepropertytypeIdent := groundtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GroundTrafficControlServiceTimeSlicePropertyType{Name: %s}).Stage(stage)", groundtrafficcontrolservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groundtrafficcontrolservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", groundtrafficcontrolservicetimeslicepropertytypeIdent, groundtrafficcontrolservicetimeslicepropertytype.Owns))
			if groundtrafficcontrolservicetimeslicepropertytype.GroundTrafficControlServiceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicepropertytype.GroundTrafficControlServiceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundTrafficControlServiceTimeSlice = %s", groundtrafficcontrolservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, groundtrafficcontrolservicetimeslicetype := range __gong__sortStageSetInstances(stage.GroundTrafficControlServiceTimeSliceTypes, stage.GroundTrafficControlServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groundtrafficcontrolservicetimeslicetypeIdent := groundtrafficcontrolservicetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GroundTrafficControlServiceTimeSliceType{Name: %s}).Stage(stage)", groundtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groundtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", groundtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", groundtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", groundtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetype.CorrectionNumber))
			if groundtrafficcontrolservicetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.Call_sign {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.ClientAirport != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.ClientAirport.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirport = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, groundtrafficcontrolservicetype := range __gong__sortStageSetInstances(stage.GroundTrafficControlServiceTypes, stage.GroundTrafficControlServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groundtrafficcontrolservicetypeIdent := groundtrafficcontrolservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GroundTrafficControlServiceType{Name: %s}).Stage(stage)", groundtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groundtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", groundtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", groundtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetype.Description)))
			for _, elem := range groundtrafficcontrolservicetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", groundtrafficcontrolservicetypeIdent, groundtrafficcontrolservicetypeIdent, targetIdent))
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
		for _, informationservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.InformationServiceTimeSlicePropertyTypes, stage.InformationServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			informationservicetimeslicepropertytypeIdent := informationservicetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InformationServiceTimeSlicePropertyType{Name: %s}).Stage(stage)", informationservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(informationservicetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", informationservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(informationservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", informationservicetimeslicepropertytypeIdent, informationservicetimeslicepropertytype.Owns))
			if informationservicetimeslicepropertytype.InformationServiceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := informationservicetimeslicepropertytype.InformationServiceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InformationServiceTimeSlice = %s", informationservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, informationservicetimeslicetype := range __gong__sortStageSetInstances(stage.InformationServiceTimeSliceTypes, stage.InformationServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			informationservicetimeslicetypeIdent := informationservicetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InformationServiceTimeSliceType{Name: %s}).Stage(stage)", informationservicetimeslicetypeIdent, __gong__toRawStringLiteral(informationservicetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", informationservicetimeslicetypeIdent, __gong__toRawStringLiteral(informationservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", informationservicetimeslicetypeIdent, __gong__toRawStringLiteral(informationservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", informationservicetimeslicetypeIdent, __gong__toRawStringLiteral(informationservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", informationservicetimeslicetypeIdent, __gong__toRawStringLiteral(informationservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", informationservicetimeslicetypeIdent, informationservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", informationservicetimeslicetypeIdent, informationservicetimeslicetype.CorrectionNumber))
			if informationservicetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := informationservicetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := informationservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.Call_sign {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.NavaidBroadcast {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidBroadcast = append(%s.NavaidBroadcast, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientAirspace {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirspace = append(%s.ClientAirspace, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientAirport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirport = append(%s.ClientAirport, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientRoute {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientRoute = append(%s.ClientRoute, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientProcedure {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientProcedure = append(%s.ClientProcedure, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientHolding {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientHolding = append(%s.ClientHolding, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientAerialRefuelling {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAerialRefuelling = append(%s.ClientAerialRefuelling, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, informationservicetype := range __gong__sortStageSetInstances(stage.InformationServiceTypes, stage.InformationServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			informationservicetypeIdent := informationservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InformationServiceType{Name: %s}).Stage(stage)", informationservicetypeIdent, __gong__toRawStringLiteral(informationservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", informationservicetypeIdent, __gong__toRawStringLiteral(informationservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", informationservicetypeIdent, __gong__toRawStringLiteral(informationservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", informationservicetypeIdent, __gong__toRawStringLiteral(informationservicetype.Description)))
			for _, elem := range informationservicetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", informationservicetypeIdent, informationservicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, lightactivationpropertytype := range __gong__sortStageSetInstances(stage.LightActivationPropertyTypes, stage.LightActivationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lightactivationpropertytypeIdent := lightactivationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightActivationPropertyType{Name: %s}).Stage(stage)", lightactivationpropertytypeIdent, __gong__toRawStringLiteral(lightactivationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightactivationpropertytypeIdent, __gong__toRawStringLiteral(lightactivationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", lightactivationpropertytypeIdent, __gong__toRawStringLiteral(lightactivationpropertytype.NilReason)))
			if lightactivationpropertytype.LightActivation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := lightactivationpropertytype.LightActivation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LightActivation = %s", lightactivationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, lightactivationtype := range __gong__sortStageSetInstances(stage.LightActivationTypes, stage.LightActivationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lightactivationtypeIdent := lightactivationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightActivationType{Name: %s}).Stage(stage)", lightactivationtypeIdent, __gong__toRawStringLiteral(lightactivationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightactivationtypeIdent, __gong__toRawStringLiteral(lightactivationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", lightactivationtypeIdent, __gong__toRawStringLiteral(lightactivationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", lightactivationtypeIdent, __gong__toRawStringLiteral(lightactivationtype.Description)))
			for _, elem := range lightactivationtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", lightactivationtypeIdent, lightactivationtypeIdent, targetIdent))
			}
			for _, elem := range lightactivationtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", lightactivationtypeIdent, lightactivationtypeIdent, targetIdent))
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
		for _, nitrogenpropertytype := range __gong__sortStageSetInstances(stage.NitrogenPropertyTypes, stage.NitrogenPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			nitrogenpropertytypeIdent := nitrogenpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NitrogenPropertyType{Name: %s}).Stage(stage)", nitrogenpropertytypeIdent, __gong__toRawStringLiteral(nitrogenpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nitrogenpropertytypeIdent, __gong__toRawStringLiteral(nitrogenpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", nitrogenpropertytypeIdent, __gong__toRawStringLiteral(nitrogenpropertytype.NilReason)))
			if nitrogenpropertytype.Nitrogen != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := nitrogenpropertytype.Nitrogen.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Nitrogen = %s", nitrogenpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, nitrogentype := range __gong__sortStageSetInstances(stage.NitrogenTypes, stage.NitrogenType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			nitrogentypeIdent := nitrogentype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NitrogenType{Name: %s}).Stage(stage)", nitrogentypeIdent, __gong__toRawStringLiteral(nitrogentype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nitrogentypeIdent, __gong__toRawStringLiteral(nitrogentype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", nitrogentypeIdent, __gong__toRawStringLiteral(nitrogentype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", nitrogentypeIdent, __gong__toRawStringLiteral(nitrogentype.Description)))
			for _, elem := range nitrogentype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", nitrogentypeIdent, nitrogentypeIdent, targetIdent))
			}
			for _, elem := range nitrogentype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", nitrogentypeIdent, nitrogentypeIdent, targetIdent))
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
		for _, oilpropertytype := range __gong__sortStageSetInstances(stage.OilPropertyTypes, stage.OilPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			oilpropertytypeIdent := oilpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OilPropertyType{Name: %s}).Stage(stage)", oilpropertytypeIdent, __gong__toRawStringLiteral(oilpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", oilpropertytypeIdent, __gong__toRawStringLiteral(oilpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", oilpropertytypeIdent, __gong__toRawStringLiteral(oilpropertytype.NilReason)))
			if oilpropertytype.Oil != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := oilpropertytype.Oil.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Oil = %s", oilpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, oiltype := range __gong__sortStageSetInstances(stage.OilTypes, stage.OilType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			oiltypeIdent := oiltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OilType{Name: %s}).Stage(stage)", oiltypeIdent, __gong__toRawStringLiteral(oiltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", oiltypeIdent, __gong__toRawStringLiteral(oiltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", oiltypeIdent, __gong__toRawStringLiteral(oiltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", oiltypeIdent, __gong__toRawStringLiteral(oiltype.Description)))
			for _, elem := range oiltype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", oiltypeIdent, oiltypeIdent, targetIdent))
			}
			for _, elem := range oiltype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", oiltypeIdent, oiltypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, onlinecontactpropertytype := range __gong__sortStageSetInstances(stage.OnlineContactPropertyTypes, stage.OnlineContactPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			onlinecontactpropertytypeIdent := onlinecontactpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OnlineContactPropertyType{Name: %s}).Stage(stage)", onlinecontactpropertytypeIdent, __gong__toRawStringLiteral(onlinecontactpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", onlinecontactpropertytypeIdent, __gong__toRawStringLiteral(onlinecontactpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", onlinecontactpropertytypeIdent, __gong__toRawStringLiteral(onlinecontactpropertytype.NilReason)))
			if onlinecontactpropertytype.OnlineContact != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := onlinecontactpropertytype.OnlineContact.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OnlineContact = %s", onlinecontactpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, onlinecontacttype := range __gong__sortStageSetInstances(stage.OnlineContactTypes, stage.OnlineContactType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			onlinecontacttypeIdent := onlinecontacttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OnlineContactType{Name: %s}).Stage(stage)", onlinecontacttypeIdent, __gong__toRawStringLiteral(onlinecontacttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", onlinecontacttypeIdent, __gong__toRawStringLiteral(onlinecontacttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", onlinecontacttypeIdent, __gong__toRawStringLiteral(onlinecontacttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", onlinecontacttypeIdent, __gong__toRawStringLiteral(onlinecontacttype.Description)))
			for _, elem := range onlinecontacttype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", onlinecontacttypeIdent, onlinecontacttypeIdent, targetIdent))
			}
			for _, elem := range onlinecontacttype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", onlinecontacttypeIdent, onlinecontacttypeIdent, targetIdent))
			}
			for _, elem := range onlinecontacttype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", onlinecontacttypeIdent, onlinecontacttypeIdent, targetIdent))
			}
			for _, elem := range onlinecontacttype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", onlinecontacttypeIdent, onlinecontacttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, organisationauthorityassociationpropertytype := range __gong__sortStageSetInstances(stage.OrganisationAuthorityAssociationPropertyTypes, stage.OrganisationAuthorityAssociationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			organisationauthorityassociationpropertytypeIdent := organisationauthorityassociationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityAssociationPropertyType{Name: %s}).Stage(stage)", organisationauthorityassociationpropertytypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthorityassociationpropertytypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", organisationauthorityassociationpropertytypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationpropertytype.NilReason)))
			if organisationauthorityassociationpropertytype.OrganisationAuthorityAssociation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := organisationauthorityassociationpropertytype.OrganisationAuthorityAssociation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OrganisationAuthorityAssociation = %s", organisationauthorityassociationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, organisationauthorityassociationtype := range __gong__sortStageSetInstances(stage.OrganisationAuthorityAssociationTypes, stage.OrganisationAuthorityAssociationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			organisationauthorityassociationtypeIdent := organisationauthorityassociationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityAssociationType{Name: %s}).Stage(stage)", organisationauthorityassociationtypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthorityassociationtypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", organisationauthorityassociationtypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", organisationauthorityassociationtypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationtype.Description)))
			for _, elem := range organisationauthorityassociationtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", organisationauthorityassociationtypeIdent, organisationauthorityassociationtypeIdent, targetIdent))
			}
			if organisationauthorityassociationtype.TheOrganisationAuthority != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := organisationauthorityassociationtype.TheOrganisationAuthority.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheOrganisationAuthority = %s", organisationauthorityassociationtypeIdent, targetIdent))
			}
			for _, elem := range organisationauthorityassociationtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", organisationauthorityassociationtypeIdent, organisationauthorityassociationtypeIdent, targetIdent))
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
		for _, organisationauthoritytimeslicepropertytype := range __gong__sortStageSetInstances(stage.OrganisationAuthorityTimeSlicePropertyTypes, stage.OrganisationAuthorityTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			organisationauthoritytimeslicepropertytypeIdent := organisationauthoritytimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityTimeSlicePropertyType{Name: %s}).Stage(stage)", organisationauthoritytimeslicepropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthoritytimeslicepropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", organisationauthoritytimeslicepropertytypeIdent, organisationauthoritytimeslicepropertytype.Owns))
			if organisationauthoritytimeslicepropertytype.OrganisationAuthorityTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := organisationauthoritytimeslicepropertytype.OrganisationAuthorityTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OrganisationAuthorityTimeSlice = %s", organisationauthoritytimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, organisationauthoritytimeslicetype := range __gong__sortStageSetInstances(stage.OrganisationAuthorityTimeSliceTypes, stage.OrganisationAuthorityTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			organisationauthoritytimeslicetypeIdent := organisationauthoritytimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityTimeSliceType{Name: %s}).Stage(stage)", organisationauthoritytimeslicetypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthoritytimeslicetypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", organisationauthoritytimeslicetypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", organisationauthoritytimeslicetypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", organisationauthoritytimeslicetypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetype.CorrectionNumber))
			for _, elem := range organisationauthoritytimeslicetype.Contact {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			for _, elem := range organisationauthoritytimeslicetype.RelatedOrganisationAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RelatedOrganisationAuthority = append(%s.RelatedOrganisationAuthority, %s)", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			for _, elem := range organisationauthoritytimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			for _, elem := range organisationauthoritytimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, organisationauthoritytype := range __gong__sortStageSetInstances(stage.OrganisationAuthorityTypes, stage.OrganisationAuthorityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			organisationauthoritytypeIdent := organisationauthoritytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityType{Name: %s}).Stage(stage)", organisationauthoritytypeIdent, __gong__toRawStringLiteral(organisationauthoritytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthoritytypeIdent, __gong__toRawStringLiteral(organisationauthoritytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", organisationauthoritytypeIdent, __gong__toRawStringLiteral(organisationauthoritytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", organisationauthoritytypeIdent, __gong__toRawStringLiteral(organisationauthoritytype.Description)))
			for _, elem := range organisationauthoritytype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", organisationauthoritytypeIdent, organisationauthoritytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, oxygenpropertytype := range __gong__sortStageSetInstances(stage.OxygenPropertyTypes, stage.OxygenPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			oxygenpropertytypeIdent := oxygenpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OxygenPropertyType{Name: %s}).Stage(stage)", oxygenpropertytypeIdent, __gong__toRawStringLiteral(oxygenpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", oxygenpropertytypeIdent, __gong__toRawStringLiteral(oxygenpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", oxygenpropertytypeIdent, __gong__toRawStringLiteral(oxygenpropertytype.NilReason)))
			if oxygenpropertytype.Oxygen != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := oxygenpropertytype.Oxygen.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Oxygen = %s", oxygenpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, oxygentype := range __gong__sortStageSetInstances(stage.OxygenTypes, stage.OxygenType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			oxygentypeIdent := oxygentype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OxygenType{Name: %s}).Stage(stage)", oxygentypeIdent, __gong__toRawStringLiteral(oxygentype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", oxygentypeIdent, __gong__toRawStringLiteral(oxygentype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", oxygentypeIdent, __gong__toRawStringLiteral(oxygentype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", oxygentypeIdent, __gong__toRawStringLiteral(oxygentype.Description)))
			for _, elem := range oxygentype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", oxygentypeIdent, oxygentypeIdent, targetIdent))
			}
			for _, elem := range oxygentype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", oxygentypeIdent, oxygentypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, passengerservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.PassengerServiceTimeSlicePropertyTypes, stage.PassengerServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			passengerservicetimeslicepropertytypeIdent := passengerservicetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PassengerServiceTimeSlicePropertyType{Name: %s}).Stage(stage)", passengerservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", passengerservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", passengerservicetimeslicepropertytypeIdent, passengerservicetimeslicepropertytype.Owns))
			if passengerservicetimeslicepropertytype.PassengerServiceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := passengerservicetimeslicepropertytype.PassengerServiceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PassengerServiceTimeSlice = %s", passengerservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, passengerservicetimeslicetype := range __gong__sortStageSetInstances(stage.PassengerServiceTimeSliceTypes, stage.PassengerServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			passengerservicetimeslicetypeIdent := passengerservicetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PassengerServiceTimeSliceType{Name: %s}).Stage(stage)", passengerservicetimeslicetypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", passengerservicetimeslicetypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", passengerservicetimeslicetypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", passengerservicetimeslicetypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", passengerservicetimeslicetypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", passengerservicetimeslicetypeIdent, passengerservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", passengerservicetimeslicetypeIdent, passengerservicetimeslicetype.CorrectionNumber))
			if passengerservicetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := passengerservicetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			if passengerservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := passengerservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.Call_sign {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.AirportHeliport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, passengerservicetype := range __gong__sortStageSetInstances(stage.PassengerServiceTypes, stage.PassengerServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			passengerservicetypeIdent := passengerservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PassengerServiceType{Name: %s}).Stage(stage)", passengerservicetypeIdent, __gong__toRawStringLiteral(passengerservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", passengerservicetypeIdent, __gong__toRawStringLiteral(passengerservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", passengerservicetypeIdent, __gong__toRawStringLiteral(passengerservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", passengerservicetypeIdent, __gong__toRawStringLiteral(passengerservicetype.Description)))
			for _, elem := range passengerservicetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", passengerservicetypeIdent, passengerservicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pilotcontrolledlightingtimeslicepropertytype := range __gong__sortStageSetInstances(stage.PilotControlledLightingTimeSlicePropertyTypes, stage.PilotControlledLightingTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pilotcontrolledlightingtimeslicepropertytypeIdent := pilotcontrolledlightingtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PilotControlledLightingTimeSlicePropertyType{Name: %s}).Stage(stage)", pilotcontrolledlightingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pilotcontrolledlightingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", pilotcontrolledlightingtimeslicepropertytypeIdent, pilotcontrolledlightingtimeslicepropertytype.Owns))
			if pilotcontrolledlightingtimeslicepropertytype.PilotControlledLightingTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := pilotcontrolledlightingtimeslicepropertytype.PilotControlledLightingTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PilotControlledLightingTimeSlice = %s", pilotcontrolledlightingtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pilotcontrolledlightingtimeslicetype := range __gong__sortStageSetInstances(stage.PilotControlledLightingTimeSliceTypes, stage.PilotControlledLightingTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pilotcontrolledlightingtimeslicetypeIdent := pilotcontrolledlightingtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PilotControlledLightingTimeSliceType{Name: %s}).Stage(stage)", pilotcontrolledlightingtimeslicetypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pilotcontrolledlightingtimeslicetypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pilotcontrolledlightingtimeslicetypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pilotcontrolledlightingtimeslicetypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", pilotcontrolledlightingtimeslicetypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetype.CorrectionNumber))
			for _, elem := range pilotcontrolledlightingtimeslicetype.ControlledLightIntensity {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlledLightIntensity = append(%s.ControlledLightIntensity, %s)", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range pilotcontrolledlightingtimeslicetype.ActivatedGroundLighting {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActivatedGroundLighting = append(%s.ActivatedGroundLighting, %s)", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range pilotcontrolledlightingtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range pilotcontrolledlightingtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pilotcontrolledlightingtype := range __gong__sortStageSetInstances(stage.PilotControlledLightingTypes, stage.PilotControlledLightingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pilotcontrolledlightingtypeIdent := pilotcontrolledlightingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PilotControlledLightingType{Name: %s}).Stage(stage)", pilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtype.Description)))
			for _, elem := range pilotcontrolledlightingtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", pilotcontrolledlightingtypeIdent, pilotcontrolledlightingtypeIdent, targetIdent))
			}
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
		for _, postaladdresspropertytype := range __gong__sortStageSetInstances(stage.PostalAddressPropertyTypes, stage.PostalAddressPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			postaladdresspropertytypeIdent := postaladdresspropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PostalAddressPropertyType{Name: %s}).Stage(stage)", postaladdresspropertytypeIdent, __gong__toRawStringLiteral(postaladdresspropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", postaladdresspropertytypeIdent, __gong__toRawStringLiteral(postaladdresspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", postaladdresspropertytypeIdent, __gong__toRawStringLiteral(postaladdresspropertytype.NilReason)))
			if postaladdresspropertytype.PostalAddress != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := postaladdresspropertytype.PostalAddress.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PostalAddress = %s", postaladdresspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, postaladdresstype := range __gong__sortStageSetInstances(stage.PostalAddressTypes, stage.PostalAddressType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			postaladdresstypeIdent := postaladdresstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PostalAddressType{Name: %s}).Stage(stage)", postaladdresstypeIdent, __gong__toRawStringLiteral(postaladdresstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", postaladdresstypeIdent, __gong__toRawStringLiteral(postaladdresstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", postaladdresstypeIdent, __gong__toRawStringLiteral(postaladdresstype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", postaladdresstypeIdent, __gong__toRawStringLiteral(postaladdresstype.Description)))
			for _, elem := range postaladdresstype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", postaladdresstypeIdent, postaladdresstypeIdent, targetIdent))
			}
			for _, elem := range postaladdresstype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", postaladdresstypeIdent, postaladdresstypeIdent, targetIdent))
			}
			for _, elem := range postaladdresstype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", postaladdresstypeIdent, postaladdresstypeIdent, targetIdent))
			}
			for _, elem := range postaladdresstype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", postaladdresstypeIdent, postaladdresstypeIdent, targetIdent))
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
		for _, radiocommunicationchannelpropertytype := range __gong__sortStageSetInstances(stage.RadioCommunicationChannelPropertyTypes, stage.RadioCommunicationChannelPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radiocommunicationchannelpropertytypeIdent := radiocommunicationchannelpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationChannelPropertyType{Name: %s}).Stage(stage)", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", radiocommunicationchannelpropertytypeIdent, radiocommunicationchannelpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, radiocommunicationchanneltimeslicepropertytype := range __gong__sortStageSetInstances(stage.RadioCommunicationChannelTimeSlicePropertyTypes, stage.RadioCommunicationChannelTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radiocommunicationchanneltimeslicepropertytypeIdent := radiocommunicationchanneltimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationChannelTimeSlicePropertyType{Name: %s}).Stage(stage)", radiocommunicationchanneltimeslicepropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationchanneltimeslicepropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", radiocommunicationchanneltimeslicepropertytypeIdent, radiocommunicationchanneltimeslicepropertytype.Owns))
			if radiocommunicationchanneltimeslicepropertytype.RadioCommunicationChannelTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := radiocommunicationchanneltimeslicepropertytype.RadioCommunicationChannelTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunicationChannelTimeSlice = %s", radiocommunicationchanneltimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, radiocommunicationchanneltimeslicetype := range __gong__sortStageSetInstances(stage.RadioCommunicationChannelTimeSliceTypes, stage.RadioCommunicationChannelTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radiocommunicationchanneltimeslicetypeIdent := radiocommunicationchanneltimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationChannelTimeSliceType{Name: %s}).Stage(stage)", radiocommunicationchanneltimeslicetypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationchanneltimeslicetypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radiocommunicationchanneltimeslicetypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radiocommunicationchanneltimeslicetypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", radiocommunicationchanneltimeslicetypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetype.CorrectionNumber))
			for _, elem := range radiocommunicationchanneltimeslicetype.Location {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = append(%s.Location, %s)", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationchanneltimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationchanneltimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationchanneltimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, radiocommunicationchanneltype := range __gong__sortStageSetInstances(stage.RadioCommunicationChannelTypes, stage.RadioCommunicationChannelType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radiocommunicationchanneltypeIdent := radiocommunicationchanneltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationChannelType{Name: %s}).Stage(stage)", radiocommunicationchanneltypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationchanneltypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radiocommunicationchanneltypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radiocommunicationchanneltypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltype.Description)))
			for _, elem := range radiocommunicationchanneltype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", radiocommunicationchanneltypeIdent, radiocommunicationchanneltypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, radiocommunicationoperationalstatuspropertytype := range __gong__sortStageSetInstances(stage.RadioCommunicationOperationalStatusPropertyTypes, stage.RadioCommunicationOperationalStatusPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radiocommunicationoperationalstatuspropertytypeIdent := radiocommunicationoperationalstatuspropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationOperationalStatusPropertyType{Name: %s}).Stage(stage)", radiocommunicationoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatuspropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", radiocommunicationoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatuspropertytype.NilReason)))
			if radiocommunicationoperationalstatuspropertytype.RadioCommunicationOperationalStatus != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := radiocommunicationoperationalstatuspropertytype.RadioCommunicationOperationalStatus.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunicationOperationalStatus = %s", radiocommunicationoperationalstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, radiocommunicationoperationalstatustype := range __gong__sortStageSetInstances(stage.RadioCommunicationOperationalStatusTypes, stage.RadioCommunicationOperationalStatusType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radiocommunicationoperationalstatustypeIdent := radiocommunicationoperationalstatustype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationOperationalStatusType{Name: %s}).Stage(stage)", radiocommunicationoperationalstatustypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatustype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationoperationalstatustypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radiocommunicationoperationalstatustypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radiocommunicationoperationalstatustypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatustype.Description)))
			for _, elem := range radiocommunicationoperationalstatustype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", radiocommunicationoperationalstatustypeIdent, radiocommunicationoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationoperationalstatustype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", radiocommunicationoperationalstatustypeIdent, radiocommunicationoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationoperationalstatustype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", radiocommunicationoperationalstatustypeIdent, radiocommunicationoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationoperationalstatustype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", radiocommunicationoperationalstatustypeIdent, radiocommunicationoperationalstatustypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routeportionpropertytype := range __gong__sortStageSetInstances(stage.RoutePortionPropertyTypes, stage.RoutePortionPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routeportionpropertytypeIdent := routeportionpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RoutePortionPropertyType{Name: %s}).Stage(stage)", routeportionpropertytypeIdent, __gong__toRawStringLiteral(routeportionpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routeportionpropertytypeIdent, __gong__toRawStringLiteral(routeportionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", routeportionpropertytypeIdent, __gong__toRawStringLiteral(routeportionpropertytype.NilReason)))
			if routeportionpropertytype.RoutePortion != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportionpropertytype.RoutePortion.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RoutePortion = %s", routeportionpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routeportiontype := range __gong__sortStageSetInstances(stage.RoutePortionTypes, stage.RoutePortionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routeportiontypeIdent := routeportiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RoutePortionType{Name: %s}).Stage(stage)", routeportiontypeIdent, __gong__toRawStringLiteral(routeportiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routeportiontypeIdent, __gong__toRawStringLiteral(routeportiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routeportiontypeIdent, __gong__toRawStringLiteral(routeportiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routeportiontypeIdent, __gong__toRawStringLiteral(routeportiontype.Description)))
			if routeportiontype.Start_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.Start_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_fixDesignatedPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.Start_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_navaidSystem = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.Start_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_position = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.Start_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_runwayPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.Start_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_aimingPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.Start_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_airportReferencePoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.IntermediatePoint_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_fixDesignatedPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.IntermediatePoint_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_navaidSystem = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.IntermediatePoint_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_position = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.IntermediatePoint_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_runwayPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.IntermediatePoint_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_aimingPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.IntermediatePoint_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_airportReferencePoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.ReferencedRoute != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.ReferencedRoute.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedRoute = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.End_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_fixDesignatedPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.End_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_navaidSystem = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.End_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_position = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.End_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_runwayPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.End_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_aimingPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeportiontype.End_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_airportReferencePoint = %s", routeportiontypeIdent, targetIdent))
			}
			for _, elem := range routeportiontype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", routeportiontypeIdent, routeportiontypeIdent, targetIdent))
			}
			for _, elem := range routeportiontype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", routeportiontypeIdent, routeportiontypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routepropertytype := range __gong__sortStageSetInstances(stage.RoutePropertyTypes, stage.RoutePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routepropertytypeIdent := routepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RoutePropertyType{Name: %s}).Stage(stage)", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", routepropertytypeIdent, routepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Actuate)))
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
		for _, searchrescueservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.SearchRescueServiceTimeSlicePropertyTypes, stage.SearchRescueServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			searchrescueservicetimeslicepropertytypeIdent := searchrescueservicetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SearchRescueServiceTimeSlicePropertyType{Name: %s}).Stage(stage)", searchrescueservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", searchrescueservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", searchrescueservicetimeslicepropertytypeIdent, searchrescueservicetimeslicepropertytype.Owns))
			if searchrescueservicetimeslicepropertytype.SearchRescueServiceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := searchrescueservicetimeslicepropertytype.SearchRescueServiceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SearchRescueServiceTimeSlice = %s", searchrescueservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, searchrescueservicetimeslicetype := range __gong__sortStageSetInstances(stage.SearchRescueServiceTimeSliceTypes, stage.SearchRescueServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			searchrescueservicetimeslicetypeIdent := searchrescueservicetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SearchRescueServiceTimeSliceType{Name: %s}).Stage(stage)", searchrescueservicetimeslicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", searchrescueservicetimeslicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", searchrescueservicetimeslicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", searchrescueservicetimeslicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", searchrescueservicetimeslicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetype.CorrectionNumber))
			if searchrescueservicetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := searchrescueservicetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			if searchrescueservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := searchrescueservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.Call_sign {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.ClientAirspace {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirspace = append(%s.ClientAirspace, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.ClientRoute {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientRoute = append(%s.ClientRoute, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, searchrescueservicetype := range __gong__sortStageSetInstances(stage.SearchRescueServiceTypes, stage.SearchRescueServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			searchrescueservicetypeIdent := searchrescueservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SearchRescueServiceType{Name: %s}).Stage(stage)", searchrescueservicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", searchrescueservicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", searchrescueservicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", searchrescueservicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetype.Description)))
			for _, elem := range searchrescueservicetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", searchrescueservicetypeIdent, searchrescueservicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, serviceoperationalstatuspropertytype := range __gong__sortStageSetInstances(stage.ServiceOperationalStatusPropertyTypes, stage.ServiceOperationalStatusPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			serviceoperationalstatuspropertytypeIdent := serviceoperationalstatuspropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ServiceOperationalStatusPropertyType{Name: %s}).Stage(stage)", serviceoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(serviceoperationalstatuspropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", serviceoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(serviceoperationalstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", serviceoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(serviceoperationalstatuspropertytype.NilReason)))
			if serviceoperationalstatuspropertytype.ServiceOperationalStatus != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := serviceoperationalstatuspropertytype.ServiceOperationalStatus.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceOperationalStatus = %s", serviceoperationalstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, serviceoperationalstatustype := range __gong__sortStageSetInstances(stage.ServiceOperationalStatusTypes, stage.ServiceOperationalStatusType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			serviceoperationalstatustypeIdent := serviceoperationalstatustype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ServiceOperationalStatusType{Name: %s}).Stage(stage)", serviceoperationalstatustypeIdent, __gong__toRawStringLiteral(serviceoperationalstatustype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", serviceoperationalstatustypeIdent, __gong__toRawStringLiteral(serviceoperationalstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", serviceoperationalstatustypeIdent, __gong__toRawStringLiteral(serviceoperationalstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", serviceoperationalstatustypeIdent, __gong__toRawStringLiteral(serviceoperationalstatustype.Description)))
			for _, elem := range serviceoperationalstatustype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", serviceoperationalstatustypeIdent, serviceoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range serviceoperationalstatustype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", serviceoperationalstatustypeIdent, serviceoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range serviceoperationalstatustype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", serviceoperationalstatustypeIdent, serviceoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range serviceoperationalstatustype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", serviceoperationalstatustypeIdent, serviceoperationalstatustypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialdatetimeslicepropertytype := range __gong__sortStageSetInstances(stage.SpecialDateTimeSlicePropertyTypes, stage.SpecialDateTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialdatetimeslicepropertytypeIdent := specialdatetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialDateTimeSlicePropertyType{Name: %s}).Stage(stage)", specialdatetimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialdatetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialdatetimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialdatetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", specialdatetimeslicepropertytypeIdent, specialdatetimeslicepropertytype.Owns))
			if specialdatetimeslicepropertytype.SpecialDateTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := specialdatetimeslicepropertytype.SpecialDateTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateTimeSlice = %s", specialdatetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialdatetimeslicetype := range __gong__sortStageSetInstances(stage.SpecialDateTimeSliceTypes, stage.SpecialDateTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialdatetimeslicetypeIdent := specialdatetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialDateTimeSliceType{Name: %s}).Stage(stage)", specialdatetimeslicetypeIdent, __gong__toRawStringLiteral(specialdatetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialdatetimeslicetypeIdent, __gong__toRawStringLiteral(specialdatetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialdatetimeslicetypeIdent, __gong__toRawStringLiteral(specialdatetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialdatetimeslicetypeIdent, __gong__toRawStringLiteral(specialdatetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", specialdatetimeslicetypeIdent, __gong__toRawStringLiteral(specialdatetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", specialdatetimeslicetypeIdent, specialdatetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", specialdatetimeslicetypeIdent, specialdatetimeslicetype.CorrectionNumber))
			for _, elem := range specialdatetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", specialdatetimeslicetypeIdent, specialdatetimeslicetypeIdent, targetIdent))
			}
			if specialdatetimeslicetype.Authority != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := specialdatetimeslicetype.Authority.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = %s", specialdatetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialdatetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", specialdatetimeslicetypeIdent, specialdatetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialdatetype := range __gong__sortStageSetInstances(stage.SpecialDateTypes, stage.SpecialDateType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialdatetypeIdent := specialdatetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialDateType{Name: %s}).Stage(stage)", specialdatetypeIdent, __gong__toRawStringLiteral(specialdatetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialdatetypeIdent, __gong__toRawStringLiteral(specialdatetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialdatetypeIdent, __gong__toRawStringLiteral(specialdatetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialdatetypeIdent, __gong__toRawStringLiteral(specialdatetype.Description)))
			for _, elem := range specialdatetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", specialdatetypeIdent, specialdatetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardlevelcolumnpropertytype := range __gong__sortStageSetInstances(stage.StandardLevelColumnPropertyTypes, stage.StandardLevelColumnPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardlevelcolumnpropertytypeIdent := standardlevelcolumnpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelColumnPropertyType{Name: %s}).Stage(stage)", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardlevelcolumnpropertytypeIdent, standardlevelcolumnpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, standardlevelcolumntimeslicepropertytype := range __gong__sortStageSetInstances(stage.StandardLevelColumnTimeSlicePropertyTypes, stage.StandardLevelColumnTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardlevelcolumntimeslicepropertytypeIdent := standardlevelcolumntimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelColumnTimeSlicePropertyType{Name: %s}).Stage(stage)", standardlevelcolumntimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelcolumntimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardlevelcolumntimeslicepropertytypeIdent, standardlevelcolumntimeslicepropertytype.Owns))
			if standardlevelcolumntimeslicepropertytype.StandardLevelColumnTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardlevelcolumntimeslicepropertytype.StandardLevelColumnTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevelColumnTimeSlice = %s", standardlevelcolumntimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardlevelcolumntimeslicetype := range __gong__sortStageSetInstances(stage.StandardLevelColumnTimeSliceTypes, stage.StandardLevelColumnTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardlevelcolumntimeslicetypeIdent := standardlevelcolumntimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelColumnTimeSliceType{Name: %s}).Stage(stage)", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetype.CorrectionNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.UnitOfMeasurement = %s", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.UnitOfMeasurement)))
			for _, elem := range standardlevelcolumntimeslicetype.Level {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = append(%s.Level, %s)", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			if standardlevelcolumntimeslicetype.LevelTable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardlevelcolumntimeslicetype.LevelTable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LevelTable = %s", standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelcolumntimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelcolumntimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardlevelcolumntype := range __gong__sortStageSetInstances(stage.StandardLevelColumnTypes, stage.StandardLevelColumnType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardlevelcolumntypeIdent := standardlevelcolumntype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelColumnType{Name: %s}).Stage(stage)", standardlevelcolumntypeIdent, __gong__toRawStringLiteral(standardlevelcolumntype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelcolumntypeIdent, __gong__toRawStringLiteral(standardlevelcolumntype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardlevelcolumntypeIdent, __gong__toRawStringLiteral(standardlevelcolumntype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardlevelcolumntypeIdent, __gong__toRawStringLiteral(standardlevelcolumntype.Description)))
			for _, elem := range standardlevelcolumntype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", standardlevelcolumntypeIdent, standardlevelcolumntypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardlevelpropertytype := range __gong__sortStageSetInstances(stage.StandardLevelPropertyTypes, stage.StandardLevelPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardlevelpropertytypeIdent := standardlevelpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelPropertyType{Name: %s}).Stage(stage)", standardlevelpropertytypeIdent, __gong__toRawStringLiteral(standardlevelpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelpropertytypeIdent, __gong__toRawStringLiteral(standardlevelpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", standardlevelpropertytypeIdent, __gong__toRawStringLiteral(standardlevelpropertytype.NilReason)))
			if standardlevelpropertytype.StandardLevel != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardlevelpropertytype.StandardLevel.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevel = %s", standardlevelpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardlevelsectortimeslicepropertytype := range __gong__sortStageSetInstances(stage.StandardLevelSectorTimeSlicePropertyTypes, stage.StandardLevelSectorTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardlevelsectortimeslicepropertytypeIdent := standardlevelsectortimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelSectorTimeSlicePropertyType{Name: %s}).Stage(stage)", standardlevelsectortimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelsectortimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardlevelsectortimeslicepropertytypeIdent, standardlevelsectortimeslicepropertytype.Owns))
			if standardlevelsectortimeslicepropertytype.StandardLevelSectorTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardlevelsectortimeslicepropertytype.StandardLevelSectorTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevelSectorTimeSlice = %s", standardlevelsectortimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardlevelsectortimeslicetype := range __gong__sortStageSetInstances(stage.StandardLevelSectorTimeSliceTypes, stage.StandardLevelSectorTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardlevelsectortimeslicetypeIdent := standardlevelsectortimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelSectorTimeSliceType{Name: %s}).Stage(stage)", standardlevelsectortimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelsectortimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardlevelsectortimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardlevelsectortimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", standardlevelsectortimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetype.CorrectionNumber))
			for _, elem := range standardlevelsectortimeslicetype.ApplicableAirspace {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApplicableAirspace = append(%s.ApplicableAirspace, %s)", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			if standardlevelsectortimeslicetype.ApplicableLevelColumn != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardlevelsectortimeslicetype.ApplicableLevelColumn.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApplicableLevelColumn = %s", standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelsectortimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelsectortimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardlevelsectortype := range __gong__sortStageSetInstances(stage.StandardLevelSectorTypes, stage.StandardLevelSectorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardlevelsectortypeIdent := standardlevelsectortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelSectorType{Name: %s}).Stage(stage)", standardlevelsectortypeIdent, __gong__toRawStringLiteral(standardlevelsectortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelsectortypeIdent, __gong__toRawStringLiteral(standardlevelsectortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardlevelsectortypeIdent, __gong__toRawStringLiteral(standardlevelsectortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardlevelsectortypeIdent, __gong__toRawStringLiteral(standardlevelsectortype.Description)))
			for _, elem := range standardlevelsectortype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", standardlevelsectortypeIdent, standardlevelsectortypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardleveltablepropertytype := range __gong__sortStageSetInstances(stage.StandardLevelTablePropertyTypes, stage.StandardLevelTablePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardleveltablepropertytypeIdent := standardleveltablepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelTablePropertyType{Name: %s}).Stage(stage)", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardleveltablepropertytypeIdent, standardleveltablepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, standardleveltabletimeslicepropertytype := range __gong__sortStageSetInstances(stage.StandardLevelTableTimeSlicePropertyTypes, stage.StandardLevelTableTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardleveltabletimeslicepropertytypeIdent := standardleveltabletimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelTableTimeSlicePropertyType{Name: %s}).Stage(stage)", standardleveltabletimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardleveltabletimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardleveltabletimeslicepropertytypeIdent, standardleveltabletimeslicepropertytype.Owns))
			if standardleveltabletimeslicepropertytype.StandardLevelTableTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := standardleveltabletimeslicepropertytype.StandardLevelTableTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevelTableTimeSlice = %s", standardleveltabletimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardleveltabletimeslicetype := range __gong__sortStageSetInstances(stage.StandardLevelTableTimeSliceTypes, stage.StandardLevelTableTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardleveltabletimeslicetypeIdent := standardleveltabletimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelTableTimeSliceType{Name: %s}).Stage(stage)", standardleveltabletimeslicetypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardleveltabletimeslicetypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardleveltabletimeslicetypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardleveltabletimeslicetypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", standardleveltabletimeslicetypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", standardleveltabletimeslicetypeIdent, standardleveltabletimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", standardleveltabletimeslicetypeIdent, standardleveltabletimeslicetype.CorrectionNumber))
			for _, elem := range standardleveltabletimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", standardleveltabletimeslicetypeIdent, standardleveltabletimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardleveltabletimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", standardleveltabletimeslicetypeIdent, standardleveltabletimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardleveltabletype := range __gong__sortStageSetInstances(stage.StandardLevelTableTypes, stage.StandardLevelTableType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardleveltabletypeIdent := standardleveltabletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelTableType{Name: %s}).Stage(stage)", standardleveltabletypeIdent, __gong__toRawStringLiteral(standardleveltabletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardleveltabletypeIdent, __gong__toRawStringLiteral(standardleveltabletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardleveltabletypeIdent, __gong__toRawStringLiteral(standardleveltabletype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardleveltabletypeIdent, __gong__toRawStringLiteral(standardleveltabletype.Description)))
			for _, elem := range standardleveltabletype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", standardleveltabletypeIdent, standardleveltabletypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, standardleveltype := range __gong__sortStageSetInstances(stage.StandardLevelTypes, stage.StandardLevelType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardleveltypeIdent := standardleveltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelType{Name: %s}).Stage(stage)", standardleveltypeIdent, __gong__toRawStringLiteral(standardleveltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardleveltypeIdent, __gong__toRawStringLiteral(standardleveltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardleveltypeIdent, __gong__toRawStringLiteral(standardleveltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardleveltypeIdent, __gong__toRawStringLiteral(standardleveltype.Description)))
			for _, elem := range standardleveltype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", standardleveltypeIdent, standardleveltypeIdent, targetIdent))
			}
			for _, elem := range standardleveltype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", standardleveltypeIdent, standardleveltypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, telephonecontactpropertytype := range __gong__sortStageSetInstances(stage.TelephoneContactPropertyTypes, stage.TelephoneContactPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			telephonecontactpropertytypeIdent := telephonecontactpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TelephoneContactPropertyType{Name: %s}).Stage(stage)", telephonecontactpropertytypeIdent, __gong__toRawStringLiteral(telephonecontactpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", telephonecontactpropertytypeIdent, __gong__toRawStringLiteral(telephonecontactpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", telephonecontactpropertytypeIdent, __gong__toRawStringLiteral(telephonecontactpropertytype.NilReason)))
			if telephonecontactpropertytype.TelephoneContact != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := telephonecontactpropertytype.TelephoneContact.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TelephoneContact = %s", telephonecontactpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, telephonecontacttype := range __gong__sortStageSetInstances(stage.TelephoneContactTypes, stage.TelephoneContactType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			telephonecontacttypeIdent := telephonecontacttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TelephoneContactType{Name: %s}).Stage(stage)", telephonecontacttypeIdent, __gong__toRawStringLiteral(telephonecontacttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", telephonecontacttypeIdent, __gong__toRawStringLiteral(telephonecontacttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", telephonecontacttypeIdent, __gong__toRawStringLiteral(telephonecontacttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", telephonecontacttypeIdent, __gong__toRawStringLiteral(telephonecontacttype.Description)))
			for _, elem := range telephonecontacttype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", telephonecontacttypeIdent, telephonecontacttypeIdent, targetIdent))
			}
			for _, elem := range telephonecontacttype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", telephonecontacttypeIdent, telephonecontacttypeIdent, targetIdent))
			}
			for _, elem := range telephonecontacttype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", telephonecontacttypeIdent, telephonecontacttypeIdent, targetIdent))
			}
			for _, elem := range telephonecontacttype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", telephonecontacttypeIdent, telephonecontacttypeIdent, targetIdent))
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
		for _, unitavailabilitypropertytype := range __gong__sortStageSetInstances(stage.UnitAvailabilityPropertyTypes, stage.UnitAvailabilityPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unitavailabilitypropertytypeIdent := unitavailabilitypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitAvailabilityPropertyType{Name: %s}).Stage(stage)", unitavailabilitypropertytypeIdent, __gong__toRawStringLiteral(unitavailabilitypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unitavailabilitypropertytypeIdent, __gong__toRawStringLiteral(unitavailabilitypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", unitavailabilitypropertytypeIdent, __gong__toRawStringLiteral(unitavailabilitypropertytype.NilReason)))
			if unitavailabilitypropertytype.UnitAvailability != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := unitavailabilitypropertytype.UnitAvailability.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UnitAvailability = %s", unitavailabilitypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, unitavailabilitytype := range __gong__sortStageSetInstances(stage.UnitAvailabilityTypes, stage.UnitAvailabilityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unitavailabilitytypeIdent := unitavailabilitytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitAvailabilityType{Name: %s}).Stage(stage)", unitavailabilitytypeIdent, __gong__toRawStringLiteral(unitavailabilitytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unitavailabilitytypeIdent, __gong__toRawStringLiteral(unitavailabilitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", unitavailabilitytypeIdent, __gong__toRawStringLiteral(unitavailabilitytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", unitavailabilitytypeIdent, __gong__toRawStringLiteral(unitavailabilitytype.Description)))
			for _, elem := range unitavailabilitytype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", unitavailabilitytypeIdent, unitavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range unitavailabilitytype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", unitavailabilitytypeIdent, unitavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range unitavailabilitytype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", unitavailabilitytypeIdent, unitavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range unitavailabilitytype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", unitavailabilitytypeIdent, unitavailabilitytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, unitdependencypropertytype := range __gong__sortStageSetInstances(stage.UnitDependencyPropertyTypes, stage.UnitDependencyPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unitdependencypropertytypeIdent := unitdependencypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitDependencyPropertyType{Name: %s}).Stage(stage)", unitdependencypropertytypeIdent, __gong__toRawStringLiteral(unitdependencypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unitdependencypropertytypeIdent, __gong__toRawStringLiteral(unitdependencypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", unitdependencypropertytypeIdent, __gong__toRawStringLiteral(unitdependencypropertytype.NilReason)))
			if unitdependencypropertytype.UnitDependency != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := unitdependencypropertytype.UnitDependency.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UnitDependency = %s", unitdependencypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, unitdependencytype := range __gong__sortStageSetInstances(stage.UnitDependencyTypes, stage.UnitDependencyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unitdependencytypeIdent := unitdependencytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitDependencyType{Name: %s}).Stage(stage)", unitdependencytypeIdent, __gong__toRawStringLiteral(unitdependencytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unitdependencytypeIdent, __gong__toRawStringLiteral(unitdependencytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", unitdependencytypeIdent, __gong__toRawStringLiteral(unitdependencytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", unitdependencytypeIdent, __gong__toRawStringLiteral(unitdependencytype.Description)))
			for _, elem := range unitdependencytype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", unitdependencytypeIdent, unitdependencytypeIdent, targetIdent))
			}
			if unitdependencytype.TheUnit != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := unitdependencytype.TheUnit.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheUnit = %s", unitdependencytypeIdent, targetIdent))
			}
			for _, elem := range unitdependencytype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", unitdependencytypeIdent, unitdependencytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, unitpropertytype := range __gong__sortStageSetInstances(stage.UnitPropertyTypes, stage.UnitPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unitpropertytypeIdent := unitpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitPropertyType{Name: %s}).Stage(stage)", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", unitpropertytypeIdent, unitpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, unittimeslicepropertytype := range __gong__sortStageSetInstances(stage.UnitTimeSlicePropertyTypes, stage.UnitTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unittimeslicepropertytypeIdent := unittimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitTimeSlicePropertyType{Name: %s}).Stage(stage)", unittimeslicepropertytypeIdent, __gong__toRawStringLiteral(unittimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unittimeslicepropertytypeIdent, __gong__toRawStringLiteral(unittimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", unittimeslicepropertytypeIdent, unittimeslicepropertytype.Owns))
			if unittimeslicepropertytype.UnitTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := unittimeslicepropertytype.UnitTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UnitTimeSlice = %s", unittimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, unittimeslicetype := range __gong__sortStageSetInstances(stage.UnitTimeSliceTypes, stage.UnitTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unittimeslicetypeIdent := unittimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitTimeSliceType{Name: %s}).Stage(stage)", unittimeslicetypeIdent, __gong__toRawStringLiteral(unittimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unittimeslicetypeIdent, __gong__toRawStringLiteral(unittimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", unittimeslicetypeIdent, __gong__toRawStringLiteral(unittimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", unittimeslicetypeIdent, __gong__toRawStringLiteral(unittimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", unittimeslicetypeIdent, __gong__toRawStringLiteral(unittimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", unittimeslicetypeIdent, unittimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", unittimeslicetypeIdent, unittimeslicetype.CorrectionNumber))
			if unittimeslicetype.Position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := unittimeslicetype.Position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Position = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.AirportLocation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := unittimeslicetype.AirportLocation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportLocation = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.OwnerOrganisation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := unittimeslicetype.OwnerOrganisation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OwnerOrganisation = %s", unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.Contact {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.RelatedUnit {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RelatedUnit = append(%s.RelatedUnit, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, unittype := range __gong__sortStageSetInstances(stage.UnitTypes, stage.UnitType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unittypeIdent := unittype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitType{Name: %s}).Stage(stage)", unittypeIdent, __gong__toRawStringLiteral(unittype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unittypeIdent, __gong__toRawStringLiteral(unittype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", unittypeIdent, __gong__toRawStringLiteral(unittype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", unittypeIdent, __gong__toRawStringLiteral(unittype.Description)))
			for _, elem := range unittype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", unittypeIdent, unittypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, vorpropertytype := range __gong__sortStageSetInstances(stage.VORPropertyTypes, stage.VORPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vorpropertytypeIdent := vorpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VORPropertyType{Name: %s}).Stage(stage)", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", vorpropertytypeIdent, vorpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.Actuate)))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures\"", fmt.Sprintf("%q", modelsPackageName))
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

func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractairportgroundservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractairportgroundservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractairportgroundservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractairportgroundservicetype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractAirportGroundServiceType", fieldName)
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

func (abstractservicetype *AbstractServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractservicetype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractServiceType", fieldName)
	}
	return
}

func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstracttrafficseparationservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstracttrafficseparationservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstracttrafficseparationservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstracttrafficseparationservicetype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractTrafficSeparationServiceType", fieldName)
	}
	return
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aerialrefuellingpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aerialrefuellingpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", aerialrefuellingpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", aerialrefuellingpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", aerialrefuellingpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", aerialrefuellingpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", aerialrefuellingpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", aerialrefuellingpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", aerialrefuellingpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", aerialrefuellingpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", aerialrefuellingpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingPropertyType", fieldName)
	}
	return
}

func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airtrafficcontrolservicetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airtrafficcontrolservicetimeslicepropertytype.Owns)

	case "AirTrafficControlServiceTimeSlice":
		if airtrafficcontrolservicetimeslicepropertytype.AirTrafficControlServiceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AirTrafficControlServiceTimeSlice", airtrafficcontrolservicetimeslicepropertytype.AirTrafficControlServiceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirTrafficControlServiceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirTrafficControlServiceTimeSlicePropertyType", fieldName)
	}
	return
}

func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airtrafficcontrolservicetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airtrafficcontrolservicetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airtrafficcontrolservicetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airtrafficcontrolservicetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", airtrafficcontrolservicetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", airtrafficcontrolservicetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", airtrafficcontrolservicetimeslicetype.CorrectionNumber)

	case "Location":
		if airtrafficcontrolservicetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", airtrafficcontrolservicetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "ServiceProvider":
		if airtrafficcontrolservicetimeslicetype.ServiceProvider != nil {
			res = __gong__marshallPointer(ident, "ServiceProvider", airtrafficcontrolservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceProvider", "nil")
		}
	case "Call_sign":
		var sb strings.Builder
		for _, _callsigndetailpropertytype := range airtrafficcontrolservicetimeslicetype.Call_sign {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Call_sign", _callsigndetailpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RadioCommunication":
		var sb strings.Builder
		for _, _radiocommunicationchannelpropertytype := range airtrafficcontrolservicetimeslicetype.RadioCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadioCommunication", _radiocommunicationchannelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundCommunication":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range airtrafficcontrolservicetimeslicetype.GroundCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundCommunication", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _serviceoperationalstatuspropertytype := range airtrafficcontrolservicetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _serviceoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airtrafficcontrolservicetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientAirport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range airtrafficcontrolservicetimeslicetype.ClientAirport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientAirport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientAirspace":
		var sb strings.Builder
		for _, _airspacepropertytype := range airtrafficcontrolservicetimeslicetype.ClientAirspace {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientAirspace", _airspacepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientRoute":
		var sb strings.Builder
		for _, _routeportionpropertytype := range airtrafficcontrolservicetimeslicetype.ClientRoute {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientRoute", _routeportionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientProcedure":
		var sb strings.Builder
		for _, _procedurepropertytype := range airtrafficcontrolservicetimeslicetype.ClientProcedure {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientProcedure", _procedurepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientHolding":
		var sb strings.Builder
		for _, _holdingpatternpropertytype := range airtrafficcontrolservicetimeslicetype.ClientHolding {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientHolding", _holdingpatternpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientAerialRefuelling":
		var sb strings.Builder
		for _, _aerialrefuellingpropertytype := range airtrafficcontrolservicetimeslicetype.ClientAerialRefuelling {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientAerialRefuelling", _aerialrefuellingpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AircraftLocator":
		if airtrafficcontrolservicetimeslicetype.AircraftLocator != nil {
			res = __gong__marshallPointer(ident, "AircraftLocator", airtrafficcontrolservicetimeslicetype.AircraftLocator.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AircraftLocator", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_13 := range airtrafficcontrolservicetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_13.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirTrafficControlServiceTimeSliceType", fieldName)
	}
	return
}

func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airtrafficcontrolservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airtrafficcontrolservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airtrafficcontrolservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airtrafficcontrolservicetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _airtrafficcontrolservicetimeslicepropertytype := range airtrafficcontrolservicetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _airtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirTrafficControlServiceType", fieldName)
	}
	return
}

func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airtrafficmanagementservicetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airtrafficmanagementservicetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airtrafficmanagementservicetimeslicepropertytype.Owns)

	case "AirTrafficManagementServiceTimeSlice":
		if airtrafficmanagementservicetimeslicepropertytype.AirTrafficManagementServiceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AirTrafficManagementServiceTimeSlice", airtrafficmanagementservicetimeslicepropertytype.AirTrafficManagementServiceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirTrafficManagementServiceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirTrafficManagementServiceTimeSlicePropertyType", fieldName)
	}
	return
}

func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airtrafficmanagementservicetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airtrafficmanagementservicetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airtrafficmanagementservicetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airtrafficmanagementservicetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", airtrafficmanagementservicetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", airtrafficmanagementservicetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", airtrafficmanagementservicetimeslicetype.CorrectionNumber)

	case "Location":
		if airtrafficmanagementservicetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", airtrafficmanagementservicetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "ServiceProvider":
		if airtrafficmanagementservicetimeslicetype.ServiceProvider != nil {
			res = __gong__marshallPointer(ident, "ServiceProvider", airtrafficmanagementservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceProvider", "nil")
		}
	case "Call_sign":
		var sb strings.Builder
		for _, _callsigndetailpropertytype := range airtrafficmanagementservicetimeslicetype.Call_sign {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Call_sign", _callsigndetailpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RadioCommunication":
		var sb strings.Builder
		for _, _radiocommunicationchannelpropertytype := range airtrafficmanagementservicetimeslicetype.RadioCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadioCommunication", _radiocommunicationchannelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundCommunication":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range airtrafficmanagementservicetimeslicetype.GroundCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundCommunication", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _serviceoperationalstatuspropertytype := range airtrafficmanagementservicetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _serviceoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airtrafficmanagementservicetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientAirspace":
		var sb strings.Builder
		for _, _airspacepropertytype := range airtrafficmanagementservicetimeslicetype.ClientAirspace {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientAirspace", _airspacepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientAerialRefuelling":
		var sb strings.Builder
		for _, _aerialrefuellingpropertytype := range airtrafficmanagementservicetimeslicetype.ClientAerialRefuelling {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientAerialRefuelling", _aerialrefuellingpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientRoute":
		var sb strings.Builder
		for _, _routeportionpropertytype := range airtrafficmanagementservicetimeslicetype.ClientRoute {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientRoute", _routeportionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_14 := range airtrafficmanagementservicetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_14.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirTrafficManagementServiceTimeSliceType", fieldName)
	}
	return
}

func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airtrafficmanagementservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airtrafficmanagementservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airtrafficmanagementservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airtrafficmanagementservicetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _airtrafficmanagementservicetimeslicepropertytype := range airtrafficmanagementservicetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _airtrafficmanagementservicetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirTrafficManagementServiceType", fieldName)
	}
	return
}

func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aircraftgroundservicetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aircraftgroundservicetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", aircraftgroundservicetimeslicepropertytype.Owns)

	case "AircraftGroundServiceTimeSlice":
		if aircraftgroundservicetimeslicepropertytype.AircraftGroundServiceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AircraftGroundServiceTimeSlice", aircraftgroundservicetimeslicepropertytype.AircraftGroundServiceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AircraftGroundServiceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AircraftGroundServiceTimeSlicePropertyType", fieldName)
	}
	return
}

func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aircraftgroundservicetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aircraftgroundservicetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aircraftgroundservicetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aircraftgroundservicetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", aircraftgroundservicetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", aircraftgroundservicetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", aircraftgroundservicetimeslicetype.CorrectionNumber)

	case "Location":
		if aircraftgroundservicetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", aircraftgroundservicetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "ServiceProvider":
		if aircraftgroundservicetimeslicetype.ServiceProvider != nil {
			res = __gong__marshallPointer(ident, "ServiceProvider", aircraftgroundservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceProvider", "nil")
		}
	case "Call_sign":
		var sb strings.Builder
		for _, _callsigndetailpropertytype := range aircraftgroundservicetimeslicetype.Call_sign {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Call_sign", _callsigndetailpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RadioCommunication":
		var sb strings.Builder
		for _, _radiocommunicationchannelpropertytype := range aircraftgroundservicetimeslicetype.RadioCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadioCommunication", _radiocommunicationchannelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundCommunication":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range aircraftgroundservicetimeslicetype.GroundCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundCommunication", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _serviceoperationalstatuspropertytype := range aircraftgroundservicetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _serviceoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range aircraftgroundservicetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AirportHeliport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range aircraftgroundservicetimeslicetype.AirportHeliport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AirportHeliport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_10 := range aircraftgroundservicetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_10.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AircraftGroundServiceTimeSliceType", fieldName)
	}
	return
}

func (aircraftgroundservicetype *AircraftGroundServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aircraftgroundservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aircraftgroundservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aircraftgroundservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aircraftgroundservicetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _aircraftgroundservicetimeslicepropertytype := range aircraftgroundservicetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _aircraftgroundservicetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AircraftGroundServiceType", fieldName)
	}
	return
}

func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airportclearanceservicetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airportclearanceservicetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airportclearanceservicetimeslicepropertytype.Owns)

	case "AirportClearanceServiceTimeSlice":
		if airportclearanceservicetimeslicepropertytype.AirportClearanceServiceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AirportClearanceServiceTimeSlice", airportclearanceservicetimeslicepropertytype.AirportClearanceServiceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirportClearanceServiceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirportClearanceServiceTimeSlicePropertyType", fieldName)
	}
	return
}

func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airportclearanceservicetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airportclearanceservicetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airportclearanceservicetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airportclearanceservicetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", airportclearanceservicetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", airportclearanceservicetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", airportclearanceservicetimeslicetype.CorrectionNumber)

	case "Location":
		if airportclearanceservicetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", airportclearanceservicetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "ServiceProvider":
		if airportclearanceservicetimeslicetype.ServiceProvider != nil {
			res = __gong__marshallPointer(ident, "ServiceProvider", airportclearanceservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceProvider", "nil")
		}
	case "Call_sign":
		var sb strings.Builder
		for _, _callsigndetailpropertytype := range airportclearanceservicetimeslicetype.Call_sign {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Call_sign", _callsigndetailpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RadioCommunication":
		var sb strings.Builder
		for _, _radiocommunicationchannelpropertytype := range airportclearanceservicetimeslicetype.RadioCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadioCommunication", _radiocommunicationchannelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundCommunication":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range airportclearanceservicetimeslicetype.GroundCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundCommunication", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _serviceoperationalstatuspropertytype := range airportclearanceservicetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _serviceoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airportclearanceservicetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AirportHeliport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range airportclearanceservicetimeslicetype.AirportHeliport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AirportHeliport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_11 := range airportclearanceservicetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_11.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirportClearanceServiceTimeSliceType", fieldName)
	}
	return
}

func (airportclearanceservicetype *AirportClearanceServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airportclearanceservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airportclearanceservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airportclearanceservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airportclearanceservicetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _airportclearanceservicetimeslicepropertytype := range airportclearanceservicetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _airportclearanceservicetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirportClearanceServiceType", fieldName)
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

func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airportsuppliesservicetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airportsuppliesservicetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airportsuppliesservicetimeslicepropertytype.Owns)

	case "AirportSuppliesServiceTimeSlice":
		if airportsuppliesservicetimeslicepropertytype.AirportSuppliesServiceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AirportSuppliesServiceTimeSlice", airportsuppliesservicetimeslicepropertytype.AirportSuppliesServiceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirportSuppliesServiceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirportSuppliesServiceTimeSlicePropertyType", fieldName)
	}
	return
}

func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airportsuppliesservicetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airportsuppliesservicetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airportsuppliesservicetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airportsuppliesservicetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", airportsuppliesservicetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", airportsuppliesservicetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", airportsuppliesservicetimeslicetype.CorrectionNumber)

	case "Location":
		if airportsuppliesservicetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", airportsuppliesservicetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "ServiceProvider":
		if airportsuppliesservicetimeslicetype.ServiceProvider != nil {
			res = __gong__marshallPointer(ident, "ServiceProvider", airportsuppliesservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceProvider", "nil")
		}
	case "Call_sign":
		var sb strings.Builder
		for _, _callsigndetailpropertytype := range airportsuppliesservicetimeslicetype.Call_sign {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Call_sign", _callsigndetailpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RadioCommunication":
		var sb strings.Builder
		for _, _radiocommunicationchannelpropertytype := range airportsuppliesservicetimeslicetype.RadioCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadioCommunication", _radiocommunicationchannelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundCommunication":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range airportsuppliesservicetimeslicetype.GroundCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundCommunication", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _serviceoperationalstatuspropertytype := range airportsuppliesservicetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _serviceoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airportsuppliesservicetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AirportHeliport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range airportsuppliesservicetimeslicetype.AirportHeliport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AirportHeliport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "FuelSupply":
		var sb strings.Builder
		for _, _fuelpropertytype := range airportsuppliesservicetimeslicetype.FuelSupply {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FuelSupply", _fuelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "OilSupply":
		var sb strings.Builder
		for _, _oilpropertytype := range airportsuppliesservicetimeslicetype.OilSupply {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "OilSupply", _oilpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NitrogenSupply":
		var sb strings.Builder
		for _, _nitrogenpropertytype := range airportsuppliesservicetimeslicetype.NitrogenSupply {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NitrogenSupply", _nitrogenpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "OxygenSupply":
		var sb strings.Builder
		for _, _oxygenpropertytype := range airportsuppliesservicetimeslicetype.OxygenSupply {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "OxygenSupply", _oxygenpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_12 := range airportsuppliesservicetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_12.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirportSuppliesServiceTimeSliceType", fieldName)
	}
	return
}

func (airportsuppliesservicetype *AirportSuppliesServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airportsuppliesservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airportsuppliesservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airportsuppliesservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airportsuppliesservicetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _airportsuppliesservicetimeslicepropertytype := range airportsuppliesservicetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _airportsuppliesservicetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirportSuppliesServiceType", fieldName)
	}
	return
}

func (airspacepropertytype *AirspacePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airspacepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airspacepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", airspacepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", airspacepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", airspacepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", airspacepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", airspacepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", airspacepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", airspacepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct AirspacePropertyType", fieldName)
	}
	return
}

func (callsigndetailpropertytype *CallsignDetailPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := callsigndetailpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", callsigndetailpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", callsigndetailpropertytype.NilReason)

	case "CallsignDetail":
		if callsigndetailpropertytype.CallsignDetail != nil {
			res = __gong__marshallPointer(ident, "CallsignDetail", callsigndetailpropertytype.CallsignDetail.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CallsignDetail", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CallsignDetailPropertyType", fieldName)
	}
	return
}

func (callsigndetailtype *CallsignDetailType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := callsigndetailtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", callsigndetailtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", callsigndetailtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", callsigndetailtype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range callsigndetailtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_15 := range callsigndetailtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_15.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CallsignDetailType", fieldName)
	}
	return
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := contactinformationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", contactinformationpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", contactinformationpropertytype.NilReason)

	case "ContactInformation":
		if contactinformationpropertytype.ContactInformation != nil {
			res = __gong__marshallPointer(ident, "ContactInformation", contactinformationpropertytype.ContactInformation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ContactInformation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ContactInformationPropertyType", fieldName)
	}
	return
}

func (contactinformationtype *ContactInformationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := contactinformationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", contactinformationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", contactinformationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", contactinformationtype.Description)

	case "Address":
		var sb strings.Builder
		for _, _postaladdresspropertytype := range contactinformationtype.Address {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Address", _postaladdresspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NetworkNode":
		var sb strings.Builder
		for _, _onlinecontactpropertytype := range contactinformationtype.NetworkNode {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NetworkNode", _onlinecontactpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PhoneFax":
		var sb strings.Builder
		for _, _telephonecontactpropertytype := range contactinformationtype.PhoneFax {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PhoneFax", _telephonecontactpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range contactinformationtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_29 := range contactinformationtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_29.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ContactInformationType", fieldName)
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

func (directionfinderpropertytype *DirectionFinderPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directionfinderpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directionfinderpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", directionfinderpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", directionfinderpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", directionfinderpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", directionfinderpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", directionfinderpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", directionfinderpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", directionfinderpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", directionfinderpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", directionfinderpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct DirectionFinderPropertyType", fieldName)
	}
	return
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevatedpointpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevatedpointpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", elevatedpointpropertytype.NilReason)

	case "ElevatedPoint":
		if elevatedpointpropertytype.ElevatedPoint != nil {
			res = __gong__marshallPointer(ident, "ElevatedPoint", elevatedpointpropertytype.ElevatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ElevatedPoint", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedPointPropertyType", fieldName)
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
		for _, _a_extension_1 := range elevatedpointtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_1.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedPointType", fieldName)
	}
	return
}

func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := firefightingservicetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", firefightingservicetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", firefightingservicetimeslicepropertytype.Owns)

	case "FireFightingServiceTimeSlice":
		if firefightingservicetimeslicepropertytype.FireFightingServiceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "FireFightingServiceTimeSlice", firefightingservicetimeslicepropertytype.FireFightingServiceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FireFightingServiceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FireFightingServiceTimeSlicePropertyType", fieldName)
	}
	return
}

func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := firefightingservicetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", firefightingservicetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", firefightingservicetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", firefightingservicetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", firefightingservicetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", firefightingservicetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", firefightingservicetimeslicetype.CorrectionNumber)

	case "Location":
		if firefightingservicetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", firefightingservicetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "ServiceProvider":
		if firefightingservicetimeslicetype.ServiceProvider != nil {
			res = __gong__marshallPointer(ident, "ServiceProvider", firefightingservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceProvider", "nil")
		}
	case "Call_sign":
		var sb strings.Builder
		for _, _callsigndetailpropertytype := range firefightingservicetimeslicetype.Call_sign {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Call_sign", _callsigndetailpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RadioCommunication":
		var sb strings.Builder
		for _, _radiocommunicationchannelpropertytype := range firefightingservicetimeslicetype.RadioCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadioCommunication", _radiocommunicationchannelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundCommunication":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range firefightingservicetimeslicetype.GroundCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundCommunication", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _serviceoperationalstatuspropertytype := range firefightingservicetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _serviceoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range firefightingservicetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AirportHeliport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range firefightingservicetimeslicetype.AirportHeliport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AirportHeliport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_16 := range firefightingservicetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_16.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FireFightingServiceTimeSliceType", fieldName)
	}
	return
}

func (firefightingservicetype *FireFightingServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := firefightingservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", firefightingservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", firefightingservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", firefightingservicetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _firefightingservicetimeslicepropertytype := range firefightingservicetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _firefightingservicetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FireFightingServiceType", fieldName)
	}
	return
}

func (fuelpropertytype *FuelPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := fuelpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", fuelpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", fuelpropertytype.NilReason)

	case "Fuel":
		if fuelpropertytype.Fuel != nil {
			res = __gong__marshallPointer(ident, "Fuel", fuelpropertytype.Fuel.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Fuel", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FuelPropertyType", fieldName)
	}
	return
}

func (fueltype *FuelType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := fueltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", fueltype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", fueltype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", fueltype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range fueltype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_17 := range fueltype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_17.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FuelType", fieldName)
	}
	return
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := groundlightsystempropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", groundlightsystempropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", groundlightsystempropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", groundlightsystempropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", groundlightsystempropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", groundlightsystempropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", groundlightsystempropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", groundlightsystempropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", groundlightsystempropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", groundlightsystempropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", groundlightsystempropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct GroundLightSystemPropertyType", fieldName)
	}
	return
}

func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := groundtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", groundtrafficcontrolservicetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", groundtrafficcontrolservicetimeslicepropertytype.Owns)

	case "GroundTrafficControlServiceTimeSlice":
		if groundtrafficcontrolservicetimeslicepropertytype.GroundTrafficControlServiceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "GroundTrafficControlServiceTimeSlice", groundtrafficcontrolservicetimeslicepropertytype.GroundTrafficControlServiceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GroundTrafficControlServiceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GroundTrafficControlServiceTimeSlicePropertyType", fieldName)
	}
	return
}

func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := groundtrafficcontrolservicetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", groundtrafficcontrolservicetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", groundtrafficcontrolservicetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", groundtrafficcontrolservicetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", groundtrafficcontrolservicetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", groundtrafficcontrolservicetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", groundtrafficcontrolservicetimeslicetype.CorrectionNumber)

	case "Location":
		if groundtrafficcontrolservicetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", groundtrafficcontrolservicetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "ServiceProvider":
		if groundtrafficcontrolservicetimeslicetype.ServiceProvider != nil {
			res = __gong__marshallPointer(ident, "ServiceProvider", groundtrafficcontrolservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceProvider", "nil")
		}
	case "Call_sign":
		var sb strings.Builder
		for _, _callsigndetailpropertytype := range groundtrafficcontrolservicetimeslicetype.Call_sign {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Call_sign", _callsigndetailpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RadioCommunication":
		var sb strings.Builder
		for _, _radiocommunicationchannelpropertytype := range groundtrafficcontrolservicetimeslicetype.RadioCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadioCommunication", _radiocommunicationchannelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundCommunication":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range groundtrafficcontrolservicetimeslicetype.GroundCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundCommunication", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _serviceoperationalstatuspropertytype := range groundtrafficcontrolservicetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _serviceoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range groundtrafficcontrolservicetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientAirport":
		if groundtrafficcontrolservicetimeslicetype.ClientAirport != nil {
			res = __gong__marshallPointer(ident, "ClientAirport", groundtrafficcontrolservicetimeslicetype.ClientAirport.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ClientAirport", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_18 := range groundtrafficcontrolservicetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_18.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GroundTrafficControlServiceTimeSliceType", fieldName)
	}
	return
}

func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := groundtrafficcontrolservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", groundtrafficcontrolservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", groundtrafficcontrolservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", groundtrafficcontrolservicetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _groundtrafficcontrolservicetimeslicepropertytype := range groundtrafficcontrolservicetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _groundtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GroundTrafficControlServiceType", fieldName)
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

func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := informationservicetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", informationservicetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", informationservicetimeslicepropertytype.Owns)

	case "InformationServiceTimeSlice":
		if informationservicetimeslicepropertytype.InformationServiceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "InformationServiceTimeSlice", informationservicetimeslicepropertytype.InformationServiceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "InformationServiceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct InformationServiceTimeSlicePropertyType", fieldName)
	}
	return
}

func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := informationservicetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", informationservicetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", informationservicetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", informationservicetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", informationservicetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", informationservicetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", informationservicetimeslicetype.CorrectionNumber)

	case "Location":
		if informationservicetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", informationservicetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "ServiceProvider":
		if informationservicetimeslicetype.ServiceProvider != nil {
			res = __gong__marshallPointer(ident, "ServiceProvider", informationservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceProvider", "nil")
		}
	case "Call_sign":
		var sb strings.Builder
		for _, _callsigndetailpropertytype := range informationservicetimeslicetype.Call_sign {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Call_sign", _callsigndetailpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RadioCommunication":
		var sb strings.Builder
		for _, _radiocommunicationchannelpropertytype := range informationservicetimeslicetype.RadioCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadioCommunication", _radiocommunicationchannelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundCommunication":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range informationservicetimeslicetype.GroundCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundCommunication", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _serviceoperationalstatuspropertytype := range informationservicetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _serviceoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range informationservicetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NavaidBroadcast":
		var sb strings.Builder
		for _, _vorpropertytype := range informationservicetimeslicetype.NavaidBroadcast {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NavaidBroadcast", _vorpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientAirspace":
		var sb strings.Builder
		for _, _airspacepropertytype := range informationservicetimeslicetype.ClientAirspace {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientAirspace", _airspacepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientAirport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range informationservicetimeslicetype.ClientAirport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientAirport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientRoute":
		var sb strings.Builder
		for _, _routeportionpropertytype := range informationservicetimeslicetype.ClientRoute {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientRoute", _routeportionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientProcedure":
		var sb strings.Builder
		for _, _procedurepropertytype := range informationservicetimeslicetype.ClientProcedure {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientProcedure", _procedurepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientHolding":
		var sb strings.Builder
		for _, _holdingpatternpropertytype := range informationservicetimeslicetype.ClientHolding {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientHolding", _holdingpatternpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientAerialRefuelling":
		var sb strings.Builder
		for _, _aerialrefuellingpropertytype := range informationservicetimeslicetype.ClientAerialRefuelling {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientAerialRefuelling", _aerialrefuellingpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_19 := range informationservicetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_19.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct InformationServiceTimeSliceType", fieldName)
	}
	return
}

func (informationservicetype *InformationServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := informationservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", informationservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", informationservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", informationservicetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _informationservicetimeslicepropertytype := range informationservicetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _informationservicetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct InformationServiceType", fieldName)
	}
	return
}

func (lightactivationpropertytype *LightActivationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := lightactivationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", lightactivationpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", lightactivationpropertytype.NilReason)

	case "LightActivation":
		if lightactivationpropertytype.LightActivation != nil {
			res = __gong__marshallPointer(ident, "LightActivation", lightactivationpropertytype.LightActivation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LightActivation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LightActivationPropertyType", fieldName)
	}
	return
}

func (lightactivationtype *LightActivationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := lightactivationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", lightactivationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", lightactivationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", lightactivationtype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range lightactivationtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension := range lightactivationtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LightActivationType", fieldName)
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
		for _, _a_extension_2 := range linguisticnotetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_2.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LinguisticNoteType", fieldName)
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

func (nitrogenpropertytype *NitrogenPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := nitrogenpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", nitrogenpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", nitrogenpropertytype.NilReason)

	case "Nitrogen":
		if nitrogenpropertytype.Nitrogen != nil {
			res = __gong__marshallPointer(ident, "Nitrogen", nitrogenpropertytype.Nitrogen.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Nitrogen", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NitrogenPropertyType", fieldName)
	}
	return
}

func (nitrogentype *NitrogenType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := nitrogentype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", nitrogentype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", nitrogentype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", nitrogentype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range nitrogentype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_20 := range nitrogentype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_20.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NitrogenType", fieldName)
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
		for _, _a_extension_3 := range notetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_3.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteType", fieldName)
	}
	return
}

func (oilpropertytype *OilPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := oilpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", oilpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", oilpropertytype.NilReason)

	case "Oil":
		if oilpropertytype.Oil != nil {
			res = __gong__marshallPointer(ident, "Oil", oilpropertytype.Oil.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Oil", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct OilPropertyType", fieldName)
	}
	return
}

func (oiltype *OilType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := oiltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", oiltype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", oiltype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", oiltype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range oiltype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_21 := range oiltype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_21.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct OilType", fieldName)
	}
	return
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := onlinecontactpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", onlinecontactpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", onlinecontactpropertytype.NilReason)

	case "OnlineContact":
		if onlinecontactpropertytype.OnlineContact != nil {
			res = __gong__marshallPointer(ident, "OnlineContact", onlinecontactpropertytype.OnlineContact.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OnlineContact", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct OnlineContactPropertyType", fieldName)
	}
	return
}

func (onlinecontacttype *OnlineContactType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := onlinecontacttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", onlinecontacttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", onlinecontacttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", onlinecontacttype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range onlinecontacttype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range onlinecontacttype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range onlinecontacttype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_30 := range onlinecontacttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_30.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct OnlineContactType", fieldName)
	}
	return
}

func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := organisationauthorityassociationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", organisationauthorityassociationpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", organisationauthorityassociationpropertytype.NilReason)

	case "OrganisationAuthorityAssociation":
		if organisationauthorityassociationpropertytype.OrganisationAuthorityAssociation != nil {
			res = __gong__marshallPointer(ident, "OrganisationAuthorityAssociation", organisationauthorityassociationpropertytype.OrganisationAuthorityAssociation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OrganisationAuthorityAssociation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct OrganisationAuthorityAssociationPropertyType", fieldName)
	}
	return
}

func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := organisationauthorityassociationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", organisationauthorityassociationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", organisationauthorityassociationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", organisationauthorityassociationtype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range organisationauthorityassociationtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheOrganisationAuthority":
		if organisationauthorityassociationtype.TheOrganisationAuthority != nil {
			res = __gong__marshallPointer(ident, "TheOrganisationAuthority", organisationauthorityassociationtype.TheOrganisationAuthority.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheOrganisationAuthority", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_5 := range organisationauthorityassociationtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_5.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct OrganisationAuthorityAssociationType", fieldName)
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

func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := organisationauthoritytimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", organisationauthoritytimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", organisationauthoritytimeslicepropertytype.Owns)

	case "OrganisationAuthorityTimeSlice":
		if organisationauthoritytimeslicepropertytype.OrganisationAuthorityTimeSlice != nil {
			res = __gong__marshallPointer(ident, "OrganisationAuthorityTimeSlice", organisationauthoritytimeslicepropertytype.OrganisationAuthorityTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OrganisationAuthorityTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct OrganisationAuthorityTimeSlicePropertyType", fieldName)
	}
	return
}

func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := organisationauthoritytimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", organisationauthoritytimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", organisationauthoritytimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", organisationauthoritytimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", organisationauthoritytimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", organisationauthoritytimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", organisationauthoritytimeslicetype.CorrectionNumber)

	case "Contact":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range organisationauthoritytimeslicetype.Contact {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Contact", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RelatedOrganisationAuthority":
		var sb strings.Builder
		for _, _organisationauthorityassociationpropertytype := range organisationauthoritytimeslicetype.RelatedOrganisationAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RelatedOrganisationAuthority", _organisationauthorityassociationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range organisationauthoritytimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_4 := range organisationauthoritytimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_4.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct OrganisationAuthorityTimeSliceType", fieldName)
	}
	return
}

func (organisationauthoritytype *OrganisationAuthorityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := organisationauthoritytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", organisationauthoritytype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", organisationauthoritytype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", organisationauthoritytype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _organisationauthoritytimeslicepropertytype := range organisationauthoritytype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _organisationauthoritytimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct OrganisationAuthorityType", fieldName)
	}
	return
}

func (oxygenpropertytype *OxygenPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := oxygenpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", oxygenpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", oxygenpropertytype.NilReason)

	case "Oxygen":
		if oxygenpropertytype.Oxygen != nil {
			res = __gong__marshallPointer(ident, "Oxygen", oxygenpropertytype.Oxygen.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Oxygen", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct OxygenPropertyType", fieldName)
	}
	return
}

func (oxygentype *OxygenType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := oxygentype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", oxygentype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", oxygentype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", oxygentype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range oxygentype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_22 := range oxygentype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_22.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct OxygenType", fieldName)
	}
	return
}

func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := passengerservicetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", passengerservicetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", passengerservicetimeslicepropertytype.Owns)

	case "PassengerServiceTimeSlice":
		if passengerservicetimeslicepropertytype.PassengerServiceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "PassengerServiceTimeSlice", passengerservicetimeslicepropertytype.PassengerServiceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PassengerServiceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PassengerServiceTimeSlicePropertyType", fieldName)
	}
	return
}

func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := passengerservicetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", passengerservicetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", passengerservicetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", passengerservicetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", passengerservicetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", passengerservicetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", passengerservicetimeslicetype.CorrectionNumber)

	case "Location":
		if passengerservicetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", passengerservicetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "ServiceProvider":
		if passengerservicetimeslicetype.ServiceProvider != nil {
			res = __gong__marshallPointer(ident, "ServiceProvider", passengerservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceProvider", "nil")
		}
	case "Call_sign":
		var sb strings.Builder
		for _, _callsigndetailpropertytype := range passengerservicetimeslicetype.Call_sign {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Call_sign", _callsigndetailpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RadioCommunication":
		var sb strings.Builder
		for _, _radiocommunicationchannelpropertytype := range passengerservicetimeslicetype.RadioCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadioCommunication", _radiocommunicationchannelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundCommunication":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range passengerservicetimeslicetype.GroundCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundCommunication", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _serviceoperationalstatuspropertytype := range passengerservicetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _serviceoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range passengerservicetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AirportHeliport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range passengerservicetimeslicetype.AirportHeliport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AirportHeliport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_23 := range passengerservicetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_23.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PassengerServiceTimeSliceType", fieldName)
	}
	return
}

func (passengerservicetype *PassengerServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := passengerservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", passengerservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", passengerservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", passengerservicetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _passengerservicetimeslicepropertytype := range passengerservicetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _passengerservicetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PassengerServiceType", fieldName)
	}
	return
}

func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pilotcontrolledlightingtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pilotcontrolledlightingtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", pilotcontrolledlightingtimeslicepropertytype.Owns)

	case "PilotControlledLightingTimeSlice":
		if pilotcontrolledlightingtimeslicepropertytype.PilotControlledLightingTimeSlice != nil {
			res = __gong__marshallPointer(ident, "PilotControlledLightingTimeSlice", pilotcontrolledlightingtimeslicepropertytype.PilotControlledLightingTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PilotControlledLightingTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PilotControlledLightingTimeSlicePropertyType", fieldName)
	}
	return
}

func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pilotcontrolledlightingtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pilotcontrolledlightingtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", pilotcontrolledlightingtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", pilotcontrolledlightingtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", pilotcontrolledlightingtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", pilotcontrolledlightingtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", pilotcontrolledlightingtimeslicetype.CorrectionNumber)

	case "ControlledLightIntensity":
		var sb strings.Builder
		for _, _lightactivationpropertytype := range pilotcontrolledlightingtimeslicetype.ControlledLightIntensity {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlledLightIntensity", _lightactivationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ActivatedGroundLighting":
		var sb strings.Builder
		for _, _groundlightsystempropertytype := range pilotcontrolledlightingtimeslicetype.ActivatedGroundLighting {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ActivatedGroundLighting", _groundlightsystempropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range pilotcontrolledlightingtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_24 := range pilotcontrolledlightingtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_24.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PilotControlledLightingTimeSliceType", fieldName)
	}
	return
}

func (pilotcontrolledlightingtype *PilotControlledLightingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pilotcontrolledlightingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pilotcontrolledlightingtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", pilotcontrolledlightingtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", pilotcontrolledlightingtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _pilotcontrolledlightingtimeslicepropertytype := range pilotcontrolledlightingtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _pilotcontrolledlightingtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PilotControlledLightingType", fieldName)
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

func (postaladdresspropertytype *PostalAddressPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := postaladdresspropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", postaladdresspropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", postaladdresspropertytype.NilReason)

	case "PostalAddress":
		if postaladdresspropertytype.PostalAddress != nil {
			res = __gong__marshallPointer(ident, "PostalAddress", postaladdresspropertytype.PostalAddress.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PostalAddress", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PostalAddressPropertyType", fieldName)
	}
	return
}

func (postaladdresstype *PostalAddressType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := postaladdresstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", postaladdresstype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", postaladdresstype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", postaladdresstype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range postaladdresstype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range postaladdresstype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range postaladdresstype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_31 := range postaladdresstype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_31.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PostalAddressType", fieldName)
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

func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radiocommunicationchannelpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radiocommunicationchannelpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", radiocommunicationchannelpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", radiocommunicationchannelpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", radiocommunicationchannelpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", radiocommunicationchannelpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", radiocommunicationchannelpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", radiocommunicationchannelpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", radiocommunicationchannelpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", radiocommunicationchannelpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", radiocommunicationchannelpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct RadioCommunicationChannelPropertyType", fieldName)
	}
	return
}

func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radiocommunicationchanneltimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radiocommunicationchanneltimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", radiocommunicationchanneltimeslicepropertytype.Owns)

	case "RadioCommunicationChannelTimeSlice":
		if radiocommunicationchanneltimeslicepropertytype.RadioCommunicationChannelTimeSlice != nil {
			res = __gong__marshallPointer(ident, "RadioCommunicationChannelTimeSlice", radiocommunicationchanneltimeslicepropertytype.RadioCommunicationChannelTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RadioCommunicationChannelTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RadioCommunicationChannelTimeSlicePropertyType", fieldName)
	}
	return
}

func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radiocommunicationchanneltimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radiocommunicationchanneltimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", radiocommunicationchanneltimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", radiocommunicationchanneltimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", radiocommunicationchanneltimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", radiocommunicationchanneltimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", radiocommunicationchanneltimeslicetype.CorrectionNumber)

	case "Location":
		var sb strings.Builder
		for _, _elevatedpointpropertytype := range radiocommunicationchanneltimeslicetype.Location {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Location", _elevatedpointpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _radiocommunicationoperationalstatuspropertytype := range radiocommunicationchanneltimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _radiocommunicationoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range radiocommunicationchanneltimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_25 := range radiocommunicationchanneltimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_25.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RadioCommunicationChannelTimeSliceType", fieldName)
	}
	return
}

func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radiocommunicationchanneltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radiocommunicationchanneltype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", radiocommunicationchanneltype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", radiocommunicationchanneltype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _radiocommunicationchanneltimeslicepropertytype := range radiocommunicationchanneltype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _radiocommunicationchanneltimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RadioCommunicationChannelType", fieldName)
	}
	return
}

func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radiocommunicationoperationalstatuspropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radiocommunicationoperationalstatuspropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", radiocommunicationoperationalstatuspropertytype.NilReason)

	case "RadioCommunicationOperationalStatus":
		if radiocommunicationoperationalstatuspropertytype.RadioCommunicationOperationalStatus != nil {
			res = __gong__marshallPointer(ident, "RadioCommunicationOperationalStatus", radiocommunicationoperationalstatuspropertytype.RadioCommunicationOperationalStatus.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RadioCommunicationOperationalStatus", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RadioCommunicationOperationalStatusPropertyType", fieldName)
	}
	return
}

func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radiocommunicationoperationalstatustype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radiocommunicationoperationalstatustype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", radiocommunicationoperationalstatustype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", radiocommunicationoperationalstatustype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range radiocommunicationoperationalstatustype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range radiocommunicationoperationalstatustype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range radiocommunicationoperationalstatustype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_26 := range radiocommunicationoperationalstatustype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_26.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RadioCommunicationOperationalStatusType", fieldName)
	}
	return
}

func (routeportionpropertytype *RoutePortionPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routeportionpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routeportionpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", routeportionpropertytype.NilReason)

	case "RoutePortion":
		if routeportionpropertytype.RoutePortion != nil {
			res = __gong__marshallPointer(ident, "RoutePortion", routeportionpropertytype.RoutePortion.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RoutePortion", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RoutePortionPropertyType", fieldName)
	}
	return
}

func (routeportiontype *RoutePortionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routeportiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routeportiontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", routeportiontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", routeportiontype.Description)

	case "Start_fixDesignatedPoint":
		if routeportiontype.Start_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "Start_fixDesignatedPoint", routeportiontype.Start_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_fixDesignatedPoint", "nil")
		}
	case "Start_navaidSystem":
		if routeportiontype.Start_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "Start_navaidSystem", routeportiontype.Start_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_navaidSystem", "nil")
		}
	case "Start_position":
		if routeportiontype.Start_position != nil {
			res = __gong__marshallPointer(ident, "Start_position", routeportiontype.Start_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_position", "nil")
		}
	case "Start_runwayPoint":
		if routeportiontype.Start_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "Start_runwayPoint", routeportiontype.Start_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_runwayPoint", "nil")
		}
	case "Start_aimingPoint":
		if routeportiontype.Start_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "Start_aimingPoint", routeportiontype.Start_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_aimingPoint", "nil")
		}
	case "Start_airportReferencePoint":
		if routeportiontype.Start_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "Start_airportReferencePoint", routeportiontype.Start_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_airportReferencePoint", "nil")
		}
	case "IntermediatePoint_fixDesignatedPoint":
		if routeportiontype.IntermediatePoint_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "IntermediatePoint_fixDesignatedPoint", routeportiontype.IntermediatePoint_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IntermediatePoint_fixDesignatedPoint", "nil")
		}
	case "IntermediatePoint_navaidSystem":
		if routeportiontype.IntermediatePoint_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "IntermediatePoint_navaidSystem", routeportiontype.IntermediatePoint_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IntermediatePoint_navaidSystem", "nil")
		}
	case "IntermediatePoint_position":
		if routeportiontype.IntermediatePoint_position != nil {
			res = __gong__marshallPointer(ident, "IntermediatePoint_position", routeportiontype.IntermediatePoint_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IntermediatePoint_position", "nil")
		}
	case "IntermediatePoint_runwayPoint":
		if routeportiontype.IntermediatePoint_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "IntermediatePoint_runwayPoint", routeportiontype.IntermediatePoint_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IntermediatePoint_runwayPoint", "nil")
		}
	case "IntermediatePoint_aimingPoint":
		if routeportiontype.IntermediatePoint_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "IntermediatePoint_aimingPoint", routeportiontype.IntermediatePoint_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IntermediatePoint_aimingPoint", "nil")
		}
	case "IntermediatePoint_airportReferencePoint":
		if routeportiontype.IntermediatePoint_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "IntermediatePoint_airportReferencePoint", routeportiontype.IntermediatePoint_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "IntermediatePoint_airportReferencePoint", "nil")
		}
	case "ReferencedRoute":
		if routeportiontype.ReferencedRoute != nil {
			res = __gong__marshallPointer(ident, "ReferencedRoute", routeportiontype.ReferencedRoute.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ReferencedRoute", "nil")
		}
	case "End_fixDesignatedPoint":
		if routeportiontype.End_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "End_fixDesignatedPoint", routeportiontype.End_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_fixDesignatedPoint", "nil")
		}
	case "End_navaidSystem":
		if routeportiontype.End_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "End_navaidSystem", routeportiontype.End_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_navaidSystem", "nil")
		}
	case "End_position":
		if routeportiontype.End_position != nil {
			res = __gong__marshallPointer(ident, "End_position", routeportiontype.End_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_position", "nil")
		}
	case "End_runwayPoint":
		if routeportiontype.End_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "End_runwayPoint", routeportiontype.End_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_runwayPoint", "nil")
		}
	case "End_aimingPoint":
		if routeportiontype.End_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "End_aimingPoint", routeportiontype.End_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_aimingPoint", "nil")
		}
	case "End_airportReferencePoint":
		if routeportiontype.End_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "End_airportReferencePoint", routeportiontype.End_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_airportReferencePoint", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range routeportiontype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_9 := range routeportiontype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_9.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RoutePortionType", fieldName)
	}
	return
}

func (routepropertytype *RoutePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", routepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", routepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", routepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", routepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", routepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", routepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", routepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", routepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", routepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct RoutePropertyType", fieldName)
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

func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := searchrescueservicetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", searchrescueservicetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", searchrescueservicetimeslicepropertytype.Owns)

	case "SearchRescueServiceTimeSlice":
		if searchrescueservicetimeslicepropertytype.SearchRescueServiceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "SearchRescueServiceTimeSlice", searchrescueservicetimeslicepropertytype.SearchRescueServiceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SearchRescueServiceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SearchRescueServiceTimeSlicePropertyType", fieldName)
	}
	return
}

func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := searchrescueservicetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", searchrescueservicetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", searchrescueservicetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", searchrescueservicetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", searchrescueservicetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", searchrescueservicetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", searchrescueservicetimeslicetype.CorrectionNumber)

	case "Location":
		if searchrescueservicetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", searchrescueservicetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "ServiceProvider":
		if searchrescueservicetimeslicetype.ServiceProvider != nil {
			res = __gong__marshallPointer(ident, "ServiceProvider", searchrescueservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceProvider", "nil")
		}
	case "Call_sign":
		var sb strings.Builder
		for _, _callsigndetailpropertytype := range searchrescueservicetimeslicetype.Call_sign {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Call_sign", _callsigndetailpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RadioCommunication":
		var sb strings.Builder
		for _, _radiocommunicationchannelpropertytype := range searchrescueservicetimeslicetype.RadioCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadioCommunication", _radiocommunicationchannelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundCommunication":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range searchrescueservicetimeslicetype.GroundCommunication {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundCommunication", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _serviceoperationalstatuspropertytype := range searchrescueservicetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _serviceoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range searchrescueservicetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientAirspace":
		var sb strings.Builder
		for _, _airspacepropertytype := range searchrescueservicetimeslicetype.ClientAirspace {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientAirspace", _airspacepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ClientRoute":
		var sb strings.Builder
		for _, _routeportionpropertytype := range searchrescueservicetimeslicetype.ClientRoute {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ClientRoute", _routeportionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_27 := range searchrescueservicetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_27.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SearchRescueServiceTimeSliceType", fieldName)
	}
	return
}

func (searchrescueservicetype *SearchRescueServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := searchrescueservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", searchrescueservicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", searchrescueservicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", searchrescueservicetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _searchrescueservicetimeslicepropertytype := range searchrescueservicetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _searchrescueservicetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SearchRescueServiceType", fieldName)
	}
	return
}

func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := serviceoperationalstatuspropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", serviceoperationalstatuspropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", serviceoperationalstatuspropertytype.NilReason)

	case "ServiceOperationalStatus":
		if serviceoperationalstatuspropertytype.ServiceOperationalStatus != nil {
			res = __gong__marshallPointer(ident, "ServiceOperationalStatus", serviceoperationalstatuspropertytype.ServiceOperationalStatus.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ServiceOperationalStatus", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ServiceOperationalStatusPropertyType", fieldName)
	}
	return
}

func (serviceoperationalstatustype *ServiceOperationalStatusType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := serviceoperationalstatustype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", serviceoperationalstatustype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", serviceoperationalstatustype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", serviceoperationalstatustype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range serviceoperationalstatustype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range serviceoperationalstatustype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range serviceoperationalstatustype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_28 := range serviceoperationalstatustype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_28.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ServiceOperationalStatusType", fieldName)
	}
	return
}

func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialdatetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialdatetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", specialdatetimeslicepropertytype.Owns)

	case "SpecialDateTimeSlice":
		if specialdatetimeslicepropertytype.SpecialDateTimeSlice != nil {
			res = __gong__marshallPointer(ident, "SpecialDateTimeSlice", specialdatetimeslicepropertytype.SpecialDateTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SpecialDateTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialDateTimeSlicePropertyType", fieldName)
	}
	return
}

func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialdatetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialdatetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", specialdatetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", specialdatetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", specialdatetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", specialdatetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", specialdatetimeslicetype.CorrectionNumber)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range specialdatetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Authority":
		if specialdatetimeslicetype.Authority != nil {
			res = __gong__marshallPointer(ident, "Authority", specialdatetimeslicetype.Authority.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Authority", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_33 := range specialdatetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_33.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialDateTimeSliceType", fieldName)
	}
	return
}

func (specialdatetype *SpecialDateType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialdatetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialdatetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", specialdatetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", specialdatetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _specialdatetimeslicepropertytype := range specialdatetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _specialdatetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialDateType", fieldName)
	}
	return
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardlevelcolumnpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardlevelcolumnpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", standardlevelcolumnpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", standardlevelcolumnpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", standardlevelcolumnpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", standardlevelcolumnpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", standardlevelcolumnpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", standardlevelcolumnpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", standardlevelcolumnpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", standardlevelcolumnpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", standardlevelcolumnpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelColumnPropertyType", fieldName)
	}
	return
}

func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardlevelcolumntimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardlevelcolumntimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", standardlevelcolumntimeslicepropertytype.Owns)

	case "StandardLevelColumnTimeSlice":
		if standardlevelcolumntimeslicepropertytype.StandardLevelColumnTimeSlice != nil {
			res = __gong__marshallPointer(ident, "StandardLevelColumnTimeSlice", standardlevelcolumntimeslicepropertytype.StandardLevelColumnTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StandardLevelColumnTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelColumnTimeSlicePropertyType", fieldName)
	}
	return
}

func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardlevelcolumntimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardlevelcolumntimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardlevelcolumntimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardlevelcolumntimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", standardlevelcolumntimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", standardlevelcolumntimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", standardlevelcolumntimeslicetype.CorrectionNumber)
	case "UnitOfMeasurement":
		res = __gong__marshallString(ident, "UnitOfMeasurement", standardlevelcolumntimeslicetype.UnitOfMeasurement)

	case "Level":
		var sb strings.Builder
		for _, _standardlevelpropertytype := range standardlevelcolumntimeslicetype.Level {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Level", _standardlevelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "LevelTable":
		if standardlevelcolumntimeslicetype.LevelTable != nil {
			res = __gong__marshallPointer(ident, "LevelTable", standardlevelcolumntimeslicetype.LevelTable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LevelTable", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range standardlevelcolumntimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_36 := range standardlevelcolumntimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_36.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelColumnTimeSliceType", fieldName)
	}
	return
}

func (standardlevelcolumntype *StandardLevelColumnType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardlevelcolumntype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardlevelcolumntype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardlevelcolumntype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardlevelcolumntype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _standardlevelcolumntimeslicepropertytype := range standardlevelcolumntype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _standardlevelcolumntimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelColumnType", fieldName)
	}
	return
}

func (standardlevelpropertytype *StandardLevelPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardlevelpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardlevelpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", standardlevelpropertytype.NilReason)

	case "StandardLevel":
		if standardlevelpropertytype.StandardLevel != nil {
			res = __gong__marshallPointer(ident, "StandardLevel", standardlevelpropertytype.StandardLevel.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StandardLevel", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelPropertyType", fieldName)
	}
	return
}

func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardlevelsectortimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardlevelsectortimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", standardlevelsectortimeslicepropertytype.Owns)

	case "StandardLevelSectorTimeSlice":
		if standardlevelsectortimeslicepropertytype.StandardLevelSectorTimeSlice != nil {
			res = __gong__marshallPointer(ident, "StandardLevelSectorTimeSlice", standardlevelsectortimeslicepropertytype.StandardLevelSectorTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StandardLevelSectorTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelSectorTimeSlicePropertyType", fieldName)
	}
	return
}

func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardlevelsectortimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardlevelsectortimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardlevelsectortimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardlevelsectortimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", standardlevelsectortimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", standardlevelsectortimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", standardlevelsectortimeslicetype.CorrectionNumber)

	case "ApplicableAirspace":
		var sb strings.Builder
		for _, _airspacepropertytype := range standardlevelsectortimeslicetype.ApplicableAirspace {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ApplicableAirspace", _airspacepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ApplicableLevelColumn":
		if standardlevelsectortimeslicetype.ApplicableLevelColumn != nil {
			res = __gong__marshallPointer(ident, "ApplicableLevelColumn", standardlevelsectortimeslicetype.ApplicableLevelColumn.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ApplicableLevelColumn", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range standardlevelsectortimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_37 := range standardlevelsectortimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_37.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelSectorTimeSliceType", fieldName)
	}
	return
}

func (standardlevelsectortype *StandardLevelSectorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardlevelsectortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardlevelsectortype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardlevelsectortype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardlevelsectortype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _standardlevelsectortimeslicepropertytype := range standardlevelsectortype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _standardlevelsectortimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelSectorType", fieldName)
	}
	return
}

func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardleveltablepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardleveltablepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", standardleveltablepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", standardleveltablepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", standardleveltablepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", standardleveltablepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", standardleveltablepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", standardleveltablepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", standardleveltablepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", standardleveltablepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", standardleveltablepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelTablePropertyType", fieldName)
	}
	return
}

func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardleveltabletimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardleveltabletimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", standardleveltabletimeslicepropertytype.Owns)

	case "StandardLevelTableTimeSlice":
		if standardleveltabletimeslicepropertytype.StandardLevelTableTimeSlice != nil {
			res = __gong__marshallPointer(ident, "StandardLevelTableTimeSlice", standardleveltabletimeslicepropertytype.StandardLevelTableTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StandardLevelTableTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelTableTimeSlicePropertyType", fieldName)
	}
	return
}

func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardleveltabletimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardleveltabletimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardleveltabletimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardleveltabletimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", standardleveltabletimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", standardleveltabletimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", standardleveltabletimeslicetype.CorrectionNumber)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range standardleveltabletimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_38 := range standardleveltabletimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_38.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelTableTimeSliceType", fieldName)
	}
	return
}

func (standardleveltabletype *StandardLevelTableType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardleveltabletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardleveltabletype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardleveltabletype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardleveltabletype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _standardleveltabletimeslicepropertytype := range standardleveltabletype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _standardleveltabletimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelTableType", fieldName)
	}
	return
}

func (standardleveltype *StandardLevelType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardleveltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardleveltype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", standardleveltype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", standardleveltype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range standardleveltype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_35 := range standardleveltype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_35.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelType", fieldName)
	}
	return
}

func (telephonecontactpropertytype *TelephoneContactPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := telephonecontactpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", telephonecontactpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", telephonecontactpropertytype.NilReason)

	case "TelephoneContact":
		if telephonecontactpropertytype.TelephoneContact != nil {
			res = __gong__marshallPointer(ident, "TelephoneContact", telephonecontactpropertytype.TelephoneContact.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TelephoneContact", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TelephoneContactPropertyType", fieldName)
	}
	return
}

func (telephonecontacttype *TelephoneContactType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := telephonecontacttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", telephonecontacttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", telephonecontacttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", telephonecontacttype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range telephonecontacttype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range telephonecontacttype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range telephonecontacttype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_32 := range telephonecontacttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_32.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TelephoneContactType", fieldName)
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
		for _, _a_extension_34 := range timesheettype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_34.GongGetIdentifier(stage)))
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

func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unitavailabilitypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unitavailabilitypropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", unitavailabilitypropertytype.NilReason)

	case "UnitAvailability":
		if unitavailabilitypropertytype.UnitAvailability != nil {
			res = __gong__marshallPointer(ident, "UnitAvailability", unitavailabilitypropertytype.UnitAvailability.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "UnitAvailability", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct UnitAvailabilityPropertyType", fieldName)
	}
	return
}

func (unitavailabilitytype *UnitAvailabilityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unitavailabilitytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unitavailabilitytype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", unitavailabilitytype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", unitavailabilitytype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range unitavailabilitytype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range unitavailabilitytype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range unitavailabilitytype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_7 := range unitavailabilitytype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_7.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct UnitAvailabilityType", fieldName)
	}
	return
}

func (unitdependencypropertytype *UnitDependencyPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unitdependencypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unitdependencypropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", unitdependencypropertytype.NilReason)

	case "UnitDependency":
		if unitdependencypropertytype.UnitDependency != nil {
			res = __gong__marshallPointer(ident, "UnitDependency", unitdependencypropertytype.UnitDependency.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "UnitDependency", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct UnitDependencyPropertyType", fieldName)
	}
	return
}

func (unitdependencytype *UnitDependencyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unitdependencytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unitdependencytype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", unitdependencytype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", unitdependencytype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range unitdependencytype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheUnit":
		if unitdependencytype.TheUnit != nil {
			res = __gong__marshallPointer(ident, "TheUnit", unitdependencytype.TheUnit.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheUnit", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_8 := range unitdependencytype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_8.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct UnitDependencyType", fieldName)
	}
	return
}

func (unitpropertytype *UnitPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unitpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unitpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", unitpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", unitpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", unitpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", unitpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", unitpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", unitpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", unitpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", unitpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", unitpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct UnitPropertyType", fieldName)
	}
	return
}

func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unittimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unittimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", unittimeslicepropertytype.Owns)

	case "UnitTimeSlice":
		if unittimeslicepropertytype.UnitTimeSlice != nil {
			res = __gong__marshallPointer(ident, "UnitTimeSlice", unittimeslicepropertytype.UnitTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "UnitTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct UnitTimeSlicePropertyType", fieldName)
	}
	return
}

func (unittimeslicetype *UnitTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unittimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unittimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", unittimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", unittimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", unittimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", unittimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", unittimeslicetype.CorrectionNumber)

	case "Position":
		if unittimeslicetype.Position != nil {
			res = __gong__marshallPointer(ident, "Position", unittimeslicetype.Position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Position", "nil")
		}
	case "AirportLocation":
		if unittimeslicetype.AirportLocation != nil {
			res = __gong__marshallPointer(ident, "AirportLocation", unittimeslicetype.AirportLocation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirportLocation", "nil")
		}
	case "OwnerOrganisation":
		if unittimeslicetype.OwnerOrganisation != nil {
			res = __gong__marshallPointer(ident, "OwnerOrganisation", unittimeslicetype.OwnerOrganisation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OwnerOrganisation", "nil")
		}
	case "Contact":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range unittimeslicetype.Contact {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Contact", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RelatedUnit":
		var sb strings.Builder
		for _, _unitdependencypropertytype := range unittimeslicetype.RelatedUnit {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RelatedUnit", _unitdependencypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _unitavailabilitypropertytype := range unittimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _unitavailabilitypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range unittimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_6 := range unittimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_6.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct UnitTimeSliceType", fieldName)
	}
	return
}

func (unittype *UnitType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unittype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unittype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", unittype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", unittype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _unittimeslicepropertytype := range unittype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _unittimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct UnitType", fieldName)
	}
	return
}

func (vorpropertytype *VORPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vorpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vorpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", vorpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", vorpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", vorpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", vorpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", vorpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", vorpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", vorpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", vorpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", vorpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct VORPropertyType", fieldName)
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
func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractairportgroundservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractairportgroundservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractairportgroundservicetype.GongMarshallField(stage, "Description"))
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
func (abstractservicetype *AbstractServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractservicetype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstracttrafficseparationservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstracttrafficseparationservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstracttrafficseparationservicetype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aerialrefuellingpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aerialrefuellingpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(aerialrefuellingpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(aerialrefuellingpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(aerialrefuellingpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(aerialrefuellingpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(aerialrefuellingpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(aerialrefuellingpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(aerialrefuellingpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(aerialrefuellingpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airtrafficcontrolservicetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airtrafficcontrolservicetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicepropertytype.GongMarshallField(stage, "AirTrafficControlServiceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "ServiceProvider"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Call_sign"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "RadioCommunication"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "GroundCommunication"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "ClientAirport"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "ClientAirspace"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "ClientRoute"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "ClientProcedure"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "ClientHolding"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "ClientAerialRefuelling"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "AircraftLocator"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airtrafficcontrolservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airtrafficcontrolservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airtrafficcontrolservicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airtrafficcontrolservicetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airtrafficmanagementservicetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airtrafficmanagementservicetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicepropertytype.GongMarshallField(stage, "AirTrafficManagementServiceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "ServiceProvider"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "Call_sign"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "RadioCommunication"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "GroundCommunication"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "ClientAirspace"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "ClientAerialRefuelling"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "ClientRoute"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airtrafficmanagementservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airtrafficmanagementservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airtrafficmanagementservicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airtrafficmanagementservicetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aircraftgroundservicetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aircraftgroundservicetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetimeslicepropertytype.GongMarshallField(stage, "AircraftGroundServiceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "ServiceProvider"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "Call_sign"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "RadioCommunication"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "GroundCommunication"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aircraftgroundservicetype *AircraftGroundServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aircraftgroundservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aircraftgroundservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aircraftgroundservicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(aircraftgroundservicetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airportclearanceservicetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airportclearanceservicetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(airportclearanceservicetimeslicepropertytype.GongMarshallField(stage, "AirportClearanceServiceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "ServiceProvider"))
		pointersInitializesStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "Call_sign"))
		pointersInitializesStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "RadioCommunication"))
		pointersInitializesStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "GroundCommunication"))
		pointersInitializesStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(airportclearanceservicetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airportclearanceservicetype *AirportClearanceServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airportclearanceservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airportclearanceservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airportclearanceservicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airportclearanceservicetype.GongMarshallField(stage, "TimeSlice"))
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
func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airportsuppliesservicetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airportsuppliesservicetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicepropertytype.GongMarshallField(stage, "AirportSuppliesServiceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "ServiceProvider"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "Call_sign"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "RadioCommunication"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "GroundCommunication"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "FuelSupply"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "OilSupply"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "NitrogenSupply"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "OxygenSupply"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airportsuppliesservicetype *AirportSuppliesServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airportsuppliesservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airportsuppliesservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airportsuppliesservicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airportsuppliesservicetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacepropertytype *AirspacePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (callsigndetailpropertytype *CallsignDetailPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(callsigndetailpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(callsigndetailpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(callsigndetailpropertytype.GongMarshallField(stage, "CallsignDetail"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (callsigndetailtype *CallsignDetailType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(callsigndetailtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(callsigndetailtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(callsigndetailtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(callsigndetailtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(callsigndetailtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (contactinformationpropertytype *ContactInformationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(contactinformationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(contactinformationpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(contactinformationpropertytype.GongMarshallField(stage, "ContactInformation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (contactinformationtype *ContactInformationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(contactinformationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(contactinformationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(contactinformationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(contactinformationtype.GongMarshallField(stage, "Address"))
		pointersInitializesStatements.WriteString(contactinformationtype.GongMarshallField(stage, "NetworkNode"))
		pointersInitializesStatements.WriteString(contactinformationtype.GongMarshallField(stage, "PhoneFax"))
		pointersInitializesStatements.WriteString(contactinformationtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(contactinformationtype.GongMarshallField(stage, "Extension"))
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
func (directionfinderpropertytype *DirectionFinderPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directionfinderpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directionfinderpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(directionfinderpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(directionfinderpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(directionfinderpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(directionfinderpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(directionfinderpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(directionfinderpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(directionfinderpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(directionfinderpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevatedpointpropertytype *ElevatedPointPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevatedpointpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevatedpointpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(elevatedpointpropertytype.GongMarshallField(stage, "ElevatedPoint"))
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
func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(firefightingservicetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(firefightingservicetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(firefightingservicetimeslicepropertytype.GongMarshallField(stage, "FireFightingServiceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "ServiceProvider"))
		pointersInitializesStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "Call_sign"))
		pointersInitializesStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "RadioCommunication"))
		pointersInitializesStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "GroundCommunication"))
		pointersInitializesStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(firefightingservicetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (firefightingservicetype *FireFightingServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(firefightingservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(firefightingservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(firefightingservicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(firefightingservicetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (fuelpropertytype *FuelPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(fuelpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(fuelpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(fuelpropertytype.GongMarshallField(stage, "Fuel"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (fueltype *FuelType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(fueltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(fueltype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(fueltype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(fueltype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(fueltype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(groundtrafficcontrolservicetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(groundtrafficcontrolservicetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetimeslicepropertytype.GongMarshallField(stage, "GroundTrafficControlServiceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "ServiceProvider"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Call_sign"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "RadioCommunication"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "GroundCommunication"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "ClientAirport"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(groundtrafficcontrolservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(groundtrafficcontrolservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(groundtrafficcontrolservicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(groundtrafficcontrolservicetype.GongMarshallField(stage, "TimeSlice"))
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
func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(informationservicetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(informationservicetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(informationservicetimeslicepropertytype.GongMarshallField(stage, "InformationServiceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "ServiceProvider"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "Call_sign"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "RadioCommunication"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "GroundCommunication"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "NavaidBroadcast"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "ClientAirspace"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "ClientAirport"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "ClientRoute"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "ClientProcedure"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "ClientHolding"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "ClientAerialRefuelling"))
		pointersInitializesStatements.WriteString(informationservicetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (informationservicetype *InformationServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(informationservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(informationservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(informationservicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(informationservicetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (lightactivationpropertytype *LightActivationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(lightactivationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(lightactivationpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(lightactivationpropertytype.GongMarshallField(stage, "LightActivation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (lightactivationtype *LightActivationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(lightactivationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(lightactivationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(lightactivationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(lightactivationtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(lightactivationtype.GongMarshallField(stage, "Extension"))
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
func (nitrogenpropertytype *NitrogenPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(nitrogenpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(nitrogenpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(nitrogenpropertytype.GongMarshallField(stage, "Nitrogen"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (nitrogentype *NitrogenType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(nitrogentype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(nitrogentype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(nitrogentype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(nitrogentype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(nitrogentype.GongMarshallField(stage, "Extension"))
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
func (oilpropertytype *OilPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(oilpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(oilpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(oilpropertytype.GongMarshallField(stage, "Oil"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (oiltype *OilType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(oiltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(oiltype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(oiltype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(oiltype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(oiltype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (onlinecontactpropertytype *OnlineContactPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(onlinecontactpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(onlinecontactpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(onlinecontactpropertytype.GongMarshallField(stage, "OnlineContact"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (onlinecontacttype *OnlineContactType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(onlinecontacttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(onlinecontacttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(onlinecontacttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(onlinecontacttype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(onlinecontacttype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(onlinecontacttype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(onlinecontacttype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(organisationauthorityassociationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(organisationauthorityassociationpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(organisationauthorityassociationpropertytype.GongMarshallField(stage, "OrganisationAuthorityAssociation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(organisationauthorityassociationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(organisationauthorityassociationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(organisationauthorityassociationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(organisationauthorityassociationtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(organisationauthorityassociationtype.GongMarshallField(stage, "TheOrganisationAuthority"))
		pointersInitializesStatements.WriteString(organisationauthorityassociationtype.GongMarshallField(stage, "Extension"))
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
func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(organisationauthoritytimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(organisationauthoritytimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(organisationauthoritytimeslicepropertytype.GongMarshallField(stage, "OrganisationAuthorityTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(organisationauthoritytimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(organisationauthoritytimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(organisationauthoritytimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(organisationauthoritytimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(organisationauthoritytimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(organisationauthoritytimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(organisationauthoritytimeslicetype.GongMarshallField(stage, "Contact"))
		pointersInitializesStatements.WriteString(organisationauthoritytimeslicetype.GongMarshallField(stage, "RelatedOrganisationAuthority"))
		pointersInitializesStatements.WriteString(organisationauthoritytimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(organisationauthoritytimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (organisationauthoritytype *OrganisationAuthorityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(organisationauthoritytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(organisationauthoritytype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(organisationauthoritytype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(organisationauthoritytype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (oxygenpropertytype *OxygenPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(oxygenpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(oxygenpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(oxygenpropertytype.GongMarshallField(stage, "Oxygen"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (oxygentype *OxygenType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(oxygentype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(oxygentype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(oxygentype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(oxygentype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(oxygentype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(passengerservicetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(passengerservicetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(passengerservicetimeslicepropertytype.GongMarshallField(stage, "PassengerServiceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "ServiceProvider"))
		pointersInitializesStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "Call_sign"))
		pointersInitializesStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "RadioCommunication"))
		pointersInitializesStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "GroundCommunication"))
		pointersInitializesStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(passengerservicetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (passengerservicetype *PassengerServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(passengerservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(passengerservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(passengerservicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(passengerservicetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pilotcontrolledlightingtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pilotcontrolledlightingtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(pilotcontrolledlightingtimeslicepropertytype.GongMarshallField(stage, "PilotControlledLightingTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pilotcontrolledlightingtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pilotcontrolledlightingtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(pilotcontrolledlightingtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(pilotcontrolledlightingtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(pilotcontrolledlightingtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(pilotcontrolledlightingtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(pilotcontrolledlightingtimeslicetype.GongMarshallField(stage, "ControlledLightIntensity"))
		pointersInitializesStatements.WriteString(pilotcontrolledlightingtimeslicetype.GongMarshallField(stage, "ActivatedGroundLighting"))
		pointersInitializesStatements.WriteString(pilotcontrolledlightingtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(pilotcontrolledlightingtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pilotcontrolledlightingtype *PilotControlledLightingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pilotcontrolledlightingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pilotcontrolledlightingtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(pilotcontrolledlightingtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(pilotcontrolledlightingtype.GongMarshallField(stage, "TimeSlice"))
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
func (postaladdresspropertytype *PostalAddressPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(postaladdresspropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(postaladdresspropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(postaladdresspropertytype.GongMarshallField(stage, "PostalAddress"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (postaladdresstype *PostalAddressType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(postaladdresstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(postaladdresstype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(postaladdresstype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(postaladdresstype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(postaladdresstype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(postaladdresstype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(postaladdresstype.GongMarshallField(stage, "Extension"))
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
func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radiocommunicationchannelpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radiocommunicationchannelpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(radiocommunicationchannelpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(radiocommunicationchannelpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(radiocommunicationchannelpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(radiocommunicationchannelpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(radiocommunicationchannelpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(radiocommunicationchannelpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(radiocommunicationchannelpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(radiocommunicationchannelpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radiocommunicationchanneltimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radiocommunicationchanneltimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(radiocommunicationchanneltimeslicepropertytype.GongMarshallField(stage, "RadioCommunicationChannelTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radiocommunicationchanneltimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radiocommunicationchanneltimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(radiocommunicationchanneltimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(radiocommunicationchanneltimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(radiocommunicationchanneltimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(radiocommunicationchanneltimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(radiocommunicationchanneltimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(radiocommunicationchanneltimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(radiocommunicationchanneltimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(radiocommunicationchanneltimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radiocommunicationchanneltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radiocommunicationchanneltype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(radiocommunicationchanneltype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(radiocommunicationchanneltype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radiocommunicationoperationalstatuspropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radiocommunicationoperationalstatuspropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(radiocommunicationoperationalstatuspropertytype.GongMarshallField(stage, "RadioCommunicationOperationalStatus"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radiocommunicationoperationalstatustype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radiocommunicationoperationalstatustype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(radiocommunicationoperationalstatustype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(radiocommunicationoperationalstatustype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(radiocommunicationoperationalstatustype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(radiocommunicationoperationalstatustype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(radiocommunicationoperationalstatustype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routeportionpropertytype *RoutePortionPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routeportionpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routeportionpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(routeportionpropertytype.GongMarshallField(stage, "RoutePortion"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routeportiontype *RoutePortionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routeportiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routeportiontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(routeportiontype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "Start_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "Start_navaidSystem"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "Start_position"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "Start_runwayPoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "Start_aimingPoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "Start_airportReferencePoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "IntermediatePoint_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "IntermediatePoint_navaidSystem"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "IntermediatePoint_position"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "IntermediatePoint_runwayPoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "IntermediatePoint_aimingPoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "IntermediatePoint_airportReferencePoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "ReferencedRoute"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "End_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "End_navaidSystem"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "End_position"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "End_runwayPoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "End_aimingPoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "End_airportReferencePoint"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(routeportiontype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routepropertytype *RoutePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Actuate"))
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
func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(searchrescueservicetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(searchrescueservicetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicepropertytype.GongMarshallField(stage, "SearchRescueServiceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "ServiceProvider"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "Call_sign"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "RadioCommunication"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "GroundCommunication"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "ClientAirspace"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "ClientRoute"))
		pointersInitializesStatements.WriteString(searchrescueservicetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (searchrescueservicetype *SearchRescueServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(searchrescueservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(searchrescueservicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(searchrescueservicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(searchrescueservicetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(serviceoperationalstatuspropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(serviceoperationalstatuspropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(serviceoperationalstatuspropertytype.GongMarshallField(stage, "ServiceOperationalStatus"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (serviceoperationalstatustype *ServiceOperationalStatusType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(serviceoperationalstatustype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(serviceoperationalstatustype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(serviceoperationalstatustype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(serviceoperationalstatustype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(serviceoperationalstatustype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(serviceoperationalstatustype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(serviceoperationalstatustype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialdatetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialdatetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(specialdatetimeslicepropertytype.GongMarshallField(stage, "SpecialDateTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialdatetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialdatetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(specialdatetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(specialdatetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(specialdatetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(specialdatetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(specialdatetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(specialdatetimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(specialdatetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialdatetype *SpecialDateType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialdatetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialdatetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(specialdatetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(specialdatetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardlevelcolumntimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardlevelcolumntimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(standardlevelcolumntimeslicepropertytype.GongMarshallField(stage, "StandardLevelColumnTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		initializerStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "UnitOfMeasurement"))
		pointersInitializesStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "Level"))
		pointersInitializesStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "LevelTable"))
		pointersInitializesStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(standardlevelcolumntimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardlevelcolumntype *StandardLevelColumnType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardlevelcolumntype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardlevelcolumntype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardlevelcolumntype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(standardlevelcolumntype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardlevelpropertytype *StandardLevelPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardlevelpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardlevelpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(standardlevelpropertytype.GongMarshallField(stage, "StandardLevel"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardlevelsectortimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardlevelsectortimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(standardlevelsectortimeslicepropertytype.GongMarshallField(stage, "StandardLevelSectorTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardlevelsectortimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardlevelsectortimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardlevelsectortimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(standardlevelsectortimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(standardlevelsectortimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(standardlevelsectortimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(standardlevelsectortimeslicetype.GongMarshallField(stage, "ApplicableAirspace"))
		pointersInitializesStatements.WriteString(standardlevelsectortimeslicetype.GongMarshallField(stage, "ApplicableLevelColumn"))
		pointersInitializesStatements.WriteString(standardlevelsectortimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(standardlevelsectortimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardlevelsectortype *StandardLevelSectorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardlevelsectortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardlevelsectortype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardlevelsectortype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(standardlevelsectortype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardleveltablepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardleveltablepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(standardleveltablepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(standardleveltablepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(standardleveltablepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(standardleveltablepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(standardleveltablepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(standardleveltablepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(standardleveltablepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(standardleveltablepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardleveltabletimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardleveltabletimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(standardleveltabletimeslicepropertytype.GongMarshallField(stage, "StandardLevelTableTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardleveltabletimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardleveltabletimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardleveltabletimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(standardleveltabletimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(standardleveltabletimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(standardleveltabletimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(standardleveltabletimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(standardleveltabletimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardleveltabletype *StandardLevelTableType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardleveltabletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardleveltabletype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardleveltabletype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(standardleveltabletype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardleveltype *StandardLevelType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardleveltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardleveltype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(standardleveltype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(standardleveltype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(standardleveltype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (telephonecontactpropertytype *TelephoneContactPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(telephonecontactpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(telephonecontactpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(telephonecontactpropertytype.GongMarshallField(stage, "TelephoneContact"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (telephonecontacttype *TelephoneContactType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(telephonecontacttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(telephonecontacttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(telephonecontacttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(telephonecontacttype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(telephonecontacttype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(telephonecontacttype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(telephonecontacttype.GongMarshallField(stage, "Extension"))
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
func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unitavailabilitypropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unitavailabilitypropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(unitavailabilitypropertytype.GongMarshallField(stage, "UnitAvailability"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unitavailabilitytype *UnitAvailabilityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unitavailabilitytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unitavailabilitytype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(unitavailabilitytype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(unitavailabilitytype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(unitavailabilitytype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(unitavailabilitytype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(unitavailabilitytype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unitdependencypropertytype *UnitDependencyPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unitdependencypropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unitdependencypropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(unitdependencypropertytype.GongMarshallField(stage, "UnitDependency"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unitdependencytype *UnitDependencyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unitdependencytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unitdependencytype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(unitdependencytype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(unitdependencytype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(unitdependencytype.GongMarshallField(stage, "TheUnit"))
		pointersInitializesStatements.WriteString(unitdependencytype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unitpropertytype *UnitPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unittimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unittimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(unittimeslicepropertytype.GongMarshallField(stage, "UnitTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unittimeslicetype *UnitTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "Position"))
		pointersInitializesStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "AirportLocation"))
		pointersInitializesStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "OwnerOrganisation"))
		pointersInitializesStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "Contact"))
		pointersInitializesStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "RelatedUnit"))
		pointersInitializesStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(unittimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unittype *UnitType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unittype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unittype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(unittype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(unittype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (vorpropertytype *VORPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vorpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(vorpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(vorpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(vorpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(vorpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(vorpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(vorpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(vorpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(vorpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(vorpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
