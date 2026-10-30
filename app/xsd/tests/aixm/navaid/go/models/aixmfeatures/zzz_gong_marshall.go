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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures"
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
		for _, abstractnavaidequipmenttype := range __gong__sortStageSetInstances(stage.AbstractNavaidEquipmentTypes, stage.AbstractNavaidEquipmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractnavaidequipmenttypeIdent := abstractnavaidequipmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractNavaidEquipmentType{Name: %s}).Stage(stage)", abstractnavaidequipmenttypeIdent, __gong__toRawStringLiteral(abstractnavaidequipmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractnavaidequipmenttypeIdent, __gong__toRawStringLiteral(abstractnavaidequipmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractnavaidequipmenttypeIdent, __gong__toRawStringLiteral(abstractnavaidequipmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractnavaidequipmenttypeIdent, __gong__toRawStringLiteral(abstractnavaidequipmenttype.Description)))
		}
	}
	if stage != nil {
		for _, abstractnavigationsystemcheckpointtype := range __gong__sortStageSetInstances(stage.AbstractNavigationSystemCheckpointTypes, stage.AbstractNavigationSystemCheckpointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractnavigationsystemcheckpointtypeIdent := abstractnavigationsystemcheckpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractNavigationSystemCheckpointType{Name: %s}).Stage(stage)", abstractnavigationsystemcheckpointtypeIdent, __gong__toRawStringLiteral(abstractnavigationsystemcheckpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractnavigationsystemcheckpointtypeIdent, __gong__toRawStringLiteral(abstractnavigationsystemcheckpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractnavigationsystemcheckpointtypeIdent, __gong__toRawStringLiteral(abstractnavigationsystemcheckpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractnavigationsystemcheckpointtypeIdent, __gong__toRawStringLiteral(abstractnavigationsystemcheckpointtype.Description)))
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
		for _, abstractradarequipmenttype := range __gong__sortStageSetInstances(stage.AbstractRadarEquipmentTypes, stage.AbstractRadarEquipmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractradarequipmenttypeIdent := abstractradarequipmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractRadarEquipmentType{Name: %s}).Stage(stage)", abstractradarequipmenttypeIdent, __gong__toRawStringLiteral(abstractradarequipmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractradarequipmenttypeIdent, __gong__toRawStringLiteral(abstractradarequipmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractradarequipmenttypeIdent, __gong__toRawStringLiteral(abstractradarequipmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractradarequipmenttypeIdent, __gong__toRawStringLiteral(abstractradarequipmenttype.Description)))
		}
	}
	if stage != nil {
		for _, abstractsurveillanceradartype := range __gong__sortStageSetInstances(stage.AbstractSurveillanceRadarTypes, stage.AbstractSurveillanceRadarType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractsurveillanceradartypeIdent := abstractsurveillanceradartype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractSurveillanceRadarType{Name: %s}).Stage(stage)", abstractsurveillanceradartypeIdent, __gong__toRawStringLiteral(abstractsurveillanceradartype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractsurveillanceradartypeIdent, __gong__toRawStringLiteral(abstractsurveillanceradartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractsurveillanceradartypeIdent, __gong__toRawStringLiteral(abstractsurveillanceradartype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractsurveillanceradartypeIdent, __gong__toRawStringLiteral(abstractsurveillanceradartype.Description)))
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
		for _, angleindicationtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AngleIndicationTimeSlicePropertyTypes, stage.AngleIndicationTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			angleindicationtimeslicepropertytypeIdent := angleindicationtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleIndicationTimeSlicePropertyType{Name: %s}).Stage(stage)", angleindicationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleindicationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", angleindicationtimeslicepropertytypeIdent, angleindicationtimeslicepropertytype.Owns))
			if angleindicationtimeslicepropertytype.AngleIndicationTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := angleindicationtimeslicepropertytype.AngleIndicationTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleIndicationTimeSlice = %s", angleindicationtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, angleindicationtimeslicetype := range __gong__sortStageSetInstances(stage.AngleIndicationTimeSliceTypes, stage.AngleIndicationTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			angleindicationtimeslicetypeIdent := angleindicationtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleIndicationTimeSliceType{Name: %s}).Stage(stage)", angleindicationtimeslicetypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleindicationtimeslicetypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", angleindicationtimeslicetypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", angleindicationtimeslicetypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", angleindicationtimeslicetypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", angleindicationtimeslicetypeIdent, angleindicationtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", angleindicationtimeslicetypeIdent, angleindicationtimeslicetype.CorrectionNumber))
			if angleindicationtimeslicetype.Fix != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := angleindicationtimeslicetype.Fix.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fix = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_fixDesignatedPoint = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_navaidSystem = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_position = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_runwayPoint = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_aimingPoint = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_airportReferencePoint = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range angleindicationtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", angleindicationtimeslicetypeIdent, angleindicationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range angleindicationtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", angleindicationtimeslicetypeIdent, angleindicationtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, angleindicationtype := range __gong__sortStageSetInstances(stage.AngleIndicationTypes, stage.AngleIndicationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			angleindicationtypeIdent := angleindicationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleIndicationType{Name: %s}).Stage(stage)", angleindicationtypeIdent, __gong__toRawStringLiteral(angleindicationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleindicationtypeIdent, __gong__toRawStringLiteral(angleindicationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", angleindicationtypeIdent, __gong__toRawStringLiteral(angleindicationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", angleindicationtypeIdent, __gong__toRawStringLiteral(angleindicationtype.Description)))
			for _, elem := range angleindicationtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", angleindicationtypeIdent, angleindicationtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, authorityfornavaidequipmentpropertytype := range __gong__sortStageSetInstances(stage.AuthorityForNavaidEquipmentPropertyTypes, stage.AuthorityForNavaidEquipmentPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityfornavaidequipmentpropertytypeIdent := authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForNavaidEquipmentPropertyType{Name: %s}).Stage(stage)", authorityfornavaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmentpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityfornavaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", authorityfornavaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmentpropertytype.NilReason)))
			if authorityfornavaidequipmentpropertytype.AuthorityForNavaidEquipment != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityfornavaidequipmentpropertytype.AuthorityForNavaidEquipment.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForNavaidEquipment = %s", authorityfornavaidequipmentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, authorityfornavaidequipmenttype := range __gong__sortStageSetInstances(stage.AuthorityForNavaidEquipmentTypes, stage.AuthorityForNavaidEquipmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityfornavaidequipmenttypeIdent := authorityfornavaidequipmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForNavaidEquipmentType{Name: %s}).Stage(stage)", authorityfornavaidequipmenttypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityfornavaidequipmenttypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityfornavaidequipmenttypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityfornavaidequipmenttypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmenttype.Description)))
			for _, elem := range authorityfornavaidequipmenttype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", authorityfornavaidequipmenttypeIdent, authorityfornavaidequipmenttypeIdent, targetIdent))
			}
			if authorityfornavaidequipmenttype.TheOrganisationAuthority != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityfornavaidequipmenttype.TheOrganisationAuthority.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheOrganisationAuthority = %s", authorityfornavaidequipmenttypeIdent, targetIdent))
			}
			for _, elem := range authorityfornavaidequipmenttype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", authorityfornavaidequipmenttypeIdent, authorityfornavaidequipmenttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, authorityforspecialnavigationstationpropertytype := range __gong__sortStageSetInstances(stage.AuthorityForSpecialNavigationStationPropertyTypes, stage.AuthorityForSpecialNavigationStationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityforspecialnavigationstationpropertytypeIdent := authorityforspecialnavigationstationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForSpecialNavigationStationPropertyType{Name: %s}).Stage(stage)", authorityforspecialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforspecialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", authorityforspecialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationpropertytype.NilReason)))
			if authorityforspecialnavigationstationpropertytype.AuthorityForSpecialNavigationStation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityforspecialnavigationstationpropertytype.AuthorityForSpecialNavigationStation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForSpecialNavigationStation = %s", authorityforspecialnavigationstationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, authorityforspecialnavigationstationtype := range __gong__sortStageSetInstances(stage.AuthorityForSpecialNavigationStationTypes, stage.AuthorityForSpecialNavigationStationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityforspecialnavigationstationtypeIdent := authorityforspecialnavigationstationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForSpecialNavigationStationType{Name: %s}).Stage(stage)", authorityforspecialnavigationstationtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforspecialnavigationstationtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityforspecialnavigationstationtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityforspecialnavigationstationtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationtype.Description)))
			for _, elem := range authorityforspecialnavigationstationtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", authorityforspecialnavigationstationtypeIdent, authorityforspecialnavigationstationtypeIdent, targetIdent))
			}
			if authorityforspecialnavigationstationtype.TheOrganisationAuthority != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityforspecialnavigationstationtype.TheOrganisationAuthority.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheOrganisationAuthority = %s", authorityforspecialnavigationstationtypeIdent, targetIdent))
			}
			for _, elem := range authorityforspecialnavigationstationtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", authorityforspecialnavigationstationtypeIdent, authorityforspecialnavigationstationtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, authorityforspecialnavigationsystempropertytype := range __gong__sortStageSetInstances(stage.AuthorityForSpecialNavigationSystemPropertyTypes, stage.AuthorityForSpecialNavigationSystemPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityforspecialnavigationsystempropertytypeIdent := authorityforspecialnavigationsystempropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForSpecialNavigationSystemPropertyType{Name: %s}).Stage(stage)", authorityforspecialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystempropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforspecialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystempropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", authorityforspecialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystempropertytype.NilReason)))
			if authorityforspecialnavigationsystempropertytype.AuthorityForSpecialNavigationSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityforspecialnavigationsystempropertytype.AuthorityForSpecialNavigationSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForSpecialNavigationSystem = %s", authorityforspecialnavigationsystempropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, authorityforspecialnavigationsystemtype := range __gong__sortStageSetInstances(stage.AuthorityForSpecialNavigationSystemTypes, stage.AuthorityForSpecialNavigationSystemType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityforspecialnavigationsystemtypeIdent := authorityforspecialnavigationsystemtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForSpecialNavigationSystemType{Name: %s}).Stage(stage)", authorityforspecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystemtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforspecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystemtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityforspecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystemtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityforspecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystemtype.Description)))
			for _, elem := range authorityforspecialnavigationsystemtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", authorityforspecialnavigationsystemtypeIdent, authorityforspecialnavigationsystemtypeIdent, targetIdent))
			}
			if authorityforspecialnavigationsystemtype.TheOrganisationAuthority != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityforspecialnavigationsystemtype.TheOrganisationAuthority.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheOrganisationAuthority = %s", authorityforspecialnavigationsystemtypeIdent, targetIdent))
			}
			for _, elem := range authorityforspecialnavigationsystemtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", authorityforspecialnavigationsystemtypeIdent, authorityforspecialnavigationsystemtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, azimuthtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AzimuthTimeSlicePropertyTypes, stage.AzimuthTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			azimuthtimeslicepropertytypeIdent := azimuthtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AzimuthTimeSlicePropertyType{Name: %s}).Stage(stage)", azimuthtimeslicepropertytypeIdent, __gong__toRawStringLiteral(azimuthtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", azimuthtimeslicepropertytypeIdent, __gong__toRawStringLiteral(azimuthtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", azimuthtimeslicepropertytypeIdent, azimuthtimeslicepropertytype.Owns))
			if azimuthtimeslicepropertytype.AzimuthTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := azimuthtimeslicepropertytype.AzimuthTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AzimuthTimeSlice = %s", azimuthtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, azimuthtimeslicetype := range __gong__sortStageSetInstances(stage.AzimuthTimeSliceTypes, stage.AzimuthTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			azimuthtimeslicetypeIdent := azimuthtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AzimuthTimeSliceType{Name: %s}).Stage(stage)", azimuthtimeslicetypeIdent, __gong__toRawStringLiteral(azimuthtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", azimuthtimeslicetypeIdent, __gong__toRawStringLiteral(azimuthtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", azimuthtimeslicetypeIdent, __gong__toRawStringLiteral(azimuthtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", azimuthtimeslicetypeIdent, __gong__toRawStringLiteral(azimuthtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", azimuthtimeslicetypeIdent, __gong__toRawStringLiteral(azimuthtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", azimuthtimeslicetypeIdent, azimuthtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", azimuthtimeslicetypeIdent, azimuthtimeslicetype.CorrectionNumber))
			if azimuthtimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := azimuthtimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, azimuthtype := range __gong__sortStageSetInstances(stage.AzimuthTypes, stage.AzimuthType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			azimuthtypeIdent := azimuthtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AzimuthType{Name: %s}).Stage(stage)", azimuthtypeIdent, __gong__toRawStringLiteral(azimuthtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", azimuthtypeIdent, __gong__toRawStringLiteral(azimuthtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", azimuthtypeIdent, __gong__toRawStringLiteral(azimuthtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", azimuthtypeIdent, __gong__toRawStringLiteral(azimuthtype.Description)))
			for _, elem := range azimuthtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", azimuthtypeIdent, azimuthtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, checkpointinstimeslicepropertytype := range __gong__sortStageSetInstances(stage.CheckpointINSTimeSlicePropertyTypes, stage.CheckpointINSTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			checkpointinstimeslicepropertytypeIdent := checkpointinstimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointINSTimeSlicePropertyType{Name: %s}).Stage(stage)", checkpointinstimeslicepropertytypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointinstimeslicepropertytypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", checkpointinstimeslicepropertytypeIdent, checkpointinstimeslicepropertytype.Owns))
			if checkpointinstimeslicepropertytype.CheckpointINSTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := checkpointinstimeslicepropertytype.CheckpointINSTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CheckpointINSTimeSlice = %s", checkpointinstimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, checkpointinstimeslicetype := range __gong__sortStageSetInstances(stage.CheckpointINSTimeSliceTypes, stage.CheckpointINSTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			checkpointinstimeslicetypeIdent := checkpointinstimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointINSTimeSliceType{Name: %s}).Stage(stage)", checkpointinstimeslicetypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointinstimeslicetypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", checkpointinstimeslicetypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", checkpointinstimeslicetypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", checkpointinstimeslicetypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", checkpointinstimeslicetypeIdent, checkpointinstimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", checkpointinstimeslicetypeIdent, checkpointinstimeslicetype.CorrectionNumber))
			if checkpointinstimeslicetype.Position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := checkpointinstimeslicetype.Position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Position = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.AirportHeliport != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := checkpointinstimeslicetype.AirportHeliport.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			for _, elem := range checkpointinstimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", checkpointinstimeslicetypeIdent, checkpointinstimeslicetypeIdent, targetIdent))
			}
			for _, elem := range checkpointinstimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", checkpointinstimeslicetypeIdent, checkpointinstimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, checkpointinstype := range __gong__sortStageSetInstances(stage.CheckpointINSTypes, stage.CheckpointINSType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			checkpointinstypeIdent := checkpointinstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointINSType{Name: %s}).Stage(stage)", checkpointinstypeIdent, __gong__toRawStringLiteral(checkpointinstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointinstypeIdent, __gong__toRawStringLiteral(checkpointinstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", checkpointinstypeIdent, __gong__toRawStringLiteral(checkpointinstype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", checkpointinstypeIdent, __gong__toRawStringLiteral(checkpointinstype.Description)))
			for _, elem := range checkpointinstype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", checkpointinstypeIdent, checkpointinstypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, checkpointvortimeslicepropertytype := range __gong__sortStageSetInstances(stage.CheckpointVORTimeSlicePropertyTypes, stage.CheckpointVORTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			checkpointvortimeslicepropertytypeIdent := checkpointvortimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointVORTimeSlicePropertyType{Name: %s}).Stage(stage)", checkpointvortimeslicepropertytypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointvortimeslicepropertytypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", checkpointvortimeslicepropertytypeIdent, checkpointvortimeslicepropertytype.Owns))
			if checkpointvortimeslicepropertytype.CheckpointVORTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := checkpointvortimeslicepropertytype.CheckpointVORTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CheckpointVORTimeSlice = %s", checkpointvortimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, checkpointvortimeslicetype := range __gong__sortStageSetInstances(stage.CheckpointVORTimeSliceTypes, stage.CheckpointVORTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			checkpointvortimeslicetypeIdent := checkpointvortimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointVORTimeSliceType{Name: %s}).Stage(stage)", checkpointvortimeslicetypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointvortimeslicetypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", checkpointvortimeslicetypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", checkpointvortimeslicetypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", checkpointvortimeslicetypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", checkpointvortimeslicetypeIdent, checkpointvortimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", checkpointvortimeslicetypeIdent, checkpointvortimeslicetype.CorrectionNumber))
			if checkpointvortimeslicetype.Position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := checkpointvortimeslicetype.Position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Position = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.AirportHeliport != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := checkpointvortimeslicetype.AirportHeliport.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range checkpointvortimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", checkpointvortimeslicetypeIdent, checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.CheckPointFacility != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := checkpointvortimeslicetype.CheckPointFacility.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CheckPointFacility = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range checkpointvortimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", checkpointvortimeslicetypeIdent, checkpointvortimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, checkpointvortype := range __gong__sortStageSetInstances(stage.CheckpointVORTypes, stage.CheckpointVORType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			checkpointvortypeIdent := checkpointvortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointVORType{Name: %s}).Stage(stage)", checkpointvortypeIdent, __gong__toRawStringLiteral(checkpointvortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointvortypeIdent, __gong__toRawStringLiteral(checkpointvortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", checkpointvortypeIdent, __gong__toRawStringLiteral(checkpointvortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", checkpointvortypeIdent, __gong__toRawStringLiteral(checkpointvortype.Description)))
			for _, elem := range checkpointvortype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", checkpointvortypeIdent, checkpointvortypeIdent, targetIdent))
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
		for _, dmetimeslicepropertytype := range __gong__sortStageSetInstances(stage.DMETimeSlicePropertyTypes, stage.DMETimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dmetimeslicepropertytypeIdent := dmetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DMETimeSlicePropertyType{Name: %s}).Stage(stage)", dmetimeslicepropertytypeIdent, __gong__toRawStringLiteral(dmetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dmetimeslicepropertytypeIdent, __gong__toRawStringLiteral(dmetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", dmetimeslicepropertytypeIdent, dmetimeslicepropertytype.Owns))
			if dmetimeslicepropertytype.DMETimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dmetimeslicepropertytype.DMETimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DMETimeSlice = %s", dmetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, dmetimeslicetype := range __gong__sortStageSetInstances(stage.DMETimeSliceTypes, stage.DMETimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dmetimeslicetypeIdent := dmetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DMETimeSliceType{Name: %s}).Stage(stage)", dmetimeslicetypeIdent, __gong__toRawStringLiteral(dmetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dmetimeslicetypeIdent, __gong__toRawStringLiteral(dmetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", dmetimeslicetypeIdent, __gong__toRawStringLiteral(dmetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dmetimeslicetypeIdent, __gong__toRawStringLiteral(dmetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", dmetimeslicetypeIdent, __gong__toRawStringLiteral(dmetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", dmetimeslicetypeIdent, dmetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", dmetimeslicetypeIdent, dmetimeslicetype.CorrectionNumber))
			if dmetimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dmetimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, dmetype := range __gong__sortStageSetInstances(stage.DMETypes, stage.DMEType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dmetypeIdent := dmetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DMEType{Name: %s}).Stage(stage)", dmetypeIdent, __gong__toRawStringLiteral(dmetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dmetypeIdent, __gong__toRawStringLiteral(dmetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", dmetypeIdent, __gong__toRawStringLiteral(dmetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dmetypeIdent, __gong__toRawStringLiteral(dmetype.Description)))
			for _, elem := range dmetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", dmetypeIdent, dmetypeIdent, targetIdent))
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
		for _, designatedpointtimeslicepropertytype := range __gong__sortStageSetInstances(stage.DesignatedPointTimeSlicePropertyTypes, stage.DesignatedPointTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			designatedpointtimeslicepropertytypeIdent := designatedpointtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DesignatedPointTimeSlicePropertyType{Name: %s}).Stage(stage)", designatedpointtimeslicepropertytypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", designatedpointtimeslicepropertytypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", designatedpointtimeslicepropertytypeIdent, designatedpointtimeslicepropertytype.Owns))
			if designatedpointtimeslicepropertytype.DesignatedPointTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := designatedpointtimeslicepropertytype.DesignatedPointTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatedPointTimeSlice = %s", designatedpointtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, designatedpointtimeslicetype := range __gong__sortStageSetInstances(stage.DesignatedPointTimeSliceTypes, stage.DesignatedPointTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			designatedpointtimeslicetypeIdent := designatedpointtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DesignatedPointTimeSliceType{Name: %s}).Stage(stage)", designatedpointtimeslicetypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", designatedpointtimeslicetypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", designatedpointtimeslicetypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", designatedpointtimeslicetypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", designatedpointtimeslicetypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", designatedpointtimeslicetypeIdent, designatedpointtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", designatedpointtimeslicetypeIdent, designatedpointtimeslicetype.CorrectionNumber))
			if designatedpointtimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := designatedpointtimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.AimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := designatedpointtimeslicetype.AimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AimingPoint = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.AirportHeliport != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := designatedpointtimeslicetype.AirportHeliport.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.RunwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := designatedpointtimeslicetype.RunwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RunwayPoint = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range designatedpointtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", designatedpointtimeslicetypeIdent, designatedpointtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range designatedpointtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", designatedpointtimeslicetypeIdent, designatedpointtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, designatedpointtype := range __gong__sortStageSetInstances(stage.DesignatedPointTypes, stage.DesignatedPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			designatedpointtypeIdent := designatedpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DesignatedPointType{Name: %s}).Stage(stage)", designatedpointtypeIdent, __gong__toRawStringLiteral(designatedpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", designatedpointtypeIdent, __gong__toRawStringLiteral(designatedpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", designatedpointtypeIdent, __gong__toRawStringLiteral(designatedpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", designatedpointtypeIdent, __gong__toRawStringLiteral(designatedpointtype.Description)))
			for _, elem := range designatedpointtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", designatedpointtypeIdent, designatedpointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, directionfindertimeslicepropertytype := range __gong__sortStageSetInstances(stage.DirectionFinderTimeSlicePropertyTypes, stage.DirectionFinderTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directionfindertimeslicepropertytypeIdent := directionfindertimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectionFinderTimeSlicePropertyType{Name: %s}).Stage(stage)", directionfindertimeslicepropertytypeIdent, __gong__toRawStringLiteral(directionfindertimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionfindertimeslicepropertytypeIdent, __gong__toRawStringLiteral(directionfindertimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", directionfindertimeslicepropertytypeIdent, directionfindertimeslicepropertytype.Owns))
			if directionfindertimeslicepropertytype.DirectionFinderTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directionfindertimeslicepropertytype.DirectionFinderTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectionFinderTimeSlice = %s", directionfindertimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, directionfindertimeslicetype := range __gong__sortStageSetInstances(stage.DirectionFinderTimeSliceTypes, stage.DirectionFinderTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directionfindertimeslicetypeIdent := directionfindertimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectionFinderTimeSliceType{Name: %s}).Stage(stage)", directionfindertimeslicetypeIdent, __gong__toRawStringLiteral(directionfindertimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionfindertimeslicetypeIdent, __gong__toRawStringLiteral(directionfindertimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directionfindertimeslicetypeIdent, __gong__toRawStringLiteral(directionfindertimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", directionfindertimeslicetypeIdent, __gong__toRawStringLiteral(directionfindertimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", directionfindertimeslicetypeIdent, __gong__toRawStringLiteral(directionfindertimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", directionfindertimeslicetypeIdent, directionfindertimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", directionfindertimeslicetypeIdent, directionfindertimeslicetype.CorrectionNumber))
			if directionfindertimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := directionfindertimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.InformationProvision {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InformationProvision = append(%s.InformationProvision, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, directionfindertype := range __gong__sortStageSetInstances(stage.DirectionFinderTypes, stage.DirectionFinderType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directionfindertypeIdent := directionfindertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectionFinderType{Name: %s}).Stage(stage)", directionfindertypeIdent, __gong__toRawStringLiteral(directionfindertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionfindertypeIdent, __gong__toRawStringLiteral(directionfindertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directionfindertypeIdent, __gong__toRawStringLiteral(directionfindertype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", directionfindertypeIdent, __gong__toRawStringLiteral(directionfindertype.Description)))
			for _, elem := range directionfindertype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", directionfindertypeIdent, directionfindertypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, distanceindicationtimeslicepropertytype := range __gong__sortStageSetInstances(stage.DistanceIndicationTimeSlicePropertyTypes, stage.DistanceIndicationTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			distanceindicationtimeslicepropertytypeIdent := distanceindicationtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DistanceIndicationTimeSlicePropertyType{Name: %s}).Stage(stage)", distanceindicationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceindicationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", distanceindicationtimeslicepropertytypeIdent, distanceindicationtimeslicepropertytype.Owns))
			if distanceindicationtimeslicepropertytype.DistanceIndicationTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := distanceindicationtimeslicepropertytype.DistanceIndicationTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DistanceIndicationTimeSlice = %s", distanceindicationtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, distanceindicationtimeslicetype := range __gong__sortStageSetInstances(stage.DistanceIndicationTimeSliceTypes, stage.DistanceIndicationTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			distanceindicationtimeslicetypeIdent := distanceindicationtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DistanceIndicationTimeSliceType{Name: %s}).Stage(stage)", distanceindicationtimeslicetypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceindicationtimeslicetypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", distanceindicationtimeslicetypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", distanceindicationtimeslicetypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", distanceindicationtimeslicetypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", distanceindicationtimeslicetypeIdent, distanceindicationtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", distanceindicationtimeslicetypeIdent, distanceindicationtimeslicetype.CorrectionNumber))
			if distanceindicationtimeslicetype.Fix != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := distanceindicationtimeslicetype.Fix.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fix = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_fixDesignatedPoint = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_navaidSystem = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_position = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_runwayPoint = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_aimingPoint = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_airportReferencePoint = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range distanceindicationtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", distanceindicationtimeslicetypeIdent, distanceindicationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range distanceindicationtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", distanceindicationtimeslicetypeIdent, distanceindicationtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, distanceindicationtype := range __gong__sortStageSetInstances(stage.DistanceIndicationTypes, stage.DistanceIndicationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			distanceindicationtypeIdent := distanceindicationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DistanceIndicationType{Name: %s}).Stage(stage)", distanceindicationtypeIdent, __gong__toRawStringLiteral(distanceindicationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceindicationtypeIdent, __gong__toRawStringLiteral(distanceindicationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", distanceindicationtypeIdent, __gong__toRawStringLiteral(distanceindicationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", distanceindicationtypeIdent, __gong__toRawStringLiteral(distanceindicationtype.Description)))
			for _, elem := range distanceindicationtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", distanceindicationtypeIdent, distanceindicationtypeIdent, targetIdent))
			}
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
		for _, elevationtimeslicepropertytype := range __gong__sortStageSetInstances(stage.ElevationTimeSlicePropertyTypes, stage.ElevationTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevationtimeslicepropertytypeIdent := elevationtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevationTimeSlicePropertyType{Name: %s}).Stage(stage)", elevationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(elevationtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(elevationtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", elevationtimeslicepropertytypeIdent, elevationtimeslicepropertytype.Owns))
			if elevationtimeslicepropertytype.ElevationTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elevationtimeslicepropertytype.ElevationTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevationTimeSlice = %s", elevationtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, elevationtimeslicetype := range __gong__sortStageSetInstances(stage.ElevationTimeSliceTypes, stage.ElevationTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevationtimeslicetypeIdent := elevationtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevationTimeSliceType{Name: %s}).Stage(stage)", elevationtimeslicetypeIdent, __gong__toRawStringLiteral(elevationtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevationtimeslicetypeIdent, __gong__toRawStringLiteral(elevationtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevationtimeslicetypeIdent, __gong__toRawStringLiteral(elevationtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevationtimeslicetypeIdent, __gong__toRawStringLiteral(elevationtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", elevationtimeslicetypeIdent, __gong__toRawStringLiteral(elevationtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", elevationtimeslicetypeIdent, elevationtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", elevationtimeslicetypeIdent, elevationtimeslicetype.CorrectionNumber))
			if elevationtimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elevationtimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, elevationtype := range __gong__sortStageSetInstances(stage.ElevationTypes, stage.ElevationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevationtypeIdent := elevationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevationType{Name: %s}).Stage(stage)", elevationtypeIdent, __gong__toRawStringLiteral(elevationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevationtypeIdent, __gong__toRawStringLiteral(elevationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevationtypeIdent, __gong__toRawStringLiteral(elevationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevationtypeIdent, __gong__toRawStringLiteral(elevationtype.Description)))
			for _, elem := range elevationtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", elevationtypeIdent, elevationtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, glidepathtimeslicepropertytype := range __gong__sortStageSetInstances(stage.GlidepathTimeSlicePropertyTypes, stage.GlidepathTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			glidepathtimeslicepropertytypeIdent := glidepathtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GlidepathTimeSlicePropertyType{Name: %s}).Stage(stage)", glidepathtimeslicepropertytypeIdent, __gong__toRawStringLiteral(glidepathtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glidepathtimeslicepropertytypeIdent, __gong__toRawStringLiteral(glidepathtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", glidepathtimeslicepropertytypeIdent, glidepathtimeslicepropertytype.Owns))
			if glidepathtimeslicepropertytype.GlidepathTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := glidepathtimeslicepropertytype.GlidepathTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GlidepathTimeSlice = %s", glidepathtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, glidepathtimeslicetype := range __gong__sortStageSetInstances(stage.GlidepathTimeSliceTypes, stage.GlidepathTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			glidepathtimeslicetypeIdent := glidepathtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GlidepathTimeSliceType{Name: %s}).Stage(stage)", glidepathtimeslicetypeIdent, __gong__toRawStringLiteral(glidepathtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glidepathtimeslicetypeIdent, __gong__toRawStringLiteral(glidepathtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", glidepathtimeslicetypeIdent, __gong__toRawStringLiteral(glidepathtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", glidepathtimeslicetypeIdent, __gong__toRawStringLiteral(glidepathtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", glidepathtimeslicetypeIdent, __gong__toRawStringLiteral(glidepathtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", glidepathtimeslicetypeIdent, glidepathtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", glidepathtimeslicetypeIdent, glidepathtimeslicetype.CorrectionNumber))
			if glidepathtimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := glidepathtimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, glidepathtype := range __gong__sortStageSetInstances(stage.GlidepathTypes, stage.GlidepathType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			glidepathtypeIdent := glidepathtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GlidepathType{Name: %s}).Stage(stage)", glidepathtypeIdent, __gong__toRawStringLiteral(glidepathtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glidepathtypeIdent, __gong__toRawStringLiteral(glidepathtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", glidepathtypeIdent, __gong__toRawStringLiteral(glidepathtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", glidepathtypeIdent, __gong__toRawStringLiteral(glidepathtype.Description)))
			for _, elem := range glidepathtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", glidepathtypeIdent, glidepathtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, informationservicepropertytype := range __gong__sortStageSetInstances(stage.InformationServicePropertyTypes, stage.InformationServicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			informationservicepropertytypeIdent := informationservicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InformationServicePropertyType{Name: %s}).Stage(stage)", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", informationservicepropertytypeIdent, informationservicepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Actuate)))
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
		for _, localizertimeslicepropertytype := range __gong__sortStageSetInstances(stage.LocalizerTimeSlicePropertyTypes, stage.LocalizerTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			localizertimeslicepropertytypeIdent := localizertimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LocalizerTimeSlicePropertyType{Name: %s}).Stage(stage)", localizertimeslicepropertytypeIdent, __gong__toRawStringLiteral(localizertimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", localizertimeslicepropertytypeIdent, __gong__toRawStringLiteral(localizertimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", localizertimeslicepropertytypeIdent, localizertimeslicepropertytype.Owns))
			if localizertimeslicepropertytype.LocalizerTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localizertimeslicepropertytype.LocalizerTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LocalizerTimeSlice = %s", localizertimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, localizertimeslicetype := range __gong__sortStageSetInstances(stage.LocalizerTimeSliceTypes, stage.LocalizerTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			localizertimeslicetypeIdent := localizertimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LocalizerTimeSliceType{Name: %s}).Stage(stage)", localizertimeslicetypeIdent, __gong__toRawStringLiteral(localizertimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", localizertimeslicetypeIdent, __gong__toRawStringLiteral(localizertimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", localizertimeslicetypeIdent, __gong__toRawStringLiteral(localizertimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", localizertimeslicetypeIdent, __gong__toRawStringLiteral(localizertimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", localizertimeslicetypeIdent, __gong__toRawStringLiteral(localizertimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", localizertimeslicetypeIdent, localizertimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", localizertimeslicetypeIdent, localizertimeslicetype.CorrectionNumber))
			if localizertimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localizertimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, localizertype := range __gong__sortStageSetInstances(stage.LocalizerTypes, stage.LocalizerType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			localizertypeIdent := localizertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LocalizerType{Name: %s}).Stage(stage)", localizertypeIdent, __gong__toRawStringLiteral(localizertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", localizertypeIdent, __gong__toRawStringLiteral(localizertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", localizertypeIdent, __gong__toRawStringLiteral(localizertype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", localizertypeIdent, __gong__toRawStringLiteral(localizertype.Description)))
			for _, elem := range localizertype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", localizertypeIdent, localizertypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, markerbeacontimeslicepropertytype := range __gong__sortStageSetInstances(stage.MarkerBeaconTimeSlicePropertyTypes, stage.MarkerBeaconTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			markerbeacontimeslicepropertytypeIdent := markerbeacontimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MarkerBeaconTimeSlicePropertyType{Name: %s}).Stage(stage)", markerbeacontimeslicepropertytypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", markerbeacontimeslicepropertytypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", markerbeacontimeslicepropertytypeIdent, markerbeacontimeslicepropertytype.Owns))
			if markerbeacontimeslicepropertytype.MarkerBeaconTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := markerbeacontimeslicepropertytype.MarkerBeaconTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MarkerBeaconTimeSlice = %s", markerbeacontimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, markerbeacontimeslicetype := range __gong__sortStageSetInstances(stage.MarkerBeaconTimeSliceTypes, stage.MarkerBeaconTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			markerbeacontimeslicetypeIdent := markerbeacontimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MarkerBeaconTimeSliceType{Name: %s}).Stage(stage)", markerbeacontimeslicetypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", markerbeacontimeslicetypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", markerbeacontimeslicetypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", markerbeacontimeslicetypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", markerbeacontimeslicetypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", markerbeacontimeslicetypeIdent, markerbeacontimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", markerbeacontimeslicetypeIdent, markerbeacontimeslicetype.CorrectionNumber))
			if markerbeacontimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := markerbeacontimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, markerbeacontype := range __gong__sortStageSetInstances(stage.MarkerBeaconTypes, stage.MarkerBeaconType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			markerbeacontypeIdent := markerbeacontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MarkerBeaconType{Name: %s}).Stage(stage)", markerbeacontypeIdent, __gong__toRawStringLiteral(markerbeacontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", markerbeacontypeIdent, __gong__toRawStringLiteral(markerbeacontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", markerbeacontypeIdent, __gong__toRawStringLiteral(markerbeacontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", markerbeacontypeIdent, __gong__toRawStringLiteral(markerbeacontype.Description)))
			for _, elem := range markerbeacontype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", markerbeacontypeIdent, markerbeacontypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ndbtimeslicepropertytype := range __gong__sortStageSetInstances(stage.NDBTimeSlicePropertyTypes, stage.NDBTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ndbtimeslicepropertytypeIdent := ndbtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NDBTimeSlicePropertyType{Name: %s}).Stage(stage)", ndbtimeslicepropertytypeIdent, __gong__toRawStringLiteral(ndbtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ndbtimeslicepropertytypeIdent, __gong__toRawStringLiteral(ndbtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", ndbtimeslicepropertytypeIdent, ndbtimeslicepropertytype.Owns))
			if ndbtimeslicepropertytype.NDBTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ndbtimeslicepropertytype.NDBTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NDBTimeSlice = %s", ndbtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ndbtimeslicetype := range __gong__sortStageSetInstances(stage.NDBTimeSliceTypes, stage.NDBTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ndbtimeslicetypeIdent := ndbtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NDBTimeSliceType{Name: %s}).Stage(stage)", ndbtimeslicetypeIdent, __gong__toRawStringLiteral(ndbtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ndbtimeslicetypeIdent, __gong__toRawStringLiteral(ndbtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ndbtimeslicetypeIdent, __gong__toRawStringLiteral(ndbtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", ndbtimeslicetypeIdent, __gong__toRawStringLiteral(ndbtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", ndbtimeslicetypeIdent, __gong__toRawStringLiteral(ndbtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", ndbtimeslicetypeIdent, ndbtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", ndbtimeslicetypeIdent, ndbtimeslicetype.CorrectionNumber))
			if ndbtimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ndbtimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ndbtype := range __gong__sortStageSetInstances(stage.NDBTypes, stage.NDBType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ndbtypeIdent := ndbtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NDBType{Name: %s}).Stage(stage)", ndbtypeIdent, __gong__toRawStringLiteral(ndbtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ndbtypeIdent, __gong__toRawStringLiteral(ndbtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ndbtypeIdent, __gong__toRawStringLiteral(ndbtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", ndbtypeIdent, __gong__toRawStringLiteral(ndbtype.Description)))
			for _, elem := range ndbtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", ndbtypeIdent, ndbtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navaidcomponentpropertytype := range __gong__sortStageSetInstances(stage.NavaidComponentPropertyTypes, stage.NavaidComponentPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidcomponentpropertytypeIdent := navaidcomponentpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidComponentPropertyType{Name: %s}).Stage(stage)", navaidcomponentpropertytypeIdent, __gong__toRawStringLiteral(navaidcomponentpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidcomponentpropertytypeIdent, __gong__toRawStringLiteral(navaidcomponentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", navaidcomponentpropertytypeIdent, __gong__toRawStringLiteral(navaidcomponentpropertytype.NilReason)))
			if navaidcomponentpropertytype.NavaidComponent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navaidcomponentpropertytype.NavaidComponent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidComponent = %s", navaidcomponentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navaidcomponenttype := range __gong__sortStageSetInstances(stage.NavaidComponentTypes, stage.NavaidComponentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidcomponenttypeIdent := navaidcomponenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidComponentType{Name: %s}).Stage(stage)", navaidcomponenttypeIdent, __gong__toRawStringLiteral(navaidcomponenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidcomponenttypeIdent, __gong__toRawStringLiteral(navaidcomponenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navaidcomponenttypeIdent, __gong__toRawStringLiteral(navaidcomponenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navaidcomponenttypeIdent, __gong__toRawStringLiteral(navaidcomponenttype.Description)))
			for _, elem := range navaidcomponenttype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navaidcomponenttypeIdent, navaidcomponenttypeIdent, targetIdent))
			}
			if navaidcomponenttype.TheNavaidEquipment != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navaidcomponenttype.TheNavaidEquipment.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheNavaidEquipment = %s", navaidcomponenttypeIdent, targetIdent))
			}
			for _, elem := range navaidcomponenttype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navaidcomponenttypeIdent, navaidcomponenttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navaidequipmentmonitoringpropertytype := range __gong__sortStageSetInstances(stage.NavaidEquipmentMonitoringPropertyTypes, stage.NavaidEquipmentMonitoringPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidequipmentmonitoringpropertytypeIdent := navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidEquipmentMonitoringPropertyType{Name: %s}).Stage(stage)", navaidequipmentmonitoringpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidequipmentmonitoringpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", navaidequipmentmonitoringpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringpropertytype.NilReason)))
			if navaidequipmentmonitoringpropertytype.NavaidEquipmentMonitoring != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navaidequipmentmonitoringpropertytype.NavaidEquipmentMonitoring.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidEquipmentMonitoring = %s", navaidequipmentmonitoringpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navaidequipmentmonitoringtype := range __gong__sortStageSetInstances(stage.NavaidEquipmentMonitoringTypes, stage.NavaidEquipmentMonitoringType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidequipmentmonitoringtypeIdent := navaidequipmentmonitoringtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidEquipmentMonitoringType{Name: %s}).Stage(stage)", navaidequipmentmonitoringtypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidequipmentmonitoringtypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navaidequipmentmonitoringtypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navaidequipmentmonitoringtypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringtype.Description)))
			for _, elem := range navaidequipmentmonitoringtype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", navaidequipmentmonitoringtypeIdent, navaidequipmentmonitoringtypeIdent, targetIdent))
			}
			for _, elem := range navaidequipmentmonitoringtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navaidequipmentmonitoringtypeIdent, navaidequipmentmonitoringtypeIdent, targetIdent))
			}
			for _, elem := range navaidequipmentmonitoringtype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", navaidequipmentmonitoringtypeIdent, navaidequipmentmonitoringtypeIdent, targetIdent))
			}
			for _, elem := range navaidequipmentmonitoringtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navaidequipmentmonitoringtypeIdent, navaidequipmentmonitoringtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navaidequipmentpropertytype := range __gong__sortStageSetInstances(stage.NavaidEquipmentPropertyTypes, stage.NavaidEquipmentPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidequipmentpropertytypeIdent := navaidequipmentpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidEquipmentPropertyType{Name: %s}).Stage(stage)", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", navaidequipmentpropertytypeIdent, navaidequipmentpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, navaidoperationalstatuspropertytype := range __gong__sortStageSetInstances(stage.NavaidOperationalStatusPropertyTypes, stage.NavaidOperationalStatusPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidoperationalstatuspropertytypeIdent := navaidoperationalstatuspropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidOperationalStatusPropertyType{Name: %s}).Stage(stage)", navaidoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(navaidoperationalstatuspropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(navaidoperationalstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", navaidoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(navaidoperationalstatuspropertytype.NilReason)))
			if navaidoperationalstatuspropertytype.NavaidOperationalStatus != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navaidoperationalstatuspropertytype.NavaidOperationalStatus.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidOperationalStatus = %s", navaidoperationalstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navaidoperationalstatustype := range __gong__sortStageSetInstances(stage.NavaidOperationalStatusTypes, stage.NavaidOperationalStatusType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidoperationalstatustypeIdent := navaidoperationalstatustype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidOperationalStatusType{Name: %s}).Stage(stage)", navaidoperationalstatustypeIdent, __gong__toRawStringLiteral(navaidoperationalstatustype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidoperationalstatustypeIdent, __gong__toRawStringLiteral(navaidoperationalstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navaidoperationalstatustypeIdent, __gong__toRawStringLiteral(navaidoperationalstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navaidoperationalstatustypeIdent, __gong__toRawStringLiteral(navaidoperationalstatustype.Description)))
			for _, elem := range navaidoperationalstatustype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", navaidoperationalstatustypeIdent, navaidoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range navaidoperationalstatustype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navaidoperationalstatustypeIdent, navaidoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range navaidoperationalstatustype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", navaidoperationalstatustypeIdent, navaidoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range navaidoperationalstatustype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navaidoperationalstatustypeIdent, navaidoperationalstatustypeIdent, targetIdent))
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
		for _, navaidtimeslicepropertytype := range __gong__sortStageSetInstances(stage.NavaidTimeSlicePropertyTypes, stage.NavaidTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidtimeslicepropertytypeIdent := navaidtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidTimeSlicePropertyType{Name: %s}).Stage(stage)", navaidtimeslicepropertytypeIdent, __gong__toRawStringLiteral(navaidtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidtimeslicepropertytypeIdent, __gong__toRawStringLiteral(navaidtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", navaidtimeslicepropertytypeIdent, navaidtimeslicepropertytype.Owns))
			if navaidtimeslicepropertytype.NavaidTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navaidtimeslicepropertytype.NavaidTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidTimeSlice = %s", navaidtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navaidtimeslicetype := range __gong__sortStageSetInstances(stage.NavaidTimeSliceTypes, stage.NavaidTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidtimeslicetypeIdent := navaidtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidTimeSliceType{Name: %s}).Stage(stage)", navaidtimeslicetypeIdent, __gong__toRawStringLiteral(navaidtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidtimeslicetypeIdent, __gong__toRawStringLiteral(navaidtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navaidtimeslicetypeIdent, __gong__toRawStringLiteral(navaidtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navaidtimeslicetypeIdent, __gong__toRawStringLiteral(navaidtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", navaidtimeslicetypeIdent, __gong__toRawStringLiteral(navaidtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", navaidtimeslicetypeIdent, navaidtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", navaidtimeslicetypeIdent, navaidtimeslicetype.CorrectionNumber))
			for _, elem := range navaidtimeslicetype.TouchDownLiftOff {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TouchDownLiftOff = append(%s.TouchDownLiftOff, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.NavaidEquipment {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidEquipment = append(%s.NavaidEquipment, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := navaidtimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.RunwayDirection {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RunwayDirection = append(%s.RunwayDirection, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.ServedAirport {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServedAirport = append(%s.ServedAirport, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, navaidtype := range __gong__sortStageSetInstances(stage.NavaidTypes, stage.NavaidType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidtypeIdent := navaidtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidType{Name: %s}).Stage(stage)", navaidtypeIdent, __gong__toRawStringLiteral(navaidtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidtypeIdent, __gong__toRawStringLiteral(navaidtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navaidtypeIdent, __gong__toRawStringLiteral(navaidtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navaidtypeIdent, __gong__toRawStringLiteral(navaidtype.Description)))
			for _, elem := range navaidtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", navaidtypeIdent, navaidtypeIdent, targetIdent))
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
		for _, precisionapproachradartimeslicepropertytype := range __gong__sortStageSetInstances(stage.PrecisionApproachRadarTimeSlicePropertyTypes, stage.PrecisionApproachRadarTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			precisionapproachradartimeslicepropertytypeIdent := precisionapproachradartimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrecisionApproachRadarTimeSlicePropertyType{Name: %s}).Stage(stage)", precisionapproachradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", precisionapproachradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", precisionapproachradartimeslicepropertytypeIdent, precisionapproachradartimeslicepropertytype.Owns))
			if precisionapproachradartimeslicepropertytype.PrecisionApproachRadarTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := precisionapproachradartimeslicepropertytype.PrecisionApproachRadarTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PrecisionApproachRadarTimeSlice = %s", precisionapproachradartimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, precisionapproachradartimeslicetype := range __gong__sortStageSetInstances(stage.PrecisionApproachRadarTimeSliceTypes, stage.PrecisionApproachRadarTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			precisionapproachradartimeslicetypeIdent := precisionapproachradartimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrecisionApproachRadarTimeSliceType{Name: %s}).Stage(stage)", precisionapproachradartimeslicetypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", precisionapproachradartimeslicetypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", precisionapproachradartimeslicetypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", precisionapproachradartimeslicetypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", precisionapproachradartimeslicetypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetype.CorrectionNumber))
			for _, elem := range precisionapproachradartimeslicetype.Contact {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := precisionapproachradartimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range precisionapproachradartimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range precisionapproachradartimeslicetype.Reflector {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Reflector = append(%s.Reflector, %s)", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range precisionapproachradartimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, precisionapproachradartype := range __gong__sortStageSetInstances(stage.PrecisionApproachRadarTypes, stage.PrecisionApproachRadarType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			precisionapproachradartypeIdent := precisionapproachradartype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrecisionApproachRadarType{Name: %s}).Stage(stage)", precisionapproachradartypeIdent, __gong__toRawStringLiteral(precisionapproachradartype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", precisionapproachradartypeIdent, __gong__toRawStringLiteral(precisionapproachradartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", precisionapproachradartypeIdent, __gong__toRawStringLiteral(precisionapproachradartype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", precisionapproachradartypeIdent, __gong__toRawStringLiteral(precisionapproachradartype.Description)))
			for _, elem := range precisionapproachradartype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", precisionapproachradartypeIdent, precisionapproachradartypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, primarysurveillanceradartimeslicepropertytype := range __gong__sortStageSetInstances(stage.PrimarySurveillanceRadarTimeSlicePropertyTypes, stage.PrimarySurveillanceRadarTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			primarysurveillanceradartimeslicepropertytypeIdent := primarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrimarySurveillanceRadarTimeSlicePropertyType{Name: %s}).Stage(stage)", primarysurveillanceradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", primarysurveillanceradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", primarysurveillanceradartimeslicepropertytypeIdent, primarysurveillanceradartimeslicepropertytype.Owns))
			if primarysurveillanceradartimeslicepropertytype.PrimarySurveillanceRadarTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := primarysurveillanceradartimeslicepropertytype.PrimarySurveillanceRadarTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PrimarySurveillanceRadarTimeSlice = %s", primarysurveillanceradartimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, primarysurveillanceradartimeslicetype := range __gong__sortStageSetInstances(stage.PrimarySurveillanceRadarTimeSliceTypes, stage.PrimarySurveillanceRadarTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			primarysurveillanceradartimeslicetypeIdent := primarysurveillanceradartimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrimarySurveillanceRadarTimeSliceType{Name: %s}).Stage(stage)", primarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", primarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", primarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", primarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", primarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetype.CorrectionNumber))
			for _, elem := range primarysurveillanceradartimeslicetype.Contact {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range primarysurveillanceradartimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range primarysurveillanceradartimeslicetype.GroundStation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundStation = append(%s.GroundStation, %s)", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range primarysurveillanceradartimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, primarysurveillanceradartype := range __gong__sortStageSetInstances(stage.PrimarySurveillanceRadarTypes, stage.PrimarySurveillanceRadarType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			primarysurveillanceradartypeIdent := primarysurveillanceradartype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrimarySurveillanceRadarType{Name: %s}).Stage(stage)", primarysurveillanceradartypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", primarysurveillanceradartypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", primarysurveillanceradartypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", primarysurveillanceradartypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartype.Description)))
			for _, elem := range primarysurveillanceradartype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", primarysurveillanceradartypeIdent, primarysurveillanceradartypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, radarcomponentpropertytype := range __gong__sortStageSetInstances(stage.RadarComponentPropertyTypes, stage.RadarComponentPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radarcomponentpropertytypeIdent := radarcomponentpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarComponentPropertyType{Name: %s}).Stage(stage)", radarcomponentpropertytypeIdent, __gong__toRawStringLiteral(radarcomponentpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarcomponentpropertytypeIdent, __gong__toRawStringLiteral(radarcomponentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", radarcomponentpropertytypeIdent, __gong__toRawStringLiteral(radarcomponentpropertytype.NilReason)))
			if radarcomponentpropertytype.RadarComponent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := radarcomponentpropertytype.RadarComponent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarComponent = %s", radarcomponentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, radarcomponenttype := range __gong__sortStageSetInstances(stage.RadarComponentTypes, stage.RadarComponentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radarcomponenttypeIdent := radarcomponenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarComponentType{Name: %s}).Stage(stage)", radarcomponenttypeIdent, __gong__toRawStringLiteral(radarcomponenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarcomponenttypeIdent, __gong__toRawStringLiteral(radarcomponenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radarcomponenttypeIdent, __gong__toRawStringLiteral(radarcomponenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radarcomponenttypeIdent, __gong__toRawStringLiteral(radarcomponenttype.Description)))
			for _, elem := range radarcomponenttype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", radarcomponenttypeIdent, radarcomponenttypeIdent, targetIdent))
			}
			if radarcomponenttype.TheRadarEquipment != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := radarcomponenttype.TheRadarEquipment.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheRadarEquipment = %s", radarcomponenttypeIdent, targetIdent))
			}
			for _, elem := range radarcomponenttype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", radarcomponenttypeIdent, radarcomponenttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, radarequipmentpropertytype := range __gong__sortStageSetInstances(stage.RadarEquipmentPropertyTypes, stage.RadarEquipmentPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radarequipmentpropertytypeIdent := radarequipmentpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarEquipmentPropertyType{Name: %s}).Stage(stage)", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", radarequipmentpropertytypeIdent, radarequipmentpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, radarsystemtimeslicepropertytype := range __gong__sortStageSetInstances(stage.RadarSystemTimeSlicePropertyTypes, stage.RadarSystemTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radarsystemtimeslicepropertytypeIdent := radarsystemtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarSystemTimeSlicePropertyType{Name: %s}).Stage(stage)", radarsystemtimeslicepropertytypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarsystemtimeslicepropertytypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", radarsystemtimeslicepropertytypeIdent, radarsystemtimeslicepropertytype.Owns))
			if radarsystemtimeslicepropertytype.RadarSystemTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := radarsystemtimeslicepropertytype.RadarSystemTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarSystemTimeSlice = %s", radarsystemtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, radarsystemtimeslicetype := range __gong__sortStageSetInstances(stage.RadarSystemTimeSliceTypes, stage.RadarSystemTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radarsystemtimeslicetypeIdent := radarsystemtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarSystemTimeSliceType{Name: %s}).Stage(stage)", radarsystemtimeslicetypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarsystemtimeslicetypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radarsystemtimeslicetypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radarsystemtimeslicetypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", radarsystemtimeslicetypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", radarsystemtimeslicetypeIdent, radarsystemtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", radarsystemtimeslicetypeIdent, radarsystemtimeslicetype.CorrectionNumber))
			for _, elem := range radarsystemtimeslicetype.RadarEquipment {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarEquipment = append(%s.RadarEquipment, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtimeslicetype.Office {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Office = append(%s.Office, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.AirportHeliport != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := radarsystemtimeslicetype.AirportHeliport.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtimeslicetype.PARRunway {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PARRunway = append(%s.PARRunway, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := radarsystemtimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, radarsystemtype := range __gong__sortStageSetInstances(stage.RadarSystemTypes, stage.RadarSystemType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			radarsystemtypeIdent := radarsystemtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarSystemType{Name: %s}).Stage(stage)", radarsystemtypeIdent, __gong__toRawStringLiteral(radarsystemtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarsystemtypeIdent, __gong__toRawStringLiteral(radarsystemtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radarsystemtypeIdent, __gong__toRawStringLiteral(radarsystemtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radarsystemtypeIdent, __gong__toRawStringLiteral(radarsystemtype.Description)))
			for _, elem := range radarsystemtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", radarsystemtypeIdent, radarsystemtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, reflectorpropertytype := range __gong__sortStageSetInstances(stage.ReflectorPropertyTypes, stage.ReflectorPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			reflectorpropertytypeIdent := reflectorpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ReflectorPropertyType{Name: %s}).Stage(stage)", reflectorpropertytypeIdent, __gong__toRawStringLiteral(reflectorpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", reflectorpropertytypeIdent, __gong__toRawStringLiteral(reflectorpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", reflectorpropertytypeIdent, __gong__toRawStringLiteral(reflectorpropertytype.NilReason)))
			if reflectorpropertytype.Reflector != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := reflectorpropertytype.Reflector.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Reflector = %s", reflectorpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, reflectortype := range __gong__sortStageSetInstances(stage.ReflectorTypes, stage.ReflectorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			reflectortypeIdent := reflectortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ReflectorType{Name: %s}).Stage(stage)", reflectortypeIdent, __gong__toRawStringLiteral(reflectortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", reflectortypeIdent, __gong__toRawStringLiteral(reflectortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", reflectortypeIdent, __gong__toRawStringLiteral(reflectortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", reflectortypeIdent, __gong__toRawStringLiteral(reflectortype.Description)))
			if reflectortype.TouchdownReflector != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := reflectortype.TouchdownReflector.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TouchdownReflector = %s", reflectortypeIdent, targetIdent))
			}
			for _, elem := range reflectortype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", reflectortypeIdent, reflectortypeIdent, targetIdent))
			}
			for _, elem := range reflectortype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", reflectortypeIdent, reflectortypeIdent, targetIdent))
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
		for _, runwaypropertytype := range __gong__sortStageSetInstances(stage.RunwayPropertyTypes, stage.RunwayPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			runwaypropertytypeIdent := runwaypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RunwayPropertyType{Name: %s}).Stage(stage)", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", runwaypropertytypeIdent, runwaypropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, sdftimeslicepropertytype := range __gong__sortStageSetInstances(stage.SDFTimeSlicePropertyTypes, stage.SDFTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sdftimeslicepropertytypeIdent := sdftimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SDFTimeSlicePropertyType{Name: %s}).Stage(stage)", sdftimeslicepropertytypeIdent, __gong__toRawStringLiteral(sdftimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sdftimeslicepropertytypeIdent, __gong__toRawStringLiteral(sdftimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", sdftimeslicepropertytypeIdent, sdftimeslicepropertytype.Owns))
			if sdftimeslicepropertytype.SDFTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := sdftimeslicepropertytype.SDFTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SDFTimeSlice = %s", sdftimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, sdftimeslicetype := range __gong__sortStageSetInstances(stage.SDFTimeSliceTypes, stage.SDFTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sdftimeslicetypeIdent := sdftimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SDFTimeSliceType{Name: %s}).Stage(stage)", sdftimeslicetypeIdent, __gong__toRawStringLiteral(sdftimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sdftimeslicetypeIdent, __gong__toRawStringLiteral(sdftimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", sdftimeslicetypeIdent, __gong__toRawStringLiteral(sdftimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", sdftimeslicetypeIdent, __gong__toRawStringLiteral(sdftimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", sdftimeslicetypeIdent, __gong__toRawStringLiteral(sdftimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", sdftimeslicetypeIdent, sdftimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", sdftimeslicetypeIdent, sdftimeslicetype.CorrectionNumber))
			if sdftimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := sdftimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, sdftype := range __gong__sortStageSetInstances(stage.SDFTypes, stage.SDFType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sdftypeIdent := sdftype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SDFType{Name: %s}).Stage(stage)", sdftypeIdent, __gong__toRawStringLiteral(sdftype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sdftypeIdent, __gong__toRawStringLiteral(sdftype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", sdftypeIdent, __gong__toRawStringLiteral(sdftype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", sdftypeIdent, __gong__toRawStringLiteral(sdftype.Description)))
			for _, elem := range sdftype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", sdftypeIdent, sdftypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, secondarysurveillanceradartimeslicepropertytype := range __gong__sortStageSetInstances(stage.SecondarySurveillanceRadarTimeSlicePropertyTypes, stage.SecondarySurveillanceRadarTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			secondarysurveillanceradartimeslicepropertytypeIdent := secondarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SecondarySurveillanceRadarTimeSlicePropertyType{Name: %s}).Stage(stage)", secondarysurveillanceradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", secondarysurveillanceradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", secondarysurveillanceradartimeslicepropertytypeIdent, secondarysurveillanceradartimeslicepropertytype.Owns))
			if secondarysurveillanceradartimeslicepropertytype.SecondarySurveillanceRadarTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := secondarysurveillanceradartimeslicepropertytype.SecondarySurveillanceRadarTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SecondarySurveillanceRadarTimeSlice = %s", secondarysurveillanceradartimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, secondarysurveillanceradartimeslicetype := range __gong__sortStageSetInstances(stage.SecondarySurveillanceRadarTimeSliceTypes, stage.SecondarySurveillanceRadarTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			secondarysurveillanceradartimeslicetypeIdent := secondarysurveillanceradartimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SecondarySurveillanceRadarTimeSliceType{Name: %s}).Stage(stage)", secondarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", secondarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", secondarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", secondarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", secondarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetype.CorrectionNumber))
			for _, elem := range secondarysurveillanceradartimeslicetype.Contact {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range secondarysurveillanceradartimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range secondarysurveillanceradartimeslicetype.GroundStation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundStation = append(%s.GroundStation, %s)", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range secondarysurveillanceradartimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, secondarysurveillanceradartype := range __gong__sortStageSetInstances(stage.SecondarySurveillanceRadarTypes, stage.SecondarySurveillanceRadarType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			secondarysurveillanceradartypeIdent := secondarysurveillanceradartype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SecondarySurveillanceRadarType{Name: %s}).Stage(stage)", secondarysurveillanceradartypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", secondarysurveillanceradartypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", secondarysurveillanceradartypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", secondarysurveillanceradartypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartype.Description)))
			for _, elem := range secondarysurveillanceradartype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", secondarysurveillanceradartypeIdent, secondarysurveillanceradartypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, significantpointinairspacetimeslicepropertytype := range __gong__sortStageSetInstances(stage.SignificantPointInAirspaceTimeSlicePropertyTypes, stage.SignificantPointInAirspaceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			significantpointinairspacetimeslicepropertytypeIdent := significantpointinairspacetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SignificantPointInAirspaceTimeSlicePropertyType{Name: %s}).Stage(stage)", significantpointinairspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", significantpointinairspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", significantpointinairspacetimeslicepropertytypeIdent, significantpointinairspacetimeslicepropertytype.Owns))
			if significantpointinairspacetimeslicepropertytype.SignificantPointInAirspaceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := significantpointinairspacetimeslicepropertytype.SignificantPointInAirspaceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointInAirspaceTimeSlice = %s", significantpointinairspacetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, significantpointinairspacetimeslicetype := range __gong__sortStageSetInstances(stage.SignificantPointInAirspaceTimeSliceTypes, stage.SignificantPointInAirspaceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			significantpointinairspacetimeslicetypeIdent := significantpointinairspacetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SignificantPointInAirspaceTimeSliceType{Name: %s}).Stage(stage)", significantpointinairspacetimeslicetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", significantpointinairspacetimeslicetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", significantpointinairspacetimeslicetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", significantpointinairspacetimeslicetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", significantpointinairspacetimeslicetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", significantpointinairspacetimeslicetypeIdent, significantpointinairspacetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", significantpointinairspacetimeslicetypeIdent, significantpointinairspacetimeslicetype.CorrectionNumber))
			if significantpointinairspacetimeslicetype.ContainingAirspace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := significantpointinairspacetimeslicetype.ContainingAirspace.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ContainingAirspace = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_fixDesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_fixDesignatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_fixDesignatedPoint = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_navaidSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_navaidSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_navaidSystem = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_position = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_runwayPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_runwayPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_runwayPoint = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_aimingPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_aimingPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_aimingPoint = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_airportReferencePoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_airportReferencePoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_airportReferencePoint = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range significantpointinairspacetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", significantpointinairspacetimeslicetypeIdent, significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range significantpointinairspacetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", significantpointinairspacetimeslicetypeIdent, significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, significantpointinairspacetype := range __gong__sortStageSetInstances(stage.SignificantPointInAirspaceTypes, stage.SignificantPointInAirspaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			significantpointinairspacetypeIdent := significantpointinairspacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SignificantPointInAirspaceType{Name: %s}).Stage(stage)", significantpointinairspacetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", significantpointinairspacetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", significantpointinairspacetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", significantpointinairspacetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetype.Description)))
			for _, elem := range significantpointinairspacetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", significantpointinairspacetypeIdent, significantpointinairspacetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialnavigationstationstatuspropertytype := range __gong__sortStageSetInstances(stage.SpecialNavigationStationStatusPropertyTypes, stage.SpecialNavigationStationStatusPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialnavigationstationstatuspropertytypeIdent := specialnavigationstationstatuspropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationStatusPropertyType{Name: %s}).Stage(stage)", specialnavigationstationstatuspropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatuspropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationstatuspropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", specialnavigationstationstatuspropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatuspropertytype.NilReason)))
			if specialnavigationstationstatuspropertytype.SpecialNavigationStationStatus != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := specialnavigationstationstatuspropertytype.SpecialNavigationStationStatus.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialNavigationStationStatus = %s", specialnavigationstationstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialnavigationstationstatustype := range __gong__sortStageSetInstances(stage.SpecialNavigationStationStatusTypes, stage.SpecialNavigationStationStatusType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialnavigationstationstatustypeIdent := specialnavigationstationstatustype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationStatusType{Name: %s}).Stage(stage)", specialnavigationstationstatustypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatustype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationstatustypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialnavigationstationstatustypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialnavigationstationstatustypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatustype.Description)))
			for _, elem := range specialnavigationstationstatustype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", specialnavigationstationstatustypeIdent, specialnavigationstationstatustypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationstatustype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", specialnavigationstationstatustypeIdent, specialnavigationstationstatustypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationstatustype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", specialnavigationstationstatustypeIdent, specialnavigationstationstatustypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationstatustype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", specialnavigationstationstatustypeIdent, specialnavigationstationstatustypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialnavigationstationtimeslicepropertytype := range __gong__sortStageSetInstances(stage.SpecialNavigationStationTimeSlicePropertyTypes, stage.SpecialNavigationStationTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialnavigationstationtimeslicepropertytypeIdent := specialnavigationstationtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationTimeSlicePropertyType{Name: %s}).Stage(stage)", specialnavigationstationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", specialnavigationstationtimeslicepropertytypeIdent, specialnavigationstationtimeslicepropertytype.Owns))
			if specialnavigationstationtimeslicepropertytype.SpecialNavigationStationTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := specialnavigationstationtimeslicepropertytype.SpecialNavigationStationTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialNavigationStationTimeSlice = %s", specialnavigationstationtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialnavigationstationtimeslicetype := range __gong__sortStageSetInstances(stage.SpecialNavigationStationTimeSliceTypes, stage.SpecialNavigationStationTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialnavigationstationtimeslicetypeIdent := specialnavigationstationtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationTimeSliceType{Name: %s}).Stage(stage)", specialnavigationstationtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialnavigationstationtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialnavigationstationtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", specialnavigationstationtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetype.CorrectionNumber))
			if specialnavigationstationtimeslicetype.SystemChain != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := specialnavigationstationtimeslicetype.SystemChain.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemChain = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.ResponsibleOrganisation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := specialnavigationstationtimeslicetype.ResponsibleOrganisation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResponsibleOrganisation = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.Position != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := specialnavigationstationtimeslicetype.Position.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Position = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationtimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialnavigationstationtype := range __gong__sortStageSetInstances(stage.SpecialNavigationStationTypes, stage.SpecialNavigationStationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialnavigationstationtypeIdent := specialnavigationstationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationType{Name: %s}).Stage(stage)", specialnavigationstationtypeIdent, __gong__toRawStringLiteral(specialnavigationstationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationtypeIdent, __gong__toRawStringLiteral(specialnavigationstationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialnavigationstationtypeIdent, __gong__toRawStringLiteral(specialnavigationstationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialnavigationstationtypeIdent, __gong__toRawStringLiteral(specialnavigationstationtype.Description)))
			for _, elem := range specialnavigationstationtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", specialnavigationstationtypeIdent, specialnavigationstationtypeIdent, targetIdent))
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
		for _, specialnavigationsystemtimeslicepropertytype := range __gong__sortStageSetInstances(stage.SpecialNavigationSystemTimeSlicePropertyTypes, stage.SpecialNavigationSystemTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialnavigationsystemtimeslicepropertytypeIdent := specialnavigationsystemtimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationSystemTimeSlicePropertyType{Name: %s}).Stage(stage)", specialnavigationsystemtimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationsystemtimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", specialnavigationsystemtimeslicepropertytypeIdent, specialnavigationsystemtimeslicepropertytype.Owns))
			if specialnavigationsystemtimeslicepropertytype.SpecialNavigationSystemTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := specialnavigationsystemtimeslicepropertytype.SpecialNavigationSystemTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialNavigationSystemTimeSlice = %s", specialnavigationsystemtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialnavigationsystemtimeslicetype := range __gong__sortStageSetInstances(stage.SpecialNavigationSystemTimeSliceTypes, stage.SpecialNavigationSystemTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialnavigationsystemtimeslicetypeIdent := specialnavigationsystemtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationSystemTimeSliceType{Name: %s}).Stage(stage)", specialnavigationsystemtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationsystemtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialnavigationsystemtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialnavigationsystemtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", specialnavigationsystemtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", specialnavigationsystemtimeslicetypeIdent, specialnavigationsystemtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", specialnavigationsystemtimeslicetypeIdent, specialnavigationsystemtimeslicetype.CorrectionNumber))
			if specialnavigationsystemtimeslicetype.ResponsibleOrganisation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := specialnavigationsystemtimeslicetype.ResponsibleOrganisation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResponsibleOrganisation = %s", specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationsystemtimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", specialnavigationsystemtimeslicetypeIdent, specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationsystemtimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", specialnavigationsystemtimeslicetypeIdent, specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specialnavigationsystemtype := range __gong__sortStageSetInstances(stage.SpecialNavigationSystemTypes, stage.SpecialNavigationSystemType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialnavigationsystemtypeIdent := specialnavigationsystemtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationSystemType{Name: %s}).Stage(stage)", specialnavigationsystemtypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationsystemtypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialnavigationsystemtypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialnavigationsystemtypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtype.Description)))
			for _, elem := range specialnavigationsystemtype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", specialnavigationsystemtypeIdent, specialnavigationsystemtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, surveillancegroundstationpropertytype := range __gong__sortStageSetInstances(stage.SurveillanceGroundStationPropertyTypes, stage.SurveillanceGroundStationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			surveillancegroundstationpropertytypeIdent := surveillancegroundstationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SurveillanceGroundStationPropertyType{Name: %s}).Stage(stage)", surveillancegroundstationpropertytypeIdent, __gong__toRawStringLiteral(surveillancegroundstationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surveillancegroundstationpropertytypeIdent, __gong__toRawStringLiteral(surveillancegroundstationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", surveillancegroundstationpropertytypeIdent, __gong__toRawStringLiteral(surveillancegroundstationpropertytype.NilReason)))
			if surveillancegroundstationpropertytype.SurveillanceGroundStation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := surveillancegroundstationpropertytype.SurveillanceGroundStation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SurveillanceGroundStation = %s", surveillancegroundstationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, surveillancegroundstationtype := range __gong__sortStageSetInstances(stage.SurveillanceGroundStationTypes, stage.SurveillanceGroundStationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			surveillancegroundstationtypeIdent := surveillancegroundstationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SurveillanceGroundStationType{Name: %s}).Stage(stage)", surveillancegroundstationtypeIdent, __gong__toRawStringLiteral(surveillancegroundstationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surveillancegroundstationtypeIdent, __gong__toRawStringLiteral(surveillancegroundstationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", surveillancegroundstationtypeIdent, __gong__toRawStringLiteral(surveillancegroundstationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", surveillancegroundstationtypeIdent, __gong__toRawStringLiteral(surveillancegroundstationtype.Description)))
			for _, elem := range surveillancegroundstationtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", surveillancegroundstationtypeIdent, surveillancegroundstationtypeIdent, targetIdent))
			}
			if surveillancegroundstationtype.TheUnit != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := surveillancegroundstationtype.TheUnit.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheUnit = %s", surveillancegroundstationtypeIdent, targetIdent))
			}
			for _, elem := range surveillancegroundstationtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", surveillancegroundstationtypeIdent, surveillancegroundstationtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, tacantimeslicepropertytype := range __gong__sortStageSetInstances(stage.TACANTimeSlicePropertyTypes, stage.TACANTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tacantimeslicepropertytypeIdent := tacantimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TACANTimeSlicePropertyType{Name: %s}).Stage(stage)", tacantimeslicepropertytypeIdent, __gong__toRawStringLiteral(tacantimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tacantimeslicepropertytypeIdent, __gong__toRawStringLiteral(tacantimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", tacantimeslicepropertytypeIdent, tacantimeslicepropertytype.Owns))
			if tacantimeslicepropertytype.TACANTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := tacantimeslicepropertytype.TACANTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TACANTimeSlice = %s", tacantimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, tacantimeslicetype := range __gong__sortStageSetInstances(stage.TACANTimeSliceTypes, stage.TACANTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tacantimeslicetypeIdent := tacantimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TACANTimeSliceType{Name: %s}).Stage(stage)", tacantimeslicetypeIdent, __gong__toRawStringLiteral(tacantimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tacantimeslicetypeIdent, __gong__toRawStringLiteral(tacantimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", tacantimeslicetypeIdent, __gong__toRawStringLiteral(tacantimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", tacantimeslicetypeIdent, __gong__toRawStringLiteral(tacantimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", tacantimeslicetypeIdent, __gong__toRawStringLiteral(tacantimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", tacantimeslicetypeIdent, tacantimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", tacantimeslicetypeIdent, tacantimeslicetype.CorrectionNumber))
			if tacantimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := tacantimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, tacantype := range __gong__sortStageSetInstances(stage.TACANTypes, stage.TACANType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tacantypeIdent := tacantype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TACANType{Name: %s}).Stage(stage)", tacantypeIdent, __gong__toRawStringLiteral(tacantype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tacantypeIdent, __gong__toRawStringLiteral(tacantype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", tacantypeIdent, __gong__toRawStringLiteral(tacantype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", tacantypeIdent, __gong__toRawStringLiteral(tacantype.Description)))
			for _, elem := range tacantype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", tacantypeIdent, tacantypeIdent, targetIdent))
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
	if stage != nil {
		for _, vortimeslicepropertytype := range __gong__sortStageSetInstances(stage.VORTimeSlicePropertyTypes, stage.VORTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vortimeslicepropertytypeIdent := vortimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VORTimeSlicePropertyType{Name: %s}).Stage(stage)", vortimeslicepropertytypeIdent, __gong__toRawStringLiteral(vortimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vortimeslicepropertytypeIdent, __gong__toRawStringLiteral(vortimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", vortimeslicepropertytypeIdent, vortimeslicepropertytype.Owns))
			if vortimeslicepropertytype.VORTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := vortimeslicepropertytype.VORTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VORTimeSlice = %s", vortimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, vortimeslicetype := range __gong__sortStageSetInstances(stage.VORTimeSliceTypes, stage.VORTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vortimeslicetypeIdent := vortimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VORTimeSliceType{Name: %s}).Stage(stage)", vortimeslicetypeIdent, __gong__toRawStringLiteral(vortimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vortimeslicetypeIdent, __gong__toRawStringLiteral(vortimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", vortimeslicetypeIdent, __gong__toRawStringLiteral(vortimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", vortimeslicetypeIdent, __gong__toRawStringLiteral(vortimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", vortimeslicetypeIdent, __gong__toRawStringLiteral(vortimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", vortimeslicetypeIdent, vortimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", vortimeslicetypeIdent, vortimeslicetype.CorrectionNumber))
			if vortimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := vortimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.Authority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.Monitoring {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, vortype := range __gong__sortStageSetInstances(stage.VORTypes, stage.VORType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vortypeIdent := vortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VORType{Name: %s}).Stage(stage)", vortypeIdent, __gong__toRawStringLiteral(vortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vortypeIdent, __gong__toRawStringLiteral(vortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", vortypeIdent, __gong__toRawStringLiteral(vortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", vortypeIdent, __gong__toRawStringLiteral(vortype.Description)))
			for _, elem := range vortype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", vortypeIdent, vortypeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures\"", fmt.Sprintf("%q", modelsPackageName))
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

func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractnavaidequipmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractnavaidequipmenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractnavaidequipmenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractnavaidequipmenttype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractNavaidEquipmentType", fieldName)
	}
	return
}

func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractnavigationsystemcheckpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractnavigationsystemcheckpointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractnavigationsystemcheckpointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractnavigationsystemcheckpointtype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractNavigationSystemCheckpointType", fieldName)
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

func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractradarequipmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractradarequipmenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractradarequipmenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractradarequipmenttype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractRadarEquipmentType", fieldName)
	}
	return
}

func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractsurveillanceradartype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractsurveillanceradartype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractsurveillanceradartype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractsurveillanceradartype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractSurveillanceRadarType", fieldName)
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

func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := angleindicationtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", angleindicationtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", angleindicationtimeslicepropertytype.Owns)

	case "AngleIndicationTimeSlice":
		if angleindicationtimeslicepropertytype.AngleIndicationTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AngleIndicationTimeSlice", angleindicationtimeslicepropertytype.AngleIndicationTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AngleIndicationTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AngleIndicationTimeSlicePropertyType", fieldName)
	}
	return
}

func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := angleindicationtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", angleindicationtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", angleindicationtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", angleindicationtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", angleindicationtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", angleindicationtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", angleindicationtimeslicetype.CorrectionNumber)

	case "Fix":
		if angleindicationtimeslicetype.Fix != nil {
			res = __gong__marshallPointer(ident, "Fix", angleindicationtimeslicetype.Fix.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Fix", "nil")
		}
	case "PointChoice_fixDesignatedPoint":
		if angleindicationtimeslicetype.PointChoice_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_fixDesignatedPoint", angleindicationtimeslicetype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_fixDesignatedPoint", "nil")
		}
	case "PointChoice_navaidSystem":
		if angleindicationtimeslicetype.PointChoice_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "PointChoice_navaidSystem", angleindicationtimeslicetype.PointChoice_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_navaidSystem", "nil")
		}
	case "PointChoice_position":
		if angleindicationtimeslicetype.PointChoice_position != nil {
			res = __gong__marshallPointer(ident, "PointChoice_position", angleindicationtimeslicetype.PointChoice_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_position", "nil")
		}
	case "PointChoice_runwayPoint":
		if angleindicationtimeslicetype.PointChoice_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_runwayPoint", angleindicationtimeslicetype.PointChoice_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_runwayPoint", "nil")
		}
	case "PointChoice_aimingPoint":
		if angleindicationtimeslicetype.PointChoice_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_aimingPoint", angleindicationtimeslicetype.PointChoice_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_aimingPoint", "nil")
		}
	case "PointChoice_airportReferencePoint":
		if angleindicationtimeslicetype.PointChoice_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_airportReferencePoint", angleindicationtimeslicetype.PointChoice_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_airportReferencePoint", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range angleindicationtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_24 := range angleindicationtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_24.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AngleIndicationTimeSliceType", fieldName)
	}
	return
}

func (angleindicationtype *AngleIndicationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := angleindicationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", angleindicationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", angleindicationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", angleindicationtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _angleindicationtimeslicepropertytype := range angleindicationtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _angleindicationtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AngleIndicationType", fieldName)
	}
	return
}

func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityfornavaidequipmentpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", authorityfornavaidequipmentpropertytype.NilReason)

	case "AuthorityForNavaidEquipment":
		if authorityfornavaidequipmentpropertytype.AuthorityForNavaidEquipment != nil {
			res = __gong__marshallPointer(ident, "AuthorityForNavaidEquipment", authorityfornavaidequipmentpropertytype.AuthorityForNavaidEquipment.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AuthorityForNavaidEquipment", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForNavaidEquipmentPropertyType", fieldName)
	}
	return
}

func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityfornavaidequipmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityfornavaidequipmenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", authorityfornavaidequipmenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", authorityfornavaidequipmenttype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range authorityfornavaidequipmenttype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheOrganisationAuthority":
		if authorityfornavaidequipmenttype.TheOrganisationAuthority != nil {
			res = __gong__marshallPointer(ident, "TheOrganisationAuthority", authorityfornavaidequipmenttype.TheOrganisationAuthority.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheOrganisationAuthority", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_1 := range authorityfornavaidequipmenttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_1.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForNavaidEquipmentType", fieldName)
	}
	return
}

func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityforspecialnavigationstationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityforspecialnavigationstationpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", authorityforspecialnavigationstationpropertytype.NilReason)

	case "AuthorityForSpecialNavigationStation":
		if authorityforspecialnavigationstationpropertytype.AuthorityForSpecialNavigationStation != nil {
			res = __gong__marshallPointer(ident, "AuthorityForSpecialNavigationStation", authorityforspecialnavigationstationpropertytype.AuthorityForSpecialNavigationStation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AuthorityForSpecialNavigationStation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForSpecialNavigationStationPropertyType", fieldName)
	}
	return
}

func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityforspecialnavigationstationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityforspecialnavigationstationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", authorityforspecialnavigationstationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", authorityforspecialnavigationstationtype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range authorityforspecialnavigationstationtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheOrganisationAuthority":
		if authorityforspecialnavigationstationtype.TheOrganisationAuthority != nil {
			res = __gong__marshallPointer(ident, "TheOrganisationAuthority", authorityforspecialnavigationstationtype.TheOrganisationAuthority.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheOrganisationAuthority", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_2 := range authorityforspecialnavigationstationtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_2.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForSpecialNavigationStationType", fieldName)
	}
	return
}

func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityforspecialnavigationsystempropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityforspecialnavigationsystempropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", authorityforspecialnavigationsystempropertytype.NilReason)

	case "AuthorityForSpecialNavigationSystem":
		if authorityforspecialnavigationsystempropertytype.AuthorityForSpecialNavigationSystem != nil {
			res = __gong__marshallPointer(ident, "AuthorityForSpecialNavigationSystem", authorityforspecialnavigationsystempropertytype.AuthorityForSpecialNavigationSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AuthorityForSpecialNavigationSystem", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForSpecialNavigationSystemPropertyType", fieldName)
	}
	return
}

func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityforspecialnavigationsystemtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityforspecialnavigationsystemtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", authorityforspecialnavigationsystemtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", authorityforspecialnavigationsystemtype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range authorityforspecialnavigationsystemtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheOrganisationAuthority":
		if authorityforspecialnavigationsystemtype.TheOrganisationAuthority != nil {
			res = __gong__marshallPointer(ident, "TheOrganisationAuthority", authorityforspecialnavigationsystemtype.TheOrganisationAuthority.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheOrganisationAuthority", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_3 := range authorityforspecialnavigationsystemtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_3.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForSpecialNavigationSystemType", fieldName)
	}
	return
}

func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := azimuthtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", azimuthtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", azimuthtimeslicepropertytype.Owns)

	case "AzimuthTimeSlice":
		if azimuthtimeslicepropertytype.AzimuthTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AzimuthTimeSlice", azimuthtimeslicepropertytype.AzimuthTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AzimuthTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AzimuthTimeSlicePropertyType", fieldName)
	}
	return
}

func (azimuthtimeslicetype *AzimuthTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := azimuthtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", azimuthtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", azimuthtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", azimuthtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", azimuthtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", azimuthtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", azimuthtimeslicetype.CorrectionNumber)

	case "Location":
		if azimuthtimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", azimuthtimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range azimuthtimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range azimuthtimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range azimuthtimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range azimuthtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_4 := range azimuthtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_4.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AzimuthTimeSliceType", fieldName)
	}
	return
}

func (azimuthtype *AzimuthType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := azimuthtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", azimuthtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", azimuthtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", azimuthtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _azimuthtimeslicepropertytype := range azimuthtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _azimuthtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AzimuthType", fieldName)
	}
	return
}

func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := checkpointinstimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", checkpointinstimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", checkpointinstimeslicepropertytype.Owns)

	case "CheckpointINSTimeSlice":
		if checkpointinstimeslicepropertytype.CheckpointINSTimeSlice != nil {
			res = __gong__marshallPointer(ident, "CheckpointINSTimeSlice", checkpointinstimeslicepropertytype.CheckpointINSTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CheckpointINSTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CheckpointINSTimeSlicePropertyType", fieldName)
	}
	return
}

func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := checkpointinstimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", checkpointinstimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", checkpointinstimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", checkpointinstimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", checkpointinstimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", checkpointinstimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", checkpointinstimeslicetype.CorrectionNumber)

	case "Position":
		if checkpointinstimeslicetype.Position != nil {
			res = __gong__marshallPointer(ident, "Position", checkpointinstimeslicetype.Position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Position", "nil")
		}
	case "AirportHeliport":
		if checkpointinstimeslicetype.AirportHeliport != nil {
			res = __gong__marshallPointer(ident, "AirportHeliport", checkpointinstimeslicetype.AirportHeliport.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirportHeliport", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range checkpointinstimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_5 := range checkpointinstimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_5.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CheckpointINSTimeSliceType", fieldName)
	}
	return
}

func (checkpointinstype *CheckpointINSType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := checkpointinstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", checkpointinstype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", checkpointinstype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", checkpointinstype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _checkpointinstimeslicepropertytype := range checkpointinstype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _checkpointinstimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CheckpointINSType", fieldName)
	}
	return
}

func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := checkpointvortimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", checkpointvortimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", checkpointvortimeslicepropertytype.Owns)

	case "CheckpointVORTimeSlice":
		if checkpointvortimeslicepropertytype.CheckpointVORTimeSlice != nil {
			res = __gong__marshallPointer(ident, "CheckpointVORTimeSlice", checkpointvortimeslicepropertytype.CheckpointVORTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CheckpointVORTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CheckpointVORTimeSlicePropertyType", fieldName)
	}
	return
}

func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := checkpointvortimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", checkpointvortimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", checkpointvortimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", checkpointvortimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", checkpointvortimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", checkpointvortimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", checkpointvortimeslicetype.CorrectionNumber)

	case "Position":
		if checkpointvortimeslicetype.Position != nil {
			res = __gong__marshallPointer(ident, "Position", checkpointvortimeslicetype.Position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Position", "nil")
		}
	case "AirportHeliport":
		if checkpointvortimeslicetype.AirportHeliport != nil {
			res = __gong__marshallPointer(ident, "AirportHeliport", checkpointvortimeslicetype.AirportHeliport.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirportHeliport", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range checkpointvortimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "CheckPointFacility":
		if checkpointvortimeslicetype.CheckPointFacility != nil {
			res = __gong__marshallPointer(ident, "CheckPointFacility", checkpointvortimeslicetype.CheckPointFacility.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CheckPointFacility", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_6 := range checkpointvortimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_6.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CheckpointVORTimeSliceType", fieldName)
	}
	return
}

func (checkpointvortype *CheckpointVORType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := checkpointvortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", checkpointvortype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", checkpointvortype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", checkpointvortype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _checkpointvortimeslicepropertytype := range checkpointvortype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _checkpointvortimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CheckpointVORType", fieldName)
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

func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := dmetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", dmetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", dmetimeslicepropertytype.Owns)

	case "DMETimeSlice":
		if dmetimeslicepropertytype.DMETimeSlice != nil {
			res = __gong__marshallPointer(ident, "DMETimeSlice", dmetimeslicepropertytype.DMETimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DMETimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DMETimeSlicePropertyType", fieldName)
	}
	return
}

func (dmetimeslicetype *DMETimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := dmetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", dmetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", dmetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", dmetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", dmetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", dmetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", dmetimeslicetype.CorrectionNumber)

	case "Location":
		if dmetimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", dmetimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range dmetimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range dmetimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range dmetimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range dmetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_8 := range dmetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_8.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DMETimeSliceType", fieldName)
	}
	return
}

func (dmetype *DMEType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := dmetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", dmetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", dmetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", dmetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _dmetimeslicepropertytype := range dmetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _dmetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DMEType", fieldName)
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

func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := designatedpointtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", designatedpointtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", designatedpointtimeslicepropertytype.Owns)

	case "DesignatedPointTimeSlice":
		if designatedpointtimeslicepropertytype.DesignatedPointTimeSlice != nil {
			res = __gong__marshallPointer(ident, "DesignatedPointTimeSlice", designatedpointtimeslicepropertytype.DesignatedPointTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DesignatedPointTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DesignatedPointTimeSlicePropertyType", fieldName)
	}
	return
}

func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := designatedpointtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", designatedpointtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", designatedpointtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", designatedpointtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", designatedpointtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", designatedpointtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", designatedpointtimeslicetype.CorrectionNumber)

	case "Location":
		if designatedpointtimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", designatedpointtimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "AimingPoint":
		if designatedpointtimeslicetype.AimingPoint != nil {
			res = __gong__marshallPointer(ident, "AimingPoint", designatedpointtimeslicetype.AimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AimingPoint", "nil")
		}
	case "AirportHeliport":
		if designatedpointtimeslicetype.AirportHeliport != nil {
			res = __gong__marshallPointer(ident, "AirportHeliport", designatedpointtimeslicetype.AirportHeliport.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirportHeliport", "nil")
		}
	case "RunwayPoint":
		if designatedpointtimeslicetype.RunwayPoint != nil {
			res = __gong__marshallPointer(ident, "RunwayPoint", designatedpointtimeslicetype.RunwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RunwayPoint", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range designatedpointtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_26 := range designatedpointtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_26.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DesignatedPointTimeSliceType", fieldName)
	}
	return
}

func (designatedpointtype *DesignatedPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := designatedpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", designatedpointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", designatedpointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", designatedpointtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _designatedpointtimeslicepropertytype := range designatedpointtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _designatedpointtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DesignatedPointType", fieldName)
	}
	return
}

func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directionfindertimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directionfindertimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", directionfindertimeslicepropertytype.Owns)

	case "DirectionFinderTimeSlice":
		if directionfindertimeslicepropertytype.DirectionFinderTimeSlice != nil {
			res = __gong__marshallPointer(ident, "DirectionFinderTimeSlice", directionfindertimeslicepropertytype.DirectionFinderTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DirectionFinderTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DirectionFinderTimeSlicePropertyType", fieldName)
	}
	return
}

func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directionfindertimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directionfindertimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", directionfindertimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", directionfindertimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", directionfindertimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", directionfindertimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", directionfindertimeslicetype.CorrectionNumber)

	case "Location":
		if directionfindertimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", directionfindertimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range directionfindertimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range directionfindertimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range directionfindertimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range directionfindertimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "InformationProvision":
		var sb strings.Builder
		for _, _informationservicepropertytype := range directionfindertimeslicetype.InformationProvision {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "InformationProvision", _informationservicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_7 := range directionfindertimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_7.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DirectionFinderTimeSliceType", fieldName)
	}
	return
}

func (directionfindertype *DirectionFinderType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directionfindertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directionfindertype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", directionfindertype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", directionfindertype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _directionfindertimeslicepropertytype := range directionfindertype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _directionfindertimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DirectionFinderType", fieldName)
	}
	return
}

func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := distanceindicationtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", distanceindicationtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", distanceindicationtimeslicepropertytype.Owns)

	case "DistanceIndicationTimeSlice":
		if distanceindicationtimeslicepropertytype.DistanceIndicationTimeSlice != nil {
			res = __gong__marshallPointer(ident, "DistanceIndicationTimeSlice", distanceindicationtimeslicepropertytype.DistanceIndicationTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DistanceIndicationTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DistanceIndicationTimeSlicePropertyType", fieldName)
	}
	return
}

func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := distanceindicationtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", distanceindicationtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", distanceindicationtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", distanceindicationtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", distanceindicationtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", distanceindicationtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", distanceindicationtimeslicetype.CorrectionNumber)

	case "Fix":
		if distanceindicationtimeslicetype.Fix != nil {
			res = __gong__marshallPointer(ident, "Fix", distanceindicationtimeslicetype.Fix.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Fix", "nil")
		}
	case "PointChoice_fixDesignatedPoint":
		if distanceindicationtimeslicetype.PointChoice_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_fixDesignatedPoint", distanceindicationtimeslicetype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_fixDesignatedPoint", "nil")
		}
	case "PointChoice_navaidSystem":
		if distanceindicationtimeslicetype.PointChoice_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "PointChoice_navaidSystem", distanceindicationtimeslicetype.PointChoice_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_navaidSystem", "nil")
		}
	case "PointChoice_position":
		if distanceindicationtimeslicetype.PointChoice_position != nil {
			res = __gong__marshallPointer(ident, "PointChoice_position", distanceindicationtimeslicetype.PointChoice_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_position", "nil")
		}
	case "PointChoice_runwayPoint":
		if distanceindicationtimeslicetype.PointChoice_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_runwayPoint", distanceindicationtimeslicetype.PointChoice_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_runwayPoint", "nil")
		}
	case "PointChoice_aimingPoint":
		if distanceindicationtimeslicetype.PointChoice_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_aimingPoint", distanceindicationtimeslicetype.PointChoice_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_aimingPoint", "nil")
		}
	case "PointChoice_airportReferencePoint":
		if distanceindicationtimeslicetype.PointChoice_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "PointChoice_airportReferencePoint", distanceindicationtimeslicetype.PointChoice_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointChoice_airportReferencePoint", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range distanceindicationtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_25 := range distanceindicationtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_25.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DistanceIndicationTimeSliceType", fieldName)
	}
	return
}

func (distanceindicationtype *DistanceIndicationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := distanceindicationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", distanceindicationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", distanceindicationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", distanceindicationtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _distanceindicationtimeslicepropertytype := range distanceindicationtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _distanceindicationtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DistanceIndicationType", fieldName)
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
		for _, _a_extension := range elevatedpointtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedPointType", fieldName)
	}
	return
}

func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevationtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevationtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", elevationtimeslicepropertytype.Owns)

	case "ElevationTimeSlice":
		if elevationtimeslicepropertytype.ElevationTimeSlice != nil {
			res = __gong__marshallPointer(ident, "ElevationTimeSlice", elevationtimeslicepropertytype.ElevationTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ElevationTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevationTimeSlicePropertyType", fieldName)
	}
	return
}

func (elevationtimeslicetype *ElevationTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevationtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevationtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", elevationtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", elevationtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", elevationtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", elevationtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", elevationtimeslicetype.CorrectionNumber)

	case "Location":
		if elevationtimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", elevationtimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range elevationtimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range elevationtimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range elevationtimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range elevationtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_9 := range elevationtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_9.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevationTimeSliceType", fieldName)
	}
	return
}

func (elevationtype *ElevationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", elevationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", elevationtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _elevationtimeslicepropertytype := range elevationtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _elevationtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevationType", fieldName)
	}
	return
}

func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := glidepathtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", glidepathtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", glidepathtimeslicepropertytype.Owns)

	case "GlidepathTimeSlice":
		if glidepathtimeslicepropertytype.GlidepathTimeSlice != nil {
			res = __gong__marshallPointer(ident, "GlidepathTimeSlice", glidepathtimeslicepropertytype.GlidepathTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GlidepathTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GlidepathTimeSlicePropertyType", fieldName)
	}
	return
}

func (glidepathtimeslicetype *GlidepathTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := glidepathtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", glidepathtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", glidepathtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", glidepathtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", glidepathtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", glidepathtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", glidepathtimeslicetype.CorrectionNumber)

	case "Location":
		if glidepathtimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", glidepathtimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range glidepathtimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range glidepathtimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range glidepathtimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range glidepathtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_10 := range glidepathtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_10.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GlidepathTimeSliceType", fieldName)
	}
	return
}

func (glidepathtype *GlidepathType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := glidepathtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", glidepathtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", glidepathtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", glidepathtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _glidepathtimeslicepropertytype := range glidepathtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _glidepathtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GlidepathType", fieldName)
	}
	return
}

func (informationservicepropertytype *InformationServicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := informationservicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", informationservicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", informationservicepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", informationservicepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", informationservicepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", informationservicepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", informationservicepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", informationservicepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", informationservicepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", informationservicepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", informationservicepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct InformationServicePropertyType", fieldName)
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
		for _, _a_extension_28 := range linguisticnotetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_28.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LinguisticNoteType", fieldName)
	}
	return
}

func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := localizertimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", localizertimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", localizertimeslicepropertytype.Owns)

	case "LocalizerTimeSlice":
		if localizertimeslicepropertytype.LocalizerTimeSlice != nil {
			res = __gong__marshallPointer(ident, "LocalizerTimeSlice", localizertimeslicepropertytype.LocalizerTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LocalizerTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LocalizerTimeSlicePropertyType", fieldName)
	}
	return
}

func (localizertimeslicetype *LocalizerTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := localizertimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", localizertimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", localizertimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", localizertimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", localizertimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", localizertimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", localizertimeslicetype.CorrectionNumber)

	case "Location":
		if localizertimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", localizertimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range localizertimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range localizertimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range localizertimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range localizertimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_11 := range localizertimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_11.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LocalizerTimeSliceType", fieldName)
	}
	return
}

func (localizertype *LocalizerType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := localizertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", localizertype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", localizertype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", localizertype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _localizertimeslicepropertytype := range localizertype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _localizertimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LocalizerType", fieldName)
	}
	return
}

func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := markerbeacontimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", markerbeacontimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", markerbeacontimeslicepropertytype.Owns)

	case "MarkerBeaconTimeSlice":
		if markerbeacontimeslicepropertytype.MarkerBeaconTimeSlice != nil {
			res = __gong__marshallPointer(ident, "MarkerBeaconTimeSlice", markerbeacontimeslicepropertytype.MarkerBeaconTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MarkerBeaconTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MarkerBeaconTimeSlicePropertyType", fieldName)
	}
	return
}

func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := markerbeacontimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", markerbeacontimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", markerbeacontimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", markerbeacontimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", markerbeacontimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", markerbeacontimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", markerbeacontimeslicetype.CorrectionNumber)

	case "Location":
		if markerbeacontimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", markerbeacontimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range markerbeacontimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range markerbeacontimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range markerbeacontimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range markerbeacontimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_12 := range markerbeacontimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_12.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MarkerBeaconTimeSliceType", fieldName)
	}
	return
}

func (markerbeacontype *MarkerBeaconType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := markerbeacontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", markerbeacontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", markerbeacontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", markerbeacontype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _markerbeacontimeslicepropertytype := range markerbeacontype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _markerbeacontimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MarkerBeaconType", fieldName)
	}
	return
}

func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ndbtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ndbtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", ndbtimeslicepropertytype.Owns)

	case "NDBTimeSlice":
		if ndbtimeslicepropertytype.NDBTimeSlice != nil {
			res = __gong__marshallPointer(ident, "NDBTimeSlice", ndbtimeslicepropertytype.NDBTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "NDBTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NDBTimeSlicePropertyType", fieldName)
	}
	return
}

func (ndbtimeslicetype *NDBTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ndbtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ndbtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", ndbtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", ndbtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", ndbtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", ndbtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", ndbtimeslicetype.CorrectionNumber)

	case "Location":
		if ndbtimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", ndbtimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range ndbtimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range ndbtimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range ndbtimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range ndbtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_17 := range ndbtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_17.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NDBTimeSliceType", fieldName)
	}
	return
}

func (ndbtype *NDBType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ndbtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ndbtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", ndbtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", ndbtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _ndbtimeslicepropertytype := range ndbtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _ndbtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NDBType", fieldName)
	}
	return
}

func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidcomponentpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidcomponentpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", navaidcomponentpropertytype.NilReason)

	case "NavaidComponent":
		if navaidcomponentpropertytype.NavaidComponent != nil {
			res = __gong__marshallPointer(ident, "NavaidComponent", navaidcomponentpropertytype.NavaidComponent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "NavaidComponent", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidComponentPropertyType", fieldName)
	}
	return
}

func (navaidcomponenttype *NavaidComponentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidcomponenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidcomponenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", navaidcomponenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", navaidcomponenttype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range navaidcomponenttype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheNavaidEquipment":
		if navaidcomponenttype.TheNavaidEquipment != nil {
			res = __gong__marshallPointer(ident, "TheNavaidEquipment", navaidcomponenttype.TheNavaidEquipment.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheNavaidEquipment", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_14 := range navaidcomponenttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_14.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidComponentType", fieldName)
	}
	return
}

func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidequipmentmonitoringpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", navaidequipmentmonitoringpropertytype.NilReason)

	case "NavaidEquipmentMonitoring":
		if navaidequipmentmonitoringpropertytype.NavaidEquipmentMonitoring != nil {
			res = __gong__marshallPointer(ident, "NavaidEquipmentMonitoring", navaidequipmentmonitoringpropertytype.NavaidEquipmentMonitoring.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "NavaidEquipmentMonitoring", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidEquipmentMonitoringPropertyType", fieldName)
	}
	return
}

func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidequipmentmonitoringtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidequipmentmonitoringtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", navaidequipmentmonitoringtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", navaidequipmentmonitoringtype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range navaidequipmentmonitoringtype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range navaidequipmentmonitoringtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range navaidequipmentmonitoringtype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_15 := range navaidequipmentmonitoringtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_15.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidEquipmentMonitoringType", fieldName)
	}
	return
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidequipmentpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidequipmentpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", navaidequipmentpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", navaidequipmentpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", navaidequipmentpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", navaidequipmentpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", navaidequipmentpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", navaidequipmentpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", navaidequipmentpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", navaidequipmentpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", navaidequipmentpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidEquipmentPropertyType", fieldName)
	}
	return
}

func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidoperationalstatuspropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidoperationalstatuspropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", navaidoperationalstatuspropertytype.NilReason)

	case "NavaidOperationalStatus":
		if navaidoperationalstatuspropertytype.NavaidOperationalStatus != nil {
			res = __gong__marshallPointer(ident, "NavaidOperationalStatus", navaidoperationalstatuspropertytype.NavaidOperationalStatus.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "NavaidOperationalStatus", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidOperationalStatusPropertyType", fieldName)
	}
	return
}

func (navaidoperationalstatustype *NavaidOperationalStatusType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidoperationalstatustype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidoperationalstatustype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", navaidoperationalstatustype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", navaidoperationalstatustype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range navaidoperationalstatustype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range navaidoperationalstatustype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range navaidoperationalstatustype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_16 := range navaidoperationalstatustype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_16.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidOperationalStatusType", fieldName)
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

func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", navaidtimeslicepropertytype.Owns)

	case "NavaidTimeSlice":
		if navaidtimeslicepropertytype.NavaidTimeSlice != nil {
			res = __gong__marshallPointer(ident, "NavaidTimeSlice", navaidtimeslicepropertytype.NavaidTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "NavaidTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidTimeSlicePropertyType", fieldName)
	}
	return
}

func (navaidtimeslicetype *NavaidTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", navaidtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", navaidtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", navaidtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", navaidtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", navaidtimeslicetype.CorrectionNumber)

	case "TouchDownLiftOff":
		var sb strings.Builder
		for _, _touchdownliftoffpropertytype := range navaidtimeslicetype.TouchDownLiftOff {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TouchDownLiftOff", _touchdownliftoffpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NavaidEquipment":
		var sb strings.Builder
		for _, _navaidcomponentpropertytype := range navaidtimeslicetype.NavaidEquipment {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NavaidEquipment", _navaidcomponentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Location":
		if navaidtimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", navaidtimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "RunwayDirection":
		var sb strings.Builder
		for _, _runwaydirectionpropertytype := range navaidtimeslicetype.RunwayDirection {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RunwayDirection", _runwaydirectionpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ServedAirport":
		var sb strings.Builder
		for _, _airportheliportpropertytype := range navaidtimeslicetype.ServedAirport {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ServedAirport", _airportheliportpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range navaidtimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range navaidtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_13 := range navaidtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_13.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidTimeSliceType", fieldName)
	}
	return
}

func (navaidtype *NavaidType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", navaidtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", navaidtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _navaidtimeslicepropertytype := range navaidtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _navaidtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidType", fieldName)
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
		for _, _a_extension_29 := range notetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_29.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteType", fieldName)
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

func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := precisionapproachradartimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", precisionapproachradartimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", precisionapproachradartimeslicepropertytype.Owns)

	case "PrecisionApproachRadarTimeSlice":
		if precisionapproachradartimeslicepropertytype.PrecisionApproachRadarTimeSlice != nil {
			res = __gong__marshallPointer(ident, "PrecisionApproachRadarTimeSlice", precisionapproachradartimeslicepropertytype.PrecisionApproachRadarTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PrecisionApproachRadarTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PrecisionApproachRadarTimeSlicePropertyType", fieldName)
	}
	return
}

func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := precisionapproachradartimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", precisionapproachradartimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", precisionapproachradartimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", precisionapproachradartimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", precisionapproachradartimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", precisionapproachradartimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", precisionapproachradartimeslicetype.CorrectionNumber)

	case "Contact":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range precisionapproachradartimeslicetype.Contact {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Contact", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Location":
		if precisionapproachradartimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", precisionapproachradartimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range precisionapproachradartimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Reflector":
		var sb strings.Builder
		for _, _reflectorpropertytype := range precisionapproachradartimeslicetype.Reflector {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Reflector", _reflectorpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_35 := range precisionapproachradartimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_35.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PrecisionApproachRadarTimeSliceType", fieldName)
	}
	return
}

func (precisionapproachradartype *PrecisionApproachRadarType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := precisionapproachradartype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", precisionapproachradartype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", precisionapproachradartype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", precisionapproachradartype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _precisionapproachradartimeslicepropertytype := range precisionapproachradartype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _precisionapproachradartimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PrecisionApproachRadarType", fieldName)
	}
	return
}

func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := primarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", primarysurveillanceradartimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", primarysurveillanceradartimeslicepropertytype.Owns)

	case "PrimarySurveillanceRadarTimeSlice":
		if primarysurveillanceradartimeslicepropertytype.PrimarySurveillanceRadarTimeSlice != nil {
			res = __gong__marshallPointer(ident, "PrimarySurveillanceRadarTimeSlice", primarysurveillanceradartimeslicepropertytype.PrimarySurveillanceRadarTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PrimarySurveillanceRadarTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PrimarySurveillanceRadarTimeSlicePropertyType", fieldName)
	}
	return
}

func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := primarysurveillanceradartimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", primarysurveillanceradartimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", primarysurveillanceradartimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", primarysurveillanceradartimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", primarysurveillanceradartimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", primarysurveillanceradartimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", primarysurveillanceradartimeslicetype.CorrectionNumber)

	case "Contact":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range primarysurveillanceradartimeslicetype.Contact {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Contact", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Location":
		if primarysurveillanceradartimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", primarysurveillanceradartimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range primarysurveillanceradartimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundStation":
		var sb strings.Builder
		for _, _surveillancegroundstationpropertytype := range primarysurveillanceradartimeslicetype.GroundStation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundStation", _surveillancegroundstationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_36 := range primarysurveillanceradartimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_36.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PrimarySurveillanceRadarTimeSliceType", fieldName)
	}
	return
}

func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := primarysurveillanceradartype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", primarysurveillanceradartype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", primarysurveillanceradartype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", primarysurveillanceradartype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _primarysurveillanceradartimeslicepropertytype := range primarysurveillanceradartype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _primarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PrimarySurveillanceRadarType", fieldName)
	}
	return
}

func (radarcomponentpropertytype *RadarComponentPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radarcomponentpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radarcomponentpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", radarcomponentpropertytype.NilReason)

	case "RadarComponent":
		if radarcomponentpropertytype.RadarComponent != nil {
			res = __gong__marshallPointer(ident, "RadarComponent", radarcomponentpropertytype.RadarComponent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RadarComponent", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RadarComponentPropertyType", fieldName)
	}
	return
}

func (radarcomponenttype *RadarComponentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radarcomponenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radarcomponenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", radarcomponenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", radarcomponenttype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range radarcomponenttype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheRadarEquipment":
		if radarcomponenttype.TheRadarEquipment != nil {
			res = __gong__marshallPointer(ident, "TheRadarEquipment", radarcomponenttype.TheRadarEquipment.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheRadarEquipment", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_37 := range radarcomponenttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_37.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RadarComponentType", fieldName)
	}
	return
}

func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radarequipmentpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radarequipmentpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", radarequipmentpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", radarequipmentpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", radarequipmentpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", radarequipmentpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", radarequipmentpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", radarequipmentpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", radarequipmentpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", radarequipmentpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", radarequipmentpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct RadarEquipmentPropertyType", fieldName)
	}
	return
}

func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radarsystemtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radarsystemtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", radarsystemtimeslicepropertytype.Owns)

	case "RadarSystemTimeSlice":
		if radarsystemtimeslicepropertytype.RadarSystemTimeSlice != nil {
			res = __gong__marshallPointer(ident, "RadarSystemTimeSlice", radarsystemtimeslicepropertytype.RadarSystemTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RadarSystemTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RadarSystemTimeSlicePropertyType", fieldName)
	}
	return
}

func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radarsystemtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radarsystemtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", radarsystemtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", radarsystemtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", radarsystemtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", radarsystemtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", radarsystemtimeslicetype.CorrectionNumber)

	case "RadarEquipment":
		var sb strings.Builder
		for _, _radarcomponentpropertytype := range radarsystemtimeslicetype.RadarEquipment {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RadarEquipment", _radarcomponentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Office":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range radarsystemtimeslicetype.Office {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Office", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AirportHeliport":
		if radarsystemtimeslicetype.AirportHeliport != nil {
			res = __gong__marshallPointer(ident, "AirportHeliport", radarsystemtimeslicetype.AirportHeliport.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirportHeliport", "nil")
		}
	case "PARRunway":
		var sb strings.Builder
		for _, _runwaypropertytype := range radarsystemtimeslicetype.PARRunway {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PARRunway", _runwaypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Location":
		if radarsystemtimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", radarsystemtimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range radarsystemtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_38 := range radarsystemtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_38.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RadarSystemTimeSliceType", fieldName)
	}
	return
}

func (radarsystemtype *RadarSystemType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := radarsystemtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", radarsystemtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", radarsystemtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", radarsystemtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _radarsystemtimeslicepropertytype := range radarsystemtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _radarsystemtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RadarSystemType", fieldName)
	}
	return
}

func (reflectorpropertytype *ReflectorPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := reflectorpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", reflectorpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", reflectorpropertytype.NilReason)

	case "Reflector":
		if reflectorpropertytype.Reflector != nil {
			res = __gong__marshallPointer(ident, "Reflector", reflectorpropertytype.Reflector.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Reflector", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ReflectorPropertyType", fieldName)
	}
	return
}

func (reflectortype *ReflectorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := reflectortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", reflectortype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", reflectortype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", reflectortype.Description)

	case "TouchdownReflector":
		if reflectortype.TouchdownReflector != nil {
			res = __gong__marshallPointer(ident, "TouchdownReflector", reflectortype.TouchdownReflector.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TouchdownReflector", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range reflectortype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_39 := range reflectortype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_39.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ReflectorType", fieldName)
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

func (runwaypropertytype *RunwayPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := runwaypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", runwaypropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", runwaypropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", runwaypropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", runwaypropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", runwaypropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", runwaypropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", runwaypropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", runwaypropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", runwaypropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", runwaypropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct RunwayPropertyType", fieldName)
	}
	return
}

func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := sdftimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", sdftimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", sdftimeslicepropertytype.Owns)

	case "SDFTimeSlice":
		if sdftimeslicepropertytype.SDFTimeSlice != nil {
			res = __gong__marshallPointer(ident, "SDFTimeSlice", sdftimeslicepropertytype.SDFTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SDFTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SDFTimeSlicePropertyType", fieldName)
	}
	return
}

func (sdftimeslicetype *SDFTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := sdftimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", sdftimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", sdftimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", sdftimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", sdftimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", sdftimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", sdftimeslicetype.CorrectionNumber)

	case "Location":
		if sdftimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", sdftimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range sdftimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range sdftimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range sdftimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range sdftimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_18 := range sdftimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_18.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SDFTimeSliceType", fieldName)
	}
	return
}

func (sdftype *SDFType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := sdftype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", sdftype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", sdftype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", sdftype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _sdftimeslicepropertytype := range sdftype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _sdftimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SDFType", fieldName)
	}
	return
}

func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := secondarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", secondarysurveillanceradartimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", secondarysurveillanceradartimeslicepropertytype.Owns)

	case "SecondarySurveillanceRadarTimeSlice":
		if secondarysurveillanceradartimeslicepropertytype.SecondarySurveillanceRadarTimeSlice != nil {
			res = __gong__marshallPointer(ident, "SecondarySurveillanceRadarTimeSlice", secondarysurveillanceradartimeslicepropertytype.SecondarySurveillanceRadarTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SecondarySurveillanceRadarTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SecondarySurveillanceRadarTimeSlicePropertyType", fieldName)
	}
	return
}

func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := secondarysurveillanceradartimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", secondarysurveillanceradartimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", secondarysurveillanceradartimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", secondarysurveillanceradartimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", secondarysurveillanceradartimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", secondarysurveillanceradartimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", secondarysurveillanceradartimeslicetype.CorrectionNumber)

	case "Contact":
		var sb strings.Builder
		for _, _contactinformationpropertytype := range secondarysurveillanceradartimeslicetype.Contact {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Contact", _contactinformationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Location":
		if secondarysurveillanceradartimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", secondarysurveillanceradartimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range secondarysurveillanceradartimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroundStation":
		var sb strings.Builder
		for _, _surveillancegroundstationpropertytype := range secondarysurveillanceradartimeslicetype.GroundStation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroundStation", _surveillancegroundstationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_40 := range secondarysurveillanceradartimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_40.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SecondarySurveillanceRadarTimeSliceType", fieldName)
	}
	return
}

func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := secondarysurveillanceradartype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", secondarysurveillanceradartype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", secondarysurveillanceradartype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", secondarysurveillanceradartype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _secondarysurveillanceradartimeslicepropertytype := range secondarysurveillanceradartype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _secondarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SecondarySurveillanceRadarType", fieldName)
	}
	return
}

func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := significantpointinairspacetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", significantpointinairspacetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", significantpointinairspacetimeslicepropertytype.Owns)

	case "SignificantPointInAirspaceTimeSlice":
		if significantpointinairspacetimeslicepropertytype.SignificantPointInAirspaceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "SignificantPointInAirspaceTimeSlice", significantpointinairspacetimeslicepropertytype.SignificantPointInAirspaceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SignificantPointInAirspaceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SignificantPointInAirspaceTimeSlicePropertyType", fieldName)
	}
	return
}

func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := significantpointinairspacetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", significantpointinairspacetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", significantpointinairspacetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", significantpointinairspacetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", significantpointinairspacetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", significantpointinairspacetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", significantpointinairspacetimeslicetype.CorrectionNumber)

	case "ContainingAirspace":
		if significantpointinairspacetimeslicetype.ContainingAirspace != nil {
			res = __gong__marshallPointer(ident, "ContainingAirspace", significantpointinairspacetimeslicetype.ContainingAirspace.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ContainingAirspace", "nil")
		}
	case "Location_fixDesignatedPoint":
		if significantpointinairspacetimeslicetype.Location_fixDesignatedPoint != nil {
			res = __gong__marshallPointer(ident, "Location_fixDesignatedPoint", significantpointinairspacetimeslicetype.Location_fixDesignatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_fixDesignatedPoint", "nil")
		}
	case "Location_navaidSystem":
		if significantpointinairspacetimeslicetype.Location_navaidSystem != nil {
			res = __gong__marshallPointer(ident, "Location_navaidSystem", significantpointinairspacetimeslicetype.Location_navaidSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_navaidSystem", "nil")
		}
	case "Location_position":
		if significantpointinairspacetimeslicetype.Location_position != nil {
			res = __gong__marshallPointer(ident, "Location_position", significantpointinairspacetimeslicetype.Location_position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_position", "nil")
		}
	case "Location_runwayPoint":
		if significantpointinairspacetimeslicetype.Location_runwayPoint != nil {
			res = __gong__marshallPointer(ident, "Location_runwayPoint", significantpointinairspacetimeslicetype.Location_runwayPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_runwayPoint", "nil")
		}
	case "Location_aimingPoint":
		if significantpointinairspacetimeslicetype.Location_aimingPoint != nil {
			res = __gong__marshallPointer(ident, "Location_aimingPoint", significantpointinairspacetimeslicetype.Location_aimingPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_aimingPoint", "nil")
		}
	case "Location_airportReferencePoint":
		if significantpointinairspacetimeslicetype.Location_airportReferencePoint != nil {
			res = __gong__marshallPointer(ident, "Location_airportReferencePoint", significantpointinairspacetimeslicetype.Location_airportReferencePoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location_airportReferencePoint", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range significantpointinairspacetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_27 := range significantpointinairspacetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_27.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SignificantPointInAirspaceTimeSliceType", fieldName)
	}
	return
}

func (significantpointinairspacetype *SignificantPointInAirspaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := significantpointinairspacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", significantpointinairspacetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", significantpointinairspacetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", significantpointinairspacetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _significantpointinairspacetimeslicepropertytype := range significantpointinairspacetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _significantpointinairspacetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SignificantPointInAirspaceType", fieldName)
	}
	return
}

func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialnavigationstationstatuspropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialnavigationstationstatuspropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", specialnavigationstationstatuspropertytype.NilReason)

	case "SpecialNavigationStationStatus":
		if specialnavigationstationstatuspropertytype.SpecialNavigationStationStatus != nil {
			res = __gong__marshallPointer(ident, "SpecialNavigationStationStatus", specialnavigationstationstatuspropertytype.SpecialNavigationStationStatus.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SpecialNavigationStationStatus", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialNavigationStationStatusPropertyType", fieldName)
	}
	return
}

func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialnavigationstationstatustype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialnavigationstationstatustype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", specialnavigationstationstatustype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", specialnavigationstationstatustype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range specialnavigationstationstatustype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range specialnavigationstationstatustype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range specialnavigationstationstatustype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_20 := range specialnavigationstationstatustype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_20.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialNavigationStationStatusType", fieldName)
	}
	return
}

func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialnavigationstationtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialnavigationstationtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", specialnavigationstationtimeslicepropertytype.Owns)

	case "SpecialNavigationStationTimeSlice":
		if specialnavigationstationtimeslicepropertytype.SpecialNavigationStationTimeSlice != nil {
			res = __gong__marshallPointer(ident, "SpecialNavigationStationTimeSlice", specialnavigationstationtimeslicepropertytype.SpecialNavigationStationTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SpecialNavigationStationTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialNavigationStationTimeSlicePropertyType", fieldName)
	}
	return
}

func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialnavigationstationtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialnavigationstationtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", specialnavigationstationtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", specialnavigationstationtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", specialnavigationstationtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", specialnavigationstationtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", specialnavigationstationtimeslicetype.CorrectionNumber)

	case "SystemChain":
		if specialnavigationstationtimeslicetype.SystemChain != nil {
			res = __gong__marshallPointer(ident, "SystemChain", specialnavigationstationtimeslicetype.SystemChain.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SystemChain", "nil")
		}
	case "ResponsibleOrganisation":
		if specialnavigationstationtimeslicetype.ResponsibleOrganisation != nil {
			res = __gong__marshallPointer(ident, "ResponsibleOrganisation", specialnavigationstationtimeslicetype.ResponsibleOrganisation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ResponsibleOrganisation", "nil")
		}
	case "Position":
		if specialnavigationstationtimeslicetype.Position != nil {
			res = __gong__marshallPointer(ident, "Position", specialnavigationstationtimeslicetype.Position.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Position", "nil")
		}
	case "Availability":
		var sb strings.Builder
		for _, _specialnavigationstationstatuspropertytype := range specialnavigationstationtimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _specialnavigationstationstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range specialnavigationstationtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_19 := range specialnavigationstationtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_19.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialNavigationStationTimeSliceType", fieldName)
	}
	return
}

func (specialnavigationstationtype *SpecialNavigationStationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialnavigationstationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialnavigationstationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", specialnavigationstationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", specialnavigationstationtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _specialnavigationstationtimeslicepropertytype := range specialnavigationstationtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _specialnavigationstationtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialNavigationStationType", fieldName)
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

func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialnavigationsystemtimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialnavigationsystemtimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", specialnavigationsystemtimeslicepropertytype.Owns)

	case "SpecialNavigationSystemTimeSlice":
		if specialnavigationsystemtimeslicepropertytype.SpecialNavigationSystemTimeSlice != nil {
			res = __gong__marshallPointer(ident, "SpecialNavigationSystemTimeSlice", specialnavigationsystemtimeslicepropertytype.SpecialNavigationSystemTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SpecialNavigationSystemTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialNavigationSystemTimeSlicePropertyType", fieldName)
	}
	return
}

func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialnavigationsystemtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialnavigationsystemtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", specialnavigationsystemtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", specialnavigationsystemtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", specialnavigationsystemtimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", specialnavigationsystemtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", specialnavigationsystemtimeslicetype.CorrectionNumber)

	case "ResponsibleOrganisation":
		if specialnavigationsystemtimeslicetype.ResponsibleOrganisation != nil {
			res = __gong__marshallPointer(ident, "ResponsibleOrganisation", specialnavigationsystemtimeslicetype.ResponsibleOrganisation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ResponsibleOrganisation", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range specialnavigationsystemtimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_21 := range specialnavigationsystemtimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_21.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialNavigationSystemTimeSliceType", fieldName)
	}
	return
}

func (specialnavigationsystemtype *SpecialNavigationSystemType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialnavigationsystemtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialnavigationsystemtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", specialnavigationsystemtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", specialnavigationsystemtype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _specialnavigationsystemtimeslicepropertytype := range specialnavigationsystemtype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _specialnavigationsystemtimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialNavigationSystemType", fieldName)
	}
	return
}

func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := surveillancegroundstationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", surveillancegroundstationpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", surveillancegroundstationpropertytype.NilReason)

	case "SurveillanceGroundStation":
		if surveillancegroundstationpropertytype.SurveillanceGroundStation != nil {
			res = __gong__marshallPointer(ident, "SurveillanceGroundStation", surveillancegroundstationpropertytype.SurveillanceGroundStation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SurveillanceGroundStation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SurveillanceGroundStationPropertyType", fieldName)
	}
	return
}

func (surveillancegroundstationtype *SurveillanceGroundStationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := surveillancegroundstationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", surveillancegroundstationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", surveillancegroundstationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", surveillancegroundstationtype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range surveillancegroundstationtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheUnit":
		if surveillancegroundstationtype.TheUnit != nil {
			res = __gong__marshallPointer(ident, "TheUnit", surveillancegroundstationtype.TheUnit.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheUnit", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_41 := range surveillancegroundstationtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_41.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SurveillanceGroundStationType", fieldName)
	}
	return
}

func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tacantimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tacantimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", tacantimeslicepropertytype.Owns)

	case "TACANTimeSlice":
		if tacantimeslicepropertytype.TACANTimeSlice != nil {
			res = __gong__marshallPointer(ident, "TACANTimeSlice", tacantimeslicepropertytype.TACANTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TACANTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TACANTimeSlicePropertyType", fieldName)
	}
	return
}

func (tacantimeslicetype *TACANTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tacantimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tacantimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", tacantimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", tacantimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", tacantimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", tacantimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", tacantimeslicetype.CorrectionNumber)

	case "Location":
		if tacantimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", tacantimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range tacantimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range tacantimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range tacantimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range tacantimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_22 := range tacantimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_22.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TACANTimeSliceType", fieldName)
	}
	return
}

func (tacantype *TACANType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tacantype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tacantype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", tacantype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", tacantype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _tacantimeslicepropertytype := range tacantype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _tacantimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TACANType", fieldName)
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

func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vortimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vortimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", vortimeslicepropertytype.Owns)

	case "VORTimeSlice":
		if vortimeslicepropertytype.VORTimeSlice != nil {
			res = __gong__marshallPointer(ident, "VORTimeSlice", vortimeslicepropertytype.VORTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VORTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct VORTimeSlicePropertyType", fieldName)
	}
	return
}

func (vortimeslicetype *VORTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vortimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vortimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", vortimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", vortimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", vortimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", vortimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", vortimeslicetype.CorrectionNumber)

	case "Location":
		if vortimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", vortimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Authority":
		var sb strings.Builder
		for _, _authorityfornavaidequipmentpropertytype := range vortimeslicetype.Authority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Authority", _authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Monitoring":
		var sb strings.Builder
		for _, _navaidequipmentmonitoringpropertytype := range vortimeslicetype.Monitoring {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Monitoring", _navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _navaidoperationalstatuspropertytype := range vortimeslicetype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _navaidoperationalstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range vortimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_23 := range vortimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_23.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct VORTimeSliceType", fieldName)
	}
	return
}

func (vortype *VORType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vortype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", vortype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", vortype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _vortimeslicepropertytype := range vortype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _vortimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct VORType", fieldName)
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
func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractnavaidequipmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractnavaidequipmenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractnavaidequipmenttype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractnavigationsystemcheckpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractnavigationsystemcheckpointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractnavigationsystemcheckpointtype.GongMarshallField(stage, "Description"))
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
func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractradarequipmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractradarequipmenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractradarequipmenttype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractsurveillanceradartype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractsurveillanceradartype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractsurveillanceradartype.GongMarshallField(stage, "Description"))
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
func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(angleindicationtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(angleindicationtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(angleindicationtimeslicepropertytype.GongMarshallField(stage, "AngleIndicationTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "Fix"))
		pointersInitializesStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "PointChoice_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "PointChoice_navaidSystem"))
		pointersInitializesStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "PointChoice_position"))
		pointersInitializesStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "PointChoice_runwayPoint"))
		pointersInitializesStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "PointChoice_aimingPoint"))
		pointersInitializesStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "PointChoice_airportReferencePoint"))
		pointersInitializesStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(angleindicationtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (angleindicationtype *AngleIndicationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(angleindicationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(angleindicationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(angleindicationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(angleindicationtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityfornavaidequipmentpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityfornavaidequipmentpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(authorityfornavaidequipmentpropertytype.GongMarshallField(stage, "AuthorityForNavaidEquipment"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityfornavaidequipmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityfornavaidequipmenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(authorityfornavaidequipmenttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(authorityfornavaidequipmenttype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(authorityfornavaidequipmenttype.GongMarshallField(stage, "TheOrganisationAuthority"))
		pointersInitializesStatements.WriteString(authorityfornavaidequipmenttype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityforspecialnavigationstationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityforspecialnavigationstationpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(authorityforspecialnavigationstationpropertytype.GongMarshallField(stage, "AuthorityForSpecialNavigationStation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityforspecialnavigationstationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityforspecialnavigationstationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(authorityforspecialnavigationstationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(authorityforspecialnavigationstationtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(authorityforspecialnavigationstationtype.GongMarshallField(stage, "TheOrganisationAuthority"))
		pointersInitializesStatements.WriteString(authorityforspecialnavigationstationtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityforspecialnavigationsystempropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityforspecialnavigationsystempropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(authorityforspecialnavigationsystempropertytype.GongMarshallField(stage, "AuthorityForSpecialNavigationSystem"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityforspecialnavigationsystemtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityforspecialnavigationsystemtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(authorityforspecialnavigationsystemtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(authorityforspecialnavigationsystemtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(authorityforspecialnavigationsystemtype.GongMarshallField(stage, "TheOrganisationAuthority"))
		pointersInitializesStatements.WriteString(authorityforspecialnavigationsystemtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(azimuthtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(azimuthtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(azimuthtimeslicepropertytype.GongMarshallField(stage, "AzimuthTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (azimuthtimeslicetype *AzimuthTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(azimuthtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (azimuthtype *AzimuthType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(azimuthtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(azimuthtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(azimuthtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(azimuthtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(checkpointinstimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(checkpointinstimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(checkpointinstimeslicepropertytype.GongMarshallField(stage, "CheckpointINSTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(checkpointinstimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(checkpointinstimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(checkpointinstimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(checkpointinstimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(checkpointinstimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(checkpointinstimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(checkpointinstimeslicetype.GongMarshallField(stage, "Position"))
		pointersInitializesStatements.WriteString(checkpointinstimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(checkpointinstimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(checkpointinstimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (checkpointinstype *CheckpointINSType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(checkpointinstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(checkpointinstype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(checkpointinstype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(checkpointinstype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(checkpointvortimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(checkpointvortimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(checkpointvortimeslicepropertytype.GongMarshallField(stage, "CheckpointVORTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "Position"))
		pointersInitializesStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "CheckPointFacility"))
		pointersInitializesStatements.WriteString(checkpointvortimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (checkpointvortype *CheckpointVORType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(checkpointvortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(checkpointvortype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(checkpointvortype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(checkpointvortype.GongMarshallField(stage, "TimeSlice"))
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
func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dmetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(dmetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(dmetimeslicepropertytype.GongMarshallField(stage, "DMETimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dmetimeslicetype *DMETimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(dmetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dmetype *DMEType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dmetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(dmetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(dmetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(dmetype.GongMarshallField(stage, "TimeSlice"))
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
func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(designatedpointtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(designatedpointtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(designatedpointtimeslicepropertytype.GongMarshallField(stage, "DesignatedPointTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "AimingPoint"))
		pointersInitializesStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "RunwayPoint"))
		pointersInitializesStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(designatedpointtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (designatedpointtype *DesignatedPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(designatedpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(designatedpointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(designatedpointtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(designatedpointtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directionfindertimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directionfindertimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(directionfindertimeslicepropertytype.GongMarshallField(stage, "DirectionFinderTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "InformationProvision"))
		pointersInitializesStatements.WriteString(directionfindertimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (directionfindertype *DirectionFinderType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directionfindertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directionfindertype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(directionfindertype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(directionfindertype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(distanceindicationtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(distanceindicationtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(distanceindicationtimeslicepropertytype.GongMarshallField(stage, "DistanceIndicationTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "Fix"))
		pointersInitializesStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "PointChoice_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "PointChoice_navaidSystem"))
		pointersInitializesStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "PointChoice_position"))
		pointersInitializesStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "PointChoice_runwayPoint"))
		pointersInitializesStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "PointChoice_aimingPoint"))
		pointersInitializesStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "PointChoice_airportReferencePoint"))
		pointersInitializesStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(distanceindicationtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (distanceindicationtype *DistanceIndicationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(distanceindicationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(distanceindicationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(distanceindicationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(distanceindicationtype.GongMarshallField(stage, "TimeSlice"))
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
func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevationtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevationtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(elevationtimeslicepropertytype.GongMarshallField(stage, "ElevationTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevationtimeslicetype *ElevationTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(elevationtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevationtype *ElevationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(elevationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(elevationtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(glidepathtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(glidepathtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(glidepathtimeslicepropertytype.GongMarshallField(stage, "GlidepathTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (glidepathtimeslicetype *GlidepathTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(glidepathtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (glidepathtype *GlidepathType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(glidepathtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(glidepathtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(glidepathtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(glidepathtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (informationservicepropertytype *InformationServicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(informationservicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(informationservicepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(informationservicepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(informationservicepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(informationservicepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(informationservicepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(informationservicepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(informationservicepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(informationservicepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(informationservicepropertytype.GongMarshallField(stage, "Actuate"))
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
func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(localizertimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(localizertimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(localizertimeslicepropertytype.GongMarshallField(stage, "LocalizerTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (localizertimeslicetype *LocalizerTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(localizertimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (localizertype *LocalizerType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(localizertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(localizertype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(localizertype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(localizertype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(markerbeacontimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(markerbeacontimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(markerbeacontimeslicepropertytype.GongMarshallField(stage, "MarkerBeaconTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(markerbeacontimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (markerbeacontype *MarkerBeaconType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(markerbeacontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(markerbeacontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(markerbeacontype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(markerbeacontype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ndbtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ndbtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(ndbtimeslicepropertytype.GongMarshallField(stage, "NDBTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ndbtimeslicetype *NDBTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(ndbtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ndbtype *NDBType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ndbtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ndbtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(ndbtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(ndbtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidcomponentpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidcomponentpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(navaidcomponentpropertytype.GongMarshallField(stage, "NavaidComponent"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidcomponenttype *NavaidComponentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidcomponenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidcomponenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(navaidcomponenttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(navaidcomponenttype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(navaidcomponenttype.GongMarshallField(stage, "TheNavaidEquipment"))
		pointersInitializesStatements.WriteString(navaidcomponenttype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidequipmentmonitoringpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidequipmentmonitoringpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(navaidequipmentmonitoringpropertytype.GongMarshallField(stage, "NavaidEquipmentMonitoring"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidequipmentmonitoringtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidequipmentmonitoringtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(navaidequipmentmonitoringtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(navaidequipmentmonitoringtype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(navaidequipmentmonitoringtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(navaidequipmentmonitoringtype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(navaidequipmentmonitoringtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidoperationalstatuspropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidoperationalstatuspropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(navaidoperationalstatuspropertytype.GongMarshallField(stage, "NavaidOperationalStatus"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidoperationalstatustype *NavaidOperationalStatusType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidoperationalstatustype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidoperationalstatustype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(navaidoperationalstatustype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(navaidoperationalstatustype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(navaidoperationalstatustype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(navaidoperationalstatustype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(navaidoperationalstatustype.GongMarshallField(stage, "Extension"))
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
func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(navaidtimeslicepropertytype.GongMarshallField(stage, "NavaidTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidtimeslicetype *NavaidTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "TouchDownLiftOff"))
		pointersInitializesStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "NavaidEquipment"))
		pointersInitializesStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "RunwayDirection"))
		pointersInitializesStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "ServedAirport"))
		pointersInitializesStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(navaidtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidtype *NavaidType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(navaidtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(navaidtype.GongMarshallField(stage, "TimeSlice"))
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
func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(precisionapproachradartimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(precisionapproachradartimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(precisionapproachradartimeslicepropertytype.GongMarshallField(stage, "PrecisionApproachRadarTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "Contact"))
		pointersInitializesStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "Reflector"))
		pointersInitializesStatements.WriteString(precisionapproachradartimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (precisionapproachradartype *PrecisionApproachRadarType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(precisionapproachradartype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(precisionapproachradartype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(precisionapproachradartype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(precisionapproachradartype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(primarysurveillanceradartimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(primarysurveillanceradartimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(primarysurveillanceradartimeslicepropertytype.GongMarshallField(stage, "PrimarySurveillanceRadarTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "Contact"))
		pointersInitializesStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "GroundStation"))
		pointersInitializesStatements.WriteString(primarysurveillanceradartimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(primarysurveillanceradartype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(primarysurveillanceradartype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(primarysurveillanceradartype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(primarysurveillanceradartype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radarcomponentpropertytype *RadarComponentPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radarcomponentpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radarcomponentpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(radarcomponentpropertytype.GongMarshallField(stage, "RadarComponent"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radarcomponenttype *RadarComponentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radarcomponenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radarcomponenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(radarcomponenttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(radarcomponenttype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(radarcomponenttype.GongMarshallField(stage, "TheRadarEquipment"))
		pointersInitializesStatements.WriteString(radarcomponenttype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radarequipmentpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radarequipmentpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(radarequipmentpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(radarequipmentpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(radarequipmentpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(radarequipmentpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(radarequipmentpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(radarequipmentpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(radarequipmentpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(radarequipmentpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radarsystemtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radarsystemtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(radarsystemtimeslicepropertytype.GongMarshallField(stage, "RadarSystemTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "RadarEquipment"))
		pointersInitializesStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "Office"))
		pointersInitializesStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "AirportHeliport"))
		pointersInitializesStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "PARRunway"))
		pointersInitializesStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(radarsystemtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (radarsystemtype *RadarSystemType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(radarsystemtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(radarsystemtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(radarsystemtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(radarsystemtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (reflectorpropertytype *ReflectorPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(reflectorpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(reflectorpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(reflectorpropertytype.GongMarshallField(stage, "Reflector"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (reflectortype *ReflectorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(reflectortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(reflectortype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(reflectortype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(reflectortype.GongMarshallField(stage, "TouchdownReflector"))
		pointersInitializesStatements.WriteString(reflectortype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(reflectortype.GongMarshallField(stage, "Extension"))
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
func (runwaypropertytype *RunwayPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(runwaypropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(runwaypropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(runwaypropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(runwaypropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(runwaypropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(runwaypropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(runwaypropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(runwaypropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(runwaypropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(runwaypropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(sdftimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(sdftimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(sdftimeslicepropertytype.GongMarshallField(stage, "SDFTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (sdftimeslicetype *SDFTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(sdftimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (sdftype *SDFType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(sdftype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(sdftype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(sdftype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(sdftype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(secondarysurveillanceradartimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(secondarysurveillanceradartimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(secondarysurveillanceradartimeslicepropertytype.GongMarshallField(stage, "SecondarySurveillanceRadarTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "Contact"))
		pointersInitializesStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "GroundStation"))
		pointersInitializesStatements.WriteString(secondarysurveillanceradartimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(secondarysurveillanceradartype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(secondarysurveillanceradartype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(secondarysurveillanceradartype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(secondarysurveillanceradartype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(significantpointinairspacetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(significantpointinairspacetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(significantpointinairspacetimeslicepropertytype.GongMarshallField(stage, "SignificantPointInAirspaceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "ContainingAirspace"))
		pointersInitializesStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Location_fixDesignatedPoint"))
		pointersInitializesStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Location_navaidSystem"))
		pointersInitializesStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Location_position"))
		pointersInitializesStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Location_runwayPoint"))
		pointersInitializesStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Location_aimingPoint"))
		pointersInitializesStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Location_airportReferencePoint"))
		pointersInitializesStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(significantpointinairspacetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (significantpointinairspacetype *SignificantPointInAirspaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(significantpointinairspacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(significantpointinairspacetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(significantpointinairspacetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(significantpointinairspacetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialnavigationstationstatuspropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialnavigationstationstatuspropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(specialnavigationstationstatuspropertytype.GongMarshallField(stage, "SpecialNavigationStationStatus"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialnavigationstationstatustype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialnavigationstationstatustype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(specialnavigationstationstatustype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(specialnavigationstationstatustype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(specialnavigationstationstatustype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(specialnavigationstationstatustype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(specialnavigationstationstatustype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialnavigationstationtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialnavigationstationtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(specialnavigationstationtimeslicepropertytype.GongMarshallField(stage, "SpecialNavigationStationTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "SystemChain"))
		pointersInitializesStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "ResponsibleOrganisation"))
		pointersInitializesStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "Position"))
		pointersInitializesStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(specialnavigationstationtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialnavigationstationtype *SpecialNavigationStationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialnavigationstationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialnavigationstationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(specialnavigationstationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(specialnavigationstationtype.GongMarshallField(stage, "TimeSlice"))
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
func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialnavigationsystemtimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialnavigationsystemtimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(specialnavigationsystemtimeslicepropertytype.GongMarshallField(stage, "SpecialNavigationSystemTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialnavigationsystemtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialnavigationsystemtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(specialnavigationsystemtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(specialnavigationsystemtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(specialnavigationsystemtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(specialnavigationsystemtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(specialnavigationsystemtimeslicetype.GongMarshallField(stage, "ResponsibleOrganisation"))
		pointersInitializesStatements.WriteString(specialnavigationsystemtimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(specialnavigationsystemtimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialnavigationsystemtype *SpecialNavigationSystemType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialnavigationsystemtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialnavigationsystemtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(specialnavigationsystemtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(specialnavigationsystemtype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(surveillancegroundstationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(surveillancegroundstationpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(surveillancegroundstationpropertytype.GongMarshallField(stage, "SurveillanceGroundStation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (surveillancegroundstationtype *SurveillanceGroundStationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(surveillancegroundstationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(surveillancegroundstationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(surveillancegroundstationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(surveillancegroundstationtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(surveillancegroundstationtype.GongMarshallField(stage, "TheUnit"))
		pointersInitializesStatements.WriteString(surveillancegroundstationtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tacantimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tacantimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(tacantimeslicepropertytype.GongMarshallField(stage, "TACANTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tacantimeslicetype *TACANTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(tacantimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tacantype *TACANType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tacantype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tacantype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(tacantype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(tacantype.GongMarshallField(stage, "TimeSlice"))
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
func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vortimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(vortimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(vortimeslicepropertytype.GongMarshallField(stage, "VORTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (vortimeslicetype *VORTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "Authority"))
		pointersInitializesStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "Monitoring"))
		pointersInitializesStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(vortimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (vortype *VORType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(vortype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(vortype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(vortype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
