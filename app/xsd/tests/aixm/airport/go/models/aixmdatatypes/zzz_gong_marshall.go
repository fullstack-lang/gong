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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes"
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
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

// insertion point for marshall all fields methods
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
