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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes"
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
		for _, codeapproachequipmentadditionaltype := range __gong__sortStageSetInstances(stage.CodeApproachEquipmentAdditionalTypes, stage.CodeApproachEquipmentAdditionalType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeapproachequipmentadditionaltypeIdent := codeapproachequipmentadditionaltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeApproachEquipmentAdditionalType{Name: %s}).Stage(stage)", codeapproachequipmentadditionaltypeIdent, __gong__toRawStringLiteral(codeapproachequipmentadditionaltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeapproachequipmentadditionaltypeIdent, __gong__toRawStringLiteral(codeapproachequipmentadditionaltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeapproachequipmentadditionaltypeIdent, __gong__toRawStringLiteral(codeapproachequipmentadditionaltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeapproachequipmentadditionaltypeIdent, __gong__toRawStringLiteral(codeapproachequipmentadditionaltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeapproachguidancetype := range __gong__sortStageSetInstances(stage.CodeApproachGuidanceTypes, stage.CodeApproachGuidanceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeapproachguidancetypeIdent := codeapproachguidancetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeApproachGuidanceType{Name: %s}).Stage(stage)", codeapproachguidancetypeIdent, __gong__toRawStringLiteral(codeapproachguidancetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeapproachguidancetypeIdent, __gong__toRawStringLiteral(codeapproachguidancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeapproachguidancetypeIdent, __gong__toRawStringLiteral(codeapproachguidancetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeapproachguidancetypeIdent, __gong__toRawStringLiteral(codeapproachguidancetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeapproachprefixtype := range __gong__sortStageSetInstances(stage.CodeApproachPrefixTypes, stage.CodeApproachPrefixType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeapproachprefixtypeIdent := codeapproachprefixtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeApproachPrefixType{Name: %s}).Stage(stage)", codeapproachprefixtypeIdent, __gong__toRawStringLiteral(codeapproachprefixtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeapproachprefixtypeIdent, __gong__toRawStringLiteral(codeapproachprefixtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeapproachprefixtypeIdent, __gong__toRawStringLiteral(codeapproachprefixtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeapproachprefixtypeIdent, __gong__toRawStringLiteral(codeapproachprefixtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeapproachtype := range __gong__sortStageSetInstances(stage.CodeApproachTypes, stage.CodeApproachType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeapproachtypeIdent := codeapproachtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeApproachType{Name: %s}).Stage(stage)", codeapproachtypeIdent, __gong__toRawStringLiteral(codeapproachtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeapproachtypeIdent, __gong__toRawStringLiteral(codeapproachtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeapproachtypeIdent, __gong__toRawStringLiteral(codeapproachtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeapproachtypeIdent, __gong__toRawStringLiteral(codeapproachtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeapprovaltype := range __gong__sortStageSetInstances(stage.CodeApprovalTypes, stage.CodeApprovalType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeapprovaltypeIdent := codeapprovaltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeApprovalType{Name: %s}).Stage(stage)", codeapprovaltypeIdent, __gong__toRawStringLiteral(codeapprovaltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeapprovaltypeIdent, __gong__toRawStringLiteral(codeapprovaltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeapprovaltypeIdent, __gong__toRawStringLiteral(codeapprovaltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeapprovaltypeIdent, __gong__toRawStringLiteral(codeapprovaltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codearcdirectiontype := range __gong__sortStageSetInstances(stage.CodeArcDirectionTypes, stage.CodeArcDirectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codearcdirectiontypeIdent := codearcdirectiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeArcDirectionType{Name: %s}).Stage(stage)", codearcdirectiontypeIdent, __gong__toRawStringLiteral(codearcdirectiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codearcdirectiontypeIdent, __gong__toRawStringLiteral(codearcdirectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codearcdirectiontypeIdent, __gong__toRawStringLiteral(codearcdirectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codearcdirectiontypeIdent, __gong__toRawStringLiteral(codearcdirectiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codebearingtype := range __gong__sortStageSetInstances(stage.CodeBearingTypes, stage.CodeBearingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codebearingtypeIdent := codebearingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeBearingType{Name: %s}).Stage(stage)", codebearingtypeIdent, __gong__toRawStringLiteral(codebearingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codebearingtypeIdent, __gong__toRawStringLiteral(codebearingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codebearingtypeIdent, __gong__toRawStringLiteral(codebearingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codebearingtypeIdent, __gong__toRawStringLiteral(codebearingtype.EnclosedText)))
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
		for _, codedesignstandardtype := range __gong__sortStageSetInstances(stage.CodeDesignStandardTypes, stage.CodeDesignStandardType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedesignstandardtypeIdent := codedesignstandardtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDesignStandardType{Name: %s}).Stage(stage)", codedesignstandardtypeIdent, __gong__toRawStringLiteral(codedesignstandardtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedesignstandardtypeIdent, __gong__toRawStringLiteral(codedesignstandardtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedesignstandardtypeIdent, __gong__toRawStringLiteral(codedesignstandardtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedesignstandardtypeIdent, __gong__toRawStringLiteral(codedesignstandardtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codedesignatedpointdesignatortype := range __gong__sortStageSetInstances(stage.CodeDesignatedPointDesignatorTypes, stage.CodeDesignatedPointDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedesignatedpointdesignatortypeIdent := codedesignatedpointdesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDesignatedPointDesignatorType{Name: %s}).Stage(stage)", codedesignatedpointdesignatortypeIdent, __gong__toRawStringLiteral(codedesignatedpointdesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedesignatedpointdesignatortypeIdent, __gong__toRawStringLiteral(codedesignatedpointdesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedesignatedpointdesignatortypeIdent, __gong__toRawStringLiteral(codedesignatedpointdesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedesignatedpointdesignatortypeIdent, __gong__toRawStringLiteral(codedesignatedpointdesignatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codedirectionreferencetype := range __gong__sortStageSetInstances(stage.CodeDirectionReferenceTypes, stage.CodeDirectionReferenceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedirectionreferencetypeIdent := codedirectionreferencetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDirectionReferenceType{Name: %s}).Stage(stage)", codedirectionreferencetypeIdent, __gong__toRawStringLiteral(codedirectionreferencetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedirectionreferencetypeIdent, __gong__toRawStringLiteral(codedirectionreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedirectionreferencetypeIdent, __gong__toRawStringLiteral(codedirectionreferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedirectionreferencetypeIdent, __gong__toRawStringLiteral(codedirectionreferencetype.EnclosedText)))
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
		for _, codeequipmentunavailabletype := range __gong__sortStageSetInstances(stage.CodeEquipmentUnavailableTypes, stage.CodeEquipmentUnavailableType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeequipmentunavailabletypeIdent := codeequipmentunavailabletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeEquipmentUnavailableType{Name: %s}).Stage(stage)", codeequipmentunavailabletypeIdent, __gong__toRawStringLiteral(codeequipmentunavailabletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeequipmentunavailabletypeIdent, __gong__toRawStringLiteral(codeequipmentunavailabletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeequipmentunavailabletypeIdent, __gong__toRawStringLiteral(codeequipmentunavailabletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeequipmentunavailabletypeIdent, __gong__toRawStringLiteral(codeequipmentunavailabletype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codefinalguidancetype := range __gong__sortStageSetInstances(stage.CodeFinalGuidanceTypes, stage.CodeFinalGuidanceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codefinalguidancetypeIdent := codefinalguidancetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFinalGuidanceType{Name: %s}).Stage(stage)", codefinalguidancetypeIdent, __gong__toRawStringLiteral(codefinalguidancetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefinalguidancetypeIdent, __gong__toRawStringLiteral(codefinalguidancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefinalguidancetypeIdent, __gong__toRawStringLiteral(codefinalguidancetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefinalguidancetypeIdent, __gong__toRawStringLiteral(codefinalguidancetype.EnclosedText)))
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
		for _, codeheightreferencetype := range __gong__sortStageSetInstances(stage.CodeHeightReferenceTypes, stage.CodeHeightReferenceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeheightreferencetypeIdent := codeheightreferencetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeHeightReferenceType{Name: %s}).Stage(stage)", codeheightreferencetypeIdent, __gong__toRawStringLiteral(codeheightreferencetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeheightreferencetypeIdent, __gong__toRawStringLiteral(codeheightreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeheightreferencetypeIdent, __gong__toRawStringLiteral(codeheightreferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeheightreferencetypeIdent, __gong__toRawStringLiteral(codeheightreferencetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeholdingusagetype := range __gong__sortStageSetInstances(stage.CodeHoldingUsageTypes, stage.CodeHoldingUsageType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeholdingusagetypeIdent := codeholdingusagetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeHoldingUsageType{Name: %s}).Stage(stage)", codeholdingusagetypeIdent, __gong__toRawStringLiteral(codeholdingusagetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeholdingusagetypeIdent, __gong__toRawStringLiteral(codeholdingusagetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeholdingusagetypeIdent, __gong__toRawStringLiteral(codeholdingusagetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeholdingusagetypeIdent, __gong__toRawStringLiteral(codeholdingusagetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeholdingusetype := range __gong__sortStageSetInstances(stage.CodeHoldingUseTypes, stage.CodeHoldingUseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeholdingusetypeIdent := codeholdingusetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeHoldingUseType{Name: %s}).Stage(stage)", codeholdingusetypeIdent, __gong__toRawStringLiteral(codeholdingusetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeholdingusetypeIdent, __gong__toRawStringLiteral(codeholdingusetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeholdingusetypeIdent, __gong__toRawStringLiteral(codeholdingusetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeholdingusetypeIdent, __gong__toRawStringLiteral(codeholdingusetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeicaocountrytype := range __gong__sortStageSetInstances(stage.CodeICAOCountryTypes, stage.CodeICAOCountryType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeicaocountrytypeIdent := codeicaocountrytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeICAOCountryType{Name: %s}).Stage(stage)", codeicaocountrytypeIdent, __gong__toRawStringLiteral(codeicaocountrytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeicaocountrytypeIdent, __gong__toRawStringLiteral(codeicaocountrytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeicaocountrytypeIdent, __gong__toRawStringLiteral(codeicaocountrytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeicaocountrytypeIdent, __gong__toRawStringLiteral(codeicaocountrytype.EnclosedText)))
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
		for _, codeminimafinalapproachpathtype := range __gong__sortStageSetInstances(stage.CodeMinimaFinalApproachPathTypes, stage.CodeMinimaFinalApproachPathType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeminimafinalapproachpathtypeIdent := codeminimafinalapproachpathtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMinimaFinalApproachPathType{Name: %s}).Stage(stage)", codeminimafinalapproachpathtypeIdent, __gong__toRawStringLiteral(codeminimafinalapproachpathtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeminimafinalapproachpathtypeIdent, __gong__toRawStringLiteral(codeminimafinalapproachpathtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeminimafinalapproachpathtypeIdent, __gong__toRawStringLiteral(codeminimafinalapproachpathtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeminimafinalapproachpathtypeIdent, __gong__toRawStringLiteral(codeminimafinalapproachpathtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeminimumaltitudetype := range __gong__sortStageSetInstances(stage.CodeMinimumAltitudeTypes, stage.CodeMinimumAltitudeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeminimumaltitudetypeIdent := codeminimumaltitudetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMinimumAltitudeType{Name: %s}).Stage(stage)", codeminimumaltitudetypeIdent, __gong__toRawStringLiteral(codeminimumaltitudetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeminimumaltitudetypeIdent, __gong__toRawStringLiteral(codeminimumaltitudetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeminimumaltitudetypeIdent, __gong__toRawStringLiteral(codeminimumaltitudetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeminimumaltitudetypeIdent, __gong__toRawStringLiteral(codeminimumaltitudetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeminimumheighttype := range __gong__sortStageSetInstances(stage.CodeMinimumHeightTypes, stage.CodeMinimumHeightType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeminimumheighttypeIdent := codeminimumheighttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMinimumHeightType{Name: %s}).Stage(stage)", codeminimumheighttypeIdent, __gong__toRawStringLiteral(codeminimumheighttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeminimumheighttypeIdent, __gong__toRawStringLiteral(codeminimumheighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeminimumheighttypeIdent, __gong__toRawStringLiteral(codeminimumheighttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeminimumheighttypeIdent, __gong__toRawStringLiteral(codeminimumheighttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codemissedapproachtype := range __gong__sortStageSetInstances(stage.CodeMissedApproachTypes, stage.CodeMissedApproachType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemissedapproachtypeIdent := codemissedapproachtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMissedApproachType{Name: %s}).Stage(stage)", codemissedapproachtypeIdent, __gong__toRawStringLiteral(codemissedapproachtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemissedapproachtypeIdent, __gong__toRawStringLiteral(codemissedapproachtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemissedapproachtypeIdent, __gong__toRawStringLiteral(codemissedapproachtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemissedapproachtypeIdent, __gong__toRawStringLiteral(codemissedapproachtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codenavigationarearestrictiontype := range __gong__sortStageSetInstances(stage.CodeNavigationAreaRestrictionTypes, stage.CodeNavigationAreaRestrictionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codenavigationarearestrictiontypeIdent := codenavigationarearestrictiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavigationAreaRestrictionType{Name: %s}).Stage(stage)", codenavigationarearestrictiontypeIdent, __gong__toRawStringLiteral(codenavigationarearestrictiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavigationarearestrictiontypeIdent, __gong__toRawStringLiteral(codenavigationarearestrictiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavigationarearestrictiontypeIdent, __gong__toRawStringLiteral(codenavigationarearestrictiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavigationarearestrictiontypeIdent, __gong__toRawStringLiteral(codenavigationarearestrictiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codenavigationareatype := range __gong__sortStageSetInstances(stage.CodeNavigationAreaTypes, stage.CodeNavigationAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codenavigationareatypeIdent := codenavigationareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavigationAreaType{Name: %s}).Stage(stage)", codenavigationareatypeIdent, __gong__toRawStringLiteral(codenavigationareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavigationareatypeIdent, __gong__toRawStringLiteral(codenavigationareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavigationareatypeIdent, __gong__toRawStringLiteral(codenavigationareatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavigationareatypeIdent, __gong__toRawStringLiteral(codenavigationareatype.EnclosedText)))
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
		for _, codeprocedureavailabilitytype := range __gong__sortStageSetInstances(stage.CodeProcedureAvailabilityTypes, stage.CodeProcedureAvailabilityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeprocedureavailabilitytypeIdent := codeprocedureavailabilitytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeProcedureAvailabilityType{Name: %s}).Stage(stage)", codeprocedureavailabilitytypeIdent, __gong__toRawStringLiteral(codeprocedureavailabilitytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeprocedureavailabilitytypeIdent, __gong__toRawStringLiteral(codeprocedureavailabilitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeprocedureavailabilitytypeIdent, __gong__toRawStringLiteral(codeprocedureavailabilitytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeprocedureavailabilitytypeIdent, __gong__toRawStringLiteral(codeprocedureavailabilitytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeprocedurecodingstandardtype := range __gong__sortStageSetInstances(stage.CodeProcedureCodingStandardTypes, stage.CodeProcedureCodingStandardType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeprocedurecodingstandardtypeIdent := codeprocedurecodingstandardtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeProcedureCodingStandardType{Name: %s}).Stage(stage)", codeprocedurecodingstandardtypeIdent, __gong__toRawStringLiteral(codeprocedurecodingstandardtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeprocedurecodingstandardtypeIdent, __gong__toRawStringLiteral(codeprocedurecodingstandardtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeprocedurecodingstandardtypeIdent, __gong__toRawStringLiteral(codeprocedurecodingstandardtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeprocedurecodingstandardtypeIdent, __gong__toRawStringLiteral(codeprocedurecodingstandardtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeproceduredistancetype := range __gong__sortStageSetInstances(stage.CodeProcedureDistanceTypes, stage.CodeProcedureDistanceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeproceduredistancetypeIdent := codeproceduredistancetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeProcedureDistanceType{Name: %s}).Stage(stage)", codeproceduredistancetypeIdent, __gong__toRawStringLiteral(codeproceduredistancetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeproceduredistancetypeIdent, __gong__toRawStringLiteral(codeproceduredistancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeproceduredistancetypeIdent, __gong__toRawStringLiteral(codeproceduredistancetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeproceduredistancetypeIdent, __gong__toRawStringLiteral(codeproceduredistancetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeprocedurefixroletype := range __gong__sortStageSetInstances(stage.CodeProcedureFixRoleTypes, stage.CodeProcedureFixRoleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeprocedurefixroletypeIdent := codeprocedurefixroletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeProcedureFixRoleType{Name: %s}).Stage(stage)", codeprocedurefixroletypeIdent, __gong__toRawStringLiteral(codeprocedurefixroletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeprocedurefixroletypeIdent, __gong__toRawStringLiteral(codeprocedurefixroletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeprocedurefixroletypeIdent, __gong__toRawStringLiteral(codeprocedurefixroletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeprocedurefixroletypeIdent, __gong__toRawStringLiteral(codeprocedurefixroletype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeprocedurephasetype := range __gong__sortStageSetInstances(stage.CodeProcedurePhaseTypes, stage.CodeProcedurePhaseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeprocedurephasetypeIdent := codeprocedurephasetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeProcedurePhaseType{Name: %s}).Stage(stage)", codeprocedurephasetypeIdent, __gong__toRawStringLiteral(codeprocedurephasetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeprocedurephasetypeIdent, __gong__toRawStringLiteral(codeprocedurephasetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeprocedurephasetypeIdent, __gong__toRawStringLiteral(codeprocedurephasetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeprocedurephasetypeIdent, __gong__toRawStringLiteral(codeprocedurephasetype.EnclosedText)))
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
		for _, codereferencepathidentifiertype := range __gong__sortStageSetInstances(stage.CodeReferencePathIdentifierTypes, stage.CodeReferencePathIdentifierType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codereferencepathidentifiertypeIdent := codereferencepathidentifiertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeReferencePathIdentifierType{Name: %s}).Stage(stage)", codereferencepathidentifiertypeIdent, __gong__toRawStringLiteral(codereferencepathidentifiertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codereferencepathidentifiertypeIdent, __gong__toRawStringLiteral(codereferencepathidentifiertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codereferencepathidentifiertypeIdent, __gong__toRawStringLiteral(codereferencepathidentifiertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codereferencepathidentifiertypeIdent, __gong__toRawStringLiteral(codereferencepathidentifiertype.EnclosedText)))
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
		for _, coderelativepositiontype := range __gong__sortStageSetInstances(stage.CodeRelativePositionTypes, stage.CodeRelativePositionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderelativepositiontypeIdent := coderelativepositiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRelativePositionType{Name: %s}).Stage(stage)", coderelativepositiontypeIdent, __gong__toRawStringLiteral(coderelativepositiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderelativepositiontypeIdent, __gong__toRawStringLiteral(coderelativepositiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderelativepositiontypeIdent, __gong__toRawStringLiteral(coderelativepositiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderelativepositiontypeIdent, __gong__toRawStringLiteral(coderelativepositiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderouteindicatortype := range __gong__sortStageSetInstances(stage.CodeRouteIndicatorTypes, stage.CodeRouteIndicatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderouteindicatortypeIdent := coderouteindicatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteIndicatorType{Name: %s}).Stage(stage)", coderouteindicatortypeIdent, __gong__toRawStringLiteral(coderouteindicatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderouteindicatortypeIdent, __gong__toRawStringLiteral(coderouteindicatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderouteindicatortypeIdent, __gong__toRawStringLiteral(coderouteindicatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderouteindicatortypeIdent, __gong__toRawStringLiteral(coderouteindicatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codesafealtitudetype := range __gong__sortStageSetInstances(stage.CodeSafeAltitudeTypes, stage.CodeSafeAltitudeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codesafealtitudetypeIdent := codesafealtitudetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSafeAltitudeType{Name: %s}).Stage(stage)", codesafealtitudetypeIdent, __gong__toRawStringLiteral(codesafealtitudetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesafealtitudetypeIdent, __gong__toRawStringLiteral(codesafealtitudetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesafealtitudetypeIdent, __gong__toRawStringLiteral(codesafealtitudetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesafealtitudetypeIdent, __gong__toRawStringLiteral(codesafealtitudetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codesegmentpathtype := range __gong__sortStageSetInstances(stage.CodeSegmentPathTypes, stage.CodeSegmentPathType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codesegmentpathtypeIdent := codesegmentpathtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSegmentPathType{Name: %s}).Stage(stage)", codesegmentpathtypeIdent, __gong__toRawStringLiteral(codesegmentpathtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesegmentpathtypeIdent, __gong__toRawStringLiteral(codesegmentpathtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesegmentpathtypeIdent, __gong__toRawStringLiteral(codesegmentpathtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesegmentpathtypeIdent, __gong__toRawStringLiteral(codesegmentpathtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codesegmentterminationtype := range __gong__sortStageSetInstances(stage.CodeSegmentTerminationTypes, stage.CodeSegmentTerminationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codesegmentterminationtypeIdent := codesegmentterminationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSegmentTerminationType{Name: %s}).Stage(stage)", codesegmentterminationtypeIdent, __gong__toRawStringLiteral(codesegmentterminationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesegmentterminationtypeIdent, __gong__toRawStringLiteral(codesegmentterminationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesegmentterminationtypeIdent, __gong__toRawStringLiteral(codesegmentterminationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesegmentterminationtypeIdent, __gong__toRawStringLiteral(codesegmentterminationtype.EnclosedText)))
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
		for _, codetaatype := range __gong__sortStageSetInstances(stage.CodeTAATypes, stage.CodeTAAType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetaatypeIdent := codetaatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTAAType{Name: %s}).Stage(stage)", codetaatypeIdent, __gong__toRawStringLiteral(codetaatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetaatypeIdent, __gong__toRawStringLiteral(codetaatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetaatypeIdent, __gong__toRawStringLiteral(codetaatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetaatypeIdent, __gong__toRawStringLiteral(codetaatype.EnclosedText)))
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
		for _, codetrajectorytype := range __gong__sortStageSetInstances(stage.CodeTrajectoryTypes, stage.CodeTrajectoryType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetrajectorytypeIdent := codetrajectorytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTrajectoryType{Name: %s}).Stage(stage)", codetrajectorytypeIdent, __gong__toRawStringLiteral(codetrajectorytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetrajectorytypeIdent, __gong__toRawStringLiteral(codetrajectorytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetrajectorytypeIdent, __gong__toRawStringLiteral(codetrajectorytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetrajectorytypeIdent, __gong__toRawStringLiteral(codetrajectorytype.EnclosedText)))
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
		for _, textsidstardesignatortype := range __gong__sortStageSetInstances(stage.TextSIDSTARDesignatorTypes, stage.TextSIDSTARDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textsidstardesignatortypeIdent := textsidstardesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextSIDSTARDesignatorType{Name: %s}).Stage(stage)", textsidstardesignatortypeIdent, __gong__toRawStringLiteral(textsidstardesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textsidstardesignatortypeIdent, __gong__toRawStringLiteral(textsidstardesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textsidstardesignatortypeIdent, __gong__toRawStringLiteral(textsidstardesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textsidstardesignatortypeIdent, __gong__toRawStringLiteral(textsidstardesignatortype.EnclosedText)))
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
		for _, valalarmlimittype := range __gong__sortStageSetInstances(stage.ValAlarmLimitTypes, stage.ValAlarmLimitType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valalarmlimittypeIdent := valalarmlimittype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValAlarmLimitType{Name: %s}).Stage(stage)", valalarmlimittypeIdent, __gong__toRawStringLiteral(valalarmlimittype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valalarmlimittypeIdent, __gong__toRawStringLiteral(valalarmlimittype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valalarmlimittypeIdent, __gong__toRawStringLiteral(valalarmlimittype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valalarmlimittypeIdent, __gong__toRawStringLiteral(valalarmlimittype.EnclosedText)))
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
		for _, valchannelnumbertype := range __gong__sortStageSetInstances(stage.ValChannelNumberTypes, stage.ValChannelNumberType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valchannelnumbertypeIdent := valchannelnumbertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValChannelNumberType{Name: %s}).Stage(stage)", valchannelnumbertypeIdent, __gong__toRawStringLiteral(valchannelnumbertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valchannelnumbertypeIdent, __gong__toRawStringLiteral(valchannelnumbertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valchannelnumbertypeIdent, __gong__toRawStringLiteral(valchannelnumbertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valchannelnumbertypeIdent, __gong__toRawStringLiteral(valchannelnumbertype.EnclosedText)))
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
		for _, valhextype := range __gong__sortStageSetInstances(stage.ValHexTypes, stage.ValHexType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valhextypeIdent := valhextype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValHexType{Name: %s}).Stage(stage)", valhextypeIdent, __gong__toRawStringLiteral(valhextype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valhextypeIdent, __gong__toRawStringLiteral(valhextype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valhextypeIdent, __gong__toRawStringLiteral(valhextype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valhextypeIdent, __gong__toRawStringLiteral(valhextype.EnclosedText)))
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
		for _, valtemperaturetype := range __gong__sortStageSetInstances(stage.ValTemperatureTypes, stage.ValTemperatureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valtemperaturetypeIdent := valtemperaturetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValTemperatureType{Name: %s}).Stage(stage)", valtemperaturetypeIdent, __gong__toRawStringLiteral(valtemperaturetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valtemperaturetypeIdent, __gong__toRawStringLiteral(valtemperaturetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valtemperaturetypeIdent, __gong__toRawStringLiteral(valtemperaturetype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valtemperaturetypeIdent, __gong__toRawStringLiteral(valtemperaturetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valtemperaturetypeIdent, __gong__toRawStringLiteral(valtemperaturetype.EnclosedText)))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes\"", fmt.Sprintf("%q", modelsPackageName))
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

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeapproachequipmentadditionaltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeapproachequipmentadditionaltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeapproachequipmentadditionaltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeapproachequipmentadditionaltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeApproachEquipmentAdditionalType", fieldName)
	}
	return
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeapproachguidancetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeapproachguidancetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeapproachguidancetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeapproachguidancetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeApproachGuidanceType", fieldName)
	}
	return
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeapproachprefixtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeapproachprefixtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeapproachprefixtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeapproachprefixtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeApproachPrefixType", fieldName)
	}
	return
}

func (codeapproachtype *CodeApproachType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeapproachtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeapproachtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeapproachtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeapproachtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeApproachType", fieldName)
	}
	return
}

func (codeapprovaltype *CodeApprovalType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeapprovaltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeapprovaltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeapprovaltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeapprovaltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeApprovalType", fieldName)
	}
	return
}

func (codearcdirectiontype *CodeArcDirectionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codearcdirectiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codearcdirectiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codearcdirectiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codearcdirectiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeArcDirectionType", fieldName)
	}
	return
}

func (codebearingtype *CodeBearingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codebearingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codebearingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codebearingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codebearingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeBearingType", fieldName)
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

func (codedesignstandardtype *CodeDesignStandardType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedesignstandardtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedesignstandardtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedesignstandardtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedesignstandardtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDesignStandardType", fieldName)
	}
	return
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedesignatedpointdesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedesignatedpointdesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedesignatedpointdesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedesignatedpointdesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDesignatedPointDesignatorType", fieldName)
	}
	return
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedirectionreferencetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedirectionreferencetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedirectionreferencetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedirectionreferencetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDirectionReferenceType", fieldName)
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

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeequipmentunavailabletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeequipmentunavailabletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeequipmentunavailabletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeequipmentunavailabletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeEquipmentUnavailableType", fieldName)
	}
	return
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codefinalguidancetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codefinalguidancetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codefinalguidancetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codefinalguidancetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFinalGuidanceType", fieldName)
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

func (codeheightreferencetype *CodeHeightReferenceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeheightreferencetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeheightreferencetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeheightreferencetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeheightreferencetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeHeightReferenceType", fieldName)
	}
	return
}

func (codeholdingusagetype *CodeHoldingUsageType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeholdingusagetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeholdingusagetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeholdingusagetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeholdingusagetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeHoldingUsageType", fieldName)
	}
	return
}

func (codeholdingusetype *CodeHoldingUseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeholdingusetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeholdingusetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeholdingusetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeholdingusetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeHoldingUseType", fieldName)
	}
	return
}

func (codeicaocountrytype *CodeICAOCountryType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeicaocountrytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeicaocountrytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeicaocountrytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeicaocountrytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeICAOCountryType", fieldName)
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

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeminimafinalapproachpathtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeminimafinalapproachpathtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeminimafinalapproachpathtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeminimafinalapproachpathtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMinimaFinalApproachPathType", fieldName)
	}
	return
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeminimumaltitudetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeminimumaltitudetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeminimumaltitudetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeminimumaltitudetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMinimumAltitudeType", fieldName)
	}
	return
}

func (codeminimumheighttype *CodeMinimumHeightType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeminimumheighttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeminimumheighttype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeminimumheighttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeminimumheighttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMinimumHeightType", fieldName)
	}
	return
}

func (codemissedapproachtype *CodeMissedApproachType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemissedapproachtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemissedapproachtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemissedapproachtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemissedapproachtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMissedApproachType", fieldName)
	}
	return
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codenavigationarearestrictiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codenavigationarearestrictiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codenavigationarearestrictiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codenavigationarearestrictiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNavigationAreaRestrictionType", fieldName)
	}
	return
}

func (codenavigationareatype *CodeNavigationAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codenavigationareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codenavigationareatype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codenavigationareatype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codenavigationareatype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNavigationAreaType", fieldName)
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

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeprocedureavailabilitytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeprocedureavailabilitytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeprocedureavailabilitytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeprocedureavailabilitytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeProcedureAvailabilityType", fieldName)
	}
	return
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeprocedurecodingstandardtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeprocedurecodingstandardtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeprocedurecodingstandardtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeprocedurecodingstandardtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeProcedureCodingStandardType", fieldName)
	}
	return
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeproceduredistancetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeproceduredistancetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeproceduredistancetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeproceduredistancetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeProcedureDistanceType", fieldName)
	}
	return
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeprocedurefixroletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeprocedurefixroletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeprocedurefixroletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeprocedurefixroletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeProcedureFixRoleType", fieldName)
	}
	return
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeprocedurephasetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeprocedurephasetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeprocedurephasetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeprocedurephasetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeProcedurePhaseType", fieldName)
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

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codereferencepathidentifiertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codereferencepathidentifiertype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codereferencepathidentifiertype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codereferencepathidentifiertype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeReferencePathIdentifierType", fieldName)
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

func (coderelativepositiontype *CodeRelativePositionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderelativepositiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderelativepositiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderelativepositiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderelativepositiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRelativePositionType", fieldName)
	}
	return
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderouteindicatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderouteindicatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderouteindicatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderouteindicatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRouteIndicatorType", fieldName)
	}
	return
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codesafealtitudetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codesafealtitudetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codesafealtitudetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codesafealtitudetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSafeAltitudeType", fieldName)
	}
	return
}

func (codesegmentpathtype *CodeSegmentPathType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codesegmentpathtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codesegmentpathtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codesegmentpathtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codesegmentpathtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSegmentPathType", fieldName)
	}
	return
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codesegmentterminationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codesegmentterminationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codesegmentterminationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codesegmentterminationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSegmentTerminationType", fieldName)
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

func (codetaatype *CodeTAAType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetaatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetaatype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetaatype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetaatype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTAAType", fieldName)
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

func (codetrajectorytype *CodeTrajectoryType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetrajectorytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetrajectorytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetrajectorytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetrajectorytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTrajectoryType", fieldName)
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

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := textsidstardesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", textsidstardesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", textsidstardesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", textsidstardesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TextSIDSTARDesignatorType", fieldName)
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

func (valalarmlimittype *ValAlarmLimitType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valalarmlimittype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valalarmlimittype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valalarmlimittype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valalarmlimittype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValAlarmLimitType", fieldName)
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

func (valchannelnumbertype *ValChannelNumberType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valchannelnumbertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valchannelnumbertype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valchannelnumbertype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valchannelnumbertype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValChannelNumberType", fieldName)
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

func (valhextype *ValHexType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valhextype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valhextype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valhextype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valhextype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValHexType", fieldName)
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

func (valtemperaturetype *ValTemperatureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valtemperaturetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valtemperaturetype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valtemperaturetype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valtemperaturetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valtemperaturetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValTemperatureType", fieldName)
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
func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeapproachequipmentadditionaltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeapproachequipmentadditionaltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeapproachequipmentadditionaltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeapproachguidancetype *CodeApproachGuidanceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeapproachguidancetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeapproachguidancetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeapproachguidancetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeapproachprefixtype *CodeApproachPrefixType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeapproachprefixtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeapproachprefixtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeapproachprefixtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeapproachtype *CodeApproachType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeapproachtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeapproachtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeapproachtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeapprovaltype *CodeApprovalType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeapprovaltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeapprovaltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeapprovaltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codearcdirectiontype *CodeArcDirectionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codearcdirectiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codearcdirectiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codearcdirectiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codebearingtype *CodeBearingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codebearingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codebearingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codebearingtype.GongMarshallField(stage, "EnclosedText"))
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
func (codedesignstandardtype *CodeDesignStandardType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedesignstandardtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedesignstandardtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedesignstandardtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedesignatedpointdesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedesignatedpointdesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedesignatedpointdesignatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codedirectionreferencetype *CodeDirectionReferenceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedirectionreferencetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedirectionreferencetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedirectionreferencetype.GongMarshallField(stage, "EnclosedText"))
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
func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeequipmentunavailabletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeequipmentunavailabletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeequipmentunavailabletype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codefinalguidancetype *CodeFinalGuidanceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codefinalguidancetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codefinalguidancetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codefinalguidancetype.GongMarshallField(stage, "EnclosedText"))
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
func (codeheightreferencetype *CodeHeightReferenceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeheightreferencetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeheightreferencetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeheightreferencetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeholdingusagetype *CodeHoldingUsageType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeholdingusagetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeholdingusagetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeholdingusagetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeholdingusetype *CodeHoldingUseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeholdingusetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeholdingusetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeholdingusetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeicaocountrytype *CodeICAOCountryType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeicaocountrytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeicaocountrytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeicaocountrytype.GongMarshallField(stage, "EnclosedText"))
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
func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeminimafinalapproachpathtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeminimafinalapproachpathtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeminimafinalapproachpathtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeminimumaltitudetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeminimumaltitudetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeminimumaltitudetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeminimumheighttype *CodeMinimumHeightType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeminimumheighttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeminimumheighttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeminimumheighttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codemissedapproachtype *CodeMissedApproachType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemissedapproachtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemissedapproachtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemissedapproachtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codenavigationarearestrictiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codenavigationarearestrictiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codenavigationarearestrictiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codenavigationareatype *CodeNavigationAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codenavigationareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codenavigationareatype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codenavigationareatype.GongMarshallField(stage, "EnclosedText"))
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
func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeprocedureavailabilitytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeprocedureavailabilitytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeprocedureavailabilitytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeprocedurecodingstandardtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeprocedurecodingstandardtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeprocedurecodingstandardtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeproceduredistancetype *CodeProcedureDistanceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeproceduredistancetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeproceduredistancetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeproceduredistancetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeprocedurefixroletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeprocedurefixroletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeprocedurefixroletype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeprocedurephasetype *CodeProcedurePhaseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeprocedurephasetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeprocedurephasetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeprocedurephasetype.GongMarshallField(stage, "EnclosedText"))
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
func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codereferencepathidentifiertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codereferencepathidentifiertype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codereferencepathidentifiertype.GongMarshallField(stage, "EnclosedText"))
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
func (coderelativepositiontype *CodeRelativePositionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderelativepositiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderelativepositiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderelativepositiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderouteindicatortype *CodeRouteIndicatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderouteindicatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderouteindicatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderouteindicatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codesafealtitudetype *CodeSafeAltitudeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codesafealtitudetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codesafealtitudetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codesafealtitudetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codesegmentpathtype *CodeSegmentPathType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codesegmentpathtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codesegmentpathtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codesegmentpathtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codesegmentterminationtype *CodeSegmentTerminationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codesegmentterminationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codesegmentterminationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codesegmentterminationtype.GongMarshallField(stage, "EnclosedText"))
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
func (codetaatype *CodeTAAType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetaatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetaatype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetaatype.GongMarshallField(stage, "EnclosedText"))
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
func (codetrajectorytype *CodeTrajectoryType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetrajectorytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetrajectorytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetrajectorytype.GongMarshallField(stage, "EnclosedText"))
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
func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(textsidstardesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(textsidstardesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(textsidstardesignatortype.GongMarshallField(stage, "EnclosedText"))
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
func (valalarmlimittype *ValAlarmLimitType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valalarmlimittype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valalarmlimittype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valalarmlimittype.GongMarshallField(stage, "EnclosedText"))
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
func (valchannelnumbertype *ValChannelNumberType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valchannelnumbertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valchannelnumbertype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valchannelnumbertype.GongMarshallField(stage, "EnclosedText"))
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
func (valhextype *ValHexType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valhextype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valhextype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valhextype.GongMarshallField(stage, "EnclosedText"))
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
func (valtemperaturetype *ValTemperatureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valtemperaturetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valtemperaturetype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valtemperaturetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valtemperaturetype.GongMarshallField(stage, "EnclosedText"))
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
