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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes"
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
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

// insertion point for marshall all fields methods
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
