// generated code - do not edit
package aixmdatatypes

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
	{{Identifier}} := (&aixmdatatypes.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *aixmdatatypes.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *aixmdatatypes.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes"
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
		for _, codeatcreportingtype := range __gong__sortStageSetInstances(stage.CodeATCReportingTypes, stage.CodeATCReportingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeatcreportingtypeIdent := codeatcreportingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeATCReportingType{Name: %s}).Stage(stage)", codeatcreportingtypeIdent, __gong__toRawStringLiteral(codeatcreportingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeatcreportingtypeIdent, __gong__toRawStringLiteral(codeatcreportingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeatcreportingtypeIdent, __gong__toRawStringLiteral(codeatcreportingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeatcreportingtypeIdent, __gong__toRawStringLiteral(codeatcreportingtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeaerialrefuellingpointtype := range __gong__sortStageSetInstances(stage.CodeAerialRefuellingPointTypes, stage.CodeAerialRefuellingPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaerialrefuellingpointtypeIdent := codeaerialrefuellingpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAerialRefuellingPointType{Name: %s}).Stage(stage)", codeaerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingpointtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingpointtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeaerialrefuellingprefixtype := range __gong__sortStageSetInstances(stage.CodeAerialRefuellingPrefixTypes, stage.CodeAerialRefuellingPrefixType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaerialrefuellingprefixtypeIdent := codeaerialrefuellingprefixtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAerialRefuellingPrefixType{Name: %s}).Stage(stage)", codeaerialrefuellingprefixtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingprefixtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaerialrefuellingprefixtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingprefixtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaerialrefuellingprefixtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingprefixtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaerialrefuellingprefixtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingprefixtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeaerialrefuellingtype := range __gong__sortStageSetInstances(stage.CodeAerialRefuellingTypes, stage.CodeAerialRefuellingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaerialrefuellingtypeIdent := codeaerialrefuellingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAerialRefuellingType{Name: %s}).Stage(stage)", codeaerialrefuellingtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaerialrefuellingtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaerialrefuellingtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaerialrefuellingtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeaircraftcategorytype := range __gong__sortStageSetInstances(stage.CodeAircraftCategoryTypes, stage.CodeAircraftCategoryType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaircraftcategorytypeIdent := codeaircraftcategorytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftCategoryType{Name: %s}).Stage(stage)", codeaircraftcategorytypeIdent, __gong__toRawStringLiteral(codeaircraftcategorytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftcategorytypeIdent, __gong__toRawStringLiteral(codeaircraftcategorytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftcategorytypeIdent, __gong__toRawStringLiteral(codeaircraftcategorytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftcategorytypeIdent, __gong__toRawStringLiteral(codeaircraftcategorytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeaircraftenginenumbertype := range __gong__sortStageSetInstances(stage.CodeAircraftEngineNumberTypes, stage.CodeAircraftEngineNumberType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaircraftenginenumbertypeIdent := codeaircraftenginenumbertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftEngineNumberType{Name: %s}).Stage(stage)", codeaircraftenginenumbertypeIdent, __gong__toRawStringLiteral(codeaircraftenginenumbertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftenginenumbertypeIdent, __gong__toRawStringLiteral(codeaircraftenginenumbertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftenginenumbertypeIdent, __gong__toRawStringLiteral(codeaircraftenginenumbertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftenginenumbertypeIdent, __gong__toRawStringLiteral(codeaircraftenginenumbertype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeaircraftenginetype := range __gong__sortStageSetInstances(stage.CodeAircraftEngineTypes, stage.CodeAircraftEngineType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaircraftenginetypeIdent := codeaircraftenginetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftEngineType{Name: %s}).Stage(stage)", codeaircraftenginetypeIdent, __gong__toRawStringLiteral(codeaircraftenginetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftenginetypeIdent, __gong__toRawStringLiteral(codeaircraftenginetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftenginetypeIdent, __gong__toRawStringLiteral(codeaircraftenginetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftenginetypeIdent, __gong__toRawStringLiteral(codeaircraftenginetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeaircrafticaotype := range __gong__sortStageSetInstances(stage.CodeAircraftICAOTypes, stage.CodeAircraftICAOType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaircrafticaotypeIdent := codeaircrafticaotype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftICAOType{Name: %s}).Stage(stage)", codeaircrafticaotypeIdent, __gong__toRawStringLiteral(codeaircrafticaotype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircrafticaotypeIdent, __gong__toRawStringLiteral(codeaircrafticaotype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircrafticaotypeIdent, __gong__toRawStringLiteral(codeaircrafticaotype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircrafticaotypeIdent, __gong__toRawStringLiteral(codeaircrafticaotype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeaircrafttype := range __gong__sortStageSetInstances(stage.CodeAircraftTypes, stage.CodeAircraftType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaircrafttypeIdent := codeaircrafttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftType{Name: %s}).Stage(stage)", codeaircrafttypeIdent, __gong__toRawStringLiteral(codeaircrafttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircrafttypeIdent, __gong__toRawStringLiteral(codeaircrafttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircrafttypeIdent, __gong__toRawStringLiteral(codeaircrafttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircrafttypeIdent, __gong__toRawStringLiteral(codeaircrafttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeaircraftwingspanclasstype := range __gong__sortStageSetInstances(stage.CodeAircraftWingspanClassTypes, stage.CodeAircraftWingspanClassType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaircraftwingspanclasstypeIdent := codeaircraftwingspanclasstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftWingspanClassType{Name: %s}).Stage(stage)", codeaircraftwingspanclasstypeIdent, __gong__toRawStringLiteral(codeaircraftwingspanclasstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftwingspanclasstypeIdent, __gong__toRawStringLiteral(codeaircraftwingspanclasstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftwingspanclasstypeIdent, __gong__toRawStringLiteral(codeaircraftwingspanclasstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftwingspanclasstypeIdent, __gong__toRawStringLiteral(codeaircraftwingspanclasstype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codealtitudeadjustmenttype := range __gong__sortStageSetInstances(stage.CodeAltitudeAdjustmentTypes, stage.CodeAltitudeAdjustmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codealtitudeadjustmenttypeIdent := codealtitudeadjustmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAltitudeAdjustmentType{Name: %s}).Stage(stage)", codealtitudeadjustmenttypeIdent, __gong__toRawStringLiteral(codealtitudeadjustmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codealtitudeadjustmenttypeIdent, __gong__toRawStringLiteral(codealtitudeadjustmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codealtitudeadjustmenttypeIdent, __gong__toRawStringLiteral(codealtitudeadjustmenttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codealtitudeadjustmenttypeIdent, __gong__toRawStringLiteral(codealtitudeadjustmenttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codealtitudeusetype := range __gong__sortStageSetInstances(stage.CodeAltitudeUseTypes, stage.CodeAltitudeUseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codealtitudeusetypeIdent := codealtitudeusetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAltitudeUseType{Name: %s}).Stage(stage)", codealtitudeusetypeIdent, __gong__toRawStringLiteral(codealtitudeusetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codealtitudeusetypeIdent, __gong__toRawStringLiteral(codealtitudeusetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codealtitudeusetypeIdent, __gong__toRawStringLiteral(codealtitudeusetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codealtitudeusetypeIdent, __gong__toRawStringLiteral(codealtitudeusetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeauthoritytype := range __gong__sortStageSetInstances(stage.CodeAuthorityTypes, stage.CodeAuthorityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeauthoritytypeIdent := codeauthoritytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAuthorityType{Name: %s}).Stage(stage)", codeauthoritytypeIdent, __gong__toRawStringLiteral(codeauthoritytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeauthoritytypeIdent, __gong__toRawStringLiteral(codeauthoritytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeauthoritytypeIdent, __gong__toRawStringLiteral(codeauthoritytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeauthoritytypeIdent, __gong__toRawStringLiteral(codeauthoritytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codecardinaldirectiontype := range __gong__sortStageSetInstances(stage.CodeCardinalDirectionTypes, stage.CodeCardinalDirectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codecardinaldirectiontypeIdent := codecardinaldirectiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCardinalDirectionType{Name: %s}).Stage(stage)", codecardinaldirectiontypeIdent, __gong__toRawStringLiteral(codecardinaldirectiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecardinaldirectiontypeIdent, __gong__toRawStringLiteral(codecardinaldirectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecardinaldirectiontypeIdent, __gong__toRawStringLiteral(codecardinaldirectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecardinaldirectiontypeIdent, __gong__toRawStringLiteral(codecardinaldirectiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codecommunicationmodetype := range __gong__sortStageSetInstances(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codecommunicationmodetypeIdent := codecommunicationmodetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCommunicationModeType{Name: %s}).Stage(stage)", codecommunicationmodetypeIdent, __gong__toRawStringLiteral(codecommunicationmodetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecommunicationmodetypeIdent, __gong__toRawStringLiteral(codecommunicationmodetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecommunicationmodetypeIdent, __gong__toRawStringLiteral(codecommunicationmodetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecommunicationmodetypeIdent, __gong__toRawStringLiteral(codecommunicationmodetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codecomparisontype := range __gong__sortStageSetInstances(stage.CodeComparisonTypes, stage.CodeComparisonType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codecomparisontypeIdent := codecomparisontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeComparisonType{Name: %s}).Stage(stage)", codecomparisontypeIdent, __gong__toRawStringLiteral(codecomparisontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecomparisontypeIdent, __gong__toRawStringLiteral(codecomparisontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecomparisontypeIdent, __gong__toRawStringLiteral(codecomparisontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecomparisontypeIdent, __gong__toRawStringLiteral(codecomparisontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codecoursetype := range __gong__sortStageSetInstances(stage.CodeCourseTypes, stage.CodeCourseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codecoursetypeIdent := codecoursetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCourseType{Name: %s}).Stage(stage)", codecoursetypeIdent, __gong__toRawStringLiteral(codecoursetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecoursetypeIdent, __gong__toRawStringLiteral(codecoursetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecoursetypeIdent, __gong__toRawStringLiteral(codecoursetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecoursetypeIdent, __gong__toRawStringLiteral(codecoursetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codedaytype := range __gong__sortStageSetInstances(stage.CodeDayTypes, stage.CodeDayType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedaytypeIdent := codedaytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDayType{Name: %s}).Stage(stage)", codedaytypeIdent, __gong__toRawStringLiteral(codedaytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedaytypeIdent, __gong__toRawStringLiteral(codedaytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedaytypeIdent, __gong__toRawStringLiteral(codedaytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedaytypeIdent, __gong__toRawStringLiteral(codedaytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codedirectionturntype := range __gong__sortStageSetInstances(stage.CodeDirectionTurnTypes, stage.CodeDirectionTurnType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedirectionturntypeIdent := codedirectionturntype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDirectionTurnType{Name: %s}).Stage(stage)", codedirectionturntypeIdent, __gong__toRawStringLiteral(codedirectionturntype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedirectionturntypeIdent, __gong__toRawStringLiteral(codedirectionturntype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedirectionturntypeIdent, __gong__toRawStringLiteral(codedirectionturntype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedirectionturntypeIdent, __gong__toRawStringLiteral(codedirectionturntype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codedirectiontype := range __gong__sortStageSetInstances(stage.CodeDirectionTypes, stage.CodeDirectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedirectiontypeIdent := codedirectiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDirectionType{Name: %s}).Stage(stage)", codedirectiontypeIdent, __gong__toRawStringLiteral(codedirectiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedirectiontypeIdent, __gong__toRawStringLiteral(codedirectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedirectiontypeIdent, __gong__toRawStringLiteral(codedirectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedirectiontypeIdent, __gong__toRawStringLiteral(codedirectiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeequipmentanticollisiontype := range __gong__sortStageSetInstances(stage.CodeEquipmentAntiCollisionTypes, stage.CodeEquipmentAntiCollisionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeequipmentanticollisiontypeIdent := codeequipmentanticollisiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeEquipmentAntiCollisionType{Name: %s}).Stage(stage)", codeequipmentanticollisiontypeIdent, __gong__toRawStringLiteral(codeequipmentanticollisiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeequipmentanticollisiontypeIdent, __gong__toRawStringLiteral(codeequipmentanticollisiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeequipmentanticollisiontypeIdent, __gong__toRawStringLiteral(codeequipmentanticollisiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeequipmentanticollisiontypeIdent, __gong__toRawStringLiteral(codeequipmentanticollisiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeflightorigintype := range __gong__sortStageSetInstances(stage.CodeFlightOriginTypes, stage.CodeFlightOriginType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeflightorigintypeIdent := codeflightorigintype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightOriginType{Name: %s}).Stage(stage)", codeflightorigintypeIdent, __gong__toRawStringLiteral(codeflightorigintype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightorigintypeIdent, __gong__toRawStringLiteral(codeflightorigintype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightorigintypeIdent, __gong__toRawStringLiteral(codeflightorigintype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightorigintypeIdent, __gong__toRawStringLiteral(codeflightorigintype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeflightpurposetype := range __gong__sortStageSetInstances(stage.CodeFlightPurposeTypes, stage.CodeFlightPurposeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeflightpurposetypeIdent := codeflightpurposetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightPurposeType{Name: %s}).Stage(stage)", codeflightpurposetypeIdent, __gong__toRawStringLiteral(codeflightpurposetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightpurposetypeIdent, __gong__toRawStringLiteral(codeflightpurposetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightpurposetypeIdent, __gong__toRawStringLiteral(codeflightpurposetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightpurposetypeIdent, __gong__toRawStringLiteral(codeflightpurposetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeflightrestrictiondesignatortype := range __gong__sortStageSetInstances(stage.CodeFlightRestrictionDesignatorTypes, stage.CodeFlightRestrictionDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeflightrestrictiondesignatortypeIdent := codeflightrestrictiondesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightRestrictionDesignatorType{Name: %s}).Stage(stage)", codeflightrestrictiondesignatortypeIdent, __gong__toRawStringLiteral(codeflightrestrictiondesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightrestrictiondesignatortypeIdent, __gong__toRawStringLiteral(codeflightrestrictiondesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightrestrictiondesignatortypeIdent, __gong__toRawStringLiteral(codeflightrestrictiondesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightrestrictiondesignatortypeIdent, __gong__toRawStringLiteral(codeflightrestrictiondesignatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeflightrestrictiontype := range __gong__sortStageSetInstances(stage.CodeFlightRestrictionTypes, stage.CodeFlightRestrictionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeflightrestrictiontypeIdent := codeflightrestrictiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightRestrictionType{Name: %s}).Stage(stage)", codeflightrestrictiontypeIdent, __gong__toRawStringLiteral(codeflightrestrictiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightrestrictiontypeIdent, __gong__toRawStringLiteral(codeflightrestrictiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightrestrictiontypeIdent, __gong__toRawStringLiteral(codeflightrestrictiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightrestrictiontypeIdent, __gong__toRawStringLiteral(codeflightrestrictiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeflightruletype := range __gong__sortStageSetInstances(stage.CodeFlightRuleTypes, stage.CodeFlightRuleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeflightruletypeIdent := codeflightruletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightRuleType{Name: %s}).Stage(stage)", codeflightruletypeIdent, __gong__toRawStringLiteral(codeflightruletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightruletypeIdent, __gong__toRawStringLiteral(codeflightruletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightruletypeIdent, __gong__toRawStringLiteral(codeflightruletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightruletypeIdent, __gong__toRawStringLiteral(codeflightruletype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeflightstatustype := range __gong__sortStageSetInstances(stage.CodeFlightStatusTypes, stage.CodeFlightStatusType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeflightstatustypeIdent := codeflightstatustype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightStatusType{Name: %s}).Stage(stage)", codeflightstatustypeIdent, __gong__toRawStringLiteral(codeflightstatustype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightstatustypeIdent, __gong__toRawStringLiteral(codeflightstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightstatustypeIdent, __gong__toRawStringLiteral(codeflightstatustype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightstatustypeIdent, __gong__toRawStringLiteral(codeflightstatustype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeflighttype := range __gong__sortStageSetInstances(stage.CodeFlightTypes, stage.CodeFlightType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeflighttypeIdent := codeflighttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightType{Name: %s}).Stage(stage)", codeflighttypeIdent, __gong__toRawStringLiteral(codeflighttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflighttypeIdent, __gong__toRawStringLiteral(codeflighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflighttypeIdent, __gong__toRawStringLiteral(codeflighttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflighttypeIdent, __gong__toRawStringLiteral(codeflighttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeflowconditionoperationtype := range __gong__sortStageSetInstances(stage.CodeFlowConditionOperationTypes, stage.CodeFlowConditionOperationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeflowconditionoperationtypeIdent := codeflowconditionoperationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlowConditionOperationType{Name: %s}).Stage(stage)", codeflowconditionoperationtypeIdent, __gong__toRawStringLiteral(codeflowconditionoperationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflowconditionoperationtypeIdent, __gong__toRawStringLiteral(codeflowconditionoperationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflowconditionoperationtypeIdent, __gong__toRawStringLiteral(codeflowconditionoperationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflowconditionoperationtypeIdent, __gong__toRawStringLiteral(codeflowconditionoperationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codefreeflighttype := range __gong__sortStageSetInstances(stage.CodeFreeFlightTypes, stage.CodeFreeFlightType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codefreeflighttypeIdent := codefreeflighttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFreeFlightType{Name: %s}).Stage(stage)", codefreeflighttypeIdent, __gong__toRawStringLiteral(codefreeflighttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefreeflighttypeIdent, __gong__toRawStringLiteral(codefreeflighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefreeflighttypeIdent, __gong__toRawStringLiteral(codefreeflighttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefreeflighttypeIdent, __gong__toRawStringLiteral(codefreeflighttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeleveltype := range __gong__sortStageSetInstances(stage.CodeLevelTypes, stage.CodeLevelType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeleveltypeIdent := codeleveltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLevelType{Name: %s}).Stage(stage)", codeleveltypeIdent, __gong__toRawStringLiteral(codeleveltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeleveltypeIdent, __gong__toRawStringLiteral(codeleveltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeleveltypeIdent, __gong__toRawStringLiteral(codeleveltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeleveltypeIdent, __gong__toRawStringLiteral(codeleveltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codelocationqualifiertype := range __gong__sortStageSetInstances(stage.CodeLocationQualifierTypes, stage.CodeLocationQualifierType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codelocationqualifiertypeIdent := codelocationqualifiertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLocationQualifierType{Name: %s}).Stage(stage)", codelocationqualifiertypeIdent, __gong__toRawStringLiteral(codelocationqualifiertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelocationqualifiertypeIdent, __gong__toRawStringLiteral(codelocationqualifiertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelocationqualifiertypeIdent, __gong__toRawStringLiteral(codelocationqualifiertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelocationqualifiertypeIdent, __gong__toRawStringLiteral(codelocationqualifiertype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codemeteoconditionstype := range __gong__sortStageSetInstances(stage.CodeMeteoConditionsTypes, stage.CodeMeteoConditionsType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemeteoconditionstypeIdent := codemeteoconditionstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMeteoConditionsType{Name: %s}).Stage(stage)", codemeteoconditionstypeIdent, __gong__toRawStringLiteral(codemeteoconditionstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemeteoconditionstypeIdent, __gong__toRawStringLiteral(codemeteoconditionstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemeteoconditionstypeIdent, __gong__toRawStringLiteral(codemeteoconditionstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemeteoconditionstypeIdent, __gong__toRawStringLiteral(codemeteoconditionstype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codemilitaryroutepointtype := range __gong__sortStageSetInstances(stage.CodeMilitaryRoutePointTypes, stage.CodeMilitaryRoutePointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemilitaryroutepointtypeIdent := codemilitaryroutepointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMilitaryRoutePointType{Name: %s}).Stage(stage)", codemilitaryroutepointtypeIdent, __gong__toRawStringLiteral(codemilitaryroutepointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemilitaryroutepointtypeIdent, __gong__toRawStringLiteral(codemilitaryroutepointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemilitaryroutepointtypeIdent, __gong__toRawStringLiteral(codemilitaryroutepointtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemilitaryroutepointtypeIdent, __gong__toRawStringLiteral(codemilitaryroutepointtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codemilitarystatustype := range __gong__sortStageSetInstances(stage.CodeMilitaryStatusTypes, stage.CodeMilitaryStatusType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemilitarystatustypeIdent := codemilitarystatustype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMilitaryStatusType{Name: %s}).Stage(stage)", codemilitarystatustypeIdent, __gong__toRawStringLiteral(codemilitarystatustype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemilitarystatustypeIdent, __gong__toRawStringLiteral(codemilitarystatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemilitarystatustypeIdent, __gong__toRawStringLiteral(codemilitarystatustype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemilitarystatustypeIdent, __gong__toRawStringLiteral(codemilitarystatustype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codemilitarytrainingtype := range __gong__sortStageSetInstances(stage.CodeMilitaryTrainingTypes, stage.CodeMilitaryTrainingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemilitarytrainingtypeIdent := codemilitarytrainingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMilitaryTrainingType{Name: %s}).Stage(stage)", codemilitarytrainingtypeIdent, __gong__toRawStringLiteral(codemilitarytrainingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemilitarytrainingtypeIdent, __gong__toRawStringLiteral(codemilitarytrainingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemilitarytrainingtypeIdent, __gong__toRawStringLiteral(codemilitarytrainingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemilitarytrainingtypeIdent, __gong__toRawStringLiteral(codemilitarytrainingtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codenavigationequipmenttype := range __gong__sortStageSetInstances(stage.CodeNavigationEquipmentTypes, stage.CodeNavigationEquipmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codenavigationequipmenttypeIdent := codenavigationequipmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavigationEquipmentType{Name: %s}).Stage(stage)", codenavigationequipmenttypeIdent, __gong__toRawStringLiteral(codenavigationequipmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavigationequipmenttypeIdent, __gong__toRawStringLiteral(codenavigationequipmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavigationequipmenttypeIdent, __gong__toRawStringLiteral(codenavigationequipmenttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavigationequipmenttypeIdent, __gong__toRawStringLiteral(codenavigationequipmenttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codenavigationspecificationtype := range __gong__sortStageSetInstances(stage.CodeNavigationSpecificationTypes, stage.CodeNavigationSpecificationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codenavigationspecificationtypeIdent := codenavigationspecificationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavigationSpecificationType{Name: %s}).Stage(stage)", codenavigationspecificationtypeIdent, __gong__toRawStringLiteral(codenavigationspecificationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavigationspecificationtypeIdent, __gong__toRawStringLiteral(codenavigationspecificationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavigationspecificationtypeIdent, __gong__toRawStringLiteral(codenavigationspecificationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavigationspecificationtypeIdent, __gong__toRawStringLiteral(codenavigationspecificationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codenotepurposetype := range __gong__sortStageSetInstances(stage.CodeNotePurposeTypes, stage.CodeNotePurposeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codenotepurposetypeIdent := codenotepurposetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNotePurposeType{Name: %s}).Stage(stage)", codenotepurposetypeIdent, __gong__toRawStringLiteral(codenotepurposetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenotepurposetypeIdent, __gong__toRawStringLiteral(codenotepurposetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenotepurposetypeIdent, __gong__toRawStringLiteral(codenotepurposetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenotepurposetypeIdent, __gong__toRawStringLiteral(codenotepurposetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeobstacleassessmentsurfacetype := range __gong__sortStageSetInstances(stage.CodeObstacleAssessmentSurfaceTypes, stage.CodeObstacleAssessmentSurfaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeobstacleassessmentsurfacetypeIdent := codeobstacleassessmentsurfacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeObstacleAssessmentSurfaceType{Name: %s}).Stage(stage)", codeobstacleassessmentsurfacetypeIdent, __gong__toRawStringLiteral(codeobstacleassessmentsurfacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeobstacleassessmentsurfacetypeIdent, __gong__toRawStringLiteral(codeobstacleassessmentsurfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeobstacleassessmentsurfacetypeIdent, __gong__toRawStringLiteral(codeobstacleassessmentsurfacetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeobstacleassessmentsurfacetypeIdent, __gong__toRawStringLiteral(codeobstacleassessmentsurfacetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeobstructionidsurfacezonetype := range __gong__sortStageSetInstances(stage.CodeObstructionIdSurfaceZoneTypes, stage.CodeObstructionIdSurfaceZoneType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeobstructionidsurfacezonetypeIdent := codeobstructionidsurfacezonetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeObstructionIdSurfaceZoneType{Name: %s}).Stage(stage)", codeobstructionidsurfacezonetypeIdent, __gong__toRawStringLiteral(codeobstructionidsurfacezonetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeobstructionidsurfacezonetypeIdent, __gong__toRawStringLiteral(codeobstructionidsurfacezonetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeobstructionidsurfacezonetypeIdent, __gong__toRawStringLiteral(codeobstructionidsurfacezonetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeobstructionidsurfacezonetypeIdent, __gong__toRawStringLiteral(codeobstructionidsurfacezonetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codernptype := range __gong__sortStageSetInstances(stage.CodeRNPTypes, stage.CodeRNPType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codernptypeIdent := codernptype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRNPType{Name: %s}).Stage(stage)", codernptypeIdent, __gong__toRawStringLiteral(codernptype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codernptypeIdent, __gong__toRawStringLiteral(codernptype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codernptypeIdent, __gong__toRawStringLiteral(codernptype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codernptypeIdent, __gong__toRawStringLiteral(codernptype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codervsmpointroletype := range __gong__sortStageSetInstances(stage.CodeRVSMPointRoleTypes, stage.CodeRVSMPointRoleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codervsmpointroletypeIdent := codervsmpointroletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRVSMPointRoleType{Name: %s}).Stage(stage)", codervsmpointroletypeIdent, __gong__toRawStringLiteral(codervsmpointroletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codervsmpointroletypeIdent, __gong__toRawStringLiteral(codervsmpointroletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codervsmpointroletypeIdent, __gong__toRawStringLiteral(codervsmpointroletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codervsmpointroletypeIdent, __gong__toRawStringLiteral(codervsmpointroletype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codervsmtype := range __gong__sortStageSetInstances(stage.CodeRVSMTypes, stage.CodeRVSMType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codervsmtypeIdent := codervsmtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRVSMType{Name: %s}).Stage(stage)", codervsmtypeIdent, __gong__toRawStringLiteral(codervsmtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codervsmtypeIdent, __gong__toRawStringLiteral(codervsmtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codervsmtypeIdent, __gong__toRawStringLiteral(codervsmtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codervsmtypeIdent, __gong__toRawStringLiteral(codervsmtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codereferenceroletype := range __gong__sortStageSetInstances(stage.CodeReferenceRoleTypes, stage.CodeReferenceRoleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codereferenceroletypeIdent := codereferenceroletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeReferenceRoleType{Name: %s}).Stage(stage)", codereferenceroletypeIdent, __gong__toRawStringLiteral(codereferenceroletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codereferenceroletypeIdent, __gong__toRawStringLiteral(codereferenceroletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codereferenceroletypeIdent, __gong__toRawStringLiteral(codereferenceroletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codereferenceroletypeIdent, __gong__toRawStringLiteral(codereferenceroletype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderouteavailabilitytype := range __gong__sortStageSetInstances(stage.CodeRouteAvailabilityTypes, stage.CodeRouteAvailabilityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderouteavailabilitytypeIdent := coderouteavailabilitytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteAvailabilityType{Name: %s}).Stage(stage)", coderouteavailabilitytypeIdent, __gong__toRawStringLiteral(coderouteavailabilitytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderouteavailabilitytypeIdent, __gong__toRawStringLiteral(coderouteavailabilitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderouteavailabilitytypeIdent, __gong__toRawStringLiteral(coderouteavailabilitytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderouteavailabilitytypeIdent, __gong__toRawStringLiteral(coderouteavailabilitytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderoutedesignatorlettertype := range __gong__sortStageSetInstances(stage.CodeRouteDesignatorLetterTypes, stage.CodeRouteDesignatorLetterType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderoutedesignatorlettertypeIdent := coderoutedesignatorlettertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteDesignatorLetterType{Name: %s}).Stage(stage)", coderoutedesignatorlettertypeIdent, __gong__toRawStringLiteral(coderoutedesignatorlettertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutedesignatorlettertypeIdent, __gong__toRawStringLiteral(coderoutedesignatorlettertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutedesignatorlettertypeIdent, __gong__toRawStringLiteral(coderoutedesignatorlettertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutedesignatorlettertypeIdent, __gong__toRawStringLiteral(coderoutedesignatorlettertype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderoutedesignatorprefixtype := range __gong__sortStageSetInstances(stage.CodeRouteDesignatorPrefixTypes, stage.CodeRouteDesignatorPrefixType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderoutedesignatorprefixtypeIdent := coderoutedesignatorprefixtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteDesignatorPrefixType{Name: %s}).Stage(stage)", coderoutedesignatorprefixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorprefixtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutedesignatorprefixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorprefixtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutedesignatorprefixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorprefixtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutedesignatorprefixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorprefixtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderoutedesignatorsuffixtype := range __gong__sortStageSetInstances(stage.CodeRouteDesignatorSuffixTypes, stage.CodeRouteDesignatorSuffixType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderoutedesignatorsuffixtypeIdent := coderoutedesignatorsuffixtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteDesignatorSuffixType{Name: %s}).Stage(stage)", coderoutedesignatorsuffixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorsuffixtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutedesignatorsuffixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorsuffixtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutedesignatorsuffixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorsuffixtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutedesignatorsuffixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorsuffixtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderoutenavigationtype := range __gong__sortStageSetInstances(stage.CodeRouteNavigationTypes, stage.CodeRouteNavigationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderoutenavigationtypeIdent := coderoutenavigationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteNavigationType{Name: %s}).Stage(stage)", coderoutenavigationtypeIdent, __gong__toRawStringLiteral(coderoutenavigationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutenavigationtypeIdent, __gong__toRawStringLiteral(coderoutenavigationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutenavigationtypeIdent, __gong__toRawStringLiteral(coderoutenavigationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutenavigationtypeIdent, __gong__toRawStringLiteral(coderoutenavigationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderouteorigintype := range __gong__sortStageSetInstances(stage.CodeRouteOriginTypes, stage.CodeRouteOriginType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderouteorigintypeIdent := coderouteorigintype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteOriginType{Name: %s}).Stage(stage)", coderouteorigintypeIdent, __gong__toRawStringLiteral(coderouteorigintype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderouteorigintypeIdent, __gong__toRawStringLiteral(coderouteorigintype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderouteorigintypeIdent, __gong__toRawStringLiteral(coderouteorigintype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderouteorigintypeIdent, __gong__toRawStringLiteral(coderouteorigintype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderoutesegmentpathtype := range __gong__sortStageSetInstances(stage.CodeRouteSegmentPathTypes, stage.CodeRouteSegmentPathType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderoutesegmentpathtypeIdent := coderoutesegmentpathtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteSegmentPathType{Name: %s}).Stage(stage)", coderoutesegmentpathtypeIdent, __gong__toRawStringLiteral(coderoutesegmentpathtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutesegmentpathtypeIdent, __gong__toRawStringLiteral(coderoutesegmentpathtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutesegmentpathtypeIdent, __gong__toRawStringLiteral(coderoutesegmentpathtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutesegmentpathtypeIdent, __gong__toRawStringLiteral(coderoutesegmentpathtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderoutetype := range __gong__sortStageSetInstances(stage.CodeRouteTypes, stage.CodeRouteType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderoutetypeIdent := coderoutetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteType{Name: %s}).Stage(stage)", coderoutetypeIdent, __gong__toRawStringLiteral(coderoutetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutetypeIdent, __gong__toRawStringLiteral(coderoutetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutetypeIdent, __gong__toRawStringLiteral(coderoutetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutetypeIdent, __gong__toRawStringLiteral(coderoutetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderuleproceduretitletype := range __gong__sortStageSetInstances(stage.CodeRuleProcedureTitleTypes, stage.CodeRuleProcedureTitleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderuleproceduretitletypeIdent := coderuleproceduretitletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRuleProcedureTitleType{Name: %s}).Stage(stage)", coderuleproceduretitletypeIdent, __gong__toRawStringLiteral(coderuleproceduretitletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderuleproceduretitletypeIdent, __gong__toRawStringLiteral(coderuleproceduretitletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderuleproceduretitletypeIdent, __gong__toRawStringLiteral(coderuleproceduretitletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderuleproceduretitletypeIdent, __gong__toRawStringLiteral(coderuleproceduretitletype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderuleproceduretype := range __gong__sortStageSetInstances(stage.CodeRuleProcedureTypes, stage.CodeRuleProcedureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderuleproceduretypeIdent := coderuleproceduretype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRuleProcedureType{Name: %s}).Stage(stage)", coderuleproceduretypeIdent, __gong__toRawStringLiteral(coderuleproceduretype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderuleproceduretypeIdent, __gong__toRawStringLiteral(coderuleproceduretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderuleproceduretypeIdent, __gong__toRawStringLiteral(coderuleproceduretype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderuleproceduretypeIdent, __gong__toRawStringLiteral(coderuleproceduretype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codesidetype := range __gong__sortStageSetInstances(stage.CodeSideTypes, stage.CodeSideType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codesidetypeIdent := codesidetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSideType{Name: %s}).Stage(stage)", codesidetypeIdent, __gong__toRawStringLiteral(codesidetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesidetypeIdent, __gong__toRawStringLiteral(codesidetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesidetypeIdent, __gong__toRawStringLiteral(codesidetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesidetypeIdent, __gong__toRawStringLiteral(codesidetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codespeedreferencetype := range __gong__sortStageSetInstances(stage.CodeSpeedReferenceTypes, stage.CodeSpeedReferenceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codespeedreferencetypeIdent := codespeedreferencetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSpeedReferenceType{Name: %s}).Stage(stage)", codespeedreferencetypeIdent, __gong__toRawStringLiteral(codespeedreferencetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codespeedreferencetypeIdent, __gong__toRawStringLiteral(codespeedreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codespeedreferencetypeIdent, __gong__toRawStringLiteral(codespeedreferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codespeedreferencetypeIdent, __gong__toRawStringLiteral(codespeedreferencetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codetacanchanneltype := range __gong__sortStageSetInstances(stage.CodeTACANChannelTypes, stage.CodeTACANChannelType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetacanchanneltypeIdent := codetacanchanneltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTACANChannelType{Name: %s}).Stage(stage)", codetacanchanneltypeIdent, __gong__toRawStringLiteral(codetacanchanneltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetacanchanneltypeIdent, __gong__toRawStringLiteral(codetacanchanneltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetacanchanneltypeIdent, __gong__toRawStringLiteral(codetacanchanneltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetacanchanneltypeIdent, __gong__toRawStringLiteral(codetacanchanneltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codetelecomnetworktype := range __gong__sortStageSetInstances(stage.CodeTelecomNetworkTypes, stage.CodeTelecomNetworkType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetelecomnetworktypeIdent := codetelecomnetworktype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTelecomNetworkType{Name: %s}).Stage(stage)", codetelecomnetworktypeIdent, __gong__toRawStringLiteral(codetelecomnetworktype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetelecomnetworktypeIdent, __gong__toRawStringLiteral(codetelecomnetworktype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetelecomnetworktypeIdent, __gong__toRawStringLiteral(codetelecomnetworktype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetelecomnetworktypeIdent, __gong__toRawStringLiteral(codetelecomnetworktype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codetimeeventcombinationtype := range __gong__sortStageSetInstances(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetimeeventcombinationtypeIdent := codetimeeventcombinationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTimeEventCombinationType{Name: %s}).Stage(stage)", codetimeeventcombinationtypeIdent, __gong__toRawStringLiteral(codetimeeventcombinationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetimeeventcombinationtypeIdent, __gong__toRawStringLiteral(codetimeeventcombinationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetimeeventcombinationtypeIdent, __gong__toRawStringLiteral(codetimeeventcombinationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetimeeventcombinationtypeIdent, __gong__toRawStringLiteral(codetimeeventcombinationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codetimeeventtype := range __gong__sortStageSetInstances(stage.CodeTimeEventTypes, stage.CodeTimeEventType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetimeeventtypeIdent := codetimeeventtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTimeEventType{Name: %s}).Stage(stage)", codetimeeventtypeIdent, __gong__toRawStringLiteral(codetimeeventtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetimeeventtypeIdent, __gong__toRawStringLiteral(codetimeeventtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetimeeventtypeIdent, __gong__toRawStringLiteral(codetimeeventtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetimeeventtypeIdent, __gong__toRawStringLiteral(codetimeeventtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codetimereferencetype := range __gong__sortStageSetInstances(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetimereferencetypeIdent := codetimereferencetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTimeReferenceType{Name: %s}).Stage(stage)", codetimereferencetypeIdent, __gong__toRawStringLiteral(codetimereferencetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetimereferencetypeIdent, __gong__toRawStringLiteral(codetimereferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetimereferencetypeIdent, __gong__toRawStringLiteral(codetimereferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetimereferencetypeIdent, __gong__toRawStringLiteral(codetimereferencetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codetranspondertype := range __gong__sortStageSetInstances(stage.CodeTransponderTypes, stage.CodeTransponderType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetranspondertypeIdent := codetranspondertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTransponderType{Name: %s}).Stage(stage)", codetranspondertypeIdent, __gong__toRawStringLiteral(codetranspondertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetranspondertypeIdent, __gong__toRawStringLiteral(codetranspondertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetranspondertypeIdent, __gong__toRawStringLiteral(codetranspondertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetranspondertypeIdent, __gong__toRawStringLiteral(codetranspondertype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeupperalphatype := range __gong__sortStageSetInstances(stage.CodeUpperAlphaTypes, stage.CodeUpperAlphaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeupperalphatypeIdent := codeupperalphatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeUpperAlphaType{Name: %s}).Stage(stage)", codeupperalphatypeIdent, __gong__toRawStringLiteral(codeupperalphatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeupperalphatypeIdent, __gong__toRawStringLiteral(codeupperalphatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeupperalphatypeIdent, __gong__toRawStringLiteral(codeupperalphatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeupperalphatypeIdent, __gong__toRawStringLiteral(codeupperalphatype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codevalueinterpretationtype := range __gong__sortStageSetInstances(stage.CodeValueInterpretationTypes, stage.CodeValueInterpretationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codevalueinterpretationtypeIdent := codevalueinterpretationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeValueInterpretationType{Name: %s}).Stage(stage)", codevalueinterpretationtypeIdent, __gong__toRawStringLiteral(codevalueinterpretationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codevalueinterpretationtypeIdent, __gong__toRawStringLiteral(codevalueinterpretationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codevalueinterpretationtypeIdent, __gong__toRawStringLiteral(codevalueinterpretationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codevalueinterpretationtypeIdent, __gong__toRawStringLiteral(codevalueinterpretationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeverticaldatumtype := range __gong__sortStageSetInstances(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeverticaldatumtypeIdent := codeverticaldatumtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVerticalDatumType{Name: %s}).Stage(stage)", codeverticaldatumtypeIdent, __gong__toRawStringLiteral(codeverticaldatumtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeverticaldatumtypeIdent, __gong__toRawStringLiteral(codeverticaldatumtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeverticaldatumtypeIdent, __gong__toRawStringLiteral(codeverticaldatumtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeverticaldatumtypeIdent, __gong__toRawStringLiteral(codeverticaldatumtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeverticalreferencetype := range __gong__sortStageSetInstances(stage.CodeVerticalReferenceTypes, stage.CodeVerticalReferenceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeverticalreferencetypeIdent := codeverticalreferencetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVerticalReferenceType{Name: %s}).Stage(stage)", codeverticalreferencetypeIdent, __gong__toRawStringLiteral(codeverticalreferencetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeverticalreferencetypeIdent, __gong__toRawStringLiteral(codeverticalreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeverticalreferencetypeIdent, __gong__toRawStringLiteral(codeverticalreferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeverticalreferencetypeIdent, __gong__toRawStringLiteral(codeverticalreferencetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codewaketurbulencetype := range __gong__sortStageSetInstances(stage.CodeWakeTurbulenceTypes, stage.CodeWakeTurbulenceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codewaketurbulencetypeIdent := codewaketurbulencetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeWakeTurbulenceType{Name: %s}).Stage(stage)", codewaketurbulencetypeIdent, __gong__toRawStringLiteral(codewaketurbulencetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codewaketurbulencetypeIdent, __gong__toRawStringLiteral(codewaketurbulencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codewaketurbulencetypeIdent, __gong__toRawStringLiteral(codewaketurbulencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codewaketurbulencetypeIdent, __gong__toRawStringLiteral(codewaketurbulencetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeyesnotype := range __gong__sortStageSetInstances(stage.CodeYesNoTypes, stage.CodeYesNoType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeyesnotypeIdent := codeyesnotype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeYesNoType{Name: %s}).Stage(stage)", codeyesnotypeIdent, __gong__toRawStringLiteral(codeyesnotype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeyesnotypeIdent, __gong__toRawStringLiteral(codeyesnotype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeyesnotypeIdent, __gong__toRawStringLiteral(codeyesnotype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeyesnotypeIdent, __gong__toRawStringLiteral(codeyesnotype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, datemonthdaytype := range __gong__sortStageSetInstances(stage.DateMonthDayTypes, stage.DateMonthDayType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datemonthdaytypeIdent := datemonthdaytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.DateMonthDayType{Name: %s}).Stage(stage)", datemonthdaytypeIdent, __gong__toRawStringLiteral(datemonthdaytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datemonthdaytypeIdent, __gong__toRawStringLiteral(datemonthdaytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", datemonthdaytypeIdent, __gong__toRawStringLiteral(datemonthdaytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", datemonthdaytypeIdent, __gong__toRawStringLiteral(datemonthdaytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, nonumbertype := range __gong__sortStageSetInstances(stage.NoNumberTypes, stage.NoNumberType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			nonumbertypeIdent := nonumbertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.NoNumberType{Name: %s}).Stage(stage)", nonumbertypeIdent, __gong__toRawStringLiteral(nonumbertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nonumbertypeIdent, __gong__toRawStringLiteral(nonumbertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", nonumbertypeIdent, __gong__toRawStringLiteral(nonumbertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", nonumbertypeIdent, nonumbertype.EnclosedText))
		}
	}
	if stage != nil {
		for _, nosequencetype := range __gong__sortStageSetInstances(stage.NoSequenceTypes, stage.NoSequenceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			nosequencetypeIdent := nosequencetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.NoSequenceType{Name: %s}).Stage(stage)", nosequencetypeIdent, __gong__toRawStringLiteral(nosequencetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nosequencetypeIdent, __gong__toRawStringLiteral(nosequencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", nosequencetypeIdent, __gong__toRawStringLiteral(nosequencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", nosequencetypeIdent, nosequencetype.EnclosedText))
		}
	}
	if stage != nil {
		for _, textaddresstype := range __gong__sortStageSetInstances(stage.TextAddressTypes, stage.TextAddressType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textaddresstypeIdent := textaddresstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextAddressType{Name: %s}).Stage(stage)", textaddresstypeIdent, __gong__toRawStringLiteral(textaddresstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textaddresstypeIdent, __gong__toRawStringLiteral(textaddresstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textaddresstypeIdent, __gong__toRawStringLiteral(textaddresstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textaddresstypeIdent, __gong__toRawStringLiteral(textaddresstype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, textdesignatortype := range __gong__sortStageSetInstances(stage.TextDesignatorTypes, stage.TextDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textdesignatortypeIdent := textdesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextDesignatorType{Name: %s}).Stage(stage)", textdesignatortypeIdent, __gong__toRawStringLiteral(textdesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textdesignatortypeIdent, __gong__toRawStringLiteral(textdesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textdesignatortypeIdent, __gong__toRawStringLiteral(textdesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textdesignatortypeIdent, __gong__toRawStringLiteral(textdesignatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, textinstructiontype := range __gong__sortStageSetInstances(stage.TextInstructionTypes, stage.TextInstructionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textinstructiontypeIdent := textinstructiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextInstructionType{Name: %s}).Stage(stage)", textinstructiontypeIdent, __gong__toRawStringLiteral(textinstructiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textinstructiontypeIdent, __gong__toRawStringLiteral(textinstructiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textinstructiontypeIdent, __gong__toRawStringLiteral(textinstructiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textinstructiontypeIdent, __gong__toRawStringLiteral(textinstructiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, textnametype := range __gong__sortStageSetInstances(stage.TextNameTypes, stage.TextNameType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textnametypeIdent := textnametype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextNameType{Name: %s}).Stage(stage)", textnametypeIdent, __gong__toRawStringLiteral(textnametype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textnametypeIdent, __gong__toRawStringLiteral(textnametype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textnametypeIdent, __gong__toRawStringLiteral(textnametype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textnametypeIdent, __gong__toRawStringLiteral(textnametype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, textnotetype := range __gong__sortStageSetInstances(stage.TextNoteTypes, stage.TextNoteType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textnotetypeIdent := textnotetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextNoteType{Name: %s}).Stage(stage)", textnotetypeIdent, __gong__toRawStringLiteral(textnotetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textnotetypeIdent, __gong__toRawStringLiteral(textnotetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textnotetypeIdent, __gong__toRawStringLiteral(textnotetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", textnotetypeIdent, __gong__toRawStringLiteral(textnotetype.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textnotetypeIdent, __gong__toRawStringLiteral(textnotetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, textphonetype := range __gong__sortStageSetInstances(stage.TextPhoneTypes, stage.TextPhoneType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textphonetypeIdent := textphonetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextPhoneType{Name: %s}).Stage(stage)", textphonetypeIdent, __gong__toRawStringLiteral(textphonetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textphonetypeIdent, __gong__toRawStringLiteral(textphonetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textphonetypeIdent, __gong__toRawStringLiteral(textphonetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textphonetypeIdent, __gong__toRawStringLiteral(textphonetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, textpropertynametype := range __gong__sortStageSetInstances(stage.TextPropertyNameTypes, stage.TextPropertyNameType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textpropertynametypeIdent := textpropertynametype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextPropertyNameType{Name: %s}).Stage(stage)", textpropertynametypeIdent, __gong__toRawStringLiteral(textpropertynametype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textpropertynametypeIdent, __gong__toRawStringLiteral(textpropertynametype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textpropertynametypeIdent, __gong__toRawStringLiteral(textpropertynametype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textpropertynametypeIdent, __gong__toRawStringLiteral(textpropertynametype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, timetype := range __gong__sortStageSetInstances(stage.TimeTypes, stage.TimeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timetypeIdent := timetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TimeType{Name: %s}).Stage(stage)", timetypeIdent, __gong__toRawStringLiteral(timetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timetypeIdent, __gong__toRawStringLiteral(timetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", timetypeIdent, __gong__toRawStringLiteral(timetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", timetypeIdent, __gong__toRawStringLiteral(timetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valangletype := range __gong__sortStageSetInstances(stage.ValAngleTypes, stage.ValAngleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valangletypeIdent := valangletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValAngleType{Name: %s}).Stage(stage)", valangletypeIdent, __gong__toRawStringLiteral(valangletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valangletypeIdent, __gong__toRawStringLiteral(valangletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valangletypeIdent, __gong__toRawStringLiteral(valangletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valangletypeIdent, __gong__toRawStringLiteral(valangletype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valbearingtype := range __gong__sortStageSetInstances(stage.ValBearingTypes, stage.ValBearingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valbearingtypeIdent := valbearingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValBearingType{Name: %s}).Stage(stage)", valbearingtypeIdent, __gong__toRawStringLiteral(valbearingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valbearingtypeIdent, __gong__toRawStringLiteral(valbearingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valbearingtypeIdent, __gong__toRawStringLiteral(valbearingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valbearingtypeIdent, __gong__toRawStringLiteral(valbearingtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valdistancesignedtype := range __gong__sortStageSetInstances(stage.ValDistanceSignedTypes, stage.ValDistanceSignedType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valdistancesignedtypeIdent := valdistancesignedtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValDistanceSignedType{Name: %s}).Stage(stage)", valdistancesignedtypeIdent, __gong__toRawStringLiteral(valdistancesignedtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valdistancesignedtypeIdent, __gong__toRawStringLiteral(valdistancesignedtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valdistancesignedtypeIdent, __gong__toRawStringLiteral(valdistancesignedtype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valdistancesignedtypeIdent, __gong__toRawStringLiteral(valdistancesignedtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valdistancesignedtypeIdent, __gong__toRawStringLiteral(valdistancesignedtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valdistancetype := range __gong__sortStageSetInstances(stage.ValDistanceTypes, stage.ValDistanceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valdistancetypeIdent := valdistancetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValDistanceType{Name: %s}).Stage(stage)", valdistancetypeIdent, __gong__toRawStringLiteral(valdistancetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valdistancetypeIdent, __gong__toRawStringLiteral(valdistancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valdistancetypeIdent, __gong__toRawStringLiteral(valdistancetype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valdistancetypeIdent, __gong__toRawStringLiteral(valdistancetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valdistancetypeIdent, __gong__toRawStringLiteral(valdistancetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valdistanceverticaltype := range __gong__sortStageSetInstances(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valdistanceverticaltypeIdent := valdistanceverticaltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValDistanceVerticalType{Name: %s}).Stage(stage)", valdistanceverticaltypeIdent, __gong__toRawStringLiteral(valdistanceverticaltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valdistanceverticaltypeIdent, __gong__toRawStringLiteral(valdistanceverticaltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valdistanceverticaltypeIdent, __gong__toRawStringLiteral(valdistanceverticaltype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valdistanceverticaltypeIdent, __gong__toRawStringLiteral(valdistanceverticaltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valdistanceverticaltypeIdent, __gong__toRawStringLiteral(valdistanceverticaltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valdurationtype := range __gong__sortStageSetInstances(stage.ValDurationTypes, stage.ValDurationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valdurationtypeIdent := valdurationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValDurationType{Name: %s}).Stage(stage)", valdurationtypeIdent, __gong__toRawStringLiteral(valdurationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valdurationtypeIdent, __gong__toRawStringLiteral(valdurationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valdurationtypeIdent, __gong__toRawStringLiteral(valdurationtype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valdurationtypeIdent, __gong__toRawStringLiteral(valdurationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valdurationtypeIdent, __gong__toRawStringLiteral(valdurationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valslopetype := range __gong__sortStageSetInstances(stage.ValSlopeTypes, stage.ValSlopeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valslopetypeIdent := valslopetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValSlopeType{Name: %s}).Stage(stage)", valslopetypeIdent, __gong__toRawStringLiteral(valslopetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valslopetypeIdent, __gong__toRawStringLiteral(valslopetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valslopetypeIdent, __gong__toRawStringLiteral(valslopetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valslopetypeIdent, __gong__toRawStringLiteral(valslopetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valspeedtype := range __gong__sortStageSetInstances(stage.ValSpeedTypes, stage.ValSpeedType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valspeedtypeIdent := valspeedtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValSpeedType{Name: %s}).Stage(stage)", valspeedtypeIdent, __gong__toRawStringLiteral(valspeedtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valspeedtypeIdent, __gong__toRawStringLiteral(valspeedtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valspeedtypeIdent, __gong__toRawStringLiteral(valspeedtype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valspeedtypeIdent, __gong__toRawStringLiteral(valspeedtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valspeedtypeIdent, __gong__toRawStringLiteral(valspeedtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valweighttype := range __gong__sortStageSetInstances(stage.ValWeightTypes, stage.ValWeightType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valweighttypeIdent := valweighttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValWeightType{Name: %s}).Stage(stage)", valweighttypeIdent, __gong__toRawStringLiteral(valweighttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valweighttypeIdent, __gong__toRawStringLiteral(valweighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valweighttypeIdent, __gong__toRawStringLiteral(valweighttype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valweighttypeIdent, __gong__toRawStringLiteral(valweighttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valweighttypeIdent, __gong__toRawStringLiteral(valweighttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, xhtmltype := range __gong__sortStageSetInstances(stage.XHTMLTypes, stage.XHTMLType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			xhtmltypeIdent := xhtmltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.XHTMLType{Name: %s}).Stage(stage)", xhtmltypeIdent, __gong__toRawStringLiteral(xhtmltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xhtmltypeIdent, __gong__toRawStringLiteral(xhtmltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", xhtmltypeIdent, __gong__toRawStringLiteral(xhtmltype.NilReason)))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *aixmdatatypes.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *aixmdatatypes.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (codeatcreportingtype *CodeATCReportingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeatcreportingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeatcreportingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeatcreportingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeatcreportingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeATCReportingType", fieldName)
	}
	return
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaerialrefuellingpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaerialrefuellingpointtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaerialrefuellingpointtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaerialrefuellingpointtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAerialRefuellingPointType", fieldName)
	}
	return
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaerialrefuellingprefixtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaerialrefuellingprefixtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaerialrefuellingprefixtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaerialrefuellingprefixtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAerialRefuellingPrefixType", fieldName)
	}
	return
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaerialrefuellingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaerialrefuellingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaerialrefuellingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaerialrefuellingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAerialRefuellingType", fieldName)
	}
	return
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaircraftcategorytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaircraftcategorytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaircraftcategorytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaircraftcategorytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAircraftCategoryType", fieldName)
	}
	return
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaircraftenginenumbertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaircraftenginenumbertype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaircraftenginenumbertype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaircraftenginenumbertype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAircraftEngineNumberType", fieldName)
	}
	return
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaircraftenginetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaircraftenginetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaircraftenginetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaircraftenginetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAircraftEngineType", fieldName)
	}
	return
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaircrafticaotype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaircrafticaotype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaircrafticaotype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaircrafticaotype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAircraftICAOType", fieldName)
	}
	return
}

func (codeaircrafttype *CodeAircraftType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaircrafttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaircrafttype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaircrafttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaircrafttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAircraftType", fieldName)
	}
	return
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaircraftwingspanclasstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaircraftwingspanclasstype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaircraftwingspanclasstype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaircraftwingspanclasstype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAircraftWingspanClassType", fieldName)
	}
	return
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codealtitudeadjustmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codealtitudeadjustmenttype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codealtitudeadjustmenttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codealtitudeadjustmenttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAltitudeAdjustmentType", fieldName)
	}
	return
}

func (codealtitudeusetype *CodeAltitudeUseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codealtitudeusetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codealtitudeusetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codealtitudeusetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codealtitudeusetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAltitudeUseType", fieldName)
	}
	return
}

func (codeauthoritytype *CodeAuthorityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeauthoritytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeauthoritytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeauthoritytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeauthoritytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAuthorityType", fieldName)
	}
	return
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codecardinaldirectiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codecardinaldirectiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codecardinaldirectiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codecardinaldirectiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeCardinalDirectionType", fieldName)
	}
	return
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codecommunicationmodetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codecommunicationmodetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codecommunicationmodetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codecommunicationmodetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeCommunicationModeType", fieldName)
	}
	return
}

func (codecomparisontype *CodeComparisonType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codecomparisontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codecomparisontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codecomparisontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codecomparisontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeComparisonType", fieldName)
	}
	return
}

func (codecoursetype *CodeCourseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codecoursetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codecoursetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codecoursetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codecoursetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeCourseType", fieldName)
	}
	return
}

func (codedaytype *CodeDayType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedaytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedaytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedaytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedaytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDayType", fieldName)
	}
	return
}

func (codedirectionturntype *CodeDirectionTurnType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedirectionturntype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedirectionturntype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedirectionturntype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedirectionturntype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDirectionTurnType", fieldName)
	}
	return
}

func (codedirectiontype *CodeDirectionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedirectiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedirectiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedirectiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedirectiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDirectionType", fieldName)
	}
	return
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeequipmentanticollisiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeequipmentanticollisiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeequipmentanticollisiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeequipmentanticollisiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeEquipmentAntiCollisionType", fieldName)
	}
	return
}

func (codeflightorigintype *CodeFlightOriginType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeflightorigintype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeflightorigintype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeflightorigintype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeflightorigintype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFlightOriginType", fieldName)
	}
	return
}

func (codeflightpurposetype *CodeFlightPurposeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeflightpurposetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeflightpurposetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeflightpurposetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeflightpurposetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFlightPurposeType", fieldName)
	}
	return
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeflightrestrictiondesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeflightrestrictiondesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeflightrestrictiondesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeflightrestrictiondesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFlightRestrictionDesignatorType", fieldName)
	}
	return
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeflightrestrictiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeflightrestrictiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeflightrestrictiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeflightrestrictiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFlightRestrictionType", fieldName)
	}
	return
}

func (codeflightruletype *CodeFlightRuleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeflightruletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeflightruletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeflightruletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeflightruletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFlightRuleType", fieldName)
	}
	return
}

func (codeflightstatustype *CodeFlightStatusType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeflightstatustype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeflightstatustype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeflightstatustype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeflightstatustype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFlightStatusType", fieldName)
	}
	return
}

func (codeflighttype *CodeFlightType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeflighttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeflighttype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeflighttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeflighttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFlightType", fieldName)
	}
	return
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeflowconditionoperationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeflowconditionoperationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeflowconditionoperationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeflowconditionoperationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFlowConditionOperationType", fieldName)
	}
	return
}

func (codefreeflighttype *CodeFreeFlightType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codefreeflighttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codefreeflighttype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codefreeflighttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codefreeflighttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFreeFlightType", fieldName)
	}
	return
}

func (codeleveltype *CodeLevelType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeleveltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeleveltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeleveltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeleveltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLevelType", fieldName)
	}
	return
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codelocationqualifiertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codelocationqualifiertype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codelocationqualifiertype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codelocationqualifiertype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLocationQualifierType", fieldName)
	}
	return
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemeteoconditionstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemeteoconditionstype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemeteoconditionstype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemeteoconditionstype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMeteoConditionsType", fieldName)
	}
	return
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemilitaryroutepointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemilitaryroutepointtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemilitaryroutepointtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemilitaryroutepointtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMilitaryRoutePointType", fieldName)
	}
	return
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemilitarystatustype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemilitarystatustype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemilitarystatustype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemilitarystatustype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMilitaryStatusType", fieldName)
	}
	return
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemilitarytrainingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemilitarytrainingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemilitarytrainingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemilitarytrainingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMilitaryTrainingType", fieldName)
	}
	return
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codenavigationequipmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codenavigationequipmenttype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codenavigationequipmenttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codenavigationequipmenttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNavigationEquipmentType", fieldName)
	}
	return
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codenavigationspecificationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codenavigationspecificationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codenavigationspecificationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codenavigationspecificationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNavigationSpecificationType", fieldName)
	}
	return
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codenotepurposetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codenotepurposetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codenotepurposetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codenotepurposetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNotePurposeType", fieldName)
	}
	return
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeobstacleassessmentsurfacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeobstacleassessmentsurfacetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeobstacleassessmentsurfacetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeobstacleassessmentsurfacetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeObstacleAssessmentSurfaceType", fieldName)
	}
	return
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeobstructionidsurfacezonetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeobstructionidsurfacezonetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeobstructionidsurfacezonetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeobstructionidsurfacezonetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeObstructionIdSurfaceZoneType", fieldName)
	}
	return
}

func (codernptype *CodeRNPType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codernptype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codernptype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codernptype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codernptype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRNPType", fieldName)
	}
	return
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codervsmpointroletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codervsmpointroletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codervsmpointroletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codervsmpointroletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRVSMPointRoleType", fieldName)
	}
	return
}

func (codervsmtype *CodeRVSMType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codervsmtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codervsmtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codervsmtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codervsmtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRVSMType", fieldName)
	}
	return
}

func (codereferenceroletype *CodeReferenceRoleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codereferenceroletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codereferenceroletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codereferenceroletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codereferenceroletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeReferenceRoleType", fieldName)
	}
	return
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderouteavailabilitytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderouteavailabilitytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderouteavailabilitytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderouteavailabilitytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRouteAvailabilityType", fieldName)
	}
	return
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderoutedesignatorlettertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderoutedesignatorlettertype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderoutedesignatorlettertype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderoutedesignatorlettertype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRouteDesignatorLetterType", fieldName)
	}
	return
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderoutedesignatorprefixtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderoutedesignatorprefixtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderoutedesignatorprefixtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderoutedesignatorprefixtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRouteDesignatorPrefixType", fieldName)
	}
	return
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderoutedesignatorsuffixtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderoutedesignatorsuffixtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderoutedesignatorsuffixtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderoutedesignatorsuffixtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRouteDesignatorSuffixType", fieldName)
	}
	return
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderoutenavigationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderoutenavigationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderoutenavigationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderoutenavigationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRouteNavigationType", fieldName)
	}
	return
}

func (coderouteorigintype *CodeRouteOriginType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderouteorigintype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderouteorigintype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderouteorigintype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderouteorigintype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRouteOriginType", fieldName)
	}
	return
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderoutesegmentpathtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderoutesegmentpathtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderoutesegmentpathtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderoutesegmentpathtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRouteSegmentPathType", fieldName)
	}
	return
}

func (coderoutetype *CodeRouteType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderoutetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderoutetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderoutetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderoutetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRouteType", fieldName)
	}
	return
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderuleproceduretitletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderuleproceduretitletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderuleproceduretitletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderuleproceduretitletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRuleProcedureTitleType", fieldName)
	}
	return
}

func (coderuleproceduretype *CodeRuleProcedureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderuleproceduretype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderuleproceduretype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderuleproceduretype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderuleproceduretype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRuleProcedureType", fieldName)
	}
	return
}

func (codesidetype *CodeSideType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codesidetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codesidetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codesidetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codesidetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSideType", fieldName)
	}
	return
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codespeedreferencetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codespeedreferencetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codespeedreferencetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codespeedreferencetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSpeedReferenceType", fieldName)
	}
	return
}

func (codetacanchanneltype *CodeTACANChannelType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetacanchanneltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetacanchanneltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetacanchanneltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetacanchanneltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTACANChannelType", fieldName)
	}
	return
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetelecomnetworktype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetelecomnetworktype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetelecomnetworktype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetelecomnetworktype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTelecomNetworkType", fieldName)
	}
	return
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetimeeventcombinationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetimeeventcombinationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetimeeventcombinationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetimeeventcombinationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTimeEventCombinationType", fieldName)
	}
	return
}

func (codetimeeventtype *CodeTimeEventType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetimeeventtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetimeeventtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetimeeventtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetimeeventtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTimeEventType", fieldName)
	}
	return
}

func (codetimereferencetype *CodeTimeReferenceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetimereferencetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetimereferencetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetimereferencetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetimereferencetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTimeReferenceType", fieldName)
	}
	return
}

func (codetranspondertype *CodeTransponderType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetranspondertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetranspondertype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetranspondertype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetranspondertype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTransponderType", fieldName)
	}
	return
}

func (codeupperalphatype *CodeUpperAlphaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeupperalphatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeupperalphatype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeupperalphatype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeupperalphatype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeUpperAlphaType", fieldName)
	}
	return
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codevalueinterpretationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codevalueinterpretationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codevalueinterpretationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codevalueinterpretationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeValueInterpretationType", fieldName)
	}
	return
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeverticaldatumtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeverticaldatumtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeverticaldatumtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeverticaldatumtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeVerticalDatumType", fieldName)
	}
	return
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeverticalreferencetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeverticalreferencetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeverticalreferencetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeverticalreferencetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeVerticalReferenceType", fieldName)
	}
	return
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codewaketurbulencetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codewaketurbulencetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codewaketurbulencetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codewaketurbulencetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeWakeTurbulenceType", fieldName)
	}
	return
}

func (codeyesnotype *CodeYesNoType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeyesnotype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeyesnotype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeyesnotype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeyesnotype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeYesNoType", fieldName)
	}
	return
}

func (datemonthdaytype *DateMonthDayType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datemonthdaytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datemonthdaytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", datemonthdaytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", datemonthdaytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct DateMonthDayType", fieldName)
	}
	return
}

func (nonumbertype *NoNumberType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := nonumbertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", nonumbertype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", nonumbertype.NilReason)
	case "EnclosedText":
		res = __gong__marshallInt(ident, "EnclosedText", nonumbertype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct NoNumberType", fieldName)
	}
	return
}

func (nosequencetype *NoSequenceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := nosequencetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", nosequencetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", nosequencetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallInt(ident, "EnclosedText", nosequencetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct NoSequenceType", fieldName)
	}
	return
}

func (textaddresstype *TextAddressType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := textaddresstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", textaddresstype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", textaddresstype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", textaddresstype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TextAddressType", fieldName)
	}
	return
}

func (textdesignatortype *TextDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := textdesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", textdesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", textdesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", textdesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TextDesignatorType", fieldName)
	}
	return
}

func (textinstructiontype *TextInstructionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := textinstructiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", textinstructiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", textinstructiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", textinstructiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TextInstructionType", fieldName)
	}
	return
}

func (textnametype *TextNameType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := textnametype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", textnametype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", textnametype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", textnametype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TextNameType", fieldName)
	}
	return
}

func (textnotetype *TextNoteType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := textnotetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", textnotetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", textnotetype.NilReason)
	case "Lang":
		res = __gong__marshallString(ident, "Lang", textnotetype.Lang)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", textnotetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TextNoteType", fieldName)
	}
	return
}

func (textphonetype *TextPhoneType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := textphonetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", textphonetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", textphonetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", textphonetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TextPhoneType", fieldName)
	}
	return
}

func (textpropertynametype *TextPropertyNameType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := textpropertynametype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", textpropertynametype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", textpropertynametype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", textpropertynametype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TextPropertyNameType", fieldName)
	}
	return
}

func (timetype *TimeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := timetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", timetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", timetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", timetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TimeType", fieldName)
	}
	return
}

func (valangletype *ValAngleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valangletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valangletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valangletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valangletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValAngleType", fieldName)
	}
	return
}

func (valbearingtype *ValBearingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valbearingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valbearingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valbearingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valbearingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValBearingType", fieldName)
	}
	return
}

func (valdistancesignedtype *ValDistanceSignedType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valdistancesignedtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valdistancesignedtype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valdistancesignedtype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valdistancesignedtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valdistancesignedtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValDistanceSignedType", fieldName)
	}
	return
}

func (valdistancetype *ValDistanceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valdistancetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valdistancetype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valdistancetype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valdistancetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valdistancetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValDistanceType", fieldName)
	}
	return
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valdistanceverticaltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valdistanceverticaltype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valdistanceverticaltype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valdistanceverticaltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valdistanceverticaltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValDistanceVerticalType", fieldName)
	}
	return
}

func (valdurationtype *ValDurationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valdurationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valdurationtype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valdurationtype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valdurationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valdurationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValDurationType", fieldName)
	}
	return
}

func (valslopetype *ValSlopeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valslopetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valslopetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valslopetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valslopetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValSlopeType", fieldName)
	}
	return
}

func (valspeedtype *ValSpeedType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valspeedtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valspeedtype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valspeedtype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valspeedtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valspeedtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValSpeedType", fieldName)
	}
	return
}

func (valweighttype *ValWeightType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valweighttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valweighttype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valweighttype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valweighttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valweighttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValWeightType", fieldName)
	}
	return
}

func (xhtmltype *XHTMLType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := xhtmltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", xhtmltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", xhtmltype.NilReason)

	default:
		log.Panicf("Unknown field %s for Gongstruct XHTMLType", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (codeatcreportingtype *CodeATCReportingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeatcreportingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeatcreportingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeatcreportingtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaerialrefuellingpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaerialrefuellingpointtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaerialrefuellingpointtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaerialrefuellingprefixtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaerialrefuellingprefixtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaerialrefuellingprefixtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaerialrefuellingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaerialrefuellingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaerialrefuellingtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaircraftcategorytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaircraftcategorytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaircraftcategorytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaircraftenginenumbertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaircraftenginenumbertype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaircraftenginenumbertype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeaircraftenginetype *CodeAircraftEngineType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaircraftenginetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaircraftenginetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaircraftenginetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeaircrafticaotype *CodeAircraftICAOType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaircrafticaotype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaircrafticaotype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaircrafticaotype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeaircrafttype *CodeAircraftType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaircrafttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaircrafttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaircrafttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaircraftwingspanclasstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaircraftwingspanclasstype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaircraftwingspanclasstype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codealtitudeadjustmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codealtitudeadjustmenttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codealtitudeadjustmenttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codealtitudeusetype *CodeAltitudeUseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codealtitudeusetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codealtitudeusetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codealtitudeusetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeauthoritytype *CodeAuthorityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeauthoritytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeauthoritytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeauthoritytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codecardinaldirectiontype *CodeCardinalDirectionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codecardinaldirectiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codecardinaldirectiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codecardinaldirectiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codecommunicationmodetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codecommunicationmodetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codecommunicationmodetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codecomparisontype *CodeComparisonType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codecomparisontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codecomparisontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codecomparisontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codecoursetype *CodeCourseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codecoursetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codecoursetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codecoursetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codedaytype *CodeDayType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedaytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedaytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedaytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codedirectionturntype *CodeDirectionTurnType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedirectionturntype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedirectionturntype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedirectionturntype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codedirectiontype *CodeDirectionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedirectiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedirectiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedirectiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeequipmentanticollisiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeequipmentanticollisiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeequipmentanticollisiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeflightorigintype *CodeFlightOriginType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeflightorigintype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeflightorigintype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeflightorigintype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeflightpurposetype *CodeFlightPurposeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeflightpurposetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeflightpurposetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeflightpurposetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeflightrestrictiondesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeflightrestrictiondesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeflightrestrictiondesignatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeflightrestrictiontype *CodeFlightRestrictionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeflightrestrictiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeflightrestrictiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeflightrestrictiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeflightruletype *CodeFlightRuleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeflightruletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeflightruletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeflightruletype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeflightstatustype *CodeFlightStatusType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeflightstatustype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeflightstatustype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeflightstatustype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeflighttype *CodeFlightType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeflighttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeflighttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeflighttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeflowconditionoperationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeflowconditionoperationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeflowconditionoperationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codefreeflighttype *CodeFreeFlightType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codefreeflighttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codefreeflighttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codefreeflighttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeleveltype *CodeLevelType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeleveltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeleveltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeleveltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codelocationqualifiertype *CodeLocationQualifierType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codelocationqualifiertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codelocationqualifiertype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codelocationqualifiertype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codemeteoconditionstype *CodeMeteoConditionsType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemeteoconditionstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemeteoconditionstype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemeteoconditionstype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemilitaryroutepointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemilitaryroutepointtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemilitaryroutepointtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codemilitarystatustype *CodeMilitaryStatusType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemilitarystatustype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemilitarystatustype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemilitarystatustype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemilitarytrainingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemilitarytrainingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemilitarytrainingtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codenavigationequipmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codenavigationequipmenttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codenavigationequipmenttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codenavigationspecificationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codenavigationspecificationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codenavigationspecificationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codenotepurposetype *CodeNotePurposeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codenotepurposetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codenotepurposetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codenotepurposetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeobstacleassessmentsurfacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeobstacleassessmentsurfacetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeobstacleassessmentsurfacetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeobstructionidsurfacezonetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeobstructionidsurfacezonetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeobstructionidsurfacezonetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codernptype *CodeRNPType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codernptype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codernptype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codernptype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codervsmpointroletype *CodeRVSMPointRoleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codervsmpointroletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codervsmpointroletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codervsmpointroletype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codervsmtype *CodeRVSMType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codervsmtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codervsmtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codervsmtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codereferenceroletype *CodeReferenceRoleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codereferenceroletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codereferenceroletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codereferenceroletype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderouteavailabilitytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderouteavailabilitytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderouteavailabilitytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderoutedesignatorlettertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderoutedesignatorlettertype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderoutedesignatorlettertype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderoutedesignatorprefixtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderoutedesignatorprefixtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderoutedesignatorprefixtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderoutedesignatorsuffixtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderoutedesignatorsuffixtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderoutedesignatorsuffixtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderoutenavigationtype *CodeRouteNavigationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderoutenavigationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderoutenavigationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderoutenavigationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderouteorigintype *CodeRouteOriginType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderouteorigintype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderouteorigintype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderouteorigintype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderoutesegmentpathtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderoutesegmentpathtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderoutesegmentpathtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderoutetype *CodeRouteType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderoutetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderoutetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderoutetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderuleproceduretitletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderuleproceduretitletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderuleproceduretitletype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderuleproceduretype *CodeRuleProcedureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderuleproceduretype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderuleproceduretype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderuleproceduretype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codesidetype *CodeSideType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codesidetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codesidetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codesidetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codespeedreferencetype *CodeSpeedReferenceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codespeedreferencetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codespeedreferencetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codespeedreferencetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codetacanchanneltype *CodeTACANChannelType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetacanchanneltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetacanchanneltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetacanchanneltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codetelecomnetworktype *CodeTelecomNetworkType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetelecomnetworktype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetelecomnetworktype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetelecomnetworktype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetimeeventcombinationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetimeeventcombinationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetimeeventcombinationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codetimeeventtype *CodeTimeEventType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetimeeventtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetimeeventtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetimeeventtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codetimereferencetype *CodeTimeReferenceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetimereferencetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetimereferencetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetimereferencetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codetranspondertype *CodeTransponderType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetranspondertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetranspondertype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetranspondertype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeupperalphatype *CodeUpperAlphaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeupperalphatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeupperalphatype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeupperalphatype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codevalueinterpretationtype *CodeValueInterpretationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codevalueinterpretationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codevalueinterpretationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codevalueinterpretationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeverticaldatumtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeverticaldatumtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeverticaldatumtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeverticalreferencetype *CodeVerticalReferenceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeverticalreferencetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeverticalreferencetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeverticalreferencetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codewaketurbulencetype *CodeWakeTurbulenceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codewaketurbulencetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codewaketurbulencetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codewaketurbulencetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeyesnotype *CodeYesNoType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeyesnotype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeyesnotype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeyesnotype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datemonthdaytype *DateMonthDayType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datemonthdaytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datemonthdaytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(datemonthdaytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (nonumbertype *NoNumberType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(nonumbertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(nonumbertype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(nonumbertype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (nosequencetype *NoSequenceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(nosequencetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(nosequencetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(nosequencetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (textaddresstype *TextAddressType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(textaddresstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(textaddresstype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(textaddresstype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (textdesignatortype *TextDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(textdesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(textdesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(textdesignatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (textinstructiontype *TextInstructionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(textinstructiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(textinstructiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(textinstructiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (textnametype *TextNameType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(textnametype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(textnametype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(textnametype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (textnotetype *TextNoteType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(textnotetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(textnotetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(textnotetype.GongMarshallField(stage, "Lang"))
		initializerStatements.WriteString(textnotetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (textphonetype *TextPhoneType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(textphonetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(textphonetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(textphonetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (textpropertynametype *TextPropertyNameType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(textpropertynametype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(textpropertynametype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(textpropertynametype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (timetype *TimeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(timetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(timetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(timetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valangletype *ValAngleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valangletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valangletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valangletype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valbearingtype *ValBearingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valbearingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valbearingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valbearingtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valdistancesignedtype *ValDistanceSignedType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valdistancesignedtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valdistancesignedtype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valdistancesignedtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valdistancesignedtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valdistancetype *ValDistanceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valdistancetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valdistancetype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valdistancetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valdistancetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valdistanceverticaltype *ValDistanceVerticalType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valdistanceverticaltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valdistanceverticaltype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valdistanceverticaltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valdistanceverticaltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valdurationtype *ValDurationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valdurationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valdurationtype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valdurationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valdurationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valslopetype *ValSlopeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valslopetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valslopetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valslopetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valspeedtype *ValSpeedType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valspeedtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valspeedtype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valspeedtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valspeedtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valweighttype *ValWeightType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valweighttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valweighttype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valweighttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valweighttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (xhtmltype *XHTMLType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(xhtmltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xhtmltype.GongMarshallField(stage, "NilReason"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
