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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes"
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
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

// insertion point for marshall all fields methods
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
