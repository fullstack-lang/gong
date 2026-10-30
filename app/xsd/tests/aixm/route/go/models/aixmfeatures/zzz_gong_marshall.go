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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures"
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
		for _, abstractdirectflighttype := range __gong__sortStageSetInstances(stage.AbstractDirectFlightTypes, stage.AbstractDirectFlightType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractdirectflighttypeIdent := abstractdirectflighttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractDirectFlightType{Name: %s}).Stage(stage)", abstractdirectflighttypeIdent, __gong__toRawStringLiteral(abstractdirectflighttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractdirectflighttypeIdent, __gong__toRawStringLiteral(abstractdirectflighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractdirectflighttypeIdent, __gong__toRawStringLiteral(abstractdirectflighttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractdirectflighttypeIdent, __gong__toRawStringLiteral(abstractdirectflighttype.Description)))
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
		for _, aerialrefuellinganchorpropertytype := range __gong__sortStageSetInstances(stage.AerialRefuellingAnchorPropertyTypes, stage.AerialRefuellingAnchorPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aerialrefuellinganchorpropertytypeIdent := aerialrefuellinganchorpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingAnchorPropertyType{Name: %s}).Stage(stage)", aerialrefuellinganchorpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchorpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellinganchorpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchorpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", aerialrefuellinganchorpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchorpropertytype.NilReason)))
			if aerialrefuellinganchorpropertytype.AerialRefuellingAnchor != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellinganchorpropertytype.AerialRefuellingAnchor.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerialRefuellingAnchor = %s", aerialrefuellinganchorpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aerialrefuellinganchortype := range __gong__sortStageSetInstances(stage.AerialRefuellingAnchorTypes, stage.AerialRefuellingAnchorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aerialrefuellinganchortypeIdent := aerialrefuellinganchortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingAnchorType{Name: %s}).Stage(stage)", aerialrefuellinganchortypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellinganchortypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aerialrefuellinganchortypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aerialrefuellinganchortypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchortype.Description)))
			if aerialrefuellinganchortype.Extent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellinganchortype.Extent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extent = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellinganchortype.VerticalExtent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalExtent = append(%s.VerticalExtent, %s)", aerialrefuellinganchortypeIdent, aerialrefuellinganchortypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellinganchortype.Point {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Point = append(%s.Point, %s)", aerialrefuellinganchortypeIdent, aerialrefuellinganchortypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellinganchortype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aerialrefuellinganchortypeIdent, aerialrefuellinganchortypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellinganchortype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aerialrefuellinganchortypeIdent, aerialrefuellinganchortypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aerialrefuellingpointpropertytype := range __gong__sortStageSetInstances(stage.AerialRefuellingPointPropertyTypes, stage.AerialRefuellingPointPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aerialrefuellingpointpropertytypeIdent := aerialrefuellingpointpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingPointPropertyType{Name: %s}).Stage(stage)", aerialrefuellingpointpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingpointpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", aerialrefuellingpointpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointpropertytype.NilReason)))
			if aerialrefuellingpointpropertytype.AerialRefuellingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingpointpropertytype.AerialRefuellingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerialRefuellingPoint = %s", aerialrefuellingpointpropertytypeIdent, targetIdent))
			}
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
		for _, aerialrefuellingtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AerialRefuellingTimeSlicePropertyTypes, stage.AerialRefuellingTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aerialrefuellingtimeslicepropertytypeIdent := aerialrefuellingtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingTimeSlicePropertyType{Name: %s}).Stage(stage)", aerialrefuellingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", aerialrefuellingtimeslicepropertytypeIdent, aerialrefuellingtimeslicepropertytype.Owns))
			if aerialrefuellingtimeslicepropertytype.AerialRefuellingTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingtimeslicepropertytype.AerialRefuellingTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerialRefuellingTimeSlice = %s", aerialrefuellingtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aerialrefuellingtimeslicetype := range __gong__sortStageSetInstances(stage.AerialRefuellingTimeSliceTypes, stage.AerialRefuellingTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aerialrefuellingtimeslicetypeIdent := aerialrefuellingtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingTimeSliceType{Name: %s}).Stage(stage)", aerialrefuellingtimeslicetypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingtimeslicetypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aerialrefuellingtimeslicetypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aerialrefuellingtimeslicetypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", aerialrefuellingtimeslicetypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetype.CorrectionNumber))
			for _, elem := range aerialrefuellingtimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.ProtectingAirspace {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProtectingAirspace = append(%s.ProtectingAirspace, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.Track {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Track = append(%s.Track, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.Anchor {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anchor = append(%s.Anchor, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.OppositeTrack != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingtimeslicetype.OppositeTrack.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OppositeTrack = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.ManagingOrganisation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ManagingOrganisation = append(%s.ManagingOrganisation, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aerialrefuellingtrackpropertytype := range __gong__sortStageSetInstances(stage.AerialRefuellingTrackPropertyTypes, stage.AerialRefuellingTrackPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aerialrefuellingtrackpropertytypeIdent := aerialrefuellingtrackpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingTrackPropertyType{Name: %s}).Stage(stage)", aerialrefuellingtrackpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingtrackpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingtrackpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingtrackpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", aerialrefuellingtrackpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingtrackpropertytype.NilReason)))
			if aerialrefuellingtrackpropertytype.AerialRefuellingTrack != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingtrackpropertytype.AerialRefuellingTrack.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerialRefuellingTrack = %s", aerialrefuellingtrackpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aerialrefuellingtracktype := range __gong__sortStageSetInstances(stage.AerialRefuellingTrackTypes, stage.AerialRefuellingTrackType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aerialrefuellingtracktypeIdent := aerialrefuellingtracktype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingTrackType{Name: %s}).Stage(stage)", aerialrefuellingtracktypeIdent, __gong__toRawStringLiteral(aerialrefuellingtracktype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingtracktypeIdent, __gong__toRawStringLiteral(aerialrefuellingtracktype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aerialrefuellingtracktypeIdent, __gong__toRawStringLiteral(aerialrefuellingtracktype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aerialrefuellingtracktypeIdent, __gong__toRawStringLiteral(aerialrefuellingtracktype.Description)))
			if aerialrefuellingtracktype.Extent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aerialrefuellingtracktype.Extent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extent = %s", aerialrefuellingtracktypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtracktype.Point {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Point = append(%s.Point, %s)", aerialrefuellingtracktypeIdent, aerialrefuellingtracktypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtracktype.VerticalExtent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalExtent = append(%s.VerticalExtent, %s)", aerialrefuellingtracktypeIdent, aerialrefuellingtracktypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtracktype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aerialrefuellingtracktypeIdent, aerialrefuellingtracktypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtracktype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aerialrefuellingtracktypeIdent, aerialrefuellingtracktypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aerialrefuellingtype := range __gong__sortStageSetInstances(stage.AerialRefuellingTypes, stage.AerialRefuellingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aerialrefuellingtypeIdent := aerialrefuellingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingType{Name: %s}).Stage(stage)", aerialrefuellingtypeIdent, __gong__toRawStringLiteral(aerialrefuellingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingtypeIdent, __gong__toRawStringLiteral(aerialrefuellingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aerialrefuellingtypeIdent, __gong__toRawStringLiteral(aerialrefuellingtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aerialrefuellingtypeIdent, __gong__toRawStringLiteral(aerialrefuellingtype.Description)))
			for _, elem := range aerialrefuellingtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", aerialrefuellingtypeIdent, aerialrefuellingtypeIdent, targetIdent))
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
		for _, airspacebordercrossingpropertytype := range __gong__sortStageSetInstances(stage.AirspaceBorderCrossingPropertyTypes, stage.AirspaceBorderCrossingPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacebordercrossingpropertytypeIdent := airspacebordercrossingpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceBorderCrossingPropertyType{Name: %s}).Stage(stage)", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airspacebordercrossingpropertytypeIdent, airspacebordercrossingpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, airspacebordercrossingtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AirspaceBorderCrossingTimeSlicePropertyTypes, stage.AirspaceBorderCrossingTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacebordercrossingtimeslicepropertytypeIdent := airspacebordercrossingtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceBorderCrossingTimeSlicePropertyType{Name: %s}).Stage(stage)", airspacebordercrossingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacebordercrossingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airspacebordercrossingtimeslicepropertytypeIdent, airspacebordercrossingtimeslicepropertytype.Owns))
			if airspacebordercrossingtimeslicepropertytype.AirspaceBorderCrossingTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacebordercrossingtimeslicepropertytype.AirspaceBorderCrossingTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceBorderCrossingTimeSlice = %s", airspacebordercrossingtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacebordercrossingtimeslicetype := range __gong__sortStageSetInstances(stage.AirspaceBorderCrossingTimeSliceTypes, stage.AirspaceBorderCrossingTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacebordercrossingtimeslicetypeIdent := airspacebordercrossingtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceBorderCrossingTimeSliceType{Name: %s}).Stage(stage)", airspacebordercrossingtimeslicetypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacebordercrossingtimeslicetypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacebordercrossingtimeslicetypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacebordercrossingtimeslicetypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airspacebordercrossingtimeslicetypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airspacebordercrossingtimeslicetypeIdent, airspacebordercrossingtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airspacebordercrossingtimeslicetypeIdent, airspacebordercrossingtimeslicetype.CorrectionNumber))
			if airspacebordercrossingtimeslicetype.ExitedAirspace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacebordercrossingtimeslicetype.ExitedAirspace.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExitedAirspace = %s", airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			if airspacebordercrossingtimeslicetype.EnteredAirspace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacebordercrossingtimeslicetype.EnteredAirspace.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnteredAirspace = %s", airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacebordercrossingtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacebordercrossingtimeslicetypeIdent, airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacebordercrossingtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacebordercrossingtimeslicetypeIdent, airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacebordercrossingtype := range __gong__sortStageSetInstances(stage.AirspaceBorderCrossingTypes, stage.AirspaceBorderCrossingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacebordercrossingtypeIdent := airspacebordercrossingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceBorderCrossingType{Name: %s}).Stage(stage)", airspacebordercrossingtypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacebordercrossingtypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacebordercrossingtypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacebordercrossingtypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtype.Description)))
			for _, elem := range airspacebordercrossingtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airspacebordercrossingtypeIdent, airspacebordercrossingtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacelayerpropertytype := range __gong__sortStageSetInstances(stage.AirspaceLayerPropertyTypes, stage.AirspaceLayerPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacelayerpropertytypeIdent := airspacelayerpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceLayerPropertyType{Name: %s}).Stage(stage)", airspacelayerpropertytypeIdent, __gong__toRawStringLiteral(airspacelayerpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacelayerpropertytypeIdent, __gong__toRawStringLiteral(airspacelayerpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacelayerpropertytypeIdent, __gong__toRawStringLiteral(airspacelayerpropertytype.NilReason)))
			if airspacelayerpropertytype.AirspaceLayer != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacelayerpropertytype.AirspaceLayer.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceLayer = %s", airspacelayerpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacelayertype := range __gong__sortStageSetInstances(stage.AirspaceLayerTypes, stage.AirspaceLayerType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacelayertypeIdent := airspacelayertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceLayerType{Name: %s}).Stage(stage)", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Description)))
			if airspacelayertype.DiscreteLevelSeries != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacelayertype.DiscreteLevelSeries.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiscreteLevelSeries = %s", airspacelayertypeIdent, targetIdent))
			}
			for _, elem := range airspacelayertype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacelayertypeIdent, airspacelayertypeIdent, targetIdent))
			}
			for _, elem := range airspacelayertype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacelayertypeIdent, airspacelayertypeIdent, targetIdent))
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
		for _, authorityforaerialrefuellingpropertytype := range __gong__sortStageSetInstances(stage.AuthorityForAerialRefuellingPropertyTypes, stage.AuthorityForAerialRefuellingPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityforaerialrefuellingpropertytypeIdent := authorityforaerialrefuellingpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForAerialRefuellingPropertyType{Name: %s}).Stage(stage)", authorityforaerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforaerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", authorityforaerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingpropertytype.NilReason)))
			if authorityforaerialrefuellingpropertytype.AuthorityForAerialRefuelling != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityforaerialrefuellingpropertytype.AuthorityForAerialRefuelling.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForAerialRefuelling = %s", authorityforaerialrefuellingpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, authorityforaerialrefuellingtype := range __gong__sortStageSetInstances(stage.AuthorityForAerialRefuellingTypes, stage.AuthorityForAerialRefuellingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityforaerialrefuellingtypeIdent := authorityforaerialrefuellingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForAerialRefuellingType{Name: %s}).Stage(stage)", authorityforaerialrefuellingtypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforaerialrefuellingtypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityforaerialrefuellingtypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityforaerialrefuellingtypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingtype.Description)))
			for _, elem := range authorityforaerialrefuellingtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", authorityforaerialrefuellingtypeIdent, authorityforaerialrefuellingtypeIdent, targetIdent))
			}
			if authorityforaerialrefuellingtype.TheOrganisationAuthority != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityforaerialrefuellingtype.TheOrganisationAuthority.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheOrganisationAuthority = %s", authorityforaerialrefuellingtypeIdent, targetIdent))
			}
			for _, elem := range authorityforaerialrefuellingtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", authorityforaerialrefuellingtypeIdent, authorityforaerialrefuellingtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, changeoverpointtimeslicepropertytype := range __gong__sortStageSetInstances(stage.ChangeOverPointTimeSlicePropertyTypes, stage.ChangeOverPointTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			changeoverpointtimeslicepropertytypeIdent := changeoverpointtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ChangeOverPointTimeSlicePropertyType{Name: %s}).Stage(stage)", changeoverpointtimeslicepropertytypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", changeoverpointtimeslicepropertytypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", changeoverpointtimeslicepropertytypeIdent, changeoverpointtimeslicepropertytype.Owns))
			if changeoverpointtimeslicepropertytype.ChangeOverPointTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := changeoverpointtimeslicepropertytype.ChangeOverPointTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ChangeOverPointTimeSlice = %s", changeoverpointtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, changeoverpointtimeslicetype := range __gong__sortStageSetInstances(stage.ChangeOverPointTimeSliceTypes, stage.ChangeOverPointTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			changeoverpointtimeslicetypeIdent := changeoverpointtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ChangeOverPointTimeSliceType{Name: %s}).Stage(stage)", changeoverpointtimeslicetypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", changeoverpointtimeslicetypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", changeoverpointtimeslicetypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", changeoverpointtimeslicetypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", changeoverpointtimeslicetypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", changeoverpointtimeslicetypeIdent, changeoverpointtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", changeoverpointtimeslicetypeIdent, changeoverpointtimeslicetype.CorrectionNumber))
			if changeoverpointtimeslicetype.Location_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_fixDesignatedPoint = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_navaidSystem = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_position = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_runwayPoint = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_aimingPoint = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_airportReferencePoint = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.ApplicableRoutePortion != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := changeoverpointtimeslicetype.ApplicableRoutePortion.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApplicableRoutePortion = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range changeoverpointtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", changeoverpointtimeslicetypeIdent, changeoverpointtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range changeoverpointtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", changeoverpointtimeslicetypeIdent, changeoverpointtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, changeoverpointtype := range __gong__sortStageSetInstances(stage.ChangeOverPointTypes, stage.ChangeOverPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			changeoverpointtypeIdent := changeoverpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ChangeOverPointType{Name: %s}).Stage(stage)", changeoverpointtypeIdent, __gong__toRawStringLiteral(changeoverpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", changeoverpointtypeIdent, __gong__toRawStringLiteral(changeoverpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", changeoverpointtypeIdent, __gong__toRawStringLiteral(changeoverpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", changeoverpointtypeIdent, __gong__toRawStringLiteral(changeoverpointtype.Description)))
			for _, elem := range changeoverpointtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", changeoverpointtypeIdent, changeoverpointtypeIdent, targetIdent))
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
		for _, directflightclasstype := range __gong__sortStageSetInstances(stage.DirectFlightClassTypes, stage.DirectFlightClassType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directflightclasstypeIdent := directflightclasstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectFlightClassType{Name: %s}).Stage(stage)", directflightclasstypeIdent, __gong__toRawStringLiteral(directflightclasstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directflightclasstypeIdent, __gong__toRawStringLiteral(directflightclasstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directflightclasstypeIdent, __gong__toRawStringLiteral(directflightclasstype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", directflightclasstypeIdent, __gong__toRawStringLiteral(directflightclasstype.Description)))
			for _, elem := range directflightclasstype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", directflightclasstypeIdent, directflightclasstypeIdent, targetIdent))
			}
			for _, elem := range directflightclasstype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", directflightclasstypeIdent, directflightclasstypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, directflightpropertytype := range __gong__sortStageSetInstances(stage.DirectFlightPropertyTypes, stage.DirectFlightPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directflightpropertytypeIdent := directflightpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectFlightPropertyType{Name: %s}).Stage(stage)", directflightpropertytypeIdent, __gong__toRawStringLiteral(directflightpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directflightpropertytypeIdent, __gong__toRawStringLiteral(directflightpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", directflightpropertytypeIdent, __gong__toRawStringLiteral(directflightpropertytype.NilReason)))
			if directflightpropertytype.DirectFlightClass != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightpropertytype.DirectFlightClass.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectFlightClass = %s", directflightpropertytypeIdent, targetIdent))
			}
			if directflightpropertytype.DirectFlightSegment != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightpropertytype.DirectFlightSegment.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectFlightSegment = %s", directflightpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, directflightsegmentpropertytype := range __gong__sortStageSetInstances(stage.DirectFlightSegmentPropertyTypes, stage.DirectFlightSegmentPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directflightsegmentpropertytypeIdent := directflightsegmentpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectFlightSegmentPropertyType{Name: %s}).Stage(stage)", directflightsegmentpropertytypeIdent, __gong__toRawStringLiteral(directflightsegmentpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directflightsegmentpropertytypeIdent, __gong__toRawStringLiteral(directflightsegmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", directflightsegmentpropertytypeIdent, __gong__toRawStringLiteral(directflightsegmentpropertytype.NilReason)))
			if directflightsegmentpropertytype.DirectFlightSegment != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmentpropertytype.DirectFlightSegment.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectFlightSegment = %s", directflightsegmentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, directflightsegmenttype := range __gong__sortStageSetInstances(stage.DirectFlightSegmentTypes, stage.DirectFlightSegmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directflightsegmenttypeIdent := directflightsegmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectFlightSegmentType{Name: %s}).Stage(stage)", directflightsegmenttypeIdent, __gong__toRawStringLiteral(directflightsegmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directflightsegmenttypeIdent, __gong__toRawStringLiteral(directflightsegmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directflightsegmenttypeIdent, __gong__toRawStringLiteral(directflightsegmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", directflightsegmenttypeIdent, __gong__toRawStringLiteral(directflightsegmenttype.Description)))
			for _, elem := range directflightsegmenttype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", directflightsegmenttypeIdent, directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.End_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_fixDesignatedPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.End_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_navaidSystem = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.End_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_position = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.End_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_runwayPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.End_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_aimingPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.End_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_airportReferencePoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.Start_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_fixDesignatedPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.Start_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_navaidSystem = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.Start_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_position = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.Start_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_runwayPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.Start_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_aimingPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directflightsegmenttype.Start_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_airportReferencePoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			for _, elem := range directflightsegmenttype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", directflightsegmenttypeIdent, directflightsegmenttypeIdent, targetIdent))
			}
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
		for _, enroutesegmentpointpropertytype := range __gong__sortStageSetInstances(stage.EnRouteSegmentPointPropertyTypes, stage.EnRouteSegmentPointPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			enroutesegmentpointpropertytypeIdent := enroutesegmentpointpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.EnRouteSegmentPointPropertyType{Name: %s}).Stage(stage)", enroutesegmentpointpropertytypeIdent, __gong__toRawStringLiteral(enroutesegmentpointpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", enroutesegmentpointpropertytypeIdent, __gong__toRawStringLiteral(enroutesegmentpointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", enroutesegmentpointpropertytypeIdent, __gong__toRawStringLiteral(enroutesegmentpointpropertytype.NilReason)))
			if enroutesegmentpointpropertytype.EnRouteSegmentPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := enroutesegmentpointpropertytype.EnRouteSegmentPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnRouteSegmentPoint = %s", enroutesegmentpointpropertytypeIdent, targetIdent))
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
		for _, flightcharacteristicpropertytype := range __gong__sortStageSetInstances(stage.FlightCharacteristicPropertyTypes, stage.FlightCharacteristicPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightcharacteristicpropertytypeIdent := flightcharacteristicpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightCharacteristicPropertyType{Name: %s}).Stage(stage)", flightcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(flightcharacteristicpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(flightcharacteristicpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(flightcharacteristicpropertytype.NilReason)))
			if flightcharacteristicpropertytype.FlightCharacteristic != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightcharacteristicpropertytype.FlightCharacteristic.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCharacteristic = %s", flightcharacteristicpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightcharacteristictype := range __gong__sortStageSetInstances(stage.FlightCharacteristicTypes, stage.FlightCharacteristicType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightcharacteristictypeIdent := flightcharacteristictype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightCharacteristicType{Name: %s}).Stage(stage)", flightcharacteristictypeIdent, __gong__toRawStringLiteral(flightcharacteristictype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightcharacteristictypeIdent, __gong__toRawStringLiteral(flightcharacteristictype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightcharacteristictypeIdent, __gong__toRawStringLiteral(flightcharacteristictype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightcharacteristictypeIdent, __gong__toRawStringLiteral(flightcharacteristictype.Description)))
			for _, elem := range flightcharacteristictype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightcharacteristictypeIdent, flightcharacteristictypeIdent, targetIdent))
			}
			for _, elem := range flightcharacteristictype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightcharacteristictypeIdent, flightcharacteristictypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightconditioncircumstancepropertytype := range __gong__sortStageSetInstances(stage.FlightConditionCircumstancePropertyTypes, stage.FlightConditionCircumstancePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightconditioncircumstancepropertytypeIdent := flightconditioncircumstancepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionCircumstancePropertyType{Name: %s}).Stage(stage)", flightconditioncircumstancepropertytypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditioncircumstancepropertytypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightconditioncircumstancepropertytypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancepropertytype.NilReason)))
			if flightconditioncircumstancepropertytype.FlightConditionCircumstance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditioncircumstancepropertytype.FlightConditionCircumstance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightConditionCircumstance = %s", flightconditioncircumstancepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightconditioncircumstancetype := range __gong__sortStageSetInstances(stage.FlightConditionCircumstanceTypes, stage.FlightConditionCircumstanceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightconditioncircumstancetypeIdent := flightconditioncircumstancetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionCircumstanceType{Name: %s}).Stage(stage)", flightconditioncircumstancetypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditioncircumstancetypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightconditioncircumstancetypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightconditioncircumstancetypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancetype.Description)))
			for _, elem := range flightconditioncircumstancetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightconditioncircumstancetypeIdent, flightconditioncircumstancetypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncircumstancetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightconditioncircumstancetypeIdent, flightconditioncircumstancetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightconditioncombinationpropertytype := range __gong__sortStageSetInstances(stage.FlightConditionCombinationPropertyTypes, stage.FlightConditionCombinationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightconditioncombinationpropertytypeIdent := flightconditioncombinationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionCombinationPropertyType{Name: %s}).Stage(stage)", flightconditioncombinationpropertytypeIdent, __gong__toRawStringLiteral(flightconditioncombinationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditioncombinationpropertytypeIdent, __gong__toRawStringLiteral(flightconditioncombinationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightconditioncombinationpropertytypeIdent, __gong__toRawStringLiteral(flightconditioncombinationpropertytype.NilReason)))
			if flightconditioncombinationpropertytype.FlightConditionCombination != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditioncombinationpropertytype.FlightConditionCombination.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightConditionCombination = %s", flightconditioncombinationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightconditioncombinationtype := range __gong__sortStageSetInstances(stage.FlightConditionCombinationTypes, stage.FlightConditionCombinationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightconditioncombinationtypeIdent := flightconditioncombinationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionCombinationType{Name: %s}).Stage(stage)", flightconditioncombinationtypeIdent, __gong__toRawStringLiteral(flightconditioncombinationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditioncombinationtypeIdent, __gong__toRawStringLiteral(flightconditioncombinationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightconditioncombinationtypeIdent, __gong__toRawStringLiteral(flightconditioncombinationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightconditioncombinationtypeIdent, __gong__toRawStringLiteral(flightconditioncombinationtype.Description)))
			for _, elem := range flightconditioncombinationtype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncombinationtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncombinationtype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncombinationtype.Element {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element = append(%s.Element, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncombinationtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightconditionelementpropertytype := range __gong__sortStageSetInstances(stage.FlightConditionElementPropertyTypes, stage.FlightConditionElementPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightconditionelementpropertytypeIdent := flightconditionelementpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionElementPropertyType{Name: %s}).Stage(stage)", flightconditionelementpropertytypeIdent, __gong__toRawStringLiteral(flightconditionelementpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditionelementpropertytypeIdent, __gong__toRawStringLiteral(flightconditionelementpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightconditionelementpropertytypeIdent, __gong__toRawStringLiteral(flightconditionelementpropertytype.NilReason)))
			if flightconditionelementpropertytype.FlightConditionElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementpropertytype.FlightConditionElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightConditionElement = %s", flightconditionelementpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightconditionelementtype := range __gong__sortStageSetInstances(stage.FlightConditionElementTypes, stage.FlightConditionElementType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightconditionelementtypeIdent := flightconditionelementtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionElementType{Name: %s}).Stage(stage)", flightconditionelementtypeIdent, __gong__toRawStringLiteral(flightconditionelementtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditionelementtypeIdent, __gong__toRawStringLiteral(flightconditionelementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightconditionelementtypeIdent, __gong__toRawStringLiteral(flightconditionelementtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightconditionelementtypeIdent, __gong__toRawStringLiteral(flightconditionelementtype.Description)))
			if flightconditionelementtype.FlightCondition_weather != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_weather.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_weather = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_aircraft != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_aircraft.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_aircraft = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_flight != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_flight.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_flight = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_directFlightCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_directFlightCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_directFlightCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_borderCrossingCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_borderCrossingCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_borderCrossingCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_operand != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_operand.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_operand = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_routePortionCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_routePortionCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_routePortionCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_standardInstrumentDepartureCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_standardInstrumentDepartureCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_standardInstrumentDepartureCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_standardInstrumentArrivalCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_standardInstrumentArrivalCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_standardInstrumentArrivalCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_organisationCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_organisationCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_organisationCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_fixDesignatedPoint = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_navaidSystem = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_position = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_runwayPoint = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_aimingPoint = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_airportReferencePoint = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_airspaceCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_airspaceCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_airspaceCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_airportHeliportCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_airportHeliportCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_airportHeliportCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_aerialRefuellingCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_aerialRefuellingCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_aerialRefuellingCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.OperationalCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightconditionelementtype.OperationalCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OperationalCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			for _, elem := range flightconditionelementtype.FlightLevel {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightLevel = append(%s.FlightLevel, %s)", flightconditionelementtypeIdent, flightconditionelementtypeIdent, targetIdent))
			}
			for _, elem := range flightconditionelementtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightconditionelementtypeIdent, flightconditionelementtypeIdent, targetIdent))
			}
			for _, elem := range flightconditionelementtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightconditionelementtypeIdent, flightconditionelementtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightrestrictionlevelpropertytype := range __gong__sortStageSetInstances(stage.FlightRestrictionLevelPropertyTypes, stage.FlightRestrictionLevelPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightrestrictionlevelpropertytypeIdent := flightrestrictionlevelpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionLevelPropertyType{Name: %s}).Stage(stage)", flightrestrictionlevelpropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionlevelpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictionlevelpropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionlevelpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightrestrictionlevelpropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionlevelpropertytype.NilReason)))
			if flightrestrictionlevelpropertytype.FlightRestrictionLevel != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightrestrictionlevelpropertytype.FlightRestrictionLevel.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRestrictionLevel = %s", flightrestrictionlevelpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightrestrictionleveltype := range __gong__sortStageSetInstances(stage.FlightRestrictionLevelTypes, stage.FlightRestrictionLevelType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightrestrictionleveltypeIdent := flightrestrictionleveltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionLevelType{Name: %s}).Stage(stage)", flightrestrictionleveltypeIdent, __gong__toRawStringLiteral(flightrestrictionleveltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictionleveltypeIdent, __gong__toRawStringLiteral(flightrestrictionleveltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightrestrictionleveltypeIdent, __gong__toRawStringLiteral(flightrestrictionleveltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightrestrictionleveltypeIdent, __gong__toRawStringLiteral(flightrestrictionleveltype.Description)))
			for _, elem := range flightrestrictionleveltype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightrestrictionleveltypeIdent, flightrestrictionleveltypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionleveltype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightrestrictionleveltypeIdent, flightrestrictionleveltypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightrestrictionroutepropertytype := range __gong__sortStageSetInstances(stage.FlightRestrictionRoutePropertyTypes, stage.FlightRestrictionRoutePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightrestrictionroutepropertytypeIdent := flightrestrictionroutepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionRoutePropertyType{Name: %s}).Stage(stage)", flightrestrictionroutepropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionroutepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictionroutepropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionroutepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightrestrictionroutepropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionroutepropertytype.NilReason)))
			if flightrestrictionroutepropertytype.FlightRestrictionRoute != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightrestrictionroutepropertytype.FlightRestrictionRoute.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRestrictionRoute = %s", flightrestrictionroutepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightrestrictionroutetype := range __gong__sortStageSetInstances(stage.FlightRestrictionRouteTypes, stage.FlightRestrictionRouteType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightrestrictionroutetypeIdent := flightrestrictionroutetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionRouteType{Name: %s}).Stage(stage)", flightrestrictionroutetypeIdent, __gong__toRawStringLiteral(flightrestrictionroutetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictionroutetypeIdent, __gong__toRawStringLiteral(flightrestrictionroutetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightrestrictionroutetypeIdent, __gong__toRawStringLiteral(flightrestrictionroutetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightrestrictionroutetypeIdent, __gong__toRawStringLiteral(flightrestrictionroutetype.Description)))
			for _, elem := range flightrestrictionroutetype.RouteElement {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteElement = append(%s.RouteElement, %s)", flightrestrictionroutetypeIdent, flightrestrictionroutetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionroutetype.Contact {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", flightrestrictionroutetypeIdent, flightrestrictionroutetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionroutetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightrestrictionroutetypeIdent, flightrestrictionroutetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionroutetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightrestrictionroutetypeIdent, flightrestrictionroutetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightrestrictiontimeslicepropertytype := range __gong__sortStageSetInstances(stage.FlightRestrictionTimeSlicePropertyTypes, stage.FlightRestrictionTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightrestrictiontimeslicepropertytypeIdent := flightrestrictiontimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionTimeSlicePropertyType{Name: %s}).Stage(stage)", flightrestrictiontimeslicepropertytypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictiontimeslicepropertytypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", flightrestrictiontimeslicepropertytypeIdent, flightrestrictiontimeslicepropertytype.Owns))
			if flightrestrictiontimeslicepropertytype.FlightRestrictionTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightrestrictiontimeslicepropertytype.FlightRestrictionTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRestrictionTimeSlice = %s", flightrestrictiontimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightrestrictiontimeslicetype := range __gong__sortStageSetInstances(stage.FlightRestrictionTimeSliceTypes, stage.FlightRestrictionTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightrestrictiontimeslicetypeIdent := flightrestrictiontimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionTimeSliceType{Name: %s}).Stage(stage)", flightrestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightrestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightrestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", flightrestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetype.CorrectionNumber))
			if flightrestrictiontimeslicetype.Flight != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightrestrictiontimeslicetype.Flight.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Flight = %s", flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictiontimeslicetype.RegulatedRoute {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RegulatedRoute = append(%s.RegulatedRoute, %s)", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictiontimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictiontimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightrestrictiontype := range __gong__sortStageSetInstances(stage.FlightRestrictionTypes, stage.FlightRestrictionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightrestrictiontypeIdent := flightrestrictiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionType{Name: %s}).Stage(stage)", flightrestrictiontypeIdent, __gong__toRawStringLiteral(flightrestrictiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictiontypeIdent, __gong__toRawStringLiteral(flightrestrictiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightrestrictiontypeIdent, __gong__toRawStringLiteral(flightrestrictiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightrestrictiontypeIdent, __gong__toRawStringLiteral(flightrestrictiontype.Description)))
			for _, elem := range flightrestrictiontype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", flightrestrictiontypeIdent, flightrestrictiontypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightroutingelementpropertytype := range __gong__sortStageSetInstances(stage.FlightRoutingElementPropertyTypes, stage.FlightRoutingElementPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightroutingelementpropertytypeIdent := flightroutingelementpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRoutingElementPropertyType{Name: %s}).Stage(stage)", flightroutingelementpropertytypeIdent, __gong__toRawStringLiteral(flightroutingelementpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightroutingelementpropertytypeIdent, __gong__toRawStringLiteral(flightroutingelementpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightroutingelementpropertytypeIdent, __gong__toRawStringLiteral(flightroutingelementpropertytype.NilReason)))
			if flightroutingelementpropertytype.FlightRoutingElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementpropertytype.FlightRoutingElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRoutingElement = %s", flightroutingelementpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, flightroutingelementtype := range __gong__sortStageSetInstances(stage.FlightRoutingElementTypes, stage.FlightRoutingElementType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			flightroutingelementtypeIdent := flightroutingelementtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRoutingElementType{Name: %s}).Stage(stage)", flightroutingelementtypeIdent, __gong__toRawStringLiteral(flightroutingelementtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightroutingelementtypeIdent, __gong__toRawStringLiteral(flightroutingelementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightroutingelementtypeIdent, __gong__toRawStringLiteral(flightroutingelementtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightroutingelementtypeIdent, __gong__toRawStringLiteral(flightroutingelementtype.Description)))
			for _, elem := range flightroutingelementtype.FlightLevel {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightLevel = append(%s.FlightLevel, %s)", flightroutingelementtypeIdent, flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_directFlightElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.Element_directFlightElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_directFlightElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_routePortionElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.Element_routePortionElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_routePortionElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_standardInstrumentDepartureElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.Element_standardInstrumentDepartureElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_standardInstrumentDepartureElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_standardInstrumentArrivalElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.Element_standardInstrumentArrivalElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_standardInstrumentArrivalElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.PointElement_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_fixDesignatedPoint = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.PointElement_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_navaidSystem = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.PointElement_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_position = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.PointElement_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_runwayPoint = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.PointElement_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_aimingPoint = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.PointElement_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_airportReferencePoint = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_airspaceElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.Element_airspaceElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_airspaceElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_airportHeliportElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.Element_airportHeliportElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_airportHeliportElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_aerialRefuellingElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := flightroutingelementtype.Element_aerialRefuellingElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_aerialRefuellingElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			for _, elem := range flightroutingelementtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightroutingelementtypeIdent, flightroutingelementtypeIdent, targetIdent))
			}
			for _, elem := range flightroutingelementtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightroutingelementtypeIdent, flightroutingelementtypeIdent, targetIdent))
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
		for _, meteorologypropertytype := range __gong__sortStageSetInstances(stage.MeteorologyPropertyTypes, stage.MeteorologyPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			meteorologypropertytypeIdent := meteorologypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MeteorologyPropertyType{Name: %s}).Stage(stage)", meteorologypropertytypeIdent, __gong__toRawStringLiteral(meteorologypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meteorologypropertytypeIdent, __gong__toRawStringLiteral(meteorologypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", meteorologypropertytypeIdent, __gong__toRawStringLiteral(meteorologypropertytype.NilReason)))
			if meteorologypropertytype.Meteorology != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := meteorologypropertytype.Meteorology.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Meteorology = %s", meteorologypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, meteorologytype := range __gong__sortStageSetInstances(stage.MeteorologyTypes, stage.MeteorologyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			meteorologytypeIdent := meteorologytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MeteorologyType{Name: %s}).Stage(stage)", meteorologytypeIdent, __gong__toRawStringLiteral(meteorologytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meteorologytypeIdent, __gong__toRawStringLiteral(meteorologytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", meteorologytypeIdent, __gong__toRawStringLiteral(meteorologytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", meteorologytypeIdent, __gong__toRawStringLiteral(meteorologytype.Description)))
			for _, elem := range meteorologytype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", meteorologytypeIdent, meteorologytypeIdent, targetIdent))
			}
			for _, elem := range meteorologytype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", meteorologytypeIdent, meteorologytypeIdent, targetIdent))
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
		for _, routeavailabilitypropertytype := range __gong__sortStageSetInstances(stage.RouteAvailabilityPropertyTypes, stage.RouteAvailabilityPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routeavailabilitypropertytypeIdent := routeavailabilitypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteAvailabilityPropertyType{Name: %s}).Stage(stage)", routeavailabilitypropertytypeIdent, __gong__toRawStringLiteral(routeavailabilitypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routeavailabilitypropertytypeIdent, __gong__toRawStringLiteral(routeavailabilitypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", routeavailabilitypropertytypeIdent, __gong__toRawStringLiteral(routeavailabilitypropertytype.NilReason)))
			if routeavailabilitypropertytype.RouteAvailability != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routeavailabilitypropertytype.RouteAvailability.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteAvailability = %s", routeavailabilitypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routeavailabilitytype := range __gong__sortStageSetInstances(stage.RouteAvailabilityTypes, stage.RouteAvailabilityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routeavailabilitytypeIdent := routeavailabilitytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteAvailabilityType{Name: %s}).Stage(stage)", routeavailabilitytypeIdent, __gong__toRawStringLiteral(routeavailabilitytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routeavailabilitytypeIdent, __gong__toRawStringLiteral(routeavailabilitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routeavailabilitytypeIdent, __gong__toRawStringLiteral(routeavailabilitytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routeavailabilitytypeIdent, __gong__toRawStringLiteral(routeavailabilitytype.Description)))
			for _, elem := range routeavailabilitytype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range routeavailabilitytype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range routeavailabilitytype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range routeavailabilitytype.Levels {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Levels = append(%s.Levels, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range routeavailabilitytype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routedmetimeslicepropertytype := range __gong__sortStageSetInstances(stage.RouteDMETimeSlicePropertyTypes, stage.RouteDMETimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routedmetimeslicepropertytypeIdent := routedmetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteDMETimeSlicePropertyType{Name: %s}).Stage(stage)", routedmetimeslicepropertytypeIdent, __gong__toRawStringLiteral(routedmetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routedmetimeslicepropertytypeIdent, __gong__toRawStringLiteral(routedmetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", routedmetimeslicepropertytypeIdent, routedmetimeslicepropertytype.Owns))
			if routedmetimeslicepropertytype.RouteDMETimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routedmetimeslicepropertytype.RouteDMETimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteDMETimeSlice = %s", routedmetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routedmetimeslicetype := range __gong__sortStageSetInstances(stage.RouteDMETimeSliceTypes, stage.RouteDMETimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routedmetimeslicetypeIdent := routedmetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteDMETimeSliceType{Name: %s}).Stage(stage)", routedmetimeslicetypeIdent, __gong__toRawStringLiteral(routedmetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routedmetimeslicetypeIdent, __gong__toRawStringLiteral(routedmetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routedmetimeslicetypeIdent, __gong__toRawStringLiteral(routedmetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routedmetimeslicetypeIdent, __gong__toRawStringLiteral(routedmetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", routedmetimeslicetypeIdent, __gong__toRawStringLiteral(routedmetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", routedmetimeslicetypeIdent, routedmetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", routedmetimeslicetypeIdent, routedmetimeslicetype.CorrectionNumber))
			if routedmetimeslicetype.ReferencedDME != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routedmetimeslicetype.ReferencedDME.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedDME = %s", routedmetimeslicetypeIdent, targetIdent))
			}
			if routedmetimeslicetype.ApplicableRoutePortion != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routedmetimeslicetype.ApplicableRoutePortion.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApplicableRoutePortion = %s", routedmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routedmetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", routedmetimeslicetypeIdent, routedmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routedmetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", routedmetimeslicetypeIdent, routedmetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routedmetype := range __gong__sortStageSetInstances(stage.RouteDMETypes, stage.RouteDMEType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routedmetypeIdent := routedmetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteDMEType{Name: %s}).Stage(stage)", routedmetypeIdent, __gong__toRawStringLiteral(routedmetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routedmetypeIdent, __gong__toRawStringLiteral(routedmetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routedmetypeIdent, __gong__toRawStringLiteral(routedmetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routedmetypeIdent, __gong__toRawStringLiteral(routedmetype.Description)))
			for _, elem := range routedmetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", routedmetypeIdent, routedmetypeIdent, targetIdent))
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
		for _, routesegmenttimeslicepropertytype := range __gong__sortStageSetInstances(stage.RouteSegmentTimeSlicePropertyTypes, stage.RouteSegmentTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routesegmenttimeslicepropertytypeIdent := routesegmenttimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteSegmentTimeSlicePropertyType{Name: %s}).Stage(stage)", routesegmenttimeslicepropertytypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routesegmenttimeslicepropertytypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", routesegmenttimeslicepropertytypeIdent, routesegmenttimeslicepropertytype.Owns))
			if routesegmenttimeslicepropertytype.RouteSegmentTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routesegmenttimeslicepropertytype.RouteSegmentTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteSegmentTimeSlice = %s", routesegmenttimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routesegmenttimeslicetype := range __gong__sortStageSetInstances(stage.RouteSegmentTimeSliceTypes, stage.RouteSegmentTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routesegmenttimeslicetypeIdent := routesegmenttimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteSegmentTimeSliceType{Name: %s}).Stage(stage)", routesegmenttimeslicetypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routesegmenttimeslicetypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routesegmenttimeslicetypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routesegmenttimeslicetypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", routesegmenttimeslicetypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", routesegmenttimeslicetypeIdent, routesegmenttimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", routesegmenttimeslicetypeIdent, routesegmenttimeslicetype.CorrectionNumber))
			if routesegmenttimeslicetype.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routesegmenttimeslicetype.Start.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.RouteFormed != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routesegmenttimeslicetype.RouteFormed.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteFormed = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.EvaluationArea != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routesegmenttimeslicetype.EvaluationArea.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EvaluationArea = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.CurveExtent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routesegmenttimeslicetype.CurveExtent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CurveExtent = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routesegmenttimeslicetype.End.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routesegmenttimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", routesegmenttimeslicetypeIdent, routesegmenttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routesegmenttimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", routesegmenttimeslicetypeIdent, routesegmenttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routesegmenttimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", routesegmenttimeslicetypeIdent, routesegmenttimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routesegmenttype := range __gong__sortStageSetInstances(stage.RouteSegmentTypes, stage.RouteSegmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routesegmenttypeIdent := routesegmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteSegmentType{Name: %s}).Stage(stage)", routesegmenttypeIdent, __gong__toRawStringLiteral(routesegmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routesegmenttypeIdent, __gong__toRawStringLiteral(routesegmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routesegmenttypeIdent, __gong__toRawStringLiteral(routesegmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routesegmenttypeIdent, __gong__toRawStringLiteral(routesegmenttype.Description)))
			for _, elem := range routesegmenttype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", routesegmenttypeIdent, routesegmenttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routetimeslicepropertytype := range __gong__sortStageSetInstances(stage.RouteTimeSlicePropertyTypes, stage.RouteTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routetimeslicepropertytypeIdent := routetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteTimeSlicePropertyType{Name: %s}).Stage(stage)", routetimeslicepropertytypeIdent, __gong__toRawStringLiteral(routetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routetimeslicepropertytypeIdent, __gong__toRawStringLiteral(routetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", routetimeslicepropertytypeIdent, routetimeslicepropertytype.Owns))
			if routetimeslicepropertytype.RouteTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routetimeslicepropertytype.RouteTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteTimeSlice = %s", routetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routetimeslicetype := range __gong__sortStageSetInstances(stage.RouteTimeSliceTypes, stage.RouteTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routetimeslicetypeIdent := routetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteTimeSliceType{Name: %s}).Stage(stage)", routetimeslicetypeIdent, __gong__toRawStringLiteral(routetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routetimeslicetypeIdent, __gong__toRawStringLiteral(routetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routetimeslicetypeIdent, __gong__toRawStringLiteral(routetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routetimeslicetypeIdent, __gong__toRawStringLiteral(routetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", routetimeslicetypeIdent, __gong__toRawStringLiteral(routetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", routetimeslicetypeIdent, routetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", routetimeslicetypeIdent, routetimeslicetype.CorrectionNumber))
			if routetimeslicetype.UserOrganisation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := routetimeslicetype.UserOrganisation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UserOrganisation = %s", routetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", routetimeslicetypeIdent, routetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", routetimeslicetypeIdent, routetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, routetype := range __gong__sortStageSetInstances(stage.RouteTypes, stage.RouteType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routetypeIdent := routetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteType{Name: %s}).Stage(stage)", routetypeIdent, __gong__toRawStringLiteral(routetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routetypeIdent, __gong__toRawStringLiteral(routetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routetypeIdent, __gong__toRawStringLiteral(routetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routetypeIdent, __gong__toRawStringLiteral(routetype.Description)))
			for _, elem := range routetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", routetypeIdent, routetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, rulesprocedurestimeslicepropertytype := range __gong__sortStageSetInstances(stage.RulesProceduresTimeSlicePropertyTypes, stage.RulesProceduresTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rulesprocedurestimeslicepropertytypeIdent := rulesprocedurestimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RulesProceduresTimeSlicePropertyType{Name: %s}).Stage(stage)", rulesprocedurestimeslicepropertytypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rulesprocedurestimeslicepropertytypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", rulesprocedurestimeslicepropertytypeIdent, rulesprocedurestimeslicepropertytype.Owns))
			if rulesprocedurestimeslicepropertytype.RulesProceduresTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := rulesprocedurestimeslicepropertytype.RulesProceduresTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RulesProceduresTimeSlice = %s", rulesprocedurestimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, rulesprocedurestimeslicetype := range __gong__sortStageSetInstances(stage.RulesProceduresTimeSliceTypes, stage.RulesProceduresTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rulesprocedurestimeslicetypeIdent := rulesprocedurestimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RulesProceduresTimeSliceType{Name: %s}).Stage(stage)", rulesprocedurestimeslicetypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rulesprocedurestimeslicetypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", rulesprocedurestimeslicetypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", rulesprocedurestimeslicetypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", rulesprocedurestimeslicetypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetype.CorrectionNumber))
			for _, elem := range rulesprocedurestimeslicetype.AffectedLocation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AffectedLocation = append(%s.AffectedLocation, %s)", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			for _, elem := range rulesprocedurestimeslicetype.AffectedArea {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AffectedArea = append(%s.AffectedArea, %s)", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			for _, elem := range rulesprocedurestimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			for _, elem := range rulesprocedurestimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, rulesprocedurestype := range __gong__sortStageSetInstances(stage.RulesProceduresTypes, stage.RulesProceduresType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rulesprocedurestypeIdent := rulesprocedurestype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RulesProceduresType{Name: %s}).Stage(stage)", rulesprocedurestypeIdent, __gong__toRawStringLiteral(rulesprocedurestype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rulesprocedurestypeIdent, __gong__toRawStringLiteral(rulesprocedurestype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", rulesprocedurestypeIdent, __gong__toRawStringLiteral(rulesprocedurestype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", rulesprocedurestypeIdent, __gong__toRawStringLiteral(rulesprocedurestype.Description)))
			for _, elem := range rulesprocedurestype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", rulesprocedurestypeIdent, rulesprocedurestypeIdent, targetIdent))
			}
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures\"", fmt.Sprintf("%q", modelsPackageName))
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

func (abstractdirectflighttype *AbstractDirectFlightType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractdirectflighttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractdirectflighttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractdirectflighttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractdirectflighttype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractDirectFlightType", fieldName)
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

func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aerialrefuellinganchorpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aerialrefuellinganchorpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", aerialrefuellinganchorpropertytype.NilReason)

	case "AerialRefuellingAnchor":
		if aerialrefuellinganchorpropertytype.AerialRefuellingAnchor != nil {
			res = __gong__marshallPointer(ident, "AerialRefuellingAnchor", aerialrefuellinganchorpropertytype.AerialRefuellingAnchor.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AerialRefuellingAnchor", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingAnchorPropertyType", fieldName)
	}
	return
}

func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aerialrefuellinganchortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aerialrefuellinganchortype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aerialrefuellinganchortype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aerialrefuellinganchortype.Description)

	case "Extent":
		if aerialrefuellinganchortype.Extent != nil {
			res = __gong__marshallPointer(ident, "Extent", aerialrefuellinganchortype.Extent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Extent", "nil")
		}
	case "VerticalExtent":
		var sb strings.Builder
		for _, _airspacelayerpropertytype := range aerialrefuellinganchortype.VerticalExtent {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "VerticalExtent", _airspacelayerpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Point":
		var sb strings.Builder
		for _, _aerialrefuellingpointpropertytype := range aerialrefuellinganchortype.Point {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Point", _aerialrefuellingpointpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range aerialrefuellinganchortype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_1 := range aerialrefuellinganchortype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_1.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingAnchorType", fieldName)
	}
	return
}

func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aerialrefuellingpointpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aerialrefuellingpointpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", aerialrefuellingpointpropertytype.NilReason)

	case "AerialRefuellingPoint":
		if aerialrefuellingpointpropertytype.AerialRefuellingPoint != nil {
			res = __gong__marshallPointer(ident, "AerialRefuellingPoint", aerialrefuellingpointpropertytype.AerialRefuellingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AerialRefuellingPoint", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingPointPropertyType", fieldName)
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
		for _, _a_extension_2 := range aerialrefuellingpointtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_2.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingPointType", fieldName)
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

func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aerialrefuellingtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aerialrefuellingtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", aerialrefuellingtimeslicepropertytype.Owns)

	case "AerialRefuellingTimeSlice":
		if aerialrefuellingtimeslicepropertytype.AerialRefuellingTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AerialRefuellingTimeSlice", aerialrefuellingtimeslicepropertytype.AerialRefuellingTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AerialRefuellingTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingTimeSlicePropertyType", fieldName)
	}
	return
}

func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aerialrefuellingtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aerialrefuellingtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aerialrefuellingtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aerialrefuellingtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", aerialrefuellingtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", aerialrefuellingtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", aerialrefuellingtimeslicetype.CorrectionNumber)

	case "Availability":
		var sb strings.Builder
		for _, _routeavailabilitypropertytype := range aerialrefuellingtimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _routeavailabilitypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ProtectingAirspace":
		var sb strings.Builder
		for _, _airspacepropertytype := range aerialrefuellingtimeslicetype.ProtectingAirspace {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ProtectingAirspace", _airspacepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Track":
		var sb strings.Builder
		for _, _aerialrefuellingtrackpropertytype := range aerialrefuellingtimeslicetype.Track {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Track", _aerialrefuellingtrackpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Anchor":
		var sb strings.Builder
		for _, _aerialrefuellinganchorpropertytype := range aerialrefuellingtimeslicetype.Anchor {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Anchor", _aerialrefuellinganchorpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "OppositeTrack":
		if aerialrefuellingtimeslicetype.OppositeTrack != nil {
			res = __gong__marshallPointer(ident, "OppositeTrack", aerialrefuellingtimeslicetype.OppositeTrack.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OppositeTrack", "nil")
		}
	case "ManagingOrganisation":
		var sb strings.Builder
		for _, _authorityforaerialrefuellingpropertytype := range aerialrefuellingtimeslicetype.ManagingOrganisation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ManagingOrganisation", _authorityforaerialrefuellingpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range aerialrefuellingtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension := range aerialrefuellingtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingTimeSliceType", fieldName)
	}
	return
}

func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aerialrefuellingtrackpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aerialrefuellingtrackpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", aerialrefuellingtrackpropertytype.NilReason)

	case "AerialRefuellingTrack":
		if aerialrefuellingtrackpropertytype.AerialRefuellingTrack != nil {
			res = __gong__marshallPointer(ident, "AerialRefuellingTrack", aerialrefuellingtrackpropertytype.AerialRefuellingTrack.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AerialRefuellingTrack", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingTrackPropertyType", fieldName)
	}
	return
}

func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aerialrefuellingtracktype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aerialrefuellingtracktype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aerialrefuellingtracktype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aerialrefuellingtracktype.Description)

	case "Extent":
		if aerialrefuellingtracktype.Extent != nil {
			res = __gong__marshallPointer(ident, "Extent", aerialrefuellingtracktype.Extent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Extent", "nil")
		}
	case "Point":
		var sb strings.Builder
		for _, _aerialrefuellingpointpropertytype := range aerialrefuellingtracktype.Point {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Point", _aerialrefuellingpointpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "VerticalExtent":
		var sb strings.Builder
		for _, _airspacelayerpropertytype := range aerialrefuellingtracktype.VerticalExtent {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "VerticalExtent", _airspacelayerpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range aerialrefuellingtracktype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_3 := range aerialrefuellingtracktype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_3.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingTrackType", fieldName)
	}
	return
}

func (aerialrefuellingtype *AerialRefuellingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aerialrefuellingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aerialrefuellingtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aerialrefuellingtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aerialrefuellingtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _aerialrefuellingtimeslicepropertytype := range aerialrefuellingtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _aerialrefuellingtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AerialRefuellingType", fieldName)
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
		for _, _a_extension_34 := range aircraftcharacteristictype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_34.GongGetIdentifier(stage)))
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

func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacebordercrossingpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacebordercrossingpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airspacebordercrossingpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airspacebordercrossingpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", airspacebordercrossingpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", airspacebordercrossingpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", airspacebordercrossingpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", airspacebordercrossingpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", airspacebordercrossingpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", airspacebordercrossingpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", airspacebordercrossingpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceBorderCrossingPropertyType", fieldName)
	}
	return
}

func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacebordercrossingtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacebordercrossingtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airspacebordercrossingtimeslicepropertytype.Owns)

	case "AirspaceBorderCrossingTimeSlice":
		if airspacebordercrossingtimeslicepropertytype.AirspaceBorderCrossingTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AirspaceBorderCrossingTimeSlice", airspacebordercrossingtimeslicepropertytype.AirspaceBorderCrossingTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirspaceBorderCrossingTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceBorderCrossingTimeSlicePropertyType", fieldName)
	}
	return
}

func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacebordercrossingtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacebordercrossingtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspacebordercrossingtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspacebordercrossingtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", airspacebordercrossingtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", airspacebordercrossingtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", airspacebordercrossingtimeslicetype.CorrectionNumber)

	case "ExitedAirspace":
		if airspacebordercrossingtimeslicetype.ExitedAirspace != nil {
			res = __gong__marshallPointer(ident, "ExitedAirspace", airspacebordercrossingtimeslicetype.ExitedAirspace.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ExitedAirspace", "nil")
		}
	case "EnteredAirspace":
		if airspacebordercrossingtimeslicetype.EnteredAirspace != nil {
			res = __gong__marshallPointer(ident, "EnteredAirspace", airspacebordercrossingtimeslicetype.EnteredAirspace.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EnteredAirspace", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airspacebordercrossingtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_19 := range airspacebordercrossingtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_19.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceBorderCrossingTimeSliceType", fieldName)
	}
	return
}

func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacebordercrossingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacebordercrossingtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspacebordercrossingtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspacebordercrossingtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _airspacebordercrossingtimeslicepropertytype := range airspacebordercrossingtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _airspacebordercrossingtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceBorderCrossingType", fieldName)
	}
	return
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacelayerpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacelayerpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airspacelayerpropertytype.NilReason)

	case "AirspaceLayer":
		if airspacelayerpropertytype.AirspaceLayer != nil {
			res = __gong__marshallPointer(ident, "AirspaceLayer", airspacelayerpropertytype.AirspaceLayer.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirspaceLayer", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceLayerPropertyType", fieldName)
	}
	return
}

func (airspacelayertype *AirspaceLayerType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacelayertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacelayertype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspacelayertype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspacelayertype.Description)

	case "DiscreteLevelSeries":
		if airspacelayertype.DiscreteLevelSeries != nil {
			res = __gong__marshallPointer(ident, "DiscreteLevelSeries", airspacelayertype.DiscreteLevelSeries.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DiscreteLevelSeries", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airspacelayertype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_36 := range airspacelayertype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_36.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceLayerType", fieldName)
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
		for _, _a_extension_39 := range altitudeadjustmenttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_39.GongGetIdentifier(stage)))
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

func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityforaerialrefuellingpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityforaerialrefuellingpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", authorityforaerialrefuellingpropertytype.NilReason)

	case "AuthorityForAerialRefuelling":
		if authorityforaerialrefuellingpropertytype.AuthorityForAerialRefuelling != nil {
			res = __gong__marshallPointer(ident, "AuthorityForAerialRefuelling", authorityforaerialrefuellingpropertytype.AuthorityForAerialRefuelling.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AuthorityForAerialRefuelling", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForAerialRefuellingPropertyType", fieldName)
	}
	return
}

func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityforaerialrefuellingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityforaerialrefuellingtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", authorityforaerialrefuellingtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", authorityforaerialrefuellingtype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range authorityforaerialrefuellingtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheOrganisationAuthority":
		if authorityforaerialrefuellingtype.TheOrganisationAuthority != nil {
			res = __gong__marshallPointer(ident, "TheOrganisationAuthority", authorityforaerialrefuellingtype.TheOrganisationAuthority.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheOrganisationAuthority", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_4 := range authorityforaerialrefuellingtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_4.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForAerialRefuellingType", fieldName)
	}
	return
}

func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := changeoverpointtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", changeoverpointtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", changeoverpointtimeslicepropertytype.Owns)

	case "ChangeOverPointTimeSlice":
		if changeoverpointtimeslicepropertytype.ChangeOverPointTimeSlice != nil {
			res = __gong__marshallPointer(ident, "ChangeOverPointTimeSlice", changeoverpointtimeslicepropertytype.ChangeOverPointTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ChangeOverPointTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ChangeOverPointTimeSlicePropertyType", fieldName)
	}
	return
}

func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := changeoverpointtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", changeoverpointtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", changeoverpointtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", changeoverpointtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", changeoverpointtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", changeoverpointtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", changeoverpointtimeslicetype.CorrectionNumber)

	case "Location_fixDesignatedPoint":
		if changeoverpointtimeslicetype.Location_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "Location_fixDesignatedPoint", changeoverpointtimeslicetype.Location_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_fixDesignatedPoint", "nil")
		}
	case "Location_navaidSystem":
		if changeoverpointtimeslicetype.Location_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "Location_navaidSystem", changeoverpointtimeslicetype.Location_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_navaidSystem", "nil")
		}
	case "Location_position":
		if changeoverpointtimeslicetype.Location_position != nil {
			res = __gong__marshallPointer(ident, "Location_position", changeoverpointtimeslicetype.Location_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_position", "nil")
		}
	case "Location_runwayPoint":
		if changeoverpointtimeslicetype.Location_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "Location_runwayPoint", changeoverpointtimeslicetype.Location_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_runwayPoint", "nil")
		}
	case "Location_aimingPoint":
		if changeoverpointtimeslicetype.Location_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "Location_aimingPoint", changeoverpointtimeslicetype.Location_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_aimingPoint", "nil")
		}
	case "Location_airportReferencePoint":
		if changeoverpointtimeslicetype.Location_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "Location_airportReferencePoint", changeoverpointtimeslicetype.Location_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_airportReferencePoint", "nil")
		}
	case "ApplicableRoutePortion":
		if changeoverpointtimeslicetype.ApplicableRoutePortion != nil {
			res = __gong__marshallPointer(ident, "ApplicableRoutePortion", changeoverpointtimeslicetype.ApplicableRoutePortion.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ApplicableRoutePortion", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range changeoverpointtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_13 := range changeoverpointtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_13.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ChangeOverPointTimeSliceType", fieldName)
	}
	return
}

func (changeoverpointtype *ChangeOverPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := changeoverpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", changeoverpointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", changeoverpointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", changeoverpointtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _changeoverpointtimeslicepropertytype := range changeoverpointtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _changeoverpointtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ChangeOverPointType", fieldName)
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
		for _, _a_extension_30 := range contactinformationtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_30.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ContactInformationType", fieldName)
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

func (directflightclasstype *DirectFlightClassType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directflightclasstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directflightclasstype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", directflightclasstype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", directflightclasstype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range directflightclasstype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_20 := range directflightclasstype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_20.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DirectFlightClassType", fieldName)
	}
	return
}

func (directflightpropertytype *DirectFlightPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directflightpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directflightpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", directflightpropertytype.NilReason)

	case "DirectFlightClass":
		if directflightpropertytype.DirectFlightClass != nil {
			res = __gong__marshallPointer(ident, "DirectFlightClass", directflightpropertytype.DirectFlightClass.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DirectFlightClass", "nil")
		}
	case "DirectFlightSegment":
		if directflightpropertytype.DirectFlightSegment != nil {
			res = __gong__marshallPointer(ident, "DirectFlightSegment", directflightpropertytype.DirectFlightSegment.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DirectFlightSegment", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DirectFlightPropertyType", fieldName)
	}
	return
}

func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directflightsegmentpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directflightsegmentpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", directflightsegmentpropertytype.NilReason)

	case "DirectFlightSegment":
		if directflightsegmentpropertytype.DirectFlightSegment != nil {
			res = __gong__marshallPointer(ident, "DirectFlightSegment", directflightsegmentpropertytype.DirectFlightSegment.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DirectFlightSegment", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DirectFlightSegmentPropertyType", fieldName)
	}
	return
}

func (directflightsegmenttype *DirectFlightSegmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directflightsegmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directflightsegmenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", directflightsegmenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", directflightsegmenttype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range directflightsegmenttype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "End_fixDesignatedPoint":
		if directflightsegmenttype.End_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "End_fixDesignatedPoint", directflightsegmenttype.End_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_fixDesignatedPoint", "nil")
		}
	case "End_navaidSystem":
		if directflightsegmenttype.End_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "End_navaidSystem", directflightsegmenttype.End_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_navaidSystem", "nil")
		}
	case "End_position":
		if directflightsegmenttype.End_position != nil {
			res = __gong__marshallPointer(ident, "End_position", directflightsegmenttype.End_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_position", "nil")
		}
	case "End_runwayPoint":
		if directflightsegmenttype.End_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "End_runwayPoint", directflightsegmenttype.End_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_runwayPoint", "nil")
		}
	case "End_aimingPoint":
		if directflightsegmenttype.End_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "End_aimingPoint", directflightsegmenttype.End_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_aimingPoint", "nil")
		}
	case "End_airportReferencePoint":
		if directflightsegmenttype.End_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "End_airportReferencePoint", directflightsegmenttype.End_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End_airportReferencePoint", "nil")
		}
	case "Start_fixDesignatedPoint":
		if directflightsegmenttype.Start_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "Start_fixDesignatedPoint", directflightsegmenttype.Start_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_fixDesignatedPoint", "nil")
		}
	case "Start_navaidSystem":
		if directflightsegmenttype.Start_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "Start_navaidSystem", directflightsegmenttype.Start_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_navaidSystem", "nil")
		}
	case "Start_position":
		if directflightsegmenttype.Start_position != nil {
			res = __gong__marshallPointer(ident, "Start_position", directflightsegmenttype.Start_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_position", "nil")
		}
	case "Start_runwayPoint":
		if directflightsegmenttype.Start_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "Start_runwayPoint", directflightsegmenttype.Start_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_runwayPoint", "nil")
		}
	case "Start_aimingPoint":
		if directflightsegmenttype.Start_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "Start_aimingPoint", directflightsegmenttype.Start_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_aimingPoint", "nil")
		}
	case "Start_airportReferencePoint":
		if directflightsegmenttype.Start_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "Start_airportReferencePoint", directflightsegmenttype.Start_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start_airportReferencePoint", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_21 := range directflightsegmenttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_21.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DirectFlightSegmentType", fieldName)
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
		for _, _a_extension_5 := range elevatedcurvetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_5.GongGetIdentifier(stage)))
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
		for _, _a_extension_6 := range elevatedpointtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_6.GongGetIdentifier(stage)))
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
		for _, _a_extension_7 := range elevatedsurfacetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_7.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedSurfaceType", fieldName)
	}
	return
}

func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := enroutesegmentpointpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", enroutesegmentpointpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", enroutesegmentpointpropertytype.NilReason)

	case "EnRouteSegmentPoint":
		if enroutesegmentpointpropertytype.EnRouteSegmentPoint != nil {
			res = __gong__marshallPointer(ident, "EnRouteSegmentPoint", enroutesegmentpointpropertytype.EnRouteSegmentPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EnRouteSegmentPoint", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct EnRouteSegmentPointPropertyType", fieldName)
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

func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightcharacteristicpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightcharacteristicpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", flightcharacteristicpropertytype.NilReason)

	case "FlightCharacteristic":
		if flightcharacteristicpropertytype.FlightCharacteristic != nil {
			res = __gong__marshallPointer(ident, "FlightCharacteristic", flightcharacteristicpropertytype.FlightCharacteristic.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCharacteristic", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightCharacteristicPropertyType", fieldName)
	}
	return
}

func (flightcharacteristictype *FlightCharacteristicType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightcharacteristictype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightcharacteristictype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", flightcharacteristictype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", flightcharacteristictype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range flightcharacteristictype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_35 := range flightcharacteristictype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_35.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightCharacteristicType", fieldName)
	}
	return
}

func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightconditioncircumstancepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightconditioncircumstancepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", flightconditioncircumstancepropertytype.NilReason)

	case "FlightConditionCircumstance":
		if flightconditioncircumstancepropertytype.FlightConditionCircumstance != nil {
			res = __gong__marshallPointer(ident, "FlightConditionCircumstance", flightconditioncircumstancepropertytype.FlightConditionCircumstance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightConditionCircumstance", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightConditionCircumstancePropertyType", fieldName)
	}
	return
}

func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightconditioncircumstancetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightconditioncircumstancetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", flightconditioncircumstancetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", flightconditioncircumstancetype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range flightconditioncircumstancetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_22 := range flightconditioncircumstancetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_22.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightConditionCircumstanceType", fieldName)
	}
	return
}

func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightconditioncombinationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightconditioncombinationpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", flightconditioncombinationpropertytype.NilReason)

	case "FlightConditionCombination":
		if flightconditioncombinationpropertytype.FlightConditionCombination != nil {
			res = __gong__marshallPointer(ident, "FlightConditionCombination", flightconditioncombinationpropertytype.FlightConditionCombination.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightConditionCombination", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightConditionCombinationPropertyType", fieldName)
	}
	return
}

func (flightconditioncombinationtype *FlightConditionCombinationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightconditioncombinationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightconditioncombinationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", flightconditioncombinationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", flightconditioncombinationtype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range flightconditioncombinationtype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range flightconditioncombinationtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range flightconditioncombinationtype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Element":
		var sb strings.Builder
		for _, _flightconditionelementpropertytype := range flightconditioncombinationtype.Element {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Element", _flightconditionelementpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_23 := range flightconditioncombinationtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_23.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightConditionCombinationType", fieldName)
	}
	return
}

func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightconditionelementpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightconditionelementpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", flightconditionelementpropertytype.NilReason)

	case "FlightConditionElement":
		if flightconditionelementpropertytype.FlightConditionElement != nil {
			res = __gong__marshallPointer(ident, "FlightConditionElement", flightconditionelementpropertytype.FlightConditionElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightConditionElement", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightConditionElementPropertyType", fieldName)
	}
	return
}

func (flightconditionelementtype *FlightConditionElementType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightconditionelementtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightconditionelementtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", flightconditionelementtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", flightconditionelementtype.Description)

	case "FlightCondition_weather":
		if flightconditionelementtype.FlightCondition_weather != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_weather", flightconditionelementtype.FlightCondition_weather.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_weather", "nil")
		}
	case "FlightCondition_aircraft":
		if flightconditionelementtype.FlightCondition_aircraft != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_aircraft", flightconditionelementtype.FlightCondition_aircraft.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_aircraft", "nil")
		}
	case "FlightCondition_flight":
		if flightconditionelementtype.FlightCondition_flight != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_flight", flightconditionelementtype.FlightCondition_flight.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_flight", "nil")
		}
	case "FlightCondition_directFlightCondition":
		if flightconditionelementtype.FlightCondition_directFlightCondition != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_directFlightCondition", flightconditionelementtype.FlightCondition_directFlightCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_directFlightCondition", "nil")
		}
	case "FlightCondition_borderCrossingCondition":
		if flightconditionelementtype.FlightCondition_borderCrossingCondition != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_borderCrossingCondition", flightconditionelementtype.FlightCondition_borderCrossingCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_borderCrossingCondition", "nil")
		}
	case "FlightCondition_operand":
		if flightconditionelementtype.FlightCondition_operand != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_operand", flightconditionelementtype.FlightCondition_operand.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_operand", "nil")
		}
	case "FlightCondition_routePortionCondition":
		if flightconditionelementtype.FlightCondition_routePortionCondition != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_routePortionCondition", flightconditionelementtype.FlightCondition_routePortionCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_routePortionCondition", "nil")
		}
	case "FlightCondition_standardInstrumentDepartureCondition":
		if flightconditionelementtype.FlightCondition_standardInstrumentDepartureCondition != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_standardInstrumentDepartureCondition", flightconditionelementtype.FlightCondition_standardInstrumentDepartureCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_standardInstrumentDepartureCondition", "nil")
		}
	case "FlightCondition_standardInstrumentArrivalCondition":
		if flightconditionelementtype.FlightCondition_standardInstrumentArrivalCondition != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_standardInstrumentArrivalCondition", flightconditionelementtype.FlightCondition_standardInstrumentArrivalCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_standardInstrumentArrivalCondition", "nil")
		}
	case "FlightCondition_organisationCondition":
		if flightconditionelementtype.FlightCondition_organisationCondition != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_organisationCondition", flightconditionelementtype.FlightCondition_organisationCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_organisationCondition", "nil")
		}
	case "SignificantPointCondition_fixDesignatedPoint":
		if flightconditionelementtype.SignificantPointCondition_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_fixDesignatedPoint", flightconditionelementtype.SignificantPointCondition_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_fixDesignatedPoint", "nil")
		}
	case "SignificantPointCondition_navaidSystem":
		if flightconditionelementtype.SignificantPointCondition_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_navaidSystem", flightconditionelementtype.SignificantPointCondition_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_navaidSystem", "nil")
		}
	case "SignificantPointCondition_position":
		if flightconditionelementtype.SignificantPointCondition_position != nil {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_position", flightconditionelementtype.SignificantPointCondition_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_position", "nil")
		}
	case "SignificantPointCondition_runwayPoint":
		if flightconditionelementtype.SignificantPointCondition_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_runwayPoint", flightconditionelementtype.SignificantPointCondition_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_runwayPoint", "nil")
		}
	case "SignificantPointCondition_aimingPoint":
		if flightconditionelementtype.SignificantPointCondition_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_aimingPoint", flightconditionelementtype.SignificantPointCondition_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_aimingPoint", "nil")
		}
	case "SignificantPointCondition_airportReferencePoint":
		if flightconditionelementtype.SignificantPointCondition_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_airportReferencePoint", flightconditionelementtype.SignificantPointCondition_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SignificantPointCondition_airportReferencePoint", "nil")
		}
	case "FlightCondition_airspaceCondition":
		if flightconditionelementtype.FlightCondition_airspaceCondition != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_airspaceCondition", flightconditionelementtype.FlightCondition_airspaceCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_airspaceCondition", "nil")
		}
	case "FlightCondition_airportHeliportCondition":
		if flightconditionelementtype.FlightCondition_airportHeliportCondition != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_airportHeliportCondition", flightconditionelementtype.FlightCondition_airportHeliportCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_airportHeliportCondition", "nil")
		}
	case "FlightCondition_aerialRefuellingCondition":
		if flightconditionelementtype.FlightCondition_aerialRefuellingCondition != nil {
			res = __gong__marshallPointer(ident, "FlightCondition_aerialRefuellingCondition", flightconditionelementtype.FlightCondition_aerialRefuellingCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightCondition_aerialRefuellingCondition", "nil")
		}
	case "OperationalCondition":
		if flightconditionelementtype.OperationalCondition != nil {
			res = __gong__marshallPointer(ident, "OperationalCondition", flightconditionelementtype.OperationalCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OperationalCondition", "nil")
		}
	case "FlightLevel":
		var sb strings.Builder
		for _, _flightrestrictionlevelpropertytype := range flightconditionelementtype.FlightLevel {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FlightLevel", _flightrestrictionlevelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range flightconditionelementtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_24 := range flightconditionelementtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_24.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightConditionElementType", fieldName)
	}
	return
}

func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightrestrictionlevelpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightrestrictionlevelpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", flightrestrictionlevelpropertytype.NilReason)

	case "FlightRestrictionLevel":
		if flightrestrictionlevelpropertytype.FlightRestrictionLevel != nil {
			res = __gong__marshallPointer(ident, "FlightRestrictionLevel", flightrestrictionlevelpropertytype.FlightRestrictionLevel.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightRestrictionLevel", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightRestrictionLevelPropertyType", fieldName)
	}
	return
}

func (flightrestrictionleveltype *FlightRestrictionLevelType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightrestrictionleveltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightrestrictionleveltype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", flightrestrictionleveltype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", flightrestrictionleveltype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range flightrestrictionleveltype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_26 := range flightrestrictionleveltype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_26.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightRestrictionLevelType", fieldName)
	}
	return
}

func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightrestrictionroutepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightrestrictionroutepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", flightrestrictionroutepropertytype.NilReason)

	case "FlightRestrictionRoute":
		if flightrestrictionroutepropertytype.FlightRestrictionRoute != nil {
			res = __gong__marshallPointer(ident, "FlightRestrictionRoute", flightrestrictionroutepropertytype.FlightRestrictionRoute.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightRestrictionRoute", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightRestrictionRoutePropertyType", fieldName)
	}
	return
}

func (flightrestrictionroutetype *FlightRestrictionRouteType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightrestrictionroutetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightrestrictionroutetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", flightrestrictionroutetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", flightrestrictionroutetype.Description)

	case "RouteElement":
		var sb strings.Builder
		for _, _flightroutingelementpropertytype := range flightrestrictionroutetype.RouteElement {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RouteElement", _flightroutingelementpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Contact":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range flightrestrictionroutetype.Contact {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Contact", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range flightrestrictionroutetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_27 := range flightrestrictionroutetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_27.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightRestrictionRouteType", fieldName)
	}
	return
}

func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightrestrictiontimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightrestrictiontimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", flightrestrictiontimeslicepropertytype.Owns)

	case "FlightRestrictionTimeSlice":
		if flightrestrictiontimeslicepropertytype.FlightRestrictionTimeSlice != nil {
			res = __gong__marshallPointer(ident, "FlightRestrictionTimeSlice", flightrestrictiontimeslicepropertytype.FlightRestrictionTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightRestrictionTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightRestrictionTimeSlicePropertyType", fieldName)
	}
	return
}

func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightrestrictiontimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightrestrictiontimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", flightrestrictiontimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", flightrestrictiontimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", flightrestrictiontimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", flightrestrictiontimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", flightrestrictiontimeslicetype.CorrectionNumber)

	case "Flight":
		if flightrestrictiontimeslicetype.Flight != nil {
			res = __gong__marshallPointer(ident, "Flight", flightrestrictiontimeslicetype.Flight.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Flight", "nil")
		}
	case "RegulatedRoute":
		var sb strings.Builder
		for _, _flightrestrictionroutepropertytype := range flightrestrictiontimeslicetype.RegulatedRoute {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RegulatedRoute", _flightrestrictionroutepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range flightrestrictiontimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_25 := range flightrestrictiontimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_25.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightRestrictionTimeSliceType", fieldName)
	}
	return
}

func (flightrestrictiontype *FlightRestrictionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightrestrictiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightrestrictiontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", flightrestrictiontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", flightrestrictiontype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _flightrestrictiontimeslicepropertytype := range flightrestrictiontype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _flightrestrictiontimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightRestrictionType", fieldName)
	}
	return
}

func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightroutingelementpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightroutingelementpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", flightroutingelementpropertytype.NilReason)

	case "FlightRoutingElement":
		if flightroutingelementpropertytype.FlightRoutingElement != nil {
			res = __gong__marshallPointer(ident, "FlightRoutingElement", flightroutingelementpropertytype.FlightRoutingElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FlightRoutingElement", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightRoutingElementPropertyType", fieldName)
	}
	return
}

func (flightroutingelementtype *FlightRoutingElementType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := flightroutingelementtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", flightroutingelementtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", flightroutingelementtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", flightroutingelementtype.Description)

	case "FlightLevel":
		var sb strings.Builder
		for _, _flightrestrictionlevelpropertytype := range flightroutingelementtype.FlightLevel {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FlightLevel", _flightrestrictionlevelpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Element_directFlightElement":
		if flightroutingelementtype.Element_directFlightElement != nil {
			res = __gong__marshallPointer(ident, "Element_directFlightElement", flightroutingelementtype.Element_directFlightElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Element_directFlightElement", "nil")
		}
	case "Element_routePortionElement":
		if flightroutingelementtype.Element_routePortionElement != nil {
			res = __gong__marshallPointer(ident, "Element_routePortionElement", flightroutingelementtype.Element_routePortionElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Element_routePortionElement", "nil")
		}
	case "Element_standardInstrumentDepartureElement":
		if flightroutingelementtype.Element_standardInstrumentDepartureElement != nil {
			res = __gong__marshallPointer(ident, "Element_standardInstrumentDepartureElement", flightroutingelementtype.Element_standardInstrumentDepartureElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Element_standardInstrumentDepartureElement", "nil")
		}
	case "Element_standardInstrumentArrivalElement":
		if flightroutingelementtype.Element_standardInstrumentArrivalElement != nil {
			res = __gong__marshallPointer(ident, "Element_standardInstrumentArrivalElement", flightroutingelementtype.Element_standardInstrumentArrivalElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Element_standardInstrumentArrivalElement", "nil")
		}
	case "PointElement_fixDesignatedPoint":
		if flightroutingelementtype.PointElement_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "PointElement_fixDesignatedPoint", flightroutingelementtype.PointElement_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointElement_fixDesignatedPoint", "nil")
		}
	case "PointElement_navaidSystem":
		if flightroutingelementtype.PointElement_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "PointElement_navaidSystem", flightroutingelementtype.PointElement_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointElement_navaidSystem", "nil")
		}
	case "PointElement_position":
		if flightroutingelementtype.PointElement_position != nil {
			res = __gong__marshallPointer(ident, "PointElement_position", flightroutingelementtype.PointElement_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointElement_position", "nil")
		}
	case "PointElement_runwayPoint":
		if flightroutingelementtype.PointElement_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "PointElement_runwayPoint", flightroutingelementtype.PointElement_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointElement_runwayPoint", "nil")
		}
	case "PointElement_aimingPoint":
		if flightroutingelementtype.PointElement_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "PointElement_aimingPoint", flightroutingelementtype.PointElement_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointElement_aimingPoint", "nil")
		}
	case "PointElement_airportReferencePoint":
		if flightroutingelementtype.PointElement_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "PointElement_airportReferencePoint", flightroutingelementtype.PointElement_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointElement_airportReferencePoint", "nil")
		}
	case "Element_airspaceElement":
		if flightroutingelementtype.Element_airspaceElement != nil {
			res = __gong__marshallPointer(ident, "Element_airspaceElement", flightroutingelementtype.Element_airspaceElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Element_airspaceElement", "nil")
		}
	case "Element_airportHeliportElement":
		if flightroutingelementtype.Element_airportHeliportElement != nil {
			res = __gong__marshallPointer(ident, "Element_airportHeliportElement", flightroutingelementtype.Element_airportHeliportElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Element_airportHeliportElement", "nil")
		}
	case "Element_aerialRefuellingElement":
		if flightroutingelementtype.Element_aerialRefuellingElement != nil {
			res = __gong__marshallPointer(ident, "Element_aerialRefuellingElement", flightroutingelementtype.Element_aerialRefuellingElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Element_aerialRefuellingElement", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range flightroutingelementtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_28 := range flightroutingelementtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_28.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FlightRoutingElementType", fieldName)
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
		for _, _a_extension_11 := range linguisticnotetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_11.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LinguisticNoteType", fieldName)
	}
	return
}

func (meteorologypropertytype *MeteorologyPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := meteorologypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", meteorologypropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", meteorologypropertytype.NilReason)

	case "Meteorology":
		if meteorologypropertytype.Meteorology != nil {
			res = __gong__marshallPointer(ident, "Meteorology", meteorologypropertytype.Meteorology.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Meteorology", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MeteorologyPropertyType", fieldName)
	}
	return
}

func (meteorologytype *MeteorologyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := meteorologytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", meteorologytype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", meteorologytype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", meteorologytype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range meteorologytype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_37 := range meteorologytype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_37.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MeteorologyType", fieldName)
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
		for _, _a_extension_12 := range notetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_12.GongGetIdentifier(stage)))
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
		for _, _a_extension_40 := range obstacleassessmentareatype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_40.GongGetIdentifier(stage)))
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
		for _, _a_extension_41 := range obstacleplacementtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_41.GongGetIdentifier(stage)))
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
		for _, _a_extension_42 := range obstructiontype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_42.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ObstructionType", fieldName)
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
		for _, _a_extension_31 := range onlinecontacttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_31.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct OnlineContactType", fieldName)
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
		for _, _a_extension_32 := range postaladdresstype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_32.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PostalAddressType", fieldName)
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

func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routeavailabilitypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routeavailabilitypropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", routeavailabilitypropertytype.NilReason)

	case "RouteAvailability":
		if routeavailabilitypropertytype.RouteAvailability != nil {
			res = __gong__marshallPointer(ident, "RouteAvailability", routeavailabilitypropertytype.RouteAvailability.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RouteAvailability", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteAvailabilityPropertyType", fieldName)
	}
	return
}

func (routeavailabilitytype *RouteAvailabilityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routeavailabilitytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routeavailabilitytype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", routeavailabilitytype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", routeavailabilitytype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range routeavailabilitytype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range routeavailabilitytype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range routeavailabilitytype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Levels":
		var sb strings.Builder
		for _, _airspacelayerpropertytype := range routeavailabilitytype.Levels {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Levels", _airspacelayerpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_15 := range routeavailabilitytype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_15.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteAvailabilityType", fieldName)
	}
	return
}

func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routedmetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routedmetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", routedmetimeslicepropertytype.Owns)

	case "RouteDMETimeSlice":
		if routedmetimeslicepropertytype.RouteDMETimeSlice != nil {
			res = __gong__marshallPointer(ident, "RouteDMETimeSlice", routedmetimeslicepropertytype.RouteDMETimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RouteDMETimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteDMETimeSlicePropertyType", fieldName)
	}
	return
}

func (routedmetimeslicetype *RouteDMETimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routedmetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routedmetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", routedmetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", routedmetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", routedmetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", routedmetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", routedmetimeslicetype.CorrectionNumber)

	case "ReferencedDME":
		if routedmetimeslicetype.ReferencedDME != nil {
			res = __gong__marshallPointer(ident, "ReferencedDME", routedmetimeslicetype.ReferencedDME.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ReferencedDME", "nil")
		}
	case "ApplicableRoutePortion":
		if routedmetimeslicetype.ApplicableRoutePortion != nil {
			res = __gong__marshallPointer(ident, "ApplicableRoutePortion", routedmetimeslicetype.ApplicableRoutePortion.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ApplicableRoutePortion", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range routedmetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_16 := range routedmetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_16.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteDMETimeSliceType", fieldName)
	}
	return
}

func (routedmetype *RouteDMEType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routedmetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routedmetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", routedmetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", routedmetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _routedmetimeslicepropertytype := range routedmetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _routedmetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteDMEType", fieldName)
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
		for _, _a_extension_17 := range routeportiontype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_17.GongGetIdentifier(stage)))
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

func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routesegmenttimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routesegmenttimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", routesegmenttimeslicepropertytype.Owns)

	case "RouteSegmentTimeSlice":
		if routesegmenttimeslicepropertytype.RouteSegmentTimeSlice != nil {
			res = __gong__marshallPointer(ident, "RouteSegmentTimeSlice", routesegmenttimeslicepropertytype.RouteSegmentTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RouteSegmentTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteSegmentTimeSlicePropertyType", fieldName)
	}
	return
}

func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routesegmenttimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routesegmenttimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", routesegmenttimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", routesegmenttimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", routesegmenttimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", routesegmenttimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", routesegmenttimeslicetype.CorrectionNumber)

	case "Start":
		if routesegmenttimeslicetype.Start != nil {
			res = __gong__marshallPointer(ident, "Start", routesegmenttimeslicetype.Start.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start", "nil")
		}
	case "RouteFormed":
		if routesegmenttimeslicetype.RouteFormed != nil {
			res = __gong__marshallPointer(ident, "RouteFormed", routesegmenttimeslicetype.RouteFormed.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RouteFormed", "nil")
		}
	case "EvaluationArea":
		if routesegmenttimeslicetype.EvaluationArea != nil {
			res = __gong__marshallPointer(ident, "EvaluationArea", routesegmenttimeslicetype.EvaluationArea.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EvaluationArea", "nil")
		}
	case "CurveExtent":
		if routesegmenttimeslicetype.CurveExtent != nil {
			res = __gong__marshallPointer(ident, "CurveExtent", routesegmenttimeslicetype.CurveExtent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CurveExtent", "nil")
		}
	case "End":
		if routesegmenttimeslicetype.End != nil {
			res = __gong__marshallPointer(ident, "End", routesegmenttimeslicetype.End.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End", "nil")
		}
	case "Availability":
		var sb strings.Builder
		for _, _routeavailabilitypropertytype := range routesegmenttimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _routeavailabilitypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range routesegmenttimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_18 := range routesegmenttimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_18.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteSegmentTimeSliceType", fieldName)
	}
	return
}

func (routesegmenttype *RouteSegmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routesegmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routesegmenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", routesegmenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", routesegmenttype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _routesegmenttimeslicepropertytype := range routesegmenttype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _routesegmenttimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteSegmentType", fieldName)
	}
	return
}

func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", routetimeslicepropertytype.Owns)

	case "RouteTimeSlice":
		if routetimeslicepropertytype.RouteTimeSlice != nil {
			res = __gong__marshallPointer(ident, "RouteTimeSlice", routetimeslicepropertytype.RouteTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RouteTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteTimeSlicePropertyType", fieldName)
	}
	return
}

func (routetimeslicetype *RouteTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", routetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", routetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", routetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", routetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", routetimeslicetype.CorrectionNumber)

	case "UserOrganisation":
		if routetimeslicetype.UserOrganisation != nil {
			res = __gong__marshallPointer(ident, "UserOrganisation", routetimeslicetype.UserOrganisation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "UserOrganisation", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range routetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_14 := range routetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_14.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteTimeSliceType", fieldName)
	}
	return
}

func (routetype *RouteType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", routetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", routetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _routetimeslicepropertytype := range routetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _routetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RouteType", fieldName)
	}
	return
}

func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rulesprocedurestimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rulesprocedurestimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", rulesprocedurestimeslicepropertytype.Owns)

	case "RulesProceduresTimeSlice":
		if rulesprocedurestimeslicepropertytype.RulesProceduresTimeSlice != nil {
			res = __gong__marshallPointer(ident, "RulesProceduresTimeSlice", rulesprocedurestimeslicepropertytype.RulesProceduresTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RulesProceduresTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RulesProceduresTimeSlicePropertyType", fieldName)
	}
	return
}

func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rulesprocedurestimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rulesprocedurestimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", rulesprocedurestimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", rulesprocedurestimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", rulesprocedurestimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", rulesprocedurestimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", rulesprocedurestimeslicetype.CorrectionNumber)

	case "AffectedLocation":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range rulesprocedurestimeslicetype.AffectedLocation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AffectedLocation", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AffectedArea":
		var sb strings.Builder
		for _, _airspacepropertytype := range rulesprocedurestimeslicetype.AffectedArea {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AffectedArea", _airspacepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range rulesprocedurestimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_29 := range rulesprocedurestimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_29.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RulesProceduresTimeSliceType", fieldName)
	}
	return
}

func (rulesprocedurestype *RulesProceduresType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rulesprocedurestype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rulesprocedurestype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", rulesprocedurestype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", rulesprocedurestype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _rulesprocedurestimeslicepropertytype := range rulesprocedurestype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _rulesprocedurestimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RulesProceduresType", fieldName)
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
		for _, _a_extension_33 := range telephonecontacttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_33.GongGetIdentifier(stage)))
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
		for _, _a_extension_38 := range timesheettype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_38.GongGetIdentifier(stage)))
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
func (abstractdirectflighttype *AbstractDirectFlightType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractdirectflighttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractdirectflighttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractdirectflighttype.GongMarshallField(stage, "Description"))
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
func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aerialrefuellinganchorpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aerialrefuellinganchorpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(aerialrefuellinganchorpropertytype.GongMarshallField(stage, "AerialRefuellingAnchor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aerialrefuellinganchortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aerialrefuellinganchortype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aerialrefuellinganchortype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(aerialrefuellinganchortype.GongMarshallField(stage, "Extent"))
		pointersInitializesStatements.WriteString(aerialrefuellinganchortype.GongMarshallField(stage, "VerticalExtent"))
		pointersInitializesStatements.WriteString(aerialrefuellinganchortype.GongMarshallField(stage, "Point"))
		pointersInitializesStatements.WriteString(aerialrefuellinganchortype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(aerialrefuellinganchortype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aerialrefuellingpointpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aerialrefuellingpointpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(aerialrefuellingpointpropertytype.GongMarshallField(stage, "AerialRefuellingPoint"))
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
func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aerialrefuellingtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aerialrefuellingtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(aerialrefuellingtimeslicepropertytype.GongMarshallField(stage, "AerialRefuellingTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "ProtectingAirspace"))
		pointersInitializesStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "Track"))
		pointersInitializesStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "Anchor"))
		pointersInitializesStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "OppositeTrack"))
		pointersInitializesStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "ManagingOrganisation"))
		pointersInitializesStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(aerialrefuellingtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aerialrefuellingtrackpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aerialrefuellingtrackpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(aerialrefuellingtrackpropertytype.GongMarshallField(stage, "AerialRefuellingTrack"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aerialrefuellingtracktype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aerialrefuellingtracktype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aerialrefuellingtracktype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(aerialrefuellingtracktype.GongMarshallField(stage, "Extent"))
		pointersInitializesStatements.WriteString(aerialrefuellingtracktype.GongMarshallField(stage, "Point"))
		pointersInitializesStatements.WriteString(aerialrefuellingtracktype.GongMarshallField(stage, "VerticalExtent"))
		pointersInitializesStatements.WriteString(aerialrefuellingtracktype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(aerialrefuellingtracktype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aerialrefuellingtype *AerialRefuellingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aerialrefuellingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aerialrefuellingtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aerialrefuellingtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(aerialrefuellingtype.GongMarshallField(stage, "TimeSlice"))
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
func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacebordercrossingpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacebordercrossingpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(airspacebordercrossingpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(airspacebordercrossingpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(airspacebordercrossingpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(airspacebordercrossingpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(airspacebordercrossingpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(airspacebordercrossingpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(airspacebordercrossingpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(airspacebordercrossingpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacebordercrossingtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacebordercrossingtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(airspacebordercrossingtimeslicepropertytype.GongMarshallField(stage, "AirspaceBorderCrossingTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacebordercrossingtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacebordercrossingtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspacebordercrossingtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(airspacebordercrossingtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(airspacebordercrossingtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(airspacebordercrossingtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(airspacebordercrossingtimeslicetype.GongMarshallField(stage, "ExitedAirspace"))
		pointersInitializesStatements.WriteString(airspacebordercrossingtimeslicetype.GongMarshallField(stage, "EnteredAirspace"))
		pointersInitializesStatements.WriteString(airspacebordercrossingtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airspacebordercrossingtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacebordercrossingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacebordercrossingtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspacebordercrossingtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airspacebordercrossingtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacelayerpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacelayerpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(airspacelayerpropertytype.GongMarshallField(stage, "AirspaceLayer"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacelayertype *AirspaceLayerType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacelayertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacelayertype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspacelayertype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airspacelayertype.GongMarshallField(stage, "DiscreteLevelSeries"))
		pointersInitializesStatements.WriteString(airspacelayertype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airspacelayertype.GongMarshallField(stage, "Extension"))
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
func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityforaerialrefuellingpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityforaerialrefuellingpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(authorityforaerialrefuellingpropertytype.GongMarshallField(stage, "AuthorityForAerialRefuelling"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityforaerialrefuellingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityforaerialrefuellingtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(authorityforaerialrefuellingtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(authorityforaerialrefuellingtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(authorityforaerialrefuellingtype.GongMarshallField(stage, "TheOrganisationAuthority"))
		pointersInitializesStatements.WriteString(authorityforaerialrefuellingtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(changeoverpointtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(changeoverpointtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(changeoverpointtimeslicepropertytype.GongMarshallField(stage, "ChangeOverPointTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Location_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Location_navaidSystem"))
		pointersInitializesStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Location_position"))
		pointersInitializesStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Location_runwayPoint"))
		pointersInitializesStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Location_aimingPoint"))
		pointersInitializesStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Location_airportReferencePoint"))
		pointersInitializesStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "ApplicableRoutePortion"))
		pointersInitializesStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(changeoverpointtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (changeoverpointtype *ChangeOverPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(changeoverpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(changeoverpointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(changeoverpointtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(changeoverpointtype.GongMarshallField(stage, "TimeSlice"))
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
func (directflightclasstype *DirectFlightClassType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directflightclasstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directflightclasstype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(directflightclasstype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(directflightclasstype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(directflightclasstype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (directflightpropertytype *DirectFlightPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directflightpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directflightpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(directflightpropertytype.GongMarshallField(stage, "DirectFlightClass"))
		pointersInitializesStatements.WriteString(directflightpropertytype.GongMarshallField(stage, "DirectFlightSegment"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directflightsegmentpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directflightsegmentpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(directflightsegmentpropertytype.GongMarshallField(stage, "DirectFlightSegment"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (directflightsegmenttype *DirectFlightSegmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "End_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "End_navaidSystem"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "End_position"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "End_runwayPoint"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "End_aimingPoint"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "End_airportReferencePoint"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Start_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Start_navaidSystem"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Start_position"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Start_runwayPoint"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Start_aimingPoint"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Start_airportReferencePoint"))
		pointersInitializesStatements.WriteString(directflightsegmenttype.GongMarshallField(stage, "Extension"))
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
func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(enroutesegmentpointpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(enroutesegmentpointpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(enroutesegmentpointpropertytype.GongMarshallField(stage, "EnRouteSegmentPoint"))
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
func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightcharacteristicpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightcharacteristicpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(flightcharacteristicpropertytype.GongMarshallField(stage, "FlightCharacteristic"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightcharacteristictype *FlightCharacteristicType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightcharacteristictype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightcharacteristictype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(flightcharacteristictype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(flightcharacteristictype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(flightcharacteristictype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightconditioncircumstancepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightconditioncircumstancepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(flightconditioncircumstancepropertytype.GongMarshallField(stage, "FlightConditionCircumstance"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightconditioncircumstancetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightconditioncircumstancetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(flightconditioncircumstancetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(flightconditioncircumstancetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(flightconditioncircumstancetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightconditioncombinationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightconditioncombinationpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(flightconditioncombinationpropertytype.GongMarshallField(stage, "FlightConditionCombination"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightconditioncombinationtype *FlightConditionCombinationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightconditioncombinationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightconditioncombinationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(flightconditioncombinationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(flightconditioncombinationtype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(flightconditioncombinationtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(flightconditioncombinationtype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(flightconditioncombinationtype.GongMarshallField(stage, "Element"))
		pointersInitializesStatements.WriteString(flightconditioncombinationtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightconditionelementpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightconditionelementpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(flightconditionelementpropertytype.GongMarshallField(stage, "FlightConditionElement"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightconditionelementtype *FlightConditionElementType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_weather"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_aircraft"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_flight"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_directFlightCondition"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_borderCrossingCondition"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_operand"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_routePortionCondition"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_standardInstrumentDepartureCondition"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_standardInstrumentArrivalCondition"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_organisationCondition"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "SignificantPointCondition_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "SignificantPointCondition_navaidSystem"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "SignificantPointCondition_position"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "SignificantPointCondition_runwayPoint"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "SignificantPointCondition_aimingPoint"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "SignificantPointCondition_airportReferencePoint"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_airspaceCondition"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_airportHeliportCondition"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightCondition_aerialRefuellingCondition"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "OperationalCondition"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "FlightLevel"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(flightconditionelementtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightrestrictionlevelpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightrestrictionlevelpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(flightrestrictionlevelpropertytype.GongMarshallField(stage, "FlightRestrictionLevel"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightrestrictionleveltype *FlightRestrictionLevelType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightrestrictionleveltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightrestrictionleveltype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(flightrestrictionleveltype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(flightrestrictionleveltype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(flightrestrictionleveltype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightrestrictionroutepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightrestrictionroutepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(flightrestrictionroutepropertytype.GongMarshallField(stage, "FlightRestrictionRoute"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightrestrictionroutetype *FlightRestrictionRouteType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightrestrictionroutetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightrestrictionroutetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(flightrestrictionroutetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(flightrestrictionroutetype.GongMarshallField(stage, "RouteElement"))
		pointersInitializesStatements.WriteString(flightrestrictionroutetype.GongMarshallField(stage, "Contact"))
		pointersInitializesStatements.WriteString(flightrestrictionroutetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(flightrestrictionroutetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightrestrictiontimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightrestrictiontimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(flightrestrictiontimeslicepropertytype.GongMarshallField(stage, "FlightRestrictionTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightrestrictiontimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightrestrictiontimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(flightrestrictiontimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(flightrestrictiontimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(flightrestrictiontimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(flightrestrictiontimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(flightrestrictiontimeslicetype.GongMarshallField(stage, "Flight"))
		pointersInitializesStatements.WriteString(flightrestrictiontimeslicetype.GongMarshallField(stage, "RegulatedRoute"))
		pointersInitializesStatements.WriteString(flightrestrictiontimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(flightrestrictiontimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightrestrictiontype *FlightRestrictionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightrestrictiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightrestrictiontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(flightrestrictiontype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(flightrestrictiontype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightroutingelementpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightroutingelementpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(flightroutingelementpropertytype.GongMarshallField(stage, "FlightRoutingElement"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (flightroutingelementtype *FlightRoutingElementType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "FlightLevel"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Element_directFlightElement"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Element_routePortionElement"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Element_standardInstrumentDepartureElement"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Element_standardInstrumentArrivalElement"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "PointElement_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "PointElement_navaidSystem"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "PointElement_position"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "PointElement_runwayPoint"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "PointElement_aimingPoint"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "PointElement_airportReferencePoint"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Element_airspaceElement"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Element_airportHeliportElement"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Element_aerialRefuellingElement"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(flightroutingelementtype.GongMarshallField(stage, "Extension"))
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
func (meteorologypropertytype *MeteorologyPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(meteorologypropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(meteorologypropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(meteorologypropertytype.GongMarshallField(stage, "Meteorology"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (meteorologytype *MeteorologyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(meteorologytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(meteorologytype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(meteorologytype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(meteorologytype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(meteorologytype.GongMarshallField(stage, "Extension"))
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
func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routeavailabilitypropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routeavailabilitypropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(routeavailabilitypropertytype.GongMarshallField(stage, "RouteAvailability"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routeavailabilitytype *RouteAvailabilityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routeavailabilitytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routeavailabilitytype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(routeavailabilitytype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(routeavailabilitytype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(routeavailabilitytype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(routeavailabilitytype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(routeavailabilitytype.GongMarshallField(stage, "Levels"))
		pointersInitializesStatements.WriteString(routeavailabilitytype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routedmetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routedmetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(routedmetimeslicepropertytype.GongMarshallField(stage, "RouteDMETimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routedmetimeslicetype *RouteDMETimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routedmetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routedmetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(routedmetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(routedmetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(routedmetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(routedmetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(routedmetimeslicetype.GongMarshallField(stage, "ReferencedDME"))
		pointersInitializesStatements.WriteString(routedmetimeslicetype.GongMarshallField(stage, "ApplicableRoutePortion"))
		pointersInitializesStatements.WriteString(routedmetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(routedmetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routedmetype *RouteDMEType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routedmetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routedmetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(routedmetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(routedmetype.GongMarshallField(stage, "TimeSlice"))
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
func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routesegmenttimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routesegmenttimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(routesegmenttimeslicepropertytype.GongMarshallField(stage, "RouteSegmentTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "Start"))
		pointersInitializesStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "RouteFormed"))
		pointersInitializesStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "EvaluationArea"))
		pointersInitializesStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "CurveExtent"))
		pointersInitializesStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "End"))
		pointersInitializesStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(routesegmenttimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routesegmenttype *RouteSegmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routesegmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routesegmenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(routesegmenttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(routesegmenttype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(routetimeslicepropertytype.GongMarshallField(stage, "RouteTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routetimeslicetype *RouteTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(routetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(routetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(routetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(routetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(routetimeslicetype.GongMarshallField(stage, "UserOrganisation"))
		pointersInitializesStatements.WriteString(routetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(routetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (routetype *RouteType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(routetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(routetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rulesprocedurestimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rulesprocedurestimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(rulesprocedurestimeslicepropertytype.GongMarshallField(stage, "RulesProceduresTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rulesprocedurestimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rulesprocedurestimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(rulesprocedurestimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(rulesprocedurestimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(rulesprocedurestimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(rulesprocedurestimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(rulesprocedurestimeslicetype.GongMarshallField(stage, "AffectedLocation"))
		pointersInitializesStatements.WriteString(rulesprocedurestimeslicetype.GongMarshallField(stage, "AffectedArea"))
		pointersInitializesStatements.WriteString(rulesprocedurestimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(rulesprocedurestimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rulesprocedurestype *RulesProceduresType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rulesprocedurestype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rulesprocedurestype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(rulesprocedurestype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(rulesprocedurestype.GongMarshallField(stage, "TimeSlice"))
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
