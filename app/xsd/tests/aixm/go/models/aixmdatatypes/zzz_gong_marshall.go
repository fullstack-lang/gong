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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmdatatypes"
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
		for _, codeaircraftgroundservicetype := range __gong__sortStageSetInstances(stage.CodeAircraftGroundServiceTypes, stage.CodeAircraftGroundServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaircraftgroundservicetypeIdent := codeaircraftgroundservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftGroundServiceType{Name: %s}).Stage(stage)", codeaircraftgroundservicetypeIdent, __gong__toRawStringLiteral(codeaircraftgroundservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftgroundservicetypeIdent, __gong__toRawStringLiteral(codeaircraftgroundservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftgroundservicetypeIdent, __gong__toRawStringLiteral(codeaircraftgroundservicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftgroundservicetypeIdent, __gong__toRawStringLiteral(codeaircraftgroundservicetype.EnclosedText)))
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
		for _, codeaircraftstandtype := range __gong__sortStageSetInstances(stage.CodeAircraftStandTypes, stage.CodeAircraftStandType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaircraftstandtypeIdent := codeaircraftstandtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftStandType{Name: %s}).Stage(stage)", codeaircraftstandtypeIdent, __gong__toRawStringLiteral(codeaircraftstandtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftstandtypeIdent, __gong__toRawStringLiteral(codeaircraftstandtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftstandtypeIdent, __gong__toRawStringLiteral(codeaircraftstandtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftstandtypeIdent, __gong__toRawStringLiteral(codeaircraftstandtype.EnclosedText)))
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
		for _, codeairportheliportcollocationtype := range __gong__sortStageSetInstances(stage.CodeAirportHeliportCollocationTypes, stage.CodeAirportHeliportCollocationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairportheliportcollocationtypeIdent := codeairportheliportcollocationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirportHeliportCollocationType{Name: %s}).Stage(stage)", codeairportheliportcollocationtypeIdent, __gong__toRawStringLiteral(codeairportheliportcollocationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairportheliportcollocationtypeIdent, __gong__toRawStringLiteral(codeairportheliportcollocationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairportheliportcollocationtypeIdent, __gong__toRawStringLiteral(codeairportheliportcollocationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairportheliportcollocationtypeIdent, __gong__toRawStringLiteral(codeairportheliportcollocationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairportheliportdesignatortype := range __gong__sortStageSetInstances(stage.CodeAirportHeliportDesignatorTypes, stage.CodeAirportHeliportDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairportheliportdesignatortypeIdent := codeairportheliportdesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirportHeliportDesignatorType{Name: %s}).Stage(stage)", codeairportheliportdesignatortypeIdent, __gong__toRawStringLiteral(codeairportheliportdesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairportheliportdesignatortypeIdent, __gong__toRawStringLiteral(codeairportheliportdesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairportheliportdesignatortypeIdent, __gong__toRawStringLiteral(codeairportheliportdesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairportheliportdesignatortypeIdent, __gong__toRawStringLiteral(codeairportheliportdesignatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairportheliporttype := range __gong__sortStageSetInstances(stage.CodeAirportHeliportTypes, stage.CodeAirportHeliportType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairportheliporttypeIdent := codeairportheliporttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirportHeliportType{Name: %s}).Stage(stage)", codeairportheliporttypeIdent, __gong__toRawStringLiteral(codeairportheliporttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairportheliporttypeIdent, __gong__toRawStringLiteral(codeairportheliporttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairportheliporttypeIdent, __gong__toRawStringLiteral(codeairportheliporttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairportheliporttypeIdent, __gong__toRawStringLiteral(codeairportheliporttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairportwarningtype := range __gong__sortStageSetInstances(stage.CodeAirportWarningTypes, stage.CodeAirportWarningType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairportwarningtypeIdent := codeairportwarningtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirportWarningType{Name: %s}).Stage(stage)", codeairportwarningtypeIdent, __gong__toRawStringLiteral(codeairportwarningtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairportwarningtypeIdent, __gong__toRawStringLiteral(codeairportwarningtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairportwarningtypeIdent, __gong__toRawStringLiteral(codeairportwarningtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairportwarningtypeIdent, __gong__toRawStringLiteral(codeairportwarningtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairspaceactivitytype := range __gong__sortStageSetInstances(stage.CodeAirspaceActivityTypes, stage.CodeAirspaceActivityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairspaceactivitytypeIdent := codeairspaceactivitytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceActivityType{Name: %s}).Stage(stage)", codeairspaceactivitytypeIdent, __gong__toRawStringLiteral(codeairspaceactivitytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspaceactivitytypeIdent, __gong__toRawStringLiteral(codeairspaceactivitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspaceactivitytypeIdent, __gong__toRawStringLiteral(codeairspaceactivitytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspaceactivitytypeIdent, __gong__toRawStringLiteral(codeairspaceactivitytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairspaceaggregationtype := range __gong__sortStageSetInstances(stage.CodeAirspaceAggregationTypes, stage.CodeAirspaceAggregationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairspaceaggregationtypeIdent := codeairspaceaggregationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceAggregationType{Name: %s}).Stage(stage)", codeairspaceaggregationtypeIdent, __gong__toRawStringLiteral(codeairspaceaggregationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspaceaggregationtypeIdent, __gong__toRawStringLiteral(codeairspaceaggregationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspaceaggregationtypeIdent, __gong__toRawStringLiteral(codeairspaceaggregationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspaceaggregationtypeIdent, __gong__toRawStringLiteral(codeairspaceaggregationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairspaceclassificationtype := range __gong__sortStageSetInstances(stage.CodeAirspaceClassificationTypes, stage.CodeAirspaceClassificationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairspaceclassificationtypeIdent := codeairspaceclassificationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceClassificationType{Name: %s}).Stage(stage)", codeairspaceclassificationtypeIdent, __gong__toRawStringLiteral(codeairspaceclassificationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspaceclassificationtypeIdent, __gong__toRawStringLiteral(codeairspaceclassificationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspaceclassificationtypeIdent, __gong__toRawStringLiteral(codeairspaceclassificationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspaceclassificationtypeIdent, __gong__toRawStringLiteral(codeairspaceclassificationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairspacedependencytype := range __gong__sortStageSetInstances(stage.CodeAirspaceDependencyTypes, stage.CodeAirspaceDependencyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairspacedependencytypeIdent := codeairspacedependencytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceDependencyType{Name: %s}).Stage(stage)", codeairspacedependencytypeIdent, __gong__toRawStringLiteral(codeairspacedependencytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspacedependencytypeIdent, __gong__toRawStringLiteral(codeairspacedependencytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspacedependencytypeIdent, __gong__toRawStringLiteral(codeairspacedependencytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspacedependencytypeIdent, __gong__toRawStringLiteral(codeairspacedependencytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairspacedesignatortype := range __gong__sortStageSetInstances(stage.CodeAirspaceDesignatorTypes, stage.CodeAirspaceDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairspacedesignatortypeIdent := codeairspacedesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceDesignatorType{Name: %s}).Stage(stage)", codeairspacedesignatortypeIdent, __gong__toRawStringLiteral(codeairspacedesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspacedesignatortypeIdent, __gong__toRawStringLiteral(codeairspacedesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspacedesignatortypeIdent, __gong__toRawStringLiteral(codeairspacedesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspacedesignatortypeIdent, __gong__toRawStringLiteral(codeairspacedesignatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairspacepointpositiontype := range __gong__sortStageSetInstances(stage.CodeAirspacePointPositionTypes, stage.CodeAirspacePointPositionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairspacepointpositiontypeIdent := codeairspacepointpositiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspacePointPositionType{Name: %s}).Stage(stage)", codeairspacepointpositiontypeIdent, __gong__toRawStringLiteral(codeairspacepointpositiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspacepointpositiontypeIdent, __gong__toRawStringLiteral(codeairspacepointpositiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspacepointpositiontypeIdent, __gong__toRawStringLiteral(codeairspacepointpositiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspacepointpositiontypeIdent, __gong__toRawStringLiteral(codeairspacepointpositiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairspacepointroletype := range __gong__sortStageSetInstances(stage.CodeAirspacePointRoleTypes, stage.CodeAirspacePointRoleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairspacepointroletypeIdent := codeairspacepointroletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspacePointRoleType{Name: %s}).Stage(stage)", codeairspacepointroletypeIdent, __gong__toRawStringLiteral(codeairspacepointroletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspacepointroletypeIdent, __gong__toRawStringLiteral(codeairspacepointroletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspacepointroletypeIdent, __gong__toRawStringLiteral(codeairspacepointroletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspacepointroletypeIdent, __gong__toRawStringLiteral(codeairspacepointroletype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeairspacetype := range __gong__sortStageSetInstances(stage.CodeAirspaceTypes, stage.CodeAirspaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeairspacetypeIdent := codeairspacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceType{Name: %s}).Stage(stage)", codeairspacetypeIdent, __gong__toRawStringLiteral(codeairspacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspacetypeIdent, __gong__toRawStringLiteral(codeairspacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspacetypeIdent, __gong__toRawStringLiteral(codeairspacetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspacetypeIdent, __gong__toRawStringLiteral(codeairspacetype.EnclosedText)))
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
		for _, codeapproachlightingicaotype := range __gong__sortStageSetInstances(stage.CodeApproachLightingICAOTypes, stage.CodeApproachLightingICAOType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeapproachlightingicaotypeIdent := codeapproachlightingicaotype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeApproachLightingICAOType{Name: %s}).Stage(stage)", codeapproachlightingicaotypeIdent, __gong__toRawStringLiteral(codeapproachlightingicaotype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeapproachlightingicaotypeIdent, __gong__toRawStringLiteral(codeapproachlightingicaotype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeapproachlightingicaotypeIdent, __gong__toRawStringLiteral(codeapproachlightingicaotype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeapproachlightingicaotypeIdent, __gong__toRawStringLiteral(codeapproachlightingicaotype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeapproachlightingtype := range __gong__sortStageSetInstances(stage.CodeApproachLightingTypes, stage.CodeApproachLightingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeapproachlightingtypeIdent := codeapproachlightingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeApproachLightingType{Name: %s}).Stage(stage)", codeapproachlightingtypeIdent, __gong__toRawStringLiteral(codeapproachlightingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeapproachlightingtypeIdent, __gong__toRawStringLiteral(codeapproachlightingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeapproachlightingtypeIdent, __gong__toRawStringLiteral(codeapproachlightingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeapproachlightingtypeIdent, __gong__toRawStringLiteral(codeapproachlightingtype.EnclosedText)))
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
		for _, codeapronelementtype := range __gong__sortStageSetInstances(stage.CodeApronElementTypes, stage.CodeApronElementType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeapronelementtypeIdent := codeapronelementtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeApronElementType{Name: %s}).Stage(stage)", codeapronelementtypeIdent, __gong__toRawStringLiteral(codeapronelementtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeapronelementtypeIdent, __gong__toRawStringLiteral(codeapronelementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeapronelementtypeIdent, __gong__toRawStringLiteral(codeapronelementtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeapronelementtypeIdent, __gong__toRawStringLiteral(codeapronelementtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeapronsectiontype := range __gong__sortStageSetInstances(stage.CodeApronSectionTypes, stage.CodeApronSectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeapronsectiontypeIdent := codeapronsectiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeApronSectionType{Name: %s}).Stage(stage)", codeapronsectiontypeIdent, __gong__toRawStringLiteral(codeapronsectiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeapronsectiontypeIdent, __gong__toRawStringLiteral(codeapronsectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeapronsectiontypeIdent, __gong__toRawStringLiteral(codeapronsectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeapronsectiontypeIdent, __gong__toRawStringLiteral(codeapronsectiontype.EnclosedText)))
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
		for _, codearrestinggearenergyabsorbtype := range __gong__sortStageSetInstances(stage.CodeArrestingGearEnergyAbsorbTypes, stage.CodeArrestingGearEnergyAbsorbType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codearrestinggearenergyabsorbtypeIdent := codearrestinggearenergyabsorbtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeArrestingGearEnergyAbsorbType{Name: %s}).Stage(stage)", codearrestinggearenergyabsorbtypeIdent, __gong__toRawStringLiteral(codearrestinggearenergyabsorbtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codearrestinggearenergyabsorbtypeIdent, __gong__toRawStringLiteral(codearrestinggearenergyabsorbtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codearrestinggearenergyabsorbtypeIdent, __gong__toRawStringLiteral(codearrestinggearenergyabsorbtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codearrestinggearenergyabsorbtypeIdent, __gong__toRawStringLiteral(codearrestinggearenergyabsorbtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codearrestinggearengagedevicetype := range __gong__sortStageSetInstances(stage.CodeArrestingGearEngageDeviceTypes, stage.CodeArrestingGearEngageDeviceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codearrestinggearengagedevicetypeIdent := codearrestinggearengagedevicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeArrestingGearEngageDeviceType{Name: %s}).Stage(stage)", codearrestinggearengagedevicetypeIdent, __gong__toRawStringLiteral(codearrestinggearengagedevicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codearrestinggearengagedevicetypeIdent, __gong__toRawStringLiteral(codearrestinggearengagedevicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codearrestinggearengagedevicetypeIdent, __gong__toRawStringLiteral(codearrestinggearengagedevicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codearrestinggearengagedevicetypeIdent, __gong__toRawStringLiteral(codearrestinggearengagedevicetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeauralmorsetype := range __gong__sortStageSetInstances(stage.CodeAuralMorseTypes, stage.CodeAuralMorseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeauralmorsetypeIdent := codeauralmorsetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAuralMorseType{Name: %s}).Stage(stage)", codeauralmorsetypeIdent, __gong__toRawStringLiteral(codeauralmorsetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeauralmorsetypeIdent, __gong__toRawStringLiteral(codeauralmorsetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeauralmorsetypeIdent, __gong__toRawStringLiteral(codeauralmorsetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeauralmorsetypeIdent, __gong__toRawStringLiteral(codeauralmorsetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeauthorityroletype := range __gong__sortStageSetInstances(stage.CodeAuthorityRoleTypes, stage.CodeAuthorityRoleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeauthorityroletypeIdent := codeauthorityroletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAuthorityRoleType{Name: %s}).Stage(stage)", codeauthorityroletypeIdent, __gong__toRawStringLiteral(codeauthorityroletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeauthorityroletypeIdent, __gong__toRawStringLiteral(codeauthorityroletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeauthorityroletypeIdent, __gong__toRawStringLiteral(codeauthorityroletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeauthorityroletypeIdent, __gong__toRawStringLiteral(codeauthorityroletype.EnclosedText)))
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
		for _, codeaviationstandardstype := range __gong__sortStageSetInstances(stage.CodeAviationStandardsTypes, stage.CodeAviationStandardsType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeaviationstandardstypeIdent := codeaviationstandardstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAviationStandardsType{Name: %s}).Stage(stage)", codeaviationstandardstypeIdent, __gong__toRawStringLiteral(codeaviationstandardstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaviationstandardstypeIdent, __gong__toRawStringLiteral(codeaviationstandardstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaviationstandardstypeIdent, __gong__toRawStringLiteral(codeaviationstandardstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaviationstandardstypeIdent, __gong__toRawStringLiteral(codeaviationstandardstype.EnclosedText)))
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
		for _, codebuoydesignatortype := range __gong__sortStageSetInstances(stage.CodeBuoyDesignatorTypes, stage.CodeBuoyDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codebuoydesignatortypeIdent := codebuoydesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeBuoyDesignatorType{Name: %s}).Stage(stage)", codebuoydesignatortypeIdent, __gong__toRawStringLiteral(codebuoydesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codebuoydesignatortypeIdent, __gong__toRawStringLiteral(codebuoydesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codebuoydesignatortypeIdent, __gong__toRawStringLiteral(codebuoydesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codebuoydesignatortypeIdent, __gong__toRawStringLiteral(codebuoydesignatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codebuoytype := range __gong__sortStageSetInstances(stage.CodeBuoyTypes, stage.CodeBuoyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codebuoytypeIdent := codebuoytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeBuoyType{Name: %s}).Stage(stage)", codebuoytypeIdent, __gong__toRawStringLiteral(codebuoytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codebuoytypeIdent, __gong__toRawStringLiteral(codebuoytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codebuoytypeIdent, __gong__toRawStringLiteral(codebuoytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codebuoytypeIdent, __gong__toRawStringLiteral(codebuoytype.EnclosedText)))
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
		for _, codecheckpointcategorytype := range __gong__sortStageSetInstances(stage.CodeCheckpointCategoryTypes, stage.CodeCheckpointCategoryType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codecheckpointcategorytypeIdent := codecheckpointcategorytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCheckpointCategoryType{Name: %s}).Stage(stage)", codecheckpointcategorytypeIdent, __gong__toRawStringLiteral(codecheckpointcategorytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecheckpointcategorytypeIdent, __gong__toRawStringLiteral(codecheckpointcategorytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecheckpointcategorytypeIdent, __gong__toRawStringLiteral(codecheckpointcategorytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecheckpointcategorytypeIdent, __gong__toRawStringLiteral(codecheckpointcategorytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codecolourtype := range __gong__sortStageSetInstances(stage.CodeColourTypes, stage.CodeColourType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codecolourtypeIdent := codecolourtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeColourType{Name: %s}).Stage(stage)", codecolourtypeIdent, __gong__toRawStringLiteral(codecolourtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecolourtypeIdent, __gong__toRawStringLiteral(codecolourtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecolourtypeIdent, __gong__toRawStringLiteral(codecolourtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecolourtypeIdent, __gong__toRawStringLiteral(codecolourtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codecommunicationchanneltype := range __gong__sortStageSetInstances(stage.CodeCommunicationChannelTypes, stage.CodeCommunicationChannelType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codecommunicationchanneltypeIdent := codecommunicationchanneltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCommunicationChannelType{Name: %s}).Stage(stage)", codecommunicationchanneltypeIdent, __gong__toRawStringLiteral(codecommunicationchanneltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecommunicationchanneltypeIdent, __gong__toRawStringLiteral(codecommunicationchanneltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecommunicationchanneltypeIdent, __gong__toRawStringLiteral(codecommunicationchanneltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecommunicationchanneltypeIdent, __gong__toRawStringLiteral(codecommunicationchanneltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codecommunicationdirectiontype := range __gong__sortStageSetInstances(stage.CodeCommunicationDirectionTypes, stage.CodeCommunicationDirectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codecommunicationdirectiontypeIdent := codecommunicationdirectiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCommunicationDirectionType{Name: %s}).Stage(stage)", codecommunicationdirectiontypeIdent, __gong__toRawStringLiteral(codecommunicationdirectiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecommunicationdirectiontypeIdent, __gong__toRawStringLiteral(codecommunicationdirectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecommunicationdirectiontypeIdent, __gong__toRawStringLiteral(codecommunicationdirectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecommunicationdirectiontypeIdent, __gong__toRawStringLiteral(codecommunicationdirectiontype.EnclosedText)))
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
		for _, codecontaminationtype := range __gong__sortStageSetInstances(stage.CodeContaminationTypes, stage.CodeContaminationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codecontaminationtypeIdent := codecontaminationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeContaminationType{Name: %s}).Stage(stage)", codecontaminationtypeIdent, __gong__toRawStringLiteral(codecontaminationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecontaminationtypeIdent, __gong__toRawStringLiteral(codecontaminationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecontaminationtypeIdent, __gong__toRawStringLiteral(codecontaminationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecontaminationtypeIdent, __gong__toRawStringLiteral(codecontaminationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codecoursequalityilstype := range __gong__sortStageSetInstances(stage.CodeCourseQualityILSTypes, stage.CodeCourseQualityILSType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codecoursequalityilstypeIdent := codecoursequalityilstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCourseQualityILSType{Name: %s}).Stage(stage)", codecoursequalityilstypeIdent, __gong__toRawStringLiteral(codecoursequalityilstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecoursequalityilstypeIdent, __gong__toRawStringLiteral(codecoursequalityilstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecoursequalityilstypeIdent, __gong__toRawStringLiteral(codecoursequalityilstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecoursequalityilstypeIdent, __gong__toRawStringLiteral(codecoursequalityilstype.EnclosedText)))
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
		for _, codedmechanneltype := range __gong__sortStageSetInstances(stage.CodeDMEChannelTypes, stage.CodeDMEChannelType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedmechanneltypeIdent := codedmechanneltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDMEChannelType{Name: %s}).Stage(stage)", codedmechanneltypeIdent, __gong__toRawStringLiteral(codedmechanneltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedmechanneltypeIdent, __gong__toRawStringLiteral(codedmechanneltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedmechanneltypeIdent, __gong__toRawStringLiteral(codedmechanneltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedmechanneltypeIdent, __gong__toRawStringLiteral(codedmechanneltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codedmetype := range __gong__sortStageSetInstances(stage.CodeDMETypes, stage.CodeDMEType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedmetypeIdent := codedmetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDMEType{Name: %s}).Stage(stage)", codedmetypeIdent, __gong__toRawStringLiteral(codedmetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedmetypeIdent, __gong__toRawStringLiteral(codedmetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedmetypeIdent, __gong__toRawStringLiteral(codedmetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedmetypeIdent, __gong__toRawStringLiteral(codedmetype.EnclosedText)))
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
		for _, codedeclareddistancetype := range __gong__sortStageSetInstances(stage.CodeDeclaredDistanceTypes, stage.CodeDeclaredDistanceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedeclareddistancetypeIdent := codedeclareddistancetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDeclaredDistanceType{Name: %s}).Stage(stage)", codedeclareddistancetypeIdent, __gong__toRawStringLiteral(codedeclareddistancetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedeclareddistancetypeIdent, __gong__toRawStringLiteral(codedeclareddistancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedeclareddistancetypeIdent, __gong__toRawStringLiteral(codedeclareddistancetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedeclareddistancetypeIdent, __gong__toRawStringLiteral(codedeclareddistancetype.EnclosedText)))
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
		for _, codedesignatedpointtype := range __gong__sortStageSetInstances(stage.CodeDesignatedPointTypes, stage.CodeDesignatedPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedesignatedpointtypeIdent := codedesignatedpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDesignatedPointType{Name: %s}).Stage(stage)", codedesignatedpointtypeIdent, __gong__toRawStringLiteral(codedesignatedpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedesignatedpointtypeIdent, __gong__toRawStringLiteral(codedesignatedpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedesignatedpointtypeIdent, __gong__toRawStringLiteral(codedesignatedpointtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedesignatedpointtypeIdent, __gong__toRawStringLiteral(codedesignatedpointtype.EnclosedText)))
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
		for _, codedistanceindicationtype := range __gong__sortStageSetInstances(stage.CodeDistanceIndicationTypes, stage.CodeDistanceIndicationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codedistanceindicationtypeIdent := codedistanceindicationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDistanceIndicationType{Name: %s}).Stage(stage)", codedistanceindicationtypeIdent, __gong__toRawStringLiteral(codedistanceindicationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedistanceindicationtypeIdent, __gong__toRawStringLiteral(codedistanceindicationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedistanceindicationtypeIdent, __gong__toRawStringLiteral(codedistanceindicationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedistanceindicationtypeIdent, __gong__toRawStringLiteral(codedistanceindicationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeemissionbandtype := range __gong__sortStageSetInstances(stage.CodeEmissionBandTypes, stage.CodeEmissionBandType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeemissionbandtypeIdent := codeemissionbandtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeEmissionBandType{Name: %s}).Stage(stage)", codeemissionbandtypeIdent, __gong__toRawStringLiteral(codeemissionbandtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeemissionbandtypeIdent, __gong__toRawStringLiteral(codeemissionbandtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeemissionbandtypeIdent, __gong__toRawStringLiteral(codeemissionbandtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeemissionbandtypeIdent, __gong__toRawStringLiteral(codeemissionbandtype.EnclosedText)))
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
		for _, codefacilityrankingtype := range __gong__sortStageSetInstances(stage.CodeFacilityRankingTypes, stage.CodeFacilityRankingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codefacilityrankingtypeIdent := codefacilityrankingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFacilityRankingType{Name: %s}).Stage(stage)", codefacilityrankingtypeIdent, __gong__toRawStringLiteral(codefacilityrankingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefacilityrankingtypeIdent, __gong__toRawStringLiteral(codefacilityrankingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefacilityrankingtypeIdent, __gong__toRawStringLiteral(codefacilityrankingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefacilityrankingtypeIdent, __gong__toRawStringLiteral(codefacilityrankingtype.EnclosedText)))
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
		for _, codefirefightingtype := range __gong__sortStageSetInstances(stage.CodeFireFightingTypes, stage.CodeFireFightingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codefirefightingtypeIdent := codefirefightingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFireFightingType{Name: %s}).Stage(stage)", codefirefightingtypeIdent, __gong__toRawStringLiteral(codefirefightingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefirefightingtypeIdent, __gong__toRawStringLiteral(codefirefightingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefirefightingtypeIdent, __gong__toRawStringLiteral(codefirefightingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefirefightingtypeIdent, __gong__toRawStringLiteral(codefirefightingtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeflightdestinationtype := range __gong__sortStageSetInstances(stage.CodeFlightDestinationTypes, stage.CodeFlightDestinationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeflightdestinationtypeIdent := codeflightdestinationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightDestinationType{Name: %s}).Stage(stage)", codeflightdestinationtypeIdent, __gong__toRawStringLiteral(codeflightdestinationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightdestinationtypeIdent, __gong__toRawStringLiteral(codeflightdestinationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightdestinationtypeIdent, __gong__toRawStringLiteral(codeflightdestinationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightdestinationtypeIdent, __gong__toRawStringLiteral(codeflightdestinationtype.EnclosedText)))
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
		for _, codefrictiondevicetype := range __gong__sortStageSetInstances(stage.CodeFrictionDeviceTypes, stage.CodeFrictionDeviceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codefrictiondevicetypeIdent := codefrictiondevicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFrictionDeviceType{Name: %s}).Stage(stage)", codefrictiondevicetypeIdent, __gong__toRawStringLiteral(codefrictiondevicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefrictiondevicetypeIdent, __gong__toRawStringLiteral(codefrictiondevicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefrictiondevicetypeIdent, __gong__toRawStringLiteral(codefrictiondevicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefrictiondevicetypeIdent, __gong__toRawStringLiteral(codefrictiondevicetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codefrictionestimatetype := range __gong__sortStageSetInstances(stage.CodeFrictionEstimateTypes, stage.CodeFrictionEstimateType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codefrictionestimatetypeIdent := codefrictionestimatetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFrictionEstimateType{Name: %s}).Stage(stage)", codefrictionestimatetypeIdent, __gong__toRawStringLiteral(codefrictionestimatetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefrictionestimatetypeIdent, __gong__toRawStringLiteral(codefrictionestimatetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefrictionestimatetypeIdent, __gong__toRawStringLiteral(codefrictionestimatetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefrictionestimatetypeIdent, __gong__toRawStringLiteral(codefrictionestimatetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codefueltype := range __gong__sortStageSetInstances(stage.CodeFuelTypes, stage.CodeFuelType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codefueltypeIdent := codefueltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFuelType{Name: %s}).Stage(stage)", codefueltypeIdent, __gong__toRawStringLiteral(codefueltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefueltypeIdent, __gong__toRawStringLiteral(codefueltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefueltypeIdent, __gong__toRawStringLiteral(codefueltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefueltypeIdent, __gong__toRawStringLiteral(codefueltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codegeobordertype := range __gong__sortStageSetInstances(stage.CodeGeoBorderTypes, stage.CodeGeoBorderType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codegeobordertypeIdent := codegeobordertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeGeoBorderType{Name: %s}).Stage(stage)", codegeobordertypeIdent, __gong__toRawStringLiteral(codegeobordertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codegeobordertypeIdent, __gong__toRawStringLiteral(codegeobordertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codegeobordertypeIdent, __gong__toRawStringLiteral(codegeobordertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codegeobordertypeIdent, __gong__toRawStringLiteral(codegeobordertype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codegradeseparationtype := range __gong__sortStageSetInstances(stage.CodeGradeSeparationTypes, stage.CodeGradeSeparationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codegradeseparationtypeIdent := codegradeseparationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeGradeSeparationType{Name: %s}).Stage(stage)", codegradeseparationtypeIdent, __gong__toRawStringLiteral(codegradeseparationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codegradeseparationtypeIdent, __gong__toRawStringLiteral(codegradeseparationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codegradeseparationtypeIdent, __gong__toRawStringLiteral(codegradeseparationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codegradeseparationtypeIdent, __gong__toRawStringLiteral(codegradeseparationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codegroundlightingtype := range __gong__sortStageSetInstances(stage.CodeGroundLightingTypes, stage.CodeGroundLightingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codegroundlightingtypeIdent := codegroundlightingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeGroundLightingType{Name: %s}).Stage(stage)", codegroundlightingtypeIdent, __gong__toRawStringLiteral(codegroundlightingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codegroundlightingtypeIdent, __gong__toRawStringLiteral(codegroundlightingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codegroundlightingtypeIdent, __gong__toRawStringLiteral(codegroundlightingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codegroundlightingtypeIdent, __gong__toRawStringLiteral(codegroundlightingtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeguidancelinetype := range __gong__sortStageSetInstances(stage.CodeGuidanceLineTypes, stage.CodeGuidanceLineType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeguidancelinetypeIdent := codeguidancelinetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeGuidanceLineType{Name: %s}).Stage(stage)", codeguidancelinetypeIdent, __gong__toRawStringLiteral(codeguidancelinetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeguidancelinetypeIdent, __gong__toRawStringLiteral(codeguidancelinetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeguidancelinetypeIdent, __gong__toRawStringLiteral(codeguidancelinetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeguidancelinetypeIdent, __gong__toRawStringLiteral(codeguidancelinetype.EnclosedText)))
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
		for _, codehelicopterperformancetype := range __gong__sortStageSetInstances(stage.CodeHelicopterPerformanceTypes, stage.CodeHelicopterPerformanceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codehelicopterperformancetypeIdent := codehelicopterperformancetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeHelicopterPerformanceType{Name: %s}).Stage(stage)", codehelicopterperformancetypeIdent, __gong__toRawStringLiteral(codehelicopterperformancetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codehelicopterperformancetypeIdent, __gong__toRawStringLiteral(codehelicopterperformancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codehelicopterperformancetypeIdent, __gong__toRawStringLiteral(codehelicopterperformancetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codehelicopterperformancetypeIdent, __gong__toRawStringLiteral(codehelicopterperformancetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeholdingcategorytype := range __gong__sortStageSetInstances(stage.CodeHoldingCategoryTypes, stage.CodeHoldingCategoryType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeholdingcategorytypeIdent := codeholdingcategorytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeHoldingCategoryType{Name: %s}).Stage(stage)", codeholdingcategorytypeIdent, __gong__toRawStringLiteral(codeholdingcategorytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeholdingcategorytypeIdent, __gong__toRawStringLiteral(codeholdingcategorytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeholdingcategorytypeIdent, __gong__toRawStringLiteral(codeholdingcategorytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeholdingcategorytypeIdent, __gong__toRawStringLiteral(codeholdingcategorytype.EnclosedText)))
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
		for _, codeiatatype := range __gong__sortStageSetInstances(stage.CodeIATATypes, stage.CodeIATAType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeiatatypeIdent := codeiatatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeIATAType{Name: %s}).Stage(stage)", codeiatatypeIdent, __gong__toRawStringLiteral(codeiatatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeiatatypeIdent, __gong__toRawStringLiteral(codeiatatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeiatatypeIdent, __gong__toRawStringLiteral(codeiatatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeiatatypeIdent, __gong__toRawStringLiteral(codeiatatype.EnclosedText)))
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
		for _, codeicaotype := range __gong__sortStageSetInstances(stage.CodeICAOTypes, stage.CodeICAOType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeicaotypeIdent := codeicaotype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeICAOType{Name: %s}).Stage(stage)", codeicaotypeIdent, __gong__toRawStringLiteral(codeicaotype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeicaotypeIdent, __gong__toRawStringLiteral(codeicaotype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeicaotypeIdent, __gong__toRawStringLiteral(codeicaotype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeicaotypeIdent, __gong__toRawStringLiteral(codeicaotype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeilsbackcoursetype := range __gong__sortStageSetInstances(stage.CodeILSBackCourseTypes, stage.CodeILSBackCourseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeilsbackcoursetypeIdent := codeilsbackcoursetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeILSBackCourseType{Name: %s}).Stage(stage)", codeilsbackcoursetypeIdent, __gong__toRawStringLiteral(codeilsbackcoursetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeilsbackcoursetypeIdent, __gong__toRawStringLiteral(codeilsbackcoursetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeilsbackcoursetypeIdent, __gong__toRawStringLiteral(codeilsbackcoursetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeilsbackcoursetypeIdent, __gong__toRawStringLiteral(codeilsbackcoursetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeintegritylevelilstype := range __gong__sortStageSetInstances(stage.CodeIntegrityLevelILSTypes, stage.CodeIntegrityLevelILSType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeintegritylevelilstypeIdent := codeintegritylevelilstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeIntegrityLevelILSType{Name: %s}).Stage(stage)", codeintegritylevelilstypeIdent, __gong__toRawStringLiteral(codeintegritylevelilstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeintegritylevelilstypeIdent, __gong__toRawStringLiteral(codeintegritylevelilstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeintegritylevelilstypeIdent, __gong__toRawStringLiteral(codeintegritylevelilstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeintegritylevelilstypeIdent, __gong__toRawStringLiteral(codeintegritylevelilstype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeintensitystandbytype := range __gong__sortStageSetInstances(stage.CodeIntensityStandByTypes, stage.CodeIntensityStandByType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeintensitystandbytypeIdent := codeintensitystandbytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeIntensityStandByType{Name: %s}).Stage(stage)", codeintensitystandbytypeIdent, __gong__toRawStringLiteral(codeintensitystandbytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeintensitystandbytypeIdent, __gong__toRawStringLiteral(codeintensitystandbytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeintensitystandbytypeIdent, __gong__toRawStringLiteral(codeintensitystandbytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeintensitystandbytypeIdent, __gong__toRawStringLiteral(codeintensitystandbytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codelanguagetype := range __gong__sortStageSetInstances(stage.CodeLanguageTypes, stage.CodeLanguageType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codelanguagetypeIdent := codelanguagetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLanguageType{Name: %s}).Stage(stage)", codelanguagetypeIdent, __gong__toRawStringLiteral(codelanguagetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelanguagetypeIdent, __gong__toRawStringLiteral(codelanguagetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelanguagetypeIdent, __gong__toRawStringLiteral(codelanguagetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelanguagetypeIdent, __gong__toRawStringLiteral(codelanguagetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codelevelseriestype := range __gong__sortStageSetInstances(stage.CodeLevelSeriesTypes, stage.CodeLevelSeriesType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codelevelseriestypeIdent := codelevelseriestype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLevelSeriesType{Name: %s}).Stage(stage)", codelevelseriestypeIdent, __gong__toRawStringLiteral(codelevelseriestype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelevelseriestypeIdent, __gong__toRawStringLiteral(codelevelseriestype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelevelseriestypeIdent, __gong__toRawStringLiteral(codelevelseriestype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelevelseriestypeIdent, __gong__toRawStringLiteral(codelevelseriestype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeleveltabledesignatortype := range __gong__sortStageSetInstances(stage.CodeLevelTableDesignatorTypes, stage.CodeLevelTableDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeleveltabledesignatortypeIdent := codeleveltabledesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLevelTableDesignatorType{Name: %s}).Stage(stage)", codeleveltabledesignatortypeIdent, __gong__toRawStringLiteral(codeleveltabledesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeleveltabledesignatortypeIdent, __gong__toRawStringLiteral(codeleveltabledesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeleveltabledesignatortypeIdent, __gong__toRawStringLiteral(codeleveltabledesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeleveltabledesignatortypeIdent, __gong__toRawStringLiteral(codeleveltabledesignatortype.EnclosedText)))
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
		for _, codelightholdingpositiontype := range __gong__sortStageSetInstances(stage.CodeLightHoldingPositionTypes, stage.CodeLightHoldingPositionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codelightholdingpositiontypeIdent := codelightholdingpositiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLightHoldingPositionType{Name: %s}).Stage(stage)", codelightholdingpositiontypeIdent, __gong__toRawStringLiteral(codelightholdingpositiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelightholdingpositiontypeIdent, __gong__toRawStringLiteral(codelightholdingpositiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelightholdingpositiontypeIdent, __gong__toRawStringLiteral(codelightholdingpositiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelightholdingpositiontypeIdent, __gong__toRawStringLiteral(codelightholdingpositiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codelightintensitytype := range __gong__sortStageSetInstances(stage.CodeLightIntensityTypes, stage.CodeLightIntensityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codelightintensitytypeIdent := codelightintensitytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLightIntensityType{Name: %s}).Stage(stage)", codelightintensitytypeIdent, __gong__toRawStringLiteral(codelightintensitytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelightintensitytypeIdent, __gong__toRawStringLiteral(codelightintensitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelightintensitytypeIdent, __gong__toRawStringLiteral(codelightintensitytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelightintensitytypeIdent, __gong__toRawStringLiteral(codelightintensitytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codelightsourcetype := range __gong__sortStageSetInstances(stage.CodeLightSourceTypes, stage.CodeLightSourceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codelightsourcetypeIdent := codelightsourcetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLightSourceType{Name: %s}).Stage(stage)", codelightsourcetypeIdent, __gong__toRawStringLiteral(codelightsourcetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelightsourcetypeIdent, __gong__toRawStringLiteral(codelightsourcetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelightsourcetypeIdent, __gong__toRawStringLiteral(codelightsourcetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelightsourcetypeIdent, __gong__toRawStringLiteral(codelightsourcetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codelightingjartype := range __gong__sortStageSetInstances(stage.CodeLightingJARTypes, stage.CodeLightingJARType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codelightingjartypeIdent := codelightingjartype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLightingJARType{Name: %s}).Stage(stage)", codelightingjartypeIdent, __gong__toRawStringLiteral(codelightingjartype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelightingjartypeIdent, __gong__toRawStringLiteral(codelightingjartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelightingjartypeIdent, __gong__toRawStringLiteral(codelightingjartype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelightingjartypeIdent, __gong__toRawStringLiteral(codelightingjartype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeloadingbridgetype := range __gong__sortStageSetInstances(stage.CodeLoadingBridgeTypes, stage.CodeLoadingBridgeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeloadingbridgetypeIdent := codeloadingbridgetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLoadingBridgeType{Name: %s}).Stage(stage)", codeloadingbridgetypeIdent, __gong__toRawStringLiteral(codeloadingbridgetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeloadingbridgetypeIdent, __gong__toRawStringLiteral(codeloadingbridgetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeloadingbridgetypeIdent, __gong__toRawStringLiteral(codeloadingbridgetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeloadingbridgetypeIdent, __gong__toRawStringLiteral(codeloadingbridgetype.EnclosedText)))
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
		for _, codelogicaloperatortype := range __gong__sortStageSetInstances(stage.CodeLogicalOperatorTypes, stage.CodeLogicalOperatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codelogicaloperatortypeIdent := codelogicaloperatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLogicalOperatorType{Name: %s}).Stage(stage)", codelogicaloperatortypeIdent, __gong__toRawStringLiteral(codelogicaloperatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelogicaloperatortypeIdent, __gong__toRawStringLiteral(codelogicaloperatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelogicaloperatortypeIdent, __gong__toRawStringLiteral(codelogicaloperatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelogicaloperatortypeIdent, __gong__toRawStringLiteral(codelogicaloperatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codemlsazimuthtype := range __gong__sortStageSetInstances(stage.CodeMLSAzimuthTypes, stage.CodeMLSAzimuthType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemlsazimuthtypeIdent := codemlsazimuthtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMLSAzimuthType{Name: %s}).Stage(stage)", codemlsazimuthtypeIdent, __gong__toRawStringLiteral(codemlsazimuthtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemlsazimuthtypeIdent, __gong__toRawStringLiteral(codemlsazimuthtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemlsazimuthtypeIdent, __gong__toRawStringLiteral(codemlsazimuthtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemlsazimuthtypeIdent, __gong__toRawStringLiteral(codemlsazimuthtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codemlschanneltype := range __gong__sortStageSetInstances(stage.CodeMLSChannelTypes, stage.CodeMLSChannelType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemlschanneltypeIdent := codemlschanneltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMLSChannelType{Name: %s}).Stage(stage)", codemlschanneltypeIdent, __gong__toRawStringLiteral(codemlschanneltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemlschanneltypeIdent, __gong__toRawStringLiteral(codemlschanneltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemlschanneltypeIdent, __gong__toRawStringLiteral(codemlschanneltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemlschanneltypeIdent, __gong__toRawStringLiteral(codemlschanneltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codemarkerbeaconsignaltype := range __gong__sortStageSetInstances(stage.CodeMarkerBeaconSignalTypes, stage.CodeMarkerBeaconSignalType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemarkerbeaconsignaltypeIdent := codemarkerbeaconsignaltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMarkerBeaconSignalType{Name: %s}).Stage(stage)", codemarkerbeaconsignaltypeIdent, __gong__toRawStringLiteral(codemarkerbeaconsignaltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemarkerbeaconsignaltypeIdent, __gong__toRawStringLiteral(codemarkerbeaconsignaltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemarkerbeaconsignaltypeIdent, __gong__toRawStringLiteral(codemarkerbeaconsignaltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemarkerbeaconsignaltypeIdent, __gong__toRawStringLiteral(codemarkerbeaconsignaltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codemarkingconditiontype := range __gong__sortStageSetInstances(stage.CodeMarkingConditionTypes, stage.CodeMarkingConditionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemarkingconditiontypeIdent := codemarkingconditiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMarkingConditionType{Name: %s}).Stage(stage)", codemarkingconditiontypeIdent, __gong__toRawStringLiteral(codemarkingconditiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemarkingconditiontypeIdent, __gong__toRawStringLiteral(codemarkingconditiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemarkingconditiontypeIdent, __gong__toRawStringLiteral(codemarkingconditiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemarkingconditiontypeIdent, __gong__toRawStringLiteral(codemarkingconditiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codemarkingstyletype := range __gong__sortStageSetInstances(stage.CodeMarkingStyleTypes, stage.CodeMarkingStyleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemarkingstyletypeIdent := codemarkingstyletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMarkingStyleType{Name: %s}).Stage(stage)", codemarkingstyletypeIdent, __gong__toRawStringLiteral(codemarkingstyletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemarkingstyletypeIdent, __gong__toRawStringLiteral(codemarkingstyletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemarkingstyletypeIdent, __gong__toRawStringLiteral(codemarkingstyletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemarkingstyletypeIdent, __gong__toRawStringLiteral(codemarkingstyletype.EnclosedText)))
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
		for _, codemilitaryoperationstype := range __gong__sortStageSetInstances(stage.CodeMilitaryOperationsTypes, stage.CodeMilitaryOperationsType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codemilitaryoperationstypeIdent := codemilitaryoperationstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMilitaryOperationsType{Name: %s}).Stage(stage)", codemilitaryoperationstypeIdent, __gong__toRawStringLiteral(codemilitaryoperationstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemilitaryoperationstypeIdent, __gong__toRawStringLiteral(codemilitaryoperationstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemilitaryoperationstypeIdent, __gong__toRawStringLiteral(codemilitaryoperationstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemilitaryoperationstypeIdent, __gong__toRawStringLiteral(codemilitaryoperationstype.EnclosedText)))
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
		for _, codendbusagetype := range __gong__sortStageSetInstances(stage.CodeNDBUsageTypes, stage.CodeNDBUsageType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codendbusagetypeIdent := codendbusagetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNDBUsageType{Name: %s}).Stage(stage)", codendbusagetypeIdent, __gong__toRawStringLiteral(codendbusagetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codendbusagetypeIdent, __gong__toRawStringLiteral(codendbusagetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codendbusagetypeIdent, __gong__toRawStringLiteral(codendbusagetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codendbusagetypeIdent, __gong__toRawStringLiteral(codendbusagetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codenavaiddesignatortype := range __gong__sortStageSetInstances(stage.CodeNavaidDesignatorTypes, stage.CodeNavaidDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codenavaiddesignatortypeIdent := codenavaiddesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavaidDesignatorType{Name: %s}).Stage(stage)", codenavaiddesignatortypeIdent, __gong__toRawStringLiteral(codenavaiddesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavaiddesignatortypeIdent, __gong__toRawStringLiteral(codenavaiddesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavaiddesignatortypeIdent, __gong__toRawStringLiteral(codenavaiddesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavaiddesignatortypeIdent, __gong__toRawStringLiteral(codenavaiddesignatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codenavaidpurposetype := range __gong__sortStageSetInstances(stage.CodeNavaidPurposeTypes, stage.CodeNavaidPurposeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codenavaidpurposetypeIdent := codenavaidpurposetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavaidPurposeType{Name: %s}).Stage(stage)", codenavaidpurposetypeIdent, __gong__toRawStringLiteral(codenavaidpurposetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavaidpurposetypeIdent, __gong__toRawStringLiteral(codenavaidpurposetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavaidpurposetypeIdent, __gong__toRawStringLiteral(codenavaidpurposetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavaidpurposetypeIdent, __gong__toRawStringLiteral(codenavaidpurposetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codenavaidservicetype := range __gong__sortStageSetInstances(stage.CodeNavaidServiceTypes, stage.CodeNavaidServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codenavaidservicetypeIdent := codenavaidservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavaidServiceType{Name: %s}).Stage(stage)", codenavaidservicetypeIdent, __gong__toRawStringLiteral(codenavaidservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavaidservicetypeIdent, __gong__toRawStringLiteral(codenavaidservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavaidservicetypeIdent, __gong__toRawStringLiteral(codenavaidservicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavaidservicetypeIdent, __gong__toRawStringLiteral(codenavaidservicetype.EnclosedText)))
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
		for _, codenitrogentype := range __gong__sortStageSetInstances(stage.CodeNitrogenTypes, stage.CodeNitrogenType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codenitrogentypeIdent := codenitrogentype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNitrogenType{Name: %s}).Stage(stage)", codenitrogentypeIdent, __gong__toRawStringLiteral(codenitrogentype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenitrogentypeIdent, __gong__toRawStringLiteral(codenitrogentype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenitrogentypeIdent, __gong__toRawStringLiteral(codenitrogentype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenitrogentypeIdent, __gong__toRawStringLiteral(codenitrogentype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codenorthreferencetype := range __gong__sortStageSetInstances(stage.CodeNorthReferenceTypes, stage.CodeNorthReferenceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codenorthreferencetypeIdent := codenorthreferencetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNorthReferenceType{Name: %s}).Stage(stage)", codenorthreferencetypeIdent, __gong__toRawStringLiteral(codenorthreferencetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenorthreferencetypeIdent, __gong__toRawStringLiteral(codenorthreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenorthreferencetypeIdent, __gong__toRawStringLiteral(codenorthreferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenorthreferencetypeIdent, __gong__toRawStringLiteral(codenorthreferencetype.EnclosedText)))
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
		for _, codeobstacleareatype := range __gong__sortStageSetInstances(stage.CodeObstacleAreaTypes, stage.CodeObstacleAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeobstacleareatypeIdent := codeobstacleareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeObstacleAreaType{Name: %s}).Stage(stage)", codeobstacleareatypeIdent, __gong__toRawStringLiteral(codeobstacleareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeobstacleareatypeIdent, __gong__toRawStringLiteral(codeobstacleareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeobstacleareatypeIdent, __gong__toRawStringLiteral(codeobstacleareatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeobstacleareatypeIdent, __gong__toRawStringLiteral(codeobstacleareatype.EnclosedText)))
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
		for _, codeoiltype := range __gong__sortStageSetInstances(stage.CodeOilTypes, stage.CodeOilType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeoiltypeIdent := codeoiltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOilType{Name: %s}).Stage(stage)", codeoiltypeIdent, __gong__toRawStringLiteral(codeoiltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeoiltypeIdent, __gong__toRawStringLiteral(codeoiltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeoiltypeIdent, __gong__toRawStringLiteral(codeoiltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeoiltypeIdent, __gong__toRawStringLiteral(codeoiltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeoperationairportheliporttype := range __gong__sortStageSetInstances(stage.CodeOperationAirportHeliportTypes, stage.CodeOperationAirportHeliportType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeoperationairportheliporttypeIdent := codeoperationairportheliporttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOperationAirportHeliportType{Name: %s}).Stage(stage)", codeoperationairportheliporttypeIdent, __gong__toRawStringLiteral(codeoperationairportheliporttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeoperationairportheliporttypeIdent, __gong__toRawStringLiteral(codeoperationairportheliporttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeoperationairportheliporttypeIdent, __gong__toRawStringLiteral(codeoperationairportheliporttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeoperationairportheliporttypeIdent, __gong__toRawStringLiteral(codeoperationairportheliporttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeoperationmanoeuvringareatype := range __gong__sortStageSetInstances(stage.CodeOperationManoeuvringAreaTypes, stage.CodeOperationManoeuvringAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeoperationmanoeuvringareatypeIdent := codeoperationmanoeuvringareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOperationManoeuvringAreaType{Name: %s}).Stage(stage)", codeoperationmanoeuvringareatypeIdent, __gong__toRawStringLiteral(codeoperationmanoeuvringareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeoperationmanoeuvringareatypeIdent, __gong__toRawStringLiteral(codeoperationmanoeuvringareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeoperationmanoeuvringareatypeIdent, __gong__toRawStringLiteral(codeoperationmanoeuvringareatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeoperationmanoeuvringareatypeIdent, __gong__toRawStringLiteral(codeoperationmanoeuvringareatype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeorganisationdesignatortype := range __gong__sortStageSetInstances(stage.CodeOrganisationDesignatorTypes, stage.CodeOrganisationDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeorganisationdesignatortypeIdent := codeorganisationdesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOrganisationDesignatorType{Name: %s}).Stage(stage)", codeorganisationdesignatortypeIdent, __gong__toRawStringLiteral(codeorganisationdesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeorganisationdesignatortypeIdent, __gong__toRawStringLiteral(codeorganisationdesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeorganisationdesignatortypeIdent, __gong__toRawStringLiteral(codeorganisationdesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeorganisationdesignatortypeIdent, __gong__toRawStringLiteral(codeorganisationdesignatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeorganisationhierarchytype := range __gong__sortStageSetInstances(stage.CodeOrganisationHierarchyTypes, stage.CodeOrganisationHierarchyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeorganisationhierarchytypeIdent := codeorganisationhierarchytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOrganisationHierarchyType{Name: %s}).Stage(stage)", codeorganisationhierarchytypeIdent, __gong__toRawStringLiteral(codeorganisationhierarchytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeorganisationhierarchytypeIdent, __gong__toRawStringLiteral(codeorganisationhierarchytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeorganisationhierarchytypeIdent, __gong__toRawStringLiteral(codeorganisationhierarchytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeorganisationhierarchytypeIdent, __gong__toRawStringLiteral(codeorganisationhierarchytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeorganisationtype := range __gong__sortStageSetInstances(stage.CodeOrganisationTypes, stage.CodeOrganisationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeorganisationtypeIdent := codeorganisationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOrganisationType{Name: %s}).Stage(stage)", codeorganisationtypeIdent, __gong__toRawStringLiteral(codeorganisationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeorganisationtypeIdent, __gong__toRawStringLiteral(codeorganisationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeorganisationtypeIdent, __gong__toRawStringLiteral(codeorganisationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeorganisationtypeIdent, __gong__toRawStringLiteral(codeorganisationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeoxygentype := range __gong__sortStageSetInstances(stage.CodeOxygenTypes, stage.CodeOxygenType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeoxygentypeIdent := codeoxygentype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOxygenType{Name: %s}).Stage(stage)", codeoxygentypeIdent, __gong__toRawStringLiteral(codeoxygentype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeoxygentypeIdent, __gong__toRawStringLiteral(codeoxygentype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeoxygentypeIdent, __gong__toRawStringLiteral(codeoxygentype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeoxygentypeIdent, __gong__toRawStringLiteral(codeoxygentype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codepartype := range __gong__sortStageSetInstances(stage.CodePARTypes, stage.CodePARType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codepartypeIdent := codepartype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePARType{Name: %s}).Stage(stage)", codepartypeIdent, __gong__toRawStringLiteral(codepartype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepartypeIdent, __gong__toRawStringLiteral(codepartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepartypeIdent, __gong__toRawStringLiteral(codepartype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepartypeIdent, __gong__toRawStringLiteral(codepartype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codepcnmethodtype := range __gong__sortStageSetInstances(stage.CodePCNMethodTypes, stage.CodePCNMethodType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codepcnmethodtypeIdent := codepcnmethodtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePCNMethodType{Name: %s}).Stage(stage)", codepcnmethodtypeIdent, __gong__toRawStringLiteral(codepcnmethodtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepcnmethodtypeIdent, __gong__toRawStringLiteral(codepcnmethodtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepcnmethodtypeIdent, __gong__toRawStringLiteral(codepcnmethodtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepcnmethodtypeIdent, __gong__toRawStringLiteral(codepcnmethodtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codepcnpavementtype := range __gong__sortStageSetInstances(stage.CodePCNPavementTypes, stage.CodePCNPavementType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codepcnpavementtypeIdent := codepcnpavementtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePCNPavementType{Name: %s}).Stage(stage)", codepcnpavementtypeIdent, __gong__toRawStringLiteral(codepcnpavementtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepcnpavementtypeIdent, __gong__toRawStringLiteral(codepcnpavementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepcnpavementtypeIdent, __gong__toRawStringLiteral(codepcnpavementtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepcnpavementtypeIdent, __gong__toRawStringLiteral(codepcnpavementtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codepcnsubgradetype := range __gong__sortStageSetInstances(stage.CodePCNSubgradeTypes, stage.CodePCNSubgradeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codepcnsubgradetypeIdent := codepcnsubgradetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePCNSubgradeType{Name: %s}).Stage(stage)", codepcnsubgradetypeIdent, __gong__toRawStringLiteral(codepcnsubgradetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepcnsubgradetypeIdent, __gong__toRawStringLiteral(codepcnsubgradetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepcnsubgradetypeIdent, __gong__toRawStringLiteral(codepcnsubgradetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepcnsubgradetypeIdent, __gong__toRawStringLiteral(codepcnsubgradetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codepcntyrepressuretype := range __gong__sortStageSetInstances(stage.CodePCNTyrePressureTypes, stage.CodePCNTyrePressureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codepcntyrepressuretypeIdent := codepcntyrepressuretype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePCNTyrePressureType{Name: %s}).Stage(stage)", codepcntyrepressuretypeIdent, __gong__toRawStringLiteral(codepcntyrepressuretype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepcntyrepressuretypeIdent, __gong__toRawStringLiteral(codepcntyrepressuretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepcntyrepressuretypeIdent, __gong__toRawStringLiteral(codepcntyrepressuretype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepcntyrepressuretypeIdent, __gong__toRawStringLiteral(codepcntyrepressuretype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codepassengerservicetype := range __gong__sortStageSetInstances(stage.CodePassengerServiceTypes, stage.CodePassengerServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codepassengerservicetypeIdent := codepassengerservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePassengerServiceType{Name: %s}).Stage(stage)", codepassengerservicetypeIdent, __gong__toRawStringLiteral(codepassengerservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepassengerservicetypeIdent, __gong__toRawStringLiteral(codepassengerservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepassengerservicetypeIdent, __gong__toRawStringLiteral(codepassengerservicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepassengerservicetypeIdent, __gong__toRawStringLiteral(codepassengerservicetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codepilotcontrolledlightingtype := range __gong__sortStageSetInstances(stage.CodePilotControlledLightingTypes, stage.CodePilotControlledLightingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codepilotcontrolledlightingtypeIdent := codepilotcontrolledlightingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePilotControlledLightingType{Name: %s}).Stage(stage)", codepilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(codepilotcontrolledlightingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(codepilotcontrolledlightingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(codepilotcontrolledlightingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(codepilotcontrolledlightingtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codepositioninilstype := range __gong__sortStageSetInstances(stage.CodePositionInILSTypes, stage.CodePositionInILSType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codepositioninilstypeIdent := codepositioninilstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePositionInILSType{Name: %s}).Stage(stage)", codepositioninilstypeIdent, __gong__toRawStringLiteral(codepositioninilstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepositioninilstypeIdent, __gong__toRawStringLiteral(codepositioninilstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepositioninilstypeIdent, __gong__toRawStringLiteral(codepositioninilstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepositioninilstypeIdent, __gong__toRawStringLiteral(codepositioninilstype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeprimaryradartype := range __gong__sortStageSetInstances(stage.CodePrimaryRadarTypes, stage.CodePrimaryRadarType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeprimaryradartypeIdent := codeprimaryradartype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePrimaryRadarType{Name: %s}).Stage(stage)", codeprimaryradartypeIdent, __gong__toRawStringLiteral(codeprimaryradartype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeprimaryradartypeIdent, __gong__toRawStringLiteral(codeprimaryradartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeprimaryradartypeIdent, __gong__toRawStringLiteral(codeprimaryradartype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeprimaryradartypeIdent, __gong__toRawStringLiteral(codeprimaryradartype.EnclosedText)))
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
		for _, codeprotectareasectiontype := range __gong__sortStageSetInstances(stage.CodeProtectAreaSectionTypes, stage.CodeProtectAreaSectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeprotectareasectiontypeIdent := codeprotectareasectiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeProtectAreaSectionType{Name: %s}).Stage(stage)", codeprotectareasectiontypeIdent, __gong__toRawStringLiteral(codeprotectareasectiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeprotectareasectiontypeIdent, __gong__toRawStringLiteral(codeprotectareasectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeprotectareasectiontypeIdent, __gong__toRawStringLiteral(codeprotectareasectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeprotectareasectiontypeIdent, __gong__toRawStringLiteral(codeprotectareasectiontype.EnclosedText)))
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
		for _, codervrreadingtype := range __gong__sortStageSetInstances(stage.CodeRVRReadingTypes, stage.CodeRVRReadingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codervrreadingtypeIdent := codervrreadingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRVRReadingType{Name: %s}).Stage(stage)", codervrreadingtypeIdent, __gong__toRawStringLiteral(codervrreadingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codervrreadingtypeIdent, __gong__toRawStringLiteral(codervrreadingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codervrreadingtypeIdent, __gong__toRawStringLiteral(codervrreadingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codervrreadingtypeIdent, __gong__toRawStringLiteral(codervrreadingtype.EnclosedText)))
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
		for _, coderadarservicetype := range __gong__sortStageSetInstances(stage.CodeRadarServiceTypes, stage.CodeRadarServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderadarservicetypeIdent := coderadarservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRadarServiceType{Name: %s}).Stage(stage)", coderadarservicetypeIdent, __gong__toRawStringLiteral(coderadarservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderadarservicetypeIdent, __gong__toRawStringLiteral(coderadarservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderadarservicetypeIdent, __gong__toRawStringLiteral(coderadarservicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderadarservicetypeIdent, __gong__toRawStringLiteral(coderadarservicetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderadioemissiontype := range __gong__sortStageSetInstances(stage.CodeRadioEmissionTypes, stage.CodeRadioEmissionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderadioemissiontypeIdent := coderadioemissiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRadioEmissionType{Name: %s}).Stage(stage)", coderadioemissiontypeIdent, __gong__toRawStringLiteral(coderadioemissiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderadioemissiontypeIdent, __gong__toRawStringLiteral(coderadioemissiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderadioemissiontypeIdent, __gong__toRawStringLiteral(coderadioemissiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderadioemissiontypeIdent, __gong__toRawStringLiteral(coderadioemissiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderadiofrequencyareatype := range __gong__sortStageSetInstances(stage.CodeRadioFrequencyAreaTypes, stage.CodeRadioFrequencyAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderadiofrequencyareatypeIdent := coderadiofrequencyareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRadioFrequencyAreaType{Name: %s}).Stage(stage)", coderadiofrequencyareatypeIdent, __gong__toRawStringLiteral(coderadiofrequencyareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderadiofrequencyareatypeIdent, __gong__toRawStringLiteral(coderadiofrequencyareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderadiofrequencyareatypeIdent, __gong__toRawStringLiteral(coderadiofrequencyareatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderadiofrequencyareatypeIdent, __gong__toRawStringLiteral(coderadiofrequencyareatype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderadiosignaltype := range __gong__sortStageSetInstances(stage.CodeRadioSignalTypes, stage.CodeRadioSignalType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderadiosignaltypeIdent := coderadiosignaltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRadioSignalType{Name: %s}).Stage(stage)", coderadiosignaltypeIdent, __gong__toRawStringLiteral(coderadiosignaltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderadiosignaltypeIdent, __gong__toRawStringLiteral(coderadiosignaltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderadiosignaltypeIdent, __gong__toRawStringLiteral(coderadiosignaltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderadiosignaltypeIdent, __gong__toRawStringLiteral(coderadiosignaltype.EnclosedText)))
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
		for _, codereflectortype := range __gong__sortStageSetInstances(stage.CodeReflectorTypes, stage.CodeReflectorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codereflectortypeIdent := codereflectortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeReflectorType{Name: %s}).Stage(stage)", codereflectortypeIdent, __gong__toRawStringLiteral(codereflectortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codereflectortypeIdent, __gong__toRawStringLiteral(codereflectortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codereflectortypeIdent, __gong__toRawStringLiteral(codereflectortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codereflectortypeIdent, __gong__toRawStringLiteral(codereflectortype.EnclosedText)))
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
		for _, coderoadtype := range __gong__sortStageSetInstances(stage.CodeRoadTypes, stage.CodeRoadType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderoadtypeIdent := coderoadtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRoadType{Name: %s}).Stage(stage)", coderoadtypeIdent, __gong__toRawStringLiteral(coderoadtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoadtypeIdent, __gong__toRawStringLiteral(coderoadtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoadtypeIdent, __gong__toRawStringLiteral(coderoadtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoadtypeIdent, __gong__toRawStringLiteral(coderoadtype.EnclosedText)))
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
		for _, coderunwayelementtype := range __gong__sortStageSetInstances(stage.CodeRunwayElementTypes, stage.CodeRunwayElementType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderunwayelementtypeIdent := coderunwayelementtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRunwayElementType{Name: %s}).Stage(stage)", coderunwayelementtypeIdent, __gong__toRawStringLiteral(coderunwayelementtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderunwayelementtypeIdent, __gong__toRawStringLiteral(coderunwayelementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderunwayelementtypeIdent, __gong__toRawStringLiteral(coderunwayelementtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderunwayelementtypeIdent, __gong__toRawStringLiteral(coderunwayelementtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderunwaymarkingtype := range __gong__sortStageSetInstances(stage.CodeRunwayMarkingTypes, stage.CodeRunwayMarkingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderunwaymarkingtypeIdent := coderunwaymarkingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRunwayMarkingType{Name: %s}).Stage(stage)", coderunwaymarkingtypeIdent, __gong__toRawStringLiteral(coderunwaymarkingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderunwaymarkingtypeIdent, __gong__toRawStringLiteral(coderunwaymarkingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderunwaymarkingtypeIdent, __gong__toRawStringLiteral(coderunwaymarkingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderunwaymarkingtypeIdent, __gong__toRawStringLiteral(coderunwaymarkingtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderunwaypointroletype := range __gong__sortStageSetInstances(stage.CodeRunwayPointRoleTypes, stage.CodeRunwayPointRoleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderunwaypointroletypeIdent := coderunwaypointroletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRunwayPointRoleType{Name: %s}).Stage(stage)", coderunwaypointroletypeIdent, __gong__toRawStringLiteral(coderunwaypointroletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderunwaypointroletypeIdent, __gong__toRawStringLiteral(coderunwaypointroletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderunwaypointroletypeIdent, __gong__toRawStringLiteral(coderunwaypointroletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderunwaypointroletypeIdent, __gong__toRawStringLiteral(coderunwaypointroletype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderunwayprotectionareatype := range __gong__sortStageSetInstances(stage.CodeRunwayProtectionAreaTypes, stage.CodeRunwayProtectionAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderunwayprotectionareatypeIdent := coderunwayprotectionareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRunwayProtectionAreaType{Name: %s}).Stage(stage)", coderunwayprotectionareatypeIdent, __gong__toRawStringLiteral(coderunwayprotectionareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderunwayprotectionareatypeIdent, __gong__toRawStringLiteral(coderunwayprotectionareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderunwayprotectionareatypeIdent, __gong__toRawStringLiteral(coderunwayprotectionareatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderunwayprotectionareatypeIdent, __gong__toRawStringLiteral(coderunwayprotectionareatype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderunwaysectiontype := range __gong__sortStageSetInstances(stage.CodeRunwaySectionTypes, stage.CodeRunwaySectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderunwaysectiontypeIdent := coderunwaysectiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRunwaySectionType{Name: %s}).Stage(stage)", coderunwaysectiontypeIdent, __gong__toRawStringLiteral(coderunwaysectiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderunwaysectiontypeIdent, __gong__toRawStringLiteral(coderunwaysectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderunwaysectiontypeIdent, __gong__toRawStringLiteral(coderunwaysectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderunwaysectiontypeIdent, __gong__toRawStringLiteral(coderunwaysectiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, coderunwaytype := range __gong__sortStageSetInstances(stage.CodeRunwayTypes, stage.CodeRunwayType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coderunwaytypeIdent := coderunwaytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRunwayType{Name: %s}).Stage(stage)", coderunwaytypeIdent, __gong__toRawStringLiteral(coderunwaytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderunwaytypeIdent, __gong__toRawStringLiteral(coderunwaytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderunwaytypeIdent, __gong__toRawStringLiteral(coderunwaytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderunwaytypeIdent, __gong__toRawStringLiteral(coderunwaytype.EnclosedText)))
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
		for _, codeserviceatctype := range __gong__sortStageSetInstances(stage.CodeServiceATCTypes, stage.CodeServiceATCType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeserviceatctypeIdent := codeserviceatctype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeServiceATCType{Name: %s}).Stage(stage)", codeserviceatctypeIdent, __gong__toRawStringLiteral(codeserviceatctype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeserviceatctypeIdent, __gong__toRawStringLiteral(codeserviceatctype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeserviceatctypeIdent, __gong__toRawStringLiteral(codeserviceatctype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeserviceatctypeIdent, __gong__toRawStringLiteral(codeserviceatctype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeserviceatfmtype := range __gong__sortStageSetInstances(stage.CodeServiceATFMTypes, stage.CodeServiceATFMType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeserviceatfmtypeIdent := codeserviceatfmtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeServiceATFMType{Name: %s}).Stage(stage)", codeserviceatfmtypeIdent, __gong__toRawStringLiteral(codeserviceatfmtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeserviceatfmtypeIdent, __gong__toRawStringLiteral(codeserviceatfmtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeserviceatfmtypeIdent, __gong__toRawStringLiteral(codeserviceatfmtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeserviceatfmtypeIdent, __gong__toRawStringLiteral(codeserviceatfmtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeservicegroundcontroltype := range __gong__sortStageSetInstances(stage.CodeServiceGroundControlTypes, stage.CodeServiceGroundControlType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeservicegroundcontroltypeIdent := codeservicegroundcontroltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeServiceGroundControlType{Name: %s}).Stage(stage)", codeservicegroundcontroltypeIdent, __gong__toRawStringLiteral(codeservicegroundcontroltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeservicegroundcontroltypeIdent, __gong__toRawStringLiteral(codeservicegroundcontroltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeservicegroundcontroltypeIdent, __gong__toRawStringLiteral(codeservicegroundcontroltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeservicegroundcontroltypeIdent, __gong__toRawStringLiteral(codeservicegroundcontroltype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeserviceinformationtype := range __gong__sortStageSetInstances(stage.CodeServiceInformationTypes, stage.CodeServiceInformationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeserviceinformationtypeIdent := codeserviceinformationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeServiceInformationType{Name: %s}).Stage(stage)", codeserviceinformationtypeIdent, __gong__toRawStringLiteral(codeserviceinformationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeserviceinformationtypeIdent, __gong__toRawStringLiteral(codeserviceinformationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeserviceinformationtypeIdent, __gong__toRawStringLiteral(codeserviceinformationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeserviceinformationtypeIdent, __gong__toRawStringLiteral(codeserviceinformationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeservicesartype := range __gong__sortStageSetInstances(stage.CodeServiceSARTypes, stage.CodeServiceSARType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeservicesartypeIdent := codeservicesartype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeServiceSARType{Name: %s}).Stage(stage)", codeservicesartypeIdent, __gong__toRawStringLiteral(codeservicesartype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeservicesartypeIdent, __gong__toRawStringLiteral(codeservicesartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeservicesartypeIdent, __gong__toRawStringLiteral(codeservicesartype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeservicesartypeIdent, __gong__toRawStringLiteral(codeservicesartype.EnclosedText)))
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
		for _, codesignalperformanceilstype := range __gong__sortStageSetInstances(stage.CodeSignalPerformanceILSTypes, stage.CodeSignalPerformanceILSType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codesignalperformanceilstypeIdent := codesignalperformanceilstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSignalPerformanceILSType{Name: %s}).Stage(stage)", codesignalperformanceilstypeIdent, __gong__toRawStringLiteral(codesignalperformanceilstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesignalperformanceilstypeIdent, __gong__toRawStringLiteral(codesignalperformanceilstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesignalperformanceilstypeIdent, __gong__toRawStringLiteral(codesignalperformanceilstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesignalperformanceilstypeIdent, __gong__toRawStringLiteral(codesignalperformanceilstype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codespecialdatetype := range __gong__sortStageSetInstances(stage.CodeSpecialDateTypes, stage.CodeSpecialDateType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codespecialdatetypeIdent := codespecialdatetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSpecialDateType{Name: %s}).Stage(stage)", codespecialdatetypeIdent, __gong__toRawStringLiteral(codespecialdatetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codespecialdatetypeIdent, __gong__toRawStringLiteral(codespecialdatetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codespecialdatetypeIdent, __gong__toRawStringLiteral(codespecialdatetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codespecialdatetypeIdent, __gong__toRawStringLiteral(codespecialdatetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codespecialnavigationchaindesignatortype := range __gong__sortStageSetInstances(stage.CodeSpecialNavigationChainDesignatorTypes, stage.CodeSpecialNavigationChainDesignatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codespecialnavigationchaindesignatortypeIdent := codespecialnavigationchaindesignatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSpecialNavigationChainDesignatorType{Name: %s}).Stage(stage)", codespecialnavigationchaindesignatortypeIdent, __gong__toRawStringLiteral(codespecialnavigationchaindesignatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codespecialnavigationchaindesignatortypeIdent, __gong__toRawStringLiteral(codespecialnavigationchaindesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codespecialnavigationchaindesignatortypeIdent, __gong__toRawStringLiteral(codespecialnavigationchaindesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codespecialnavigationchaindesignatortypeIdent, __gong__toRawStringLiteral(codespecialnavigationchaindesignatortype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codespecialnavigationstationtype := range __gong__sortStageSetInstances(stage.CodeSpecialNavigationStationTypes, stage.CodeSpecialNavigationStationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codespecialnavigationstationtypeIdent := codespecialnavigationstationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSpecialNavigationStationType{Name: %s}).Stage(stage)", codespecialnavigationstationtypeIdent, __gong__toRawStringLiteral(codespecialnavigationstationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codespecialnavigationstationtypeIdent, __gong__toRawStringLiteral(codespecialnavigationstationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codespecialnavigationstationtypeIdent, __gong__toRawStringLiteral(codespecialnavigationstationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codespecialnavigationstationtypeIdent, __gong__toRawStringLiteral(codespecialnavigationstationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codespecialnavigationsystemtype := range __gong__sortStageSetInstances(stage.CodeSpecialNavigationSystemTypes, stage.CodeSpecialNavigationSystemType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codespecialnavigationsystemtypeIdent := codespecialnavigationsystemtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSpecialNavigationSystemType{Name: %s}).Stage(stage)", codespecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(codespecialnavigationsystemtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codespecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(codespecialnavigationsystemtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codespecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(codespecialnavigationsystemtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codespecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(codespecialnavigationsystemtype.EnclosedText)))
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
		for _, codestandbypowertype := range __gong__sortStageSetInstances(stage.CodeStandbyPowerTypes, stage.CodeStandbyPowerType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codestandbypowertypeIdent := codestandbypowertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStandbyPowerType{Name: %s}).Stage(stage)", codestandbypowertypeIdent, __gong__toRawStringLiteral(codestandbypowertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestandbypowertypeIdent, __gong__toRawStringLiteral(codestandbypowertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestandbypowertypeIdent, __gong__toRawStringLiteral(codestandbypowertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestandbypowertypeIdent, __gong__toRawStringLiteral(codestandbypowertype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codestatusairporttype := range __gong__sortStageSetInstances(stage.CodeStatusAirportTypes, stage.CodeStatusAirportType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codestatusairporttypeIdent := codestatusairporttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusAirportType{Name: %s}).Stage(stage)", codestatusairporttypeIdent, __gong__toRawStringLiteral(codestatusairporttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusairporttypeIdent, __gong__toRawStringLiteral(codestatusairporttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusairporttypeIdent, __gong__toRawStringLiteral(codestatusairporttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusairporttypeIdent, __gong__toRawStringLiteral(codestatusairporttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codestatusairspacetype := range __gong__sortStageSetInstances(stage.CodeStatusAirspaceTypes, stage.CodeStatusAirspaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codestatusairspacetypeIdent := codestatusairspacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusAirspaceType{Name: %s}).Stage(stage)", codestatusairspacetypeIdent, __gong__toRawStringLiteral(codestatusairspacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusairspacetypeIdent, __gong__toRawStringLiteral(codestatusairspacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusairspacetypeIdent, __gong__toRawStringLiteral(codestatusairspacetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusairspacetypeIdent, __gong__toRawStringLiteral(codestatusairspacetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codestatusconstructiontype := range __gong__sortStageSetInstances(stage.CodeStatusConstructionTypes, stage.CodeStatusConstructionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codestatusconstructiontypeIdent := codestatusconstructiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusConstructionType{Name: %s}).Stage(stage)", codestatusconstructiontypeIdent, __gong__toRawStringLiteral(codestatusconstructiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusconstructiontypeIdent, __gong__toRawStringLiteral(codestatusconstructiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusconstructiontypeIdent, __gong__toRawStringLiteral(codestatusconstructiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusconstructiontypeIdent, __gong__toRawStringLiteral(codestatusconstructiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codestatusnavaidtype := range __gong__sortStageSetInstances(stage.CodeStatusNavaidTypes, stage.CodeStatusNavaidType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codestatusnavaidtypeIdent := codestatusnavaidtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusNavaidType{Name: %s}).Stage(stage)", codestatusnavaidtypeIdent, __gong__toRawStringLiteral(codestatusnavaidtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusnavaidtypeIdent, __gong__toRawStringLiteral(codestatusnavaidtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusnavaidtypeIdent, __gong__toRawStringLiteral(codestatusnavaidtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusnavaidtypeIdent, __gong__toRawStringLiteral(codestatusnavaidtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codestatusoperationstype := range __gong__sortStageSetInstances(stage.CodeStatusOperationsTypes, stage.CodeStatusOperationsType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codestatusoperationstypeIdent := codestatusoperationstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusOperationsType{Name: %s}).Stage(stage)", codestatusoperationstypeIdent, __gong__toRawStringLiteral(codestatusoperationstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusoperationstypeIdent, __gong__toRawStringLiteral(codestatusoperationstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusoperationstypeIdent, __gong__toRawStringLiteral(codestatusoperationstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusoperationstypeIdent, __gong__toRawStringLiteral(codestatusoperationstype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codestatusservicetype := range __gong__sortStageSetInstances(stage.CodeStatusServiceTypes, stage.CodeStatusServiceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codestatusservicetypeIdent := codestatusservicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusServiceType{Name: %s}).Stage(stage)", codestatusservicetypeIdent, __gong__toRawStringLiteral(codestatusservicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusservicetypeIdent, __gong__toRawStringLiteral(codestatusservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusservicetypeIdent, __gong__toRawStringLiteral(codestatusservicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusservicetypeIdent, __gong__toRawStringLiteral(codestatusservicetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codesurfacecompositiontype := range __gong__sortStageSetInstances(stage.CodeSurfaceCompositionTypes, stage.CodeSurfaceCompositionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codesurfacecompositiontypeIdent := codesurfacecompositiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSurfaceCompositionType{Name: %s}).Stage(stage)", codesurfacecompositiontypeIdent, __gong__toRawStringLiteral(codesurfacecompositiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesurfacecompositiontypeIdent, __gong__toRawStringLiteral(codesurfacecompositiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesurfacecompositiontypeIdent, __gong__toRawStringLiteral(codesurfacecompositiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesurfacecompositiontypeIdent, __gong__toRawStringLiteral(codesurfacecompositiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codesurfaceconditiontype := range __gong__sortStageSetInstances(stage.CodeSurfaceConditionTypes, stage.CodeSurfaceConditionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codesurfaceconditiontypeIdent := codesurfaceconditiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSurfaceConditionType{Name: %s}).Stage(stage)", codesurfaceconditiontypeIdent, __gong__toRawStringLiteral(codesurfaceconditiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesurfaceconditiontypeIdent, __gong__toRawStringLiteral(codesurfaceconditiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesurfaceconditiontypeIdent, __gong__toRawStringLiteral(codesurfaceconditiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesurfaceconditiontypeIdent, __gong__toRawStringLiteral(codesurfaceconditiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codesurfacepreparationtype := range __gong__sortStageSetInstances(stage.CodeSurfacePreparationTypes, stage.CodeSurfacePreparationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codesurfacepreparationtypeIdent := codesurfacepreparationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSurfacePreparationType{Name: %s}).Stage(stage)", codesurfacepreparationtypeIdent, __gong__toRawStringLiteral(codesurfacepreparationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesurfacepreparationtypeIdent, __gong__toRawStringLiteral(codesurfacepreparationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesurfacepreparationtypeIdent, __gong__toRawStringLiteral(codesurfacepreparationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesurfacepreparationtypeIdent, __gong__toRawStringLiteral(codesurfacepreparationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codesystemactivationtype := range __gong__sortStageSetInstances(stage.CodeSystemActivationTypes, stage.CodeSystemActivationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codesystemactivationtypeIdent := codesystemactivationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSystemActivationType{Name: %s}).Stage(stage)", codesystemactivationtypeIdent, __gong__toRawStringLiteral(codesystemactivationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesystemactivationtypeIdent, __gong__toRawStringLiteral(codesystemactivationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesystemactivationtypeIdent, __gong__toRawStringLiteral(codesystemactivationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesystemactivationtypeIdent, __gong__toRawStringLiteral(codesystemactivationtype.EnclosedText)))
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
		for _, codetlofsectiontype := range __gong__sortStageSetInstances(stage.CodeTLOFSectionTypes, stage.CodeTLOFSectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetlofsectiontypeIdent := codetlofsectiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTLOFSectionType{Name: %s}).Stage(stage)", codetlofsectiontypeIdent, __gong__toRawStringLiteral(codetlofsectiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetlofsectiontypeIdent, __gong__toRawStringLiteral(codetlofsectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetlofsectiontypeIdent, __gong__toRawStringLiteral(codetlofsectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetlofsectiontypeIdent, __gong__toRawStringLiteral(codetlofsectiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codetaxiwayelementtype := range __gong__sortStageSetInstances(stage.CodeTaxiwayElementTypes, stage.CodeTaxiwayElementType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetaxiwayelementtypeIdent := codetaxiwayelementtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTaxiwayElementType{Name: %s}).Stage(stage)", codetaxiwayelementtypeIdent, __gong__toRawStringLiteral(codetaxiwayelementtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetaxiwayelementtypeIdent, __gong__toRawStringLiteral(codetaxiwayelementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetaxiwayelementtypeIdent, __gong__toRawStringLiteral(codetaxiwayelementtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetaxiwayelementtypeIdent, __gong__toRawStringLiteral(codetaxiwayelementtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codetaxiwaysectiontype := range __gong__sortStageSetInstances(stage.CodeTaxiwaySectionTypes, stage.CodeTaxiwaySectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetaxiwaysectiontypeIdent := codetaxiwaysectiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTaxiwaySectionType{Name: %s}).Stage(stage)", codetaxiwaysectiontypeIdent, __gong__toRawStringLiteral(codetaxiwaysectiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetaxiwaysectiontypeIdent, __gong__toRawStringLiteral(codetaxiwaysectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetaxiwaysectiontypeIdent, __gong__toRawStringLiteral(codetaxiwaysectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetaxiwaysectiontypeIdent, __gong__toRawStringLiteral(codetaxiwaysectiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codetaxiwaytype := range __gong__sortStageSetInstances(stage.CodeTaxiwayTypes, stage.CodeTaxiwayType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetaxiwaytypeIdent := codetaxiwaytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTaxiwayType{Name: %s}).Stage(stage)", codetaxiwaytypeIdent, __gong__toRawStringLiteral(codetaxiwaytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetaxiwaytypeIdent, __gong__toRawStringLiteral(codetaxiwaytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetaxiwaytypeIdent, __gong__toRawStringLiteral(codetaxiwaytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetaxiwaytypeIdent, __gong__toRawStringLiteral(codetaxiwaytype.EnclosedText)))
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
		for _, codeuuidtype := range __gong__sortStageSetInstances(stage.CodeUUIDTypes, stage.CodeUUIDType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeuuidtypeIdent := codeuuidtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeUUIDType{Name: %s}).Stage(stage)", codeuuidtypeIdent, __gong__toRawStringLiteral(codeuuidtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeuuidtypeIdent, __gong__toRawStringLiteral(codeuuidtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeuuidtypeIdent, __gong__toRawStringLiteral(codeuuidtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeuuidtypeIdent, __gong__toRawStringLiteral(codeuuidtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeunitdependencytype := range __gong__sortStageSetInstances(stage.CodeUnitDependencyTypes, stage.CodeUnitDependencyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeunitdependencytypeIdent := codeunitdependencytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeUnitDependencyType{Name: %s}).Stage(stage)", codeunitdependencytypeIdent, __gong__toRawStringLiteral(codeunitdependencytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeunitdependencytypeIdent, __gong__toRawStringLiteral(codeunitdependencytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeunitdependencytypeIdent, __gong__toRawStringLiteral(codeunitdependencytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeunitdependencytypeIdent, __gong__toRawStringLiteral(codeunitdependencytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeunittype := range __gong__sortStageSetInstances(stage.CodeUnitTypes, stage.CodeUnitType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeunittypeIdent := codeunittype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeUnitType{Name: %s}).Stage(stage)", codeunittypeIdent, __gong__toRawStringLiteral(codeunittype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeunittypeIdent, __gong__toRawStringLiteral(codeunittype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeunittypeIdent, __gong__toRawStringLiteral(codeunittype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeunittypeIdent, __gong__toRawStringLiteral(codeunittype.EnclosedText)))
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
		for _, codeusagelimitationtype := range __gong__sortStageSetInstances(stage.CodeUsageLimitationTypes, stage.CodeUsageLimitationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeusagelimitationtypeIdent := codeusagelimitationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeUsageLimitationType{Name: %s}).Stage(stage)", codeusagelimitationtypeIdent, __gong__toRawStringLiteral(codeusagelimitationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeusagelimitationtypeIdent, __gong__toRawStringLiteral(codeusagelimitationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeusagelimitationtypeIdent, __gong__toRawStringLiteral(codeusagelimitationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeusagelimitationtypeIdent, __gong__toRawStringLiteral(codeusagelimitationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codevasistype := range __gong__sortStageSetInstances(stage.CodeVASISTypes, stage.CodeVASISType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codevasistypeIdent := codevasistype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVASISType{Name: %s}).Stage(stage)", codevasistypeIdent, __gong__toRawStringLiteral(codevasistype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codevasistypeIdent, __gong__toRawStringLiteral(codevasistype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codevasistypeIdent, __gong__toRawStringLiteral(codevasistype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codevasistypeIdent, __gong__toRawStringLiteral(codevasistype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codevortype := range __gong__sortStageSetInstances(stage.CodeVORTypes, stage.CodeVORType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codevortypeIdent := codevortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVORType{Name: %s}).Stage(stage)", codevortypeIdent, __gong__toRawStringLiteral(codevortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codevortypeIdent, __gong__toRawStringLiteral(codevortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codevortypeIdent, __gong__toRawStringLiteral(codevortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codevortypeIdent, __gong__toRawStringLiteral(codevortype.EnclosedText)))
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
		for _, codeverticalstructuremarkingtype := range __gong__sortStageSetInstances(stage.CodeVerticalStructureMarkingTypes, stage.CodeVerticalStructureMarkingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeverticalstructuremarkingtypeIdent := codeverticalstructuremarkingtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVerticalStructureMarkingType{Name: %s}).Stage(stage)", codeverticalstructuremarkingtypeIdent, __gong__toRawStringLiteral(codeverticalstructuremarkingtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeverticalstructuremarkingtypeIdent, __gong__toRawStringLiteral(codeverticalstructuremarkingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeverticalstructuremarkingtypeIdent, __gong__toRawStringLiteral(codeverticalstructuremarkingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeverticalstructuremarkingtypeIdent, __gong__toRawStringLiteral(codeverticalstructuremarkingtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeverticalstructurematerialtype := range __gong__sortStageSetInstances(stage.CodeVerticalStructureMaterialTypes, stage.CodeVerticalStructureMaterialType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeverticalstructurematerialtypeIdent := codeverticalstructurematerialtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVerticalStructureMaterialType{Name: %s}).Stage(stage)", codeverticalstructurematerialtypeIdent, __gong__toRawStringLiteral(codeverticalstructurematerialtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeverticalstructurematerialtypeIdent, __gong__toRawStringLiteral(codeverticalstructurematerialtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeverticalstructurematerialtypeIdent, __gong__toRawStringLiteral(codeverticalstructurematerialtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeverticalstructurematerialtypeIdent, __gong__toRawStringLiteral(codeverticalstructurematerialtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeverticalstructuretype := range __gong__sortStageSetInstances(stage.CodeVerticalStructureTypes, stage.CodeVerticalStructureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeverticalstructuretypeIdent := codeverticalstructuretype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVerticalStructureType{Name: %s}).Stage(stage)", codeverticalstructuretypeIdent, __gong__toRawStringLiteral(codeverticalstructuretype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeverticalstructuretypeIdent, __gong__toRawStringLiteral(codeverticalstructuretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeverticalstructuretypeIdent, __gong__toRawStringLiteral(codeverticalstructuretype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeverticalstructuretypeIdent, __gong__toRawStringLiteral(codeverticalstructuretype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codevisualdockingguidancetype := range __gong__sortStageSetInstances(stage.CodeVisualDockingGuidanceTypes, stage.CodeVisualDockingGuidanceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codevisualdockingguidancetypeIdent := codevisualdockingguidancetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVisualDockingGuidanceType{Name: %s}).Stage(stage)", codevisualdockingguidancetypeIdent, __gong__toRawStringLiteral(codevisualdockingguidancetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codevisualdockingguidancetypeIdent, __gong__toRawStringLiteral(codevisualdockingguidancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codevisualdockingguidancetypeIdent, __gong__toRawStringLiteral(codevisualdockingguidancetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codevisualdockingguidancetypeIdent, __gong__toRawStringLiteral(codevisualdockingguidancetype.EnclosedText)))
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
		for _, codeworkareatype := range __gong__sortStageSetInstances(stage.CodeWorkAreaTypes, stage.CodeWorkAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeworkareatypeIdent := codeworkareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeWorkAreaType{Name: %s}).Stage(stage)", codeworkareatypeIdent, __gong__toRawStringLiteral(codeworkareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeworkareatypeIdent, __gong__toRawStringLiteral(codeworkareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeworkareatypeIdent, __gong__toRawStringLiteral(codeworkareatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeworkareatypeIdent, __gong__toRawStringLiteral(codeworkareatype.EnclosedText)))
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
		for _, datetimetype := range __gong__sortStageSetInstances(stage.DateTimeTypes, stage.DateTimeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datetimetypeIdent := datetimetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.DateTimeType{Name: %s}).Stage(stage)", datetimetypeIdent, __gong__toRawStringLiteral(datetimetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datetimetypeIdent, __gong__toRawStringLiteral(datetimetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", datetimetypeIdent, __gong__toRawStringLiteral(datetimetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", datetimetypeIdent, __gong__toRawStringLiteral(datetimetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, datetype := range __gong__sortStageSetInstances(stage.DateTypes, stage.DateType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datetypeIdent := datetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.DateType{Name: %s}).Stage(stage)", datetypeIdent, __gong__toRawStringLiteral(datetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datetypeIdent, __gong__toRawStringLiteral(datetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", datetypeIdent, __gong__toRawStringLiteral(datetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", datetypeIdent, __gong__toRawStringLiteral(datetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, dateyeartype := range __gong__sortStageSetInstances(stage.DateYearTypes, stage.DateYearType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dateyeartypeIdent := dateyeartype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.DateYearType{Name: %s}).Stage(stage)", dateyeartypeIdent, __gong__toRawStringLiteral(dateyeartype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dateyeartypeIdent, __gong__toRawStringLiteral(dateyeartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", dateyeartypeIdent, __gong__toRawStringLiteral(dateyeartype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", dateyeartypeIdent, __gong__toRawStringLiteral(dateyeartype.EnclosedText)))
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
		for _, textremarktype := range __gong__sortStageSetInstances(stage.TextRemarkTypes, stage.TextRemarkType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textremarktypeIdent := textremarktype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextRemarkType{Name: %s}).Stage(stage)", textremarktypeIdent, __gong__toRawStringLiteral(textremarktype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textremarktypeIdent, __gong__toRawStringLiteral(textremarktype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textremarktypeIdent, __gong__toRawStringLiteral(textremarktype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textremarktypeIdent, __gong__toRawStringLiteral(textremarktype.EnclosedText)))
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
		for _, valdepthtype := range __gong__sortStageSetInstances(stage.ValDepthTypes, stage.ValDepthType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valdepthtypeIdent := valdepthtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValDepthType{Name: %s}).Stage(stage)", valdepthtypeIdent, __gong__toRawStringLiteral(valdepthtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valdepthtypeIdent, __gong__toRawStringLiteral(valdepthtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valdepthtypeIdent, __gong__toRawStringLiteral(valdepthtype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valdepthtypeIdent, __gong__toRawStringLiteral(valdepthtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valdepthtypeIdent, __gong__toRawStringLiteral(valdepthtype.EnclosedText)))
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
		for _, valfltype := range __gong__sortStageSetInstances(stage.ValFLTypes, stage.ValFLType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valfltypeIdent := valfltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValFLType{Name: %s}).Stage(stage)", valfltypeIdent, __gong__toRawStringLiteral(valfltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valfltypeIdent, __gong__toRawStringLiteral(valfltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valfltypeIdent, __gong__toRawStringLiteral(valfltype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valfltypeIdent, __gong__toRawStringLiteral(valfltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", valfltypeIdent, valfltype.EnclosedText))
		}
	}
	if stage != nil {
		for _, valfrequencytype := range __gong__sortStageSetInstances(stage.ValFrequencyTypes, stage.ValFrequencyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valfrequencytypeIdent := valfrequencytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValFrequencyType{Name: %s}).Stage(stage)", valfrequencytypeIdent, __gong__toRawStringLiteral(valfrequencytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valfrequencytypeIdent, __gong__toRawStringLiteral(valfrequencytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valfrequencytypeIdent, __gong__toRawStringLiteral(valfrequencytype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valfrequencytypeIdent, __gong__toRawStringLiteral(valfrequencytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valfrequencytypeIdent, __gong__toRawStringLiteral(valfrequencytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valfrictiontype := range __gong__sortStageSetInstances(stage.ValFrictionTypes, stage.ValFrictionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valfrictiontypeIdent := valfrictiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValFrictionType{Name: %s}).Stage(stage)", valfrictiontypeIdent, __gong__toRawStringLiteral(valfrictiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valfrictiontypeIdent, __gong__toRawStringLiteral(valfrictiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valfrictiontypeIdent, __gong__toRawStringLiteral(valfrictiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valfrictiontypeIdent, __gong__toRawStringLiteral(valfrictiontype.EnclosedText)))
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
		for _, vallcntype := range __gong__sortStageSetInstances(stage.ValLCNTypes, stage.ValLCNType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vallcntypeIdent := vallcntype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValLCNType{Name: %s}).Stage(stage)", vallcntypeIdent, __gong__toRawStringLiteral(vallcntype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vallcntypeIdent, __gong__toRawStringLiteral(vallcntype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", vallcntypeIdent, __gong__toRawStringLiteral(vallcntype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", vallcntypeIdent, __gong__toRawStringLiteral(vallcntype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, vallightintensitytype := range __gong__sortStageSetInstances(stage.ValLightIntensityTypes, stage.ValLightIntensityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vallightintensitytypeIdent := vallightintensitytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValLightIntensityType{Name: %s}).Stage(stage)", vallightintensitytypeIdent, __gong__toRawStringLiteral(vallightintensitytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vallightintensitytypeIdent, __gong__toRawStringLiteral(vallightintensitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", vallightintensitytypeIdent, __gong__toRawStringLiteral(vallightintensitytype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", vallightintensitytypeIdent, __gong__toRawStringLiteral(vallightintensitytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", vallightintensitytypeIdent, __gong__toRawStringLiteral(vallightintensitytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valmagneticvariationchangetype := range __gong__sortStageSetInstances(stage.ValMagneticVariationChangeTypes, stage.ValMagneticVariationChangeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valmagneticvariationchangetypeIdent := valmagneticvariationchangetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValMagneticVariationChangeType{Name: %s}).Stage(stage)", valmagneticvariationchangetypeIdent, __gong__toRawStringLiteral(valmagneticvariationchangetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valmagneticvariationchangetypeIdent, __gong__toRawStringLiteral(valmagneticvariationchangetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valmagneticvariationchangetypeIdent, __gong__toRawStringLiteral(valmagneticvariationchangetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valmagneticvariationchangetypeIdent, __gong__toRawStringLiteral(valmagneticvariationchangetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valmagneticvariationtype := range __gong__sortStageSetInstances(stage.ValMagneticVariationTypes, stage.ValMagneticVariationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valmagneticvariationtypeIdent := valmagneticvariationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValMagneticVariationType{Name: %s}).Stage(stage)", valmagneticvariationtypeIdent, __gong__toRawStringLiteral(valmagneticvariationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valmagneticvariationtypeIdent, __gong__toRawStringLiteral(valmagneticvariationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valmagneticvariationtypeIdent, __gong__toRawStringLiteral(valmagneticvariationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valmagneticvariationtypeIdent, __gong__toRawStringLiteral(valmagneticvariationtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valpcntype := range __gong__sortStageSetInstances(stage.ValPCNTypes, stage.ValPCNType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valpcntypeIdent := valpcntype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValPCNType{Name: %s}).Stage(stage)", valpcntypeIdent, __gong__toRawStringLiteral(valpcntype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valpcntypeIdent, __gong__toRawStringLiteral(valpcntype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valpcntypeIdent, __gong__toRawStringLiteral(valpcntype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valpcntypeIdent, __gong__toRawStringLiteral(valpcntype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valpercenttype := range __gong__sortStageSetInstances(stage.ValPercentTypes, stage.ValPercentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valpercenttypeIdent := valpercenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValPercentType{Name: %s}).Stage(stage)", valpercenttypeIdent, __gong__toRawStringLiteral(valpercenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valpercenttypeIdent, __gong__toRawStringLiteral(valpercenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valpercenttypeIdent, __gong__toRawStringLiteral(valpercenttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valpercenttypeIdent, __gong__toRawStringLiteral(valpercenttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, valpressuretype := range __gong__sortStageSetInstances(stage.ValPressureTypes, stage.ValPressureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			valpressuretypeIdent := valpressuretype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValPressureType{Name: %s}).Stage(stage)", valpressuretypeIdent, __gong__toRawStringLiteral(valpressuretype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valpressuretypeIdent, __gong__toRawStringLiteral(valpressuretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valpressuretypeIdent, __gong__toRawStringLiteral(valpressuretype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valpressuretypeIdent, __gong__toRawStringLiteral(valpressuretype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valpressuretypeIdent, __gong__toRawStringLiteral(valpressuretype.EnclosedText)))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmdatatypes/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmdatatypes"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmdatatypes" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmdatatypes\"", fmt.Sprintf("%q", modelsPackageName))
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

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaircraftgroundservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaircraftgroundservicetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaircraftgroundservicetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaircraftgroundservicetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAircraftGroundServiceType", fieldName)
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

func (codeaircraftstandtype *CodeAircraftStandType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaircraftstandtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaircraftstandtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaircraftstandtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaircraftstandtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAircraftStandType", fieldName)
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

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairportheliportcollocationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairportheliportcollocationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairportheliportcollocationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairportheliportcollocationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirportHeliportCollocationType", fieldName)
	}
	return
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairportheliportdesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairportheliportdesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairportheliportdesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairportheliportdesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirportHeliportDesignatorType", fieldName)
	}
	return
}

func (codeairportheliporttype *CodeAirportHeliportType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairportheliporttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairportheliporttype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairportheliporttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairportheliporttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirportHeliportType", fieldName)
	}
	return
}

func (codeairportwarningtype *CodeAirportWarningType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairportwarningtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairportwarningtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairportwarningtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairportwarningtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirportWarningType", fieldName)
	}
	return
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairspaceactivitytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairspaceactivitytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairspaceactivitytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairspaceactivitytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirspaceActivityType", fieldName)
	}
	return
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairspaceaggregationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairspaceaggregationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairspaceaggregationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairspaceaggregationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirspaceAggregationType", fieldName)
	}
	return
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairspaceclassificationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairspaceclassificationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairspaceclassificationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairspaceclassificationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirspaceClassificationType", fieldName)
	}
	return
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairspacedependencytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairspacedependencytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairspacedependencytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairspacedependencytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirspaceDependencyType", fieldName)
	}
	return
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairspacedesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairspacedesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairspacedesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairspacedesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirspaceDesignatorType", fieldName)
	}
	return
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairspacepointpositiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairspacepointpositiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairspacepointpositiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairspacepointpositiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirspacePointPositionType", fieldName)
	}
	return
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairspacepointroletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairspacepointroletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairspacepointroletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairspacepointroletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirspacePointRoleType", fieldName)
	}
	return
}

func (codeairspacetype *CodeAirspaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeairspacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeairspacetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeairspacetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeairspacetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAirspaceType", fieldName)
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

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeapproachlightingicaotype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeapproachlightingicaotype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeapproachlightingicaotype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeapproachlightingicaotype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeApproachLightingICAOType", fieldName)
	}
	return
}

func (codeapproachlightingtype *CodeApproachLightingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeapproachlightingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeapproachlightingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeapproachlightingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeapproachlightingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeApproachLightingType", fieldName)
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

func (codeapronelementtype *CodeApronElementType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeapronelementtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeapronelementtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeapronelementtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeapronelementtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeApronElementType", fieldName)
	}
	return
}

func (codeapronsectiontype *CodeApronSectionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeapronsectiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeapronsectiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeapronsectiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeapronsectiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeApronSectionType", fieldName)
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

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codearrestinggearenergyabsorbtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codearrestinggearenergyabsorbtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codearrestinggearenergyabsorbtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codearrestinggearenergyabsorbtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeArrestingGearEnergyAbsorbType", fieldName)
	}
	return
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codearrestinggearengagedevicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codearrestinggearengagedevicetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codearrestinggearengagedevicetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codearrestinggearengagedevicetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeArrestingGearEngageDeviceType", fieldName)
	}
	return
}

func (codeauralmorsetype *CodeAuralMorseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeauralmorsetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeauralmorsetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeauralmorsetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeauralmorsetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAuralMorseType", fieldName)
	}
	return
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeauthorityroletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeauthorityroletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeauthorityroletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeauthorityroletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAuthorityRoleType", fieldName)
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

func (codeaviationstandardstype *CodeAviationStandardsType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeaviationstandardstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeaviationstandardstype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeaviationstandardstype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeaviationstandardstype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeAviationStandardsType", fieldName)
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

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codebuoydesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codebuoydesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codebuoydesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codebuoydesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeBuoyDesignatorType", fieldName)
	}
	return
}

func (codebuoytype *CodeBuoyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codebuoytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codebuoytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codebuoytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codebuoytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeBuoyType", fieldName)
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

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codecheckpointcategorytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codecheckpointcategorytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codecheckpointcategorytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codecheckpointcategorytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeCheckpointCategoryType", fieldName)
	}
	return
}

func (codecolourtype *CodeColourType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codecolourtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codecolourtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codecolourtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codecolourtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeColourType", fieldName)
	}
	return
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codecommunicationchanneltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codecommunicationchanneltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codecommunicationchanneltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codecommunicationchanneltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeCommunicationChannelType", fieldName)
	}
	return
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codecommunicationdirectiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codecommunicationdirectiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codecommunicationdirectiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codecommunicationdirectiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeCommunicationDirectionType", fieldName)
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

func (codecontaminationtype *CodeContaminationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codecontaminationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codecontaminationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codecontaminationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codecontaminationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeContaminationType", fieldName)
	}
	return
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codecoursequalityilstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codecoursequalityilstype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codecoursequalityilstype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codecoursequalityilstype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeCourseQualityILSType", fieldName)
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

func (codedmechanneltype *CodeDMEChannelType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedmechanneltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedmechanneltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedmechanneltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedmechanneltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDMEChannelType", fieldName)
	}
	return
}

func (codedmetype *CodeDMEType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedmetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedmetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedmetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedmetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDMEType", fieldName)
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

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedeclareddistancetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedeclareddistancetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedeclareddistancetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedeclareddistancetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDeclaredDistanceType", fieldName)
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

func (codedesignatedpointtype *CodeDesignatedPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedesignatedpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedesignatedpointtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedesignatedpointtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedesignatedpointtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDesignatedPointType", fieldName)
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

func (codedistanceindicationtype *CodeDistanceIndicationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codedistanceindicationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codedistanceindicationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codedistanceindicationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codedistanceindicationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeDistanceIndicationType", fieldName)
	}
	return
}

func (codeemissionbandtype *CodeEmissionBandType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeemissionbandtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeemissionbandtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeemissionbandtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeemissionbandtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeEmissionBandType", fieldName)
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

func (codefacilityrankingtype *CodeFacilityRankingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codefacilityrankingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codefacilityrankingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codefacilityrankingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codefacilityrankingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFacilityRankingType", fieldName)
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

func (codefirefightingtype *CodeFireFightingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codefirefightingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codefirefightingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codefirefightingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codefirefightingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFireFightingType", fieldName)
	}
	return
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeflightdestinationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeflightdestinationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeflightdestinationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeflightdestinationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFlightDestinationType", fieldName)
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

func (codefrictiondevicetype *CodeFrictionDeviceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codefrictiondevicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codefrictiondevicetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codefrictiondevicetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codefrictiondevicetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFrictionDeviceType", fieldName)
	}
	return
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codefrictionestimatetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codefrictionestimatetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codefrictionestimatetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codefrictionestimatetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFrictionEstimateType", fieldName)
	}
	return
}

func (codefueltype *CodeFuelType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codefueltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codefueltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codefueltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codefueltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeFuelType", fieldName)
	}
	return
}

func (codegeobordertype *CodeGeoBorderType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codegeobordertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codegeobordertype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codegeobordertype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codegeobordertype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeGeoBorderType", fieldName)
	}
	return
}

func (codegradeseparationtype *CodeGradeSeparationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codegradeseparationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codegradeseparationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codegradeseparationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codegradeseparationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeGradeSeparationType", fieldName)
	}
	return
}

func (codegroundlightingtype *CodeGroundLightingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codegroundlightingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codegroundlightingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codegroundlightingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codegroundlightingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeGroundLightingType", fieldName)
	}
	return
}

func (codeguidancelinetype *CodeGuidanceLineType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeguidancelinetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeguidancelinetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeguidancelinetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeguidancelinetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeGuidanceLineType", fieldName)
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

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codehelicopterperformancetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codehelicopterperformancetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codehelicopterperformancetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codehelicopterperformancetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeHelicopterPerformanceType", fieldName)
	}
	return
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeholdingcategorytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeholdingcategorytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeholdingcategorytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeholdingcategorytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeHoldingCategoryType", fieldName)
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

func (codeiatatype *CodeIATAType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeiatatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeiatatype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeiatatype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeiatatype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeIATAType", fieldName)
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

func (codeicaotype *CodeICAOType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeicaotype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeicaotype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeicaotype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeicaotype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeICAOType", fieldName)
	}
	return
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeilsbackcoursetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeilsbackcoursetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeilsbackcoursetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeilsbackcoursetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeILSBackCourseType", fieldName)
	}
	return
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeintegritylevelilstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeintegritylevelilstype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeintegritylevelilstype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeintegritylevelilstype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeIntegrityLevelILSType", fieldName)
	}
	return
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeintensitystandbytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeintensitystandbytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeintensitystandbytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeintensitystandbytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeIntensityStandByType", fieldName)
	}
	return
}

func (codelanguagetype *CodeLanguageType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codelanguagetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codelanguagetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codelanguagetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codelanguagetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLanguageType", fieldName)
	}
	return
}

func (codelevelseriestype *CodeLevelSeriesType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codelevelseriestype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codelevelseriestype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codelevelseriestype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codelevelseriestype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLevelSeriesType", fieldName)
	}
	return
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeleveltabledesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeleveltabledesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeleveltabledesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeleveltabledesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLevelTableDesignatorType", fieldName)
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

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codelightholdingpositiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codelightholdingpositiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codelightholdingpositiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codelightholdingpositiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLightHoldingPositionType", fieldName)
	}
	return
}

func (codelightintensitytype *CodeLightIntensityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codelightintensitytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codelightintensitytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codelightintensitytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codelightintensitytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLightIntensityType", fieldName)
	}
	return
}

func (codelightsourcetype *CodeLightSourceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codelightsourcetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codelightsourcetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codelightsourcetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codelightsourcetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLightSourceType", fieldName)
	}
	return
}

func (codelightingjartype *CodeLightingJARType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codelightingjartype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codelightingjartype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codelightingjartype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codelightingjartype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLightingJARType", fieldName)
	}
	return
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeloadingbridgetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeloadingbridgetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeloadingbridgetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeloadingbridgetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLoadingBridgeType", fieldName)
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

func (codelogicaloperatortype *CodeLogicalOperatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codelogicaloperatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codelogicaloperatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codelogicaloperatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codelogicaloperatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeLogicalOperatorType", fieldName)
	}
	return
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemlsazimuthtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemlsazimuthtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemlsazimuthtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemlsazimuthtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMLSAzimuthType", fieldName)
	}
	return
}

func (codemlschanneltype *CodeMLSChannelType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemlschanneltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemlschanneltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemlschanneltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemlschanneltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMLSChannelType", fieldName)
	}
	return
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemarkerbeaconsignaltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemarkerbeaconsignaltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemarkerbeaconsignaltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemarkerbeaconsignaltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMarkerBeaconSignalType", fieldName)
	}
	return
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemarkingconditiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemarkingconditiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemarkingconditiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemarkingconditiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMarkingConditionType", fieldName)
	}
	return
}

func (codemarkingstyletype *CodeMarkingStyleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemarkingstyletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemarkingstyletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemarkingstyletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemarkingstyletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMarkingStyleType", fieldName)
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

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codemilitaryoperationstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codemilitaryoperationstype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codemilitaryoperationstype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codemilitaryoperationstype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeMilitaryOperationsType", fieldName)
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

func (codendbusagetype *CodeNDBUsageType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codendbusagetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codendbusagetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codendbusagetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codendbusagetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNDBUsageType", fieldName)
	}
	return
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codenavaiddesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codenavaiddesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codenavaiddesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codenavaiddesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNavaidDesignatorType", fieldName)
	}
	return
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codenavaidpurposetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codenavaidpurposetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codenavaidpurposetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codenavaidpurposetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNavaidPurposeType", fieldName)
	}
	return
}

func (codenavaidservicetype *CodeNavaidServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codenavaidservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codenavaidservicetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codenavaidservicetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codenavaidservicetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNavaidServiceType", fieldName)
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

func (codenitrogentype *CodeNitrogenType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codenitrogentype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codenitrogentype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codenitrogentype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codenitrogentype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNitrogenType", fieldName)
	}
	return
}

func (codenorthreferencetype *CodeNorthReferenceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codenorthreferencetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codenorthreferencetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codenorthreferencetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codenorthreferencetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeNorthReferenceType", fieldName)
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

func (codeobstacleareatype *CodeObstacleAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeobstacleareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeobstacleareatype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeobstacleareatype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeobstacleareatype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeObstacleAreaType", fieldName)
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

func (codeoiltype *CodeOilType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeoiltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeoiltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeoiltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeoiltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeOilType", fieldName)
	}
	return
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeoperationairportheliporttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeoperationairportheliporttype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeoperationairportheliporttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeoperationairportheliporttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeOperationAirportHeliportType", fieldName)
	}
	return
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeoperationmanoeuvringareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeoperationmanoeuvringareatype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeoperationmanoeuvringareatype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeoperationmanoeuvringareatype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeOperationManoeuvringAreaType", fieldName)
	}
	return
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeorganisationdesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeorganisationdesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeorganisationdesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeorganisationdesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeOrganisationDesignatorType", fieldName)
	}
	return
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeorganisationhierarchytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeorganisationhierarchytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeorganisationhierarchytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeorganisationhierarchytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeOrganisationHierarchyType", fieldName)
	}
	return
}

func (codeorganisationtype *CodeOrganisationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeorganisationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeorganisationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeorganisationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeorganisationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeOrganisationType", fieldName)
	}
	return
}

func (codeoxygentype *CodeOxygenType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeoxygentype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeoxygentype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeoxygentype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeoxygentype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeOxygenType", fieldName)
	}
	return
}

func (codepartype *CodePARType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codepartype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codepartype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codepartype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codepartype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodePARType", fieldName)
	}
	return
}

func (codepcnmethodtype *CodePCNMethodType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codepcnmethodtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codepcnmethodtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codepcnmethodtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codepcnmethodtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodePCNMethodType", fieldName)
	}
	return
}

func (codepcnpavementtype *CodePCNPavementType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codepcnpavementtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codepcnpavementtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codepcnpavementtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codepcnpavementtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodePCNPavementType", fieldName)
	}
	return
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codepcnsubgradetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codepcnsubgradetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codepcnsubgradetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codepcnsubgradetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodePCNSubgradeType", fieldName)
	}
	return
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codepcntyrepressuretype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codepcntyrepressuretype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codepcntyrepressuretype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codepcntyrepressuretype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodePCNTyrePressureType", fieldName)
	}
	return
}

func (codepassengerservicetype *CodePassengerServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codepassengerservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codepassengerservicetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codepassengerservicetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codepassengerservicetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodePassengerServiceType", fieldName)
	}
	return
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codepilotcontrolledlightingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codepilotcontrolledlightingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codepilotcontrolledlightingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codepilotcontrolledlightingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodePilotControlledLightingType", fieldName)
	}
	return
}

func (codepositioninilstype *CodePositionInILSType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codepositioninilstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codepositioninilstype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codepositioninilstype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codepositioninilstype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodePositionInILSType", fieldName)
	}
	return
}

func (codeprimaryradartype *CodePrimaryRadarType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeprimaryradartype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeprimaryradartype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeprimaryradartype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeprimaryradartype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodePrimaryRadarType", fieldName)
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

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeprotectareasectiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeprotectareasectiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeprotectareasectiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeprotectareasectiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeProtectAreaSectionType", fieldName)
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

func (codervrreadingtype *CodeRVRReadingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codervrreadingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codervrreadingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codervrreadingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codervrreadingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRVRReadingType", fieldName)
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

func (coderadarservicetype *CodeRadarServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderadarservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderadarservicetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderadarservicetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderadarservicetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRadarServiceType", fieldName)
	}
	return
}

func (coderadioemissiontype *CodeRadioEmissionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderadioemissiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderadioemissiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderadioemissiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderadioemissiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRadioEmissionType", fieldName)
	}
	return
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderadiofrequencyareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderadiofrequencyareatype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderadiofrequencyareatype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderadiofrequencyareatype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRadioFrequencyAreaType", fieldName)
	}
	return
}

func (coderadiosignaltype *CodeRadioSignalType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderadiosignaltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderadiosignaltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderadiosignaltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderadiosignaltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRadioSignalType", fieldName)
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

func (codereflectortype *CodeReflectorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codereflectortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codereflectortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codereflectortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codereflectortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeReflectorType", fieldName)
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

func (coderoadtype *CodeRoadType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderoadtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderoadtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderoadtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderoadtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRoadType", fieldName)
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

func (coderunwayelementtype *CodeRunwayElementType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderunwayelementtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderunwayelementtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderunwayelementtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderunwayelementtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRunwayElementType", fieldName)
	}
	return
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderunwaymarkingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderunwaymarkingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderunwaymarkingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderunwaymarkingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRunwayMarkingType", fieldName)
	}
	return
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderunwaypointroletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderunwaypointroletype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderunwaypointroletype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderunwaypointroletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRunwayPointRoleType", fieldName)
	}
	return
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderunwayprotectionareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderunwayprotectionareatype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderunwayprotectionareatype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderunwayprotectionareatype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRunwayProtectionAreaType", fieldName)
	}
	return
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderunwaysectiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderunwaysectiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderunwaysectiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderunwaysectiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRunwaySectionType", fieldName)
	}
	return
}

func (coderunwaytype *CodeRunwayType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coderunwaytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coderunwaytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", coderunwaytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coderunwaytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeRunwayType", fieldName)
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

func (codeserviceatctype *CodeServiceATCType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeserviceatctype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeserviceatctype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeserviceatctype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeserviceatctype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeServiceATCType", fieldName)
	}
	return
}

func (codeserviceatfmtype *CodeServiceATFMType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeserviceatfmtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeserviceatfmtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeserviceatfmtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeserviceatfmtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeServiceATFMType", fieldName)
	}
	return
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeservicegroundcontroltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeservicegroundcontroltype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeservicegroundcontroltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeservicegroundcontroltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeServiceGroundControlType", fieldName)
	}
	return
}

func (codeserviceinformationtype *CodeServiceInformationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeserviceinformationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeserviceinformationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeserviceinformationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeserviceinformationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeServiceInformationType", fieldName)
	}
	return
}

func (codeservicesartype *CodeServiceSARType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeservicesartype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeservicesartype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeservicesartype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeservicesartype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeServiceSARType", fieldName)
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

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codesignalperformanceilstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codesignalperformanceilstype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codesignalperformanceilstype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codesignalperformanceilstype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSignalPerformanceILSType", fieldName)
	}
	return
}

func (codespecialdatetype *CodeSpecialDateType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codespecialdatetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codespecialdatetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codespecialdatetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codespecialdatetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSpecialDateType", fieldName)
	}
	return
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codespecialnavigationchaindesignatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codespecialnavigationchaindesignatortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codespecialnavigationchaindesignatortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codespecialnavigationchaindesignatortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSpecialNavigationChainDesignatorType", fieldName)
	}
	return
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codespecialnavigationstationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codespecialnavigationstationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codespecialnavigationstationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codespecialnavigationstationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSpecialNavigationStationType", fieldName)
	}
	return
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codespecialnavigationsystemtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codespecialnavigationsystemtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codespecialnavigationsystemtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codespecialnavigationsystemtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSpecialNavigationSystemType", fieldName)
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

func (codestandbypowertype *CodeStandbyPowerType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codestandbypowertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codestandbypowertype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codestandbypowertype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codestandbypowertype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeStandbyPowerType", fieldName)
	}
	return
}

func (codestatusairporttype *CodeStatusAirportType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codestatusairporttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codestatusairporttype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codestatusairporttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codestatusairporttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeStatusAirportType", fieldName)
	}
	return
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codestatusairspacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codestatusairspacetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codestatusairspacetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codestatusairspacetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeStatusAirspaceType", fieldName)
	}
	return
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codestatusconstructiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codestatusconstructiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codestatusconstructiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codestatusconstructiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeStatusConstructionType", fieldName)
	}
	return
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codestatusnavaidtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codestatusnavaidtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codestatusnavaidtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codestatusnavaidtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeStatusNavaidType", fieldName)
	}
	return
}

func (codestatusoperationstype *CodeStatusOperationsType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codestatusoperationstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codestatusoperationstype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codestatusoperationstype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codestatusoperationstype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeStatusOperationsType", fieldName)
	}
	return
}

func (codestatusservicetype *CodeStatusServiceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codestatusservicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codestatusservicetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codestatusservicetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codestatusservicetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeStatusServiceType", fieldName)
	}
	return
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codesurfacecompositiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codesurfacecompositiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codesurfacecompositiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codesurfacecompositiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSurfaceCompositionType", fieldName)
	}
	return
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codesurfaceconditiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codesurfaceconditiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codesurfaceconditiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codesurfaceconditiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSurfaceConditionType", fieldName)
	}
	return
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codesurfacepreparationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codesurfacepreparationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codesurfacepreparationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codesurfacepreparationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSurfacePreparationType", fieldName)
	}
	return
}

func (codesystemactivationtype *CodeSystemActivationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codesystemactivationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codesystemactivationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codesystemactivationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codesystemactivationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeSystemActivationType", fieldName)
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

func (codetlofsectiontype *CodeTLOFSectionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetlofsectiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetlofsectiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetlofsectiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetlofsectiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTLOFSectionType", fieldName)
	}
	return
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetaxiwayelementtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetaxiwayelementtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetaxiwayelementtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetaxiwayelementtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTaxiwayElementType", fieldName)
	}
	return
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetaxiwaysectiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetaxiwaysectiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetaxiwaysectiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetaxiwaysectiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTaxiwaySectionType", fieldName)
	}
	return
}

func (codetaxiwaytype *CodeTaxiwayType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetaxiwaytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetaxiwaytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codetaxiwaytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetaxiwaytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeTaxiwayType", fieldName)
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

func (codeuuidtype *CodeUUIDType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeuuidtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeuuidtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeuuidtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeuuidtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeUUIDType", fieldName)
	}
	return
}

func (codeunitdependencytype *CodeUnitDependencyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeunitdependencytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeunitdependencytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeunitdependencytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeunitdependencytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeUnitDependencyType", fieldName)
	}
	return
}

func (codeunittype *CodeUnitType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeunittype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeunittype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeunittype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeunittype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeUnitType", fieldName)
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

func (codeusagelimitationtype *CodeUsageLimitationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeusagelimitationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeusagelimitationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeusagelimitationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeusagelimitationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeUsageLimitationType", fieldName)
	}
	return
}

func (codevasistype *CodeVASISType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codevasistype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codevasistype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codevasistype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codevasistype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeVASISType", fieldName)
	}
	return
}

func (codevortype *CodeVORType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codevortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codevortype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codevortype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codevortype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeVORType", fieldName)
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

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeverticalstructuremarkingtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeverticalstructuremarkingtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeverticalstructuremarkingtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeverticalstructuremarkingtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeVerticalStructureMarkingType", fieldName)
	}
	return
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeverticalstructurematerialtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeverticalstructurematerialtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeverticalstructurematerialtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeverticalstructurematerialtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeVerticalStructureMaterialType", fieldName)
	}
	return
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeverticalstructuretype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeverticalstructuretype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeverticalstructuretype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeverticalstructuretype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeVerticalStructureType", fieldName)
	}
	return
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codevisualdockingguidancetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codevisualdockingguidancetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codevisualdockingguidancetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codevisualdockingguidancetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeVisualDockingGuidanceType", fieldName)
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

func (codeworkareatype *CodeWorkAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeworkareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeworkareatype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", codeworkareatype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeworkareatype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeWorkAreaType", fieldName)
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

func (datetimetype *DateTimeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datetimetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datetimetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", datetimetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", datetimetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct DateTimeType", fieldName)
	}
	return
}

func (datetype *DateType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", datetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", datetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct DateType", fieldName)
	}
	return
}

func (dateyeartype *DateYearType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := dateyeartype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", dateyeartype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", dateyeartype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", dateyeartype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct DateYearType", fieldName)
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

func (textremarktype *TextRemarkType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := textremarktype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", textremarktype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", textremarktype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", textremarktype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TextRemarkType", fieldName)
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

func (valdepthtype *ValDepthType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valdepthtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valdepthtype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valdepthtype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valdepthtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valdepthtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValDepthType", fieldName)
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

func (valfltype *ValFLType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valfltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valfltype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valfltype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valfltype.NilReason)
	case "EnclosedText":
		res = __gong__marshallInt(ident, "EnclosedText", valfltype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValFLType", fieldName)
	}
	return
}

func (valfrequencytype *ValFrequencyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valfrequencytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valfrequencytype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valfrequencytype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valfrequencytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valfrequencytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValFrequencyType", fieldName)
	}
	return
}

func (valfrictiontype *ValFrictionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valfrictiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valfrictiontype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valfrictiontype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valfrictiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValFrictionType", fieldName)
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

func (vallcntype *ValLCNType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vallcntype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vallcntype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", vallcntype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", vallcntype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValLCNType", fieldName)
	}
	return
}

func (vallightintensitytype *ValLightIntensityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vallightintensitytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vallightintensitytype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", vallightintensitytype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", vallightintensitytype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", vallightintensitytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValLightIntensityType", fieldName)
	}
	return
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valmagneticvariationchangetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valmagneticvariationchangetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valmagneticvariationchangetype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valmagneticvariationchangetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValMagneticVariationChangeType", fieldName)
	}
	return
}

func (valmagneticvariationtype *ValMagneticVariationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valmagneticvariationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valmagneticvariationtype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valmagneticvariationtype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valmagneticvariationtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValMagneticVariationType", fieldName)
	}
	return
}

func (valpcntype *ValPCNType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valpcntype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valpcntype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valpcntype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valpcntype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValPCNType", fieldName)
	}
	return
}

func (valpercenttype *ValPercentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valpercenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valpercenttype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valpercenttype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valpercenttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValPercentType", fieldName)
	}
	return
}

func (valpressuretype *ValPressureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := valpressuretype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", valpressuretype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", valpressuretype.Uom)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", valpressuretype.NilReason)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", valpressuretype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct ValPressureType", fieldName)
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
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaircraftgroundservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaircraftgroundservicetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaircraftgroundservicetype.GongMarshallField(stage, "EnclosedText"))
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
func (codeaircraftstandtype *CodeAircraftStandType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaircraftstandtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaircraftstandtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaircraftstandtype.GongMarshallField(stage, "EnclosedText"))
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
func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairportheliportcollocationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairportheliportcollocationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairportheliportcollocationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairportheliportdesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairportheliportdesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairportheliportdesignatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairportheliporttype *CodeAirportHeliportType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairportheliporttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairportheliporttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairportheliporttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairportwarningtype *CodeAirportWarningType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairportwarningtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairportwarningtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairportwarningtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairspaceactivitytype *CodeAirspaceActivityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairspaceactivitytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairspaceactivitytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairspaceactivitytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairspaceaggregationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairspaceaggregationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairspaceaggregationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairspaceclassificationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairspaceclassificationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairspaceclassificationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairspacedependencytype *CodeAirspaceDependencyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairspacedependencytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairspacedependencytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairspacedependencytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairspacedesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairspacedesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairspacedesignatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairspacepointpositiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairspacepointpositiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairspacepointpositiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairspacepointroletype *CodeAirspacePointRoleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairspacepointroletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairspacepointroletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairspacepointroletype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeairspacetype *CodeAirspaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeairspacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeairspacetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeairspacetype.GongMarshallField(stage, "EnclosedText"))
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
func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeapproachlightingicaotype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeapproachlightingicaotype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeapproachlightingicaotype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeapproachlightingtype *CodeApproachLightingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeapproachlightingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeapproachlightingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeapproachlightingtype.GongMarshallField(stage, "EnclosedText"))
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
func (codeapronelementtype *CodeApronElementType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeapronelementtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeapronelementtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeapronelementtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeapronsectiontype *CodeApronSectionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeapronsectiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeapronsectiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeapronsectiontype.GongMarshallField(stage, "EnclosedText"))
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
func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codearrestinggearenergyabsorbtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codearrestinggearenergyabsorbtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codearrestinggearenergyabsorbtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codearrestinggearengagedevicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codearrestinggearengagedevicetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codearrestinggearengagedevicetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeauralmorsetype *CodeAuralMorseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeauralmorsetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeauralmorsetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeauralmorsetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeauthorityroletype *CodeAuthorityRoleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeauthorityroletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeauthorityroletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeauthorityroletype.GongMarshallField(stage, "EnclosedText"))
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
func (codeaviationstandardstype *CodeAviationStandardsType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeaviationstandardstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeaviationstandardstype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeaviationstandardstype.GongMarshallField(stage, "EnclosedText"))
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
func (codebuoydesignatortype *CodeBuoyDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codebuoydesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codebuoydesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codebuoydesignatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codebuoytype *CodeBuoyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codebuoytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codebuoytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codebuoytype.GongMarshallField(stage, "EnclosedText"))
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
func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codecheckpointcategorytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codecheckpointcategorytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codecheckpointcategorytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codecolourtype *CodeColourType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codecolourtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codecolourtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codecolourtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codecommunicationchanneltype *CodeCommunicationChannelType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codecommunicationchanneltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codecommunicationchanneltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codecommunicationchanneltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codecommunicationdirectiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codecommunicationdirectiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codecommunicationdirectiontype.GongMarshallField(stage, "EnclosedText"))
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
func (codecontaminationtype *CodeContaminationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codecontaminationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codecontaminationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codecontaminationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codecoursequalityilstype *CodeCourseQualityILSType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codecoursequalityilstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codecoursequalityilstype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codecoursequalityilstype.GongMarshallField(stage, "EnclosedText"))
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
func (codedmechanneltype *CodeDMEChannelType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedmechanneltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedmechanneltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedmechanneltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codedmetype *CodeDMEType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedmetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedmetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedmetype.GongMarshallField(stage, "EnclosedText"))
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
func (codedeclareddistancetype *CodeDeclaredDistanceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedeclareddistancetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedeclareddistancetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedeclareddistancetype.GongMarshallField(stage, "EnclosedText"))
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
func (codedesignatedpointtype *CodeDesignatedPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedesignatedpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedesignatedpointtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedesignatedpointtype.GongMarshallField(stage, "EnclosedText"))
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
func (codedistanceindicationtype *CodeDistanceIndicationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codedistanceindicationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codedistanceindicationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codedistanceindicationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeemissionbandtype *CodeEmissionBandType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeemissionbandtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeemissionbandtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeemissionbandtype.GongMarshallField(stage, "EnclosedText"))
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
func (codefacilityrankingtype *CodeFacilityRankingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codefacilityrankingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codefacilityrankingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codefacilityrankingtype.GongMarshallField(stage, "EnclosedText"))
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
func (codefirefightingtype *CodeFireFightingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codefirefightingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codefirefightingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codefirefightingtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeflightdestinationtype *CodeFlightDestinationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeflightdestinationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeflightdestinationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeflightdestinationtype.GongMarshallField(stage, "EnclosedText"))
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
func (codefrictiondevicetype *CodeFrictionDeviceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codefrictiondevicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codefrictiondevicetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codefrictiondevicetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codefrictionestimatetype *CodeFrictionEstimateType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codefrictionestimatetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codefrictionestimatetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codefrictionestimatetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codefueltype *CodeFuelType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codefueltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codefueltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codefueltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codegeobordertype *CodeGeoBorderType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codegeobordertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codegeobordertype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codegeobordertype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codegradeseparationtype *CodeGradeSeparationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codegradeseparationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codegradeseparationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codegradeseparationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codegroundlightingtype *CodeGroundLightingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codegroundlightingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codegroundlightingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codegroundlightingtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeguidancelinetype *CodeGuidanceLineType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeguidancelinetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeguidancelinetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeguidancelinetype.GongMarshallField(stage, "EnclosedText"))
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
func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codehelicopterperformancetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codehelicopterperformancetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codehelicopterperformancetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeholdingcategorytype *CodeHoldingCategoryType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeholdingcategorytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeholdingcategorytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeholdingcategorytype.GongMarshallField(stage, "EnclosedText"))
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
func (codeiatatype *CodeIATAType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeiatatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeiatatype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeiatatype.GongMarshallField(stage, "EnclosedText"))
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
func (codeicaotype *CodeICAOType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeicaotype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeicaotype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeicaotype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeilsbackcoursetype *CodeILSBackCourseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeilsbackcoursetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeilsbackcoursetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeilsbackcoursetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeintegritylevelilstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeintegritylevelilstype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeintegritylevelilstype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeintensitystandbytype *CodeIntensityStandByType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeintensitystandbytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeintensitystandbytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeintensitystandbytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codelanguagetype *CodeLanguageType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codelanguagetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codelanguagetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codelanguagetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codelevelseriestype *CodeLevelSeriesType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codelevelseriestype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codelevelseriestype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codelevelseriestype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeleveltabledesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeleveltabledesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeleveltabledesignatortype.GongMarshallField(stage, "EnclosedText"))
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
func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codelightholdingpositiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codelightholdingpositiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codelightholdingpositiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codelightintensitytype *CodeLightIntensityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codelightintensitytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codelightintensitytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codelightintensitytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codelightsourcetype *CodeLightSourceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codelightsourcetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codelightsourcetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codelightsourcetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codelightingjartype *CodeLightingJARType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codelightingjartype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codelightingjartype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codelightingjartype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeloadingbridgetype *CodeLoadingBridgeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeloadingbridgetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeloadingbridgetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeloadingbridgetype.GongMarshallField(stage, "EnclosedText"))
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
func (codelogicaloperatortype *CodeLogicalOperatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codelogicaloperatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codelogicaloperatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codelogicaloperatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codemlsazimuthtype *CodeMLSAzimuthType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemlsazimuthtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemlsazimuthtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemlsazimuthtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codemlschanneltype *CodeMLSChannelType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemlschanneltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemlschanneltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemlschanneltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemarkerbeaconsignaltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemarkerbeaconsignaltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemarkerbeaconsignaltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codemarkingconditiontype *CodeMarkingConditionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemarkingconditiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemarkingconditiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemarkingconditiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codemarkingstyletype *CodeMarkingStyleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemarkingstyletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemarkingstyletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemarkingstyletype.GongMarshallField(stage, "EnclosedText"))
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
func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codemilitaryoperationstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codemilitaryoperationstype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codemilitaryoperationstype.GongMarshallField(stage, "EnclosedText"))
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
func (codendbusagetype *CodeNDBUsageType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codendbusagetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codendbusagetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codendbusagetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codenavaiddesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codenavaiddesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codenavaiddesignatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codenavaidpurposetype *CodeNavaidPurposeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codenavaidpurposetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codenavaidpurposetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codenavaidpurposetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codenavaidservicetype *CodeNavaidServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codenavaidservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codenavaidservicetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codenavaidservicetype.GongMarshallField(stage, "EnclosedText"))
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
func (codenitrogentype *CodeNitrogenType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codenitrogentype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codenitrogentype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codenitrogentype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codenorthreferencetype *CodeNorthReferenceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codenorthreferencetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codenorthreferencetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codenorthreferencetype.GongMarshallField(stage, "EnclosedText"))
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
func (codeobstacleareatype *CodeObstacleAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeobstacleareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeobstacleareatype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeobstacleareatype.GongMarshallField(stage, "EnclosedText"))
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
func (codeoiltype *CodeOilType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeoiltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeoiltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeoiltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeoperationairportheliporttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeoperationairportheliporttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeoperationairportheliporttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeoperationmanoeuvringareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeoperationmanoeuvringareatype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeoperationmanoeuvringareatype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeorganisationdesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeorganisationdesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeorganisationdesignatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeorganisationhierarchytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeorganisationhierarchytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeorganisationhierarchytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeorganisationtype *CodeOrganisationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeorganisationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeorganisationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeorganisationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeoxygentype *CodeOxygenType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeoxygentype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeoxygentype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeoxygentype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codepartype *CodePARType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codepartype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codepartype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codepartype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codepcnmethodtype *CodePCNMethodType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codepcnmethodtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codepcnmethodtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codepcnmethodtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codepcnpavementtype *CodePCNPavementType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codepcnpavementtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codepcnpavementtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codepcnpavementtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codepcnsubgradetype *CodePCNSubgradeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codepcnsubgradetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codepcnsubgradetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codepcnsubgradetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codepcntyrepressuretype *CodePCNTyrePressureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codepcntyrepressuretype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codepcntyrepressuretype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codepcntyrepressuretype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codepassengerservicetype *CodePassengerServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codepassengerservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codepassengerservicetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codepassengerservicetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codepilotcontrolledlightingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codepilotcontrolledlightingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codepilotcontrolledlightingtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codepositioninilstype *CodePositionInILSType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codepositioninilstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codepositioninilstype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codepositioninilstype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeprimaryradartype *CodePrimaryRadarType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeprimaryradartype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeprimaryradartype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeprimaryradartype.GongMarshallField(stage, "EnclosedText"))
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
func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeprotectareasectiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeprotectareasectiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeprotectareasectiontype.GongMarshallField(stage, "EnclosedText"))
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
func (codervrreadingtype *CodeRVRReadingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codervrreadingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codervrreadingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codervrreadingtype.GongMarshallField(stage, "EnclosedText"))
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
func (coderadarservicetype *CodeRadarServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderadarservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderadarservicetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderadarservicetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderadioemissiontype *CodeRadioEmissionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderadioemissiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderadioemissiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderadioemissiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderadiofrequencyareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderadiofrequencyareatype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderadiofrequencyareatype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderadiosignaltype *CodeRadioSignalType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderadiosignaltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderadiosignaltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderadiosignaltype.GongMarshallField(stage, "EnclosedText"))
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
func (codereflectortype *CodeReflectorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codereflectortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codereflectortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codereflectortype.GongMarshallField(stage, "EnclosedText"))
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
func (coderoadtype *CodeRoadType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderoadtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderoadtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderoadtype.GongMarshallField(stage, "EnclosedText"))
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
func (coderunwayelementtype *CodeRunwayElementType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderunwayelementtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderunwayelementtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderunwayelementtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderunwaymarkingtype *CodeRunwayMarkingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderunwaymarkingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderunwaymarkingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderunwaymarkingtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderunwaypointroletype *CodeRunwayPointRoleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderunwaypointroletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderunwaypointroletype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderunwaypointroletype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderunwayprotectionareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderunwayprotectionareatype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderunwayprotectionareatype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderunwaysectiontype *CodeRunwaySectionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderunwaysectiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderunwaysectiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderunwaysectiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (coderunwaytype *CodeRunwayType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coderunwaytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coderunwaytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(coderunwaytype.GongMarshallField(stage, "EnclosedText"))
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
func (codeserviceatctype *CodeServiceATCType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeserviceatctype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeserviceatctype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeserviceatctype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeserviceatfmtype *CodeServiceATFMType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeserviceatfmtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeserviceatfmtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeserviceatfmtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeservicegroundcontroltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeservicegroundcontroltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeservicegroundcontroltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeserviceinformationtype *CodeServiceInformationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeserviceinformationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeserviceinformationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeserviceinformationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeservicesartype *CodeServiceSARType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeservicesartype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeservicesartype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeservicesartype.GongMarshallField(stage, "EnclosedText"))
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
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codesignalperformanceilstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codesignalperformanceilstype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codesignalperformanceilstype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codespecialdatetype *CodeSpecialDateType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codespecialdatetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codespecialdatetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codespecialdatetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codespecialnavigationchaindesignatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codespecialnavigationchaindesignatortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codespecialnavigationchaindesignatortype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codespecialnavigationstationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codespecialnavigationstationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codespecialnavigationstationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codespecialnavigationsystemtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codespecialnavigationsystemtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codespecialnavigationsystemtype.GongMarshallField(stage, "EnclosedText"))
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
func (codestandbypowertype *CodeStandbyPowerType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codestandbypowertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codestandbypowertype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codestandbypowertype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codestatusairporttype *CodeStatusAirportType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codestatusairporttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codestatusairporttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codestatusairporttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codestatusairspacetype *CodeStatusAirspaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codestatusairspacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codestatusairspacetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codestatusairspacetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codestatusconstructiontype *CodeStatusConstructionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codestatusconstructiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codestatusconstructiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codestatusconstructiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codestatusnavaidtype *CodeStatusNavaidType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codestatusnavaidtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codestatusnavaidtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codestatusnavaidtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codestatusoperationstype *CodeStatusOperationsType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codestatusoperationstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codestatusoperationstype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codestatusoperationstype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codestatusservicetype *CodeStatusServiceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codestatusservicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codestatusservicetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codestatusservicetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codesurfacecompositiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codesurfacecompositiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codesurfacecompositiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codesurfaceconditiontype *CodeSurfaceConditionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codesurfaceconditiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codesurfaceconditiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codesurfaceconditiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codesurfacepreparationtype *CodeSurfacePreparationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codesurfacepreparationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codesurfacepreparationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codesurfacepreparationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codesystemactivationtype *CodeSystemActivationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codesystemactivationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codesystemactivationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codesystemactivationtype.GongMarshallField(stage, "EnclosedText"))
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
func (codetlofsectiontype *CodeTLOFSectionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetlofsectiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetlofsectiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetlofsectiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codetaxiwayelementtype *CodeTaxiwayElementType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetaxiwayelementtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetaxiwayelementtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetaxiwayelementtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetaxiwaysectiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetaxiwaysectiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetaxiwaysectiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codetaxiwaytype *CodeTaxiwayType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetaxiwaytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetaxiwaytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codetaxiwaytype.GongMarshallField(stage, "EnclosedText"))
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
func (codeuuidtype *CodeUUIDType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeuuidtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeuuidtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeuuidtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeunitdependencytype *CodeUnitDependencyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeunitdependencytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeunitdependencytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeunitdependencytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeunittype *CodeUnitType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeunittype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeunittype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeunittype.GongMarshallField(stage, "EnclosedText"))
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
func (codeusagelimitationtype *CodeUsageLimitationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeusagelimitationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeusagelimitationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeusagelimitationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codevasistype *CodeVASISType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codevasistype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codevasistype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codevasistype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codevortype *CodeVORType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codevortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codevortype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codevortype.GongMarshallField(stage, "EnclosedText"))
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
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeverticalstructuremarkingtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeverticalstructuremarkingtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeverticalstructuremarkingtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeverticalstructurematerialtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeverticalstructurematerialtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeverticalstructurematerialtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeverticalstructuretype *CodeVerticalStructureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeverticalstructuretype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeverticalstructuretype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeverticalstructuretype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codevisualdockingguidancetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codevisualdockingguidancetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codevisualdockingguidancetype.GongMarshallField(stage, "EnclosedText"))
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
func (codeworkareatype *CodeWorkAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeworkareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeworkareatype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(codeworkareatype.GongMarshallField(stage, "EnclosedText"))
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
func (datetimetype *DateTimeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datetimetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datetimetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(datetimetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datetype *DateType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(datetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dateyeartype *DateYearType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dateyeartype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(dateyeartype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(dateyeartype.GongMarshallField(stage, "EnclosedText"))
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
func (textremarktype *TextRemarkType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(textremarktype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(textremarktype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(textremarktype.GongMarshallField(stage, "EnclosedText"))
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
func (valdepthtype *ValDepthType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valdepthtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valdepthtype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valdepthtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valdepthtype.GongMarshallField(stage, "EnclosedText"))
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
func (valfltype *ValFLType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valfltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valfltype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valfltype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valfltype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valfrequencytype *ValFrequencyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valfrequencytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valfrequencytype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valfrequencytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valfrequencytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valfrictiontype *ValFrictionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valfrictiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valfrictiontype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valfrictiontype.GongMarshallField(stage, "EnclosedText"))
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
func (vallcntype *ValLCNType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vallcntype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(vallcntype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(vallcntype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (vallightintensitytype *ValLightIntensityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vallightintensitytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(vallightintensitytype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(vallightintensitytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(vallightintensitytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valmagneticvariationchangetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valmagneticvariationchangetype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valmagneticvariationchangetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valmagneticvariationtype *ValMagneticVariationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valmagneticvariationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valmagneticvariationtype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valmagneticvariationtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valpcntype *ValPCNType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valpcntype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valpcntype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valpcntype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valpercenttype *ValPercentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valpercenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valpercenttype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valpercenttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (valpressuretype *ValPressureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(valpressuretype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(valpressuretype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(valpressuretype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(valpressuretype.GongMarshallField(stage, "EnclosedText"))
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
