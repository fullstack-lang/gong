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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models"
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
		for _, codecolourtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeColourTypes, stage.AixmdatatypesStage.CodeColourType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codecolourtypeIdent := codecolourtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeColourType{Name: %s}).Stage(stage.AixmdatatypesStage)", codecolourtypeIdent, __gong__toRawStringLiteral(codecolourtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecolourtypeIdent, __gong__toRawStringLiteral(codecolourtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecolourtypeIdent, __gong__toRawStringLiteral(codecolourtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecolourtypeIdent, __gong__toRawStringLiteral(codecolourtype.EnclosedText)))
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
		for _, codegroundlightingtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeGroundLightingTypes, stage.AixmdatatypesStage.CodeGroundLightingType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codegroundlightingtypeIdent := codegroundlightingtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeGroundLightingType{Name: %s}).Stage(stage.AixmdatatypesStage)", codegroundlightingtypeIdent, __gong__toRawStringLiteral(codegroundlightingtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codegroundlightingtypeIdent, __gong__toRawStringLiteral(codegroundlightingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codegroundlightingtypeIdent, __gong__toRawStringLiteral(codegroundlightingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codegroundlightingtypeIdent, __gong__toRawStringLiteral(codegroundlightingtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codelightintensitytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeLightIntensityTypes, stage.AixmdatatypesStage.CodeLightIntensityType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codelightintensitytypeIdent := codelightintensitytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLightIntensityType{Name: %s}).Stage(stage.AixmdatatypesStage)", codelightintensitytypeIdent, __gong__toRawStringLiteral(codelightintensitytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelightintensitytypeIdent, __gong__toRawStringLiteral(codelightintensitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelightintensitytypeIdent, __gong__toRawStringLiteral(codelightintensitytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelightintensitytypeIdent, __gong__toRawStringLiteral(codelightintensitytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codelightsourcetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeLightSourceTypes, stage.AixmdatatypesStage.CodeLightSourceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codelightsourcetypeIdent := codelightsourcetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLightSourceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codelightsourcetypeIdent, __gong__toRawStringLiteral(codelightsourcetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelightsourcetypeIdent, __gong__toRawStringLiteral(codelightsourcetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelightsourcetypeIdent, __gong__toRawStringLiteral(codelightsourcetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelightsourcetypeIdent, __gong__toRawStringLiteral(codelightsourcetype.EnclosedText)))
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
		for _, codeobstacleareatype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeObstacleAreaTypes, stage.AixmdatatypesStage.CodeObstacleAreaType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeobstacleareatypeIdent := codeobstacleareatype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeObstacleAreaType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeobstacleareatypeIdent, __gong__toRawStringLiteral(codeobstacleareatype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeobstacleareatypeIdent, __gong__toRawStringLiteral(codeobstacleareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeobstacleareatypeIdent, __gong__toRawStringLiteral(codeobstacleareatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeobstacleareatypeIdent, __gong__toRawStringLiteral(codeobstacleareatype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeobstacleassessmentsurfacetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeObstacleAssessmentSurfaceTypes, stage.AixmdatatypesStage.CodeObstacleAssessmentSurfaceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeobstacleassessmentsurfacetypeIdent := codeobstacleassessmentsurfacetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeObstacleAssessmentSurfaceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeobstacleassessmentsurfacetypeIdent, __gong__toRawStringLiteral(codeobstacleassessmentsurfacetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeobstacleassessmentsurfacetypeIdent, __gong__toRawStringLiteral(codeobstacleassessmentsurfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeobstacleassessmentsurfacetypeIdent, __gong__toRawStringLiteral(codeobstacleassessmentsurfacetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeobstacleassessmentsurfacetypeIdent, __gong__toRawStringLiteral(codeobstacleassessmentsurfacetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codestatusconstructiontype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeStatusConstructionTypes, stage.AixmdatatypesStage.CodeStatusConstructionType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codestatusconstructiontypeIdent := codestatusconstructiontype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusConstructionType{Name: %s}).Stage(stage.AixmdatatypesStage)", codestatusconstructiontypeIdent, __gong__toRawStringLiteral(codestatusconstructiontype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusconstructiontypeIdent, __gong__toRawStringLiteral(codestatusconstructiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusconstructiontypeIdent, __gong__toRawStringLiteral(codestatusconstructiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusconstructiontypeIdent, __gong__toRawStringLiteral(codestatusconstructiontype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codestatusoperationstype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeStatusOperationsTypes, stage.AixmdatatypesStage.CodeStatusOperationsType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codestatusoperationstypeIdent := codestatusoperationstype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusOperationsType{Name: %s}).Stage(stage.AixmdatatypesStage)", codestatusoperationstypeIdent, __gong__toRawStringLiteral(codestatusoperationstype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusoperationstypeIdent, __gong__toRawStringLiteral(codestatusoperationstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusoperationstypeIdent, __gong__toRawStringLiteral(codestatusoperationstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusoperationstypeIdent, __gong__toRawStringLiteral(codestatusoperationstype.EnclosedText)))
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
		for _, codeverticalstructuremarkingtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeVerticalStructureMarkingTypes, stage.AixmdatatypesStage.CodeVerticalStructureMarkingType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeverticalstructuremarkingtypeIdent := codeverticalstructuremarkingtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVerticalStructureMarkingType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeverticalstructuremarkingtypeIdent, __gong__toRawStringLiteral(codeverticalstructuremarkingtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeverticalstructuremarkingtypeIdent, __gong__toRawStringLiteral(codeverticalstructuremarkingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeverticalstructuremarkingtypeIdent, __gong__toRawStringLiteral(codeverticalstructuremarkingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeverticalstructuremarkingtypeIdent, __gong__toRawStringLiteral(codeverticalstructuremarkingtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeverticalstructurematerialtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeVerticalStructureMaterialTypes, stage.AixmdatatypesStage.CodeVerticalStructureMaterialType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeverticalstructurematerialtypeIdent := codeverticalstructurematerialtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVerticalStructureMaterialType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeverticalstructurematerialtypeIdent, __gong__toRawStringLiteral(codeverticalstructurematerialtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeverticalstructurematerialtypeIdent, __gong__toRawStringLiteral(codeverticalstructurematerialtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeverticalstructurematerialtypeIdent, __gong__toRawStringLiteral(codeverticalstructurematerialtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeverticalstructurematerialtypeIdent, __gong__toRawStringLiteral(codeverticalstructurematerialtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeverticalstructuretype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeVerticalStructureTypes, stage.AixmdatatypesStage.CodeVerticalStructureType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeverticalstructuretypeIdent := codeverticalstructuretype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVerticalStructureType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeverticalstructuretypeIdent, __gong__toRawStringLiteral(codeverticalstructuretype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeverticalstructuretypeIdent, __gong__toRawStringLiteral(codeverticalstructuretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeverticalstructuretypeIdent, __gong__toRawStringLiteral(codeverticalstructuretype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeverticalstructuretypeIdent, __gong__toRawStringLiteral(codeverticalstructuretype.EnclosedText)))
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
		for _, textdesignatortype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.TextDesignatorTypes, stage.AixmdatatypesStage.TextDesignatorType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			textdesignatortypeIdent := textdesignatortype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextDesignatorType{Name: %s}).Stage(stage.AixmdatatypesStage)", textdesignatortypeIdent, __gong__toRawStringLiteral(textdesignatortype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textdesignatortypeIdent, __gong__toRawStringLiteral(textdesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textdesignatortypeIdent, __gong__toRawStringLiteral(textdesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textdesignatortypeIdent, __gong__toRawStringLiteral(textdesignatortype.EnclosedText)))
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
		for _, vallightintensitytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValLightIntensityTypes, stage.AixmdatatypesStage.ValLightIntensityType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			vallightintensitytypeIdent := vallightintensitytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValLightIntensityType{Name: %s}).Stage(stage.AixmdatatypesStage)", vallightintensitytypeIdent, __gong__toRawStringLiteral(vallightintensitytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vallightintensitytypeIdent, __gong__toRawStringLiteral(vallightintensitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", vallightintensitytypeIdent, __gong__toRawStringLiteral(vallightintensitytype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", vallightintensitytypeIdent, __gong__toRawStringLiteral(vallightintensitytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", vallightintensitytypeIdent, __gong__toRawStringLiteral(vallightintensitytype.EnclosedText)))
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
			for _, elem := range a_extension.AbstractElevatedCurveExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractElevatedCurveExtension = append(%s.AbstractElevatedCurveExtension, %s)", a_extensionIdent, a_extensionIdent, targetIdent))
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
			for _, elem := range a_extension_1.AbstractElevatedPointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractElevatedPointExtension = append(%s.AbstractElevatedPointExtension, %s)", a_extension_1Ident, a_extension_1Ident, targetIdent))
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
			for _, elem := range a_extension_10.AbstractLightElementExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractLightElementExtension = append(%s.AbstractLightElementExtension, %s)", a_extension_10Ident, a_extension_10Ident, targetIdent))
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
			for _, elem := range a_extension_11.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_11Ident, a_extension_11Ident, targetIdent))
			}
			for _, elem := range a_extension_11.AbstractLightElementStatusExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractLightElementStatusExtension = append(%s.AbstractLightElementStatusExtension, %s)", a_extension_11Ident, a_extension_11Ident, targetIdent))
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
			for _, elem := range a_extension_12.AbstractTimesheetExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractTimesheetExtension = append(%s.AbstractTimesheetExtension, %s)", a_extension_12Ident, a_extension_12Ident, targetIdent))
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
			for _, elem := range a_extension_2.AbstractElevatedSurfaceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractElevatedSurfaceExtension = append(%s.AbstractElevatedSurfaceExtension, %s)", a_extension_2Ident, a_extension_2Ident, targetIdent))
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
			for _, elem := range a_extension_3.AbstractAeronauticalGroundLightExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAeronauticalGroundLightExtension = append(%s.AbstractAeronauticalGroundLightExtension, %s)", a_extension_3Ident, a_extension_3Ident, targetIdent))
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
			for _, elem := range a_extension_4.AbstractLinguisticNoteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractLinguisticNoteExtension = append(%s.AbstractLinguisticNoteExtension, %s)", a_extension_4Ident, a_extension_4Ident, targetIdent))
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
			for _, elem := range a_extension_5.AbstractNoteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNoteExtension = append(%s.AbstractNoteExtension, %s)", a_extension_5Ident, a_extension_5Ident, targetIdent))
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
			for _, elem := range a_extension_6.AbstractObstacleAreaExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractObstacleAreaExtension = append(%s.AbstractObstacleAreaExtension, %s)", a_extension_6Ident, a_extension_6Ident, targetIdent))
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
			for _, elem := range a_extension_7.AbstractVerticalStructureExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractVerticalStructureExtension = append(%s.AbstractVerticalStructureExtension, %s)", a_extension_7Ident, a_extension_7Ident, targetIdent))
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
			for _, elem := range a_extension_8.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_8Ident, a_extension_8Ident, targetIdent))
			}
			for _, elem := range a_extension_8.AbstractVerticalStructureLightingStatusExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractVerticalStructureLightingStatusExtension = append(%s.AbstractVerticalStructureLightingStatusExtension, %s)", a_extension_8Ident, a_extension_8Ident, targetIdent))
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
			for _, elem := range a_extension_9.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_9Ident, a_extension_9Ident, targetIdent))
			}
			for _, elem := range a_extension_9.AbstractVerticalStructurePartExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractVerticalStructurePartExtension = append(%s.AbstractVerticalStructurePartExtension, %s)", a_extension_9Ident, a_extension_9Ident, targetIdent))
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
		for _, aeronauticalgroundlighttimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AeronauticalGroundLightTimeSlicePropertyTypes, stage.AixmfeaturesStage.AeronauticalGroundLightTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aeronauticalgroundlighttimeslicepropertytypeIdent := aeronauticalgroundlighttimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", aeronauticalgroundlighttimeslicepropertytypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aeronauticalgroundlighttimeslicepropertytypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", aeronauticalgroundlighttimeslicepropertytypeIdent, aeronauticalgroundlighttimeslicepropertytype.Owns))
			if aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AeronauticalGroundLightTimeSlice = %s", aeronauticalgroundlighttimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aeronauticalgroundlighttimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes, stage.AixmfeaturesStage.AeronauticalGroundLightTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aeronauticalgroundlighttimeslicetypeIdent := aeronauticalgroundlighttimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AeronauticalGroundLightTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", aeronauticalgroundlighttimeslicetypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aeronauticalgroundlighttimeslicetypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aeronauticalgroundlighttimeslicetypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aeronauticalgroundlighttimeslicetypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", aeronauticalgroundlighttimeslicetypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", aeronauticalgroundlighttimeslicetypeIdent, aeronauticalgroundlighttimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", aeronauticalgroundlighttimeslicetypeIdent, aeronauticalgroundlighttimeslicetype.CorrectionNumber))
			if aeronauticalgroundlighttimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aeronauticalgroundlighttimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aeronauticalgroundlighttimeslicetypeIdent, aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.Colour != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.Colour.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Colour = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.Flashing != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.Flashing.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Flashing = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.StructureBeacon != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.StructureBeacon.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StructureBeacon = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.AerodromeBeacon != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.AerodromeBeacon.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerodromeBeacon = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aeronauticalgroundlighttimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aeronauticalgroundlighttimeslicetypeIdent, aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aeronauticalgroundlighttimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aeronauticalgroundlighttimeslicetypeIdent, aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aeronauticalgroundlighttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AeronauticalGroundLightTypes, stage.AixmfeaturesStage.AeronauticalGroundLightType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aeronauticalgroundlighttypeIdent := aeronauticalgroundlighttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AeronauticalGroundLightType{Name: %s}).Stage(stage.AixmfeaturesStage)", aeronauticalgroundlighttypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aeronauticalgroundlighttypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aeronauticalgroundlighttypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aeronauticalgroundlighttypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttype.Description)))
			if aeronauticalgroundlighttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aeronauticalgroundlighttypeIdent, targetIdent))
			}
			for _, elem := range aeronauticalgroundlighttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aeronauticalgroundlighttypeIdent, aeronauticalgroundlighttypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", aeronauticalgroundlighttypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aeronauticalgroundlighttype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", aeronauticalgroundlighttypeIdent, targetIdent))
			}
			for _, elem := range aeronauticalgroundlighttype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", aeronauticalgroundlighttypeIdent, aeronauticalgroundlighttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airportheliportpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirportHeliportPropertyTypes, stage.AixmfeaturesStage.AirportHeliportPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airportheliportpropertytypeIdent := airportheliportpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportHeliportPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
		for _, elevatedcurvepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ElevatedCurvePropertyTypes, stage.AixmfeaturesStage.ElevatedCurvePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			elevatedcurvepropertytypeIdent := elevatedcurvepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedCurvePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", elevatedcurvepropertytypeIdent, __gong__toRawStringLiteral(elevatedcurvepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedcurvepropertytypeIdent, __gong__toRawStringLiteral(elevatedcurvepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", elevatedcurvepropertytypeIdent, __gong__toRawStringLiteral(elevatedcurvepropertytype.NilReason)))
			if elevatedcurvepropertytype.ElevatedCurve != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedcurvepropertytype.ElevatedCurve.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedCurve = %s", elevatedcurvepropertytypeIdent, targetIdent))
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
		for _, elevatedpointpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ElevatedPointPropertyTypes, stage.AixmfeaturesStage.ElevatedPointPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			elevatedpointpropertytypeIdent := elevatedpointpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedPointPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", elevatedpointpropertytypeIdent, __gong__toRawStringLiteral(elevatedpointpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedpointpropertytypeIdent, __gong__toRawStringLiteral(elevatedpointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", elevatedpointpropertytypeIdent, __gong__toRawStringLiteral(elevatedpointpropertytype.NilReason)))
			if elevatedpointpropertytype.ElevatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedpointpropertytype.ElevatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedPoint = %s", elevatedpointpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, elevatedpointtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ElevatedPointTypes, stage.AixmfeaturesStage.ElevatedPointType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			elevatedpointtypeIdent := elevatedpointtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedPointType{Name: %s}).Stage(stage.AixmfeaturesStage)", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.SrsName)))
			if elevatedpointtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedpointtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", elevatedpointtypeIdent, targetIdent))
			}
			for _, elem := range elevatedpointtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", elevatedpointtypeIdent, elevatedpointtypeIdent, targetIdent))
			}
			if elevatedpointtype.Pos != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedpointtype.Pos.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", elevatedpointtypeIdent, targetIdent))
			}
			if elevatedpointtype.HorizontalAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedpointtype.HorizontalAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalAccuracy = %s", elevatedpointtypeIdent, targetIdent))
			}
			for _, elem := range elevatedpointtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevatedpointtypeIdent, elevatedpointtypeIdent, targetIdent))
			}
			if elevatedpointtype.Elevation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedpointtype.Elevation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elevation = %s", elevatedpointtypeIdent, targetIdent))
			}
			if elevatedpointtype.GeoidUndulation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedpointtype.GeoidUndulation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoidUndulation = %s", elevatedpointtypeIdent, targetIdent))
			}
			if elevatedpointtype.VerticalDatum != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedpointtype.VerticalDatum.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalDatum = %s", elevatedpointtypeIdent, targetIdent))
			}
			if elevatedpointtype.VerticalAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedpointtype.VerticalAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalAccuracy = %s", elevatedpointtypeIdent, targetIdent))
			}
			for _, elem := range elevatedpointtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevatedpointtypeIdent, elevatedpointtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, elevatedsurfacepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ElevatedSurfacePropertyTypes, stage.AixmfeaturesStage.ElevatedSurfacePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			elevatedsurfacepropertytypeIdent := elevatedsurfacepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedSurfacePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", elevatedsurfacepropertytypeIdent, __gong__toRawStringLiteral(elevatedsurfacepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedsurfacepropertytypeIdent, __gong__toRawStringLiteral(elevatedsurfacepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", elevatedsurfacepropertytypeIdent, __gong__toRawStringLiteral(elevatedsurfacepropertytype.NilReason)))
			if elevatedsurfacepropertytype.ElevatedSurface != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevatedsurfacepropertytype.ElevatedSurface.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedSurface = %s", elevatedsurfacepropertytypeIdent, targetIdent))
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
		for _, groundlightsystempropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.GroundLightSystemPropertyTypes, stage.AixmfeaturesStage.GroundLightSystemPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			groundlightsystempropertytypeIdent := groundlightsystempropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GroundLightSystemPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, lightelementpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LightElementPropertyTypes, stage.AixmfeaturesStage.LightElementPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			lightelementpropertytypeIdent := lightelementpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightElementPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", lightelementpropertytypeIdent, __gong__toRawStringLiteral(lightelementpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightelementpropertytypeIdent, __gong__toRawStringLiteral(lightelementpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", lightelementpropertytypeIdent, __gong__toRawStringLiteral(lightelementpropertytype.NilReason)))
			if lightelementpropertytype.LightElement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightelementpropertytype.LightElement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LightElement = %s", lightelementpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, lightelementstatuspropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LightElementStatusPropertyTypes, stage.AixmfeaturesStage.LightElementStatusPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			lightelementstatuspropertytypeIdent := lightelementstatuspropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightElementStatusPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", lightelementstatuspropertytypeIdent, __gong__toRawStringLiteral(lightelementstatuspropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightelementstatuspropertytypeIdent, __gong__toRawStringLiteral(lightelementstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", lightelementstatuspropertytypeIdent, __gong__toRawStringLiteral(lightelementstatuspropertytype.NilReason)))
			if lightelementstatuspropertytype.LightElementStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightelementstatuspropertytype.LightElementStatus.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LightElementStatus = %s", lightelementstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, lightelementstatustype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LightElementStatusTypes, stage.AixmfeaturesStage.LightElementStatusType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			lightelementstatustypeIdent := lightelementstatustype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightElementStatusType{Name: %s}).Stage(stage.AixmfeaturesStage)", lightelementstatustypeIdent, __gong__toRawStringLiteral(lightelementstatustype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightelementstatustypeIdent, __gong__toRawStringLiteral(lightelementstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", lightelementstatustypeIdent, __gong__toRawStringLiteral(lightelementstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", lightelementstatustypeIdent, __gong__toRawStringLiteral(lightelementstatustype.Description)))
			if lightelementstatustype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightelementstatustype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", lightelementstatustypeIdent, targetIdent))
			}
			for _, elem := range lightelementstatustype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", lightelementstatustypeIdent, lightelementstatustypeIdent, targetIdent))
			}
			for _, elem := range lightelementstatustype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", lightelementstatustypeIdent, lightelementstatustypeIdent, targetIdent))
			}
			for _, elem := range lightelementstatustype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", lightelementstatustypeIdent, lightelementstatustypeIdent, targetIdent))
			}
			for _, elem := range lightelementstatustype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", lightelementstatustypeIdent, lightelementstatustypeIdent, targetIdent))
			}
			if lightelementstatustype.Status != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightelementstatustype.Status.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Status = %s", lightelementstatustypeIdent, targetIdent))
			}
			for _, elem := range lightelementstatustype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", lightelementstatustypeIdent, lightelementstatustypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, lightelementtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LightElementTypes, stage.AixmfeaturesStage.LightElementType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			lightelementtypeIdent := lightelementtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightElementType{Name: %s}).Stage(stage.AixmfeaturesStage)", lightelementtypeIdent, __gong__toRawStringLiteral(lightelementtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightelementtypeIdent, __gong__toRawStringLiteral(lightelementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", lightelementtypeIdent, __gong__toRawStringLiteral(lightelementtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", lightelementtypeIdent, __gong__toRawStringLiteral(lightelementtype.Description)))
			if lightelementtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightelementtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", lightelementtypeIdent, targetIdent))
			}
			for _, elem := range lightelementtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", lightelementtypeIdent, lightelementtypeIdent, targetIdent))
			}
			if lightelementtype.Colour != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightelementtype.Colour.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Colour = %s", lightelementtypeIdent, targetIdent))
			}
			if lightelementtype.IntensityLevel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightelementtype.IntensityLevel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntensityLevel = %s", lightelementtypeIdent, targetIdent))
			}
			if lightelementtype.Intensity != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightelementtype.Intensity.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Intensity = %s", lightelementtypeIdent, targetIdent))
			}
			if lightelementtype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightelementtype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", lightelementtypeIdent, targetIdent))
			}
			if lightelementtype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightelementtype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", lightelementtypeIdent, targetIdent))
			}
			for _, elem := range lightelementtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", lightelementtypeIdent, lightelementtypeIdent, targetIdent))
			}
			for _, elem := range lightelementtype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", lightelementtypeIdent, lightelementtypeIdent, targetIdent))
			}
			for _, elem := range lightelementtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", lightelementtypeIdent, lightelementtypeIdent, targetIdent))
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
		for _, markerbeaconpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.MarkerBeaconPropertyTypes, stage.AixmfeaturesStage.MarkerBeaconPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			markerbeaconpropertytypeIdent := markerbeaconpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MarkerBeaconPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", markerbeaconpropertytypeIdent, markerbeaconpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Actuate)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, navaidequipmentpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidEquipmentPropertyTypes, stage.AixmfeaturesStage.NavaidEquipmentPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidequipmentpropertytypeIdent := navaidequipmentpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidEquipmentPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
		for _, obstacleareatimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ObstacleAreaTimeSlicePropertyTypes, stage.AixmfeaturesStage.ObstacleAreaTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			obstacleareatimeslicepropertytypeIdent := obstacleareatimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstacleAreaTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", obstacleareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", obstacleareatimeslicepropertytypeIdent, obstacleareatimeslicepropertytype.Owns))
			if obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstacleAreaTimeSlice = %s", obstacleareatimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, obstacleareatimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ObstacleAreaTimeSliceTypes, stage.AixmfeaturesStage.ObstacleAreaTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			obstacleareatimeslicetypeIdent := obstacleareatimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstacleAreaTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", obstacleareatimeslicetypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleareatimeslicetypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", obstacleareatimeslicetypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", obstacleareatimeslicetypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", obstacleareatimeslicetypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", obstacleareatimeslicetypeIdent, obstacleareatimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", obstacleareatimeslicetypeIdent, obstacleareatimeslicetype.CorrectionNumber))
			if obstacleareatimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range obstacleareatimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", obstacleareatimeslicetypeIdent, obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.ObstructionIdSurfaceCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicetype.ObstructionIdSurfaceCondition.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstructionIdSurfaceCondition = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.Reference_ownerOrganisation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicetype.Reference_ownerOrganisation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Reference_ownerOrganisation = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.Reference_ownerRunway != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicetype.Reference_ownerRunway.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Reference_ownerRunway = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.Reference_ownerAirport != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicetype.Reference_ownerAirport.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Reference_ownerAirport = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.SurfaceExtent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatimeslicetype.SurfaceExtent.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SurfaceExtent = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range obstacleareatimeslicetype.Obstacle {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Obstacle = append(%s.Obstacle, %s)", obstacleareatimeslicetypeIdent, obstacleareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range obstacleareatimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", obstacleareatimeslicetypeIdent, obstacleareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range obstacleareatimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", obstacleareatimeslicetypeIdent, obstacleareatimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, obstacleareatype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ObstacleAreaTypes, stage.AixmfeaturesStage.ObstacleAreaType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			obstacleareatypeIdent := obstacleareatype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstacleAreaType{Name: %s}).Stage(stage.AixmfeaturesStage)", obstacleareatypeIdent, __gong__toRawStringLiteral(obstacleareatype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleareatypeIdent, __gong__toRawStringLiteral(obstacleareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", obstacleareatypeIdent, __gong__toRawStringLiteral(obstacleareatype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", obstacleareatypeIdent, __gong__toRawStringLiteral(obstacleareatype.Description)))
			if obstacleareatype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", obstacleareatypeIdent, targetIdent))
			}
			for _, elem := range obstacleareatype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", obstacleareatypeIdent, obstacleareatypeIdent, targetIdent))
			}
			if obstacleareatype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", obstacleareatypeIdent, targetIdent))
			}
			if obstacleareatype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleareatype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", obstacleareatypeIdent, targetIdent))
			}
			for _, elem := range obstacleareatype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", obstacleareatypeIdent, obstacleareatypeIdent, targetIdent))
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
		for _, passengerservicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PassengerServicePropertyTypes, stage.AixmfeaturesStage.PassengerServicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			passengerservicepropertytypeIdent := passengerservicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PassengerServicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", passengerservicepropertytypeIdent, passengerservicepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Actuate)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, pointtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PointTypes, stage.AixmfeaturesStage.PointType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			pointtypeIdent := pointtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PointType{Name: %s}).Stage(stage.AixmfeaturesStage)", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.SrsName)))
			if pointtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", pointtypeIdent, targetIdent))
			}
			for _, elem := range pointtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", pointtypeIdent, pointtypeIdent, targetIdent))
			}
			if pointtype.Pos != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointtype.Pos.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", pointtypeIdent, targetIdent))
			}
			if pointtype.HorizontalAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointtype.HorizontalAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalAccuracy = %s", pointtypeIdent, targetIdent))
			}
			for _, elem := range pointtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", pointtypeIdent, pointtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, runwaydirectionpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RunwayDirectionPropertyTypes, stage.AixmfeaturesStage.RunwayDirectionPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			runwaydirectionpropertytypeIdent := runwaydirectionpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RunwayDirectionPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, servicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ServicePropertyTypes, stage.AixmfeaturesStage.ServicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			servicepropertytypeIdent := servicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ServicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", servicepropertytypeIdent, servicepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Actuate)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialnavigationstationpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialNavigationStationPropertyTypes, stage.AixmfeaturesStage.SpecialNavigationStationPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialnavigationstationpropertytypeIdent := specialnavigationstationpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", specialnavigationstationpropertytypeIdent, specialnavigationstationpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Actuate)))
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
	if stage.AixmfeaturesStage != nil {
		for _, unitpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.UnitPropertyTypes, stage.AixmfeaturesStage.UnitPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			unitpropertytypeIdent := unitpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, verticalstructurelightingstatuspropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VerticalStructureLightingStatusPropertyTypes, stage.AixmfeaturesStage.VerticalStructureLightingStatusPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			verticalstructurelightingstatuspropertytypeIdent := verticalstructurelightingstatuspropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructureLightingStatusPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", verticalstructurelightingstatuspropertytypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatuspropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructurelightingstatuspropertytypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", verticalstructurelightingstatuspropertytypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatuspropertytype.NilReason)))
			if verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalStructureLightingStatus = %s", verticalstructurelightingstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, verticalstructurelightingstatustype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VerticalStructureLightingStatusTypes, stage.AixmfeaturesStage.VerticalStructureLightingStatusType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			verticalstructurelightingstatustypeIdent := verticalstructurelightingstatustype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructureLightingStatusType{Name: %s}).Stage(stage.AixmfeaturesStage)", verticalstructurelightingstatustypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatustype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructurelightingstatustypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", verticalstructurelightingstatustypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", verticalstructurelightingstatustypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatustype.Description)))
			if verticalstructurelightingstatustype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructurelightingstatustype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", verticalstructurelightingstatustypeIdent, targetIdent))
			}
			for _, elem := range verticalstructurelightingstatustype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", verticalstructurelightingstatustypeIdent, verticalstructurelightingstatustypeIdent, targetIdent))
			}
			for _, elem := range verticalstructurelightingstatustype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", verticalstructurelightingstatustypeIdent, verticalstructurelightingstatustypeIdent, targetIdent))
			}
			for _, elem := range verticalstructurelightingstatustype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", verticalstructurelightingstatustypeIdent, verticalstructurelightingstatustypeIdent, targetIdent))
			}
			for _, elem := range verticalstructurelightingstatustype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", verticalstructurelightingstatustypeIdent, verticalstructurelightingstatustypeIdent, targetIdent))
			}
			if verticalstructurelightingstatustype.Status != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructurelightingstatustype.Status.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Status = %s", verticalstructurelightingstatustypeIdent, targetIdent))
			}
			for _, elem := range verticalstructurelightingstatustype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", verticalstructurelightingstatustypeIdent, verticalstructurelightingstatustypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, verticalstructurepartpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VerticalStructurePartPropertyTypes, stage.AixmfeaturesStage.VerticalStructurePartPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			verticalstructurepartpropertytypeIdent := verticalstructurepartpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructurePartPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", verticalstructurepartpropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepartpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructurepartpropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepartpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", verticalstructurepartpropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepartpropertytype.NilReason)))
			if verticalstructurepartpropertytype.VerticalStructurePart != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructurepartpropertytype.VerticalStructurePart.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalStructurePart = %s", verticalstructurepartpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, verticalstructureparttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VerticalStructurePartTypes, stage.AixmfeaturesStage.VerticalStructurePartType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			verticalstructureparttypeIdent := verticalstructureparttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructurePartType{Name: %s}).Stage(stage.AixmfeaturesStage)", verticalstructureparttypeIdent, __gong__toRawStringLiteral(verticalstructureparttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructureparttypeIdent, __gong__toRawStringLiteral(verticalstructureparttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", verticalstructureparttypeIdent, __gong__toRawStringLiteral(verticalstructureparttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", verticalstructureparttypeIdent, __gong__toRawStringLiteral(verticalstructureparttype.Description)))
			if verticalstructureparttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", verticalstructureparttypeIdent, targetIdent))
			}
			for _, elem := range verticalstructureparttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
			for _, elem := range verticalstructureparttype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
			for _, elem := range verticalstructureparttype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
			for _, elem := range verticalstructureparttype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.VerticalExtent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.VerticalExtent.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalExtent = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.VerticalExtentAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.VerticalExtentAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalExtentAccuracy = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.ConstructionStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.ConstructionStatus.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConstructionStatus = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.MarkingPattern != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.MarkingPattern.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MarkingPattern = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.MarkingFirstColour != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.MarkingFirstColour.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MarkingFirstColour = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.MarkingSecondColour != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.MarkingSecondColour.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MarkingSecondColour = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.Frangible != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.Frangible.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frangible = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.VisibleMaterial != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.VisibleMaterial.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VisibleMaterial = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.HorizontalProjection_surfaceExtent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.HorizontalProjection_surfaceExtent.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalProjection_surfaceExtent = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.HorizontalProjection_linearExtent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.HorizontalProjection_linearExtent.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalProjection_linearExtent = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.HorizontalProjection_location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructureparttype.HorizontalProjection_location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalProjection_location = %s", verticalstructureparttypeIdent, targetIdent))
			}
			for _, elem := range verticalstructureparttype.Lighting {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lighting = append(%s.Lighting, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
			for _, elem := range verticalstructureparttype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, verticalstructurepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VerticalStructurePropertyTypes, stage.AixmfeaturesStage.VerticalStructurePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			verticalstructurepropertytypeIdent := verticalstructurepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructurePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, verticalstructuretimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VerticalStructureTimeSlicePropertyTypes, stage.AixmfeaturesStage.VerticalStructureTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			verticalstructuretimeslicepropertytypeIdent := verticalstructuretimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructureTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", verticalstructuretimeslicepropertytypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructuretimeslicepropertytypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", verticalstructuretimeslicepropertytypeIdent, verticalstructuretimeslicepropertytype.Owns))
			if verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalStructureTimeSlice = %s", verticalstructuretimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, verticalstructuretimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VerticalStructureTimeSliceTypes, stage.AixmfeaturesStage.VerticalStructureTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			verticalstructuretimeslicetypeIdent := verticalstructuretimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructureTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", verticalstructuretimeslicetypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructuretimeslicetypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", verticalstructuretimeslicetypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", verticalstructuretimeslicetypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", verticalstructuretimeslicetypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetype.CorrectionNumber))
			if verticalstructuretimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.Lighted != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.Lighted.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lighted = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.MarkingICAOStandard != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.MarkingICAOStandard.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MarkingICAOStandard = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.Group != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.Group.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.Length != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.Length.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Length = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.Width != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.Width.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Width = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.Radius != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.Radius.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Radius = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.LightingICAOStandard != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.LightingICAOStandard.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LightingICAOStandard = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.SynchronisedLighting != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.SynchronisedLighting.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SynchronisedLighting = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			if verticalstructuretimeslicetype.Marker != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretimeslicetype.Marker.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Marker = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.Part {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = append(%s.Part, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.HostedPassengerService {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HostedPassengerService = append(%s.HostedPassengerService, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.SupportedGroundLight {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SupportedGroundLight = append(%s.SupportedGroundLight, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.HostedNavaidEquipment {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HostedNavaidEquipment = append(%s.HostedNavaidEquipment, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.HostedSpecialNavStation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HostedSpecialNavStation = append(%s.HostedSpecialNavStation, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.HostedUnit {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HostedUnit = append(%s.HostedUnit, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.HostedOrganisation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HostedOrganisation = append(%s.HostedOrganisation, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.SupportedService {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SupportedService = append(%s.SupportedService, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.LightingAvailability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LightingAvailability = append(%s.LightingAvailability, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, verticalstructuretype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VerticalStructureTypes, stage.AixmfeaturesStage.VerticalStructureType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			verticalstructuretypeIdent := verticalstructuretype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructureType{Name: %s}).Stage(stage.AixmfeaturesStage)", verticalstructuretypeIdent, __gong__toRawStringLiteral(verticalstructuretype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructuretypeIdent, __gong__toRawStringLiteral(verticalstructuretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", verticalstructuretypeIdent, __gong__toRawStringLiteral(verticalstructuretype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", verticalstructuretypeIdent, __gong__toRawStringLiteral(verticalstructuretype.Description)))
			if verticalstructuretype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", verticalstructuretypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", verticalstructuretypeIdent, verticalstructuretypeIdent, targetIdent))
			}
			if verticalstructuretype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", verticalstructuretypeIdent, targetIdent))
			}
			if verticalstructuretype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := verticalstructuretype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", verticalstructuretypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", verticalstructuretypeIdent, verticalstructuretypeIdent, targetIdent))
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
			if basicmessagememberaixmpropertytype.VerticalStructure != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.VerticalStructure.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalStructure = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.ObstacleArea != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.ObstacleArea.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstacleArea = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.AeronauticalGroundLight != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.AeronauticalGroundLight.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AeronauticalGroundLight = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/xlink"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models\"", fmt.Sprintf("%q", modelsPackageName))
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
