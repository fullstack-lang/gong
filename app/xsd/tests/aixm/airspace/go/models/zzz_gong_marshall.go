// generated code - do not edit
package models

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
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *models.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *models.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models"
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

	if stage.XlinkStage != nil {
		for _, xlink_schemaelement := range __gong__sortStageSetInstances(stage.XlinkStage.Xlink_SchemaElements, stage.XlinkStage.Xlink_SchemaElement_stagedOrder) {
			if lastStageDecl != "XlinkStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "XlinkStage"
			}
			xlink_schemaelementIdent := xlink_schemaelement.GongGetIdentifier(stage.XlinkStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.Xlink_SchemaElement{Name: %s}).Stage(stage.XlinkStage)", xlink_schemaelementIdent, __gong__toRawStringLiteral(xlink_schemaelement.Name)))
			if lastStageVal != "XlinkStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "XlinkStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xlink_schemaelementIdent, __gong__toRawStringLiteral(xlink_schemaelement.Name)))
		}
	}
	if stage.GmlStage != nil {
		for _, abstractcurvesegmenttype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractCurveSegmentTypes, stage.GmlStage.AbstractCurveSegmentType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractcurvesegmenttypeIdent := abstractcurvesegmenttype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractCurveSegmentType{Name: %s}).Stage(stage.GmlStage)", abstractcurvesegmenttypeIdent, __gong__toRawStringLiteral(abstractcurvesegmenttype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractcurvesegmenttypeIdent, __gong__toRawStringLiteral(abstractcurvesegmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtStart = %d", abstractcurvesegmenttypeIdent, abstractcurvesegmenttype.NumDerivativesAtStart))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtEnd = %d", abstractcurvesegmenttypeIdent, abstractcurvesegmenttype.NumDerivativesAtEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativeInterior = %d", abstractcurvesegmenttypeIdent, abstractcurvesegmenttype.NumDerivativeInterior))
		}
	}
	if stage.GmlStage != nil {
		for _, abstractcurvetype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractCurveTypes, stage.GmlStage.AbstractCurveType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractcurvetypeIdent := abstractcurvetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractCurveType{Name: %s}).Stage(stage.GmlStage)", abstractcurvetypeIdent, __gong__toRawStringLiteral(abstractcurvetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractcurvetypeIdent, __gong__toRawStringLiteral(abstractcurvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractcurvetypeIdent, __gong__toRawStringLiteral(abstractcurvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractcurvetypeIdent, __gong__toRawStringLiteral(abstractcurvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", abstractcurvetypeIdent, __gong__toRawStringLiteral(abstractcurvetype.SrsName)))
			if abstractcurvetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstractcurvetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractcurvetypeIdent, targetIdent))
			}
			for _, elem := range abstractcurvetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractcurvetypeIdent, abstractcurvetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, abstractfeaturemembertype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractFeatureMemberTypes, stage.GmlStage.AbstractFeatureMemberType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractfeaturemembertypeIdent := abstractfeaturemembertype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractFeatureMemberType{Name: %s}).Stage(stage.GmlStage)", abstractfeaturemembertypeIdent, __gong__toRawStringLiteral(abstractfeaturemembertype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractfeaturemembertypeIdent, __gong__toRawStringLiteral(abstractfeaturemembertype.Name)))
		}
	}
	if stage.GmlStage != nil {
		for _, abstractfeaturetype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractFeatureTypes, stage.GmlStage.AbstractFeatureType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractfeaturetypeIdent := abstractfeaturetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractFeatureType{Name: %s}).Stage(stage.GmlStage)", abstractfeaturetypeIdent, __gong__toRawStringLiteral(abstractfeaturetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractfeaturetypeIdent, __gong__toRawStringLiteral(abstractfeaturetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractfeaturetypeIdent, __gong__toRawStringLiteral(abstractfeaturetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractfeaturetypeIdent, __gong__toRawStringLiteral(abstractfeaturetype.Description)))
			if abstractfeaturetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstractfeaturetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractfeaturetypeIdent, targetIdent))
			}
			for _, elem := range abstractfeaturetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractfeaturetypeIdent, abstractfeaturetypeIdent, targetIdent))
			}
			if abstractfeaturetype.BoundedBy != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstractfeaturetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractfeaturetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, abstractgmltype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractGMLTypes, stage.GmlStage.AbstractGMLType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractgmltypeIdent := abstractgmltype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractGMLType{Name: %s}).Stage(stage.GmlStage)", abstractgmltypeIdent, __gong__toRawStringLiteral(abstractgmltype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractgmltypeIdent, __gong__toRawStringLiteral(abstractgmltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractgmltypeIdent, __gong__toRawStringLiteral(abstractgmltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractgmltypeIdent, __gong__toRawStringLiteral(abstractgmltype.Description)))
			if abstractgmltype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstractgmltype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractgmltypeIdent, targetIdent))
			}
			for _, elem := range abstractgmltype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractgmltypeIdent, abstractgmltypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, abstractgeometricprimitivetype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractGeometricPrimitiveTypes, stage.GmlStage.AbstractGeometricPrimitiveType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractgeometricprimitivetypeIdent := abstractgeometricprimitivetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractGeometricPrimitiveType{Name: %s}).Stage(stage.GmlStage)", abstractgeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstractgeometricprimitivetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractgeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstractgeometricprimitivetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractgeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstractgeometricprimitivetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractgeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstractgeometricprimitivetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", abstractgeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstractgeometricprimitivetype.SrsName)))
			if abstractgeometricprimitivetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstractgeometricprimitivetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractgeometricprimitivetypeIdent, targetIdent))
			}
			for _, elem := range abstractgeometricprimitivetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractgeometricprimitivetypeIdent, abstractgeometricprimitivetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, abstractgeometrytype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractGeometryTypes, stage.GmlStage.AbstractGeometryType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractgeometrytypeIdent := abstractgeometrytype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractGeometryType{Name: %s}).Stage(stage.GmlStage)", abstractgeometrytypeIdent, __gong__toRawStringLiteral(abstractgeometrytype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractgeometrytypeIdent, __gong__toRawStringLiteral(abstractgeometrytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractgeometrytypeIdent, __gong__toRawStringLiteral(abstractgeometrytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractgeometrytypeIdent, __gong__toRawStringLiteral(abstractgeometrytype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", abstractgeometrytypeIdent, __gong__toRawStringLiteral(abstractgeometrytype.SrsName)))
			if abstractgeometrytype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstractgeometrytype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractgeometrytypeIdent, targetIdent))
			}
			for _, elem := range abstractgeometrytype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractgeometrytypeIdent, abstractgeometrytypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, abstractmetadatapropertytype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractMetadataPropertyTypes, stage.GmlStage.AbstractMetadataPropertyType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractmetadatapropertytypeIdent := abstractmetadatapropertytype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractMetadataPropertyType{Name: %s}).Stage(stage.GmlStage)", abstractmetadatapropertytypeIdent, __gong__toRawStringLiteral(abstractmetadatapropertytype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractmetadatapropertytypeIdent, __gong__toRawStringLiteral(abstractmetadatapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", abstractmetadatapropertytypeIdent, abstractmetadatapropertytype.Owns))
		}
	}
	if stage.GmlStage != nil {
		for _, abstractringpropertytype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractRingPropertyTypes, stage.GmlStage.AbstractRingPropertyType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractringpropertytypeIdent := abstractringpropertytype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractRingPropertyType{Name: %s}).Stage(stage.GmlStage)", abstractringpropertytypeIdent, __gong__toRawStringLiteral(abstractringpropertytype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractringpropertytypeIdent, __gong__toRawStringLiteral(abstractringpropertytype.Name)))
			if abstractringpropertytype.Ring != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstractringpropertytype.Ring.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ring = %s", abstractringpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, abstractringtype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractRingTypes, stage.GmlStage.AbstractRingType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractringtypeIdent := abstractringtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractRingType{Name: %s}).Stage(stage.GmlStage)", abstractringtypeIdent, __gong__toRawStringLiteral(abstractringtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractringtypeIdent, __gong__toRawStringLiteral(abstractringtype.Name)))
		}
	}
	if stage.GmlStage != nil {
		for _, abstractsurfacepatchtype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractSurfacePatchTypes, stage.GmlStage.AbstractSurfacePatchType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractsurfacepatchtypeIdent := abstractsurfacepatchtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractSurfacePatchType{Name: %s}).Stage(stage.GmlStage)", abstractsurfacepatchtypeIdent, __gong__toRawStringLiteral(abstractsurfacepatchtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractsurfacepatchtypeIdent, __gong__toRawStringLiteral(abstractsurfacepatchtype.Name)))
		}
	}
	if stage.GmlStage != nil {
		for _, abstractsurfacetype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractSurfaceTypes, stage.GmlStage.AbstractSurfaceType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstractsurfacetypeIdent := abstractsurfacetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractSurfaceType{Name: %s}).Stage(stage.GmlStage)", abstractsurfacetypeIdent, __gong__toRawStringLiteral(abstractsurfacetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractsurfacetypeIdent, __gong__toRawStringLiteral(abstractsurfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractsurfacetypeIdent, __gong__toRawStringLiteral(abstractsurfacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractsurfacetypeIdent, __gong__toRawStringLiteral(abstractsurfacetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", abstractsurfacetypeIdent, __gong__toRawStringLiteral(abstractsurfacetype.SrsName)))
			if abstractsurfacetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstractsurfacetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractsurfacetypeIdent, targetIdent))
			}
			for _, elem := range abstractsurfacetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractsurfacetypeIdent, abstractsurfacetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, abstracttimegeometricprimitivetype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractTimeGeometricPrimitiveTypes, stage.GmlStage.AbstractTimeGeometricPrimitiveType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstracttimegeometricprimitivetypeIdent := abstracttimegeometricprimitivetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractTimeGeometricPrimitiveType{Name: %s}).Stage(stage.GmlStage)", abstracttimegeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimegeometricprimitivetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstracttimegeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimegeometricprimitivetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstracttimegeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimegeometricprimitivetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstracttimegeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimegeometricprimitivetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", abstracttimegeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimegeometricprimitivetype.Frame)))
			if abstracttimegeometricprimitivetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstracttimegeometricprimitivetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstracttimegeometricprimitivetypeIdent, targetIdent))
			}
			for _, elem := range abstracttimegeometricprimitivetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstracttimegeometricprimitivetypeIdent, abstracttimegeometricprimitivetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, abstracttimeobjecttype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractTimeObjectTypes, stage.GmlStage.AbstractTimeObjectType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstracttimeobjecttypeIdent := abstracttimeobjecttype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractTimeObjectType{Name: %s}).Stage(stage.GmlStage)", abstracttimeobjecttypeIdent, __gong__toRawStringLiteral(abstracttimeobjecttype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstracttimeobjecttypeIdent, __gong__toRawStringLiteral(abstracttimeobjecttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstracttimeobjecttypeIdent, __gong__toRawStringLiteral(abstracttimeobjecttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstracttimeobjecttypeIdent, __gong__toRawStringLiteral(abstracttimeobjecttype.Description)))
			if abstracttimeobjecttype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstracttimeobjecttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstracttimeobjecttypeIdent, targetIdent))
			}
			for _, elem := range abstracttimeobjecttype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstracttimeobjecttypeIdent, abstracttimeobjecttypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, abstracttimeprimitivetype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractTimePrimitiveTypes, stage.GmlStage.AbstractTimePrimitiveType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstracttimeprimitivetypeIdent := abstracttimeprimitivetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractTimePrimitiveType{Name: %s}).Stage(stage.GmlStage)", abstracttimeprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimeprimitivetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstracttimeprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimeprimitivetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstracttimeprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimeprimitivetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstracttimeprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimeprimitivetype.Description)))
			if abstracttimeprimitivetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstracttimeprimitivetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstracttimeprimitivetypeIdent, targetIdent))
			}
			for _, elem := range abstracttimeprimitivetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstracttimeprimitivetypeIdent, abstracttimeprimitivetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, abstracttimeslicetype := range __gong__sortStageSetInstances(stage.GmlStage.AbstractTimeSliceTypes, stage.GmlStage.AbstractTimeSliceType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			abstracttimeslicetypeIdent := abstracttimeslicetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractTimeSliceType{Name: %s}).Stage(stage.GmlStage)", abstracttimeslicetypeIdent, __gong__toRawStringLiteral(abstracttimeslicetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstracttimeslicetypeIdent, __gong__toRawStringLiteral(abstracttimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstracttimeslicetypeIdent, __gong__toRawStringLiteral(abstracttimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstracttimeslicetypeIdent, __gong__toRawStringLiteral(abstracttimeslicetype.Description)))
			if abstracttimeslicetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstracttimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstracttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range abstracttimeslicetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstracttimeslicetypeIdent, abstracttimeslicetypeIdent, targetIdent))
			}
			if abstracttimeslicetype.ValidTime != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := abstracttimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", abstracttimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, angletype := range __gong__sortStageSetInstances(stage.GmlStage.AngleTypes, stage.GmlStage.AngleType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			angletypeIdent := angletype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AngleType{Name: %s}).Stage(stage.GmlStage)", angletypeIdent, __gong__toRawStringLiteral(angletype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angletypeIdent, __gong__toRawStringLiteral(angletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", angletypeIdent, __gong__toRawStringLiteral(angletype.EnclosedText)))
		}
	}
	if stage.GmlStage != nil {
		for _, arcbycenterpointtype := range __gong__sortStageSetInstances(stage.GmlStage.ArcByCenterPointTypes, stage.GmlStage.ArcByCenterPointType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			arcbycenterpointtypeIdent := arcbycenterpointtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.ArcByCenterPointType{Name: %s}).Stage(stage.GmlStage)", arcbycenterpointtypeIdent, __gong__toRawStringLiteral(arcbycenterpointtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arcbycenterpointtypeIdent, __gong__toRawStringLiteral(arcbycenterpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtStart = %d", arcbycenterpointtypeIdent, arcbycenterpointtype.NumDerivativesAtStart))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtEnd = %d", arcbycenterpointtypeIdent, arcbycenterpointtype.NumDerivativesAtEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativeInterior = %d", arcbycenterpointtypeIdent, arcbycenterpointtype.NumDerivativeInterior))
			if arcbycenterpointtype.Interpolation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = gml.%s", arcbycenterpointtypeIdent, arcbycenterpointtype.Interpolation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = \"\"", arcbycenterpointtypeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.NumArc = %d", arcbycenterpointtypeIdent, arcbycenterpointtype.NumArc))
			if arcbycenterpointtype.Pos != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := arcbycenterpointtype.Pos.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", arcbycenterpointtypeIdent, targetIdent))
			}
			if arcbycenterpointtype.PointProperty != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := arcbycenterpointtype.PointProperty.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = %s", arcbycenterpointtypeIdent, targetIdent))
			}
			if arcbycenterpointtype.PosList != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := arcbycenterpointtype.PosList.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", arcbycenterpointtypeIdent, targetIdent))
			}
			if arcbycenterpointtype.Radius != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := arcbycenterpointtype.Radius.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Radius = %s", arcbycenterpointtypeIdent, targetIdent))
			}
			if arcbycenterpointtype.StartAngle != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := arcbycenterpointtype.StartAngle.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartAngle = %s", arcbycenterpointtypeIdent, targetIdent))
			}
			if arcbycenterpointtype.EndAngle != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := arcbycenterpointtype.EndAngle.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndAngle = %s", arcbycenterpointtypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, arcstringtype := range __gong__sortStageSetInstances(stage.GmlStage.ArcStringTypes, stage.GmlStage.ArcStringType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			arcstringtypeIdent := arcstringtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.ArcStringType{Name: %s}).Stage(stage.GmlStage)", arcstringtypeIdent, __gong__toRawStringLiteral(arcstringtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arcstringtypeIdent, __gong__toRawStringLiteral(arcstringtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtStart = %d", arcstringtypeIdent, arcstringtype.NumDerivativesAtStart))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtEnd = %d", arcstringtypeIdent, arcstringtype.NumDerivativesAtEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativeInterior = %d", arcstringtypeIdent, arcstringtype.NumDerivativeInterior))
			if arcstringtype.Interpolation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = gml.%s", arcstringtypeIdent, arcstringtype.Interpolation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = \"\"", arcstringtypeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.NumArc = %d", arcstringtypeIdent, arcstringtype.NumArc))
			for _, elem := range arcstringtype.Pos {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = append(%s.Pos, %s)", arcstringtypeIdent, arcstringtypeIdent, targetIdent))
			}
			for _, elem := range arcstringtype.PointProperty {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = append(%s.PointProperty, %s)", arcstringtypeIdent, arcstringtypeIdent, targetIdent))
			}
			if arcstringtype.PosList != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := arcstringtype.PosList.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", arcstringtypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, arctype := range __gong__sortStageSetInstances(stage.GmlStage.ArcTypes, stage.GmlStage.ArcType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			arctypeIdent := arctype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.ArcType{Name: %s}).Stage(stage.GmlStage)", arctypeIdent, __gong__toRawStringLiteral(arctype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtStart = %d", arctypeIdent, arctype.NumDerivativesAtStart))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtEnd = %d", arctypeIdent, arctype.NumDerivativesAtEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativeInterior = %d", arctypeIdent, arctype.NumDerivativeInterior))
			if arctype.Interpolation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = gml.%s", arctypeIdent, arctype.Interpolation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = \"\"", arctypeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.NumArc = %d", arctypeIdent, arctype.NumArc))
			if arctype.Pos != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := arctype.Pos.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", arctypeIdent, targetIdent))
			}
			if arctype.PointProperty != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := arctype.PointProperty.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = %s", arctypeIdent, targetIdent))
			}
			if arctype.PosList != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := arctype.PosList.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", arctypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, boundingshapetype := range __gong__sortStageSetInstances(stage.GmlStage.BoundingShapeTypes, stage.GmlStage.BoundingShapeType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			boundingshapetypeIdent := boundingshapetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.BoundingShapeType{Name: %s}).Stage(stage.GmlStage)", boundingshapetypeIdent, __gong__toRawStringLiteral(boundingshapetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", boundingshapetypeIdent, __gong__toRawStringLiteral(boundingshapetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", boundingshapetypeIdent, __gong__toRawStringLiteral(boundingshapetype.NilReason)))
			if boundingshapetype.Envelope != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := boundingshapetype.Envelope.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Envelope = %s", boundingshapetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, circlebycenterpointtype := range __gong__sortStageSetInstances(stage.GmlStage.CircleByCenterPointTypes, stage.GmlStage.CircleByCenterPointType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			circlebycenterpointtypeIdent := circlebycenterpointtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CircleByCenterPointType{Name: %s}).Stage(stage.GmlStage)", circlebycenterpointtypeIdent, __gong__toRawStringLiteral(circlebycenterpointtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circlebycenterpointtypeIdent, __gong__toRawStringLiteral(circlebycenterpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtStart = %d", circlebycenterpointtypeIdent, circlebycenterpointtype.NumDerivativesAtStart))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtEnd = %d", circlebycenterpointtypeIdent, circlebycenterpointtype.NumDerivativesAtEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativeInterior = %d", circlebycenterpointtypeIdent, circlebycenterpointtype.NumDerivativeInterior))
			if circlebycenterpointtype.Interpolation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = gml.%s", circlebycenterpointtypeIdent, circlebycenterpointtype.Interpolation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = \"\"", circlebycenterpointtypeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.NumArc = %d", circlebycenterpointtypeIdent, circlebycenterpointtype.NumArc))
			if circlebycenterpointtype.Pos != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := circlebycenterpointtype.Pos.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", circlebycenterpointtypeIdent, targetIdent))
			}
			if circlebycenterpointtype.PointProperty != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := circlebycenterpointtype.PointProperty.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = %s", circlebycenterpointtypeIdent, targetIdent))
			}
			if circlebycenterpointtype.PosList != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := circlebycenterpointtype.PosList.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", circlebycenterpointtypeIdent, targetIdent))
			}
			if circlebycenterpointtype.Radius != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := circlebycenterpointtype.Radius.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Radius = %s", circlebycenterpointtypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, circletype := range __gong__sortStageSetInstances(stage.GmlStage.CircleTypes, stage.GmlStage.CircleType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			circletypeIdent := circletype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CircleType{Name: %s}).Stage(stage.GmlStage)", circletypeIdent, __gong__toRawStringLiteral(circletype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circletypeIdent, __gong__toRawStringLiteral(circletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtStart = %d", circletypeIdent, circletype.NumDerivativesAtStart))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtEnd = %d", circletypeIdent, circletype.NumDerivativesAtEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativeInterior = %d", circletypeIdent, circletype.NumDerivativeInterior))
			if circletype.Interpolation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = gml.%s", circletypeIdent, circletype.Interpolation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = \"\"", circletypeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.NumArc = %d", circletypeIdent, circletype.NumArc))
			for _, elem := range circletype.Pos {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = append(%s.Pos, %s)", circletypeIdent, circletypeIdent, targetIdent))
			}
			for _, elem := range circletype.PointProperty {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = append(%s.PointProperty, %s)", circletypeIdent, circletypeIdent, targetIdent))
			}
			if circletype.PosList != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := circletype.PosList.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", circletypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, codetype := range __gong__sortStageSetInstances(stage.GmlStage.CodeTypes, stage.GmlStage.CodeType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			codetypeIdent := codetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CodeType{Name: %s}).Stage(stage.GmlStage)", codetypeIdent, __gong__toRawStringLiteral(codetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetypeIdent, __gong__toRawStringLiteral(codetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CodeSpace = %s", codetypeIdent, __gong__toRawStringLiteral(codetype.CodeSpace)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetypeIdent, __gong__toRawStringLiteral(codetype.EnclosedText)))
		}
	}
	if stage.GmlStage != nil {
		for _, codewithauthoritytype := range __gong__sortStageSetInstances(stage.GmlStage.CodeWithAuthorityTypes, stage.GmlStage.CodeWithAuthorityType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			codewithauthoritytypeIdent := codewithauthoritytype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CodeWithAuthorityType{Name: %s}).Stage(stage.GmlStage)", codewithauthoritytypeIdent, __gong__toRawStringLiteral(codewithauthoritytype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codewithauthoritytypeIdent, __gong__toRawStringLiteral(codewithauthoritytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CodeSpace = %s", codewithauthoritytypeIdent, __gong__toRawStringLiteral(codewithauthoritytype.CodeSpace)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codewithauthoritytypeIdent, __gong__toRawStringLiteral(codewithauthoritytype.EnclosedText)))
		}
	}
	if stage.GmlStage != nil {
		for _, compositecurvetype := range __gong__sortStageSetInstances(stage.GmlStage.CompositeCurveTypes, stage.GmlStage.CompositeCurveType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			compositecurvetypeIdent := compositecurvetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CompositeCurveType{Name: %s}).Stage(stage.GmlStage)", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.SrsName)))
			values.WriteString(fmt.Sprintf("\n\t%s.AggregationType = %s", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.AggregationType)))
			if compositecurvetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := compositecurvetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", compositecurvetypeIdent, targetIdent))
			}
			for _, elem := range compositecurvetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", compositecurvetypeIdent, compositecurvetypeIdent, targetIdent))
			}
			for _, elem := range compositecurvetype.CurveMember {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CurveMember = append(%s.CurveMember, %s)", compositecurvetypeIdent, compositecurvetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, curvepropertytype := range __gong__sortStageSetInstances(stage.GmlStage.CurvePropertyTypes, stage.GmlStage.CurvePropertyType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			curvepropertytypeIdent := curvepropertytype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CurvePropertyType{Name: %s}).Stage(stage.GmlStage)", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Actuate)))
			if curvepropertytype.Curve != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := curvepropertytype.Curve.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Curve = %s", curvepropertytypeIdent, targetIdent))
			}
			if curvepropertytype.OrientableCurve != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := curvepropertytype.OrientableCurve.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OrientableCurve = %s", curvepropertytypeIdent, targetIdent))
			}
			if curvepropertytype.CompositeCurve != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := curvepropertytype.CompositeCurve.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompositeCurve = %s", curvepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, curvesegmentarraypropertytype := range __gong__sortStageSetInstances(stage.GmlStage.CurveSegmentArrayPropertyTypes, stage.GmlStage.CurveSegmentArrayPropertyType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			curvesegmentarraypropertytypeIdent := curvesegmentarraypropertytype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CurveSegmentArrayPropertyType{Name: %s}).Stage(stage.GmlStage)", curvesegmentarraypropertytypeIdent, __gong__toRawStringLiteral(curvesegmentarraypropertytype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvesegmentarraypropertytypeIdent, __gong__toRawStringLiteral(curvesegmentarraypropertytype.Name)))
			for _, elem := range curvesegmentarraypropertytype.LineStringSegment {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LineStringSegment = append(%s.LineStringSegment, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.ArcString {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcString = append(%s.ArcString, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.Arc {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arc = append(%s.Arc, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.Circle {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Circle = append(%s.Circle, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.ArcByCenterPoint {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcByCenterPoint = append(%s.ArcByCenterPoint, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.CircleByCenterPoint {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CircleByCenterPoint = append(%s.CircleByCenterPoint, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.GeodesicString {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeodesicString = append(%s.GeodesicString, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.Geodesic {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Geodesic = append(%s.Geodesic, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, curvetype := range __gong__sortStageSetInstances(stage.GmlStage.CurveTypes, stage.GmlStage.CurveType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			curvetypeIdent := curvetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CurveType{Name: %s}).Stage(stage.GmlStage)", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.SrsName)))
			if curvetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := curvetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", curvetypeIdent, targetIdent))
			}
			for _, elem := range curvetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", curvetypeIdent, curvetypeIdent, targetIdent))
			}
			if curvetype.Segments != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := curvetype.Segments.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Segments = %s", curvetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, directpositionlisttype := range __gong__sortStageSetInstances(stage.GmlStage.DirectPositionListTypes, stage.GmlStage.DirectPositionListType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			directpositionlisttypeIdent := directpositionlisttype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.DirectPositionListType{Name: %s}).Stage(stage.GmlStage)", directpositionlisttypeIdent, __gong__toRawStringLiteral(directpositionlisttype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directpositionlisttypeIdent, __gong__toRawStringLiteral(directpositionlisttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", directpositionlisttypeIdent, __gong__toRawStringLiteral(directpositionlisttype.SrsName)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", directpositionlisttypeIdent, __gong__toRawStringLiteral(directpositionlisttype.EnclosedText)))
		}
	}
	if stage.GmlStage != nil {
		for _, directpositiontype := range __gong__sortStageSetInstances(stage.GmlStage.DirectPositionTypes, stage.GmlStage.DirectPositionType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			directpositiontypeIdent := directpositiontype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.DirectPositionType{Name: %s}).Stage(stage.GmlStage)", directpositiontypeIdent, __gong__toRawStringLiteral(directpositiontype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directpositiontypeIdent, __gong__toRawStringLiteral(directpositiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", directpositiontypeIdent, __gong__toRawStringLiteral(directpositiontype.SrsName)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", directpositiontypeIdent, __gong__toRawStringLiteral(directpositiontype.EnclosedText)))
		}
	}
	if stage.GmlStage != nil {
		for _, dynamicfeaturetype := range __gong__sortStageSetInstances(stage.GmlStage.DynamicFeatureTypes, stage.GmlStage.DynamicFeatureType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			dynamicfeaturetypeIdent := dynamicfeaturetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.DynamicFeatureType{Name: %s}).Stage(stage.GmlStage)", dynamicfeaturetypeIdent, __gong__toRawStringLiteral(dynamicfeaturetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dynamicfeaturetypeIdent, __gong__toRawStringLiteral(dynamicfeaturetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", dynamicfeaturetypeIdent, __gong__toRawStringLiteral(dynamicfeaturetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dynamicfeaturetypeIdent, __gong__toRawStringLiteral(dynamicfeaturetype.Description)))
			if dynamicfeaturetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := dynamicfeaturetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", dynamicfeaturetypeIdent, targetIdent))
			}
			for _, elem := range dynamicfeaturetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", dynamicfeaturetypeIdent, dynamicfeaturetypeIdent, targetIdent))
			}
			if dynamicfeaturetype.BoundedBy != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := dynamicfeaturetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", dynamicfeaturetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, envelopetype := range __gong__sortStageSetInstances(stage.GmlStage.EnvelopeTypes, stage.GmlStage.EnvelopeType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			envelopetypeIdent := envelopetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.EnvelopeType{Name: %s}).Stage(stage.GmlStage)", envelopetypeIdent, __gong__toRawStringLiteral(envelopetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", envelopetypeIdent, __gong__toRawStringLiteral(envelopetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", envelopetypeIdent, __gong__toRawStringLiteral(envelopetype.SrsName)))
			if envelopetype.LowerCorner != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := envelopetype.LowerCorner.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerCorner = %s", envelopetypeIdent, targetIdent))
			}
			if envelopetype.UpperCorner != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := envelopetype.UpperCorner.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperCorner = %s", envelopetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, geodesicstringtype := range __gong__sortStageSetInstances(stage.GmlStage.GeodesicStringTypes, stage.GmlStage.GeodesicStringType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			geodesicstringtypeIdent := geodesicstringtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.GeodesicStringType{Name: %s}).Stage(stage.GmlStage)", geodesicstringtypeIdent, __gong__toRawStringLiteral(geodesicstringtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geodesicstringtypeIdent, __gong__toRawStringLiteral(geodesicstringtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtStart = %d", geodesicstringtypeIdent, geodesicstringtype.NumDerivativesAtStart))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtEnd = %d", geodesicstringtypeIdent, geodesicstringtype.NumDerivativesAtEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativeInterior = %d", geodesicstringtypeIdent, geodesicstringtype.NumDerivativeInterior))
			if geodesicstringtype.Interpolation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = gml.%s", geodesicstringtypeIdent, geodesicstringtype.Interpolation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = \"\"", geodesicstringtypeIdent))
			}
			if geodesicstringtype.PosList != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := geodesicstringtype.PosList.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", geodesicstringtypeIdent, targetIdent))
			}
			if geodesicstringtype.Pos != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := geodesicstringtype.Pos.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", geodesicstringtypeIdent, targetIdent))
			}
			if geodesicstringtype.PointProperty != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := geodesicstringtype.PointProperty.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = %s", geodesicstringtypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, geodesictype := range __gong__sortStageSetInstances(stage.GmlStage.GeodesicTypes, stage.GmlStage.GeodesicType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			geodesictypeIdent := geodesictype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.GeodesicType{Name: %s}).Stage(stage.GmlStage)", geodesictypeIdent, __gong__toRawStringLiteral(geodesictype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geodesictypeIdent, __gong__toRawStringLiteral(geodesictype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtStart = %d", geodesictypeIdent, geodesictype.NumDerivativesAtStart))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtEnd = %d", geodesictypeIdent, geodesictype.NumDerivativesAtEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativeInterior = %d", geodesictypeIdent, geodesictype.NumDerivativeInterior))
			if geodesictype.Interpolation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = gml.%s", geodesictypeIdent, geodesictype.Interpolation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = \"\"", geodesictypeIdent))
			}
			if geodesictype.PosList != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := geodesictype.PosList.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", geodesictypeIdent, targetIdent))
			}
			if geodesictype.Pos != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := geodesictype.Pos.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", geodesictypeIdent, targetIdent))
			}
			if geodesictype.PointProperty != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := geodesictype.PointProperty.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = %s", geodesictypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, lengthtype := range __gong__sortStageSetInstances(stage.GmlStage.LengthTypes, stage.GmlStage.LengthType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			lengthtypeIdent := lengthtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.LengthType{Name: %s}).Stage(stage.GmlStage)", lengthtypeIdent, __gong__toRawStringLiteral(lengthtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lengthtypeIdent, __gong__toRawStringLiteral(lengthtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", lengthtypeIdent, __gong__toRawStringLiteral(lengthtype.EnclosedText)))
		}
	}
	if stage.GmlStage != nil {
		for _, linestringsegmenttype := range __gong__sortStageSetInstances(stage.GmlStage.LineStringSegmentTypes, stage.GmlStage.LineStringSegmentType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			linestringsegmenttypeIdent := linestringsegmenttype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.LineStringSegmentType{Name: %s}).Stage(stage.GmlStage)", linestringsegmenttypeIdent, __gong__toRawStringLiteral(linestringsegmenttype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linestringsegmenttypeIdent, __gong__toRawStringLiteral(linestringsegmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtStart = %d", linestringsegmenttypeIdent, linestringsegmenttype.NumDerivativesAtStart))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtEnd = %d", linestringsegmenttypeIdent, linestringsegmenttype.NumDerivativesAtEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativeInterior = %d", linestringsegmenttypeIdent, linestringsegmenttype.NumDerivativeInterior))
			if linestringsegmenttype.Interpolation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = gml.%s", linestringsegmenttypeIdent, linestringsegmenttype.Interpolation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = \"\"", linestringsegmenttypeIdent))
			}
			for _, elem := range linestringsegmenttype.Pos {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = append(%s.Pos, %s)", linestringsegmenttypeIdent, linestringsegmenttypeIdent, targetIdent))
			}
			for _, elem := range linestringsegmenttype.PointProperty {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = append(%s.PointProperty, %s)", linestringsegmenttypeIdent, linestringsegmenttypeIdent, targetIdent))
			}
			if linestringsegmenttype.PosList != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := linestringsegmenttype.PosList.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", linestringsegmenttypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, measuretype := range __gong__sortStageSetInstances(stage.GmlStage.MeasureTypes, stage.GmlStage.MeasureType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			measuretypeIdent := measuretype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.MeasureType{Name: %s}).Stage(stage.GmlStage)", measuretypeIdent, __gong__toRawStringLiteral(measuretype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measuretypeIdent, __gong__toRawStringLiteral(measuretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", measuretypeIdent, __gong__toRawStringLiteral(measuretype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %f", measuretypeIdent, measuretype.EnclosedText))
		}
	}
	if stage.GmlStage != nil {
		for _, orientablecurvetype := range __gong__sortStageSetInstances(stage.GmlStage.OrientableCurveTypes, stage.GmlStage.OrientableCurveType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			orientablecurvetypeIdent := orientablecurvetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.OrientableCurveType{Name: %s}).Stage(stage.GmlStage)", orientablecurvetypeIdent, __gong__toRawStringLiteral(orientablecurvetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", orientablecurvetypeIdent, __gong__toRawStringLiteral(orientablecurvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", orientablecurvetypeIdent, __gong__toRawStringLiteral(orientablecurvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", orientablecurvetypeIdent, __gong__toRawStringLiteral(orientablecurvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", orientablecurvetypeIdent, __gong__toRawStringLiteral(orientablecurvetype.SrsName)))
			if orientablecurvetype.Orientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Orientation = gml.%s", orientablecurvetypeIdent, orientablecurvetype.Orientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Orientation = \"\"", orientablecurvetypeIdent))
			}
			if orientablecurvetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := orientablecurvetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", orientablecurvetypeIdent, targetIdent))
			}
			for _, elem := range orientablecurvetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", orientablecurvetypeIdent, orientablecurvetypeIdent, targetIdent))
			}
			if orientablecurvetype.BaseCurve != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := orientablecurvetype.BaseCurve.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BaseCurve = %s", orientablecurvetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, pointpropertytype := range __gong__sortStageSetInstances(stage.GmlStage.PointPropertyTypes, stage.GmlStage.PointPropertyType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			pointpropertytypeIdent := pointpropertytype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.PointPropertyType{Name: %s}).Stage(stage.GmlStage)", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Actuate)))
			if pointpropertytype.Point != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := pointpropertytype.Point.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Point = %s", pointpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, pointtype := range __gong__sortStageSetInstances(stage.GmlStage.PointTypes, stage.GmlStage.PointType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			pointtypeIdent := pointtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.PointType{Name: %s}).Stage(stage.GmlStage)", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.SrsName)))
			if pointtype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := pointtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", pointtypeIdent, targetIdent))
			}
			for _, elem := range pointtype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", pointtypeIdent, pointtypeIdent, targetIdent))
			}
			if pointtype.Pos != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := pointtype.Pos.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", pointtypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, polygonpatchtype := range __gong__sortStageSetInstances(stage.GmlStage.PolygonPatchTypes, stage.GmlStage.PolygonPatchType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			polygonpatchtypeIdent := polygonpatchtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.PolygonPatchType{Name: %s}).Stage(stage.GmlStage)", polygonpatchtypeIdent, __gong__toRawStringLiteral(polygonpatchtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", polygonpatchtypeIdent, __gong__toRawStringLiteral(polygonpatchtype.Name)))
			if polygonpatchtype.Interpolation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = gml.%s", polygonpatchtypeIdent, polygonpatchtype.Interpolation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = \"\"", polygonpatchtypeIdent))
			}
			if polygonpatchtype.Exterior != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := polygonpatchtype.Exterior.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Exterior = %s", polygonpatchtypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, ringtype := range __gong__sortStageSetInstances(stage.GmlStage.RingTypes, stage.GmlStage.RingType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			ringtypeIdent := ringtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.RingType{Name: %s}).Stage(stage.GmlStage)", ringtypeIdent, __gong__toRawStringLiteral(ringtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ringtypeIdent, __gong__toRawStringLiteral(ringtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.AggregationType = %s", ringtypeIdent, __gong__toRawStringLiteral(ringtype.AggregationType)))
			for _, elem := range ringtype.CurveMember {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CurveMember = append(%s.CurveMember, %s)", ringtypeIdent, ringtypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, surfacepatcharraypropertytype := range __gong__sortStageSetInstances(stage.GmlStage.SurfacePatchArrayPropertyTypes, stage.GmlStage.SurfacePatchArrayPropertyType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			surfacepatcharraypropertytypeIdent := surfacepatcharraypropertytype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.SurfacePatchArrayPropertyType{Name: %s}).Stage(stage.GmlStage)", surfacepatcharraypropertytypeIdent, __gong__toRawStringLiteral(surfacepatcharraypropertytype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surfacepatcharraypropertytypeIdent, __gong__toRawStringLiteral(surfacepatcharraypropertytype.Name)))
			for _, elem := range surfacepatcharraypropertytype.PolygonPatch {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PolygonPatch = append(%s.PolygonPatch, %s)", surfacepatcharraypropertytypeIdent, surfacepatcharraypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, surfacetype := range __gong__sortStageSetInstances(stage.GmlStage.SurfaceTypes, stage.GmlStage.SurfaceType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			surfacetypeIdent := surfacetype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.SurfaceType{Name: %s}).Stage(stage.GmlStage)", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.SrsName)))
			if surfacetype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := surfacetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", surfacetypeIdent, targetIdent))
			}
			for _, elem := range surfacetype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", surfacetypeIdent, surfacetypeIdent, targetIdent))
			}
			if surfacetype.Patches != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := surfacetype.Patches.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Patches = %s", surfacetypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, timeinstanttype := range __gong__sortStageSetInstances(stage.GmlStage.TimeInstantTypes, stage.GmlStage.TimeInstantType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			timeinstanttypeIdent := timeinstanttype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.TimeInstantType{Name: %s}).Stage(stage.GmlStage)", timeinstanttypeIdent, __gong__toRawStringLiteral(timeinstanttype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timeinstanttypeIdent, __gong__toRawStringLiteral(timeinstanttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", timeinstanttypeIdent, __gong__toRawStringLiteral(timeinstanttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", timeinstanttypeIdent, __gong__toRawStringLiteral(timeinstanttype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", timeinstanttypeIdent, __gong__toRawStringLiteral(timeinstanttype.Frame)))
			if timeinstanttype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := timeinstanttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", timeinstanttypeIdent, targetIdent))
			}
			for _, elem := range timeinstanttype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", timeinstanttypeIdent, timeinstanttypeIdent, targetIdent))
			}
			if timeinstanttype.TimePosition != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := timeinstanttype.TimePosition.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimePosition = %s", timeinstanttypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, timeperiodtype := range __gong__sortStageSetInstances(stage.GmlStage.TimePeriodTypes, stage.GmlStage.TimePeriodType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			timeperiodtypeIdent := timeperiodtype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.TimePeriodType{Name: %s}).Stage(stage.GmlStage)", timeperiodtypeIdent, __gong__toRawStringLiteral(timeperiodtype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timeperiodtypeIdent, __gong__toRawStringLiteral(timeperiodtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", timeperiodtypeIdent, __gong__toRawStringLiteral(timeperiodtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", timeperiodtypeIdent, __gong__toRawStringLiteral(timeperiodtype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", timeperiodtypeIdent, __gong__toRawStringLiteral(timeperiodtype.Frame)))
			if timeperiodtype.Identifier != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := timeperiodtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", timeperiodtypeIdent, targetIdent))
			}
			for _, elem := range timeperiodtype.NameXSD {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", timeperiodtypeIdent, timeperiodtypeIdent, targetIdent))
			}
			if timeperiodtype.BeginPosition != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := timeperiodtype.BeginPosition.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BeginPosition = %s", timeperiodtypeIdent, targetIdent))
			}
			if timeperiodtype.EndPosition != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := timeperiodtype.EndPosition.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPosition = %s", timeperiodtypeIdent, targetIdent))
			}
		}
	}
	if stage.GmlStage != nil {
		for _, timepositiontype := range __gong__sortStageSetInstances(stage.GmlStage.TimePositionTypes, stage.GmlStage.TimePositionType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			timepositiontypeIdent := timepositiontype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.TimePositionType{Name: %s}).Stage(stage.GmlStage)", timepositiontypeIdent, __gong__toRawStringLiteral(timepositiontype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timepositiontypeIdent, __gong__toRawStringLiteral(timepositiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", timepositiontypeIdent, __gong__toRawStringLiteral(timepositiontype.Frame)))
			if timepositiontype.IndeterminatePosition.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.IndeterminatePosition = gml.%s", timepositiontypeIdent, timepositiontype.IndeterminatePosition.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.IndeterminatePosition = \"\"", timepositiontypeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", timepositiontypeIdent, __gong__toRawStringLiteral(timepositiontype.EnclosedText)))
		}
	}
	if stage.GmlStage != nil {
		for _, timeprimitivepropertytype := range __gong__sortStageSetInstances(stage.GmlStage.TimePrimitivePropertyTypes, stage.GmlStage.TimePrimitivePropertyType_stagedOrder) {
			if lastStageDecl != "GmlStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "GmlStage"
			}
			timeprimitivepropertytypeIdent := timeprimitivepropertytype.GongGetIdentifier(stage.GmlStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.TimePrimitivePropertyType{Name: %s}).Stage(stage.GmlStage)", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.Name)))
			if lastStageVal != "GmlStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "GmlStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.Actuate)))
			if timeprimitivepropertytype.TimeInstant != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := timeprimitivepropertytype.TimeInstant.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInstant = %s", timeprimitivepropertytypeIdent, targetIdent))
			}
			if timeprimitivepropertytype.TimePeriod != nil {
				if lastStagePtr != "GmlStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "GmlStage"
				}
				targetIdent := timeprimitivepropertytype.TimePeriod.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimePeriod = %s", timeprimitivepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, abstractaixmfeaturebasetype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureBaseTypes, stage.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureBaseType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			abstractaixmfeaturebasetypeIdent := abstractaixmfeaturebasetype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", abstractaixmfeaturebasetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturebasetype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmfeaturebasetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturebasetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmfeaturebasetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturebasetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractaixmfeaturebasetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturebasetype.Description)))
			if abstractaixmfeaturebasetype.Identifier != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmfeaturebasetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractaixmfeaturebasetypeIdent, targetIdent))
			}
			for _, elem := range abstractaixmfeaturebasetype.NameXSD {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractaixmfeaturebasetypeIdent, abstractaixmfeaturebasetypeIdent, targetIdent))
			}
			if abstractaixmfeaturebasetype.BoundedBy != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmfeaturebasetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractaixmfeaturebasetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, abstractaixmfeaturetype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureTypes, stage.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			abstractaixmfeaturetypeIdent := abstractaixmfeaturetype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMFeatureType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", abstractaixmfeaturetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturetype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmfeaturetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmfeaturetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractaixmfeaturetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturetype.Description)))
			if abstractaixmfeaturetype.Identifier != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmfeaturetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractaixmfeaturetypeIdent, targetIdent))
			}
			for _, elem := range abstractaixmfeaturetype.NameXSD {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractaixmfeaturetypeIdent, abstractaixmfeaturetypeIdent, targetIdent))
			}
			if abstractaixmfeaturetype.BoundedBy != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmfeaturetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractaixmfeaturetypeIdent, targetIdent))
			}
			if abstractaixmfeaturetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmfeaturetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", abstractaixmfeaturetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, abstractaixmmessagebasetype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseTypes, stage.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			abstractaixmmessagebasetypeIdent := abstractaixmmessagebasetype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", abstractaixmmessagebasetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagebasetype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmmessagebasetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagebasetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmmessagebasetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagebasetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractaixmmessagebasetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagebasetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", abstractaixmmessagebasetypeIdent, abstractaixmmessagebasetype.SequenceNumber))
			if abstractaixmmessagebasetype.Identifier != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmmessagebasetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractaixmmessagebasetypeIdent, targetIdent))
			}
			for _, elem := range abstractaixmmessagebasetype.NameXSD {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractaixmmessagebasetypeIdent, abstractaixmmessagebasetypeIdent, targetIdent))
			}
			if abstractaixmmessagebasetype.BoundedBy != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmmessagebasetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractaixmmessagebasetypeIdent, targetIdent))
			}
			if abstractaixmmessagebasetype.MessageMetadata != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmmessagebasetype.MessageMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MessageMetadata = %s", abstractaixmmessagebasetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, abstractaixmmessagetype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageTypes, stage.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			abstractaixmmessagetypeIdent := abstractaixmmessagetype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMMessageType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", abstractaixmmessagetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagetype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmmessagetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmmessagetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractaixmmessagetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", abstractaixmmessagetypeIdent, abstractaixmmessagetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.AggregationType = %s", abstractaixmmessagetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagetype.AggregationType)))
			if abstractaixmmessagetype.Identifier != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmmessagetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractaixmmessagetypeIdent, targetIdent))
			}
			for _, elem := range abstractaixmmessagetype.NameXSD {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractaixmmessagetypeIdent, abstractaixmmessagetypeIdent, targetIdent))
			}
			if abstractaixmmessagetype.BoundedBy != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmmessagetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractaixmmessagetypeIdent, targetIdent))
			}
			if abstractaixmmessagetype.MessageMetadata != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmmessagetype.MessageMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MessageMetadata = %s", abstractaixmmessagetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, abstractaixmobjecttype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.AbstractAIXMObjectTypes, stage.AixmabstractgmlobjecttypesStage.AbstractAIXMObjectType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			abstractaixmobjecttypeIdent := abstractaixmobjecttype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMObjectType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", abstractaixmobjecttypeIdent, __gong__toRawStringLiteral(abstractaixmobjecttype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmobjecttypeIdent, __gong__toRawStringLiteral(abstractaixmobjecttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmobjecttypeIdent, __gong__toRawStringLiteral(abstractaixmobjecttype.Id)))
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, abstractaixmpropertytype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.AbstractAIXMPropertyTypes, stage.AixmabstractgmlobjecttypesStage.AbstractAIXMPropertyType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			abstractaixmpropertytypeIdent := abstractaixmpropertytype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMPropertyType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", abstractaixmpropertytypeIdent, __gong__toRawStringLiteral(abstractaixmpropertytype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmpropertytypeIdent, __gong__toRawStringLiteral(abstractaixmpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", abstractaixmpropertytypeIdent, __gong__toRawStringLiteral(abstractaixmpropertytype.NilReason)))
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, abstractaixmtimeslicebasetype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceBaseTypes, stage.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceBaseType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			abstractaixmtimeslicebasetypeIdent := abstractaixmtimeslicebasetype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", abstractaixmtimeslicebasetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicebasetype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmtimeslicebasetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicebasetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmtimeslicebasetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicebasetype.Id)))
			if abstractaixmtimeslicebasetype.ValidTime != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmtimeslicebasetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", abstractaixmtimeslicebasetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, abstractaixmtimeslicetype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes, stage.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			abstractaixmtimeslicetypeIdent := abstractaixmtimeslicetype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", abstractaixmtimeslicetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicetype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmtimeslicetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmtimeslicetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractaixmtimeslicetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicetype.Description)))
			if abstractaixmtimeslicetype.Interpretation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = aixmabstractgmlobjecttypes.%s", abstractaixmtimeslicetypeIdent, abstractaixmtimeslicetype.Interpretation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = \"\"", abstractaixmtimeslicetypeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", abstractaixmtimeslicetypeIdent, abstractaixmtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", abstractaixmtimeslicetypeIdent, abstractaixmtimeslicetype.CorrectionNumber))
			if abstractaixmtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractaixmtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range abstractaixmtimeslicetype.NameXSD {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractaixmtimeslicetypeIdent, abstractaixmtimeslicetypeIdent, targetIdent))
			}
			if abstractaixmtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", abstractaixmtimeslicetypeIdent, targetIdent))
			}
			if abstractaixmtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", abstractaixmtimeslicetypeIdent, targetIdent))
			}
			if abstractaixmtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractaixmtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", abstractaixmtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, abstractextensiontype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.AbstractExtensionTypes, stage.AixmabstractgmlobjecttypesStage.AbstractExtensionType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			abstractextensiontypeIdent := abstractextensiontype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractExtensionType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", abstractextensiontypeIdent, __gong__toRawStringLiteral(abstractextensiontype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractextensiontypeIdent, __gong__toRawStringLiteral(abstractextensiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractextensiontypeIdent, __gong__toRawStringLiteral(abstractextensiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractextensiontypeIdent, __gong__toRawStringLiteral(abstractextensiontype.Description)))
			if abstractextensiontype.Identifier != nil {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := abstractextensiontype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractextensiontypeIdent, targetIdent))
			}
			for _, elem := range abstractextensiontype.NameXSD {
				if lastStagePtr != "AixmabstractgmlobjecttypesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmabstractgmlobjecttypesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractextensiontypeIdent, abstractextensiontypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, featuremetadatapropertytype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.FeatureMetadataPropertyTypes, stage.AixmabstractgmlobjecttypesStage.FeatureMetadataPropertyType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			featuremetadatapropertytypeIdent := featuremetadatapropertytype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.FeatureMetadataPropertyType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", featuremetadatapropertytypeIdent, __gong__toRawStringLiteral(featuremetadatapropertytype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", featuremetadatapropertytypeIdent, __gong__toRawStringLiteral(featuremetadatapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", featuremetadatapropertytypeIdent, featuremetadatapropertytype.Owns))
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, featuretimeslicemetadatapropertytype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.FeatureTimeSliceMetadataPropertyTypes, stage.AixmabstractgmlobjecttypesStage.FeatureTimeSliceMetadataPropertyType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			featuretimeslicemetadatapropertytypeIdent := featuretimeslicemetadatapropertytype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", featuretimeslicemetadatapropertytypeIdent, __gong__toRawStringLiteral(featuretimeslicemetadatapropertytype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", featuretimeslicemetadatapropertytypeIdent, __gong__toRawStringLiteral(featuretimeslicemetadatapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", featuretimeslicemetadatapropertytypeIdent, featuretimeslicemetadatapropertytype.Owns))
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, messagemetadatapropertytype := range __gong__sortStageSetInstances(stage.AixmabstractgmlobjecttypesStage.MessageMetadataPropertyTypes, stage.AixmabstractgmlobjecttypesStage.MessageMetadataPropertyType_stagedOrder) {
			if lastStageDecl != "AixmabstractgmlobjecttypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmabstractgmlobjecttypesStage"
			}
			messagemetadatapropertytypeIdent := messagemetadatapropertytype.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.MessageMetadataPropertyType{Name: %s}).Stage(stage.AixmabstractgmlobjecttypesStage)", messagemetadatapropertytypeIdent, __gong__toRawStringLiteral(messagemetadatapropertytype.Name)))
			if lastStageVal != "AixmabstractgmlobjecttypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmabstractgmlobjecttypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", messagemetadatapropertytypeIdent, __gong__toRawStringLiteral(messagemetadatapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", messagemetadatapropertytypeIdent, messagemetadatapropertytype.Owns))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeaircraftcategorytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAircraftCategoryTypes, stage.AixmdatatypesStage.CodeAircraftCategoryType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaircraftcategorytypeIdent := codeaircraftcategorytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftCategoryType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaircraftcategorytypeIdent, __gong__toRawStringLiteral(codeaircraftcategorytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftcategorytypeIdent, __gong__toRawStringLiteral(codeaircraftcategorytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftcategorytypeIdent, __gong__toRawStringLiteral(codeaircraftcategorytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftcategorytypeIdent, __gong__toRawStringLiteral(codeaircraftcategorytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeaircraftenginenumbertype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAircraftEngineNumberTypes, stage.AixmdatatypesStage.CodeAircraftEngineNumberType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaircraftenginenumbertypeIdent := codeaircraftenginenumbertype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftEngineNumberType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaircraftenginenumbertypeIdent, __gong__toRawStringLiteral(codeaircraftenginenumbertype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftenginenumbertypeIdent, __gong__toRawStringLiteral(codeaircraftenginenumbertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftenginenumbertypeIdent, __gong__toRawStringLiteral(codeaircraftenginenumbertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftenginenumbertypeIdent, __gong__toRawStringLiteral(codeaircraftenginenumbertype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeaircraftenginetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAircraftEngineTypes, stage.AixmdatatypesStage.CodeAircraftEngineType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaircraftenginetypeIdent := codeaircraftenginetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftEngineType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaircraftenginetypeIdent, __gong__toRawStringLiteral(codeaircraftenginetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftenginetypeIdent, __gong__toRawStringLiteral(codeaircraftenginetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftenginetypeIdent, __gong__toRawStringLiteral(codeaircraftenginetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftenginetypeIdent, __gong__toRawStringLiteral(codeaircraftenginetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeaircrafticaotype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAircraftICAOTypes, stage.AixmdatatypesStage.CodeAircraftICAOType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaircrafticaotypeIdent := codeaircrafticaotype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftICAOType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaircrafticaotypeIdent, __gong__toRawStringLiteral(codeaircrafticaotype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircrafticaotypeIdent, __gong__toRawStringLiteral(codeaircrafticaotype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircrafticaotypeIdent, __gong__toRawStringLiteral(codeaircrafticaotype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircrafticaotypeIdent, __gong__toRawStringLiteral(codeaircrafticaotype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeaircrafttype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAircraftTypes, stage.AixmdatatypesStage.CodeAircraftType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaircrafttypeIdent := codeaircrafttype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaircrafttypeIdent, __gong__toRawStringLiteral(codeaircrafttype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircrafttypeIdent, __gong__toRawStringLiteral(codeaircrafttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircrafttypeIdent, __gong__toRawStringLiteral(codeaircrafttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircrafttypeIdent, __gong__toRawStringLiteral(codeaircrafttype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeaircraftwingspanclasstype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAircraftWingspanClassTypes, stage.AixmdatatypesStage.CodeAircraftWingspanClassType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaircraftwingspanclasstypeIdent := codeaircraftwingspanclasstype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftWingspanClassType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaircraftwingspanclasstypeIdent, __gong__toRawStringLiteral(codeaircraftwingspanclasstype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftwingspanclasstypeIdent, __gong__toRawStringLiteral(codeaircraftwingspanclasstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftwingspanclasstypeIdent, __gong__toRawStringLiteral(codeaircraftwingspanclasstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftwingspanclasstypeIdent, __gong__toRawStringLiteral(codeaircraftwingspanclasstype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeairspaceactivitytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAirspaceActivityTypes, stage.AixmdatatypesStage.CodeAirspaceActivityType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeairspaceactivitytypeIdent := codeairspaceactivitytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceActivityType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeairspaceactivitytypeIdent, __gong__toRawStringLiteral(codeairspaceactivitytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspaceactivitytypeIdent, __gong__toRawStringLiteral(codeairspaceactivitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspaceactivitytypeIdent, __gong__toRawStringLiteral(codeairspaceactivitytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspaceactivitytypeIdent, __gong__toRawStringLiteral(codeairspaceactivitytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeairspaceaggregationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAirspaceAggregationTypes, stage.AixmdatatypesStage.CodeAirspaceAggregationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeairspaceaggregationtypeIdent := codeairspaceaggregationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceAggregationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeairspaceaggregationtypeIdent, __gong__toRawStringLiteral(codeairspaceaggregationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspaceaggregationtypeIdent, __gong__toRawStringLiteral(codeairspaceaggregationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspaceaggregationtypeIdent, __gong__toRawStringLiteral(codeairspaceaggregationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspaceaggregationtypeIdent, __gong__toRawStringLiteral(codeairspaceaggregationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeairspaceclassificationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAirspaceClassificationTypes, stage.AixmdatatypesStage.CodeAirspaceClassificationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeairspaceclassificationtypeIdent := codeairspaceclassificationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceClassificationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeairspaceclassificationtypeIdent, __gong__toRawStringLiteral(codeairspaceclassificationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspaceclassificationtypeIdent, __gong__toRawStringLiteral(codeairspaceclassificationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspaceclassificationtypeIdent, __gong__toRawStringLiteral(codeairspaceclassificationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspaceclassificationtypeIdent, __gong__toRawStringLiteral(codeairspaceclassificationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeairspacedependencytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAirspaceDependencyTypes, stage.AixmdatatypesStage.CodeAirspaceDependencyType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeairspacedependencytypeIdent := codeairspacedependencytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceDependencyType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeairspacedependencytypeIdent, __gong__toRawStringLiteral(codeairspacedependencytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspacedependencytypeIdent, __gong__toRawStringLiteral(codeairspacedependencytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspacedependencytypeIdent, __gong__toRawStringLiteral(codeairspacedependencytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspacedependencytypeIdent, __gong__toRawStringLiteral(codeairspacedependencytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeairspacedesignatortype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAirspaceDesignatorTypes, stage.AixmdatatypesStage.CodeAirspaceDesignatorType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeairspacedesignatortypeIdent := codeairspacedesignatortype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceDesignatorType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeairspacedesignatortypeIdent, __gong__toRawStringLiteral(codeairspacedesignatortype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspacedesignatortypeIdent, __gong__toRawStringLiteral(codeairspacedesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspacedesignatortypeIdent, __gong__toRawStringLiteral(codeairspacedesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspacedesignatortypeIdent, __gong__toRawStringLiteral(codeairspacedesignatortype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeairspacetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAirspaceTypes, stage.AixmdatatypesStage.CodeAirspaceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeairspacetypeIdent := codeairspacetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspaceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeairspacetypeIdent, __gong__toRawStringLiteral(codeairspacetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspacetypeIdent, __gong__toRawStringLiteral(codeairspacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspacetypeIdent, __gong__toRawStringLiteral(codeairspacetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspacetypeIdent, __gong__toRawStringLiteral(codeairspacetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codealtitudeusetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAltitudeUseTypes, stage.AixmdatatypesStage.CodeAltitudeUseType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codealtitudeusetypeIdent := codealtitudeusetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAltitudeUseType{Name: %s}).Stage(stage.AixmdatatypesStage)", codealtitudeusetypeIdent, __gong__toRawStringLiteral(codealtitudeusetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codealtitudeusetypeIdent, __gong__toRawStringLiteral(codealtitudeusetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codealtitudeusetypeIdent, __gong__toRawStringLiteral(codealtitudeusetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codealtitudeusetypeIdent, __gong__toRawStringLiteral(codealtitudeusetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeauthoritytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAuthorityTypes, stage.AixmdatatypesStage.CodeAuthorityType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeauthoritytypeIdent := codeauthoritytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAuthorityType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeauthoritytypeIdent, __gong__toRawStringLiteral(codeauthoritytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeauthoritytypeIdent, __gong__toRawStringLiteral(codeauthoritytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeauthoritytypeIdent, __gong__toRawStringLiteral(codeauthoritytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeauthoritytypeIdent, __gong__toRawStringLiteral(codeauthoritytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codecommunicationmodetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeCommunicationModeTypes, stage.AixmdatatypesStage.CodeCommunicationModeType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codecommunicationmodetypeIdent := codecommunicationmodetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCommunicationModeType{Name: %s}).Stage(stage.AixmdatatypesStage)", codecommunicationmodetypeIdent, __gong__toRawStringLiteral(codecommunicationmodetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecommunicationmodetypeIdent, __gong__toRawStringLiteral(codecommunicationmodetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecommunicationmodetypeIdent, __gong__toRawStringLiteral(codecommunicationmodetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecommunicationmodetypeIdent, __gong__toRawStringLiteral(codecommunicationmodetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codedaytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeDayTypes, stage.AixmdatatypesStage.CodeDayType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codedaytypeIdent := codedaytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDayType{Name: %s}).Stage(stage.AixmdatatypesStage)", codedaytypeIdent, __gong__toRawStringLiteral(codedaytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedaytypeIdent, __gong__toRawStringLiteral(codedaytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedaytypeIdent, __gong__toRawStringLiteral(codedaytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedaytypeIdent, __gong__toRawStringLiteral(codedaytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeequipmentanticollisiontype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeEquipmentAntiCollisionTypes, stage.AixmdatatypesStage.CodeEquipmentAntiCollisionType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeequipmentanticollisiontypeIdent := codeequipmentanticollisiontype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeEquipmentAntiCollisionType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeequipmentanticollisiontypeIdent, __gong__toRawStringLiteral(codeequipmentanticollisiontype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeequipmentanticollisiontypeIdent, __gong__toRawStringLiteral(codeequipmentanticollisiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeequipmentanticollisiontypeIdent, __gong__toRawStringLiteral(codeequipmentanticollisiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeequipmentanticollisiontypeIdent, __gong__toRawStringLiteral(codeequipmentanticollisiontype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codegeobordertype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeGeoBorderTypes, stage.AixmdatatypesStage.CodeGeoBorderType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codegeobordertypeIdent := codegeobordertype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeGeoBorderType{Name: %s}).Stage(stage.AixmdatatypesStage)", codegeobordertypeIdent, __gong__toRawStringLiteral(codegeobordertype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codegeobordertypeIdent, __gong__toRawStringLiteral(codegeobordertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codegeobordertypeIdent, __gong__toRawStringLiteral(codegeobordertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codegeobordertypeIdent, __gong__toRawStringLiteral(codegeobordertype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codemilitaryoperationstype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeMilitaryOperationsTypes, stage.AixmdatatypesStage.CodeMilitaryOperationsType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codemilitaryoperationstypeIdent := codemilitaryoperationstype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMilitaryOperationsType{Name: %s}).Stage(stage.AixmdatatypesStage)", codemilitaryoperationstypeIdent, __gong__toRawStringLiteral(codemilitaryoperationstype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemilitaryoperationstypeIdent, __gong__toRawStringLiteral(codemilitaryoperationstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemilitaryoperationstypeIdent, __gong__toRawStringLiteral(codemilitaryoperationstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemilitaryoperationstypeIdent, __gong__toRawStringLiteral(codemilitaryoperationstype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codenavigationequipmenttype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeNavigationEquipmentTypes, stage.AixmdatatypesStage.CodeNavigationEquipmentType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codenavigationequipmenttypeIdent := codenavigationequipmenttype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavigationEquipmentType{Name: %s}).Stage(stage.AixmdatatypesStage)", codenavigationequipmenttypeIdent, __gong__toRawStringLiteral(codenavigationequipmenttype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavigationequipmenttypeIdent, __gong__toRawStringLiteral(codenavigationequipmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavigationequipmenttypeIdent, __gong__toRawStringLiteral(codenavigationequipmenttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavigationequipmenttypeIdent, __gong__toRawStringLiteral(codenavigationequipmenttype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codenavigationspecificationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeNavigationSpecificationTypes, stage.AixmdatatypesStage.CodeNavigationSpecificationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codenavigationspecificationtypeIdent := codenavigationspecificationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavigationSpecificationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codenavigationspecificationtypeIdent, __gong__toRawStringLiteral(codenavigationspecificationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavigationspecificationtypeIdent, __gong__toRawStringLiteral(codenavigationspecificationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavigationspecificationtypeIdent, __gong__toRawStringLiteral(codenavigationspecificationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavigationspecificationtypeIdent, __gong__toRawStringLiteral(codenavigationspecificationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codenotepurposetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeNotePurposeTypes, stage.AixmdatatypesStage.CodeNotePurposeType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codenotepurposetypeIdent := codenotepurposetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNotePurposeType{Name: %s}).Stage(stage.AixmdatatypesStage)", codenotepurposetypeIdent, __gong__toRawStringLiteral(codenotepurposetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenotepurposetypeIdent, __gong__toRawStringLiteral(codenotepurposetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenotepurposetypeIdent, __gong__toRawStringLiteral(codenotepurposetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenotepurposetypeIdent, __gong__toRawStringLiteral(codenotepurposetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codervsmtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRVSMTypes, stage.AixmdatatypesStage.CodeRVSMType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codervsmtypeIdent := codervsmtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRVSMType{Name: %s}).Stage(stage.AixmdatatypesStage)", codervsmtypeIdent, __gong__toRawStringLiteral(codervsmtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codervsmtypeIdent, __gong__toRawStringLiteral(codervsmtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codervsmtypeIdent, __gong__toRawStringLiteral(codervsmtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codervsmtypeIdent, __gong__toRawStringLiteral(codervsmtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codestatusairspacetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeStatusAirspaceTypes, stage.AixmdatatypesStage.CodeStatusAirspaceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codestatusairspacetypeIdent := codestatusairspacetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusAirspaceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codestatusairspacetypeIdent, __gong__toRawStringLiteral(codestatusairspacetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusairspacetypeIdent, __gong__toRawStringLiteral(codestatusairspacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusairspacetypeIdent, __gong__toRawStringLiteral(codestatusairspacetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusairspacetypeIdent, __gong__toRawStringLiteral(codestatusairspacetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codetimeeventcombinationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeTimeEventCombinationTypes, stage.AixmdatatypesStage.CodeTimeEventCombinationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codetimeeventcombinationtypeIdent := codetimeeventcombinationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTimeEventCombinationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codetimeeventcombinationtypeIdent, __gong__toRawStringLiteral(codetimeeventcombinationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetimeeventcombinationtypeIdent, __gong__toRawStringLiteral(codetimeeventcombinationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetimeeventcombinationtypeIdent, __gong__toRawStringLiteral(codetimeeventcombinationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetimeeventcombinationtypeIdent, __gong__toRawStringLiteral(codetimeeventcombinationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codetimeeventtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeTimeEventTypes, stage.AixmdatatypesStage.CodeTimeEventType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codetimeeventtypeIdent := codetimeeventtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTimeEventType{Name: %s}).Stage(stage.AixmdatatypesStage)", codetimeeventtypeIdent, __gong__toRawStringLiteral(codetimeeventtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetimeeventtypeIdent, __gong__toRawStringLiteral(codetimeeventtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetimeeventtypeIdent, __gong__toRawStringLiteral(codetimeeventtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetimeeventtypeIdent, __gong__toRawStringLiteral(codetimeeventtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codetimereferencetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeTimeReferenceTypes, stage.AixmdatatypesStage.CodeTimeReferenceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codetimereferencetypeIdent := codetimereferencetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTimeReferenceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codetimereferencetypeIdent, __gong__toRawStringLiteral(codetimereferencetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetimereferencetypeIdent, __gong__toRawStringLiteral(codetimereferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetimereferencetypeIdent, __gong__toRawStringLiteral(codetimereferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetimereferencetypeIdent, __gong__toRawStringLiteral(codetimereferencetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codetranspondertype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeTransponderTypes, stage.AixmdatatypesStage.CodeTransponderType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codetranspondertypeIdent := codetranspondertype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTransponderType{Name: %s}).Stage(stage.AixmdatatypesStage)", codetranspondertypeIdent, __gong__toRawStringLiteral(codetranspondertype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetranspondertypeIdent, __gong__toRawStringLiteral(codetranspondertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetranspondertypeIdent, __gong__toRawStringLiteral(codetranspondertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetranspondertypeIdent, __gong__toRawStringLiteral(codetranspondertype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codevalueinterpretationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeValueInterpretationTypes, stage.AixmdatatypesStage.CodeValueInterpretationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codevalueinterpretationtypeIdent := codevalueinterpretationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeValueInterpretationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codevalueinterpretationtypeIdent, __gong__toRawStringLiteral(codevalueinterpretationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codevalueinterpretationtypeIdent, __gong__toRawStringLiteral(codevalueinterpretationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codevalueinterpretationtypeIdent, __gong__toRawStringLiteral(codevalueinterpretationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codevalueinterpretationtypeIdent, __gong__toRawStringLiteral(codevalueinterpretationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeverticaldatumtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeVerticalDatumTypes, stage.AixmdatatypesStage.CodeVerticalDatumType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeverticaldatumtypeIdent := codeverticaldatumtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVerticalDatumType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeverticaldatumtypeIdent, __gong__toRawStringLiteral(codeverticaldatumtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeverticaldatumtypeIdent, __gong__toRawStringLiteral(codeverticaldatumtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeverticaldatumtypeIdent, __gong__toRawStringLiteral(codeverticaldatumtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeverticaldatumtypeIdent, __gong__toRawStringLiteral(codeverticaldatumtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeverticalreferencetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeVerticalReferenceTypes, stage.AixmdatatypesStage.CodeVerticalReferenceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeverticalreferencetypeIdent := codeverticalreferencetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVerticalReferenceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeverticalreferencetypeIdent, __gong__toRawStringLiteral(codeverticalreferencetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeverticalreferencetypeIdent, __gong__toRawStringLiteral(codeverticalreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeverticalreferencetypeIdent, __gong__toRawStringLiteral(codeverticalreferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeverticalreferencetypeIdent, __gong__toRawStringLiteral(codeverticalreferencetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codewaketurbulencetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeWakeTurbulenceTypes, stage.AixmdatatypesStage.CodeWakeTurbulenceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codewaketurbulencetypeIdent := codewaketurbulencetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeWakeTurbulenceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codewaketurbulencetypeIdent, __gong__toRawStringLiteral(codewaketurbulencetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codewaketurbulencetypeIdent, __gong__toRawStringLiteral(codewaketurbulencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codewaketurbulencetypeIdent, __gong__toRawStringLiteral(codewaketurbulencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codewaketurbulencetypeIdent, __gong__toRawStringLiteral(codewaketurbulencetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeyesnotype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeYesNoTypes, stage.AixmdatatypesStage.CodeYesNoType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeyesnotypeIdent := codeyesnotype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeYesNoType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeyesnotypeIdent, __gong__toRawStringLiteral(codeyesnotype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeyesnotypeIdent, __gong__toRawStringLiteral(codeyesnotype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeyesnotypeIdent, __gong__toRawStringLiteral(codeyesnotype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeyesnotypeIdent, __gong__toRawStringLiteral(codeyesnotype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, datemonthdaytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.DateMonthDayTypes, stage.AixmdatatypesStage.DateMonthDayType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			datemonthdaytypeIdent := datemonthdaytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.DateMonthDayType{Name: %s}).Stage(stage.AixmdatatypesStage)", datemonthdaytypeIdent, __gong__toRawStringLiteral(datemonthdaytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datemonthdaytypeIdent, __gong__toRawStringLiteral(datemonthdaytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", datemonthdaytypeIdent, __gong__toRawStringLiteral(datemonthdaytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", datemonthdaytypeIdent, __gong__toRawStringLiteral(datemonthdaytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, nonumbertype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.NoNumberTypes, stage.AixmdatatypesStage.NoNumberType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			nonumbertypeIdent := nonumbertype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.NoNumberType{Name: %s}).Stage(stage.AixmdatatypesStage)", nonumbertypeIdent, __gong__toRawStringLiteral(nonumbertype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nonumbertypeIdent, __gong__toRawStringLiteral(nonumbertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", nonumbertypeIdent, __gong__toRawStringLiteral(nonumbertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", nonumbertypeIdent, nonumbertype.EnclosedText))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, nosequencetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.NoSequenceTypes, stage.AixmdatatypesStage.NoSequenceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			nosequencetypeIdent := nosequencetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.NoSequenceType{Name: %s}).Stage(stage.AixmdatatypesStage)", nosequencetypeIdent, __gong__toRawStringLiteral(nosequencetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nosequencetypeIdent, __gong__toRawStringLiteral(nosequencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", nosequencetypeIdent, __gong__toRawStringLiteral(nosequencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", nosequencetypeIdent, nosequencetype.EnclosedText))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, textnametype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.TextNameTypes, stage.AixmdatatypesStage.TextNameType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			textnametypeIdent := textnametype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextNameType{Name: %s}).Stage(stage.AixmdatatypesStage)", textnametypeIdent, __gong__toRawStringLiteral(textnametype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textnametypeIdent, __gong__toRawStringLiteral(textnametype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textnametypeIdent, __gong__toRawStringLiteral(textnametype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textnametypeIdent, __gong__toRawStringLiteral(textnametype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, textnotetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.TextNoteTypes, stage.AixmdatatypesStage.TextNoteType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			textnotetypeIdent := textnotetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextNoteType{Name: %s}).Stage(stage.AixmdatatypesStage)", textnotetypeIdent, __gong__toRawStringLiteral(textnotetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textnotetypeIdent, __gong__toRawStringLiteral(textnotetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textnotetypeIdent, __gong__toRawStringLiteral(textnotetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", textnotetypeIdent, __gong__toRawStringLiteral(textnotetype.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textnotetypeIdent, __gong__toRawStringLiteral(textnotetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, textpropertynametype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.TextPropertyNameTypes, stage.AixmdatatypesStage.TextPropertyNameType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			textpropertynametypeIdent := textpropertynametype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextPropertyNameType{Name: %s}).Stage(stage.AixmdatatypesStage)", textpropertynametypeIdent, __gong__toRawStringLiteral(textpropertynametype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textpropertynametypeIdent, __gong__toRawStringLiteral(textpropertynametype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textpropertynametypeIdent, __gong__toRawStringLiteral(textpropertynametype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textpropertynametypeIdent, __gong__toRawStringLiteral(textpropertynametype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, timetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.TimeTypes, stage.AixmdatatypesStage.TimeType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			timetypeIdent := timetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TimeType{Name: %s}).Stage(stage.AixmdatatypesStage)", timetypeIdent, __gong__toRawStringLiteral(timetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timetypeIdent, __gong__toRawStringLiteral(timetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", timetypeIdent, __gong__toRawStringLiteral(timetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", timetypeIdent, __gong__toRawStringLiteral(timetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, valdistancesignedtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValDistanceSignedTypes, stage.AixmdatatypesStage.ValDistanceSignedType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valdistancesignedtypeIdent := valdistancesignedtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValDistanceSignedType{Name: %s}).Stage(stage.AixmdatatypesStage)", valdistancesignedtypeIdent, __gong__toRawStringLiteral(valdistancesignedtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valdistancesignedtypeIdent, __gong__toRawStringLiteral(valdistancesignedtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valdistancesignedtypeIdent, __gong__toRawStringLiteral(valdistancesignedtype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valdistancesignedtypeIdent, __gong__toRawStringLiteral(valdistancesignedtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valdistancesignedtypeIdent, __gong__toRawStringLiteral(valdistancesignedtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, valdistancetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValDistanceTypes, stage.AixmdatatypesStage.ValDistanceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valdistancetypeIdent := valdistancetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValDistanceType{Name: %s}).Stage(stage.AixmdatatypesStage)", valdistancetypeIdent, __gong__toRawStringLiteral(valdistancetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valdistancetypeIdent, __gong__toRawStringLiteral(valdistancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valdistancetypeIdent, __gong__toRawStringLiteral(valdistancetype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valdistancetypeIdent, __gong__toRawStringLiteral(valdistancetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valdistancetypeIdent, __gong__toRawStringLiteral(valdistancetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, valdistanceverticaltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValDistanceVerticalTypes, stage.AixmdatatypesStage.ValDistanceVerticalType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valdistanceverticaltypeIdent := valdistanceverticaltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValDistanceVerticalType{Name: %s}).Stage(stage.AixmdatatypesStage)", valdistanceverticaltypeIdent, __gong__toRawStringLiteral(valdistanceverticaltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valdistanceverticaltypeIdent, __gong__toRawStringLiteral(valdistanceverticaltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valdistanceverticaltypeIdent, __gong__toRawStringLiteral(valdistanceverticaltype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valdistanceverticaltypeIdent, __gong__toRawStringLiteral(valdistanceverticaltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valdistanceverticaltypeIdent, __gong__toRawStringLiteral(valdistanceverticaltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, valdurationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValDurationTypes, stage.AixmdatatypesStage.ValDurationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valdurationtypeIdent := valdurationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValDurationType{Name: %s}).Stage(stage.AixmdatatypesStage)", valdurationtypeIdent, __gong__toRawStringLiteral(valdurationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valdurationtypeIdent, __gong__toRawStringLiteral(valdurationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valdurationtypeIdent, __gong__toRawStringLiteral(valdurationtype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valdurationtypeIdent, __gong__toRawStringLiteral(valdurationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valdurationtypeIdent, __gong__toRawStringLiteral(valdurationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, valfltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValFLTypes, stage.AixmdatatypesStage.ValFLType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valfltypeIdent := valfltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValFLType{Name: %s}).Stage(stage.AixmdatatypesStage)", valfltypeIdent, __gong__toRawStringLiteral(valfltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valfltypeIdent, __gong__toRawStringLiteral(valfltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valfltypeIdent, __gong__toRawStringLiteral(valfltype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valfltypeIdent, __gong__toRawStringLiteral(valfltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", valfltypeIdent, valfltype.EnclosedText))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, valspeedtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValSpeedTypes, stage.AixmdatatypesStage.ValSpeedType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valspeedtypeIdent := valspeedtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValSpeedType{Name: %s}).Stage(stage.AixmdatatypesStage)", valspeedtypeIdent, __gong__toRawStringLiteral(valspeedtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valspeedtypeIdent, __gong__toRawStringLiteral(valspeedtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valspeedtypeIdent, __gong__toRawStringLiteral(valspeedtype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valspeedtypeIdent, __gong__toRawStringLiteral(valspeedtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valspeedtypeIdent, __gong__toRawStringLiteral(valspeedtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, valweighttype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValWeightTypes, stage.AixmdatatypesStage.ValWeightType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valweighttypeIdent := valweighttype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValWeightType{Name: %s}).Stage(stage.AixmdatatypesStage)", valweighttypeIdent, __gong__toRawStringLiteral(valweighttype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valweighttypeIdent, __gong__toRawStringLiteral(valweighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valweighttypeIdent, __gong__toRawStringLiteral(valweighttype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valweighttypeIdent, __gong__toRawStringLiteral(valweighttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valweighttypeIdent, __gong__toRawStringLiteral(valweighttype.EnclosedText)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extensions, stage.AixmfeaturesStage.A_extension_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extensionIdent := a_extension.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extensionIdent, __gong__toRawStringLiteral(a_extension.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extensionIdent, __gong__toRawStringLiteral(a_extension.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extensionIdent, a_extension.Owns))
			for _, elem := range a_extension.AbstractAirspaceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirspaceExtension = append(%s.AbstractAirspaceExtension, %s)", a_extensionIdent, a_extensionIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_1 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_1s, stage.AixmfeaturesStage.A_extension_1_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_1Ident := a_extension_1.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_1{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_1Ident, __gong__toRawStringLiteral(a_extension_1.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_1Ident, __gong__toRawStringLiteral(a_extension_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_1Ident, a_extension_1.Owns))
			for _, elem := range a_extension_1.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_1Ident, a_extension_1Ident, targetIdent))
			}
			for _, elem := range a_extension_1.AbstractAirspaceActivationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirspaceActivationExtension = append(%s.AbstractAirspaceActivationExtension, %s)", a_extension_1Ident, a_extension_1Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_10 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_10s, stage.AixmfeaturesStage.A_extension_10_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_10Ident := a_extension_10.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_10{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_10Ident, __gong__toRawStringLiteral(a_extension_10.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_10Ident, __gong__toRawStringLiteral(a_extension_10.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_10Ident, a_extension_10.Owns))
			for _, elem := range a_extension_10.AbstractLinguisticNoteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractLinguisticNoteExtension = append(%s.AbstractLinguisticNoteExtension, %s)", a_extension_10Ident, a_extension_10Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_11 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_11s, stage.AixmfeaturesStage.A_extension_11_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_11Ident := a_extension_11.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_11{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_11Ident, __gong__toRawStringLiteral(a_extension_11.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_11Ident, __gong__toRawStringLiteral(a_extension_11.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_11Ident, a_extension_11.Owns))
			for _, elem := range a_extension_11.AbstractNoteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNoteExtension = append(%s.AbstractNoteExtension, %s)", a_extension_11Ident, a_extension_11Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_12 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_12s, stage.AixmfeaturesStage.A_extension_12_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_12Ident := a_extension_12.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_12{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_12Ident, __gong__toRawStringLiteral(a_extension_12.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_12Ident, __gong__toRawStringLiteral(a_extension_12.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_12Ident, a_extension_12.Owns))
			for _, elem := range a_extension_12.AbstractAircraftCharacteristicExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAircraftCharacteristicExtension = append(%s.AbstractAircraftCharacteristicExtension, %s)", a_extension_12Ident, a_extension_12Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_13 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_13s, stage.AixmfeaturesStage.A_extension_13_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_13Ident := a_extension_13.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_13{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_13Ident, __gong__toRawStringLiteral(a_extension_13.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_13Ident, __gong__toRawStringLiteral(a_extension_13.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_13Ident, a_extension_13.Owns))
			for _, elem := range a_extension_13.AbstractAirspaceLayerExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirspaceLayerExtension = append(%s.AbstractAirspaceLayerExtension, %s)", a_extension_13Ident, a_extension_13Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_14 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_14s, stage.AixmfeaturesStage.A_extension_14_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_14Ident := a_extension_14.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_14{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_14Ident, __gong__toRawStringLiteral(a_extension_14.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_14Ident, __gong__toRawStringLiteral(a_extension_14.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_14Ident, a_extension_14.Owns))
			for _, elem := range a_extension_14.AbstractTimesheetExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractTimesheetExtension = append(%s.AbstractTimesheetExtension, %s)", a_extension_14Ident, a_extension_14Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_2 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_2s, stage.AixmfeaturesStage.A_extension_2_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_2Ident := a_extension_2.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_2{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_2Ident, __gong__toRawStringLiteral(a_extension_2.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_2Ident, __gong__toRawStringLiteral(a_extension_2.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_2Ident, a_extension_2.Owns))
			for _, elem := range a_extension_2.AbstractAirspaceGeometryComponentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirspaceGeometryComponentExtension = append(%s.AbstractAirspaceGeometryComponentExtension, %s)", a_extension_2Ident, a_extension_2Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_3 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_3s, stage.AixmfeaturesStage.A_extension_3_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_3Ident := a_extension_3.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_3{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_3Ident, __gong__toRawStringLiteral(a_extension_3.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_3Ident, __gong__toRawStringLiteral(a_extension_3.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_3Ident, a_extension_3.Owns))
			for _, elem := range a_extension_3.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_3Ident, a_extension_3Ident, targetIdent))
			}
			for _, elem := range a_extension_3.AbstractAirspaceLayerClassExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirspaceLayerClassExtension = append(%s.AbstractAirspaceLayerClassExtension, %s)", a_extension_3Ident, a_extension_3Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_4 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_4s, stage.AixmfeaturesStage.A_extension_4_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_4Ident := a_extension_4.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_4{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_4Ident, __gong__toRawStringLiteral(a_extension_4.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_4Ident, __gong__toRawStringLiteral(a_extension_4.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_4Ident, a_extension_4.Owns))
			for _, elem := range a_extension_4.AbstractAirspaceVolumeExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirspaceVolumeExtension = append(%s.AbstractAirspaceVolumeExtension, %s)", a_extension_4Ident, a_extension_4Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_5 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_5s, stage.AixmfeaturesStage.A_extension_5_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_5Ident := a_extension_5.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_5{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_5Ident, __gong__toRawStringLiteral(a_extension_5.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_5Ident, __gong__toRawStringLiteral(a_extension_5.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_5Ident, a_extension_5.Owns))
			for _, elem := range a_extension_5.AbstractAirspaceVolumeDependencyExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirspaceVolumeDependencyExtension = append(%s.AbstractAirspaceVolumeDependencyExtension, %s)", a_extension_5Ident, a_extension_5Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_6 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_6s, stage.AixmfeaturesStage.A_extension_6_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_6Ident := a_extension_6.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_6{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_6Ident, __gong__toRawStringLiteral(a_extension_6.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_6Ident, __gong__toRawStringLiteral(a_extension_6.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_6Ident, a_extension_6.Owns))
			for _, elem := range a_extension_6.AbstractAuthorityForAirspaceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAuthorityForAirspaceExtension = append(%s.AbstractAuthorityForAirspaceExtension, %s)", a_extension_6Ident, a_extension_6Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_7 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_7s, stage.AixmfeaturesStage.A_extension_7_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_7Ident := a_extension_7.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_7{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_7Ident, __gong__toRawStringLiteral(a_extension_7.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_7Ident, __gong__toRawStringLiteral(a_extension_7.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_7Ident, a_extension_7.Owns))
			for _, elem := range a_extension_7.AbstractGeoBorderExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractGeoBorderExtension = append(%s.AbstractGeoBorderExtension, %s)", a_extension_7Ident, a_extension_7Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_8 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_8s, stage.AixmfeaturesStage.A_extension_8_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_8Ident := a_extension_8.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_8{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_8Ident, __gong__toRawStringLiteral(a_extension_8.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_8Ident, __gong__toRawStringLiteral(a_extension_8.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_8Ident, a_extension_8.Owns))
			for _, elem := range a_extension_8.AbstractElevatedCurveExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractElevatedCurveExtension = append(%s.AbstractElevatedCurveExtension, %s)", a_extension_8Ident, a_extension_8Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_9 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_9s, stage.AixmfeaturesStage.A_extension_9_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_9Ident := a_extension_9.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_9{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_9Ident, __gong__toRawStringLiteral(a_extension_9.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_9Ident, __gong__toRawStringLiteral(a_extension_9.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_9Ident, a_extension_9.Owns))
			for _, elem := range a_extension_9.AbstractElevatedSurfaceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractElevatedSurfaceExtension = append(%s.AbstractElevatedSurfaceExtension, %s)", a_extension_9Ident, a_extension_9Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, abstractpropertieswithscheduletype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AbstractPropertiesWithScheduleTypes, stage.AixmfeaturesStage.AbstractPropertiesWithScheduleType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			abstractpropertieswithscheduletypeIdent := abstractpropertieswithscheduletype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractPropertiesWithScheduleType{Name: %s}).Stage(stage.AixmfeaturesStage)", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Description)))
			if abstractpropertieswithscheduletype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractpropertieswithscheduletype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractpropertieswithscheduletypeIdent, targetIdent))
			}
			for _, elem := range abstractpropertieswithscheduletype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractpropertieswithscheduletypeIdent, abstractpropertieswithscheduletypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aircraftcharacteristicpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AircraftCharacteristicPropertyTypes, stage.AixmfeaturesStage.AircraftCharacteristicPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aircraftcharacteristicpropertytypeIdent := aircraftcharacteristicpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftCharacteristicPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", aircraftcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(aircraftcharacteristicpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(aircraftcharacteristicpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", aircraftcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(aircraftcharacteristicpropertytype.NilReason)))
			if aircraftcharacteristicpropertytype.AircraftCharacteristic != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristicpropertytype.AircraftCharacteristic.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCharacteristic = %s", aircraftcharacteristicpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aircraftcharacteristictype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AircraftCharacteristicTypes, stage.AixmfeaturesStage.AircraftCharacteristicType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aircraftcharacteristictypeIdent := aircraftcharacteristictype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftCharacteristicType{Name: %s}).Stage(stage.AixmfeaturesStage)", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Description)))
			if aircraftcharacteristictype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			for _, elem := range aircraftcharacteristictype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aircraftcharacteristictypeIdent, aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.Engine != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.Engine.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Engine = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.NumberEngine != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.NumberEngine.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NumberEngine = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.TypeAircraftICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.TypeAircraftICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TypeAircraftICAO = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.AircraftLandingCategory != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.AircraftLandingCategory.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftLandingCategory = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.WingSpan != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.WingSpan.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WingSpan = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.WingSpanInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.WingSpanInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WingSpanInterpretation = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.ClassWingSpan != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.ClassWingSpan.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClassWingSpan = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.Weight != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.Weight.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Weight = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.WeightInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.WeightInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WeightInterpretation = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.Passengers != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.Passengers.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Passengers = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.PassengersInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.PassengersInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PassengersInterpretation = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.Speed != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.Speed.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Speed = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.SpeedInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.SpeedInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpeedInterpretation = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.WakeTurbulence != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.WakeTurbulence.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WakeTurbulence = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.NavigationEquipment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.NavigationEquipment.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavigationEquipment = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.NavigationSpecification != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.NavigationSpecification.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavigationSpecification = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.VerticalSeparationCapability != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.VerticalSeparationCapability.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalSeparationCapability = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.AntiCollisionAndSeparationEquipment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.AntiCollisionAndSeparationEquipment.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AntiCollisionAndSeparationEquipment = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.CommunicationEquipment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.CommunicationEquipment.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CommunicationEquipment = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			if aircraftcharacteristictype.SurveillanceEquipment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftcharacteristictype.SurveillanceEquipment.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SurveillanceEquipment = %s", aircraftcharacteristictypeIdent, targetIdent))
			}
			for _, elem := range aircraftcharacteristictype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aircraftcharacteristictypeIdent, aircraftcharacteristictypeIdent, targetIdent))
			}
			for _, elem := range aircraftcharacteristictype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aircraftcharacteristictypeIdent, aircraftcharacteristictypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspaceactivationpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceActivationPropertyTypes, stage.AixmfeaturesStage.AirspaceActivationPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspaceactivationpropertytypeIdent := airspaceactivationpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceActivationPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspaceactivationpropertytypeIdent, __gong__toRawStringLiteral(airspaceactivationpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspaceactivationpropertytypeIdent, __gong__toRawStringLiteral(airspaceactivationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspaceactivationpropertytypeIdent, __gong__toRawStringLiteral(airspaceactivationpropertytype.NilReason)))
			if airspaceactivationpropertytype.AirspaceActivation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspaceactivationpropertytype.AirspaceActivation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceActivation = %s", airspaceactivationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspaceactivationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceActivationTypes, stage.AixmfeaturesStage.AirspaceActivationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspaceactivationtypeIdent := airspaceactivationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceActivationType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspaceactivationtypeIdent, __gong__toRawStringLiteral(airspaceactivationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspaceactivationtypeIdent, __gong__toRawStringLiteral(airspaceactivationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspaceactivationtypeIdent, __gong__toRawStringLiteral(airspaceactivationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspaceactivationtypeIdent, __gong__toRawStringLiteral(airspaceactivationtype.Description)))
			if airspaceactivationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspaceactivationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			if airspaceactivationtype.Activity != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspaceactivationtype.Activity.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Activity = %s", airspaceactivationtypeIdent, targetIdent))
			}
			if airspaceactivationtype.Status != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspaceactivationtype.Status.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Status = %s", airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.Levels {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Levels = append(%s.Levels, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.User {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.User = append(%s.User, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.Aircraft {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Aircraft = append(%s.Aircraft, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacegeometrycomponentpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceGeometryComponentPropertyTypes, stage.AixmfeaturesStage.AirspaceGeometryComponentPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacegeometrycomponentpropertytypeIdent := airspacegeometrycomponentpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceGeometryComponentPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacegeometrycomponentpropertytypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponentpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacegeometrycomponentpropertytypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacegeometrycomponentpropertytypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponentpropertytype.NilReason)))
			if airspacegeometrycomponentpropertytype.AirspaceGeometryComponent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacegeometrycomponentpropertytype.AirspaceGeometryComponent.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceGeometryComponent = %s", airspacegeometrycomponentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacegeometrycomponenttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceGeometryComponentTypes, stage.AixmfeaturesStage.AirspaceGeometryComponentType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacegeometrycomponenttypeIdent := airspacegeometrycomponenttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceGeometryComponentType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacegeometrycomponenttypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponenttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacegeometrycomponenttypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacegeometrycomponenttypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacegeometrycomponenttypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponenttype.Description)))
			if airspacegeometrycomponenttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacegeometrycomponenttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airspacegeometrycomponenttypeIdent, targetIdent))
			}
			for _, elem := range airspacegeometrycomponenttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airspacegeometrycomponenttypeIdent, airspacegeometrycomponenttypeIdent, targetIdent))
			}
			if airspacegeometrycomponenttype.Operation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacegeometrycomponenttype.Operation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Operation = %s", airspacegeometrycomponenttypeIdent, targetIdent))
			}
			if airspacegeometrycomponenttype.OperationSequence != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacegeometrycomponenttype.OperationSequence.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OperationSequence = %s", airspacegeometrycomponenttypeIdent, targetIdent))
			}
			for _, elem := range airspacegeometrycomponenttype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacegeometrycomponenttypeIdent, airspacegeometrycomponenttypeIdent, targetIdent))
			}
			if airspacegeometrycomponenttype.TheAirspaceVolume != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacegeometrycomponenttype.TheAirspaceVolume.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheAirspaceVolume = %s", airspacegeometrycomponenttypeIdent, targetIdent))
			}
			for _, elem := range airspacegeometrycomponenttype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacegeometrycomponenttypeIdent, airspacegeometrycomponenttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacelayerclasspropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceLayerClassPropertyTypes, stage.AixmfeaturesStage.AirspaceLayerClassPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacelayerclasspropertytypeIdent := airspacelayerclasspropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceLayerClassPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacelayerclasspropertytypeIdent, __gong__toRawStringLiteral(airspacelayerclasspropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacelayerclasspropertytypeIdent, __gong__toRawStringLiteral(airspacelayerclasspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacelayerclasspropertytypeIdent, __gong__toRawStringLiteral(airspacelayerclasspropertytype.NilReason)))
			if airspacelayerclasspropertytype.AirspaceLayerClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayerclasspropertytype.AirspaceLayerClass.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceLayerClass = %s", airspacelayerclasspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacelayerclasstype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceLayerClassTypes, stage.AixmfeaturesStage.AirspaceLayerClassType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacelayerclasstypeIdent := airspacelayerclasstype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceLayerClassType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacelayerclasstypeIdent, __gong__toRawStringLiteral(airspacelayerclasstype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacelayerclasstypeIdent, __gong__toRawStringLiteral(airspacelayerclasstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacelayerclasstypeIdent, __gong__toRawStringLiteral(airspacelayerclasstype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacelayerclasstypeIdent, __gong__toRawStringLiteral(airspacelayerclasstype.Description)))
			if airspacelayerclasstype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayerclasstype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airspacelayerclasstypeIdent, targetIdent))
			}
			for _, elem := range airspacelayerclasstype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
			for _, elem := range airspacelayerclasstype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
			for _, elem := range airspacelayerclasstype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
			for _, elem := range airspacelayerclasstype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
			if airspacelayerclasstype.Classification != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayerclasstype.Classification.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Classification = %s", airspacelayerclasstypeIdent, targetIdent))
			}
			for _, elem := range airspacelayerclasstype.AssociatedLevels {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AssociatedLevels = append(%s.AssociatedLevels, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
			for _, elem := range airspacelayerclasstype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacelayerpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceLayerPropertyTypes, stage.AixmfeaturesStage.AirspaceLayerPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacelayerpropertytypeIdent := airspacelayerpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceLayerPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacelayerpropertytypeIdent, __gong__toRawStringLiteral(airspacelayerpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacelayerpropertytypeIdent, __gong__toRawStringLiteral(airspacelayerpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacelayerpropertytypeIdent, __gong__toRawStringLiteral(airspacelayerpropertytype.NilReason)))
			if airspacelayerpropertytype.AirspaceLayer != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayerpropertytype.AirspaceLayer.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceLayer = %s", airspacelayerpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacelayertype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceLayerTypes, stage.AixmfeaturesStage.AirspaceLayerType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacelayertypeIdent := airspacelayertype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceLayerType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Description)))
			if airspacelayertype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayertype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airspacelayertypeIdent, targetIdent))
			}
			for _, elem := range airspacelayertype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airspacelayertypeIdent, airspacelayertypeIdent, targetIdent))
			}
			if airspacelayertype.UpperLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayertype.UpperLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLimit = %s", airspacelayertypeIdent, targetIdent))
			}
			if airspacelayertype.UpperLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayertype.UpperLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLimitReference = %s", airspacelayertypeIdent, targetIdent))
			}
			if airspacelayertype.LowerLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayertype.LowerLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLimit = %s", airspacelayertypeIdent, targetIdent))
			}
			if airspacelayertype.LowerLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayertype.LowerLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLimitReference = %s", airspacelayertypeIdent, targetIdent))
			}
			if airspacelayertype.AltitudeInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayertype.AltitudeInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AltitudeInterpretation = %s", airspacelayertypeIdent, targetIdent))
			}
			if airspacelayertype.DiscreteLevelSeries != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacelayertype.DiscreteLevelSeries.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiscreteLevelSeries = %s", airspacelayertypeIdent, targetIdent))
			}
			for _, elem := range airspacelayertype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacelayertypeIdent, airspacelayertypeIdent, targetIdent))
			}
			for _, elem := range airspacelayertype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacelayertypeIdent, airspacelayertypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspacePropertyTypes, stage.AixmfeaturesStage.AirspacePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacepropertytypeIdent := airspacepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspacePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, airspacetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceTimeSlicePropertyTypes, stage.AixmfeaturesStage.AirspaceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacetimeslicepropertytypeIdent := airspacetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airspacetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airspacetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airspacetimeslicepropertytypeIdent, airspacetimeslicepropertytype.Owns))
			if airspacetimeslicepropertytype.AirspaceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicepropertytype.AirspaceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceTimeSlice = %s", airspacetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceTimeSliceTypes, stage.AixmfeaturesStage.AirspaceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacetimeslicetypeIdent := airspacetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacetimeslicetypeIdent, __gong__toRawStringLiteral(airspacetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacetimeslicetypeIdent, __gong__toRawStringLiteral(airspacetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacetimeslicetypeIdent, __gong__toRawStringLiteral(airspacetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacetimeslicetypeIdent, __gong__toRawStringLiteral(airspacetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airspacetimeslicetypeIdent, __gong__toRawStringLiteral(airspacetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airspacetimeslicetypeIdent, airspacetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airspacetimeslicetypeIdent, airspacetimeslicetype.CorrectionNumber))
			if airspacetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.LocalType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.LocalType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LocalType = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.DesignatorICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.DesignatorICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatorICAO = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.ControlType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.ControlType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlType = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.UpperLowerSeparation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.UpperLowerSeparation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLowerSeparation = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacetimeslicetype.Class {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Class = append(%s.Class, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.ProtectedRoute != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetimeslicetype.ProtectedRoute.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProtectedRoute = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacetimeslicetype.GeometryComponent {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeometryComponent = append(%s.GeometryComponent, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacetimeslicetype.Activation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Activation = append(%s.Activation, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceTypes, stage.AixmfeaturesStage.AirspaceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacetypeIdent := airspacetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacetypeIdent, __gong__toRawStringLiteral(airspacetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacetypeIdent, __gong__toRawStringLiteral(airspacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacetypeIdent, __gong__toRawStringLiteral(airspacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacetypeIdent, __gong__toRawStringLiteral(airspacetype.Description)))
			if airspacetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airspacetypeIdent, targetIdent))
			}
			for _, elem := range airspacetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airspacetypeIdent, airspacetypeIdent, targetIdent))
			}
			if airspacetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", airspacetypeIdent, targetIdent))
			}
			if airspacetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", airspacetypeIdent, targetIdent))
			}
			for _, elem := range airspacetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airspacetypeIdent, airspacetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacevolumedependencypropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceVolumeDependencyPropertyTypes, stage.AixmfeaturesStage.AirspaceVolumeDependencyPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacevolumedependencypropertytypeIdent := airspacevolumedependencypropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceVolumeDependencyPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacevolumedependencypropertytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencypropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacevolumedependencypropertytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacevolumedependencypropertytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencypropertytype.NilReason)))
			if airspacevolumedependencypropertytype.AirspaceVolumeDependency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumedependencypropertytype.AirspaceVolumeDependency.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceVolumeDependency = %s", airspacevolumedependencypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacevolumedependencytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceVolumeDependencyTypes, stage.AixmfeaturesStage.AirspaceVolumeDependencyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacevolumedependencytypeIdent := airspacevolumedependencytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceVolumeDependencyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacevolumedependencytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacevolumedependencytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacevolumedependencytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacevolumedependencytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencytype.Description)))
			if airspacevolumedependencytype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumedependencytype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airspacevolumedependencytypeIdent, targetIdent))
			}
			for _, elem := range airspacevolumedependencytype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airspacevolumedependencytypeIdent, airspacevolumedependencytypeIdent, targetIdent))
			}
			if airspacevolumedependencytype.Dependency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumedependencytype.Dependency.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dependency = %s", airspacevolumedependencytypeIdent, targetIdent))
			}
			for _, elem := range airspacevolumedependencytype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacevolumedependencytypeIdent, airspacevolumedependencytypeIdent, targetIdent))
			}
			if airspacevolumedependencytype.TheAirspace != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumedependencytype.TheAirspace.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheAirspace = %s", airspacevolumedependencytypeIdent, targetIdent))
			}
			for _, elem := range airspacevolumedependencytype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacevolumedependencytypeIdent, airspacevolumedependencytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacevolumepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceVolumePropertyTypes, stage.AixmfeaturesStage.AirspaceVolumePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacevolumepropertytypeIdent := airspacevolumepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceVolumePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacevolumepropertytypeIdent, __gong__toRawStringLiteral(airspacevolumepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacevolumepropertytypeIdent, __gong__toRawStringLiteral(airspacevolumepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacevolumepropertytypeIdent, __gong__toRawStringLiteral(airspacevolumepropertytype.NilReason)))
			if airspacevolumepropertytype.AirspaceVolume != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumepropertytype.AirspaceVolume.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceVolume = %s", airspacevolumepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacevolumetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceVolumeTypes, stage.AixmfeaturesStage.AirspaceVolumeType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacevolumetypeIdent := airspacevolumetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceVolumeType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacevolumetypeIdent, __gong__toRawStringLiteral(airspacevolumetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacevolumetypeIdent, __gong__toRawStringLiteral(airspacevolumetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacevolumetypeIdent, __gong__toRawStringLiteral(airspacevolumetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacevolumetypeIdent, __gong__toRawStringLiteral(airspacevolumetype.Description)))
			if airspacevolumetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airspacevolumetypeIdent, targetIdent))
			}
			for _, elem := range airspacevolumetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airspacevolumetypeIdent, airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.UpperLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.UpperLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLimit = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.UpperLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.UpperLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLimitReference = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.MaximumLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.MaximumLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MaximumLimit = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.MaximumLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.MaximumLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MaximumLimitReference = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.LowerLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.LowerLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLimit = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.LowerLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.LowerLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLimitReference = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.MinimumLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.MinimumLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinimumLimit = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.MinimumLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.MinimumLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinimumLimitReference = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.Width != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.Width.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Width = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.HorizontalProjection != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.HorizontalProjection.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalProjection = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.Centreline != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.Centreline.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Centreline = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.ContributorAirspace != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacevolumetype.ContributorAirspace.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ContributorAirspace = %s", airspacevolumetypeIdent, targetIdent))
			}
			for _, elem := range airspacevolumetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacevolumetypeIdent, airspacevolumetypeIdent, targetIdent))
			}
			for _, elem := range airspacevolumetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacevolumetypeIdent, airspacevolumetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityforairspacetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForAirspaceTimeSlicePropertyTypes, stage.AixmfeaturesStage.AuthorityForAirspaceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityforairspacetimeslicepropertytypeIdent := authorityforairspacetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForAirspaceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityforairspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforairspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", authorityforairspacetimeslicepropertytypeIdent, authorityforairspacetimeslicepropertytype.Owns))
			if authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForAirspaceTimeSlice = %s", authorityforairspacetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityforairspacetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes, stage.AixmfeaturesStage.AuthorityForAirspaceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityforairspacetimeslicetypeIdent := authorityforairspacetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForAirspaceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityforairspacetimeslicetypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforairspacetimeslicetypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityforairspacetimeslicetypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityforairspacetimeslicetypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", authorityforairspacetimeslicetypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", authorityforairspacetimeslicetypeIdent, authorityforairspacetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", authorityforairspacetimeslicetypeIdent, authorityforairspacetimeslicetype.CorrectionNumber))
			if authorityforairspacetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range authorityforairspacetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", authorityforairspacetimeslicetypeIdent, authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			if authorityforairspacetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			if authorityforairspacetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			if authorityforairspacetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			if authorityforairspacetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			if authorityforairspacetimeslicetype.ResponsibleOrganisation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetimeslicetype.ResponsibleOrganisation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResponsibleOrganisation = %s", authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			if authorityforairspacetimeslicetype.AssignedAirspace != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetimeslicetype.AssignedAirspace.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AssignedAirspace = %s", authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range authorityforairspacetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", authorityforairspacetimeslicetypeIdent, authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range authorityforairspacetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", authorityforairspacetimeslicetypeIdent, authorityforairspacetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityforairspacetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForAirspaceTypes, stage.AixmfeaturesStage.AuthorityForAirspaceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityforairspacetypeIdent := authorityforairspacetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForAirspaceType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityforairspacetypeIdent, __gong__toRawStringLiteral(authorityforairspacetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforairspacetypeIdent, __gong__toRawStringLiteral(authorityforairspacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityforairspacetypeIdent, __gong__toRawStringLiteral(authorityforairspacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityforairspacetypeIdent, __gong__toRawStringLiteral(authorityforairspacetype.Description)))
			if authorityforairspacetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", authorityforairspacetypeIdent, targetIdent))
			}
			for _, elem := range authorityforairspacetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", authorityforairspacetypeIdent, authorityforairspacetypeIdent, targetIdent))
			}
			if authorityforairspacetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", authorityforairspacetypeIdent, targetIdent))
			}
			if authorityforairspacetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforairspacetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", authorityforairspacetypeIdent, targetIdent))
			}
			for _, elem := range authorityforairspacetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", authorityforairspacetypeIdent, authorityforairspacetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, curvepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.CurvePropertyTypes, stage.AixmfeaturesStage.CurvePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			curvepropertytypeIdent := curvepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CurvePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.NilReason)))
			if curvepropertytype.Curve != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := curvepropertytype.Curve.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Curve = %s", curvepropertytypeIdent, targetIdent))
			}
			if curvepropertytype.ElevatedCurve != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := curvepropertytype.ElevatedCurve.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedCurve = %s", curvepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, curvetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.CurveTypes, stage.AixmfeaturesStage.CurveType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			curvetypeIdent := curvetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CurveType{Name: %s}).Stage(stage.AixmfeaturesStage)", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.SrsName)))
			if curvetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := curvetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", curvetypeIdent, targetIdent))
			}
			for _, elem := range curvetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", curvetypeIdent, curvetypeIdent, targetIdent))
			}
			if curvetype.Segments != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := curvetype.Segments.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Segments = %s", curvetypeIdent, targetIdent))
			}
			if curvetype.HorizontalAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := curvetype.HorizontalAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalAccuracy = %s", curvetypeIdent, targetIdent))
			}
			for _, elem := range curvetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", curvetypeIdent, curvetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, elevatedcurvetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ElevatedCurveTypes, stage.AixmfeaturesStage.ElevatedCurveType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			elevatedcurvetypeIdent := elevatedcurvetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedCurveType{Name: %s}).Stage(stage.AixmfeaturesStage)", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.SrsName)))
			if elevatedcurvetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedcurvetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", elevatedcurvetypeIdent, targetIdent))
			}
			for _, elem := range elevatedcurvetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", elevatedcurvetypeIdent, elevatedcurvetypeIdent, targetIdent))
			}
			if elevatedcurvetype.Segments != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedcurvetype.Segments.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Segments = %s", elevatedcurvetypeIdent, targetIdent))
			}
			if elevatedcurvetype.HorizontalAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedcurvetype.HorizontalAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalAccuracy = %s", elevatedcurvetypeIdent, targetIdent))
			}
			for _, elem := range elevatedcurvetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevatedcurvetypeIdent, elevatedcurvetypeIdent, targetIdent))
			}
			if elevatedcurvetype.Elevation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedcurvetype.Elevation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elevation = %s", elevatedcurvetypeIdent, targetIdent))
			}
			if elevatedcurvetype.GeoidUndulation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedcurvetype.GeoidUndulation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoidUndulation = %s", elevatedcurvetypeIdent, targetIdent))
			}
			if elevatedcurvetype.VerticalDatum != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedcurvetype.VerticalDatum.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalDatum = %s", elevatedcurvetypeIdent, targetIdent))
			}
			if elevatedcurvetype.VerticalAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedcurvetype.VerticalAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalAccuracy = %s", elevatedcurvetypeIdent, targetIdent))
			}
			for _, elem := range elevatedcurvetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevatedcurvetypeIdent, elevatedcurvetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, elevatedsurfacetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ElevatedSurfaceTypes, stage.AixmfeaturesStage.ElevatedSurfaceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			elevatedsurfacetypeIdent := elevatedsurfacetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedSurfaceType{Name: %s}).Stage(stage.AixmfeaturesStage)", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.SrsName)))
			if elevatedsurfacetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedsurfacetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", elevatedsurfacetypeIdent, targetIdent))
			}
			for _, elem := range elevatedsurfacetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", elevatedsurfacetypeIdent, elevatedsurfacetypeIdent, targetIdent))
			}
			if elevatedsurfacetype.Patches != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedsurfacetype.Patches.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Patches = %s", elevatedsurfacetypeIdent, targetIdent))
			}
			if elevatedsurfacetype.HorizontalAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedsurfacetype.HorizontalAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalAccuracy = %s", elevatedsurfacetypeIdent, targetIdent))
			}
			for _, elem := range elevatedsurfacetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevatedsurfacetypeIdent, elevatedsurfacetypeIdent, targetIdent))
			}
			if elevatedsurfacetype.Elevation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedsurfacetype.Elevation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elevation = %s", elevatedsurfacetypeIdent, targetIdent))
			}
			if elevatedsurfacetype.GeoidUndulation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedsurfacetype.GeoidUndulation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoidUndulation = %s", elevatedsurfacetypeIdent, targetIdent))
			}
			if elevatedsurfacetype.VerticalDatum != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedsurfacetype.VerticalDatum.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalDatum = %s", elevatedsurfacetypeIdent, targetIdent))
			}
			if elevatedsurfacetype.VerticalAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedsurfacetype.VerticalAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalAccuracy = %s", elevatedsurfacetypeIdent, targetIdent))
			}
			for _, elem := range elevatedsurfacetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevatedsurfacetypeIdent, elevatedsurfacetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, geobordertimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.GeoBorderTimeSlicePropertyTypes, stage.AixmfeaturesStage.GeoBorderTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			geobordertimeslicepropertytypeIdent := geobordertimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GeoBorderTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", geobordertimeslicepropertytypeIdent, __gong__toRawStringLiteral(geobordertimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geobordertimeslicepropertytypeIdent, __gong__toRawStringLiteral(geobordertimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", geobordertimeslicepropertytypeIdent, geobordertimeslicepropertytype.Owns))
			if geobordertimeslicepropertytype.GeoBorderTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := geobordertimeslicepropertytype.GeoBorderTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoBorderTimeSlice = %s", geobordertimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, geobordertimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.GeoBorderTimeSliceTypes, stage.AixmfeaturesStage.GeoBorderTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			geobordertimeslicetypeIdent := geobordertimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GeoBorderTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", geobordertimeslicetypeIdent, __gong__toRawStringLiteral(geobordertimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geobordertimeslicetypeIdent, __gong__toRawStringLiteral(geobordertimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", geobordertimeslicetypeIdent, __gong__toRawStringLiteral(geobordertimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", geobordertimeslicetypeIdent, __gong__toRawStringLiteral(geobordertimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", geobordertimeslicetypeIdent, __gong__toRawStringLiteral(geobordertimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", geobordertimeslicetypeIdent, geobordertimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", geobordertimeslicetypeIdent, geobordertimeslicetype.CorrectionNumber))
			if geobordertimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := geobordertimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", geobordertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range geobordertimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", geobordertimeslicetypeIdent, geobordertimeslicetypeIdent, targetIdent))
			}
			if geobordertimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := geobordertimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", geobordertimeslicetypeIdent, targetIdent))
			}
			if geobordertimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := geobordertimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", geobordertimeslicetypeIdent, targetIdent))
			}
			if geobordertimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := geobordertimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", geobordertimeslicetypeIdent, targetIdent))
			}
			if geobordertimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := geobordertimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", geobordertimeslicetypeIdent, targetIdent))
			}
			if geobordertimeslicetype.Border != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := geobordertimeslicetype.Border.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Border = %s", geobordertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range geobordertimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", geobordertimeslicetypeIdent, geobordertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range geobordertimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", geobordertimeslicetypeIdent, geobordertimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, geobordertype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.GeoBorderTypes, stage.AixmfeaturesStage.GeoBorderType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			geobordertypeIdent := geobordertype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GeoBorderType{Name: %s}).Stage(stage.AixmfeaturesStage)", geobordertypeIdent, __gong__toRawStringLiteral(geobordertype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geobordertypeIdent, __gong__toRawStringLiteral(geobordertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", geobordertypeIdent, __gong__toRawStringLiteral(geobordertype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", geobordertypeIdent, __gong__toRawStringLiteral(geobordertype.Description)))
			if geobordertype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := geobordertype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", geobordertypeIdent, targetIdent))
			}
			for _, elem := range geobordertype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", geobordertypeIdent, geobordertypeIdent, targetIdent))
			}
			if geobordertype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := geobordertype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", geobordertypeIdent, targetIdent))
			}
			if geobordertype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := geobordertype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", geobordertypeIdent, targetIdent))
			}
			for _, elem := range geobordertype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", geobordertypeIdent, geobordertypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, linguisticnotepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LinguisticNotePropertyTypes, stage.AixmfeaturesStage.LinguisticNotePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			linguisticnotepropertytypeIdent := linguisticnotepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LinguisticNotePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", linguisticnotepropertytypeIdent, __gong__toRawStringLiteral(linguisticnotepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linguisticnotepropertytypeIdent, __gong__toRawStringLiteral(linguisticnotepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", linguisticnotepropertytypeIdent, __gong__toRawStringLiteral(linguisticnotepropertytype.NilReason)))
			if linguisticnotepropertytype.LinguisticNote != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := linguisticnotepropertytype.LinguisticNote.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LinguisticNote = %s", linguisticnotepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, linguisticnotetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LinguisticNoteTypes, stage.AixmfeaturesStage.LinguisticNoteType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			linguisticnotetypeIdent := linguisticnotetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LinguisticNoteType{Name: %s}).Stage(stage.AixmfeaturesStage)", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Description)))
			if linguisticnotetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := linguisticnotetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", linguisticnotetypeIdent, targetIdent))
			}
			for _, elem := range linguisticnotetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", linguisticnotetypeIdent, linguisticnotetypeIdent, targetIdent))
			}
			if linguisticnotetype.Note != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := linguisticnotetype.Note.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", linguisticnotetypeIdent, targetIdent))
			}
			for _, elem := range linguisticnotetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", linguisticnotetypeIdent, linguisticnotetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, notepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NotePropertyTypes, stage.AixmfeaturesStage.NotePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			notepropertytypeIdent := notepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NotePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", notepropertytypeIdent, __gong__toRawStringLiteral(notepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notepropertytypeIdent, __gong__toRawStringLiteral(notepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", notepropertytypeIdent, __gong__toRawStringLiteral(notepropertytype.NilReason)))
			if notepropertytype.Note != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := notepropertytype.Note.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, notetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NoteTypes, stage.AixmfeaturesStage.NoteType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			notetypeIdent := notetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NoteType{Name: %s}).Stage(stage.AixmfeaturesStage)", notetypeIdent, __gong__toRawStringLiteral(notetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Description)))
			if notetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := notetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", notetypeIdent, targetIdent))
			}
			for _, elem := range notetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", notetypeIdent, notetypeIdent, targetIdent))
			}
			if notetype.PropertyName != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := notetype.PropertyName.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PropertyName = %s", notetypeIdent, targetIdent))
			}
			if notetype.Purpose != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := notetype.Purpose.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Purpose = %s", notetypeIdent, targetIdent))
			}
			for _, elem := range notetype.TranslatedNote {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TranslatedNote = append(%s.TranslatedNote, %s)", notetypeIdent, notetypeIdent, targetIdent))
			}
			for _, elem := range notetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", notetypeIdent, notetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, organisationauthoritypropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OrganisationAuthorityPropertyTypes, stage.AixmfeaturesStage.OrganisationAuthorityPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			organisationauthoritypropertytypeIdent := organisationauthoritypropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, routepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RoutePropertyTypes, stage.AixmfeaturesStage.RoutePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routepropertytypeIdent := routepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RoutePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, standardlevelcolumnpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelColumnPropertyTypes, stage.AixmfeaturesStage.StandardLevelColumnPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardlevelcolumnpropertytypeIdent := standardlevelcolumnpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelColumnPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, surfacepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SurfacePropertyTypes, stage.AixmfeaturesStage.SurfacePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			surfacepropertytypeIdent := surfacepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SurfacePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", surfacepropertytypeIdent, __gong__toRawStringLiteral(surfacepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surfacepropertytypeIdent, __gong__toRawStringLiteral(surfacepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", surfacepropertytypeIdent, __gong__toRawStringLiteral(surfacepropertytype.NilReason)))
			if surfacepropertytype.Surface != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := surfacepropertytype.Surface.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Surface = %s", surfacepropertytypeIdent, targetIdent))
			}
			if surfacepropertytype.ElevatedSurface != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := surfacepropertytype.ElevatedSurface.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedSurface = %s", surfacepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, surfacetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SurfaceTypes, stage.AixmfeaturesStage.SurfaceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			surfacetypeIdent := surfacetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SurfaceType{Name: %s}).Stage(stage.AixmfeaturesStage)", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.SrsName)))
			if surfacetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := surfacetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", surfacetypeIdent, targetIdent))
			}
			for _, elem := range surfacetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", surfacetypeIdent, surfacetypeIdent, targetIdent))
			}
			if surfacetype.Patches != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := surfacetype.Patches.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Patches = %s", surfacetypeIdent, targetIdent))
			}
			if surfacetype.HorizontalAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := surfacetype.HorizontalAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalAccuracy = %s", surfacetypeIdent, targetIdent))
			}
			for _, elem := range surfacetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", surfacetypeIdent, surfacetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, timesheetpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.TimesheetPropertyTypes, stage.AixmfeaturesStage.TimesheetPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			timesheetpropertytypeIdent := timesheetpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TimesheetPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", timesheetpropertytypeIdent, __gong__toRawStringLiteral(timesheetpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timesheetpropertytypeIdent, __gong__toRawStringLiteral(timesheetpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", timesheetpropertytypeIdent, __gong__toRawStringLiteral(timesheetpropertytype.NilReason)))
			if timesheetpropertytype.Timesheet != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheetpropertytype.Timesheet.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Timesheet = %s", timesheetpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, timesheettype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.TimesheetTypes, stage.AixmfeaturesStage.TimesheetType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			timesheettypeIdent := timesheettype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TimesheetType{Name: %s}).Stage(stage.AixmfeaturesStage)", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Description)))
			if timesheettype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", timesheettypeIdent, targetIdent))
			}
			for _, elem := range timesheettype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", timesheettypeIdent, timesheettypeIdent, targetIdent))
			}
			if timesheettype.TimeReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.TimeReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeReference = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.StartDate != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.StartDate.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartDate = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.EndDate != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.EndDate.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndDate = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.Day != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.Day.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Day = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.DayTil != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.DayTil.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DayTil = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.StartTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.StartTime.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartTime = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.StartEvent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.StartEvent.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartEvent = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.StartTimeRelativeEvent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.StartTimeRelativeEvent.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartTimeRelativeEvent = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.StartEventInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.StartEventInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartEventInterpretation = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.EndTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.EndTime.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndTime = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.EndEvent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.EndEvent.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndEvent = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.EndTimeRelativeEvent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.EndTimeRelativeEvent.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndTimeRelativeEvent = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.EndEventInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.EndEventInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndEventInterpretation = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.DaylightSavingAdjust != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.DaylightSavingAdjust.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DaylightSavingAdjust = %s", timesheettypeIdent, targetIdent))
			}
			if timesheettype.Excluded != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := timesheettype.Excluded.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Excluded = %s", timesheettypeIdent, targetIdent))
			}
			for _, elem := range timesheettype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", timesheettypeIdent, timesheettypeIdent, targetIdent))
			}
			for _, elem := range timesheettype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", timesheettypeIdent, timesheettypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aixmbasicmessage := range __gong__sortStageSetInstances(stage.AIXMBasicMessages, stage.AIXMBasicMessage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aixmbasicmessageIdent := aixmbasicmessage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AIXMBasicMessage{Name: %s}).Stage(stage)", aixmbasicmessageIdent, __gong__toRawStringLiteral(aixmbasicmessage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aixmbasicmessageIdent, __gong__toRawStringLiteral(aixmbasicmessage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aixmbasicmessageIdent, __gong__toRawStringLiteral(aixmbasicmessage.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aixmbasicmessageIdent, __gong__toRawStringLiteral(aixmbasicmessage.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", aixmbasicmessageIdent, aixmbasicmessage.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.AggregationType = %s", aixmbasicmessageIdent, __gong__toRawStringLiteral(aixmbasicmessage.AggregationType)))
			if aixmbasicmessage.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aixmbasicmessage.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aixmbasicmessageIdent, targetIdent))
			}
			for _, elem := range aixmbasicmessage.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aixmbasicmessageIdent, aixmbasicmessageIdent, targetIdent))
			}
			if aixmbasicmessage.BoundedBy != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aixmbasicmessage.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", aixmbasicmessageIdent, targetIdent))
			}
			if aixmbasicmessage.MessageMetadata != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aixmbasicmessage.MessageMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MessageMetadata = %s", aixmbasicmessageIdent, targetIdent))
			}
			for _, elem := range aixmbasicmessage.HasMember {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HasMember = append(%s.HasMember, %s)", aixmbasicmessageIdent, aixmbasicmessageIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aixmbasicmessagetype := range __gong__sortStageSetInstances(stage.AIXMBasicMessageTypes, stage.AIXMBasicMessageType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aixmbasicmessagetypeIdent := aixmbasicmessagetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AIXMBasicMessageType{Name: %s}).Stage(stage)", aixmbasicmessagetypeIdent, __gong__toRawStringLiteral(aixmbasicmessagetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aixmbasicmessagetypeIdent, __gong__toRawStringLiteral(aixmbasicmessagetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aixmbasicmessagetypeIdent, __gong__toRawStringLiteral(aixmbasicmessagetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aixmbasicmessagetypeIdent, __gong__toRawStringLiteral(aixmbasicmessagetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", aixmbasicmessagetypeIdent, aixmbasicmessagetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.AggregationType = %s", aixmbasicmessagetypeIdent, __gong__toRawStringLiteral(aixmbasicmessagetype.AggregationType)))
			if aixmbasicmessagetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aixmbasicmessagetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aixmbasicmessagetypeIdent, targetIdent))
			}
			for _, elem := range aixmbasicmessagetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aixmbasicmessagetypeIdent, aixmbasicmessagetypeIdent, targetIdent))
			}
			if aixmbasicmessagetype.BoundedBy != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aixmbasicmessagetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", aixmbasicmessagetypeIdent, targetIdent))
			}
			if aixmbasicmessagetype.MessageMetadata != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aixmbasicmessagetype.MessageMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MessageMetadata = %s", aixmbasicmessagetypeIdent, targetIdent))
			}
			for _, elem := range aixmbasicmessagetype.HasMember {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HasMember = append(%s.HasMember, %s)", aixmbasicmessagetypeIdent, aixmbasicmessagetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, basicmessagememberaixmpropertytype := range __gong__sortStageSetInstances(stage.BasicMessageMemberAIXMPropertyTypes, stage.BasicMessageMemberAIXMPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			basicmessagememberaixmpropertytypeIdent := basicmessagememberaixmpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.BasicMessageMemberAIXMPropertyType{Name: %s}).Stage(stage)", basicmessagememberaixmpropertytypeIdent, __gong__toRawStringLiteral(basicmessagememberaixmpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", basicmessagememberaixmpropertytypeIdent, __gong__toRawStringLiteral(basicmessagememberaixmpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", basicmessagememberaixmpropertytypeIdent, __gong__toRawStringLiteral(basicmessagememberaixmpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", basicmessagememberaixmpropertytypeIdent, __gong__toRawStringLiteral(basicmessagememberaixmpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", basicmessagememberaixmpropertytypeIdent, __gong__toRawStringLiteral(basicmessagememberaixmpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", basicmessagememberaixmpropertytypeIdent, __gong__toRawStringLiteral(basicmessagememberaixmpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", basicmessagememberaixmpropertytypeIdent, __gong__toRawStringLiteral(basicmessagememberaixmpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", basicmessagememberaixmpropertytypeIdent, __gong__toRawStringLiteral(basicmessagememberaixmpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", basicmessagememberaixmpropertytypeIdent, __gong__toRawStringLiteral(basicmessagememberaixmpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", basicmessagememberaixmpropertytypeIdent, __gong__toRawStringLiteral(basicmessagememberaixmpropertytype.Actuate)))
			if basicmessagememberaixmpropertytype.Airspace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.Airspace.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Airspace = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.AuthorityForAirspace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.AuthorityForAirspace.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForAirspace = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.GeoBorder != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.GeoBorder.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoBorder = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/xlink"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
	_ *aixmfeatures.Stage
	_ *aixmdatatypes.Stage
	_ *aixmabstractgmlobjecttypes.Stage
	_ *gml.Stage
	_ *xlink.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *models.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (aixmbasicmessage *AIXMBasicMessage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aixmbasicmessage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aixmbasicmessage.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aixmbasicmessage.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aixmbasicmessage.Description)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", aixmbasicmessage.SequenceNumber)
	case "AggregationType":
		res = __gong__marshallString(ident, "AggregationType", aixmbasicmessage.AggregationType)

	case "HasMember":
		var sb strings.Builder
		for _, _basicmessagememberaixmpropertytype := range aixmbasicmessage.HasMember {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "HasMember", _basicmessagememberaixmpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AIXMBasicMessage", fieldName)
	}
	return
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aixmbasicmessagetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aixmbasicmessagetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aixmbasicmessagetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aixmbasicmessagetype.Description)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", aixmbasicmessagetype.SequenceNumber)
	case "AggregationType":
		res = __gong__marshallString(ident, "AggregationType", aixmbasicmessagetype.AggregationType)

	case "HasMember":
		var sb strings.Builder
		for _, _basicmessagememberaixmpropertytype := range aixmbasicmessagetype.HasMember {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "HasMember", _basicmessagememberaixmpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AIXMBasicMessageType", fieldName)
	}
	return
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := basicmessagememberaixmpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", basicmessagememberaixmpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", basicmessagememberaixmpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", basicmessagememberaixmpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", basicmessagememberaixmpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", basicmessagememberaixmpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", basicmessagememberaixmpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", basicmessagememberaixmpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", basicmessagememberaixmpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", basicmessagememberaixmpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct BasicMessageMemberAIXMPropertyType", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (aixmbasicmessage *AIXMBasicMessage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aixmbasicmessage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aixmbasicmessage.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aixmbasicmessage.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(aixmbasicmessage.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(aixmbasicmessage.GongMarshallField(stage, "AggregationType"))
		pointersInitializesStatements.WriteString(aixmbasicmessage.GongMarshallField(stage, "HasMember"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aixmbasicmessagetype *AIXMBasicMessageType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aixmbasicmessagetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aixmbasicmessagetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aixmbasicmessagetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(aixmbasicmessagetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(aixmbasicmessagetype.GongMarshallField(stage, "AggregationType"))
		pointersInitializesStatements.WriteString(aixmbasicmessagetype.GongMarshallField(stage, "HasMember"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(basicmessagememberaixmpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(basicmessagememberaixmpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(basicmessagememberaixmpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(basicmessagememberaixmpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(basicmessagememberaixmpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(basicmessagememberaixmpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(basicmessagememberaixmpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(basicmessagememberaixmpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(basicmessagememberaixmpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
