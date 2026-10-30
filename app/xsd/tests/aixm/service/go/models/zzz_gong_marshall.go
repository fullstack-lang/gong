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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models"
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
		for _, codeaircraftgroundservicetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAircraftGroundServiceTypes, stage.AixmdatatypesStage.CodeAircraftGroundServiceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaircraftgroundservicetypeIdent := codeaircraftgroundservicetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAircraftGroundServiceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaircraftgroundservicetypeIdent, __gong__toRawStringLiteral(codeaircraftgroundservicetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaircraftgroundservicetypeIdent, __gong__toRawStringLiteral(codeaircraftgroundservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaircraftgroundservicetypeIdent, __gong__toRawStringLiteral(codeaircraftgroundservicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaircraftgroundservicetypeIdent, __gong__toRawStringLiteral(codeaircraftgroundservicetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeaviationstandardstype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAviationStandardsTypes, stage.AixmdatatypesStage.CodeAviationStandardsType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeaviationstandardstypeIdent := codeaviationstandardstype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAviationStandardsType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeaviationstandardstypeIdent, __gong__toRawStringLiteral(codeaviationstandardstype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeaviationstandardstypeIdent, __gong__toRawStringLiteral(codeaviationstandardstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeaviationstandardstypeIdent, __gong__toRawStringLiteral(codeaviationstandardstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeaviationstandardstypeIdent, __gong__toRawStringLiteral(codeaviationstandardstype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codecommunicationchanneltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeCommunicationChannelTypes, stage.AixmdatatypesStage.CodeCommunicationChannelType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codecommunicationchanneltypeIdent := codecommunicationchanneltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCommunicationChannelType{Name: %s}).Stage(stage.AixmdatatypesStage)", codecommunicationchanneltypeIdent, __gong__toRawStringLiteral(codecommunicationchanneltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecommunicationchanneltypeIdent, __gong__toRawStringLiteral(codecommunicationchanneltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecommunicationchanneltypeIdent, __gong__toRawStringLiteral(codecommunicationchanneltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecommunicationchanneltypeIdent, __gong__toRawStringLiteral(codecommunicationchanneltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codecommunicationdirectiontype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeCommunicationDirectionTypes, stage.AixmdatatypesStage.CodeCommunicationDirectionType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codecommunicationdirectiontypeIdent := codecommunicationdirectiontype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCommunicationDirectionType{Name: %s}).Stage(stage.AixmdatatypesStage)", codecommunicationdirectiontypeIdent, __gong__toRawStringLiteral(codecommunicationdirectiontype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecommunicationdirectiontypeIdent, __gong__toRawStringLiteral(codecommunicationdirectiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecommunicationdirectiontypeIdent, __gong__toRawStringLiteral(codecommunicationdirectiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecommunicationdirectiontypeIdent, __gong__toRawStringLiteral(codecommunicationdirectiontype.EnclosedText)))
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
		for _, codefacilityrankingtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFacilityRankingTypes, stage.AixmdatatypesStage.CodeFacilityRankingType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codefacilityrankingtypeIdent := codefacilityrankingtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFacilityRankingType{Name: %s}).Stage(stage.AixmdatatypesStage)", codefacilityrankingtypeIdent, __gong__toRawStringLiteral(codefacilityrankingtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefacilityrankingtypeIdent, __gong__toRawStringLiteral(codefacilityrankingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefacilityrankingtypeIdent, __gong__toRawStringLiteral(codefacilityrankingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefacilityrankingtypeIdent, __gong__toRawStringLiteral(codefacilityrankingtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codefirefightingtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFireFightingTypes, stage.AixmdatatypesStage.CodeFireFightingType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codefirefightingtypeIdent := codefirefightingtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFireFightingType{Name: %s}).Stage(stage.AixmdatatypesStage)", codefirefightingtypeIdent, __gong__toRawStringLiteral(codefirefightingtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefirefightingtypeIdent, __gong__toRawStringLiteral(codefirefightingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefirefightingtypeIdent, __gong__toRawStringLiteral(codefirefightingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefirefightingtypeIdent, __gong__toRawStringLiteral(codefirefightingtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeflightdestinationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFlightDestinationTypes, stage.AixmdatatypesStage.CodeFlightDestinationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeflightdestinationtypeIdent := codeflightdestinationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFlightDestinationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeflightdestinationtypeIdent, __gong__toRawStringLiteral(codeflightdestinationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeflightdestinationtypeIdent, __gong__toRawStringLiteral(codeflightdestinationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeflightdestinationtypeIdent, __gong__toRawStringLiteral(codeflightdestinationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeflightdestinationtypeIdent, __gong__toRawStringLiteral(codeflightdestinationtype.EnclosedText)))
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
		for _, codefueltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeFuelTypes, stage.AixmdatatypesStage.CodeFuelType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codefueltypeIdent := codefueltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeFuelType{Name: %s}).Stage(stage.AixmdatatypesStage)", codefueltypeIdent, __gong__toRawStringLiteral(codefueltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codefueltypeIdent, __gong__toRawStringLiteral(codefueltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codefueltypeIdent, __gong__toRawStringLiteral(codefueltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codefueltypeIdent, __gong__toRawStringLiteral(codefueltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeintensitystandbytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeIntensityStandByTypes, stage.AixmdatatypesStage.CodeIntensityStandByType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeintensitystandbytypeIdent := codeintensitystandbytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeIntensityStandByType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeintensitystandbytypeIdent, __gong__toRawStringLiteral(codeintensitystandbytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeintensitystandbytypeIdent, __gong__toRawStringLiteral(codeintensitystandbytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeintensitystandbytypeIdent, __gong__toRawStringLiteral(codeintensitystandbytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeintensitystandbytypeIdent, __gong__toRawStringLiteral(codeintensitystandbytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codelanguagetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeLanguageTypes, stage.AixmdatatypesStage.CodeLanguageType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codelanguagetypeIdent := codelanguagetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLanguageType{Name: %s}).Stage(stage.AixmdatatypesStage)", codelanguagetypeIdent, __gong__toRawStringLiteral(codelanguagetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelanguagetypeIdent, __gong__toRawStringLiteral(codelanguagetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelanguagetypeIdent, __gong__toRawStringLiteral(codelanguagetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelanguagetypeIdent, __gong__toRawStringLiteral(codelanguagetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codelevelseriestype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeLevelSeriesTypes, stage.AixmdatatypesStage.CodeLevelSeriesType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codelevelseriestypeIdent := codelevelseriestype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLevelSeriesType{Name: %s}).Stage(stage.AixmdatatypesStage)", codelevelseriestypeIdent, __gong__toRawStringLiteral(codelevelseriestype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelevelseriestypeIdent, __gong__toRawStringLiteral(codelevelseriestype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codelevelseriestypeIdent, __gong__toRawStringLiteral(codelevelseriestype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelevelseriestypeIdent, __gong__toRawStringLiteral(codelevelseriestype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeleveltabledesignatortype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeLevelTableDesignatorTypes, stage.AixmdatatypesStage.CodeLevelTableDesignatorType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeleveltabledesignatortypeIdent := codeleveltabledesignatortype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeLevelTableDesignatorType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeleveltabledesignatortypeIdent, __gong__toRawStringLiteral(codeleveltabledesignatortype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeleveltabledesignatortypeIdent, __gong__toRawStringLiteral(codeleveltabledesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeleveltabledesignatortypeIdent, __gong__toRawStringLiteral(codeleveltabledesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeleveltabledesignatortypeIdent, __gong__toRawStringLiteral(codeleveltabledesignatortype.EnclosedText)))
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
		for _, codenitrogentype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeNitrogenTypes, stage.AixmdatatypesStage.CodeNitrogenType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codenitrogentypeIdent := codenitrogentype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNitrogenType{Name: %s}).Stage(stage.AixmdatatypesStage)", codenitrogentypeIdent, __gong__toRawStringLiteral(codenitrogentype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenitrogentypeIdent, __gong__toRawStringLiteral(codenitrogentype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenitrogentypeIdent, __gong__toRawStringLiteral(codenitrogentype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenitrogentypeIdent, __gong__toRawStringLiteral(codenitrogentype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codenorthreferencetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeNorthReferenceTypes, stage.AixmdatatypesStage.CodeNorthReferenceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codenorthreferencetypeIdent := codenorthreferencetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNorthReferenceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codenorthreferencetypeIdent, __gong__toRawStringLiteral(codenorthreferencetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenorthreferencetypeIdent, __gong__toRawStringLiteral(codenorthreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenorthreferencetypeIdent, __gong__toRawStringLiteral(codenorthreferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenorthreferencetypeIdent, __gong__toRawStringLiteral(codenorthreferencetype.EnclosedText)))
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
		for _, codeoiltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeOilTypes, stage.AixmdatatypesStage.CodeOilType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeoiltypeIdent := codeoiltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOilType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeoiltypeIdent, __gong__toRawStringLiteral(codeoiltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeoiltypeIdent, __gong__toRawStringLiteral(codeoiltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeoiltypeIdent, __gong__toRawStringLiteral(codeoiltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeoiltypeIdent, __gong__toRawStringLiteral(codeoiltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeorganisationdesignatortype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeOrganisationDesignatorTypes, stage.AixmdatatypesStage.CodeOrganisationDesignatorType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeorganisationdesignatortypeIdent := codeorganisationdesignatortype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOrganisationDesignatorType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeorganisationdesignatortypeIdent, __gong__toRawStringLiteral(codeorganisationdesignatortype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeorganisationdesignatortypeIdent, __gong__toRawStringLiteral(codeorganisationdesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeorganisationdesignatortypeIdent, __gong__toRawStringLiteral(codeorganisationdesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeorganisationdesignatortypeIdent, __gong__toRawStringLiteral(codeorganisationdesignatortype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeorganisationhierarchytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeOrganisationHierarchyTypes, stage.AixmdatatypesStage.CodeOrganisationHierarchyType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeorganisationhierarchytypeIdent := codeorganisationhierarchytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOrganisationHierarchyType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeorganisationhierarchytypeIdent, __gong__toRawStringLiteral(codeorganisationhierarchytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeorganisationhierarchytypeIdent, __gong__toRawStringLiteral(codeorganisationhierarchytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeorganisationhierarchytypeIdent, __gong__toRawStringLiteral(codeorganisationhierarchytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeorganisationhierarchytypeIdent, __gong__toRawStringLiteral(codeorganisationhierarchytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeorganisationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeOrganisationTypes, stage.AixmdatatypesStage.CodeOrganisationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeorganisationtypeIdent := codeorganisationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOrganisationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeorganisationtypeIdent, __gong__toRawStringLiteral(codeorganisationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeorganisationtypeIdent, __gong__toRawStringLiteral(codeorganisationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeorganisationtypeIdent, __gong__toRawStringLiteral(codeorganisationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeorganisationtypeIdent, __gong__toRawStringLiteral(codeorganisationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeoxygentype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeOxygenTypes, stage.AixmdatatypesStage.CodeOxygenType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeoxygentypeIdent := codeoxygentype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeOxygenType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeoxygentypeIdent, __gong__toRawStringLiteral(codeoxygentype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeoxygentypeIdent, __gong__toRawStringLiteral(codeoxygentype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeoxygentypeIdent, __gong__toRawStringLiteral(codeoxygentype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeoxygentypeIdent, __gong__toRawStringLiteral(codeoxygentype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codepassengerservicetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodePassengerServiceTypes, stage.AixmdatatypesStage.CodePassengerServiceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codepassengerservicetypeIdent := codepassengerservicetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePassengerServiceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codepassengerservicetypeIdent, __gong__toRawStringLiteral(codepassengerservicetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepassengerservicetypeIdent, __gong__toRawStringLiteral(codepassengerservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepassengerservicetypeIdent, __gong__toRawStringLiteral(codepassengerservicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepassengerservicetypeIdent, __gong__toRawStringLiteral(codepassengerservicetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codepilotcontrolledlightingtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodePilotControlledLightingTypes, stage.AixmdatatypesStage.CodePilotControlledLightingType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codepilotcontrolledlightingtypeIdent := codepilotcontrolledlightingtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePilotControlledLightingType{Name: %s}).Stage(stage.AixmdatatypesStage)", codepilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(codepilotcontrolledlightingtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(codepilotcontrolledlightingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(codepilotcontrolledlightingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(codepilotcontrolledlightingtype.EnclosedText)))
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
		for _, coderadioemissiontype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRadioEmissionTypes, stage.AixmdatatypesStage.CodeRadioEmissionType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderadioemissiontypeIdent := coderadioemissiontype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRadioEmissionType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderadioemissiontypeIdent, __gong__toRawStringLiteral(coderadioemissiontype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderadioemissiontypeIdent, __gong__toRawStringLiteral(coderadioemissiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderadioemissiontypeIdent, __gong__toRawStringLiteral(coderadioemissiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderadioemissiontypeIdent, __gong__toRawStringLiteral(coderadioemissiontype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeserviceatctype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeServiceATCTypes, stage.AixmdatatypesStage.CodeServiceATCType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeserviceatctypeIdent := codeserviceatctype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeServiceATCType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeserviceatctypeIdent, __gong__toRawStringLiteral(codeserviceatctype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeserviceatctypeIdent, __gong__toRawStringLiteral(codeserviceatctype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeserviceatctypeIdent, __gong__toRawStringLiteral(codeserviceatctype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeserviceatctypeIdent, __gong__toRawStringLiteral(codeserviceatctype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeserviceatfmtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeServiceATFMTypes, stage.AixmdatatypesStage.CodeServiceATFMType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeserviceatfmtypeIdent := codeserviceatfmtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeServiceATFMType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeserviceatfmtypeIdent, __gong__toRawStringLiteral(codeserviceatfmtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeserviceatfmtypeIdent, __gong__toRawStringLiteral(codeserviceatfmtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeserviceatfmtypeIdent, __gong__toRawStringLiteral(codeserviceatfmtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeserviceatfmtypeIdent, __gong__toRawStringLiteral(codeserviceatfmtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeservicegroundcontroltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeServiceGroundControlTypes, stage.AixmdatatypesStage.CodeServiceGroundControlType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeservicegroundcontroltypeIdent := codeservicegroundcontroltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeServiceGroundControlType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeservicegroundcontroltypeIdent, __gong__toRawStringLiteral(codeservicegroundcontroltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeservicegroundcontroltypeIdent, __gong__toRawStringLiteral(codeservicegroundcontroltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeservicegroundcontroltypeIdent, __gong__toRawStringLiteral(codeservicegroundcontroltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeservicegroundcontroltypeIdent, __gong__toRawStringLiteral(codeservicegroundcontroltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeserviceinformationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeServiceInformationTypes, stage.AixmdatatypesStage.CodeServiceInformationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeserviceinformationtypeIdent := codeserviceinformationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeServiceInformationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeserviceinformationtypeIdent, __gong__toRawStringLiteral(codeserviceinformationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeserviceinformationtypeIdent, __gong__toRawStringLiteral(codeserviceinformationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeserviceinformationtypeIdent, __gong__toRawStringLiteral(codeserviceinformationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeserviceinformationtypeIdent, __gong__toRawStringLiteral(codeserviceinformationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeservicesartype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeServiceSARTypes, stage.AixmdatatypesStage.CodeServiceSARType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeservicesartypeIdent := codeservicesartype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeServiceSARType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeservicesartypeIdent, __gong__toRawStringLiteral(codeservicesartype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeservicesartypeIdent, __gong__toRawStringLiteral(codeservicesartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeservicesartypeIdent, __gong__toRawStringLiteral(codeservicesartype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeservicesartypeIdent, __gong__toRawStringLiteral(codeservicesartype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codespecialdatetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeSpecialDateTypes, stage.AixmdatatypesStage.CodeSpecialDateType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codespecialdatetypeIdent := codespecialdatetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSpecialDateType{Name: %s}).Stage(stage.AixmdatatypesStage)", codespecialdatetypeIdent, __gong__toRawStringLiteral(codespecialdatetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codespecialdatetypeIdent, __gong__toRawStringLiteral(codespecialdatetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codespecialdatetypeIdent, __gong__toRawStringLiteral(codespecialdatetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codespecialdatetypeIdent, __gong__toRawStringLiteral(codespecialdatetype.EnclosedText)))
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
		for _, codestatusservicetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeStatusServiceTypes, stage.AixmdatatypesStage.CodeStatusServiceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codestatusservicetypeIdent := codestatusservicetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusServiceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codestatusservicetypeIdent, __gong__toRawStringLiteral(codestatusservicetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusservicetypeIdent, __gong__toRawStringLiteral(codestatusservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusservicetypeIdent, __gong__toRawStringLiteral(codestatusservicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusservicetypeIdent, __gong__toRawStringLiteral(codestatusservicetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codesystemactivationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeSystemActivationTypes, stage.AixmdatatypesStage.CodeSystemActivationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codesystemactivationtypeIdent := codesystemactivationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSystemActivationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codesystemactivationtypeIdent, __gong__toRawStringLiteral(codesystemactivationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesystemactivationtypeIdent, __gong__toRawStringLiteral(codesystemactivationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesystemactivationtypeIdent, __gong__toRawStringLiteral(codesystemactivationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesystemactivationtypeIdent, __gong__toRawStringLiteral(codesystemactivationtype.EnclosedText)))
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
		for _, codeunitdependencytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeUnitDependencyTypes, stage.AixmdatatypesStage.CodeUnitDependencyType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeunitdependencytypeIdent := codeunitdependencytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeUnitDependencyType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeunitdependencytypeIdent, __gong__toRawStringLiteral(codeunitdependencytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeunitdependencytypeIdent, __gong__toRawStringLiteral(codeunitdependencytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeunitdependencytypeIdent, __gong__toRawStringLiteral(codeunitdependencytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeunitdependencytypeIdent, __gong__toRawStringLiteral(codeunitdependencytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeunittype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeUnitTypes, stage.AixmdatatypesStage.CodeUnitType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeunittypeIdent := codeunittype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeUnitType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeunittypeIdent, __gong__toRawStringLiteral(codeunittype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeunittypeIdent, __gong__toRawStringLiteral(codeunittype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeunittypeIdent, __gong__toRawStringLiteral(codeunittype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeunittypeIdent, __gong__toRawStringLiteral(codeunittype.EnclosedText)))
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
		for _, dateyeartype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.DateYearTypes, stage.AixmdatatypesStage.DateYearType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			dateyeartypeIdent := dateyeartype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.DateYearType{Name: %s}).Stage(stage.AixmdatatypesStage)", dateyeartypeIdent, __gong__toRawStringLiteral(dateyeartype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dateyeartypeIdent, __gong__toRawStringLiteral(dateyeartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", dateyeartypeIdent, __gong__toRawStringLiteral(dateyeartype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", dateyeartypeIdent, __gong__toRawStringLiteral(dateyeartype.EnclosedText)))
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
		for _, valfrequencytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValFrequencyTypes, stage.AixmdatatypesStage.ValFrequencyType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valfrequencytypeIdent := valfrequencytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValFrequencyType{Name: %s}).Stage(stage.AixmdatatypesStage)", valfrequencytypeIdent, __gong__toRawStringLiteral(valfrequencytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valfrequencytypeIdent, __gong__toRawStringLiteral(valfrequencytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", valfrequencytypeIdent, __gong__toRawStringLiteral(valfrequencytype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valfrequencytypeIdent, __gong__toRawStringLiteral(valfrequencytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valfrequencytypeIdent, __gong__toRawStringLiteral(valfrequencytype.EnclosedText)))
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
			for _, elem := range a_extension.AbstractLightActivationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractLightActivationExtension = append(%s.AbstractLightActivationExtension, %s)", a_extensionIdent, a_extensionIdent, targetIdent))
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
			for _, elem := range a_extension_10.AbstractAircraftGroundServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAircraftGroundServiceExtension = append(%s.AbstractAircraftGroundServiceExtension, %s)", a_extension_10Ident, a_extension_10Ident, targetIdent))
			}
			for _, elem := range a_extension_10.AbstractAirportGroundServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirportGroundServiceExtension = append(%s.AbstractAirportGroundServiceExtension, %s)", a_extension_10Ident, a_extension_10Ident, targetIdent))
			}
			for _, elem := range a_extension_10.AbstractServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceExtension = append(%s.AbstractServiceExtension, %s)", a_extension_10Ident, a_extension_10Ident, targetIdent))
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
			for _, elem := range a_extension_11.AbstractAirportClearanceServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirportClearanceServiceExtension = append(%s.AbstractAirportClearanceServiceExtension, %s)", a_extension_11Ident, a_extension_11Ident, targetIdent))
			}
			for _, elem := range a_extension_11.AbstractAirportGroundServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirportGroundServiceExtension = append(%s.AbstractAirportGroundServiceExtension, %s)", a_extension_11Ident, a_extension_11Ident, targetIdent))
			}
			for _, elem := range a_extension_11.AbstractServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceExtension = append(%s.AbstractServiceExtension, %s)", a_extension_11Ident, a_extension_11Ident, targetIdent))
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
			for _, elem := range a_extension_12.AbstractAirportSuppliesServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirportSuppliesServiceExtension = append(%s.AbstractAirportSuppliesServiceExtension, %s)", a_extension_12Ident, a_extension_12Ident, targetIdent))
			}
			for _, elem := range a_extension_12.AbstractAirportGroundServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirportGroundServiceExtension = append(%s.AbstractAirportGroundServiceExtension, %s)", a_extension_12Ident, a_extension_12Ident, targetIdent))
			}
			for _, elem := range a_extension_12.AbstractServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceExtension = append(%s.AbstractServiceExtension, %s)", a_extension_12Ident, a_extension_12Ident, targetIdent))
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
			for _, elem := range a_extension_13.AbstractAirTrafficControlServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirTrafficControlServiceExtension = append(%s.AbstractAirTrafficControlServiceExtension, %s)", a_extension_13Ident, a_extension_13Ident, targetIdent))
			}
			for _, elem := range a_extension_13.AbstractTrafficSeparationServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractTrafficSeparationServiceExtension = append(%s.AbstractTrafficSeparationServiceExtension, %s)", a_extension_13Ident, a_extension_13Ident, targetIdent))
			}
			for _, elem := range a_extension_13.AbstractServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceExtension = append(%s.AbstractServiceExtension, %s)", a_extension_13Ident, a_extension_13Ident, targetIdent))
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
			for _, elem := range a_extension_14.AbstractAirTrafficManagementServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirTrafficManagementServiceExtension = append(%s.AbstractAirTrafficManagementServiceExtension, %s)", a_extension_14Ident, a_extension_14Ident, targetIdent))
			}
			for _, elem := range a_extension_14.AbstractServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceExtension = append(%s.AbstractServiceExtension, %s)", a_extension_14Ident, a_extension_14Ident, targetIdent))
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
			for _, elem := range a_extension_15.AbstractCallsignDetailExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractCallsignDetailExtension = append(%s.AbstractCallsignDetailExtension, %s)", a_extension_15Ident, a_extension_15Ident, targetIdent))
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
			for _, elem := range a_extension_16.AbstractFireFightingServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractFireFightingServiceExtension = append(%s.AbstractFireFightingServiceExtension, %s)", a_extension_16Ident, a_extension_16Ident, targetIdent))
			}
			for _, elem := range a_extension_16.AbstractAirportGroundServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirportGroundServiceExtension = append(%s.AbstractAirportGroundServiceExtension, %s)", a_extension_16Ident, a_extension_16Ident, targetIdent))
			}
			for _, elem := range a_extension_16.AbstractServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceExtension = append(%s.AbstractServiceExtension, %s)", a_extension_16Ident, a_extension_16Ident, targetIdent))
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
			for _, elem := range a_extension_17.AbstractFuelExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractFuelExtension = append(%s.AbstractFuelExtension, %s)", a_extension_17Ident, a_extension_17Ident, targetIdent))
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
			for _, elem := range a_extension_18.AbstractGroundTrafficControlServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractGroundTrafficControlServiceExtension = append(%s.AbstractGroundTrafficControlServiceExtension, %s)", a_extension_18Ident, a_extension_18Ident, targetIdent))
			}
			for _, elem := range a_extension_18.AbstractTrafficSeparationServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractTrafficSeparationServiceExtension = append(%s.AbstractTrafficSeparationServiceExtension, %s)", a_extension_18Ident, a_extension_18Ident, targetIdent))
			}
			for _, elem := range a_extension_18.AbstractServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceExtension = append(%s.AbstractServiceExtension, %s)", a_extension_18Ident, a_extension_18Ident, targetIdent))
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
			for _, elem := range a_extension_19.AbstractInformationServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractInformationServiceExtension = append(%s.AbstractInformationServiceExtension, %s)", a_extension_19Ident, a_extension_19Ident, targetIdent))
			}
			for _, elem := range a_extension_19.AbstractServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceExtension = append(%s.AbstractServiceExtension, %s)", a_extension_19Ident, a_extension_19Ident, targetIdent))
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
			for _, elem := range a_extension_2.AbstractLinguisticNoteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractLinguisticNoteExtension = append(%s.AbstractLinguisticNoteExtension, %s)", a_extension_2Ident, a_extension_2Ident, targetIdent))
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
			for _, elem := range a_extension_20.AbstractNitrogenExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNitrogenExtension = append(%s.AbstractNitrogenExtension, %s)", a_extension_20Ident, a_extension_20Ident, targetIdent))
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
			for _, elem := range a_extension_21.AbstractOilExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractOilExtension = append(%s.AbstractOilExtension, %s)", a_extension_21Ident, a_extension_21Ident, targetIdent))
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
			for _, elem := range a_extension_22.AbstractOxygenExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractOxygenExtension = append(%s.AbstractOxygenExtension, %s)", a_extension_22Ident, a_extension_22Ident, targetIdent))
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
			for _, elem := range a_extension_23.AbstractPassengerServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPassengerServiceExtension = append(%s.AbstractPassengerServiceExtension, %s)", a_extension_23Ident, a_extension_23Ident, targetIdent))
			}
			for _, elem := range a_extension_23.AbstractAirportGroundServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAirportGroundServiceExtension = append(%s.AbstractAirportGroundServiceExtension, %s)", a_extension_23Ident, a_extension_23Ident, targetIdent))
			}
			for _, elem := range a_extension_23.AbstractServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceExtension = append(%s.AbstractServiceExtension, %s)", a_extension_23Ident, a_extension_23Ident, targetIdent))
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
			for _, elem := range a_extension_24.AbstractPilotControlledLightingExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPilotControlledLightingExtension = append(%s.AbstractPilotControlledLightingExtension, %s)", a_extension_24Ident, a_extension_24Ident, targetIdent))
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
			for _, elem := range a_extension_25.AbstractRadioCommunicationChannelExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRadioCommunicationChannelExtension = append(%s.AbstractRadioCommunicationChannelExtension, %s)", a_extension_25Ident, a_extension_25Ident, targetIdent))
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
			for _, elem := range a_extension_26.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_26Ident, a_extension_26Ident, targetIdent))
			}
			for _, elem := range a_extension_26.AbstractRadioCommunicationOperationalStatusExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRadioCommunicationOperationalStatusExtension = append(%s.AbstractRadioCommunicationOperationalStatusExtension, %s)", a_extension_26Ident, a_extension_26Ident, targetIdent))
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
			for _, elem := range a_extension_27.AbstractSearchRescueServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSearchRescueServiceExtension = append(%s.AbstractSearchRescueServiceExtension, %s)", a_extension_27Ident, a_extension_27Ident, targetIdent))
			}
			for _, elem := range a_extension_27.AbstractServiceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceExtension = append(%s.AbstractServiceExtension, %s)", a_extension_27Ident, a_extension_27Ident, targetIdent))
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
			for _, elem := range a_extension_28.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_28Ident, a_extension_28Ident, targetIdent))
			}
			for _, elem := range a_extension_28.AbstractServiceOperationalStatusExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractServiceOperationalStatusExtension = append(%s.AbstractServiceOperationalStatusExtension, %s)", a_extension_28Ident, a_extension_28Ident, targetIdent))
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
			for _, elem := range a_extension_29.AbstractContactInformationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractContactInformationExtension = append(%s.AbstractContactInformationExtension, %s)", a_extension_29Ident, a_extension_29Ident, targetIdent))
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
			for _, elem := range a_extension_3.AbstractNoteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNoteExtension = append(%s.AbstractNoteExtension, %s)", a_extension_3Ident, a_extension_3Ident, targetIdent))
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
			for _, elem := range a_extension_30.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_30Ident, a_extension_30Ident, targetIdent))
			}
			for _, elem := range a_extension_30.AbstractOnlineContactExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractOnlineContactExtension = append(%s.AbstractOnlineContactExtension, %s)", a_extension_30Ident, a_extension_30Ident, targetIdent))
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
			for _, elem := range a_extension_31.AbstractPostalAddressExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPostalAddressExtension = append(%s.AbstractPostalAddressExtension, %s)", a_extension_31Ident, a_extension_31Ident, targetIdent))
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
			for _, elem := range a_extension_32.AbstractTelephoneContactExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractTelephoneContactExtension = append(%s.AbstractTelephoneContactExtension, %s)", a_extension_32Ident, a_extension_32Ident, targetIdent))
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
			for _, elem := range a_extension_33.AbstractSpecialDateExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSpecialDateExtension = append(%s.AbstractSpecialDateExtension, %s)", a_extension_33Ident, a_extension_33Ident, targetIdent))
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
			for _, elem := range a_extension_34.AbstractTimesheetExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractTimesheetExtension = append(%s.AbstractTimesheetExtension, %s)", a_extension_34Ident, a_extension_34Ident, targetIdent))
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
			for _, elem := range a_extension_35.AbstractStandardLevelExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractStandardLevelExtension = append(%s.AbstractStandardLevelExtension, %s)", a_extension_35Ident, a_extension_35Ident, targetIdent))
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
			for _, elem := range a_extension_36.AbstractStandardLevelColumnExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractStandardLevelColumnExtension = append(%s.AbstractStandardLevelColumnExtension, %s)", a_extension_36Ident, a_extension_36Ident, targetIdent))
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
			for _, elem := range a_extension_37.AbstractStandardLevelSectorExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractStandardLevelSectorExtension = append(%s.AbstractStandardLevelSectorExtension, %s)", a_extension_37Ident, a_extension_37Ident, targetIdent))
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
			for _, elem := range a_extension_38.AbstractStandardLevelTableExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractStandardLevelTableExtension = append(%s.AbstractStandardLevelTableExtension, %s)", a_extension_38Ident, a_extension_38Ident, targetIdent))
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
			for _, elem := range a_extension_4.AbstractOrganisationAuthorityExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractOrganisationAuthorityExtension = append(%s.AbstractOrganisationAuthorityExtension, %s)", a_extension_4Ident, a_extension_4Ident, targetIdent))
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
			for _, elem := range a_extension_5.AbstractOrganisationAuthorityAssociationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractOrganisationAuthorityAssociationExtension = append(%s.AbstractOrganisationAuthorityAssociationExtension, %s)", a_extension_5Ident, a_extension_5Ident, targetIdent))
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
			for _, elem := range a_extension_6.AbstractUnitExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractUnitExtension = append(%s.AbstractUnitExtension, %s)", a_extension_6Ident, a_extension_6Ident, targetIdent))
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
			for _, elem := range a_extension_7.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_7Ident, a_extension_7Ident, targetIdent))
			}
			for _, elem := range a_extension_7.AbstractUnitAvailabilityExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractUnitAvailabilityExtension = append(%s.AbstractUnitAvailabilityExtension, %s)", a_extension_7Ident, a_extension_7Ident, targetIdent))
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
			for _, elem := range a_extension_8.AbstractUnitDependencyExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractUnitDependencyExtension = append(%s.AbstractUnitDependencyExtension, %s)", a_extension_8Ident, a_extension_8Ident, targetIdent))
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
			for _, elem := range a_extension_9.AbstractRoutePortionExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRoutePortionExtension = append(%s.AbstractRoutePortionExtension, %s)", a_extension_9Ident, a_extension_9Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, abstractairportgroundservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AbstractAirportGroundServiceTypes, stage.AixmfeaturesStage.AbstractAirportGroundServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			abstractairportgroundservicetypeIdent := abstractairportgroundservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractAirportGroundServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", abstractairportgroundservicetypeIdent, __gong__toRawStringLiteral(abstractairportgroundservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractairportgroundservicetypeIdent, __gong__toRawStringLiteral(abstractairportgroundservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractairportgroundservicetypeIdent, __gong__toRawStringLiteral(abstractairportgroundservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractairportgroundservicetypeIdent, __gong__toRawStringLiteral(abstractairportgroundservicetype.Description)))
			if abstractairportgroundservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractairportgroundservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractairportgroundservicetypeIdent, targetIdent))
			}
			for _, elem := range abstractairportgroundservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractairportgroundservicetypeIdent, abstractairportgroundservicetypeIdent, targetIdent))
			}
			if abstractairportgroundservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractairportgroundservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractairportgroundservicetypeIdent, targetIdent))
			}
			if abstractairportgroundservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractairportgroundservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", abstractairportgroundservicetypeIdent, targetIdent))
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
		for _, abstractservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AbstractServiceTypes, stage.AixmfeaturesStage.AbstractServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			abstractservicetypeIdent := abstractservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", abstractservicetypeIdent, __gong__toRawStringLiteral(abstractservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractservicetypeIdent, __gong__toRawStringLiteral(abstractservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractservicetypeIdent, __gong__toRawStringLiteral(abstractservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractservicetypeIdent, __gong__toRawStringLiteral(abstractservicetype.Description)))
			if abstractservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractservicetypeIdent, targetIdent))
			}
			for _, elem := range abstractservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractservicetypeIdent, abstractservicetypeIdent, targetIdent))
			}
			if abstractservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractservicetypeIdent, targetIdent))
			}
			if abstractservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", abstractservicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, abstracttrafficseparationservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AbstractTrafficSeparationServiceTypes, stage.AixmfeaturesStage.AbstractTrafficSeparationServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			abstracttrafficseparationservicetypeIdent := abstracttrafficseparationservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractTrafficSeparationServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", abstracttrafficseparationservicetypeIdent, __gong__toRawStringLiteral(abstracttrafficseparationservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstracttrafficseparationservicetypeIdent, __gong__toRawStringLiteral(abstracttrafficseparationservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstracttrafficseparationservicetypeIdent, __gong__toRawStringLiteral(abstracttrafficseparationservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstracttrafficseparationservicetypeIdent, __gong__toRawStringLiteral(abstracttrafficseparationservicetype.Description)))
			if abstracttrafficseparationservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstracttrafficseparationservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstracttrafficseparationservicetypeIdent, targetIdent))
			}
			for _, elem := range abstracttrafficseparationservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstracttrafficseparationservicetypeIdent, abstracttrafficseparationservicetypeIdent, targetIdent))
			}
			if abstracttrafficseparationservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstracttrafficseparationservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstracttrafficseparationservicetypeIdent, targetIdent))
			}
			if abstracttrafficseparationservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstracttrafficseparationservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", abstracttrafficseparationservicetypeIdent, targetIdent))
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
		for _, airtrafficcontrolservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirTrafficControlServiceTimeSlicePropertyTypes, stage.AixmfeaturesStage.AirTrafficControlServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airtrafficcontrolservicetimeslicepropertytypeIdent := airtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficControlServiceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airtrafficcontrolservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficcontrolservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airtrafficcontrolservicetimeslicepropertytypeIdent, airtrafficcontrolservicetimeslicepropertytype.Owns))
			if airtrafficcontrolservicetimeslicepropertytype.AirTrafficControlServiceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicepropertytype.AirTrafficControlServiceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirTrafficControlServiceTimeSlice = %s", airtrafficcontrolservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airtrafficcontrolservicetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes, stage.AixmfeaturesStage.AirTrafficControlServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airtrafficcontrolservicetimeslicetypeIdent := airtrafficcontrolservicetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficControlServiceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetype.CorrectionNumber))
			if airtrafficcontrolservicetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.FlightOperations != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.FlightOperations.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightOperations = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.Call_sign {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.RadarAssisted != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.RadarAssisted.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarAssisted = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.DataLinkEnabled != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.DataLinkEnabled.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataLinkEnabled = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.DataLinkChannel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.DataLinkChannel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataLinkChannel = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientAirport {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirport = append(%s.ClientAirport, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientAirspace {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirspace = append(%s.ClientAirspace, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientRoute {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientRoute = append(%s.ClientRoute, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientProcedure {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientProcedure = append(%s.ClientProcedure, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientHolding {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientHolding = append(%s.ClientHolding, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.ClientAerialRefuelling {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAerialRefuelling = append(%s.ClientAerialRefuelling, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetimeslicetype.AircraftLocator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetimeslicetype.AircraftLocator.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftLocator = %s", airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airtrafficcontrolservicetimeslicetypeIdent, airtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airtrafficcontrolservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirTrafficControlServiceTypes, stage.AixmfeaturesStage.AirTrafficControlServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airtrafficcontrolservicetypeIdent := airtrafficcontrolservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficControlServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(airtrafficcontrolservicetype.Description)))
			if airtrafficcontrolservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airtrafficcontrolservicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airtrafficcontrolservicetypeIdent, airtrafficcontrolservicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", airtrafficcontrolservicetypeIdent, targetIdent))
			}
			if airtrafficcontrolservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficcontrolservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", airtrafficcontrolservicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficcontrolservicetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airtrafficcontrolservicetypeIdent, airtrafficcontrolservicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airtrafficmanagementservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirTrafficManagementServiceTimeSlicePropertyTypes, stage.AixmfeaturesStage.AirTrafficManagementServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airtrafficmanagementservicetimeslicepropertytypeIdent := airtrafficmanagementservicetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficManagementServiceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airtrafficmanagementservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficmanagementservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airtrafficmanagementservicetimeslicepropertytypeIdent, airtrafficmanagementservicetimeslicepropertytype.Owns))
			if airtrafficmanagementservicetimeslicepropertytype.AirTrafficManagementServiceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicepropertytype.AirTrafficManagementServiceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirTrafficManagementServiceTimeSlice = %s", airtrafficmanagementservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airtrafficmanagementservicetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes, stage.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airtrafficmanagementservicetimeslicetypeIdent := airtrafficmanagementservicetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficManagementServiceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airtrafficmanagementservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficmanagementservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airtrafficmanagementservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airtrafficmanagementservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airtrafficmanagementservicetimeslicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetype.CorrectionNumber))
			if airtrafficmanagementservicetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetimeslicetype.FlightOperations != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.FlightOperations.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightOperations = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.Call_sign {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.ClientAirspace {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirspace = append(%s.ClientAirspace, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.ClientAerialRefuelling {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAerialRefuelling = append(%s.ClientAerialRefuelling, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.ClientRoute {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientRoute = append(%s.ClientRoute, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airtrafficmanagementservicetimeslicetypeIdent, airtrafficmanagementservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airtrafficmanagementservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirTrafficManagementServiceTypes, stage.AixmfeaturesStage.AirTrafficManagementServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airtrafficmanagementservicetypeIdent := airtrafficmanagementservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirTrafficManagementServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airtrafficmanagementservicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airtrafficmanagementservicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airtrafficmanagementservicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airtrafficmanagementservicetypeIdent, __gong__toRawStringLiteral(airtrafficmanagementservicetype.Description)))
			if airtrafficmanagementservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airtrafficmanagementservicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airtrafficmanagementservicetypeIdent, airtrafficmanagementservicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", airtrafficmanagementservicetypeIdent, targetIdent))
			}
			if airtrafficmanagementservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airtrafficmanagementservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", airtrafficmanagementservicetypeIdent, targetIdent))
			}
			for _, elem := range airtrafficmanagementservicetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airtrafficmanagementservicetypeIdent, airtrafficmanagementservicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aircraftgroundservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AircraftGroundServiceTimeSlicePropertyTypes, stage.AixmfeaturesStage.AircraftGroundServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aircraftgroundservicetimeslicepropertytypeIdent := aircraftgroundservicetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftGroundServiceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", aircraftgroundservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftgroundservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", aircraftgroundservicetimeslicepropertytypeIdent, aircraftgroundservicetimeslicepropertytype.Owns))
			if aircraftgroundservicetimeslicepropertytype.AircraftGroundServiceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicepropertytype.AircraftGroundServiceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftGroundServiceTimeSlice = %s", aircraftgroundservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aircraftgroundservicetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes, stage.AixmfeaturesStage.AircraftGroundServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aircraftgroundservicetimeslicetypeIdent := aircraftgroundservicetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftGroundServiceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", aircraftgroundservicetimeslicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftgroundservicetimeslicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aircraftgroundservicetimeslicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aircraftgroundservicetimeslicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", aircraftgroundservicetimeslicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetype.CorrectionNumber))
			if aircraftgroundservicetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetimeslicetype.FlightOperations != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.FlightOperations.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightOperations = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.Call_sign {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.AirportHeliport {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aircraftgroundservicetimeslicetypeIdent, aircraftgroundservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, aircraftgroundservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AircraftGroundServiceTypes, stage.AixmfeaturesStage.AircraftGroundServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			aircraftgroundservicetypeIdent := aircraftgroundservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftGroundServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", aircraftgroundservicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftgroundservicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aircraftgroundservicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aircraftgroundservicetypeIdent, __gong__toRawStringLiteral(aircraftgroundservicetype.Description)))
			if aircraftgroundservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", aircraftgroundservicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", aircraftgroundservicetypeIdent, aircraftgroundservicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", aircraftgroundservicetypeIdent, targetIdent))
			}
			if aircraftgroundservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := aircraftgroundservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", aircraftgroundservicetypeIdent, targetIdent))
			}
			for _, elem := range aircraftgroundservicetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", aircraftgroundservicetypeIdent, aircraftgroundservicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airportclearanceservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirportClearanceServiceTimeSlicePropertyTypes, stage.AixmfeaturesStage.AirportClearanceServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airportclearanceservicetimeslicepropertytypeIdent := airportclearanceservicetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportClearanceServiceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airportclearanceservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportclearanceservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airportclearanceservicetimeslicepropertytypeIdent, airportclearanceservicetimeslicepropertytype.Owns))
			if airportclearanceservicetimeslicepropertytype.AirportClearanceServiceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicepropertytype.AirportClearanceServiceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportClearanceServiceTimeSlice = %s", airportclearanceservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airportclearanceservicetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes, stage.AixmfeaturesStage.AirportClearanceServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airportclearanceservicetimeslicetypeIdent := airportclearanceservicetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportClearanceServiceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airportclearanceservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportclearanceservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airportclearanceservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airportclearanceservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airportclearanceservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetype.CorrectionNumber))
			if airportclearanceservicetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			if airportclearanceservicetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			if airportclearanceservicetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			if airportclearanceservicetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			if airportclearanceservicetimeslicetype.FlightOperations != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicetype.FlightOperations.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightOperations = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			if airportclearanceservicetimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			if airportclearanceservicetimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			if airportclearanceservicetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			if airportclearanceservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.Call_sign {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.AirportHeliport {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			if airportclearanceservicetimeslicetype.SnowPlan != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetimeslicetype.SnowPlan.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SnowPlan = %s", airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airportclearanceservicetimeslicetypeIdent, airportclearanceservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airportclearanceservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirportClearanceServiceTypes, stage.AixmfeaturesStage.AirportClearanceServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airportclearanceservicetypeIdent := airportclearanceservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportClearanceServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airportclearanceservicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportclearanceservicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airportclearanceservicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airportclearanceservicetypeIdent, __gong__toRawStringLiteral(airportclearanceservicetype.Description)))
			if airportclearanceservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airportclearanceservicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airportclearanceservicetypeIdent, airportclearanceservicetypeIdent, targetIdent))
			}
			if airportclearanceservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", airportclearanceservicetypeIdent, targetIdent))
			}
			if airportclearanceservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportclearanceservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", airportclearanceservicetypeIdent, targetIdent))
			}
			for _, elem := range airportclearanceservicetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airportclearanceservicetypeIdent, airportclearanceservicetypeIdent, targetIdent))
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
		for _, airportsuppliesservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirportSuppliesServiceTimeSlicePropertyTypes, stage.AixmfeaturesStage.AirportSuppliesServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airportsuppliesservicetimeslicepropertytypeIdent := airportsuppliesservicetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportSuppliesServiceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", airportsuppliesservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportsuppliesservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airportsuppliesservicetimeslicepropertytypeIdent, airportsuppliesservicetimeslicepropertytype.Owns))
			if airportsuppliesservicetimeslicepropertytype.AirportSuppliesServiceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetimeslicepropertytype.AirportSuppliesServiceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportSuppliesServiceTimeSlice = %s", airportsuppliesservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airportsuppliesservicetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes, stage.AixmfeaturesStage.AirportSuppliesServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airportsuppliesservicetimeslicetypeIdent := airportsuppliesservicetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportSuppliesServiceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airportsuppliesservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportsuppliesservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airportsuppliesservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airportsuppliesservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airportsuppliesservicetimeslicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetype.CorrectionNumber))
			if airportsuppliesservicetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetimeslicetype.FlightOperations != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.FlightOperations.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightOperations = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.Call_sign {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.AirportHeliport {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.FuelSupply {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FuelSupply = append(%s.FuelSupply, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.OilSupply {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OilSupply = append(%s.OilSupply, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.NitrogenSupply {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NitrogenSupply = append(%s.NitrogenSupply, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.OxygenSupply {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OxygenSupply = append(%s.OxygenSupply, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airportsuppliesservicetimeslicetypeIdent, airportsuppliesservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, airportsuppliesservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AirportSuppliesServiceTypes, stage.AixmfeaturesStage.AirportSuppliesServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			airportsuppliesservicetypeIdent := airportsuppliesservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportSuppliesServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", airportsuppliesservicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportsuppliesservicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airportsuppliesservicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airportsuppliesservicetypeIdent, __gong__toRawStringLiteral(airportsuppliesservicetype.Description)))
			if airportsuppliesservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", airportsuppliesservicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", airportsuppliesservicetypeIdent, airportsuppliesservicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", airportsuppliesservicetypeIdent, targetIdent))
			}
			if airportsuppliesservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := airportsuppliesservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", airportsuppliesservicetypeIdent, targetIdent))
			}
			for _, elem := range airportsuppliesservicetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airportsuppliesservicetypeIdent, airportsuppliesservicetypeIdent, targetIdent))
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
		for _, callsigndetailpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.CallsignDetailPropertyTypes, stage.AixmfeaturesStage.CallsignDetailPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			callsigndetailpropertytypeIdent := callsigndetailpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CallsignDetailPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", callsigndetailpropertytypeIdent, __gong__toRawStringLiteral(callsigndetailpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", callsigndetailpropertytypeIdent, __gong__toRawStringLiteral(callsigndetailpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", callsigndetailpropertytypeIdent, __gong__toRawStringLiteral(callsigndetailpropertytype.NilReason)))
			if callsigndetailpropertytype.CallsignDetail != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := callsigndetailpropertytype.CallsignDetail.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CallsignDetail = %s", callsigndetailpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, callsigndetailtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.CallsignDetailTypes, stage.AixmfeaturesStage.CallsignDetailType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			callsigndetailtypeIdent := callsigndetailtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CallsignDetailType{Name: %s}).Stage(stage.AixmfeaturesStage)", callsigndetailtypeIdent, __gong__toRawStringLiteral(callsigndetailtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", callsigndetailtypeIdent, __gong__toRawStringLiteral(callsigndetailtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", callsigndetailtypeIdent, __gong__toRawStringLiteral(callsigndetailtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", callsigndetailtypeIdent, __gong__toRawStringLiteral(callsigndetailtype.Description)))
			if callsigndetailtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := callsigndetailtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", callsigndetailtypeIdent, targetIdent))
			}
			for _, elem := range callsigndetailtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", callsigndetailtypeIdent, callsigndetailtypeIdent, targetIdent))
			}
			if callsigndetailtype.CallSign != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := callsigndetailtype.CallSign.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CallSign = %s", callsigndetailtypeIdent, targetIdent))
			}
			if callsigndetailtype.Language != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := callsigndetailtype.Language.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Language = %s", callsigndetailtypeIdent, targetIdent))
			}
			for _, elem := range callsigndetailtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", callsigndetailtypeIdent, callsigndetailtypeIdent, targetIdent))
			}
			for _, elem := range callsigndetailtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", callsigndetailtypeIdent, callsigndetailtypeIdent, targetIdent))
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
		for _, directionfinderpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DirectionFinderPropertyTypes, stage.AixmfeaturesStage.DirectionFinderPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			directionfinderpropertytypeIdent := directionfinderpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectionFinderPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", directionfinderpropertytypeIdent, directionfinderpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", directionfinderpropertytypeIdent, __gong__toRawStringLiteral(directionfinderpropertytype.Actuate)))
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
		for _, firefightingservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FireFightingServiceTimeSlicePropertyTypes, stage.AixmfeaturesStage.FireFightingServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			firefightingservicetimeslicepropertytypeIdent := firefightingservicetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FireFightingServiceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", firefightingservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", firefightingservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", firefightingservicetimeslicepropertytypeIdent, firefightingservicetimeslicepropertytype.Owns))
			if firefightingservicetimeslicepropertytype.FireFightingServiceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicepropertytype.FireFightingServiceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FireFightingServiceTimeSlice = %s", firefightingservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, firefightingservicetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FireFightingServiceTimeSliceTypes, stage.AixmfeaturesStage.FireFightingServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			firefightingservicetimeslicetypeIdent := firefightingservicetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FireFightingServiceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", firefightingservicetimeslicetypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", firefightingservicetimeslicetypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", firefightingservicetimeslicetypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", firefightingservicetimeslicetypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", firefightingservicetimeslicetypeIdent, __gong__toRawStringLiteral(firefightingservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetype.CorrectionNumber))
			if firefightingservicetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.FlightOperations != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.FlightOperations.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightOperations = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.Call_sign {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.AirportHeliport {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.Category != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.Category.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Category = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			if firefightingservicetimeslicetype.Standard != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetimeslicetype.Standard.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Standard = %s", firefightingservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", firefightingservicetimeslicetypeIdent, firefightingservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, firefightingservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FireFightingServiceTypes, stage.AixmfeaturesStage.FireFightingServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			firefightingservicetypeIdent := firefightingservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FireFightingServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", firefightingservicetypeIdent, __gong__toRawStringLiteral(firefightingservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", firefightingservicetypeIdent, __gong__toRawStringLiteral(firefightingservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", firefightingservicetypeIdent, __gong__toRawStringLiteral(firefightingservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", firefightingservicetypeIdent, __gong__toRawStringLiteral(firefightingservicetype.Description)))
			if firefightingservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", firefightingservicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", firefightingservicetypeIdent, firefightingservicetypeIdent, targetIdent))
			}
			if firefightingservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", firefightingservicetypeIdent, targetIdent))
			}
			if firefightingservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := firefightingservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", firefightingservicetypeIdent, targetIdent))
			}
			for _, elem := range firefightingservicetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", firefightingservicetypeIdent, firefightingservicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, fuelpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FuelPropertyTypes, stage.AixmfeaturesStage.FuelPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			fuelpropertytypeIdent := fuelpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FuelPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", fuelpropertytypeIdent, __gong__toRawStringLiteral(fuelpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fuelpropertytypeIdent, __gong__toRawStringLiteral(fuelpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", fuelpropertytypeIdent, __gong__toRawStringLiteral(fuelpropertytype.NilReason)))
			if fuelpropertytype.Fuel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := fuelpropertytype.Fuel.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fuel = %s", fuelpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, fueltype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.FuelTypes, stage.AixmfeaturesStage.FuelType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			fueltypeIdent := fueltype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.FuelType{Name: %s}).Stage(stage.AixmfeaturesStage)", fueltypeIdent, __gong__toRawStringLiteral(fueltype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fueltypeIdent, __gong__toRawStringLiteral(fueltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", fueltypeIdent, __gong__toRawStringLiteral(fueltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", fueltypeIdent, __gong__toRawStringLiteral(fueltype.Description)))
			if fueltype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := fueltype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", fueltypeIdent, targetIdent))
			}
			for _, elem := range fueltype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", fueltypeIdent, fueltypeIdent, targetIdent))
			}
			if fueltype.Category != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := fueltype.Category.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Category = %s", fueltypeIdent, targetIdent))
			}
			for _, elem := range fueltype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", fueltypeIdent, fueltypeIdent, targetIdent))
			}
			for _, elem := range fueltype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", fueltypeIdent, fueltypeIdent, targetIdent))
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
		for _, groundtrafficcontrolservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.GroundTrafficControlServiceTimeSlicePropertyTypes, stage.AixmfeaturesStage.GroundTrafficControlServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			groundtrafficcontrolservicetimeslicepropertytypeIdent := groundtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GroundTrafficControlServiceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", groundtrafficcontrolservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groundtrafficcontrolservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", groundtrafficcontrolservicetimeslicepropertytypeIdent, groundtrafficcontrolservicetimeslicepropertytype.Owns))
			if groundtrafficcontrolservicetimeslicepropertytype.GroundTrafficControlServiceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicepropertytype.GroundTrafficControlServiceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundTrafficControlServiceTimeSlice = %s", groundtrafficcontrolservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, groundtrafficcontrolservicetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes, stage.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			groundtrafficcontrolservicetimeslicetypeIdent := groundtrafficcontrolservicetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GroundTrafficControlServiceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", groundtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groundtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", groundtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", groundtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", groundtrafficcontrolservicetimeslicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetype.CorrectionNumber))
			if groundtrafficcontrolservicetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.FlightOperations != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.FlightOperations.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightOperations = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.Call_sign {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.RadarAssisted != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.RadarAssisted.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarAssisted = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.DataLinkEnabled != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.DataLinkEnabled.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataLinkEnabled = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.DataLinkChannel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.DataLinkChannel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataLinkChannel = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetimeslicetype.ClientAirport != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetimeslicetype.ClientAirport.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirport = %s", groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", groundtrafficcontrolservicetimeslicetypeIdent, groundtrafficcontrolservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, groundtrafficcontrolservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.GroundTrafficControlServiceTypes, stage.AixmfeaturesStage.GroundTrafficControlServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			groundtrafficcontrolservicetypeIdent := groundtrafficcontrolservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GroundTrafficControlServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", groundtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groundtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", groundtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", groundtrafficcontrolservicetypeIdent, __gong__toRawStringLiteral(groundtrafficcontrolservicetype.Description)))
			if groundtrafficcontrolservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", groundtrafficcontrolservicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", groundtrafficcontrolservicetypeIdent, groundtrafficcontrolservicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", groundtrafficcontrolservicetypeIdent, targetIdent))
			}
			if groundtrafficcontrolservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := groundtrafficcontrolservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", groundtrafficcontrolservicetypeIdent, targetIdent))
			}
			for _, elem := range groundtrafficcontrolservicetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", groundtrafficcontrolservicetypeIdent, groundtrafficcontrolservicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, holdingpatternpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.HoldingPatternPropertyTypes, stage.AixmfeaturesStage.HoldingPatternPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			holdingpatternpropertytypeIdent := holdingpatternpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.HoldingPatternPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", holdingpatternpropertytypeIdent, holdingpatternpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", holdingpatternpropertytypeIdent, __gong__toRawStringLiteral(holdingpatternpropertytype.Actuate)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, informationservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.InformationServiceTimeSlicePropertyTypes, stage.AixmfeaturesStage.InformationServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			informationservicetimeslicepropertytypeIdent := informationservicetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InformationServiceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", informationservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(informationservicetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", informationservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(informationservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", informationservicetimeslicepropertytypeIdent, informationservicetimeslicepropertytype.Owns))
			if informationservicetimeslicepropertytype.InformationServiceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicepropertytype.InformationServiceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InformationServiceTimeSlice = %s", informationservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, informationservicetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.InformationServiceTimeSliceTypes, stage.AixmfeaturesStage.InformationServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			informationservicetimeslicetypeIdent := informationservicetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InformationServiceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", informationservicetimeslicetypeIdent, __gong__toRawStringLiteral(informationservicetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", informationservicetimeslicetypeIdent, __gong__toRawStringLiteral(informationservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", informationservicetimeslicetypeIdent, __gong__toRawStringLiteral(informationservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", informationservicetimeslicetypeIdent, __gong__toRawStringLiteral(informationservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", informationservicetimeslicetypeIdent, __gong__toRawStringLiteral(informationservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", informationservicetimeslicetypeIdent, informationservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", informationservicetimeslicetypeIdent, informationservicetimeslicetype.CorrectionNumber))
			if informationservicetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.FlightOperations != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.FlightOperations.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightOperations = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.Call_sign {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.Voice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.Voice.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Voice = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.DataLink != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.DataLink.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataLink = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			if informationservicetimeslicetype.Recorded != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetimeslicetype.Recorded.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Recorded = %s", informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.NavaidBroadcast {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidBroadcast = append(%s.NavaidBroadcast, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientAirspace {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirspace = append(%s.ClientAirspace, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientAirport {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirport = append(%s.ClientAirport, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientRoute {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientRoute = append(%s.ClientRoute, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientProcedure {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientProcedure = append(%s.ClientProcedure, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientHolding {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientHolding = append(%s.ClientHolding, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.ClientAerialRefuelling {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAerialRefuelling = append(%s.ClientAerialRefuelling, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", informationservicetimeslicetypeIdent, informationservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, informationservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.InformationServiceTypes, stage.AixmfeaturesStage.InformationServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			informationservicetypeIdent := informationservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InformationServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", informationservicetypeIdent, __gong__toRawStringLiteral(informationservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", informationservicetypeIdent, __gong__toRawStringLiteral(informationservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", informationservicetypeIdent, __gong__toRawStringLiteral(informationservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", informationservicetypeIdent, __gong__toRawStringLiteral(informationservicetype.Description)))
			if informationservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", informationservicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", informationservicetypeIdent, informationservicetypeIdent, targetIdent))
			}
			if informationservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", informationservicetypeIdent, targetIdent))
			}
			if informationservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := informationservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", informationservicetypeIdent, targetIdent))
			}
			for _, elem := range informationservicetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", informationservicetypeIdent, informationservicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, lightactivationpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LightActivationPropertyTypes, stage.AixmfeaturesStage.LightActivationPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			lightactivationpropertytypeIdent := lightactivationpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightActivationPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", lightactivationpropertytypeIdent, __gong__toRawStringLiteral(lightactivationpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightactivationpropertytypeIdent, __gong__toRawStringLiteral(lightactivationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", lightactivationpropertytypeIdent, __gong__toRawStringLiteral(lightactivationpropertytype.NilReason)))
			if lightactivationpropertytype.LightActivation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightactivationpropertytype.LightActivation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LightActivation = %s", lightactivationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, lightactivationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LightActivationTypes, stage.AixmfeaturesStage.LightActivationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			lightactivationtypeIdent := lightactivationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightActivationType{Name: %s}).Stage(stage.AixmfeaturesStage)", lightactivationtypeIdent, __gong__toRawStringLiteral(lightactivationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightactivationtypeIdent, __gong__toRawStringLiteral(lightactivationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", lightactivationtypeIdent, __gong__toRawStringLiteral(lightactivationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", lightactivationtypeIdent, __gong__toRawStringLiteral(lightactivationtype.Description)))
			if lightactivationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightactivationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", lightactivationtypeIdent, targetIdent))
			}
			for _, elem := range lightactivationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", lightactivationtypeIdent, lightactivationtypeIdent, targetIdent))
			}
			if lightactivationtype.Clicks != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightactivationtype.Clicks.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Clicks = %s", lightactivationtypeIdent, targetIdent))
			}
			if lightactivationtype.IntensityLevel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightactivationtype.IntensityLevel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntensityLevel = %s", lightactivationtypeIdent, targetIdent))
			}
			if lightactivationtype.Activation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := lightactivationtype.Activation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Activation = %s", lightactivationtypeIdent, targetIdent))
			}
			for _, elem := range lightactivationtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", lightactivationtypeIdent, lightactivationtypeIdent, targetIdent))
			}
			for _, elem := range lightactivationtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", lightactivationtypeIdent, lightactivationtypeIdent, targetIdent))
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
		for _, nitrogenpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NitrogenPropertyTypes, stage.AixmfeaturesStage.NitrogenPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			nitrogenpropertytypeIdent := nitrogenpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NitrogenPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", nitrogenpropertytypeIdent, __gong__toRawStringLiteral(nitrogenpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nitrogenpropertytypeIdent, __gong__toRawStringLiteral(nitrogenpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", nitrogenpropertytypeIdent, __gong__toRawStringLiteral(nitrogenpropertytype.NilReason)))
			if nitrogenpropertytype.Nitrogen != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := nitrogenpropertytype.Nitrogen.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Nitrogen = %s", nitrogenpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, nitrogentype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NitrogenTypes, stage.AixmfeaturesStage.NitrogenType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			nitrogentypeIdent := nitrogentype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NitrogenType{Name: %s}).Stage(stage.AixmfeaturesStage)", nitrogentypeIdent, __gong__toRawStringLiteral(nitrogentype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nitrogentypeIdent, __gong__toRawStringLiteral(nitrogentype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", nitrogentypeIdent, __gong__toRawStringLiteral(nitrogentype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", nitrogentypeIdent, __gong__toRawStringLiteral(nitrogentype.Description)))
			if nitrogentype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := nitrogentype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", nitrogentypeIdent, targetIdent))
			}
			for _, elem := range nitrogentype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", nitrogentypeIdent, nitrogentypeIdent, targetIdent))
			}
			if nitrogentype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := nitrogentype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", nitrogentypeIdent, targetIdent))
			}
			for _, elem := range nitrogentype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", nitrogentypeIdent, nitrogentypeIdent, targetIdent))
			}
			for _, elem := range nitrogentype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", nitrogentypeIdent, nitrogentypeIdent, targetIdent))
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
		for _, oilpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OilPropertyTypes, stage.AixmfeaturesStage.OilPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			oilpropertytypeIdent := oilpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OilPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", oilpropertytypeIdent, __gong__toRawStringLiteral(oilpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", oilpropertytypeIdent, __gong__toRawStringLiteral(oilpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", oilpropertytypeIdent, __gong__toRawStringLiteral(oilpropertytype.NilReason)))
			if oilpropertytype.Oil != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := oilpropertytype.Oil.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Oil = %s", oilpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, oiltype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OilTypes, stage.AixmfeaturesStage.OilType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			oiltypeIdent := oiltype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OilType{Name: %s}).Stage(stage.AixmfeaturesStage)", oiltypeIdent, __gong__toRawStringLiteral(oiltype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", oiltypeIdent, __gong__toRawStringLiteral(oiltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", oiltypeIdent, __gong__toRawStringLiteral(oiltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", oiltypeIdent, __gong__toRawStringLiteral(oiltype.Description)))
			if oiltype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := oiltype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", oiltypeIdent, targetIdent))
			}
			for _, elem := range oiltype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", oiltypeIdent, oiltypeIdent, targetIdent))
			}
			if oiltype.Category != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := oiltype.Category.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Category = %s", oiltypeIdent, targetIdent))
			}
			for _, elem := range oiltype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", oiltypeIdent, oiltypeIdent, targetIdent))
			}
			for _, elem := range oiltype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", oiltypeIdent, oiltypeIdent, targetIdent))
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
		for _, organisationauthorityassociationpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OrganisationAuthorityAssociationPropertyTypes, stage.AixmfeaturesStage.OrganisationAuthorityAssociationPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			organisationauthorityassociationpropertytypeIdent := organisationauthorityassociationpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityAssociationPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", organisationauthorityassociationpropertytypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthorityassociationpropertytypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", organisationauthorityassociationpropertytypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationpropertytype.NilReason)))
			if organisationauthorityassociationpropertytype.OrganisationAuthorityAssociation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthorityassociationpropertytype.OrganisationAuthorityAssociation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OrganisationAuthorityAssociation = %s", organisationauthorityassociationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, organisationauthorityassociationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OrganisationAuthorityAssociationTypes, stage.AixmfeaturesStage.OrganisationAuthorityAssociationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			organisationauthorityassociationtypeIdent := organisationauthorityassociationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityAssociationType{Name: %s}).Stage(stage.AixmfeaturesStage)", organisationauthorityassociationtypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthorityassociationtypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", organisationauthorityassociationtypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", organisationauthorityassociationtypeIdent, __gong__toRawStringLiteral(organisationauthorityassociationtype.Description)))
			if organisationauthorityassociationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthorityassociationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", organisationauthorityassociationtypeIdent, targetIdent))
			}
			for _, elem := range organisationauthorityassociationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", organisationauthorityassociationtypeIdent, organisationauthorityassociationtypeIdent, targetIdent))
			}
			if organisationauthorityassociationtype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthorityassociationtype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", organisationauthorityassociationtypeIdent, targetIdent))
			}
			for _, elem := range organisationauthorityassociationtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", organisationauthorityassociationtypeIdent, organisationauthorityassociationtypeIdent, targetIdent))
			}
			if organisationauthorityassociationtype.TheOrganisationAuthority != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthorityassociationtype.TheOrganisationAuthority.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheOrganisationAuthority = %s", organisationauthorityassociationtypeIdent, targetIdent))
			}
			for _, elem := range organisationauthorityassociationtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", organisationauthorityassociationtypeIdent, organisationauthorityassociationtypeIdent, targetIdent))
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
		for _, organisationauthoritytimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OrganisationAuthorityTimeSlicePropertyTypes, stage.AixmfeaturesStage.OrganisationAuthorityTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			organisationauthoritytimeslicepropertytypeIdent := organisationauthoritytimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", organisationauthoritytimeslicepropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthoritytimeslicepropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", organisationauthoritytimeslicepropertytypeIdent, organisationauthoritytimeslicepropertytype.Owns))
			if organisationauthoritytimeslicepropertytype.OrganisationAuthorityTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytimeslicepropertytype.OrganisationAuthorityTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OrganisationAuthorityTimeSlice = %s", organisationauthoritytimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, organisationauthoritytimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes, stage.AixmfeaturesStage.OrganisationAuthorityTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			organisationauthoritytimeslicetypeIdent := organisationauthoritytimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", organisationauthoritytimeslicetypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthoritytimeslicetypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", organisationauthoritytimeslicetypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", organisationauthoritytimeslicetypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", organisationauthoritytimeslicetypeIdent, __gong__toRawStringLiteral(organisationauthoritytimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetype.CorrectionNumber))
			if organisationauthoritytimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			for _, elem := range organisationauthoritytimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			if organisationauthoritytimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			if organisationauthoritytimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			if organisationauthoritytimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			if organisationauthoritytimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			if organisationauthoritytimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			if organisationauthoritytimeslicetype.Military != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytimeslicetype.Military.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Military = %s", organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			for _, elem := range organisationauthoritytimeslicetype.Contact {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			for _, elem := range organisationauthoritytimeslicetype.RelatedOrganisationAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RelatedOrganisationAuthority = append(%s.RelatedOrganisationAuthority, %s)", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			for _, elem := range organisationauthoritytimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetypeIdent, targetIdent))
			}
			for _, elem := range organisationauthoritytimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", organisationauthoritytimeslicetypeIdent, organisationauthoritytimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, organisationauthoritytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OrganisationAuthorityTypes, stage.AixmfeaturesStage.OrganisationAuthorityType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			organisationauthoritytypeIdent := organisationauthoritytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityType{Name: %s}).Stage(stage.AixmfeaturesStage)", organisationauthoritytypeIdent, __gong__toRawStringLiteral(organisationauthoritytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthoritytypeIdent, __gong__toRawStringLiteral(organisationauthoritytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", organisationauthoritytypeIdent, __gong__toRawStringLiteral(organisationauthoritytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", organisationauthoritytypeIdent, __gong__toRawStringLiteral(organisationauthoritytype.Description)))
			if organisationauthoritytype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", organisationauthoritytypeIdent, targetIdent))
			}
			for _, elem := range organisationauthoritytype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", organisationauthoritytypeIdent, organisationauthoritytypeIdent, targetIdent))
			}
			if organisationauthoritytype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", organisationauthoritytypeIdent, targetIdent))
			}
			if organisationauthoritytype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := organisationauthoritytype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", organisationauthoritytypeIdent, targetIdent))
			}
			for _, elem := range organisationauthoritytype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", organisationauthoritytypeIdent, organisationauthoritytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, oxygenpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OxygenPropertyTypes, stage.AixmfeaturesStage.OxygenPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			oxygenpropertytypeIdent := oxygenpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OxygenPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", oxygenpropertytypeIdent, __gong__toRawStringLiteral(oxygenpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", oxygenpropertytypeIdent, __gong__toRawStringLiteral(oxygenpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", oxygenpropertytypeIdent, __gong__toRawStringLiteral(oxygenpropertytype.NilReason)))
			if oxygenpropertytype.Oxygen != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := oxygenpropertytype.Oxygen.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Oxygen = %s", oxygenpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, oxygentype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.OxygenTypes, stage.AixmfeaturesStage.OxygenType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			oxygentypeIdent := oxygentype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OxygenType{Name: %s}).Stage(stage.AixmfeaturesStage)", oxygentypeIdent, __gong__toRawStringLiteral(oxygentype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", oxygentypeIdent, __gong__toRawStringLiteral(oxygentype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", oxygentypeIdent, __gong__toRawStringLiteral(oxygentype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", oxygentypeIdent, __gong__toRawStringLiteral(oxygentype.Description)))
			if oxygentype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := oxygentype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", oxygentypeIdent, targetIdent))
			}
			for _, elem := range oxygentype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", oxygentypeIdent, oxygentypeIdent, targetIdent))
			}
			if oxygentype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := oxygentype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", oxygentypeIdent, targetIdent))
			}
			for _, elem := range oxygentype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", oxygentypeIdent, oxygentypeIdent, targetIdent))
			}
			for _, elem := range oxygentype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", oxygentypeIdent, oxygentypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, passengerservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PassengerServiceTimeSlicePropertyTypes, stage.AixmfeaturesStage.PassengerServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			passengerservicetimeslicepropertytypeIdent := passengerservicetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PassengerServiceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", passengerservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", passengerservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", passengerservicetimeslicepropertytypeIdent, passengerservicetimeslicepropertytype.Owns))
			if passengerservicetimeslicepropertytype.PassengerServiceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicepropertytype.PassengerServiceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PassengerServiceTimeSlice = %s", passengerservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, passengerservicetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PassengerServiceTimeSliceTypes, stage.AixmfeaturesStage.PassengerServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			passengerservicetimeslicetypeIdent := passengerservicetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PassengerServiceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", passengerservicetimeslicetypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", passengerservicetimeslicetypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", passengerservicetimeslicetypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", passengerservicetimeslicetypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", passengerservicetimeslicetypeIdent, __gong__toRawStringLiteral(passengerservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", passengerservicetimeslicetypeIdent, passengerservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", passengerservicetimeslicetypeIdent, passengerservicetimeslicetype.CorrectionNumber))
			if passengerservicetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			if passengerservicetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			if passengerservicetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			if passengerservicetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			if passengerservicetimeslicetype.FlightOperations != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicetype.FlightOperations.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightOperations = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			if passengerservicetimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			if passengerservicetimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			if passengerservicetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			if passengerservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.Call_sign {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.AirportHeliport {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = append(%s.AirportHeliport, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
			if passengerservicetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", passengerservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", passengerservicetimeslicetypeIdent, passengerservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, passengerservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PassengerServiceTypes, stage.AixmfeaturesStage.PassengerServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			passengerservicetypeIdent := passengerservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PassengerServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", passengerservicetypeIdent, __gong__toRawStringLiteral(passengerservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", passengerservicetypeIdent, __gong__toRawStringLiteral(passengerservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", passengerservicetypeIdent, __gong__toRawStringLiteral(passengerservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", passengerservicetypeIdent, __gong__toRawStringLiteral(passengerservicetype.Description)))
			if passengerservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", passengerservicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", passengerservicetypeIdent, passengerservicetypeIdent, targetIdent))
			}
			if passengerservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", passengerservicetypeIdent, targetIdent))
			}
			if passengerservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := passengerservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", passengerservicetypeIdent, targetIdent))
			}
			for _, elem := range passengerservicetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", passengerservicetypeIdent, passengerservicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, pilotcontrolledlightingtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PilotControlledLightingTimeSlicePropertyTypes, stage.AixmfeaturesStage.PilotControlledLightingTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			pilotcontrolledlightingtimeslicepropertytypeIdent := pilotcontrolledlightingtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PilotControlledLightingTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", pilotcontrolledlightingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pilotcontrolledlightingtimeslicepropertytypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", pilotcontrolledlightingtimeslicepropertytypeIdent, pilotcontrolledlightingtimeslicepropertytype.Owns))
			if pilotcontrolledlightingtimeslicepropertytype.PilotControlledLightingTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicepropertytype.PilotControlledLightingTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PilotControlledLightingTimeSlice = %s", pilotcontrolledlightingtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, pilotcontrolledlightingtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes, stage.AixmfeaturesStage.PilotControlledLightingTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			pilotcontrolledlightingtimeslicetypeIdent := pilotcontrolledlightingtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PilotControlledLightingTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", pilotcontrolledlightingtimeslicetypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pilotcontrolledlightingtimeslicetypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pilotcontrolledlightingtimeslicetypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pilotcontrolledlightingtimeslicetypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", pilotcontrolledlightingtimeslicetypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetype.CorrectionNumber))
			if pilotcontrolledlightingtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range pilotcontrolledlightingtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtimeslicetype.Duration != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicetype.Duration.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtimeslicetype.IntensitySteps != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicetype.IntensitySteps.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntensitySteps = %s", pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtimeslicetype.StandByIntensity != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicetype.StandByIntensity.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandByIntensity = %s", pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtimeslicetype.RadioFrequency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicetype.RadioFrequency.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioFrequency = %s", pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtimeslicetype.ActivationInstruction != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtimeslicetype.ActivationInstruction.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActivationInstruction = %s", pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range pilotcontrolledlightingtimeslicetype.ControlledLightIntensity {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlledLightIntensity = append(%s.ControlledLightIntensity, %s)", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range pilotcontrolledlightingtimeslicetype.ActivatedGroundLighting {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActivatedGroundLighting = append(%s.ActivatedGroundLighting, %s)", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range pilotcontrolledlightingtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range pilotcontrolledlightingtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", pilotcontrolledlightingtimeslicetypeIdent, pilotcontrolledlightingtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, pilotcontrolledlightingtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PilotControlledLightingTypes, stage.AixmfeaturesStage.PilotControlledLightingType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			pilotcontrolledlightingtypeIdent := pilotcontrolledlightingtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PilotControlledLightingType{Name: %s}).Stage(stage.AixmfeaturesStage)", pilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pilotcontrolledlightingtypeIdent, __gong__toRawStringLiteral(pilotcontrolledlightingtype.Description)))
			if pilotcontrolledlightingtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", pilotcontrolledlightingtypeIdent, targetIdent))
			}
			for _, elem := range pilotcontrolledlightingtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", pilotcontrolledlightingtypeIdent, pilotcontrolledlightingtypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", pilotcontrolledlightingtypeIdent, targetIdent))
			}
			if pilotcontrolledlightingtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := pilotcontrolledlightingtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", pilotcontrolledlightingtypeIdent, targetIdent))
			}
			for _, elem := range pilotcontrolledlightingtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", pilotcontrolledlightingtypeIdent, pilotcontrolledlightingtypeIdent, targetIdent))
			}
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
		for _, procedurepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ProcedurePropertyTypes, stage.AixmfeaturesStage.ProcedurePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			procedurepropertytypeIdent := procedurepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ProcedurePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", procedurepropertytypeIdent, procedurepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", procedurepropertytypeIdent, __gong__toRawStringLiteral(procedurepropertytype.Actuate)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radiocommunicationchannelpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadioCommunicationChannelPropertyTypes, stage.AixmfeaturesStage.RadioCommunicationChannelPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radiocommunicationchannelpropertytypeIdent := radiocommunicationchannelpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationChannelPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", radiocommunicationchannelpropertytypeIdent, radiocommunicationchannelpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", radiocommunicationchannelpropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchannelpropertytype.Actuate)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radiocommunicationchanneltimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadioCommunicationChannelTimeSlicePropertyTypes, stage.AixmfeaturesStage.RadioCommunicationChannelTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radiocommunicationchanneltimeslicepropertytypeIdent := radiocommunicationchanneltimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationChannelTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", radiocommunicationchanneltimeslicepropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationchanneltimeslicepropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", radiocommunicationchanneltimeslicepropertytypeIdent, radiocommunicationchanneltimeslicepropertytype.Owns))
			if radiocommunicationchanneltimeslicepropertytype.RadioCommunicationChannelTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicepropertytype.RadioCommunicationChannelTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunicationChannelTimeSlice = %s", radiocommunicationchanneltimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radiocommunicationchanneltimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes, stage.AixmfeaturesStage.RadioCommunicationChannelTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radiocommunicationchanneltimeslicetypeIdent := radiocommunicationchanneltimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationChannelTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", radiocommunicationchanneltimeslicetypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationchanneltimeslicetypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radiocommunicationchanneltimeslicetypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radiocommunicationchanneltimeslicetypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", radiocommunicationchanneltimeslicetypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetype.CorrectionNumber))
			if radiocommunicationchanneltimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationchanneltimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.Mode != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.Mode.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mode = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.FrequencyTransmission != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.FrequencyTransmission.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FrequencyTransmission = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.FrequencyReception != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.FrequencyReception.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FrequencyReception = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.Channel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.Channel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Channel = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.Logon != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.Logon.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Logon = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.EmissionType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.EmissionType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionType = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.SelectiveCall != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.SelectiveCall.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SelectiveCall = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			if radiocommunicationchanneltimeslicetype.TrafficDirection != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltimeslicetype.TrafficDirection.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TrafficDirection = %s", radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationchanneltimeslicetype.Location {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = append(%s.Location, %s)", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationchanneltimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationchanneltimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationchanneltimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", radiocommunicationchanneltimeslicetypeIdent, radiocommunicationchanneltimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radiocommunicationchanneltype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadioCommunicationChannelTypes, stage.AixmfeaturesStage.RadioCommunicationChannelType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radiocommunicationchanneltypeIdent := radiocommunicationchanneltype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationChannelType{Name: %s}).Stage(stage.AixmfeaturesStage)", radiocommunicationchanneltypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationchanneltypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radiocommunicationchanneltypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radiocommunicationchanneltypeIdent, __gong__toRawStringLiteral(radiocommunicationchanneltype.Description)))
			if radiocommunicationchanneltype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", radiocommunicationchanneltypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationchanneltype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", radiocommunicationchanneltypeIdent, radiocommunicationchanneltypeIdent, targetIdent))
			}
			if radiocommunicationchanneltype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", radiocommunicationchanneltypeIdent, targetIdent))
			}
			if radiocommunicationchanneltype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationchanneltype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", radiocommunicationchanneltypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationchanneltype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", radiocommunicationchanneltypeIdent, radiocommunicationchanneltypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radiocommunicationoperationalstatuspropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadioCommunicationOperationalStatusPropertyTypes, stage.AixmfeaturesStage.RadioCommunicationOperationalStatusPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radiocommunicationoperationalstatuspropertytypeIdent := radiocommunicationoperationalstatuspropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationOperationalStatusPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", radiocommunicationoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatuspropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", radiocommunicationoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatuspropertytype.NilReason)))
			if radiocommunicationoperationalstatuspropertytype.RadioCommunicationOperationalStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationoperationalstatuspropertytype.RadioCommunicationOperationalStatus.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunicationOperationalStatus = %s", radiocommunicationoperationalstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radiocommunicationoperationalstatustype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadioCommunicationOperationalStatusTypes, stage.AixmfeaturesStage.RadioCommunicationOperationalStatusType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radiocommunicationoperationalstatustypeIdent := radiocommunicationoperationalstatustype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadioCommunicationOperationalStatusType{Name: %s}).Stage(stage.AixmfeaturesStage)", radiocommunicationoperationalstatustypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatustype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radiocommunicationoperationalstatustypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radiocommunicationoperationalstatustypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radiocommunicationoperationalstatustypeIdent, __gong__toRawStringLiteral(radiocommunicationoperationalstatustype.Description)))
			if radiocommunicationoperationalstatustype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationoperationalstatustype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", radiocommunicationoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationoperationalstatustype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", radiocommunicationoperationalstatustypeIdent, radiocommunicationoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationoperationalstatustype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", radiocommunicationoperationalstatustypeIdent, radiocommunicationoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationoperationalstatustype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", radiocommunicationoperationalstatustypeIdent, radiocommunicationoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationoperationalstatustype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", radiocommunicationoperationalstatustypeIdent, radiocommunicationoperationalstatustypeIdent, targetIdent))
			}
			if radiocommunicationoperationalstatustype.OperationalStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radiocommunicationoperationalstatustype.OperationalStatus.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OperationalStatus = %s", radiocommunicationoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range radiocommunicationoperationalstatustype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", radiocommunicationoperationalstatustypeIdent, radiocommunicationoperationalstatustypeIdent, targetIdent))
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
		for _, searchrescueservicetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SearchRescueServiceTimeSlicePropertyTypes, stage.AixmfeaturesStage.SearchRescueServiceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			searchrescueservicetimeslicepropertytypeIdent := searchrescueservicetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SearchRescueServiceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", searchrescueservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", searchrescueservicetimeslicepropertytypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", searchrescueservicetimeslicepropertytypeIdent, searchrescueservicetimeslicepropertytype.Owns))
			if searchrescueservicetimeslicepropertytype.SearchRescueServiceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicepropertytype.SearchRescueServiceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SearchRescueServiceTimeSlice = %s", searchrescueservicetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, searchrescueservicetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes, stage.AixmfeaturesStage.SearchRescueServiceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			searchrescueservicetimeslicetypeIdent := searchrescueservicetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SearchRescueServiceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", searchrescueservicetimeslicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", searchrescueservicetimeslicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", searchrescueservicetimeslicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", searchrescueservicetimeslicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", searchrescueservicetimeslicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetype.CorrectionNumber))
			if searchrescueservicetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			if searchrescueservicetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			if searchrescueservicetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			if searchrescueservicetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			if searchrescueservicetimeslicetype.FlightOperations != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicetype.FlightOperations.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightOperations = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			if searchrescueservicetimeslicetype.Rank != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicetype.Rank.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rank = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			if searchrescueservicetimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			if searchrescueservicetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			if searchrescueservicetimeslicetype.ServiceProvider != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicetype.ServiceProvider.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceProvider = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.Call_sign {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Call_sign = append(%s.Call_sign, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.RadioCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunication = append(%s.RadioCommunication, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.GroundCommunication {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundCommunication = append(%s.GroundCommunication, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			if searchrescueservicetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.ClientAirspace {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientAirspace = append(%s.ClientAirspace, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.ClientRoute {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClientRoute = append(%s.ClientRoute, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", searchrescueservicetimeslicetypeIdent, searchrescueservicetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, searchrescueservicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SearchRescueServiceTypes, stage.AixmfeaturesStage.SearchRescueServiceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			searchrescueservicetypeIdent := searchrescueservicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SearchRescueServiceType{Name: %s}).Stage(stage.AixmfeaturesStage)", searchrescueservicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", searchrescueservicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", searchrescueservicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", searchrescueservicetypeIdent, __gong__toRawStringLiteral(searchrescueservicetype.Description)))
			if searchrescueservicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", searchrescueservicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", searchrescueservicetypeIdent, searchrescueservicetypeIdent, targetIdent))
			}
			if searchrescueservicetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", searchrescueservicetypeIdent, targetIdent))
			}
			if searchrescueservicetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := searchrescueservicetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", searchrescueservicetypeIdent, targetIdent))
			}
			for _, elem := range searchrescueservicetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", searchrescueservicetypeIdent, searchrescueservicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, serviceoperationalstatuspropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ServiceOperationalStatusPropertyTypes, stage.AixmfeaturesStage.ServiceOperationalStatusPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			serviceoperationalstatuspropertytypeIdent := serviceoperationalstatuspropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ServiceOperationalStatusPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", serviceoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(serviceoperationalstatuspropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", serviceoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(serviceoperationalstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", serviceoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(serviceoperationalstatuspropertytype.NilReason)))
			if serviceoperationalstatuspropertytype.ServiceOperationalStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := serviceoperationalstatuspropertytype.ServiceOperationalStatus.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServiceOperationalStatus = %s", serviceoperationalstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, serviceoperationalstatustype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ServiceOperationalStatusTypes, stage.AixmfeaturesStage.ServiceOperationalStatusType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			serviceoperationalstatustypeIdent := serviceoperationalstatustype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ServiceOperationalStatusType{Name: %s}).Stage(stage.AixmfeaturesStage)", serviceoperationalstatustypeIdent, __gong__toRawStringLiteral(serviceoperationalstatustype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", serviceoperationalstatustypeIdent, __gong__toRawStringLiteral(serviceoperationalstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", serviceoperationalstatustypeIdent, __gong__toRawStringLiteral(serviceoperationalstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", serviceoperationalstatustypeIdent, __gong__toRawStringLiteral(serviceoperationalstatustype.Description)))
			if serviceoperationalstatustype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := serviceoperationalstatustype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", serviceoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range serviceoperationalstatustype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", serviceoperationalstatustypeIdent, serviceoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range serviceoperationalstatustype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", serviceoperationalstatustypeIdent, serviceoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range serviceoperationalstatustype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", serviceoperationalstatustypeIdent, serviceoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range serviceoperationalstatustype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", serviceoperationalstatustypeIdent, serviceoperationalstatustypeIdent, targetIdent))
			}
			if serviceoperationalstatustype.OperationalStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := serviceoperationalstatustype.OperationalStatus.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OperationalStatus = %s", serviceoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range serviceoperationalstatustype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", serviceoperationalstatustypeIdent, serviceoperationalstatustypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialdatetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialDateTimeSlicePropertyTypes, stage.AixmfeaturesStage.SpecialDateTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialdatetimeslicepropertytypeIdent := specialdatetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialDateTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialdatetimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialdatetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialdatetimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialdatetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", specialdatetimeslicepropertytypeIdent, specialdatetimeslicepropertytype.Owns))
			if specialdatetimeslicepropertytype.SpecialDateTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetimeslicepropertytype.SpecialDateTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateTimeSlice = %s", specialdatetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialdatetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialDateTimeSliceTypes, stage.AixmfeaturesStage.SpecialDateTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialdatetimeslicetypeIdent := specialdatetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialDateTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialdatetimeslicetypeIdent, __gong__toRawStringLiteral(specialdatetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialdatetimeslicetypeIdent, __gong__toRawStringLiteral(specialdatetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialdatetimeslicetypeIdent, __gong__toRawStringLiteral(specialdatetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialdatetimeslicetypeIdent, __gong__toRawStringLiteral(specialdatetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", specialdatetimeslicetypeIdent, __gong__toRawStringLiteral(specialdatetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", specialdatetimeslicetypeIdent, specialdatetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", specialdatetimeslicetypeIdent, specialdatetimeslicetype.CorrectionNumber))
			if specialdatetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", specialdatetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialdatetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", specialdatetimeslicetypeIdent, specialdatetimeslicetypeIdent, targetIdent))
			}
			if specialdatetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", specialdatetimeslicetypeIdent, targetIdent))
			}
			if specialdatetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", specialdatetimeslicetypeIdent, targetIdent))
			}
			if specialdatetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", specialdatetimeslicetypeIdent, targetIdent))
			}
			if specialdatetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", specialdatetimeslicetypeIdent, targetIdent))
			}
			if specialdatetimeslicetype.DateDay != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetimeslicetype.DateDay.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateDay = %s", specialdatetimeslicetypeIdent, targetIdent))
			}
			if specialdatetimeslicetype.DateYear != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetimeslicetype.DateYear.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateYear = %s", specialdatetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialdatetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", specialdatetimeslicetypeIdent, specialdatetimeslicetypeIdent, targetIdent))
			}
			if specialdatetimeslicetype.Authority != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetimeslicetype.Authority.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = %s", specialdatetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialdatetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", specialdatetimeslicetypeIdent, specialdatetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialdatetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialDateTypes, stage.AixmfeaturesStage.SpecialDateType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialdatetypeIdent := specialdatetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialDateType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialdatetypeIdent, __gong__toRawStringLiteral(specialdatetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialdatetypeIdent, __gong__toRawStringLiteral(specialdatetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialdatetypeIdent, __gong__toRawStringLiteral(specialdatetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialdatetypeIdent, __gong__toRawStringLiteral(specialdatetype.Description)))
			if specialdatetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", specialdatetypeIdent, targetIdent))
			}
			for _, elem := range specialdatetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", specialdatetypeIdent, specialdatetypeIdent, targetIdent))
			}
			if specialdatetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", specialdatetypeIdent, targetIdent))
			}
			if specialdatetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialdatetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", specialdatetypeIdent, targetIdent))
			}
			for _, elem := range specialdatetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", specialdatetypeIdent, specialdatetypeIdent, targetIdent))
			}
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
		for _, standardlevelcolumntimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelColumnTimeSlicePropertyTypes, stage.AixmfeaturesStage.StandardLevelColumnTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardlevelcolumntimeslicepropertytypeIdent := standardlevelcolumntimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelColumnTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardlevelcolumntimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelcolumntimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardlevelcolumntimeslicepropertytypeIdent, standardlevelcolumntimeslicepropertytype.Owns))
			if standardlevelcolumntimeslicepropertytype.StandardLevelColumnTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntimeslicepropertytype.StandardLevelColumnTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevelColumnTimeSlice = %s", standardlevelcolumntimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardlevelcolumntimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes, stage.AixmfeaturesStage.StandardLevelColumnTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardlevelcolumntimeslicetypeIdent := standardlevelcolumntimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelColumnTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetype.CorrectionNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.UnitOfMeasurement = %s", standardlevelcolumntimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelcolumntimeslicetype.UnitOfMeasurement)))
			if standardlevelcolumntimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelcolumntimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			if standardlevelcolumntimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			if standardlevelcolumntimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			if standardlevelcolumntimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			if standardlevelcolumntimeslicetype.Series != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntimeslicetype.Series.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Series = %s", standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			if standardlevelcolumntimeslicetype.Separation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntimeslicetype.Separation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Separation = %s", standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelcolumntimeslicetype.Level {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = append(%s.Level, %s)", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			if standardlevelcolumntimeslicetype.LevelTable != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntimeslicetype.LevelTable.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LevelTable = %s", standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelcolumntimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelcolumntimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", standardlevelcolumntimeslicetypeIdent, standardlevelcolumntimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardlevelcolumntype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelColumnTypes, stage.AixmfeaturesStage.StandardLevelColumnType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardlevelcolumntypeIdent := standardlevelcolumntype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelColumnType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardlevelcolumntypeIdent, __gong__toRawStringLiteral(standardlevelcolumntype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelcolumntypeIdent, __gong__toRawStringLiteral(standardlevelcolumntype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardlevelcolumntypeIdent, __gong__toRawStringLiteral(standardlevelcolumntype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardlevelcolumntypeIdent, __gong__toRawStringLiteral(standardlevelcolumntype.Description)))
			if standardlevelcolumntype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", standardlevelcolumntypeIdent, targetIdent))
			}
			for _, elem := range standardlevelcolumntype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", standardlevelcolumntypeIdent, standardlevelcolumntypeIdent, targetIdent))
			}
			if standardlevelcolumntype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", standardlevelcolumntypeIdent, targetIdent))
			}
			if standardlevelcolumntype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelcolumntype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", standardlevelcolumntypeIdent, targetIdent))
			}
			for _, elem := range standardlevelcolumntype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", standardlevelcolumntypeIdent, standardlevelcolumntypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardlevelpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelPropertyTypes, stage.AixmfeaturesStage.StandardLevelPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardlevelpropertytypeIdent := standardlevelpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardlevelpropertytypeIdent, __gong__toRawStringLiteral(standardlevelpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelpropertytypeIdent, __gong__toRawStringLiteral(standardlevelpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", standardlevelpropertytypeIdent, __gong__toRawStringLiteral(standardlevelpropertytype.NilReason)))
			if standardlevelpropertytype.StandardLevel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelpropertytype.StandardLevel.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevel = %s", standardlevelpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardlevelsectortimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelSectorTimeSlicePropertyTypes, stage.AixmfeaturesStage.StandardLevelSectorTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardlevelsectortimeslicepropertytypeIdent := standardlevelsectortimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelSectorTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardlevelsectortimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelsectortimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardlevelsectortimeslicepropertytypeIdent, standardlevelsectortimeslicepropertytype.Owns))
			if standardlevelsectortimeslicepropertytype.StandardLevelSectorTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortimeslicepropertytype.StandardLevelSectorTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevelSectorTimeSlice = %s", standardlevelsectortimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardlevelsectortimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes, stage.AixmfeaturesStage.StandardLevelSectorTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardlevelsectortimeslicetypeIdent := standardlevelsectortimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelSectorTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardlevelsectortimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelsectortimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardlevelsectortimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardlevelsectortimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", standardlevelsectortimeslicetypeIdent, __gong__toRawStringLiteral(standardlevelsectortimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetype.CorrectionNumber))
			if standardlevelsectortimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelsectortimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			if standardlevelsectortimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			if standardlevelsectortimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			if standardlevelsectortimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			if standardlevelsectortimeslicetype.FlightRule != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortimeslicetype.FlightRule.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightRule = %s", standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			if standardlevelsectortimeslicetype.FromTrack != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortimeslicetype.FromTrack.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FromTrack = %s", standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			if standardlevelsectortimeslicetype.ToTrack != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortimeslicetype.ToTrack.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ToTrack = %s", standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			if standardlevelsectortimeslicetype.AngleType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortimeslicetype.AngleType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleType = %s", standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelsectortimeslicetype.ApplicableAirspace {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApplicableAirspace = append(%s.ApplicableAirspace, %s)", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			if standardlevelsectortimeslicetype.ApplicableLevelColumn != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortimeslicetype.ApplicableLevelColumn.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ApplicableLevelColumn = %s", standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelsectortimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardlevelsectortimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", standardlevelsectortimeslicetypeIdent, standardlevelsectortimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardlevelsectortype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelSectorTypes, stage.AixmfeaturesStage.StandardLevelSectorType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardlevelsectortypeIdent := standardlevelsectortype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelSectorType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardlevelsectortypeIdent, __gong__toRawStringLiteral(standardlevelsectortype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelsectortypeIdent, __gong__toRawStringLiteral(standardlevelsectortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardlevelsectortypeIdent, __gong__toRawStringLiteral(standardlevelsectortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardlevelsectortypeIdent, __gong__toRawStringLiteral(standardlevelsectortype.Description)))
			if standardlevelsectortype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", standardlevelsectortypeIdent, targetIdent))
			}
			for _, elem := range standardlevelsectortype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", standardlevelsectortypeIdent, standardlevelsectortypeIdent, targetIdent))
			}
			if standardlevelsectortype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", standardlevelsectortypeIdent, targetIdent))
			}
			if standardlevelsectortype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardlevelsectortype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", standardlevelsectortypeIdent, targetIdent))
			}
			for _, elem := range standardlevelsectortype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", standardlevelsectortypeIdent, standardlevelsectortypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardleveltablepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelTablePropertyTypes, stage.AixmfeaturesStage.StandardLevelTablePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardleveltablepropertytypeIdent := standardleveltablepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelTablePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardleveltablepropertytypeIdent, standardleveltablepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", standardleveltablepropertytypeIdent, __gong__toRawStringLiteral(standardleveltablepropertytype.Actuate)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardleveltabletimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelTableTimeSlicePropertyTypes, stage.AixmfeaturesStage.StandardLevelTableTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardleveltabletimeslicepropertytypeIdent := standardleveltabletimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelTableTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardleveltabletimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardleveltabletimeslicepropertytypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardleveltabletimeslicepropertytypeIdent, standardleveltabletimeslicepropertytype.Owns))
			if standardleveltabletimeslicepropertytype.StandardLevelTableTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltabletimeslicepropertytype.StandardLevelTableTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevelTableTimeSlice = %s", standardleveltabletimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardleveltabletimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelTableTimeSliceTypes, stage.AixmfeaturesStage.StandardLevelTableTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardleveltabletimeslicetypeIdent := standardleveltabletimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelTableTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardleveltabletimeslicetypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardleveltabletimeslicetypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardleveltabletimeslicetypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardleveltabletimeslicetypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", standardleveltabletimeslicetypeIdent, __gong__toRawStringLiteral(standardleveltabletimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", standardleveltabletimeslicetypeIdent, standardleveltabletimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", standardleveltabletimeslicetypeIdent, standardleveltabletimeslicetype.CorrectionNumber))
			if standardleveltabletimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltabletimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", standardleveltabletimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardleveltabletimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", standardleveltabletimeslicetypeIdent, standardleveltabletimeslicetypeIdent, targetIdent))
			}
			if standardleveltabletimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltabletimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", standardleveltabletimeslicetypeIdent, targetIdent))
			}
			if standardleveltabletimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltabletimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", standardleveltabletimeslicetypeIdent, targetIdent))
			}
			if standardleveltabletimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltabletimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", standardleveltabletimeslicetypeIdent, targetIdent))
			}
			if standardleveltabletimeslicetype.StandardICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltabletimeslicetype.StandardICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardICAO = %s", standardleveltabletimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardleveltabletimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", standardleveltabletimeslicetypeIdent, standardleveltabletimeslicetypeIdent, targetIdent))
			}
			for _, elem := range standardleveltabletimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", standardleveltabletimeslicetypeIdent, standardleveltabletimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardleveltabletype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelTableTypes, stage.AixmfeaturesStage.StandardLevelTableType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardleveltabletypeIdent := standardleveltabletype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelTableType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardleveltabletypeIdent, __gong__toRawStringLiteral(standardleveltabletype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardleveltabletypeIdent, __gong__toRawStringLiteral(standardleveltabletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardleveltabletypeIdent, __gong__toRawStringLiteral(standardleveltabletype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardleveltabletypeIdent, __gong__toRawStringLiteral(standardleveltabletype.Description)))
			if standardleveltabletype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltabletype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", standardleveltabletypeIdent, targetIdent))
			}
			for _, elem := range standardleveltabletype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", standardleveltabletypeIdent, standardleveltabletypeIdent, targetIdent))
			}
			if standardleveltabletype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltabletype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", standardleveltabletypeIdent, targetIdent))
			}
			if standardleveltabletype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltabletype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", standardleveltabletypeIdent, targetIdent))
			}
			for _, elem := range standardleveltabletype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", standardleveltabletypeIdent, standardleveltabletypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, standardleveltype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.StandardLevelTypes, stage.AixmfeaturesStage.StandardLevelType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			standardleveltypeIdent := standardleveltype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelType{Name: %s}).Stage(stage.AixmfeaturesStage)", standardleveltypeIdent, __gong__toRawStringLiteral(standardleveltype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardleveltypeIdent, __gong__toRawStringLiteral(standardleveltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", standardleveltypeIdent, __gong__toRawStringLiteral(standardleveltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", standardleveltypeIdent, __gong__toRawStringLiteral(standardleveltype.Description)))
			if standardleveltype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", standardleveltypeIdent, targetIdent))
			}
			for _, elem := range standardleveltype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", standardleveltypeIdent, standardleveltypeIdent, targetIdent))
			}
			if standardleveltype.VerticalDistance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := standardleveltype.VerticalDistance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalDistance = %s", standardleveltypeIdent, targetIdent))
			}
			for _, elem := range standardleveltype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", standardleveltypeIdent, standardleveltypeIdent, targetIdent))
			}
			for _, elem := range standardleveltype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", standardleveltypeIdent, standardleveltypeIdent, targetIdent))
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
		for _, unitavailabilitypropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.UnitAvailabilityPropertyTypes, stage.AixmfeaturesStage.UnitAvailabilityPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			unitavailabilitypropertytypeIdent := unitavailabilitypropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitAvailabilityPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", unitavailabilitypropertytypeIdent, __gong__toRawStringLiteral(unitavailabilitypropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unitavailabilitypropertytypeIdent, __gong__toRawStringLiteral(unitavailabilitypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", unitavailabilitypropertytypeIdent, __gong__toRawStringLiteral(unitavailabilitypropertytype.NilReason)))
			if unitavailabilitypropertytype.UnitAvailability != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unitavailabilitypropertytype.UnitAvailability.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UnitAvailability = %s", unitavailabilitypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, unitavailabilitytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.UnitAvailabilityTypes, stage.AixmfeaturesStage.UnitAvailabilityType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			unitavailabilitytypeIdent := unitavailabilitytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitAvailabilityType{Name: %s}).Stage(stage.AixmfeaturesStage)", unitavailabilitytypeIdent, __gong__toRawStringLiteral(unitavailabilitytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unitavailabilitytypeIdent, __gong__toRawStringLiteral(unitavailabilitytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", unitavailabilitytypeIdent, __gong__toRawStringLiteral(unitavailabilitytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", unitavailabilitytypeIdent, __gong__toRawStringLiteral(unitavailabilitytype.Description)))
			if unitavailabilitytype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unitavailabilitytype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", unitavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range unitavailabilitytype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", unitavailabilitytypeIdent, unitavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range unitavailabilitytype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", unitavailabilitytypeIdent, unitavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range unitavailabilitytype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", unitavailabilitytypeIdent, unitavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range unitavailabilitytype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", unitavailabilitytypeIdent, unitavailabilitytypeIdent, targetIdent))
			}
			if unitavailabilitytype.OperationalStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unitavailabilitytype.OperationalStatus.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OperationalStatus = %s", unitavailabilitytypeIdent, targetIdent))
			}
			for _, elem := range unitavailabilitytype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", unitavailabilitytypeIdent, unitavailabilitytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, unitdependencypropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.UnitDependencyPropertyTypes, stage.AixmfeaturesStage.UnitDependencyPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			unitdependencypropertytypeIdent := unitdependencypropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitDependencyPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", unitdependencypropertytypeIdent, __gong__toRawStringLiteral(unitdependencypropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unitdependencypropertytypeIdent, __gong__toRawStringLiteral(unitdependencypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", unitdependencypropertytypeIdent, __gong__toRawStringLiteral(unitdependencypropertytype.NilReason)))
			if unitdependencypropertytype.UnitDependency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unitdependencypropertytype.UnitDependency.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UnitDependency = %s", unitdependencypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, unitdependencytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.UnitDependencyTypes, stage.AixmfeaturesStage.UnitDependencyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			unitdependencytypeIdent := unitdependencytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitDependencyType{Name: %s}).Stage(stage.AixmfeaturesStage)", unitdependencytypeIdent, __gong__toRawStringLiteral(unitdependencytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unitdependencytypeIdent, __gong__toRawStringLiteral(unitdependencytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", unitdependencytypeIdent, __gong__toRawStringLiteral(unitdependencytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", unitdependencytypeIdent, __gong__toRawStringLiteral(unitdependencytype.Description)))
			if unitdependencytype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unitdependencytype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", unitdependencytypeIdent, targetIdent))
			}
			for _, elem := range unitdependencytype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", unitdependencytypeIdent, unitdependencytypeIdent, targetIdent))
			}
			if unitdependencytype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unitdependencytype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", unitdependencytypeIdent, targetIdent))
			}
			for _, elem := range unitdependencytype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", unitdependencytypeIdent, unitdependencytypeIdent, targetIdent))
			}
			if unitdependencytype.TheUnit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unitdependencytype.TheUnit.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheUnit = %s", unitdependencytypeIdent, targetIdent))
			}
			for _, elem := range unitdependencytype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", unitdependencytypeIdent, unitdependencytypeIdent, targetIdent))
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
		for _, unittimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.UnitTimeSlicePropertyTypes, stage.AixmfeaturesStage.UnitTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			unittimeslicepropertytypeIdent := unittimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", unittimeslicepropertytypeIdent, __gong__toRawStringLiteral(unittimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unittimeslicepropertytypeIdent, __gong__toRawStringLiteral(unittimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", unittimeslicepropertytypeIdent, unittimeslicepropertytype.Owns))
			if unittimeslicepropertytype.UnitTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicepropertytype.UnitTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UnitTimeSlice = %s", unittimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, unittimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.UnitTimeSliceTypes, stage.AixmfeaturesStage.UnitTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			unittimeslicetypeIdent := unittimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", unittimeslicetypeIdent, __gong__toRawStringLiteral(unittimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unittimeslicetypeIdent, __gong__toRawStringLiteral(unittimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", unittimeslicetypeIdent, __gong__toRawStringLiteral(unittimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", unittimeslicetypeIdent, __gong__toRawStringLiteral(unittimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", unittimeslicetypeIdent, __gong__toRawStringLiteral(unittimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", unittimeslicetypeIdent, unittimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", unittimeslicetypeIdent, unittimeslicetype.CorrectionNumber))
			if unittimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.CompliantICAO != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.CompliantICAO.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompliantICAO = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.Military != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.Military.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Military = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.Position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.Position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Position = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.AirportLocation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.AirportLocation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportLocation = %s", unittimeslicetypeIdent, targetIdent))
			}
			if unittimeslicetype.OwnerOrganisation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittimeslicetype.OwnerOrganisation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OwnerOrganisation = %s", unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.Contact {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.RelatedUnit {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RelatedUnit = append(%s.RelatedUnit, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
			for _, elem := range unittimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", unittimeslicetypeIdent, unittimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, unittype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.UnitTypes, stage.AixmfeaturesStage.UnitType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			unittypeIdent := unittype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitType{Name: %s}).Stage(stage.AixmfeaturesStage)", unittypeIdent, __gong__toRawStringLiteral(unittype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unittypeIdent, __gong__toRawStringLiteral(unittype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", unittypeIdent, __gong__toRawStringLiteral(unittype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", unittypeIdent, __gong__toRawStringLiteral(unittype.Description)))
			if unittype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", unittypeIdent, targetIdent))
			}
			for _, elem := range unittype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", unittypeIdent, unittypeIdent, targetIdent))
			}
			if unittype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", unittypeIdent, targetIdent))
			}
			if unittype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := unittype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", unittypeIdent, targetIdent))
			}
			for _, elem := range unittype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", unittypeIdent, unittypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, vorpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VORPropertyTypes, stage.AixmfeaturesStage.VORPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			vorpropertytypeIdent := vorpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VORPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", vorpropertytypeIdent, __gong__toRawStringLiteral(vorpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
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
			if basicmessagememberaixmpropertytype.OrganisationAuthority != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.OrganisationAuthority.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OrganisationAuthority = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.Unit != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.Unit.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Unit = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.AirTrafficControlService != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.AirTrafficControlService.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirTrafficControlService = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.AirTrafficManagementService != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.AirTrafficManagementService.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirTrafficManagementService = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.GroundTrafficControlService != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.GroundTrafficControlService.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundTrafficControlService = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.InformationService != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.InformationService.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InformationService = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.SearchRescueService != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.SearchRescueService.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SearchRescueService = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.AircraftGroundService != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.AircraftGroundService.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftGroundService = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.AirportClearanceService != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.AirportClearanceService.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportClearanceService = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.AirportSuppliesService != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.AirportSuppliesService.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportSuppliesService = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.FireFightingService != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.FireFightingService.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FireFightingService = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.PassengerService != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.PassengerService.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PassengerService = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.PilotControlledLighting != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.PilotControlledLighting.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PilotControlledLighting = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.RadioCommunicationChannel != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.RadioCommunicationChannel.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadioCommunicationChannel = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.SpecialDate != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.SpecialDate.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDate = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.StandardLevelTable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.StandardLevelTable.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevelTable = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.StandardLevelColumn != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.StandardLevelColumn.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevelColumn = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.StandardLevelSector != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.StandardLevelSector.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevelSector = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.StandardLevel != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.StandardLevel.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandardLevel = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/xlink"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models\"", fmt.Sprintf("%q", modelsPackageName))
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
