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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models"
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
		for _, codeatcreportingtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeATCReportingTypes, stage.AixmdatatypesStage.CodeATCReportingType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeatcreportingtypeIdent := codeatcreportingtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeATCReportingType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeatcreportingtypeIdent, __gong__toRawStringLiteral(codeatcreportingtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeatcreportingtypeIdent, __gong__toRawStringLiteral(codeatcreportingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeatcreportingtypeIdent, __gong__toRawStringLiteral(codeatcreportingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeatcreportingtypeIdent, __gong__toRawStringLiteral(codeatcreportingtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeaerialrefuellingpointtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAerialRefuellingPointTypes, stage.AixmdatatypesStage.CodeAerialRefuellingPointType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaerialrefuellingpointtypeIdent := codeaerialrefuellingpointtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAerialRefuellingPointType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingpointtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingpointtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingpointtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeaerialrefuellingprefixtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAerialRefuellingPrefixTypes, stage.AixmdatatypesStage.CodeAerialRefuellingPrefixType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaerialrefuellingprefixtypeIdent := codeaerialrefuellingprefixtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAerialRefuellingPrefixType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaerialrefuellingprefixtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingprefixtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaerialrefuellingprefixtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingprefixtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaerialrefuellingprefixtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingprefixtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaerialrefuellingprefixtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingprefixtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeaerialrefuellingtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAerialRefuellingTypes, stage.AixmdatatypesStage.CodeAerialRefuellingType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaerialrefuellingtypeIdent := codeaerialrefuellingtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAerialRefuellingType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaerialrefuellingtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaerialrefuellingtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaerialrefuellingtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaerialrefuellingtypeIdent, __gong__toRawStringLiteral(codeaerialrefuellingtype.EnclosedText)))
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
		for _, codealtitudeadjustmenttype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAltitudeAdjustmentTypes, stage.AixmdatatypesStage.CodeAltitudeAdjustmentType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codealtitudeadjustmenttypeIdent := codealtitudeadjustmenttype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAltitudeAdjustmentType{Name: %s}).Stage(stage.AixmdatatypesStage)", codealtitudeadjustmenttypeIdent, __gong__toRawStringLiteral(codealtitudeadjustmenttype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codealtitudeadjustmenttypeIdent, __gong__toRawStringLiteral(codealtitudeadjustmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codealtitudeadjustmenttypeIdent, __gong__toRawStringLiteral(codealtitudeadjustmenttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codealtitudeadjustmenttypeIdent, __gong__toRawStringLiteral(codealtitudeadjustmenttype.EnclosedText)))
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
		for _, codecardinaldirectiontype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeCardinalDirectionTypes, stage.AixmdatatypesStage.CodeCardinalDirectionType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codecardinaldirectiontypeIdent := codecardinaldirectiontype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCardinalDirectionType{Name: %s}).Stage(stage.AixmdatatypesStage)", codecardinaldirectiontypeIdent, __gong__toRawStringLiteral(codecardinaldirectiontype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecardinaldirectiontypeIdent, __gong__toRawStringLiteral(codecardinaldirectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecardinaldirectiontypeIdent, __gong__toRawStringLiteral(codecardinaldirectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecardinaldirectiontypeIdent, __gong__toRawStringLiteral(codecardinaldirectiontype.EnclosedText)))
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
		for _, codecomparisontype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeComparisonTypes, stage.AixmdatatypesStage.CodeComparisonType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codecomparisontypeIdent := codecomparisontype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeComparisonType{Name: %s}).Stage(stage.AixmdatatypesStage)", codecomparisontypeIdent, __gong__toRawStringLiteral(codecomparisontype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecomparisontypeIdent, __gong__toRawStringLiteral(codecomparisontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecomparisontypeIdent, __gong__toRawStringLiteral(codecomparisontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecomparisontypeIdent, __gong__toRawStringLiteral(codecomparisontype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codecoursetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeCourseTypes, stage.AixmdatatypesStage.CodeCourseType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codecoursetypeIdent := codecoursetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCourseType{Name: %s}).Stage(stage.AixmdatatypesStage)", codecoursetypeIdent, __gong__toRawStringLiteral(codecoursetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecoursetypeIdent, __gong__toRawStringLiteral(codecoursetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecoursetypeIdent, __gong__toRawStringLiteral(codecoursetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecoursetypeIdent, __gong__toRawStringLiteral(codecoursetype.EnclosedText)))
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
		for _, codedirectionturntype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeDirectionTurnTypes, stage.AixmdatatypesStage.CodeDirectionTurnType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codedirectionturntypeIdent := codedirectionturntype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDirectionTurnType{Name: %s}).Stage(stage.AixmdatatypesStage)", codedirectionturntypeIdent, __gong__toRawStringLiteral(codedirectionturntype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedirectionturntypeIdent, __gong__toRawStringLiteral(codedirectionturntype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedirectionturntypeIdent, __gong__toRawStringLiteral(codedirectionturntype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedirectionturntypeIdent, __gong__toRawStringLiteral(codedirectionturntype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codedirectiontype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeDirectionTypes, stage.AixmdatatypesStage.CodeDirectionType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codedirectiontypeIdent := codedirectiontype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDirectionType{Name: %s}).Stage(stage.AixmdatatypesStage)", codedirectiontypeIdent, __gong__toRawStringLiteral(codedirectiontype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedirectiontypeIdent, __gong__toRawStringLiteral(codedirectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedirectiontypeIdent, __gong__toRawStringLiteral(codedirectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedirectiontypeIdent, __gong__toRawStringLiteral(codedirectiontype.EnclosedText)))
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
		for _, codeflightorigintype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFlightOriginTypes, stage.AixmdatatypesStage.CodeFlightOriginType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeflightorigintypeIdent := codeflightorigintype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightOriginType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeflightorigintypeIdent, __gong__toRawStringLiteral(codeflightorigintype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightorigintypeIdent, __gong__toRawStringLiteral(codeflightorigintype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightorigintypeIdent, __gong__toRawStringLiteral(codeflightorigintype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightorigintypeIdent, __gong__toRawStringLiteral(codeflightorigintype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeflightpurposetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFlightPurposeTypes, stage.AixmdatatypesStage.CodeFlightPurposeType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeflightpurposetypeIdent := codeflightpurposetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightPurposeType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeflightpurposetypeIdent, __gong__toRawStringLiteral(codeflightpurposetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightpurposetypeIdent, __gong__toRawStringLiteral(codeflightpurposetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightpurposetypeIdent, __gong__toRawStringLiteral(codeflightpurposetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightpurposetypeIdent, __gong__toRawStringLiteral(codeflightpurposetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeflightrestrictiondesignatortype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFlightRestrictionDesignatorTypes, stage.AixmdatatypesStage.CodeFlightRestrictionDesignatorType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeflightrestrictiondesignatortypeIdent := codeflightrestrictiondesignatortype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightRestrictionDesignatorType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeflightrestrictiondesignatortypeIdent, __gong__toRawStringLiteral(codeflightrestrictiondesignatortype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightrestrictiondesignatortypeIdent, __gong__toRawStringLiteral(codeflightrestrictiondesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightrestrictiondesignatortypeIdent, __gong__toRawStringLiteral(codeflightrestrictiondesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightrestrictiondesignatortypeIdent, __gong__toRawStringLiteral(codeflightrestrictiondesignatortype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeflightrestrictiontype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFlightRestrictionTypes, stage.AixmdatatypesStage.CodeFlightRestrictionType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeflightrestrictiontypeIdent := codeflightrestrictiontype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightRestrictionType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeflightrestrictiontypeIdent, __gong__toRawStringLiteral(codeflightrestrictiontype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightrestrictiontypeIdent, __gong__toRawStringLiteral(codeflightrestrictiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightrestrictiontypeIdent, __gong__toRawStringLiteral(codeflightrestrictiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightrestrictiontypeIdent, __gong__toRawStringLiteral(codeflightrestrictiontype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeflightruletype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFlightRuleTypes, stage.AixmdatatypesStage.CodeFlightRuleType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeflightruletypeIdent := codeflightruletype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightRuleType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeflightruletypeIdent, __gong__toRawStringLiteral(codeflightruletype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightruletypeIdent, __gong__toRawStringLiteral(codeflightruletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightruletypeIdent, __gong__toRawStringLiteral(codeflightruletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightruletypeIdent, __gong__toRawStringLiteral(codeflightruletype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeflightstatustype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFlightStatusTypes, stage.AixmdatatypesStage.CodeFlightStatusType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeflightstatustypeIdent := codeflightstatustype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightStatusType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeflightstatustypeIdent, __gong__toRawStringLiteral(codeflightstatustype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightstatustypeIdent, __gong__toRawStringLiteral(codeflightstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightstatustypeIdent, __gong__toRawStringLiteral(codeflightstatustype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightstatustypeIdent, __gong__toRawStringLiteral(codeflightstatustype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeflighttype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFlightTypes, stage.AixmdatatypesStage.CodeFlightType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeflighttypeIdent := codeflighttype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeflighttypeIdent, __gong__toRawStringLiteral(codeflighttype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflighttypeIdent, __gong__toRawStringLiteral(codeflighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflighttypeIdent, __gong__toRawStringLiteral(codeflighttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflighttypeIdent, __gong__toRawStringLiteral(codeflighttype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeflowconditionoperationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFlowConditionOperationTypes, stage.AixmdatatypesStage.CodeFlowConditionOperationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeflowconditionoperationtypeIdent := codeflowconditionoperationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlowConditionOperationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeflowconditionoperationtypeIdent, __gong__toRawStringLiteral(codeflowconditionoperationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflowconditionoperationtypeIdent, __gong__toRawStringLiteral(codeflowconditionoperationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflowconditionoperationtypeIdent, __gong__toRawStringLiteral(codeflowconditionoperationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflowconditionoperationtypeIdent, __gong__toRawStringLiteral(codeflowconditionoperationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codefreeflighttype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFreeFlightTypes, stage.AixmdatatypesStage.CodeFreeFlightType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codefreeflighttypeIdent := codefreeflighttype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFreeFlightType{Name: %s}).Stage(stage.AixmdatatypesStage)", codefreeflighttypeIdent, __gong__toRawStringLiteral(codefreeflighttype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefreeflighttypeIdent, __gong__toRawStringLiteral(codefreeflighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefreeflighttypeIdent, __gong__toRawStringLiteral(codefreeflighttype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefreeflighttypeIdent, __gong__toRawStringLiteral(codefreeflighttype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeleveltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeLevelTypes, stage.AixmdatatypesStage.CodeLevelType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeleveltypeIdent := codeleveltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLevelType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeleveltypeIdent, __gong__toRawStringLiteral(codeleveltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeleveltypeIdent, __gong__toRawStringLiteral(codeleveltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeleveltypeIdent, __gong__toRawStringLiteral(codeleveltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeleveltypeIdent, __gong__toRawStringLiteral(codeleveltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codelocationqualifiertype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeLocationQualifierTypes, stage.AixmdatatypesStage.CodeLocationQualifierType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codelocationqualifiertypeIdent := codelocationqualifiertype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLocationQualifierType{Name: %s}).Stage(stage.AixmdatatypesStage)", codelocationqualifiertypeIdent, __gong__toRawStringLiteral(codelocationqualifiertype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelocationqualifiertypeIdent, __gong__toRawStringLiteral(codelocationqualifiertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelocationqualifiertypeIdent, __gong__toRawStringLiteral(codelocationqualifiertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelocationqualifiertypeIdent, __gong__toRawStringLiteral(codelocationqualifiertype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codemeteoconditionstype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeMeteoConditionsTypes, stage.AixmdatatypesStage.CodeMeteoConditionsType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codemeteoconditionstypeIdent := codemeteoconditionstype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMeteoConditionsType{Name: %s}).Stage(stage.AixmdatatypesStage)", codemeteoconditionstypeIdent, __gong__toRawStringLiteral(codemeteoconditionstype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemeteoconditionstypeIdent, __gong__toRawStringLiteral(codemeteoconditionstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemeteoconditionstypeIdent, __gong__toRawStringLiteral(codemeteoconditionstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemeteoconditionstypeIdent, __gong__toRawStringLiteral(codemeteoconditionstype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codemilitaryroutepointtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeMilitaryRoutePointTypes, stage.AixmdatatypesStage.CodeMilitaryRoutePointType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codemilitaryroutepointtypeIdent := codemilitaryroutepointtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMilitaryRoutePointType{Name: %s}).Stage(stage.AixmdatatypesStage)", codemilitaryroutepointtypeIdent, __gong__toRawStringLiteral(codemilitaryroutepointtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemilitaryroutepointtypeIdent, __gong__toRawStringLiteral(codemilitaryroutepointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemilitaryroutepointtypeIdent, __gong__toRawStringLiteral(codemilitaryroutepointtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemilitaryroutepointtypeIdent, __gong__toRawStringLiteral(codemilitaryroutepointtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codemilitarystatustype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeMilitaryStatusTypes, stage.AixmdatatypesStage.CodeMilitaryStatusType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codemilitarystatustypeIdent := codemilitarystatustype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMilitaryStatusType{Name: %s}).Stage(stage.AixmdatatypesStage)", codemilitarystatustypeIdent, __gong__toRawStringLiteral(codemilitarystatustype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemilitarystatustypeIdent, __gong__toRawStringLiteral(codemilitarystatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemilitarystatustypeIdent, __gong__toRawStringLiteral(codemilitarystatustype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemilitarystatustypeIdent, __gong__toRawStringLiteral(codemilitarystatustype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codemilitarytrainingtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeMilitaryTrainingTypes, stage.AixmdatatypesStage.CodeMilitaryTrainingType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codemilitarytrainingtypeIdent := codemilitarytrainingtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMilitaryTrainingType{Name: %s}).Stage(stage.AixmdatatypesStage)", codemilitarytrainingtypeIdent, __gong__toRawStringLiteral(codemilitarytrainingtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemilitarytrainingtypeIdent, __gong__toRawStringLiteral(codemilitarytrainingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemilitarytrainingtypeIdent, __gong__toRawStringLiteral(codemilitarytrainingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemilitarytrainingtypeIdent, __gong__toRawStringLiteral(codemilitarytrainingtype.EnclosedText)))
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
		for _, codeobstructionidsurfacezonetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeObstructionIdSurfaceZoneTypes, stage.AixmdatatypesStage.CodeObstructionIdSurfaceZoneType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeobstructionidsurfacezonetypeIdent := codeobstructionidsurfacezonetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeObstructionIdSurfaceZoneType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeobstructionidsurfacezonetypeIdent, __gong__toRawStringLiteral(codeobstructionidsurfacezonetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeobstructionidsurfacezonetypeIdent, __gong__toRawStringLiteral(codeobstructionidsurfacezonetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeobstructionidsurfacezonetypeIdent, __gong__toRawStringLiteral(codeobstructionidsurfacezonetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeobstructionidsurfacezonetypeIdent, __gong__toRawStringLiteral(codeobstructionidsurfacezonetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codernptype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRNPTypes, stage.AixmdatatypesStage.CodeRNPType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codernptypeIdent := codernptype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRNPType{Name: %s}).Stage(stage.AixmdatatypesStage)", codernptypeIdent, __gong__toRawStringLiteral(codernptype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codernptypeIdent, __gong__toRawStringLiteral(codernptype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codernptypeIdent, __gong__toRawStringLiteral(codernptype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codernptypeIdent, __gong__toRawStringLiteral(codernptype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codervsmpointroletype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRVSMPointRoleTypes, stage.AixmdatatypesStage.CodeRVSMPointRoleType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codervsmpointroletypeIdent := codervsmpointroletype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRVSMPointRoleType{Name: %s}).Stage(stage.AixmdatatypesStage)", codervsmpointroletypeIdent, __gong__toRawStringLiteral(codervsmpointroletype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codervsmpointroletypeIdent, __gong__toRawStringLiteral(codervsmpointroletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codervsmpointroletypeIdent, __gong__toRawStringLiteral(codervsmpointroletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codervsmpointroletypeIdent, __gong__toRawStringLiteral(codervsmpointroletype.EnclosedText)))
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
		for _, codereferenceroletype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeReferenceRoleTypes, stage.AixmdatatypesStage.CodeReferenceRoleType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codereferenceroletypeIdent := codereferenceroletype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeReferenceRoleType{Name: %s}).Stage(stage.AixmdatatypesStage)", codereferenceroletypeIdent, __gong__toRawStringLiteral(codereferenceroletype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codereferenceroletypeIdent, __gong__toRawStringLiteral(codereferenceroletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codereferenceroletypeIdent, __gong__toRawStringLiteral(codereferenceroletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codereferenceroletypeIdent, __gong__toRawStringLiteral(codereferenceroletype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderouteavailabilitytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRouteAvailabilityTypes, stage.AixmdatatypesStage.CodeRouteAvailabilityType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderouteavailabilitytypeIdent := coderouteavailabilitytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteAvailabilityType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderouteavailabilitytypeIdent, __gong__toRawStringLiteral(coderouteavailabilitytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderouteavailabilitytypeIdent, __gong__toRawStringLiteral(coderouteavailabilitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderouteavailabilitytypeIdent, __gong__toRawStringLiteral(coderouteavailabilitytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderouteavailabilitytypeIdent, __gong__toRawStringLiteral(coderouteavailabilitytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderoutedesignatorlettertype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRouteDesignatorLetterTypes, stage.AixmdatatypesStage.CodeRouteDesignatorLetterType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderoutedesignatorlettertypeIdent := coderoutedesignatorlettertype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteDesignatorLetterType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderoutedesignatorlettertypeIdent, __gong__toRawStringLiteral(coderoutedesignatorlettertype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutedesignatorlettertypeIdent, __gong__toRawStringLiteral(coderoutedesignatorlettertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutedesignatorlettertypeIdent, __gong__toRawStringLiteral(coderoutedesignatorlettertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutedesignatorlettertypeIdent, __gong__toRawStringLiteral(coderoutedesignatorlettertype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderoutedesignatorprefixtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRouteDesignatorPrefixTypes, stage.AixmdatatypesStage.CodeRouteDesignatorPrefixType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderoutedesignatorprefixtypeIdent := coderoutedesignatorprefixtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteDesignatorPrefixType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderoutedesignatorprefixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorprefixtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutedesignatorprefixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorprefixtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutedesignatorprefixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorprefixtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutedesignatorprefixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorprefixtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderoutedesignatorsuffixtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRouteDesignatorSuffixTypes, stage.AixmdatatypesStage.CodeRouteDesignatorSuffixType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderoutedesignatorsuffixtypeIdent := coderoutedesignatorsuffixtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteDesignatorSuffixType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderoutedesignatorsuffixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorsuffixtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutedesignatorsuffixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorsuffixtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutedesignatorsuffixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorsuffixtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutedesignatorsuffixtypeIdent, __gong__toRawStringLiteral(coderoutedesignatorsuffixtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderoutenavigationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRouteNavigationTypes, stage.AixmdatatypesStage.CodeRouteNavigationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderoutenavigationtypeIdent := coderoutenavigationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteNavigationType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderoutenavigationtypeIdent, __gong__toRawStringLiteral(coderoutenavigationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutenavigationtypeIdent, __gong__toRawStringLiteral(coderoutenavigationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutenavigationtypeIdent, __gong__toRawStringLiteral(coderoutenavigationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutenavigationtypeIdent, __gong__toRawStringLiteral(coderoutenavigationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderouteorigintype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRouteOriginTypes, stage.AixmdatatypesStage.CodeRouteOriginType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderouteorigintypeIdent := coderouteorigintype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteOriginType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderouteorigintypeIdent, __gong__toRawStringLiteral(coderouteorigintype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderouteorigintypeIdent, __gong__toRawStringLiteral(coderouteorigintype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderouteorigintypeIdent, __gong__toRawStringLiteral(coderouteorigintype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderouteorigintypeIdent, __gong__toRawStringLiteral(coderouteorigintype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderoutesegmentpathtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRouteSegmentPathTypes, stage.AixmdatatypesStage.CodeRouteSegmentPathType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderoutesegmentpathtypeIdent := coderoutesegmentpathtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteSegmentPathType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderoutesegmentpathtypeIdent, __gong__toRawStringLiteral(coderoutesegmentpathtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutesegmentpathtypeIdent, __gong__toRawStringLiteral(coderoutesegmentpathtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutesegmentpathtypeIdent, __gong__toRawStringLiteral(coderoutesegmentpathtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutesegmentpathtypeIdent, __gong__toRawStringLiteral(coderoutesegmentpathtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderoutetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRouteTypes, stage.AixmdatatypesStage.CodeRouteType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderoutetypeIdent := coderoutetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRouteType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderoutetypeIdent, __gong__toRawStringLiteral(coderoutetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderoutetypeIdent, __gong__toRawStringLiteral(coderoutetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderoutetypeIdent, __gong__toRawStringLiteral(coderoutetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderoutetypeIdent, __gong__toRawStringLiteral(coderoutetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderuleproceduretitletype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRuleProcedureTitleTypes, stage.AixmdatatypesStage.CodeRuleProcedureTitleType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderuleproceduretitletypeIdent := coderuleproceduretitletype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRuleProcedureTitleType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderuleproceduretitletypeIdent, __gong__toRawStringLiteral(coderuleproceduretitletype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderuleproceduretitletypeIdent, __gong__toRawStringLiteral(coderuleproceduretitletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderuleproceduretitletypeIdent, __gong__toRawStringLiteral(coderuleproceduretitletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderuleproceduretitletypeIdent, __gong__toRawStringLiteral(coderuleproceduretitletype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderuleproceduretype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRuleProcedureTypes, stage.AixmdatatypesStage.CodeRuleProcedureType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderuleproceduretypeIdent := coderuleproceduretype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRuleProcedureType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderuleproceduretypeIdent, __gong__toRawStringLiteral(coderuleproceduretype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderuleproceduretypeIdent, __gong__toRawStringLiteral(coderuleproceduretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderuleproceduretypeIdent, __gong__toRawStringLiteral(coderuleproceduretype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderuleproceduretypeIdent, __gong__toRawStringLiteral(coderuleproceduretype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codesidetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeSideTypes, stage.AixmdatatypesStage.CodeSideType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codesidetypeIdent := codesidetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSideType{Name: %s}).Stage(stage.AixmdatatypesStage)", codesidetypeIdent, __gong__toRawStringLiteral(codesidetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesidetypeIdent, __gong__toRawStringLiteral(codesidetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesidetypeIdent, __gong__toRawStringLiteral(codesidetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesidetypeIdent, __gong__toRawStringLiteral(codesidetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codespeedreferencetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeSpeedReferenceTypes, stage.AixmdatatypesStage.CodeSpeedReferenceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codespeedreferencetypeIdent := codespeedreferencetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSpeedReferenceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codespeedreferencetypeIdent, __gong__toRawStringLiteral(codespeedreferencetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codespeedreferencetypeIdent, __gong__toRawStringLiteral(codespeedreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codespeedreferencetypeIdent, __gong__toRawStringLiteral(codespeedreferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codespeedreferencetypeIdent, __gong__toRawStringLiteral(codespeedreferencetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codetacanchanneltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeTACANChannelTypes, stage.AixmdatatypesStage.CodeTACANChannelType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codetacanchanneltypeIdent := codetacanchanneltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTACANChannelType{Name: %s}).Stage(stage.AixmdatatypesStage)", codetacanchanneltypeIdent, __gong__toRawStringLiteral(codetacanchanneltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetacanchanneltypeIdent, __gong__toRawStringLiteral(codetacanchanneltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetacanchanneltypeIdent, __gong__toRawStringLiteral(codetacanchanneltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetacanchanneltypeIdent, __gong__toRawStringLiteral(codetacanchanneltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codetelecomnetworktype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeTelecomNetworkTypes, stage.AixmdatatypesStage.CodeTelecomNetworkType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codetelecomnetworktypeIdent := codetelecomnetworktype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeTelecomNetworkType{Name: %s}).Stage(stage.AixmdatatypesStage)", codetelecomnetworktypeIdent, __gong__toRawStringLiteral(codetelecomnetworktype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetelecomnetworktypeIdent, __gong__toRawStringLiteral(codetelecomnetworktype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codetelecomnetworktypeIdent, __gong__toRawStringLiteral(codetelecomnetworktype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetelecomnetworktypeIdent, __gong__toRawStringLiteral(codetelecomnetworktype.EnclosedText)))
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
		for _, codeupperalphatype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeUpperAlphaTypes, stage.AixmdatatypesStage.CodeUpperAlphaType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeupperalphatypeIdent := codeupperalphatype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeUpperAlphaType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeupperalphatypeIdent, __gong__toRawStringLiteral(codeupperalphatype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeupperalphatypeIdent, __gong__toRawStringLiteral(codeupperalphatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeupperalphatypeIdent, __gong__toRawStringLiteral(codeupperalphatype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeupperalphatypeIdent, __gong__toRawStringLiteral(codeupperalphatype.EnclosedText)))
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
		for _, textaddresstype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.TextAddressTypes, stage.AixmdatatypesStage.TextAddressType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			textaddresstypeIdent := textaddresstype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextAddressType{Name: %s}).Stage(stage.AixmdatatypesStage)", textaddresstypeIdent, __gong__toRawStringLiteral(textaddresstype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textaddresstypeIdent, __gong__toRawStringLiteral(textaddresstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textaddresstypeIdent, __gong__toRawStringLiteral(textaddresstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textaddresstypeIdent, __gong__toRawStringLiteral(textaddresstype.EnclosedText)))
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
		for _, textinstructiontype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.TextInstructionTypes, stage.AixmdatatypesStage.TextInstructionType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			textinstructiontypeIdent := textinstructiontype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextInstructionType{Name: %s}).Stage(stage.AixmdatatypesStage)", textinstructiontypeIdent, __gong__toRawStringLiteral(textinstructiontype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textinstructiontypeIdent, __gong__toRawStringLiteral(textinstructiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textinstructiontypeIdent, __gong__toRawStringLiteral(textinstructiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textinstructiontypeIdent, __gong__toRawStringLiteral(textinstructiontype.EnclosedText)))
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
		for _, textphonetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.TextPhoneTypes, stage.AixmdatatypesStage.TextPhoneType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			textphonetypeIdent := textphonetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.TextPhoneType{Name: %s}).Stage(stage.AixmdatatypesStage)", textphonetypeIdent, __gong__toRawStringLiteral(textphonetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textphonetypeIdent, __gong__toRawStringLiteral(textphonetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", textphonetypeIdent, __gong__toRawStringLiteral(textphonetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", textphonetypeIdent, __gong__toRawStringLiteral(textphonetype.EnclosedText)))
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
		for _, valangletype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValAngleTypes, stage.AixmdatatypesStage.ValAngleType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valangletypeIdent := valangletype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValAngleType{Name: %s}).Stage(stage.AixmdatatypesStage)", valangletypeIdent, __gong__toRawStringLiteral(valangletype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valangletypeIdent, __gong__toRawStringLiteral(valangletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valangletypeIdent, __gong__toRawStringLiteral(valangletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valangletypeIdent, __gong__toRawStringLiteral(valangletype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, valbearingtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValBearingTypes, stage.AixmdatatypesStage.ValBearingType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valbearingtypeIdent := valbearingtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValBearingType{Name: %s}).Stage(stage.AixmdatatypesStage)", valbearingtypeIdent, __gong__toRawStringLiteral(valbearingtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valbearingtypeIdent, __gong__toRawStringLiteral(valbearingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valbearingtypeIdent, __gong__toRawStringLiteral(valbearingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valbearingtypeIdent, __gong__toRawStringLiteral(valbearingtype.EnclosedText)))
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
		for _, valslopetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValSlopeTypes, stage.AixmdatatypesStage.ValSlopeType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valslopetypeIdent := valslopetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValSlopeType{Name: %s}).Stage(stage.AixmdatatypesStage)", valslopetypeIdent, __gong__toRawStringLiteral(valslopetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valslopetypeIdent, __gong__toRawStringLiteral(valslopetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valslopetypeIdent, __gong__toRawStringLiteral(valslopetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valslopetypeIdent, __gong__toRawStringLiteral(valslopetype.EnclosedText)))
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
	if stage.AixmdatatypesStage != nil {
		for _, xhtmltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.XHTMLTypes, stage.AixmdatatypesStage.XHTMLType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			xhtmltypeIdent := xhtmltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.XHTMLType{Name: %s}).Stage(stage.AixmdatatypesStage)", xhtmltypeIdent, __gong__toRawStringLiteral(xhtmltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xhtmltypeIdent, __gong__toRawStringLiteral(xhtmltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", xhtmltypeIdent, __gong__toRawStringLiteral(xhtmltype.NilReason)))
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
			for _, elem := range a_extension.AbstractAerialRefuellingExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAerialRefuellingExtension = append(%s.AbstractAerialRefuellingExtension, %s)", a_extensionIdent, a_extensionIdent, targetIdent))
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
			for _, elem := range a_extension_1.AbstractAerialRefuellingAnchorExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAerialRefuellingAnchorExtension = append(%s.AbstractAerialRefuellingAnchorExtension, %s)", a_extension_1Ident, a_extension_1Ident, targetIdent))
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
			for _, elem := range a_extension_10.AbstractPointReferenceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPointReferenceExtension = append(%s.AbstractPointReferenceExtension, %s)", a_extension_10Ident, a_extension_10Ident, targetIdent))
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
			for _, elem := range a_extension_11.AbstractLinguisticNoteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractLinguisticNoteExtension = append(%s.AbstractLinguisticNoteExtension, %s)", a_extension_11Ident, a_extension_11Ident, targetIdent))
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
			for _, elem := range a_extension_12.AbstractNoteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNoteExtension = append(%s.AbstractNoteExtension, %s)", a_extension_12Ident, a_extension_12Ident, targetIdent))
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
			for _, elem := range a_extension_13.AbstractChangeOverPointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractChangeOverPointExtension = append(%s.AbstractChangeOverPointExtension, %s)", a_extension_13Ident, a_extension_13Ident, targetIdent))
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
			for _, elem := range a_extension_14.AbstractRouteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRouteExtension = append(%s.AbstractRouteExtension, %s)", a_extension_14Ident, a_extension_14Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_15 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_15s, stage.AixmfeaturesStage.A_extension_15_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_15Ident := a_extension_15.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_15{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_15Ident, __gong__toRawStringLiteral(a_extension_15.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_15Ident, __gong__toRawStringLiteral(a_extension_15.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_15Ident, a_extension_15.Owns))
			for _, elem := range a_extension_15.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_15Ident, a_extension_15Ident, targetIdent))
			}
			for _, elem := range a_extension_15.AbstractRouteAvailabilityExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRouteAvailabilityExtension = append(%s.AbstractRouteAvailabilityExtension, %s)", a_extension_15Ident, a_extension_15Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_16 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_16s, stage.AixmfeaturesStage.A_extension_16_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_16Ident := a_extension_16.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_16{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_16Ident, __gong__toRawStringLiteral(a_extension_16.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_16Ident, __gong__toRawStringLiteral(a_extension_16.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_16Ident, a_extension_16.Owns))
			for _, elem := range a_extension_16.AbstractRouteDMEExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRouteDMEExtension = append(%s.AbstractRouteDMEExtension, %s)", a_extension_16Ident, a_extension_16Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_17 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_17s, stage.AixmfeaturesStage.A_extension_17_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_17Ident := a_extension_17.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_17{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_17Ident, __gong__toRawStringLiteral(a_extension_17.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_17Ident, __gong__toRawStringLiteral(a_extension_17.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_17Ident, a_extension_17.Owns))
			for _, elem := range a_extension_17.AbstractRoutePortionExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRoutePortionExtension = append(%s.AbstractRoutePortionExtension, %s)", a_extension_17Ident, a_extension_17Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_18 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_18s, stage.AixmfeaturesStage.A_extension_18_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_18Ident := a_extension_18.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_18{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_18Ident, __gong__toRawStringLiteral(a_extension_18.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_18Ident, __gong__toRawStringLiteral(a_extension_18.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_18Ident, a_extension_18.Owns))
			for _, elem := range a_extension_18.AbstractRouteSegmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRouteSegmentExtension = append(%s.AbstractRouteSegmentExtension, %s)", a_extension_18Ident, a_extension_18Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_19 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_19s, stage.AixmfeaturesStage.A_extension_19_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_19Ident := a_extension_19.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_19{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_19Ident, __gong__toRawStringLiteral(a_extension_19.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_19Ident, __gong__toRawStringLiteral(a_extension_19.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_19Ident, a_extension_19.Owns))
			for _, elem := range a_extension_19.AbstractAirspaceBorderCrossingExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirspaceBorderCrossingExtension = append(%s.AbstractAirspaceBorderCrossingExtension, %s)", a_extension_19Ident, a_extension_19Ident, targetIdent))
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
			for _, elem := range a_extension_2.AbstractSegmentPointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSegmentPointExtension = append(%s.AbstractSegmentPointExtension, %s)", a_extension_2Ident, a_extension_2Ident, targetIdent))
			}
			for _, elem := range a_extension_2.AbstractAerialRefuellingPointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAerialRefuellingPointExtension = append(%s.AbstractAerialRefuellingPointExtension, %s)", a_extension_2Ident, a_extension_2Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_20 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_20s, stage.AixmfeaturesStage.A_extension_20_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_20Ident := a_extension_20.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_20{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_20Ident, __gong__toRawStringLiteral(a_extension_20.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_20Ident, __gong__toRawStringLiteral(a_extension_20.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_20Ident, a_extension_20.Owns))
			for _, elem := range a_extension_20.AbstractDirectFlightExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractDirectFlightExtension = append(%s.AbstractDirectFlightExtension, %s)", a_extension_20Ident, a_extension_20Ident, targetIdent))
			}
			for _, elem := range a_extension_20.AbstractDirectFlightClassExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractDirectFlightClassExtension = append(%s.AbstractDirectFlightClassExtension, %s)", a_extension_20Ident, a_extension_20Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_21 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_21s, stage.AixmfeaturesStage.A_extension_21_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_21Ident := a_extension_21.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_21{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_21Ident, __gong__toRawStringLiteral(a_extension_21.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_21Ident, __gong__toRawStringLiteral(a_extension_21.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_21Ident, a_extension_21.Owns))
			for _, elem := range a_extension_21.AbstractDirectFlightExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractDirectFlightExtension = append(%s.AbstractDirectFlightExtension, %s)", a_extension_21Ident, a_extension_21Ident, targetIdent))
			}
			for _, elem := range a_extension_21.AbstractDirectFlightSegmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractDirectFlightSegmentExtension = append(%s.AbstractDirectFlightSegmentExtension, %s)", a_extension_21Ident, a_extension_21Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_22 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_22s, stage.AixmfeaturesStage.A_extension_22_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_22Ident := a_extension_22.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_22{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_22Ident, __gong__toRawStringLiteral(a_extension_22.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_22Ident, __gong__toRawStringLiteral(a_extension_22.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_22Ident, a_extension_22.Owns))
			for _, elem := range a_extension_22.AbstractFlightConditionCircumstanceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractFlightConditionCircumstanceExtension = append(%s.AbstractFlightConditionCircumstanceExtension, %s)", a_extension_22Ident, a_extension_22Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_23 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_23s, stage.AixmfeaturesStage.A_extension_23_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_23Ident := a_extension_23.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_23{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_23Ident, __gong__toRawStringLiteral(a_extension_23.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_23Ident, __gong__toRawStringLiteral(a_extension_23.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_23Ident, a_extension_23.Owns))
			for _, elem := range a_extension_23.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_23Ident, a_extension_23Ident, targetIdent))
			}
			for _, elem := range a_extension_23.AbstractFlightConditionCombinationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractFlightConditionCombinationExtension = append(%s.AbstractFlightConditionCombinationExtension, %s)", a_extension_23Ident, a_extension_23Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_24 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_24s, stage.AixmfeaturesStage.A_extension_24_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_24Ident := a_extension_24.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_24{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_24Ident, __gong__toRawStringLiteral(a_extension_24.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_24Ident, __gong__toRawStringLiteral(a_extension_24.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_24Ident, a_extension_24.Owns))
			for _, elem := range a_extension_24.AbstractFlightConditionElementExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractFlightConditionElementExtension = append(%s.AbstractFlightConditionElementExtension, %s)", a_extension_24Ident, a_extension_24Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_25 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_25s, stage.AixmfeaturesStage.A_extension_25_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_25Ident := a_extension_25.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_25{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_25Ident, __gong__toRawStringLiteral(a_extension_25.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_25Ident, __gong__toRawStringLiteral(a_extension_25.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_25Ident, a_extension_25.Owns))
			for _, elem := range a_extension_25.AbstractFlightRestrictionExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractFlightRestrictionExtension = append(%s.AbstractFlightRestrictionExtension, %s)", a_extension_25Ident, a_extension_25Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_26 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_26s, stage.AixmfeaturesStage.A_extension_26_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_26Ident := a_extension_26.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_26{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_26Ident, __gong__toRawStringLiteral(a_extension_26.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_26Ident, __gong__toRawStringLiteral(a_extension_26.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_26Ident, a_extension_26.Owns))
			for _, elem := range a_extension_26.AbstractFlightRestrictionLevelExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractFlightRestrictionLevelExtension = append(%s.AbstractFlightRestrictionLevelExtension, %s)", a_extension_26Ident, a_extension_26Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_27 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_27s, stage.AixmfeaturesStage.A_extension_27_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_27Ident := a_extension_27.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_27{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_27Ident, __gong__toRawStringLiteral(a_extension_27.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_27Ident, __gong__toRawStringLiteral(a_extension_27.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_27Ident, a_extension_27.Owns))
			for _, elem := range a_extension_27.AbstractFlightRestrictionRouteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractFlightRestrictionRouteExtension = append(%s.AbstractFlightRestrictionRouteExtension, %s)", a_extension_27Ident, a_extension_27Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_28 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_28s, stage.AixmfeaturesStage.A_extension_28_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_28Ident := a_extension_28.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_28{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_28Ident, __gong__toRawStringLiteral(a_extension_28.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_28Ident, __gong__toRawStringLiteral(a_extension_28.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_28Ident, a_extension_28.Owns))
			for _, elem := range a_extension_28.AbstractFlightRoutingElementExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractFlightRoutingElementExtension = append(%s.AbstractFlightRoutingElementExtension, %s)", a_extension_28Ident, a_extension_28Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_29 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_29s, stage.AixmfeaturesStage.A_extension_29_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_29Ident := a_extension_29.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_29{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_29Ident, __gong__toRawStringLiteral(a_extension_29.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_29Ident, __gong__toRawStringLiteral(a_extension_29.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_29Ident, a_extension_29.Owns))
			for _, elem := range a_extension_29.AbstractRulesProceduresExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRulesProceduresExtension = append(%s.AbstractRulesProceduresExtension, %s)", a_extension_29Ident, a_extension_29Ident, targetIdent))
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
			for _, elem := range a_extension_3.AbstractAerialRefuellingTrackExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAerialRefuellingTrackExtension = append(%s.AbstractAerialRefuellingTrackExtension, %s)", a_extension_3Ident, a_extension_3Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_30 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_30s, stage.AixmfeaturesStage.A_extension_30_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_30Ident := a_extension_30.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_30{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_30Ident, __gong__toRawStringLiteral(a_extension_30.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_30Ident, __gong__toRawStringLiteral(a_extension_30.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_30Ident, a_extension_30.Owns))
			for _, elem := range a_extension_30.AbstractContactInformationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractContactInformationExtension = append(%s.AbstractContactInformationExtension, %s)", a_extension_30Ident, a_extension_30Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_31 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_31s, stage.AixmfeaturesStage.A_extension_31_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_31Ident := a_extension_31.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_31{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_31Ident, __gong__toRawStringLiteral(a_extension_31.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_31Ident, __gong__toRawStringLiteral(a_extension_31.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_31Ident, a_extension_31.Owns))
			for _, elem := range a_extension_31.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_31Ident, a_extension_31Ident, targetIdent))
			}
			for _, elem := range a_extension_31.AbstractOnlineContactExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractOnlineContactExtension = append(%s.AbstractOnlineContactExtension, %s)", a_extension_31Ident, a_extension_31Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_32 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_32s, stage.AixmfeaturesStage.A_extension_32_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_32Ident := a_extension_32.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_32{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_32Ident, __gong__toRawStringLiteral(a_extension_32.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_32Ident, __gong__toRawStringLiteral(a_extension_32.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_32Ident, a_extension_32.Owns))
			for _, elem := range a_extension_32.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_32Ident, a_extension_32Ident, targetIdent))
			}
			for _, elem := range a_extension_32.AbstractPostalAddressExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPostalAddressExtension = append(%s.AbstractPostalAddressExtension, %s)", a_extension_32Ident, a_extension_32Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_33 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_33s, stage.AixmfeaturesStage.A_extension_33_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_33Ident := a_extension_33.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_33{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_33Ident, __gong__toRawStringLiteral(a_extension_33.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_33Ident, __gong__toRawStringLiteral(a_extension_33.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_33Ident, a_extension_33.Owns))
			for _, elem := range a_extension_33.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_33Ident, a_extension_33Ident, targetIdent))
			}
			for _, elem := range a_extension_33.AbstractTelephoneContactExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractTelephoneContactExtension = append(%s.AbstractTelephoneContactExtension, %s)", a_extension_33Ident, a_extension_33Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_34 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_34s, stage.AixmfeaturesStage.A_extension_34_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_34Ident := a_extension_34.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_34{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_34Ident, __gong__toRawStringLiteral(a_extension_34.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_34Ident, __gong__toRawStringLiteral(a_extension_34.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_34Ident, a_extension_34.Owns))
			for _, elem := range a_extension_34.AbstractAircraftCharacteristicExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAircraftCharacteristicExtension = append(%s.AbstractAircraftCharacteristicExtension, %s)", a_extension_34Ident, a_extension_34Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_35 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_35s, stage.AixmfeaturesStage.A_extension_35_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_35Ident := a_extension_35.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_35{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_35Ident, __gong__toRawStringLiteral(a_extension_35.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_35Ident, __gong__toRawStringLiteral(a_extension_35.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_35Ident, a_extension_35.Owns))
			for _, elem := range a_extension_35.AbstractFlightCharacteristicExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractFlightCharacteristicExtension = append(%s.AbstractFlightCharacteristicExtension, %s)", a_extension_35Ident, a_extension_35Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_36 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_36s, stage.AixmfeaturesStage.A_extension_36_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_36Ident := a_extension_36.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_36{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_36Ident, __gong__toRawStringLiteral(a_extension_36.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_36Ident, __gong__toRawStringLiteral(a_extension_36.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_36Ident, a_extension_36.Owns))
			for _, elem := range a_extension_36.AbstractAirspaceLayerExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirspaceLayerExtension = append(%s.AbstractAirspaceLayerExtension, %s)", a_extension_36Ident, a_extension_36Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_37 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_37s, stage.AixmfeaturesStage.A_extension_37_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_37Ident := a_extension_37.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_37{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_37Ident, __gong__toRawStringLiteral(a_extension_37.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_37Ident, __gong__toRawStringLiteral(a_extension_37.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_37Ident, a_extension_37.Owns))
			for _, elem := range a_extension_37.AbstractMeteorologyExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractMeteorologyExtension = append(%s.AbstractMeteorologyExtension, %s)", a_extension_37Ident, a_extension_37Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_38 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_38s, stage.AixmfeaturesStage.A_extension_38_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_38Ident := a_extension_38.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_38{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_38Ident, __gong__toRawStringLiteral(a_extension_38.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_38Ident, __gong__toRawStringLiteral(a_extension_38.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_38Ident, a_extension_38.Owns))
			for _, elem := range a_extension_38.AbstractTimesheetExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractTimesheetExtension = append(%s.AbstractTimesheetExtension, %s)", a_extension_38Ident, a_extension_38Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_39 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_39s, stage.AixmfeaturesStage.A_extension_39_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_39Ident := a_extension_39.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_39{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_39Ident, __gong__toRawStringLiteral(a_extension_39.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_39Ident, __gong__toRawStringLiteral(a_extension_39.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_39Ident, a_extension_39.Owns))
			for _, elem := range a_extension_39.AbstractAltitudeAdjustmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAltitudeAdjustmentExtension = append(%s.AbstractAltitudeAdjustmentExtension, %s)", a_extension_39Ident, a_extension_39Ident, targetIdent))
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
			for _, elem := range a_extension_4.AbstractAuthorityForAerialRefuellingExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAuthorityForAerialRefuellingExtension = append(%s.AbstractAuthorityForAerialRefuellingExtension, %s)", a_extension_4Ident, a_extension_4Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_40 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_40s, stage.AixmfeaturesStage.A_extension_40_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_40Ident := a_extension_40.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_40{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_40Ident, __gong__toRawStringLiteral(a_extension_40.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_40Ident, __gong__toRawStringLiteral(a_extension_40.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_40Ident, a_extension_40.Owns))
			for _, elem := range a_extension_40.AbstractObstacleAssessmentAreaExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractObstacleAssessmentAreaExtension = append(%s.AbstractObstacleAssessmentAreaExtension, %s)", a_extension_40Ident, a_extension_40Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_41 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_41s, stage.AixmfeaturesStage.A_extension_41_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_41Ident := a_extension_41.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_41{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_41Ident, __gong__toRawStringLiteral(a_extension_41.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_41Ident, __gong__toRawStringLiteral(a_extension_41.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_41Ident, a_extension_41.Owns))
			for _, elem := range a_extension_41.AbstractObstaclePlacementExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractObstaclePlacementExtension = append(%s.AbstractObstaclePlacementExtension, %s)", a_extension_41Ident, a_extension_41Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, a_extension_42 := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.A_extension_42s, stage.AixmfeaturesStage.A_extension_42_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			a_extension_42Ident := a_extension_42.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_42{Name: %s}).Stage(stage.AixmfeaturesStage)", a_extension_42Ident, __gong__toRawStringLiteral(a_extension_42.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_42Ident, __gong__toRawStringLiteral(a_extension_42.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_42Ident, a_extension_42.Owns))
			for _, elem := range a_extension_42.AbstractObstructionExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractObstructionExtension = append(%s.AbstractObstructionExtension, %s)", a_extension_42Ident, a_extension_42Ident, targetIdent))
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
			for _, elem := range a_extension_5.AbstractElevatedCurveExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractElevatedCurveExtension = append(%s.AbstractElevatedCurveExtension, %s)", a_extension_5Ident, a_extension_5Ident, targetIdent))
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
			for _, elem := range a_extension_6.AbstractElevatedPointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractElevatedPointExtension = append(%s.AbstractElevatedPointExtension, %s)", a_extension_6Ident, a_extension_6Ident, targetIdent))
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
			for _, elem := range a_extension_7.AbstractElevatedSurfaceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractElevatedSurfaceExtension = append(%s.AbstractElevatedSurfaceExtension, %s)", a_extension_7Ident, a_extension_7Ident, targetIdent))
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
			for _, elem := range a_extension_8.AbstractAngleUseExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAngleUseExtension = append(%s.AbstractAngleUseExtension, %s)", a_extension_8Ident, a_extension_8Ident, targetIdent))
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
			for _, elem := range a_extension_9.AbstractSegmentPointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSegmentPointExtension = append(%s.AbstractSegmentPointExtension, %s)", a_extension_9Ident, a_extension_9Ident, targetIdent))
			}
			for _, elem := range a_extension_9.AbstractEnRouteSegmentPointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractEnRouteSegmentPointExtension = append(%s.AbstractEnRouteSegmentPointExtension, %s)", a_extension_9Ident, a_extension_9Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, abstractdirectflighttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AbstractDirectFlightTypes, stage.AixmfeaturesStage.AbstractDirectFlightType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			abstractdirectflighttypeIdent := abstractdirectflighttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractDirectFlightType{Name: %s}).Stage(stage.AixmfeaturesStage)", abstractdirectflighttypeIdent, __gong__toRawStringLiteral(abstractdirectflighttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractdirectflighttypeIdent, __gong__toRawStringLiteral(abstractdirectflighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractdirectflighttypeIdent, __gong__toRawStringLiteral(abstractdirectflighttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractdirectflighttypeIdent, __gong__toRawStringLiteral(abstractdirectflighttype.Description)))
			if abstractdirectflighttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractdirectflighttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractdirectflighttypeIdent, targetIdent))
			}
			for _, elem := range abstractdirectflighttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractdirectflighttypeIdent, abstractdirectflighttypeIdent, targetIdent))
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
		for _, abstractsegmentpointtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AbstractSegmentPointTypes, stage.AixmfeaturesStage.AbstractSegmentPointType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			abstractsegmentpointtypeIdent := abstractsegmentpointtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractSegmentPointType{Name: %s}).Stage(stage.AixmfeaturesStage)", abstractsegmentpointtypeIdent, __gong__toRawStringLiteral(abstractsegmentpointtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractsegmentpointtypeIdent, __gong__toRawStringLiteral(abstractsegmentpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractsegmentpointtypeIdent, __gong__toRawStringLiteral(abstractsegmentpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractsegmentpointtypeIdent, __gong__toRawStringLiteral(abstractsegmentpointtype.Description)))
			if abstractsegmentpointtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractsegmentpointtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractsegmentpointtypeIdent, targetIdent))
			}
			for _, elem := range abstractsegmentpointtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractsegmentpointtypeIdent, abstractsegmentpointtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aerialrefuellinganchorpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AerialRefuellingAnchorPropertyTypes, stage.AixmfeaturesStage.AerialRefuellingAnchorPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aerialrefuellinganchorpropertytypeIdent := aerialrefuellinganchorpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingAnchorPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", aerialrefuellinganchorpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchorpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellinganchorpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchorpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", aerialrefuellinganchorpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchorpropertytype.NilReason)))
			if aerialrefuellinganchorpropertytype.AerialRefuellingAnchor != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchorpropertytype.AerialRefuellingAnchor.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerialRefuellingAnchor = %s", aerialrefuellinganchorpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aerialrefuellinganchortype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AerialRefuellingAnchorTypes, stage.AixmfeaturesStage.AerialRefuellingAnchorType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aerialrefuellinganchortypeIdent := aerialrefuellinganchortype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingAnchorType{Name: %s}).Stage(stage.AixmfeaturesStage)", aerialrefuellinganchortypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchortype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellinganchortypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aerialrefuellinganchortypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aerialrefuellinganchortypeIdent, __gong__toRawStringLiteral(aerialrefuellinganchortype.Description)))
			if aerialrefuellinganchortype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellinganchortype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aerialrefuellinganchortypeIdent, aerialrefuellinganchortypeIdent, targetIdent))
			}
			if aerialrefuellinganchortype.OutboundCourse != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.OutboundCourse.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OutboundCourse = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			if aerialrefuellinganchortype.OutboundCourseType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.OutboundCourseType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OutboundCourseType = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			if aerialrefuellinganchortype.InboundCourse != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.InboundCourse.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InboundCourse = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			if aerialrefuellinganchortype.TurnDirection != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.TurnDirection.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TurnDirection = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			if aerialrefuellinganchortype.SpeedLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.SpeedLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpeedLimit = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			if aerialrefuellinganchortype.LegSeparation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.LegSeparation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LegSeparation = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			if aerialrefuellinganchortype.LegLength != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.LegLength.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LegLength = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			if aerialrefuellinganchortype.RefuellingBaseLevel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.RefuellingBaseLevel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RefuellingBaseLevel = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			if aerialrefuellinganchortype.RefuellingBaseLevelReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.RefuellingBaseLevelReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RefuellingBaseLevelReference = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			if aerialrefuellinganchortype.Extent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellinganchortype.Extent.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extent = %s", aerialrefuellinganchortypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellinganchortype.VerticalExtent {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalExtent = append(%s.VerticalExtent, %s)", aerialrefuellinganchortypeIdent, aerialrefuellinganchortypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellinganchortype.Point {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Point = append(%s.Point, %s)", aerialrefuellinganchortypeIdent, aerialrefuellinganchortypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellinganchortype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aerialrefuellinganchortypeIdent, aerialrefuellinganchortypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellinganchortype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aerialrefuellinganchortypeIdent, aerialrefuellinganchortypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aerialrefuellingpointpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AerialRefuellingPointPropertyTypes, stage.AixmfeaturesStage.AerialRefuellingPointPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aerialrefuellingpointpropertytypeIdent := aerialrefuellingpointpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingPointPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", aerialrefuellingpointpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingpointpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", aerialrefuellingpointpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointpropertytype.NilReason)))
			if aerialrefuellingpointpropertytype.AerialRefuellingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointpropertytype.AerialRefuellingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerialRefuellingPoint = %s", aerialrefuellingpointpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aerialrefuellingpointtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AerialRefuellingPointTypes, stage.AixmfeaturesStage.AerialRefuellingPointType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aerialrefuellingpointtypeIdent := aerialrefuellingpointtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingPointType{Name: %s}).Stage(stage.AixmfeaturesStage)", aerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aerialrefuellingpointtypeIdent, __gong__toRawStringLiteral(aerialrefuellingpointtype.Description)))
			if aerialrefuellingpointtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingpointtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aerialrefuellingpointtypeIdent, aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.ReportingATC != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.ReportingATC.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReportingATC = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.FlyOver != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.FlyOver.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlyOver = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.Waypoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.Waypoint.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Waypoint = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.RadarGuidance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.RadarGuidance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarGuidance = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingpointtype.FacilityMakeup {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FacilityMakeup = append(%s.FacilityMakeup, %s)", aerialrefuellingpointtypeIdent, aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_fixDesignatedPoint = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_navaidSystem = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_position = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_runwayPoint = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_aimingPoint = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.PointChoice_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.PointChoice_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_airportReferencePoint = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.ExtendedServiceVolume != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.ExtendedServiceVolume.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtendedServiceVolume = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingpointtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aerialrefuellingpointtypeIdent, aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.Sequence != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.Sequence.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequence = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			if aerialrefuellingpointtype.UsageType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingpointtype.UsageType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UsageType = %s", aerialrefuellingpointtypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingpointtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aerialrefuellingpointtypeIdent, aerialrefuellingpointtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aerialrefuellingpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AerialRefuellingPropertyTypes, stage.AixmfeaturesStage.AerialRefuellingPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aerialrefuellingpropertytypeIdent := aerialrefuellingpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", aerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, aerialrefuellingtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AerialRefuellingTimeSlicePropertyTypes, stage.AixmfeaturesStage.AerialRefuellingTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aerialrefuellingtimeslicepropertytypeIdent := aerialrefuellingtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", aerialrefuellingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", aerialrefuellingtimeslicepropertytypeIdent, aerialrefuellingtimeslicepropertytype.Owns))
			if aerialrefuellingtimeslicepropertytype.AerialRefuellingTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicepropertytype.AerialRefuellingTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerialRefuellingTimeSlice = %s", aerialrefuellingtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aerialrefuellingtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AerialRefuellingTimeSliceTypes, stage.AixmfeaturesStage.AerialRefuellingTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aerialrefuellingtimeslicetypeIdent := aerialrefuellingtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", aerialrefuellingtimeslicetypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingtimeslicetypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aerialrefuellingtimeslicetypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aerialrefuellingtimeslicetypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", aerialrefuellingtimeslicetypeIdent, __gong__toRawStringLiteral(aerialrefuellingtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetype.CorrectionNumber))
			if aerialrefuellingtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.DesignatorPrefix != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.DesignatorPrefix.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatorPrefix = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.DesignatorNumber != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.DesignatorNumber.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatorNumber = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.DesignatorSuffix != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.DesignatorSuffix.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatorSuffix = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.DesignatorDirection != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.DesignatorDirection.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatorDirection = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.RadarBeaconSetting != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.RadarBeaconSetting.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarBeaconSetting = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.XbandRadarSetting != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.XbandRadarSetting.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.XbandRadarSetting = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.TankerChannel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.TankerChannel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TankerChannel = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.ReceiverChannel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.ReceiverChannel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReceiverChannel = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.HelicopterRoute != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.HelicopterRoute.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HelicopterRoute = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.SpecialRefuelling != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.SpecialRefuelling.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialRefuelling = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.BidirectionalUse != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.BidirectionalUse.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BidirectionalUse = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.ReverseDirectionTurn != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.ReverseDirectionTurn.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReverseDirectionTurn = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.ProtectingAirspace {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProtectingAirspace = append(%s.ProtectingAirspace, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.Track {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Track = append(%s.Track, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.Anchor {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anchor = append(%s.Anchor, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			if aerialrefuellingtimeslicetype.OppositeTrack != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtimeslicetype.OppositeTrack.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OppositeTrack = %s", aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.ManagingOrganisation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ManagingOrganisation = append(%s.ManagingOrganisation, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aerialrefuellingtimeslicetypeIdent, aerialrefuellingtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aerialrefuellingtrackpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AerialRefuellingTrackPropertyTypes, stage.AixmfeaturesStage.AerialRefuellingTrackPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aerialrefuellingtrackpropertytypeIdent := aerialrefuellingtrackpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingTrackPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", aerialrefuellingtrackpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingtrackpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingtrackpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingtrackpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", aerialrefuellingtrackpropertytypeIdent, __gong__toRawStringLiteral(aerialrefuellingtrackpropertytype.NilReason)))
			if aerialrefuellingtrackpropertytype.AerialRefuellingTrack != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtrackpropertytype.AerialRefuellingTrack.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerialRefuellingTrack = %s", aerialrefuellingtrackpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aerialrefuellingtracktype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AerialRefuellingTrackTypes, stage.AixmfeaturesStage.AerialRefuellingTrackType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aerialrefuellingtracktypeIdent := aerialrefuellingtracktype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingTrackType{Name: %s}).Stage(stage.AixmfeaturesStage)", aerialrefuellingtracktypeIdent, __gong__toRawStringLiteral(aerialrefuellingtracktype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingtracktypeIdent, __gong__toRawStringLiteral(aerialrefuellingtracktype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aerialrefuellingtracktypeIdent, __gong__toRawStringLiteral(aerialrefuellingtracktype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aerialrefuellingtracktypeIdent, __gong__toRawStringLiteral(aerialrefuellingtracktype.Description)))
			if aerialrefuellingtracktype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtracktype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aerialrefuellingtracktypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtracktype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aerialrefuellingtracktypeIdent, aerialrefuellingtracktypeIdent, targetIdent))
			}
			if aerialrefuellingtracktype.Extent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtracktype.Extent.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extent = %s", aerialrefuellingtracktypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtracktype.Point {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Point = append(%s.Point, %s)", aerialrefuellingtracktypeIdent, aerialrefuellingtracktypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtracktype.VerticalExtent {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalExtent = append(%s.VerticalExtent, %s)", aerialrefuellingtracktypeIdent, aerialrefuellingtracktypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtracktype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aerialrefuellingtracktypeIdent, aerialrefuellingtracktypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtracktype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aerialrefuellingtracktypeIdent, aerialrefuellingtracktypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aerialrefuellingtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AerialRefuellingTypes, stage.AixmfeaturesStage.AerialRefuellingType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aerialrefuellingtypeIdent := aerialrefuellingtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AerialRefuellingType{Name: %s}).Stage(stage.AixmfeaturesStage)", aerialrefuellingtypeIdent, __gong__toRawStringLiteral(aerialrefuellingtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aerialrefuellingtypeIdent, __gong__toRawStringLiteral(aerialrefuellingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aerialrefuellingtypeIdent, __gong__toRawStringLiteral(aerialrefuellingtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aerialrefuellingtypeIdent, __gong__toRawStringLiteral(aerialrefuellingtype.Description)))
			if aerialrefuellingtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aerialrefuellingtypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aerialrefuellingtypeIdent, aerialrefuellingtypeIdent, targetIdent))
			}
			if aerialrefuellingtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", aerialrefuellingtypeIdent, targetIdent))
			}
			if aerialrefuellingtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aerialrefuellingtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", aerialrefuellingtypeIdent, targetIdent))
			}
			for _, elem := range aerialrefuellingtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", aerialrefuellingtypeIdent, aerialrefuellingtypeIdent, targetIdent))
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
		for _, airspacebordercrossingpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceBorderCrossingPropertyTypes, stage.AixmfeaturesStage.AirspaceBorderCrossingPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacebordercrossingpropertytypeIdent := airspacebordercrossingpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceBorderCrossingPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacebordercrossingpropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, airspacebordercrossingtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceBorderCrossingTimeSlicePropertyTypes, stage.AixmfeaturesStage.AirspaceBorderCrossingTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacebordercrossingtimeslicepropertytypeIdent := airspacebordercrossingtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceBorderCrossingTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacebordercrossingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacebordercrossingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airspacebordercrossingtimeslicepropertytypeIdent, airspacebordercrossingtimeslicepropertytype.Owns))
			if airspacebordercrossingtimeslicepropertytype.AirspaceBorderCrossingTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacebordercrossingtimeslicepropertytype.AirspaceBorderCrossingTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceBorderCrossingTimeSlice = %s", airspacebordercrossingtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacebordercrossingtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes, stage.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacebordercrossingtimeslicetypeIdent := airspacebordercrossingtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceBorderCrossingTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacebordercrossingtimeslicetypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacebordercrossingtimeslicetypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacebordercrossingtimeslicetypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacebordercrossingtimeslicetypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airspacebordercrossingtimeslicetypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airspacebordercrossingtimeslicetypeIdent, airspacebordercrossingtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airspacebordercrossingtimeslicetypeIdent, airspacebordercrossingtimeslicetype.CorrectionNumber))
			if airspacebordercrossingtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacebordercrossingtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacebordercrossingtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airspacebordercrossingtimeslicetypeIdent, airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			if airspacebordercrossingtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacebordercrossingtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			if airspacebordercrossingtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacebordercrossingtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			if airspacebordercrossingtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacebordercrossingtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			if airspacebordercrossingtimeslicetype.ExitedAirspace != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacebordercrossingtimeslicetype.ExitedAirspace.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExitedAirspace = %s", airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			if airspacebordercrossingtimeslicetype.EnteredAirspace != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacebordercrossingtimeslicetype.EnteredAirspace.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnteredAirspace = %s", airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacebordercrossingtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacebordercrossingtimeslicetypeIdent, airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacebordercrossingtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacebordercrossingtimeslicetypeIdent, airspacebordercrossingtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airspacebordercrossingtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirspaceBorderCrossingTypes, stage.AixmfeaturesStage.AirspaceBorderCrossingType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airspacebordercrossingtypeIdent := airspacebordercrossingtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceBorderCrossingType{Name: %s}).Stage(stage.AixmfeaturesStage)", airspacebordercrossingtypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacebordercrossingtypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacebordercrossingtypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacebordercrossingtypeIdent, __gong__toRawStringLiteral(airspacebordercrossingtype.Description)))
			if airspacebordercrossingtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacebordercrossingtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airspacebordercrossingtypeIdent, targetIdent))
			}
			for _, elem := range airspacebordercrossingtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airspacebordercrossingtypeIdent, airspacebordercrossingtypeIdent, targetIdent))
			}
			if airspacebordercrossingtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacebordercrossingtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", airspacebordercrossingtypeIdent, targetIdent))
			}
			if airspacebordercrossingtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airspacebordercrossingtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", airspacebordercrossingtypeIdent, targetIdent))
			}
			for _, elem := range airspacebordercrossingtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airspacebordercrossingtypeIdent, airspacebordercrossingtypeIdent, targetIdent))
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
		for _, altitudeadjustmentpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AltitudeAdjustmentPropertyTypes, stage.AixmfeaturesStage.AltitudeAdjustmentPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			altitudeadjustmentpropertytypeIdent := altitudeadjustmentpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AltitudeAdjustmentPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", altitudeadjustmentpropertytypeIdent, __gong__toRawStringLiteral(altitudeadjustmentpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", altitudeadjustmentpropertytypeIdent, __gong__toRawStringLiteral(altitudeadjustmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", altitudeadjustmentpropertytypeIdent, __gong__toRawStringLiteral(altitudeadjustmentpropertytype.NilReason)))
			if altitudeadjustmentpropertytype.AltitudeAdjustment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := altitudeadjustmentpropertytype.AltitudeAdjustment.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AltitudeAdjustment = %s", altitudeadjustmentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, altitudeadjustmenttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AltitudeAdjustmentTypes, stage.AixmfeaturesStage.AltitudeAdjustmentType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			altitudeadjustmenttypeIdent := altitudeadjustmenttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AltitudeAdjustmentType{Name: %s}).Stage(stage.AixmfeaturesStage)", altitudeadjustmenttypeIdent, __gong__toRawStringLiteral(altitudeadjustmenttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", altitudeadjustmenttypeIdent, __gong__toRawStringLiteral(altitudeadjustmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", altitudeadjustmenttypeIdent, __gong__toRawStringLiteral(altitudeadjustmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", altitudeadjustmenttypeIdent, __gong__toRawStringLiteral(altitudeadjustmenttype.Description)))
			if altitudeadjustmenttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := altitudeadjustmenttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", altitudeadjustmenttypeIdent, targetIdent))
			}
			for _, elem := range altitudeadjustmenttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", altitudeadjustmenttypeIdent, altitudeadjustmenttypeIdent, targetIdent))
			}
			if altitudeadjustmenttype.AltitudeAdjustmentType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := altitudeadjustmenttype.AltitudeAdjustmentType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AltitudeAdjustmentType = %s", altitudeadjustmenttypeIdent, targetIdent))
			}
			if altitudeadjustmenttype.PrimaryAlternateMinimum != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := altitudeadjustmenttype.PrimaryAlternateMinimum.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PrimaryAlternateMinimum = %s", altitudeadjustmenttypeIdent, targetIdent))
			}
			if altitudeadjustmenttype.AltitudeAdjustment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := altitudeadjustmenttype.AltitudeAdjustment.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AltitudeAdjustment = %s", altitudeadjustmenttypeIdent, targetIdent))
			}
			if altitudeadjustmenttype.LocalRemoteCode != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := altitudeadjustmenttype.LocalRemoteCode.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LocalRemoteCode = %s", altitudeadjustmenttypeIdent, targetIdent))
			}
			for _, elem := range altitudeadjustmenttype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", altitudeadjustmenttypeIdent, altitudeadjustmenttypeIdent, targetIdent))
			}
			for _, elem := range altitudeadjustmenttype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", altitudeadjustmenttypeIdent, altitudeadjustmenttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, angleindicationpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AngleIndicationPropertyTypes, stage.AixmfeaturesStage.AngleIndicationPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			angleindicationpropertytypeIdent := angleindicationpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleIndicationPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", angleindicationpropertytypeIdent, __gong__toRawStringLiteral(angleindicationpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, angleusepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AngleUsePropertyTypes, stage.AixmfeaturesStage.AngleUsePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			angleusepropertytypeIdent := angleusepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleUsePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", angleusepropertytypeIdent, __gong__toRawStringLiteral(angleusepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleusepropertytypeIdent, __gong__toRawStringLiteral(angleusepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", angleusepropertytypeIdent, __gong__toRawStringLiteral(angleusepropertytype.NilReason)))
			if angleusepropertytype.AngleUse != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleusepropertytype.AngleUse.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleUse = %s", angleusepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, angleusetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AngleUseTypes, stage.AixmfeaturesStage.AngleUseType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			angleusetypeIdent := angleusetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleUseType{Name: %s}).Stage(stage.AixmfeaturesStage)", angleusetypeIdent, __gong__toRawStringLiteral(angleusetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleusetypeIdent, __gong__toRawStringLiteral(angleusetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", angleusetypeIdent, __gong__toRawStringLiteral(angleusetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", angleusetypeIdent, __gong__toRawStringLiteral(angleusetype.Description)))
			if angleusetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleusetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", angleusetypeIdent, targetIdent))
			}
			for _, elem := range angleusetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", angleusetypeIdent, angleusetypeIdent, targetIdent))
			}
			if angleusetype.AlongCourseGuidance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleusetype.AlongCourseGuidance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AlongCourseGuidance = %s", angleusetypeIdent, targetIdent))
			}
			for _, elem := range angleusetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", angleusetypeIdent, angleusetypeIdent, targetIdent))
			}
			if angleusetype.TheAngleIndication != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleusetype.TheAngleIndication.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheAngleIndication = %s", angleusetypeIdent, targetIdent))
			}
			for _, elem := range angleusetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", angleusetypeIdent, angleusetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityforaerialrefuellingpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForAerialRefuellingPropertyTypes, stage.AixmfeaturesStage.AuthorityForAerialRefuellingPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityforaerialrefuellingpropertytypeIdent := authorityforaerialrefuellingpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForAerialRefuellingPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityforaerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforaerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", authorityforaerialrefuellingpropertytypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingpropertytype.NilReason)))
			if authorityforaerialrefuellingpropertytype.AuthorityForAerialRefuelling != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforaerialrefuellingpropertytype.AuthorityForAerialRefuelling.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForAerialRefuelling = %s", authorityforaerialrefuellingpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityforaerialrefuellingtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForAerialRefuellingTypes, stage.AixmfeaturesStage.AuthorityForAerialRefuellingType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityforaerialrefuellingtypeIdent := authorityforaerialrefuellingtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForAerialRefuellingType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityforaerialrefuellingtypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforaerialrefuellingtypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityforaerialrefuellingtypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityforaerialrefuellingtypeIdent, __gong__toRawStringLiteral(authorityforaerialrefuellingtype.Description)))
			if authorityforaerialrefuellingtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforaerialrefuellingtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", authorityforaerialrefuellingtypeIdent, targetIdent))
			}
			for _, elem := range authorityforaerialrefuellingtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", authorityforaerialrefuellingtypeIdent, authorityforaerialrefuellingtypeIdent, targetIdent))
			}
			if authorityforaerialrefuellingtype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforaerialrefuellingtype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", authorityforaerialrefuellingtypeIdent, targetIdent))
			}
			for _, elem := range authorityforaerialrefuellingtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", authorityforaerialrefuellingtypeIdent, authorityforaerialrefuellingtypeIdent, targetIdent))
			}
			if authorityforaerialrefuellingtype.TheOrganisationAuthority != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforaerialrefuellingtype.TheOrganisationAuthority.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheOrganisationAuthority = %s", authorityforaerialrefuellingtypeIdent, targetIdent))
			}
			for _, elem := range authorityforaerialrefuellingtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", authorityforaerialrefuellingtypeIdent, authorityforaerialrefuellingtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, changeoverpointtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ChangeOverPointTimeSlicePropertyTypes, stage.AixmfeaturesStage.ChangeOverPointTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			changeoverpointtimeslicepropertytypeIdent := changeoverpointtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ChangeOverPointTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", changeoverpointtimeslicepropertytypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", changeoverpointtimeslicepropertytypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", changeoverpointtimeslicepropertytypeIdent, changeoverpointtimeslicepropertytype.Owns))
			if changeoverpointtimeslicepropertytype.ChangeOverPointTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicepropertytype.ChangeOverPointTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ChangeOverPointTimeSlice = %s", changeoverpointtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, changeoverpointtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ChangeOverPointTimeSliceTypes, stage.AixmfeaturesStage.ChangeOverPointTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			changeoverpointtimeslicetypeIdent := changeoverpointtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ChangeOverPointTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", changeoverpointtimeslicetypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", changeoverpointtimeslicetypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", changeoverpointtimeslicetypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", changeoverpointtimeslicetypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", changeoverpointtimeslicetypeIdent, __gong__toRawStringLiteral(changeoverpointtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", changeoverpointtimeslicetypeIdent, changeoverpointtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", changeoverpointtimeslicetypeIdent, changeoverpointtimeslicetype.CorrectionNumber))
			if changeoverpointtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range changeoverpointtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", changeoverpointtimeslicetypeIdent, changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Distance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.Distance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_fixDesignatedPoint = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_navaidSystem = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_position = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_runwayPoint = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_aimingPoint = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.Location_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.Location_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_airportReferencePoint = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			if changeoverpointtimeslicetype.ApplicableRoutePortion != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtimeslicetype.ApplicableRoutePortion.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApplicableRoutePortion = %s", changeoverpointtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range changeoverpointtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", changeoverpointtimeslicetypeIdent, changeoverpointtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range changeoverpointtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", changeoverpointtimeslicetypeIdent, changeoverpointtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, changeoverpointtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ChangeOverPointTypes, stage.AixmfeaturesStage.ChangeOverPointType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			changeoverpointtypeIdent := changeoverpointtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ChangeOverPointType{Name: %s}).Stage(stage.AixmfeaturesStage)", changeoverpointtypeIdent, __gong__toRawStringLiteral(changeoverpointtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", changeoverpointtypeIdent, __gong__toRawStringLiteral(changeoverpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", changeoverpointtypeIdent, __gong__toRawStringLiteral(changeoverpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", changeoverpointtypeIdent, __gong__toRawStringLiteral(changeoverpointtype.Description)))
			if changeoverpointtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", changeoverpointtypeIdent, targetIdent))
			}
			for _, elem := range changeoverpointtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", changeoverpointtypeIdent, changeoverpointtypeIdent, targetIdent))
			}
			if changeoverpointtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", changeoverpointtypeIdent, targetIdent))
			}
			if changeoverpointtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := changeoverpointtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", changeoverpointtypeIdent, targetIdent))
			}
			for _, elem := range changeoverpointtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", changeoverpointtypeIdent, changeoverpointtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, contactinformationpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ContactInformationPropertyTypes, stage.AixmfeaturesStage.ContactInformationPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			contactinformationpropertytypeIdent := contactinformationpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ContactInformationPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", contactinformationpropertytypeIdent, __gong__toRawStringLiteral(contactinformationpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", contactinformationpropertytypeIdent, __gong__toRawStringLiteral(contactinformationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", contactinformationpropertytypeIdent, __gong__toRawStringLiteral(contactinformationpropertytype.NilReason)))
			if contactinformationpropertytype.ContactInformation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := contactinformationpropertytype.ContactInformation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ContactInformation = %s", contactinformationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, contactinformationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ContactInformationTypes, stage.AixmfeaturesStage.ContactInformationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			contactinformationtypeIdent := contactinformationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ContactInformationType{Name: %s}).Stage(stage.AixmfeaturesStage)", contactinformationtypeIdent, __gong__toRawStringLiteral(contactinformationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", contactinformationtypeIdent, __gong__toRawStringLiteral(contactinformationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", contactinformationtypeIdent, __gong__toRawStringLiteral(contactinformationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", contactinformationtypeIdent, __gong__toRawStringLiteral(contactinformationtype.Description)))
			if contactinformationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := contactinformationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", contactinformationtypeIdent, targetIdent))
			}
			for _, elem := range contactinformationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
			}
			if contactinformationtype.Title != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := contactinformationtype.Title.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title = %s", contactinformationtypeIdent, targetIdent))
			}
			for _, elem := range contactinformationtype.Address {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Address = append(%s.Address, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
			}
			for _, elem := range contactinformationtype.NetworkNode {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NetworkNode = append(%s.NetworkNode, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
			}
			for _, elem := range contactinformationtype.PhoneFax {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PhoneFax = append(%s.PhoneFax, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
			}
			for _, elem := range contactinformationtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
			}
			for _, elem := range contactinformationtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", contactinformationtypeIdent, contactinformationtypeIdent, targetIdent))
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
		for _, dmepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DMEPropertyTypes, stage.AixmfeaturesStage.DMEPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			dmepropertytypeIdent := dmepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DMEPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", dmepropertytypeIdent, __gong__toRawStringLiteral(dmepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, designatedpointpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DesignatedPointPropertyTypes, stage.AixmfeaturesStage.DesignatedPointPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			designatedpointpropertytypeIdent := designatedpointpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DesignatedPointPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", designatedpointpropertytypeIdent, __gong__toRawStringLiteral(designatedpointpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, directflightclasstype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DirectFlightClassTypes, stage.AixmfeaturesStage.DirectFlightClassType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			directflightclasstypeIdent := directflightclasstype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectFlightClassType{Name: %s}).Stage(stage.AixmfeaturesStage)", directflightclasstypeIdent, __gong__toRawStringLiteral(directflightclasstype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directflightclasstypeIdent, __gong__toRawStringLiteral(directflightclasstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directflightclasstypeIdent, __gong__toRawStringLiteral(directflightclasstype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", directflightclasstypeIdent, __gong__toRawStringLiteral(directflightclasstype.Description)))
			if directflightclasstype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightclasstype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", directflightclasstypeIdent, targetIdent))
			}
			for _, elem := range directflightclasstype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", directflightclasstypeIdent, directflightclasstypeIdent, targetIdent))
			}
			for _, elem := range directflightclasstype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", directflightclasstypeIdent, directflightclasstypeIdent, targetIdent))
			}
			if directflightclasstype.ExceedLength != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightclasstype.ExceedLength.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExceedLength = %s", directflightclasstypeIdent, targetIdent))
			}
			for _, elem := range directflightclasstype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", directflightclasstypeIdent, directflightclasstypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, directflightpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DirectFlightPropertyTypes, stage.AixmfeaturesStage.DirectFlightPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			directflightpropertytypeIdent := directflightpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectFlightPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", directflightpropertytypeIdent, __gong__toRawStringLiteral(directflightpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directflightpropertytypeIdent, __gong__toRawStringLiteral(directflightpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", directflightpropertytypeIdent, __gong__toRawStringLiteral(directflightpropertytype.NilReason)))
			if directflightpropertytype.DirectFlightClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightpropertytype.DirectFlightClass.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectFlightClass = %s", directflightpropertytypeIdent, targetIdent))
			}
			if directflightpropertytype.DirectFlightSegment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightpropertytype.DirectFlightSegment.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectFlightSegment = %s", directflightpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, directflightsegmentpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DirectFlightSegmentPropertyTypes, stage.AixmfeaturesStage.DirectFlightSegmentPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			directflightsegmentpropertytypeIdent := directflightsegmentpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectFlightSegmentPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", directflightsegmentpropertytypeIdent, __gong__toRawStringLiteral(directflightsegmentpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directflightsegmentpropertytypeIdent, __gong__toRawStringLiteral(directflightsegmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", directflightsegmentpropertytypeIdent, __gong__toRawStringLiteral(directflightsegmentpropertytype.NilReason)))
			if directflightsegmentpropertytype.DirectFlightSegment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmentpropertytype.DirectFlightSegment.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectFlightSegment = %s", directflightsegmentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, directflightsegmenttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DirectFlightSegmentTypes, stage.AixmfeaturesStage.DirectFlightSegmentType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			directflightsegmenttypeIdent := directflightsegmenttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectFlightSegmentType{Name: %s}).Stage(stage.AixmfeaturesStage)", directflightsegmenttypeIdent, __gong__toRawStringLiteral(directflightsegmenttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directflightsegmenttypeIdent, __gong__toRawStringLiteral(directflightsegmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directflightsegmenttypeIdent, __gong__toRawStringLiteral(directflightsegmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", directflightsegmenttypeIdent, __gong__toRawStringLiteral(directflightsegmenttype.Description)))
			if directflightsegmenttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", directflightsegmenttypeIdent, targetIdent))
			}
			for _, elem := range directflightsegmenttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", directflightsegmenttypeIdent, directflightsegmenttypeIdent, targetIdent))
			}
			for _, elem := range directflightsegmenttype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", directflightsegmenttypeIdent, directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.End_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_fixDesignatedPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.End_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_navaidSystem = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.End_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_position = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.End_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_runwayPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.End_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_aimingPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.End_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.End_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_airportReferencePoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.Start_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_fixDesignatedPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.Start_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_navaidSystem = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.Start_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_position = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.Start_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_runwayPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.Start_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_aimingPoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			if directflightsegmenttype.Start_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directflightsegmenttype.Start_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_airportReferencePoint = %s", directflightsegmenttypeIdent, targetIdent))
			}
			for _, elem := range directflightsegmenttype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", directflightsegmenttypeIdent, directflightsegmenttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, distanceindicationpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DistanceIndicationPropertyTypes, stage.AixmfeaturesStage.DistanceIndicationPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			distanceindicationpropertytypeIdent := distanceindicationpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DistanceIndicationPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", distanceindicationpropertytypeIdent, __gong__toRawStringLiteral(distanceindicationpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
		for _, enroutesegmentpointpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.EnRouteSegmentPointPropertyTypes, stage.AixmfeaturesStage.EnRouteSegmentPointPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			enroutesegmentpointpropertytypeIdent := enroutesegmentpointpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.EnRouteSegmentPointPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", enroutesegmentpointpropertytypeIdent, __gong__toRawStringLiteral(enroutesegmentpointpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", enroutesegmentpointpropertytypeIdent, __gong__toRawStringLiteral(enroutesegmentpointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", enroutesegmentpointpropertytypeIdent, __gong__toRawStringLiteral(enroutesegmentpointpropertytype.NilReason)))
			if enroutesegmentpointpropertytype.EnRouteSegmentPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointpropertytype.EnRouteSegmentPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnRouteSegmentPoint = %s", enroutesegmentpointpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, enroutesegmentpointtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.EnRouteSegmentPointTypes, stage.AixmfeaturesStage.EnRouteSegmentPointType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			enroutesegmentpointtypeIdent := enroutesegmentpointtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.EnRouteSegmentPointType{Name: %s}).Stage(stage.AixmfeaturesStage)", enroutesegmentpointtypeIdent, __gong__toRawStringLiteral(enroutesegmentpointtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", enroutesegmentpointtypeIdent, __gong__toRawStringLiteral(enroutesegmentpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", enroutesegmentpointtypeIdent, __gong__toRawStringLiteral(enroutesegmentpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", enroutesegmentpointtypeIdent, __gong__toRawStringLiteral(enroutesegmentpointtype.Description)))
			if enroutesegmentpointtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			for _, elem := range enroutesegmentpointtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", enroutesegmentpointtypeIdent, enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.ReportingATC != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.ReportingATC.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReportingATC = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.FlyOver != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.FlyOver.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlyOver = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.Waypoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.Waypoint.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Waypoint = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.RadarGuidance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.RadarGuidance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarGuidance = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			for _, elem := range enroutesegmentpointtype.FacilityMakeup {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FacilityMakeup = append(%s.FacilityMakeup, %s)", enroutesegmentpointtypeIdent, enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_fixDesignatedPoint = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_navaidSystem = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_position = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_runwayPoint = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_aimingPoint = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.PointChoice_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.PointChoice_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_airportReferencePoint = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.ExtendedServiceVolume != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.ExtendedServiceVolume.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtendedServiceVolume = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			for _, elem := range enroutesegmentpointtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", enroutesegmentpointtypeIdent, enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.RoleFreeFlight != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.RoleFreeFlight.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RoleFreeFlight = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.RoleRVSM != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.RoleRVSM.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RoleRVSM = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.TurnRadius != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.TurnRadius.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TurnRadius = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			if enroutesegmentpointtype.RoleMilitaryTraining != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := enroutesegmentpointtype.RoleMilitaryTraining.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RoleMilitaryTraining = %s", enroutesegmentpointtypeIdent, targetIdent))
			}
			for _, elem := range enroutesegmentpointtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", enroutesegmentpointtypeIdent, enroutesegmentpointtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightcharacteristicpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightCharacteristicPropertyTypes, stage.AixmfeaturesStage.FlightCharacteristicPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightcharacteristicpropertytypeIdent := flightcharacteristicpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightCharacteristicPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(flightcharacteristicpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(flightcharacteristicpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(flightcharacteristicpropertytype.NilReason)))
			if flightcharacteristicpropertytype.FlightCharacteristic != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightcharacteristicpropertytype.FlightCharacteristic.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCharacteristic = %s", flightcharacteristicpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightcharacteristictype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightCharacteristicTypes, stage.AixmfeaturesStage.FlightCharacteristicType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightcharacteristictypeIdent := flightcharacteristictype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightCharacteristicType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightcharacteristictypeIdent, __gong__toRawStringLiteral(flightcharacteristictype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightcharacteristictypeIdent, __gong__toRawStringLiteral(flightcharacteristictype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightcharacteristictypeIdent, __gong__toRawStringLiteral(flightcharacteristictype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightcharacteristictypeIdent, __gong__toRawStringLiteral(flightcharacteristictype.Description)))
			if flightcharacteristictype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightcharacteristictype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", flightcharacteristictypeIdent, targetIdent))
			}
			for _, elem := range flightcharacteristictype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", flightcharacteristictypeIdent, flightcharacteristictypeIdent, targetIdent))
			}
			if flightcharacteristictype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightcharacteristictype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", flightcharacteristictypeIdent, targetIdent))
			}
			if flightcharacteristictype.Rule != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightcharacteristictype.Rule.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rule = %s", flightcharacteristictypeIdent, targetIdent))
			}
			if flightcharacteristictype.Status != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightcharacteristictype.Status.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Status = %s", flightcharacteristictypeIdent, targetIdent))
			}
			if flightcharacteristictype.Military != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightcharacteristictype.Military.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Military = %s", flightcharacteristictypeIdent, targetIdent))
			}
			if flightcharacteristictype.Origin != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightcharacteristictype.Origin.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Origin = %s", flightcharacteristictypeIdent, targetIdent))
			}
			if flightcharacteristictype.Purpose != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightcharacteristictype.Purpose.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Purpose = %s", flightcharacteristictypeIdent, targetIdent))
			}
			for _, elem := range flightcharacteristictype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightcharacteristictypeIdent, flightcharacteristictypeIdent, targetIdent))
			}
			for _, elem := range flightcharacteristictype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightcharacteristictypeIdent, flightcharacteristictypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightconditioncircumstancepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightConditionCircumstancePropertyTypes, stage.AixmfeaturesStage.FlightConditionCircumstancePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightconditioncircumstancepropertytypeIdent := flightconditioncircumstancepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionCircumstancePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightconditioncircumstancepropertytypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditioncircumstancepropertytypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightconditioncircumstancepropertytypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancepropertytype.NilReason)))
			if flightconditioncircumstancepropertytype.FlightConditionCircumstance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditioncircumstancepropertytype.FlightConditionCircumstance.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightConditionCircumstance = %s", flightconditioncircumstancepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightconditioncircumstancetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightConditionCircumstanceTypes, stage.AixmfeaturesStage.FlightConditionCircumstanceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightconditioncircumstancetypeIdent := flightconditioncircumstancetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionCircumstanceType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightconditioncircumstancetypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditioncircumstancetypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightconditioncircumstancetypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightconditioncircumstancetypeIdent, __gong__toRawStringLiteral(flightconditioncircumstancetype.Description)))
			if flightconditioncircumstancetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditioncircumstancetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", flightconditioncircumstancetypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncircumstancetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", flightconditioncircumstancetypeIdent, flightconditioncircumstancetypeIdent, targetIdent))
			}
			if flightconditioncircumstancetype.ReferenceLocation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditioncircumstancetype.ReferenceLocation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferenceLocation = %s", flightconditioncircumstancetypeIdent, targetIdent))
			}
			if flightconditioncircumstancetype.RelationWithLocation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditioncircumstancetype.RelationWithLocation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RelationWithLocation = %s", flightconditioncircumstancetypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncircumstancetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightconditioncircumstancetypeIdent, flightconditioncircumstancetypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncircumstancetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightconditioncircumstancetypeIdent, flightconditioncircumstancetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightconditioncombinationpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightConditionCombinationPropertyTypes, stage.AixmfeaturesStage.FlightConditionCombinationPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightconditioncombinationpropertytypeIdent := flightconditioncombinationpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionCombinationPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightconditioncombinationpropertytypeIdent, __gong__toRawStringLiteral(flightconditioncombinationpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditioncombinationpropertytypeIdent, __gong__toRawStringLiteral(flightconditioncombinationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightconditioncombinationpropertytypeIdent, __gong__toRawStringLiteral(flightconditioncombinationpropertytype.NilReason)))
			if flightconditioncombinationpropertytype.FlightConditionCombination != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditioncombinationpropertytype.FlightConditionCombination.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightConditionCombination = %s", flightconditioncombinationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightconditioncombinationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightConditionCombinationTypes, stage.AixmfeaturesStage.FlightConditionCombinationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightconditioncombinationtypeIdent := flightconditioncombinationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionCombinationType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightconditioncombinationtypeIdent, __gong__toRawStringLiteral(flightconditioncombinationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditioncombinationtypeIdent, __gong__toRawStringLiteral(flightconditioncombinationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightconditioncombinationtypeIdent, __gong__toRawStringLiteral(flightconditioncombinationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightconditioncombinationtypeIdent, __gong__toRawStringLiteral(flightconditioncombinationtype.Description)))
			if flightconditioncombinationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditioncombinationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", flightconditioncombinationtypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncombinationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncombinationtype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncombinationtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncombinationtype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
			if flightconditioncombinationtype.LogicalOperator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditioncombinationtype.LogicalOperator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LogicalOperator = %s", flightconditioncombinationtypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncombinationtype.Element {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element = append(%s.Element, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
			for _, elem := range flightconditioncombinationtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightconditioncombinationtypeIdent, flightconditioncombinationtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightconditionelementpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightConditionElementPropertyTypes, stage.AixmfeaturesStage.FlightConditionElementPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightconditionelementpropertytypeIdent := flightconditionelementpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionElementPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightconditionelementpropertytypeIdent, __gong__toRawStringLiteral(flightconditionelementpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditionelementpropertytypeIdent, __gong__toRawStringLiteral(flightconditionelementpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightconditionelementpropertytypeIdent, __gong__toRawStringLiteral(flightconditionelementpropertytype.NilReason)))
			if flightconditionelementpropertytype.FlightConditionElement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementpropertytype.FlightConditionElement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightConditionElement = %s", flightconditionelementpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightconditionelementtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightConditionElementTypes, stage.AixmfeaturesStage.FlightConditionElementType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightconditionelementtypeIdent := flightconditionelementtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightConditionElementType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightconditionelementtypeIdent, __gong__toRawStringLiteral(flightconditionelementtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightconditionelementtypeIdent, __gong__toRawStringLiteral(flightconditionelementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightconditionelementtypeIdent, __gong__toRawStringLiteral(flightconditionelementtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightconditionelementtypeIdent, __gong__toRawStringLiteral(flightconditionelementtype.Description)))
			if flightconditionelementtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", flightconditionelementtypeIdent, targetIdent))
			}
			for _, elem := range flightconditionelementtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", flightconditionelementtypeIdent, flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.Index != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.Index.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Index = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_weather != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_weather.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_weather = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_aircraft != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_aircraft.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_aircraft = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_flight != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_flight.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_flight = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_directFlightCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_directFlightCondition.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_directFlightCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_borderCrossingCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_borderCrossingCondition.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_borderCrossingCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_operand != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_operand.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_operand = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_routePortionCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_routePortionCondition.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_routePortionCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_standardInstrumentDepartureCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_standardInstrumentDepartureCondition.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_standardInstrumentDepartureCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_standardInstrumentArrivalCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_standardInstrumentArrivalCondition.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_standardInstrumentArrivalCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_organisationCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_organisationCondition.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_organisationCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_fixDesignatedPoint = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_navaidSystem = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_position = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_runwayPoint = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_aimingPoint = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.SignificantPointCondition_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.SignificantPointCondition_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointCondition_airportReferencePoint = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_airspaceCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_airspaceCondition.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_airspaceCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_airportHeliportCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_airportHeliportCondition.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_airportHeliportCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.FlightCondition_aerialRefuellingCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.FlightCondition_aerialRefuellingCondition.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightCondition_aerialRefuellingCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			if flightconditionelementtype.OperationalCondition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightconditionelementtype.OperationalCondition.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OperationalCondition = %s", flightconditionelementtypeIdent, targetIdent))
			}
			for _, elem := range flightconditionelementtype.FlightLevel {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightLevel = append(%s.FlightLevel, %s)", flightconditionelementtypeIdent, flightconditionelementtypeIdent, targetIdent))
			}
			for _, elem := range flightconditionelementtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightconditionelementtypeIdent, flightconditionelementtypeIdent, targetIdent))
			}
			for _, elem := range flightconditionelementtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightconditionelementtypeIdent, flightconditionelementtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightrestrictionlevelpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightRestrictionLevelPropertyTypes, stage.AixmfeaturesStage.FlightRestrictionLevelPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightrestrictionlevelpropertytypeIdent := flightrestrictionlevelpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionLevelPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightrestrictionlevelpropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionlevelpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictionlevelpropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionlevelpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightrestrictionlevelpropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionlevelpropertytype.NilReason)))
			if flightrestrictionlevelpropertytype.FlightRestrictionLevel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictionlevelpropertytype.FlightRestrictionLevel.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRestrictionLevel = %s", flightrestrictionlevelpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightrestrictionleveltype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightRestrictionLevelTypes, stage.AixmfeaturesStage.FlightRestrictionLevelType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightrestrictionleveltypeIdent := flightrestrictionleveltype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionLevelType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightrestrictionleveltypeIdent, __gong__toRawStringLiteral(flightrestrictionleveltype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictionleveltypeIdent, __gong__toRawStringLiteral(flightrestrictionleveltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightrestrictionleveltypeIdent, __gong__toRawStringLiteral(flightrestrictionleveltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightrestrictionleveltypeIdent, __gong__toRawStringLiteral(flightrestrictionleveltype.Description)))
			if flightrestrictionleveltype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictionleveltype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", flightrestrictionleveltypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionleveltype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", flightrestrictionleveltypeIdent, flightrestrictionleveltypeIdent, targetIdent))
			}
			if flightrestrictionleveltype.UpperLevel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictionleveltype.UpperLevel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLevel = %s", flightrestrictionleveltypeIdent, targetIdent))
			}
			if flightrestrictionleveltype.UpperLevelReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictionleveltype.UpperLevelReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLevelReference = %s", flightrestrictionleveltypeIdent, targetIdent))
			}
			if flightrestrictionleveltype.LowerLevel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictionleveltype.LowerLevel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLevel = %s", flightrestrictionleveltypeIdent, targetIdent))
			}
			if flightrestrictionleveltype.LowerLevelReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictionleveltype.LowerLevelReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLevelReference = %s", flightrestrictionleveltypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionleveltype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightrestrictionleveltypeIdent, flightrestrictionleveltypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionleveltype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightrestrictionleveltypeIdent, flightrestrictionleveltypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightrestrictionroutepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightRestrictionRoutePropertyTypes, stage.AixmfeaturesStage.FlightRestrictionRoutePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightrestrictionroutepropertytypeIdent := flightrestrictionroutepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionRoutePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightrestrictionroutepropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionroutepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictionroutepropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionroutepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightrestrictionroutepropertytypeIdent, __gong__toRawStringLiteral(flightrestrictionroutepropertytype.NilReason)))
			if flightrestrictionroutepropertytype.FlightRestrictionRoute != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictionroutepropertytype.FlightRestrictionRoute.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRestrictionRoute = %s", flightrestrictionroutepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightrestrictionroutetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightRestrictionRouteTypes, stage.AixmfeaturesStage.FlightRestrictionRouteType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightrestrictionroutetypeIdent := flightrestrictionroutetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionRouteType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightrestrictionroutetypeIdent, __gong__toRawStringLiteral(flightrestrictionroutetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictionroutetypeIdent, __gong__toRawStringLiteral(flightrestrictionroutetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightrestrictionroutetypeIdent, __gong__toRawStringLiteral(flightrestrictionroutetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightrestrictionroutetypeIdent, __gong__toRawStringLiteral(flightrestrictionroutetype.Description)))
			if flightrestrictionroutetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictionroutetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", flightrestrictionroutetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionroutetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", flightrestrictionroutetypeIdent, flightrestrictionroutetypeIdent, targetIdent))
			}
			if flightrestrictionroutetype.PriorPermission != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictionroutetype.PriorPermission.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PriorPermission = %s", flightrestrictionroutetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionroutetype.RouteElement {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteElement = append(%s.RouteElement, %s)", flightrestrictionroutetypeIdent, flightrestrictionroutetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionroutetype.Contact {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", flightrestrictionroutetypeIdent, flightrestrictionroutetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionroutetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightrestrictionroutetypeIdent, flightrestrictionroutetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictionroutetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightrestrictionroutetypeIdent, flightrestrictionroutetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightrestrictiontimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightRestrictionTimeSlicePropertyTypes, stage.AixmfeaturesStage.FlightRestrictionTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightrestrictiontimeslicepropertytypeIdent := flightrestrictiontimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightrestrictiontimeslicepropertytypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictiontimeslicepropertytypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", flightrestrictiontimeslicepropertytypeIdent, flightrestrictiontimeslicepropertytype.Owns))
			if flightrestrictiontimeslicepropertytype.FlightRestrictionTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontimeslicepropertytype.FlightRestrictionTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRestrictionTimeSlice = %s", flightrestrictiontimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightrestrictiontimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightRestrictionTimeSliceTypes, stage.AixmfeaturesStage.FlightRestrictionTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightrestrictiontimeslicetypeIdent := flightrestrictiontimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightrestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightrestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightrestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", flightrestrictiontimeslicetypeIdent, __gong__toRawStringLiteral(flightrestrictiontimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetype.CorrectionNumber))
			if flightrestrictiontimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictiontimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			if flightrestrictiontimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			if flightrestrictiontimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			if flightrestrictiontimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			if flightrestrictiontimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			if flightrestrictiontimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			if flightrestrictiontimeslicetype.Instruction != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontimeslicetype.Instruction.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Instruction = %s", flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			if flightrestrictiontimeslicetype.Flight != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontimeslicetype.Flight.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Flight = %s", flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictiontimeslicetype.RegulatedRoute {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RegulatedRoute = append(%s.RegulatedRoute, %s)", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictiontimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictiontimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightrestrictiontimeslicetypeIdent, flightrestrictiontimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightrestrictiontype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightRestrictionTypes, stage.AixmfeaturesStage.FlightRestrictionType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightrestrictiontypeIdent := flightrestrictiontype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRestrictionType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightrestrictiontypeIdent, __gong__toRawStringLiteral(flightrestrictiontype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightrestrictiontypeIdent, __gong__toRawStringLiteral(flightrestrictiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightrestrictiontypeIdent, __gong__toRawStringLiteral(flightrestrictiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightrestrictiontypeIdent, __gong__toRawStringLiteral(flightrestrictiontype.Description)))
			if flightrestrictiontype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", flightrestrictiontypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictiontype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", flightrestrictiontypeIdent, flightrestrictiontypeIdent, targetIdent))
			}
			if flightrestrictiontype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", flightrestrictiontypeIdent, targetIdent))
			}
			if flightrestrictiontype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightrestrictiontype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", flightrestrictiontypeIdent, targetIdent))
			}
			for _, elem := range flightrestrictiontype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", flightrestrictiontypeIdent, flightrestrictiontypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightroutingelementpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightRoutingElementPropertyTypes, stage.AixmfeaturesStage.FlightRoutingElementPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightroutingelementpropertytypeIdent := flightroutingelementpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRoutingElementPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightroutingelementpropertytypeIdent, __gong__toRawStringLiteral(flightroutingelementpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightroutingelementpropertytypeIdent, __gong__toRawStringLiteral(flightroutingelementpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", flightroutingelementpropertytypeIdent, __gong__toRawStringLiteral(flightroutingelementpropertytype.NilReason)))
			if flightroutingelementpropertytype.FlightRoutingElement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementpropertytype.FlightRoutingElement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRoutingElement = %s", flightroutingelementpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, flightroutingelementtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FlightRoutingElementTypes, stage.AixmfeaturesStage.FlightRoutingElementType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			flightroutingelementtypeIdent := flightroutingelementtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FlightRoutingElementType{Name: %s}).Stage(stage.AixmfeaturesStage)", flightroutingelementtypeIdent, __gong__toRawStringLiteral(flightroutingelementtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", flightroutingelementtypeIdent, __gong__toRawStringLiteral(flightroutingelementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", flightroutingelementtypeIdent, __gong__toRawStringLiteral(flightroutingelementtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", flightroutingelementtypeIdent, __gong__toRawStringLiteral(flightroutingelementtype.Description)))
			if flightroutingelementtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", flightroutingelementtypeIdent, targetIdent))
			}
			for _, elem := range flightroutingelementtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", flightroutingelementtypeIdent, flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.OrderNumber != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.OrderNumber.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OrderNumber = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Speed != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.Speed.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Speed = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.SpeedReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.SpeedReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpeedReference = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.SpeedCriteria != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.SpeedCriteria.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpeedCriteria = %s", flightroutingelementtypeIdent, targetIdent))
			}
			for _, elem := range flightroutingelementtype.FlightLevel {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightLevel = append(%s.FlightLevel, %s)", flightroutingelementtypeIdent, flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_directFlightElement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.Element_directFlightElement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_directFlightElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_routePortionElement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.Element_routePortionElement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_routePortionElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_standardInstrumentDepartureElement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.Element_standardInstrumentDepartureElement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_standardInstrumentDepartureElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_standardInstrumentArrivalElement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.Element_standardInstrumentArrivalElement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_standardInstrumentArrivalElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.PointElement_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_fixDesignatedPoint = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.PointElement_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_navaidSystem = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.PointElement_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_position = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.PointElement_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_runwayPoint = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.PointElement_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_aimingPoint = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.PointElement_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.PointElement_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointElement_airportReferencePoint = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_airspaceElement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.Element_airspaceElement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_airspaceElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_airportHeliportElement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.Element_airportHeliportElement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_airportHeliportElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			if flightroutingelementtype.Element_aerialRefuellingElement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := flightroutingelementtype.Element_aerialRefuellingElement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Element_aerialRefuellingElement = %s", flightroutingelementtypeIdent, targetIdent))
			}
			for _, elem := range flightroutingelementtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", flightroutingelementtypeIdent, flightroutingelementtypeIdent, targetIdent))
			}
			for _, elem := range flightroutingelementtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", flightroutingelementtypeIdent, flightroutingelementtypeIdent, targetIdent))
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
		for _, meteorologypropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.MeteorologyPropertyTypes, stage.AixmfeaturesStage.MeteorologyPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			meteorologypropertytypeIdent := meteorologypropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MeteorologyPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", meteorologypropertytypeIdent, __gong__toRawStringLiteral(meteorologypropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meteorologypropertytypeIdent, __gong__toRawStringLiteral(meteorologypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", meteorologypropertytypeIdent, __gong__toRawStringLiteral(meteorologypropertytype.NilReason)))
			if meteorologypropertytype.Meteorology != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := meteorologypropertytype.Meteorology.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Meteorology = %s", meteorologypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, meteorologytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.MeteorologyTypes, stage.AixmfeaturesStage.MeteorologyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			meteorologytypeIdent := meteorologytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MeteorologyType{Name: %s}).Stage(stage.AixmfeaturesStage)", meteorologytypeIdent, __gong__toRawStringLiteral(meteorologytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meteorologytypeIdent, __gong__toRawStringLiteral(meteorologytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", meteorologytypeIdent, __gong__toRawStringLiteral(meteorologytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", meteorologytypeIdent, __gong__toRawStringLiteral(meteorologytype.Description)))
			if meteorologytype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := meteorologytype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", meteorologytypeIdent, targetIdent))
			}
			for _, elem := range meteorologytype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", meteorologytypeIdent, meteorologytypeIdent, targetIdent))
			}
			if meteorologytype.FlightConditions != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := meteorologytype.FlightConditions.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightConditions = %s", meteorologytypeIdent, targetIdent))
			}
			if meteorologytype.Visibility != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := meteorologytype.Visibility.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Visibility = %s", meteorologytypeIdent, targetIdent))
			}
			if meteorologytype.VisibilityInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := meteorologytype.VisibilityInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VisibilityInterpretation = %s", meteorologytypeIdent, targetIdent))
			}
			if meteorologytype.RunwayVisualRange != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := meteorologytype.RunwayVisualRange.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RunwayVisualRange = %s", meteorologytypeIdent, targetIdent))
			}
			if meteorologytype.RunwayVisualRangeInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := meteorologytype.RunwayVisualRangeInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RunwayVisualRangeInterpretation = %s", meteorologytypeIdent, targetIdent))
			}
			for _, elem := range meteorologytype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", meteorologytypeIdent, meteorologytypeIdent, targetIdent))
			}
			for _, elem := range meteorologytype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", meteorologytypeIdent, meteorologytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, navaidpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidPropertyTypes, stage.AixmfeaturesStage.NavaidPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidpropertytypeIdent := navaidpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidpropertytypeIdent, __gong__toRawStringLiteral(navaidpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
		for _, obstacleassessmentareapropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ObstacleAssessmentAreaPropertyTypes, stage.AixmfeaturesStage.ObstacleAssessmentAreaPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			obstacleassessmentareapropertytypeIdent := obstacleassessmentareapropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstacleAssessmentAreaPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", obstacleassessmentareapropertytypeIdent, __gong__toRawStringLiteral(obstacleassessmentareapropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleassessmentareapropertytypeIdent, __gong__toRawStringLiteral(obstacleassessmentareapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", obstacleassessmentareapropertytypeIdent, __gong__toRawStringLiteral(obstacleassessmentareapropertytype.NilReason)))
			if obstacleassessmentareapropertytype.ObstacleAssessmentArea != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareapropertytype.ObstacleAssessmentArea.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstacleAssessmentArea = %s", obstacleassessmentareapropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, obstacleassessmentareatype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ObstacleAssessmentAreaTypes, stage.AixmfeaturesStage.ObstacleAssessmentAreaType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			obstacleassessmentareatypeIdent := obstacleassessmentareatype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstacleAssessmentAreaType{Name: %s}).Stage(stage.AixmfeaturesStage)", obstacleassessmentareatypeIdent, __gong__toRawStringLiteral(obstacleassessmentareatype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleassessmentareatypeIdent, __gong__toRawStringLiteral(obstacleassessmentareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", obstacleassessmentareatypeIdent, __gong__toRawStringLiteral(obstacleassessmentareatype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", obstacleassessmentareatypeIdent, __gong__toRawStringLiteral(obstacleassessmentareatype.Description)))
			if obstacleassessmentareatype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			for _, elem := range obstacleassessmentareatype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", obstacleassessmentareatypeIdent, obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.SectionNumber != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.SectionNumber.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SectionNumber = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.Slope != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.Slope.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slope = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.AssessedAltitude != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.AssessedAltitude.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AssessedAltitude = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.SlopeLowerAltitude != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.SlopeLowerAltitude.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SlopeLowerAltitude = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.GradientLowHigh != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.GradientLowHigh.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GradientLowHigh = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.SurfaceZone != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.SurfaceZone.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SurfaceZone = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.SafetyRegulation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.SafetyRegulation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SafetyRegulation = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			for _, elem := range obstacleassessmentareatype.AircraftCategory {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCategory = append(%s.AircraftCategory, %s)", obstacleassessmentareatypeIdent, obstacleassessmentareatypeIdent, targetIdent))
			}
			for _, elem := range obstacleassessmentareatype.SignificantObstacle {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantObstacle = append(%s.SignificantObstacle, %s)", obstacleassessmentareatypeIdent, obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.Surface != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.Surface.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Surface = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			if obstacleassessmentareatype.StartingCurve != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleassessmentareatype.StartingCurve.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartingCurve = %s", obstacleassessmentareatypeIdent, targetIdent))
			}
			for _, elem := range obstacleassessmentareatype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", obstacleassessmentareatypeIdent, obstacleassessmentareatypeIdent, targetIdent))
			}
			for _, elem := range obstacleassessmentareatype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", obstacleassessmentareatypeIdent, obstacleassessmentareatypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, obstacleplacementpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ObstaclePlacementPropertyTypes, stage.AixmfeaturesStage.ObstaclePlacementPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			obstacleplacementpropertytypeIdent := obstacleplacementpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstaclePlacementPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", obstacleplacementpropertytypeIdent, __gong__toRawStringLiteral(obstacleplacementpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleplacementpropertytypeIdent, __gong__toRawStringLiteral(obstacleplacementpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", obstacleplacementpropertytypeIdent, __gong__toRawStringLiteral(obstacleplacementpropertytype.NilReason)))
			if obstacleplacementpropertytype.ObstaclePlacement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleplacementpropertytype.ObstaclePlacement.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstaclePlacement = %s", obstacleplacementpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, obstacleplacementtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ObstaclePlacementTypes, stage.AixmfeaturesStage.ObstaclePlacementType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			obstacleplacementtypeIdent := obstacleplacementtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstaclePlacementType{Name: %s}).Stage(stage.AixmfeaturesStage)", obstacleplacementtypeIdent, __gong__toRawStringLiteral(obstacleplacementtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleplacementtypeIdent, __gong__toRawStringLiteral(obstacleplacementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", obstacleplacementtypeIdent, __gong__toRawStringLiteral(obstacleplacementtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", obstacleplacementtypeIdent, __gong__toRawStringLiteral(obstacleplacementtype.Description)))
			if obstacleplacementtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleplacementtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", obstacleplacementtypeIdent, targetIdent))
			}
			for _, elem := range obstacleplacementtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", obstacleplacementtypeIdent, obstacleplacementtypeIdent, targetIdent))
			}
			if obstacleplacementtype.ObstacleBearing != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleplacementtype.ObstacleBearing.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstacleBearing = %s", obstacleplacementtypeIdent, targetIdent))
			}
			if obstacleplacementtype.ObstacleDistance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleplacementtype.ObstacleDistance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstacleDistance = %s", obstacleplacementtypeIdent, targetIdent))
			}
			if obstacleplacementtype.PointType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleplacementtype.PointType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointType = %s", obstacleplacementtypeIdent, targetIdent))
			}
			if obstacleplacementtype.ObstaclePlacement != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstacleplacementtype.ObstaclePlacement.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstaclePlacement = %s", obstacleplacementtypeIdent, targetIdent))
			}
			for _, elem := range obstacleplacementtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", obstacleplacementtypeIdent, obstacleplacementtypeIdent, targetIdent))
			}
			for _, elem := range obstacleplacementtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", obstacleplacementtypeIdent, obstacleplacementtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, obstructionpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ObstructionPropertyTypes, stage.AixmfeaturesStage.ObstructionPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			obstructionpropertytypeIdent := obstructionpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstructionPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", obstructionpropertytypeIdent, __gong__toRawStringLiteral(obstructionpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstructionpropertytypeIdent, __gong__toRawStringLiteral(obstructionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", obstructionpropertytypeIdent, __gong__toRawStringLiteral(obstructionpropertytype.NilReason)))
			if obstructionpropertytype.Obstruction != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstructionpropertytype.Obstruction.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Obstruction = %s", obstructionpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, obstructiontype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ObstructionTypes, stage.AixmfeaturesStage.ObstructionType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			obstructiontypeIdent := obstructiontype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstructionType{Name: %s}).Stage(stage.AixmfeaturesStage)", obstructiontypeIdent, __gong__toRawStringLiteral(obstructiontype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstructiontypeIdent, __gong__toRawStringLiteral(obstructiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", obstructiontypeIdent, __gong__toRawStringLiteral(obstructiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", obstructiontypeIdent, __gong__toRawStringLiteral(obstructiontype.Description)))
			if obstructiontype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstructiontype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", obstructiontypeIdent, targetIdent))
			}
			for _, elem := range obstructiontype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", obstructiontypeIdent, obstructiontypeIdent, targetIdent))
			}
			if obstructiontype.RequiredClearance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstructiontype.RequiredClearance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequiredClearance = %s", obstructiontypeIdent, targetIdent))
			}
			if obstructiontype.MinimumAltitude != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstructiontype.MinimumAltitude.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinimumAltitude = %s", obstructiontypeIdent, targetIdent))
			}
			if obstructiontype.SurfacePenetration != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstructiontype.SurfacePenetration.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SurfacePenetration = %s", obstructiontypeIdent, targetIdent))
			}
			if obstructiontype.SlopePenetration != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstructiontype.SlopePenetration.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SlopePenetration = %s", obstructiontypeIdent, targetIdent))
			}
			if obstructiontype.Controlling != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstructiontype.Controlling.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Controlling = %s", obstructiontypeIdent, targetIdent))
			}
			if obstructiontype.CloseIn != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstructiontype.CloseIn.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CloseIn = %s", obstructiontypeIdent, targetIdent))
			}
			if obstructiontype.TheVerticalStructure != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := obstructiontype.TheVerticalStructure.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheVerticalStructure = %s", obstructiontypeIdent, targetIdent))
			}
			for _, elem := range obstructiontype.Adjustment {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Adjustment = append(%s.Adjustment, %s)", obstructiontypeIdent, obstructiontypeIdent, targetIdent))
			}
			for _, elem := range obstructiontype.ObstaclePlacement {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstaclePlacement = append(%s.ObstaclePlacement, %s)", obstructiontypeIdent, obstructiontypeIdent, targetIdent))
			}
			for _, elem := range obstructiontype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", obstructiontypeIdent, obstructiontypeIdent, targetIdent))
			}
			for _, elem := range obstructiontype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", obstructiontypeIdent, obstructiontypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, onlinecontactpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OnlineContactPropertyTypes, stage.AixmfeaturesStage.OnlineContactPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			onlinecontactpropertytypeIdent := onlinecontactpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OnlineContactPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", onlinecontactpropertytypeIdent, __gong__toRawStringLiteral(onlinecontactpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", onlinecontactpropertytypeIdent, __gong__toRawStringLiteral(onlinecontactpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", onlinecontactpropertytypeIdent, __gong__toRawStringLiteral(onlinecontactpropertytype.NilReason)))
			if onlinecontactpropertytype.OnlineContact != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := onlinecontactpropertytype.OnlineContact.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OnlineContact = %s", onlinecontactpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, onlinecontacttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OnlineContactTypes, stage.AixmfeaturesStage.OnlineContactType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			onlinecontacttypeIdent := onlinecontacttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OnlineContactType{Name: %s}).Stage(stage.AixmfeaturesStage)", onlinecontacttypeIdent, __gong__toRawStringLiteral(onlinecontacttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", onlinecontacttypeIdent, __gong__toRawStringLiteral(onlinecontacttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", onlinecontacttypeIdent, __gong__toRawStringLiteral(onlinecontacttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", onlinecontacttypeIdent, __gong__toRawStringLiteral(onlinecontacttype.Description)))
			if onlinecontacttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := onlinecontacttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", onlinecontacttypeIdent, targetIdent))
			}
			for _, elem := range onlinecontacttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", onlinecontacttypeIdent, onlinecontacttypeIdent, targetIdent))
			}
			for _, elem := range onlinecontacttype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", onlinecontacttypeIdent, onlinecontacttypeIdent, targetIdent))
			}
			for _, elem := range onlinecontacttype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", onlinecontacttypeIdent, onlinecontacttypeIdent, targetIdent))
			}
			for _, elem := range onlinecontacttype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", onlinecontacttypeIdent, onlinecontacttypeIdent, targetIdent))
			}
			if onlinecontacttype.Network != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := onlinecontacttype.Network.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Network = %s", onlinecontacttypeIdent, targetIdent))
			}
			if onlinecontacttype.Linkage != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := onlinecontacttype.Linkage.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Linkage = %s", onlinecontacttypeIdent, targetIdent))
			}
			if onlinecontacttype.Protocol != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := onlinecontacttype.Protocol.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Protocol = %s", onlinecontacttypeIdent, targetIdent))
			}
			if onlinecontacttype.EMail != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := onlinecontacttype.EMail.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EMail = %s", onlinecontacttypeIdent, targetIdent))
			}
			for _, elem := range onlinecontacttype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", onlinecontacttypeIdent, onlinecontacttypeIdent, targetIdent))
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
		for _, pointpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PointPropertyTypes, stage.AixmfeaturesStage.PointPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			pointpropertytypeIdent := pointpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PointPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.NilReason)))
			if pointpropertytype.Point != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointpropertytype.Point.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Point = %s", pointpropertytypeIdent, targetIdent))
			}
			if pointpropertytype.ElevatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointpropertytype.ElevatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedPoint = %s", pointpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, pointreferencepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PointReferencePropertyTypes, stage.AixmfeaturesStage.PointReferencePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			pointreferencepropertytypeIdent := pointreferencepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PointReferencePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", pointreferencepropertytypeIdent, __gong__toRawStringLiteral(pointreferencepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointreferencepropertytypeIdent, __gong__toRawStringLiteral(pointreferencepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", pointreferencepropertytypeIdent, __gong__toRawStringLiteral(pointreferencepropertytype.NilReason)))
			if pointreferencepropertytype.PointReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointreferencepropertytype.PointReference.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointReference = %s", pointreferencepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, pointreferencetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PointReferenceTypes, stage.AixmfeaturesStage.PointReferenceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			pointreferencetypeIdent := pointreferencetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PointReferenceType{Name: %s}).Stage(stage.AixmfeaturesStage)", pointreferencetypeIdent, __gong__toRawStringLiteral(pointreferencetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointreferencetypeIdent, __gong__toRawStringLiteral(pointreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pointreferencetypeIdent, __gong__toRawStringLiteral(pointreferencetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pointreferencetypeIdent, __gong__toRawStringLiteral(pointreferencetype.Description)))
			if pointreferencetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointreferencetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", pointreferencetypeIdent, targetIdent))
			}
			for _, elem := range pointreferencetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", pointreferencetypeIdent, pointreferencetypeIdent, targetIdent))
			}
			if pointreferencetype.Role != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointreferencetype.Role.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Role = %s", pointreferencetypeIdent, targetIdent))
			}
			if pointreferencetype.PriorFixTolerance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointreferencetype.PriorFixTolerance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PriorFixTolerance = %s", pointreferencetypeIdent, targetIdent))
			}
			if pointreferencetype.PostFixTolerance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointreferencetype.PostFixTolerance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PostFixTolerance = %s", pointreferencetypeIdent, targetIdent))
			}
			if pointreferencetype.Point != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointreferencetype.Point.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Point = %s", pointreferencetypeIdent, targetIdent))
			}
			for _, elem := range pointreferencetype.FacilityAngle {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FacilityAngle = append(%s.FacilityAngle, %s)", pointreferencetypeIdent, pointreferencetypeIdent, targetIdent))
			}
			for _, elem := range pointreferencetype.FacilityDistance {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FacilityDistance = append(%s.FacilityDistance, %s)", pointreferencetypeIdent, pointreferencetypeIdent, targetIdent))
			}
			if pointreferencetype.FixToleranceArea != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pointreferencetype.FixToleranceArea.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FixToleranceArea = %s", pointreferencetypeIdent, targetIdent))
			}
			for _, elem := range pointreferencetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", pointreferencetypeIdent, pointreferencetypeIdent, targetIdent))
			}
			for _, elem := range pointreferencetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", pointreferencetypeIdent, pointreferencetypeIdent, targetIdent))
			}
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
		for _, postaladdresspropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PostalAddressPropertyTypes, stage.AixmfeaturesStage.PostalAddressPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			postaladdresspropertytypeIdent := postaladdresspropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PostalAddressPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", postaladdresspropertytypeIdent, __gong__toRawStringLiteral(postaladdresspropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", postaladdresspropertytypeIdent, __gong__toRawStringLiteral(postaladdresspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", postaladdresspropertytypeIdent, __gong__toRawStringLiteral(postaladdresspropertytype.NilReason)))
			if postaladdresspropertytype.PostalAddress != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := postaladdresspropertytype.PostalAddress.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PostalAddress = %s", postaladdresspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, postaladdresstype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PostalAddressTypes, stage.AixmfeaturesStage.PostalAddressType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			postaladdresstypeIdent := postaladdresstype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PostalAddressType{Name: %s}).Stage(stage.AixmfeaturesStage)", postaladdresstypeIdent, __gong__toRawStringLiteral(postaladdresstype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", postaladdresstypeIdent, __gong__toRawStringLiteral(postaladdresstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", postaladdresstypeIdent, __gong__toRawStringLiteral(postaladdresstype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", postaladdresstypeIdent, __gong__toRawStringLiteral(postaladdresstype.Description)))
			if postaladdresstype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := postaladdresstype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", postaladdresstypeIdent, targetIdent))
			}
			for _, elem := range postaladdresstype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", postaladdresstypeIdent, postaladdresstypeIdent, targetIdent))
			}
			for _, elem := range postaladdresstype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", postaladdresstypeIdent, postaladdresstypeIdent, targetIdent))
			}
			for _, elem := range postaladdresstype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", postaladdresstypeIdent, postaladdresstypeIdent, targetIdent))
			}
			for _, elem := range postaladdresstype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", postaladdresstypeIdent, postaladdresstypeIdent, targetIdent))
			}
			if postaladdresstype.DeliveryPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := postaladdresstype.DeliveryPoint.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliveryPoint = %s", postaladdresstypeIdent, targetIdent))
			}
			if postaladdresstype.City != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := postaladdresstype.City.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.City = %s", postaladdresstypeIdent, targetIdent))
			}
			if postaladdresstype.AdministrativeArea != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := postaladdresstype.AdministrativeArea.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AdministrativeArea = %s", postaladdresstypeIdent, targetIdent))
			}
			if postaladdresstype.PostalCode != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := postaladdresstype.PostalCode.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PostalCode = %s", postaladdresstypeIdent, targetIdent))
			}
			if postaladdresstype.Country != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := postaladdresstype.Country.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Country = %s", postaladdresstypeIdent, targetIdent))
			}
			for _, elem := range postaladdresstype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", postaladdresstypeIdent, postaladdresstypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radiofrequencyareapropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadioFrequencyAreaPropertyTypes, stage.AixmfeaturesStage.RadioFrequencyAreaPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radiofrequencyareapropertytypeIdent := radiofrequencyareapropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioFrequencyAreaPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", radiofrequencyareapropertytypeIdent, __gong__toRawStringLiteral(radiofrequencyareapropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, routeavailabilitypropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteAvailabilityPropertyTypes, stage.AixmfeaturesStage.RouteAvailabilityPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routeavailabilitypropertytypeIdent := routeavailabilitypropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteAvailabilityPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", routeavailabilitypropertytypeIdent, __gong__toRawStringLiteral(routeavailabilitypropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routeavailabilitypropertytypeIdent, __gong__toRawStringLiteral(routeavailabilitypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", routeavailabilitypropertytypeIdent, __gong__toRawStringLiteral(routeavailabilitypropertytype.NilReason)))
			if routeavailabilitypropertytype.RouteAvailability != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeavailabilitypropertytype.RouteAvailability.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteAvailability = %s", routeavailabilitypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routeavailabilitytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteAvailabilityTypes, stage.AixmfeaturesStage.RouteAvailabilityType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routeavailabilitytypeIdent := routeavailabilitytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteAvailabilityType{Name: %s}).Stage(stage.AixmfeaturesStage)", routeavailabilitytypeIdent, __gong__toRawStringLiteral(routeavailabilitytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routeavailabilitytypeIdent, __gong__toRawStringLiteral(routeavailabilitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routeavailabilitytypeIdent, __gong__toRawStringLiteral(routeavailabilitytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routeavailabilitytypeIdent, __gong__toRawStringLiteral(routeavailabilitytype.Description)))
			if routeavailabilitytype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeavailabilitytype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", routeavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range routeavailabilitytype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range routeavailabilitytype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range routeavailabilitytype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range routeavailabilitytype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
			if routeavailabilitytype.Direction != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeavailabilitytype.Direction.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", routeavailabilitytypeIdent, targetIdent))
			}
			if routeavailabilitytype.CardinalDirection != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeavailabilitytype.CardinalDirection.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CardinalDirection = %s", routeavailabilitytypeIdent, targetIdent))
			}
			if routeavailabilitytype.Status != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeavailabilitytype.Status.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Status = %s", routeavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range routeavailabilitytype.Levels {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Levels = append(%s.Levels, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range routeavailabilitytype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", routeavailabilitytypeIdent, routeavailabilitytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routedmetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteDMETimeSlicePropertyTypes, stage.AixmfeaturesStage.RouteDMETimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routedmetimeslicepropertytypeIdent := routedmetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteDMETimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", routedmetimeslicepropertytypeIdent, __gong__toRawStringLiteral(routedmetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routedmetimeslicepropertytypeIdent, __gong__toRawStringLiteral(routedmetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", routedmetimeslicepropertytypeIdent, routedmetimeslicepropertytype.Owns))
			if routedmetimeslicepropertytype.RouteDMETimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetimeslicepropertytype.RouteDMETimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteDMETimeSlice = %s", routedmetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routedmetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteDMETimeSliceTypes, stage.AixmfeaturesStage.RouteDMETimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routedmetimeslicetypeIdent := routedmetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteDMETimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", routedmetimeslicetypeIdent, __gong__toRawStringLiteral(routedmetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routedmetimeslicetypeIdent, __gong__toRawStringLiteral(routedmetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routedmetimeslicetypeIdent, __gong__toRawStringLiteral(routedmetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routedmetimeslicetypeIdent, __gong__toRawStringLiteral(routedmetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", routedmetimeslicetypeIdent, __gong__toRawStringLiteral(routedmetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", routedmetimeslicetypeIdent, routedmetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", routedmetimeslicetypeIdent, routedmetimeslicetype.CorrectionNumber))
			if routedmetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", routedmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routedmetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", routedmetimeslicetypeIdent, routedmetimeslicetypeIdent, targetIdent))
			}
			if routedmetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", routedmetimeslicetypeIdent, targetIdent))
			}
			if routedmetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", routedmetimeslicetypeIdent, targetIdent))
			}
			if routedmetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", routedmetimeslicetypeIdent, targetIdent))
			}
			if routedmetimeslicetype.CriticalDME != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetimeslicetype.CriticalDME.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CriticalDME = %s", routedmetimeslicetypeIdent, targetIdent))
			}
			if routedmetimeslicetype.Satisfactory != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetimeslicetype.Satisfactory.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Satisfactory = %s", routedmetimeslicetypeIdent, targetIdent))
			}
			if routedmetimeslicetype.ReferencedDME != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetimeslicetype.ReferencedDME.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedDME = %s", routedmetimeslicetypeIdent, targetIdent))
			}
			if routedmetimeslicetype.ApplicableRoutePortion != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetimeslicetype.ApplicableRoutePortion.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApplicableRoutePortion = %s", routedmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routedmetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", routedmetimeslicetypeIdent, routedmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routedmetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", routedmetimeslicetypeIdent, routedmetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routedmetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteDMETypes, stage.AixmfeaturesStage.RouteDMEType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routedmetypeIdent := routedmetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteDMEType{Name: %s}).Stage(stage.AixmfeaturesStage)", routedmetypeIdent, __gong__toRawStringLiteral(routedmetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routedmetypeIdent, __gong__toRawStringLiteral(routedmetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routedmetypeIdent, __gong__toRawStringLiteral(routedmetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routedmetypeIdent, __gong__toRawStringLiteral(routedmetype.Description)))
			if routedmetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", routedmetypeIdent, targetIdent))
			}
			for _, elem := range routedmetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", routedmetypeIdent, routedmetypeIdent, targetIdent))
			}
			if routedmetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", routedmetypeIdent, targetIdent))
			}
			if routedmetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routedmetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", routedmetypeIdent, targetIdent))
			}
			for _, elem := range routedmetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", routedmetypeIdent, routedmetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routeportionpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RoutePortionPropertyTypes, stage.AixmfeaturesStage.RoutePortionPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routeportionpropertytypeIdent := routeportionpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RoutePortionPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", routeportionpropertytypeIdent, __gong__toRawStringLiteral(routeportionpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routeportionpropertytypeIdent, __gong__toRawStringLiteral(routeportionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", routeportionpropertytypeIdent, __gong__toRawStringLiteral(routeportionpropertytype.NilReason)))
			if routeportionpropertytype.RoutePortion != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportionpropertytype.RoutePortion.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RoutePortion = %s", routeportionpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routeportiontype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RoutePortionTypes, stage.AixmfeaturesStage.RoutePortionType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routeportiontypeIdent := routeportiontype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RoutePortionType{Name: %s}).Stage(stage.AixmfeaturesStage)", routeportiontypeIdent, __gong__toRawStringLiteral(routeportiontype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routeportiontypeIdent, __gong__toRawStringLiteral(routeportiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routeportiontypeIdent, __gong__toRawStringLiteral(routeportiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routeportiontypeIdent, __gong__toRawStringLiteral(routeportiontype.Description)))
			if routeportiontype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", routeportiontypeIdent, targetIdent))
			}
			for _, elem := range routeportiontype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", routeportiontypeIdent, routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.Start_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_fixDesignatedPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.Start_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_navaidSystem = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.Start_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_position = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.Start_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_runwayPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.Start_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_aimingPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.Start_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.Start_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start_airportReferencePoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.IntermediatePoint_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_fixDesignatedPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.IntermediatePoint_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_navaidSystem = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.IntermediatePoint_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_position = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.IntermediatePoint_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_runwayPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.IntermediatePoint_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_aimingPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.IntermediatePoint_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.IntermediatePoint_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntermediatePoint_airportReferencePoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.ReferencedRoute != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.ReferencedRoute.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedRoute = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.End_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_fixDesignatedPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.End_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_navaidSystem = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.End_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_position = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.End_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_runwayPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.End_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_aimingPoint = %s", routeportiontypeIdent, targetIdent))
			}
			if routeportiontype.End_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routeportiontype.End_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End_airportReferencePoint = %s", routeportiontypeIdent, targetIdent))
			}
			for _, elem := range routeportiontype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", routeportiontypeIdent, routeportiontypeIdent, targetIdent))
			}
			for _, elem := range routeportiontype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", routeportiontypeIdent, routeportiontypeIdent, targetIdent))
			}
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
		for _, routesegmenttimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteSegmentTimeSlicePropertyTypes, stage.AixmfeaturesStage.RouteSegmentTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routesegmenttimeslicepropertytypeIdent := routesegmenttimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteSegmentTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", routesegmenttimeslicepropertytypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routesegmenttimeslicepropertytypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", routesegmenttimeslicepropertytypeIdent, routesegmenttimeslicepropertytype.Owns))
			if routesegmenttimeslicepropertytype.RouteSegmentTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicepropertytype.RouteSegmentTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteSegmentTimeSlice = %s", routesegmenttimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routesegmenttimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteSegmentTimeSliceTypes, stage.AixmfeaturesStage.RouteSegmentTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routesegmenttimeslicetypeIdent := routesegmenttimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteSegmentTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", routesegmenttimeslicetypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routesegmenttimeslicetypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routesegmenttimeslicetypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routesegmenttimeslicetypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", routesegmenttimeslicetypeIdent, __gong__toRawStringLiteral(routesegmenttimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", routesegmenttimeslicetypeIdent, routesegmenttimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", routesegmenttimeslicetypeIdent, routesegmenttimeslicetype.CorrectionNumber))
			if routesegmenttimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routesegmenttimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", routesegmenttimeslicetypeIdent, routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.Level != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.Level.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.UpperLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.UpperLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLimit = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.UpperLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.UpperLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLimitReference = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.LowerLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.LowerLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLimit = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.LowerLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.LowerLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLimitReference = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.MinimumObstacleClearanceAltitude != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.MinimumObstacleClearanceAltitude.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinimumObstacleClearanceAltitude = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.PathType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.PathType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PathType = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.TrueTrack != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.TrueTrack.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TrueTrack = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.MagneticTrack != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.MagneticTrack.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticTrack = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.ReverseTrueTrack != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.ReverseTrueTrack.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReverseTrueTrack = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.ReverseMagneticTrack != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.ReverseMagneticTrack.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReverseMagneticTrack = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.Length != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.Length.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Length = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.WidthLeft != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.WidthLeft.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WidthLeft = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.WidthRight != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.WidthRight.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WidthRight = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.TurnDirection != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.TurnDirection.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TurnDirection = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.SignalGap != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.SignalGap.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignalGap = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.MinimumEnrouteAltitude != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.MinimumEnrouteAltitude.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinimumEnrouteAltitude = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.MinimumCrossingAtEnd != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.MinimumCrossingAtEnd.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinimumCrossingAtEnd = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.MinimumCrossingAtEndReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.MinimumCrossingAtEndReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinimumCrossingAtEndReference = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.MaximumCrossingAtEnd != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.MaximumCrossingAtEnd.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MaximumCrossingAtEnd = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.MaximumCrossingAtEndReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.MaximumCrossingAtEndReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MaximumCrossingAtEndReference = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.NavigationType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.NavigationType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavigationType = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.RequiredNavigationPerformance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.RequiredNavigationPerformance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequiredNavigationPerformance = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.DesignatorSuffix != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.DesignatorSuffix.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatorSuffix = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.Start != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.Start.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.RouteFormed != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.RouteFormed.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteFormed = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.EvaluationArea != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.EvaluationArea.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EvaluationArea = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.CurveExtent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.CurveExtent.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CurveExtent = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			if routesegmenttimeslicetype.End != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttimeslicetype.End.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", routesegmenttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routesegmenttimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", routesegmenttimeslicetypeIdent, routesegmenttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routesegmenttimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", routesegmenttimeslicetypeIdent, routesegmenttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routesegmenttimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", routesegmenttimeslicetypeIdent, routesegmenttimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routesegmenttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteSegmentTypes, stage.AixmfeaturesStage.RouteSegmentType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routesegmenttypeIdent := routesegmenttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteSegmentType{Name: %s}).Stage(stage.AixmfeaturesStage)", routesegmenttypeIdent, __gong__toRawStringLiteral(routesegmenttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routesegmenttypeIdent, __gong__toRawStringLiteral(routesegmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routesegmenttypeIdent, __gong__toRawStringLiteral(routesegmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routesegmenttypeIdent, __gong__toRawStringLiteral(routesegmenttype.Description)))
			if routesegmenttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", routesegmenttypeIdent, targetIdent))
			}
			for _, elem := range routesegmenttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", routesegmenttypeIdent, routesegmenttypeIdent, targetIdent))
			}
			if routesegmenttype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", routesegmenttypeIdent, targetIdent))
			}
			if routesegmenttype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routesegmenttype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", routesegmenttypeIdent, targetIdent))
			}
			for _, elem := range routesegmenttype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", routesegmenttypeIdent, routesegmenttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteTimeSlicePropertyTypes, stage.AixmfeaturesStage.RouteTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routetimeslicepropertytypeIdent := routetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", routetimeslicepropertytypeIdent, __gong__toRawStringLiteral(routetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routetimeslicepropertytypeIdent, __gong__toRawStringLiteral(routetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", routetimeslicepropertytypeIdent, routetimeslicepropertytype.Owns))
			if routetimeslicepropertytype.RouteTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicepropertytype.RouteTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteTimeSlice = %s", routetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteTimeSliceTypes, stage.AixmfeaturesStage.RouteTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routetimeslicetypeIdent := routetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", routetimeslicetypeIdent, __gong__toRawStringLiteral(routetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routetimeslicetypeIdent, __gong__toRawStringLiteral(routetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routetimeslicetypeIdent, __gong__toRawStringLiteral(routetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routetimeslicetypeIdent, __gong__toRawStringLiteral(routetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", routetimeslicetypeIdent, __gong__toRawStringLiteral(routetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", routetimeslicetypeIdent, routetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", routetimeslicetypeIdent, routetimeslicetype.CorrectionNumber))
			if routetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", routetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", routetimeslicetypeIdent, routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.DesignatorPrefix != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.DesignatorPrefix.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatorPrefix = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.DesignatorSecondLetter != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.DesignatorSecondLetter.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatorSecondLetter = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.DesignatorNumber != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.DesignatorNumber.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatorNumber = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.MultipleIdentifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.MultipleIdentifier.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MultipleIdentifier = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.LocationDesignator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.LocationDesignator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LocationDesignator = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.FlightRule != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.FlightRule.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRule = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.InternationalUse != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.InternationalUse.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InternationalUse = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.MilitaryUse != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.MilitaryUse.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MilitaryUse = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.MilitaryTrainingType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.MilitaryTrainingType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MilitaryTrainingType = %s", routetimeslicetypeIdent, targetIdent))
			}
			if routetimeslicetype.UserOrganisation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetimeslicetype.UserOrganisation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UserOrganisation = %s", routetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", routetimeslicetypeIdent, routetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range routetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", routetimeslicetypeIdent, routetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, routetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RouteTypes, stage.AixmfeaturesStage.RouteType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			routetypeIdent := routetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RouteType{Name: %s}).Stage(stage.AixmfeaturesStage)", routetypeIdent, __gong__toRawStringLiteral(routetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routetypeIdent, __gong__toRawStringLiteral(routetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", routetypeIdent, __gong__toRawStringLiteral(routetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", routetypeIdent, __gong__toRawStringLiteral(routetype.Description)))
			if routetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", routetypeIdent, targetIdent))
			}
			for _, elem := range routetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", routetypeIdent, routetypeIdent, targetIdent))
			}
			if routetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", routetypeIdent, targetIdent))
			}
			if routetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := routetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", routetypeIdent, targetIdent))
			}
			for _, elem := range routetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", routetypeIdent, routetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, rulesprocedurestimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RulesProceduresTimeSlicePropertyTypes, stage.AixmfeaturesStage.RulesProceduresTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			rulesprocedurestimeslicepropertytypeIdent := rulesprocedurestimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RulesProceduresTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", rulesprocedurestimeslicepropertytypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rulesprocedurestimeslicepropertytypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", rulesprocedurestimeslicepropertytypeIdent, rulesprocedurestimeslicepropertytype.Owns))
			if rulesprocedurestimeslicepropertytype.RulesProceduresTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestimeslicepropertytype.RulesProceduresTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RulesProceduresTimeSlice = %s", rulesprocedurestimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, rulesprocedurestimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RulesProceduresTimeSliceTypes, stage.AixmfeaturesStage.RulesProceduresTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			rulesprocedurestimeslicetypeIdent := rulesprocedurestimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RulesProceduresTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", rulesprocedurestimeslicetypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rulesprocedurestimeslicetypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", rulesprocedurestimeslicetypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", rulesprocedurestimeslicetypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", rulesprocedurestimeslicetypeIdent, __gong__toRawStringLiteral(rulesprocedurestimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetype.CorrectionNumber))
			if rulesprocedurestimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			for _, elem := range rulesprocedurestimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			if rulesprocedurestimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			if rulesprocedurestimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			if rulesprocedurestimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			if rulesprocedurestimeslicetype.Category != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestimeslicetype.Category.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Category = %s", rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			if rulesprocedurestimeslicetype.Title != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestimeslicetype.Title.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title = %s", rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			if rulesprocedurestimeslicetype.Content != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestimeslicetype.Content.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Content = %s", rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			for _, elem := range rulesprocedurestimeslicetype.AffectedLocation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AffectedLocation = append(%s.AffectedLocation, %s)", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			for _, elem := range rulesprocedurestimeslicetype.AffectedArea {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AffectedArea = append(%s.AffectedArea, %s)", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			for _, elem := range rulesprocedurestimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetypeIdent, targetIdent))
			}
			for _, elem := range rulesprocedurestimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", rulesprocedurestimeslicetypeIdent, rulesprocedurestimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, rulesprocedurestype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RulesProceduresTypes, stage.AixmfeaturesStage.RulesProceduresType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			rulesprocedurestypeIdent := rulesprocedurestype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RulesProceduresType{Name: %s}).Stage(stage.AixmfeaturesStage)", rulesprocedurestypeIdent, __gong__toRawStringLiteral(rulesprocedurestype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rulesprocedurestypeIdent, __gong__toRawStringLiteral(rulesprocedurestype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", rulesprocedurestypeIdent, __gong__toRawStringLiteral(rulesprocedurestype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", rulesprocedurestypeIdent, __gong__toRawStringLiteral(rulesprocedurestype.Description)))
			if rulesprocedurestype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", rulesprocedurestypeIdent, targetIdent))
			}
			for _, elem := range rulesprocedurestype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", rulesprocedurestypeIdent, rulesprocedurestypeIdent, targetIdent))
			}
			if rulesprocedurestype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", rulesprocedurestypeIdent, targetIdent))
			}
			if rulesprocedurestype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := rulesprocedurestype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", rulesprocedurestypeIdent, targetIdent))
			}
			for _, elem := range rulesprocedurestype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", rulesprocedurestypeIdent, rulesprocedurestypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, runwaycentrelinepointpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RunwayCentrelinePointPropertyTypes, stage.AixmfeaturesStage.RunwayCentrelinePointPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			runwaycentrelinepointpropertytypeIdent := runwaycentrelinepointpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RunwayCentrelinePointPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", runwaycentrelinepointpropertytypeIdent, __gong__toRawStringLiteral(runwaycentrelinepointpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, standardinstrumentarrivalpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardInstrumentArrivalPropertyTypes, stage.AixmfeaturesStage.StandardInstrumentArrivalPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardinstrumentarrivalpropertytypeIdent := standardinstrumentarrivalpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardInstrumentArrivalPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardinstrumentarrivalpropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentarrivalpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
	if stage.AixmfeaturesStage != nil {
		for _, standardinstrumentdeparturepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardInstrumentDeparturePropertyTypes, stage.AixmfeaturesStage.StandardInstrumentDeparturePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardinstrumentdeparturepropertytypeIdent := standardinstrumentdeparturepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardInstrumentDeparturePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardinstrumentdeparturepropertytypeIdent, __gong__toRawStringLiteral(standardinstrumentdeparturepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
		for _, telephonecontactpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.TelephoneContactPropertyTypes, stage.AixmfeaturesStage.TelephoneContactPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			telephonecontactpropertytypeIdent := telephonecontactpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TelephoneContactPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", telephonecontactpropertytypeIdent, __gong__toRawStringLiteral(telephonecontactpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", telephonecontactpropertytypeIdent, __gong__toRawStringLiteral(telephonecontactpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", telephonecontactpropertytypeIdent, __gong__toRawStringLiteral(telephonecontactpropertytype.NilReason)))
			if telephonecontactpropertytype.TelephoneContact != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := telephonecontactpropertytype.TelephoneContact.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TelephoneContact = %s", telephonecontactpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, telephonecontacttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.TelephoneContactTypes, stage.AixmfeaturesStage.TelephoneContactType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			telephonecontacttypeIdent := telephonecontacttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TelephoneContactType{Name: %s}).Stage(stage.AixmfeaturesStage)", telephonecontacttypeIdent, __gong__toRawStringLiteral(telephonecontacttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", telephonecontacttypeIdent, __gong__toRawStringLiteral(telephonecontacttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", telephonecontacttypeIdent, __gong__toRawStringLiteral(telephonecontacttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", telephonecontacttypeIdent, __gong__toRawStringLiteral(telephonecontacttype.Description)))
			if telephonecontacttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := telephonecontacttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", telephonecontacttypeIdent, targetIdent))
			}
			for _, elem := range telephonecontacttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", telephonecontacttypeIdent, telephonecontacttypeIdent, targetIdent))
			}
			for _, elem := range telephonecontacttype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", telephonecontacttypeIdent, telephonecontacttypeIdent, targetIdent))
			}
			for _, elem := range telephonecontacttype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", telephonecontacttypeIdent, telephonecontacttypeIdent, targetIdent))
			}
			for _, elem := range telephonecontacttype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", telephonecontacttypeIdent, telephonecontacttypeIdent, targetIdent))
			}
			if telephonecontacttype.Voice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := telephonecontacttype.Voice.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Voice = %s", telephonecontacttypeIdent, targetIdent))
			}
			if telephonecontacttype.Facsimile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := telephonecontacttype.Facsimile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Facsimile = %s", telephonecontacttypeIdent, targetIdent))
			}
			for _, elem := range telephonecontacttype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", telephonecontacttypeIdent, telephonecontacttypeIdent, targetIdent))
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
		for _, touchdownliftoffpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.TouchDownLiftOffPropertyTypes, stage.AixmfeaturesStage.TouchDownLiftOffPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			touchdownliftoffpropertytypeIdent := touchdownliftoffpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TouchDownLiftOffPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", touchdownliftoffpropertytypeIdent, __gong__toRawStringLiteral(touchdownliftoffpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
			if basicmessagememberaixmpropertytype.Route != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.Route.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Route = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.RouteSegment != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.RouteSegment.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteSegment = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.RouteDME != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.RouteDME.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RouteDME = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.ChangeOverPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.ChangeOverPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ChangeOverPoint = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.FlightRestriction != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.FlightRestriction.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRestriction = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.AirspaceBorderCrossing != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.AirspaceBorderCrossing.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceBorderCrossing = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.RulesProcedures != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.RulesProcedures.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RulesProcedures = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.AerialRefuelling != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.AerialRefuelling.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerialRefuelling = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/xlink"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models\"", fmt.Sprintf("%q", modelsPackageName))
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
